/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package pymutate

import (
	"strings"

	"abate/lang/frontkit"
	"abate/lang/mutcore"
)

// pyClassState enumerates the states of the oracle-exclusion machine. The
// machine tracks def/class headers and INDENT/DEDENT pairs so tokens
// inside a test block are never classified as mutable: rewriting the
// assertion oracle itself would make every kill verdict meaningless.
type pyClassState int

const (
	// pyStateNormal classifies production tokens against the shared sets.
	pyStateNormal pyClassState = iota
	// pyStateSawDef waits for the function name after a def keyword.
	pyStateSawDef
	// pyStateSawClass waits for the class name after a class keyword.
	pyStateSawClass
	// pyStateHeaderTest waits for the INDENT opening the suite of a def
	// or class whose name pytest would collect (test_* / Test*).
	pyStateHeaderTest
	// pyStateSkipping swallows every token of a test block until the
	// DEDENT matching the INDENT that opened it.
	pyStateSkipping
)

// pyClassifier is the stateful per-stream oracle-exclusion classifier.
// The lexer emits INDENT/DEDENT synthetic tokens, so the classifier can
// track block entry and exit without parsing the source text.
type pyClassifier struct {
	state pyClassState
	depth int
}

// newPyClassifier returns the classifier driving one workspace scan. The
// state resets on every EOF marker, so one instance can classify the
// token streams of many files in sequence without leaking block depth
// across file boundaries.
func newPyClassifier() *pyClassifier {
	return &pyClassifier{}
}

// classify inspects one token in stream order and returns its mutation
// class, or false when the token is excluded from the mutable set. The
// path argument is accepted to match the LexicalScan callback shape; the
// oracle exclusion is purely structural and never depends on the file.
func (c *pyClassifier) classify(_ string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch c.state {
	case pyStateNormal:
		return c.classifyNormal(lt)
	case pyStateSawDef:
		return c.classifyDefName(lt, "test_")
	case pyStateSawClass:
		return c.classifyDefName(lt, "Test")
	case pyStateHeaderTest:
		return c.classifyHeaderTest(lt)
	default:
		return c.classifySkipping(lt)
	}
}

// classifyNormal handles production code: def/class keywords arm the
// header machine, every other token faces the shared lexical sets.
func (c *pyClassifier) classifyNormal(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.reset()
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
		return mutcore.ClassUnknown, false
	}
	switch lt.Text {
	case "def":
		c.state = pyStateSawDef
		return mutcore.ClassUnknown, false
	case "class":
		c.state = pyStateSawClass
		return mutcore.ClassUnknown, false
	}
	return mutcore.ClassifyLexical("", lt)
}

// classifyDefName inspects the declared name after def/class: a name
// pytest would collect (test_* functions, Test* classes) moves the
// machine into the header-waiting state; any other name returns to
// production classification. Names themselves are never mutable.
func (c *pyClassifier) classifyDefName(lt frontkit.Token, prefix string) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.reset()
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
		c.state = pyStateNormal
		return mutcore.ClassUnknown, false
	}
	if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, prefix) {
		c.state = pyStateHeaderTest
		return mutcore.ClassUnknown, false
	}
	c.state = pyStateNormal
	return mutcore.ClassUnknown, false
}

// classifyHeaderTest waits out the remainder of a test def/class header
// (parameter list, decorators already consumed, return annotations) until
// the INDENT opening its suite, which starts the skipped block.
func (c *pyClassifier) classifyHeaderTest(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenIndent:
		c.state = pyStateSkipping
		c.depth = 1
	case frontkit.TokenEOF, frontkit.TokenDedent:
		c.reset()
	}
	return mutcore.ClassUnknown, false
}

// classifySkipping swallows the tokens of a test block. Nested INDENT
// pairs deepen the block; the DEDENT closing the outermost suite returns
// the machine to production classification.
func (c *pyClassifier) classifySkipping(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenIndent:
		c.depth++
	case frontkit.TokenDedent:
		c.depth--
		if c.depth <= 0 {
			c.state = pyStateNormal
		}
	case frontkit.TokenEOF:
		c.reset()
	}
	return mutcore.ClassUnknown, false
}

// reset returns the machine to its initial state, used at file boundaries.
func (c *pyClassifier) reset() {
	c.state = pyStateNormal
	c.depth = 0
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .py file through the shared Python lexical tables, skipping the
// test blocks that form the assertion oracle, in deterministic
// (path, offset) order.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	classifier := newPyClassifier()
	return mutcore.LexicalScan(root, map[string]bool{".py": true}, isTestFile,
		pyLexSpec{}, classifier.classify)
}
