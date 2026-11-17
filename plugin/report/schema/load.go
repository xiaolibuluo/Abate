/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"encoding/json"
	"fmt"
	"os"
)

// readFile reads the raw bytes of a schema document from disk. It is isolated so
// the validator has a single, testable I/O boundary.
func readFile(path string) ([]byte, error) {
	//nolint:gosec // the schema path is supplied explicitly by the caller.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("schema: cannot read %q: %w", path, err)
	}
	return data, nil
}

// unmarshal decodes JSON bytes into an any tree. The standard decoder
// represents numbers as float64, which is the assumption the validator relies on
// for integer and number checks.
func unmarshal(data []byte, out *any) error {
	return json.Unmarshal(data, out)
}
