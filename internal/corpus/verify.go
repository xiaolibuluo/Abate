/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"errors"
	"fmt"
	"io/fs"
)

// Verify recomputes the SHA-256 of every registered entry from its original
// bytes and classifies every recorded SPDX expression. Every file is visited;
// there is no skip, whitelist or provenance-dependent branch.
func Verify(fsys fs.FS, m Manifest) (VerifyResult, error) {
	result := newVerifyResult(len(m.Entries))
	failures := make([]error, 0)
	for _, entry := range m.Entries {
		data, err := fs.ReadFile(fsys, entry.RelPath)
		if err != nil {
			failures = append(failures,
				fmt.Errorf("%w: %s", errEntryUnreadable, entry.RelPath))
		} else {
			checkEntry(&result, entry, hashBytes(data))
		}
	}
	return result, errors.Join(failures...)
}

// checkEntry folds one successfully read entry into the aggregate result.
func checkEntry(result *VerifyResult, entry Entry, actual string) {
	result.Recomputed++
	if actual != entry.SHA256 {
		result.SHA256Mismatches++
		result.Mismatches = append(result.Mismatches, fmt.Sprintf(
			"%s|expect=%s|actual=%s", entry.RelPath, entry.SHA256, actual))
	}
	if _, tracked := result.ParseablePerLanguage[entry.Language]; tracked && actual == entry.SHA256 {
		result.ParseablePerLanguage[entry.Language]++
	}
	verdict := ClassifySPDX(entry.SPDX)
	switch verdict {
	case spdxGPL, spdxUNKNOWN:
		result.SPDXViolations++
		result.SPDXDetails = append(result.SPDXDetails,
			entry.RelPath+"|spdx="+entry.SPDX)
	case spdxBoundary:
		result.BoundarySPDXCount++
	}
}
