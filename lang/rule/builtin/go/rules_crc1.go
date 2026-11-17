/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goCRC1Rules declares the naming and comment rules drawn from the Go Code
// Review Comments initialisms, variable-name and package-name entries.
func goCRC1Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-CRC-001", Reference: "Go Code Review Comments, Initialisms: TCP",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Tcp`, Message: "write TCP in all capitals",
		},
		{
			ID: "GO-CRC-002", Reference: "Go Code Review Comments, Initialisms: IP",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Ip`, Message: "write IP in all capitals",
		},
		{
			ID: "GO-CRC-003", Reference: "Go Code Review Comments, Initialisms: CPU",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Cpu`, Message: "write CPU in all capitals",
		},
		{
			ID: "GO-CRC-004", Reference: "Go Code Review Comments, Initialisms: UTF8",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Utf8`, Message: "write UTF8 in all capitals",
		},
		{
			ID: "GO-CRC-005", Reference: "Go Code Review Comments, Initialisms: ASCII",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Ascii`, Message: "write ASCII in all capitals",
		},
		{
			ID: "GO-CRC-006", Reference: "Go Code Review Comments, Initialisms: CLI",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Cli`, Message: "write CLI in all capitals",
		},
		{
			ID: "GO-CRC-007", Reference: "Go Code Review Comments, Initialisms: GUI",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Gui`, Message: "write GUI in all capitals",
		},
		{
			ID: "GO-CRC-008", Reference: "Go Code Review Comments, Variable Names",
			Severity: rule.SeverityInfo, Category: rule.CategoryNaming,
			Pattern: `Ident/^[a-z][A-Za-z0-9]*[0-9]{2,}$/`,
			Message: "avoid numeric suffixes in variable names",
		},
		{
			ID: "GO-CRC-009", Reference: "Go Code Review Comments, Error Strings",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `StrLit/^.[A-Z]/`,
			Message: "string literals that read as messages should not start with a capital",
		},
		{
			ID: "GO-CRC-010", Reference: "Go Code Review Comments, Package Names",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Package/^[A-Z]/`,
			Message: "package names should be lowercase",
		},
		{
			ID: "GO-CRC-011", Reference: "Go Code Review Comments, Initialisms: CSV",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Csv`, Message: "write CSV in all capitals",
		},
	}
}
