/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "errors"

// Sentinel errors of the air package. Every failure wraps exactly one of these
// values so callers can classify an archive with errors.Is.
var (
	// ErrInvalidVersionRange reports a malformed reader or writer version range.
	ErrInvalidVersionRange = errors.New("air: invalid version range")
	// ErrNilUnit reports a nil compilation unit passed to an operation.
	ErrNilUnit = errors.New("air: nil compilation unit")
	// ErrDuplicatePath reports two files registered under the same path.
	ErrDuplicatePath = errors.New("air: duplicate file path")
	// ErrBadMagic reports an archive that does not start with the ABAT magic.
	ErrBadMagic = errors.New("air: bad archive magic")
	// ErrTruncatedArchive reports an archive shorter than its declared fields.
	ErrTruncatedArchive = errors.New("air: truncated archive")
	// ErrCorruptChecksum reports a SHA-256 mismatch over the whole archive.
	ErrCorruptChecksum = errors.New("air: archive checksum mismatch")
	// ErrDuplicateSection reports the same section identifier twice.
	ErrDuplicateSection = errors.New("air: duplicate archive section")
	// ErrMalformedNode reports an invalid node table or an invalid child link.
	ErrMalformedNode = errors.New("air: malformed node table")
	// ErrMalformedStrings reports an invalid string table.
	ErrMalformedStrings = errors.New("air: malformed string table")
	// ErrMalformedSemantics reports an invalid semantic annotation table.
	ErrMalformedSemantics = errors.New("air: malformed semantic table")
	// ErrMalformedUnit reports invalid package or file metadata.
	ErrMalformedUnit = errors.New("air: malformed unit header")
	// ErrIndexOutOfRange reports an index outside its referenced table.
	ErrIndexOutOfRange = errors.New("air: index out of range")
	// ErrTreeTooDeep reports a node tree deeper than the decoder allows.
	ErrTreeTooDeep = errors.New("air: node tree too deep")
	// ErrBadPosition reports an invalid source interval on a node.
	ErrBadPosition = errors.New("air: invalid node position")
	// ErrInvalidEdit reports an edit outside the source buffer or overlapping another edit.
	ErrInvalidEdit = errors.New("air: invalid source edit")
	// ErrUnitMismatch reports two units that differ after a round trip.
	ErrUnitMismatch = errors.New("air: unit comparison mismatch")
)
