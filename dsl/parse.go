// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package dsl 实现 CES 第 2 章实验 DSL 的加载与校验。
// 零依赖红线：仅使用标准库，解析 YAML 严格子集（映射 + 标量 + 列表），
// 未知字段不拒绝（CES-2.1.1），校验失败输出结构化错误（CES-2.1.2）。
//
// Package dsl implements CES Chapter 2 experiment DSL loading and validation.
// Zero-dependency red line: stdlib only, parsing a strict YAML subset
// (mappings + scalars + lists). Unknown fields are tolerated (CES-2.1.1);
// validation failures produce structured errors (CES-2.1.2).
package dsl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qingdujian/abate/core"
)

// Error 是结构化 DSL 错误（CES-2.1.2）：字段路径 + 原因 + 规范条款号。
// Error is a structured DSL error (CES-2.1.2): field path, reason, clause ID.
type Error struct {
	Path   string `json:"path"`
	Reason string `json:"reason"`
	Clause string `json:"clause"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Path, e.Reason, e.Clause)
}

// Parse 将 DSL 文本解析为 core.Experiment 并执行规范校验。
// Parse parses DSL text into a core.Experiment and runs specification checks.
func Parse(text string) (core.Experiment, error) {
	e, err := decode(text)
	if err != nil {
		return core.Experiment{}, err
	}
	if err := e.Validate(); err != nil {
		return core.Experiment{}, err
	}
	return e, nil
}

// decode 解析严格行式 YAML 子集：两个空格缩进表达层级。
// decode parses a strict line-oriented YAML subset: two-space indentation.
func decode(text string) (core.Experiment, error) {
	var e core.Experiment
	section := ""
	listKey := ""
	for ln, raw := range strings.Split(text, "\n") {
		line := strings.TrimRight(raw, " \t\r")
		trim := strings.TrimSpace(line)
		if trim == "" || strings.HasPrefix(trim, "#") {
			continue
		}
		indent := len(line) - len(strings.TrimLeft(line, " "))
		if indent%2 != 0 {
			return e, Error{Path: fmt.Sprintf("line %d", ln+1),
				Reason: "indent must be multiples of two spaces", Clause: "CES-2.1.2"}
		}
		key, val, isItem, err := splitKV(trim)
		if err != nil {
			return e, Error{Path: fmt.Sprintf("line %d", ln+1),
				Reason: err.Error(), Clause: "CES-2.1.2"}
		}
		switch indent {
		case 0:
			section, listKey = key, ""
			if err := applyTop(&e, key, val); err != nil {
				return e, err
			}
		case 2:
			listKey = ""
			if isItem {
				if err := applyListItem(&e, section, key, val, ln+1); err != nil {
					return e, err
				}
				listKey = section
				continue
			}
			if err := applySection(&e, section, key, val); err != nil {
				return e, err
			}
		case 4:
			if listKey == "" {
				return e, Error{Path: fmt.Sprintf("line %d", ln+1),
					Reason: "nested key outside list item", Clause: "CES-2.1.2"}
			}
			if err := applyListField(&e, listKey, key, val); err != nil {
				return e, err
			}
		default:
			return e, Error{Path: fmt.Sprintf("line %d", ln+1),
				Reason: "nesting deeper than three levels is not supported",
				Clause: "CES-2.1.2"}
		}
	}
	return e, nil
}

// splitKV 拆分 key: value 或列表项 - key: value。
// splitKV splits a 'key: value' line or a '- key: value' list item.
func splitKV(s string) (key, val string, isItem bool, err error) {
	if strings.HasPrefix(s, "- ") {
		isItem = true
		s = strings.TrimPrefix(s, "- ")
	}
	i := strings.Index(s, ":")
	if i < 0 {
		return "", "", false, fmt.Errorf("expected 'key: value', got %q", s)
	}
	key = strings.TrimSpace(s[:i])
	val = unquote(strings.TrimSpace(s[i+1:]))
	if key == "" {
		return "", "", false, fmt.Errorf("empty key in %q", s)
	}
	return key, val, isItem, nil
}

// unquote 剥离标量值的首尾匹配引号（YAML 标量语义）。
// unquote strips matching surrounding quotes from a scalar value.
func unquote(v string) string {
	if len(v) >= 2 {
		if (v[0] == '"' && v[len(v)-1] == '"') ||
			(v[0] == '\'' && v[len(v)-1] == '\'') {
			return v[1 : len(v)-1]
		}
	}
	return v
}

// parseInt 解析整数字段并包装为结构化错误。
// parseInt parses an integer field into a structured error on failure.
func parseInt(path, raw string) (int64, error) {
	n, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, Error{Path: path, Reason: "expected integer, got " + strconv.Quote(raw),
			Clause: "CES-2.1.2"}
	}
	return n, nil
}

// applyTop 处理顶层字段；未知字段不拒绝（CES-2.1.1）。
// applyTop handles top-level fields; unknown fields are tolerated (CES-2.1.1).
func applyTop(e *core.Experiment, key, val string) error {
	switch key {
	case "id":
		e.ID = val
	case "seed":
		n, err := parseInt("seed", val)
		if err != nil {
			return err
		}
		e.Seed = n
	case "criticality":
		n, err := parseInt("criticality", val)
		if err != nil {
			return err
		}
		e.Criticality = int(n)
	case "recoveryBudgetMs":
		n, err := parseInt("recoveryBudgetMs", val)
		if err != nil {
			return err
		}
		e.RecoveryBudgetMs = n
	case "target", "fault", "blastRadius", "steadyState", "abortOn", "params":
		// 节标题，值由更深层级填充。
	default:
		// CES-2.1.1：未知字段容忍，不拒绝。
	}
	return nil
}

// applySection 处理二级字段。
// applySection handles second-level fields.
func applySection(e *core.Experiment, section, key, val string) error {
	switch section {
	case "target":
		switch key {
		case "kind":
			e.Target.Kind = core.TargetKind(val)
		case "name":
			e.Target.Name = val
		default:
			if e.Target.Selector == nil {
				e.Target.Selector = map[string]string{}
			}
			e.Target.Selector[key] = val
		}
	case "fault":
		if key == "type" {
			e.Fault.Class = core.FaultClass(val)
		} else {
			if e.Fault.Params == nil {
				e.Fault.Params = map[string]string{}
			}
			e.Fault.Params[key] = val
		}
	case "blastRadius":
		return applyBlast(e, key, val)
	}
	return nil
}

// applyBlast 填充爆炸半径字段（CES-2.5.1）。
// applyBlast fills blast-radius fields (CES-2.5.1).
func applyBlast(e *core.Experiment, key, val string) error {
	n, err := parseInt("blastRadius."+key, val)
	if err != nil {
		return err
	}
	switch key {
	case "maxInstances":
		e.Blast.MaxInstances = int(n)
	case "maxDurationMs":
		e.Blast.MaxDurationMs = n
	case "maxTraffic":
		e.Blast.MaxTraffic = core.BP(n)
	}
	return nil
}

// applyListItem 开启 steadyState / abortOn 列表项（CES-2.6.1）。
// applyListItem opens a steadyState / abortOn list item (CES-2.6.1).
func applyListItem(e *core.Experiment, section, key, val string, ln int) error {
	if section != "steadyState" && section != "abortOn" {
		return Error{Path: fmt.Sprintf("line %d", ln),
			Reason: "list items only allowed under steadyState/abortOn",
			Clause: "CES-2.1.2"}
	}
	a := core.SLIAssert{}
	if err := assignAssert(&a, key, val); err != nil {
		return err
	}
	if section == "steadyState" {
		e.Steady = append(e.Steady, a)
	} else {
		e.AbortOn = append(e.AbortOn, a)
	}
	return nil
}

// applyListField 填充当前列表项的后续字段。
// applyListField fills subsequent fields of the current list item.
func applyListField(e *core.Experiment, listKey, key, val string) error {
	var list *[]core.SLIAssert
	if listKey == "steadyState" {
		list = &e.Steady
	} else {
		list = &e.AbortOn
	}
	if len(*list) == 0 {
		return Error{Path: listKey + "." + key,
			Reason: "field without preceding list item", Clause: "CES-2.1.2"}
	}
	return assignAssert(&(*list)[len(*list)-1], key, val)
}

// assignAssert 将字段写入断言结构；windowMs 非整数时返回结构化错误（CES-2.1.2 严格语义）。
// assignAssert writes one field into an assertion; a non-integer windowMs
// yields a structured error (strict CES-2.1.2 semantics).
func assignAssert(a *core.SLIAssert, key, val string) error {
	switch key {
	case "sli":
		a.SLI = val
	case "assert":
		a.Assert = val
	case "method":
		a.Method = val
	case "windowMs":
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return Error{Path: "windowMs",
				Reason: "expected integer, got " + strconv.Quote(val),
				Clause: "CES-2.1.2"}
		}
		a.WindowMs = n
	}
	return nil
}
