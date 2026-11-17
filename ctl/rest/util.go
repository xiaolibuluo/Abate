/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"io"
	"net/http"
)

// maxRequestBodyBounds request bodies so a malicious client cannot exhaust memory.
const maxRequestBody = 1 << 20

// readBody reads a request body up to the fixed cap, returning an empty slice on
// any read error so the caller's validation logic can report a clean problem.
func readBody(req *http.Request) []byte {
	if req.Body == nil {
		return nil
	}
	limited := io.LimitReader(req.Body, maxRequestBody)
	data, err := io.ReadAll(limited)
	if err != nil {
		return nil
	}
	return data
}
