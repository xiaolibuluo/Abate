/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"fmt"
	"io/fs"
	"slices"
	"sort"
)

// ProjectRef describes one read-only upstream tree and the provenance shared
// by every source file registered beneath it.
type ProjectRef struct {
	Language       string
	RelDir         string
	UpstreamURL    string
	Commit         string
	SPDX           string
	LicenseRelPath string
	Project        string
	Source         Source
	ConstructTags  []string
}

// Register walks each project tree in lexical order, records every file whose
// extension belongs to one of the seven languages and returns a manifest whose
// entries are sorted by RelPath ascending.
func Register(fsys fs.FS, projects []ProjectRef) (Manifest, error) {
	entries := make([]Entry, 0)
	for _, ref := range projects {
		evidence, err := evidenceFor(fsys, ref)
		if err != nil {
			return Manifest{}, err
		}
		batch, err := walkProject(fsys, ref, evidence)
		if err != nil {
			return Manifest{}, err
		}
		entries = append(entries, batch...)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].RelPath < entries[j].RelPath
	})
	return withDigest(Manifest{
		SchemaVersion: manifestSchemaVersion,
		Entries:       entries,
	}), nil
}

// evidenceFor binds the licence declaration to the SHA-256 of the real file
// that carries it, so every entry can be traced back without fabrication.
func evidenceFor(fsys fs.FS, ref ProjectRef) (string, error) {
	digest, _, err := fileDigest(fsys, ref.LicenseRelPath)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errEvidenceUnreadable, ref.LicenseRelPath)
	}
	return ref.LicenseRelPath + ":" + digest, nil
}

// walkProject registers one tree. Non-source extensions are outside the
// registration scope; no file is dropped based on parseability or provenance.
func walkProject(fsys fs.FS, ref ProjectRef, evidence string) ([]Entry, error) {
	out := make([]Entry, 0)
	err := fs.WalkDir(fsys, ref.RelDir, func(relPath string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		lang := languageOfExt(relPath)
		if lang == "" {
			return nil
		}
		digest, size, readErr := fileDigest(fsys, relPath)
		if readErr != nil {
			return fmt.Errorf("%w: %s", errEntryUnreadable, relPath)
		}
		out = append(out, Entry{
			Language:      lang,
			Project:       ref.Project,
			UpstreamURL:   ref.UpstreamURL,
			Commit:        ref.Commit,
			SPDX:          ref.SPDX,
			SPDXEvidence:  evidence,
			RelPath:       relPath,
			Bytes:         size,
			SHA256:        digest,
			ConstructTags: slices.Clone(ref.ConstructTags),
			Source:        ref.Source,
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errEntryUnreadable, ref.RelDir)
	}
	return out, nil
}
