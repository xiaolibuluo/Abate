// Abate —— 止息混沌 · 日志适配器
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package log

import (
	"github.com/qingdujian/abate/core"
)

// coreLogger 桥接 log.Logger 到 core.Logger 接口。
// 注意：adapter 多一层调用，caller skip 需要额外 +1。
type coreLogger struct {
	inner Logger
}

// Adapt 将 log.Logger 适配为 core.Logger 接口。
func Adapt(l Logger) core.Logger {
	return &coreLogger{inner: l}
}

func toFields(fields []core.LogField) []Field {
	out := make([]Field, len(fields))
	for i, f := range fields {
		out[i] = Field{Key: f.Key, Value: f.Value}
	}
	return out
}

func (c *coreLogger) Debug(msg string, fields ...core.LogField) {
	c.inner.Debug(msg, toFields(fields)...)
}

func (c *coreLogger) Info(msg string, fields ...core.LogField) {
	c.inner.Info(msg, toFields(fields)...)
}

func (c *coreLogger) Warn(msg string, fields ...core.LogField) {
	c.inner.Warn(msg, toFields(fields)...)
}

func (c *coreLogger) Error(msg string, fields ...core.LogField) {
	c.inner.Error(msg, toFields(fields)...)
}

func (c *coreLogger) With(fields ...core.LogField) core.Logger {
	return &coreLogger{inner: c.inner.With(toFields(fields)...)}
}

func (c *coreLogger) SetLevel(level int) {
	c.inner.SetLevel(Level(level))
}

func (c *coreLogger) Level() int {
	return int(c.inner.Level())
}
