/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// CaseExpect is the neutral expected outcome of a case.
type CaseExpect struct {
	// Accept is true when the implementation should accept the document.
	Accept bool `yaml:"accept"`
	// RuleIDs are the expected CES rule identifiers on rejection.
	//
	// For parse cases exactly one RuleID is expected; for validate cases the
	// set must match exactly (order independent).
	RuleIDs []string `yaml:"ruleIds"`
}
