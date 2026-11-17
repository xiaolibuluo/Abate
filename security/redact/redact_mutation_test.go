/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

import (
	"strings"
	"testing"
)

// TestMutationKeyNamesRequireExactWhitelist proves that single-character
// mutations of a sensitive key are preserved, so a mutant input cannot make the
// scanner silently swallow unrelated text.
func TestMutationKeyNamesRequireExactWhitelist(t *testing.T) {
	mutants := []string{
		"toke=LEAKVALUE",
		"tokens=LEAKVALUE",
		"tokenX=LEAKVALUE",
		"x-token=LEAKVALUE",
		"token:LEAKVALUE",
		"token LEAKVALUE",
	}
	for _, mutant := range mutants {
		if got := Text(mutant); got != mutant {
			t.Fatalf("mutant %q was altered: %q", mutant, got)
		}
	}
	base := "token=LEAKVALUE"
	if got := Text(base); got != "token="+Placeholder {
		t.Fatalf("wild type was not redacted: %q", got)
	}
}

// TestMutationBearerWordBoundary proves that mutating the Bearer marker word
// disables the scheme rule, while a genuine marker is always redacted.
func TestMutationBearerWordBoundary(t *testing.T) {
	cases := []struct {
		input   string
		redacts bool
	}{
		{"Bearer abc", true},
		{"BEARER abc", true},
		{"Bearer\tabc", true},
		{"BearerX abc", false},
		{"xBearer abc", false},
		{"Bearerabc", false},
		{"Bearer=abc", true},
	}
	for _, tc := range cases {
		got := Text(tc.input)
		hasMarker := strings.Contains(got, Placeholder)
		if hasMarker != tc.redacts {
			t.Fatalf("Text(%q) = %q, redacted=%v want %v", tc.input, got, hasMarker, tc.redacts)
		}
	}
}

// TestMutationSeparatorControlsValueExtent proves that changing one terminator
// byte changes exactly how much of the value is swallowed.
func TestMutationSeparatorControlsValueExtent(t *testing.T) {
	cases := []struct {
		name  string
		input string
		call  func(string) string
		want  string
	}{
		{"space", "token=ABCDEFG tail", Text, "token=" + Placeholder + " tail"},
		{"ampersand", "token=ABCDEFG&x=1", Text, "token=" + Placeholder + "&x=1"},
		{"semicolon", "token=ABCDEFG;x=1", Text, "token=" + Placeholder + ";x=1"},
		{"quote", "token='ABCDEFG'", Text, "token=" + Placeholder},
		{"url path", "token=ABCDEFG/next", Path, "token=" + Placeholder + "/next"},
		{"free text path", "token=ABCDEFG/next", Text, "token=" + Placeholder},
		{"no terminator", "token=ABCDEFG", Text, "token=" + Placeholder},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.call(tc.input); got != tc.want {
				t.Fatalf("%s: got %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}
