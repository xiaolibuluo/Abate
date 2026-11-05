// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package reporter 实现 CES 第 6 章结构化报告生成器。
//
// Package reporter implements the CES Chapter 6 structured reporter.
package reporter

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/qingdujian/abate/core"
)

// Report 是 CES 第 6 章定义的结构化报告顶层 Schema。
type Report struct {
	SchemaVersion string     `json:"schemaVersion"`
	Experiment     core.Experiment `json:"experiment"`
	Score          core.Score `json:"score"`
	ContentHash    string     `json:"contentHash"`
}

// marshalIndent 可在测试中替换以触发序列化错误分支。
var marshalIndent = json.MarshalIndent

// Reporter 是文件系统报告生成器，实现 core.Reporter 接口。
type Reporter struct {
	cap  core.Capability
	dir  string
}

// New 构造指定输出目录的报告生成器。
func New(dir string) (*Reporter, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("reporter: mkdir %s: %w", dir, err)
	}
	return &Reporter{
		cap: core.Capability{Name: "file-reporter", Version: "1.0.0", Protocol: core.ProtocolVersion},
		dir: dir,
	}, nil
}

// Capability 返回能力声明。
func (r *Reporter) Capability() core.Capability { return r.cap }

// Emit 将评分结果写入 {dir}/{experimentID}.json 并返回文件路径。
func (r *Reporter) Emit(_ context.Context, s core.Score) (string, error) {
	rep := Report{
		SchemaVersion: "ces/v0.1",
		Score:          s,
	}
	data, err := marshalIndent(rep, "", "  ")
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	rep.ContentHash = hex.EncodeToString(hash[:])
	data, err = marshalIndent(rep, "", "  ")
	if err != nil {
		return "", err
	}
	path := filepath.Join(r.dir, "report.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}
	return path, nil
}
