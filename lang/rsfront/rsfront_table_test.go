/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// TestRsTableLookups covers both hit and miss paths of the keyword/infix/item
// tables so the mutation tests above are known to exercise real table-driven
// behavior.
func TestRsTableLookups(t *testing.T) {
	if !rsContains(rsKeywords, "fn") || rsContains(rsKeywords, "nope") {
		t.Fatal("rsContains broken")
	}
	if !rsHasPrefix("/// doc", rsDocPrefixes) ||
		rsHasPrefix("// plain", rsDirectivePrefixes) {
		t.Fatal("rsHasPrefix broken")
	}
	if rsInfixPower(rsInfixBindings, "*") != 70 ||
		rsInfixPower(rsInfixBindings, ",") != 0 {
		t.Fatal("rsInfixPower broken")
	}
	if !rsInfixRight(rsInfixBindings, "=") ||
		rsInfixRight(rsInfixBindings, "+") ||
		rsInfixRight(rsInfixBindings, ",") {
		t.Fatal("rsInfixRight broken")
	}
	if rule, ok := rsItemRuleFor(rsItemRules, "trait"); !ok ||
		rule.kind != air.KindTraitDecl {
		t.Fatal("rsItemRuleFor broken")
	}
	if _, ok := rsItemRuleFor(rsItemRules, "nope"); ok {
		t.Fatal("rsItemRuleFor accepted an unknown keyword")
	}
}

// TestRsTokenLookups covers both hit and miss paths of the token/kind/note
// lookup helpers.
func TestRsTokenLookups(t *testing.T) {
	if rsTokenNodeKind(frontkit.TokenStr) != air.KindStrLit ||
		rsTokenNodeKind(frontkit.TokenEOF) != air.KindRegion {
		t.Fatal("rsTokenNodeKind broken")
	}
	if !rsIsLexicalKind(frontkit.TokenNum) ||
		rsIsLexicalKind(frontkit.TokenEOF) {
		t.Fatal("rsIsLexicalKind broken")
	}
	if !rsIsTrivia(frontkit.TokenNewline) || rsIsTrivia(frontkit.TokenIdent) {
		t.Fatal("rsIsTrivia broken")
	}
	if !rsIsOpenBracket("{") || rsIsOpenBracket("}") {
		t.Fatal("rsIsOpenBracket broken")
	}
	if rsCloseOf("[") != "]" || rsCloseOf("{") != "}" || rsCloseOf("?") != ")" {
		t.Fatal("rsCloseOf broken")
	}
	if rsNoteFor(rsNotes, "macro") != "macro" || rsNoteFor(rsNotes, "nope") != "" {
		t.Fatal("rsNoteFor broken")
	}
	if rsNormalizeDialectKey("  RUST 2021 ") != "2021" ||
		rsNormalizeDialectKey("") != "" {
		t.Fatal("rsNormalizeDialectKey broken")
	}
	if rsEffectiveLead(rsTokenize(t, "pub(crate) fn f() {}\n"), 0, 12) != "fn" {
		t.Fatal("rsEffectiveLead did not skip a visibility modifier")
	}
}
