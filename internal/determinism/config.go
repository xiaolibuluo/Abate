/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

// Config is the deterministic configuration of a guest execution instance.
//
// Config is pure data: it carries no file descriptor, no handle and no process
// reference, so it can be compared, hashed and replayed byte for byte. The
// configuration model mirrors plugin/sandbox: a guest is described entirely by
// data, never by capability-bearing handles.
type Config struct {
	// FuelLimit is the initial fuel budget charged for every top-level call.
	// A value of zero disables the meter and allows unbounded execution.
	FuelLimit uint64

	// MemoryMinPages is the hint for the smallest linear memory in 64KiB pages.
	MemoryMinPages uint32

	// MemoryMaxPages caps linear memory growth in 64KiB pages. It must be
	// greater than zero; loading rejects guests whose memory exceeds it.
	MemoryMaxPages uint32

	// ClockEnabled exposes the deterministic virtual clock host functions.
	ClockEnabled bool

	// ClockEpoch is the virtual time origin in nanoseconds for the clock.
	ClockEpoch uint64
}

// DefaultConfig returns the deterministic default configuration.
//
// It enables the fuel meter with a generous default budget, sizes the linear
// memory to one page and enables the virtual clock anchored at the zero epoch.
func DefaultConfig() Config {
	return Config{
		FuelLimit:      1 << 20,
		MemoryMinPages: 1,
		MemoryMaxPages: 256,
		ClockEnabled:   true,
		ClockEpoch:     0,
	}
}

// Validate reports whether the configuration is internally consistent.
func (c Config) Validate() error {
	if c.MemoryMaxPages == 0 {
		return errConfigMemoryMax
	}
	if c.MemoryMinPages > c.MemoryMaxPages {
		return errConfigMemoryMin
	}
	return nil
}
