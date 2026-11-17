/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

// Violation is one deterministic rule failure of a CES document.
type Violation struct {
	RuleID    string   `json:"ruleId" yaml:"ruleId"`
	FieldPath string   `json:"fieldPath" yaml:"fieldPath"`
	Message   string   `json:"message" yaml:"message"`
	Severity  Severity `json:"severity" yaml:"severity"`
}
