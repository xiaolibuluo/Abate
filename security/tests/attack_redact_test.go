/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tests

import (
	"strings"
	"testing"

	"abate/security/redact"
)

// TestThreatT9SecretRedaction proves T9 at the package surface: every known
// secret shape is neutralized by redact.Text or redact.Path, clean text is
// preserved, and the fixed placeholder is never recursively expanded.
func TestThreatT9SecretRedaction(t *testing.T) {
	const leaked = "LEAKVALUE-t9-9f8e7d6c"
	inputs := []struct {
		name  string
		input string
		call  func(string) string
	}{
		{"bearer header", "Bearer " + leaked, redact.Text},
		{"query token", "/v1/run?token=" + leaked + "&x=1", redact.Path},
		{"authorization header", "Authorization: Bearer " + leaked, redact.Text},
		{"password pair", "password=" + leaked, redact.Text},
		{"upper case key", "API_KEY=" + leaked, redact.Text},
		{"path segment", "/v1/secret=" + leaked + "/detail", redact.Path},
	}
	for _, tc := range inputs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call(tc.input)
			if strings.Contains(got, leaked) {
				t.Fatalf("secret survived redaction: %q", got)
			}
			if !strings.Contains(got, redact.Placeholder) {
				t.Fatalf("redaction marker missing: %q", got)
			}
			if got == tc.input {
				t.Fatal("input was not transformed")
			}
		})
	}

	once := redact.Text("token=" + leaked)
	if redact.Text(once) != once {
		t.Fatalf("redaction is not idempotent: %q -> %q", once, redact.Text(once))
	}
	const clean = "experiment-42 completed with 0 violations"
	if redact.Text(clean) != clean || redact.Path(clean) != clean {
		t.Fatal("clean log text was altered")
	}
}
