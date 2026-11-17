/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/json"
	"sort"

	"abate/lang/rule"
)

// oasisSchemaURI is the canonical OASIS SARIF 2.1.0 schema URI embedded in every
// generated document. The value is a valid absolute URI so it passes the schema's
// "uri" format constraint.
const oasisSchemaURI = "https://docs.oasis-open.org/sarif/sarif/v2.1.0/cos02/schemas/sarif-schema-2.1.0.json"

// sarifDocument is the root SARIF 2.1.0 log object.
type sarifDocument struct {
	Schema  string     `json:"$schema"`
	Version string     `json:"version"`
	Runs    []sarifRun `json:"runs"`
}

// sarifRun is a single run produced by one tool driver.
type sarifRun struct {
	Tool    sarifTool     `json:"tool"`
	Results []sarifResult `json:"results"`
}

// sarifTool wraps the driver that produced the results.
type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

// sarifDriver identifies the Abate reporter.
type sarifDriver struct {
	Name           string      `json:"name"`
	Version        string      `json:"version,omitempty"`
	InformationURI string      `json:"informationUri,omitempty"`
	Rules          []sarifRule `json:"rules,omitempty"`
}

// sarifRule is the descriptor for one reported rule.
type sarifRule struct {
	ID               string       `json:"id"`
	Name             string       `json:"name,omitempty"`
	ShortDescription sarifMessage `json:"shortDescription,omitempty"`
}

// sarifMessage carries human-readable text.
type sarifMessage struct {
	Text string `json:"text"`
}

// sarifResult is one SARIF result (a single finding).
type sarifResult struct {
	RuleID    string          `json:"ruleId,omitempty"`
	Level     string          `json:"level"`
	Kind      string          `json:"kind"`
	Message   sarifMessage    `json:"message"`
	Locations []sarifLocation `json:"locations"`
}

// sarifLocation associates a result with a source location.
type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

// sarifPhysicalLocation holds the artifact and region of a location.
type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           sarifRegion           `json:"region,omitempty"`
}

// sarifArtifactLocation names the source artifact.
type sarifArtifactLocation struct {
	URI string `json:"uri"`
}

// sarifRegion pinpoints a line/column in the artifact.
type sarifRegion struct {
	StartLine   int `json:"startLine,omitempty"`
	StartColumn int `json:"startColumn,omitempty"`
}

// BuildSARIF renders findings into a deterministic SARIF 2.1.0 log. The document
// is byte-stable for identical inputs: findings are sorted by a canonical key and
// the embedded JSON preserves a fixed field order, so the same findings always
// produce identical bytes. The returned bytes satisfy the OASIS 2.1.0 schema and
// can be validated with ValidateSARIFAgainstOASIS.
func BuildSARIF(findings []rule.Finding) ([]byte, error) {
	sorted := append([]rule.Finding(nil), findings...)
	sort.Slice(sorted, func(i, j int) bool {
		return findingLess(sorted[i], sorted[j])
	})

	rules, seen := make([]sarifRule, 0, len(sorted)), map[string]bool{}
	results := make([]sarifResult, 0, len(sorted))
	for _, f := range sorted {
		id := string(f.RuleID)
		if !seen[id] {
			seen[id] = true
			rules = append(rules, sarifRule{
				ID:               id,
				Name:             id,
				ShortDescription: sarifMessage{Text: "Abate static rule " + id},
			})
		}
		results = append(results, sarifResult{
			RuleID:  id,
			Level:   sarifLevel(f.Severity),
			Kind:    sarifKind(f.Severity),
			Message: sarifMessage{Text: f.Message},
			Locations: []sarifLocation{{
				PhysicalLocation: sarifPhysicalLocation{
					ArtifactLocation: sarifArtifactLocation{URI: f.Path},
					Region: sarifRegion{
						StartLine:   f.Line,
						StartColumn: f.Column,
					},
				},
			}},
		})
	}

	doc := sarifDocument{
		Schema:  oasisSchemaURI,
		Version: "2.1.0",
		Runs: []sarifRun{{
			Tool: sarifTool{Driver: sarifDriver{
				Name:           "Abate RPT",
				Version:        "0.1.0",
				InformationURI: "https://github.com/xiaolibuluo/Abate",
				Rules:          rules,
			}},
			Results: results,
		}},
	}
	return json.Marshal(doc)
}

// sarifLevel maps a finding severity to the SARIF result level enum.
func sarifLevel(s rule.Severity) string {
	switch s {
	case rule.SeverityError:
		return "error"
	case rule.SeverityWarning:
		return "warning"
	case rule.SeverityInfo:
		return "note"
	default:
		return "warning"
	}
}

// sarifKind maps a finding severity to the SARIF result kind enum.
func sarifKind(s rule.Severity) string {
	switch s {
	case rule.SeverityError:
		return "fail"
	case rule.SeverityWarning:
		return "review"
	case rule.SeverityInfo:
		return "informational"
	default:
		return "review"
	}
}

// findingLess provides the canonical total order used to sort findings so SARIF
// output is deterministic regardless of the caller's input order.
func findingLess(a, b rule.Finding) bool {
	if a.RuleID != b.RuleID {
		return a.RuleID < b.RuleID
	}
	if a.Path != b.Path {
		return a.Path < b.Path
	}
	if a.Line != b.Line {
		return a.Line < b.Line
	}
	if a.Column != b.Column {
		return a.Column < b.Column
	}
	return a.Message < b.Message
}
