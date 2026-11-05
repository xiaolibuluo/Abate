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

var litExecFn = func(args []string, input string) ([]byte, error) {
	cmd := exec.Command("kubectl", args...)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	return cmd.CombinedOutput()
}

var litApplyFn = func(yaml string) (string, error) {
	out, err := litExecFn([]string{"apply", "-f", "-"}, yaml)
	if err != nil {
		return "", fmt.Errorf("kubectl apply: %w: %s", err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

var litDeleteFn = func(name string) error {
	out, err := litExecFn([]string{"delete", "chaosengine", name}, "")
	if err != nil {
		// CES-3.1.2：Cleanup 必须可重复调用；资源已不存在视为成功。
		if strings.Contains(strings.ToLower(string(out)), "not found") {
			return nil
		}
		return fmt.Errorf("kubectl delete: %w: %s", err, string(out))
	}
	return nil
}

var litCounter int64

// LitmusBackend 是 LitmusChaos 平台的 Backend 实现（CES-7.2.3）。
// LitmusBackend is the Backend implementation for LitmusChaos.
type LitmusBackend struct {
	namespace   string
	appLabel    string
	appKind     string
}

// NewLitmusBackend 构造 Litmus 后端。
func NewLitmusBackend(namespace, appLabel, appKind string) *LitmusBackend {
	return &LitmusBackend{namespace: namespace, appLabel: appLabel, appKind: appKind}
}

// Apply 生成 ChaosEngine YAML 并通过 kubectl apply 创建故障。
func (l *LitmusBackend) Apply(mapped MappedFault) (string, error) {
	idx := atomic.AddInt64(&litCounter, 1)
	resName := fmt.Sprintf("abate-litmus-%d", idx)

	yaml := l.generateChaosEngine(resName, mapped)
	if _, err := litApplyFn(yaml); err != nil {
		return "", fmt.Errorf("litmus: apply: %w", err)
	}
	return resName, nil
}

// Cleanup 通过 kubectl delete chaosengine 清除故障。
func (l *LitmusBackend) Cleanup(upstreamID string) error {
	if err := litDeleteFn(upstreamID); err != nil {
		return fmt.Errorf("litmus: cleanup: %w", err)
	}
	return nil
}

func (l *LitmusBackend) generateChaosEngine(name string, mapped MappedFault) string {
	var sb strings.Builder
	sb.WriteString("apiVersion: litmuschaos.io/v1alpha1\n")
	sb.WriteString("kind: ChaosEngine\n")
	sb.WriteString(fmt.Sprintf("metadata:\n  name: %s\n  namespace: %s\n", name, l.namespace))
	sb.WriteString("spec:\n")
	sb.WriteString(fmt.Sprintf("  appinfo:\n    appns: %s\n    applabel: %s\n    appkind: %s\n", l.namespace, l.appLabel, l.appKind))
	sb.WriteString(fmt.Sprintf("  chaosServiceAccount: chaos-operator\n"))
	sb.WriteString(fmt.Sprintf("  experiments:\n  - name: %s\n", mapped.UpstreamType))
	sb.WriteString("    spec:\n      components:\n        env:\n")
	for k, v := range mapped.Params {
		sb.WriteString(fmt.Sprintf("        - name: %s\n          value: %s\n", k, v))
	}
	return sb.String()
}
