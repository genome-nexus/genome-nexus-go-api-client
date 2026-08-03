package genome_nexus_public_api

import (
	"encoding/json"
	"testing"
)

// TestVariantAnnotationUnmarshalMissingId is a regression test for a bug where
// UnmarshalJSON treated "id" as unconditionally required on every element of a
// batch annotation response. Genome Nexus only includes "id" for successfully
// annotated variants; a variant that fails to annotate (e.g. a reference-allele
// mismatch) comes back with successfully_annotated=false, an errorMessage,
// originalVariantQuery, and variant, but no id. Since a batch call decodes the
// entire response array at once, one such variant failed JSON decoding for the
// whole array, silently discarding every other (otherwise valid) variant
// annotation returned in the same batch.
func TestVariantAnnotationUnmarshalMissingId(t *testing.T) {
	// Actual response observed from Genome Nexus for a variant with a
	// reference-allele mismatch (11:57004345-57004346, GC>TT, ref genome has CC).
	data := []byte(`{
		"variant": "11:g.57004345_57004346delinsTT",
		"errorMessage": "Reference allele extracted from response (CC) does not match given reference allele (GC)",
		"originalVariantQuery": "11,57004345,57004346,GC,TT",
		"successfully_annotated": false
	}`)

	var va VariantAnnotation
	if err := json.Unmarshal(data, &va); err != nil {
		t.Fatalf("UnmarshalJSON failed for a variant missing \"id\": %v", err)
	}

	if va.Id != nil {
		t.Errorf("Id = %v, want nil (never provided by Genome Nexus for this response)", va.Id)
	}
	if va.OriginalVariantQuery != "11,57004345,57004346,GC,TT" {
		t.Errorf("OriginalVariantQuery = %q, want the query string", va.OriginalVariantQuery)
	}
	if va.SuccessfullyAnnotated == nil || *va.SuccessfullyAnnotated {
		t.Errorf("SuccessfullyAnnotated = %v, want false", va.SuccessfullyAnnotated)
	}
}

// TestVariantAnnotationUnmarshalWithId confirms a normal, successful response
// (which does include "id") still decodes correctly.
func TestVariantAnnotationUnmarshalWithId(t *testing.T) {
	data := []byte(`{
		"variant": "10:g.63852470C>T",
		"originalVariantQuery": "10,63852470,63852470,C,T",
		"id": "10:g.63852470C>T",
		"successfully_annotated": true
	}`)

	var va VariantAnnotation
	if err := json.Unmarshal(data, &va); err != nil {
		t.Fatalf("UnmarshalJSON failed for a normal successful response: %v", err)
	}
	if va.Id == nil || *va.Id != "10:g.63852470C>T" {
		t.Errorf("Id = %v, want %q", va.Id, "10:g.63852470C>T")
	}
}

// TestVariantAnnotationUnmarshalMissingOriginalVariantQuery confirms the
// remaining required properties are still enforced.
func TestVariantAnnotationUnmarshalMissingOriginalVariantQuery(t *testing.T) {
	data := []byte(`{"variant": "10:g.63852470C>T"}`)

	var va VariantAnnotation
	if err := json.Unmarshal(data, &va); err == nil {
		t.Fatalf("UnmarshalJSON succeeded without originalVariantQuery, want an error")
	}
}
