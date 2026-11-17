/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"abate/plugin/hotswap"
	"abate/plugin/signature"
)

// pluginSwapResult is the deterministic hot-swap receipt.
type pluginSwapResult struct {
	ID    string `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	State string `json:"state"`
}

// cmdPluginHotswap loads a newer signed manifest and atomically replaces the
// registered instance without interrupting in-flight calls.
func cmdPluginHotswap(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writePluginHelp(stdout)
		return exitOK
	}
	fs := flag.NewFlagSet("plugin hotswap", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	id := fs.String("id", "", "registered plugin id")
	newManifest := fs.String("new-manifest", "", "path to the replacement manifest envelope")
	if err := fs.Parse(reorderArgs(args, pluginValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitUsageError
	}
	if *id == "" || *newManifest == "" || fs.NArg() != 0 {
		writePluginHelp(stderr)
		return exitUsageError
	}
	file, err := readPluginManifestFile(*newManifest)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitUsageError
	}
	if string(file.Manifest.ID) != *id {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: manifest id %q does not match %q\n", file.Manifest.ID, *id)
		return exitUsageError
	}
	factory, ok := pluginFactory(file.Manifest.Type)
	if !ok {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: unsupported plugin type %q\n", file.Manifest.Type)
		return exitUsageError
	}
	caps, err := pluginGrantedCaps(file.Manifest.CapabilitiesDecl)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitUsageError
	}

	session, err := loadPluginSession()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitInternalError
	}
	old, ok := session.h.Lookup(signature.PluginID(*id))
	if !ok {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: plugin %q is not registered\n", *id)
		return exitBusinessFailure
	}
	if err := registerPluginAnchor(session, file); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitUsageError
	}
	next, err := session.h.Load(file.Manifest, factory, caps)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: load replacement: %v\n", err)
		return exitBusinessFailure
	}
	swap, err := hotswap.Begin(context.Background(), old, next)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitBusinessFailure
	}
	session.envelopes = upsertEnvelope(session.envelopes, file)
	if err := savePluginSession(session); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin hotswap: %v\n", err)
		return exitInternalError
	}
	result := pluginSwapResult{
		ID:    *id,
		From:  old.Manifest.Version.String(),
		To:    next.Manifest.Version.String(),
		State: swap.State().String(),
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
