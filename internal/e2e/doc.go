/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package e2e runs the deterministic end-to-end corpus gate. It parses a
// stratified sample with all seven frontends, applies real in-memory mutations
// and damaged inputs, and emits a fully reproducible report with four digest
// chains. No wall-clock value, absolute path, runtime version string or map
// iteration ever enters the compared report bytes.
package e2e
