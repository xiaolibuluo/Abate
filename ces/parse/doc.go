/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package parse turns CES YAML/JSON byte streams into executable experiments.
//
// Parsing is strict (unknown fields are rejected), normalization fills
// defaults and resolves references, and the operator expression compiler
// assembles a kernel algebra tree. All work is pure and reentrant.
package parse
