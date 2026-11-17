/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"slices"
	"strings"

	"abate/lang/air"
)

// jsOrderChildren 返回按源码位置升序排列的子节点切片.
//
// AIR 的兄弟节点按位置升序存放，前端若按解析顺序追加子节点，就必须先排序。
// 否则像 import 语句这种先补出子表达式再补出关键字的归约，会产出兄弟区间
// 逆序的树，位置往返校验会把它判成兄弟重叠。
func jsOrderChildren(children []air.Node) []air.Node {
	out := slices.Clone(children)
	slices.SortStableFunc(out, func(a, b air.Node) int {
		ap, bp := a.Pos(), b.Pos()
		if ap.Offset != bp.Offset {
			if ap.Offset < bp.Offset {
				return -1
			}
			return 1
		}
		if ap.EndOffset != bp.EndOffset {
			if ap.EndOffset < bp.EndOffset {
				return -1
			}
			return 1
		}
		return strings.Compare(air.AttributeOf(a), air.AttributeOf(b))
	})
	return out
}

// jsCompareNodes 与 jsOrderChildren 使用同一排序键，便于二分插入复用.
func jsCompareNodes(a, b air.Node) int {
	ap, bp := a.Pos(), b.Pos()
	if ap.Offset != bp.Offset {
		if ap.Offset < bp.Offset {
			return -1
		}
		return 1
	}
	if ap.EndOffset != bp.EndOffset {
		if ap.EndOffset < bp.EndOffset {
			return -1
		}
		return 1
	}
	return strings.Compare(air.AttributeOf(a), air.AttributeOf(b))
}

// jsInsertOrdered 把单个节点插入已按 jsOrderChildren 排好序的切片.
//
// 插入点取「首个严格大于该键」的位置，与对 append(kept, node) 做稳定排序的
// 结果一致（新节点原本排在末尾，同键元素保持既有相对次序）。这样每轮归约
// 只做一次二分查找与一次移动，避免对扁平大数组反复做 O(k log k) 全排序。
func jsInsertOrdered(kept []air.Node, node air.Node) []air.Node {
	lo, hi := 0, len(kept)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if jsCompareNodes(kept[mid], node) <= 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	at := lo
	out := make([]air.Node, 0, len(kept)+1)
	out = append(out, kept[:at]...)
	out = append(out, node)
	out = append(out, kept[at:]...)
	return out
}

// jsNestChildren 把被其它子节点完全包含的子节点嵌入到那个子节点内部.
//
// AIR 的兄弟节点必须互不重叠，所以像 TS 类体这种「容器节点与容器内的类型
// 注记」被当成兄弟节点拼在一起时会违反该不变量。嵌入之后信息不丢失，被
// 重建节点的语义注记会一并复制过去。
func jsNestChildren(children []air.Node) []air.Node {
	return jsNestDepth(children, 0)
}

// jsNestDepth 递归地完成嵌入，depth 限制嵌套层数以保证终止.
func jsNestDepth(children []air.Node, depth int) []air.Node {
	out := jsOrderChildren(children)
	if depth > 8 {
		return out
	}
	for {
		split := jsSplitNested(out)
		if !split.changed {
			return out
		}
		merged := append(split.host.Children(), split.nested...)
		rebuilt := jsRebuild(split.host, jsNestDepth(merged, depth+1))
		out = jsInsertOrdered(split.kept, rebuilt)
	}
}

// jsSplitOutcome 是 jsSplitNested 的返回结果，把多个返回值收敛为单一结构体，
// 以满足 function-result-limit（单函数返回结果不超过 3 个）的精准红线。
type jsSplitOutcome struct {
	nested  []air.Node
	host    air.Node
	kept    []air.Node
	changed bool
}

// Span table caches the source intervals of one sibling group.
//
// 区间由 air.Span 提供，每次调用都要走 Node 接口的 Pos() 方法；对压缩单行
// 大数组（数万个兄弟节点）而言，O(k²) 次接口调用是解析耗时的主要来源。
// 预计算把接口调用降到每个节点恰好一次，内层循环只读两个 int 数组。
type jsSpanTable struct {
	starts       []int
	ends         []int
	suffixMinEnd []int
}

// newJsSpanTable 预计算全部节点的 start/end 以及 end 的后缀最小值.
//
// 兄弟节点已按 (start, end) 升序排列，因此对候选 i 而言：
//   - j > i 的节点必然满足 start[j] >= start[i]，其中存在被包含节点当且仅当
//     suffixMinEnd[i+1] <= end[i]；
//   - j < i 的节点只有在 start 与 i 相等时才可能满足 start[j] >= start[i]，
//     此时按 end 升序排列保证 end[j] <= end[i]，必然被包含。
//
// 这两条判据与原双重循环的「存在其它节点满足 otherStart >= start &&
// otherEnd <= end」完全等价，但把每次探测从 O(k) 降到 O(1)。
func newJsSpanTable(children []air.Node) jsSpanTable {
	size := len(children)
	table := jsSpanTable{
		starts:       make([]int, size),
		ends:         make([]int, size),
		suffixMinEnd: make([]int, size+1),
	}
	for i, child := range children {
		table.starts[i], table.ends[i] = air.Span(child)
	}
	table.suffixMinEnd[size] = int(^uint(0) >> 1)
	for i := size - 1; i >= 0; i-- {
		minEnd := table.ends[i]
		if table.suffixMinEnd[i+1] < minEnd {
			minEnd = table.suffixMinEnd[i+1]
		}
		table.suffixMinEnd[i] = minEnd
	}
	return table
}

// containsOther 报告候选 i 是否包含至少一个其它兄弟节点.
func (t jsSpanTable) containsOther(i int) bool {
	if t.ends[i] <= t.starts[i] {
		return false
	}
	if i+1 < len(t.starts) && t.suffixMinEnd[i+1] <= t.ends[i] {
		return true
	}
	return i > 0 && t.starts[i-1] == t.starts[i]
}

// jsSplitNested 找出第一个包含其它子节点的宿主节点，并把它与被包含节点分开.
//
// 返回的 changed 为 false 表示没有节点互相包含，调用方可以直接结束。
//
// 复杂度修复：曾经对每个候选 i 都在内层循环里对全部 k 个兄弟调用 air.Span，
// 形成 O(k²) 次 Node 接口调用（叠加外层归约轮次可到 O(k³)）。现在区间只
// 预计算一次，先用后缀最小值 O(1) 判断候选是否需要展开，仅在确找到一个宿主
// 时才做一次 O(k) 扫描收集被包含节点，因此「扁平无嵌套」的大数组单轮即返回。
func jsSplitNested(children []air.Node) jsSplitOutcome {
	table := newJsSpanTable(children)
	for i := range children {
		if !table.containsOther(i) {
			continue
		}
		start, end := table.starts[i], table.ends[i]
		nested := make([]air.Node, 0, 2)
		kept := make([]air.Node, 0, len(children)-1)
		for j, other := range children {
			if j == i {
				continue
			}
			if table.starts[j] >= start && table.ends[j] <= end {
				nested = append(nested, other)
			} else {
				kept = append(kept, other)
			}
		}
		if len(nested) == 0 {
			continue
		}
		return jsSplitOutcome{nested: nested, host: children[i], kept: kept, changed: true}
	}
	return jsSplitOutcome{kept: children}
}

// jsRebuild 复制一个节点并用给定子节点替换原有子节点，语义注记一并复制.
func jsRebuild(node air.Node, children []air.Node) air.Node {
	rebuilt := air.NewNodeAt(node.Kind(), node.Pos(), node.End(),
		air.AttributeOf(node), jsOrderChildren(children)...)
	semantics := air.SemanticsOf(node)
	for _, key := range semantics.Keys() {
		value, _ := semantics.Get(key)
		_ = air.Annotate(rebuilt, key, value)
	}
	return rebuilt
}
