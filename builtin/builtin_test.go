// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package builtin

import (
	"fmt"
	"os"
	"testing"

	"github.com/qingdujian/abate/adapter"
	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/plugin"
	"github.com/qingdujian/abate/probe"
	"github.com/qingdujian/abate/reporter"
	"github.com/qingdujian/abate/validator"
)

type mockBackend struct{}

func (m *mockBackend) Apply(mapped adapter.MappedFault) (string, error) {
	return "upstream-mock", nil
}

func (m *mockBackend) Cleanup(upstreamID string) error {
	return nil
}

func TestRegister(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	if err := Register(r, dir); err != nil {
		t.Fatalf("Register: %v", err)
	}

	names := r.List()
	if len(names) < 4 {
		t.Fatalf("expected >= 4 plugins, got %d: %v", len(names), names)
	}
}

func TestRegisterAdapter(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	if err := Register(r, dir); err != nil {
		t.Fatalf("Register: %v", err)
	}

	mt := adapter.NewMappingTable(adapter.DefaultChaosMeshMappings())
	if err := RegisterAdapter(r, "chaosmesh", "0.1.0", &mockBackend{}, mt); err != nil {
		t.Fatalf("RegisterAdapter: %v", err)
	}

	_, err := r.FindInjector(adapter.NewMappingTable(adapter.DefaultChaosMeshMappings()).MappedClasses()[0], core.TargetService)
	if err != nil {
		t.Errorf("FindInjector after adapter register: %v", err)
	}
}

func TestEnsureCapability(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	if err := Register(r, dir); err != nil {
		t.Fatalf("Register: %v", err)
	}

	for _, role := range []plugin.Role{plugin.RoleInjector, plugin.RoleProbe, plugin.RoleValidator, plugin.RoleReporter} {
		if err := EnsureCapability(r, role); err != nil {
			t.Errorf("EnsureCapability(%q): %v", role, err)
		}
	}
}

func TestRegister_InvalidDir(t *testing.T) {
	tmpFile, _ := os.CreateTemp("", "builtin-blocker")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	r := plugin.NewRegistry()
	if err := Register(r, tmpFile.Name()); err == nil {
		t.Fatal("expected error for invalid dir")
	}
}

func TestRegister_Duplicate(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	if err := Register(r, dir); err != nil {
		t.Fatalf("first Register: %v", err)
	}
	if err := Register(r, dir); err == nil {
		t.Fatal("expected error for duplicate register")
	}
}

func TestRegister_ProbeConflict(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	p := probe.New(42)
	_ = r.Register(plugin.Plugin{Probe: p})
	if err := Register(r, dir); err == nil {
		t.Fatal("expected error for probe conflict")
	}
}

func TestRegister_ValidatorConflict(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	v := validator.New()
	_ = r.Register(plugin.Plugin{Validator: v})
	if err := Register(r, dir); err == nil {
		t.Fatal("expected error for validator conflict")
	}
}

func TestRegister_ReporterConflict(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	rep, _ := reporter.New(dir)
	_ = r.Register(plugin.Plugin{Reporter: rep})
	if err := Register(r, dir); err == nil {
		t.Fatal("expected error for reporter conflict")
	}
}

func TestRegister_ReporterFail(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	old := reporterNew
	reporterNew = func(string) (*reporter.Reporter, error) {
		return nil, fmt.Errorf("injected failure")
	}
	defer func() { reporterNew = old }()

	r := plugin.NewRegistry()
	if err := Register(r, dir); err == nil {
		t.Fatal("expected error for reporter failure")
	}
}

func TestRegisterAdapter_NilBackend(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	_ = Register(r, dir)
	mt := adapter.NewMappingTable(adapter.DefaultChaosMeshMappings())
	if err := RegisterAdapter(r, "bad", "1", nil, mt); err == nil {
		t.Fatal("expected error for nil backend")
	}
}

func TestRegisterAdapter_Duplicate(t *testing.T) {
	dir, _ := os.MkdirTemp("", "builtin-test")
	defer os.RemoveAll(dir)

	r := plugin.NewRegistry()
	_ = Register(r, dir)
	mt := adapter.NewMappingTable(adapter.DefaultChaosMeshMappings())
	_ = RegisterAdapter(r, "chaosmesh", "0.1.0", &mockBackend{}, mt)
	if err := RegisterAdapter(r, "chaosmesh", "0.1.0", &mockBackend{}, mt); err == nil {
		t.Fatal("expected error for duplicate adapter")
	}
}

func TestEnsureCapability_Empty(t *testing.T) {
	r := plugin.NewRegistry()
	if err := EnsureCapability(r, plugin.RoleInjector); err == nil {
		t.Error("expected error for empty registry")
	}
}

func TestEnsureCapability_UnknownRole(t *testing.T) {
	r := plugin.NewRegistry()
	if err := EnsureCapability(r, plugin.Role("unknown")); err == nil {
		t.Error("expected error for unknown role")
	}
}