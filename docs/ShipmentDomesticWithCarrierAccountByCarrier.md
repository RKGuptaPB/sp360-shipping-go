# ShipmentDomesticWithCarrierAccountByCarrier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | Pointer to **string** | This is a unique identifier associated with the specific sub-carrier account, which must be valid.&lt;br /&gt; This is used in the shipment creation (if this value is defined, Carrier properties will be skipped). | [optional] 
**Service** | Pointer to **string** | Indicates a unique identifier associated with the carrier specific service, that is ServiceID, which must be valid. | [optional] 

## Methods

### NewShipmentDomesticWithCarrierAccountByCarrier

`func NewShipmentDomesticWithCarrierAccountByCarrier() *ShipmentDomesticWithCarrierAccountByCarrier`

NewShipmentDomesticWithCarrierAccountByCarrier instantiates a new ShipmentDomesticWithCarrierAccountByCarrier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticWithCarrierAccountByCarrierWithDefaults

`func NewShipmentDomesticWithCarrierAccountByCarrierWithDefaults() *ShipmentDomesticWithCarrierAccountByCarrier`

NewShipmentDomesticWithCarrierAccountByCarrierWithDefaults instantiates a new ShipmentDomesticWithCarrierAccountByCarrier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetService

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ShipmentDomesticWithCarrierAccountByCarrier) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


