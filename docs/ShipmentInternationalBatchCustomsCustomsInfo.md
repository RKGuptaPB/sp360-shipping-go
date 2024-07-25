# ShipmentInternationalBatchCustomsCustomsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonForExport** | **string** | The reason the commodity is being exported. Valid values are | 
**CustomsDeclaredValue** | **float32** | Item value in mentioned currencyCode | 
**CurrencyCode** | **string** | The currency used for declared value. Use three uppercase letters, per ISO 4217 | 
**EELPFC** | Pointer to **string** | A number provided by the Automated Export System (AES). This is required if the item is valued at more than $2,500 USD per Schedule B export codes. | [optional] 
**CertificateNumber** | Pointer to **string** | The certificate number associated with the commodity. | [optional] 
**Comments** | Pointer to **string** | Free-form comments regarding the exported shipment. | [optional] 
**FromCustomsReference** | Pointer to **string** | Free-form reference information provided by the requestor of the shipment. Depending on the carrier this information may or may not be rendered on the customs documents. | [optional] 
**ImporterCustomsReference** | Pointer to **string** | A reference number used by the importer, such as a VAT number, PO number, or insured number. | [optional] 
**InvoiceNumber** | Pointer to **string** | The commercial invoice number assigned by the exporter. | [optional] 
**LicenseNumber** | Pointer to **string** | The export license number associated with the commodity. | [optional] 
**SdrValue** | Pointer to **float32** | When an international parcel is insured, the insured value must be expressed in Special Drawing Rights values | [optional] 

## Methods

### NewShipmentInternationalBatchCustomsCustomsInfo

`func NewShipmentInternationalBatchCustomsCustomsInfo(reasonForExport string, customsDeclaredValue float32, currencyCode string, ) *ShipmentInternationalBatchCustomsCustomsInfo`

NewShipmentInternationalBatchCustomsCustomsInfo instantiates a new ShipmentInternationalBatchCustomsCustomsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalBatchCustomsCustomsInfoWithDefaults

`func NewShipmentInternationalBatchCustomsCustomsInfoWithDefaults() *ShipmentInternationalBatchCustomsCustomsInfo`

NewShipmentInternationalBatchCustomsCustomsInfoWithDefaults instantiates a new ShipmentInternationalBatchCustomsCustomsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonForExport

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetReasonForExport() string`

GetReasonForExport returns the ReasonForExport field if non-nil, zero value otherwise.

### GetReasonForExportOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetReasonForExportOk() (*string, bool)`

GetReasonForExportOk returns a tuple with the ReasonForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForExport

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetReasonForExport(v string)`

SetReasonForExport sets ReasonForExport field to given value.


### GetCustomsDeclaredValue

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCustomsDeclaredValue() float32`

GetCustomsDeclaredValue returns the CustomsDeclaredValue field if non-nil, zero value otherwise.

### GetCustomsDeclaredValueOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCustomsDeclaredValueOk() (*float32, bool)`

GetCustomsDeclaredValueOk returns a tuple with the CustomsDeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsDeclaredValue

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetCustomsDeclaredValue(v float32)`

SetCustomsDeclaredValue sets CustomsDeclaredValue field to given value.


### GetCurrencyCode

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetEELPFC

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetEELPFC() string`

GetEELPFC returns the EELPFC field if non-nil, zero value otherwise.

### GetEELPFCOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetEELPFCOk() (*string, bool)`

GetEELPFCOk returns a tuple with the EELPFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEELPFC

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetEELPFC(v string)`

SetEELPFC sets EELPFC field to given value.

### HasEELPFC

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasEELPFC() bool`

HasEELPFC returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.

### GetComments

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFromCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetFromCustomsReference() string`

GetFromCustomsReference returns the FromCustomsReference field if non-nil, zero value otherwise.

### GetFromCustomsReferenceOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetFromCustomsReferenceOk() (*string, bool)`

GetFromCustomsReferenceOk returns a tuple with the FromCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetFromCustomsReference(v string)`

SetFromCustomsReference sets FromCustomsReference field to given value.

### HasFromCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasFromCustomsReference() bool`

HasFromCustomsReference returns a boolean if a field has been set.

### GetImporterCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetImporterCustomsReference() string`

GetImporterCustomsReference returns the ImporterCustomsReference field if non-nil, zero value otherwise.

### GetImporterCustomsReferenceOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetImporterCustomsReferenceOk() (*string, bool)`

GetImporterCustomsReferenceOk returns a tuple with the ImporterCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetImporterCustomsReference(v string)`

SetImporterCustomsReference sets ImporterCustomsReference field to given value.

### HasImporterCustomsReference

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasImporterCustomsReference() bool`

HasImporterCustomsReference returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### GetSdrValue

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetSdrValue() float32`

GetSdrValue returns the SdrValue field if non-nil, zero value otherwise.

### GetSdrValueOk

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) GetSdrValueOk() (*float32, bool)`

GetSdrValueOk returns a tuple with the SdrValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrValue

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) SetSdrValue(v float32)`

SetSdrValue sets SdrValue field to given value.

### HasSdrValue

`func (o *ShipmentInternationalBatchCustomsCustomsInfo) HasSdrValue() bool`

HasSdrValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


