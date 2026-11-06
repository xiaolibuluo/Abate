//go:build integration

// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// 适配器集成测试：需要真实 Kubernetes 集群（k3s）与 kubectl。
// Adapter integration tests: require a real Kubernetes cluster (k3s) and kubectl.
// 运行 Run: go test -tags integration ./adapter/ -v -count=1
package adapter_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qingdujian/abate/adapter"
)

// kubectlRun 执行 kubectl 并返回输出。
func kubectlRun(t *testing.T, args []string, stdin string) string {
	t.Helper()
	cmd := exec.Command("kubectl", args...)
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("kubectl %s: %v: %s", strings.Join(args, " "), err, out)
	}
	return string(out)
}

const chaosMeshCRD = `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: networkchaos.chaos-mesh.org
spec:
  group: chaos-mesh.org
  scope: Namespaced
  names:
    plural: networkchaos
    singular: networkchaos
    kind: NetworkChaos
  versions:
    - name: v1alpha1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          x-kubernetes-preserve-unknown-fields: true
`

const litmusCRD = `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: chaosengines.litmuschaos.io
spec:
  group: litmuschaos.io
  scope: Namespaced
  names:
    plural: chaosengines
    singular: chaosengine
    kind: ChaosEngine
  versions:
    - name: v1alpha1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          x-kubernetes-preserve-unknown-fields: true
`

// TestChaosMeshIntegration 在真实 k3s 上验证 ChaosMesh 适配器全链路。
func TestChaosMeshIntegration(t *testing.T) {
	kubectlRun(t, []string{"apply", "-f", "-"}, chaosMeshCRD)
	t.Cleanup(func() {
		kubectlRun(t, []string{"delete", "crd", "networkchaos.chaos-mesh.org", "--ignore-not-found"}, "")
	})
	kubectlRun(t, []string{"wait", "--for=condition=established", "--timeout=60s", "crd/networkchaos.chaos-mesh.org"}, "")

	backend := adapter.NewChaosMeshBackend("default")
	id, err := backend.Apply(adapter.MappedFault{
		UpstreamType: "NetworkChaos/delay",
		Params:       map[string]string{"latency": "100ms"},
	})
	if err != nil {
		t.Fatalf("chaosmesh apply: %v", err)
	}
	if !strings.HasPrefix(id, "networkchaos/abate-") {
		t.Fatalf("chaosmesh unexpected id: %q", id)
	}
	name := strings.TrimPrefix(id, "networkchaos/")
	if out := kubectlRun(t, []string{"get", "networkchaos", name, "-o", "name"}, ""); !strings.Contains(out, name) {
		t.Fatalf("chaosmesh CR %s not found: %s", name, out)
	}
	if err := backend.Cleanup(id); err != nil {
		t.Fatalf("chaosmesh cleanup: %v", err)
	}
	if out := kubectlRun(t, []string{"get", "networkchaos", name, "-o", "name", "--ignore-not-found"}, ""); strings.Contains(out, name) {
		t.Fatalf("chaosmesh CR %s still exists after cleanup: %s", name, out)
	}
	// CES-3.1.2：Cleanup 必须可重复调用。
	if err := backend.Cleanup(id); err != nil {
		t.Fatalf("chaosmesh cleanup idempotency: %v", err)
	}
}

// TestLitmusIntegration 在真实 k3s 上验证 Litmus 适配器全链路。
func TestLitmusIntegration(t *testing.T) {
	kubectlRun(t, []string{"apply", "-f", "-"}, litmusCRD)
	t.Cleanup(func() {
		kubectlRun(t, []string{"delete", "crd", "chaosengines.litmuschaos.io", "--ignore-not-found"}, "")
	})
	kubectlRun(t, []string{"wait", "--for=condition=established", "--timeout=60s", "crd/chaosengines.litmuschaos.io"}, "")

	backend := adapter.NewLitmusBackend("default", "app=nginx", "deployment")
	id, err := backend.Apply(adapter.MappedFault{
		UpstreamType: "pod-network-latency",
		Params:       map[string]string{"TOTAL_CHAOS_DURATION": "30"},
	})
	if err != nil {
		t.Fatalf("litmus apply: %v", err)
	}
	if !strings.HasPrefix(id, "abate-litmus-") {
		t.Fatalf("litmus unexpected id: %q", id)
	}
	if out := kubectlRun(t, []string{"get", "chaosengine", id, "-o", "name"}, ""); !strings.Contains(out, id) {
		t.Fatalf("litmus ChaosEngine %s not found: %s", id, out)
	}
	if err := backend.Cleanup(id); err != nil {
		t.Fatalf("litmus cleanup: %v", err)
	}
	if out := kubectlRun(t, []string{"get", "chaosengine", id, "-o", "name", "--ignore-not-found"}, ""); strings.Contains(out, id) {
		t.Fatalf("litmus ChaosEngine %s still exists after cleanup: %s", id, out)
	}
	// CES-3.1.2：Cleanup 必须可重复调用。
	if err := backend.Cleanup(id); err != nil {
		t.Fatalf("litmus cleanup idempotency: %v", err)
	}
}

// TestGremlinCLIContract 用 CLI shim 验证 Gremlin 适配器的命令行契约。
// 真实 Gremlin SaaS 环境无法本地补齐，此处验证适配器与 gremlin CLI 的集成契约。
func TestGremlinCLIContract(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "args.log")
	shim := "#!/bin/sh\necho \"$@\" >> " + logFile + "\nexit 0\n"
	if err := os.WriteFile(filepath.Join(dir, "gremlin"), []byte(shim), 0o755); err != nil {
		t.Fatalf("write shim: %v", err)
	}
	t.Setenv("PATH", dir+":"+os.Getenv("PATH"))

	backend := adapter.NewGremlinBackend("team-1")
	id, err := backend.Apply(adapter.MappedFault{
		UpstreamType: "latency",
		Params:       map[string]string{"length": "60"},
	})
	if err != nil {
		t.Fatalf("gremlin apply: %v", err)
	}
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read args log: %v", err)
	}
	log := string(data)
	if !strings.Contains(log, "attack latency") || !strings.Contains(log, "-l 60") || strings.Contains(log, "--guid") {
		t.Fatalf("gremlin apply args mismatch: %s", log)
	}
	if err := backend.Cleanup(id); err != nil {
		t.Fatalf("gremlin cleanup: %v", err)
	}
	data, err = os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("reread args log: %v", err)
	}
	if !strings.Contains(string(data), "rollback") {
		t.Fatalf("gremlin cleanup args mismatch: %s", string(data))
	}
}
