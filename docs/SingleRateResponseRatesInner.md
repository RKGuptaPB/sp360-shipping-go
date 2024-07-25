# SingleRateResponseRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | The base service charge is payable to the carrier, excluding special service charges. | [optional] 
**Carrier** | Pointer to **string** | Carrier is a service used to transport the parcels or couriers from one place to another. | [optional] 
**CarrierAccount** | Pointer to **string** | Carrier Account is the account associated to specific carrier service. | [optional] 
**DeliveryCommitment** | Pointer to [**RateShopResponseRatesInnerDeliveryCommitment**](RateShopResponseRatesInnerDeliveryCommitment.md) |  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc.  | [optional] 
**RateTypeId** | Pointer to **string** | Its value can be CONTRACT_RATES, COMMERCIAL or COMMERCIAL_BASE for USPS and COMMERCIAL for other carriers depending on the Pitney Bowes contract/subscription | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**SpecialServices** | Pointer to [**[]SingleRateResponseRatesInnerSpecialServicesInner**](SingleRateResponseRatesInnerSpecialServicesInner.md) |  It provides a carrier-service based special or extra service. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties. | [optional] 

## Methods

### NewSingleRateResponseRatesInner

`func NewSingleRateResponseRatesInner() *SingleRateResponseRatesInner`

NewSingleRateResponseRatesInner instantiates a new SingleRateResponseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRateResponseRatesInnerWithDefaults

`func NewSingleRateResponseRatesInnerWithDefaults() *SingleRateResponseRatesInner`

NewSingleRateResponseRatesInnerWithDefaults instantiates a new SingleRateResponseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *SingleRateResponseRatesInner) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *SingleRateResponseRatesInner) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *SingleRateResponseRatesInner) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *SingleRateResponseRatesInner) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *SingleRateResponseRatesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SingleRateResponseRatesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SingleRateResponseRatesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SingleRateResponseRatesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccount

`func (o *SingleRateResponseRatesInner) GetCarrierAccount() string`

GetCarrierAccount returns the CarrierAccount field if non-nil, zero value otherwise.

### GetCarrierAccountOk

`func (o *SingleRateResponseRatesInner) GetCarrierAccountOk() (*string, bool)`

GetCarrierAccountOk returns a tuple with the CarrierAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccount

`func (o *SingleRateResponseRatesInner) SetCarrierAccount(v string)`

SetCarrierAccount sets CarrierAccount field to given value.

### HasCarrierAccount

`func (o *SingleRateResponseRatesInner) HasCarrierAccount() bool`

HasCarrierAccount returns a boolean if a field has been set.

### GetDeliveryCommitment

`func (o *SingleRateResponseRatesInner) GetDeliveryCommitment() RateShopResponseRatesInnerDeliveryCommitment`

GetDeliveryCommitment returns the DeliveryCommitment field if non-nil, zero value otherwise.

### GetDeliveryCommitmentOk

`func (o *SingleRateResponseRatesInner) GetDeliveryCommitmentOk() (*RateShopResponseRatesInnerDeliveryCommitment, bool)`

GetDeliveryCommitmentOk returns a tuple with the DeliveryCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCommitment

`func (o *SingleRateResponseRatesInner) SetDeliveryCommitment(v RateShopResponseRatesInnerDeliveryCommitment)`

SetDeliveryCommitment sets DeliveryCommitment field to given value.

### HasDeliveryCommitment

`func (o *SingleRateResponseRatesInner) HasDeliveryCommitment() bool`

HasDeliveryCommitment returns a boolean if a field has been set.

### GetParcelType

`func (o *SingleRateResponseRatesInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SingleRateResponseRatesInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SingleRateResponseRatesInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *SingleRateResponseRatesInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateTypeId

`func (o *SingleRateResponseRatesInner) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *SingleRateResponseRatesInner) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *SingleRateResponseRatesInner) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *SingleRateResponseRatesInner) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *SingleRateResponseRatesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SingleRateResponseRatesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SingleRateResponseRatesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SingleRateResponseRatesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSpecialServices

`func (o *SingleRateResponseRatesInner) GetSpecialServices() []SingleRateResponseRatesInnerSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *SingleRateResponseRatesInner) GetSpecialServicesOk() (*[]SingleRateResponseRatesInnerSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *SingleRateResponseRatesInner) SetSpecialServices(v []SingleRateResponseRatesInnerSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *SingleRateResponseRatesInner) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *SingleRateResponseRatesInner) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *SingleRateResponseRatesInner) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *SingleRateResponseRatesInner) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *SingleRateResponseRatesInner) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


