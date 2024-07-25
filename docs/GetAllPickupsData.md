# GetAllPickupsData

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

### NewGetAllPickupsData

`func NewGetAllPickupsData() *GetAllPickupsData`

NewGetAllPickupsData instantiates a new GetAllPickupsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPickupsDataWithDefaults

`func NewGetAllPickupsDataWithDefaults() *GetAllPickupsData`

NewGetAllPickupsDataWithDefaults instantiates a new GetAllPickupsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *GetAllPickupsData) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *GetAllPickupsData) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *GetAllPickupsData) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *GetAllPickupsData) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *GetAllPickupsData) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *GetAllPickupsData) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *GetAllPickupsData) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *GetAllPickupsData) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *GetAllPickupsData) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *GetAllPickupsData) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *GetAllPickupsData) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *GetAllPickupsData) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *GetAllPickupsData) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetAllPickupsData) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetAllPickupsData) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetAllPickupsData) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *GetAllPickupsData) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *GetAllPickupsData) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *GetAllPickupsData) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *GetAllPickupsData) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *GetAllPickupsData) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *GetAllPickupsData) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *GetAllPickupsData) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *GetAllPickupsData) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetPickupSummary

`func (o *GetAllPickupsData) GetPickupSummary() []SchedulePickupUPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *GetAllPickupsData) GetPickupSummaryOk() (*[]SchedulePickupUPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *GetAllPickupsData) SetPickupSummary(v []SchedulePickupUPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *GetAllPickupsData) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *GetAllPickupsData) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *GetAllPickupsData) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *GetAllPickupsData) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *GetAllPickupsData) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *GetAllPickupsData) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *GetAllPickupsData) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *GetAllPickupsData) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *GetAllPickupsData) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *GetAllPickupsData) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *GetAllPickupsData) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *GetAllPickupsData) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *GetAllPickupsData) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *GetAllPickupsData) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *GetAllPickupsData) GetPickupOptionsOk() (*[]SchedulePickupUPSRequestPickupOptionsInner, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *GetAllPickupsData) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *GetAllPickupsData) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPickupRate

`func (o *GetAllPickupsData) GetPickupRate() SchedulePickupUPSResponsePickupRate`

GetPickupRate returns the PickupRate field if non-nil, zero value otherwise.

### GetPickupRateOk

`func (o *GetAllPickupsData) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool)`

GetPickupRateOk returns a tuple with the PickupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRate

`func (o *GetAllPickupsData) SetPickupRate(v SchedulePickupUPSResponsePickupRate)`

SetPickupRate sets PickupRate field to given value.

### HasPickupRate

`func (o *GetAllPickupsData) HasPickupRate() bool`

HasPickupRate returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *GetAllPickupsData) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *GetAllPickupsData) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *GetAllPickupsData) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *GetAllPickupsData) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


