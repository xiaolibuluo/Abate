// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

import (
	"context"
	"fmt"
)

// ProtocolVersion 是插件协议当前主版本（CES-7.1 握手协商）。
// ProtocolVersion is the current plugin protocol major version.
const ProtocolVersion = "v1"

// Capability 是插件启动时的能力声明（CES-7.1）。
// Capability is the plugin capability declaration made at startup (CES-7.1).
type Capability struct {
	// Name 插件名。
	Name string `json:"name"`
	// Version 插件自身语义化版本。
	Version string `json:"version"`
	// Protocol 插件支持的协议主版本；与核心不一致即拒绝加载。
	Protocol string `json:"protocol"`
	// FaultClasses 插件支持注入的故障分类编码（F1-F6）。
	FaultClasses []FaultClass `json:"faultClasses,omitempty"`
	// TargetKinds 插件支持的目标寻址模式。
	TargetKinds []TargetKind `json:"targetKinds,omitempty"`
}

// CheckProtocol 执行版本协商：不一致返回带错误码语义的明确错误。
// CheckProtocol performs version negotiation; mismatches are refused with an
// explicit error.
func (c Capability) CheckProtocol() error {
	if c.Protocol != ProtocolVersion {
		return fmt.Errorf("plugin %q: protocol %q incompatible with %q (CES-7.1, E1001)",
			c.Name, c.Protocol, ProtocolVersion)
	}
	return nil
}

// Injection 是一次故障注入的句柄，用于幂等回滚（CES-2.4.1）。
// Injection is a handle to one fault injection for idempotent rollback.
type Injection struct {
	// ID 注入实例唯一标识。
	ID string `json:"id"`
	// ExperimentID 所属实验。
	ExperimentID string `json:"experimentId"`
}

// Injector 执行故障注入与恢复（CES-2.4.1 幂等；CES-3.1.2 无残留）。
// Injector performs fault injection and recovery; idempotent (CES-2.4.1),
// residue-free (CES-3.1.2).
type Injector interface {
	// Capability 返回能力声明。
	Capability() Capability
	// Inject 开始注入，返回句柄；重复调用同一参数必须不产生叠加效果。
	Inject(ctx context.Context, e Experiment) (Injection, error)
	// Recover 回滚指定注入；必须可重复调用且最终无残留。
	Recover(ctx context.Context, in Injection) error
}

// Sample 是一次 SLI 采样点。
// Sample is one SLI sample point.
type Sample struct {
	// SLI 指标名。
	SLI string `json:"sli"`
	// Value 万分位比率或毫秒整数，含义按指标定义。
	Value int64 `json:"value"`
	// TimestampMs 采样时刻（Unix 毫秒）。
	TimestampMs int64 `json:"timestampMs"`
}

// Probe 采集稳态与红线所需的 SLI 样本（CES 第 5 章观测）。
// Probe collects SLI samples for steady-state and red-line observation.
type Probe interface {
	// Capability 返回能力声明。
	Capability() Capability
	// Observe 对目标执行一次采样。
	Observe(ctx context.Context, t Target) (Sample, error)
}

// Validator 执行实验前准入与实验后准出校验（CES 第 5、6 章判定）。
// Validator performs pre-flight admission and post-flight exit checks.
type Validator interface {
	// Capability 返回能力声明。
	Capability() Capability
	// Preflight 准入检查；返回的每个问题是阻止实验的硬伤（CES-3.1.4）。
	Preflight(ctx context.Context, e Experiment) ([]string, error)
	// Postflight 准出检查：恢复验证与残留扫描（CES-3.1.2）。
	Postflight(ctx context.Context, e Experiment) ([]string, error)
}

// Reporter 产出结构化报告（CES 第 6 章：自包含、可脱机验证）。
// Reporter produces structured reports (CES Chapter 6).
type Reporter interface {
	// Capability 返回能力声明。
	Capability() Capability
	// Emit 输出一次评分运行的报告；返回产物定位符。
	Emit(ctx context.Context, s Score) (string, error)
}

// LogField 是结构化日志的字段，与 log.Field 对齐。
// LogField is a structured log field, aligned with log.Field.
type LogField struct {
	Key   string
	Value any
}

// Logger 全链路结构化日志接口（CES-9.6 可观测性）。
// 各模块通过依赖注入获取日志器，实现完整闭环的详细日志输出链。
// Logger is the full-chain structured logging interface (CES-9.6 observability).
// Modules obtain a logger via dependency injection for complete closed-loop logging.
type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	With(fields ...LogField) Logger
	SetLevel(level int)
	Level() int
}
