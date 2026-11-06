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
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestFlywheel_Process_KillOnRetry(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{
		seq: []string{"package main", "package main\nfunc TestKill(t *testing.T) { t.Fail() }"},
	}
	runner := &mockRunner{
		passMap: map[string]bool{
			"package main": true,
			"package main\nfunc TestKill(t *testing.T) { t.Fail() }": false,
		},
	}
	f, _ := NewFlywheel(llm, runner, FlywheelConfig{MaxRetries: 3, MaxMutants: 0})

	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if !results[0].Killed {
		t.Error("mutant should be killed on retry")
	}
}

func TestFlywheel_Results(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{tests: map[string]string{m.ID(): "package main"}}
	runner := &mockRunner{passMap: map[string]bool{"package main": false}}
	f, _ := NewFlywheel(llm, runner, DefaultFlywheelConfig())

	_, _ = f.Process([]core.Mutation{m}, "source")
	results := f.Results()
	if len(results) != 1 {
		t.Fatalf("Results = %d, want 1", len(results))
	}
	if !results[0].Killed {
		t.Error("should be killed")
	}
}

func TestFlywheel_FeedbackLoop_SyntaxError(t *testing.T) {
	m := makeMutation("m1")
	llm := &mockLLM{
		seq: []string{
			"func broken(",
			"package main\nfunc TestKill(t *testing.T) { t.Fail() }",
		},
	}
	runner := &mockRunner{passMap: map[string]bool{
		"package main\nfunc TestKill(t *testing.T) { t.Fail() }": false,
	}}
	f, _ := NewFlywheel(llm, runner, FlywheelConfig{MaxRetries: 3, MaxMutants: 10})
	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if !results[0].Killed {
		t.Error("mutant should be killed on retry after syntax fix")
	}
	if len(llm.feedbacks) != 2 {
		t.Fatalf("feedbacks = %d, want 2", len(llm.feedbacks))
	}
	if llm.feedbacks[0] != "" {
		t.Errorf("first feedback = %q, want empty", llm.feedbacks[0])
	}
	if !strings.Contains(llm.feedbacks[1], "syntax error") {
		t.Errorf("second feedback = %q, want syntax error feedback", llm.feedbacks[1])
	}
}

func TestFlywheel_FeedbackLoop_StrongerTest(t *testing.T) {
	m := makeMutation("m1")
	weak := "package main\nfunc TestWeak(t *testing.T) {}"
	strong := "package main\nfunc TestKill(t *testing.T) { t.Fail() }"
	llm := &mockLLM{seq: []string{weak, strong}}
	runner := &mockRunner{passMap: map[string]bool{
		weak:  true,
		strong: false,
	}}
	f, _ := NewFlywheel(llm, runner, FlywheelConfig{MaxRetries: 3, MaxMutants: 10})
	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if !results[0].Killed {
		t.Error("mutant should be killed by stronger test")
	}
	if len(llm.feedbacks) != 2 {
		t.Fatalf("feedbacks = %d, want 2", len(llm.feedbacks))
	}
	if !strings.Contains(llm.feedbacks[1], "stronger test") {
		t.Errorf("second feedback = %q, want stronger-test feedback", llm.feedbacks[1])
	}
}

func TestDefaultFlywheelConfig(t *testing.T) {
	cfg := DefaultFlywheelConfig()
	if cfg.MaxRetries != 3 {
		t.Errorf("MaxRetries = %d", cfg.MaxRetries)
	}
	if cfg.MaxMutants != 100 {
		t.Errorf("MaxMutants = %d", cfg.MaxMutants)
	}
}
