// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package engine 实现 CES 第 3 章的实验执行语义与状态机。
// Package engine implements the CES Chapter 3 execution semantics and state machine.
package engine

import "fmt"

// State 是实验生命周期状态（CES-3.1）。
// State is an experiment lifecycle state (CES-3.1).
type State string

const (
	// StatePending 实验已受理，等待准入。
	StatePending State = "Pending"
	// StatePreflight 实验前准入检查中（CES-3.1.4）。
	StatePreflight State = "Preflight"
	// StateInjecting 故障注入进行中。
	StateInjecting State = "Injecting"
	// StateObserving 稳态观测中。
	StateObserving State = "Observing"
	// StateRecovering 回滚与残留清理中（CES-3.1.2）。
	StateRecovering State = "Recovering"
	// StateVerdict 终态：判定完成，产出报告。
	StateVerdict State = "Verdict"
)

// transitions 是规范强制的合法迁移表（CES-3.1）。
// 键为源状态，值为可达目标状态集合。
//
// transitions is the normative legal transition table (CES-3.1).
var transitions = map[State]map[State]bool{
	StatePending: {
		StatePreflight: true,
		StateVerdict:   true, // 准入前取消，判定为 CANCELLED
	},
	StatePreflight: {
		StateInjecting:  true,
		StateVerdict:    true, // 准入失败，判定为 REJECTED（CES-3.1.4）
		StateRecovering: true, // 红线触发（CES-3.1.1）
	},
	StateInjecting: {
		StateObserving:  true,
		StateRecovering: true,
	},
	StateObserving: {
		StateRecovering: true, // 观测结束或红线触发，必经 Recovering
	},
	StateRecovering: {
		StateVerdict: true,
	},
	StateVerdict: {},
}

// CanTransit 报告 from→to 是否为合法迁移（CES-3.1）。
// CanTransit reports whether from→to is a legal transition (CES-3.1).
func CanTransit(from, to State) bool {
	return transitions[from][to]
}

// Transition 校验并执行一次状态迁移；非法迁移视为实现缺陷（CES-3.1）。
// Transition validates and performs one state transition; an illegal
// transition is an implementation defect and returns an error.
func Transition(from, to State) (State, error) {
	if !CanTransit(from, to) {
		return from, fmt.Errorf("illegal transition %q -> %q (CES-3.1)", from, to)
	}
	return to, nil
}

// MustRecover 报告状态 s 下红线触发时是否必须迁移至 Recovering（CES-3.1.1）。
// MustRecover reports whether a red-line trigger in state s must transition
// to Recovering (CES-3.1.1).
func MustRecover(s State) bool {
	switch s {
	case StatePreflight, StateInjecting, StateObserving:
		return true
	}
	return false
}

// Terminal 报告状态是否为终态。
// Terminal reports whether the state is terminal.
func (s State) Terminal() bool {
	return s == StateVerdict
}

// States 返回规范定义的全部状态，顺序固定。
// States returns all specification-defined states in a fixed order.
func States() []State {
	return []State{
		StatePending, StatePreflight, StateInjecting,
		StateObserving, StateRecovering, StateVerdict,
	}
}
