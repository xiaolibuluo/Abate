/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Experiment is the top-level CES document.
//
// APIVersion and Kind are pinned by the version negotiator; the executable
// content lives in Spec. CES v1 rejects unknown fields so that a typo can
// never be silently ignored.
type Experiment struct {
	APIVersion string         `json:"apiVersion" yaml:"apiVersion"`
	Kind       string         `json:"kind" yaml:"kind"`
	Metadata   Metadata       `json:"metadata" yaml:"metadata"`
	Spec       ExperimentSpec `json:"spec" yaml:"spec"`
}
