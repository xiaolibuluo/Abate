/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"errors"
	"fmt"
)

// APIVersion is the CES API version produced by this build.
const APIVersion = "ces.tencent.com/v1"

// KindExperiment is the only document kind accepted in CES v1.
const KindExperiment = "Experiment"

// ErrUnsupportedVersion is returned when a client asks for an unknown API version.
var ErrUnsupportedVersion = errors.New("unsupported CES API version")

// Version negotiates the CES API version of a document.
type Version struct{}

// NewVersion returns the version negotiator for this build.
func NewVersion() Version {
	return Version{}
}

// Current returns the API version produced by this build.
func (Version) Current() string {
	return APIVersion
}

// SupportedVersions returns a fresh list of negotiable API versions.
func (Version) SupportedVersions() []string {
	return []string{APIVersion}
}

// Negotiate resolves a requested API version to a supported one.
//
// An empty request selects the current version. Any other value must exactly
// match a supported version, otherwise ErrUnsupportedVersion is returned with
// the full supported list in the message.
func (v Version) Negotiate(requested string) (string, error) {
	if requested == "" {
		return v.Current(), nil
	}
	for _, s := range v.SupportedVersions() {
		if s == requested {
			return s, nil
		}
	}
	return "", fmt.Errorf("%w: requested %q, supported %v",
		ErrUnsupportedVersion, requested, v.SupportedVersions())
}
