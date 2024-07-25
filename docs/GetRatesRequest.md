# GetRatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | [**RateShopFromAddress**](RateShopFromAddress.md) |  | 
**Parcel** | [**RateShopParcel**](RateShopParcel.md) |  | 
**CarrierAccounts** | **[]string** | It provides one or more carrier accounts Ids for rate shop. It can be referred from &#x60;Get Carrier Accounts&#x60; API | 
**ParcelType** | **string** | Parcel Type required for rating, it&#39;s value can be referred from response of &#x60;Get Parcel Types&#x60; API | 
**ServiceId** | **string** | Service to be used for rating, it can be referred from response of &#x60;Get Services&#x60; API | 
**SpecialServices** | Pointer to [**[]SpecialService**](SpecialService.md) | Special services to be used for rating, it can be referred from response of &#x60;Get Special Services&#x60; API | [optional] 
**ToAddress** | [**SingleRateToAddress**](SingleRateToAddress.md) |  | 
**DateOfShipment** | Pointer to **string** | Defines the date of the Shipment in the format YYYY:MM: DD., required for future shipment rating | [optional] 
**IsHazmat** | Pointer to **bool** | isHazmat if set to true, only services which support Hazardous material shipment would be visible in the response | [optional] 

## Methods

### NewGetRatesRequest

`func NewGetRatesRequest(fromAddress RateShopFromAddress, parcel RateShopParcel, carrierAccounts []string, parcelType string, serviceId string, toAddress SingleRateToAddress, ) *GetRatesRequest`

NewGetRatesRequest instantiates a new GetRatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRatesRequestWithDefaults

`func NewGetRatesRequestWithDefaults() *GetRatesRequest`

NewGetRatesRequestWithDefaults instantiates a new GetRatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *GetRatesRequest) GetFromAddress() RateShopFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetRatesRequest) GetFromAddressOk() (*RateShopFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetRatesRequest) SetFromAddress(v RateShopFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *GetRatesRequest) GetParcel() RateShopParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *GetRatesRequest) GetParcelOk() (*RateShopParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *GetRatesRequest) SetParcel(v RateShopParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccounts

`func (o *GetRatesRequest) GetCarrierAccounts() []string`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *GetRatesRequest) GetCarrierAccountsOk() (*[]string, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *GetRatesRequest) SetCarrierAccounts(v []string)`

SetCarrierAccounts sets CarrierAccounts field to given value.


### GetParcelType

`func (o *GetRatesRequest) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *GetRatesRequest) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *GetRatesRequest) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *GetRatesRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetRatesRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetRatesRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *GetRatesRequest) GetSpecialServices() []SpecialService`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *GetRatesRequest) GetSpecialServicesOk() (*[]SpecialService, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *GetRatesRequest) SetSpecialServices(v []SpecialService)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *GetRatesRequest) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetToAddress

`func (o *GetRatesRequest) GetToAddress() SingleRateToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetRatesRequest) GetToAddressOk() (*SingleRateToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetRatesRequest) SetToAddress(v SingleRateToAddress)`

SetToAddress sets ToAddress field to given value.


### GetDateOfShipment

`func (o *GetRatesRequest) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *GetRatesRequest) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *GetRatesRequest) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *GetRatesRequest) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetIsHazmat

`func (o *GetRatesRequest) GetIsHazmat() bool`

GetIsHazmat returns the IsHazmat field if non-nil, zero value otherwise.

### GetIsHazmatOk

`func (o *GetRatesRequest) GetIsHazmatOk() (*bool, bool)`

GetIsHazmatOk returns a tuple with the IsHazmat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHazmat

`func (o *GetRatesRequest) SetIsHazmat(v bool)`

SetIsHazmat sets IsHazmat field to given value.

### HasIsHazmat

`func (o *GetRatesRequest) HasIsHazmat() bool`

HasIsHazmat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


