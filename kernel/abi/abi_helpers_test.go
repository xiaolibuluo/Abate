/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi_test

import "testing"

func assertPanics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected a panic, but the call returned normally")
		}
	}()
	fn()
}

func assertNoPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	fn()
}

type chaosSource struct {
	state uint64
}

func newChaosSource(seed uint64) *chaosSource {
	return &chaosSource{state: seed}
}

func (s *chaosSource) next() uint64 {
	s.state += 0x9e3779b97f4a7c15
	z := s.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func chaosUint64Values() []uint64 {
	values := []uint64{
		0, 1, 2, 62, 63, 64, 65,
		4094, 4095, 4096, 4097, 8191, 8192, 8193,
		2097150, 2097151, 2097152, 2097153, 4194304,
		1073741823, 1073741824,
		4294967295, 4294967296,
		^uint64(0), ^uint64(0) - 1, ^uint64(0) - 62,
		^uint64(0) - 4095, ^uint64(0) - 2097151,
	}
	gen := newChaosSource(0xab17c0de5eed1234)
	for i := 0; i < 4096; i++ {
		values = append(values, gen.next())
	}
	return values
}
