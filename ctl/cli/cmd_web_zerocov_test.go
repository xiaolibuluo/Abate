/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"strings"
	"testing"

	"abate/ctl/web"
)

// TestCmdWebDelegatesToWebRun proves the CLI adapter forwards the argument
// slice and both output streams to web.Run without transformation, and that
// each deterministic, non-blocking web branch keeps its documented exit code
// and diagnostic text.
func TestCmdWebDelegatesToWebRun(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		args    []string
		want    int
		outPart string
		errPart string
	}{
		{"long-help", []string{"--help"}, exitOK, "Usage: abate web", ""},
		{"short-help", []string{"-h"}, exitOK, "Usage: abate web", ""},
		{"word-help", []string{"-help"}, exitOK, "Usage: abate web", ""},
		{"invalid-flag", []string{"--bogus"}, exitUsageError, "", "Usage: abate web"},
		{"invalid-rest", []string{"--rest", "://bad", "--listen", "127.0.0.1:0"},
			exitUsageError, "", "invalid --rest"},
		{"bad-port", []string{"--listen", "127.0.0.1:99999"},
			exitInternalError, "", "cannot bind"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotOut, gotErr bytes.Buffer
			got := cmdWeb(tc.args, &gotOut, &gotErr)
			if got != tc.want {
				t.Fatalf("cmdWeb(%v) = %d, want %d (stderr=%q)",
					tc.args, got, tc.want, gotErr.String())
			}
			var refOut, refErr bytes.Buffer
			ref := web.Run(tc.args, &refOut, &refErr)
			if got != ref || gotOut.String() != refOut.String() ||
				gotErr.String() != refErr.String() {
				t.Fatalf("cmdWeb diverged from web.Run: code %d/%d stdout %q/%q stderr %q/%q",
					got, ref, gotOut.String(), refOut.String(),
					gotErr.String(), refErr.String())
			}
			if tc.outPart != "" && !strings.Contains(gotOut.String(), tc.outPart) {
				t.Fatalf("stdout %q missing %q", gotOut.String(), tc.outPart)
			}
			if tc.errPart != "" && !strings.Contains(gotErr.String(), tc.errPart) {
				t.Fatalf("stderr %q missing %q", gotErr.String(), tc.errPart)
			}
		})
	}
}
