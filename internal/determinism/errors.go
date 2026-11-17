/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import "errors"

// errConfigMemoryMax indicates a configuration with no memory ceiling.
//
// The textual value is kept byte-identical to the frozen T01 golden
// observation (docs/refactor/golden/plugin-semantics-before-migration.txt).
var errConfigMemoryMax = errors.New("wasm: config: memory max pages must be positive")

// errConfigMemoryMin indicates a configuration whose minimum exceeds its max.
//
// The textual value is kept byte-identical to the frozen T01 golden observation.
var errConfigMemoryMin = errors.New("wasm: config: memory min pages exceeds max")
