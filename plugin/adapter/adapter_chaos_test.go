/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"sync"
	"testing"
)

// TestConcurrentDetectAndEmitChaos exercises the air-gapped adapter from many
// goroutines to surface hidden shared state or data races.
func TestConcurrentDetectAndEmitChaos(t *testing.T) {
	a := newTestAdapter()
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			env := map[string]string{"CI": "true"}
			id, ok := a.Detect(env)
			if !ok || id != ADPAdapterID {
				t.Errorf("concurrent Detect = (%q,%v)", id, ok)
			}
			if err := a.Emit([]byte("payload")); err != nil {
				t.Errorf("concurrent Emit: %v", err)
			}
			if err := a.Emit([]byte{0xff}); !errors.Is(err, ErrAdapterPayloadRejected) {
				t.Errorf("concurrent invalid Emit: %v", err)
			}
			_ = a.Manifest()
			_ = a.Type()
		}()
	}
	wg.Wait()
}

// TestDetectDeterministic100 proves detection is a pure function of the
// environment across many independent adapters.
func TestDetectDeterministic100(t *testing.T) {
	env := map[string]string{"CI": "true", "GITHUB_ACTIONS": "false"}
	for i := 0; i < 100; i++ {
		id, ok := newTestAdapter().Detect(env)
		if !ok || id != ADPAdapterID {
			t.Fatalf("iteration %d Detect = (%q,%v)", i, id, ok)
		}
	}
}
