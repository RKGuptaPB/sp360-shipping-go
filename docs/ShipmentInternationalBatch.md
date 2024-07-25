# ShipmentInternationalBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | [**ShipmentInternationalBatchFromAddress**](ShipmentInternationalBatchFromAddress.md) |  | 
**Parcel** | [**ShipmentInternationalBatchParcel**](ShipmentInternationalBatchParcel.md) |  | 
**CarrierAccountId** | **string** | A unique identifier associated with the Carrier account used by client users during shipment process. | 
**ParcelType** | **string** | &gt;-Packaging type specific to the carrier, e.g., FRPKG, LGENV, TUBE,PKG. | 
**ServiceId** | **string** | &gt;-The abbreviated name of the carrier-specific service. Required for creating a shipment. Optional for rating a parcel. | 
**DateOfShipment** | Pointer to **string** | The date of the shipment. The format must be YYY:MM:DD. | [optional] 
**SpecialServices** | Pointer to [**[]SpecialService**](SpecialService.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentDomesticMetadataInner**](ShipmentDomesticMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ToAddress** | [**ShipmentInternationalBatchToAddress**](ShipmentInternationalBatchToAddress.md) |  | 
**Customs** | [**ShipmentInternationalBatchCustoms**](ShipmentInternationalBatchCustoms.md) |  | 

## Methods

### NewShipmentInternationalBatch

`func NewShipmentInternationalBatch(fromAddress ShipmentInternationalBatchFromAddress, parcel ShipmentInternationalBatchParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentInternationalBatchToAddress, customs ShipmentInternationalBatchCustoms, ) *ShipmentInternationalBatch`

NewShipmentInternationalBatch instantiates a new ShipmentInternationalBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalBatchWithDefaults

`func NewShipmentInternationalBatchWithDefaults() *ShipmentInternationalBatch`

NewShipmentInternationalBatchWithDefaults instantiates a new ShipmentInternationalBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *ShipmentInternationalBatch) GetFromAddress() ShipmentInternationalBatchFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentInternationalBatch) GetFromAddressOk() (*ShipmentInternationalBatchFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentInternationalBatch) SetFromAddress(v ShipmentInternationalBatchFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *ShipmentInternationalBatch) GetParcel() ShipmentInternationalBatchParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentInternationalBatch) GetParcelOk() (*ShipmentInternationalBatchParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentInternationalBatch) SetParcel(v ShipmentInternationalBatchParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccountId

`func (o *ShipmentInternationalBatch) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ShipmentInternationalBatch) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ShipmentInternationalBatch) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *ShipmentInternationalBatch) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentInternationalBatch) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentInternationalBatch) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *ShipmentInternationalBatch) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShipmentInternationalBatch) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShipmentInternationalBatch) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetDateOfShipment

`func (o *ShipmentInternationalBatch) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentInternationalBatch) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentInternationalBatch) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentInternationalBatch) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetSpecialServices

`func (o *ShipmentInternationalBatch) GetSpecialServices() []SpecialService`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ShipmentInternationalBatch) GetSpecialServicesOk() (*[]SpecialService, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ShipmentInternationalBatch) SetSpecialServices(v []SpecialService)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ShipmentInternationalBatch) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentInternationalBatch) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentInternationalBatch) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentInternationalBatch) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentInternationalBatch) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentInternationalBatch) GetMetadata() []ShipmentDomesticMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentInternationalBatch) GetMetadataOk() (*[]ShipmentDomesticMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentInternationalBatch) SetMetadata(v []ShipmentDomesticMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentInternationalBatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToAddress

`func (o *ShipmentInternationalBatch) GetToAddress() ShipmentInternationalBatchToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentInternationalBatch) GetToAddressOk() (*ShipmentInternationalBatchToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentInternationalBatch) SetToAddress(v ShipmentInternationalBatchToAddress)`

SetToAddress sets ToAddress field to given value.


### GetCustoms

`func (o *ShipmentInternationalBatch) GetCustoms() ShipmentInternationalBatchCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *ShipmentInternationalBatch) GetCustomsOk() (*ShipmentInternationalBatchCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *ShipmentInternationalBatch) SetCustoms(v ShipmentInternationalBatchCustoms)`

SetCustoms sets Customs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


