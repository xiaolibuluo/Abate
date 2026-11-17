/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

// Category groups rules by the class of defect they detect.
type Category string

// Category values used by the built-in rule sets.
const (
	// CategoryNaming covers identifier and package naming conventions.
	CategoryNaming Category = "naming"
	// CategoryStyle covers formatting, comments and idiom rules.
	CategoryStyle Category = "style"
	// CategoryErrorProne covers constructs that commonly hide defects.
	CategoryErrorProne Category = "error-prone"
	// CategoryPerformance covers avoidable allocations and copying.
	CategoryPerformance Category = "performance"
	// CategoryConcurrency covers goroutine, channel and lock usage.
	CategoryConcurrency Category = "concurrency"
	// CategoryBoundary covers API and interface boundary conventions.
	CategoryBoundary Category = "boundary"
)

// Valid reports whether the category is one of the allocated values.
func (c Category) Valid() bool {
	switch c {
	case CategoryNaming, CategoryStyle, CategoryErrorProne,
		CategoryPerformance, CategoryConcurrency, CategoryBoundary:
		return true
	default:
		return false
	}
}

// String returns the stable wire text of the category.
func (c Category) String() string {
	return string(c)
}
