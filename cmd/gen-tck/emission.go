/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"errors"
	"fmt"
	"strings"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/ces/validate"
	"abate/internal/yamlmini"
)

type caseFile struct {
	id   string
	kind string
	body string
}

// outcome is the neutral pipeline result for a scenario.
type outcome struct {
	accept  bool
	ruleIDs []string
	// stage is where the document first failed: "parse", "validate" or "".
	stage string
}

func pipeline(e schema.Experiment) outcome {
	data, err := yamlmini.Marshal(e)
	if err != nil {
		return outcome{accept: false, ruleIDs: []string{"CES-PARSE-005"}, stage: "parse"}
	}
	exp, perr := parse.Parse(data, parse.FormatYAML)
	if perr != nil {
		return outcome{accept: false, ruleIDs: []string{ruleIDOf(perr)}, stage: "parse"}
	}
	norm, nerr := parse.Normalize(exp)
	if nerr != nil {
		return outcome{accept: false, ruleIDs: []string{ruleIDOf(nerr)}, stage: "parse"}
	}
	viols, verr := validate.Validate(norm)
	if verr != nil {
		return outcome{accept: false, ruleIDs: []string{"CES-VALID-INTERNAL"}, stage: "validate"}
	}
	if len(viols) == 0 {
		return outcome{accept: true}
	}
	ids := make([]string, 0, len(viols))
	for _, v := range viols {
		ids = append(ids, v.RuleID)
	}
	return outcome{accept: false, ruleIDs: ids, stage: "validate"}
}

func pipelineRaw(raw string) outcome {
	exp, perr := parse.Parse([]byte(raw), parse.FormatYAML)
	if perr != nil {
		return outcome{accept: false, ruleIDs: []string{ruleIDOf(perr)}, stage: "parse"}
	}
	if _, nerr := parse.Normalize(exp); nerr != nil {
		return outcome{accept: false, ruleIDs: []string{ruleIDOf(nerr)}, stage: "parse"}
	}
	return outcome{accept: true}
}

func ruleIDOf(err error) string {
	var pe *parse.ParseError
	if errors.As(err, &pe) {
		return pe.RuleID
	}
	return "CES-PARSE-005"
}

func buildCase(s scenario, get func(string) int) caseFile {
	var raw string
	var o outcome
	if s.raw != "" {
		raw = s.raw
		o = pipelineRaw(raw)
	} else {
		data, err := yamlmini.Marshal(s.e)
		if err != nil {
			panic(err)
		}
		raw = string(data)
		o = pipeline(s.e)
	}

	// Derive the case kind from where the document first fails.
	kind := ""
	switch {
	case o.stage == "parse":
		kind = "parse"
	case o.stage == "validate":
		kind = "validate"
	case s.cat == "roundtrip":
		kind = "roundtrip"
	case s.cat == "validate-ok":
		kind = "validate"
	default:
		kind = "parse"
	}

	var id string
	switch kind {
	case "parse":
		id = fmt.Sprintf("TCK-PARSE-%03d", get("parse"))
	case "validate":
		id = fmt.Sprintf("TCK-VALIDATE-%03d", get("validate"))
	default:
		id = fmt.Sprintf("TCK-ROUNDTRIP-%03d", get("roundtrip"))
	}

	doc := indent(raw)
	body := fmt.Sprintf("id: %s\nspec: %q\nkind: %s\ninput: |\n%s\nexpect:\n  accept: %t\n",
		id, s.spec, kind, doc, o.accept)
	if o.accept {
		body += "  ruleIds: []\n"
	} else if len(o.ruleIDs) == 1 {
		body += fmt.Sprintf("  ruleIds:\n    - %s\n", o.ruleIDs[0])
	} else {
		body += "  ruleIds:\n"
		for _, r := range o.ruleIDs {
			body += fmt.Sprintf("    - %s\n", r)
		}
	}
	return caseFile{id: id, kind: kind, body: body}
}

func indent(s string) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	for i, l := range lines {
		if l != "" {
			lines[i] = "  " + l
		}
	}
	return strings.Join(lines, "\n")
}
