# MultipieceShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | description | [optional] 
**Type** | Pointer to **string** | description | [optional] 
**FromAddress** | Pointer to [**MultipieceShipmentRequestFromAddress**](MultipieceShipmentRequestFromAddress.md) |  | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**ShipmentOptions** | Pointer to [**MultipieceShipmentRequestShipmentOptions**](MultipieceShipmentRequestShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]MultipieceShipmentRequestMetadataInner**](MultipieceShipmentRequestMetadataInner.md) | description | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceShipmentRequestMultiPieceParcelsInner**](MultipieceShipmentRequestMultiPieceParcelsInner.md) | description | [optional] 
**ToAddress** | Pointer to [**MultipieceShipmentRequestToAddress**](MultipieceShipmentRequestToAddress.md) |  | [optional] 

## Methods

### NewMultipieceShipmentRequest

`func NewMultipieceShipmentRequest() *MultipieceShipmentRequest`

NewMultipieceShipmentRequest instantiates a new MultipieceShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceShipmentRequestWithDefaults

`func NewMultipieceShipmentRequestWithDefaults() *MultipieceShipmentRequest`

NewMultipieceShipmentRequestWithDefaults instantiates a new MultipieceShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *MultipieceShipmentRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MultipieceShipmentRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MultipieceShipmentRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MultipieceShipmentRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *MultipieceShipmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultipieceShipmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultipieceShipmentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MultipieceShipmentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceShipmentRequest) GetFromAddress() MultipieceShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceShipmentRequest) GetFromAddressOk() (*MultipieceShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceShipmentRequest) SetFromAddress(v MultipieceShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceShipmentRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *MultipieceShipmentRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *MultipieceShipmentRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *MultipieceShipmentRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *MultipieceShipmentRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceShipmentRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceShipmentRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceShipmentRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceShipmentRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceShipmentRequest) GetShipmentOptions() MultipieceShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceShipmentRequest) GetShipmentOptionsOk() (*MultipieceShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceShipmentRequest) SetShipmentOptions(v MultipieceShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceShipmentRequest) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *MultipieceShipmentRequest) GetMetadata() []MultipieceShipmentRequestMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MultipieceShipmentRequest) GetMetadataOk() (*[]MultipieceShipmentRequestMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MultipieceShipmentRequest) SetMetadata(v []MultipieceShipmentRequestMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MultipieceShipmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceShipmentRequest) GetMultiPieceParcels() []MultipieceShipmentRequestMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceShipmentRequest) GetMultiPieceParcelsOk() (*[]MultipieceShipmentRequestMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceShipmentRequest) SetMultiPieceParcels(v []MultipieceShipmentRequestMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceShipmentRequest) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceShipmentRequest) GetToAddress() MultipieceShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceShipmentRequest) GetToAddressOk() (*MultipieceShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceShipmentRequest) SetToAddress(v MultipieceShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceShipmentRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


