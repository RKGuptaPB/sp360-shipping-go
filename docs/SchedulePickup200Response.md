# SchedulePickup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | description | [optional] 
**PickupConfirmationNumber** | Pointer to **string** | description | [optional] 
**PickupId** | Pointer to **string** | description | [optional] 
**Carrier** | Pointer to **string** | description | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupUPSResponsePickupSummaryInner**](SchedulePickupUPSResponsePickupSummaryInner.md) | description | [optional] 
**SpecialInstructions** | Pointer to **string** | description | [optional] 
**PickupStartTime** | Pointer to **string** | description | [optional] 
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**Reference** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupUPSRequestPickupOptionsInner**](SchedulePickupUPSRequestPickupOptionsInner.md) | description | [optional] 
**PickupRate** | Pointer to [**SchedulePickupUPSResponsePickupRate**](SchedulePickupUPSResponsePickupRate.md) |  | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | description | [optional] 

## Methods

### NewSchedulePickup200Response

`func NewSchedulePickup200Response() *SchedulePickup200Response`

NewSchedulePickup200Response instantiates a new SchedulePickup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickup200ResponseWithDefaults

`func NewSchedulePickup200ResponseWithDefaults() *SchedulePickup200Response`

NewSchedulePickup200ResponseWithDefaults instantiates a new SchedulePickup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickup200Response) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickup200Response) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickup200Response) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickup200Response) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickup200Response) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickup200Response) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickup200Response) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickup200Response) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickup200Response) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickup200Response) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickup200Response) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickup200Response) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickup200Response) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickup200Response) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickup200Response) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickup200Response) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickup200Response) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickup200Response) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickup200Response) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickup200Response) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickup200Response) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickup200Response) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickup200Response) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickup200Response) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickup200Response) GetPickupSummary() []SchedulePickupUPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickup200Response) GetPickupSummaryOk() (*[]SchedulePickupUPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickup200Response) SetPickupSummary(v []SchedulePickupUPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickup200Response) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *SchedulePickup200Response) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SchedulePickup200Response) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SchedulePickup200Response) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SchedulePickup200Response) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### GetPickupStartTime

`func (o *SchedulePickup200Response) GetPickupStartTime() string`

GetPickupStartTime returns the PickupStartTime field if non-nil, zero value otherwise.

### GetPickupStartTimeOk

`func (o *SchedulePickup200Response) GetPickupStartTimeOk() (*string, bool)`

GetPickupStartTimeOk returns a tuple with the PickupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartTime

`func (o *SchedulePickup200Response) SetPickupStartTime(v string)`

SetPickupStartTime sets PickupStartTime field to given value.

### HasPickupStartTime

`func (o *SchedulePickup200Response) HasPickupStartTime() bool`

HasPickupStartTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickup200Response) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickup200Response) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickup200Response) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickup200Response) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickup200Response) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickup200Response) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickup200Response) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickup200Response) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickup200Response) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickup200Response) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickup200Response) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickup200Response) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickup200Response) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickup200Response) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickup200Response) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickup200Response) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPickupRate

`func (o *SchedulePickup200Response) GetPickupRate() SchedulePickupUPSResponsePickupRate`

GetPickupRate returns the PickupRate field if non-nil, zero value otherwise.

### GetPickupRateOk

`func (o *SchedulePickup200Response) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool)`

GetPickupRateOk returns a tuple with the PickupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRate

`func (o *SchedulePickup200Response) SetPickupRate(v SchedulePickupUPSResponsePickupRate)`

SetPickupRate sets PickupRate field to given value.

### HasPickupRate

`func (o *SchedulePickup200Response) HasPickupRate() bool`

HasPickupRate returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *SchedulePickup200Response) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *SchedulePickup200Response) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *SchedulePickup200Response) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *SchedulePickup200Response) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


