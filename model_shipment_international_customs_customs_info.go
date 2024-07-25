/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ShipmentInternationalCustomsCustomsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInternationalCustomsCustomsInfo{}

// ShipmentInternationalCustomsCustomsInfo This is additional customs information required along with item details.
type ShipmentInternationalCustomsCustomsInfo struct {
	// The reason the commodity is being exported.
	ReasonForExport string `json:"reasonForExport"`
	// Item value in mentioned currencyCode
	CustomsDeclaredValue float32 `json:"customsDeclaredValue"`
	// The currency used for declared value. Use three uppercase letters, per ISO 4217
	CurrencyCode string `json:"currencyCode"`
	// A number provided by the Automated Export System (AES). This is required if the item is valued at more than $2,500 USD per Schedule B export codes.
	EELPFC *string `json:"EELPFC,omitempty"`
	// The certificate number associated with the commodity.
	CertificateNumber *string `json:"certificateNumber,omitempty"`
	// Free-form comments regarding the exported shipment.
	Comments *string `json:"comments,omitempty"`
	// Free-form reference information provided by the requestor of the shipment. Depending on the carrier this information may or may not be rendered on the customs documents.
	FromCustomsReference *string `json:"fromCustomsReference,omitempty"`
	// A reference number used by the importer, such as a VAT number, PO number, or insured number.
	ImporterCustomsReference *string `json:"importerCustomsReference,omitempty"`
	// The commercial invoice number assigned by the exporter.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// The export license number associated with the commodity.
	LicenseNumber *string `json:"licenseNumber,omitempty"`
	// When an international parcel is insured, the insured value must be expressed in Special Drawing Rights values.
	SdrValue *float32 `json:"sdrValue,omitempty"`
}

type _ShipmentInternationalCustomsCustomsInfo ShipmentInternationalCustomsCustomsInfo

// NewShipmentInternationalCustomsCustomsInfo instantiates a new ShipmentInternationalCustomsCustomsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInternationalCustomsCustomsInfo(reasonForExport string, customsDeclaredValue float32, currencyCode string) *ShipmentInternationalCustomsCustomsInfo {
	this := ShipmentInternationalCustomsCustomsInfo{}
	this.ReasonForExport = reasonForExport
	this.CustomsDeclaredValue = customsDeclaredValue
	this.CurrencyCode = currencyCode
	return &this
}

// NewShipmentInternationalCustomsCustomsInfoWithDefaults instantiates a new ShipmentInternationalCustomsCustomsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInternationalCustomsCustomsInfoWithDefaults() *ShipmentInternationalCustomsCustomsInfo {
	this := ShipmentInternationalCustomsCustomsInfo{}
	return &this
}

// GetReasonForExport returns the ReasonForExport field value
func (o *ShipmentInternationalCustomsCustomsInfo) GetReasonForExport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonForExport
}

// GetReasonForExportOk returns a tuple with the ReasonForExport field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetReasonForExportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonForExport, true
}

// SetReasonForExport sets field value
func (o *ShipmentInternationalCustomsCustomsInfo) SetReasonForExport(v string) {
	o.ReasonForExport = v
}

// GetCustomsDeclaredValue returns the CustomsDeclaredValue field value
func (o *ShipmentInternationalCustomsCustomsInfo) GetCustomsDeclaredValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CustomsDeclaredValue
}

// GetCustomsDeclaredValueOk returns a tuple with the CustomsDeclaredValue field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetCustomsDeclaredValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomsDeclaredValue, true
}

// SetCustomsDeclaredValue sets field value
func (o *ShipmentInternationalCustomsCustomsInfo) SetCustomsDeclaredValue(v float32) {
	o.CustomsDeclaredValue = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *ShipmentInternationalCustomsCustomsInfo) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *ShipmentInternationalCustomsCustomsInfo) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetEELPFC returns the EELPFC field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetEELPFC() string {
	if o == nil || IsNil(o.EELPFC) {
		var ret string
		return ret
	}
	return *o.EELPFC
}

// GetEELPFCOk returns a tuple with the EELPFC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetEELPFCOk() (*string, bool) {
	if o == nil || IsNil(o.EELPFC) {
		return nil, false
	}
	return o.EELPFC, true
}

// HasEELPFC returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasEELPFC() bool {
	if o != nil && !IsNil(o.EELPFC) {
		return true
	}

	return false
}

// SetEELPFC gets a reference to the given string and assigns it to the EELPFC field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetEELPFC(v string) {
	o.EELPFC = &v
}

// GetCertificateNumber returns the CertificateNumber field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetCertificateNumber() string {
	if o == nil || IsNil(o.CertificateNumber) {
		var ret string
		return ret
	}
	return *o.CertificateNumber
}

// GetCertificateNumberOk returns a tuple with the CertificateNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetCertificateNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateNumber) {
		return nil, false
	}
	return o.CertificateNumber, true
}

// HasCertificateNumber returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasCertificateNumber() bool {
	if o != nil && !IsNil(o.CertificateNumber) {
		return true
	}

	return false
}

// SetCertificateNumber gets a reference to the given string and assigns it to the CertificateNumber field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetCertificateNumber(v string) {
	o.CertificateNumber = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetComments(v string) {
	o.Comments = &v
}

// GetFromCustomsReference returns the FromCustomsReference field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetFromCustomsReference() string {
	if o == nil || IsNil(o.FromCustomsReference) {
		var ret string
		return ret
	}
	return *o.FromCustomsReference
}

// GetFromCustomsReferenceOk returns a tuple with the FromCustomsReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetFromCustomsReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.FromCustomsReference) {
		return nil, false
	}
	return o.FromCustomsReference, true
}

// HasFromCustomsReference returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasFromCustomsReference() bool {
	if o != nil && !IsNil(o.FromCustomsReference) {
		return true
	}

	return false
}

// SetFromCustomsReference gets a reference to the given string and assigns it to the FromCustomsReference field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetFromCustomsReference(v string) {
	o.FromCustomsReference = &v
}

// GetImporterCustomsReference returns the ImporterCustomsReference field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetImporterCustomsReference() string {
	if o == nil || IsNil(o.ImporterCustomsReference) {
		var ret string
		return ret
	}
	return *o.ImporterCustomsReference
}

// GetImporterCustomsReferenceOk returns a tuple with the ImporterCustomsReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetImporterCustomsReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ImporterCustomsReference) {
		return nil, false
	}
	return o.ImporterCustomsReference, true
}

// HasImporterCustomsReference returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasImporterCustomsReference() bool {
	if o != nil && !IsNil(o.ImporterCustomsReference) {
		return true
	}

	return false
}

// SetImporterCustomsReference gets a reference to the given string and assigns it to the ImporterCustomsReference field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetImporterCustomsReference(v string) {
	o.ImporterCustomsReference = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetLicenseNumber returns the LicenseNumber field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetLicenseNumber() string {
	if o == nil || IsNil(o.LicenseNumber) {
		var ret string
		return ret
	}
	return *o.LicenseNumber
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetLicenseNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseNumber) {
		return nil, false
	}
	return o.LicenseNumber, true
}

// HasLicenseNumber returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasLicenseNumber() bool {
	if o != nil && !IsNil(o.LicenseNumber) {
		return true
	}

	return false
}

// SetLicenseNumber gets a reference to the given string and assigns it to the LicenseNumber field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetLicenseNumber(v string) {
	o.LicenseNumber = &v
}

// GetSdrValue returns the SdrValue field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsInfo) GetSdrValue() float32 {
	if o == nil || IsNil(o.SdrValue) {
		var ret float32
		return ret
	}
	return *o.SdrValue
}

// GetSdrValueOk returns a tuple with the SdrValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) GetSdrValueOk() (*float32, bool) {
	if o == nil || IsNil(o.SdrValue) {
		return nil, false
	}
	return o.SdrValue, true
}

// HasSdrValue returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsInfo) HasSdrValue() bool {
	if o != nil && !IsNil(o.SdrValue) {
		return true
	}

	return false
}

// SetSdrValue gets a reference to the given float32 and assigns it to the SdrValue field.
func (o *ShipmentInternationalCustomsCustomsInfo) SetSdrValue(v float32) {
	o.SdrValue = &v
}

func (o ShipmentInternationalCustomsCustomsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInternationalCustomsCustomsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reasonForExport"] = o.ReasonForExport
	toSerialize["customsDeclaredValue"] = o.CustomsDeclaredValue
	toSerialize["currencyCode"] = o.CurrencyCode
	if !IsNil(o.EELPFC) {
		toSerialize["EELPFC"] = o.EELPFC
	}
	if !IsNil(o.CertificateNumber) {
		toSerialize["certificateNumber"] = o.CertificateNumber
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.FromCustomsReference) {
		toSerialize["fromCustomsReference"] = o.FromCustomsReference
	}
	if !IsNil(o.ImporterCustomsReference) {
		toSerialize["importerCustomsReference"] = o.ImporterCustomsReference
	}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.LicenseNumber) {
		toSerialize["licenseNumber"] = o.LicenseNumber
	}
	if !IsNil(o.SdrValue) {
		toSerialize["sdrValue"] = o.SdrValue
	}
	return toSerialize, nil
}

func (o *ShipmentInternationalCustomsCustomsInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reasonForExport",
		"customsDeclaredValue",
		"currencyCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varShipmentInternationalCustomsCustomsInfo := _ShipmentInternationalCustomsCustomsInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentInternationalCustomsCustomsInfo)

	if err != nil {
		return err
	}

	*o = ShipmentInternationalCustomsCustomsInfo(varShipmentInternationalCustomsCustomsInfo)

	return err
}

type NullableShipmentInternationalCustomsCustomsInfo struct {
	value *ShipmentInternationalCustomsCustomsInfo
	isSet bool
}

func (v NullableShipmentInternationalCustomsCustomsInfo) Get() *ShipmentInternationalCustomsCustomsInfo {
	return v.value
}

func (v *NullableShipmentInternationalCustomsCustomsInfo) Set(val *ShipmentInternationalCustomsCustomsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInternationalCustomsCustomsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInternationalCustomsCustomsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInternationalCustomsCustomsInfo(val *ShipmentInternationalCustomsCustomsInfo) *NullableShipmentInternationalCustomsCustomsInfo {
	return &NullableShipmentInternationalCustomsCustomsInfo{value: val, isSet: true}
}

func (v NullableShipmentInternationalCustomsCustomsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInternationalCustomsCustomsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


