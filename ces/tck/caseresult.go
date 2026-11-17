/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// CaseResult is the observed outcome of running one case.
type CaseResult struct {
	ID       string
	Kind     CaseKind
	Passed   bool
	Expected CaseExpect
	Actual   CaseExpect
	Detail   string
}
