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

var cmExecFn = func(args []string, input string) ([]byte, error) {
	cmd := exec.Command("kubectl", args...)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	return cmd.CombinedOutput()
}

var cmApplyFn = func(yaml string) (string, error) {
	out, err := cmExecFn([]string{"apply", "-f", "-"}, yaml)
	if err != nil {
		return "", fmt.Errorf("kubectl apply: %w: %s", err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

var cmDeleteFn = func(name string) error {
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid resource name %q", name)
	}
	out, err := cmExecFn([]string{"delete", parts[0], parts[1]}, "")
	if err != nil {
		// CES-3.1.2：Cleanup 必须可重复调用；资源已不存在视为成功。
		if strings.Contains(strings.ToLower(string(out)), "not found") {
			return nil
		}
		return fmt.Errorf("kubectl delete: %w: %s", err, string(out))
	}
	return nil
}

var cmCounter int64

// ChaosMeshBackend 是 ChaosMesh 平台的 Backend 实现（CES-7.2.3）。
// ChaosMeshBackend is the Backend implementation for ChaosMesh.
type ChaosMeshBackend struct {
	namespace string
}

// NewChaosMeshBackend 构造 ChaosMesh 后端，namespace 指定 CRD 命名空间。
func NewChaosMeshBackend(namespace string) *ChaosMeshBackend {
	return &ChaosMeshBackend{namespace: namespace}
}

// Apply 生成 CRD YAML 并通过 kubectl apply 创建故障。
func (c *ChaosMeshBackend) Apply(mapped MappedFault) (string, error) {
	idx := atomic.AddInt64(&cmCounter, 1)
	resType := strings.SplitN(mapped.UpstreamType, "/", 2)
	if len(resType) != 2 {
		return "", fmt.Errorf("chaosmesh: invalid upstream type %q", mapped.UpstreamType)
	}
	crdKind := resType[0]
	action := resType[1]
	resName := fmt.Sprintf("abate-%d", idx)

	yaml := c.generateCRD(crdKind, action, resName, mapped)
	if _, err := cmApplyFn(yaml); err != nil {
		return "", fmt.Errorf("chaosmesh: apply: %w", err)
	}

	resourceID := fmt.Sprintf("%s/%s", crdKindLower(crdKind), resName)
	return resourceID, nil
}

// Cleanup 通过 kubectl delete 清除故障。
func (c *ChaosMeshBackend) Cleanup(upstreamID string) error {
	if err := cmDeleteFn(upstreamID); err != nil {
		return fmt.Errorf("chaosmesh: cleanup: %w", err)
	}
	return nil
}

func (c *ChaosMeshBackend) generateCRD(kind, action, name string, mapped MappedFault) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("apiVersion: chaos-mesh.org/v1alpha1\n"))
	sb.WriteString(fmt.Sprintf("kind: %s\n", kind))
	sb.WriteString(fmt.Sprintf("metadata:\n  name: %s\n  namespace: %s\n", name, c.namespace))
	sb.WriteString("spec:\n")
	sb.WriteString(fmt.Sprintf("  action: %s\n", action))
	sb.WriteString(fmt.Sprintf("  mode: all\n"))
	sb.WriteString(fmt.Sprintf("  selector:\n    namespaces:\n    - %s\n", c.namespace))
	for k, v := range mapped.Params {
		sb.WriteString(fmt.Sprintf("  %s: %s\n", k, v))
	}
	return sb.String()
}

func crdKindLower(kind string) string {
	switch kind {
	case "PodChaos":
		return "podchaos"
	case "NetworkChaos":
		return "networkchaos"
	case "IOChaos":
		return "iochaos"
	case "StressChaos":
		return "stresschaos"
	case "TimeChaos":
		return "timechaos"
	default:
		return strings.ToLower(kind)
	}
}
