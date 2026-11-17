/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/signature"
)

// pluginCallResult is the deterministic, binary-safe call receipt. The result
// is base64 encoded and hashed so two identical calls can be compared byte for
// byte by CI without ever interpreting plugin output as text.
type pluginCallResult struct {
	ID           string `json:"id"`
	Entry        string `json:"entry"`
	ResultBase64 string `json:"resultBase64"`
	ResultBytes  int    `json:"resultBytes"`
	ResultSHA256 string `json:"resultSha256"`
}

// pluginRequiredCaps maps an entry to the authority the CLI demands. Reporter
// and adapter entries are observation-bound; the operator entry needs the
// deterministic random stream.
func pluginRequiredCaps(entry string) ([]capability.Capability, bool) {
	switch entry {
	case host.EntryApply:
		return []capability.Capability{capability.CapRand}, true
	case host.EntryBuild, host.EntryDetect, host.EntryEmit:
		return []capability.Capability{capability.CapObserve}, true
	default:
		return nil, false
	}
}

// cmdPluginCall invokes one capability-gated entry on a registered instance.
func cmdPluginCall(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writePluginHelp(stdout)
		return exitOK
	}
	fs := flag.NewFlagSet("plugin call", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	id := fs.String("id", "", "registered plugin id")
	entry := fs.String("entry", "", "entry name")
	argsB64 := fs.String("args", "", "base64 encoded call arguments")
	if err := fs.Parse(reorderArgs(args, pluginValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin call: %v\n", err)
		return exitUsageError
	}
	if *id == "" || fs.NArg() != 0 {
		writePluginHelp(stderr)
		return exitUsageError
	}
	wants, known := pluginRequiredCaps(*entry)
	if !known {
		_, _ = fmt.Fprintf(stderr, "abate plugin call: unknown entry %q\n", *entry)
		return exitUsageError
	}
	var payload []byte
	if *argsB64 != "" {
		decoded, err := base64.StdEncoding.DecodeString(*argsB64)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "abate plugin call: decode args: %v\n", err)
			return exitUsageError
		}
		payload = decoded
	}

	session, err := loadPluginSession()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin call: %v\n", err)
		return exitInternalError
	}
	inst, ok := session.h.Lookup(signature.PluginID(*id))
	if !ok {
		_, _ = fmt.Fprintf(stderr, "abate plugin call: plugin %q is not registered\n", *id)
		return exitBusinessFailure
	}
	out, err := session.h.Call(inst, *entry, payload, wants)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin call: %v\n", err)
		return exitBusinessFailure
	}
	sum := sha256.Sum256(out)
	result := pluginCallResult{
		ID:           *id,
		Entry:        *entry,
		ResultBase64: base64.StdEncoding.EncodeToString(out),
		ResultBytes:  len(out),
		ResultSHA256: hex.EncodeToString(sum[:]),
	}
	data, err := json.Marshal(result)
	if err != nil {
		return exitInternalError
	}
	if _, err := stdout.Write(append(data, '\n')); err != nil {
		return exitInternalError
	}
	return exitOK
}
