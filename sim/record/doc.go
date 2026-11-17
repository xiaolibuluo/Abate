/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package record implements the SIM-REC module: a write-ahead recorder and a
// frame-by-frame replayer that seal a deterministic experiment run into a
// tamper-evident, byte-stable archive.
//
// The Recorder subscribes to read-only sources (scheduler, observation pipeline,
// random stream, virtual clock and fault injectors) and wraps the scheduler
// handler, so that no recorded frame ever alters the event stream it observes.
// The resulting Archive is a versioned binary with a header, an append-only frame
// block, a sequential index and a SHA-256 tail checksum; every frame also carries
// the hash of its predecessor, forming a chain that breaks on any tampering.
//
// Compare and Diff are a standard-library rewrite of the deep-recursive
// comparison algorithm from google/go-cmp. The improvement is a deterministic
// structural difference: the frame domain is walked field by field without
// reflection, the first divergence carries its complete path and expected or
// actual values, statistics are aggregated, and all map-derived output is
// sorted by key so no iteration order can leak into the report.
package record
