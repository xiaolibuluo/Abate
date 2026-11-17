/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"errors"
	"fmt"
)

// errBadEquivalentDef reports an equivalent definition missing one of its four
// required elements: operator, context, proof and rationale.
var errBadEquivalentDef = errors.New("mutate: incomplete equivalent definition")

// EquivalentDef declares one semantically equivalent mutation shape. The four
// fields are the operator, the context pattern, the proof identifier and the
// written equivalence argument; all four are required, so an equivalent
// mutation is always an explicit, auditable declaration.
type EquivalentDef struct {
	OperatorID     string `json:"operatorID"`
	ContextPattern string `json:"contextPattern"`
	ProofID        string `json:"proofID"`
	Rationale      string `json:"rationale"`
}

// Validate checks that every element of the definition is recorded.
func (d EquivalentDef) Validate() error {
	if d.OperatorID == "" || d.ContextPattern == "" ||
		d.ProofID == "" || d.Rationale == "" {
		return fmt.Errorf("%w: %q", errBadEquivalentDef, d.ProofID)
	}
	return nil
}
