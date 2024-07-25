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

// checks if the PrinterMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrinterMappingRequest{}

// PrinterMappingRequest struct for PrinterMappingRequest
type PrinterMappingRequest struct {
	// Refers to a printer connected (directly or via network) to a computer.
	Alias string `json:"alias"`
	// Indicates the Device Serial number.
	SerialNumber string `json:"serialNumber"`
	// The Printer name which is used for mapping.
	PrinterName string `json:"printerName"`
}

type _PrinterMappingRequest PrinterMappingRequest

// NewPrinterMappingRequest instantiates a new PrinterMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrinterMappingRequest(alias string, serialNumber string, printerName string) *PrinterMappingRequest {
	this := PrinterMappingRequest{}
	this.Alias = alias
	this.SerialNumber = serialNumber
	this.PrinterName = printerName
	return &this
}

// NewPrinterMappingRequestWithDefaults instantiates a new PrinterMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrinterMappingRequestWithDefaults() *PrinterMappingRequest {
	this := PrinterMappingRequest{}
	return &this
}

// GetAlias returns the Alias field value
func (o *PrinterMappingRequest) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *PrinterMappingRequest) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *PrinterMappingRequest) SetAlias(v string) {
	o.Alias = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *PrinterMappingRequest) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *PrinterMappingRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *PrinterMappingRequest) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetPrinterName returns the PrinterName field value
func (o *PrinterMappingRequest) GetPrinterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrinterName
}

// GetPrinterNameOk returns a tuple with the PrinterName field value
// and a boolean to check if the value has been set.
func (o *PrinterMappingRequest) GetPrinterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrinterName, true
}

// SetPrinterName sets field value
func (o *PrinterMappingRequest) SetPrinterName(v string) {
	o.PrinterName = v
}

func (o PrinterMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrinterMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["printerName"] = o.PrinterName
	return toSerialize, nil
}

func (o *PrinterMappingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"serialNumber",
		"printerName",
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

	varPrinterMappingRequest := _PrinterMappingRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrinterMappingRequest)

	if err != nil {
		return err
	}

	*o = PrinterMappingRequest(varPrinterMappingRequest)

	return err
}

type NullablePrinterMappingRequest struct {
	value *PrinterMappingRequest
	isSet bool
}

func (v NullablePrinterMappingRequest) Get() *PrinterMappingRequest {
	return v.value
}

func (v *NullablePrinterMappingRequest) Set(val *PrinterMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrinterMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrinterMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrinterMappingRequest(val *PrinterMappingRequest) *NullablePrinterMappingRequest {
	return &NullablePrinterMappingRequest{value: val, isSet: true}
}

func (v NullablePrinterMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrinterMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


