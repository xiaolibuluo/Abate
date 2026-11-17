/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

// IsAlpha reports whether b is an ASCII letter.
func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// IsDigit reports whether b is an ASCII decimal digit.
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// IsKeyByte reports whether b may appear inside a sensitive key or marker word.
func isKeyByte(b byte) bool {
	return isAlpha(b) || isDigit(b) || b == '_' || b == '-' || b == '.'
}

// IsSpace reports whether b is ASCII whitespace.
func isSpace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r', '\v', '\f':
		return true
	default:
		return false
	}
}

// FoldLower maps one ASCII upper-case letter to lower case and leaves every
// other byte unchanged.
func foldLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// HasFoldPrefix reports whether s starts with word, ASCII case-insensitively.
func hasFoldPrefix(s, word string) bool {
	if len(s) < len(word) {
		return false
	}
	for i := 0; i < len(word); i++ {
		if foldLower(s[i]) != foldLower(word[i]) {
			return false
		}
	}
	return true
}

// IsValueSeparator reports whether b ends a key/value secret value. In urlMode,
// '/' and '?' also end the value so URL structure is preserved. The placeholder
// itself contains none of these bytes, which keeps redaction idempotent.
func isValueSeparator(b byte, urlMode bool) bool {
	switch b {
	case ' ', '\t', '\n', '\r', '\v', '\f', '&', '"', '\'', ';', '#', '`':
		return true
	case '/', '?':
		return urlMode
	default:
		return false
	}
}

// IsHeaderSeparator reports whether b ends an Authorization header value.
// Spaces and quotes are deliberately absent: a canonical "Bearer <token>" value
// carries a space, and a quote in the middle of a value must not terminate the
// scan early and leave the rest of the credential in the log. A quoted value at
// the very start is handled separately by skipQuoted.
func isHeaderSeparator(b byte) bool {
	switch b {
	case '\n', '\r', '&', ';', '#', '`':
		return true
	default:
		return false
	}
}

// SkipQuoted returns the width of a leading quoted value, including both
// quotes, so a quoted secret never leaves its payload behind in the log.
func skipQuoted(v string) (int, bool) {
	if len(v) == 0 || (v[0] != '"' && v[0] != '\'') {
		return 0, false
	}
	quote := v[0]
	n := 1
	for n < len(v) && v[n] != quote {
		n++
	}
	if n < len(v) {
		n++
	}
	return n, true
}

// SkipValue returns the number of leading bytes of v that belong to the secret
// value. A nested "Bearer <token>" value is consumed in full so no fragment of
// a credential survives redaction.
func skipValue(v string, urlMode bool) int {
	if n, ok := skipQuoted(v); ok {
		return n
	}
	n := 0
	if hasFoldPrefix(v, bearerWord) {
		p := len(bearerWord)
		if p < len(v) && isSpace(v[p]) {
			n = p
			for n < len(v) && isSpace(v[n]) {
				n++
			}
		}
	}
	for n < len(v) && !isValueSeparator(v[n], urlMode) {
		n++
	}
	return n
}

// SkipHeaderValue returns the number of leading bytes of v that belong to an
// Authorization header value, stopping at a structural separator or line end.
func skipHeaderValue(v string) int {
	if n, ok := skipQuoted(v); ok {
		return n
	}
	n := 0
	for n < len(v) && !isHeaderSeparator(v[n]) {
		n++
	}
	return n
}
