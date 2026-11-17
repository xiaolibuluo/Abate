/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

// VerifyResult is the full recomputation report for one manifest. Mismatches
// and SPDXDetails carry per-file localisers for the tenth gate segment.
type VerifyResult struct {
	Registered           int            `json:"registered"`
	Recomputed           int            `json:"recomputed"`
	SHA256Mismatches     int            `json:"sha256Mismatches"`
	SPDXViolations       int            `json:"spdxViolations"`
	BoundarySPDXCount    int            `json:"boundarySPDXCount"`
	ParseablePerLanguage map[string]int `json:"parseablePerLanguage"`
	Mismatches           []string       `json:"mismatches"`
	SPDXDetails          []string       `json:"spdxDetails"`
}

// newVerifyResult builds a result with all seven language counters present.
func newVerifyResult(registered int) VerifyResult {
	r := VerifyResult{
		Registered:           registered,
		ParseablePerLanguage: make(map[string]int, len(languageOrder)),
		Mismatches:           make([]string, 0),
		SPDXDetails:          make([]string, 0),
	}
	for _, lang := range languageOrder {
		r.ParseablePerLanguage[lang] = 0
	}
	return r
}
