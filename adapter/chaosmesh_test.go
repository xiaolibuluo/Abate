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
	"errors"
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestChaosMeshBackend_Apply_Success(t *testing.T) {
	orig := cmApplyFn
	defer func() { cmApplyFn = orig }()
	var gotYAML string
	cmApplyFn = func(yaml string) (string, error) {
		gotYAML = yaml
		return "created", nil
	}

	b := NewChaosMeshBackend("default")
	mapped := MappedFault{
		UpstreamType: "NetworkChaos/loss",
		Target:       core.Target{Kind: core.TargetService, Name: "svc-a"},
		Params:       map[string]string{"loss": "50", "direction": "to"},
	}

	id, err := b.Apply(mapped)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if !strings.HasPrefix(id, "networkchaos/abate-") {
		t.Errorf("id = %q", id)
	}
	if !strings.Contains(gotYAML, "kind: NetworkChaos") {
		t.Errorf("YAML missing kind: %s", gotYAML)
		t.Errorf("YAML missing action: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "namespace: default") {
		t.Errorf("YAML missing namespace: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "loss: 50") {
		t.Errorf("YAML missing param: %s", gotYAML)
	}
}

func TestChaosMeshBackend_Apply_InvalidType(t *testing.T) {
	b := NewChaosMeshBackend("default")
	_, err := b.Apply(MappedFault{UpstreamType: "invalid"})
	if err == nil {
		t.Fatal("expected error for invalid upstream type")
	}
}

func TestChaosMeshBackend_Apply_ApplyError(t *testing.T) {
	orig := cmApplyFn
	defer func() { cmApplyFn = orig }()
	cmApplyFn = func(yaml string) (string, error) { return "", errors.New("kubectl failed") }

	b := NewChaosMeshBackend("default")
	_, err := b.Apply(MappedFault{UpstreamType: "PodChaos/kill", Params: map[string]string{}})
	if err == nil {
		t.Fatal("expected error for apply failure")
	}
}

func TestChaosMeshBackend_Cleanup_Success(t *testing.T) {
	orig := cmDeleteFn
	defer func() { cmDeleteFn = orig }()
	var deleted string
	cmDeleteFn = func(name string) error {
		deleted = name
		return nil
	}

	b := NewChaosMeshBackend("default")
	if err := b.Cleanup("podchaos/abate-1"); err != nil {
		t.Fatalf("Cleanup: %v", err)
	}
	if deleted != "podchaos/abate-1" {
		t.Errorf("deleted = %q", deleted)
	}
}

func TestChaosMeshBackend_Cleanup_DeleteError(t *testing.T) {
	orig := cmDeleteFn
	defer func() { cmDeleteFn = orig }()
	cmDeleteFn = func(name string) error { return errors.New("delete failed") }

	b := NewChaosMeshBackend("default")
	if err := b.Cleanup("podchaos/abate-1"); err == nil {
		t.Fatal("expected error for delete failure")
	}
}

func TestCmApplyFn_Success(t *testing.T) {
	orig := cmExecFn
	defer func() { cmExecFn = orig }()
	cmExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("created"), nil
	}
	result, err := cmApplyFn("yaml")
	if err != nil {
		t.Fatalf("cmApplyFn: %v", err)
	}
	if result != "created" {
		t.Errorf("result = %q", result)
	}
}

func TestCmApplyFn_Error(t *testing.T) {
	_, err := cmApplyFn("invalid yaml")
	if err == nil {
		t.Fatal("expected error from kubectl")
	}
}

func TestCmDeleteFn_InvalidName(t *testing.T) {
	err := cmDeleteFn("no-slash")
	if err == nil {
		t.Fatal("expected error for invalid resource name")
	}
}

func TestCmDeleteFn_Success(t *testing.T) {
	orig := cmExecFn
	defer func() { cmExecFn = orig }()
	cmExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("deleted"), nil
	}
	if err := cmDeleteFn("podchaos/abate-1"); err != nil {
		t.Fatalf("cmDeleteFn: %v", err)
	}
}

func TestCmDeleteFn_KubectlError(t *testing.T) {
	err := cmDeleteFn("podchaos/nonexistent")
	if err == nil {
		t.Fatal("expected error from kubectl delete")
	}
}

// TestCmDeleteFn_NotFoundIdempotent 覆盖 CES-3.1.2 幂等分支：
// kubectl 返回 not found 错误时 cmDeleteFn 视为成功返回 nil。
func TestCmDeleteFn_NotFoundIdempotent(t *testing.T) {
	orig := cmExecFn
	defer func() { cmExecFn = orig }()
	cmExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("Error from server (NotFound): podchaos \"abate-1\" not found"), errors.New("exit status 1")
	}
	if err := cmDeleteFn("podchaos/abate-1"); err != nil {
		t.Fatalf("cmDeleteFn should return nil for not found, got: %v", err)
	}
}

func TestCrdKindLower(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"PodChaos", "podchaos"},
		{"NetworkChaos", "networkchaos"},
		{"IOChaos", "iochaos"},
		{"StressChaos", "stresschaos"},
		{"TimeChaos", "timechaos"},
		{"HTTPChaos", "httpchaos"},
	}
	for _, tt := range tests {
		if got := crdKindLower(tt.input); got != tt.want {
			t.Errorf("crdKindLower(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
