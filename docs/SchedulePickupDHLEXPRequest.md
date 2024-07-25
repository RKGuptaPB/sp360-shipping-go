# SchedulePickupDHLEXPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupSummaryInner**](SchedulePickupDHLEXPRequestPickupSummaryInner.md) | description | [optional] 
**Additionalnotes** | Pointer to **string** | description | [optional] 
**PickupStartTime** | Pointer to **string** | description | [optional] 
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupOptionsInner**](SchedulePickupDHLEXPRequestPickupOptionsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupDHLEXPRequest

`func NewSchedulePickupDHLEXPRequest() *SchedulePickupDHLEXPRequest`

NewSchedulePickupDHLEXPRequest instantiates a new SchedulePickupDHLEXPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestWithDefaults

`func NewSchedulePickupDHLEXPRequestWithDefaults() *SchedulePickupDHLEXPRequest`

NewSchedulePickupDHLEXPRequestWithDefaults instantiates a new SchedulePickupDHLEXPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupDHLEXPRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupDHLEXPRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupDHLEXPRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupDHLEXPRequest) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupDHLEXPRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupDHLEXPRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupDHLEXPRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupDHLEXPRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupDHLEXPRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupDHLEXPRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupDHLEXPRequest) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupDHLEXPRequest) GetPickupSummary() []SchedulePickupDHLEXPRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupSummaryOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupDHLEXPRequest) SetPickupSummary(v []SchedulePickupDHLEXPRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupDHLEXPRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupDHLEXPRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetPickupStartTime

`func (o *SchedulePickupDHLEXPRequest) GetPickupStartTime() string`

GetPickupStartTime returns the PickupStartTime field if non-nil, zero value otherwise.

### GetPickupStartTimeOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupStartTimeOk() (*string, bool)`

GetPickupStartTimeOk returns a tuple with the PickupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartTime

`func (o *SchedulePickupDHLEXPRequest) SetPickupStartTime(v string)`

SetPickupStartTime sets PickupStartTime field to given value.

### HasPickupStartTime

`func (o *SchedulePickupDHLEXPRequest) HasPickupStartTime() bool`

HasPickupStartTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupDHLEXPRequest) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupDHLEXPRequest) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupDHLEXPRequest) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPRequest) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPRequest) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPRequest) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupDHLEXPRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupDHLEXPRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupDHLEXPRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupDHLEXPRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupDHLEXPRequest) GetPickupOptions() []SchedulePickupDHLEXPRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupOptionsOk() (*[]SchedulePickupDHLEXPRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupDHLEXPRequest) SetPickupOptions(v []SchedulePickupDHLEXPRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupDHLEXPRequest) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


