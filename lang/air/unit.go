/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"crypto/sha256"
	"slices"
	"sort"
)

// Unit is one versioned compilation unit.
//
// Files and Roots are parallel slices kept in ascending path order, so the
// encoder never depends on the order in which a frontend registered them.
// Annotations are unit-level extension fields introduced in schema version 2;
// version 1 archives simply have none.
type Unit struct {
	// Package is the package metadata of the unit.
	Package PackageRef
	// Files is the path-ordered file metadata list.
	Files []FileRef
	// Roots holds the AIR root node of each file, aligned with Files.
	Roots []Node
	// Annotations holds additive unit-level metadata, serialized in sorted order.
	Annotations map[string]string
}

// NewUnit returns an empty unit for the given package.
func NewUnit(pkg PackageRef) *Unit {
	return &Unit{
		Package:     pkg.Clone(),
		Files:       []FileRef{},
		Roots:       []Node{},
		Annotations: map[string]string{},
	}
}

// AddFile registers a file and its root node, keeping both slices path-sorted.
func (u *Unit) AddFile(file FileRef, root Node) error {
	if u == nil {
		return ErrNilUnit
	}
	if err := file.Validate(); err != nil {
		return err
	}
	if root == nil {
		return ErrMalformedNode
	}
	at := sort.Search(len(u.Files), func(i int) bool {
		return u.Files[i].Path >= file.Path
	})
	if at < len(u.Files) && u.Files[at].Path == file.Path {
		return ErrDuplicatePath
	}
	u.Files = slices.Insert(u.Files, at, file)
	u.Roots = slices.Insert(u.Roots, at, root)
	return nil
}

// Validate checks the structural invariants of the unit.
func (u *Unit) Validate() error {
	if u == nil {
		return ErrNilUnit
	}
	if err := u.Package.Validate(); err != nil {
		return err
	}
	if len(u.Files) != len(u.Roots) {
		return ErrMalformedUnit
	}
	for i, file := range u.Files {
		if err := file.Validate(); err != nil {
			return err
		}
		if u.Roots[i] == nil || !u.Roots[i].Pos().Valid() {
			return ErrMalformedNode
		}
	}
	return nil
}

// Fingerprint returns the SHA-256 digest of the canonical section bytes.
//
// The fingerprint is computed from the same section encoder used by the wire
// format, so a stable unit always produces a stable hash without any I/O.
func (u *Unit) Fingerprint() [32]byte {
	if u == nil {
		return [32]byte{}
	}
	return sha256.Sum256(encodeSections(u, NewVersion()))
}

// NodeCount returns the total number of nodes reachable from every root.
func (u *Unit) NodeCount() int {
	if u == nil {
		return 0
	}
	count := 0
	for _, root := range u.Roots {
		Walk(root, func(_ Node) bool {
			count++
			return true
		})
	}
	return count
}

// AnnotationKeys returns the extension keys in ascending order.
func (u *Unit) AnnotationKeys() []string {
	keys := make([]string, 0, len(u.Annotations))
	for key := range u.Annotations {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
