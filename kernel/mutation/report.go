/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"strconv"

	"abate/kernel/verify"
)

// Report is the deterministic outcome of one mutation run.
//
// 参考 PITest 报告（逐变异体状态 + 变异分数），改进点=指纹由
// 种子 + 目标指纹 + 套件指纹 + 逐变异体规范串计算，且 Bytes 为固定字段序的
// 无浮点、无 map 迭代序编码，可直接做 SHA-256 回放校验。
type Report struct {
	// Seed is the experiment seed that drives point selection.
	Seed uint64

	// TargetFingerprint is the digest reported by the Target port.
	TargetFingerprint [32]byte

	// SuiteFingerprint is the digest reported by the Suite port.
	SuiteFingerprint [32]byte

	// Fingerprint is the SHA-256 digest of the canonical report body.
	Fingerprint [32]byte

	// Points lists every scheduled mutant in total order.
	Points []MutationPoint

	// Verdicts lists the fate of each mutant in the same order.
	Verdicts []verify.MutationVerdict

	// KillingTests lists the minimized killing test for each mutant.
	KillingTests []string

	// Score summarizes the verdicts through the KRN-VER verifier.
	Score verify.Score
}

// NewReport assembles a report and derives its deterministic fingerprint.
func NewReport(seed uint64, targetFP, suiteFP [32]byte, points []MutationPoint,
	verdicts []verify.MutationVerdict, killing []string,
) Report {
	rep := Report{
		Seed: seed, TargetFingerprint: targetFP, SuiteFingerprint: suiteFP,
		Points:       append([]MutationPoint(nil), points...),
		Verdicts:     append([]verify.MutationVerdict(nil), verdicts...),
		KillingTests: append([]string(nil), killing...),
		Score:        verify.Summarize(verdicts),
	}
	rep.Fingerprint = sha256.Sum256(rep.canonical())
	return rep
}

// Bytes returns the canonical byte-stable encoding of the whole report.
func (r Report) Bytes() []byte {
	body := r.canonical()
	out := make([]byte, 0, len(body)+sha256.Size)
	out = append(out, body...)
	out = append(out, r.Fingerprint[:]...)
	return out
}

// canonical encodes every field except the derived fingerprint.
func (r Report) canonical() []byte {
	var b bytes.Buffer
	writeUint(&b, r.Seed)
	_, _ = b.Write(r.TargetFingerprint[:])
	_, _ = b.Write(r.SuiteFingerprint[:])
	n := len(r.Points)
	if len(r.Verdicts) < n {
		n = len(r.Verdicts)
	}
	if len(r.KillingTests) < n {
		n = len(r.KillingTests)
	}
	writeUint(&b, uint64(n)) //nolint:gosec // slice lengths are non-negative and bounded.
	for i := 0; i < n; i++ {
		point := r.Points[i]
		writeText(&b, point.Node)
		writeText(&b, point.Original)
		writeText(&b, point.Mutated)
		writeText(&b, point.operatorKey())
		writeUint(&b, point.Priority)
		writeUint(&b, point.Selector)
		_, _ = b.Write(point.Fingerprint[:])
		writeText(&b, string(r.Verdicts[i]))
		writeText(&b, r.KillingTests[i])
	}
	writeUint(&b, r.Score.Total)
	writeUint(&b, r.Score.Killed)
	writeUint(&b, r.Score.Survived)
	writeUint(&b, r.Score.Equivalent)
	writeUint(&b, r.Score.Timeout)
	return b.Bytes()
}

// writeUint appends one big-endian uint64 to the buffer.
func writeUint(b *bytes.Buffer, v uint64) {
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[:], v)
	_, _ = b.Write(tmp[:])
}

// writeText appends a length-prefixed string to the buffer.
func writeText(b *bytes.Buffer, s string) {
	writeUint(b, uint64(len(s))) //nolint:gosec // string length is non-negative.
	_, _ = b.WriteString(s)
}

// ReportFingerprint returns the SHA-256 digest of a canonical report body.
//
// It is exposed for replay checks that must recompute a fingerprint from parts.
func ReportFingerprint(rep Report) [32]byte {
	return sha256.Sum256(rep.canonical())
}

// Describe returns a one-line deterministic summary of the report.
func (r Report) Describe() string {
	return "seed=" + strconv.FormatUint(r.Seed, 10) +
		" mutants=" + strconv.FormatUint(r.Score.Total, 10) +
		" killed=" + strconv.FormatUint(r.Score.Killed, 10) +
		" survived=" + strconv.FormatUint(r.Score.Survived, 10)
}
