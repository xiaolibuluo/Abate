/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package reporter

import (
	"errors"
	"fmt"
	"strings"

	"abate/plugin/host"
)

// ErrUnknownFormat reports a manifest whose BuildRef names no known format.
var ErrUnknownFormat = errors.New("reporter: unknown artifact format")

// ErrWrongPluginType reports a manifest that is not a reporter manifest.
var ErrWrongPluginType = errors.New("reporter: manifest is not a reporter")

// Compile-time proof that Factory matches the host.Factory signature.
var _ host.Factory = Factory

// Factory instantiates one artifactReporter from a verified manifest.
//
// It fails closed on both a wrong plugin type and an unrecognized format. The
// format is carried by Manifest.BuildRef, either as a bare format name or with
// an explicit "format:" / "format=" prefix.
func Factory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeReporter {
		return nil, fmt.Errorf("%w: manifest %q has type %q",
			ErrWrongPluginType, m.ID, m.Type)
	}
	format, err := parseFormat(m.BuildRef)
	if err != nil {
		return nil, err
	}
	return newArtifactReporter(m, format), nil
}

// parseFormat resolves the format token from a manifest build reference.
func parseFormat(text string) (Format, error) {
	value := strings.TrimSpace(text)
	if cut, ok := strings.CutPrefix(value, "format:"); ok {
		value = cut
	}
	if cut, ok := strings.CutPrefix(value, "format="); ok {
		value = cut
	}
	normalized := Format(strings.ToLower(strings.TrimSpace(value)))
	for _, known := range AllFormats() {
		if normalized == known {
			return normalized, nil
		}
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownFormat, text)
}
