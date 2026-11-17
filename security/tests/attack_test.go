/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tests

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"testing"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/ces/validate"
	"abate/ctl/rest"
	"abate/internal/yamlmini"
	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/report"
	"abate/plugin/signature"
	"abate/sim/record"
)

// baseExpr is one CES expression leaf. The expression parser treats it as an
// atomic operator reference, so composing it N times yields a chain of depth N.
const baseExpr = "ε"

// buildEnvelope marshals a fully valid CES document and overwrites its method
// expression, so the only thing a test varies is the expression string.
func buildEnvelope(t *testing.T, expr string) []byte {
	t.Helper()
	const base = int64(1_000_000_000) // 1e9 ns == 1s.
	e := schema.Experiment{
		APIVersion: schema.APIVersion,
		Kind:       schema.KindExperiment,
		Metadata:   schema.Metadata{Name: "security-attack"},
		Spec: schema.ExperimentSpec{
			Target: schema.Target{Kind: schema.TargetSimulator, Selector: map[string]string{"node": "n1"}},
			SteadyState: []schema.SteadyState{{
				ProbeRef:    "p1",
				Expectation: schema.Expectation{Type: schema.ExpectValue, Value: "ok"},
				Tolerance:   schema.Tolerance{Type: schema.ToleranceAbsolute, Value: 1},
				HoldFor:     schema.VirtualDuration(base * 2),
			}},
			Method:   schema.Method{Expression: expr},
			Probes:   []schema.Probe{{ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: schema.VirtualDuration(base)}},
			Rollback: schema.Rollback{Strategy: schema.StrategyAuto},
			Schedule: schema.Schedule{Continuous: &schema.ContinuousSpec{Duration: schema.VirtualDuration(base)}},
		},
	}
	data, err := yamlmini.Marshal(e)
	if err != nil {
		t.Fatalf("marshal envelope: %v", err)
	}
	return data
}

// TestThreatT1SignatureRejection proves T1: a malicious or unsigned plugin is
// refused by the trust-anchor verification chain. There is no fall-through to an
// unsigned state.
func TestThreatT1SignatureRejection(t *testing.T) {
	store := signature.NewTrustStore()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	if err := store.Add("abate/release", pub); err != nil {
		t.Fatalf("add trust anchor: %v", err)
	}

	m := signature.Manifest{
		ID:               "abate/evil",
		Version:          mustSemVer(t, "1.0.0"),
		Type:             signature.PluginType("reporter"),
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      "abate/release",
		BuildRef:         "deadbeef",
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		t.Fatalf("sign manifest: %v", err)
	}

	// A correctly signed manifest verifies.
	if err := signature.Verify(store, m); err != nil {
		t.Fatalf("signed manifest rejected: %v", err)
	}

	// T1a: a tampered signature is rejected.
	tampered := m
	tampered.Signature = append([]byte(nil), m.Signature...)
	tampered.Signature[0] ^= 0xff
	if err := signature.Verify(store, tampered); err == nil {
		t.Fatal("tampered signature accepted")
	}

	// T1b: a manifest with a corrupt content hash is rejected (distinct sentinel).
	corruptHash := m
	corruptHash.SHA256 = sha256.Sum256([]byte("tamper"))
	if err := signature.Verify(store, corruptHash); err == nil {
		t.Fatal("corrupt content hash accepted")
	}

	// T1c: an unknown trust anchor is rejected before any crypto runs.
	unknown := m
	unknown.TrustAnchor = "abate/ghost"
	if err := signature.Verify(store, unknown); err == nil {
		t.Fatal("unknown trust anchor accepted")
	}

	// T1d: an unsigned manifest (empty signature) is rejected, never trusted.
	unsigned := m
	unsigned.Signature = nil
	if err := signature.Verify(store, unsigned); err == nil {
		t.Fatal("unsigned manifest accepted")
	}
}

// TestThreatT6ParseInjection proves T6: a CES document carrying an expression
// that exceeds the operator-tree depth cap (64) is rejected during validation,
// neutralizing document-injection denial of service via deeply bracketed
// expressions.
func TestThreatT6ParseInjection(t *testing.T) {
	// 200 levels of parenthesis nesting: far beyond the 64-level cap. The depth
	// guard counts parentheses, so an attacker abusing document nesting must use
	// deeply bracketed expressions, which the compiler refuses.
	deep := strings.Repeat("(", 200) + baseExpr + strings.Repeat(")", 200)
	viols, err := validate.Validate(mustParseNormalize(t, deep))
	if err != nil {
		t.Fatalf("validate over-deep document: %v", err)
	}
	if !hasRule(viols, validate.RuleMethod003) {
		t.Fatalf("over-deep expression not rejected; violations = %v", viols)
	}

	// A valid shallow expression validates cleanly (the cap only rejects abuse).
	sviols, err := validate.Validate(mustParseNormalize(t, baseExpr))
	if err != nil {
		t.Fatalf("validate shallow document: %v", err)
	}
	if len(sviols) != 0 {
		t.Fatalf("shallow expression wrongly rejected: %v", sviols)
	}
}

// mustParseNormalize parses and normalizes a CES document carrying the given
// method expression, failing the test on any stage error.
func mustParseNormalize(t *testing.T, expr string) *schema.Experiment {
	t.Helper()
	data := buildEnvelope(t, expr)
	exp, err := parse.Parse(data, parse.FormatYAML)
	if err != nil {
		t.Fatalf("parse envelope: %v", err)
	}
	norm, err := parse.Normalize(exp)
	if err != nil {
		t.Fatalf("normalize envelope: %v", err)
	}
	return norm
}

// hasRule reports whether a violation set contains the given rule identifier.
func hasRule(viols []validate.Violation, id string) bool {
	for _, v := range viols {
		if v.RuleID == id {
			return true
		}
	}
	return false
}

// TestThreatT7ReportXSS proves T7: attacker-controlled message text rendered into
// the HTML report is neutralized by html/template auto-escaping. A raw <script>
// payload must never reach the document body.
func TestThreatT7ReportXSS(t *testing.T) {
	payload := "<script>alert('xss')</script>"
	findings := []rule.Finding{{
		RuleID:    "GO-EG-001",
		Path:      "attacker.go",
		Line:      42,
		Column:    7,
		Severity:  rule.SeverityError,
		Category:  rule.CategoryErrorProne,
		Message:   payload,
		Reference: "ref",
	}}

	out, err := report.BuildHTML(findings, mutation.Report{}, report.DefaultHTMLOptions())
	if err != nil {
		t.Fatalf("build html: %v", err)
	}
	html := string(out)

	if strings.Contains(html, payload) {
		t.Fatalf("raw script payload reached HTML output: %q", payload)
	}
	if !strings.Contains(html, "&lt;script&gt;") {
		t.Fatalf("payload was not escaped: %q", html)
	}
}

// TestThreatT8UnauthorizedAccess proves T8: the control plane refuses to bind a
// remote-listen address without an auth token, so the API is never exposed
// unauthenticated. (The per-request Bearer check is covered by ctl/rest tests.)
func TestThreatT8UnauthorizedAccess(t *testing.T) {
	// A remote bind without a token is refused at startup: the server can never
	// come up exposed and unauthenticated (PLAN-07 §4, threat T8).
	cfg := rest.DefaultServerConfig()
	cfg.Listen = "0.0.0.0:8420" // remote bind, no token set.
	if _, err := rest.New(cfg, rest.ServerOptions{}); err == nil {
		t.Fatal("remote bind accepted without an auth token")
	}
}

// TestThreatT10ArchiveTamper proves T10: a replay archive whose bytes have been
// altered fails verification on load, defeating evidence forgery. The frame hash
// chain plus the header linkage detect any single-byte mutation.
func TestThreatT10ArchiveTamper(t *testing.T) {
	var fp [32]byte
	fp[0] = 0xab

	rec := record.NewRecorder(fp)
	rec.Record(record.Frame{At: 1, Kind: record.KindSched, Payload: []byte("event-1")})
	rec.Record(record.Frame{At: 2, Kind: record.KindSched, Payload: []byte("event-2")})
	rec.Record(record.Frame{At: 3, Kind: record.KindRNG, Payload: []byte("draw-1")})
	arch, err := rec.Stop()
	if err != nil {
		t.Fatalf("seal archive: %v", err)
	}

	var clean bytes.Buffer
	if err := arch.Save(&clean); err != nil {
		t.Fatalf("save archive: %v", err)
	}

	// A clean archive loads and verifies.
	reloaded, err := record.Load(bytes.NewReader(clean.Bytes()))
	if err != nil {
		t.Fatalf("load clean archive: %v", err)
	}
	if err := reloaded.Verify(); err != nil {
		t.Fatalf("clean archive failed verify: %v", err)
	}

	// Flip a single byte of the on-disk bytes (an attacker forging a frame).
	tampered := clean.Bytes()
	tampered[len(tampered)-1] ^= 0xff

	loaded, err := record.Load(bytes.NewReader(tampered))
	if err == nil {
		// Load succeeded but the chain/header hash must still reject the mutation.
		if verr := loaded.Verify(); verr == nil {
			t.Fatal("tampered archive passed verification")
		}
	}
	// Either Load or Verify must refuse the altered bytes; reaching here means the
	// attack was intercepted.
}

func mustSemVer(t *testing.T, s string) signature.SemVer {
	t.Helper()
	v, err := signature.ParseSemVer(s)
	if err != nil {
		t.Fatalf("parse semver %q: %v", s, err)
	}
	return v
}
