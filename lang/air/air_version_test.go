/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"encoding/binary"
	"errors"
	"strings"
	"testing"
)

// mustVersion builds a test version range or fails.
func mustVersion(t *testing.T, minimum, current uint16) Version {
	t.Helper()
	v, err := NewVersionRange(minimum, current)
	if err != nil {
		t.Fatalf("NewVersionRange(%d,%d): %v", minimum, current, err)
	}
	return v
}

// reorderSections rewrites the framed section stream in a different order.
func reorderSections(t *testing.T, archive []byte, order []byte) []byte {
	t.Helper()
	body := archiveBody(archive)
	parts := map[byte][]byte{}
	cur := &cursor{data: body[headerSize:]}
	for cur.remaining() > 0 {
		id := cur.u8()
		size := cur.u32()
		payload := cur.take(int(size))
		if cur.err != nil {
			t.Fatalf("framing: %v", cur.err)
		}
		parts[id] = appendSection(nil, id, payload)
	}
	out := append([]byte(nil), body[:headerSize]...)
	for _, id := range order {
		part, ok := parts[id]
		if !ok {
			t.Fatalf("missing section %d", id)
		}
		out = append(out, part...)
	}
	return reseal(out)
}

// TestVersionNegotiate covers empty, boundary, and out-of-range requests.
func TestVersionNegotiate(t *testing.T) {
	t.Parallel()
	current := NewVersion()
	if current.Current() != SchemaVersion || current.MinSupported() != MinSchemaVersion {
		t.Fatalf("build version: %s", current.String())
	}
	if got, err := current.Negotiate(0); err != nil || got != SchemaVersion {
		t.Fatalf("empty request: %d %v", got, err)
	}
	for _, version := range []uint16{MinSchemaVersion, SchemaVersion} {
		if got, err := current.Negotiate(version); err != nil || got != version {
			t.Fatalf("negotiate v%d: %d %v", version, got, err)
		}
	}
	for _, version := range []uint16{SchemaVersion + 1, 0xFFFF} {
		if _, err := current.Negotiate(version); !errors.Is(err, ErrUnsupportedAIRVersion) {
			t.Fatalf("v%d must be rejected: %v", version, err)
		}
	}
	if current.String() != "v1..v2" {
		t.Fatalf("version rendering: %s", current.String())
	}
}

// TestVersionAccepts exercises every rejection branch of version negotiation.
func TestVersionAccepts(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		minimum    uint16
		version    uint16
		readerMin  uint16
		readerMax  uint16
		wantErr    bool
		wantReason string
	}{
		{"v1 on current", 1, 1, 1, 2, false, ""},
		{"v2 on current", 1, 2, 1, 2, false, ""},
		{"v2 on legacy", 1, 2, 1, 1, true, "newer than the reader"},
		{"future minimum", 3, 3, 1, 2, true, "newer than the reader"},
		{"old archive on new reader", 1, 1, 2, 2, true, "older than the reader"},
		{"zero version", 1, 0, 1, 2, true, "zero writer version"},
		{"zero minimum", 0, 1, 1, 2, true, "zero writer minimum"},
		{"reversed header", 2, 1, 1, 2, true, "writer minimum exceeds writer version"},
	}
	for _, tc := range cases {
		reader := mustVersion(t, tc.readerMin, tc.readerMax)
		err := reader.Accepts(tc.minimum, tc.version)
		if tc.wantErr {
			if !errors.Is(err, ErrUnsupportedAIRVersion) {
				t.Fatalf("%s: wrong sentinel %v", tc.name, err)
			}
			if !strings.Contains(err.Error(), tc.wantReason) {
				t.Fatalf("%s: missing reason %q in %q", tc.name, tc.wantReason, err)
			}
			if !strings.Contains(err.Error(), "reader supports") {
				t.Fatalf("%s: reader version not reported: %q", tc.name, err)
			}
			continue
		}
		if err != nil {
			t.Fatalf("%s: unexpected error %v", tc.name, err)
		}
	}
}

// TestVersionRangeInvalid rejects malformed ranges.
func TestVersionRangeInvalid(t *testing.T) {
	t.Parallel()
	for _, pair := range [][2]uint16{{0, 1}, {1, 0}, {3, 2}, {0, 0}} {
		if _, err := NewVersionRange(pair[0], pair[1]); !errors.Is(err, ErrInvalidVersionRange) {
			t.Fatalf("range %v accepted: %v", pair, err)
		}
	}
}

// TestOldArchiveNewDecoder proves a v1 archive loads on the current reader.
func TestOldArchiveNewDecoder(t *testing.T) {
	t.Parallel()
	plain := mustUnit(t)
	plain.Annotations = map[string]string{}
	legacy := mustEncode(t, NewEncoderAt(mustVersion(t, 1, 1)), plain)
	back := mustDecode(t, NewDecoder(), legacy)
	if err := CompareUnits(plain, back); err != nil {
		t.Fatalf("old archive mismatch: %v", err)
	}
	if sectionPayload(t, legacy, sectionExt) != nil {
		t.Fatal("v1 encoder wrote an extension section")
	}
	withAnnotations := mustEncode(t, NewEncoderAt(mustVersion(t, 1, 1)), mustUnit(t))
	downgraded := mustDecode(t, NewDecoder(), withAnnotations)
	if len(downgraded.Annotations) != 0 || downgraded.NodeCount() != plain.NodeCount() {
		t.Fatal("v1 archive must drop annotations but keep every node")
	}
}

// TestNewArchiveOldDecoderRejected proves a v2 writer is refused by a v1 reader
// and that the error names both sides of the negotiation.
func TestNewArchiveOldDecoderRejected(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoderAt(mustVersion(t, 2, 2)), mustUnit(t))
	oldReader := NewDecoderAt(mustVersion(t, 1, 1))
	unit, err := oldReader.Decode(archive)
	if err == nil || unit != nil {
		t.Fatal("v1 reader accepted a v2 archive")
	}
	if !errors.Is(err, ErrUnsupportedAIRVersion) {
		t.Fatalf("wrong sentinel: %v", err)
	}
	message := err.Error()
	for _, fragment := range []string{"writer version v2", "writer minimum v2", "reader supports v1..v1"} {
		if !strings.Contains(message, fragment) {
			t.Fatalf("error omits %q: %q", fragment, message)
		}
	}
}

// TestFutureTailSectionSkipped models v1 to v2 additive compatibility.
//
// Boundary: an unknown trailing section is skipped only after version
// negotiation succeeds. A writer that raises its version or minimum beyond the
// reader range is rejected as a whole, as TestNewArchiveOldDecoderRejected
// demonstrates; tail skipping never overrides the version gate.
func TestFutureTailSectionSkipped(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	legacy := mustEncode(t, NewEncoderAt(mustVersion(t, 1, 1)), u)
	before := mustDecode(t, NewDecoder(), legacy)
	body := archiveBody(legacy)
	tail := appendSection(nil, unknownTailID, []byte("v2-additive-fields"))
	body = append(body, tail...)
	future := reseal(body)

	after := mustDecode(t, NewDecoder(), future)
	if err := CompareUnits(before, after); err != nil {
		t.Fatalf("unknown tail section changed the unit: %v", err)
	}
	if got := sectionPayload(t, future, unknownTailID); string(got) != "v2-additive-fields" {
		t.Fatalf("tail payload lost: %q", got)
	}
	legacyUnit, err := NewDecoderAt(mustVersion(t, 1, 1)).Decode(future)
	if err != nil {
		t.Fatalf("v1 reader failed to skip an unknown tail: %v", err)
	}
	if err := CompareUnits(before, legacyUnit); err != nil {
		t.Fatalf("v1 reader changed the unit: %v", err)
	}
}

// TestV2ExtensionRoundTrip proves the version 2 extension carries annotations.
func TestV2ExtensionRoundTrip(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	archive := mustEncode(t, NewEncoder(), u)
	if sectionPayload(t, archive, sectionExt) == nil {
		t.Fatal("v2 archive lacks the extension section")
	}
	back := mustDecode(t, NewDecoder(), archive)
	if err := CompareUnits(u, back); err != nil {
		t.Fatalf("v2 round trip: %v", err)
	}
	if len(back.AnnotationKeys()) != 2 || back.Annotations["generator"] != "abate-air-test" {
		t.Fatalf("annotations lost: %v", back.Annotations)
	}
}

// TestSectionOrderIndependence proves section framing, not position, drives
// decoding, so additive sections can be appended without reordering risk.
func TestSectionOrderIndependence(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	archive := mustEncode(t, NewEncoder(), u)
	shuffled := reorderSections(t, archive, []byte{
		sectionExt, sectionSem, sectionStrs, sectionNodes, sectionUnit,
	})
	back := mustDecode(t, NewDecoder(), shuffled)
	if err := CompareUnits(u, back); err != nil {
		t.Fatalf("reordered sections changed the unit: %v", err)
	}
}

// TestArchiveMagicAndVersionRejection reseals hand-edited headers so each
// field is tested independently of the checksum.
func TestArchiveMagicAndVersionRejection(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	badMagic := archiveBody(archive)
	badMagic[0] = 'X'
	if _, err := NewDecoder().Decode(reseal(badMagic)); !errors.Is(err, ErrBadMagic) {
		t.Fatalf("bad magic: %v", err)
	}
	badVersion := archiveBody(archive)
	binary.BigEndian.PutUint16(badVersion[4:6], 99)
	if _, err := NewDecoder().Decode(reseal(badVersion)); !errors.Is(err, ErrUnsupportedAIRVersion) {
		t.Fatalf("future version: %v", err)
	}
	badMinimum := archiveBody(archive)
	binary.BigEndian.PutUint16(badMinimum[6:8], 99)
	if _, err := NewDecoder().Decode(reseal(badMinimum)); !errors.Is(err, ErrUnsupportedAIRVersion) {
		t.Fatalf("future minimum: %v", err)
	}
	swapped := archiveBody(archive)
	binary.BigEndian.PutUint16(swapped[4:6], 1)
	binary.BigEndian.PutUint16(swapped[6:8], 2)
	if _, err := NewDecoder().Decode(reseal(swapped)); !errors.Is(err, ErrUnsupportedAIRVersion) {
		t.Fatalf("reversed version header: %v", err)
	}
}
