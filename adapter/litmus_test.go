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

func TestLitmusBackend_Apply_Success(t *testing.T) {
	orig := litApplyFn
	defer func() { litApplyFn = orig }()
	var gotYAML string
	litApplyFn = func(yaml string) (string, error) {
		gotYAML = yaml
		return "created", nil
	}

	b := NewLitmusBackend("litmus", "app=nginx", "deployment")
	mapped := MappedFault{
		UpstreamType: "pod-delete",
		Target:       core.Target{Kind: core.TargetService, Name: "svc-a"},
		Params:       map[string]string{"signal": "SIGKILL"},
	}

	id, err := b.Apply(mapped)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if !strings.HasPrefix(id, "abate-litmus-") {
		t.Errorf("id = %q", id)
	}
	if !strings.Contains(gotYAML, "kind: ChaosEngine") {
		t.Errorf("YAML missing kind: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "name: pod-delete") {
		t.Errorf("YAML missing experiment name: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "appns: litmus") {
		t.Errorf("YAML missing appns: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "applabel: app=nginx") {
		t.Errorf("YAML missing applabel: %s", gotYAML)
	}
	if !strings.Contains(gotYAML, "value: SIGKILL") {
		t.Errorf("YAML missing param value: %s", gotYAML)
	}
}

func TestLitmusBackend_Apply_ApplyError(t *testing.T) {
	orig := litApplyFn
	defer func() { litApplyFn = orig }()
	litApplyFn = func(yaml string) (string, error) { return "", errors.New("kubectl failed") }

	b := NewLitmusBackend("litmus", "app=nginx", "deployment")
	_, err := b.Apply(MappedFault{UpstreamType: "cpu-hog", Params: map[string]string{}})
	if err == nil {
		t.Fatal("expected error for apply failure")
	}
}

func TestLitmusBackend_Cleanup_Success(t *testing.T) {
	orig := litDeleteFn
	defer func() { litDeleteFn = orig }()
	var deleted string
	litDeleteFn = func(name string) error {
		deleted = name
		return nil
	}

	b := NewLitmusBackend("litmus", "app=nginx", "deployment")
	if err := b.Cleanup("abate-litmus-1"); err != nil {
		t.Fatalf("Cleanup: %v", err)
	}
	if deleted != "abate-litmus-1" {
		t.Errorf("deleted = %q", deleted)
	}
}

func TestLitmusBackend_Cleanup_DeleteError(t *testing.T) {
	orig := litDeleteFn
	defer func() { litDeleteFn = orig }()
	litDeleteFn = func(name string) error { return errors.New("delete failed") }

	b := NewLitmusBackend("litmus", "app=nginx", "deployment")
	if err := b.Cleanup("abate-litmus-1"); err == nil {
		t.Fatal("expected error for delete failure")
	}
}

func TestLitApplyFn_Success(t *testing.T) {
	orig := litExecFn
	defer func() { litExecFn = orig }()
	litExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("created"), nil
	}
	result, err := litApplyFn("yaml")
	if err != nil {
		t.Fatalf("litApplyFn: %v", err)
	}
	if result != "created" {
		t.Errorf("result = %q", result)
	}
}

func TestLitApplyFn_Error(t *testing.T) {
	_, err := litApplyFn("invalid yaml")
	if err == nil {
		t.Fatal("expected error from kubectl")
	}
}

func TestLitDeleteFn_Success(t *testing.T) {
	orig := litExecFn
	defer func() { litExecFn = orig }()
	litExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("deleted"), nil
	}
	if err := litDeleteFn("abate-litmus-1"); err != nil {
		t.Fatalf("litDeleteFn: %v", err)
	}
}

func TestLitDeleteFn_Error(t *testing.T) {
	err := litDeleteFn("nonexistent")
	if err == nil {
		t.Fatal("expected error from kubectl delete")
	}
}

// TestLitDeleteFn_NotFoundIdempotent 覆盖 CES-3.1.2 幂等分支：
// kubectl 返回 not found 错误时 litDeleteFn 视为成功返回 nil。
func TestLitDeleteFn_NotFoundIdempotent(t *testing.T) {
	orig := litExecFn
	defer func() { litExecFn = orig }()
	litExecFn = func(args []string, input string) ([]byte, error) {
		return []byte("Error from server (NotFound): chaosengine \"abate-litmus-1\" not found"), errors.New("exit status 1")
	}
	if err := litDeleteFn("abate-litmus-1"); err != nil {
		t.Fatalf("litDeleteFn should return nil for not found, got: %v", err)
	}
}
