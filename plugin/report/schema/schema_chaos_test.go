/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"sync"
	"testing"
)

// TestConcurrentValidationChaos validates one compiled OASIS schema from many
// goroutines. The regexp cache is a sync.Map, so this also proves concurrent
// pattern compilation is race-free.
func TestConcurrentValidationChaos(t *testing.T) {
	root, err := LoadSchema(oasisSchemaPath)
	if err != nil {
		t.Fatalf("LoadSchema: %v", err)
	}
	doc := decodeDocument(t, sarifDocument)
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if err := root.Validate(doc); err != nil {
				t.Errorf("concurrent OASIS validation: %v", err)
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentPatternCacheChaos forces many distinct patterns through the
// cache simultaneously, exercising both the Load and LoadOrStore paths.
func TestConcurrentPatternCacheChaos(t *testing.T) {
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			s := &Schema{typ: "string", pattern: "^a+$"}
			if err := s.Validate("aaa"); err != nil {
				t.Errorf("iteration %d pattern validation: %v", idx, err)
			}
			if err := s.Validate("b"); err == nil {
				t.Errorf("iteration %d accepted a non-match", idx)
			}
		}(i)
	}
	wg.Wait()
}

// TestConcurrentLoadSchemaChaos compiles the bundled schema concurrently.
func TestConcurrentLoadSchemaChaos(t *testing.T) {
	const n = 8
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			root, err := LoadSchema(oasisSchemaPath)
			if err != nil {
				t.Errorf("concurrent LoadSchema: %v", err)
				return
			}
			if err := root.Validate(decodeDocument(t, sarifDocument)); err != nil {
				t.Errorf("concurrent document validation: %v", err)
			}
		}()
	}
	wg.Wait()
}
