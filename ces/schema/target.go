/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// TargetKind constants enumerate the supported target domains.
const (
	TargetSimulator = "simulator"
	TargetSource    = "source"
	TargetExternal  = "external"
)

// Target identifies what an experiment perturbs.
//
// Kind is one of simulator, source or external; external is a reserved field
// in CES v1 and is rejected by validation because no runtime support exists.
// Selector is a stable key/value map (package, function, node, process).
type Target struct {
	Kind     string            `json:"kind" yaml:"kind"`
	Selector map[string]string `json:"selector" yaml:"selector"`
}
