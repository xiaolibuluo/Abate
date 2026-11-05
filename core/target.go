// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

import (
	"fmt"
	"sort"
)

// TargetKind 是 CES-2.2.1 定义的三类目标寻址模式。
// TargetKind is one of the three addressing modes defined by CES-2.2.1.
type TargetKind string

const (
	// TargetService 服务寻址：按服务名选择一组实例。
	TargetService TargetKind = "service"
	// TargetInstance 实例寻址：直接指定实例标识。
	TargetInstance TargetKind = "instance"
	// TargetDependency 依赖面寻址：选择某服务的外部依赖（DB、缓存、下游 API）。
	TargetDependency TargetKind = "dependency"
)

// Target 描述实验作用目标（CES-2.2.1：选择器语义必须确定性）。
// Target describes the experiment target (CES-2.2.1: selector semantics must
// be deterministic).
type Target struct {
	// Kind 寻址模式，必填。
	Kind TargetKind `yaml:"kind" json:"kind"`
	// Name 目标名：服务名、实例 ID 或依赖名。
	Name string `yaml:"name" json:"name"`
	// Selector 键值选择器；多键之间为逻辑与，匹配结果按字典序稳定排序。
	// Selector is a key-value matcher; keys are ANDed, matches are sorted
	// lexicographically for deterministic iteration.
	Selector map[string]string `yaml:"selector,omitempty" json:"selector,omitempty"`
}

// Validate 校验目标定义：寻址模式合法、名称非空。
// Validate checks the target: a valid addressing mode and a non-empty name.
func (t Target) Validate() error {
	switch t.Kind {
	case TargetService, TargetInstance, TargetDependency:
	default:
		return fmt.Errorf("target.kind: invalid kind %q (CES-2.2.1)", t.Kind)
	}
	if t.Name == "" {
		return fmt.Errorf("target.name: must not be empty (CES-2.2.1)")
	}
	return nil
}

// SelectorKeys 返回字典序升序的选择器键，保证序列化与日志输出确定性。
// SelectorKeys returns selector keys in ascending lexicographic order so that
// serialization and logging are deterministic.
func (t Target) SelectorKeys() []string {
	keys := make([]string, 0, len(t.Selector))
	for k := range t.Selector {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Matches 报告给定属性集合是否满足选择器（全键等值匹配，逻辑与）。
// Matches reports whether the given attribute set satisfies the selector
// (equality on every key, logical AND).
func (t Target) Matches(attrs map[string]string) bool {
	for k, v := range t.Selector {
		if attrs[k] != v {
			return false
		}
	}
	return true
}
