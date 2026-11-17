/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// ErrNilArchive reports a comparison against a missing archive.
var ErrNilArchive = errors.New("record: nil archive")

// Compare recursively compares two archives and returns a structured Diff.
//
// The comparison follows the deep-recursive model of google/go-cmp but is a
// standard-library rewrite for the frame domain: reflection and its
// traversal-order dependence are removed, every field is visited in a fixed
// order, and the earliest divergence is reported with its complete path.
// Header fields, frame counts, sequence numbers, virtual instants, kinds,
// payload bytes and chain hashes are all compared. Both archives are verified
// first, so a corrupted chain is reported as an error rather than a Diff.
func Compare(a, b *Archive) (Diff, error) {
	if a == nil || b == nil {
		return Diff{}, ErrNilArchive
	}
	rp := Replayer{}
	if err := rp.Verify(a); err != nil {
		return Diff{}, fmt.Errorf("record: compare left archive: %w", err)
	}
	if err := rp.Verify(b); err != nil {
		return Diff{}, fmt.Errorf("record: compare right archive: %w", err)
	}
	w := &diffWalker{diff: Diff{
		Equal: true,
		Summary: DiffSummary{
			LeftFrames:          a.Len(),
			RightFrames:         b.Len(),
			FirstDivergentFrame: -1,
			KindCounts:          make(map[FrameKind]int),
		},
	}}
	w.compareHeader(a, b)
	n := a.Len()
	if b.Len() < n {
		n = b.Len()
	}
	for i := 0; i < n; i++ {
		w.compareFrame(i, a.Frame(i), b.Frame(i), a.Hash(i), b.Hash(i))
	}
	w.compareLengths(a.Len(), b.Len())
	w.compareAnchor(a, b)
	w.diff.Equal = w.diff.Summary.FieldsDifferent == 0
	return w.diff, nil
}

// diffWalker accumulates one structured comparison without shared state.
type diffWalker struct {
	diff     Diff
	firstSet bool
}

// compareHeader visits the fixed header fields of both archives.
func (w *diffWalker) compareHeader(a, b *Archive) {
	wantV := strconv.FormatUint(uint64(a.Version), 10)
	gotV := strconv.FormatUint(uint64(b.Version), 10)
	if a.Version == b.Version {
		w.reportMatch()
	} else {
		w.reportDifferent("Header.Version", -1, Frame{}, "Version", wantV, gotV, -1)
	}
	wantF := hex.EncodeToString(a.Fingerprint[:])
	gotF := hex.EncodeToString(b.Fingerprint[:])
	if a.Fingerprint == b.Fingerprint {
		w.reportMatch()
	} else {
		w.reportDifferent("Header.Fingerprint", -1, Frame{}, "Fingerprint", wantF, gotF, -1)
	}
}

// compareAnchor visits the chain anchor last, so a frame-count difference is
// reported at the frame level instead of being masked by a derived header hash.
func (w *diffWalker) compareAnchor(a, b *Archive) {
	want, got := a.RootHash(), b.RootHash()
	wantH := hex.EncodeToString(want[:])
	gotH := hex.EncodeToString(got[:])
	if want == got {
		w.reportMatch()
	} else {
		w.reportDifferent("Header.RootHash", -1, Frame{}, "RootHash", wantH, gotH, -1)
	}
}

// compareFrame visits every field of one frame pair in a fixed order.
func (w *diffWalker) compareFrame(index int, want, got Frame, wantHash, gotHash [32]byte) {
	w.diff.Summary.FramesCompared++
	w.diff.Summary.KindCounts[want.Kind]++
	base := fmt.Sprintf("Frames[%d]", index)
	wantSeq := strconv.FormatUint(want.Seq, 10)
	gotSeq := strconv.FormatUint(got.Seq, 10)
	if want.Seq == got.Seq {
		w.reportMatch()
	} else {
		w.reportDifferent(base+".Seq", index, want, "Seq", wantSeq, gotSeq, -1)
	}
	wantAt := strconv.FormatInt(int64(want.At), 10)
	gotAt := strconv.FormatInt(int64(got.At), 10)
	if want.At == got.At {
		w.reportMatch()
	} else {
		w.reportDifferent(base+".At", index, want, "At", wantAt, gotAt, -1)
	}
	wantKind := string(want.Kind)
	gotKind := string(got.Kind)
	if want.Kind == got.Kind {
		w.reportMatch()
	} else {
		w.reportDifferent(base+".Kind", index, want, "Kind", wantKind, gotKind, -1)
	}
	w.comparePayload(base, index, want, got)
	wantH := hex.EncodeToString(wantHash[:])
	gotH := hex.EncodeToString(gotHash[:])
	if wantHash == gotHash {
		w.reportMatch()
	} else {
		w.reportDifferent(base+".Hash", index, want, "Hash", wantH, gotH, -1)
	}
}

// comparePayload visits the payload length and each payload byte in order.
func (w *diffWalker) comparePayload(base string, index int, want, got Frame) {
	wantLen, gotLen := len(want.Payload), len(got.Payload)
	if wantLen == gotLen {
		w.reportMatch()
	} else {
		w.reportDifferent(base+".Payload.len", index, want, "Payload.len",
			strconv.Itoa(wantLen), strconv.Itoa(gotLen), -1)
	}
	n := wantLen
	if gotLen < n {
		n = gotLen
	}
	for j := 0; j < n; j++ {
		wantB := strconv.Itoa(int(want.Payload[j]))
		gotB := strconv.Itoa(int(got.Payload[j]))
		if want.Payload[j] == got.Payload[j] {
			w.reportMatch()
		} else {
			w.reportDifferent(fmt.Sprintf("%s.Payload[%d]", base, j), index, want, "Payload",
				wantB, gotB, j)
		}
	}
}

// compareLengths reports a missing or extra tail frame when counts disagree.
func (w *diffWalker) compareLengths(left, right int) {
	if left == right {
		return
	}
	field := "FrameCount.extra"
	path := fmt.Sprintf("Frames[%d]", right)
	index := right
	if right > left {
		field = "FrameCount.missing"
		path = fmt.Sprintf("Frames[%d]", left)
		index = left
	}
	w.reportDifferent(path, index, Frame{}, field,
		strconv.Itoa(left), strconv.Itoa(right), -1)
}

// reportDifferent records one diverging field comparison and captures the
// earliest divergence for diagnostics.
func (w *diffWalker) reportDifferent(path string, frameIndex int, f Frame,
	field, want, got string, byteOffset int,
) {
	w.diff.Summary.FieldsCompared++
	w.diff.Summary.FieldsDifferent++
	if w.firstSet {
		return
	}
	w.firstSet = true
	w.diff.Summary.FirstDivergentFrame = frameIndex
	w.diff.First = &DiffPoint{
		Path:       path,
		FrameIndex: frameIndex,
		Seq:        f.Seq,
		Kind:       f.Kind,
		Field:      field,
		Want:       want,
		Got:        got,
		ByteOffset: byteOffset,
	}
}

// reportMatch records one field comparison that held equal.
func (w *diffWalker) reportMatch() {
	w.diff.Summary.FieldsCompared++
	w.diff.Summary.FieldsEqual++
}
