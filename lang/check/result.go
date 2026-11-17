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
	"strings"

	"abate/lang/rule"
)

// Result is the complete, deterministic outcome of one check run.
type Result struct {
	// Findings is the deduplicated, canonically ordered finding list.
	Findings []rule.Finding
	// Summary holds the run statistics.
	Summary Summary
	// Fingerprint is the SHA-256 digest of the fixed-order result fields.
	Fingerprint [32]byte
	// Format is the default rendering format.
	Format string
}

// HasViolations reports whether the run produced at least one finding.
func (r Result) HasViolations() bool {
	return len(r.Findings) > 0
}

// FingerprintHex returns the lowercase hexadecimal digest of the result.
func (r Result) FingerprintHex() string {
	return hex.EncodeToString(r.Fingerprint[:])
}

// Canonical renders every result field in a fixed order.
func (r Result) Canonical() string {
	var builder strings.Builder
	_, _ = builder.WriteString(r.Summary.Canonical())
	_, _ = builder.WriteString("\x1e")
	_, _ = builder.WriteString(r.Format)
	for _, finding := range r.Findings {
		_, _ = builder.WriteString("\x1e")
		_, _ = builder.WriteString(finding.Canonical())
	}
	return builder.String()
}

// Render produces the report bytes for a format, falling back to console for
// an unknown format so output remains deterministic.
func (r Result) Render(format string) []byte {
	switch format {
	case FormatJSON:
		return renderJSON(r)
	case FormatSARIF:
		return renderSARIF(r)
	default:
		return renderConsole(r)
	}
}

// computeFingerprint hashes the fixed-order result fields.
func computeFingerprint(findings []rule.Finding, summary Summary,
	format string) [32]byte {
	partial := Result{Findings: findings, Summary: summary, Format: format}
	return sha256.Sum256([]byte(partial.Canonical()))
}
