# SchedulePickupUSPSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupUSPSRequestPickupSummaryInner**](SchedulePickupUSPSRequestPickupSummaryInner.md) | description | [optional] 
**Additionalnotes** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 

## Methods

### NewSchedulePickupUSPSRequest

`func NewSchedulePickupUSPSRequest() *SchedulePickupUSPSRequest`

NewSchedulePickupUSPSRequest instantiates a new SchedulePickupUSPSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUSPSRequestWithDefaults

`func NewSchedulePickupUSPSRequestWithDefaults() *SchedulePickupUSPSRequest`

NewSchedulePickupUSPSRequestWithDefaults instantiates a new SchedulePickupUSPSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUSPSRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUSPSRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUSPSRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupUSPSRequest) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupUSPSRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUSPSRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUSPSRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupUSPSRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupUSPSRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUSPSRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUSPSRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupUSPSRequest) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUSPSRequest) GetPickupSummary() []SchedulePickupUSPSRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUSPSRequest) GetPickupSummaryOk() (*[]SchedulePickupUSPSRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUSPSRequest) SetPickupSummary(v []SchedulePickupUSPSRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUSPSRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupUSPSRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupUSPSRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupUSPSRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupUSPSRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupUSPSRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupUSPSRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupUSPSRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupUSPSRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


