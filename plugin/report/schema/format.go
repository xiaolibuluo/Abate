/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// checkFormat enforces a small set of JSON-Schema format keywords. Unknown
// formats are accepted so the validator remains forward-compatible.
func checkFormat(name, str string) error {
	switch name {
	case "uri":
		if _, err := url.ParseRequestURI(str); err != nil {
			return fmt.Errorf("value %q is not a valid URI: %w", str, err)
		}
	case "uri-reference":
		if strings.TrimSpace(str) == "" {
			return errors.New("value is an empty uri-reference")
		}
	case "date-time":
		if _, err := time.Parse(time.RFC3339, str); err != nil {
			if _, err2 := time.Parse(time.RFC3339Nano, str); err2 != nil {
				return fmt.Errorf("value %q is not a valid RFC3339 date-time", str)
			}
		}
	default:
		// Forward-compatible: unrecognized formats are not enforced.
	}
	return nil
}

// summarize renders a scalar value as a short, stable string for error messages.
func summarize(v any) string {
	switch x := v.(type) {
	case string:
		if len(x) > 32 {
			return x[:32] + "..."
		}
		return x
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", x)
	}
}
