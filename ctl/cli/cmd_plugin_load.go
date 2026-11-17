/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/sample"
	"abate/plugin/signature"
)

// pluginManifestFile is the on-disk envelope for a signed plugin manifest. It
// carries the public half of the trust anchor so the CLI can verify a signature
// without ever handling a private key. The format is shared with the local
// publish pipeline (scripts/plugin-publish.sh).
type pluginManifestFile struct {
	Manifest        signature.Manifest `json:"manifest"`
	AnchorPublicKey string             `json:"anchorPublicKey"`
}

// pluginLoadResult is the deterministic load receipt.
type pluginLoadResult struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Anchor  string `json:"anchor"`
}

// writePluginManifestFile serializes a signed manifest envelope. It is the
// single writer used by tests and by the publish evidence pipeline.
func writePluginManifestFile(path string, m signature.Manifest, pub ed25519.PublicKey) error {
	payload := pluginManifestFile{Manifest: m, AnchorPublicKey: hex.EncodeToString(pub)}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("plugin: encode manifest envelope: %w", err)
	}
	//nolint:gosec // the destination path is supplied explicitly by the caller.
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("plugin: write manifest envelope: %w", err)
	}
	return nil
}

// readPluginManifestFile parses a signed manifest envelope from disk.
func readPluginManifestFile(path string) (pluginManifestFile, error) {
	var file pluginManifestFile
	//nolint:gosec // the manifest path is supplied explicitly by the caller.
	data, err := os.ReadFile(path)
	if err != nil {
		return file, fmt.Errorf("plugin: read manifest %q: %w", path, err)
	}
	if err := json.Unmarshal(data, &file); err != nil {
		return file, fmt.Errorf("plugin: decode manifest %q: %w", path, err)
	}
	return file, nil
}

// registerPluginAnchor adds the envelope public key to the session trust store.
// Re-registering the same anchor with a different key is rejected so a manifest
// cannot silently swap the trust root of a running session.
func registerPluginAnchor(session *pluginSession, file pluginManifestFile) error {
	anchor := file.Manifest.TrustAnchor
	if anchor == "" {
		return errors.New("plugin: manifest has no trust anchor")
	}
	pub, err := hex.DecodeString(file.AnchorPublicKey)
	if err != nil {
		return fmt.Errorf("plugin: decode anchor public key: %w", err)
	}
	if len(pub) != ed25519.PublicKeySize {
		return fmt.Errorf("plugin: anchor public key has size %d", len(pub))
	}
	key := ed25519.PublicKey(pub)
	if existing := session.store.Resolve(anchor); existing != nil {
		if session.store.Fingerprint(existing) != session.store.Fingerprint(key) {
			return fmt.Errorf("plugin: trust anchor %q already bound to another key", anchor)
		}
		return nil
	}
	if err := session.store.Add(anchor, key); err != nil {
		return fmt.Errorf("plugin: register trust anchor: %w", err)
	}
	return nil
}

// pluginFactory maps a manifest type to its reference factory.
func pluginFactory(t signature.PluginType) (host.Factory, bool) {
	switch t {
	case signature.PluginTypeOperator:
		return sample.OperatorFactory, true
	case signature.PluginTypeReporter:
		return sample.ReporterFactory, true
	case signature.PluginTypeAdapter:
		return sample.AdapterFactory, true
	default:
		return nil, false
	}
}

// pluginGrantedCaps derives the effective grant from the manifest declaration.
// The host intersects it with this value, so a plugin only receives what it
// declared: an undeclared authority can never be exercised through the CLI.
func pluginGrantedCaps(decl string) (host.Capabilities, error) {
	return capability.ParseGrant(decl)
}

// cmdPluginLoad verifies and loads a signed manifest into the CLI session.
func cmdPluginLoad(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writePluginHelp(stdout)
		return exitOK
	}
	fs := flag.NewFlagSet("plugin load", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", "", "path to a signed manifest envelope")
	typeFlag := fs.String("type", "", "optional expected plugin type")
	if err := fs.Parse(reorderArgs(args, pluginValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitUsageError
	}
	if *manifestPath == "" || fs.NArg() != 0 {
		writePluginHelp(stderr)
		return exitUsageError
	}
	file, err := readPluginManifestFile(*manifestPath)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitUsageError
	}
	if *typeFlag != "" && *typeFlag != string(file.Manifest.Type) {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: declared type %q disagrees with %q\n", file.Manifest.Type, *typeFlag)
		return exitUsageError
	}
	factory, ok := pluginFactory(file.Manifest.Type)
	if !ok {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: unsupported plugin type %q\n", file.Manifest.Type)
		return exitUsageError
	}
	caps, err := pluginGrantedCaps(file.Manifest.CapabilitiesDecl)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitUsageError
	}

	session, err := loadPluginSession()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitInternalError
	}
	if err := registerPluginAnchor(session, file); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitUsageError
	}
	inst, err := session.h.Load(file.Manifest, factory, caps)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitBusinessFailure
	}
	session.envelopes = upsertEnvelope(session.envelopes, file)
	if err := savePluginSession(session); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin load: %v\n", err)
		return exitInternalError
	}
	result := pluginLoadResult{
		ID:      string(inst.Manifest.ID),
		Type:    string(inst.Manifest.Type),
		Version: inst.Manifest.Version.String(),
		Anchor:  inst.Manifest.TrustAnchor,
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
