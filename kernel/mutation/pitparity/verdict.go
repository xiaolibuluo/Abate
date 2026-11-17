/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"fmt"

	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// verdictIndex indexes the engine verdicts by mutation point and by node.
type verdictIndex struct {
	points map[string]verify.MutationVerdict
	nodes  map[string]verify.MutationVerdict
}

// newVerdictIndex indexes the verdicts of one engine run.
func newVerdictIndex(run mutation.Report) (verdictIndex, error) {
	index := verdictIndex{
		points: make(map[string]verify.MutationVerdict, len(run.Points)),
		nodes:  make(map[string]verify.MutationVerdict, len(run.Points)),
	}
	for i, point := range run.Points {
		if i >= len(run.Verdicts) {
			return verdictIndex{}, fmt.Errorf("%w: point %d has no verdict",
				ErrMissingVerdict, i)
		}
		verdict := run.Verdicts[i]
		class := ""
		if point.Operator != nil {
			class = point.Operator.Params()["class"]
		}
		key := point.Node + "|" + class
		if prior, ok := index.points[key]; ok && prior != verdict {
			return verdictIndex{}, fmt.Errorf("%w: node %q class %q has two verdicts",
				ErrMissingVerdict, point.Node, class)
		}
		index.points[key] = verdict
		if prior, ok := index.nodes[point.Node]; ok && prior != verdict {
			return verdictIndex{}, fmt.Errorf("%w: node %q has two verdicts",
				ErrMissingVerdict, point.Node)
		}
		index.nodes[point.Node] = verdict
	}
	return index, nil
}

// lookup returns the verdict of one mapping, preferring the exact point.
func (v verdictIndex) lookup(mapping Mapping) (verify.MutationVerdict, string, bool) {
	if verdict, ok := v.points[mapping.Node+"|"+mapping.Class]; ok {
		return verdict, VerdictModePoint, true
	}
	if verdict, ok := v.nodes[mapping.Node]; ok {
		return verdict, VerdictModeNode, true
	}
	return "", "", false
}
