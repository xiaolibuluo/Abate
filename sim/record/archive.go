/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"abate/sim/clock"
)

// archiveVersion is the on-disk format version understood by this build.
const archiveVersion uint16 = 1

// archiveMagic prefixes every archive so a corrupted or wrong-format file is
// rejected at load time.
var archiveMagic = [8]byte{'A', 'B', 'A', 'T', 'E', 'R', 'E', 'C'}

// archiveHeaderLen is the fixed byte length of the header preceding the frame block.
const archiveHeaderLen = 8 + 2 + 32 + 8

// Sentinel errors reported by archive load and verification.
var (
	// ErrInvalidFormat reports a structurally invalid or corrupted archive.
	ErrInvalidFormat = errors.New("record: invalid archive format")
	// ErrCorruptChain reports a broken frame hash chain.
	ErrCorruptChain = errors.New("record: frame hash chain broken")
)

// Archive is a sealed, versioned and integrity-checked recording of an experiment
// run. It is produced by the Recorder and replayed or compared by the Replayer.
type Archive struct {
	Version     uint16
	Fingerprint [32]byte
	entries     []frameEntry
}

// frameEntry is one decoded frame together with its chain hash and the hash of the
// preceding frame, forming the tamper-evident linkage between frames.
type frameEntry struct {
	frame Frame
	prev  [32]byte
	hash  [32]byte
}

// Len returns the number of frames in the archive.
func (a *Archive) Len() int { return len(a.entries) }

// Frame returns the frame at index i.
func (a *Archive) Frame(i int) Frame { return a.entries[i].frame }

// Hash returns the chain hash of the frame at index i.
func (a *Archive) Hash(i int) [32]byte { return a.entries[i].hash }

// RootHash returns the header hash that anchors the first frame of the chain.
func (a *Archive) RootHash() [32]byte { return a.headerHash() }

// headerHash derives the chain anchor from the header fields recorded on disk.
func (a *Archive) headerHash() [32]byte {
	buf := make([]byte, archiveHeaderLen)
	copy(buf[0:8], archiveMagic[:])
	binary.BigEndian.PutUint16(buf[8:10], a.Version)
	copy(buf[10:42], a.Fingerprint[:])
	binary.BigEndian.PutUint64(buf[42:50], uint64(len(a.entries))) //nolint:gosec
	return sha256.Sum256(buf)
}

// newArchive assembles the chain of frame entries from an ordered frame list.
func newArchive(fp [32]byte, frames []Frame) (*Archive, error) {
	a := &Archive{Version: archiveVersion, Fingerprint: fp, entries: make([]frameEntry, len(frames))}
	prev := a.headerHash()
	var buf []byte
	for i := range frames {
		body := encodeFrameBody(frames[i])
		buf = buf[:0]
		buf = append(buf, prev[:]...)
		buf = append(buf, body...)
		h := sha256.Sum256(buf)
		a.entries[i] = frameEntry{frame: frames[i], prev: prev, hash: h}
		prev = h
	}
	return a, nil
}

// Save writes the archive in the versioned binary format to w.
func (a *Archive) Save(w io.Writer) error {
	var out bytes.Buffer
	_, _ = out.Write(archiveMagic[:])
	writeUint16(&out, a.Version)
	_, _ = out.Write(a.Fingerprint[:])
	writeUint64(&out, uint64(len(a.entries))) //nolint:gosec // frame count is bounded by the experiment size.
	offsets := make([]uint64, len(a.entries))
	for i := range a.entries {
		offsets[i] = uint64(out.Len()) //nolint:gosec // buffer length is a non-negative offset.
		e := &a.entries[i]
		_, _ = out.Write(e.prev[:])
		_, _ = out.Write(encodeFrameBody(e.frame))
		_, _ = out.Write(e.hash[:])
	}
	for i := range a.entries {
		writeUint64(&out, a.entries[i].frame.Seq)
		writeUint64(&out, offsets[i])
	}
	sum := sha256.Sum256(out.Bytes())
	_, _ = out.Write(sum[:])
	if _, err := w.Write(out.Bytes()); err != nil {
		return fmt.Errorf("record: write archive: %w", err)
	}
	return nil
}

// Load reads a versioned archive from r, verifying its integrity on the way in.
func Load(r io.Reader) (*Archive, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("record: read archive: %w", err)
	}
	if len(data) < archiveHeaderLen+sha256.Size {
		return nil, ErrInvalidFormat
	}
	if !bytes.Equal(data[0:8], archiveMagic[:]) {
		return nil, ErrInvalidFormat
	}
	version := binary.BigEndian.Uint16(data[8:10])
	if version != archiveVersion {
		return nil, ErrInvalidFormat
	}
	var fp [32]byte
	copy(fp[:], data[10:42])
	count := binary.BigEndian.Uint64(data[42:50])
	bodyEnd := len(data) - sha256.Size
	sum := sha256.Sum256(data[:bodyEnd])
	if !bytes.Equal(sum[:], data[bodyEnd:]) {
		return nil, ErrInvalidFormat
	}
	pos := archiveHeaderLen
	entries := make([]frameEntry, 0, count)
	for i := uint64(0); i < count; i++ { //nolint:gosec
		var prev [32]byte
		if pos+32 > bodyEnd {
			return nil, ErrInvalidFormat
		}
		copy(prev[:], data[pos:pos+32])
		pos += 32
		if pos+25 > bodyEnd { //nolint:gosec // minimum frame body length.
			return nil, ErrInvalidFormat
		}
		seq := binary.BigEndian.Uint64(data[pos : pos+8])
		pos += 8
		at := clock.Instant(int64(binary.BigEndian.Uint64(data[pos : pos+8]))) //nolint:gosec
		pos += 8
		tag := data[pos]
		pos++
		kind, ok := FrameKindFromTag(tag)
		if !ok {
			return nil, ErrInvalidFormat
		}
		plen := binary.BigEndian.Uint64(data[pos : pos+8])
		pos += 8
		// Bound plen against the remaining span instead of adding an untrusted
		// uint64 to pos: that addition can wrap and bypass the length check.
		if plen > uint64(bodyEnd-pos) { //nolint:gosec // the preceding bounds check keeps the remainder non-negative.
			return nil, ErrInvalidFormat
		}
		payload := data[pos : pos+int(plen)] //nolint:gosec // plen is bounded by the remaining span.
		pos += int(plen)                     //nolint:gosec // plen is bounded by the remaining span.
		if pos+32 > bodyEnd {
			return nil, ErrInvalidFormat
		}
		var h [32]byte
		copy(h[:], data[pos:pos+32])
		pos += 32
		entries = append(entries, frameEntry{
			frame: Frame{Seq: seq, At: at, Kind: kind, Payload: payload},
			prev:  prev,
			hash:  h,
		})
	}
	if pos+int(count)*16 != bodyEnd { //nolint:gosec
		return nil, ErrInvalidFormat
	}
	a := &Archive{Version: version, Fingerprint: fp, entries: entries}
	if err := a.Verify(); err != nil {
		return nil, ErrInvalidFormat
	}
	return a, nil
}

// Verify checks the chain of frame hashes and the header linkage from scratch.
func (a *Archive) Verify() error {
	prev := a.headerHash()
	for i := range a.entries {
		chain := append(append([]byte(nil), a.entries[i].prev[:]...), encodeFrameBody(a.entries[i].frame)...)
		recomputed := sha256.Sum256(chain)
		if recomputed != a.entries[i].hash {
			return ErrCorruptChain
		}
		if a.entries[i].prev != prev {
			return ErrCorruptChain
		}
		prev = a.entries[i].hash
	}
	return nil
}
