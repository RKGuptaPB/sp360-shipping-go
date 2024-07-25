# ShipmentERRCoversheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the shipment. User can provide any custom value to it for its own reference | [optional] 
**FromAddress** | [**ShipmentFromAddress**](ShipmentFromAddress.md) |  | 
**Parcel** | [**ShipmentParcel**](ShipmentParcel.md) |  | 
**CarrierAccountId** | **string** | A unique identifier associated with the Carrier account used by client users during shipment process. | 
**ParcelType** | **string** | &gt;-Packaging type specific to the carrier, e.g., LTR, LGENV. | 
**ServiceId** | **string** | &gt;-The abbreviated name of the carrier-specific service. Required for creating a shipment. Optional for rating a parcel. | 
**DateOfShipment** | Pointer to **string** | Current Shipment date | [optional] 
**SpecialServices** | Pointer to [**[]SpecialServiceERRInner**](SpecialServiceERRInner.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentMetadataInner**](ShipmentMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ToAddress** | [**ShipmentToAddress**](ShipmentToAddress.md) |  | 

## Methods

### NewShipmentERRCoversheet

`func NewShipmentERRCoversheet(fromAddress ShipmentFromAddress, parcel ShipmentParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentToAddress, ) *ShipmentERRCoversheet`

NewShipmentERRCoversheet instantiates a new ShipmentERRCoversheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentERRCoversheetWithDefaults

`func NewShipmentERRCoversheetWithDefaults() *ShipmentERRCoversheet`

NewShipmentERRCoversheetWithDefaults instantiates a new ShipmentERRCoversheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ShipmentERRCoversheet) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ShipmentERRCoversheet) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ShipmentERRCoversheet) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ShipmentERRCoversheet) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFromAddress

`func (o *ShipmentERRCoversheet) GetFromAddress() ShipmentFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentERRCoversheet) GetFromAddressOk() (*ShipmentFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentERRCoversheet) SetFromAddress(v ShipmentFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *ShipmentERRCoversheet) GetParcel() ShipmentParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentERRCoversheet) GetParcelOk() (*ShipmentParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentERRCoversheet) SetParcel(v ShipmentParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccountId

`func (o *ShipmentERRCoversheet) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ShipmentERRCoversheet) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ShipmentERRCoversheet) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *ShipmentERRCoversheet) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentERRCoversheet) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentERRCoversheet) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *ShipmentERRCoversheet) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShipmentERRCoversheet) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShipmentERRCoversheet) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetDateOfShipment

`func (o *ShipmentERRCoversheet) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentERRCoversheet) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentERRCoversheet) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentERRCoversheet) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetSpecialServices

`func (o *ShipmentERRCoversheet) GetSpecialServices() []SpecialServiceERRInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ShipmentERRCoversheet) GetSpecialServicesOk() (*[]SpecialServiceERRInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ShipmentERRCoversheet) SetSpecialServices(v []SpecialServiceERRInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ShipmentERRCoversheet) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentERRCoversheet) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentERRCoversheet) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentERRCoversheet) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentERRCoversheet) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentERRCoversheet) GetMetadata() []ShipmentMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentERRCoversheet) GetMetadataOk() (*[]ShipmentMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentERRCoversheet) SetMetadata(v []ShipmentMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentERRCoversheet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToAddress

`func (o *ShipmentERRCoversheet) GetToAddress() ShipmentToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentERRCoversheet) GetToAddressOk() (*ShipmentToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentERRCoversheet) SetToAddress(v ShipmentToAddress)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


