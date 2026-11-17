/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"crypto/sha256"
	"errors"
	"testing"
)

// mustUnit builds a two-file unit with comments, semantics, and annotations.
func mustUnit(t *testing.T) *Unit {
	t.Helper()
	u := NewUnit(NewPackageRef("demo", []string{"os", "fmt", "fmt"}))
	name := NewNode(KindIdent, NewPosition(1, 1, 0, 4), "name")
	if !Annotate(name, SemTypeRef, "string") {
		t.Fatal("Annotate rejected a package node")
	}
	lit := NewNode(KindNumLit, NewPosition(1, 6, 5, 6), "1")
	call := NewNode(KindCall, NewPosition(1, 1, 0, 9), "+", name, lit)
	Annotate(call, SemPurity, "pure")
	root := NewNode(KindFuncDecl, NewPosition(1, 1, 0, 9), "main", call)
	if err := u.AddFile(NewFileRef("b.go", []byte("name + 1\n")), root); err != nil {
		t.Fatalf("AddFile b.go: %v", err)
	}
	doc := NewNode(KindComment, NewPosition(1, 1, 0, 3), "//x")
	root2 := NewNode(KindPackage, NewPosition(1, 1, 0, 3), "demo", doc)
	if err := u.AddFile(NewFileRef("a.go", []byte("//x\n")), root2); err != nil {
		t.Fatalf("AddFile a.go: %v", err)
	}
	u.Annotations["generator"] = "abate-air-test"
	u.Annotations["alpha"] = "1"
	return u
}

// mustEncode encodes a unit or fails the test.
func mustEncode(t *testing.T, encoder Encoder, u *Unit) []byte {
	t.Helper()
	out, err := encoder.Encode(u)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	return out
}

// mustDecode decodes an archive or fails the test.
func mustDecode(t *testing.T, decoder Decoder, data []byte) *Unit {
	t.Helper()
	u, err := decoder.Decode(data)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	return u
}

// reseal recomputes the trailing checksum of a hand-edited archive body.
func reseal(body []byte) []byte {
	sum := sha256.Sum256(body)
	out := make([]byte, 0, len(body)+checksumSize)
	out = append(out, body...)
	return append(out, sum[:]...)
}

// archiveBody strips the trailing checksum from a valid archive.
func archiveBody(archive []byte) []byte {
	return append([]byte(nil), archive[:len(archive)-checksumSize]...)
}

// sectionPayload returns the first section payload with the given identifier.
func sectionPayload(t *testing.T, archive []byte, id byte) []byte {
	t.Helper()
	body := archive[:len(archive)-checksumSize]
	cur := &cursor{data: body[headerSize:]}
	for cur.remaining() > 0 {
		sid := cur.u8()
		size := cur.u32()
		payload := cur.take(int(size))
		if cur.err != nil {
			t.Fatalf("section framing: %v", cur.err)
		}
		if sid == id {
			return payload
		}
	}
	return nil
}

// sectionOffset returns the absolute start offset of a section payload.
func sectionOffset(t *testing.T, archive []byte, id byte) int {
	t.Helper()
	body := archive[:len(archive)-checksumSize]
	cur := &cursor{data: body[headerSize:]}
	for cur.remaining() > 0 {
		sid := cur.u8()
		size := cur.u32()
		_ = cur.take(int(size))
		if cur.err != nil {
			t.Fatalf("framing: %v", cur.err)
		}
		if sid == id {
			return headerSize + cur.off - int(size)
		}
	}
	return -1
}

// mutateByte replaces one body byte and reseals the archive checksum.
func mutateByte(t *testing.T, archive []byte, at int, value byte) []byte {
	t.Helper()
	body := archiveBody(archive)
	if at < 0 || at >= len(body) {
		t.Fatalf("mutation offset %d out of range", at)
	}
	body[at] = value
	return reseal(body)
}

// archiveSentinels is the complete set of classified AIR failures.
var archiveSentinels = []error{
	ErrUnsupportedAIRVersion, ErrInvalidVersionRange, ErrNilUnit,
	ErrDuplicatePath, ErrBadMagic, ErrTruncatedArchive, ErrCorruptChecksum,
	ErrDuplicateSection, ErrMalformedNode, ErrMalformedStrings,
	ErrMalformedSemantics, ErrMalformedUnit, ErrIndexOutOfRange,
	ErrTreeTooDeep, ErrBadPosition, ErrInvalidEdit, ErrUnitMismatch,
}

// isAIRSentinel reports whether an error wraps one of the documented sentinels.
func isAIRSentinel(err error) bool {
	for _, sentinel := range archiveSentinels {
		if errors.Is(err, sentinel) {
			return true
		}
	}
	return false
}

// errorText renders an error for deterministic comparison.
func errorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// digestOfUnit returns the stable fingerprint of a decoded unit.
func digestOfUnit(u *Unit) [32]byte {
	if u == nil {
		return [32]byte{}
	}
	return u.Fingerprint()
}
