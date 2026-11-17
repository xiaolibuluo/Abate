/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// errBroken reports a QA fixture that could not be derived from a real file.
var errBroken = errors.New("corpusgen: broken fixture failed")

// brokenMeta is project-owned metadata about one intentional syntax error.
type brokenMeta struct {
	Language  string `json:"language"`
	Kind      string `json:"kind"`
	Expected  string `json:"expectedDiagnostic"`
	SourceRel string `json:"sourceRelPath"`
}

// brokenExtensions maps each language to its generated source extension.
var brokenExtensions = map[string]string{
	"go": "gotxt", "py": "py", "js": "js", "c": "c",
	"cs": "cs", "rs": "rs", "kt": "kt",
}

// damageSource applies one precise syntax-level edit to a real clean file.
func damageSource(language string, content []byte) ([]byte, string, error) {
	data := append([]byte(nil), content...)
	if language == "py" {
		marker := strings.Index(string(data), "def ")
		if marker < 0 {
			return nil, "", fmt.Errorf("%w: %s has no def", errBroken, language)
		}
		colon := strings.IndexByte(string(data[marker:]), ':')
		if colon < 0 {
			return nil, "", fmt.Errorf("%w: %s has no def colon", errBroken, language)
		}
		at := marker + colon
		out := append(append([]byte(nil), data[:at]...), data[at+1:]...)
		return out, "missing-colon", nil
	}
	last := strings.LastIndexByte(string(data), '}')
	if last < 0 {
		return nil, "", fmt.Errorf("%w: %s has no closing brace", errBroken, language)
	}
	out := append(append([]byte(nil), data[:last]...), data[last+1:]...)
	return out, "unclosed-delimiter", nil
}

// GenerateBroken writes QA-only damaged fixtures. They are never registered
// in the main manifest, so they cannot pollute first-unit or capture metrics.
func GenerateBroken(brokenRoot, corpusRoot string, ref PlanRef) (int, error) {
	total := 0
	for _, language := range ref.Languages {
		sourceDir := filepath.Join(corpusRoot, language, "s0")
		entries, err := os.ReadDir(sourceDir)
		if err != nil {
			return total, fmt.Errorf("%w: read %s: %w", errBroken, sourceDir, err)
		}
		if len(entries) < ref.BrokenPerLanguage {
			return total, fmt.Errorf("%w: %s has %d clean fixtures, need %d",
				errBroken, language, len(entries), ref.BrokenPerLanguage)
		}
		targetDir := filepath.Join(brokenRoot, language)
		if err = os.MkdirAll(targetDir, generatedDirPerm); err != nil {
			return total, fmt.Errorf("%w: mkdir %s: %w", errBroken, targetDir, err)
		}
		ext, found := brokenExtensions[language]
		if !found {
			return total, fmt.Errorf("%w: no extension for %s", errBroken, language)
		}
		for i := 0; i < ref.BrokenPerLanguage; i++ {
			name := entries[i].Name()
			sourceRel := filepath.ToSlash(filepath.Join(language, "s0", name))
			// The path is built from a read-only directory listing of our own
			// generated corpus, not from untrusted input.
			clean, readErr := os.ReadFile(filepath.Join(sourceDir, name)) //nolint:gosec
			if readErr != nil {
				return total, fmt.Errorf("%w: read %s: %w", errBroken, sourceRel, readErr)
			}
			damaged, kind, damageErr := damageSource(language, clean)
			if damageErr != nil {
				return total, damageErr
			}
			base := filepath.Join(targetDir, fmt.Sprintf("%s_%05d", language, i))
			// QA fixtures carry a distinct .broken suffix: they deliberately
			// contain syntax errors and must never be parsed as Go source by
			// repository-wide tools such as the dogfood self-check. They are
			// not registered in the main manifest either.
			if err = os.WriteFile(base+"."+ext+".broken", damaged, generatedFilePerm); err != nil {
				return total, fmt.Errorf("%w: write %s: %w", errBroken, base, err)
			}
			meta := brokenMeta{
				Language: language, Kind: kind,
				Expected: "ParseError", SourceRel: sourceRel,
			}
			encoded, marshalErr := json.MarshalIndent(meta, "", "  ")
			if marshalErr != nil {
				return total, fmt.Errorf("%w: encode meta: %w", errBroken, marshalErr)
			}
			if err = os.WriteFile(base+".meta.json", encoded, generatedFilePerm); err != nil {
				return total, fmt.Errorf("%w: write meta: %w", errBroken, err)
			}
			total++
		}
	}
	return total, nil
}
