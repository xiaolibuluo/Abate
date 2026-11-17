/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"testing"

	"abate/lang/air"
)

func TestDiagnosticSinkOrderingAndCap(t *testing.T) {
	t.Parallel()
	sink := NewDiagnosticSink("a.fk", "python")
	if sink.Tag() != "PY" {
		t.Fatalf("tag = %q, want PY", sink.Tag())
	}
	if err := sink.Warning(2, 30, 3, 4, "late"); err != nil {
		t.Fatal(err)
	}
	if err := sink.Error(1, 5, 1, 2, "early"); err != nil {
		t.Fatal(err)
	}
	if err := sink.Warning(3, 5, 1, 9, "same offset"); err != nil {
		t.Fatal(err)
	}
	list := sink.Diagnostics()
	if len(list) != 3 || list[0].Code != "LANG-PY-E-001" ||
		list[1].Code != "LANG-PY-W-003" || list[2].Code != "LANG-PY-W-002" {
		t.Fatalf("diagnostic ordering = %v", list)
	}
	if !sink.IsError() || sink.Len() != 3 || sink.Truncated() {
		t.Fatal("sink flags are wrong")
	}
	full := NewDiagnosticSink("b.fk", "fake")
	var limitError error
	for i := 0; i < air.MaxDiagnostics+5; i++ {
		if err := full.Warning(i, i, 1, 1, "msg"); err != nil {
			limitError = err
			break
		}
	}
	if !errors.Is(limitError, ErrDiagnosticLimit) ||
		full.Len() != air.MaxDiagnostics || !full.Truncated() ||
		len(full.Diagnostics()) != air.MaxDiagnostics || full.IsError() {
		t.Fatalf("cap handling failed: err=%v len=%d truncated=%v",
			limitError, full.Len(), full.Truncated())
	}
}

func TestDiagnosticSinkEdgeCases(t *testing.T) {
	t.Parallel()
	if NewDiagnosticSink("c.fk", "!!").Tag() != "LANG" {
		t.Fatal("unusable tag must fall back to LANG")
	}
	var nilSink *DiagnosticSink
	if nilSink.IsError() || nilSink.Len() != 0 || nilSink.Truncated() ||
		len(nilSink.Diagnostics()) != 0 || nilSink.Tag() != "LANG" {
		t.Fatal("nil sink must be inert")
	}
	if err := nilSink.Error(1, 1, 1, 1, "x"); !errors.Is(err, ErrDiagnosticLimit) {
		t.Fatalf("nil sink add error = %v", err)
	}
}
