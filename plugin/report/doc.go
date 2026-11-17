/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package report renders static rule findings and a mutation report into portable
// artifact formats. It currently produces two deterministic encodings:
//
//   - SARIF 2.1.0 (OASIS standard) via the embedded and validated schema, and
//   - a compact, field-ordered JSON summary of findings plus the mutation report.
//
// Both encoders are pure functions of their inputs: identical findings and
// reports always yield byte-identical output, which lets the control plane hash
// and replay artifacts. The SARIF encoder is validated against the official OASIS
// 2.1.0 JSON-Schema through the dependency-free validator in the schema
// subpackage, replacing the external Microsoft SARIF validator.
package report
