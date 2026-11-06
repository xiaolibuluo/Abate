// Abate —— 止息混沌 · 日志格式化器
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package log

import (
	"fmt"
	"time"
)

// textFormatter 文本格式化器。
type textFormatter struct{}

func (textFormatter) Format(e Entry) []byte {
	s := fmt.Sprintf("[%s] %s %s", e.Level, e.Time.Format("2006-01-02T15:04:05.000Z07:00"), e.Message)
	// 精准代码位置：包名/文件名:行号 函数名
	if e.Caller.File != "" {
		s += fmt.Sprintf(" @%s", e.Caller.String())
	}
	for _, f := range e.Fields {
		s += fmt.Sprintf(" %s=%v", f.Key, f.Value)
	}
	s += "\n"
	return []byte(s)
}

// jsonFormatter JSON 格式化器。
type jsonFormatter struct{}

func (jsonFormatter) Format(e Entry) []byte {
	s := fmt.Sprintf(`{"level":"%s","time":"%s","msg":%q`,
		e.Level, e.Time.Format(time.RFC3339Nano), e.Message)
	if e.Caller.File != "" {
		s += fmt.Sprintf(","+`"caller":{"pkg":%q,"file":%q,"line":%d,"func":%q}`,
			e.Caller.Pkg, e.Caller.File, e.Caller.Line, e.Caller.Func)
	}
	for _, f := range e.Fields {
		s += fmt.Sprintf(",%q:%q", f.Key, fmt.Sprint(f.Value))
	}
	s += "}\n"
	return []byte(s)
}
