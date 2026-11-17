/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"

	"abate/kernel/algebra"
	"abate/sim/clock"
	"abate/sim/record"
)

// Injector executes C-family faults and their dual compensations on one World.
//
// 参考 chaosblade 的 Inject/Rollback 执行器，改进点=纯库形态、全虚拟故障原语、
// 对偶补偿统一经 Inverse + 基线快照，且事件流带确定性指纹、可并发调用。
// A mutex serializes every mutation of the virtual world, so concurrent Inject
// and Rollback calls are race-free without leaking global mutable state.
type Injector struct {
	mu       sync.Mutex
	world    *World
	registry *Registry
	frames   []record.Frame
	lastAt   clock.Instant
	step     uint64
}

// NewInjector binds an injector to one virtual world.
func NewInjector(world *World) (*Injector, error) {
	if world == nil {
		return nil, ErrNilWorld
	}
	return &Injector{world: world, registry: NewRegistry()}, nil
}

// World returns the virtual world driven by the injector.
func (i *Injector) World() *World { return i.world }

// Registry returns the read-only twelve-operator registry.
func (i *Injector) Registry() *Registry { return i.registry }

// Frames returns a defensive copy of the recorded fault event stream.
func (i *Injector) Frames() []record.Frame {
	i.mu.Lock()
	defer i.mu.Unlock()
	out := make([]record.Frame, 0, len(i.frames))
	for _, frame := range i.frames {
		copied := frame
		copied.Payload = append([]byte(nil), frame.Payload...)
		out = append(out, copied)
	}
	return out
}

// Digest returns the SHA-256 digest of the canonical injection event stream.
func (i *Injector) Digest() [32]byte {
	i.mu.Lock()
	defer i.mu.Unlock()
	var b bytes.Buffer
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[:], i.step)
	_, _ = b.Write(tmp[:])
	for _, frame := range i.frames {
		_, _ = b.WriteString(strconv.FormatInt(int64(frame.At), 10))
		_, _ = b.Write(frame.Kind.Marshal())
		_, _ = b.WriteString(string(frame.Kind))
		_, _ = b.Write(frame.Payload)
		_ = b.WriteByte(0)
	}
	return sha256.Sum256(b.Bytes())
}

// Inject applies one registered C-family fault to the virtual world.
func (i *Injector) Inject(ctx context.Context, op algebra.Operator) error {
	if err := i.validate(ctx, op); err != nil {
		return err
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	if err := dispatchApply(i.world, op); err != nil {
		return err
	}
	i.recordFrames(op, "inject")
	return nil
}

// Rollback applies the operator dual and restores the baseline snapshot.
//
// 参考 chaosblade 的故障清理，改进点=强制经过 Inverse() 取对偶算子，再执行
// SIM 对偶调用，最后用基线快照兜底，回滚后 Snapshot 必须与 Baseline 逐字节相等。
func (i *Injector) Rollback(ctx context.Context, op algebra.Operator) error {
	if err := i.validate(ctx, op); err != nil {
		return err
	}
	dual, err := op.Inverse()
	if err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	if !strings.HasSuffix(string(dual.ID()), "~inv") {
		return ErrNotDual
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	if err := dispatchRecover(i.world, dual); err != nil {
		return err
	}
	i.world.RestoreBaseline()
	if i.world.Snapshot() != i.world.Baseline() {
		return ErrRollbackIncomplete
	}
	i.recordFrames(op, "rollback")
	return nil
}

// validate checks the context, family, registration, invertibility and target.
func (i *Injector) validate(ctx context.Context, op algebra.Operator) error {
	if ctx == nil {
		return ErrCanceled
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	if op == nil {
		return ErrNilOperator
	}
	if op.Family() != algebra.FamilyChaos {
		return ErrWrongFamily
	}
	if !i.registry.Contains(op.ID()) {
		return ErrUnknownOperator
	}
	if !op.Invertible() {
		return ErrNotInvertible
	}
	if _, ok := i.registry.specByClass(op.Params()["class"]); !ok {
		return ErrUnknownOperator
	}
	if !slices.Contains(i.world.KnownTargets(), op.Params()["target"]) {
		return ErrUnknownTarget
	}
	return nil
}

// recordFrames appends the four real SIM-derived frames of one inject/rollback.
func (i *Injector) recordFrames(op algebra.Operator, phase string) {
	at := i.world.clk.Now()
	if at < i.lastAt {
		at = i.lastAt
	}
	i.lastAt = at
	i.step++
	fingerprint := op.Fingerprint()
	fault := make([]byte, 0, len(phase)+1+len(op.ID())+len(fingerprint))
	fault = append(fault, phase...)
	fault = append(fault, ':')
	fault = append(fault, string(op.ID())...)
	fault = append(fault, fingerprint[:]...)
	i.frames = append(i.frames,
		record.Frame{At: at, Kind: record.KindFault, Payload: fault},
		record.Frame{At: at, Kind: record.KindNet, Payload: []byte(i.world.netState())},
		record.Frame{At: at, Kind: record.KindFS, Payload: []byte(i.world.fsState())},
		record.Frame{At: at, Kind: record.KindVTask, Payload: []byte(i.world.taskState())},
	)
}

// dispatchApply routes one fault class to its virtual primitive implementation.
func dispatchApply(world *World, op algebra.Operator) error {
	switch op.Params()["class"] {
	case "network-latency", "network-partition", "network-loss", "network-bandwidth":
		return applyNetwork(world, op)
	case "disk-io-latency", "disk-full", "cpu-pressure", "memory-pressure":
		return applyResource(world, op)
	case "clock-offset":
		return applyClock(world, op)
	case "process-kill":
		return applyProcess(world, op)
	case "dependency-error", "dependency-timeout":
		return applyDependency(world, op)
	default:
		return ErrUnknownOperator
	}
}

// dispatchRecover routes the dual operator to its compensation implementation.
func dispatchRecover(world *World, dual algebra.Operator) error {
	switch dual.Params()["class"] {
	case "network-latency", "network-partition", "network-loss", "network-bandwidth":
		return recoverNetwork(world, dual)
	case "disk-io-latency", "disk-full", "cpu-pressure", "memory-pressure":
		return recoverResource(world, dual)
	case "clock-offset":
		return recoverClock(world, dual)
	case "process-kill":
		return recoverProcess(world, dual)
	case "dependency-error", "dependency-timeout":
		return recoverDependency(world, dual)
	default:
		return ErrUnknownOperator
	}
}

// faultParams is the parsed canonical parameter set of one fault operator.
type faultParams struct {
	class  string
	target string
	value  int64
}

// paramsOf parses the canonical class, target and integer value of one operator.
func paramsOf(op algebra.Operator) (faultParams, error) {
	params := op.Params()
	value, err := strconv.ParseInt(params["value"], 10, 64)
	if err != nil {
		return faultParams{}, ErrInvalidParams
	}
	return faultParams{class: params["class"], target: params["target"], value: value}, nil
}
