/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package observe normalizes probe samples into a totally ordered observation
// stream consumed by the verifier and the recorder.
//
// Probes execute inside the simulator and never touch real I/O. Collection
// order is always the lexicographic probe-ID order, so registration order can
// never influence the stream or its digest.
package observe
