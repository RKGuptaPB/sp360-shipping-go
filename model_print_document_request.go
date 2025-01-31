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

// checks if the PrintDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrintDocumentRequest{}

// PrintDocumentRequest struct for PrintDocumentRequest
type PrintDocumentRequest struct {
	// It refers to a Printer connected (directly or via network) to a Computer.
	PrinterAliasName string `json:"printerAliasName"`
	// Content/Identifier of document e.g. DOCUMENT_REFERECE_ID. Actual document name e.g. abc.pdf. [IN] i.e base64 string, URL, file path
	Data string `json:"data"`
	// Data Type of the document e.g. DOCUMENT_REFERENCE. [IN/OUT]
	DataType string `json:"dataType"`
	// The format of the document file the print takes.
	DocumentType string `json:"documentType"`
	// It refers to a form size
	FormName string `json:"formName"`
	// The orientation of the document layout: Portrait or Landscape.
	Orientation *string `json:"orientation,omitempty"`
	Reference *PrintDocumentRequestReference `json:"reference,omitempty"`
}

type _PrintDocumentRequest PrintDocumentRequest

// NewPrintDocumentRequest instantiates a new PrintDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintDocumentRequest(printerAliasName string, data string, dataType string, documentType string, formName string) *PrintDocumentRequest {
	this := PrintDocumentRequest{}
	this.PrinterAliasName = printerAliasName
	this.Data = data
	this.DataType = dataType
	this.DocumentType = documentType
	this.FormName = formName
	return &this
}

// NewPrintDocumentRequestWithDefaults instantiates a new PrintDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintDocumentRequestWithDefaults() *PrintDocumentRequest {
	this := PrintDocumentRequest{}
	return &this
}

// GetPrinterAliasName returns the PrinterAliasName field value
func (o *PrintDocumentRequest) GetPrinterAliasName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrinterAliasName
}

// GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field value
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetPrinterAliasNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrinterAliasName, true
}

// SetPrinterAliasName sets field value
func (o *PrintDocumentRequest) SetPrinterAliasName(v string) {
	o.PrinterAliasName = v
}

// GetData returns the Data field value
func (o *PrintDocumentRequest) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PrintDocumentRequest) SetData(v string) {
	o.Data = v
}

// GetDataType returns the DataType field value
func (o *PrintDocumentRequest) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *PrintDocumentRequest) SetDataType(v string) {
	o.DataType = v
}

// GetDocumentType returns the DocumentType field value
func (o *PrintDocumentRequest) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *PrintDocumentRequest) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetFormName returns the FormName field value
func (o *PrintDocumentRequest) GetFormName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetFormNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormName, true
}

// SetFormName sets field value
func (o *PrintDocumentRequest) SetFormName(v string) {
	o.FormName = v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *PrintDocumentRequest) GetOrientation() string {
	if o == nil || IsNil(o.Orientation) {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetOrientationOk() (*string, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *PrintDocumentRequest) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *PrintDocumentRequest) SetOrientation(v string) {
	o.Orientation = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PrintDocumentRequest) GetReference() PrintDocumentRequestReference {
	if o == nil || IsNil(o.Reference) {
		var ret PrintDocumentRequestReference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintDocumentRequest) GetReferenceOk() (*PrintDocumentRequestReference, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PrintDocumentRequest) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given PrintDocumentRequestReference and assigns it to the Reference field.
func (o *PrintDocumentRequest) SetReference(v PrintDocumentRequestReference) {
	o.Reference = &v
}

func (o PrintDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrintDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["printerAliasName"] = o.PrinterAliasName
	toSerialize["data"] = o.Data
	toSerialize["dataType"] = o.DataType
	toSerialize["documentType"] = o.DocumentType
	toSerialize["formName"] = o.FormName
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

func (o *PrintDocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"printerAliasName",
		"data",
		"dataType",
		"documentType",
		"formName",
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

	varPrintDocumentRequest := _PrintDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrintDocumentRequest)

	if err != nil {
		return err
	}

	*o = PrintDocumentRequest(varPrintDocumentRequest)

	return err
}

type NullablePrintDocumentRequest struct {
	value *PrintDocumentRequest
	isSet bool
}

func (v NullablePrintDocumentRequest) Get() *PrintDocumentRequest {
	return v.value
}

func (v *NullablePrintDocumentRequest) Set(val *PrintDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintDocumentRequest(val *PrintDocumentRequest) *NullablePrintDocumentRequest {
	return &NullablePrintDocumentRequest{value: val, isSet: true}
}

func (v NullablePrintDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


