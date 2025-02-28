/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_internal_api

import (
	"encoding/json"
)

// PostTranslationalModification struct for PostTranslationalModification
type PostTranslationalModification struct {
	EnsemblTranscriptIds []string `json:"ensemblTranscriptIds,omitempty"`
	Position *int32 `json:"position,omitempty"`
	PubmedIds []string `json:"pubmedIds,omitempty"`
	Sequence *string `json:"sequence,omitempty"`
	Type *string `json:"type,omitempty"`
	UniprotAccession *string `json:"uniprotAccession,omitempty"`
	UniprotEntry *string `json:"uniprotEntry,omitempty"`
}

// NewPostTranslationalModification instantiates a new PostTranslationalModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTranslationalModification() *PostTranslationalModification {
	this := PostTranslationalModification{}
	return &this
}

// NewPostTranslationalModificationWithDefaults instantiates a new PostTranslationalModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTranslationalModificationWithDefaults() *PostTranslationalModification {
	this := PostTranslationalModification{}
	return &this
}

// GetEnsemblTranscriptIds returns the EnsemblTranscriptIds field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetEnsemblTranscriptIds() []string {
	if o == nil || isNil(o.EnsemblTranscriptIds) {
		var ret []string
		return ret
	}
	return o.EnsemblTranscriptIds
}

// GetEnsemblTranscriptIdsOk returns a tuple with the EnsemblTranscriptIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetEnsemblTranscriptIdsOk() ([]string, bool) {
	if o == nil || isNil(o.EnsemblTranscriptIds) {
    return nil, false
	}
	return o.EnsemblTranscriptIds, true
}

// HasEnsemblTranscriptIds returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasEnsemblTranscriptIds() bool {
	if o != nil && !isNil(o.EnsemblTranscriptIds) {
		return true
	}

	return false
}

// SetEnsemblTranscriptIds gets a reference to the given []string and assigns it to the EnsemblTranscriptIds field.
func (o *PostTranslationalModification) SetEnsemblTranscriptIds(v []string) {
	o.EnsemblTranscriptIds = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetPosition() int32 {
	if o == nil || isNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetPositionOk() (*int32, bool) {
	if o == nil || isNil(o.Position) {
    return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasPosition() bool {
	if o != nil && !isNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *PostTranslationalModification) SetPosition(v int32) {
	o.Position = &v
}

// GetPubmedIds returns the PubmedIds field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetPubmedIds() []string {
	if o == nil || isNil(o.PubmedIds) {
		var ret []string
		return ret
	}
	return o.PubmedIds
}

// GetPubmedIdsOk returns a tuple with the PubmedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetPubmedIdsOk() ([]string, bool) {
	if o == nil || isNil(o.PubmedIds) {
    return nil, false
	}
	return o.PubmedIds, true
}

// HasPubmedIds returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasPubmedIds() bool {
	if o != nil && !isNil(o.PubmedIds) {
		return true
	}

	return false
}

// SetPubmedIds gets a reference to the given []string and assigns it to the PubmedIds field.
func (o *PostTranslationalModification) SetPubmedIds(v []string) {
	o.PubmedIds = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetSequence() string {
	if o == nil || isNil(o.Sequence) {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetSequenceOk() (*string, bool) {
	if o == nil || isNil(o.Sequence) {
    return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasSequence() bool {
	if o != nil && !isNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *PostTranslationalModification) SetSequence(v string) {
	o.Sequence = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostTranslationalModification) SetType(v string) {
	o.Type = &v
}

// GetUniprotAccession returns the UniprotAccession field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetUniprotAccession() string {
	if o == nil || isNil(o.UniprotAccession) {
		var ret string
		return ret
	}
	return *o.UniprotAccession
}

// GetUniprotAccessionOk returns a tuple with the UniprotAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetUniprotAccessionOk() (*string, bool) {
	if o == nil || isNil(o.UniprotAccession) {
    return nil, false
	}
	return o.UniprotAccession, true
}

// HasUniprotAccession returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasUniprotAccession() bool {
	if o != nil && !isNil(o.UniprotAccession) {
		return true
	}

	return false
}

// SetUniprotAccession gets a reference to the given string and assigns it to the UniprotAccession field.
func (o *PostTranslationalModification) SetUniprotAccession(v string) {
	o.UniprotAccession = &v
}

// GetUniprotEntry returns the UniprotEntry field value if set, zero value otherwise.
func (o *PostTranslationalModification) GetUniprotEntry() string {
	if o == nil || isNil(o.UniprotEntry) {
		var ret string
		return ret
	}
	return *o.UniprotEntry
}

// GetUniprotEntryOk returns a tuple with the UniprotEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTranslationalModification) GetUniprotEntryOk() (*string, bool) {
	if o == nil || isNil(o.UniprotEntry) {
    return nil, false
	}
	return o.UniprotEntry, true
}

// HasUniprotEntry returns a boolean if a field has been set.
func (o *PostTranslationalModification) HasUniprotEntry() bool {
	if o != nil && !isNil(o.UniprotEntry) {
		return true
	}

	return false
}

// SetUniprotEntry gets a reference to the given string and assigns it to the UniprotEntry field.
func (o *PostTranslationalModification) SetUniprotEntry(v string) {
	o.UniprotEntry = &v
}

func (o PostTranslationalModification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnsemblTranscriptIds) {
		toSerialize["ensemblTranscriptIds"] = o.EnsemblTranscriptIds
	}
	if !isNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !isNil(o.PubmedIds) {
		toSerialize["pubmedIds"] = o.PubmedIds
	}
	if !isNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UniprotAccession) {
		toSerialize["uniprotAccession"] = o.UniprotAccession
	}
	if !isNil(o.UniprotEntry) {
		toSerialize["uniprotEntry"] = o.UniprotEntry
	}
	return json.Marshal(toSerialize)
}

type NullablePostTranslationalModification struct {
	value *PostTranslationalModification
	isSet bool
}

func (v NullablePostTranslationalModification) Get() *PostTranslationalModification {
	return v.value
}

func (v *NullablePostTranslationalModification) Set(val *PostTranslationalModification) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTranslationalModification) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTranslationalModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTranslationalModification(val *PostTranslationalModification) *NullablePostTranslationalModification {
	return &NullablePostTranslationalModification{value: val, isSet: true}
}

func (v NullablePostTranslationalModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTranslationalModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


