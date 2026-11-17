/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"abate/internal/e2e/metrics"
)

var (
	errBadEquivalentFile = errors.New("mutate: invalid equivalent definitions file")
	errDuplicateProof    = errors.New("mutate: duplicate equivalence proof")
)

// equivalentFile is the on-disk projection of ci/equivalent_mutations.json.
type equivalentFile struct {
	SchemaVersion int             `json:"schemaVersion"`
	Definitions   []EquivalentDef `json:"definitions"`
}

// EquivalentSet is the explicit, closed set of equivalent mutation definitions.
// A mutation is equivalent only when it matches one of these declarations.
type EquivalentSet struct {
	defs []EquivalentDef
}

// NewEquivalentSet validates and orders a set of definitions by proof id.
func NewEquivalentSet(defs []EquivalentDef) (EquivalentSet, error) {
	out := EquivalentSet{defs: slices.Clone(defs)}
	slices.SortFunc(out.defs, func(a, b EquivalentDef) int {
		if a.ProofID < b.ProofID {
			return -1
		}
		if a.ProofID > b.ProofID {
			return 1
		}
		return 0
	})
	for i, def := range out.defs {
		if err := def.Validate(); err != nil {
			return EquivalentSet{}, err
		}
		if i > 0 && out.defs[i-1].ProofID == def.ProofID {
			return EquivalentSet{}, fmt.Errorf("%w: %s", errDuplicateProof,
				def.ProofID)
		}
	}
	return out, nil
}

// All returns the definitions in ascending proof-id order.
func (s EquivalentSet) All() []EquivalentDef {
	return slices.Clone(s.defs)
}

// Classify returns the proof id when a drilldown record matches a declaration.
// A record matches by its recorded proof id, or by its operator id combined
// with the equivalent flag. Equivalent mutations are tagged, never removed.
func (s EquivalentSet) Classify(rec metrics.DrilldownRecord) (string, bool) {
	for _, def := range s.defs {
		if rec.EquivalentProofID == def.ProofID {
			return def.ProofID, true
		}
		if rec.OperatorID == def.OperatorID && rec.Equivalent {
			return def.ProofID, true
		}
	}
	return "", false
}

// LoadEquivalentSet reads the explicit set from a JSON document. The file is
// the single on-disk source of the equivalent-mutation declarations.
func LoadEquivalentSet(path string) (EquivalentSet, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return EquivalentSet{}, fmt.Errorf("%w: %s", errBadEquivalentFile, path)
	}
	var file equivalentFile
	if err = json.Unmarshal(data, &file); err != nil {
		return EquivalentSet{}, fmt.Errorf("%w: %s", errBadEquivalentFile, path)
	}
	if file.SchemaVersion < 1 || len(file.Definitions) == 0 {
		return EquivalentSet{}, fmt.Errorf("%w: empty definitions in %s",
			errBadEquivalentFile, path)
	}
	return NewEquivalentSet(file.Definitions)
}
