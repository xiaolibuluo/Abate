/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package chaos builds deterministic damaged corpora from in-memory source
// buffers: half truncations, byte splices, token drops and sha256-deduplicated
// merges. A random generator only chooses damage offsets, never which upstream
// file enters the corpus, and no damaged byte is ever written back to the
// read-only source tree.
package chaos
