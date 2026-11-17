/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import "abate/sim/clock"

// InjectFault applies a fault to the file system or to a specific file.
//
// FaultLatency overrides the virtual nanoseconds charged per operation, while
// FaultBadBlock marks a byte range of one file as unreadable. The fault methods
// live apart from the tree operations so fs.go stays within the one-type rule.
func (f *FileSystem) InjectFault(fault Fault) error {
	switch fault.Kind {
	case FaultLatency:
		if fault.Value < 0 {
			return ErrOffset
		}
		f.ioLatency = clock.Duration(fault.Value)
		return nil
	case FaultBadBlock:
		r := f.resolve(fault.Path)
		if r.err != nil {
			return r.err
		}
		if r.node == nil {
			return ErrNoEntry
		}
		fl, ok := r.node.(*File)
		if !ok {
			return ErrNotFile
		}
		fl.markBad(fault.Offset, fault.Value)
		return nil
	default:
		return ErrUnknownFault
	}
}

// RecoverFault clears any bad-block faults recorded on the file at path.
//
// Resolving a directory is a valid no-op; a missing entry is reported as
// ErrNoEntry so callers can distinguish a recovery that did nothing.
func (f *FileSystem) RecoverFault(path string) error {
	r := f.resolve(path)
	if r.err != nil {
		return r.err
	}
	if r.node == nil {
		return ErrNoEntry
	}
	if fl, ok := r.node.(*File); ok {
		fl.clearBad()
	}
	return nil
}
