// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package injector

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/qingdujian/abate/core"
)

var defaultTruncate = func(f *os.File, size int64) error { return f.Truncate(size) }

// DiskInjector 是磁盘注入器，通过创建大文件实现磁盘空间耗尽（F4）。
// DiskInjector performs disk exhaustion by creating large files (F4).
type DiskInjector struct {
	cap        core.Capability
	mu         sync.Mutex
	active     map[string]string
	dir        string
	truncateFn func(*os.File, int64) error
}

// NewDiskInjector 构造磁盘注入器，dir 为临时文件目录。
func NewDiskInjector(dir string) (*DiskInjector, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("disk-injector: mkdir %s: %w", dir, err)
	}
	return &DiskInjector{
		cap: core.Capability{
			Name:         "disk-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F4Exhaustion},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService},
		},
		active:     make(map[string]string),
		dir:        dir,
		truncateFn: defaultTruncate,
	}, nil
}

// Capability 返回能力声明。
func (d *DiskInjector) Capability() core.Capability { return d.cap }

// Inject 创建填充文件实现磁盘耗尽（CES-2.4.1 幂等）。
func (d *DiskInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	sizeStr := e.Fault.Params["size"]
	if sizeStr == "" {
		return core.Injection{}, fmt.Errorf("disk-injector: size parameter is required")
	}

	path := filepath.Join(d.dir, e.ID+".fill")
	f, err := os.Create(path)
	if err != nil {
		return core.Injection{}, fmt.Errorf("disk-injector: create %s: %w", path, err)
	}
	if err := d.truncateFn(f, parseSize(sizeStr)); err != nil {
		f.Close()
		os.Remove(path)
		return core.Injection{}, fmt.Errorf("disk-injector: truncate: %w", err)
	}
	f.Close()

	d.active[e.ID] = path
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 删除填充文件（CES-3.1.2 可重复调用）。
func (d *DiskInjector) Recover(_ context.Context, inj core.Injection) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	path, ok := d.active[inj.ID]
	if !ok {
		return nil
	}

	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("disk-injector: remove %s: %w", path, err)
	}

	delete(d.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (d *DiskInjector) IsActive() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.active) > 0
}

// parseSize 将字符串大小（如 100M、1G）解析为字节数。
func parseSize(s string) int64 {
	if len(s) == 0 {
		return 0
	}
	unit := s[len(s)-1]
	num := int64(0)
	for _, c := range s[:len(s)-1] {
		if c < '0' || c > '9' {
			return 0
		}
		num = num*10 + int64(c-'0')
	}
	switch unit {
	case 'K', 'k':
		return num * 1024
	case 'M', 'm':
		return num * 1024 * 1024
	case 'G', 'g':
		return num * 1024 * 1024 * 1024
	default:
		for _, c := range s {
			if c < '0' || c > '9' {
				return 0
			}
		}
		var n int64
		for _, c := range s {
			n = n*10 + int64(c-'0')
		}
		return n
	}
}