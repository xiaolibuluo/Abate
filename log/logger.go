// Abate —— 止息混沌 · 结构化日志引擎
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

// Package log 提供 Abate 框架的全链路结构化日志能力。
// 支持 DEBUG/INFO/WARN/ERROR 四级日志、JSON/Text 双格式输出、
// stdout/file/both 多输出目标、上下文字段注入、caller 精准追踪
// （包名/文件名/行号/函数名），出错可立即定位代码位置。
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level 日志级别。
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelOff
)

// String 返回级别字符串表示。
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelOff:
		return "OFF"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel 从字符串解析日志级别。
func ParseLevel(s string) Level {
	switch s {
	case "debug", "DEBUG":
		return LevelDebug
	case "info", "INFO":
		return LevelInfo
	case "warn", "WARN":
		return LevelWarn
	case "error", "ERROR":
		return LevelError
	case "off", "OFF":
		return LevelOff
	default:
		return LevelInfo
	}
}

// Field 是结构化日志的字段。
type Field struct {
	Key   string
	Value any
}

// CallerInfo 记录日志调用的精确代码位置。
type CallerInfo struct {
	Pkg      string // 包名，如 github.com/qingdujian/abate/sim
	File     string // 文件名，如 sim/sim.go
	Line     int    // 行号
	Func     string // 函数名，如 (*Simulator).Run
}

// String 返回可读的调用位置字符串。
func (c CallerInfo) String() string {
	return fmt.Sprintf("%s:%d %s", c.File, c.Line, c.Func)
}

// Entry 是一条日志记录。
type Entry struct {
	Time    time.Time
	Level   Level
	Message string
	Fields  []Field
	Caller  CallerInfo
}

// Formatter 日志格式化器接口。
type Formatter interface {
	Format(entry Entry) []byte
}

// Logger 结构化日志器接口。
// 各模块通过此接口输出全链路日志。
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	With(fields ...Field) Logger
	SetLevel(level Level)
	Level() Level
}

// stdLogger 标准日志器实现。
type stdLogger struct {
	mu        sync.Mutex
	level     Level
	formatter Formatter
	writers   []io.Writer
	fields    []Field
}

// New 创建日志器，输出到指定 writer。
func New(w io.Writer, level Level, format string) Logger {
	var f Formatter = textFormatter{}
	if format == "json" {
		f = jsonFormatter{}
	}
	return &stdLogger{
		level:     level,
		formatter: f,
		writers:   []io.Writer{w},
	}
}

// NewFile 创建输出到文件的日志器。
func NewFile(path string, level Level, format string) (Logger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("log: open %s: %w", path, err)
	}
	return New(f, level, format), nil
}

// NewDual 创建同时输出到 stdout 和文件的日志器。
func NewDual(path string, level Level, format string) (Logger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("log: open %s: %w", path, err)
	}
	var fmtl Formatter = textFormatter{}
	if format == "json" {
		fmtl = jsonFormatter{}
	}
	return &stdLogger{
		level:     level,
		formatter: fmtl,
		writers:   []io.Writer{os.Stdout, f},
	}, nil
}

// Default 创建默认日志器（stdout, INFO, text）。
func Default() Logger {
	return New(os.Stdout, LevelInfo, "text")
}

// Nop 返回空操作日志器。
func Nop() Logger {
	return &stdLogger{level: LevelOff, formatter: textFormatter{}, writers: []io.Writer{io.Discard}}
}

// captureCaller 获取调用方的精确代码位置。
// 自动跳过 log 包内部所有帧（captureCaller/log/Debug/Info/Warn/Error/adapter），
// 无论经过多少层 adapter 包装，都能定位到真正的业务调用方代码。
func captureCaller(skip int) CallerInfo {
	// 从 skip 开始向上遍历，跳过所有 log 包内部的帧
	for i := skip; i < skip+10; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// 跳过 log 包自身的帧
		fn := runtime.FuncForPC(pc)
		if fn != nil && strings.Contains(fn.Name(), "/log") {
			continue
		}
		// 提取相对于模块根的路径
		if idx := strings.Index(file, "qingdujian/abate/"); idx >= 0 {
			file = file[idx+len("qingdujian/abate/"):]
		} else {
			file = filepath.Base(file)
		}
		funcName := ""
		pkg := ""
		if fn != nil {
			funcName = fn.Name()
			if dot := strings.LastIndex(funcName, "."); dot >= 0 {
				pkg = funcName[:dot]
			}
		}
		return CallerInfo{
			Pkg:  pkg,
			File: file,
			Line: line,
			Func: funcName,
		}
	}
	return CallerInfo{}
}

func (l *stdLogger) log(level Level, msg string, fields []Field) {
	if level < l.level {
		return
	}
	// skip=3: captureCaller -> log -> Debug/Info/Warn/Error -> 业务代码
	caller := captureCaller(3)
	allFields := make([]Field, 0, len(l.fields)+len(fields))
	allFields = append(allFields, l.fields...)
	allFields = append(allFields, fields...)
	entry := Entry{
		Time:    time.Now(),
		Level:   level,
		Message: msg,
		Fields:  allFields,
		Caller:  caller,
	}
	data := l.formatter.Format(entry)
	l.mu.Lock()
	for _, w := range l.writers {
		w.Write(data)
	}
	l.mu.Unlock()
}

func (l *stdLogger) Debug(msg string, fields ...Field) { l.log(LevelDebug, msg, fields) }
func (l *stdLogger) Info(msg string, fields ...Field)  { l.log(LevelInfo, msg, fields) }
func (l *stdLogger) Warn(msg string, fields ...Field)  { l.log(LevelWarn, msg, fields) }
func (l *stdLogger) Error(msg string, fields ...Field) { l.log(LevelError, msg, fields) }

func (l *stdLogger) With(fields ...Field) Logger {
	newFields := make([]Field, 0, len(l.fields)+len(fields))
	newFields = append(newFields, l.fields...)
	newFields = append(newFields, fields...)
	return &stdLogger{
		level:     l.level,
		formatter: l.formatter,
		writers:   l.writers,
		fields:    newFields,
	}
}

func (l *stdLogger) SetLevel(level Level) {
	l.mu.Lock()
	l.level = level
	l.mu.Unlock()
}

func (l *stdLogger) Level() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// F 创建字段的便捷函数。
func F(key string, value any) Field {
	return Field{Key: key, Value: value}
}
