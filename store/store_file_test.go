// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileStoreSaveAndGetRecord(t *testing.T) {
	dir := t.TempDir()
	f, err := NewFileStore(dir)
	if err != nil {
		t.Fatalf("new: %v", err)
	}
	r := testRecord("exp-b")
	if err := f.SaveRecord(r); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, ok, err := f.GetRecord("exp-b")
	if err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if got.Exp.ID != "exp-b" {
		t.Errorf("id = %q", got.Exp.ID)
	}
}

func TestFileStoreGetRecordNotFound(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_, ok, err := f.GetRecord("ghost")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ok {
		t.Error("should not be found")
	}
}

func TestFileStoreListRecords(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = f.SaveRecord(testRecord("zeta"))
	_ = f.SaveRecord(testRecord("alpha"))
	list, err := f.ListRecords()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("len = %d", len(list))
	}
	if list[0].Exp.ID != "alpha" || list[1].Exp.ID != "zeta" {
		t.Errorf("order: %s, %s", list[0].Exp.ID, list[1].Exp.ID)
	}
}

func TestFileStoreListRecordsEmpty(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	list, err := f.ListRecords()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if list != nil {
		t.Errorf("expected nil, got %d", len(list))
	}
}

func TestFileStoreSaveAndGetScore(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	s := testScore()
	if err := f.SaveScore(s); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, ok, err := f.GetScore()
	if err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if got.Final != 8900 {
		t.Errorf("final = %d", got.Final)
	}
}

func TestFileStoreGetScoreNotFound(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_, ok, err := f.GetScore()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ok {
		t.Error("should not be found")
	}
}

func TestFileStoreNewMkdirError(t *testing.T) {
	_, err := NewFileStore("/proc/cannot/create")
	if err == nil {
		t.Error("expected mkdir error")
	}
}

func TestFileStoreListRecordsSkipsNonJSON(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = f.SaveRecord(testRecord("real"))
	_ = os.WriteFile(filepath.Join(dir, "records", "not-json.txt"), []byte("x"), 0644)
	list, err := f.ListRecords()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("len = %d, want 1", len(list))
	}
}

func TestFileStoreSaveRecordMarshalError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	orig := marshalIndent
	marshalIndent = func(_ interface{}, _, _ string) ([]byte, error) {
		return nil, errTestStore
	}
	defer func() { marshalIndent = orig }()
	if err := f.SaveRecord(testRecord("fail")); err == nil {
		t.Error("expected marshal error")
	}
}

func TestFileStoreGetRecordUnmarshalError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.MkdirAll(filepath.Join(dir, "records"), 0755)
	_ = os.WriteFile(filepath.Join(dir, "records", "bad.json"), []byte("not json"), 0644)
	if _, _, err := f.GetRecord("bad"); err == nil {
		t.Error("expected unmarshal error")
	}
}

func TestFileStoreListRecordsUnmarshalError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.MkdirAll(filepath.Join(dir, "records"), 0755)
	_ = os.WriteFile(filepath.Join(dir, "records", "bad.json"), []byte("not json"), 0644)
	if _, err := f.ListRecords(); err == nil {
		t.Error("expected unmarshal error")
	}
}

func TestFileStoreSaveScoreMarshalError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	orig := marshalIndent
	marshalIndent = func(_ interface{}, _, _ string) ([]byte, error) {
		return nil, errTestStore
	}
	defer func() { marshalIndent = orig }()
	if err := f.SaveScore(testScore()); err == nil {
		t.Error("expected marshal error")
	}
}

func TestFileStoreGetScoreUnmarshalError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.WriteFile(filepath.Join(dir, "score.json"), []byte("not json"), 0644)
	if _, _, err := f.GetScore(); err == nil {
		t.Error("expected unmarshal error")
	}
}

func TestFileStoreSaveRecordMkdirError(t *testing.T) {
	f := &FileStore{dir: "/proc/cannot"}
	if err := f.SaveRecord(testRecord("fail")); err == nil {
		t.Error("expected mkdir error")
	}
}

func TestFileStoreGetRecordReadError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.MkdirAll(filepath.Join(dir, "records", "dir.json"), 0755)
	if _, _, err := f.GetRecord("dir"); err == nil {
		t.Error("expected read error")
	}
}

func TestFileStoreListRecordsReadDirError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.MkdirAll(filepath.Join(dir, "records"), 0755)
	_ = os.WriteFile(filepath.Join(dir, "records", "a.json"), []byte("{}"), 0644)
	_ = os.RemoveAll(filepath.Join(dir, "records"))
	_ = os.WriteFile(filepath.Join(dir, "records"), []byte("x"), 0644)
	if _, err := f.ListRecords(); err == nil {
		t.Error("expected readdir error")
	}
}

func TestFileStoreListRecordsReadFileError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	rdir := filepath.Join(dir, "records")
	_ = os.MkdirAll(rdir, 0755)
	_ = os.Symlink(dir, filepath.Join(rdir, "link.json"))
	if _, err := f.ListRecords(); err == nil {
		t.Error("expected readfile error")
	}
}

func TestFileStoreGetScoreReadError(t *testing.T) {
	dir := t.TempDir()
	f, _ := NewFileStore(dir)
	_ = os.MkdirAll(filepath.Join(dir, "score.json"), 0755)
	if _, _, err := f.GetScore(); err == nil {
		t.Error("expected read error")
	}
}
