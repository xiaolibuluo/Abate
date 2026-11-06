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

func TestGremlinBackend_Apply_Success(t *testing.T) {
	orig := gremlinExecFn
	defer func() { gremlinExecFn = orig }()
	var gotArgs []string
	gremlinExecFn = func(args []string) ([]byte, error) {
		gotArgs = args
		return []byte("ok"), nil
	}

	b := NewGremlinBackend("team-1")
	mapped := MappedFault{
		UpstreamType: "latency/delay",
		Target:       core.Target{Kind: core.TargetService, Name: "svc-a"},
		Params:       map[string]string{"ms": "100"},
	}

	id, err := b.Apply(mapped)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if !strings.HasPrefix(id, "gremlin-") {
		t.Errorf("id = %q", id)
	}
	if gotArgs[0] != "attack" || gotArgs[1] != "delay" {
		t.Errorf("args = %v", gotArgs)
	}
	if contains(gotArgs, "--guid") {
		t.Errorf("args must not contain --guid: %v", gotArgs)
	}
	if !containsPair(gotArgs, "-m", "100") {
		t.Errorf("args missing -m 100: %v", gotArgs)
	}
}

func TestGremlinBackend_Apply_NoSlash(t *testing.T) {
	orig := gremlinExecFn
	defer func() { gremlinExecFn = orig }()
	gremlinExecFn = func(args []string) ([]byte, error) {
		return []byte("ok"), nil
	}

	b := NewGremlinBackend("team-1")
	id, err := b.Apply(MappedFault{UpstreamType: "cpu", Params: map[string]string{}})
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if id == "" {
		t.Error("id should not be empty")
	}
}

func TestGremlinBackend_Apply_ExecError(t *testing.T) {
	orig := gremlinExecFn
	defer func() { gremlinExecFn = orig }()
	gremlinExecFn = func(args []string) ([]byte, error) {
		return []byte("fail"), errors.New("gremlin not found")
	}

	b := NewGremlinBackend("team-1")
	_, err := b.Apply(MappedFault{UpstreamType: "latency/delay", Params: map[string]string{}})
	if err == nil {
		t.Fatal("expected error for exec failure")
	}
}

func TestGremlinBackend_Cleanup_Success(t *testing.T) {
	orig := gremlinExecFn
	defer func() { gremlinExecFn = orig }()
	var gotArgs []string
	gremlinExecFn = func(args []string) ([]byte, error) {
		gotArgs = args
		return []byte("ok"), nil
	}

	b := NewGremlinBackend("team-1")
	if err := b.Cleanup("gremlin-1"); err != nil {
		t.Fatalf("Cleanup: %v", err)
	}
	if len(gotArgs) != 1 || gotArgs[0] != "rollback" {
		t.Errorf("args = %v, want [rollback]", gotArgs)
	}
}

func TestGremlinBackend_Cleanup_ExecError(t *testing.T) {
	orig := gremlinExecFn
	defer func() { gremlinExecFn = orig }()
	gremlinExecFn = func(args []string) ([]byte, error) {
		return []byte("fail"), errors.New("stop failed")
	}

	b := NewGremlinBackend("team-1")
	if err := b.Cleanup("gremlin-1"); err == nil {
		t.Fatal("expected error for cleanup failure")
	}
}

func TestGremlinExecFn(t *testing.T) {
	_, err := gremlinExecFn([]string{"--version"})
	if err == nil {
		t.Fatal("expected error from gremlin CLI")
	}
}

func TestBuildGremlinArgs_WithSlash(t *testing.T) {
	args := buildGremlinArgs(MappedFault{
		UpstreamType: "network/loss",
		Params:       map[string]string{"percent": "50"},
	})
	if args[1] != "loss" {
		t.Errorf("method = %q, want loss", args[1])
	}
	if !containsPair(args, "--percent", "50") {
		t.Errorf("args missing --percent 50: %v", args)
	}
}

// TestBuildGremlinArgs_ShortFlags 验证已确认参数映射为真实 CLI 短参数，
// 且 flag 与值拆分为独立数组元素（官方 latency 参数表）。
func TestBuildGremlinArgs_ShortFlags(t *testing.T) {
	args := buildGremlinArgs(MappedFault{
		UpstreamType: "latency/delay",
		Params: map[string]string{
			"ip":     "10.0.0.1",
			"port":   "8080",
			"ms":     "200",
			"jitter": "50",
			"length": "120",
		},
	})
	if args[0] != "attack" || args[1] != "delay" {
		t.Errorf("args = %v", args)
	}
	for _, p := range [][2]string{
		{"-i", "10.0.0.1"}, {"-p", "8080"}, {"-m", "200"}, {"-j", "50"}, {"-l", "120"},
	} {
		if !containsPair(args, p[0], p[1]) {
			t.Errorf("args missing %s %s: %v", p[0], p[1], args)
		}
	}
	if contains(args, "--guid") {
		t.Errorf("args must not contain --guid: %v", args)
	}
}

func TestDefaultGremlinMappings(t *testing.T) {
	mt := NewMappingTable(DefaultGremlinMappings())
	fm, err := mt.Lookup("F1.1")
	if err != nil {
		t.Fatalf("Lookup F1.1: %v", err)
	}
	if fm.UpstreamType != "latency/delay" {
		t.Errorf("UpstreamType = %q", fm.UpstreamType)
	}
	degs := mt.Degradations()
	if len(degs) != 10 {
		t.Errorf("degradations = %d, want 10", len(degs))
	}
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

// containsPair 校验 flag 与其值作为相邻的独立数组元素出现。
// containsPair asserts a flag and its value appear as adjacent elements.
func containsPair(slice []string, flag, value string) bool {
	for i := 0; i+1 < len(slice); i++ {
		if slice[i] == flag && slice[i+1] == value {
			return true
		}
	}
	return false
}
