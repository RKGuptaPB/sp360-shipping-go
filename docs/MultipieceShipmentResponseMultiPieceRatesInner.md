# MultipieceShipmentResponseMultiPieceRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCarrierCharge** | Pointer to **float32** | description | [optional] 
**Carrier** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**RateTypeId** | Pointer to **string** | description | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**ServiceName** | Pointer to **string** | description | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner**](MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner.md) | description | [optional] 
**Surcharges** | Pointer to [**[]MultipieceShipmentResponseMultiPieceRatesInnerSurchargesInner**](MultipieceShipmentResponseMultiPieceRatesInnerSurchargesInner.md) | description | [optional] 

## Methods

### NewMultipieceShipmentResponseMultiPieceRatesInner

`func NewMultipieceShipmentResponseMultiPieceRatesInner() *MultipieceShipmentResponseMultiPieceRatesInner`

NewMultipieceShipmentResponseMultiPieceRatesInner instantiates a new MultipieceShipmentResponseMultiPieceRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceShipmentResponseMultiPieceRatesInnerWithDefaults

`func NewMultipieceShipmentResponseMultiPieceRatesInnerWithDefaults() *MultipieceShipmentResponseMultiPieceRatesInner`

NewMultipieceShipmentResponseMultiPieceRatesInnerWithDefaults instantiates a new MultipieceShipmentResponseMultiPieceRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCarrierCharge

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRateTypeId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetMultiPieceParcels() []MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetMultiPieceParcelsOk() (*[]MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetMultiPieceParcels(v []MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetSurcharges

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetSurcharges() []MultipieceShipmentResponseMultiPieceRatesInnerSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) GetSurchargesOk() (*[]MultipieceShipmentResponseMultiPieceRatesInnerSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) SetSurcharges(v []MultipieceShipmentResponseMultiPieceRatesInnerSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *MultipieceShipmentResponseMultiPieceRatesInner) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


