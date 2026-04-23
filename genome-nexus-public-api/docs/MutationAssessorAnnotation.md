# MutationAssessorAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**MutationAssessor**](MutationAssessor.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**FunctionalImpactPrediction** | Pointer to **string** | Functional impact prediction | [optional] 
**FunctionalImpactScore** | Pointer to **float64** | Functional impact score | [optional] 
**HgvspShort** | Pointer to **string** | HGVSp short notation | [optional] 
**Mav** | Pointer to **int32** | MAV score | [optional] 
**Msa** | Pointer to **string** | MSA identifier | [optional] 
**Sv** | Pointer to **int32** | SV score | [optional] 
**UniprotId** | Pointer to **string** | UniProt identifier | [optional] 

## Methods

### NewMutationAssessorAnnotation

`func NewMutationAssessorAnnotation() *MutationAssessorAnnotation`

NewMutationAssessorAnnotation instantiates a new MutationAssessorAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutationAssessorAnnotationWithDefaults

`func NewMutationAssessorAnnotationWithDefaults() *MutationAssessorAnnotation`

NewMutationAssessorAnnotationWithDefaults instantiates a new MutationAssessorAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *MutationAssessorAnnotation) GetAnnotation() MutationAssessor`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *MutationAssessorAnnotation) GetAnnotationOk() (*MutationAssessor, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *MutationAssessorAnnotation) SetAnnotation(v MutationAssessor)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *MutationAssessorAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetLicense

`func (o *MutationAssessorAnnotation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MutationAssessorAnnotation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MutationAssessorAnnotation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *MutationAssessorAnnotation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetFunctionalImpactPrediction

`func (o *MutationAssessorAnnotation) GetFunctionalImpactPrediction() string`

GetFunctionalImpactPrediction returns the FunctionalImpactPrediction field if non-nil, zero value otherwise.

### GetFunctionalImpactPredictionOk

`func (o *MutationAssessorAnnotation) GetFunctionalImpactPredictionOk() (*string, bool)`

GetFunctionalImpactPredictionOk returns a tuple with the FunctionalImpactPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalImpactPrediction

`func (o *MutationAssessorAnnotation) SetFunctionalImpactPrediction(v string)`

SetFunctionalImpactPrediction sets FunctionalImpactPrediction field to given value.

### HasFunctionalImpactPrediction

`func (o *MutationAssessorAnnotation) HasFunctionalImpactPrediction() bool`

HasFunctionalImpactPrediction returns a boolean if a field has been set.

### GetFunctionalImpactScore

`func (o *MutationAssessorAnnotation) GetFunctionalImpactScore() float64`

GetFunctionalImpactScore returns the FunctionalImpactScore field if non-nil, zero value otherwise.

### GetFunctionalImpactScoreOk

`func (o *MutationAssessorAnnotation) GetFunctionalImpactScoreOk() (*float64, bool)`

GetFunctionalImpactScoreOk returns a tuple with the FunctionalImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalImpactScore

`func (o *MutationAssessorAnnotation) SetFunctionalImpactScore(v float64)`

SetFunctionalImpactScore sets FunctionalImpactScore field to given value.

### HasFunctionalImpactScore

`func (o *MutationAssessorAnnotation) HasFunctionalImpactScore() bool`

HasFunctionalImpactScore returns a boolean if a field has been set.

### GetHgvspShort

`func (o *MutationAssessorAnnotation) GetHgvspShort() string`

GetHgvspShort returns the HgvspShort field if non-nil, zero value otherwise.

### GetHgvspShortOk

`func (o *MutationAssessorAnnotation) GetHgvspShortOk() (*string, bool)`

GetHgvspShortOk returns a tuple with the HgvspShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvspShort

`func (o *MutationAssessorAnnotation) SetHgvspShort(v string)`

SetHgvspShort sets HgvspShort field to given value.

### HasHgvspShort

`func (o *MutationAssessorAnnotation) HasHgvspShort() bool`

HasHgvspShort returns a boolean if a field has been set.

### GetMav

`func (o *MutationAssessorAnnotation) GetMav() int32`

GetMav returns the Mav field if non-nil, zero value otherwise.

### GetMavOk

`func (o *MutationAssessorAnnotation) GetMavOk() (*int32, bool)`

GetMavOk returns a tuple with the Mav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMav

`func (o *MutationAssessorAnnotation) SetMav(v int32)`

SetMav sets Mav field to given value.

### HasMav

`func (o *MutationAssessorAnnotation) HasMav() bool`

HasMav returns a boolean if a field has been set.

### GetMsa

`func (o *MutationAssessorAnnotation) GetMsa() string`

GetMsa returns the Msa field if non-nil, zero value otherwise.

### GetMsaOk

`func (o *MutationAssessorAnnotation) GetMsaOk() (*string, bool)`

GetMsaOk returns a tuple with the Msa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsa

`func (o *MutationAssessorAnnotation) SetMsa(v string)`

SetMsa sets Msa field to given value.

### HasMsa

`func (o *MutationAssessorAnnotation) HasMsa() bool`

HasMsa returns a boolean if a field has been set.

### GetSv

`func (o *MutationAssessorAnnotation) GetSv() int32`

GetSv returns the Sv field if non-nil, zero value otherwise.

### GetSvOk

`func (o *MutationAssessorAnnotation) GetSvOk() (*int32, bool)`

GetSvOk returns a tuple with the Sv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSv

`func (o *MutationAssessorAnnotation) SetSv(v int32)`

SetSv sets Sv field to given value.

### HasSv

`func (o *MutationAssessorAnnotation) HasSv() bool`

HasSv returns a boolean if a field has been set.

### GetUniprotId

`func (o *MutationAssessorAnnotation) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *MutationAssessorAnnotation) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *MutationAssessorAnnotation) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *MutationAssessorAnnotation) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


