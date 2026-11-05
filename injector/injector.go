// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package injector 实现 CES 第 7 章 Injector 角色的文件系统适配器。
//
// Package injector implements the CES Chapter 7 Injector role
// as a file-system adapter.
package injector

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/qingdujian/abate/core"
)

// Injector 是文件系统注入器，实现 core.Injector 接口。
// 注入状态以标记文件记录，Recover 删除标记文件。
type Injector struct {
	cap  core.Capability
	mu   sync.Mutex
	dir  string
	active map[string]bool
}

// New 构造指定目录的注入器。
func New(dir string) (*Injector, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("injector: mkdir %s: %w", dir, err)
	}
	return &Injector{
		cap: core.Capability{
			Name:        "file-injector",
			Version:     "1.0.0",
			Protocol:    core.ProtocolVersion,
			FaultClasses: core.FaultClasses(),
			TargetKinds:  []core.TargetKind{core.TargetService, core.TargetInstance, core.TargetDependency},
		},
		dir:     dir,
		active:  make(map[string]bool),
	}, nil
}

// Capability 返回能力声明。
func (in *Injector) Capability() core.Capability { return in.cap }

// Inject 创建注入标记文件，返回句柄（CES-2.4.1 幂等）。
func (in *Injector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	in.mu.Lock()
	defer in.mu.Unlock()
	path := filepath.Join(in.dir, e.ID+".inj")
	if err := os.WriteFile(path, []byte(e.ID), 0644); err != nil {
		return core.Injection{}, err
	}
	in.active[e.ID] = true
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 删除注入标记文件（CES-3.1.2 可重复调用）。
func (in *Injector) Recover(_ context.Context, inj core.Injection) error {
	in.mu.Lock()
	defer in.mu.Unlock()
	path := filepath.Join(in.dir, inj.ID+".inj")
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	delete(in.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (in *Injector) IsActive() bool {
	in.mu.Lock()
	defer in.mu.Unlock()
	for _, a := range in.active {
		if a {
			return true
		}
	}
	return false
}
