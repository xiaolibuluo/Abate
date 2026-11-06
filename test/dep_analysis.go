// Abate —— 止息混沌 · 依赖分析插件（精层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// DepAnalysisPlugin 跨文件依赖分析插件（精层）。
type DepAnalysisPlugin struct{}

func (DepAnalysisPlugin) Capability() core.Capability {
	return core.Capability{Name: "dep-analysis", Version: "0.1.0", Protocol: "1"}
}

func (DepAnalysisPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (DepAnalysisPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "dep-analysis", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	// 根据规模决定扫描上限（稳层协同）。
	maxFiles := 0
	plan := PlanScan(ctx, target)
	if plan.Mode == ModeSampled {
		maxFiles = plan.FileBudget
	} else if plan.Mode == ModeMinimal {
		maxFiles = plan.FileBudget
	}
	g := BuildDepGraph(target, maxFiles)
	a := g.Analyze()
	dur := time.Since(start)
	// 环检测：有环则 FAIL。
	status := StatusPass
	if len(a.Cycles) > 0 {
		status = StatusFail
	}
	// 构建摘要。
	var b strings.Builder
	fmt.Fprintf(&b, "files=%d parsed=%d edges=%d cycles=%d hubs=%d orphans=%d depth=%d",
		a.TotalFiles, a.ParsedFiles, a.TotalEdges, len(a.Cycles), len(a.Hubs), len(a.Orphans), a.MaxDepth)
	if len(a.Cycles) > 0 {
		b.WriteString(" | cycles: ")
		for i, c := range a.Cycles {
			if i >= 3 {
				fmt.Fprintf(&b, "... +%d more", len(a.Cycles)-3)
				break
			}
			if i > 0 {
				b.WriteString("; ")
			}
			b.WriteString(c.String())
		}
	}
	if len(a.Hubs) > 0 {
		b.WriteString(" | hubs: ")
		for i, h := range a.Hubs {
			if i >= 3 {
				fmt.Fprintf(&b, "... +%d more", len(a.Hubs)-3)
				break
			}
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("%s(%d)", shortPath(g.Root, h.File), h.InDegree))
		}
	}
	return TestResult{
		Name:    "dep-analysis",
		Status:  status,
		Duration: dur,
		Details: b.String(),
		Metrics: map[string]float64{
			"files":   float64(a.TotalFiles),
			"edges":   float64(a.TotalEdges),
			"cycles":  float64(len(a.Cycles)),
			"hubs":    float64(len(a.Hubs)),
			"orphans": float64(len(a.Orphans)),
			"depth":   float64(a.MaxDepth),
		},
	}
}

func (DepAnalysisPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(DepAnalysisPlugin{})
}

// shortPath 返回相对根目录的短路径，超长时截断。
func shortPath(root, p string) string {
	rel, err := filepath.Rel(root, p)
	if err != nil {
		return filepath.Base(p)
	}
	return rel
}
