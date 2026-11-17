/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// expandLanguagesPaths turns files and directories into a sorted file list
// restricted to a language's extensions, skipping vendored trees.
func expandLanguagesPaths(paths, extensions []string) ([]string, error) {
	out := make([]string, 0, 16)
	for _, raw := range paths {
		if strings.TrimSpace(raw) == "" {
			continue
		}
		clean := filepath.Clean(raw)
		info, err := os.Stat(clean)
		if err != nil {
			return nil, fmt.Errorf("%w: %s: %w", errReadSource, clean, err)
		}
		if !info.IsDir() {
			if hasExtension(clean, extensions) {
				out = append(out, clean)
			}
			continue
		}
		if err := walkLanguageDir(clean, extensions, &out); err != nil {
			return nil, err
		}
	}
	slices.Sort(out)
	return slices.Compact(out), nil
}

// walkLanguageDir collects matching files below a directory in lexical order.
func walkLanguageDir(dir string, extensions []string, out *[]string) error {
	err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if isSkipDirectory(entry.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if hasExtension(path, extensions) {
			*out = append(*out, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: %w", errReadSource, err)
	}
	return nil
}

// hasExtension reports whether a path ends with one of the extensions.
func hasExtension(path string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

// isSkipDirectory reports directories excluded from a check walk.
func isSkipDirectory(name string) bool {
	switch name {
	case "vendor", "third_party", ".git", "node_modules":
		return true
	default:
		return false
	}
}
