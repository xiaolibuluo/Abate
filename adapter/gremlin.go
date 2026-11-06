// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package adapter

import (
	"fmt"
	"os/exec"
	"strings"
	"sync/atomic"
)

var gremlinExecFn = func(args []string) ([]byte, error) {
	return exec.Command("gremlin", args...).CombinedOutput()
}

var gremlinCounter int64

// GremlinBackend 是 Gremlin 平台的 Backend 实现（CES-7.2.3）。
// GremlinBackend is the Backend implementation for Gremlin.
type GremlinBackend struct {
	teamID string
}

// NewGremlinBackend 构造 Gremlin 后端。
func NewGremlinBackend(teamID string) *GremlinBackend {
	return &GremlinBackend{teamID: teamID}
}

// gremlinShortFlags 将已确认的 Gremlin 上游参数名映射为真实 CLI 短参数
// （官方文档 latency 与 cpu 实验）。未列入的参数保持长参数传递：Gremlin CLI
// 长短参数等价，其余实验的短参数待官方文档逐一确认后补充。
// gremlinShortFlags maps confirmed upstream param names to real CLI short
// flags (official docs, latency & cpu experiments). Params not listed fall
// back to long flags, which the Gremlin CLI accepts equivalently.
var gremlinShortFlags = map[string]string{
	"length":     "-l", // 实验持续时长（秒）/ duration in seconds
	"ms":         "-m", // 网络延迟（毫秒）/ added latency in milliseconds
	"jitter":     "-j", // 抖动（毫秒）/ jitter in milliseconds
	"ip":         "-i", // 目标 IP/CIDR / target IP or CIDR
	"device":     "-d", // 网卡设备 / network device
	"hostname":   "-h", // 目标域名 / target hostname
	"port":       "-p", // 目的端口 / destination port
	"sourcePort": "-s", // 源端口 / source port
	"protocol":   "-P", // 协议 TCP/UDP/ICMP / protocol
	"cores":      "-c", // CPU 核数 / CPU cores
}

// Apply 通过 gremlin CLI 创建攻击。攻击 GUID 由 Gremlin CLI 自动生成，
// 不作为参数传入；返回值为本地生成的追踪句柄。
// Apply creates an attack via the gremlin CLI. The attack GUID is generated
// by the CLI itself and never passed as a flag; the return value is a
// locally generated tracking handle.
func (g *GremlinBackend) Apply(mapped MappedFault) (string, error) {
	idx := atomic.AddInt64(&gremlinCounter, 1)
	guid := fmt.Sprintf("gremlin-%d", idx)

	args := buildGremlinArgs(mapped)
	out, err := gremlinExecFn(args)
	if err != nil {
		return "", fmt.Errorf("gremlin: apply: %w: %s", err, string(out))
	}
	return guid, nil
}

// Cleanup 通过 gremlin rollback 回滚攻击。真实 CLI 的回滚命令为
// gremlin rollback，不接受攻击 ID；upstreamID 仅用于错误日志追踪。
// Cleanup rolls back attacks via gremlin rollback. The real CLI rollback
// command takes no attack ID; upstreamID is kept for error logging only.
func (g *GremlinBackend) Cleanup(upstreamID string) error {
	out, err := gremlinExecFn([]string{"rollback"})
	if err != nil {
		return fmt.Errorf("gremlin: cleanup %s: %w: %s", upstreamID, err, string(out))
	}
	return nil
}

// buildGremlinArgs 构造 gremlin attack 命令行参数。flag 与值必须拆分为
// 独立的数组元素（如 ["-l","60"]），不得合并为单个字符串。
// buildGremlinArgs builds gremlin attack CLI args. Each flag and its value
// are separate array elements (e.g. ["-l","60"]), never merged into one.
func buildGremlinArgs(mapped MappedFault) []string {
	parts := strings.SplitN(mapped.UpstreamType, "/", 2)
	method := mapped.UpstreamType
	if len(parts) == 2 {
		method = parts[1]
	}

	args := []string{"attack", method}
	for k, v := range mapped.Params {
		flag := "--" + k
		if short, ok := gremlinShortFlags[k]; ok {
			flag = short
		}
		args = append(args, flag, v)
	}
	return args
}
