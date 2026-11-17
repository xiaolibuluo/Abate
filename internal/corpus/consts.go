/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import "errors"

// Kib is the number of bytes in one kibibyte.
const Kib = 1024

// Mib is the number of bytes in one mebibyte.
const Mib = 1024 * Kib

// DefaultCorpusRoot is the single source of truth for the repository-relative
// path of the self-researched E2E corpus. Both the E2E runner and corpusctl
// resolve their corpus from this constant, so the two paths cannot drift.
const DefaultCorpusRoot = "testdata/corpus"

// Named parameters of the corpus registry and the deterministic sampler.
const (
	manifestSchemaVersion   = 1
	minFileBytes            = 256
	maxFileBytes            = 180 * Kib
	strataCount             = 8
	defaultFilesPerLanguage = 200
	defaultSeed             = 20261114
)

// Canonical report order of the seven registered languages.
var languageOrder = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// Extension-to-language registration table. Both .c and .h map to c by design.
var extensionLanguage = map[string]string{
	".go":    "go",
	".gotxt": "go",
	".py":    "py",
	".js":    "js",
	".c":     "c",
	".h":     "c",
	".cs":    "cs",
	".rs":    "rs",
	".kt":    "kt",
}

// Upper-case SPDX identifiers accepted as permissive upstream declarations.
var allowlistedSPDX = map[string]bool{
	"MIT":           true,
	"APACHE-2.0":    true,
	"BSD-3-CLAUSE":  true,
	"BSD-2-CLAUSE":  true,
	"ISC":           true,
	"CC0-1.0":       true,
	"UNLICENSE":     true,
	"PUBLIC-DOMAIN": true,
	"PUBLIC DOMAIN": true,
}

// Markers that require a licence to be recorded as a boundary case.
var boundarySPDXMarkers = []string{"EXCEPTION", "CLASSPATH", "WITH"}

// Markers that identify copyleft licences rejected by verification.
var gplSPDXMarkers = []string{
	"GPL", "LGPL", "AGPL",
	"GNU GENERAL PUBLIC", "GNU LESSER GENERAL PUBLIC",
	"GNU AFFERO", "GNU LIBRARY GENERAL PUBLIC",
}

// Connector words removed before an SPDX expression is tokenised.
var spdxConnectors = map[string]bool{"OR": true, "AND": true}

// Package-level sentinels keep wrapped failures comparable through errors.Is.
var (
	errEntryUnreadable    = errors.New("corpus: entry unreadable")
	errEvidenceUnreadable = errors.New("corpus: licence evidence unreadable")
	errManifestUnreadable = errors.New("corpus: manifest unreadable")
	errManifestMalformed  = errors.New("corpus: manifest malformed")
	errUnsortedManifest   = errors.New("corpus: entries are not sorted by relPath")
)

// ErrVerifyFailed reports a non-zero zero-tolerance counter.
var ErrVerifyFailed = errors.New("corpus: verification failed")

// ErrUnknownCommand reports an unsupported corpusctl subcommand.
var ErrUnknownCommand = errors.New("corpusctl: expected register or verify")
