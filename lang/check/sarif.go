/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"abate/lang/rule"
)

// sarifVersion is the fixed SARIF schema version rendered by this build.
const sarifVersion = "2.1.0"

// sarifSchema is the canonical SARIF 2.1.0 JSON schema identifier.
const sarifSchema = "https://json.schemastore.org/sarif-2.1.0.json"

// sarifText is the SARIF message object.
type sarifText struct {
	Text string `json:"text"`
}

// sarifRuleProperties carries Abate rule metadata in a fixed field order.
type sarifRuleProperties struct {
	Category  string `json:"category"`
	Severity  string `json:"severity"`
	Reference string `json:"reference"`
}

// sarifRuleDesc describes one rule in the tool driver.
type sarifRuleDesc struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	ShortDescription sarifText           `json:"shortDescription"`
	Properties       sarifRuleProperties `json:"properties"`
}

// sarifDriver is the SARIF tool driver component.
type sarifDriver struct {
	Name           string          `json:"name"`
	Version        string          `json:"version"`
	InformationURI string          `json:"informationUri"`
	Rules          []sarifRuleDesc `json:"rules"`
}

// sarifTool wraps the driver.
type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

// sarifArtifact identifies the source file of a result.
type sarifArtifact struct {
	URI string `json:"uri"`
}

// sarifRegion locates a result inside its source file.
type sarifRegion struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}

// sarifPhysicalLocation combines the artifact and region.
type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifact `json:"artifactLocation"`
	Region           sarifRegion   `json:"region"`
}

// sarifLocation is the SARIF location wrapper.
type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

// sarifFingerprints holds the deterministic per-finding digest.
type sarifFingerprints struct {
	Abate string `json:"abateFinding/v1"`
}

// sarifResult is one reported match.
type sarifResult struct {
	RuleID              string            `json:"ruleId"`
	RuleIndex           int               `json:"ruleIndex"`
	Level               string            `json:"level"`
	Message             sarifText         `json:"message"`
	Locations           []sarifLocation   `json:"locations"`
	PartialFingerprints sarifFingerprints `json:"partialFingerprints"`
}

// sarifRun is one tool invocation with its results.
type sarifRun struct {
	Tool    sarifTool     `json:"tool"`
	Results []sarifResult `json:"results"`
}

// sarifReport is the minimal, deterministic SARIF 2.1.0 document.
type sarifReport struct {
	Schema  string     `json:"$schema"`
	Version string     `json:"version"`
	Runs    []sarifRun `json:"runs"`
}

// renderSARIF produces a minimal SARIF 2.1.0 document.
//
// Field order is fixed by struct declaration order, no timestamps are emitted,
// and every result carries a deterministic SHA-256 partial fingerprint. This
// is a self-contained M3 renderer; validation with the official SARIF
// validator is an M4 deliverable.
func renderSARIF(r Result) []byte {
	ids := r.Summary.SortedRules()
	index := make(map[string]int, len(ids))
	rules := make([]sarifRuleDesc, 0, len(ids))
	for i, id := range ids {
		index[id] = i
		rules = append(rules, ruleDescriptor(r.Findings, id))
	}
	results := make([]sarifResult, 0, len(r.Findings))
	for _, finding := range r.Findings {
		results = append(results,
			sarifFinding(finding, index[string(finding.RuleID)]))
	}
	report := sarifReport{
		Schema: sarifSchema, Version: sarifVersion,
		Runs: []sarifRun{{
			Tool: sarifTool{Driver: sarifDriver{
				Name: "Abate", Version: "3.0.0",
				InformationURI: "https://github.com/xiaolibuluo/Abate",
				Rules:          rules,
			}},
			Results: results,
		}},
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return renderConsole(r)
	}
	return append(data, '\n')
}

// ruleDescriptor builds the driver entry for one rule identifier.
func ruleDescriptor(findings []rule.Finding, id string) sarifRuleDesc {
	out := sarifRuleDesc{ID: id, Name: id}
	for _, finding := range findings {
		if string(finding.RuleID) != id {
			continue
		}
		out.ShortDescription = sarifText{Text: finding.Message}
		out.Properties = sarifRuleProperties{
			Category:  finding.Category.String(),
			Severity:  finding.Severity.String(),
			Reference: finding.Reference,
		}
		return out
	}
	return out
}

// sarifFinding projects one finding into a SARIF result.
func sarifFinding(finding rule.Finding, index int) sarifResult {
	line, column := finding.Line, finding.Column
	if line < 1 {
		line = 1
	}
	if column < 1 {
		column = 1
	}
	digest := sha256.Sum256([]byte(finding.Canonical()))
	return sarifResult{
		RuleID: string(finding.RuleID), RuleIndex: index,
		Level:   sarifLevel(finding.Severity),
		Message: sarifText{Text: finding.Message},
		Locations: []sarifLocation{{PhysicalLocation: sarifPhysicalLocation{
			ArtifactLocation: sarifArtifact{URI: finding.Path},
			Region:           sarifRegion{StartLine: line, StartColumn: column},
		}}},
		PartialFingerprints: sarifFingerprints{
			Abate: hex.EncodeToString(digest[:]),
		},
	}
}

// sarifLevel maps an Abate severity to a SARIF result level.
func sarifLevel(severity rule.Severity) string {
	switch severity {
	case rule.SeverityError:
		return "error"
	case rule.SeverityWarning:
		return "warning"
	default:
		return "note"
	}
}
