/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"crypto/subtle"
	"net/http"
	"strings"
)

// verifyBearer checks the Authorization header against the configured token
// using a constant-time comparison so timing does not leak the secret.
func verifyBearer(req *http.Request, token string) bool {
	if strings.TrimSpace(token) == "" {
		return false
	}
	header := req.Header.Get("Authorization")
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return false
	}
	present := []byte(strings.TrimSpace(header[len(prefix):]))
	want := []byte(token)
	return subtle.ConstantTimeCompare(present, want) == 1
}
