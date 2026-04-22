# TranscriptConsequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AminoAcids** | Pointer to **string** | Amino acids | [optional] 
**Canonical** | Pointer to **string** | Canonical transcript indicator | [optional] 
**Codons** | Pointer to **string** | Codons | [optional] 
**ConsequenceTerms** | Pointer to **[]string** | List of consequence terms | [optional] 
**Exon** | Pointer to **string** |  | [optional] 
**GeneId** | Pointer to **string** | Ensembl gene id | [optional] 
**GeneSymbol** | Pointer to **string** | Hugo gene symbol | [optional] 
**HgncId** | Pointer to **string** | HGNC id | [optional] 
**Hgvsc** | Pointer to **string** | HGVSc | [optional] 
**Hgvsg** | Pointer to **string** | HGVSg | [optional] 
**Hgvsp** | Pointer to **string** | HGVSp | [optional] 
**PolyphenPrediction** | Pointer to **string** | Polyphen Prediction | [optional] 
**PolyphenScore** | Pointer to **float64** | Polyphen Score | [optional] 
**ProteinEnd** | Pointer to **int32** | Protein end position | [optional] 
**ProteinId** | Pointer to **string** | Ensembl protein id | [optional] 
**ProteinStart** | Pointer to **int32** | Protein start position | [optional] 
**RefseqTranscriptIds** | Pointer to **[]string** | List of RefSeq transcript ids | [optional] 
**SiftPrediction** | Pointer to **string** | Sift Prediction | [optional] 
**SiftScore** | Pointer to **float64** | Sift Score | [optional] 
**TranscriptId** | **string** | Ensembl transcript id | 
**UniprotId** | Pointer to **string** |  | [optional] 
**VariantAllele** | Pointer to **string** | Variant allele | [optional] 
**AllEffects** | Pointer to **string** | All effects (VEP) | [optional] 
**Biotype** | Pointer to **string** | Biotype of the transcript | [optional] 
**Ccds** | Pointer to **string** | CCDS transcript identifier | [optional] 
**CdnaPosition** | Pointer to **string** | cDNA position | [optional] 
**CdsPosition** | Pointer to **string** | CDS position | [optional] 
**ClinSig** | Pointer to **[]string** | ClinVar clinical significance | [optional] 
**Distance** | Pointer to **int32** | Distance to transcript | [optional] 
**Domains** | Pointer to **[]map[string]string** | Overlapping protein domains | [optional] 
**GenePheno** | Pointer to **int32** | Indicates if gene is associated with a phenotype | [optional] 
**HgvsOffset** | Pointer to **int32** | HGVS offset | [optional] 
**HighInfPos** | Pointer to **string** | High information position in motif | [optional] 
**Impact** | Pointer to **string** | VEP predicted impact | [optional] 
**Intron** | Pointer to **string** | Intron number | [optional] 
**Minimised** | Pointer to **int32** | Allele minimisation flag | [optional] 
**MotifName** | Pointer to **string** | Overlapping regulatory motif name | [optional] 
**MotifPos** | Pointer to **int32** | Position in motif | [optional] 
**MotifScoreChange** | Pointer to **float64** | Motif score change | [optional] 
**Pheno** | Pointer to **[]int32** | Variant associated with a phenotype | [optional] 
**Pick** | Pointer to **int32** | VEP pick flag (chosen transcript) | [optional] 
**Pubmed** | Pointer to **[]int32** | PubMed IDs for publications | [optional] 
**Somatic** | Pointer to **[]int32** | Somatic status flags | [optional] 
**Swissprot** | Pointer to **[]string** | UniProtKB/Swiss-Prot accessions | [optional] 
**SymbolSource** | Pointer to **string** | Source of gene symbol | [optional] 
**Trembl** | Pointer to **[]string** | UniProtKB/TrEMBL accessions | [optional] 
**Tsl** | Pointer to **int32** | Transcript support level | [optional] 
**Uniparc** | Pointer to **string** | UniParc identifier | [optional] 
**VariantClass** | Pointer to **string** | Sequence Ontology variant class | [optional] 

## Methods

### NewTranscriptConsequence

`func NewTranscriptConsequence(transcriptId string, ) *TranscriptConsequence`

NewTranscriptConsequence instantiates a new TranscriptConsequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptConsequenceWithDefaults

`func NewTranscriptConsequenceWithDefaults() *TranscriptConsequence`

NewTranscriptConsequenceWithDefaults instantiates a new TranscriptConsequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAminoAcids

`func (o *TranscriptConsequence) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *TranscriptConsequence) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *TranscriptConsequence) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *TranscriptConsequence) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetCanonical

`func (o *TranscriptConsequence) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TranscriptConsequence) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TranscriptConsequence) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *TranscriptConsequence) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetCodons

`func (o *TranscriptConsequence) GetCodons() string`

GetCodons returns the Codons field if non-nil, zero value otherwise.

### GetCodonsOk

`func (o *TranscriptConsequence) GetCodonsOk() (*string, bool)`

GetCodonsOk returns a tuple with the Codons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodons

`func (o *TranscriptConsequence) SetCodons(v string)`

SetCodons sets Codons field to given value.

### HasCodons

`func (o *TranscriptConsequence) HasCodons() bool`

HasCodons returns a boolean if a field has been set.

### GetConsequenceTerms

`func (o *TranscriptConsequence) GetConsequenceTerms() []string`

GetConsequenceTerms returns the ConsequenceTerms field if non-nil, zero value otherwise.

### GetConsequenceTermsOk

`func (o *TranscriptConsequence) GetConsequenceTermsOk() (*[]string, bool)`

GetConsequenceTermsOk returns a tuple with the ConsequenceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequenceTerms

`func (o *TranscriptConsequence) SetConsequenceTerms(v []string)`

SetConsequenceTerms sets ConsequenceTerms field to given value.

### HasConsequenceTerms

`func (o *TranscriptConsequence) HasConsequenceTerms() bool`

HasConsequenceTerms returns a boolean if a field has been set.

### GetExon

`func (o *TranscriptConsequence) GetExon() string`

GetExon returns the Exon field if non-nil, zero value otherwise.

### GetExonOk

`func (o *TranscriptConsequence) GetExonOk() (*string, bool)`

GetExonOk returns a tuple with the Exon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExon

`func (o *TranscriptConsequence) SetExon(v string)`

SetExon sets Exon field to given value.

### HasExon

`func (o *TranscriptConsequence) HasExon() bool`

HasExon returns a boolean if a field has been set.

### GetGeneId

`func (o *TranscriptConsequence) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *TranscriptConsequence) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *TranscriptConsequence) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *TranscriptConsequence) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetGeneSymbol

`func (o *TranscriptConsequence) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *TranscriptConsequence) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *TranscriptConsequence) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *TranscriptConsequence) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### GetHgncId

`func (o *TranscriptConsequence) GetHgncId() string`

GetHgncId returns the HgncId field if non-nil, zero value otherwise.

### GetHgncIdOk

`func (o *TranscriptConsequence) GetHgncIdOk() (*string, bool)`

GetHgncIdOk returns a tuple with the HgncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgncId

`func (o *TranscriptConsequence) SetHgncId(v string)`

SetHgncId sets HgncId field to given value.

### HasHgncId

`func (o *TranscriptConsequence) HasHgncId() bool`

HasHgncId returns a boolean if a field has been set.

### GetHgvsc

`func (o *TranscriptConsequence) GetHgvsc() string`

GetHgvsc returns the Hgvsc field if non-nil, zero value otherwise.

### GetHgvscOk

`func (o *TranscriptConsequence) GetHgvscOk() (*string, bool)`

GetHgvscOk returns a tuple with the Hgvsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsc

`func (o *TranscriptConsequence) SetHgvsc(v string)`

SetHgvsc sets Hgvsc field to given value.

### HasHgvsc

`func (o *TranscriptConsequence) HasHgvsc() bool`

HasHgvsc returns a boolean if a field has been set.

### GetHgvsg

`func (o *TranscriptConsequence) GetHgvsg() string`

GetHgvsg returns the Hgvsg field if non-nil, zero value otherwise.

### GetHgvsgOk

`func (o *TranscriptConsequence) GetHgvsgOk() (*string, bool)`

GetHgvsgOk returns a tuple with the Hgvsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsg

`func (o *TranscriptConsequence) SetHgvsg(v string)`

SetHgvsg sets Hgvsg field to given value.

### HasHgvsg

`func (o *TranscriptConsequence) HasHgvsg() bool`

HasHgvsg returns a boolean if a field has been set.

### GetHgvsp

`func (o *TranscriptConsequence) GetHgvsp() string`

GetHgvsp returns the Hgvsp field if non-nil, zero value otherwise.

### GetHgvspOk

`func (o *TranscriptConsequence) GetHgvspOk() (*string, bool)`

GetHgvspOk returns a tuple with the Hgvsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsp

`func (o *TranscriptConsequence) SetHgvsp(v string)`

SetHgvsp sets Hgvsp field to given value.

### HasHgvsp

`func (o *TranscriptConsequence) HasHgvsp() bool`

HasHgvsp returns a boolean if a field has been set.

### GetPolyphenPrediction

`func (o *TranscriptConsequence) GetPolyphenPrediction() string`

GetPolyphenPrediction returns the PolyphenPrediction field if non-nil, zero value otherwise.

### GetPolyphenPredictionOk

`func (o *TranscriptConsequence) GetPolyphenPredictionOk() (*string, bool)`

GetPolyphenPredictionOk returns a tuple with the PolyphenPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenPrediction

`func (o *TranscriptConsequence) SetPolyphenPrediction(v string)`

SetPolyphenPrediction sets PolyphenPrediction field to given value.

### HasPolyphenPrediction

`func (o *TranscriptConsequence) HasPolyphenPrediction() bool`

HasPolyphenPrediction returns a boolean if a field has been set.

### GetPolyphenScore

`func (o *TranscriptConsequence) GetPolyphenScore() float64`

GetPolyphenScore returns the PolyphenScore field if non-nil, zero value otherwise.

### GetPolyphenScoreOk

`func (o *TranscriptConsequence) GetPolyphenScoreOk() (*float64, bool)`

GetPolyphenScoreOk returns a tuple with the PolyphenScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenScore

`func (o *TranscriptConsequence) SetPolyphenScore(v float64)`

SetPolyphenScore sets PolyphenScore field to given value.

### HasPolyphenScore

`func (o *TranscriptConsequence) HasPolyphenScore() bool`

HasPolyphenScore returns a boolean if a field has been set.

### GetProteinEnd

`func (o *TranscriptConsequence) GetProteinEnd() int32`

GetProteinEnd returns the ProteinEnd field if non-nil, zero value otherwise.

### GetProteinEndOk

`func (o *TranscriptConsequence) GetProteinEndOk() (*int32, bool)`

GetProteinEndOk returns a tuple with the ProteinEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinEnd

`func (o *TranscriptConsequence) SetProteinEnd(v int32)`

SetProteinEnd sets ProteinEnd field to given value.

### HasProteinEnd

`func (o *TranscriptConsequence) HasProteinEnd() bool`

HasProteinEnd returns a boolean if a field has been set.

### GetProteinId

`func (o *TranscriptConsequence) GetProteinId() string`

GetProteinId returns the ProteinId field if non-nil, zero value otherwise.

### GetProteinIdOk

`func (o *TranscriptConsequence) GetProteinIdOk() (*string, bool)`

GetProteinIdOk returns a tuple with the ProteinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinId

`func (o *TranscriptConsequence) SetProteinId(v string)`

SetProteinId sets ProteinId field to given value.

### HasProteinId

`func (o *TranscriptConsequence) HasProteinId() bool`

HasProteinId returns a boolean if a field has been set.

### GetProteinStart

`func (o *TranscriptConsequence) GetProteinStart() int32`

GetProteinStart returns the ProteinStart field if non-nil, zero value otherwise.

### GetProteinStartOk

`func (o *TranscriptConsequence) GetProteinStartOk() (*int32, bool)`

GetProteinStartOk returns a tuple with the ProteinStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinStart

`func (o *TranscriptConsequence) SetProteinStart(v int32)`

SetProteinStart sets ProteinStart field to given value.

### HasProteinStart

`func (o *TranscriptConsequence) HasProteinStart() bool`

HasProteinStart returns a boolean if a field has been set.

### GetRefseqTranscriptIds

`func (o *TranscriptConsequence) GetRefseqTranscriptIds() []string`

GetRefseqTranscriptIds returns the RefseqTranscriptIds field if non-nil, zero value otherwise.

### GetRefseqTranscriptIdsOk

`func (o *TranscriptConsequence) GetRefseqTranscriptIdsOk() (*[]string, bool)`

GetRefseqTranscriptIdsOk returns a tuple with the RefseqTranscriptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqTranscriptIds

`func (o *TranscriptConsequence) SetRefseqTranscriptIds(v []string)`

SetRefseqTranscriptIds sets RefseqTranscriptIds field to given value.

### HasRefseqTranscriptIds

`func (o *TranscriptConsequence) HasRefseqTranscriptIds() bool`

HasRefseqTranscriptIds returns a boolean if a field has been set.

### GetSiftPrediction

`func (o *TranscriptConsequence) GetSiftPrediction() string`

GetSiftPrediction returns the SiftPrediction field if non-nil, zero value otherwise.

### GetSiftPredictionOk

`func (o *TranscriptConsequence) GetSiftPredictionOk() (*string, bool)`

GetSiftPredictionOk returns a tuple with the SiftPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftPrediction

`func (o *TranscriptConsequence) SetSiftPrediction(v string)`

SetSiftPrediction sets SiftPrediction field to given value.

### HasSiftPrediction

`func (o *TranscriptConsequence) HasSiftPrediction() bool`

HasSiftPrediction returns a boolean if a field has been set.

### GetSiftScore

`func (o *TranscriptConsequence) GetSiftScore() float64`

GetSiftScore returns the SiftScore field if non-nil, zero value otherwise.

### GetSiftScoreOk

`func (o *TranscriptConsequence) GetSiftScoreOk() (*float64, bool)`

GetSiftScoreOk returns a tuple with the SiftScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftScore

`func (o *TranscriptConsequence) SetSiftScore(v float64)`

SetSiftScore sets SiftScore field to given value.

### HasSiftScore

`func (o *TranscriptConsequence) HasSiftScore() bool`

HasSiftScore returns a boolean if a field has been set.

### GetTranscriptId

`func (o *TranscriptConsequence) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *TranscriptConsequence) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *TranscriptConsequence) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.


### GetUniprotId

`func (o *TranscriptConsequence) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *TranscriptConsequence) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *TranscriptConsequence) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *TranscriptConsequence) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.

### GetVariantAllele

`func (o *TranscriptConsequence) GetVariantAllele() string`

GetVariantAllele returns the VariantAllele field if non-nil, zero value otherwise.

### GetVariantAlleleOk

`func (o *TranscriptConsequence) GetVariantAlleleOk() (*string, bool)`

GetVariantAlleleOk returns a tuple with the VariantAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantAllele

`func (o *TranscriptConsequence) SetVariantAllele(v string)`

SetVariantAllele sets VariantAllele field to given value.

### HasVariantAllele

`func (o *TranscriptConsequence) HasVariantAllele() bool`

HasVariantAllele returns a boolean if a field has been set.

### GetAllEffects

`func (o *TranscriptConsequence) GetAllEffects() string`

GetAllEffects returns the AllEffects field if non-nil, zero value otherwise.

### GetAllEffectsOk

`func (o *TranscriptConsequence) GetAllEffectsOk() (*string, bool)`

GetAllEffectsOk returns a tuple with the AllEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEffects

`func (o *TranscriptConsequence) SetAllEffects(v string)`

SetAllEffects sets AllEffects field to given value.

### HasAllEffects

`func (o *TranscriptConsequence) HasAllEffects() bool`

HasAllEffects returns a boolean if a field has been set.

### GetBiotype

`func (o *TranscriptConsequence) GetBiotype() string`

GetBiotype returns the Biotype field if non-nil, zero value otherwise.

### GetBiotypeOk

`func (o *TranscriptConsequence) GetBiotypeOk() (*string, bool)`

GetBiotypeOk returns a tuple with the Biotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiotype

`func (o *TranscriptConsequence) SetBiotype(v string)`

SetBiotype sets Biotype field to given value.

### HasBiotype

`func (o *TranscriptConsequence) HasBiotype() bool`

HasBiotype returns a boolean if a field has been set.

### GetCcds

`func (o *TranscriptConsequence) GetCcds() string`

GetCcds returns the Ccds field if non-nil, zero value otherwise.

### GetCcdsOk

`func (o *TranscriptConsequence) GetCcdsOk() (*string, bool)`

GetCcdsOk returns a tuple with the Ccds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcds

`func (o *TranscriptConsequence) SetCcds(v string)`

SetCcds sets Ccds field to given value.

### HasCcds

`func (o *TranscriptConsequence) HasCcds() bool`

HasCcds returns a boolean if a field has been set.

### GetCdnaPosition

`func (o *TranscriptConsequence) GetCdnaPosition() string`

GetCdnaPosition returns the CdnaPosition field if non-nil, zero value otherwise.

### GetCdnaPositionOk

`func (o *TranscriptConsequence) GetCdnaPositionOk() (*string, bool)`

GetCdnaPositionOk returns a tuple with the CdnaPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnaPosition

`func (o *TranscriptConsequence) SetCdnaPosition(v string)`

SetCdnaPosition sets CdnaPosition field to given value.

### HasCdnaPosition

`func (o *TranscriptConsequence) HasCdnaPosition() bool`

HasCdnaPosition returns a boolean if a field has been set.

### GetCdsPosition

`func (o *TranscriptConsequence) GetCdsPosition() string`

GetCdsPosition returns the CdsPosition field if non-nil, zero value otherwise.

### GetCdsPositionOk

`func (o *TranscriptConsequence) GetCdsPositionOk() (*string, bool)`

GetCdsPositionOk returns a tuple with the CdsPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdsPosition

`func (o *TranscriptConsequence) SetCdsPosition(v string)`

SetCdsPosition sets CdsPosition field to given value.

### HasCdsPosition

`func (o *TranscriptConsequence) HasCdsPosition() bool`

HasCdsPosition returns a boolean if a field has been set.

### GetClinSig

`func (o *TranscriptConsequence) GetClinSig() []string`

GetClinSig returns the ClinSig field if non-nil, zero value otherwise.

### GetClinSigOk

`func (o *TranscriptConsequence) GetClinSigOk() (*[]string, bool)`

GetClinSigOk returns a tuple with the ClinSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinSig

`func (o *TranscriptConsequence) SetClinSig(v []string)`

SetClinSig sets ClinSig field to given value.

### HasClinSig

`func (o *TranscriptConsequence) HasClinSig() bool`

HasClinSig returns a boolean if a field has been set.

### GetDistance

`func (o *TranscriptConsequence) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *TranscriptConsequence) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *TranscriptConsequence) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *TranscriptConsequence) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetDomains

`func (o *TranscriptConsequence) GetDomains() []map[string]string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *TranscriptConsequence) GetDomainsOk() (*[]map[string]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *TranscriptConsequence) SetDomains(v []map[string]string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *TranscriptConsequence) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGenePheno

`func (o *TranscriptConsequence) GetGenePheno() int32`

GetGenePheno returns the GenePheno field if non-nil, zero value otherwise.

### GetGenePhenoOk

`func (o *TranscriptConsequence) GetGenePhenoOk() (*int32, bool)`

GetGenePhenoOk returns a tuple with the GenePheno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenePheno

`func (o *TranscriptConsequence) SetGenePheno(v int32)`

SetGenePheno sets GenePheno field to given value.

### HasGenePheno

`func (o *TranscriptConsequence) HasGenePheno() bool`

HasGenePheno returns a boolean if a field has been set.

### GetHgvsOffset

`func (o *TranscriptConsequence) GetHgvsOffset() int32`

GetHgvsOffset returns the HgvsOffset field if non-nil, zero value otherwise.

### GetHgvsOffsetOk

`func (o *TranscriptConsequence) GetHgvsOffsetOk() (*int32, bool)`

GetHgvsOffsetOk returns a tuple with the HgvsOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsOffset

`func (o *TranscriptConsequence) SetHgvsOffset(v int32)`

SetHgvsOffset sets HgvsOffset field to given value.

### HasHgvsOffset

`func (o *TranscriptConsequence) HasHgvsOffset() bool`

HasHgvsOffset returns a boolean if a field has been set.

### GetHighInfPos

`func (o *TranscriptConsequence) GetHighInfPos() string`

GetHighInfPos returns the HighInfPos field if non-nil, zero value otherwise.

### GetHighInfPosOk

`func (o *TranscriptConsequence) GetHighInfPosOk() (*string, bool)`

GetHighInfPosOk returns a tuple with the HighInfPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighInfPos

`func (o *TranscriptConsequence) SetHighInfPos(v string)`

SetHighInfPos sets HighInfPos field to given value.

### HasHighInfPos

`func (o *TranscriptConsequence) HasHighInfPos() bool`

HasHighInfPos returns a boolean if a field has been set.

### GetImpact

`func (o *TranscriptConsequence) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *TranscriptConsequence) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *TranscriptConsequence) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *TranscriptConsequence) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetIntron

`func (o *TranscriptConsequence) GetIntron() string`

GetIntron returns the Intron field if non-nil, zero value otherwise.

### GetIntronOk

`func (o *TranscriptConsequence) GetIntronOk() (*string, bool)`

GetIntronOk returns a tuple with the Intron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntron

`func (o *TranscriptConsequence) SetIntron(v string)`

SetIntron sets Intron field to given value.

### HasIntron

`func (o *TranscriptConsequence) HasIntron() bool`

HasIntron returns a boolean if a field has been set.

### GetMinimised

`func (o *TranscriptConsequence) GetMinimised() int32`

GetMinimised returns the Minimised field if non-nil, zero value otherwise.

### GetMinimisedOk

`func (o *TranscriptConsequence) GetMinimisedOk() (*int32, bool)`

GetMinimisedOk returns a tuple with the Minimised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimised

`func (o *TranscriptConsequence) SetMinimised(v int32)`

SetMinimised sets Minimised field to given value.

### HasMinimised

`func (o *TranscriptConsequence) HasMinimised() bool`

HasMinimised returns a boolean if a field has been set.

### GetMotifName

`func (o *TranscriptConsequence) GetMotifName() string`

GetMotifName returns the MotifName field if non-nil, zero value otherwise.

### GetMotifNameOk

`func (o *TranscriptConsequence) GetMotifNameOk() (*string, bool)`

GetMotifNameOk returns a tuple with the MotifName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifName

`func (o *TranscriptConsequence) SetMotifName(v string)`

SetMotifName sets MotifName field to given value.

### HasMotifName

`func (o *TranscriptConsequence) HasMotifName() bool`

HasMotifName returns a boolean if a field has been set.

### GetMotifPos

`func (o *TranscriptConsequence) GetMotifPos() int32`

GetMotifPos returns the MotifPos field if non-nil, zero value otherwise.

### GetMotifPosOk

`func (o *TranscriptConsequence) GetMotifPosOk() (*int32, bool)`

GetMotifPosOk returns a tuple with the MotifPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifPos

`func (o *TranscriptConsequence) SetMotifPos(v int32)`

SetMotifPos sets MotifPos field to given value.

### HasMotifPos

`func (o *TranscriptConsequence) HasMotifPos() bool`

HasMotifPos returns a boolean if a field has been set.

### GetMotifScoreChange

`func (o *TranscriptConsequence) GetMotifScoreChange() float64`

GetMotifScoreChange returns the MotifScoreChange field if non-nil, zero value otherwise.

### GetMotifScoreChangeOk

`func (o *TranscriptConsequence) GetMotifScoreChangeOk() (*float64, bool)`

GetMotifScoreChangeOk returns a tuple with the MotifScoreChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifScoreChange

`func (o *TranscriptConsequence) SetMotifScoreChange(v float64)`

SetMotifScoreChange sets MotifScoreChange field to given value.

### HasMotifScoreChange

`func (o *TranscriptConsequence) HasMotifScoreChange() bool`

HasMotifScoreChange returns a boolean if a field has been set.

### GetPheno

`func (o *TranscriptConsequence) GetPheno() []int32`

GetPheno returns the Pheno field if non-nil, zero value otherwise.

### GetPhenoOk

`func (o *TranscriptConsequence) GetPhenoOk() (*[]int32, bool)`

GetPhenoOk returns a tuple with the Pheno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPheno

`func (o *TranscriptConsequence) SetPheno(v []int32)`

SetPheno sets Pheno field to given value.

### HasPheno

`func (o *TranscriptConsequence) HasPheno() bool`

HasPheno returns a boolean if a field has been set.

### GetPick

`func (o *TranscriptConsequence) GetPick() int32`

GetPick returns the Pick field if non-nil, zero value otherwise.

### GetPickOk

`func (o *TranscriptConsequence) GetPickOk() (*int32, bool)`

GetPickOk returns a tuple with the Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPick

`func (o *TranscriptConsequence) SetPick(v int32)`

SetPick sets Pick field to given value.

### HasPick

`func (o *TranscriptConsequence) HasPick() bool`

HasPick returns a boolean if a field has been set.

### GetPubmed

`func (o *TranscriptConsequence) GetPubmed() []int32`

GetPubmed returns the Pubmed field if non-nil, zero value otherwise.

### GetPubmedOk

`func (o *TranscriptConsequence) GetPubmedOk() (*[]int32, bool)`

GetPubmedOk returns a tuple with the Pubmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubmed

`func (o *TranscriptConsequence) SetPubmed(v []int32)`

SetPubmed sets Pubmed field to given value.

### HasPubmed

`func (o *TranscriptConsequence) HasPubmed() bool`

HasPubmed returns a boolean if a field has been set.

### GetSomatic

`func (o *TranscriptConsequence) GetSomatic() []int32`

GetSomatic returns the Somatic field if non-nil, zero value otherwise.

### GetSomaticOk

`func (o *TranscriptConsequence) GetSomaticOk() (*[]int32, bool)`

GetSomaticOk returns a tuple with the Somatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSomatic

`func (o *TranscriptConsequence) SetSomatic(v []int32)`

SetSomatic sets Somatic field to given value.

### HasSomatic

`func (o *TranscriptConsequence) HasSomatic() bool`

HasSomatic returns a boolean if a field has been set.

### GetSwissprot

`func (o *TranscriptConsequence) GetSwissprot() []string`

GetSwissprot returns the Swissprot field if non-nil, zero value otherwise.

### GetSwissprotOk

`func (o *TranscriptConsequence) GetSwissprotOk() (*[]string, bool)`

GetSwissprotOk returns a tuple with the Swissprot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissprot

`func (o *TranscriptConsequence) SetSwissprot(v []string)`

SetSwissprot sets Swissprot field to given value.

### HasSwissprot

`func (o *TranscriptConsequence) HasSwissprot() bool`

HasSwissprot returns a boolean if a field has been set.

### GetSymbolSource

`func (o *TranscriptConsequence) GetSymbolSource() string`

GetSymbolSource returns the SymbolSource field if non-nil, zero value otherwise.

### GetSymbolSourceOk

`func (o *TranscriptConsequence) GetSymbolSourceOk() (*string, bool)`

GetSymbolSourceOk returns a tuple with the SymbolSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolSource

`func (o *TranscriptConsequence) SetSymbolSource(v string)`

SetSymbolSource sets SymbolSource field to given value.

### HasSymbolSource

`func (o *TranscriptConsequence) HasSymbolSource() bool`

HasSymbolSource returns a boolean if a field has been set.

### GetTrembl

`func (o *TranscriptConsequence) GetTrembl() []string`

GetTrembl returns the Trembl field if non-nil, zero value otherwise.

### GetTremblOk

`func (o *TranscriptConsequence) GetTremblOk() (*[]string, bool)`

GetTremblOk returns a tuple with the Trembl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrembl

`func (o *TranscriptConsequence) SetTrembl(v []string)`

SetTrembl sets Trembl field to given value.

### HasTrembl

`func (o *TranscriptConsequence) HasTrembl() bool`

HasTrembl returns a boolean if a field has been set.

### GetTsl

`func (o *TranscriptConsequence) GetTsl() int32`

GetTsl returns the Tsl field if non-nil, zero value otherwise.

### GetTslOk

`func (o *TranscriptConsequence) GetTslOk() (*int32, bool)`

GetTslOk returns a tuple with the Tsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsl

`func (o *TranscriptConsequence) SetTsl(v int32)`

SetTsl sets Tsl field to given value.

### HasTsl

`func (o *TranscriptConsequence) HasTsl() bool`

HasTsl returns a boolean if a field has been set.

### GetUniparc

`func (o *TranscriptConsequence) GetUniparc() string`

GetUniparc returns the Uniparc field if non-nil, zero value otherwise.

### GetUniparcOk

`func (o *TranscriptConsequence) GetUniparcOk() (*string, bool)`

GetUniparcOk returns a tuple with the Uniparc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniparc

`func (o *TranscriptConsequence) SetUniparc(v string)`

SetUniparc sets Uniparc field to given value.

### HasUniparc

`func (o *TranscriptConsequence) HasUniparc() bool`

HasUniparc returns a boolean if a field has been set.

### GetVariantClass

`func (o *TranscriptConsequence) GetVariantClass() string`

GetVariantClass returns the VariantClass field if non-nil, zero value otherwise.

### GetVariantClassOk

`func (o *TranscriptConsequence) GetVariantClassOk() (*string, bool)`

GetVariantClassOk returns a tuple with the VariantClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantClass

`func (o *TranscriptConsequence) SetVariantClass(v string)`

SetVariantClass sets VariantClass field to given value.

### HasVariantClass

`func (o *TranscriptConsequence) HasVariantClass() bool`

HasVariantClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


