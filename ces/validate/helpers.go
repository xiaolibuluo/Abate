/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"fmt"
	"strings"
	"unicode"

	"abate/ces/schema"
)

// idRune mirrors the CES identifier alphabet used by the parser's lexer.
func idRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) ||
		r == '-' || r == '_' || r == '.'
}

// referencedIDs extracts every operator:// reference deterministically,
// in first-appearance order, without depending on a full expression parse.
func referencedIDs(expr string) []string {
	const marker = "operator://"
	var out []string
	for rest := expr; ; {
		at := strings.Index(rest, marker)
		if at < 0 {
			return out
		}
		tail := rest[at+len(marker):]
		end := strings.IndexFunc(tail, func(r rune) bool { return !idRune(r) })
		if end < 0 {
			end = len(tail)
		}
		if end > 0 {
			out = append(out, tail[:end])
		}
		rest = tail[end:]
	}
}

// path builds an indexed field path without depending on map iteration order.
func path(base string, i int) string {
	return fmt.Sprintf("%s[%d]", base, i)
}

// probeIndex returns the index of a probe id, or -1 when absent.
func probeIndex(exp *schema.Experiment, id string) int {
	for i := range exp.Spec.Probes {
		if exp.Spec.Probes[i].ID == id {
			return i
		}
	}
	return -1
}

// isSimulatorURL implements CES-VALID-PROBE-003 (the simulator iron rule):
// only relative paths and the in-simulator hosts are reachable; any other
// http(s) host is an external target and must be rejected.
func isSimulatorURL(raw string) bool {
	if raw == "" {
		return false
	}
	if strings.HasPrefix(raw, "/") || strings.HasPrefix(raw, "sim://") {
		return true
	}
	rest, ok := trimScheme(raw)
	if !ok {
		return false
	}
	host := rest
	if slash := strings.IndexByte(host, '/'); slash >= 0 {
		host = host[:slash]
	}
	if colon := strings.IndexByte(host, ':'); colon >= 0 {
		host = host[:colon]
	}
	switch host {
	case "sim", "localhost", "127.0.0.1", "::1":
		return true
	default:
		return false
	}
}

// trimScheme strips an http or https prefix from a URL.
func trimScheme(raw string) (string, bool) {
	for _, prefix := range []string{"http://", "https://"} {
		if strings.HasPrefix(raw, prefix) {
			return strings.TrimPrefix(raw, prefix), true
		}
	}
	return "", false
}

// isSafeShellCommand implements CES-VALID-PROBE-004: a command is accepted
// only when it has no shell metacharacters and its first token belongs to the
// read-only in-simulator shell subset.
func isSafeShellCommand(command string) bool {
	if strings.TrimSpace(command) == "" {
		return false
	}
	for _, meta := range []string{";", "|", "&", ">", "<", "`", "$", "\n", "(", ")"} {
		if strings.Contains(command, meta) {
			return false
		}
	}
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return false
	}
	return readOnlyCommand(fields[0])
}

// readOnlyCommand reports whether a token is an allowed read-only command.
func readOnlyCommand(token string) bool {
	switch token {
	case "echo", "true", "false", "printf", "test", "cat", "grep",
		"ls", "pwd", "date", "sleep", "wc", "head", "tail", "sort", "uniq":
		return true
	default:
		return false
	}
}

// validPredicate implements CES-VALID-STEADY-005: a predicate is a comparison
// operator followed by an integer or decimal literal, matching the verifier's
// predicate grammar (>5, >=5, <5, <=5, ==5, !=5, -3.5 accepted).
func validPredicate(s string) bool {
	if s == "" {
		return false
	}
	opLen := 0
	for _, op := range []string{">=", "<=", "==", "!=", ">", "<"} {
		if strings.HasPrefix(s, op) {
			opLen = len(op)
			break
		}
	}
	if opLen == 0 {
		return false
	}
	return validNumber(s[opLen:])
}

// validNumber reports whether s is a signed integer or decimal literal.
func validNumber(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	if s == "" {
		return false
	}
	dots := 0
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
		case s[i] == '.' && dots == 0:
			dots++
		default:
			return false
		}
	}
	return s != "."
}

// validCron implements CES-VALID-SCHEDULE-002: exactly five fields, each
// built from digits, '*', ',', '-' and '/', with numeric values inside the
// field's range.
func validCron(expression string) bool {
	fields := strings.Fields(expression)
	if len(fields) != 5 {
		return false
	}
	bounds := [5]int{59, 23, 31, 12, 7}
	for i, f := range fields {
		if !validCronField(f, bounds[i]) {
			return false
		}
	}
	return true
}

// validCronField checks one cron field's characters and numeric range.
func validCronField(field string, upper int) bool {
	if field == "" {
		return false
	}
	for _, part := range strings.Split(field, ",") {
		if part == "" {
			return false
		}
		base := part
		if slash := strings.IndexByte(base, '/'); slash >= 0 {
			if !validCronDigits(base[slash+1:]) {
				return false
			}
			base = base[:slash]
		}
		if base == "" {
			return false
		}
		for _, atom := range strings.Split(base, "-") {
			if atom == "*" {
				continue
			}
			if !validCronDigits(atom) {
				return false
			}
			n := 0
			for _, c := range atom {
				n = n*10 + int(c-'0')
			}
			if n > upper {
				return false
			}
		}
	}
	return true
}

// validCronDigits reports whether s is a non-empty decimal digit string.
func validCronDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// validVirtualTimezone implements the virtual-time half of SCHEDULE-002: the
// simulator never reads a host clock, so only UTC or a Virtual/* label is valid.
func validVirtualTimezone(tz string) bool {
	if tz == "" || tz == "UTC" {
		return true
	}
	return strings.HasPrefix(tz, "Virtual/") && len(tz) > len("Virtual/")
}

// toleranceAllowed implements the STEADY-004 type matrix. Range expectations
// require an absolute tolerance (per the specification example); value
// expectations accept all three numeric tolerances; predicates are boolean
// and are valid only when no numeric tolerance is declared.
func toleranceAllowed(expectType, tolType string) bool {
	switch expectType {
	case schema.ExpectValue:
		switch tolType {
		case schema.ToleranceAbsolute, schema.ToleranceRelative, schema.TolerancePercent:
			return true
		default:
			return false
		}
	case schema.ExpectRange:
		return tolType == schema.ToleranceAbsolute
	case schema.ExpectPredicate:
		return tolType == ""
	default:
		return false
	}
}

// knownExpectation reports whether an expectation type is defined by CES v1.
func knownExpectation(t string) bool {
	return t == schema.ExpectValue || t == schema.ExpectRange || t == schema.ExpectPredicate
}

// knownTolerance reports whether a tolerance type is defined by CES v1.
func knownTolerance(t string) bool {
	return t == schema.ToleranceAbsolute || t == schema.ToleranceRelative ||
		t == schema.TolerancePercent
}

// selectorKeyValid checks a selector key/value against the CES selector grammar
// (non-empty, DNS-style characters, bounded length), used by TARGET-003.
func selectorKeyValid(s string) bool {
	if s == "" || len(s) > 253 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		ok := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c == '/' || c == ':'
		if !ok {
			return false
		}
	}
	return true
}
