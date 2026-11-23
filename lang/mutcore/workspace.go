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

// skippedDirs names the directories that never participate in a staged
// mutant workspace: VCS state, vendored dependencies and build outputs.
// Staging them would multiply the per-mutant cost without changing any
// toolchain verdict.
var skippedDirs = map[string]bool{
	".git":          true,
	"vendor":        true,
	"node_modules":  true,
	"target":        true,
	"bin":           true,
	"obj":           true,
	"__pycache__":   true,
	".pytest_cache": true,
}

// CopyWorkspace materializes the toolchain-visible subset of the workspace
// tree at src into dst: every file the driver requires, with VCS and build
// output directories skipped. Permissions are preserved and the copy is
// fully synchronous, so the caller never observes a partially materialized
// tree.
func CopyWorkspace(src, dst string, required func(string) bool) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	walk := func(path string, d fs.DirEntry, err error) error {
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
			if skippedDirs[d.Name()] {
				return filepath.SkipDir
			}
			return os.MkdirAll(filepath.Join(dst, rel), 0o755)
		}
		if !required(d.Name()) {
			return nil
		}
		return copyFile(path, filepath.Join(dst, rel), d)
	}
	if err := filepath.WalkDir(src, walk); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	return nil
}

// copyFile copies one regular file preserving its permission bits. The
// destination always sits inside a freshly staged directory, so the file
// never exists yet and WriteFile applies the source permissions exactly.
func copyFile(src, dst string, d fs.DirEntry) error {
	info, err := d.Info()
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return nil
	}
	data, err := os.ReadFile(src) //nolint:gosec // the path is enumerated from the workspace tree.
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, info.Mode().Perm()) //nolint:gosec // dst is derived from the caller-supplied workspace root.
}

// RewriteToken returns the file content with exactly one token replaced by
// the mutated text. The token interval must fall inside the file, which
// the scanner guarantees because every token was discovered on this
// snapshot.
func RewriteToken(content []byte, tok Token, mutated string) ([]byte, error) {
	if tok.Offset < 0 || tok.End > len(content) || tok.Offset >= tok.End {
		return nil, fmt.Errorf("%w: token interval [%d,%d) outside %s (%d bytes)",
			ErrWorkspaceFailed, tok.Offset, tok.End, tok.Path, len(content))
	}
	if got := string(content[tok.Offset:tok.End]); got != tok.Text {
		return nil, fmt.Errorf("%w: token drift in %s: expected %q, found %q",
			ErrWorkspaceFailed, tok.Path, tok.Text, got)
	}
	out := make([]byte, 0, len(content)-(tok.End-tok.Offset)+len(mutated))
	out = append(out, content[:tok.Offset]...)
	out = append(out, mutated...)
	out = append(out, content[tok.End:]...)
	return out, nil
}

// ApplyMutation rewrites exactly one token of the staged workspace to the
// mutated text. The rewrite is verified against the snapshot discovered by
// the scanner: if the staged file does not carry the expected original
// text at the recorded byte interval, the workspace drifted and the
// evaluation aborts instead of judging the wrong source.
func ApplyMutation(dir string, tok Token, mutated string) error {
	target := filepath.Join(dir, filepath.FromSlash(tok.Path))
	content, err := os.ReadFile(target) //nolint:gosec // target is derived from the scanned snapshot inside our own staging tree.
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	rewritten, err := RewriteToken(content, tok, mutated)
	if err != nil {
		return err
	}
	if err := os.WriteFile(target, rewritten, 0o644); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	return nil
}
