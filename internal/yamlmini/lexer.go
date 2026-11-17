/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// A line is one logical input line after indentation measurement and comment
// stripping. Text excludes leading indentation; indent is the leading space
// count. Blank lines are retained because they may belong to a block scalar.
type line struct {
	text   string
	indent int
	num    int
	blank  bool
}

// lexLines splits data into logical lines, strips comments that appear outside
// quotes and rejects tabs used as indentation. CRLF is normalised to LF.
func lexLines(data []byte) ([]line, error) {
	if bytes.Contains(data, []byte("\x00")) {
		return nil, errors.New("yamlmini: NUL byte in document")
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	raw := strings.Split(text, "\n")
	out := make([]line, 0, len(raw))
	for i, s := range raw {
		num := i + 1
		indent := 0
		for indent < len(s) && s[indent] == ' ' {
			indent++
		}
		if indent < len(s) && s[indent] == '\t' {
			return nil, fmt.Errorf("yamlmini: line %d: tab in indentation", num)
		}
		content := stripComment(s[indent:])
		trimmed := strings.TrimRight(content, " \t")
		// Multi-document markers are explicitly out of scope; fail rather than
		// silently mis-lexing them as a scalar.
		if trimmed == "---" || strings.HasPrefix(trimmed, "--- ") ||
			trimmed == "..." || strings.HasPrefix(trimmed, "... ") {
			return nil, fmt.Errorf("yamlmini: line %d: multi-document markers are not supported", num)
		}
		// A complex key ("? key") is outside the string-key whitelist.
		if trimmed == "?" || strings.HasPrefix(trimmed, "? ") {
			return nil, fmt.Errorf("yamlmini: line %d: complex keys are not supported", num)
		}
		out = append(out, line{
			text:   trimmed,
			indent: indent,
			num:    num,
			blank:  strings.TrimSpace(trimmed) == "",
		})
	}
	return out, nil
}

// stripComment removes a trailing "# ..." comment. A '#' starts a comment only
// at the beginning of the content or after whitespace, and never inside quotes.
func stripComment(s string) string {
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if quote == '"' && c == '\\' && i+1 < len(s) {
				i++ // skip the escaped byte
				continue
			}
			if c == quote {
				if quote == '\'' && i+1 < len(s) && s[i+1] == '\'' {
					i++ // doubled quote inside a single-quoted scalar
					continue
				}
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '#' && (i == 0 || s[i-1] == ' ' || s[i-1] == '\t') {
			return s[:i]
		}
	}
	return s
}
