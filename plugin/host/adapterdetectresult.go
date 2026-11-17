/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

// AdapterDetectResult is the serializable outcome of an EntryDetect call. It
// records the matched adapter identifier and whether detection succeeded, so the
// host can return a stable, JSON-encoded result for the control plane.
type AdapterDetectResult struct {
	// ID is the stable adapter identifier produced by detection.
	ID string

	// OK reports whether the adapter matched the supplied environment.
	OK bool
}
