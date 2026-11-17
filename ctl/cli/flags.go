/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import "strings"

// Supported report formats.
const (
	formatConsole = "console"
	formatJSON    = "json"
)

// reportSchema identifies the deterministic report layout.
const reportSchema = "abate.report.v1"

// validFormat reports whether a report format is implemented.
func validFormat(format string) bool {
	return format == formatConsole || format == formatJSON
}

// reorderArgs moves flags ahead of positional arguments so the standard flag
// package accepts both "run file.yaml --seed 42" and "--seed 42 file.yaml".
func reorderArgs(args []string, valueFlags []string) []string {
	flags := make([]string, 0, len(args))
	positional := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		a := args[i]
		if !strings.HasPrefix(a, "-") || a == "-" {
			positional = append(positional, a)
			continue
		}
		flags = append(flags, a)
		if strings.Contains(a, "=") || !takesValue(a, valueFlags) || i+1 >= len(args) {
			continue
		}
		i++
		flags = append(flags, args[i])
	}
	return append(flags, positional...)
}

// takesValue reports whether a flag consumes the following argument.
func takesValue(flagName string, valueFlags []string) bool {
	name := strings.TrimLeft(flagName, "-")
	for _, candidate := range valueFlags {
		if name == candidate {
			return true
		}
	}
	return false
}

// runValueFlags lists the run flags that consume the following argument.
func runValueFlags() []string {
	return []string{"seed", "format", "out", "record"}
}

// replayValueFlags lists the replay flags that consume the following argument.
func replayValueFlags() []string {
	return []string{"format", "out"}
}

// checkValueFlags lists the check flags that consume the following argument.
func checkValueFlags() []string {
	return []string{"format", "severity", "rules", "out"}
}

// deriveValueFlags lists the derive flags that consume the following argument.
func deriveValueFlags() []string {
	return []string{"seed", "format", "out", "source"}
}
