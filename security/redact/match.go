/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

// BearerWidth returns the width of a "Bearer" scheme prefix at i, including
// the whitespace that follows it. The caller enforces the word boundary.
func bearerWidth(s string, i int) (int, bool) {
	rest := s[i:]
	if !hasFoldPrefix(rest, bearerWord) {
		return 0, false
	}
	p := len(bearerWord)
	if p >= len(rest) || !isSpace(rest[p]) {
		return 0, false
	}
	for p < len(rest) && isSpace(rest[p]) {
		p++
	}
	return p, true
}

// AuthorizationWidth returns the width of an "Authorization:" header prefix at
// i, including the colon and the surrounding optional whitespace.
func authorizationWidth(s string, i int) (int, bool) {
	rest := s[i:]
	if !hasFoldPrefix(rest, authorizationWord) {
		return 0, false
	}
	p := len(authorizationWord)
	for p < len(rest) && isSpace(rest[p]) {
		p++
	}
	if p >= len(rest) || rest[p] != ':' {
		return 0, false
	}
	p++
	for p < len(rest) && isSpace(rest[p]) {
		p++
	}
	return p, true
}

// SensitiveKeyEnd returns the index of '=' when s[i:] starts with a known
// sensitive key immediately followed by '='. Unknown keys fail open to the
// original text, which is the documented fail-safe behaviour.
func sensitiveKeyEnd(s string, i int) (int, bool) {
	j := i
	for j < len(s) && isKeyByte(s[j]) {
		j++
	}
	if j == i || j >= len(s) || s[j] != '=' {
		return 0, false
	}
	if !isSensitiveKey(s[i:j]) {
		return 0, false
	}
	return j, true
}
