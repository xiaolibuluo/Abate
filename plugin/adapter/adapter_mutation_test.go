/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"regexp"
	"testing"
)

// TestDetectIsSensitiveToProviderValue proves detection is not vacuous: a
// single-character environment mutation flips the result.
func TestDetectIsSensitiveToProviderValue(t *testing.T) {
	a := newTestAdapter()
	if _, ok := a.Detect(map[string]string{"CI": "true"}); !ok {
		t.Fatal("canonical CI marker did not match")
	}
	for _, mutated := range []string{"TRUE", "True", "1", "", "true "} {
		if _, ok := a.Detect(map[string]string{"CI": mutated}); ok {
			t.Fatalf("mutated CI value %q still matched", mutated)
		}
	}
}

// TestEmitIsSensitiveToPayload proves payload assertions respond to input
// mutations rather than always succeeding or always failing.
func TestEmitIsSensitiveToPayload(t *testing.T) {
	a := newTestAdapter()
	pattern, err := regexp.Compile(`^release-[0-9]+$`)
	if err != nil {
		t.Fatalf("compile payload pattern: %v", err)
	}
	a.payloadPattern = pattern
	if err := a.Emit([]byte("release-7")); err != nil {
		t.Fatalf("valid payload rejected: %v", err)
	}
	if err := a.Emit([]byte("release-x")); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("mutated payload accepted: %v", err)
	}
	if err := a.Emit([]byte("release-7\xff")); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("invalid UTF-8 payload accepted: %v", err)
	}
}

// TestProviderKeyOrderIsStable proves the provider key list is ordered and its
// first match is independent of map iteration order.
func TestProviderKeyOrderIsStable(t *testing.T) {
	if len(ciProviderKeys) == 0 {
		t.Fatal("no provider keys configured")
	}
	for i := 1; i < len(ciProviderKeys); i++ {
		if ciProviderKeys[i-1] >= ciProviderKeys[i] {
			t.Fatalf("provider keys are not in a stable order: %v", ciProviderKeys)
		}
	}
	env := map[string]string{}
	for _, k := range ciProviderKeys {
		env[k] = "true"
	}
	for i := 0; i < 50; i++ {
		id, ok := newTestAdapter().Detect(env)
		if !ok || id != ADPAdapterID {
			t.Fatalf("iteration %d Detect = (%q,%v)", i, id, ok)
		}
	}
}
