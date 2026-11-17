/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"abate/kernel/scheduler"
	"abate/sim/record"
)

// lockedReportSHA is the byte-stable JSON report digest for ok.yaml with seed 42.
const lockedReportSHA = "c00c280272c7611da4a79de07f7223f7ab76bfe6db5e97b0c44ea899ef9fea57"

// TestRunRecordE2E proves --record seals a real, verifiable archive without
// changing a single byte of the deterministic report.
func TestRunRecordE2E(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	first := filepath.Join(dir, "a.arch")
	second := filepath.Join(dir, "b.arch")
	code, withRecord, errs := invoke("run", "testdata/ok.yaml",
		"--seed", "42", "--format", "json", "--record", first)
	if code != exitOK {
		t.Fatalf("run --record exit = %d (stderr=%s)", code, errs)
	}
	code, _, errs = invoke("run", "testdata/ok.yaml",
		"--seed", "42", "--format", "json", "--record", second)
	if code != exitOK {
		t.Fatalf("second run --record exit = %d (stderr=%s)", code, errs)
	}
	_, plain, _ := invoke("run", "testdata/ok.yaml", "--seed", "42", "--format", "json")
	if withRecord != plain || shaOf(withRecord) != lockedReportSHA {
		t.Fatalf("recording changed the report bytes: %s", shaOf(withRecord))
	}
	left, err := os.ReadFile(first) //nolint:gosec // the path was created by this test.
	if err != nil {
		t.Fatalf("read first archive: %v", err)
	}
	right, err := os.ReadFile(second) //nolint:gosec // the path was created by this test.
	if err != nil || !bytes.Equal(left, right) {
		t.Fatalf("same-seed archives differ: err=%v", err)
	}
	arch, err := record.Load(bytes.NewReader(left))
	if err != nil {
		t.Fatalf("load recorded archive: %v", err)
	}
	if err := (record.Replayer{}).Verify(arch); err != nil {
		t.Fatalf("verify recorded archive: %v", err)
	}
	kinds := make(map[record.FrameKind]int)
	for i := 0; i < arch.Len(); i++ {
		kinds[arch.Frame(i).Kind]++
	}
	for _, k := range []record.FrameKind{
		record.KindSched, record.KindObs, record.KindRNG, record.KindClk,
		record.KindFault, record.KindNet, record.KindFS, record.KindVTask,
	} {
		if kinds[k] == 0 {
			t.Fatalf("archive has no %s frames: %v", k, kinds)
		}
	}
	collector := &countingHooks{}
	if err := (record.Replayer{}).Replay(arch, collector); err != nil ||
		collector.count != arch.Len() {
		t.Fatalf("replay recorded archive: %d frames, err=%v", collector.count, err)
	}
}

// countingHooks counts replayed frames.
type countingHooks struct {
	count int
}

// OnFrame records one replayed frame.
func (h *countingHooks) OnFrame(_ record.Frame) error {
	h.count++
	return nil
}

// TestRunRecordFailurePaths covers unwritable paths, dry-run and nil archives.
func TestRunRecordFailurePaths(t *testing.T) {
	t.Parallel()
	bad := filepath.Join(t.TempDir(), "absent-dir", "x.arch")
	code, out, _ := invoke("run", "testdata/ok.yaml", "--record", bad)
	if code != exitInternalError || out != "" {
		t.Fatalf("unwritable archive exit = %d out = %q", code, out)
	}
	dry := filepath.Join(t.TempDir(), "dry.arch")
	code, _, errs := invoke("run", "testdata/ok.yaml", "--dry-run", "--record", dry)
	if code != exitOK {
		t.Fatalf("dry-run with --record exit = %d (stderr=%s)", code, errs)
	}
	if _, err := os.Stat(dry); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("dry-run wrote an archive: %v", err)
	}
	if err := writeArchive(nil, dry); err == nil {
		t.Fatal("writeArchive accepted a nil archive")
	}
}

// TestRunRecordFlagPermutation proves --record works before the positional path.
func TestRunRecordFlagPermutation(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "perm.arch")
	code, _, errs := invoke("run", "--record", path, "testdata/ok.yaml", "--seed", "7")
	if code != exitOK {
		t.Fatalf("permuted flags exit = %d (stderr=%s)", code, errs)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("permuted flags did not write the archive: %v", err)
	}
	if !slices.Contains(runValueFlags(), "record") ||
		!slices.Contains(replayValueFlags(), "format") {
		t.Fatal("value-flag tables lost an entry")
	}
	if !strings.Contains(runHelp, "Examples:") {
		t.Fatal("run help lost its copyable examples")
	}
}

// TestExperimentHandlerWiring proves the handler records without changing the
// event stream that the report digests.
func TestExperimentHandlerWiring(t *testing.T) {
	t.Parallel()
	exp := loadFixture(t, "ok.yaml")
	times := probeSchedule(exp)
	h := newExperimentHandler(exp, times, nil)
	children, err := h.handler()(scheduler.Event{
		At:      scheduler.VirtualTime(1000),
		Kind:    scheduler.EventApply,
		Payload: []byte("apply"),
	})
	if err != nil || len(children) != probeSamples+1 {
		t.Fatalf("wrapped handler produced %d children, err=%v", len(children), err)
	}
	arch, err := h.stop()
	if err != nil {
		t.Fatalf("stop handler: %v", err)
	}
	if arch.Len() < 3 {
		t.Fatalf("handler recorded only %d frames", arch.Len())
	}
	if err := (record.Replayer{}).Verify(arch); err != nil {
		t.Fatalf("handler archive is invalid: %v", err)
	}
}
