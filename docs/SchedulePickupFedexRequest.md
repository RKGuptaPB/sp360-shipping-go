# SchedulePickupFedexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupFedexRequestPickupSummaryInner**](SchedulePickupFedexRequestPickupSummaryInner.md) | description | [optional] 
**Additionalnotes** | Pointer to **string** | description | [optional] 
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupUPSRequestPickupOptionsInner**](SchedulePickupUPSRequestPickupOptionsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupFedexRequest

`func NewSchedulePickupFedexRequest() *SchedulePickupFedexRequest`

NewSchedulePickupFedexRequest instantiates a new SchedulePickupFedexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupFedexRequestWithDefaults

`func NewSchedulePickupFedexRequestWithDefaults() *SchedulePickupFedexRequest`

NewSchedulePickupFedexRequestWithDefaults instantiates a new SchedulePickupFedexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupFedexRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupFedexRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupFedexRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupFedexRequest) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupFedexRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupFedexRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupFedexRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupFedexRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupFedexRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupFedexRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupFedexRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupFedexRequest) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupFedexRequest) GetPickupSummary() []SchedulePickupFedexRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupFedexRequest) GetPickupSummaryOk() (*[]SchedulePickupFedexRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupFedexRequest) SetPickupSummary(v []SchedulePickupFedexRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupFedexRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupFedexRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupFedexRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupFedexRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupFedexRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupFedexRequest) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupFedexRequest) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupFedexRequest) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupFedexRequest) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupFedexRequest) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupFedexRequest) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupFedexRequest) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupFedexRequest) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupFedexRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupFedexRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupFedexRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupFedexRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupFedexRequest) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupFedexRequest) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupFedexRequest) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupFedexRequest) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


