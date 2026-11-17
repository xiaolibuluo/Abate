/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// MutationVerdict is the fate of one mutant under the test suite.
type MutationVerdict string

// Mutation verdict values.
const (
	MutantKilled     MutationVerdict = "KILLED"
	MutantSurvived   MutationVerdict = "SURVIVED"
	MutantEquivalent MutationVerdict = "EQUIVALENT"
	MutantTimeout    MutationVerdict = "TIMEOUT"
)
