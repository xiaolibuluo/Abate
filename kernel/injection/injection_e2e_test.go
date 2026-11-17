/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/kernel/verify"
	"abate/sim/clock"
	"abate/sim/rand"
	"abate/sim/record"
)

// e2eTarget is a synthetic M-side target used by the mixed M⊗C experiment.
type e2eTarget struct {
	nodes  []string
	tokens map[string]string
}

func newE2ETarget() *e2eTarget {
	return &e2eTarget{
		nodes: []string{"a", "b", "c", "d", "e", "f", "g"},
		tokens: map[string]string{
			"a": "+", "b": "<", "c": "1", "d": "break",
			"e": "true", "f": "==", "g": "call",
		},
	}
}

func (t *e2eTarget) Fingerprint() [32]byte {
	var b strings.Builder
	for _, node := range t.nodes {
		_, _ = b.WriteString(node + "=" + t.tokens[node] + ";")
	}
	return sha256.Sum256([]byte(b.String()))
}

func (t *e2eTarget) Nodes() []string          { return append([]string(nil), t.nodes...) }
func (t *e2eTarget) Token(node string) string { return t.tokens[node] }
func (*e2eTarget) Equivalent(p mutation.MutationPoint) bool {
	return p.Mutated == p.Original
}
func (*e2eTarget) Evaluate(_ mutation.MutationPoint, test string) verify.Outcome {
	if test >= "t2" {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// e2eSuite is the synthetic M-side test selector.
type e2eSuite struct{}

func (e2eSuite) Fingerprint() [32]byte {
	return sha256.Sum256([]byte("m2-e2e-suite"))
}
func (e2eSuite) Tests() []string          { return []string{"t1", "t2", "t3"} }
func (e2eSuite) Coverage(_ string) uint64 { return 1 }

// e2eHooks collects replayed frames for a frame-by-frame comparison.
type e2eHooks struct {
	frames []record.Frame
}

func (h *e2eHooks) OnFrame(f record.Frame) error {
	h.frames = append(h.frames, f)
	return nil
}

// buildM2Archive orchestrates one M⊗C experiment: the mutation Engine derives
// M-family events while the Injector applies and compensates all twelve
// C-family faults, and the recorder seals eight FrameKind categories.
func buildM2Archive(seed uint64) (*record.Archive, error) {
	ctx := context.Background()
	report, err := mutation.NewEngine(seed).Run(ctx, newE2ETarget(), e2eSuite{})
	if err != nil {
		return nil, fmt.Errorf("mutation engine: %w", err)
	}
	world, err := NewWorld(seed)
	if err != nil {
		return nil, err
	}
	injector, err := NewInjector(world)
	if err != nil {
		return nil, err
	}
	for _, op := range injector.Registry().All() {
		if err := injector.Inject(ctx, op); err != nil {
			return nil, fmt.Errorf("inject %s: %w", op.ID(), err)
		}
		if err := injector.Rollback(ctx, op); err != nil {
			return nil, fmt.Errorf("rollback %s: %w", op.ID(), err)
		}
	}
	stream := rand.New(seed)
	for i := 0; i < 4; i++ {
		_ = stream.Uint64()
	}
	at := clock.Instant(9000)
	observations := []observe.Observation{{
		At: 9000, Seq: 1, ProbeID: "health",
		Value: observe.NewNumber(200), Status: observe.StatusOK,
	}}
	fingerprint := sha256.Sum256([]byte(fmt.Sprintf("m2-mixed/%d", seed)))
	recorder := record.NewRecorder(fingerprint)
	recorder.Subscribe(
		record.NewRNGSource(at, stream),
		record.NewCLKSource(world.Clock()),
		record.NewOBSSource(at, observations),
	)
	recorder.Start(ctx)
	handler := recorder.WrapHandler(func(_ scheduler.Event) ([]scheduler.Event, error) {
		return nil, nil
	})
	for index, when := range []scheduler.VirtualTime{1000, 2000, 3000} {
		if _, err := handler(scheduler.Event{
			At: when, Kind: scheduler.EventApply, Payload: []byte{byte(index)},
		}); err != nil {
			return nil, err
		}
	}
	frames := injector.Frames()
	for _, frame := range frames {
		recorder.Record(frame)
	}
	lastAt := frames[len(frames)-1].At
	for index, verdict := range report.Verdicts {
		payload := fmt.Sprintf("M/%d/%s/%s", index,
			report.Points[index].Operator.ID(), verdict)
		recorder.Record(record.Frame{
			At: lastAt, Kind: record.KindFault, Payload: []byte(payload),
		})
	}
	return recorder.Stop()
}

func TestMixedPerturbationFrameReplay(t *testing.T) {
	t.Parallel()
	first, err := buildM2Archive(42)
	if err != nil {
		t.Fatalf("first archive: %v", err)
	}
	second, err := buildM2Archive(42)
	if err != nil {
		t.Fatalf("second archive: %v", err)
	}
	replayer := record.Replayer{}
	if err := replayer.Verify(first); err != nil {
		t.Fatalf("archive verification failed: %v", err)
	}
	kinds := make(map[record.FrameKind]int)
	for i := 0; i < first.Len(); i++ {
		kinds[first.Frame(i).Kind]++
	}
	want := []record.FrameKind{
		record.KindSched, record.KindObs, record.KindRNG, record.KindClk,
		record.KindFault, record.KindNet, record.KindFS, record.KindVTask,
	}
	for _, kind := range want {
		if kinds[kind] == 0 {
			t.Fatalf("archive is missing %s frames: %v", kind, kinds)
		}
	}
	t.Logf("M⊗C archive: %d frames, 8 kinds = %v", first.Len(), kinds)
	left, right := &e2eHooks{}, &e2eHooks{}
	if err := replayer.Replay(first, left); err != nil {
		t.Fatalf("first replay: %v", err)
	}
	if err := replayer.Replay(second, right); err != nil {
		t.Fatalf("second replay: %v", err)
	}
	if len(left.frames) != len(right.frames) || len(left.frames) != first.Len() {
		t.Fatalf("replay frame counts differ: %d vs %d vs %d",
			len(left.frames), len(right.frames), first.Len())
	}
	for i := range left.frames {
		a, b := left.frames[i], right.frames[i]
		if a.Seq != b.Seq || a.At != b.At || a.Kind != b.Kind ||
			!bytes.Equal(a.Payload, b.Payload) {
			t.Fatalf("frame %d differs between replays: %+v vs %+v", i, a, b)
		}
	}
	diff, err := record.Compare(first, second)
	if err != nil || !diff.Equal {
		t.Fatalf("structured compare disagrees: equal=%v err=%v", diff.Equal, err)
	}
}
