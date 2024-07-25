# CreateBulkShipmentsERRLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**Size** | **string** | This indicates the label size of the Bulk Shipment, e.g., DocSize can be 8&#39; X 11&#39; or 4&#39; X 6&#39;. | 
**Type** | **string** | This indicates the type of the Batch Shipment, e.g., Shipping Label. | 
**Format** | Pointer to **string** | This defines the type of the shipment which is printed. For example Shipping label prints in PDF form. | [optional] 
**Name** | **string** |  | 
**CarrierAccountId** | **string** | Default &#x60;carrierAccountId&#x60; to be used for this batch. You can override this value by defining it at shipment level. | 
**ParcelType** | **string** | Default &#x60;parcelType&#x60; specific to the carrier, e.g., PKG, SOFTPACK to be used for this batch. You can override this value by defining it at shipment level. | 
**ServiceId** | **string** | Default abbreviated name &#x60;serviceId&#x60; of the carrier-specific service to be used for this batch. You can override this value by defining it at shipment level. | 
**SpecialServices** | Pointer to [**[]SpecialServiceERRInner**](SpecialServiceERRInner.md) |  | [optional] 
**Shipments** | [**[]ShipmentERR**](ShipmentERR.md) |  | 

## Methods

### NewCreateBulkShipmentsERRLabel

`func NewCreateBulkShipmentsERRLabel(size string, type_ string, name string, carrierAccountId string, parcelType string, serviceId string, shipments []ShipmentERR, ) *CreateBulkShipmentsERRLabel`

NewCreateBulkShipmentsERRLabel instantiates a new CreateBulkShipmentsERRLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBulkShipmentsERRLabelWithDefaults

`func NewCreateBulkShipmentsERRLabelWithDefaults() *CreateBulkShipmentsERRLabel`

NewCreateBulkShipmentsERRLabelWithDefaults instantiates a new CreateBulkShipmentsERRLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *CreateBulkShipmentsERRLabel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CreateBulkShipmentsERRLabel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CreateBulkShipmentsERRLabel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CreateBulkShipmentsERRLabel) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSize

`func (o *CreateBulkShipmentsERRLabel) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBulkShipmentsERRLabel) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBulkShipmentsERRLabel) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *CreateBulkShipmentsERRLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBulkShipmentsERRLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBulkShipmentsERRLabel) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *CreateBulkShipmentsERRLabel) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateBulkShipmentsERRLabel) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateBulkShipmentsERRLabel) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateBulkShipmentsERRLabel) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *CreateBulkShipmentsERRLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBulkShipmentsERRLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBulkShipmentsERRLabel) SetName(v string)`

SetName sets Name field to given value.


### GetCarrierAccountId

`func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateBulkShipmentsERRLabel) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *CreateBulkShipmentsERRLabel) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateBulkShipmentsERRLabel) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateBulkShipmentsERRLabel) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *CreateBulkShipmentsERRLabel) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBulkShipmentsERRLabel) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBulkShipmentsERRLabel) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *CreateBulkShipmentsERRLabel) GetSpecialServices() []SpecialServiceERRInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *CreateBulkShipmentsERRLabel) GetSpecialServicesOk() (*[]SpecialServiceERRInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *CreateBulkShipmentsERRLabel) SetSpecialServices(v []SpecialServiceERRInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *CreateBulkShipmentsERRLabel) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipments

`func (o *CreateBulkShipmentsERRLabel) GetShipments() []ShipmentERR`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *CreateBulkShipmentsERRLabel) GetShipmentsOk() (*[]ShipmentERR, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *CreateBulkShipmentsERRLabel) SetShipments(v []ShipmentERR)`

SetShipments sets Shipments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


