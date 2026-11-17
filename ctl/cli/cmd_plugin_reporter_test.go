/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	reportpkg "abate/plugin/report"
	"abate/plugin/signature"
)

func TestPluginReporterAndAdapterEntries(t *testing.T) {
	pluginTestReset(t)
	pub, priv := pluginTestKeys(t)
	repID := "abate/cli-reporter"
	rep := pluginTestSignedManifest(t, pub, priv, repID, pluginTestVersion(1, 0, 0), signature.PluginTypeReporter, "observe")
	repPath := pluginTestFile(t, "reporter.json", rep, pub)
	if code, _, errs := invoke("plugin", "load", "--manifest", repPath); code != exitOK {
		t.Fatalf("load reporter = %d (stderr=%s)", code, errs)
	}
	buildArgs, err := json.Marshal(map[string]any{"Findings": []any{}, "Report": map[string]any{}})
	if err != nil {
		t.Fatalf("encode build args: %v", err)
	}
	code, out, errs := invoke("plugin", "call", "--id", repID, "--entry", "build",
		"--args", base64.StdEncoding.EncodeToString(buildArgs))
	if code != exitOK {
		t.Fatalf("call build = %d (stderr=%s)", code, errs)
	}
	var built pluginCallResult
	if err := json.Unmarshal([]byte(out), &built); err != nil {
		t.Fatalf("decode build result: %v", err)
	}
	sarif, err := base64.StdEncoding.DecodeString(built.ResultBase64)
	if err != nil {
		t.Fatalf("decode SARIF: %v", err)
	}
	if err := reportpkg.ValidateSARIFAgainstOASIS(sarif); err != nil {
		t.Fatalf("CLI SARIF failed OASIS validation: %v", err)
	}

	adpID := "abate/cli-adapter"
	adp := pluginTestSignedManifest(t, pub, priv, adpID, pluginTestVersion(1, 0, 0), signature.PluginTypeAdapter, "observe")
	adpPath := pluginTestFile(t, "adapter.json", adp, pub)
	if code, _, errs := invoke("plugin", "load", "--manifest", adpPath); code != exitOK {
		t.Fatalf("load adapter = %d (stderr=%s)", code, errs)
	}
	env, err := json.Marshal(map[string]string{"CI": "true"})
	if err != nil {
		t.Fatalf("encode env: %v", err)
	}
	code, detectOut, errs := invoke("plugin", "call", "--id", adpID, "--entry", "detect",
		"--args", base64.StdEncoding.EncodeToString(env))
	if code != exitOK {
		t.Fatalf("call detect = %d (stderr=%s)", code, errs)
	}
	var detected pluginCallResult
	if err := json.Unmarshal([]byte(detectOut), &detected); err != nil {
		t.Fatalf("decode detect result: %v", err)
	}
	detectedBytes, err := base64.StdEncoding.DecodeString(detected.ResultBase64)
	if err != nil || !strings.Contains(string(detectedBytes), "adp-ci") {
		t.Fatalf("detect result = %s (%v)", detectedBytes, err)
	}
	code, _, errs = invoke("plugin", "call", "--id", adpID, "--entry", "emit",
		"--args", base64.StdEncoding.EncodeToString([]byte("valid-payload")))
	if code != exitOK {
		t.Fatalf("call emit = %d (stderr=%s)", code, errs)
	}
	code, _, errs = invoke("plugin", "call", "--id", adpID, "--entry", "emit",
		"--args", base64.StdEncoding.EncodeToString([]byte{0xff}))
	if code != exitBusinessFailure {
		t.Fatalf("invalid emit exit = %d, want business failure (stderr=%s)", code, errs)
	}
}
