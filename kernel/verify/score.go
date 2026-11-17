/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// Score summarizes a mutation run with the industry-standard denominator:
// equivalent mutants are excluded from the total.
type Score struct {
	Total      uint64
	Killed     uint64
	Survived   uint64
	Equivalent uint64
	Timeout    uint64
}

// Summarize counts a list of mutation verdicts into a score.
func Summarize(verdicts []MutationVerdict) Score {
	var s Score
	for _, v := range verdicts {
		s.Total++
		switch v {
		case MutantKilled:
			s.Killed++
		case MutantSurvived:
			s.Survived++
		case MutantEquivalent:
			s.Equivalent++
		case MutantTimeout:
			s.Timeout++
		}
	}
	return s
}

// Value returns Killed / (Total - Equivalent), or zero for an empty denominator.
func (s Score) Value() float64 {
	excluded := s.Equivalent
	if excluded > s.Total {
		excluded = s.Total
	}
	denom := s.Total - excluded
	if denom == 0 {
		return 0
	}
	return float64(s.Killed) / float64(denom)
}
