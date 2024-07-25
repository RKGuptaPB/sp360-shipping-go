# SchedulePickupUSPSResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupConfirmationNumber** | Pointer to **string** | description | [optional] 
**PickupId** | Pointer to **string** | description | [optional] 
**Carrier** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupUSPSRequestPickupSummaryInner**](SchedulePickupUSPSRequestPickupSummaryInner.md) | description | [optional] 
**SpecialInstructions** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 

## Methods

### NewSchedulePickupUSPSResponse

`func NewSchedulePickupUSPSResponse() *SchedulePickupUSPSResponse`

NewSchedulePickupUSPSResponse instantiates a new SchedulePickupUSPSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUSPSResponseWithDefaults

`func NewSchedulePickupUSPSResponseWithDefaults() *SchedulePickupUSPSResponse`

NewSchedulePickupUSPSResponseWithDefaults instantiates a new SchedulePickupUSPSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUSPSResponse) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUSPSResponse) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUSPSResponse) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupUSPSResponse) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupUSPSResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUSPSResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUSPSResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupUSPSResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickupUSPSResponse) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickupUSPSResponse) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickupUSPSResponse) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickupUSPSResponse) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickupUSPSResponse) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickupUSPSResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickupUSPSResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickupUSPSResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickupUSPSResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupUSPSResponse) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUSPSResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUSPSResponse) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupUSPSResponse) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUSPSResponse) GetPickupSummary() []SchedulePickupUSPSRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUSPSResponse) GetPickupSummaryOk() (*[]SchedulePickupUSPSRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUSPSResponse) SetPickupSummary(v []SchedulePickupUSPSRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUSPSResponse) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *SchedulePickupUSPSResponse) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SchedulePickupUSPSResponse) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SchedulePickupUSPSResponse) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SchedulePickupUSPSResponse) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupUSPSResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupUSPSResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupUSPSResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupUSPSResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


