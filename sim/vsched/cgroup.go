/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

// Cgroup is a virtual control group that bounds CPU weight and memory for its
// member tasks. Memory usage is the aggregate of all attached, non-terminated
// tasks.
type Cgroup struct {
	name       string
	shares     int64
	baseShares int64
	memLimit   int64
	memUsed    int64
	tasks      int
}

// SetShares updates the active CPU share weight, clamped to a minimum of one.
func (c *Cgroup) SetShares(v int64) {
	if v < 1 {
		v = 1
	}
	c.shares = v
}

// Name returns the cgroup name.
func (c *Cgroup) Name() string { return c.name }

// Shares returns the active CPU share weight.
func (c *Cgroup) Shares() int64 { return c.shares }

// BaseShares returns the original CPU share weight before any throttling.
func (c *Cgroup) BaseShares() int64 { return c.baseShares }

// MemLimit returns the configured memory limit in bytes.
func (c *Cgroup) MemLimit() int64 { return c.memLimit }

// MemUsed returns the aggregate memory charged to member tasks.
func (c *Cgroup) MemUsed() int64 { return c.memUsed }

// TaskCount returns the number of member tasks.
func (c *Cgroup) TaskCount() int { return c.tasks }
