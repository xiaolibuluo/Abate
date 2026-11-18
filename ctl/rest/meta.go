/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

// Build metadata is fixed at compile time so the version endpoint stays
// deterministic and offline (no network, no build-time injection dependency).
const (
	// MetaVersion is the Abate release label.
	MetaVersion = "1.2.28"
	// MetaSchema is the CES schema identifier this build speaks.
	MetaSchema = "ces.tencent.com/v1"
	// MetaBuild marks the build provenance.
	MetaBuild = "deterministic"
	// MetaLicense is the dual-license label.
	MetaLicense = "AGPLv3 or commercial"
	// MetaSource is the canonical source repository.
	MetaSource = "https://github.com/xiaolibuluo/Abate"
	// MetaWatermark is the provenance watermark.
	MetaWatermark = "小黎的部落"
	// MetaAPIVersion is the REST resource prefix version.
	MetaAPIVersion = "v1"
)
