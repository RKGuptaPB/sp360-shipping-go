# CustomsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonForExport** | Pointer to **string** | &gt;-The reason the commodity is being exported. Valid values are: Valid Values: [GIFT COMMERCIAL_SAMPLE MERCHANDISE DOCUMENTS RETURNED_GOODS OTHER SOLD NOT_SOLD] | [optional] 
**CustomsDeclaredValue** | Pointer to **float64** | &gt;- Enter the value to declare in customs for the shipment. Enter the value in the currency specified in the currencyCode field. | [optional] 
**CurrencyCode** | **string** | &gt;-The type of currency used for the monetary values in this API request. Use three uppercase letters, per ISO 4217. For example, use USD for US Dollars, CAD for Canadian Dollars, and EUR for Euros. | 
**EELPFC** | Pointer to **string** | &gt;- A number provided by the Automated Export System (AES). This field is required if the item is valued at more than $2,500 USD per Schedule B export codes. | [optional] 
**CertificateNumber** | Pointer to **string** | The certificate number associated with the commodity. | [optional] 
**Comments** | Pointer to **string** | &gt;-Free form comments regarding the exported shipment entered by the shipper. | [optional] 
**FromCustomsReference** | Pointer to **string** | &gt;-Free form reference information provided by the requestor of the shipment. Depending on the carrier this information may or may not be rendered on the customs documents. | [optional] 
**ImporterCustomsReference** | Pointer to **string** | &gt;- A reference number used by the importer, such as a VAT number, PO number, or insured number. &#x60;Shipments to the EU&#x60;: Merchants shipping to the European Union (EU) from outside the EU must provide a VAT or IOSS number. Enter the number in this field and specify the reference type in the importerCustomsReferenceType field. | [optional] 
**InvoiceNumber** | Pointer to **string** | The commercial invoice number assigned by the exporter. | [optional] 
**LicenseNumber** | Pointer to **string** | The export license number associated with the commodity. | [optional] 
**DeclarationStatement** | Pointer to **string** | This is declaration statement from the Shipper for the items being shipped. | [optional] 
**ImporterCustomsReferenceType** | Pointer to **string** | &gt;-A reference type used by the importer to determine the type of importerCustomsReference. Valid Values are TAX_CODE, IMPORTER_CODE, VAT_NUMBER, IOSS_NUMBER. | [optional] 
**InsuredAmount** | Pointer to **float64** | Enter the insurance fee, if applicable. | [optional] 
**InsuredNumber** | Pointer to **string** | &gt;- If the sender wishes to ensure the contents, they complete an insurance receipt and affix the insured numbered label to the package. The insured number label is what this field represents. | [optional] 
**SdrValue** | Pointer to **float64** | &gt;-When an international parcel is insured, the insured value must be expressed in Special Drawing Rights values. E.g., $100 USD &#x3D; 66.87 SDR. | [optional] 

## Methods

### NewCustomsInfo

`func NewCustomsInfo(currencyCode string, ) *CustomsInfo`

NewCustomsInfo instantiates a new CustomsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomsInfoWithDefaults

`func NewCustomsInfoWithDefaults() *CustomsInfo`

NewCustomsInfoWithDefaults instantiates a new CustomsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonForExport

`func (o *CustomsInfo) GetReasonForExport() string`

GetReasonForExport returns the ReasonForExport field if non-nil, zero value otherwise.

### GetReasonForExportOk

`func (o *CustomsInfo) GetReasonForExportOk() (*string, bool)`

GetReasonForExportOk returns a tuple with the ReasonForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForExport

`func (o *CustomsInfo) SetReasonForExport(v string)`

SetReasonForExport sets ReasonForExport field to given value.

### HasReasonForExport

`func (o *CustomsInfo) HasReasonForExport() bool`

HasReasonForExport returns a boolean if a field has been set.

### GetCustomsDeclaredValue

`func (o *CustomsInfo) GetCustomsDeclaredValue() float64`

GetCustomsDeclaredValue returns the CustomsDeclaredValue field if non-nil, zero value otherwise.

### GetCustomsDeclaredValueOk

`func (o *CustomsInfo) GetCustomsDeclaredValueOk() (*float64, bool)`

GetCustomsDeclaredValueOk returns a tuple with the CustomsDeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsDeclaredValue

`func (o *CustomsInfo) SetCustomsDeclaredValue(v float64)`

SetCustomsDeclaredValue sets CustomsDeclaredValue field to given value.

### HasCustomsDeclaredValue

`func (o *CustomsInfo) HasCustomsDeclaredValue() bool`

HasCustomsDeclaredValue returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CustomsInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CustomsInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CustomsInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetEELPFC

`func (o *CustomsInfo) GetEELPFC() string`

GetEELPFC returns the EELPFC field if non-nil, zero value otherwise.

### GetEELPFCOk

`func (o *CustomsInfo) GetEELPFCOk() (*string, bool)`

GetEELPFCOk returns a tuple with the EELPFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEELPFC

`func (o *CustomsInfo) SetEELPFC(v string)`

SetEELPFC sets EELPFC field to given value.

### HasEELPFC

`func (o *CustomsInfo) HasEELPFC() bool`

HasEELPFC returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *CustomsInfo) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *CustomsInfo) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *CustomsInfo) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *CustomsInfo) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.

### GetComments

`func (o *CustomsInfo) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CustomsInfo) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CustomsInfo) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CustomsInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFromCustomsReference

`func (o *CustomsInfo) GetFromCustomsReference() string`

GetFromCustomsReference returns the FromCustomsReference field if non-nil, zero value otherwise.

### GetFromCustomsReferenceOk

`func (o *CustomsInfo) GetFromCustomsReferenceOk() (*string, bool)`

GetFromCustomsReferenceOk returns a tuple with the FromCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCustomsReference

`func (o *CustomsInfo) SetFromCustomsReference(v string)`

SetFromCustomsReference sets FromCustomsReference field to given value.

### HasFromCustomsReference

`func (o *CustomsInfo) HasFromCustomsReference() bool`

HasFromCustomsReference returns a boolean if a field has been set.

### GetImporterCustomsReference

`func (o *CustomsInfo) GetImporterCustomsReference() string`

GetImporterCustomsReference returns the ImporterCustomsReference field if non-nil, zero value otherwise.

### GetImporterCustomsReferenceOk

`func (o *CustomsInfo) GetImporterCustomsReferenceOk() (*string, bool)`

GetImporterCustomsReferenceOk returns a tuple with the ImporterCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterCustomsReference

`func (o *CustomsInfo) SetImporterCustomsReference(v string)`

SetImporterCustomsReference sets ImporterCustomsReference field to given value.

### HasImporterCustomsReference

`func (o *CustomsInfo) HasImporterCustomsReference() bool`

HasImporterCustomsReference returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *CustomsInfo) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CustomsInfo) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CustomsInfo) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *CustomsInfo) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *CustomsInfo) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *CustomsInfo) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *CustomsInfo) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *CustomsInfo) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### GetDeclarationStatement

`func (o *CustomsInfo) GetDeclarationStatement() string`

GetDeclarationStatement returns the DeclarationStatement field if non-nil, zero value otherwise.

### GetDeclarationStatementOk

`func (o *CustomsInfo) GetDeclarationStatementOk() (*string, bool)`

GetDeclarationStatementOk returns a tuple with the DeclarationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationStatement

`func (o *CustomsInfo) SetDeclarationStatement(v string)`

SetDeclarationStatement sets DeclarationStatement field to given value.

### HasDeclarationStatement

`func (o *CustomsInfo) HasDeclarationStatement() bool`

HasDeclarationStatement returns a boolean if a field has been set.

### GetImporterCustomsReferenceType

`func (o *CustomsInfo) GetImporterCustomsReferenceType() string`

GetImporterCustomsReferenceType returns the ImporterCustomsReferenceType field if non-nil, zero value otherwise.

### GetImporterCustomsReferenceTypeOk

`func (o *CustomsInfo) GetImporterCustomsReferenceTypeOk() (*string, bool)`

GetImporterCustomsReferenceTypeOk returns a tuple with the ImporterCustomsReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterCustomsReferenceType

`func (o *CustomsInfo) SetImporterCustomsReferenceType(v string)`

SetImporterCustomsReferenceType sets ImporterCustomsReferenceType field to given value.

### HasImporterCustomsReferenceType

`func (o *CustomsInfo) HasImporterCustomsReferenceType() bool`

HasImporterCustomsReferenceType returns a boolean if a field has been set.

### GetInsuredAmount

`func (o *CustomsInfo) GetInsuredAmount() float64`

GetInsuredAmount returns the InsuredAmount field if non-nil, zero value otherwise.

### GetInsuredAmountOk

`func (o *CustomsInfo) GetInsuredAmountOk() (*float64, bool)`

GetInsuredAmountOk returns a tuple with the InsuredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuredAmount

`func (o *CustomsInfo) SetInsuredAmount(v float64)`

SetInsuredAmount sets InsuredAmount field to given value.

### HasInsuredAmount

`func (o *CustomsInfo) HasInsuredAmount() bool`

HasInsuredAmount returns a boolean if a field has been set.

### GetInsuredNumber

`func (o *CustomsInfo) GetInsuredNumber() string`

GetInsuredNumber returns the InsuredNumber field if non-nil, zero value otherwise.

### GetInsuredNumberOk

`func (o *CustomsInfo) GetInsuredNumberOk() (*string, bool)`

GetInsuredNumberOk returns a tuple with the InsuredNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuredNumber

`func (o *CustomsInfo) SetInsuredNumber(v string)`

SetInsuredNumber sets InsuredNumber field to given value.

### HasInsuredNumber

`func (o *CustomsInfo) HasInsuredNumber() bool`

HasInsuredNumber returns a boolean if a field has been set.

### GetSdrValue

`func (o *CustomsInfo) GetSdrValue() float64`

GetSdrValue returns the SdrValue field if non-nil, zero value otherwise.

### GetSdrValueOk

`func (o *CustomsInfo) GetSdrValueOk() (*float64, bool)`

GetSdrValueOk returns a tuple with the SdrValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrValue

`func (o *CustomsInfo) SetSdrValue(v float64)`

SetSdrValue sets SdrValue field to given value.

### HasSdrValue

`func (o *CustomsInfo) HasSdrValue() bool`

HasSdrValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


