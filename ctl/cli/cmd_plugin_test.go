/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/plugin/signature"
)

// pluginTestKeys generates one release key pair for a test case.
func pluginTestKeys(t *testing.T) (ed25519.PublicKey, ed25519.PrivateKey) {
	t.Helper()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	return pub, priv
}

// pluginTestManifest builds an unsigned manifest under the given anchor.
func pluginTestManifest(t *testing.T,
	id string, version signature.SemVer, typ signature.PluginType, decl string) signature.Manifest {
	t.Helper()
	return signature.Manifest{
		ID:               signature.PluginID(id),
		Version:          version,
		Type:             typ,
		CapabilitiesDecl: decl,
		TrustAnchor:      "plugin-test-anchor",
		BuildRef:         "build-plugin-test",
	}
}

// pluginTestSignedManifest builds and signs a manifest under the given anchor.
func pluginTestSignedManifest(t *testing.T, pub ed25519.PublicKey, priv ed25519.PrivateKey,
	id string, version signature.SemVer, typ signature.PluginType, decl string) signature.Manifest {
	t.Helper()
	m := pluginTestManifest(t, id, version, typ, decl)
	store := signature.NewTrustStore()
	if err := store.Add("plugin-test-anchor", pub); err != nil {
		t.Fatalf("store.Add: %v", err)
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		t.Fatalf("Sign: %v", err)
	}
	return m
}

// pluginTestFile writes a manifest envelope and returns its path.
func pluginTestFile(t *testing.T, name string, m signature.Manifest, pub ed25519.PublicKey) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := writePluginManifestFile(path, m, pub); err != nil {
		t.Fatalf("writePluginManifestFile: %v", err)
	}
	return path
}

// pluginTestReset isolates the persisted CLI session for one test case. The
// temporary directory stays inside the package directory so the bundled OASIS
// SARIF schema remains reachable by upward search during reporter tests.
func pluginTestReset(t *testing.T) {
	t.Helper()
	dir, err := os.MkdirTemp(".", "plugin-state-")
	if err != nil {
		t.Fatalf("MkdirTemp: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Errorf("remove plugin state directory: %v", err)
		}
	})
	t.Chdir(dir)
}

// pluginTestVersion returns a semantic version literal.
func pluginTestVersion(major, minor, patch uint) signature.SemVer {
	return signature.SemVer{Major: major, Minor: minor, Patch: patch}
}

func TestPluginHelpAndDispatch(t *testing.T) {
	pluginTestReset(t)
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"help", []string{"plugin", "--help"}, exitOK},
		{"short-help", []string{"plugin", "-h"}, exitOK},
		{"help-word", []string{"plugin", "help"}, exitOK},
		{"load-help", []string{"plugin", "load", "--help"}, exitOK},
		{"list-help", []string{"plugin", "list", "--help"}, exitOK},
		{"call-help", []string{"plugin", "call", "--help"}, exitOK},
		{"hotswap-help", []string{"plugin", "hotswap", "--help"}, exitOK},
		{"bare", []string{"plugin"}, exitUsageError},
		{"unknown", []string{"plugin", "bogus"}, exitUsageError},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if out == "" && errs == "" {
				t.Fatal("plugin produced no output")
			}
		})
	}
	if !strings.Contains(rootHelp, "abate plugin load") {
		t.Fatal("root help does not document the plugin command")
	}
}

func TestPluginLoadCallListDeterministic(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	id := "abate/cli-operator"
	m := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 0), signature.PluginTypeOperator, "clock,rand")
	path := pluginTestFile(t, "operator.json", m, pub)

	code, out, errs := invoke("plugin", "load", "--manifest", path, "--type", "operator")
	if code != exitOK {
		t.Fatalf("load exit = %d (stderr=%s)", code, errs)
	}
	var loaded pluginLoadResult
	if err := json.Unmarshal([]byte(out), &loaded); err != nil {
		t.Fatalf("decode load result: %v", err)
	}
	if loaded.ID != id || loaded.Version != "1.0.0" || loaded.Anchor != "plugin-test-anchor" {
		t.Fatalf("load result = %+v", loaded)
	}

	code, listOut, errs := invoke("plugin", "list")
	if code != exitOK {
		t.Fatalf("list exit = %d (stderr=%s)", code, errs)
	}
	var items []pluginListItem
	if err := json.Unmarshal([]byte(listOut), &items); err != nil {
		t.Fatalf("decode list: %v", err)
	}
	if len(items) != 1 || items[0].ID != id || items[0].Version != "1.0.0" {
		t.Fatalf("list = %+v", items)
	}
	code, listAgain, _ := invoke("plugin", "list")
	if code != exitOK || listAgain != listOut {
		t.Fatal("plugin list is not deterministic")
	}

	encoded := base64.StdEncoding.EncodeToString([]byte("deterministic-input"))
	code, first, errs := invoke("plugin", "call", "--id", id, "--entry", "apply", "--args", encoded)
	if code != exitOK {
		t.Fatalf("call exit = %d (stderr=%s)", code, errs)
	}
	code, second, errs := invoke("plugin", "call", "--id", id, "--entry", "apply", "--args", encoded)
	if code != exitOK {
		t.Fatalf("call again exit = %d (stderr=%s)", code, errs)
	}
	if first != second {
		t.Fatal("operator call output is not deterministic")
	}
	var result pluginCallResult
	if err := json.Unmarshal([]byte(first), &result); err != nil {
		t.Fatalf("decode call result: %v", err)
	}
	if result.ResultBytes != len("deterministic-input")+1 || result.ResultSHA256 == "" {
		t.Fatalf("call result = %+v", result)
	}
	decoded, err := base64.StdEncoding.DecodeString(result.ResultBase64)
	if err != nil || len(decoded) != result.ResultBytes {
		t.Fatalf("result payload is not valid base64: %v", err)
	}
}

func TestPluginRejectsUnsignedManifestWithoutRegistration(t *testing.T) {
	pluginTestReset(t)
	pub, _ := pluginTestKeys(t)
	id := "abate/cli-unsigned"
	m := pluginTestManifest(t, id, pluginTestVersion(1, 0, 0), signature.PluginTypeOperator, "clock,rand")
	path := pluginTestFile(t, "unsigned.json", m, pub)
	code, _, errs := invoke("plugin", "load", "--manifest", path)
	if code != exitBusinessFailure {
		t.Fatalf("unsigned load exit = %d, want %d (stderr=%s)", code, exitBusinessFailure, errs)
	}
	if !strings.Contains(errs, "signature invalid") {
		t.Fatalf("missing signature diagnostic: %s", errs)
	}
	code, listOut, _ := invoke("plugin", "list")
	if code != exitOK || strings.Contains(listOut, id) {
		t.Fatalf("rejected manifest was registered: %s", listOut)
	}
}

func TestPluginHotswapEndToEnd(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	id := "abate/cli-hotswap"
	old := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 0), signature.PluginTypeOperator, "clock,rand")
	oldPath := pluginTestFile(t, "v1.json", old, pub)
	if code, _, errs := invoke("plugin", "load", "--manifest", oldPath); code != exitOK {
		t.Fatalf("load v1 exit = %d (stderr=%s)", code, errs)
	}
	upgraded := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 1), signature.PluginTypeOperator, "clock,rand")
	upPath := pluginTestFile(t, "v2.json", upgraded, pub)
	code, out, errs := invoke("plugin", "hotswap", "--id", id, "--new-manifest", upPath)
	if code != exitOK {
		t.Fatalf("hotswap exit = %d (stderr=%s)", code, errs)
	}
	var swap pluginSwapResult
	if err := json.Unmarshal([]byte(out), &swap); err != nil {
		t.Fatalf("decode swap result: %v", err)
	}
	if swap.From != "1.0.0" || swap.To != "1.0.1" || swap.State != "UNLOAD" {
		t.Fatalf("swap result = %+v", swap)
	}
	code, listOut, _ := invoke("plugin", "list")
	if code != exitOK || !strings.Contains(listOut, `"version":"1.0.1"`) {
		t.Fatalf("registry did not serve the new version: %s", listOut)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte("after-swap"))
	code, _, errs = invoke("plugin", "call", "--id", id, "--entry", "apply", "--args", encoded)
	if code != exitOK {
		t.Fatalf("call after hotswap = %d (stderr=%s)", code, errs)
	}
	rollback := pluginTestSignedManifest(t, pub, priv, id, pluginTestVersion(1, 0, 0), signature.PluginTypeOperator, "clock,rand")
	rollbackPath := pluginTestFile(t, "rollback.json", rollback, pub)
	code, _, errs = invoke("plugin", "hotswap", "--id", id, "--new-manifest", rollbackPath)
	if code != exitBusinessFailure {
		t.Fatalf("downgrade hotswap exit = %d, want business failure (stderr=%s)", code, errs)
	}
}
