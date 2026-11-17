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

// TestTextRedaction is the table-driven contract for free-form text covering
// normal secrets, boundary shapes and malformed inputs.
func TestTextRedaction(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"bearer standard", "Bearer abc123", "Bearer [REDACTED]"},
		{"bearer lower case", "bearer AbC.XYZ-9", "bearer [REDACTED]"},
		{"bearer upper case and wide space", "BEARER   tok", "BEARER   [REDACTED]"},
		{"bearer without value", "Bearer ", "Bearer [REDACTED]"},
		{"bare bearer word kept", "Bearer", "Bearer"},
		{"bearer glued word kept", "BearerX abc", "BearerX abc"},
		{"token pair", "token=abc123", "token=[REDACTED]"},
		{"token pair upper key", "TOKEN=abc123", "TOKEN=[REDACTED]"},
		{"query with ampersand", "?token=abc&x=1", "?token=[REDACTED]&x=1"},
		{"two pairs", "secret=xyz&password=pwd", "secret=[REDACTED]&password=[REDACTED]"},
		{"space terminates value", "api_key=key123 next", "api_key=[REDACTED] next"},
		{"authorization header", "Authorization: Bearer abc.def", "Authorization: [REDACTED]"},
		{"authorization spaced", "authorization :  Basic dXNlcjpwYXNz", "authorization :  [REDACTED]"},
		{"authorization header fragment", "Authorization: Bearer abc&x=1", "Authorization: [REDACTED]&x=1"},
		{"authorization as pair", "Authorization=Bearer abc", "Authorization=[REDACTED]"},
		{"hyphen key", "api-key=abc", "api-key=[REDACTED]"},
		{"compact key", "apikey=abc", "apikey=[REDACTED]"},
		{"client secret", "client_secret=abc", "client_secret=[REDACTED]"},
		{"access token", "access_token=abc", "access_token=[REDACTED]"},
		{"private key", "private_key=abc", "private_key=[REDACTED]"},
		{"passwd key", "passwd=abc", "passwd=[REDACTED]"},
		{"pwd key", "pwd=abc", "pwd=[REDACTED]"},
		{"session key", "session=abc", "session=[REDACTED]"},
		{"single quoted value", "token='quoted'", "token=[REDACTED]"},
		{"double quoted value", "token=\"abc\"", "token=[REDACTED]"},
		{"quoted bearer", "Bearer \"abc\"", "Bearer [REDACTED]"},
		{"unclosed quote", "token='abc", "token=[REDACTED]"},
		{"empty string", "", ""},
		{"clean text", "plain log line", "plain log line"},
		{"unknown key kept", "mytoken=abc", "mytoken=abc"},
		{"key without equals", "token", "token"},
		{"empty value", "token=", "token=[REDACTED]"},
		{"repeated pair", "token=abc token=def", "token=[REDACTED] token=[REDACTED]"},
		{"placeholder not recursive", "token=[REDACTED]", "token=[REDACTED]"},
		{"header placeholder not recursive", "Authorization: [REDACTED]", "Authorization: [REDACTED]"},
		{"non secret punctuation", "a=b; c=d", "a=b; c=d"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Text(tc.input); got != tc.want {
				t.Fatalf("Text(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

// TestPathRedaction proves the URL-aware scanner preserves route structure.
func TestPathRedaction(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"path segment preserved", "/v1/token=SECRET/detail", "/v1/token=[REDACTED]/detail"},
		{"query and extra key", "/v1/run?token=SECRET&x=1", "/v1/run?token=[REDACTED]&x=1"},
		{"fragment separator", "/v1/token=SECRET#frag", "/v1/token=[REDACTED]#frag"},
		{"multiple path keys", "/a/secret=one/b/api_key=two", "/a/secret=[REDACTED]/b/api_key=[REDACTED]"},
		{"encoded bearer preserved", "/v1/Bearer%20TOKENVALUE", "/v1/Bearer%20TOKENVALUE"},
		{"clean path", "/v1/experiments/42", "/v1/experiments/42"},
		{"empty path", "", ""},
		{"path placeholder not recursive", "/v1/token=[REDACTED]/detail", "/v1/token=[REDACTED]/detail"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Path(tc.input); got != tc.want {
				t.Fatalf("Path(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

// TestTextLongInput proves the scanner stays linear and still redacts a secret
// embedded at the end of a very large log line.
func TestTextLongInput(t *testing.T) {
	prefix := strings.Repeat("x", 100000)
	input := prefix + " token=LEAKVALUE"
	want := prefix + " token=[REDACTED]"
	got := Text(input)
	if got != want {
		t.Fatalf("long input mismatch: len(got)=%d len(want)=%d", len(got), len(want))
	}
	if strings.Contains(got, "LEAKVALUE") {
		t.Fatal("secret survived redaction in a long input")
	}
}

// TestTextMalformedAndBoundaryInput covers malformed and adversarial shapes:
// invalid UTF-8, separator-only strings, nested schemes and each terminator.
func TestTextMalformedAndBoundaryInput(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"invalid utf8 preserved", string([]byte{'t', 'o', 'k', 'e', 'n', '=', 0xff, 0xfe}), "token=[REDACTED]"},
		{"only separators", "&&& \"\"\n", "&&& \"\"\n"},
		{"colon without authorization", "note: Bearer abc", "note: Bearer [REDACTED]"},
		{"nested bearer in value", "token=Bearer abc", "token=[REDACTED]"},
		{"semicolon terminates", "token=abc;next=1", "token=[REDACTED];next=1"},
		{"tab terminates", "token=abc\tdef", "token=[REDACTED]\tdef"},
		{"newline terminates", "token=abc\nline", "token=[REDACTED]\nline"},
		{"hash terminates", "token=abc#frag", "token=[REDACTED]#frag"},
		{"backtick terminates", "token=abc`x", "token=[REDACTED]`x"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Text(tc.input); got != tc.want {
				t.Fatalf("Text(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

// TestPathAndTextDifferOnURLStructure pins the intentional difference between
// the two exported entry points.
func TestPathAndTextDifferOnURLStructure(t *testing.T) {
	input := "/v1/token=SECRET/detail"
	if got, want := Text(input), "/v1/token=[REDACTED]"; got != want {
		t.Fatalf("Text(%q) = %q, want %q", input, got, want)
	}
	if got, want := Path(input), "/v1/token=[REDACTED]/detail"; got != want {
		t.Fatalf("Path(%q) = %q, want %q", input, got, want)
	}
}
