/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gotarget

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// Write errors: the target directory policy and the self-check gate.
var (
	ErrDirExists   = errors.New("gotarget: target directory already exists")
	ErrSelfCheck   = errors.New("gotarget: generated module failed self-check")
	ErrPackageName = errors.New("gotarget: invalid Go package name")
)

// Fingerprint returns the deterministic SHA-256 digest of the generated
// file set: file names are hashed in ascending order together with their
// bytes, so two generations agree on the fingerprint if and only if every
// file agrees byte for byte.
func (t Target) Fingerprint() [32]byte {
	names := make([]string, 0, len(t.Files))
	for name := range t.Files {
		names = append(names, name)
	}
	slices.Sort(names)
	h := sha256.New()
	for _, name := range names {
		h.Write([]byte(name))
		h.Write([]byte{0})
		h.Write(t.Files[name])
		h.Write([]byte{0})
	}
	var sum [32]byte
	copy(sum[:], h.Sum(nil))
	return sum
}

// ByteSize returns the total byte count of every generated file.
func (t Target) ByteSize() int {
	total := 0
	for _, data := range t.Files {
		total += len(data)
	}
	return total
}

// Write materializes the target module into dir atomically: the files land
// in a sibling temporary directory first, the optional self-check runs the
// real toolchain against that staging tree, and only a fully green module
// is renamed into place. A failed write never leaves a partial target
// directory behind. dir must not exist yet, so a generation can never
// silently overwrite user sources.
func Write(dir string, target Target, selfCheck bool) error {
	if _, err := os.Stat(dir); err == nil {
		return fmt.Errorf("%w: %s", ErrDirExists, dir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("gotarget: stat %s: %v", dir, err)
	}
	parent := filepath.Dir(dir)
	base := filepath.Base(dir)
	staging, err := os.MkdirTemp(parent, "."+base+".targetgen-")
	if err != nil {
		return fmt.Errorf("gotarget: staging directory: %v", err)
	}
	defer func() { _ = os.RemoveAll(staging) }()
	names := make([]string, 0, len(target.Files))
	for name := range target.Files {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		path := filepath.Join(staging, filepath.FromSlash(name))
		if err := os.WriteFile(path, target.Files[name], 0o644); err != nil {
			return fmt.Errorf("gotarget: write %s: %v", name, err)
		}
	}
	if selfCheck {
		if output, err := runSelfCheck(staging); err != nil {
			return fmt.Errorf("%w: %v: %s", ErrSelfCheck, err, tailOutput(output))
		}
	}
	if err := os.Rename(staging, dir); err != nil {
		return fmt.Errorf("gotarget: move into place %s: %v", dir, err)
	}
	return nil
}

// runSelfCheck executes `go test ./...` inside the staged module. The
// generated module has zero dependencies beyond the standard library, so
// the check runs fully offline; a red module is a generator bug and is
// reported through ErrSelfCheck instead of being delivered.
func runSelfCheck(dir string) (string, error) {
	cmd := exec.Command("go", "test", "-count=1", "./...")
	cmd.Dir = dir
	var buf strings.Builder
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	return buf.String(), err
}

// tailOutput truncates self-check output so an error message stays readable
// while still carrying the failing test diagnostics.
func tailOutput(output string) string {
	text := strings.TrimSpace(output)
	if len(text) > 512 {
		return text[:512]
	}
	return text
}

// ValidPackageName reports whether name is a legal ASCII Go package
// identifier. The generator rejects anything else up front so the
// self-check never has to surface a trivial syntax error.
func ValidPackageName(name string) bool {
	if name == "" {
		return false
	}
	for i, r := range name {
		switch {
		case r >= 'a' && r <= 'z':
		case r >= 'A' && r <= 'Z':
		case r == '_':
		case r >= '0' && r <= '9' && i > 0:
		default:
			return false
		}
	}
	return true
}
