/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"encoding/json"
)

// MutationAssessor struct for MutationAssessor
type MutationAssessor struct {
	// Codon start position
	CodonStartPosition *string `json:"codonStartPosition,omitempty"`
	// Number of mutations in COSMIC for this protein
	CosmicCount *int32 `json:"cosmicCount,omitempty"`
	// Functional impact
	FunctionalImpact *string `json:"functionalImpact,omitempty"`
	// Functional impact score
	FunctionalImpactScore *float64 `json:"functionalImpactScore,omitempty"`
	Hgvs *string `json:"hgvs,omitempty"`
	// Hugo gene symbol
	HugoSymbol *string `json:"hugoSymbol,omitempty"`
	// User-input variants
	Input string `json:"input"`
	// Mapping issue info
	MappingIssue *string `json:"mappingIssue,omitempty"`
	// Portion of gaps in variant position in multiple sequence alignment
	MsaGaps *float64 `json:"msaGaps,omitempty"`
	// Number of diverse sequences in multiple sequence alignment (identical or highly similar sequences filtered out)
	MsaHeight *int32 `json:"msaHeight,omitempty"`
	// Link to multiple sequence alignment
	MsaLink *string `json:"msaLink,omitempty"`
	// Link to 3d structure browser
	PdbLink *string `json:"pdbLink,omitempty"`
	// Reference genome variant
	ReferenceGenomeVariant *string `json:"referenceGenomeVariant,omitempty"`
	// Reference genome variant type
	ReferenceGenomeVariantType *string `json:"referenceGenomeVariantType,omitempty"`
	// Refseq protein ID
	RefseqId *string `json:"refseqId,omitempty"`
	// Variant position in Refseq protein, can be different from the one in Uniprot
	RefseqPosition *int32 `json:"refseqPosition,omitempty"`
	// Reference residue in Refseq protein, can be different from the one in Uniprot
	RefseqResidue *string `json:"refseqResidue,omitempty"`
	// Number of SNPs in dbSNP for this protein
	SnpCount *int32 `json:"snpCount,omitempty"`
	// Uniprot protein accession ID
	UniprotId *string `json:"uniprotId,omitempty"`
	// Variant position in Uniprot protein, can be different from the one in Refseq
	UniprotPosition *int32 `json:"uniprotPosition,omitempty"`
	// Reference residue in Uniprot protein, can be different from the one in Refseq
	UniprotResidue *string `json:"uniprotResidue,omitempty"`
	// Amino acid substitution
	Variant *string `json:"variant,omitempty"`
	// Variant conservation score
	VariantConservationScore *float64 `json:"variantConservationScore,omitempty"`
	// Variant specificity score
	VariantSpecificityScore *float64 `json:"variantSpecificityScore,omitempty"`
}

// NewMutationAssessor instantiates a new MutationAssessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutationAssessor(input string) *MutationAssessor {
	this := MutationAssessor{}
	this.Input = input
	return &this
}

// NewMutationAssessorWithDefaults instantiates a new MutationAssessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutationAssessorWithDefaults() *MutationAssessor {
	this := MutationAssessor{}
	return &this
}

// GetCodonStartPosition returns the CodonStartPosition field value if set, zero value otherwise.
func (o *MutationAssessor) GetCodonStartPosition() string {
	if o == nil || isNil(o.CodonStartPosition) {
		var ret string
		return ret
	}
	return *o.CodonStartPosition
}

// GetCodonStartPositionOk returns a tuple with the CodonStartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetCodonStartPositionOk() (*string, bool) {
	if o == nil || isNil(o.CodonStartPosition) {
    return nil, false
	}
	return o.CodonStartPosition, true
}

// HasCodonStartPosition returns a boolean if a field has been set.
func (o *MutationAssessor) HasCodonStartPosition() bool {
	if o != nil && !isNil(o.CodonStartPosition) {
		return true
	}

	return false
}

// SetCodonStartPosition gets a reference to the given string and assigns it to the CodonStartPosition field.
func (o *MutationAssessor) SetCodonStartPosition(v string) {
	o.CodonStartPosition = &v
}

// GetCosmicCount returns the CosmicCount field value if set, zero value otherwise.
func (o *MutationAssessor) GetCosmicCount() int32 {
	if o == nil || isNil(o.CosmicCount) {
		var ret int32
		return ret
	}
	return *o.CosmicCount
}

// GetCosmicCountOk returns a tuple with the CosmicCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetCosmicCountOk() (*int32, bool) {
	if o == nil || isNil(o.CosmicCount) {
    return nil, false
	}
	return o.CosmicCount, true
}

// HasCosmicCount returns a boolean if a field has been set.
func (o *MutationAssessor) HasCosmicCount() bool {
	if o != nil && !isNil(o.CosmicCount) {
		return true
	}

	return false
}

// SetCosmicCount gets a reference to the given int32 and assigns it to the CosmicCount field.
func (o *MutationAssessor) SetCosmicCount(v int32) {
	o.CosmicCount = &v
}

// GetFunctionalImpact returns the FunctionalImpact field value if set, zero value otherwise.
func (o *MutationAssessor) GetFunctionalImpact() string {
	if o == nil || isNil(o.FunctionalImpact) {
		var ret string
		return ret
	}
	return *o.FunctionalImpact
}

// GetFunctionalImpactOk returns a tuple with the FunctionalImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetFunctionalImpactOk() (*string, bool) {
	if o == nil || isNil(o.FunctionalImpact) {
    return nil, false
	}
	return o.FunctionalImpact, true
}

// HasFunctionalImpact returns a boolean if a field has been set.
func (o *MutationAssessor) HasFunctionalImpact() bool {
	if o != nil && !isNil(o.FunctionalImpact) {
		return true
	}

	return false
}

// SetFunctionalImpact gets a reference to the given string and assigns it to the FunctionalImpact field.
func (o *MutationAssessor) SetFunctionalImpact(v string) {
	o.FunctionalImpact = &v
}

// GetFunctionalImpactScore returns the FunctionalImpactScore field value if set, zero value otherwise.
func (o *MutationAssessor) GetFunctionalImpactScore() float64 {
	if o == nil || isNil(o.FunctionalImpactScore) {
		var ret float64
		return ret
	}
	return *o.FunctionalImpactScore
}

// GetFunctionalImpactScoreOk returns a tuple with the FunctionalImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetFunctionalImpactScoreOk() (*float64, bool) {
	if o == nil || isNil(o.FunctionalImpactScore) {
    return nil, false
	}
	return o.FunctionalImpactScore, true
}

// HasFunctionalImpactScore returns a boolean if a field has been set.
func (o *MutationAssessor) HasFunctionalImpactScore() bool {
	if o != nil && !isNil(o.FunctionalImpactScore) {
		return true
	}

	return false
}

// SetFunctionalImpactScore gets a reference to the given float64 and assigns it to the FunctionalImpactScore field.
func (o *MutationAssessor) SetFunctionalImpactScore(v float64) {
	o.FunctionalImpactScore = &v
}

// GetHgvs returns the Hgvs field value if set, zero value otherwise.
func (o *MutationAssessor) GetHgvs() string {
	if o == nil || isNil(o.Hgvs) {
		var ret string
		return ret
	}
	return *o.Hgvs
}

// GetHgvsOk returns a tuple with the Hgvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetHgvsOk() (*string, bool) {
	if o == nil || isNil(o.Hgvs) {
    return nil, false
	}
	return o.Hgvs, true
}

// HasHgvs returns a boolean if a field has been set.
func (o *MutationAssessor) HasHgvs() bool {
	if o != nil && !isNil(o.Hgvs) {
		return true
	}

	return false
}

// SetHgvs gets a reference to the given string and assigns it to the Hgvs field.
func (o *MutationAssessor) SetHgvs(v string) {
	o.Hgvs = &v
}

// GetHugoSymbol returns the HugoSymbol field value if set, zero value otherwise.
func (o *MutationAssessor) GetHugoSymbol() string {
	if o == nil || isNil(o.HugoSymbol) {
		var ret string
		return ret
	}
	return *o.HugoSymbol
}

// GetHugoSymbolOk returns a tuple with the HugoSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetHugoSymbolOk() (*string, bool) {
	if o == nil || isNil(o.HugoSymbol) {
    return nil, false
	}
	return o.HugoSymbol, true
}

// HasHugoSymbol returns a boolean if a field has been set.
func (o *MutationAssessor) HasHugoSymbol() bool {
	if o != nil && !isNil(o.HugoSymbol) {
		return true
	}

	return false
}

// SetHugoSymbol gets a reference to the given string and assigns it to the HugoSymbol field.
func (o *MutationAssessor) SetHugoSymbol(v string) {
	o.HugoSymbol = &v
}

// GetInput returns the Input field value
func (o *MutationAssessor) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetInputOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *MutationAssessor) SetInput(v string) {
	o.Input = v
}

// GetMappingIssue returns the MappingIssue field value if set, zero value otherwise.
func (o *MutationAssessor) GetMappingIssue() string {
	if o == nil || isNil(o.MappingIssue) {
		var ret string
		return ret
	}
	return *o.MappingIssue
}

// GetMappingIssueOk returns a tuple with the MappingIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetMappingIssueOk() (*string, bool) {
	if o == nil || isNil(o.MappingIssue) {
    return nil, false
	}
	return o.MappingIssue, true
}

// HasMappingIssue returns a boolean if a field has been set.
func (o *MutationAssessor) HasMappingIssue() bool {
	if o != nil && !isNil(o.MappingIssue) {
		return true
	}

	return false
}

// SetMappingIssue gets a reference to the given string and assigns it to the MappingIssue field.
func (o *MutationAssessor) SetMappingIssue(v string) {
	o.MappingIssue = &v
}

// GetMsaGaps returns the MsaGaps field value if set, zero value otherwise.
func (o *MutationAssessor) GetMsaGaps() float64 {
	if o == nil || isNil(o.MsaGaps) {
		var ret float64
		return ret
	}
	return *o.MsaGaps
}

// GetMsaGapsOk returns a tuple with the MsaGaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetMsaGapsOk() (*float64, bool) {
	if o == nil || isNil(o.MsaGaps) {
    return nil, false
	}
	return o.MsaGaps, true
}

// HasMsaGaps returns a boolean if a field has been set.
func (o *MutationAssessor) HasMsaGaps() bool {
	if o != nil && !isNil(o.MsaGaps) {
		return true
	}

	return false
}

// SetMsaGaps gets a reference to the given float64 and assigns it to the MsaGaps field.
func (o *MutationAssessor) SetMsaGaps(v float64) {
	o.MsaGaps = &v
}

// GetMsaHeight returns the MsaHeight field value if set, zero value otherwise.
func (o *MutationAssessor) GetMsaHeight() int32 {
	if o == nil || isNil(o.MsaHeight) {
		var ret int32
		return ret
	}
	return *o.MsaHeight
}

// GetMsaHeightOk returns a tuple with the MsaHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetMsaHeightOk() (*int32, bool) {
	if o == nil || isNil(o.MsaHeight) {
    return nil, false
	}
	return o.MsaHeight, true
}

// HasMsaHeight returns a boolean if a field has been set.
func (o *MutationAssessor) HasMsaHeight() bool {
	if o != nil && !isNil(o.MsaHeight) {
		return true
	}

	return false
}

// SetMsaHeight gets a reference to the given int32 and assigns it to the MsaHeight field.
func (o *MutationAssessor) SetMsaHeight(v int32) {
	o.MsaHeight = &v
}

// GetMsaLink returns the MsaLink field value if set, zero value otherwise.
func (o *MutationAssessor) GetMsaLink() string {
	if o == nil || isNil(o.MsaLink) {
		var ret string
		return ret
	}
	return *o.MsaLink
}

// GetMsaLinkOk returns a tuple with the MsaLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetMsaLinkOk() (*string, bool) {
	if o == nil || isNil(o.MsaLink) {
    return nil, false
	}
	return o.MsaLink, true
}

// HasMsaLink returns a boolean if a field has been set.
func (o *MutationAssessor) HasMsaLink() bool {
	if o != nil && !isNil(o.MsaLink) {
		return true
	}

	return false
}

// SetMsaLink gets a reference to the given string and assigns it to the MsaLink field.
func (o *MutationAssessor) SetMsaLink(v string) {
	o.MsaLink = &v
}

// GetPdbLink returns the PdbLink field value if set, zero value otherwise.
func (o *MutationAssessor) GetPdbLink() string {
	if o == nil || isNil(o.PdbLink) {
		var ret string
		return ret
	}
	return *o.PdbLink
}

// GetPdbLinkOk returns a tuple with the PdbLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetPdbLinkOk() (*string, bool) {
	if o == nil || isNil(o.PdbLink) {
    return nil, false
	}
	return o.PdbLink, true
}

// HasPdbLink returns a boolean if a field has been set.
func (o *MutationAssessor) HasPdbLink() bool {
	if o != nil && !isNil(o.PdbLink) {
		return true
	}

	return false
}

// SetPdbLink gets a reference to the given string and assigns it to the PdbLink field.
func (o *MutationAssessor) SetPdbLink(v string) {
	o.PdbLink = &v
}

// GetReferenceGenomeVariant returns the ReferenceGenomeVariant field value if set, zero value otherwise.
func (o *MutationAssessor) GetReferenceGenomeVariant() string {
	if o == nil || isNil(o.ReferenceGenomeVariant) {
		var ret string
		return ret
	}
	return *o.ReferenceGenomeVariant
}

// GetReferenceGenomeVariantOk returns a tuple with the ReferenceGenomeVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetReferenceGenomeVariantOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceGenomeVariant) {
    return nil, false
	}
	return o.ReferenceGenomeVariant, true
}

// HasReferenceGenomeVariant returns a boolean if a field has been set.
func (o *MutationAssessor) HasReferenceGenomeVariant() bool {
	if o != nil && !isNil(o.ReferenceGenomeVariant) {
		return true
	}

	return false
}

// SetReferenceGenomeVariant gets a reference to the given string and assigns it to the ReferenceGenomeVariant field.
func (o *MutationAssessor) SetReferenceGenomeVariant(v string) {
	o.ReferenceGenomeVariant = &v
}

// GetReferenceGenomeVariantType returns the ReferenceGenomeVariantType field value if set, zero value otherwise.
func (o *MutationAssessor) GetReferenceGenomeVariantType() string {
	if o == nil || isNil(o.ReferenceGenomeVariantType) {
		var ret string
		return ret
	}
	return *o.ReferenceGenomeVariantType
}

// GetReferenceGenomeVariantTypeOk returns a tuple with the ReferenceGenomeVariantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetReferenceGenomeVariantTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceGenomeVariantType) {
    return nil, false
	}
	return o.ReferenceGenomeVariantType, true
}

// HasReferenceGenomeVariantType returns a boolean if a field has been set.
func (o *MutationAssessor) HasReferenceGenomeVariantType() bool {
	if o != nil && !isNil(o.ReferenceGenomeVariantType) {
		return true
	}

	return false
}

// SetReferenceGenomeVariantType gets a reference to the given string and assigns it to the ReferenceGenomeVariantType field.
func (o *MutationAssessor) SetReferenceGenomeVariantType(v string) {
	o.ReferenceGenomeVariantType = &v
}

// GetRefseqId returns the RefseqId field value if set, zero value otherwise.
func (o *MutationAssessor) GetRefseqId() string {
	if o == nil || isNil(o.RefseqId) {
		var ret string
		return ret
	}
	return *o.RefseqId
}

// GetRefseqIdOk returns a tuple with the RefseqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetRefseqIdOk() (*string, bool) {
	if o == nil || isNil(o.RefseqId) {
    return nil, false
	}
	return o.RefseqId, true
}

// HasRefseqId returns a boolean if a field has been set.
func (o *MutationAssessor) HasRefseqId() bool {
	if o != nil && !isNil(o.RefseqId) {
		return true
	}

	return false
}

// SetRefseqId gets a reference to the given string and assigns it to the RefseqId field.
func (o *MutationAssessor) SetRefseqId(v string) {
	o.RefseqId = &v
}

// GetRefseqPosition returns the RefseqPosition field value if set, zero value otherwise.
func (o *MutationAssessor) GetRefseqPosition() int32 {
	if o == nil || isNil(o.RefseqPosition) {
		var ret int32
		return ret
	}
	return *o.RefseqPosition
}

// GetRefseqPositionOk returns a tuple with the RefseqPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetRefseqPositionOk() (*int32, bool) {
	if o == nil || isNil(o.RefseqPosition) {
    return nil, false
	}
	return o.RefseqPosition, true
}

// HasRefseqPosition returns a boolean if a field has been set.
func (o *MutationAssessor) HasRefseqPosition() bool {
	if o != nil && !isNil(o.RefseqPosition) {
		return true
	}

	return false
}

// SetRefseqPosition gets a reference to the given int32 and assigns it to the RefseqPosition field.
func (o *MutationAssessor) SetRefseqPosition(v int32) {
	o.RefseqPosition = &v
}

// GetRefseqResidue returns the RefseqResidue field value if set, zero value otherwise.
func (o *MutationAssessor) GetRefseqResidue() string {
	if o == nil || isNil(o.RefseqResidue) {
		var ret string
		return ret
	}
	return *o.RefseqResidue
}

// GetRefseqResidueOk returns a tuple with the RefseqResidue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetRefseqResidueOk() (*string, bool) {
	if o == nil || isNil(o.RefseqResidue) {
    return nil, false
	}
	return o.RefseqResidue, true
}

// HasRefseqResidue returns a boolean if a field has been set.
func (o *MutationAssessor) HasRefseqResidue() bool {
	if o != nil && !isNil(o.RefseqResidue) {
		return true
	}

	return false
}

// SetRefseqResidue gets a reference to the given string and assigns it to the RefseqResidue field.
func (o *MutationAssessor) SetRefseqResidue(v string) {
	o.RefseqResidue = &v
}

// GetSnpCount returns the SnpCount field value if set, zero value otherwise.
func (o *MutationAssessor) GetSnpCount() int32 {
	if o == nil || isNil(o.SnpCount) {
		var ret int32
		return ret
	}
	return *o.SnpCount
}

// GetSnpCountOk returns a tuple with the SnpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetSnpCountOk() (*int32, bool) {
	if o == nil || isNil(o.SnpCount) {
    return nil, false
	}
	return o.SnpCount, true
}

// HasSnpCount returns a boolean if a field has been set.
func (o *MutationAssessor) HasSnpCount() bool {
	if o != nil && !isNil(o.SnpCount) {
		return true
	}

	return false
}

// SetSnpCount gets a reference to the given int32 and assigns it to the SnpCount field.
func (o *MutationAssessor) SetSnpCount(v int32) {
	o.SnpCount = &v
}

// GetUniprotId returns the UniprotId field value if set, zero value otherwise.
func (o *MutationAssessor) GetUniprotId() string {
	if o == nil || isNil(o.UniprotId) {
		var ret string
		return ret
	}
	return *o.UniprotId
}

// GetUniprotIdOk returns a tuple with the UniprotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetUniprotIdOk() (*string, bool) {
	if o == nil || isNil(o.UniprotId) {
    return nil, false
	}
	return o.UniprotId, true
}

// HasUniprotId returns a boolean if a field has been set.
func (o *MutationAssessor) HasUniprotId() bool {
	if o != nil && !isNil(o.UniprotId) {
		return true
	}

	return false
}

// SetUniprotId gets a reference to the given string and assigns it to the UniprotId field.
func (o *MutationAssessor) SetUniprotId(v string) {
	o.UniprotId = &v
}

// GetUniprotPosition returns the UniprotPosition field value if set, zero value otherwise.
func (o *MutationAssessor) GetUniprotPosition() int32 {
	if o == nil || isNil(o.UniprotPosition) {
		var ret int32
		return ret
	}
	return *o.UniprotPosition
}

// GetUniprotPositionOk returns a tuple with the UniprotPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetUniprotPositionOk() (*int32, bool) {
	if o == nil || isNil(o.UniprotPosition) {
    return nil, false
	}
	return o.UniprotPosition, true
}

// HasUniprotPosition returns a boolean if a field has been set.
func (o *MutationAssessor) HasUniprotPosition() bool {
	if o != nil && !isNil(o.UniprotPosition) {
		return true
	}

	return false
}

// SetUniprotPosition gets a reference to the given int32 and assigns it to the UniprotPosition field.
func (o *MutationAssessor) SetUniprotPosition(v int32) {
	o.UniprotPosition = &v
}

// GetUniprotResidue returns the UniprotResidue field value if set, zero value otherwise.
func (o *MutationAssessor) GetUniprotResidue() string {
	if o == nil || isNil(o.UniprotResidue) {
		var ret string
		return ret
	}
	return *o.UniprotResidue
}

// GetUniprotResidueOk returns a tuple with the UniprotResidue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetUniprotResidueOk() (*string, bool) {
	if o == nil || isNil(o.UniprotResidue) {
    return nil, false
	}
	return o.UniprotResidue, true
}

// HasUniprotResidue returns a boolean if a field has been set.
func (o *MutationAssessor) HasUniprotResidue() bool {
	if o != nil && !isNil(o.UniprotResidue) {
		return true
	}

	return false
}

// SetUniprotResidue gets a reference to the given string and assigns it to the UniprotResidue field.
func (o *MutationAssessor) SetUniprotResidue(v string) {
	o.UniprotResidue = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *MutationAssessor) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
    return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *MutationAssessor) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *MutationAssessor) SetVariant(v string) {
	o.Variant = &v
}

// GetVariantConservationScore returns the VariantConservationScore field value if set, zero value otherwise.
func (o *MutationAssessor) GetVariantConservationScore() float64 {
	if o == nil || isNil(o.VariantConservationScore) {
		var ret float64
		return ret
	}
	return *o.VariantConservationScore
}

// GetVariantConservationScoreOk returns a tuple with the VariantConservationScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetVariantConservationScoreOk() (*float64, bool) {
	if o == nil || isNil(o.VariantConservationScore) {
    return nil, false
	}
	return o.VariantConservationScore, true
}

// HasVariantConservationScore returns a boolean if a field has been set.
func (o *MutationAssessor) HasVariantConservationScore() bool {
	if o != nil && !isNil(o.VariantConservationScore) {
		return true
	}

	return false
}

// SetVariantConservationScore gets a reference to the given float64 and assigns it to the VariantConservationScore field.
func (o *MutationAssessor) SetVariantConservationScore(v float64) {
	o.VariantConservationScore = &v
}

// GetVariantSpecificityScore returns the VariantSpecificityScore field value if set, zero value otherwise.
func (o *MutationAssessor) GetVariantSpecificityScore() float64 {
	if o == nil || isNil(o.VariantSpecificityScore) {
		var ret float64
		return ret
	}
	return *o.VariantSpecificityScore
}

// GetVariantSpecificityScoreOk returns a tuple with the VariantSpecificityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutationAssessor) GetVariantSpecificityScoreOk() (*float64, bool) {
	if o == nil || isNil(o.VariantSpecificityScore) {
    return nil, false
	}
	return o.VariantSpecificityScore, true
}

// HasVariantSpecificityScore returns a boolean if a field has been set.
func (o *MutationAssessor) HasVariantSpecificityScore() bool {
	if o != nil && !isNil(o.VariantSpecificityScore) {
		return true
	}

	return false
}

// SetVariantSpecificityScore gets a reference to the given float64 and assigns it to the VariantSpecificityScore field.
func (o *MutationAssessor) SetVariantSpecificityScore(v float64) {
	o.VariantSpecificityScore = &v
}

func (o MutationAssessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CodonStartPosition) {
		toSerialize["codonStartPosition"] = o.CodonStartPosition
	}
	if !isNil(o.CosmicCount) {
		toSerialize["cosmicCount"] = o.CosmicCount
	}
	if !isNil(o.FunctionalImpact) {
		toSerialize["functionalImpact"] = o.FunctionalImpact
	}
	if !isNil(o.FunctionalImpactScore) {
		toSerialize["functionalImpactScore"] = o.FunctionalImpactScore
	}
	if !isNil(o.Hgvs) {
		toSerialize["hgvs"] = o.Hgvs
	}
	if !isNil(o.HugoSymbol) {
		toSerialize["hugoSymbol"] = o.HugoSymbol
	}
	if true {
		toSerialize["input"] = o.Input
	}
	if !isNil(o.MappingIssue) {
		toSerialize["mappingIssue"] = o.MappingIssue
	}
	if !isNil(o.MsaGaps) {
		toSerialize["msaGaps"] = o.MsaGaps
	}
	if !isNil(o.MsaHeight) {
		toSerialize["msaHeight"] = o.MsaHeight
	}
	if !isNil(o.MsaLink) {
		toSerialize["msaLink"] = o.MsaLink
	}
	if !isNil(o.PdbLink) {
		toSerialize["pdbLink"] = o.PdbLink
	}
	if !isNil(o.ReferenceGenomeVariant) {
		toSerialize["referenceGenomeVariant"] = o.ReferenceGenomeVariant
	}
	if !isNil(o.ReferenceGenomeVariantType) {
		toSerialize["referenceGenomeVariantType"] = o.ReferenceGenomeVariantType
	}
	if !isNil(o.RefseqId) {
		toSerialize["refseqId"] = o.RefseqId
	}
	if !isNil(o.RefseqPosition) {
		toSerialize["refseqPosition"] = o.RefseqPosition
	}
	if !isNil(o.RefseqResidue) {
		toSerialize["refseqResidue"] = o.RefseqResidue
	}
	if !isNil(o.SnpCount) {
		toSerialize["snpCount"] = o.SnpCount
	}
	if !isNil(o.UniprotId) {
		toSerialize["uniprotId"] = o.UniprotId
	}
	if !isNil(o.UniprotPosition) {
		toSerialize["uniprotPosition"] = o.UniprotPosition
	}
	if !isNil(o.UniprotResidue) {
		toSerialize["uniprotResidue"] = o.UniprotResidue
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !isNil(o.VariantConservationScore) {
		toSerialize["variantConservationScore"] = o.VariantConservationScore
	}
	if !isNil(o.VariantSpecificityScore) {
		toSerialize["variantSpecificityScore"] = o.VariantSpecificityScore
	}
	return json.Marshal(toSerialize)
}

type NullableMutationAssessor struct {
	value *MutationAssessor
	isSet bool
}

func (v NullableMutationAssessor) Get() *MutationAssessor {
	return v.value
}

func (v *NullableMutationAssessor) Set(val *MutationAssessor) {
	v.value = val
	v.isSet = true
}

func (v NullableMutationAssessor) IsSet() bool {
	return v.isSet
}

func (v *NullableMutationAssessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutationAssessor(val *MutationAssessor) *NullableMutationAssessor {
	return &NullableMutationAssessor{value: val, isSet: true}
}

func (v NullableMutationAssessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutationAssessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


