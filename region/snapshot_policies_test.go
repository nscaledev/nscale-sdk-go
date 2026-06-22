package region_test

import (
	"encoding/json"
	"testing"

	"github.com/nscaledev/nscale-sdk-go/region"
)

func TestFileStorageCreateWithUnsetSnapshotPoliciesEncodesNull(t *testing.T) {
	var body region.StorageV2Create
	body.Spec.SnapshotPolicies = nil

	snapshotPolicies, ok := encodedSpecSnapshotPolicies(t, body, "File Storage create request")
	if !ok {
		t.Fatal("encoded request spec omitted snapshotPolicies, want explicit null")
	}

	if snapshotPolicies != nil {
		t.Fatalf("encoded request spec snapshotPolicies = %#v, want nil", snapshotPolicies)
	}
}

func TestFileStorageCreateWithEmptySnapshotPolicySetEncodesEmptyArray(t *testing.T) {
	var body region.StorageV2Create
	snapshotPolicies := region.StorageSnapshotPolicyListV2Spec{}
	body.Spec.SnapshotPolicies = &snapshotPolicies

	snapshotPoliciesJSON, ok := encodedSpecSnapshotPolicies(t, body, "File Storage create request")
	if !ok {
		t.Fatal("encoded request spec omitted snapshotPolicies, want empty array")
	}

	snapshotPoliciesArray, ok := snapshotPoliciesJSON.([]any)
	if !ok {
		t.Fatalf("encoded request spec snapshotPolicies = %T, want array", snapshotPoliciesJSON)
	}

	if len(snapshotPoliciesArray) != 0 {
		t.Fatalf("encoded request spec snapshotPolicies length = %d, want 0", len(snapshotPoliciesArray))
	}
}

func TestFileStorageUpdateWithUnsetSnapshotPoliciesEncodesNull(t *testing.T) {
	var body region.StorageV2Update
	body.Spec.SnapshotPolicies = nil

	snapshotPolicies, ok := encodedSpecSnapshotPolicies(t, body, "File Storage update request")
	if !ok {
		t.Fatal("encoded request spec omitted snapshotPolicies, want explicit null")
	}

	if snapshotPolicies != nil {
		t.Fatalf("encoded request spec snapshotPolicies = %#v, want nil", snapshotPolicies)
	}
}

func TestFileStorageUpdateWithEmptySnapshotPolicySetEncodesEmptyArray(t *testing.T) {
	var body region.StorageV2Update
	snapshotPolicies := region.StorageSnapshotPolicyListV2Spec{}
	body.Spec.SnapshotPolicies = &snapshotPolicies

	snapshotPoliciesJSON, ok := encodedSpecSnapshotPolicies(t, body, "File Storage update request")
	if !ok {
		t.Fatal("encoded request spec omitted snapshotPolicies, want empty array")
	}

	snapshotPoliciesArray, ok := snapshotPoliciesJSON.([]any)
	if !ok {
		t.Fatalf("encoded request spec snapshotPolicies = %T, want array", snapshotPoliciesJSON)
	}

	if len(snapshotPoliciesArray) != 0 {
		t.Fatalf("encoded request spec snapshotPolicies length = %d, want 0", len(snapshotPoliciesArray))
	}
}

func encodedSpecSnapshotPolicies(t *testing.T, body any, requestName string) (any, bool) {
	t.Helper()

	encoded, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal %s: %v", requestName, err)
	}

	var request map[string]any
	if err := json.Unmarshal(encoded, &request); err != nil {
		t.Fatalf("unmarshal %s JSON: %v", requestName, err)
	}

	spec, ok := request["spec"].(map[string]any)
	if !ok {
		t.Fatalf("encoded request spec = %T, want object", request["spec"])
	}

	snapshotPolicies, ok := spec["snapshotPolicies"]
	return snapshotPolicies, ok
}
