/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"crypto/sha256"
	"slices"
	"strings"
)

// AIRSuite adapts the deterministic test list to the mutation.Suite port.
type AIRSuite struct {
	tests []string
	cover map[string]uint64
	fp    [32]byte
}

// NewAIRSuite builds a suite that covers every node with a fixed test list.
func NewAIRSuite(nodes []string, tests []string) (*AIRSuite, error) {
	if len(tests) == 0 {
		return nil, ErrEmptySuite
	}
	sortedTests := append([]string(nil), tests...)
	slices.Sort(sortedTests)
	sortedNodes := append([]string(nil), nodes...)
	slices.Sort(sortedNodes)
	cover := make(map[string]uint64, len(sortedNodes))
	for _, node := range sortedNodes {
		cover[node] = uint64(len(sortedTests)) //nolint:gosec // a slice length is non-negative.
	}
	var b strings.Builder
	for _, test := range sortedTests {
		_, _ = b.WriteString(test)
		_ = b.WriteByte('|')
	}
	_ = b.WriteByte('#')
	for _, node := range sortedNodes {
		_, _ = b.WriteString(node)
		_ = b.WriteByte(';')
	}
	return &AIRSuite{
		tests: sortedTests, cover: cover,
		fp: sha256.Sum256([]byte(b.String())),
	}, nil
}

// Fingerprint returns the stable digest of the test list and covered nodes.
func (s *AIRSuite) Fingerprint() [32]byte {
	return s.fp
}

// Tests returns the sorted test identifiers.
func (s *AIRSuite) Tests() []string {
	return append([]string(nil), s.tests...)
}

// Coverage returns the number of tests covering one node.
func (s *AIRSuite) Coverage(node string) uint64 {
	return s.cover[node]
}
