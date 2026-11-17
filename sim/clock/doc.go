/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package clock provides the deterministic virtual clock that is the single
// time source of the Abate simulator.
//
// All simulator time is an integer nanosecond count. Floating-point time and any
// real wall-clock access (time.Now) are forbidden so that event ordering and
// hashing are reproducible across runs.
package clock
