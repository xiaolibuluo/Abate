/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"strings"
	"unicode"
)

// SPDX classification outcomes returned by ClassifySPDX.
const (
	spdxAllowlisted = "allowlisted"
	spdxBoundary    = "boundary"
	spdxGPL         = "GPL"
	spdxUNKNOWN     = "UNKNOWN"
)

// ClassifySPDX classifies an upstream licence expression verbatim. A compound
// expression is permissive only when every licence token is known permissive;
// copyleft and boundary markers take precedence over token checks.
func ClassifySPDX(expr string) string {
	upper := strings.ToUpper(strings.TrimSpace(expr))
	if upper == "" {
		return spdxUNKNOWN
	}
	if containsMarker(upper, gplSPDXMarkers) {
		return spdxGPL
	}
	if containsMarker(upper, boundarySPDXMarkers) {
		return spdxBoundary
	}
	tokens := spdxTokens(upper)
	if len(tokens) == 0 {
		return spdxUNKNOWN
	}
	for _, token := range tokens {
		if !allowlistedSPDX[token] {
			return spdxUNKNOWN
		}
	}
	return spdxAllowlisted
}

// containsMarker reports whether the expression contains any listed marker.
func containsMarker(upper string, markers []string) bool {
	for _, marker := range markers {
		if strings.Contains(upper, marker) {
			return true
		}
	}
	return false
}

// spdxTokens splits an expression into licence identifiers, discarding
// connectors and punctuation so dual-licence forms remain verifiable.
func spdxTokens(upper string) []string {
	fields := strings.FieldsFunc(upper, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '.' && r != '_'
	})
	out := make([]string, 0, len(fields))
	for _, field := range fields {
		if !spdxConnectors[field] {
			out = append(out, field)
		}
	}
	return out
}
