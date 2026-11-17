/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/plugin/signature"
)

// writeRawEnvelope writes an arbitrary envelope so malformed-key cases can be
// exercised without going through the well-formed writer.
func writeRawEnvelope(t *testing.T, raw string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "raw.json")
	//nolint:gosec // the path is created by this test.
	if err := os.WriteFile(path, []byte(raw), 0o600); err != nil {
		t.Fatalf("write raw envelope: %v", err)
	}
	return path
}

func TestPluginLoadUsageErrors(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	good := pluginTestSignedManifest(t, pub, priv, "abate/cli-usage", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	goodPath := pluginTestFile(t, "good.json", good, pub)
	unknownCap := pluginTestSignedManifest(t, pub, priv, "abate/cli-badcap", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "bogus")
	badCapPath := pluginTestFile(t, "badcap.json", unknownCap, pub)
	unknownType := pluginTestSignedManifest(t, pub, priv, "abate/cli-badtype", pluginTestVersion(1, 0, 0),
		signature.PluginType("sandbox"), "clock,rand")
	badTypePath := pluginTestFile(t, "badtype.json", unknownType, pub)
	noAnchor := pluginTestManifest(t, "abate/cli-noanchor", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	noAnchor.TrustAnchor = ""
	noAnchorPath := pluginTestFile(t, "noanchor.json", noAnchor, pub)

	cases := []struct {
		name string
		args []string
	}{
		{"missing-manifest", []string{"plugin", "load"}},
		{"extra-positional", []string{"plugin", "load", "--manifest", goodPath, "extra"}},
		{"unknown-flag", []string{"plugin", "load", "--bogus"}},
		{"wrong-type", []string{"plugin", "load", "--manifest", goodPath, "--type", "reporter"}},
		{"missing-file", []string{"plugin", "load", "--manifest", filepath.Join(t.TempDir(), "absent")}},
		{"malformed-json", []string{"plugin", "load", "--manifest", writeRawEnvelope(t, "{broken")}},
		{"bad-key-hex", []string{"plugin", "load", "--manifest", writeRawEnvelope(t,
			`{"manifest":{"ID":"x","TrustAnchor":"plugin-test-anchor","Type":"operator"},"anchorPublicKey":"zz"}`)}},
		{"bad-key-size", []string{"plugin", "load", "--manifest", writeRawEnvelope(t,
			`{"manifest":{"ID":"x","TrustAnchor":"plugin-test-anchor","Type":"operator"},"anchorPublicKey":"616263"}`)}},
		{"empty-anchor", []string{"plugin", "load", "--manifest", noAnchorPath}},
		{"unknown-capability", []string{"plugin", "load", "--manifest", badCapPath}},
		{"unknown-type", []string{"plugin", "load", "--manifest", badTypePath}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != exitUsageError {
				t.Fatalf("exit = %d, want usage error (stderr=%s)", code, errs)
			}
		})
	}
}

func TestPluginAnchorRebindingRejected(t *testing.T) {
	pluginTestReset(t)
	pubA, privA := pluginTestKeys(t)
	pubB, privB := pluginTestKeys(t)
	first := pluginTestSignedManifest(t, pubA, privA, "abate/cli-anchor-a", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	firstPath := pluginTestFile(t, "anchor-a.json", first, pubA)
	if code, _, errs := invoke("plugin", "load", "--manifest", firstPath); code != exitOK {
		t.Fatalf("first load = %d (stderr=%s)", code, errs)
	}
	if code, _, _ := invoke("plugin", "load", "--manifest", firstPath); code != exitOK {
		t.Fatal("reloading the same anchor and key must succeed")
	}
	second := pluginTestSignedManifest(t, pubB, privB, "abate/cli-anchor-b", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	secondPath := pluginTestFile(t, "anchor-b.json", second, pubB)
	code, _, errs := invoke("plugin", "load", "--manifest", secondPath)
	if code != exitUsageError || !strings.Contains(errs, "another key") {
		t.Fatalf("anchor rebinding = %d (stderr=%s)", code, errs)
	}
}

func TestPluginCallAndListErrors(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	noRand := pluginTestSignedManifest(t, pub, priv, "abate/cli-norand", pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock")
	path := pluginTestFile(t, "norand.json", noRand, pub)
	if code, _, errs := invoke("plugin", "load", "--manifest", path); code != exitOK {
		t.Fatalf("load = %d (stderr=%s)", code, errs)
	}
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"missing-id", []string{"plugin", "call", "--entry", "apply"}, exitUsageError},
		{"unknown-entry", []string{"plugin", "call", "--id", "abate/cli-norand", "--entry", "nope"}, exitUsageError},
		{"bad-base64", []string{"plugin", "call", "--id", "abate/cli-norand", "--entry", "apply", "--args", "!!!"}, exitUsageError},
		{"unknown-flag", []string{"plugin", "call", "--bogus"}, exitUsageError},
		{"unregistered", []string{"plugin", "call", "--id", "abate/absent", "--entry", "apply"}, exitBusinessFailure},
		{"missing-capability", []string{"plugin", "call", "--id", "abate/cli-norand", "--entry", "apply",
			"--args", base64.StdEncoding.EncodeToString([]byte("x"))}, exitBusinessFailure},
		{"list-extra-arg", []string{"plugin", "list", "extra"}, exitUsageError},
		{"list-unknown-flag", []string{"plugin", "list", "--bogus"}, exitUsageError},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
		})
	}
}

func TestPluginHotswapErrors(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	id := "abate/cli-swap-errors"
	old := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	oldPath := pluginTestFile(t, "swap-old.json", old, pub)
	if code, _, errs := invoke("plugin", "load", "--manifest", oldPath); code != exitOK {
		t.Fatalf("load = %d (stderr=%s)", code, errs)
	}
	wrongID := pluginTestSignedManifest(t, pub, priv, "abate/other", pluginTestVersion(1, 0, 1),
		signature.PluginTypeOperator, "clock,rand")
	wrongIDPath := pluginTestFile(t, "swap-wrong-id.json", wrongID, pub)
	unsigned := pluginTestManifest(t, id, pluginTestVersion(1, 0, 1),
		signature.PluginTypeOperator, "clock,rand")
	unsignedPath := pluginTestFile(t, "swap-unsigned.json", unsigned, pub)
	absent := pluginTestSignedManifest(t, pub, priv, "abate/absent", pluginTestVersion(1, 0, 1),
		signature.PluginTypeOperator, "clock,rand")
	absentPath := pluginTestFile(t, "swap-absent.json", absent, pub)
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"missing-flags", []string{"plugin", "hotswap", "--id", id}, exitUsageError},
		{"unknown-flag", []string{"plugin", "hotswap", "--bogus"}, exitUsageError},
		{"missing-file", []string{"plugin", "hotswap", "--id", id, "--new-manifest", "absent.json"}, exitUsageError},
		{"id-mismatch", []string{"plugin", "hotswap", "--id", id, "--new-manifest", wrongIDPath}, exitUsageError},
		{"unregistered", []string{"plugin", "hotswap", "--id", "abate/absent", "--new-manifest", absentPath}, exitBusinessFailure},
		{"unsigned-replacement", []string{"plugin", "hotswap", "--id", id, "--new-manifest", unsignedPath}, exitBusinessFailure},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
		})
	}
}

func TestPluginOutputFailureIsInternalError(t *testing.T) {
	pluginTestReset(t)
	if code := cmdPluginList(nil, failingWriter{}, io.Discard); code != exitInternalError {
		t.Fatalf("list with a failing writer = %d", code)
	}
	pub, priv := pluginTestKeys(t)
	id := "abate/cli-writefail"
	m := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 0),
		signature.PluginTypeOperator, "clock,rand")
	path := pluginTestFile(t, "writefail.json", m, pub)
	if code := cmdPluginLoad([]string{"--manifest", path}, failingWriter{}, io.Discard); code != exitInternalError {
		t.Fatalf("load with a failing writer = %d", code)
	}
	callArgs := []string{"--id", id, "--entry", "apply", "--args", base64.StdEncoding.EncodeToString([]byte("x"))}
	if code := cmdPluginCall(callArgs, failingWriter{}, io.Discard); code != exitInternalError {
		t.Fatalf("call with a failing writer = %d", code)
	}
	next := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 1),
		signature.PluginTypeOperator, "clock,rand")
	nextPath := pluginTestFile(t, "writefail-next.json", next, pub)
	swapArgs := []string{"--id", id, "--new-manifest", nextPath}
	if code := cmdPluginHotswap(swapArgs, failingWriter{}, io.Discard); code != exitInternalError {
		t.Fatalf("hotswap with a failing writer = %d", code)
	}
}
