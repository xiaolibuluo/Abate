/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

import (
	"sort"
	"strings"
)

// Bearer and Authorization are the two header-style secret markers, stored in
// lower case so comparison can fold ASCII case without allocation.
const (
	bearerWord        = "bearer"
	authorizationWord = "authorization"
)

// The sensitive key table is a fixed, sorted, lower-case slice rather than a
// map, so lookup never depends on map iteration order. The order must stay
// sorted for the binary search in isSensitiveKey to remain valid.
var sensitiveKeys = []string{
	"access_token",
	"api-key",
	"api_key",
	"apikey",
	"authorization",
	"bearer",
	"client_secret",
	"passwd",
	"password",
	"private_key",
	"pwd",
	"secret",
	"session",
	"token",
}

// IsSensitiveKey reports whether key names a secret value in any letter case.
// The lookup is a deterministic binary search over the sorted table.
func isSensitiveKey(key string) bool {
	lower := strings.ToLower(key)
	i := sort.SearchStrings(sensitiveKeys, lower)
	return i < len(sensitiveKeys) && sensitiveKeys[i] == lower
}
