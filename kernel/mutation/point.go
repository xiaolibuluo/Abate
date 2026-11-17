/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"strings"

	"abate/kernel/algebra"
)

// MutationPoint is one discovered mutant of a single synthetic node.
//
// 参考 PITest 的 MutationPoint（字节码位置 + 算子 + 优先级），改进点=节点引用改为
// 语言无关的字符串，优先级改为整数覆盖数，并附加种子派生的 Selector 与确定性指纹。
// The fingerprint covers every field, so two runs of one seed cannot disagree.
type MutationPoint struct { //nolint:revive // PLAN-01 §4 fixes this domain name.
	// Node is the language-independent node reference.
	Node string

	// Original is the lexical token before mutation.
	Original string

	// Mutated is the lexical token after applying the operator.
	Mutated string

	// Operator is the M-family algebra operator applied to the node.
	Operator algebra.Operator

	// Priority is the integer coverage count of the node.
	Priority uint64

	// Selector is the SIM-RNG draw that seeds deterministic point selection.
	Selector uint64

	// Fingerprint is the SHA-256 digest of the canonical point fields.
	Fingerprint [32]byte
}

// NewMutationPoint builds a point and derives its deterministic fingerprint.
//
// A nil operator is tolerated and hashed as a zero fingerprint, which lets the
// caller validate malformed points before the engine dispatches them.
func NewMutationPoint(node, original, mutated string, op algebra.Operator,
	priority, selector uint64,
) MutationPoint {
	var b strings.Builder
	_, _ = b.WriteString(node)
	_ = b.WriteByte(0)
	_, _ = b.WriteString(original)
	_ = b.WriteByte(0)
	_, _ = b.WriteString(mutated)
	_ = b.WriteByte(0)
	if op != nil {
		fp := op.Fingerprint()
		_, _ = b.Write(fp[:])
	} else {
		_, _ = b.Write(make([]byte, sha256.Size))
	}
	_, _ = b.WriteString(strconv.FormatUint(priority, 10))
	_ = b.WriteByte(0)
	_, _ = b.WriteString(strconv.FormatUint(selector, 10))
	return MutationPoint{
		Node: node, Original: original, Mutated: mutated, Operator: op,
		Priority: priority, Selector: selector,
		Fingerprint: sha256.Sum256([]byte(b.String())),
	}
}

// Less reports the coverage-guided total order of two points.
//
// 参考 PITest 线程优先级调度，改进点=全序裁决：覆盖数降序优先，其后依次按
// 算子 ID、节点引用、变异记号、种子选择子、指纹字典序比较，杜绝并发序泄漏。
// A nil operator sorts before every registered operator.
func (p MutationPoint) Less(other MutationPoint) bool {
	if p.Priority != other.Priority {
		return p.Priority > other.Priority
	}
	if p.operatorKey() != other.operatorKey() {
		return p.operatorKey() < other.operatorKey()
	}
	if p.Node != other.Node {
		return p.Node < other.Node
	}
	if p.Mutated != other.Mutated {
		return p.Mutated < other.Mutated
	}
	if p.Selector != other.Selector {
		return p.Selector < other.Selector
	}
	return bytes.Compare(p.Fingerprint[:], other.Fingerprint[:]) < 0
}

// operatorKey returns the operator identifier, or an empty string for nil.
func (p MutationPoint) operatorKey() string {
	if p.Operator == nil {
		return ""
	}
	return string(p.Operator.ID())
}
