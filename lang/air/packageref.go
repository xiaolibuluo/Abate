/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"slices"
	"strings"
)

// PackageRef names the package of a unit and holds its canonical import set.
type PackageRef struct {
	// Name is the declared package name.
	Name string
	// Imports is the deduplicated, lexicographically sorted import path list.
	Imports []string
}

// NewPackageRef normalizes imports by copying, sorting, and deduplicating them.
func NewPackageRef(name string, imports []string) PackageRef {
	sorted := slices.Clone(imports)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	return PackageRef{Name: name, Imports: sorted}
}

// Validate checks that every import path is non-empty.
func (p PackageRef) Validate() error {
	for _, path := range p.Imports {
		if path == "" {
			return ErrMalformedUnit
		}
	}
	return nil
}

// Canonical returns a stable comparison key for the package reference.
func (p PackageRef) Canonical() string {
	return p.Name + "\x00" + strings.Join(p.Imports, "\x00")
}

// Clone returns an independent copy of the package reference.
func (p PackageRef) Clone() PackageRef {
	return PackageRef{Name: p.Name, Imports: slices.Clone(p.Imports)}
}
