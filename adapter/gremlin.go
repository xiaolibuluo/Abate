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

// Apply 通过 gremlin CLI 创建攻击。
func (g *GremlinBackend) Apply(mapped MappedFault) (string, error) {
	idx := atomic.AddInt64(&gremlinCounter, 1)
	guid := fmt.Sprintf("gremlin-%d", idx)

	args := buildGremlinArgs(guid, mapped)
	out, err := gremlinExecFn(args)
	if err != nil {
		return "", fmt.Errorf("gremlin: apply: %w: %s", err, string(out))
	}
	return guid, nil
}

// Cleanup 通过 gremlin CLI 停止攻击。
func (g *GremlinBackend) Cleanup(upstreamID string) error {
	out, err := gremlinExecFn([]string{"attack", "stop", upstreamID})
	if err != nil {
		return fmt.Errorf("gremlin: cleanup: %w: %s", err, string(out))
	}
	return nil
}

func buildGremlinArgs(guid string, mapped MappedFault) []string {
	parts := strings.SplitN(mapped.UpstreamType, "/", 2)
	method := mapped.UpstreamType
	if len(parts) == 2 {
		method = parts[1]
	}

	args := []string{"attack", method, "--guid", guid}
	for k, v := range mapped.Params {
		args = append(args, "--"+k, v)
	}
	return args
}
