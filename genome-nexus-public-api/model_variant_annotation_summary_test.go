package genome_nexus_public_api

import (
	"encoding/json"
	"testing"
)

// TestVariantAnnotationSummaryUnmarshalMissingSingularField is a regression test
// for a bug where UnmarshalJSON treated "transcriptConsequenceSummary" (singular)
// as unconditionally required. Genome Nexus's actual API never sends that field —
// only the pluralized "transcriptConsequenceSummaries" array — for variants
// resolved as intergenic (no single transcript to report a canonical consequence
// for). The spec still listed the old singular field as required, so every
// annotation_summary for an intergenic variant failed JSON decoding entirely, even
// though the variant itself annotated successfully.
func TestVariantAnnotationSummaryUnmarshalMissingSingularField(t *testing.T) {
	// Actual response observed from Genome Nexus for an intergenic variant
	// (successfully_annotated=true at the top level; this is its annotation_summary).
	data := []byte(`{
		"variant": "7:g.1394614_1394615delinsCT",
		"genomicLocation": {"chromosome": "7", "start": 1394614, "end": 1394615, "referenceAllele": "GG", "variantAllele": "CT"},
		"strandSign": "+",
		"variantType": "DNP",
		"assemblyName": "GRCh37",
		"transcriptConsequences": [],
		"transcriptConsequenceSummaries": []
	}`)

	var vas VariantAnnotationSummary
	if err := json.Unmarshal(data, &vas); err != nil {
		t.Fatalf("UnmarshalJSON failed for a response missing \"transcriptConsequenceSummary\": %v", err)
	}

	if vas.TranscriptConsequenceSummary != nil {
		t.Errorf("TranscriptConsequenceSummary = %v, want nil (never provided by Genome Nexus for intergenic variants)", vas.TranscriptConsequenceSummary)
	}
	if vas.Variant != "7:g.1394614_1394615delinsCT" {
		t.Errorf("Variant = %q, want the variant string", vas.Variant)
	}
}

// TestVariantAnnotationSummaryUnmarshalMissingVariant confirms the remaining
// required properties are still enforced.
func TestVariantAnnotationSummaryUnmarshalMissingVariant(t *testing.T) {
	data := []byte(`{"transcriptConsequenceSummaries": [], "transcriptConsequences": []}`)

	var vas VariantAnnotationSummary
	if err := json.Unmarshal(data, &vas); err == nil {
		t.Fatalf("UnmarshalJSON succeeded without variant, want an error")
	}
}
