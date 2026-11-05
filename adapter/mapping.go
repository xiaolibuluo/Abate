// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package adapter 实现 CES-7.2 第三方混沌平台适配器。
//
// Package adapter implements the CES-7.2 third-party chaos platform adapter.
package adapter

import (
	"fmt"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/hub"
)

// MappedFault 是从 CES 编码翻译到上游平台格式的故障描述。
// MappedFault is a fault translated from CES encoding to upstream format.
type MappedFault struct {
	// UpstreamType 上游平台的故障类型标识。
	UpstreamType string
	// Target 作用目标（CES-2.2.1）。
	Target core.Target
	// Params 翻译后的参数集合。
	Params map[string]string
}

// Backend 是第三方混沌平台须实现的接口（CES-7.2.3）。
// Backend is the interface third-party chaos platforms must implement.
type Backend interface {
	// Apply 在上游平台创建故障，返回上游句柄。
	Apply(mapped MappedFault) (upstreamID string, err error)
	// Cleanup 清除上游平台故障；必须可重复调用。
	Cleanup(upstreamID string) error
}

// FaultMapping 是一条 CES 子编码到上游标识的映射（CES-7.2.2）。
// FaultMapping maps a CES sub-code to an upstream identifier.
type FaultMapping struct {
	// SubCode CES 子编码（如 F1.1）。
	SubCode string
	// UpstreamType 上游平台故障类型标识。
	UpstreamType string
	// ParamMap CES 参数名 → 上游参数名的映射。
	ParamMap map[string]string
}

// Degradation 记录一条显式降级（CES-7.2.2）。
// Degradation records an explicit capability degradation.
type Degradation struct {
	// SubCode 无法完整映射的 CES 子编码。
	SubCode string
	// Reason 降级原因。
	Reason string
}

// MappingTable 持有全部映射与降级记录。
// MappingTable holds all mappings and degradation records.
type MappingTable struct {
	mappings     map[string]FaultMapping
	degradations []Degradation
}

// NewMappingTable 从映射列表构造映射表，自动推导降级列表。
// NewMappingTable builds a mapping table from the given mappings; sub-codes
// not present in the list are recorded as degradations.
func NewMappingTable(entries []FaultMapping) *MappingTable {
	m := make(map[string]FaultMapping, len(entries))
	for _, e := range entries {
		m[e.SubCode] = e
	}
	allSubCodes := allChaosHubSubCodes()
	var degs []Degradation
	for _, sc := range allSubCodes {
		if _, ok := m[sc]; !ok {
			degs = append(degs, Degradation{SubCode: sc, Reason: "no upstream mapping (CES-7.2.2)"})
		}
	}
	return &MappingTable{mappings: m, degradations: degs}
}

// Lookup 查找子编码的映射；未找到返回错误。
// Lookup returns the mapping for a sub-code or an error.
func (mt *MappingTable) Lookup(subCode string) (FaultMapping, error) {
	fm, ok := mt.mappings[subCode]
	if !ok {
		return FaultMapping{}, fmt.Errorf("adapter: sub-code %q not mapped (CES-7.2.2 degradation)", subCode)
	}
	return fm, nil
}

// MappedClasses 返回有完整映射的故障分类编码（去重、字典序）。
// MappedClasses returns the deduplicated, sorted fault classes with mappings.
func (mt *MappingTable) MappedClasses() []core.FaultClass {
	seen := make(map[core.FaultClass]bool)
	for sc := range mt.mappings {
		if len(sc) >= 2 {
			fc := core.FaultClass(sc[:2])
			seen[fc] = true
		}
	}
	result := make([]core.FaultClass, 0, len(seen))
	for fc := range seen {
		result = append(result, fc)
	}
	return result
}

// Degradations 返回降级记录列表。
// Degradations returns the degradation records.
func (mt *MappingTable) Degradations() []Degradation {
	return mt.degradations
}

// Translate 将 CES Fault 翻译为 MappedFault。
// Translate converts a CES Fault into a MappedFault.
func (mt *MappingTable) Translate(f core.Fault, t core.Target) (MappedFault, error) {
	subCode := string(f.Class)
	fm, err := mt.Lookup(subCode)
	if err != nil {
		return MappedFault{}, err
	}
	translated := make(map[string]string, len(f.Params))
	for cesKey, val := range f.Params {
		upstreamKey, ok := fm.ParamMap[cesKey]
		if !ok {
			upstreamKey = cesKey
		}
		translated[upstreamKey] = val
	}
	return MappedFault{
		UpstreamType: fm.UpstreamType,
		Target:       t,
		Params:       translated,
	}, nil
}

// allChaosHubSubCodes 返回附录 B 种子表全部 20 项子编码，委托 hub 包。
func allChaosHubSubCodes() []string {
	return hub.NewSeedRegistry().AllCodeStrings()
}

// DefaultLitmusMappings 返回 LitmusChaos 适配器的种子映射表。
// DefaultLitmusMappings returns seed mappings for a LitmusChaos adapter.
func DefaultLitmusMappings() []FaultMapping {
	return []FaultMapping{
		{SubCode: "F1.1", UpstreamType: "http-latency", ParamMap: map[string]string{"delay": "latency"}},
		{SubCode: "F1.2", UpstreamType: "http-latency", ParamMap: map[string]string{"delay": "latency", "jitter": "jitter"}},
		{SubCode: "F1.3", UpstreamType: "http-latency", ParamMap: map[string]string{"timeoutMs": "timeout"}},
		{SubCode: "F2.1", UpstreamType: "pod-delete", ParamMap: map[string]string{"signal": "signal"}},
		{SubCode: "F2.2", UpstreamType: "pod-restart", ParamMap: map[string]string{"signal": "signal", "gracePeriodMs": "gracePeriod"}},
		{SubCode: "F3.1", UpstreamType: "network-loss", ParamMap: map[string]string{"loss": "loss", "direction": "direction"}},
		{SubCode: "F3.3", UpstreamType: "network-partition", ParamMap: map[string]string{"peer": "target", "direction": "direction"}},
		{SubCode: "F4.1", UpstreamType: "cpu-hog", ParamMap: map[string]string{"utilization": "load"}},
		{SubCode: "F4.2", UpstreamType: "memory-hog", ParamMap: map[string]string{"utilization": "load", "fillRate": "workers"}},
		{SubCode: "F4.3", UpstreamType: "disk-fill", ParamMap: map[string]string{"utilization": "size"}},
	}
}

// DefaultGremlinMappings 返回 Gremlin 适配器的种子映射表。
// DefaultGremlinMappings returns seed mappings for a Gremlin adapter.
func DefaultGremlinMappings() []FaultMapping {
	return []FaultMapping{
		{SubCode: "F1.1", UpstreamType: "latency/delay", ParamMap: map[string]string{"delay": "ms"}},
		{SubCode: "F1.2", UpstreamType: "latency/jitter", ParamMap: map[string]string{"delay": "ms", "jitter": "jitter"}},
		{SubCode: "F2.1", UpstreamType: "process/signal", ParamMap: map[string]string{"signal": "signal"}},
		{SubCode: "F2.2", UpstreamType: "process/restart", ParamMap: map[string]string{"signal": "signal", "gracePeriodMs": "delay"}},
		{SubCode: "F3.1", UpstreamType: "network/loss", ParamMap: map[string]string{"loss": "percent"}},
		{SubCode: "F3.3", UpstreamType: "network/blackhole", ParamMap: map[string]string{"peer": "ip"}},
		{SubCode: "F4.1", UpstreamType: "cpu/stress", ParamMap: map[string]string{"utilization": "load"}},
		{SubCode: "F4.2", UpstreamType: "memory/stress", ParamMap: map[string]string{"utilization": "megabytes"}},
		{SubCode: "F4.3", UpstreamType: "disk/fill", ParamMap: map[string]string{"utilization": "percent"}},
		{SubCode: "F6.1", UpstreamType: "time/travel", ParamMap: map[string]string{"offsetMs": "offset"}},
	}
}

// DefaultChaosMeshMappings 返回 Chaos Mesh 适配器的种子映射表。
// DefaultChaosMeshMappings returns seed mappings for a Chaos Mesh adapter.
func DefaultChaosMeshMappings() []FaultMapping {
	return []FaultMapping{
		{SubCode: "F1.1", UpstreamType: "IOChaos/delay", ParamMap: map[string]string{"delay": "latency"}},
		{SubCode: "F1.2", UpstreamType: "IOChaos/delay", ParamMap: map[string]string{"delay": "latency", "jitter": "jitter"}},
		{SubCode: "F1.3", UpstreamType: "HTTPChaos/abort", ParamMap: map[string]string{"timeoutMs": "delay"}},
		{SubCode: "F2.1", UpstreamType: "PodChaos/pod-kill", ParamMap: map[string]string{"signal": "signal"}},
		{SubCode: "F2.2", UpstreamType: "PodChaos/container-kill", ParamMap: map[string]string{"signal": "signal", "gracePeriodMs": "gracePeriod"}},
		{SubCode: "F3.1", UpstreamType: "NetworkChaos/loss", ParamMap: map[string]string{"loss": "loss", "direction": "direction"}},
		{SubCode: "F3.3", UpstreamType: "NetworkChaos/partition", ParamMap: map[string]string{"peer": "target", "direction": "direction"}},
		{SubCode: "F4.1", UpstreamType: "StressChaos/cpu", ParamMap: map[string]string{"utilization": "load"}},
		{SubCode: "F4.2", UpstreamType: "StressChaos/mem", ParamMap: map[string]string{"utilization": "load", "fillRate": "workers"}},
		{SubCode: "F4.3", UpstreamType: "IOChaos/fill", ParamMap: map[string]string{"utilization": "size"}},
		{SubCode: "F6.1", UpstreamType: "TimeChaos/skew", ParamMap: map[string]string{"offsetMs": "offset"}},
	}
}