// Abate —— 止息混沌 · 依赖图分析（精层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"sort"
)

// Cycle 是依赖图中检测到的环。
type Cycle struct {
	Files []string
}

// String 返回环的路径描述。
func (c Cycle) String() string {
	if len(c.Files) == 0 {
		return ""
	}
	s := c.Files[0]
	for _, f := range c.Files[1:] {
		s += " -> " + f
	}
	return s + " -> " + c.Files[0]
}

// HubFile 是被多个文件依赖的枢纽文件。
type HubFile struct {
	File     string
	InDegree int
}

// GraphAnalysis 是依赖图分析结果。
type GraphAnalysis struct {
	TotalFiles  int
	ParsedFiles int
	TotalEdges  int
	Cycles      []Cycle
	Hubs        []HubFile
	Orphans    []string
	MaxDepth    int
}

// Analyze 对依赖图执行环检测、枢纽发现、孤儿识别和深度计算。
func (g *DepGraph) Analyze() GraphAnalysis {
	a := GraphAnalysis{TotalFiles: len(g.Files), ParsedFiles: g.Parsed}
	// 统计边数和入度。
	inDegree := map[string]int{}
	for _, f := range g.Files {
		inDegree[f] = 0
	}
	for _, deps := range g.Edges {
		a.TotalEdges += len(deps)
		for dep := range deps {
			inDegree[dep]++
		}
	}
	// 环检测：DFS 三色标记。
	a.Cycles = g.detectCycles()
	// 枢纽文件：入度 >= 3 的文件，按入度降序。
	for f, d := range inDegree {
		if d >= 3 {
			a.Hubs = append(a.Hubs, HubFile{File: f, InDegree: d})
		}
	}
	sort.Slice(a.Hubs, func(i, j int) bool { return a.Hubs[i].InDegree > a.Hubs[j].InDegree })
	if len(a.Hubs) > 10 {
		a.Hubs = a.Hubs[:10]
	}
	// 孤儿文件：无入边也无出边。
	for _, f := range g.Files {
		if inDegree[f] == 0 && len(g.Edges[f]) == 0 {
			a.Orphans = append(a.Orphans, f)
		}
	}
	sort.Strings(a.Orphans)
	// 最大依赖深度：从入度为 0 的节点出发 DFS。
	a.MaxDepth = g.maxDepth(inDegree)
	return a
}

// detectCycles 使用 DFS 三色标记法检测所有环。
func (g *DepGraph) detectCycles() []Cycle {
	const (
		white = 0 // 未访问
		gray  = 1 // 访问中（在当前 DFS 路径上）
		black = 2 // 已完成
	)
	color := map[string]int{}
	stack := []string{}
	var cycles []Cycle
	var dfs func(f string)
	dfs = func(f string) {
		color[f] = gray
		stack = append(stack, f)
		for dep := range g.Edges[f] {
			switch color[dep] {
			case white:
				dfs(dep)
			case gray:
				// 找到环：从 stack 中 dep 的位置开始截取。
				for i := len(stack) - 1; i >= 0; i-- {
					if stack[i] == dep {
						c := make([]string, len(stack)-i)
						copy(c, stack[i:])
						cycles = append(cycles, Cycle{Files: c})
						break
					}
				}
			}
		}
		stack = stack[:len(stack)-1]
	color[f] = black
	}
	for _, f := range g.Files {
		if color[f] == white {
			dfs(f)
		}
	}
	return cycles
}

// maxDepth 计算从入度为 0 的节点出发的最大依赖链深度。
func (g *DepGraph) maxDepth(inDegree map[string]int) int {
	memo := map[string]int{}
	var depth func(f string) int
	depth = func(f string) int {
		if d, ok := memo[f]; ok {
			return d
		}
		deps := g.Edges[f]
		if len(deps) == 0 {
			memo[f] = 1
			return 1
		}
		max := 0
		for dep := range deps {
			if d := depth(dep); d > max {
				max = d
			}
		}
		memo[f] = max + 1
		return max + 1
	}
	result := 0
	for _, f := range g.Files {
		if inDegree[f] == 0 {
			if d := depth(f); d > result {
				result = d
			}
		}
	}
	return result
}
