/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

// Format selects the encoding of a CES document.
type Format string

// Supported document formats.
const (
	FormatAuto Format = "auto"
	FormatYAML Format = "yaml"
	FormatJSON Format = "json"
)
