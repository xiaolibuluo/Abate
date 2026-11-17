/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

// requestCounter yields deterministic, monotonic request identifiers. A counter
// (not a random UUID) keeps error instances reproducible across runs and avoids
// any entropy source (PLAN-08 determinism invariant).
var requestCounter atomic.Uint64

// requestInstance returns a stable request identifier for problem documents.
func requestInstance(*http.Request) string {
	id := requestCounter.Add(1)
	return "/requests/" + strconv.FormatUint(id, 10)
}
