/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"slices"
	"strconv"

	"abate/kernel/algebra"
)

// chaosSpec describes one of the twelve C-family fault operators.
//
// 参考 chaosblade 的故障参数分类（网络/资源/时钟/进程/依赖），改进点=每个故障
// 直接实例化为 KRN-ALG 的 C 族 Operator，参数与补偿策略可枚举、指纹确定性。
type chaosSpec struct {
	name         string
	target       string
	value        int64
	compensation string
}

// Registry enumerates the twelve C-family fault operators deterministically.
//
// 参考 chaosblade 的 Model 注册表，改进点=去掉 CLI/exec/渠道胶水，收敛为纯库
// 形态的十二算子表，且每个算子带 Inverse 对偶与确定性指纹。
// The registry is immutable after construction and safe for concurrent reads.
type Registry struct {
	specs []chaosSpec
}

// NewRegistry builds the twelve-operator registry in a fixed declaration order.
func NewRegistry() *Registry {
	return &Registry{specs: []chaosSpec{
		{"network-latency", linkA + "-" + linkB, 9000, "recover-fault"},
		{"network-partition", linkA + "-" + linkB, 0, "recover-fault"},
		{"network-loss", linkA + "-" + linkB, 250000, "recover-fault"},
		{"network-bandwidth", linkA + "-" + linkB, 1000, "recover-fault"},
		{"disk-io-latency", filePath, 5000, "reset-latency"},
		{"disk-full", filePath, probeBytes, "recover-badblock"},
		{"cpu-pressure", cgroupName, 1, "recover-throttle"},
		{"memory-pressure", "task-1", baseMemLimit + 1, "rebuild-scheduler"},
		{"clock-offset", "clock", 5000, "inverse-offset"},
		{"process-kill", "task-1", 0, "rebuild-scheduler"},
		{"dependency-error", "svc", 0, "clear-flag"},
		{"dependency-timeout", "svc", 0, "clear-flag"},
	}}
}

// Classes returns the twelve fault class names in declaration order.
func (r *Registry) Classes() []string {
	out := make([]string, 0, len(r.specs))
	for _, spec := range r.specs {
		out = append(out, spec.name)
	}
	return out
}

// All returns the twelve registered C-family operators sorted by identifier.
func (r *Registry) All() []algebra.Operator {
	out := make([]algebra.Operator, 0, len(r.specs))
	for _, spec := range r.specs {
		out = append(out, buildOperator(spec, algebra.Params{}))
	}
	slices.SortFunc(out, byOperatorID)
	return out
}

// Contains reports whether an identifier belongs to the registered faults.
func (r *Registry) Contains(id algebra.OperatorID) bool {
	_, ok := r.specByID(id)
	return ok
}

// Variant rebuilds one operator with overridden parameters.
//
// 参考 chaosblade 的 ExpModel 参数注入，改进点=覆盖键先排序再合并，杜绝 map
// 迭代序泄漏；未知算子以 sentinel error 拒绝。
func (r *Registry) Variant(id algebra.OperatorID,
	overrides algebra.Params,
) (algebra.Operator, error) {
	spec, ok := r.specByID(id)
	if !ok {
		return nil, ErrUnknownOperator
	}
	return buildOperator(spec, overrides), nil
}

// specByID resolves one spec from its formatted operator identifier.
func (r *Registry) specByID(id algebra.OperatorID) (chaosSpec, bool) {
	for _, spec := range r.specs {
		if formatID(spec.name) == id {
			return spec, true
		}
	}
	return chaosSpec{}, false
}

// specByClass resolves one spec from its fault class name.
func (r *Registry) specByClass(class string) (chaosSpec, bool) {
	for _, spec := range r.specs {
		if spec.name == class {
			return spec, true
		}
	}
	return chaosSpec{}, false
}

// formatID builds the globally unique identifier of a fault class.
func formatID(name string) algebra.OperatorID {
	return algebra.FormatOperatorID(algebra.FamilyChaos, name, "1")
}

// buildOperator instantiates one C-family operator from its spec and overrides.
func buildOperator(spec chaosSpec, overrides algebra.Params) algebra.Operator {
	params := algebra.Params{
		"class":        spec.name,
		"target":       spec.target,
		"value":        strconv.FormatInt(spec.value, 10),
		"compensation": spec.compensation,
	}
	keys := make([]string, 0, len(overrides))
	for key := range overrides {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, key := range keys {
		params[key] = overrides[key]
	}
	return algebra.NewBasic(formatID(spec.name), algebra.FamilyChaos, params,
		[]algebra.TargetKey{algebra.TargetKey(params["target"])}, true)
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
