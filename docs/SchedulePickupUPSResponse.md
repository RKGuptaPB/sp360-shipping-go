# SchedulePickupUPSResponse

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
**PickupTotalWeight** | Pointer to **float32** | description | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | description | [optional] 
**PickupOptions** | Pointer to [**[]SchedulePickupUPSRequestPickupOptionsInner**](SchedulePickupUPSRequestPickupOptionsInner.md) | description | [optional] 
**PickupRate** | Pointer to [**SchedulePickupUPSResponsePickupRate**](SchedulePickupUPSResponsePickupRate.md) |  | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | description | [optional] 

## Methods

### NewSchedulePickupUPSResponse

`func NewSchedulePickupUPSResponse() *SchedulePickupUPSResponse`

NewSchedulePickupUPSResponse instantiates a new SchedulePickupUPSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSResponseWithDefaults

`func NewSchedulePickupUPSResponseWithDefaults() *SchedulePickupUPSResponse`

NewSchedulePickupUPSResponseWithDefaults instantiates a new SchedulePickupUPSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUPSResponse) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUPSResponse) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUPSResponse) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupUPSResponse) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickupUPSResponse) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickupUPSResponse) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickupUPSResponse) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickupUPSResponse) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickupUPSResponse) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickupUPSResponse) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickupUPSResponse) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickupUPSResponse) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickupUPSResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickupUPSResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickupUPSResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickupUPSResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupUPSResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUPSResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUPSResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupUPSResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupUPSResponse) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUPSResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUPSResponse) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupUPSResponse) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUPSResponse) GetPickupSummary() []SchedulePickupUPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUPSResponse) GetPickupSummaryOk() (*[]SchedulePickupUPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUPSResponse) SetPickupSummary(v []SchedulePickupUPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUPSResponse) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *SchedulePickupUPSResponse) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SchedulePickupUPSResponse) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SchedulePickupUPSResponse) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SchedulePickupUPSResponse) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupUPSResponse) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupUPSResponse) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupUPSResponse) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupUPSResponse) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupUPSResponse) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupUPSResponse) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupUPSResponse) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupUPSResponse) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupUPSResponse) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupUPSResponse) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupUPSResponse) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupUPSResponse) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPickupRate

`func (o *SchedulePickupUPSResponse) GetPickupRate() SchedulePickupUPSResponsePickupRate`

GetPickupRate returns the PickupRate field if non-nil, zero value otherwise.

### GetPickupRateOk

`func (o *SchedulePickupUPSResponse) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool)`

GetPickupRateOk returns a tuple with the PickupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRate

`func (o *SchedulePickupUPSResponse) SetPickupRate(v SchedulePickupUPSResponsePickupRate)`

SetPickupRate sets PickupRate field to given value.

### HasPickupRate

`func (o *SchedulePickupUPSResponse) HasPickupRate() bool`

HasPickupRate returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *SchedulePickupUPSResponse) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *SchedulePickupUPSResponse) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *SchedulePickupUPSResponse) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *SchedulePickupUPSResponse) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


