/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"abate/ces/schema"
	"abate/internal/yamlmini"
)

// MaxDocumentBytes bounds a CES document to 10 MiB to prevent memory abuse.
const MaxDocumentBytes = 10 << 20

// Parse converts a CES byte stream into a strong-typed Experiment.
//
// Format auto sniffs JSON from a leading '{' and treats everything else as
// YAML. Both decoders reject unknown fields so a typo can never be ignored.
func Parse(data []byte, format Format) (*schema.Experiment, error) {
	if len(data) == 0 {
		return nil, newError(RuleEmptyDocument, "$", "empty document")
	}
	if len(data) > MaxDocumentBytes {
		return nil, newError(RuleInputTooLarge, "$",
			fmt.Sprintf("document exceeds %d bytes", MaxDocumentBytes))
	}
	f := format
	if f == FormatAuto {
		f = sniffFormat(data)
	}

	var exp schema.Experiment
	switch f {
	case FormatJSON:
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		if err := dec.Decode(&exp); err != nil {
			return nil, decodeFailure(err)
		}
	case FormatYAML:
		dec := yamlmini.NewDecoder(bytes.NewReader(data))
		dec.KnownFields(true)
		if err := dec.Decode(&exp); err != nil {
			return nil, decodeFailure(err)
		}
	default:
		return nil, newError(RuleInvalidFormat, "$", "unsupported format "+string(f))
	}
	if err := checkEnvelope(&exp); err != nil {
		return nil, err
	}
	return &exp, nil
}

// sniffFormat picks a decoder from the first non-space byte.
func sniffFormat(data []byte) Format {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) > 0 && trimmed[0] == '{' {
		return FormatJSON
	}
	return FormatYAML
}

// decodeFailure classifies a strict-decoder error into a rule identifier.
func decodeFailure(err error) *ParseError {
	msg := err.Error()
	rule := RuleMalformedDocument
	if strings.Contains(msg, "unknown field") || strings.Contains(msg, "not found in type") {
		rule = RuleUnknownField
	}
	return newError(rule, "$", msg)
}

// checkEnvelope validates the fixed document envelope fields.
func checkEnvelope(exp *schema.Experiment) error {
	if exp.APIVersion == "" {
		return newError(RuleMissingField, "$.apiVersion", "apiVersion is required")
	}
	if _, err := schema.NewVersion().Negotiate(exp.APIVersion); err != nil {
		return newError(RuleUnsupportedVersion, "$.apiVersion", err.Error())
	}
	if exp.Kind != schema.KindExperiment {
		return newError(RuleMalformedDocument, "$.kind", "kind must be Experiment")
	}
	if exp.Metadata.Name == "" {
		return newError(RuleMissingField, "$.metadata.name", "metadata.name is required")
	}
	return nil
}
