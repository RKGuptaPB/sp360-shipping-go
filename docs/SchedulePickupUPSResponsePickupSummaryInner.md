# SchedulePickupUPSResponsePickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | description | [optional] 
**PackageCount** | Pointer to **float32** | description | [optional] 
**TotalWeight** | Pointer to **float32** | description | [optional] 
**WeightUnit** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**ParcelType** | Pointer to **string** | description | [optional] 
**PackageDetails** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner**](SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupUPSResponsePickupSummaryInner

`func NewSchedulePickupUPSResponsePickupSummaryInner() *SchedulePickupUPSResponsePickupSummaryInner`

NewSchedulePickupUPSResponsePickupSummaryInner instantiates a new SchedulePickupUPSResponsePickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults

`func NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults() *SchedulePickupUPSResponsePickupSummaryInner`

NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults instantiates a new SchedulePickupUPSResponsePickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetPackageDetails

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageDetails() []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageDetailsOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetPackageDetails(v []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


