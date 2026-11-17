/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"strconv"
)

// Capture binds one AIR node to a metavariable of a structural pattern.
type Capture struct {
	// Name is the metavariable name without its leading dollar sign.
	Name string
	// Kind is the AIR node kind wire name.
	Kind string
	// Text is the node's own attribute text.
	Text string
	// Line is the one-based source line of the node.
	Line int
	// Column is the one-based source column of the node.
	Column int
}

// String renders the capture in a stable one-line form.
func (c Capture) String() string {
	return fmt.Sprintf("%s=%s(%s:%d:%d)", c.Name, c.Text,
		c.Kind, c.Line, c.Column)
}

// Canonical renders the capture in a byte-stable form used by fingerprints.
func (c Capture) Canonical() string {
	return c.Name + "\x00" + c.Kind + "\x00" + c.Text + "\x00" +
		strconv.Itoa(c.Line) + "\x00" + strconv.Itoa(c.Column)
}
