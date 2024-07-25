# ShipmentDomesticWithCarrierAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | [**FromAddressV2**](FromAddressV2.md) |  | 
**ToAddress** | [**ToAddressV2**](ToAddressV2.md) |  | 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection.&lt;br /&gt; ParcelType can have categories like Package, Envelopes, Paks, Boxes, Tube, etc. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;  | [optional] 
**RateShopBy** | Pointer to **string** | By RateShop is a way to restrict the carrier services returned in a Rate Shop experience. &lt;br /&gt;  Without this Rate Shop Group, which is attached to an Enterprise or Location, the carrier services shown to the end user, are all that will be available for each carrier, based on Addresses, Weight, Dimensions, and Package Type. | [optional] 
**ByCarrier** | Pointer to [**ShipmentDomesticWithCarrierAccountByCarrier**](ShipmentDomesticWithCarrierAccountByCarrier.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentDomesticWithCarrierAccountShipmentOptions**](ShipmentDomesticWithCarrierAccountShipmentOptions.md) |  | [optional] 
**References** | Pointer to [**ReferenceV2**](ReferenceV2.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentDomesticWithCarrierAccountMetadataInner**](ShipmentDomesticWithCarrierAccountMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here.&lt;br /&gt; For now, &#39;Cost Account Name&#39; is supported. | [optional] 
**LabelSize** | Pointer to **string** | Defines the label size of the Shipment, that is, the Shipping Label is available in different Doc Size. | [optional] 
**LabelType** | Pointer to **string** | Defines the type of the Shipment. | [optional] 
**LabelFormat** | Pointer to **string** | Defines the file/format in which the label is printed.&lt;br /&gt; For ZPL2, DOC_4X6 will be supported, while for PDF, both the sizes are supported. | [optional] 
**PrinterAliasName** | Pointer to **string** | Refers to a printer connected (directly or via network) to a computer. | [optional] 

## Methods

### NewShipmentDomesticWithCarrierAccount

`func NewShipmentDomesticWithCarrierAccount(fromAddress FromAddressV2, toAddress ToAddressV2, ) *ShipmentDomesticWithCarrierAccount`

NewShipmentDomesticWithCarrierAccount instantiates a new ShipmentDomesticWithCarrierAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticWithCarrierAccountWithDefaults

`func NewShipmentDomesticWithCarrierAccountWithDefaults() *ShipmentDomesticWithCarrierAccount`

NewShipmentDomesticWithCarrierAccountWithDefaults instantiates a new ShipmentDomesticWithCarrierAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *ShipmentDomesticWithCarrierAccount) GetFromAddress() FromAddressV2`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentDomesticWithCarrierAccount) GetFromAddressOk() (*FromAddressV2, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentDomesticWithCarrierAccount) SetFromAddress(v FromAddressV2)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *ShipmentDomesticWithCarrierAccount) GetToAddress() ToAddressV2`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentDomesticWithCarrierAccount) GetToAddressOk() (*ToAddressV2, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentDomesticWithCarrierAccount) SetToAddress(v ToAddressV2)`

SetToAddress sets ToAddress field to given value.


### GetParcel

`func (o *ShipmentDomesticWithCarrierAccount) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentDomesticWithCarrierAccount) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentDomesticWithCarrierAccount) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ShipmentDomesticWithCarrierAccount) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelType

`func (o *ShipmentDomesticWithCarrierAccount) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentDomesticWithCarrierAccount) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentDomesticWithCarrierAccount) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *ShipmentDomesticWithCarrierAccount) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateShopBy

`func (o *ShipmentDomesticWithCarrierAccount) GetRateShopBy() string`

GetRateShopBy returns the RateShopBy field if non-nil, zero value otherwise.

### GetRateShopByOk

`func (o *ShipmentDomesticWithCarrierAccount) GetRateShopByOk() (*string, bool)`

GetRateShopByOk returns a tuple with the RateShopBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateShopBy

`func (o *ShipmentDomesticWithCarrierAccount) SetRateShopBy(v string)`

SetRateShopBy sets RateShopBy field to given value.

### HasRateShopBy

`func (o *ShipmentDomesticWithCarrierAccount) HasRateShopBy() bool`

HasRateShopBy returns a boolean if a field has been set.

### GetByCarrier

`func (o *ShipmentDomesticWithCarrierAccount) GetByCarrier() ShipmentDomesticWithCarrierAccountByCarrier`

GetByCarrier returns the ByCarrier field if non-nil, zero value otherwise.

### GetByCarrierOk

`func (o *ShipmentDomesticWithCarrierAccount) GetByCarrierOk() (*ShipmentDomesticWithCarrierAccountByCarrier, bool)`

GetByCarrierOk returns a tuple with the ByCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCarrier

`func (o *ShipmentDomesticWithCarrierAccount) SetByCarrier(v ShipmentDomesticWithCarrierAccountByCarrier)`

SetByCarrier sets ByCarrier field to given value.

### HasByCarrier

`func (o *ShipmentDomesticWithCarrierAccount) HasByCarrier() bool`

HasByCarrier returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentDomesticWithCarrierAccount) GetShipmentOptions() ShipmentDomesticWithCarrierAccountShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentDomesticWithCarrierAccount) GetShipmentOptionsOk() (*ShipmentDomesticWithCarrierAccountShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentDomesticWithCarrierAccount) SetShipmentOptions(v ShipmentDomesticWithCarrierAccountShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentDomesticWithCarrierAccount) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetReferences

`func (o *ShipmentDomesticWithCarrierAccount) GetReferences() ReferenceV2`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentDomesticWithCarrierAccount) GetReferencesOk() (*ReferenceV2, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentDomesticWithCarrierAccount) SetReferences(v ReferenceV2)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentDomesticWithCarrierAccount) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentDomesticWithCarrierAccount) GetMetadata() []ShipmentDomesticWithCarrierAccountMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentDomesticWithCarrierAccount) GetMetadataOk() (*[]ShipmentDomesticWithCarrierAccountMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentDomesticWithCarrierAccount) SetMetadata(v []ShipmentDomesticWithCarrierAccountMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentDomesticWithCarrierAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLabelSize

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelSize() string`

GetLabelSize returns the LabelSize field if non-nil, zero value otherwise.

### GetLabelSizeOk

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelSizeOk() (*string, bool)`

GetLabelSizeOk returns a tuple with the LabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSize

`func (o *ShipmentDomesticWithCarrierAccount) SetLabelSize(v string)`

SetLabelSize sets LabelSize field to given value.

### HasLabelSize

`func (o *ShipmentDomesticWithCarrierAccount) HasLabelSize() bool`

HasLabelSize returns a boolean if a field has been set.

### GetLabelType

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *ShipmentDomesticWithCarrierAccount) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.

### HasLabelType

`func (o *ShipmentDomesticWithCarrierAccount) HasLabelType() bool`

HasLabelType returns a boolean if a field has been set.

### GetLabelFormat

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *ShipmentDomesticWithCarrierAccount) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *ShipmentDomesticWithCarrierAccount) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *ShipmentDomesticWithCarrierAccount) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetPrinterAliasName

`func (o *ShipmentDomesticWithCarrierAccount) GetPrinterAliasName() string`

GetPrinterAliasName returns the PrinterAliasName field if non-nil, zero value otherwise.

### GetPrinterAliasNameOk

`func (o *ShipmentDomesticWithCarrierAccount) GetPrinterAliasNameOk() (*string, bool)`

GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterAliasName

`func (o *ShipmentDomesticWithCarrierAccount) SetPrinterAliasName(v string)`

SetPrinterAliasName sets PrinterAliasName field to given value.

### HasPrinterAliasName

`func (o *ShipmentDomesticWithCarrierAccount) HasPrinterAliasName() bool`

HasPrinterAliasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


