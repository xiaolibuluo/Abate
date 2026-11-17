/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// statusNotMeasured is the honest label of the self-built Go inventory.
const statusNotMeasured = "NOT_MEASURED"

var (
	// ErrNotMeasured reports that the Go construct inventory is a self-built
	// list with no seventh embedded node-types document and no measured
	// coverage. Callers must report NOT_MEASURED instead of a fabricated ratio.
	ErrNotMeasured    = errors.New("coverage: go constructs are NOT_MEASURED")
	errBadGoInventory = errors.New("coverage: invalid go construct inventory")
)

// goConstructJSON is one row of the on-disk Go construct table.
type goConstructJSON struct {
	NodeType string `json:"nodeType"`
	Category string `json:"category"`
}

// goInventoryJSON is the on-disk projection of ci/constructs_go.json.
type goInventoryJSON struct {
	SchemaVersion int               `json:"schemaVersion"`
	Language      string            `json:"language"`
	Provenance    string            `json:"source"`
	Status        string            `json:"status"`
	Measured      bool              `json:"measured"`
	Constructs    []goConstructJSON `json:"constructs"`
}

// GoInventory is the explicit self-built Go construct inventory. Only six
// upstream node-types documents are embedded, so this inventory enumerates
// go/ast node types instead of a measured seventh data set.
type GoInventory struct {
	SchemaVersion int
	Language      string
	Provenance    string
	Status        string
	Measured      bool
	Constructs    []string
	categories    map[string]string
}

// LoadGoInventory reads and validates the self-built inventory. A document that
// claims to be measured is rejected, because no real measurement backs it.
func LoadGoInventory(path string) (GoInventory, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return GoInventory{}, fmt.Errorf("%w: %s", errBadGoInventory, path)
	}
	var raw goInventoryJSON
	if err = json.Unmarshal(data, &raw); err != nil {
		return GoInventory{}, fmt.Errorf("%w: %s", errBadGoInventory, path)
	}
	if raw.Language != "go" || raw.Status != statusNotMeasured ||
		raw.Measured || raw.SchemaVersion < 1 || len(raw.Constructs) == 0 {
		return GoInventory{}, fmt.Errorf("%w: %s", errBadGoInventory, path)
	}
	out := GoInventory{
		SchemaVersion: raw.SchemaVersion,
		Language:      raw.Language,
		Provenance:    raw.Provenance,
		Status:        raw.Status,
		Measured:      raw.Measured,
		Constructs:    make([]string, 0, len(raw.Constructs)),
		categories:    make(map[string]string, len(raw.Constructs)),
	}
	seen := make(map[string]struct{}, len(raw.Constructs))
	for _, item := range raw.Constructs {
		if item.NodeType == "" {
			return GoInventory{}, fmt.Errorf("%w: empty node type",
				errBadGoInventory)
		}
		if _, duplicate := seen[item.NodeType]; duplicate {
			return GoInventory{}, fmt.Errorf("%w: duplicate %s",
				errBadGoInventory, item.NodeType)
		}
		seen[item.NodeType] = struct{}{}
		out.Constructs = append(out.Constructs, item.NodeType)
		out.categories[item.NodeType] = item.Category
	}
	slices.Sort(out.Constructs)
	return out, nil
}

// Nodes returns the construct names in ascending order.
func (i GoInventory) Nodes() []string {
	out := slices.Clone(i.Constructs)
	slices.Sort(out)
	return out
}

// CategoryOf returns the recorded category of one construct.
func (i GoInventory) CategoryOf(node string) string {
	return i.categories[node]
}

// IsMeasured reports the measurement status without relaxing it. A document
// labelled NOT_MEASURED is never treated as measured.
func (i GoInventory) IsMeasured() bool {
	return i.Measured && i.Status != statusNotMeasured
}
