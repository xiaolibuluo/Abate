// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package hub 实现 ChaosHub 子编码注册表、远程同步与版本公告机制。
// Package hub implements the ChaosHub sub-code registry, remote sync and announcement.
package hub

import (
	"fmt"
	"sort"
	"sync"
)

// SubCode 描述一条 ChaosHub 子编码（CES 附录 B）。
type SubCode struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Params      string `json:"params"`
	Description string `json:"description,omitempty"`
}

// Registry 是子编码注册表，线程安全。
type Registry struct {
	mu       sync.RWMutex
	version  string
	codes    map[string]SubCode
}

// NewRegistry 构造空注册表。
func NewRegistry(version string) *Registry {
	return &Registry{version: version, codes: make(map[string]SubCode)}
}

// NewSeedRegistry 返回附录 B v0.1 种子表（20 项）。
func NewSeedRegistry() *Registry {
	r := NewRegistry("v0.1")
	for _, sc := range seedSubCodes() {
		r.codes[sc.Code] = sc
	}
	return r
}

// Register 注册或更新一条子编码。
func (r *Registry) Register(sc SubCode) error {
	if sc.Code == "" {
		return fmt.Errorf("hub: sub-code must not be empty")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.codes[sc.Code] = sc
	return nil
}

// Lookup 查找子编码；未找到返回错误。
func (r *Registry) Lookup(code string) (SubCode, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sc, ok := r.codes[code]
	if !ok {
		return SubCode{}, fmt.Errorf("hub: sub-code %q not found", code)
	}
	return sc, nil
}

// AllCodes 返回全部子编码（字典序）。
func (r *Registry) AllCodes() []SubCode {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]SubCode, 0, len(r.codes))
	for _, sc := range r.codes {
		result = append(result, sc)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Code < result[j].Code
	})
	return result
}

// AllCodeStrings 返回全部子编码字符串（字典序）。
func (r *Registry) AllCodeStrings() []string {
	codes := r.AllCodes()
	out := make([]string, len(codes))
	for i, sc := range codes {
		out[i] = sc.Code
	}
	return out
}

// Version 返回当前注册表版本。
func (r *Registry) Version() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.version
}

// SetVersion 更新注册表版本。
func (r *Registry) SetVersion(v string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.version = v
}

// Count 返回子编码数量。
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.codes)
}

// Merge 将其他注册表的子编码合并进来，返回新增和更新的编码列表。
func (r *Registry) Merge(other *Registry) (added, updated []string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for code, sc := range other.codes {
		if old, exists := r.codes[code]; exists {
			if old != sc {
				updated = append(updated, code)
			}
		} else {
			added = append(added, code)
		}
		r.codes[code] = sc
	}
	sort.Strings(added)
	sort.Strings(updated)
	return
}

// seedSubCodes 返回附录 B v0.1 种子表 20 项。
func seedSubCodes() []SubCode {
	return []SubCode{
		{Code: "F1.1", Name: "网络延迟", Params: "delayMs"},
		{Code: "F1.2", Name: "延迟抖动", Params: "delayMs,jitterMs"},
		{Code: "F1.3", Name: "超时", Params: "timeoutMs"},
		{Code: "F2.1", Name: "进程终止", Params: "signal"},
		{Code: "F2.2", Name: "进程重启", Params: "signal,gracePeriodMs"},
		{Code: "F3.1", Name: "丢包", Params: "loss,direction"},
		{Code: "F3.2", Name: "连接重置", Params: "direction"},
		{Code: "F3.3", Name: "网络分区", Params: "peer,direction"},
		{Code: "F3.4", Name: "带宽限制", Params: "rateKbps"},
		{Code: "F4.1", Name: "CPU 耗尽", Params: "utilization"},
		{Code: "F4.2", Name: "内存耗尽", Params: "utilization,fillRate"},
		{Code: "F4.3", Name: "磁盘填充", Params: "utilization"},
		{Code: "F4.4", Name: "文件描述符耗尽", Params: "count"},
		{Code: "F4.5", Name: "inode 耗尽", Params: "count"},
		{Code: "F4.6", Name: "连接数耗尽", Params: "count"},
		{Code: "F5.1", Name: "位翻转", Params: "bitIndex"},
		{Code: "F5.2", Name: "数据截断", Params: "length"},
		{Code: "F5.3", Name: "校验和篡改", Params: "algorithm"},
		{Code: "F6.1", Name: "时钟偏移", Params: "offsetMs"},
		{Code: "F6.2", Name: "时钟回拨", Params: "offsetMs"},
	}
}
