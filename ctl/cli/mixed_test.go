/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/scheduler"
	"abate/kernel/verify"
	"abate/sim/record"
)

// errMixedTest is a test-only sentinel injected into the mixed ports.
var errMixedTest = errors.New("cli: injected mixed failure")

// fakePoint builds a mutation point with a nil operator, which the synthetic
// target contract only inspects by its lexical fields.
func fakePoint(original, mutated string) mutation.MutationPoint {
	return mutation.NewMutationPoint("n-arith", original, mutated, nil, 1, 0)
}

// applyEventForTest returns one scheduler apply event at a fixed instant.
func applyEventForTest() scheduler.Event {
	return scheduler.Event{
		At:      scheduler.VirtualTime(1_000_000_000),
		Kind:    scheduler.EventApply,
		Payload: []byte("apply"),
	}
}

// replayCollector captures replayed frames for a frame-by-frame comparison.
type replayCollector struct {
	frames []record.Frame
}

// OnFrame records one replayed frame.
func (c *replayCollector) OnFrame(frame record.Frame) error {
	c.frames = append(c.frames, frame)
	return nil
}

// TestMixedFramesRealSIMKinds proves mixedFrames folds four real SIM-derived
// frame categories plus the L6 M-family verdict events into one deterministic
// stream, and that every payload is taken from an actual SIM API or engine.
func TestMixedFramesRealSIMKinds(t *testing.T) {
	t.Parallel()
	frames, err := mixedFrames(42)
	if err != nil {
		t.Fatalf("mixedFrames: %v", err)
	}
	kinds := make(map[record.FrameKind]int)
	var mEvents, injectEvents, rollbackEvents int
	netPrefix, fsPrefix, taskMarker := true, true, false
	for _, frame := range frames {
		kinds[frame.Kind]++
		payload := string(frame.Payload)
		switch frame.Kind {
		case record.KindNet:
			netPrefix = netPrefix && strings.HasPrefix(payload, "net/")
		case record.KindFS:
			fsPrefix = fsPrefix && strings.HasPrefix(payload, "fs/")
		case record.KindVTask:
			taskMarker = taskMarker || strings.Contains(payload, "task-1/")
		case record.KindFault:
			switch {
			case strings.HasPrefix(payload, "M/"):
				mEvents++
			case strings.HasPrefix(payload, "inject:"):
				injectEvents++
			case strings.HasPrefix(payload, "rollback:"):
				rollbackEvents++
			}
		}
	}
	for _, kind := range []record.FrameKind{
		record.KindFault, record.KindNet, record.KindFS, record.KindVTask,
	} {
		if kinds[kind] == 0 {
			t.Fatalf("mixedFrames produced no %s frames: %v", kind, kinds)
		}
	}
	if !netPrefix || !fsPrefix || !taskMarker {
		t.Fatalf("SIM payloads are not real state: net=%t fs=%t task=%t",
			netPrefix, fsPrefix, taskMarker)
	}
	if mEvents == 0 || injectEvents != 12 || rollbackEvents != 12 {
		t.Fatalf("M⊗C events wrong: M=%d inject=%d rollback=%d",
			mEvents, injectEvents, rollbackEvents)
	}
}

// TestMixedFramesDeterministic proves 100 same-seed runs of the L6/L7
// orchestration yield exactly one canonical byte stream.
func TestMixedFramesDeterministic(t *testing.T) {
	t.Parallel()
	hashOf := func(seed uint64) string {
		frames, err := mixedFrames(seed)
		if err != nil {
			t.Fatalf("mixedFrames(%d): %v", seed, err)
		}
		var b bytes.Buffer
		for _, frame := range frames {
			_, _ = b.WriteString(frame.Kind.String())
			_, _ = b.Write(frame.Payload)
			_ = b.WriteByte(0)
		}
		sum := sha256.Sum256(b.Bytes())
		return hex.EncodeToString(sum[:])
	}
	want := hashOf(42)
	unique := map[string]int{want: 0}
	for i := 0; i < 100; i++ {
		unique[hashOf(42)]++
	}
	if len(unique) != 1 || unique[want] != 100 {
		t.Fatalf("same-seed mixedFrames produced %d distinct hashes", len(unique))
	}
	if hashOf(43) == want {
		t.Fatal("different seeds produced identical mixed streams")
	}
}

// TestMixedTargetSuiteContract locks the synthetic L6 ports to a stable,
// ordered, deterministic contract that the engine can enumerate.
func TestMixedTargetSuiteContract(t *testing.T) {
	t.Parallel()
	target := newMixedTarget()
	nodes := target.Nodes()
	if len(nodes) != 7 || nodes[0] != "n-arith" || nodes[6] != "n-return" {
		t.Fatalf("nodes not sorted/stable: %v", nodes)
	}
	for _, node := range nodes {
		if target.Token(node) == "" {
			t.Fatalf("node %s has no token", node)
		}
	}
	wantFingerprint := target.Fingerprint()
	if got := target.Fingerprint(); got != wantFingerprint {
		t.Fatal("target fingerprint unstable")
	}
	if !target.Equivalent(fakePoint("+", "+")) ||
		target.Equivalent(fakePoint("+", "-")) {
		t.Fatal("Equivalent disagrees with the mutation")
	}
	mutated := fakePoint("+", "-")
	if target.Evaluate(mutated, "t1") != verify.OutcomeHeld ||
		target.Evaluate(mutated, "t2") != verify.OutcomeViolated {
		t.Fatal("Evaluate disagrees with the killing-test boundary")
	}
	if target.Evaluate(fakePoint("+", "+"), "t3") != verify.OutcomeHeld {
		t.Fatal("equivalent mutant must hold")
	}
	suite := mixedSuite{}
	if len(suite.Tests()) != 3 || suite.Coverage("any") != 1 {
		t.Fatal("suite contract drifted")
	}
	wantSuiteFingerprint := suite.Fingerprint()
	if got := suite.Fingerprint(); got != wantSuiteFingerprint {
		t.Fatal("suite fingerprint unstable")
	}
}

// TestExperimentHandlerMixedReplay proves the upgraded experimentHandler seals
// all eight FrameKinds and replays frame by frame without divergence.
func TestExperimentHandlerMixedReplay(t *testing.T) {
	t.Parallel()
	build := func() *record.Archive {
		exp := loadFixture(t, "ok.yaml")
		times := probeSchedule(exp)
		pipe, err := buildPipeline(exp)
		if err != nil {
			t.Fatalf("buildPipeline: %v", err)
		}
		obs := collectObservations(pipe, times)
		h := newExperimentHandler(exp, times, obs)
		if _, err := h.handler()(applyEventForTest()); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		arch, err := h.stop()
		if err != nil {
			t.Fatalf("stop: %v", err)
		}
		return arch
	}
	first, second := build(), build()
	if err := (record.Replayer{}).Verify(first); err != nil {
		t.Fatalf("archive invalid: %v", err)
	}
	kinds := make(map[record.FrameKind]int)
	for i := 0; i < first.Len(); i++ {
		kinds[first.Frame(i).Kind]++
	}
	for _, kind := range []record.FrameKind{
		record.KindSched, record.KindObs, record.KindRNG, record.KindClk,
		record.KindFault, record.KindNet, record.KindFS, record.KindVTask,
	} {
		if kinds[kind] == 0 {
			t.Fatalf("handler archive missing %s frames: %v", kind, kinds)
		}
	}
	left, right := &replayCollector{}, &replayCollector{}
	rp := record.Replayer{}
	if err := rp.Replay(first, left); err != nil {
		t.Fatalf("replay left: %v", err)
	}
	if err := rp.Replay(second, right); err != nil {
		t.Fatalf("replay right: %v", err)
	}
	if len(left.frames) != len(right.frames) || len(left.frames) != first.Len() {
		t.Fatalf("replay frame counts differ: %d/%d/%d",
			len(left.frames), len(right.frames), first.Len())
	}
	for i := range left.frames {
		a, b := left.frames[i], right.frames[i]
		if a.Seq != b.Seq || a.At != b.At || a.Kind != b.Kind ||
			!bytes.Equal(a.Payload, b.Payload) {
			t.Fatalf("frame %d diverges between handlers", i)
		}
	}
	t.Logf("M⊗C handler archive: %d frames, 8 kinds = %v", first.Len(), kinds)
}
