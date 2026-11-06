// Abate —— 止息混沌 · 结构化日志引擎测试
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestLevelString(t *testing.T) {
	cases := []struct {
		level Level
		want  string
	}{
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{LevelOff, "OFF"},
	}
	for _, c := range cases {
		if got := c.level.String(); got != c.want {
			t.Errorf("Level(%d).String() = %q, want %q", c.level, got, c.want)
		}
	}
}

func TestParseLevel(t *testing.T) {
	cases := []struct {
		input string
		want  Level
	}{
		{"debug", LevelDebug},
		{"INFO", LevelInfo},
		{"warn", LevelWarn},
		{"ERROR", LevelError},
		{"off", LevelOff},
		{"unknown", LevelInfo},
	}
	for _, c := range cases {
		if got := ParseLevel(c.input); got != c.want {
			t.Errorf("ParseLevel(%q) = %d, want %d", c.input, got, c.want)
		}
	}
}

func TestTextFormatter(t *testing.T) {
	f := textFormatter{}
	e := Entry{
		Level:   LevelInfo,
		Message: "test message",
		Fields:  []Field{F("key", "value")},
	}
	out := f.Format(e)
	s := string(out)
	if !strings.Contains(s, "[INFO]") {
		t.Errorf("text formatter missing level: %s", s)
	}
	if !strings.Contains(s, "test message") {
		t.Errorf("text formatter missing message: %s", s)
	}
	if !strings.Contains(s, "key=value") {
		t.Errorf("text formatter missing field: %s", s)
	}
}

func TestJSONFormatter(t *testing.T) {
	f := jsonFormatter{}
	e := Entry{
		Level:   LevelError,
		Message: "error occurred",
		Fields:  []Field{F("code", 500)},
	}
	out := f.Format(e)
	s := string(out)
	if !strings.Contains(s, `"level":"ERROR"`) {
		t.Errorf("json formatter missing level: %s", s)
	}
	if !strings.Contains(s, `"msg":"error occurred"`) {
		t.Errorf("json formatter missing message: %s", s)
	}
	if !strings.Contains(s, `"code":"500"`) {
		t.Errorf("json formatter missing field: %s", s)
	}
}

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelWarn, "text")
	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")
	s := buf.String()
	if strings.Contains(s, "debug msg") {
		t.Error("DEBUG should be filtered at WARN level")
	}
	if strings.Contains(s, "info msg") {
		t.Error("INFO should be filtered at WARN level")
	}
	if !strings.Contains(s, "warn msg") {
		t.Error("WARN should appear at WARN level")
	}
	if !strings.Contains(s, "error msg") {
		t.Error("ERROR should appear at WARN level")
	}
}

func TestLoggerWith(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelDebug, "text")
	child := logger.With(F("module", "sim"), F("exp", "test-1"))
	child.Info("experiment started")
	s := buf.String()
	if !strings.Contains(s, "module=sim") {
		t.Errorf("With fields missing module: %s", s)
	}
	if !strings.Contains(s, "exp=test-1") {
		t.Errorf("With fields missing exp: %s", s)
	}
	// parent should not have child fields
	buf.Reset()
	logger.Info("parent msg")
	s = buf.String()
	if strings.Contains(s, "module=sim") {
		t.Error("parent logger should not have child fields")
	}
}

func TestNopLogger(t *testing.T) {
	logger := Nop()
	logger.Debug("should not appear")
	logger.Info("should not appear")
	logger.Error("should not appear")
	if logger.Level() != LevelOff {
		t.Errorf("Nop level = %d, want %d", logger.Level(), LevelOff)
	}
}

func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelError, "text")
	logger.Info("should not appear")
	logger.SetLevel(LevelDebug)
	logger.Info("should appear")
	s := buf.String()
	if strings.Contains(s, "should not appear") {
		t.Error("INFO should be filtered at ERROR level")
	}
	if !strings.Contains(s, "should appear") {
		t.Error("INFO should appear after SetLevel(DEBUG)")
	}
}

func TestFieldHelper(t *testing.T) {
	f := F("count", 42)
	if f.Key != "count" {
		t.Errorf("F key = %q, want %q", f.Key, "count")
	}
	if f.Value != 42 {
		t.Errorf("F value = %v, want 42", f.Value)
	}
}
