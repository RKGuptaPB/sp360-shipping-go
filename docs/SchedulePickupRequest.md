# SchedulePickupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupFedexRequestPickupSummaryInner**](SchedulePickupFedexRequestPickupSummaryInner.md) | description | [optional] 
**Additionalnotes** | Pointer to **string** | description | [optional] 
**PickupStartTime** | Pointer to **string** | description | [optional] 
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupUPSRequestPickupOptionsInner**](SchedulePickupUPSRequestPickupOptionsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupRequest

`func NewSchedulePickupRequest() *SchedulePickupRequest`

NewSchedulePickupRequest instantiates a new SchedulePickupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupRequestWithDefaults

`func NewSchedulePickupRequestWithDefaults() *SchedulePickupRequest`

NewSchedulePickupRequestWithDefaults instantiates a new SchedulePickupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupRequest) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupRequest) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupRequest) GetPickupSummary() []SchedulePickupFedexRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupRequest) GetPickupSummaryOk() (*[]SchedulePickupFedexRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupRequest) SetPickupSummary(v []SchedulePickupFedexRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetPickupStartTime

`func (o *SchedulePickupRequest) GetPickupStartTime() string`

GetPickupStartTime returns the PickupStartTime field if non-nil, zero value otherwise.

### GetPickupStartTimeOk

`func (o *SchedulePickupRequest) GetPickupStartTimeOk() (*string, bool)`

GetPickupStartTimeOk returns a tuple with the PickupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartTime

`func (o *SchedulePickupRequest) SetPickupStartTime(v string)`

SetPickupStartTime sets PickupStartTime field to given value.

### HasPickupStartTime

`func (o *SchedulePickupRequest) HasPickupStartTime() bool`

HasPickupStartTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupRequest) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupRequest) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupRequest) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupRequest) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupRequest) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupRequest) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupRequest) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupRequest) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupRequest) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupRequest) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupRequest) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupRequest) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


