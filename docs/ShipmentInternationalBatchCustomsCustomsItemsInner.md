# ShipmentInternationalBatchCustomsCustomsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A detailed description of the commodity, up to 255 characters. The description will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. Each description will appear on the form. | 
**HSTariffCode** | Pointer to **string** | The destination country’s tariff-classification number (HS code) for the commodity. Most countries use the six-digit Harmonized System (HS) as the basis for their tariff classifications and then add digits for more detail. The maximum length for an HS code is 14 characters. The HS code will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. | [optional] 
**OriginCountryCode** | Pointer to **string** | The two-character ISO country code of the shipment’s origin country. Use ISO 3166-1 alpha-2 standard values. | [optional] 
**Quantity** | **float32** | Enter the total number of items of this type of commodity. | 
**UnitPrice** | **float32** | The price of one item of this type of commodity. | 
**WeightUnit** | Pointer to **string** | The unit of measurement. This field is required by the unitWeight object. | [optional] 
**Weight** | Pointer to **float32** | The weight of the item. | [optional] 

## Methods

### NewShipmentInternationalBatchCustomsCustomsItemsInner

`func NewShipmentInternationalBatchCustomsCustomsItemsInner(description string, quantity float32, unitPrice float32, ) *ShipmentInternationalBatchCustomsCustomsItemsInner`

NewShipmentInternationalBatchCustomsCustomsItemsInner instantiates a new ShipmentInternationalBatchCustomsCustomsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalBatchCustomsCustomsItemsInnerWithDefaults

`func NewShipmentInternationalBatchCustomsCustomsItemsInnerWithDefaults() *ShipmentInternationalBatchCustomsCustomsItemsInner`

NewShipmentInternationalBatchCustomsCustomsItemsInnerWithDefaults instantiates a new ShipmentInternationalBatchCustomsCustomsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHSTariffCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetHSTariffCode() string`

GetHSTariffCode returns the HSTariffCode field if non-nil, zero value otherwise.

### GetHSTariffCodeOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetHSTariffCodeOk() (*string, bool)`

GetHSTariffCodeOk returns a tuple with the HSTariffCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSTariffCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetHSTariffCode(v string)`

SetHSTariffCode sets HSTariffCode field to given value.

### HasHSTariffCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) HasHSTariffCode() bool`

HasHSTariffCode returns a boolean if a field has been set.

### GetOriginCountryCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetOriginCountryCode() string`

GetOriginCountryCode returns the OriginCountryCode field if non-nil, zero value otherwise.

### GetOriginCountryCodeOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetOriginCountryCodeOk() (*string, bool)`

GetOriginCountryCodeOk returns a tuple with the OriginCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetOriginCountryCode(v string)`

SetOriginCountryCode sets OriginCountryCode field to given value.

### HasOriginCountryCode

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) HasOriginCountryCode() bool`

HasOriginCountryCode returns a boolean if a field has been set.

### GetQuantity

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetWeightUnit

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWeight

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ShipmentInternationalBatchCustomsCustomsItemsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


