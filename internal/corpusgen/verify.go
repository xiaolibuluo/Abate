/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// errVerify reports a failed generator self-check.
var errVerify = errors.New("corpusgen: self-check failed")

// VerifyStats is the raw evidence collected by the post-generation self-check.
type VerifyStats struct {
	Files       int
	PerLanguage map[string]int
	PerStratum  map[string]int
	DistinctSHA int
	Dedication  string
}

// thirdPartyMarkers must never appear in the generated source bytes.
var thirdPartyMarkers = []string{
	"Copyright (c)", "SPDX-License-Identifier: MIT", "Apache License",
	"The MIT License", "GNU GENERAL PUBLIC", "github.com/",
	"http://", "https://", "gin-gonic", "lodash", "cjson",
	"newtonsoft", "gvisor", "tree-sitter", "okhttp",
}

// verifyContext holds the immutable inputs of one verification run.
type verifyContext struct {
	ref    PlanRef
	bands  SizeBands
	quotas map[string]int
	stats  VerifyStats
	seen   map[string]bool
}

// parseRelPath extracts the language and stratum of a generated source file.
func (c *verifyContext) parseRelPath(relPath string) (string, int, error) {
	ext := filepath.Ext(relPath)
	language := ""
	for _, candidate := range c.ref.Languages {
		if ext == "."+brokenExtensions[candidate] {
			language = candidate
			break
		}
	}
	if language == "" {
		return "", 0, nil
	}
	parts := strings.Split(filepath.ToSlash(relPath), "/")
	if len(parts) != 3 || !strings.HasPrefix(parts[1], "s") {
		return "", 0, fmt.Errorf("%w: unexpected relPath %s", errVerify, relPath)
	}
	stratum := 0
	if _, err := fmt.Sscanf(parts[1], "s%d", &stratum); err != nil {
		return "", 0, fmt.Errorf("%w: stratum parse %s: %w", errVerify, relPath, err)
	}
	if stratum < 0 || stratum >= c.ref.StrataCount {
		return "", 0, fmt.Errorf("%w: stratum %d out of range in %s",
			errVerify, stratum, relPath)
	}
	return language, stratum, nil
}

// checkCaptureAnchor proves that the first mutation tokens belong to the file's
// own anchor and that non-clean anchors carry the first call token.
func checkCaptureAnchor(data []byte, relPath, ext string) error {
	marker := bytes.Index(data, []byte("1 + 2 == 3"))
	if marker < 0 {
		return fmt.Errorf("%w: %s has no capture anchor", errVerify, relPath)
	}
	classes := []struct {
		name   string
		tokens []string
	}{
		{"digit", anchorDigitTokens}, {"op", anchorOpTokens}, {"cmp", anchorCmpTokens},
	}
	for _, class := range classes {
		if at := firstTokenOffset(data, class.tokens); at < marker {
			return fmt.Errorf("%w: %s first %s at %d before anchor %d",
				errVerify, relPath, class.name, at, marker)
		}
	}
	clean := strings.HasSuffix(filepath.Base(relPath), "_"+firstFileLabel+ext)
	if !clean && firstTokenOffset(data, anchorCallTokens) < marker {
		return fmt.Errorf("%w: %s first call token before anchor", errVerify, relPath)
	}
	lower := strings.ToLower(string(data))
	for _, markerText := range thirdPartyMarkers {
		if strings.Contains(lower, strings.ToLower(markerText)) {
			return fmt.Errorf("%w: %s contains third-party marker %q",
				errVerify, relPath, markerText)
		}
	}
	return nil
}

// inspectOneFile validates one source file and records its raw numbers.
func (c *verifyContext) inspectOneFile(relPath string, data []byte) error {
	language, stratum, err := c.parseRelPath(relPath)
	if err != nil || language == "" {
		return err
	}
	if got := c.bands.StratumOf(len(data)); got != stratum {
		return fmt.Errorf("%w: %s byte band %d != %d", errVerify, relPath, got, stratum)
	}
	sum := sha256.Sum256(data)
	digest := hex.EncodeToString(sum[:])
	if c.seen[digest] {
		return fmt.Errorf("%w: duplicate sha %s", errVerify, relPath)
	}
	c.seen[digest] = true
	if err = checkCaptureAnchor(data, relPath, filepath.Ext(relPath)); err != nil {
		return err
	}
	c.stats.Files++
	c.stats.PerLanguage[language]++
	c.stats.PerStratum[fmt.Sprintf("%s:s%d", language, stratum)]++
	return nil
}

// checkQuotas proves that every language and every stratum is completely full.
func (c *verifyContext) checkQuotas() error {
	for _, language := range c.ref.Languages {
		want := c.quotas[language]
		if c.stats.PerLanguage[language] != want {
			return fmt.Errorf("%w: %s files=%d want %d",
				errVerify, language, c.stats.PerLanguage[language], want)
		}
		for stratum := 0; stratum < c.ref.StrataCount; stratum++ {
			key := fmt.Sprintf("%s:s%d", language, stratum)
			bandWant := c.bands.Quota(want, stratum)
			if c.stats.PerStratum[key] != bandWant {
				return fmt.Errorf("%w: %s files=%d want %d",
					errVerify, key, c.stats.PerStratum[key], bandWant)
			}
		}
	}
	return nil
}

// Verify re-reads every generated file and enforces all generator invariants.
func Verify(opts GenOptions) (VerifyStats, error) {
	ctx := verifyContext{stats: VerifyStats{
		PerLanguage: map[string]int{}, PerStratum: map[string]int{},
	}, seen: map[string]bool{}}
	ref, err := LoadPlanRef(opts.PlanPath)
	if err != nil {
		return ctx.stats, err
	}
	bands, err := NewSizeBands(ref.MinFileBytes, ref.MaxFileBytes, ref.StrataCount)
	if err != nil {
		return ctx.stats, err
	}
	ctx.ref = ref
	ctx.bands = bands
	ctx.quotas = languageQuotas(opts.Count, ref.Languages)
	fsys := os.DirFS(opts.Root)
	err = fs.WalkDir(fsys, ".", func(relPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() {
			return walkErr
		}
		data, readErr := fs.ReadFile(fsys, relPath)
		if readErr != nil {
			return readErr
		}
		return ctx.inspectOneFile(relPath, data)
	})
	if err != nil {
		return ctx.stats, err
	}
	if err = ctx.checkQuotas(); err != nil {
		return ctx.stats, err
	}
	dedication, err := os.ReadFile(filepath.Join(opts.Root, DedicationRelPath))
	if err != nil {
		return ctx.stats, fmt.Errorf("%w: read dedication: %w", errVerify, err)
	}
	sum := sha256.Sum256(dedication)
	ctx.stats.Dedication = hex.EncodeToString(sum[:])
	ctx.stats.DistinctSHA = len(ctx.seen)
	return ctx.stats, nil
}
