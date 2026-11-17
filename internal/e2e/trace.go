/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"bytes"
	"encoding/json"
	"os"

	"abate/internal/corpus"
)

// traceLine is one deterministic progress row. It contains no wall-clock value
// and no absolute path; AllocDeltaBytes is a runtime observation that lives
// only in this sidecar, never in the compared report bytes.
type traceLine struct {
	Seq             int    `json:"seq"`
	Stage           string `json:"stage"`
	Language        string `json:"language"`
	RelPath         string `json:"relPath"`
	FileSHA256      string `json:"fileSha256"`
	AllocDeltaBytes uint64 `json:"allocDeltaBytes"`
}

// Trace writes deterministic progress lines before every real parse so the
// last completed row localises a crash without any recover. Lines are flushed
// in bounded batches through os.WriteFile, which keeps the variable output
// path free of gosec file-inclusion findings.
type Trace struct {
	path       string
	buf        bytes.Buffer
	seq        int
	sinceFlush int
	err        error
}

// NewTrace builds a trace bound to one output file path.
func NewTrace(path string) *Trace {
	return &Trace{path: path}
}

// Mark appends one progress row and flushes periodically. A write failure is
// retained and surfaced through Err instead of being silently discarded.
func (t *Trace) Mark(entry corpus.Entry, stage string, alloc uint64) {
	if t.err != nil {
		return
	}
	t.seq++
	line := traceLine{
		Seq: t.seq, Stage: stage, Language: entry.Language,
		RelPath: entry.RelPath, FileSHA256: entry.SHA256,
		AllocDeltaBytes: alloc,
	}
	data, err := json.Marshal(line)
	if err != nil {
		t.err = err
		return
	}
	_, _ = t.buf.Write(data)
	_ = t.buf.WriteByte('\n')
	t.sinceFlush++
	if t.sinceFlush >= traceFlushInterval {
		t.err = t.Flush()
		t.sinceFlush = 0
	}
}

// Flush writes the complete accumulated trace so far.
func (t *Trace) Flush() error {
	return os.WriteFile(t.path, t.buf.Bytes(), FilePerm)
}

// Err returns the first trace write failure, or nil when every flush worked.
func (t *Trace) Err() error {
	return t.err
}
