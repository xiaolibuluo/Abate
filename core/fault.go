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

import "fmt"

// FaultClass 是 CES 第 1 章定义的故障分类编码，永久固定，只可标记废弃，不得回收复用。
// FaultClass is the fault classification code defined in CES Chapter 1.
// Codes are permanent: they may be deprecated, never recycled.
type FaultClass string

const (
	// F1Latency 延迟：调用、I/O、调度的时延注入与抖动。
	F1Latency FaultClass = "F1"
	// F2Crash 崩溃：进程、实例、组件的终止与重启。
	F2Crash FaultClass = "F2"
	// F3Partition 分区：网络断连、丢包、DNS 污染、路由隔离。
	F3Partition FaultClass = "F3"
	// F4Exhaustion 耗尽：CPU、内存、磁盘、文件描述符、连接池耗尽。
	F4Exhaustion FaultClass = "F4"
	// F5Corruption 损坏：数据位翻转、乱序、重复、丢失。
	F5Corruption FaultClass = "F5"
	// F6ClockSkew 时钟漂移：系统时钟偏移与回拨。
	F6ClockSkew FaultClass = "F6"
)

// FaultClasses 返回全部合法故障分类，顺序固定（确定性要求）。
// FaultClasses returns all valid fault classes in a fixed order.
func FaultClasses() []FaultClass {
	return []FaultClass{
		F1Latency, F2Crash, F3Partition,
		F4Exhaustion, F5Corruption, F6ClockSkew,
	}
}

// Valid 报告分类编码是否合法。
// Valid reports whether the class code is defined by the specification.
func (c FaultClass) Valid() bool {
	switch c {
	case F1Latency, F2Crash, F3Partition,
		F4Exhaustion, F5Corruption, F6ClockSkew:
		return true
	}
	return false
}

// Fault 描述一次故障注入：分类编码 + 该分类下的参数集合（CES-2.3.1）。
// Fault describes one fault injection: a class code plus the parameter set
// defined for that class (CES-2.3.1).
type Fault struct {
	// Class 故障分类编码（CES-2.3.1 强制使用第 1 章编码）。
	Class FaultClass `yaml:"type" json:"type"`
	// Params 分类参数，键值均为字符串以保证跨插件可序列化。
	// Params holds class parameters as strings for cross-plugin serialization.
	Params map[string]string `yaml:"params" json:"params"`
}

// Validate 校验故障定义：编码合法、参数非空。
// Validate checks the fault definition: a valid class code and a non-empty
// parameter set.
func (f Fault) Validate() error {
	if !f.Class.Valid() {
		return fmt.Errorf("fault.type: invalid class %q (CES-2.3.1)", f.Class)
	}
	if len(f.Params) == 0 {
		return fmt.Errorf("fault.params: must not be empty (CES-2.3.1)")
	}
	return nil
}
