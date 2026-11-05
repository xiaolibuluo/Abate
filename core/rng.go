// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

// RNG 是确定性随机数生成器：splitmix64 算法，纯整数运算，无外部依赖。
// 零抖动底座：同一 seed 在任何机器、任何时刻产生完全相同的序列，
// 变异体选择与混沌注入时序由此驱动（CES-3.1.3 确定性重放的根基）。
//
// RNG is a deterministic pseudo-random generator (splitmix64): pure integer
// arithmetic, no external dependencies. The same seed yields the same
// sequence on any machine at any time, driving mutant selection and chaos
// injection timing (the foundation of CES-3.1.3 deterministic replay).
type RNG struct {
	state uint64
}

// NewRNG 以 seed 构造生成器；seed 为 0 时归一为固定非零种子，
// 避免 splitmix64 的零状态退化。
// NewRNG constructs a generator from seed; a zero seed is normalized to a
// fixed non-zero value to avoid degenerate zero state.
func NewRNG(seed int64) RNG {
	s := uint64(seed)
	if s == 0 {
		s = 0x9E3779B97F4A7C15
	}
	return RNG{state: s}
}

// Uint64 产出序列中的下一个 uint64。
// Uint64 yields the next uint64 in the sequence.
func (r *RNG) Uint64() uint64 {
	r.state += 0x9E3779B97F4A7C15
	z := r.state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

// Intn 产出 [0, n) 区间的确定性整数；n <= 0 时返回 0。
// Intn yields a deterministic integer in [0, n); non-positive n yields 0.
func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.Uint64() % uint64(n))
}

// Shuffle 对切片做确定性 Fisher-Yates 洗牌（原地，返回同一切片）。
// 等长输入 + 同 seed ⇒ 输出排列完全一致。
// Shuffle performs a deterministic in-place Fisher-Yates shuffle; equal
// inputs with the same seed yield identical permutations.
func (r *RNG) Shuffle(n int, swap func(i, j int)) {
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		swap(i, j)
	}
}

// Perm 返回 [0, n) 的确定性随机排列。
// Perm returns a deterministic random permutation of [0, n).
func (r *RNG) Perm(n int) []int {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	r.Shuffle(n, func(i, j int) { p[i], p[j] = p[j], p[i] })
	return p
}
