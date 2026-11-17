/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"strings"
	"testing"
)

func TestItoaAndPlaceholders(t *testing.T) {
	t.Parallel()
	if itoa(0) != "0" || itoa(20520) != "20520" {
		t.Fatal("itoa failed")
	}
	rendered := renderPlaceholders("n@@I@@s@@S@@a@@N@@b@@M@@", 12, 3)
	if !strings.HasPrefix(rendered, "n12s3a") {
		t.Fatalf("placeholders = %q", rendered)
	}
	if strings.Contains(rendered, "@@") {
		t.Fatalf("unresolved placeholder: %q", rendered)
	}
}

func TestAllRecipesHaveRealConstructs(t *testing.T) {
	t.Parallel()
	recipes := AllRecipes()
	if len(recipes) != 7 {
		t.Fatalf("recipe count = %d", len(recipes))
	}
	for _, recipe := range recipes {
		if recipe.Language == "" || recipe.Ext == "" || len(recipe.Tags) < 5 {
			t.Fatalf("incomplete recipe %+v", recipe)
		}
		if len(recipe.Blocks) < 5 || recipe.CleanBlock == "" {
			t.Fatalf("recipe %s has too few constructs", recipe.Language)
		}
		if HeaderHasToken(recipe.Header) {
			t.Fatalf("recipe %s header has mutation token", recipe.Language)
		}
		if !strings.Contains(recipe.AnchorClean, "1 + 2 == 3") ||
			!strings.Contains(recipe.AnchorFull, "f()") {
			t.Fatalf("recipe %s anchor invalid", recipe.Language)
		}
		rng := NewLabeledRNG(7, recipe.Language)
		block := recipe.RenderBlock(&rng, 3, 2)
		if strings.Contains(block, "@@") || len(block) < 40 {
			t.Fatalf("recipe %s block = %q", recipe.Language, block)
		}
		clean := recipe.RenderCleanBlock(&rng, 3, 2, AnchorClean)
		if strings.Contains(clean, "@@") || len(clean) < 10 {
			t.Fatalf("recipe %s clean block = %q", recipe.Language, clean)
		}
	}
}

func TestConvergenceIdentifiersAreUniqueAndLegal(t *testing.T) {
	t.Parallel()
	for _, language := range []string{"go", "py", "js", "c", "cs", "rs", "kt"} {
		first := convergenceBlock(language, 3, 4, 400)
		second := convergenceBlock(language, 3, 5, 400)
		if first == second || strings.Contains(first, "@@") {
			t.Fatalf("%s convergence blocks not unique", language)
		}
		if len(first) < 380 || len(first) > 460 {
			t.Fatalf("%s convergence span = %d", language, len(first))
		}
	}
	if convergenceBlock("unknown", 1, 1, 100) != "" {
		t.Fatal("unknown language produced a block")
	}
	if longIdentifier(1, 2, 50, StyleSnake) == longIdentifier(1, 2, 50, StyleCamel) {
		t.Fatal("identifier styles collide")
	}
}

func TestManifestRefsAndDedication(t *testing.T) {
	t.Parallel()
	refs := SelfResearchedRefs()
	spdx, upstream, source, evidence := refs.SPDX, refs.UpstreamURL,
		refs.Source, refs.EvidenceRelPath
	if spdx != "CC0-1.0" || upstream != "" ||
		source != "self_researched" || evidence != DedicationRelPath {
		t.Fatalf("refs = %+v", refs)
	}
	if !strings.Contains(DedicationText, "CC0-1.0") ||
		!strings.Contains(DedicationText, "waives all copyright") {
		t.Fatal("dedication text incomplete")
	}
}
