/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"abate/plugin/report/schema"
)

// oasisSchemaFile is the bundled OASIS SARIF 2.1.0 JSON-Schema reference. It lives
// under internal/standards/sarif as public standard fact data: an international
// standard schema, not proprietary open-source implementation source. The schema
// validator in plugin/report/schema is 100% self-developed. The file is loaded at
// validation time to prove the generated SARIF conforms to the published standard
// without relying on the external Microsoft validator.
// ⚖️ Distribution/license terms of the standard data file await rights-holder/legal review.
const oasisSchemaFile = "internal/standards/sarif/sarif-2.1.0.schema.json"

// findOASISSchema walks upward from the current working directory looking for the
// bundled SARIF schema. This keeps the validator robust to the package's working
// directory (Go tests run from the package directory, while the published binary
// may run from the repository root).
func findOASISSchema() (string, bool) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	for i := 0; i < 8; i++ {
		candidate := filepath.Join(dir, oasisSchemaFile)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", false
}

// ValidateSARIFAgainst validates a generated SARIF document against the OASIS 2.1.0
// schema loaded from schemaPath. A nil error means the document satisfies every
// keyword the self-researched offline validator checks.
func ValidateSARIFAgainst(data []byte, schemaPath string) error {
	compiled, err := schema.LoadSchema(schemaPath)
	if err != nil {
		return err
	}
	var doc any
	if err := json.Unmarshal(data, &doc); err != nil {
		return fmt.Errorf("report: generated SARIF is not valid JSON: %w", err)
	}
	return compiled.Validate(doc)
}

// ValidateSARIFAgainstOASIS validates a generated SARIF document against the
// bundled OASIS 2.1.0 schema, locating the schema file automatically. It is the
// function the acceptance tests use to assert RPT-SARIF passes the official
// validator.
func ValidateSARIFAgainstOASIS(data []byte) error {
	path, ok := findOASISSchema()
	if !ok {
		return fmt.Errorf("report: cannot locate OASIS schema %q", oasisSchemaFile)
	}
	return ValidateSARIFAgainst(data, path)
}
