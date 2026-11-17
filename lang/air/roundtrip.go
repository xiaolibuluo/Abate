/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"fmt"
	"slices"
)

// Walk visits a tree in deterministic pre-order without recursion.
//
// Returning false from visit prunes the children of the current node. The
// iterative stack keeps pathological tree depths from exhausting the goroutine
// stack.
func Walk(root Node, visit func(Node) bool) {
	if root == nil {
		return
	}
	stack := []Node{root}
	for len(stack) > 0 {
		last := len(stack) - 1
		current := stack[last]
		stack = stack[:last]
		if !visit(current) {
			continue
		}
		children := current.Children()
		for i := len(children) - 1; i >= 0; i-- {
			if children[i] != nil {
				stack = append(stack, children[i])
			}
		}
	}
}

// Collect returns every node of a unit in file order and pre-order.
func Collect(u *Unit) []Node {
	if u == nil {
		return nil
	}
	out := make([]Node, 0, 16)
	for _, root := range u.Roots {
		Walk(root, func(current Node) bool {
			out = append(out, current)
			return true
		})
	}
	return out
}

// Span returns the half-open byte interval of a node.
func Span(target Node) (start, end int) {
	if target == nil {
		return 0, 0
	}
	pos := target.Pos()
	return pos.Offset, pos.EndOffset
}

// CheckIntervals verifies that every node interval fits both the source file
// and its parent, and that siblings never overlap.
//
// Gaps between siblings are expected and allowed: they contain the separators
// and keywords that a neutral IR does not model as vertices.
func CheckIntervals(root Node, fileSize int) error {
	type frame struct {
		current Node
		parent  Position
	}
	stack := []frame{{current: root, parent: Position{EndOffset: fileSize}}}
	for len(stack) > 0 {
		last := len(stack) - 1
		top := stack[last]
		stack = stack[:last]
		if top.current == nil {
			continue
		}
		pos := top.current.Pos()
		if !pos.Valid() || pos.Offset < top.parent.Offset ||
			pos.EndOffset > top.parent.EndOffset {
			return fmt.Errorf("%w: %s escapes its parent or the file",
				ErrBadPosition, pos.String())
		}
		children := top.current.Children()
		previousEnd := pos.Offset
		for i := len(children) - 1; i >= 0; i-- {
			child := children[i]
			if child == nil {
				continue
			}
			stack = append(stack, frame{current: child, parent: pos})
		}
		for _, child := range children {
			if child == nil {
				continue
			}
			cpos := child.Pos()
			if cpos.Offset < previousEnd {
				return fmt.Errorf("%w: sibling overlap at %s",
					ErrBadPosition, cpos.String())
			}
			previousEnd = cpos.EndOffset
		}
	}
	return nil
}

// CompareUnits reports the first difference between two units.
//
// It compares package metadata, file metadata, sorted annotations, and every
// node field including its position, attribute, child count, and annotations.
func CompareUnits(left, right *Unit) error {
	if left == nil || right == nil {
		if left == right {
			return nil
		}
		return fmt.Errorf("%w: nil unit on one side", ErrUnitMismatch)
	}
	if left.Package.Canonical() != right.Package.Canonical() {
		return fmt.Errorf("%w: package %q vs %q", ErrUnitMismatch,
			left.Package.Canonical(), right.Package.Canonical())
	}
	if len(left.Files) != len(right.Files) {
		return fmt.Errorf("%w: %d vs %d files", ErrUnitMismatch,
			len(left.Files), len(right.Files))
	}
	for i := range left.Files {
		if left.Files[i].Canonical() != right.Files[i].Canonical() {
			return fmt.Errorf("%w: file %d %q vs %q", ErrUnitMismatch,
				i, left.Files[i].Canonical(), right.Files[i].Canonical())
		}
	}
	if err := compareAnnotations(left, right); err != nil {
		return err
	}
	if len(left.Roots) != len(right.Roots) {
		return fmt.Errorf("%w: root count", ErrUnitMismatch)
	}
	for i := range left.Roots {
		if err := compareNodes(left.Roots[i], right.Roots[i]); err != nil {
			return fmt.Errorf("%w: root %d: %w", ErrUnitMismatch, i, err)
		}
	}
	return nil
}

// compareAnnotations compares the sorted unit-level extension maps.
func compareAnnotations(left, right *Unit) error {
	keysLeft, keysRight := left.AnnotationKeys(), right.AnnotationKeys()
	if !slices.Equal(keysLeft, keysRight) {
		return fmt.Errorf("%w: annotation keys %v vs %v",
			ErrUnitMismatch, keysLeft, keysRight)
	}
	for _, key := range keysLeft {
		if left.Annotations[key] != right.Annotations[key] {
			return fmt.Errorf("%w: annotation %q", ErrUnitMismatch, key)
		}
	}
	return nil
}

// compareNodes compares two subtrees field by field.
func compareNodes(left, right Node) error {
	if left == nil || right == nil {
		if left == right {
			return nil
		}
		return fmt.Errorf("%w: nil node", ErrUnitMismatch)
	}
	if left.Kind() != right.Kind() ||
		left.Pos().Canonical() != right.Pos().Canonical() ||
		left.End().Canonical() != right.End().Canonical() ||
		AttributeOf(left) != AttributeOf(right) ||
		IDOf(left) != IDOf(right) ||
		SemanticsOf(left).Canonical() != SemanticsOf(right).Canonical() {
		return fmt.Errorf("%w: node %s", ErrUnitMismatch, left.Kind())
	}
	lc, rc := left.Children(), right.Children()
	if len(lc) != len(rc) {
		return fmt.Errorf("%w: child count of %s", ErrUnitMismatch, left.Kind())
	}
	for i := range lc {
		if err := compareNodes(lc[i], rc[i]); err != nil {
			return err
		}
	}
	return nil
}

// ComputeDiff derives a deterministic node-level difference from identifier
// multisets, so reused subtrees are counted even when the unit was rebuilt.
func ComputeDiff(old, fresh *Unit) Diff {
	before, after := idCounts(old), idCounts(fresh)
	keys := make([]NodeID, 0, len(before)+len(after))
	for id := range before {
		keys = append(keys, id)
	}
	for id := range after {
		if _, ok := before[id]; !ok {
			keys = append(keys, id)
		}
	}
	slices.Sort(keys)

	out := Diff{FullRebuild: old == nil || len(before) == 0}
	for _, id := range keys {
		left, right := before[id], after[id]
		shared := min(left, right)
		out.Reused += shared
		for i := shared; i < right; i++ {
			out.Added = append(out.Added, id)
		}
		for i := shared; i < left; i++ {
			out.Removed = append(out.Removed, id)
		}
	}
	return out
}

// idCounts tallies node identifiers over a whole unit.
func idCounts(u *Unit) map[NodeID]int {
	counts := make(map[NodeID]int)
	for _, n := range Collect(u) {
		counts[IDOf(n)]++
	}
	return counts
}
