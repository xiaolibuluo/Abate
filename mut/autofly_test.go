// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package mut

import (
	"errors"
	"testing"

	"github.com/qingdujian/abate/core"
)

type mockLLM struct {
	tests     map[string]string
	seq       []string
	err       error
	callCount int
	feedbacks []string
}

func (m *mockLLM) GenerateTest(mutant core.Mutation, _ string, feedback string) (string, error) {
	m.callCount++
	m.feedbacks = append(m.feedbacks, feedback)
	if m.err != nil {
		return "", m.err
	}
	if m.callCount <= len(m.seq) {
		return m.seq[m.callCount-1], nil
	}
	if t, ok := m.tests[mutant.ID()]; ok {
		return t, nil
	}
	return "package main", nil
}

type mockRunner struct {
	passMap map[string]bool
	err     error
}

func (r *mockRunner) RunTest(_, testCode string) (bool, error) {
	if r.err != nil {
		return false, r.err
	}
	if pass, ok := r.passMap[testCode]; ok {
		return pass, nil
	}
	return true, nil
}

func makeMutation(id string) core.Mutation {
	return core.Mutation{
		File:  "test.go",
		Line:  1,
		Col:   1,
		Class: core.OpArithmetic,
		From:  "+",
		To:    "-",
	}
}

func TestNewFlywheel_Success(t *testing.T) {
	llm := &mockLLM{}
	runner := &mockRunner{}
	f, err := NewFlywheel(llm, runner, DefaultFlywheelConfig())
	if err != nil {
		t.Fatalf("NewFlywheel: %v", err)
	}
	if f.KillRate() != 0 {
		t.Error("KillRate should be 0")
	}
	if f.SurvivedCount() != 0 {
		t.Error("SurvivedCount should be 0")
	}
}

func TestNewFlywheel_NilLLM(t *testing.T) {
	runner := &mockRunner{}
	_, err := NewFlywheel(nil, runner, DefaultFlywheelConfig())
	if err == nil {
		t.Fatal("expected error for nil LLM")
	}
}

func TestNewFlywheel_NilRunner(t *testing.T) {
	llm := &mockLLM{}
	_, err := NewFlywheel(llm, nil, DefaultFlywheelConfig())
	if err == nil {
		t.Fatal("expected error for nil runner")
	}
}

func TestFlywheel_Process_Empty(t *testing.T) {
	llm := &mockLLM{}
	runner := &mockRunner{}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())
	result, err := f.Process(nil, "")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if result != nil {
		t.Error("result should be nil for empty input")
	}
}

func TestFlywheel_Process_KillOnFirstTry(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{tests: map[string]string{m.ID(): "package main\nfunc TestKill(t *testing.T) { t.Fail() }"}}
	runner := &mockRunner{passMap: map[string]bool{"package main\nfunc TestKill(t *testing.T) { t.Fail() }": false}}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())

	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	if !results[0].Killed {
		t.Error("mutant should be killed")
	}
	if results[0].Attempts != 1 {
		t.Errorf("attempts = %d, want 1", results[0].Attempts)
	}
	if f.KillRate() != 1.0 {
		t.Errorf("KillRate = %f, want 1.0", f.KillRate())
	}
	if f.SurvivedCount() != 0 {
		t.Errorf("SurvivedCount = %d, want 0", f.SurvivedCount())
	}
}

func TestFlywheel_Process_Survived(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{tests: map[string]string{m.ID(): "package main\nfunc TestNoKill(t *testing.T) {}"}}
	runner := &mockRunner{passMap: map[string]bool{"package main\nfunc TestNoKill(t *testing.T) {}": true}}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())

	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if results[0].Killed {
		t.Error("mutant should survive")
	}
	if results[0].Attempts != 3 {
		t.Errorf("attempts = %d, want 3", results[0].Attempts)
	}
	if f.KillRate() != 0 {
		t.Errorf("KillRate = %f, want 0", f.KillRate())
	}
	if f.SurvivedCount() != 1 {
		t.Errorf("SurvivedCount = %d, want 1", f.SurvivedCount())
	}
}

func TestFlywheel_Process_LLMError(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{err: errors.New("LLM unavailable")}
	runner := &mockRunner{}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())

	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if results[0].Killed {
		t.Error("mutant should survive when LLM fails")
	}
}

func TestFlywheel_Process_RunnerError(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{tests: map[string]string{m.ID(): "package main"}}
	runner := &mockRunner{err: errors.New("runner failed")}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())

	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if results[0].Killed {
		t.Error("mutant should survive when runner fails")
	}
}

func TestFlywheel_Process_MaxMutants(t *testing.T) {
	llm := &mockLLM{}
	runner := &mockRunner{}
	f, _ := NewFlywheel(llm, runner, FlywheelConfig{MaxRetries: 1, MaxMutants: 2})

	mutants := []core.Mutation{
		makeMutation("m1"),
		makeMutation("m2"),
		makeMutation("m3"),
		makeMutation("m4"),
	}
	results, err := f.Process(mutants, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("results = %d, want 2 (MaxMutants)", len(results))
	}
}
