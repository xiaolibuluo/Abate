/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"fmt"
)

// ErrUnmappedMutation reports a PIT mutation the mapping table does not cover.
var ErrUnmappedMutation = errors.New("pitparity: mutation is not registered in the mapping table")

// ErrMissingVerdict reports that the engine produced no verdict for a mapping.
var ErrMissingVerdict = errors.New("pitparity: no engine verdict for the mapped node")

// Mapping binds one PIT mutation site to the Abate fixture node and operator
// class the parity comparison uses.
//
// 参考 PITest 变异体与 Abate 变异点的对应关系，改进点=映射显式登记且带语义说明，
// 未登记的变异一律以 ErrUnmappedMutation 拒绝，杜绝静默丢弃。
type Mapping struct {
	// Method is the mutated method of the Java baseline project.
	Method string `json:"method"`

	// Mutator is the PIT mutator family of the mutation.
	Mutator string `json:"mutator"`

	// Line is the mutated source line of the Java baseline project.
	Line int `json:"line"`

	// Node is the Abate fixture node the mutation is compared against.
	Node string `json:"node"`

	// Class is the Abate operator class the mutation is compared against.
	Class string `json:"class"`

	// Excluded marks a PIT specific operator without an Abate counterpart.
	Excluded bool `json:"excluded"`

	// Reason states the semantic correspondence or the exclusion cause.
	Reason string `json:"reason"`
}

// lookupMapping returns the mapping registered for one PIT mutation site.
//
// The table is scanned linearly so the result never depends on map iteration
// order and the registration order stays the order of the reference report.
func lookupMapping(method, mutator string, line int) (Mapping, bool) {
	for _, entry := range mappingTable {
		if entry.Method == method && entry.Mutator == mutator && entry.Line == line {
			return entry, true
		}
	}
	return Mapping{}, false
}

// requireMapping returns the mapping of one site and fails when it is missing.
func requireMapping(method, mutator string, line int) (Mapping, error) {
	mapping, ok := lookupMapping(method, mutator, line)
	if !ok {
		return Mapping{}, fmt.Errorf("%w: %s:%d %s",
			ErrUnmappedMutation, method, line, mutator)
	}
	return mapping, nil
}

// resolveSite returns the mapping and the semantic witness of one site.
//
// Both tables must cover a site: a registered mapping without a witness would
// silently drop the mutation, so the resolution fails instead.
func resolveSite(method, mutator string, line int) (Mapping, witnessCase, error) {
	mapping, err := requireMapping(method, mutator, line)
	if err != nil {
		return Mapping{}, witnessCase{}, err
	}
	witness, ok := witnessFor(method, mutator, line)
	if !ok {
		return Mapping{}, witnessCase{}, fmt.Errorf("%w: %s:%d %s has no semantic witness",
			ErrUnmappedMutation, method, line, mutator)
	}
	return mapping, witness, nil
}
