/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"strings"
	"testing"
)

// TestHelpersPredicateAndNumber covers the predicate grammar and every
// branch of the signed decimal literal scanner.
func TestHelpersPredicateAndNumber(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {">5", true}, {">=5", true}, {"<5", true},
		{"<=5", true}, {"==5", true}, {"!=5", true}, {">-3.5", true},
		{"5", false}, {">", false}, {">.", false}, {">1a", false},
		{"..", false}, {"abc", false}, {"!=+", false}, {">+2", true},
	} {
		if got := validPredicate(c.in); got != c.want {
			t.Fatalf("validPredicate(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"-", false}, {"+", false}, {".", false},
		{"1", true}, {"-3.5", true}, {"+2", true}, {"12", true},
		{"1.2.3", false}, {"1a", false}, {"1..2", false}, {"1-2", false},
	} {
		if got := validNumber(c.in); got != c.want {
			t.Fatalf("validNumber(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestHelpersCronGrammar covers the five-field cron grammar and its range,
// step, list and separator branches.
func TestHelpersCronGrammar(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"* * * * *", true}, {"*/5 0-23 * * 1-5", true},
		{"0 0 1 1 0", true}, {"* * * *", false}, {"* * * * * *", false},
		{"60 * * * *", false}, {"* 24 * * *", false}, {"* * 32 * *", false},
		{"* * * 13 *", false}, {"* * * * 8", false},
		{"1,,2 * * * *", false}, {"*/ * * * *", false}, {"/5 * * * *", false},
		{"*/x * * * *", false}, {"- * * * *", false}, {"1- * * * *", false},
		{"", false},
	} {
		if got := validCron(c.in); got != c.want {
			t.Fatalf("validCron(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"0", true}, {"007", true}, {"12a", false}, {"1.5", false},
	} {
		if got := validCronDigits(c.in); got != c.want {
			t.Fatalf("validCronDigits(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestHelpersVirtualTimezone covers the virtual-time timezone grammar.
func TestHelpersVirtualTimezone(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", true}, {"UTC", true}, {"Virtual/Tokyo", true},
		{"Virtual/", false}, {"Virtual", false},
		{"Asia/Shanghai", false}, {"utc", false},
	} {
		if got := validVirtualTimezone(c.in); got != c.want {
			t.Fatalf("validVirtualTimezone(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestHelpersSimulatorURL covers the simulator iron rule for every URL shape.
func TestHelpersSimulatorURL(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"/health", true}, {"sim://probe", true},
		{"http://sim/health", true}, {"https://sim:8443/health", true},
		{"http://localhost/health", true}, {"https://127.0.0.1:9443/x", true},
		{"https://evil.example.com/x", false}, {"ftp://sim/x", false},
		{"sim", false}, {"http://SIM/x", false}, {"https://::1/x", false},
	} {
		if got := isSimulatorURL(c.in); got != c.want {
			t.Fatalf("isSimulatorURL(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	for _, c := range []struct {
		in   string
		want string
		ok   bool
	}{
		{"http://sim/x", "sim/x", true}, {"https://sim/x", "sim/x", true},
		{"ftp://sim/x", "", false}, {"sim://x", "", false},
	} {
		got, ok := trimScheme(c.in)
		if got != c.want || ok != c.ok {
			t.Fatalf("trimScheme(%q) = (%q,%v), want (%q,%v)",
				c.in, got, ok, c.want, c.ok)
		}
	}
}

// TestHelpersShellCommand covers the read-only simulator shell subset.
func TestHelpersShellCommand(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"   ", false}, {"echo ok", true}, {"true", true},
		{"ls -la /tmp", true}, {"cat /etc/hosts", true},
		{"echo ok; rm -rf /", false}, {"a | b", false}, {"a & b", false},
		{"a > b", false}, {"a < b", false}, {"a `b`", false},
		{"$(whoami)", false}, {"a\nb", false}, {"(echo)", false},
		{"rm -rf /", false}, {"/bin/sh", false},
	} {
		if got := isSafeShellCommand(c.in); got != c.want {
			t.Fatalf("isSafeShellCommand(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	for _, tok := range []string{
		"echo", "true", "false", "printf", "test", "cat", "grep",
		"ls", "pwd", "date", "sleep", "wc", "head", "tail", "sort", "uniq",
	} {
		if !readOnlyCommand(tok) {
			t.Fatalf("readOnlyCommand(%q) = false", tok)
		}
	}
	if readOnlyCommand("rm") {
		t.Fatal("readOnlyCommand accepted a mutating command")
	}
}

// TestHelpersSelector covers the CES selector key/value grammar.
func TestHelpersSelector(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"", false}, {"node", true}, {"a/b:c_d.e-1", true},
		{"bad key", false}, {"bad@key", false},
		{strings.Repeat("a", 254), false}, {strings.Repeat("a", 253), true},
	} {
		if got := selectorKeyValid(c.in); got != c.want {
			t.Fatalf("selectorKeyValid(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestHelpersMisc covers the remaining pure helpers and violation builders.
func TestHelpersMisc(t *testing.T) {
	t.Parallel()
	if got := path("spec.probes", 3); got != "spec.probes[3]" {
		t.Fatalf("path = %q", got)
	}
	if got := probeIndex(fresh(t), "p1"); got != 0 {
		t.Fatalf("probeIndex(p1) = %d, want 0", got)
	}
	if got := probeIndex(fresh(t), "ghost"); got != -1 {
		t.Fatalf("probeIndex(ghost) = %d, want -1", got)
	}
	for _, r := range []rune{'ε', '9', '-', '_', '.'} {
		if !idRune(r) {
			t.Fatalf("idRune(%q) = false", r)
		}
	}
	if idRune(' ') || idRune('@') {
		t.Fatal("idRune accepted a forbidden rune")
	}
	v := violation("X", "p", "m", SeverityWarning)
	if v.RuleID != "X" || v.FieldPath != "p" || v.Message != "m" ||
		v.Severity != SeverityWarning {
		t.Fatalf("violation = %+v", v)
	}
	w := warningAt("Y", "q", "n")
	if w.Severity != SeverityWarning || w.RuleID != "Y" {
		t.Fatalf("warningAt = %+v", w)
	}
	e := errorAt("Z", "r", "o")
	if e.Severity != SeverityError || e.RuleID != "Z" {
		t.Fatalf("errorAt = %+v", e)
	}
}
