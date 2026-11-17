/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"errors"
	"fmt"
	"slices"

	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// ErrUnsupportedLanguage reports a language tag with no grammar facts.
var ErrUnsupportedLanguage = errors.New("coverage: unsupported language")

// indexedLanguages is the canonical order of the seven registered languages.
var indexedLanguages = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// Index joins the six embedded upstream grammars with the self-built
// Go inventory. Coverage numbers are computed from a real AIR unit, never from
// a static guess.
type Index struct {
	grammars  map[string]*grammar.Grammar
	mappings  map[string][]grammar.Mapping
	inventory GoInventory
}

// NewCoverageIndex loads the six embedded grammars and the Go inventory.
func NewCoverageIndex(goInventoryPath string) (Index, error) {
	inventory, err := LoadGoInventory(goInventoryPath)
	if err != nil {
		return Index{}, err
	}
	out := Index{
		grammars:  make(map[string]*grammar.Grammar),
		mappings:  make(map[string][]grammar.Mapping),
		inventory: inventory,
	}
	for _, tag := range grammar.SupportedLanguages() {
		facts, loadErr := grammar.Load(tag)
		if loadErr != nil {
			return Index{}, loadErr
		}
		out.grammars[tag] = facts
		out.mappings[tag] = grammar.Map(facts, grammar.DefaultRules(tag))
	}
	return out, nil
}

// Languages returns the seven language tags in canonical order.
func (Index) Languages() []string {
	return slices.Clone(indexedLanguages)
}

// Expected returns the construct names of one language in ascending order. The
// Go row comes from the self-built inventory, not from an embedded grammar.
func (c Index) Expected(lang string) []string {
	if lang == "go" {
		return c.inventory.Nodes()
	}
	out := make([]string, 0, len(c.mappings[lang]))
	for _, mapping := range c.mappings[lang] {
		out = append(out, mapping.Node)
	}
	slices.Sort(out)
	return out
}

// IsMeasured reports whether one language has real embedded grammar facts.
func (c Index) IsMeasured(lang string) bool {
	if lang == "go" {
		return c.inventory.IsMeasured()
	}
	_, found := c.grammars[lang]
	return found
}

// Status returns the explicit measurement label of one language.
func (c Index) Status(lang string) string {
	if !c.IsMeasured(lang) {
		return statusNotMeasured
	}
	return "MEASURED"
}

// Coverage computes the construct coverage of one language from a real parsed
// unit. The Go row is refused with ErrNotMeasured instead of a fabricated ratio.
func (c Index) Coverage(lang string, u *air.Unit) (fkcov.Coverage, error) {
	if lang == "go" {
		return fkcov.Coverage{}, ErrNotMeasured
	}
	facts, found := c.grammars[lang]
	if !found {
		return fkcov.Coverage{}, fmt.Errorf("%w: %s",
			ErrUnsupportedLanguage, lang)
	}
	return fkcov.Compute(facts, c.mappings[lang], fkcov.Observe(u)), nil
}

// Degradation returns the degraded-mapping share of one language. Go has no
// measured mappings, so its share is zero and callers must use Status first.
func (c Index) Degradation(lang string, u *air.Unit) float64 {
	if lang == "go" {
		return 0
	}
	return Degradation(c.mappings[lang], fkcov.Observe(u))
}
