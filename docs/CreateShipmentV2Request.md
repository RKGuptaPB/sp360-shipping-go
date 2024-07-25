# CreateShipmentV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | [**FromAddressV2**](FromAddressV2.md) |  | 
**ToAddress** | [**ToAddressV2**](ToAddressV2.md) |  | 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection.&lt;br /&gt; ParcelType can have categories like Package, Envelopes, Paks, Boxes, Tube, etc.  | [optional] 
**RateShopBy** | Pointer to **string** | By RateShop is a way to restrict the carrier services returned in a Rate Shop experience.&lt;br /&gt;  Without this Rate Shop Group, which is attached to an Enterprise or Location, the carrier services shown to the end user, are all that will be available for each carrier, based on Addresses, Weight, Dimensions, and Package Type. | [optional] 
**ByCarrier** | Pointer to [**ShipmentDomesticWithCarrierByCarrier**](ShipmentDomesticWithCarrierByCarrier.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentDomesticByRulesetCarrierShipmentOptions**](ShipmentDomesticByRulesetCarrierShipmentOptions.md) |  | [optional] 
**References** | Pointer to [**ReferenceV2**](ReferenceV2.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentDomesticByRulesetCarrierMetadataInner**](ShipmentDomesticByRulesetCarrierMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here.&lt;br /&gt; For now, &#39;Cost Account Name&#39; is supported. | [optional] 
**LabelSize** | Pointer to **string** | Defines the label size of the Shipment, that is, the Shipping Label is available in different Doc Size. | [optional] 
**LabelType** | Pointer to **string** | Defines the type of the Shipment. | [optional] 
**LabelFormat** | Pointer to **string** | Defines the file/format in which the label is printed. | [optional] 
**PrinterAliasName** | Pointer to **string** | Refers to a printer connected (directly or via network) to a computer. | [optional] 
**ByRateGroup** | Pointer to [**ShipmentDomesticByRateGroupByRateGroup**](ShipmentDomesticByRateGroupByRateGroup.md) |  | [optional] 
**ByRuleSet** | Pointer to [**ShipmentDomesticByRulesetCarrierByRuleSet**](ShipmentDomesticByRulesetCarrierByRuleSet.md) |  | [optional] 

## Methods

### NewCreateShipmentV2Request

`func NewCreateShipmentV2Request(fromAddress FromAddressV2, toAddress ToAddressV2, ) *CreateShipmentV2Request`

NewCreateShipmentV2Request instantiates a new CreateShipmentV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipmentV2RequestWithDefaults

`func NewCreateShipmentV2RequestWithDefaults() *CreateShipmentV2Request`

NewCreateShipmentV2RequestWithDefaults instantiates a new CreateShipmentV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *CreateShipmentV2Request) GetFromAddress() FromAddressV2`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateShipmentV2Request) GetFromAddressOk() (*FromAddressV2, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateShipmentV2Request) SetFromAddress(v FromAddressV2)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *CreateShipmentV2Request) GetToAddress() ToAddressV2`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateShipmentV2Request) GetToAddressOk() (*ToAddressV2, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateShipmentV2Request) SetToAddress(v ToAddressV2)`

SetToAddress sets ToAddress field to given value.


### GetParcel

`func (o *CreateShipmentV2Request) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *CreateShipmentV2Request) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *CreateShipmentV2Request) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *CreateShipmentV2Request) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelType

`func (o *CreateShipmentV2Request) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateShipmentV2Request) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateShipmentV2Request) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *CreateShipmentV2Request) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateShopBy

`func (o *CreateShipmentV2Request) GetRateShopBy() string`

GetRateShopBy returns the RateShopBy field if non-nil, zero value otherwise.

### GetRateShopByOk

`func (o *CreateShipmentV2Request) GetRateShopByOk() (*string, bool)`

GetRateShopByOk returns a tuple with the RateShopBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateShopBy

`func (o *CreateShipmentV2Request) SetRateShopBy(v string)`

SetRateShopBy sets RateShopBy field to given value.

### HasRateShopBy

`func (o *CreateShipmentV2Request) HasRateShopBy() bool`

HasRateShopBy returns a boolean if a field has been set.

### GetByCarrier

`func (o *CreateShipmentV2Request) GetByCarrier() ShipmentDomesticWithCarrierByCarrier`

GetByCarrier returns the ByCarrier field if non-nil, zero value otherwise.

### GetByCarrierOk

`func (o *CreateShipmentV2Request) GetByCarrierOk() (*ShipmentDomesticWithCarrierByCarrier, bool)`

GetByCarrierOk returns a tuple with the ByCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCarrier

`func (o *CreateShipmentV2Request) SetByCarrier(v ShipmentDomesticWithCarrierByCarrier)`

SetByCarrier sets ByCarrier field to given value.

### HasByCarrier

`func (o *CreateShipmentV2Request) HasByCarrier() bool`

HasByCarrier returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *CreateShipmentV2Request) GetShipmentOptions() ShipmentDomesticByRulesetCarrierShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *CreateShipmentV2Request) GetShipmentOptionsOk() (*ShipmentDomesticByRulesetCarrierShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *CreateShipmentV2Request) SetShipmentOptions(v ShipmentDomesticByRulesetCarrierShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *CreateShipmentV2Request) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetReferences

`func (o *CreateShipmentV2Request) GetReferences() ReferenceV2`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CreateShipmentV2Request) GetReferencesOk() (*ReferenceV2, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CreateShipmentV2Request) SetReferences(v ReferenceV2)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CreateShipmentV2Request) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateShipmentV2Request) GetMetadata() []ShipmentDomesticByRulesetCarrierMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateShipmentV2Request) GetMetadataOk() (*[]ShipmentDomesticByRulesetCarrierMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateShipmentV2Request) SetMetadata(v []ShipmentDomesticByRulesetCarrierMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateShipmentV2Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLabelSize

`func (o *CreateShipmentV2Request) GetLabelSize() string`

GetLabelSize returns the LabelSize field if non-nil, zero value otherwise.

### GetLabelSizeOk

`func (o *CreateShipmentV2Request) GetLabelSizeOk() (*string, bool)`

GetLabelSizeOk returns a tuple with the LabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSize

`func (o *CreateShipmentV2Request) SetLabelSize(v string)`

SetLabelSize sets LabelSize field to given value.

### HasLabelSize

`func (o *CreateShipmentV2Request) HasLabelSize() bool`

HasLabelSize returns a boolean if a field has been set.

### GetLabelType

`func (o *CreateShipmentV2Request) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *CreateShipmentV2Request) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *CreateShipmentV2Request) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.

### HasLabelType

`func (o *CreateShipmentV2Request) HasLabelType() bool`

HasLabelType returns a boolean if a field has been set.

### GetLabelFormat

`func (o *CreateShipmentV2Request) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *CreateShipmentV2Request) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *CreateShipmentV2Request) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *CreateShipmentV2Request) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetPrinterAliasName

`func (o *CreateShipmentV2Request) GetPrinterAliasName() string`

GetPrinterAliasName returns the PrinterAliasName field if non-nil, zero value otherwise.

### GetPrinterAliasNameOk

`func (o *CreateShipmentV2Request) GetPrinterAliasNameOk() (*string, bool)`

GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterAliasName

`func (o *CreateShipmentV2Request) SetPrinterAliasName(v string)`

SetPrinterAliasName sets PrinterAliasName field to given value.

### HasPrinterAliasName

`func (o *CreateShipmentV2Request) HasPrinterAliasName() bool`

HasPrinterAliasName returns a boolean if a field has been set.

### GetByRateGroup

`func (o *CreateShipmentV2Request) GetByRateGroup() ShipmentDomesticByRateGroupByRateGroup`

GetByRateGroup returns the ByRateGroup field if non-nil, zero value otherwise.

### GetByRateGroupOk

`func (o *CreateShipmentV2Request) GetByRateGroupOk() (*ShipmentDomesticByRateGroupByRateGroup, bool)`

GetByRateGroupOk returns a tuple with the ByRateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByRateGroup

`func (o *CreateShipmentV2Request) SetByRateGroup(v ShipmentDomesticByRateGroupByRateGroup)`

SetByRateGroup sets ByRateGroup field to given value.

### HasByRateGroup

`func (o *CreateShipmentV2Request) HasByRateGroup() bool`

HasByRateGroup returns a boolean if a field has been set.

### GetByRuleSet

`func (o *CreateShipmentV2Request) GetByRuleSet() ShipmentDomesticByRulesetCarrierByRuleSet`

GetByRuleSet returns the ByRuleSet field if non-nil, zero value otherwise.

### GetByRuleSetOk

`func (o *CreateShipmentV2Request) GetByRuleSetOk() (*ShipmentDomesticByRulesetCarrierByRuleSet, bool)`

GetByRuleSetOk returns a tuple with the ByRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByRuleSet

`func (o *CreateShipmentV2Request) SetByRuleSet(v ShipmentDomesticByRulesetCarrierByRuleSet)`

SetByRuleSet sets ByRuleSet field to given value.

### HasByRuleSet

`func (o *CreateShipmentV2Request) HasByRuleSet() bool`

HasByRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


