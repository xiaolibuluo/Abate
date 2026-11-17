/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package adapter provides concrete AdapterPlugin implementations. The ADP-CI adapter
// bridges host events to a CI system without ever performing an external API call:
// detection inspects the environment only, and Emit validates the payload against an
// internal assertion schema. This keeps the plugin fully air-gapped, satisfying the
// project rule that plugins must not reach external services.
package adapter

import (
	"errors"
	"regexp"
	"unicode/utf8"

	"abate/plugin/host"
)

// ADPAdapterID is the stable adapter identifier returned when CI detection succeeds.
const ADPAdapterID = "adp-ci"

// ciProviderKeys are the environment variable names that, when present and set to
// "true", indicate a known CI provider. Detection consults only these keys and never
// contacts the network.
var ciProviderKeys = []string{"CI", "GITHUB_ACTIONS", "GITLAB_CI", "JENKINS_URL"}

// ErrAdapterPayloadRejected is returned by Emit when the payload violates the
// adapter's air-gapped assertion contract.
var ErrAdapterPayloadRejected = errors.New("plugin: adapter payload rejected")

// CIAdapter is the ADP-CI adapter. It is intentionally air-gapped: it performs no
// external API call. Detect inspects the process environment only, and Emit validates
// the payload locally against an internal assertion schema (non-empty, valid UTF-8,
// and optionally a regular expression) instead of contacting any network endpoint.
type CIAdapter struct {
	manifest host.Manifest

	// payloadPattern, when non-nil, is the assertion the emitted payload must match.
	payloadPattern *regexp.Regexp
}

// NewCIAdapter constructs a CI adapter from its manifest.
func NewCIAdapter(m host.Manifest) *CIAdapter {
	return &CIAdapter{manifest: m}
}

// Manifest returns the signed metadata of the plugin.
func (a *CIAdapter) Manifest() host.Manifest { return a.manifest }

// Type returns the plugin kind.
func (*CIAdapter) Type() host.PluginType { return host.PluginTypeAdapter }

// Close releases no resources for the in-memory CI adapter.
func (*CIAdapter) Close() error { return nil }

// Detect inspects the environment and returns a stable adapter id plus whether a CI
// provider was recognized. It never contacts the network.
func (*CIAdapter) Detect(env map[string]string) (string, bool) {
	for _, k := range ciProviderKeys {
		if v, ok := env[k]; ok && v == "true" {
			return ADPAdapterID, true
		}
	}
	return "", false
}

// Emit validates the payload against the internal assertion schema and returns an
// error on violation. It performs no I/O: the payload is checked locally so the
// adapter can run fully air-gapped. A non-empty, valid UTF-8 payload satisfies the
// baseline assertion; an optional payloadPattern tightens the contract further.
func (a *CIAdapter) Emit(payload []byte) error {
	if len(payload) == 0 {
		return ErrAdapterPayloadRejected
	}
	if !utf8.Valid(payload) {
		return ErrAdapterPayloadRejected
	}
	if a.payloadPattern != nil && !a.payloadPattern.Match(payload) {
		return ErrAdapterPayloadRejected
	}
	return nil
}
