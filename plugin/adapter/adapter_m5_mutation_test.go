/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"testing"
)

// TestM5MutationAllowListChangesRejection proves the field allow-list is a
// real assertion and not a decorative field.
func TestM5MutationAllowListChangesRejection(t *testing.T) {
	t.Parallel()
	base := NewPayloadAssertion(itsRequiredFields)
	payload := m5ITSPayload()
	if err := base.Assert(payload); err != nil {
		t.Fatalf("baseline assertion rejected a valid payload: %v", err)
	}
	mutated := base
	mutated.Allowed = []string{"type"}
	if err := mutated.Assert(payload); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("mutated allow-list error = %v, want a rejection", err)
	}
	complete := base
	complete.Allowed = append([]string(nil), itsRequiredFields...)
	if err := complete.Assert(payload); err != nil {
		t.Fatalf("complete allow-list rejected a valid payload: %v", err)
	}
}

// TestM5MutationRequiredFieldsChangesRejection proves adding a mandatory
// field changes the decision for one fixed payload.
func TestM5MutationRequiredFieldsChangesRejection(t *testing.T) {
	t.Parallel()
	base := NewPayloadAssertion(itsRequiredFields)
	if err := base.Assert(m5ITSPayload()); err != nil {
		t.Fatalf("baseline assertion: %v", err)
	}
	mutated := NewPayloadAssertion(append(append([]string(nil), itsRequiredFields...), "evidence"))
	if err := mutated.Assert(m5ITSPayload()); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("mutated required list error = %v, want a rejection", err)
	}
}

// TestM5MutationSizeCapChangesRejection proves the size cap is enforced.
func TestM5MutationSizeCapChangesRejection(t *testing.T) {
	t.Parallel()
	base := NewPayloadAssertion(itsRequiredFields)
	payload := m5ITSPayload()
	if err := base.Assert(payload); err != nil {
		t.Fatalf("baseline assertion: %v", err)
	}
	tight := base
	tight.MaxBytes = 8
	if err := tight.Assert(payload); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("tight cap error = %v, want a rejection", err)
	}
	zero := base
	zero.MaxBytes = 0
	if err := zero.Assert(payload); err != nil {
		t.Fatalf("zero cap should select the default: %v", err)
	}
	negative := base
	negative.MaxBytes = -1
	if err := negative.Assert(payload); err != nil {
		t.Fatalf("negative cap should select the default: %v", err)
	}
}

// TestM5MutationPayloadFieldChangesRejection proves a single outbound field
// insertion flips the decision from accept to reject.
func TestM5MutationPayloadFieldChangesRejection(t *testing.T) {
	t.Parallel()
	assertion := NewPayloadAssertion(msgRequiredFields)
	if err := assertion.Assert(m5MSGPayload()); err != nil {
		t.Fatalf("baseline payload rejected: %v", err)
	}
	mutated := []byte(`{"kind":"report","level":"info","summary":"deterministic build","url":"x"}`)
	if err := assertion.Assert(mutated); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("mutated payload error = %v, want a rejection", err)
	}
}

// TestM5MutationEnvValueChangesDetection proves detection is driven solely by
// the injected map value rather than by any ambient process state.
func TestM5MutationEnvValueChangesDetection(t *testing.T) {
	t.Parallel()
	its := newITSAdapter(m5AdapterManifest("abate/adp-its"))
	env := map[string]string{"ABATE_ADP_ITS": "true"}
	if _, ok := its.Detect(env); !ok {
		t.Fatal("enabled value was not detected")
	}
	env["ABATE_ADP_ITS"] = "false"
	if _, ok := its.Detect(env); ok {
		t.Fatal("disabled value was still detected")
	}
	delete(env, "ABATE_ADP_ITS")
	if _, ok := its.Detect(env); ok {
		t.Fatal("a removed marker was still detected")
	}
}
