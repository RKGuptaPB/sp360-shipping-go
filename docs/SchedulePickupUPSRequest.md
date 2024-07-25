# SchedulePickupUPSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupUPSRequestPickupSummaryInner**](SchedulePickupUPSRequestPickupSummaryInner.md) | description | [optional] 
**Additionalnotes** | Pointer to **string** | description | [optional] 
**PickupStartTime** | Pointer to **string** | description | [optional] 
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupUPSRequestPickupOptionsInner**](SchedulePickupUPSRequestPickupOptionsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupUPSRequest

`func NewSchedulePickupUPSRequest() *SchedulePickupUPSRequest`

NewSchedulePickupUPSRequest instantiates a new SchedulePickupUPSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSRequestWithDefaults

`func NewSchedulePickupUPSRequestWithDefaults() *SchedulePickupUPSRequest`

NewSchedulePickupUPSRequestWithDefaults instantiates a new SchedulePickupUPSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUPSRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUPSRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUPSRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupUPSRequest) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupUPSRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUPSRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUPSRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupUPSRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupUPSRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUPSRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUPSRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupUPSRequest) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUPSRequest) GetPickupSummary() []SchedulePickupUPSRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUPSRequest) GetPickupSummaryOk() (*[]SchedulePickupUPSRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUPSRequest) SetPickupSummary(v []SchedulePickupUPSRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUPSRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupUPSRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupUPSRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupUPSRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupUPSRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetPickupStartTime

`func (o *SchedulePickupUPSRequest) GetPickupStartTime() string`

GetPickupStartTime returns the PickupStartTime field if non-nil, zero value otherwise.

### GetPickupStartTimeOk

`func (o *SchedulePickupUPSRequest) GetPickupStartTimeOk() (*string, bool)`

GetPickupStartTimeOk returns a tuple with the PickupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartTime

`func (o *SchedulePickupUPSRequest) SetPickupStartTime(v string)`

SetPickupStartTime sets PickupStartTime field to given value.

### HasPickupStartTime

`func (o *SchedulePickupUPSRequest) HasPickupStartTime() bool`

HasPickupStartTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupUPSRequest) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupUPSRequest) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupUPSRequest) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupUPSRequest) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupUPSRequest) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupUPSRequest) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupUPSRequest) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupUPSRequest) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupUPSRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupUPSRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupUPSRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupUPSRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupUPSRequest) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupUPSRequest) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupUPSRequest) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupUPSRequest) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


