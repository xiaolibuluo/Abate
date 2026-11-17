/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"abate/plugin/host"
	"abate/plugin/signature"
)

// pluginSession is a rebuilt plugin runtime for one CLI invocation. It owns a
// fresh trust store and host, plus the signed manifest envelopes that were
// loaded by previous invocations. State lives in an explicit file rather than
// in process memory, so "abate plugin load", "list", "call" and "hotswap" work
// as separate processes while remaining fully deterministic and replayable.
type pluginSession struct {
	store     *signature.TrustStore
	h         *host.Host
	envelopes []pluginManifestFile
}

// newPluginSession constructs an empty session with a fresh trust store.
func newPluginSession() *pluginSession {
	store := signature.NewTrustStore()
	return &pluginSession{store: store, h: host.NewHost(store)}
}

// pluginStatePath returns the deterministic session state path, anchored at
// the current working directory. CI runs every plugin command from the same
// directory; tests use t.Chdir to isolate each case. No process environment is
// read or mutated, so the command tree stays deterministic and replayable.
func pluginStatePath() string {
	return filepath.Join(".abate", "plugin-state.json")
}

// loadPluginSession rebuilds a session from the state file. Every stored
// envelope is re-verified through Host.Load, so a tampered state file can never
// turn an unsigned plugin into a registered instance.
func loadPluginSession() (*pluginSession, error) {
	session := newPluginSession()
	//nolint:gosec // the state path is an explicit, caller-controlled location.
	data, err := os.ReadFile(pluginStatePath())
	if errors.Is(err, os.ErrNotExist) {
		return session, nil
	}
	if err != nil {
		return nil, fmt.Errorf("plugin: read state: %w", err)
	}
	var envelopes []pluginManifestFile
	if err := json.Unmarshal(data, &envelopes); err != nil {
		return nil, fmt.Errorf("plugin: decode state: %w", err)
	}
	for i := range envelopes {
		envelope := envelopes[i]
		if err := registerPluginAnchor(session, envelope); err != nil {
			return nil, err
		}
		factory, ok := pluginFactory(envelope.Manifest.Type)
		if !ok {
			return nil, fmt.Errorf("plugin: state holds unsupported type %q", envelope.Manifest.Type)
		}
		caps, err := pluginGrantedCaps(envelope.Manifest.CapabilitiesDecl)
		if err != nil {
			return nil, err
		}
		if _, err := session.h.Load(envelope.Manifest, factory, caps); err != nil {
			return nil, fmt.Errorf("plugin: state holds an invalid manifest: %w", err)
		}
		session.envelopes = append(session.envelopes, envelope)
	}
	return session, nil
}

// savePluginSession persists the verified envelopes for the next invocation.
func savePluginSession(session *pluginSession) error {
	data, err := json.Marshal(session.envelopes)
	if err != nil {
		return fmt.Errorf("plugin: encode state: %w", err)
	}
	path := pluginStatePath()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("plugin: create state directory: %w", err)
	}
	//nolint:gosec // the state path is an explicit, caller-controlled location.
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("plugin: write state: %w", err)
	}
	return nil
}

// upsertEnvelope replaces the envelope with the same plugin id, or appends it.
// Replace keeps hot-swaps and repeated loads from accumulating stale versions.
func upsertEnvelope(envelopes []pluginManifestFile, next pluginManifestFile) []pluginManifestFile {
	out := make([]pluginManifestFile, 0, len(envelopes)+1)
	for _, existing := range envelopes {
		if existing.Manifest.ID != next.Manifest.ID {
			out = append(out, existing)
		}
	}
	return append(out, next)
}

// pluginHelp documents the plugin command tree and its copyable examples.
const pluginHelp = `Usage: abate plugin <subcommand> [flags]

Manage signed plugins in the native plugin host. Every load verifies the
Ed25519 signature against the public key shipped in the manifest file; an
unsigned or tampered manifest is rejected and never registered.

Subcommands:
  load --manifest <file> [--type operator|reporter|adapter]
                        Verify and load a signed manifest.
  list                  List registered plugin ids in deterministic order.
  call --id <id> --entry <entry> [--args <base64>]
                        Call an entry; entries: apply, build, detect, emit.
  hotswap --id <id> --new-manifest <file>
                        Atomically replace an instance with a newer version
                        without interrupting in-flight calls.

Exit codes (PLAN-07 §4):
  0  success
  1  business failure (signature, capability or hot-swap rejection)
  2  internal error (output or encoding failure)
  3  usage error (bad flags, unreadable or malformed manifest file)

Examples:
  abate plugin load --manifest plugin-manifest.json --type operator
  abate plugin list
  abate plugin call --id abate/operator --entry apply --args ZWJhdGU=
  abate plugin hotswap --id abate/operator --new-manifest plugin-v2.json
  abate plugin --help
`

// writePluginHelp prints the deterministic plugin usage text.
func writePluginHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, pluginHelp)
}

// pluginValueFlags lists the value-taking flags of every plugin subcommand.
func pluginValueFlags() []string {
	return []string{"manifest", "type", "id", "entry", "args", "new-manifest"}
}

// cmdPlugin dispatches the plugin subcommands.
func cmdPlugin(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		writePluginHelp(stderr)
		return exitUsageError
	}
	switch args[0] {
	case "load":
		return cmdPluginLoad(args[1:], stdout, stderr)
	case "list":
		return cmdPluginList(args[1:], stdout, stderr)
	case "call":
		return cmdPluginCall(args[1:], stdout, stderr)
	case "hotswap":
		return cmdPluginHotswap(args[1:], stdout, stderr)
	case "help", "--help", "-h":
		writePluginHelp(stdout)
		return exitOK
	default:
		_, _ = fmt.Fprintf(stderr, "abate plugin: unknown subcommand %q\n", args[0])
		writePluginHelp(stderr)
		return exitUsageError
	}
}
