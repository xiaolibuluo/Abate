/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"testing"

	"abate/internal/corpus"
	"abate/internal/e2e"
)

// e2eFrontendsForTest builds the real seven-language frontend table.
func e2eFrontendsForTest() (e2e.FrontendTable, error) {
	return e2e.NewFrontendTable()
}

// corpusEntryForTest makes a minimal manifest entry for one language and path.
func corpusEntryForTest(lang, relPath string) corpus.Entry {
	return corpus.Entry{Language: lang, RelPath: relPath}
}

// TestHelpersSanity guards the test helpers themselves against silent drift.
func TestHelpersSanity(t *testing.T) {
	ft, err := e2eFrontendsForTest()
	if err != nil {
		t.Fatalf("frontend table: %v", err)
	}
	if len(ft.Languages()) != len(reportLanguages) {
		t.Fatalf("frontend languages=%d want %d",
			len(ft.Languages()), len(reportLanguages))
	}
}
