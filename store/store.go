// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package store 实现实验记录与评分结果的持久化存储（CES 第 6 章）。
//
// Package store implements persistence for experiment records and
// scoring results per CES Chapter 6.
package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/qingdujian/abate/core"
)

// Store 是实验记录与评分的存储接口。
type Store interface {
	SaveRecord(r core.Record) error
	GetRecord(id string) (core.Record, bool, error)
	ListRecords() ([]core.Record, error)
	SaveScore(s core.Score) error
	GetScore() (core.Score, bool, error)
}

// MemoryStore 是线程安全的内存存储。
type MemoryStore struct {
	mu      sync.RWMutex
	records map[string]core.Record
	score   *core.Score
}

// NewMemoryStore 构造空内存存储。
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{records: make(map[string]core.Record)}
}

// SaveRecord 保存实验记录。
func (m *MemoryStore) SaveRecord(r core.Record) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.records[r.Exp.ID] = r
	return nil
}

// GetRecord 按 ID 查找实验记录。
func (m *MemoryStore) GetRecord(id string) (core.Record, bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.records[id]
	return r, ok, nil
}

// ListRecords 返回全部记录，按实验 ID 字典序排列。
func (m *MemoryStore) ListRecords() ([]core.Record, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]core.Record, 0, len(m.records))
	for _, r := range m.records {
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Exp.ID < result[j].Exp.ID
	})
	return result, nil
}

// SaveScore 保存评分结果。
func (m *MemoryStore) SaveScore(s core.Score) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.score = &s
	return nil
}

// GetScore 读取评分结果。
func (m *MemoryStore) GetScore() (core.Score, bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.score == nil {
		return core.Score{}, false, nil
	}
	return *m.score, true, nil
}

// marshalIndent 和 unmarshal 可在测试中替换以触发序列化错误分支。
var marshalIndent = json.MarshalIndent
var unmarshal = json.Unmarshal

// FileStore 是基于文件系统的持久化存储。
type FileStore struct {
	dir string
}

// NewFileStore 构造指定目录的文件存储；目录不存在时自动创建。
func NewFileStore(dir string) (*FileStore, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("store: mkdir %s: %w", dir, err)
	}
	return &FileStore{dir: dir}, nil
}

// SaveRecord 将记录写入 {dir}/records/{id}.json。
func (f *FileStore) SaveRecord(r core.Record) error {
	path := filepath.Join(f.dir, "records", r.Exp.ID+".json")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := marshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GetRecord 按 ID 从文件读取记录。
func (f *FileStore) GetRecord(id string) (core.Record, bool, error) {
	path := filepath.Join(f.dir, "records", id+".json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return core.Record{}, false, nil
	}
	if err != nil {
		return core.Record{}, false, err
	}
	var r core.Record
	if err := unmarshal(data, &r); err != nil {
		return core.Record{}, false, err
	}
	return r, true, nil
}

// ListRecords 遍历 records 目录加载全部记录。
func (f *FileStore) ListRecords() ([]core.Record, error) {
	rdir := filepath.Join(f.dir, "records")
	entries, err := os.ReadDir(rdir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var result []core.Record
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(rdir, e.Name()))
		if err != nil {
			return nil, err
		}
		var r core.Record
		if err := unmarshal(data, &r); err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Exp.ID < result[j].Exp.ID
	})
	return result, nil
}

// SaveScore 将评分写入 {dir}/score.json。
func (f *FileStore) SaveScore(s core.Score) error {
	path := filepath.Join(f.dir, "score.json")
	data, err := marshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GetScore 从文件读取评分。
func (f *FileStore) GetScore() (core.Score, bool, error) {
	path := filepath.Join(f.dir, "score.json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return core.Score{}, false, nil
	}
	if err != nil {
		return core.Score{}, false, err
	}
	var s core.Score
	if err := unmarshal(data, &s); err != nil {
		return core.Score{}, false, err
	}
	return s, true, nil
}
