// Abate —— 止息混沌 · 测试插件生态
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

// Package test 实现测试插件接口、注册表与矩阵编排器。
package test

import (
	"context"
	"time"

	"github.com/qingdujian/abate/core"
)

// TestStatus 表示测试结果状态。
type TestStatus string

const (
	StatusPass TestStatus = "PASS"
	StatusFail TestStatus = "FAIL"
	StatusSkip TestStatus = "SKIP"
)

// TestResult 是单个测试插件的运行结果。
type TestResult struct {
	Name     string             `json:"name"`
	Status   TestStatus         `json:"status"`
	Duration time.Duration      `json:"duration"`
	Details  string             `json:"details,omitempty"`
	Metrics  map[string]float64 `json:"metrics,omitempty"`
}

// TestPlugin 是所有测试插件必须实现的接口。
type TestPlugin interface {
	// Capability 返回能力声明（复用 core.Capability）。
	Capability() core.Capability
	// Setup 准备测试环境，返回清理函数。
	Setup(ctx context.Context) (cleanup func(), err error)
	// Run 执行测试，返回结果。
	Run(ctx context.Context, target string) TestResult
	// Teardown 清理测试残留。
	Teardown(ctx context.Context) error
}
