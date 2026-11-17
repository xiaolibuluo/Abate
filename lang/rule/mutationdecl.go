/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"strings"
)

// MutationDecl is the optional declaration that lets a rule derive one M-family
// mutation operator. A zero value means the rule derives no operator.
type MutationDecl struct {
	// Family names the derived operator family, for example "M-LIT".
	Family string
	// Kind selects the derivation strategy applied to a matched node.
	Kind string
	// Params holds the deterministic derivation parameters.
	Params []string
}

// Enabled reports whether the declaration describes a real derived operator.
func (m MutationDecl) Enabled() bool {
	return m.Family != "" && m.Kind != ""
}

// Valid reports whether the declaration is either empty or fully specified.
func (m MutationDecl) Valid() bool {
	if !m.Enabled() {
		return !m.EnabledPartially()
	}
	for _, param := range m.Params {
		if strings.TrimSpace(param) == "" {
			return false
		}
	}
	return true
}

// EnabledPartially reports a declaration with only one of family and kind set.
func (m MutationDecl) EnabledPartially() bool {
	return (m.Family == "") != (m.Kind == "")
}

// Canonical renders the declaration in a byte-stable form.
func (m MutationDecl) Canonical() string {
	parts := make([]string, 0, 2+len(m.Params))
	parts = append(parts, m.Family, m.Kind)
	parts = append(parts, m.Params...)
	return strings.Join(parts, "\x00")
}
