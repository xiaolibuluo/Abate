/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"encoding/hex"
	"math"
	"slices"
	"strconv"

	"abate/kernel/algebra"
)

// tokenDeleted marks a statement or call removed by a mutation operator.
const tokenDeleted = "<deleted>"

// classSpec is one of the seven M-family operator classes.
//
// 参考 PITest / go-gremlins 的算子注册模型（按 token 类别分派变异），
// 改进点=每个算子实例化为 algebra.Operator（M 族），参数与指纹确定性可枚举。
type classSpec struct {
	name           string
	representative string
	invertible     bool
	variants       func(token string) []string
}

// Registry enumerates the seven M-family operator classes deterministically.
//
// 参考 PITest 算子族注册表，改进点=全部算子由纯 Go 表驱动生成，实现 KRN-ALG 的
// Operator 契约，带确定性 Fingerprint；删除类算子显式声明不可逆。
// The registry is immutable after construction and safe for concurrent reads.
type Registry struct {
	specs   []classSpec
	names   []string
	derived []algebra.Operator
	byToken map[string][]algebra.Operator
}

// NewRegistry builds the seven-class registry in a fixed declaration order.
func NewRegistry() *Registry {
	specs := []classSpec{
		{"arith-replace", "+", true, arithVariants},
		{"boundary-value", "<", true, boundaryVariants},
		{"constant-perturb", "1", true, constantVariants},
		{"control-flow", "break", false, controlVariants},
		{"return-value", "true", true, returnVariants},
		{"negate-condition", "==", true, negateVariants},
		{"remove-call", "call", false, callVariants},
	}
	names := make([]string, 0, len(specs))
	for _, spec := range specs {
		names = append(names, spec.name)
	}
	return &Registry{specs: specs, names: names}
}

// Classes returns the seven class names in fixed order.
func (r *Registry) Classes() []string {
	return append([]string(nil), r.names...)
}

// All returns one representative operator per class plus every derived
// operator, sorted by identifier.
func (r *Registry) All() []algebra.Operator {
	out := make([]algebra.Operator, 0, len(r.specs)+len(r.derived))
	for _, spec := range r.specs {
		variants := spec.variants(spec.representative)
		if len(variants) == 0 {
			continue
		}
		out = append(out, buildOperator(spec, spec.representative, variants[0], 0))
	}
	out = append(out, r.derived...)
	slices.SortFunc(out, byOperatorID)
	return out
}

// For returns every operator applicable to one token, sorted by identifier.
func (r *Registry) For(token string) []algebra.Operator {
	if token == "" {
		return nil
	}
	out := make([]algebra.Operator, 0, 8)
	for _, spec := range r.specs {
		for index, mutated := range spec.variants(token) {
			out = append(out, buildOperator(spec, token, mutated, index))
		}
	}
	out = append(out, r.derivedFor(token)...)
	slices.SortFunc(out, byOperatorID)
	return out
}

// derivedFor returns a defensive copy of the derived operators bound to one
// token; the frozen index itself is never mutated after construction.
func (r *Registry) derivedFor(token string) []algebra.Operator {
	return append([]algebra.Operator(nil), r.byToken[token]...)
}

// Valid reports whether op belongs to one of the registered M-family classes.
func (r *Registry) Valid(op algebra.Operator) bool {
	if op == nil || op.Family() != algebra.FamilyMutation {
		return false
	}
	return slices.Contains(r.names, op.Params()["class"])
}

// byOperatorID orders operators by their globally unique identifier.
func byOperatorID(a, b algebra.Operator) int {
	if a.ID() < b.ID() {
		return -1
	}
	if a.ID() > b.ID() {
		return 1
	}
	return 0
}

// buildOperator instantiates one parameterized M-family algebra operator.
func buildOperator(spec classSpec, token, mutated string, index int) algebra.Operator {
	suffix := spec.name + "-" + hex.EncodeToString([]byte(token)) + "-" + strconv.Itoa(index)
	id := algebra.FormatOperatorID(algebra.FamilyMutation, suffix, "1")
	params := algebra.Params{
		"class": spec.name,
		"from":  token,
		"to":    mutated,
		"index": strconv.Itoa(index),
	}
	return algebra.NewBasic(id, algebra.FamilyMutation, params,
		[]algebra.TargetKey{algebra.TargetKey(token)}, spec.invertible)
}

// arithVariants swaps arithmetic and logical binary operators.
func arithVariants(token string) []string {
	switch token {
	case "+":
		return []string{"-"}
	case "-":
		return []string{"+"}
	case "*":
		return []string{"/"}
	case "/":
		return []string{"*"}
	case "&&":
		return []string{"||"}
	case "||":
		return []string{"&&"}
	default:
		return nil
	}
}

// boundaryVariants shifts strict comparisons across their boundary.
func boundaryVariants(token string) []string {
	switch token {
	case "<":
		return []string{"<="}
	case "<=":
		return []string{"<"}
	case ">":
		return []string{">="}
	case ">=":
		return []string{">"}
	default:
		return nil
	}
}

// constantVariants perturbs an integer constant by +1, -1 and zero.
func constantVariants(token string) []string {
	n, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return nil
	}
	out := make([]string, 0, 3)
	if n < math.MaxInt64 {
		out = append(out, strconv.FormatInt(n+1, 10))
	}
	if n > math.MinInt64 {
		out = append(out, strconv.FormatInt(n-1, 10))
	}
	out = append(out, "0")
	return out
}

// controlVariants deletes a break or continue statement.
func controlVariants(token string) []string {
	if token == "break" || token == "continue" {
		return []string{tokenDeleted}
	}
	return nil
}

// returnVariants flips boolean and nil return values.
func returnVariants(token string) []string {
	switch token {
	case "true":
		return []string{"false"}
	case "false":
		return []string{"true"}
	case "nil":
		return []string{"nonnil"}
	case "nonnil":
		return []string{"nil"}
	default:
		return nil
	}
}

// negateVariants inverts a comparison or boolean condition.
func negateVariants(token string) []string {
	switch token {
	case "==":
		return []string{"!="}
	case "!=":
		return []string{"=="}
	case "<":
		return []string{">="}
	case ">":
		return []string{"<="}
	case "<=":
		return []string{">"}
	case ">=":
		return []string{"<"}
	default:
		return nil
	}
}

// callVariants deletes a call expression.
func callVariants(token string) []string {
	if token == "call" {
		return []string{tokenDeleted}
	}
	return nil
}
