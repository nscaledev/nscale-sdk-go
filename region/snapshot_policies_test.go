package region_test

import (
	"encoding/json"
	"testing"

	"github.com/nscaledev/nscale-sdk-go/region"
)

func TestFileStorageCreateWithUnsetSnapshotPoliciesOmitsField(t *testing.T) {
	var body region.StorageV2Create
	body.Spec.SnapshotPolicies = nil

	_, ok := encodedSpecSnapshotPolicies(t, body, "File Storage create request")
	if ok {
		t.Fatal("encoded request spec included snapshotPolicies, want omitted")
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

func TestFileStorageUpdateWithUnsetSnapshotPoliciesOmitsField(t *testing.T) {
	var body region.StorageV2Update
	body.Spec.SnapshotPolicies = nil

	_, ok := encodedSpecSnapshotPolicies(t, body, "File Storage update request")
	if ok {
		t.Fatal("encoded request spec included snapshotPolicies, want omitted")
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

func TestFileStorageCreateWithUnsetDefaultSnapshotProtectionOmitsField(t *testing.T) {
	var body region.StorageV2Create
	body.Spec.DefaultSnapshotProtectionEnabled = nil

	_, ok := encodedSpecDefaultSnapshotProtectionEnabled(t, body, "File Storage create request")
	if ok {
		t.Fatal("encoded request spec included defaultSnapshotProtectionEnabled, want omitted")
	}
}

func TestFileStorageCreateWithDisabledDefaultSnapshotProtectionEncodesFalse(t *testing.T) {
	var body region.StorageV2Create
	defaultSnapshotProtectionEnabled := false
	body.Spec.DefaultSnapshotProtectionEnabled = &defaultSnapshotProtectionEnabled

	encodedDefaultSnapshotProtectionEnabled, ok := encodedSpecDefaultSnapshotProtectionEnabled(t, body, "File Storage create request")
	if !ok {
		t.Fatal("encoded request spec omitted defaultSnapshotProtectionEnabled, want false")
	}

	if encodedDefaultSnapshotProtectionEnabled != false {
		t.Fatalf("encoded request spec defaultSnapshotProtectionEnabled = %#v, want false", encodedDefaultSnapshotProtectionEnabled)
	}
}

func TestFileStorageUpdateWithUnsetDefaultSnapshotProtectionOmitsField(t *testing.T) {
	var body region.StorageV2Update
	body.Spec.DefaultSnapshotProtectionEnabled = nil

	_, ok := encodedSpecDefaultSnapshotProtectionEnabled(t, body, "File Storage update request")
	if ok {
		t.Fatal("encoded request spec included defaultSnapshotProtectionEnabled, want omitted")
	}
}

func TestFileStorageUpdateWithEnabledDefaultSnapshotProtectionEncodesTrue(t *testing.T) {
	var body region.StorageV2Update
	defaultSnapshotProtectionEnabled := true
	body.Spec.DefaultSnapshotProtectionEnabled = &defaultSnapshotProtectionEnabled

	encodedDefaultSnapshotProtectionEnabled, ok := encodedSpecDefaultSnapshotProtectionEnabled(t, body, "File Storage update request")
	if !ok {
		t.Fatal("encoded request spec omitted defaultSnapshotProtectionEnabled, want true")
	}

	if encodedDefaultSnapshotProtectionEnabled != true {
		t.Fatalf("encoded request spec defaultSnapshotProtectionEnabled = %#v, want true", encodedDefaultSnapshotProtectionEnabled)
	}
}

func encodedSpecSnapshotPolicies(t *testing.T, body any, requestName string) (any, bool) {
	t.Helper()

	return encodedSpecField(t, body, requestName, "snapshotPolicies")
}

func encodedSpecDefaultSnapshotProtectionEnabled(t *testing.T, body any, requestName string) (any, bool) {
	t.Helper()

	return encodedSpecField(t, body, requestName, "defaultSnapshotProtectionEnabled")
}

func encodedSpecField(t *testing.T, body any, requestName string, field string) (any, bool) {
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

	value, ok := spec[field]
	return value, ok
}
