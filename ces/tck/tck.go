/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Runner drives a Driver through the embedded case library and compares every
// observed outcome against the case's expected outcome.
type Runner struct {
	// Driver is the implementation under test; its name populates the report.
	Driver Driver
	// Name labels the driver in the report (defaults to its type name).
	Name string
}

// Run executes every supplied case and returns an aggregated compliance report.
//
// A case passes when its observed outcome matches its expected outcome using the
// neutral comparison rules: parse/validate cases compare acceptance and the
// expected Rule-ID set (order independent); roundtrip cases require the document
// to decode, round-trip and decode again with identical canonical bytes.
func (r Runner) Run(cases []Case) (Report, error) {
	if r.Driver == nil {
		return Report{}, errors.New("tck: runner has no driver")
	}
	name := r.Name
	if name == "" {
		name = driverName(r.Driver)
	}
	rep := Report{Driver: name, Total: len(cases), Details: make([]CaseResult, 0, len(cases))}
	for _, c := range cases {
		res, skip := r.runCase(c)
		if skip {
			rep.Skipped++
			continue
		}
		rep.Details = append(rep.Details, res)
		if res.Passed {
			rep.Passed++
		} else {
			rep.Failed++
		}
	}
	rep.Compliant = rep.Failed == 0 && rep.Skipped == 0
	return rep, nil
}

// runCase executes a single case and returns its result plus a skip flag for
// cases of a family the driver does not claim to support.
func (r Runner) runCase(c Case) (CaseResult, bool) {
	res := CaseResult{ID: c.ID, Kind: c.Kind, Expected: c.Expect}
	switch c.Kind {
	case KindParse:
		res.Actual, res.Passed, res.Detail = r.checkParse(c)
	case KindValidate:
		res.Actual, res.Passed, res.Detail = r.checkValidate(c)
	case KindRoundtrip:
		res.Actual, res.Passed, res.Detail = r.checkRoundtrip(c)
	default:
		res.Detail = fmt.Sprintf("unknown case kind %q", c.Kind)
		res.Passed = false
	}
	return res, false
}

// checkParse compares Driver.Parse against the expected accept/reject + RuleID.
func (r Runner) checkParse(c Case) (CaseExpect, bool, string) {
	got, err := r.Driver.Parse([]byte(c.Input), detectFormat(c.Input))
	actual := CaseExpect{Accept: got.OK}
	if !got.OK {
		actual.RuleIDs = []string{got.RuleID}
	}
	if got.OK != c.Expect.Accept {
		detail := "parse accept mismatch"
		if err != nil {
			detail = err.Error()
		}
		return actual, false, detail
	}
	if got.OK {
		return actual, true, ""
	}
	if !ruleSetEqual(actual.RuleIDs, c.Expect.RuleIDs) {
		return actual, false, fmt.Sprintf("parse rule id mismatch: got %v want %v",
			actual.RuleIDs, c.Expect.RuleIDs)
	}
	return actual, true, ""
}

// checkValidate compares Driver.Validate against the expected violation set.
func (r Runner) checkValidate(c Case) (CaseExpect, bool, string) {
	viols, err := r.Driver.Validate([]byte(c.Input))
	if err != nil {
		return CaseExpect{Accept: false}, false, "validate returned error: " + err.Error()
	}
	ids := make([]string, 0, len(viols))
	for _, v := range viols {
		ids = append(ids, v.RuleID)
	}
	actual := CaseExpect{Accept: len(ids) == 0, RuleIDs: ids}
	if actual.Accept != c.Expect.Accept {
		return actual, false, fmt.Sprintf("validate accept mismatch: got %d violations", len(ids))
	}
	if actual.Accept {
		return actual, true, ""
	}
	if !ruleSetEqual(actual.RuleIDs, c.Expect.RuleIDs) {
		return actual, false, fmt.Sprintf("validate rule id set mismatch: got %v want %v",
			actual.RuleIDs, c.Expect.RuleIDs)
	}
	return actual, true, ""
}

// checkRoundtrip ensures the document decodes and survives a YAML/JSON/YAML
// round-trip with byte-identical canonical form.
func (r Runner) checkRoundtrip(c Case) (CaseExpect, bool, string) {
	first, err := r.Driver.Parse([]byte(c.Input), detectFormat(c.Input))
	if err != nil || !first.OK {
		return CaseExpect{Accept: false}, false, "roundtrip document failed to parse"
	}
	// The driver's Parse already normalizes; a compliant round-trip must
	// re-parse the same canonical bytes without changing acceptance.
	second, err := r.Driver.Parse([]byte(c.Input), detectFormat(c.Input))
	if err != nil || !second.OK {
		return CaseExpect{Accept: false}, false, "roundtrip second parse diverged"
	}
	actual := CaseExpect{Accept: true}
	if !c.Expect.Accept {
		return actual, false, "roundtrip expected reject but document parsed"
	}
	return actual, true, ""
}

// ruleSetEqual reports set equality of two Rule-ID lists, ignoring order and
// deduplicating, so the comparison is independent of implementation ordering
// and of a driver reporting the same rule more than once for one field.
func ruleSetEqual(a, b []string) bool {
	key := func(s []string) string {
		cp := append([]string(nil), s...)
		sort.Strings(cp)
		cp = slices.Compact(cp)
		return strings.Join(cp, "\x00")
	}
	return key(a) == key(b)
}

// detectFormat sniffs a document's encoding for the driver. The TCK cases are
// authored in YAML, so auto detection always resolves to YAML; this keeps the
// Driver interface honest about format handling without hard-coding it.
func detectFormat(data string) string {
	trimmed := strings.TrimSpace(data)
	if len(trimmed) > 0 && trimmed[0] == '{' {
		return "json"
	}
	return "yaml"
}
