# ShipmentInternational

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **string** | Defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8&#39; X 11&#39;). | 
**Type** | **string** | Defines the type of the Shipment, e.g., Shipping Label. | 
**Format** | Pointer to **string** | Defines the type of the shipment which is printed. For example, Shipping label prints in PDF form. | [optional] 
**DateOfShipment** | Pointer to **string** | Defines the date of the Shipment in the format YYYY:MM: DD. | [optional] 
**FromAddress** | [**ShipmentInternationalFromAddress**](ShipmentInternationalFromAddress.md) |  | 
**Parcel** | [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | 
**CarrierAccountId** | **string** | The unique identifier associated with the Carrier account used by client users during shipment process. | 
**ParcelType** | **string** | &gt;-Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. It has categories like Package, Envelopes, Paks, Boxes, Tube, defined per specific carrier and used in abbreviated form, e.g., FRPKG, LGENV, TUBE, PKG. | 
**ServiceId** | **string** | &gt;-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel. | 
**SpecialServices** | Pointer to [**[]SpecialService**](SpecialService.md) | It provides a carrier-service based special or extra service. | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentDomesticMetadataInner**](ShipmentDomesticMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#39;Cost Account Name&#39; is supported. | [optional] 
**ToAddress** | [**ShipmentInternationalToAddress**](ShipmentInternationalToAddress.md) |  | 
**Customs** | [**ShipmentInternationalCustoms**](ShipmentInternationalCustoms.md) |  | 

## Methods

### NewShipmentInternational

`func NewShipmentInternational(size string, type_ string, fromAddress ShipmentInternationalFromAddress, parcel ShipmentDomesticParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentInternationalToAddress, customs ShipmentInternationalCustoms, ) *ShipmentInternational`

NewShipmentInternational instantiates a new ShipmentInternational object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalWithDefaults

`func NewShipmentInternationalWithDefaults() *ShipmentInternational`

NewShipmentInternationalWithDefaults instantiates a new ShipmentInternational object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *ShipmentInternational) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ShipmentInternational) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ShipmentInternational) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *ShipmentInternational) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipmentInternational) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipmentInternational) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *ShipmentInternational) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ShipmentInternational) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ShipmentInternational) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ShipmentInternational) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDateOfShipment

`func (o *ShipmentInternational) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentInternational) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentInternational) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentInternational) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetFromAddress

`func (o *ShipmentInternational) GetFromAddress() ShipmentInternationalFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentInternational) GetFromAddressOk() (*ShipmentInternationalFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentInternational) SetFromAddress(v ShipmentInternationalFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *ShipmentInternational) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentInternational) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentInternational) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccountId

`func (o *ShipmentInternational) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ShipmentInternational) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ShipmentInternational) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *ShipmentInternational) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentInternational) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentInternational) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *ShipmentInternational) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShipmentInternational) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShipmentInternational) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *ShipmentInternational) GetSpecialServices() []SpecialService`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ShipmentInternational) GetSpecialServicesOk() (*[]SpecialService, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ShipmentInternational) SetSpecialServices(v []SpecialService)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ShipmentInternational) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentInternational) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentInternational) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentInternational) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentInternational) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentInternational) GetMetadata() []ShipmentDomesticMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentInternational) GetMetadataOk() (*[]ShipmentDomesticMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentInternational) SetMetadata(v []ShipmentDomesticMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentInternational) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToAddress

`func (o *ShipmentInternational) GetToAddress() ShipmentInternationalToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentInternational) GetToAddressOk() (*ShipmentInternationalToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentInternational) SetToAddress(v ShipmentInternationalToAddress)`

SetToAddress sets ToAddress field to given value.


### GetCustoms

`func (o *ShipmentInternational) GetCustoms() ShipmentInternationalCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *ShipmentInternational) GetCustomsOk() (*ShipmentInternationalCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *ShipmentInternational) SetCustoms(v ShipmentInternationalCustoms)`

SetCustoms sets Customs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


