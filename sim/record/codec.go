/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/sim/clock"
	"abate/sim/rand"
)

const (
	sizeUint64 = 8
	sizeUint16 = 2
)

// putUint64 writes v as big-endian into a pre-sized slice.
func putUint64(b []byte, v uint64) { binary.BigEndian.PutUint64(b, v) }

// writeUint64 appends v as big-endian bytes to a buffer, discarding the always
// nil buffer write error.
func writeUint64(b *bytes.Buffer, v uint64) {
	var t [sizeUint64]byte
	binary.BigEndian.PutUint64(t[:], v)
	_, _ = b.Write(t[:])
}

// writeUint16 appends v as big-endian bytes to a buffer, discarding the always
// nil buffer write error.
func writeUint16(b *bytes.Buffer, v uint16) {
	var t [sizeUint16]byte
	binary.BigEndian.PutUint16(t[:], v)
	_, _ = b.Write(t[:])
}

// encodeFrameBody serializes the deterministic, fixed-layout body of a frame.
func encodeFrameBody(f Frame) []byte {
	payloadLen := uint64(len(f.Payload)) //nolint:gosec // payload length is bounded by the frame encoding.
	body := make([]byte, 2*sizeUint64+1+sizeUint64+len(f.Payload))
	pos := 0
	putUint64(body[pos:], f.Seq)
	pos += sizeUint64
	putUint64(body[pos:], uint64(f.At)) //nolint:gosec // virtual time is an int64 count and converts losslessly.
	pos += sizeUint64
	body[pos] = f.Kind.Marshal()[0]
	pos++
	putUint64(body[pos:], payloadLen)
	pos += sizeUint64
	copy(body[pos:], f.Payload)
	return body
}

// encodeEvent serializes a scheduler event into a SCHED frame payload.
func encodeEvent(e scheduler.Event) []byte {
	kind := []byte(e.Kind)
	payloadLen := uint64(len(e.Payload)) //nolint:gosec
	buf := make([]byte, 3*sizeUint64+sizeUint64+len(kind)+sizeUint64+len(e.Payload))
	pos := 0
	putUint64(buf[pos:], uint64(e.At)) //nolint:gosec
	pos += sizeUint64
	putUint64(buf[pos:], e.Seq)
	pos += sizeUint64
	putUint64(buf[pos:], uint64(len(kind))) //nolint:gosec
	pos += sizeUint64
	copy(buf[pos:], kind)
	pos += len(kind)
	putUint64(buf[pos:], payloadLen)
	pos += sizeUint64
	copy(buf[pos:], e.Payload)
	return buf
}

// encodeDraw serializes one random draw record into an RNG frame payload.
func encodeDraw(d rand.DrawRecord) []byte {
	kind := []byte(d.Kind)
	argCount := uint64(len(d.Args)) //nolint:gosec
	buf := make([]byte, sizeUint64+sizeUint64+sizeUint64+len(kind)+sizeUint64+len(d.Args)*sizeUint64+sizeUint64)
	pos := 0
	putUint64(buf[pos:], uint64(d.Index)) //nolint:gosec
	pos += sizeUint64
	putUint64(buf[pos:], uint64(len(kind))) //nolint:gosec
	pos += sizeUint64
	copy(buf[pos:], kind)
	pos += len(kind)
	putUint64(buf[pos:], argCount)
	pos += sizeUint64
	for _, a := range d.Args {
		putUint64(buf[pos:], a)
		pos += sizeUint64
	}
	putUint64(buf[pos:], d.Result)
	return buf
}

// encodeClock serializes the virtual clock instant and four-state mode.
func encodeClock(at clock.Instant, st clock.State) []byte {
	return []byte(clockStateString(at, st))
}

// clockStateString renders the clock state deterministically as text.
func clockStateString(at clock.Instant, st clock.State) string {
	return fmt.Sprintf("%d %s", int64(at), st.String()) //nolint:gosec // virtual time is an int64 count.
}

// encodeObservation serializes one observation into an OBS frame payload.
func encodeObservation(o observe.Observation) []byte {
	b, _ := o.Encode()
	return b
}
