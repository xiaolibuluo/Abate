/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package main is the offline plugin publishing signer. It generates an
// Ed25519 release key, signs one plugin manifest, self-verifies the signature
// and writes the manifest envelope consumed by "abate plugin load". The private
// key never leaves the process and is never written to disk; only the public
// half is shipped so the CLI can verify the artifact without a key server.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"abate/plugin/signature"
)

// releaseAnchorID is the trust anchor id used for locally published artifacts.
const releaseAnchorID = "plugin-release-anchor"

// manifestEnvelope mirrors ctl/cli.pluginManifestFile. It is duplicated here so
// the publishing tool stays a standalone main package with no CLI dependency.
type manifestEnvelope struct {
	Manifest        signature.Manifest `json:"manifest"`
	AnchorPublicKey string             `json:"anchorPublicKey"`
}

// main parses the publishing flags and runs the signer.
func main() {
	out := flag.String("out", "", "destination path of the signed manifest envelope")
	id := flag.String("id", "abate/operator", "plugin id to publish")
	version := flag.String("version", "1.0.0", "semantic version to publish")
	buildRef := flag.String("build-ref", "local", "commit or build reference")
	flag.Parse()
	if err := publish(*out, *id, *version, *buildRef); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "plugin-publish:", err)
		os.Exit(1)
	}
}

// publish generates a key pair, signs and self-verifies a manifest, then writes
// the envelope. Every step uses only the Go standard library and the in-repo
// signature package; no network, subprocess or vendored third-party module is
// involved.
func publish(out, id, version, buildRef string) error {
	if out == "" {
		return errors.New("an output path is required")
	}
	if id == "" {
		return errors.New("a plugin id is required")
	}
	ver, err := signature.ParseSemVer(version)
	if err != nil {
		return err
	}
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		return err
	}
	store := signature.NewTrustStore()
	if err := store.Add(releaseAnchorID, pub); err != nil {
		return err
	}
	m := signature.Manifest{
		ID:               signature.PluginID(id),
		Version:          ver,
		Type:             signature.PluginTypeOperator,
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      releaseAnchorID,
		BuildRef:         buildRef,
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		return err
	}
	if err := signature.Verify(store, m); err != nil {
		return fmt.Errorf("self-verification failed: %w", err)
	}
	envelope := manifestEnvelope{
		Manifest:        m,
		AnchorPublicKey: hex.EncodeToString(pub),
	}
	data, err := json.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	//nolint:gosec // the output path is supplied explicitly by the pipeline.
	if err := os.WriteFile(out, data, 0o600); err != nil {
		return err
	}
	sum := sha256.Sum256(data)
	_, _ = fmt.Printf("plugin-publish: id=%s version=%s anchor=%s sha256=%x out=%s\n",
		id, version, releaseAnchorID, sum, out)
	return nil
}
