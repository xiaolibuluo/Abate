/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package scheduler implements the deterministic discrete-event loop.
//
// The virtual clock is the only time source, every event carries a scheduler
// assigned total-order sequence, and the run digest makes replay byte-stable.
// The package depends only on the standard library and kernel-local types.
package scheduler
