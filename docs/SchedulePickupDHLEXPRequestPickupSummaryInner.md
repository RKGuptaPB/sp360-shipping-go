# SchedulePickupDHLEXPRequestPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | description | [optional] 
**PackageCount** | Pointer to **float32** | description | [optional] 
**TotalWeight** | Pointer to **float32** | description | [optional] 
**WeightUnit** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**PackageDetails** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner**](SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupDHLEXPRequestPickupSummaryInner

`func NewSchedulePickupDHLEXPRequestPickupSummaryInner() *SchedulePickupDHLEXPRequestPickupSummaryInner`

NewSchedulePickupDHLEXPRequestPickupSummaryInner instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults

`func NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults() *SchedulePickupDHLEXPRequestPickupSummaryInner`

NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPackageDetails

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageDetails() []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageDetailsOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetPackageDetails(v []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


