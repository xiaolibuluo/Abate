/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutcore

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyPolicy declares which files the sandbox must materialize for every
// staged mutant workspace, which heavy directories are copied once and
// shared across mutants through link farming, and which paths are always
// excluded. It is the contract every native-mode backend supplies to the
// session core so a real third-party project (with its dependency tree)
// can be judged without copying node_modules once per mutant.
type CopyPolicy struct {
	// Always lists the glob patterns for files that must be physically
	// materialized in every per-mutant workspace. These are the source
	// files and build manifests whose contents change when a token is
	// mutated—the staged tree must own its own copy so ApplyMutation
	// can rewrite them without corrupting the pristine sandbox root.
	Always []string

	// Once lists the relative directory paths that are copied into the
	// sandbox root exactly once at session creation. Per-mutant
	// workspaces receive symlinks (or hardlinks) to these directories
	// instead of a physical copy, so a 500 MB node_modules or vendor/
	// incurs a one-time cost at setup and near-zero cost per mutant.
	Once []string

	// Exclude lists the glob patterns for paths that are never
	// materialized, even when they match Always. It is additive on top
	// of the hard-coded skippedDirs list (e.g. .git, target/, bin/).
	Exclude []string
}

// Validate reports whether the policy is usable: at least one source
// pattern in Always, and no Once directory that overlaps with Exclude.
func (p CopyPolicy) Validate() error {
	if len(p.Always) == 0 {
		return fmt.Errorf("%w: CopyPolicy.Always must have at least one pattern", ErrInvalidRoot)
	}
	for _, once := range p.Once {
		for _, excl := range p.Exclude {
			if once == excl {
				return fmt.Errorf("%w: CopyPolicy.Once %q conflicts with Exclude %q",
					ErrInvalidRoot, once, excl)
			}
		}
	}
	return nil
}

// CopyWorkspaceNative materializes a mutant workspace for sandbox mode.
//
// Unlike CopyWorkspace which copies every required file per mutant, this
// function expects that a pristine sandbox root (sandboxRoot) was already
// populated with the Once directories before the session started judging.
// It copies only the Always files into the per-mutant staging directory
// and creates symlinks back to sandboxRoot for the Once directories. The
// mutation is then applied on top of this hybrid tree.
func CopyWorkspaceNative(sandboxRoot, dst string, policy CopyPolicy, required func(string) bool) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	// Copy Always files from sandboxRoot into dst.
	walk := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(sandboxRoot, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			if skippedDirs[d.Name()] || matchAny(rel, policy.Exclude) {
				return filepath.SkipDir
			}
			// Check if this directory path matches any Once entry.
			for _, once := range policy.Once {
				if rel == once {
					return filepath.SkipDir
				}
			}
			return os.MkdirAll(filepath.Join(dst, rel), 0o755)
		}
		if !required(d.Name()) {
			return nil
		}
		if matchAny(rel, policy.Exclude) {
			return nil
		}
		return copyFile(path, filepath.Join(dst, rel), d)
	}
	if err := filepath.WalkDir(sandboxRoot, walk); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	// Link Once directories from sandboxRoot into dst.
	for _, once := range policy.Once {
		src := filepath.Join(sandboxRoot, once)
		if info, err := os.Stat(src); err != nil || !info.IsDir() {
			continue
		}
		dstDir := filepath.Join(dst, once)
		// Ensure parent directory exists.
		if parent := filepath.Dir(dstDir); parent != dst {
			if err := os.MkdirAll(parent, 0o755); err != nil {
				return fmt.Errorf("%w: mkdir parent %s: %v", ErrWorkspaceFailed, parent, err)
			}
		}
		if err := os.Symlink(src, dstDir); err != nil {
			return fmt.Errorf("%w: symlink %s -> %s: %v", ErrWorkspaceFailed, once, src, err)
		}
	}
	return nil
}

// InitSandbox creates the sandbox root directory from the user's project
// directory by copying the Once directories once. The returned sandboxRoot
// path must be kept alive for the lifetime of the session.
func InitSandbox(projectRoot string, policy CopyPolicy) (string, error) {
	sandRoot, err := os.MkdirTemp("", "abate-sandbox-")
	if err != nil {
		return "", fmt.Errorf("%w: create sandbox root: %v", ErrWorkspaceFailed, err)
	}
	for _, once := range policy.Once {
		src := filepath.Join(projectRoot, once)
		info, err := os.Stat(src)
		if err != nil {
			if os.IsNotExist(err) {
				// Missing Once directory is not fatal: the user may
				// not have pre-built dependencies.
				continue
			}
			_ = os.RemoveAll(sandRoot)
			return "", fmt.Errorf("%w: stat %s: %v", ErrWorkspaceFailed, src, err)
		}
		if !info.IsDir() {
			_ = os.RemoveAll(sandRoot)
			return "", fmt.Errorf("%w: Once entry %s is not a directory", ErrWorkspaceFailed, src)
		}
		dst := filepath.Join(sandRoot, once)
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			_ = os.RemoveAll(sandRoot)
			return "", fmt.Errorf("%w: mkdir %s: %v", ErrWorkspaceFailed, filepath.Dir(dst), err)
		}
		if err := copyDir(src, dst); err != nil {
			_ = os.RemoveAll(sandRoot)
			return "", fmt.Errorf("%w: copy Once %s: %v", ErrWorkspaceFailed, once, err)
		}
	}
	return sandRoot, nil
}

// copyDir recursively copies a directory tree.
func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			return os.MkdirAll(filepath.Join(dst, rel), 0o755)
		}
		return copyFile(path, filepath.Join(dst, rel), d)
	})
}

// matchAny reports whether name matches any of the literal path prefixes
// in list. It performs a simple lexical prefix check, not a glob expand.
func matchAny(name string, list []string) bool {
	for _, pat := range list {
		if name == pat {
			return true
		}
	}
	return false
}
