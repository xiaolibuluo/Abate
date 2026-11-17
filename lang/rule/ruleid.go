/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"strconv"
	"strings"
)

// RuleID is a namespaced rule identifier in LANG-SOURCE-NNN form, for example
// GO-EG-001, GO-CRC-014 or GO-VET-003.
type RuleID string //nolint:revive // RuleID is the frozen public contract name.

// ParseRuleID validates text and returns it as a RuleID.
func ParseRuleID(text string) (RuleID, error) {
	parts := strings.Split(text, "-")
	if len(parts) != 3 || !isNamespace(parts[0]) || !isNamespace(parts[1]) ||
		!isSequence(parts[2]) {
		return "", fmt.Errorf("%w: %q", ErrInvalidRuleID, text)
	}
	return RuleID(text), nil
}

// Valid reports whether the identifier has the required three-part shape.
func (id RuleID) Valid() bool {
	_, err := ParseRuleID(string(id))
	return err == nil
}

// Language returns the lowercase language tag of the identifier.
func (id RuleID) Language() string {
	parts := strings.Split(string(id), "-")
	if len(parts) != 3 {
		return ""
	}
	return strings.ToLower(parts[0])
}

// Source returns the specification family encoded in the identifier.
func (id RuleID) Source() string {
	parts := strings.Split(string(id), "-")
	if len(parts) != 3 {
		return ""
	}
	return parts[1]
}

// Sequence returns the three-digit ordinal of the identifier, or zero when
// the identifier is malformed.
func (id RuleID) Sequence() int {
	parts := strings.Split(string(id), "-")
	if len(parts) != 3 {
		return 0
	}
	value, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0
	}
	return value
}

// String returns the identifier text.
func (id RuleID) String() string {
	return string(id)
}

// isNamespace validates one uppercase namespace segment.
func isNamespace(text string) bool {
	if text == "" || !isUpperLetter(text[0]) {
		return false
	}
	for i := 1; i < len(text); i++ {
		if !isUpperLetter(text[i]) && !isDigit(text[i]) {
			return false
		}
	}
	return true
}

// isSequence validates a three-digit ordinal.
func isSequence(text string) bool {
	if len(text) != 3 {
		return false
	}
	for i := 0; i < len(text); i++ {
		if !isDigit(text[i]) {
			return false
		}
	}
	return true
}

// isUpperLetter reports whether a byte is an ASCII uppercase letter.
func isUpperLetter(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

// isDigit reports whether a byte is an ASCII digit.
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
