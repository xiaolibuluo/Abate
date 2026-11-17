/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

// Decoder validates and loads a versioned AIR archive.
//
// Loading is validation: the checksum, magic, version range, section framing,
// and every table index are checked before a Unit is returned. Unknown section
// identifiers are skipped by their declared length, which is how additive
// schema changes stay forward compatible.
type Decoder struct {
	version Version
}

// NewDecoder returns a decoder for the version range of this build.
func NewDecoder() Decoder {
	return Decoder{version: NewVersion()}
}

// NewDecoderAt returns a decoder for an explicit version range, used to
// emulate older or future readers in compatibility tests.
func NewDecoderAt(version Version) Decoder {
	return Decoder{version: version}
}

// Version returns the version range accepted by this decoder.
func (d Decoder) Version() Version {
	return d.version
}

// Load is an alias for Decode: loading always performs full validation.
func (d Decoder) Load(data []byte) (*Unit, error) {
	return d.Decode(data)
}

// Decode parses and validates an archive into a Unit.
func (d Decoder) Decode(data []byte) (*Unit, error) {
	if len(data) < headerSize+checksumSize {
		return nil, ErrTruncatedArchive
	}
	body := data[:len(data)-checksumSize]
	sum := sha256.Sum256(body)
	if !bytes.Equal(sum[:], data[len(data)-checksumSize:]) {
		return nil, ErrCorruptChecksum
	}
	if string(body[:4]) != archiveMagic {
		return nil, ErrBadMagic
	}
	writerVersion := binary.BigEndian.Uint16(body[4:6])
	writerMin := binary.BigEndian.Uint16(body[6:8])
	if err := d.version.Accepts(writerMin, writerVersion); err != nil {
		return nil, err
	}

	sections, err := readSections(body[headerSize:])
	if err != nil {
		return nil, err
	}
	stringsPayload, ok := sections[sectionStrs]
	if !ok {
		return nil, ErrMalformedStrings
	}
	strings, err := decodeStrings(stringsPayload)
	if err != nil {
		return nil, err
	}
	nodesPayload, ok := sections[sectionNodes]
	if !ok {
		return nil, ErrMalformedNode
	}
	nodes, err := decodeNodes(nodesPayload, strings)
	if err != nil {
		return nil, err
	}
	unitPayload, ok := sections[sectionUnit]
	if !ok {
		return nil, ErrMalformedUnit
	}
	meta, err := decodeUnit(unitPayload, strings, nodes)
	if err != nil {
		return nil, err
	}
	if semPayload, present := sections[sectionSem]; present {
		if err := decodeSemantics(semPayload, strings, nodes); err != nil {
			return nil, err
		}
	}
	annotations := make(map[string]string)
	if extPayload, present := sections[sectionExt]; present {
		annotations, err = decodeExtension(extPayload, strings)
		if err != nil {
			return nil, err
		}
	}
	unit := &Unit{
		Package:     meta.pkg,
		Files:       meta.files,
		Roots:       meta.roots,
		Annotations: annotations,
	}
	if err := unit.Validate(); err != nil {
		return nil, err
	}
	return unit, nil
}

// readSections splits a framed section stream, rejecting duplicate identifiers.
func readSections(data []byte) (map[byte][]byte, error) {
	sections := make(map[byte][]byte)
	cur := &cursor{data: data}
	for cur.remaining() > 0 {
		id := cur.u8()
		size := cur.u32()
		payload := cur.take(int(size))
		if cur.err != nil {
			return nil, cur.err
		}
		if _, exists := sections[id]; exists {
			return nil, ErrDuplicateSection
		}
		sections[id] = payload
	}
	if cur.err != nil {
		return nil, cur.err
	}
	return sections, nil
}
