/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

// defaultMaxPayloadBytes caps one local payload at 64 KiB.
const defaultMaxPayloadBytes = 1 << 16

// forbiddenPayloadMarkers lists outbound-field names that an air-gapped
// payload must never carry. Both exact keys and keys containing a marker are
// rejected, recursively, so nested transport configuration cannot slip in.
var forbiddenPayloadMarkers = []string{
	"endpoint", "url", "token", "host", "address", "webhook",
	"uri", "secret", "credential", "apikey", "api_key", "authorization",
}

// PayloadAssertion is the shared, zero-I/O assertion for adapter payloads.
//
// Assert checks non-emptiness, size, UTF-8 validity, JSON object shape,
// required fields, a forbidden outbound-field set and an optional field
// allow-list. It never reads the environment, the clock, the filesystem or the
// network, so its result is a pure function of the payload bytes.
type PayloadAssertion struct {
	// Required lists the top-level fields every payload must carry.
	Required []string
	// Allowed optionally restricts the permitted top-level field names.
	Allowed []string
	// MaxBytes caps the payload size; a non-positive value selects the default.
	MaxBytes int
}

// NewPayloadAssertion builds an assertion with the default size cap.
func NewPayloadAssertion(required []string) PayloadAssertion {
	return PayloadAssertion{
		Required: append([]string(nil), required...),
		MaxBytes: defaultMaxPayloadBytes,
	}
}

// Assert validates one payload and returns a wrapped rejection on any breach.
func (p PayloadAssertion) Assert(payload []byte) error {
	if len(payload) == 0 {
		return rejectPayload("empty payload")
	}
	limit := p.MaxBytes
	if limit <= 0 {
		limit = defaultMaxPayloadBytes
	}
	if len(payload) > limit {
		return rejectPayload("payload exceeds the size cap")
	}
	if !utf8.Valid(payload) {
		return rejectPayload("payload is not valid UTF-8")
	}
	var doc map[string]any
	if err := json.Unmarshal(payload, &doc); err != nil {
		return rejectPayload("payload is not a JSON object")
	}
	if doc == nil {
		return rejectPayload("payload is a null object")
	}
	if err := p.checkForbidden(doc); err != nil {
		return err
	}
	if err := p.checkRequired(doc); err != nil {
		return err
	}
	return p.checkAllowed(doc)
}

// checkForbidden walks every map key at every nesting depth in sorted order.
func (p PayloadAssertion) checkForbidden(value any) error {
	switch typed := value.(type) {
	case map[string]any:
		keys := sortedKeys(typed)
		for _, key := range keys {
			if forbiddenKey(key) {
				return rejectPayload("outbound field " + key)
			}
			if err := p.checkForbidden(typed[key]); err != nil {
				return err
			}
		}
	case []any:
		for _, item := range typed {
			if err := p.checkForbidden(item); err != nil {
				return err
			}
		}
	default:
		return nil
	}
	return nil
}

// checkRequired verifies every mandatory field is present and non-nil.
func (p PayloadAssertion) checkRequired(doc map[string]any) error {
	for _, name := range p.Required {
		value, ok := doc[name]
		if !ok {
			return rejectPayload("missing required field " + name)
		}
		if value == nil {
			return rejectPayload("nil required field " + name)
		}
	}
	return nil
}

// checkAllowed enforces the optional top-level allow-list when it is set.
func (p PayloadAssertion) checkAllowed(doc map[string]any) error {
	if len(p.Allowed) == 0 {
		return nil
	}
	for _, key := range sortedKeys(doc) {
		if !slices.Contains(p.Allowed, key) {
			return rejectPayload("field not in allow-list: " + key)
		}
	}
	return nil
}

// sortedKeys returns map keys in ascending order without map-order leakage.
func sortedKeys(doc map[string]any) []string {
	keys := make([]string, 0, len(doc))
	for key := range doc {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

// forbiddenKey reports whether a key names or embeds an outbound field.
func forbiddenKey(key string) bool {
	lower := strings.ToLower(key)
	for _, marker := range forbiddenPayloadMarkers {
		if lower == marker || strings.Contains(lower, marker) {
			return true
		}
	}
	return false
}

// rejectPayload wraps the shared air-gapped rejection sentinel with a reason.
func rejectPayload(reason string) error {
	return fmt.Errorf("%w: %s", ErrAdapterPayloadRejected, reason)
}
