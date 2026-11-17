/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// AnchorMode selects the minimal first-unit anchor or the full capture anchor.
// It is a named string, never a boolean function parameter.
type AnchorMode string

// AnchorClean is the minimal first-unit anchor: no call tokens.
const AnchorClean AnchorMode = "clean"

// AnchorFull is the capture anchor carrying f() and g() call tokens.
const AnchorFull AnchorMode = "full"
