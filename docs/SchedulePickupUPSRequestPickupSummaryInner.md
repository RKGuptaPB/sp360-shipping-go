# SchedulePickupUPSRequestPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | description | [optional] 
**PackageCount** | Pointer to **float32** | description | [optional] 
**TotalWeight** | Pointer to **float32** | description | [optional] 
**WeightUnit** | Pointer to **string** | description | [optional] 
**ParcelType** | Pointer to **string** | description | [optional] 
**ToAddressCountryCode** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**PackageDetails** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner**](SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupUPSRequestPickupSummaryInner

`func NewSchedulePickupUPSRequestPickupSummaryInner() *SchedulePickupUPSRequestPickupSummaryInner`

NewSchedulePickupUPSRequestPickupSummaryInner instantiates a new SchedulePickupUPSRequestPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults

`func NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults() *SchedulePickupUPSRequestPickupSummaryInner`

NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupUPSRequestPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetParcelType

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetToAddressCountryCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetToAddressCountryCode() string`

GetToAddressCountryCode returns the ToAddressCountryCode field if non-nil, zero value otherwise.

### GetToAddressCountryCodeOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetToAddressCountryCodeOk() (*string, bool)`

GetToAddressCountryCodeOk returns a tuple with the ToAddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressCountryCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetToAddressCountryCode(v string)`

SetToAddressCountryCode sets ToAddressCountryCode field to given value.

### HasToAddressCountryCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasToAddressCountryCode() bool`

HasToAddressCountryCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPackageDetails

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageDetails() []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageDetailsOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetPackageDetails(v []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *SchedulePickupUPSRequestPickupSummaryInner) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


