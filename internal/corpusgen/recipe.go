/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import "strings"

// Recipe is the complete rendering contract of one source language.
type Recipe struct {
	Language    string
	Ext         string
	Tags        []string
	Header      string
	Footer      string
	AnchorClean string
	AnchorFull  string
	CleanBlock  string
	Blocks      []BlockTemplate
}

// renderPlaceholders substitutes the five deterministic recipe placeholders.
func renderPlaceholders(text string, index, seq int) string {
	number := (index*13+seq*7)%97 + 5
	other := (index*7+seq*11)%61 + 9
	return strings.NewReplacer(
		"@@I@@", itoa(index),
		"@@S@@", itoa(seq),
		"@@N@@", itoa(number),
		"@@M@@", itoa(other),
	).Replace(text)
}

// itoa renders a non-negative decimal without importing formatting machinery.
func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := make([]byte, 0, 12)
	for value > 0 {
		digits = append(digits, byte('0'+value%10))
		value /= 10
	}
	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}
	return string(digits)
}

// RenderBlock selects one construct family deterministically and fills its
// unique identifiers. Every call advances the same labeled RNG stream.
func (r Recipe) RenderBlock(rng *LabeledRNG, index, seq int) string {
	chosen := rng.Intn(len(r.Blocks))
	return renderPlaceholders(r.Blocks[chosen].Text, index, seq)
}

// RenderCleanBlock renders the minimal construct used by degrade-clean first
// units; all other files draw from the full high-fidelity construct pool.
func (r Recipe) RenderCleanBlock(rng *LabeledRNG, index, seq int,
	mode AnchorMode) string {
	if mode == AnchorClean {
		return renderPlaceholders(r.CleanBlock, index, seq)
	}
	return r.RenderBlock(rng, index, seq)
}

// AllRecipes returns the seven recipes in canonical report order.
func AllRecipes() []Recipe {
	return []Recipe{
		goRecipe(), pyRecipe(), jsRecipe(), cRecipe(),
		csRecipe(), rsRecipe(), ktRecipe(),
	}
}
