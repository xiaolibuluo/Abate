/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"fmt"
	"math"
	"strings"

	"abate/sim/clock"
	"abate/sim/fs"
	"abate/sim/net"
	"abate/sim/vsched"
)

// Fixed identifiers of the virtual fault substrate.
const (
	linkA         = "n1"
	linkB         = "n2"
	filePath      = "/data/blob"
	cgroupName    = "cg"
	baseShares    = int64(1024)
	baseMemLimit  = int64(1 << 20)
	baseLatency   = clock.Duration(1000)
	baseBandwidth = uint64(1_000_000)
	probeBytes    = int64(16)
)

// World is the fully virtual substrate on which every C-family fault acts.
//
// 参考 chaosblade 对真实主机的故障作用面，改进点=所有作用面改为 SIM 内存模型
// （虚拟链路、内存 FS、虚拟任务、虚拟时钟），加上纯内存的依赖故障与配额账本。
// Snapshot renders every observable field in a fixed order with sorted keys,
// so the baseline and the post-rollback state are comparable byte for byte.
type World struct {
	clk        *clock.VirtualClock
	network    *net.Network
	filesystem *fs.FileSystem
	scheduler  *vsched.Scheduler
	link       *net.Link
	tasks      []string
	quotaBytes int64
	offsetNS   clock.Duration
	depTimeout bool
	depError   bool
	baseline   string
}

// NewWorld builds a seeded virtual world and records its baseline snapshot.
func NewWorld(seed uint64) (*World, error) {
	// Every fixed identifier and capacity below is valid by construction, so
	// the infallible setup calls are explicitly acknowledged rather than
	// wrapped in unreachable error branches.
	clk := clock.NewVirtualClock()
	network := net.NewNetworkWithSeed(clk, seed)
	link := network.Link(linkA, linkB, net.LinkConfig{
		BaseLatency: baseLatency, BandwidthBps: baseBandwidth,
	})
	filesystem := fs.NewFileSystem(clk)
	_ = filesystem.MkdirAll("/data")
	handle, err := filesystem.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("injection: create: %w", err)
	}
	_, _ = handle.Write(make([]byte, probeBytes))
	_ = handle.Close()
	tasks := []string{"task-1", "task-2"}
	scheduler := vsched.NewScheduler(clk)
	_, _ = scheduler.NewCgroup(cgroupName, baseShares, baseMemLimit)
	for _, id := range tasks {
		_, _ = scheduler.NewTask(id, cgroupName)
		_ = scheduler.Enqueue(id)
	}
	world := &World{
		clk: clk, network: network, filesystem: filesystem, scheduler: scheduler,
		link: link, tasks: tasks, quotaBytes: fs.MaxTotalBytes,
	}
	world.baseline = world.Snapshot()
	return world, nil
}

// Clock returns the single virtual clock of the world.
func (w *World) Clock() *clock.VirtualClock { return w.clk }

// Network returns the virtual network of the world.
func (w *World) Network() *net.Network { return w.network }

// FileSystem returns the in-memory virtual file system of the world.
func (w *World) FileSystem() *fs.FileSystem { return w.filesystem }

// Scheduler returns the virtual task scheduler of the world.
func (w *World) Scheduler() *vsched.Scheduler { return w.scheduler }

// Baseline returns the snapshot captured before any fault was injected.
func (w *World) Baseline() string { return w.baseline }

// KnownTargets returns the fault targets of the world in fixed order.
func (w *World) KnownTargets() []string {
	return []string{filePath, cgroupName, "clock", w.link.A() + "-" + w.link.B(), "svc", "task-1"}
}

// Dependency resolves one virtual dependency call under the active faults.
func (w *World) Dependency(name string) error {
	if name == "" {
		return ErrUnknownTarget
	}
	if w.depTimeout {
		return ErrDependencyTimeout
	}
	if w.depError {
		return ErrDependencyFailure
	}
	return nil
}

// Snapshot renders every observable SIM field in a deterministic order.
func (w *World) Snapshot() string {
	var b strings.Builder
	write := func(s string) { _, _ = b.WriteString(s) }
	write(w.netState())
	write(w.fsState())
	write(w.taskState())
	write(fmt.Sprintf("clk/%s/%d;", w.clk.State().String(), int64(w.offsetNS)))
	write(fmt.Sprintf("dep/%t/%t;", w.depTimeout, w.depError))
	return b.String()
}

// RestoreBaseline applies the baseline snapshot to every SIM component.
//
// 参考 chaosblade 的 rollback 清理脚本，改进点=纯内存恢复（链路 Recover、
// 坏块清除、配额复原、时钟偏移对偶、任务调度器按基线重建），无信号或 CLI。
func (w *World) RestoreBaseline() {
	w.link.Recover()
	_ = w.filesystem.RecoverFault(filePath)
	_ = w.filesystem.InjectFault(fs.Fault{Kind: fs.FaultLatency, Value: int64(fs.DefaultIOLatency)})
	w.quotaBytes = fs.MaxTotalBytes
	if w.offsetNS != 0 {
		w.clk.Offset(-w.offsetNS)
		w.offsetNS = 0
	}
	w.depTimeout = false
	w.depError = false
	w.rebuildScheduler()
}

// netState reports the real virtual-link configuration and partition flag.
func (w *World) netState() string {
	cfg := w.link.Config()
	return fmt.Sprintf("net/%d/%d/%d/%t;", int64(cfg.BaseLatency), cfg.BandwidthBps,
		math.Float64bits(cfg.LossRate), w.link.Partitioned())
}

// fsState reports the measured IO latency, quota, write and read outcomes.
func (w *World) fsState() string {
	return fmt.Sprintf("fs/%d/%d/%s/%s;",
		w.ioProbeNS(), w.quotaBytes, w.writeProbe(), w.readProbe())
}

// taskState reports the real virtual task and cgroup state in sorted id order.
func (w *World) taskState() string {
	var b strings.Builder
	write := func(s string) { _, _ = b.WriteString(s) }
	for _, id := range w.scheduler.TaskIDs() {
		task, _ := w.scheduler.Task(id)
		write(fmt.Sprintf("%s/%d/%t/%d;", id, task.State(), task.OOM(), task.MemBytes()))
	}
	cgroup, _ := w.scheduler.Cgroup(cgroupName)
	write(fmt.Sprintf("cg/%d/%d;", cgroup.Shares(), cgroup.MemUsed()))
	return b.String()
}

// ioProbeNS measures the virtual IO latency charged by one real fs.Open call.
func (w *World) ioProbeNS() int64 {
	before := w.clk.Now().UnixNano()
	handle, err := w.filesystem.Open(filePath)
	if err != nil {
		return -1
	}
	_ = handle.Close()
	return w.clk.Now().UnixNano() - before
}

// writeProbe reports the outcome of one real virtual file write.
func (w *World) writeProbe() string {
	if w.quotaBytes <= 0 {
		return "quota"
	}
	handle, err := w.filesystem.Open(filePath)
	if err != nil {
		return "open-error"
	}
	_, err = handle.Write([]byte{0x5a})
	_ = handle.Close()
	if err != nil {
		return "write:" + err.Error()
	}
	return "ok"
}

// readProbe reports the outcome of one real virtual file read.
func (w *World) readProbe() string {
	handle, err := w.filesystem.Open(filePath)
	if err != nil {
		return "open-error"
	}
	buf := make([]byte, 8)
	_, err = handle.Read(buf)
	_ = handle.Close()
	if err != nil {
		return "read:" + err.Error()
	}
	return "ok"
}

// rebuildScheduler constructs a fresh scheduler from the baseline task set.
func (w *World) rebuildScheduler() {
	scheduler := vsched.NewScheduler(w.clk)
	_, _ = scheduler.NewCgroup(cgroupName, baseShares, baseMemLimit)
	for _, id := range w.tasks {
		_, _ = scheduler.NewTask(id, cgroupName)
		_ = scheduler.Enqueue(id)
	}
	w.scheduler = scheduler
}
