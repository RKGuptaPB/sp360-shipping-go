/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"encoding/json"
)

// checks if the DomesticShipmentResponseRateSpecialServicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomesticShipmentResponseRateSpecialServicesInner{}

// DomesticShipmentResponseRateSpecialServicesInner struct for DomesticShipmentResponseRateSpecialServicesInner
type DomesticShipmentResponseRateSpecialServicesInner struct {
	// The amount of the special service.
	Fee *float32 `json:"fee,omitempty"`
	// >-The parameters to set for the special service, such as an insurance value or a receipt-number format. This is required if the special service requires input parameters. If a special service does not require input parameters, you can either leave out the array or pass an empty array.
	InputParameters []DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner `json:"inputParameters,omitempty"`
	// A unique identifier associated to the Special Service, which depends on the carrier-based service.
	SpecialServiceId *string `json:"specialServiceId,omitempty"`
}

// NewDomesticShipmentResponseRateSpecialServicesInner instantiates a new DomesticShipmentResponseRateSpecialServicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomesticShipmentResponseRateSpecialServicesInner() *DomesticShipmentResponseRateSpecialServicesInner {
	this := DomesticShipmentResponseRateSpecialServicesInner{}
	return &this
}

// NewDomesticShipmentResponseRateSpecialServicesInnerWithDefaults instantiates a new DomesticShipmentResponseRateSpecialServicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomesticShipmentResponseRateSpecialServicesInnerWithDefaults() *DomesticShipmentResponseRateSpecialServicesInner {
	this := DomesticShipmentResponseRateSpecialServicesInner{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetFee() float32 {
	if o == nil || IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *DomesticShipmentResponseRateSpecialServicesInner) SetFee(v float32) {
	o.Fee = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetInputParameters() []DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner {
	if o == nil || IsNil(o.InputParameters) {
		var ret []DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetInputParametersOk() ([]DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner, bool) {
	if o == nil || IsNil(o.InputParameters) {
		return nil, false
	}
	return o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) HasInputParameters() bool {
	if o != nil && !IsNil(o.InputParameters) {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given []DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner and assigns it to the InputParameters field.
func (o *DomesticShipmentResponseRateSpecialServicesInner) SetInputParameters(v []DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) {
	o.InputParameters = v
}

// GetSpecialServiceId returns the SpecialServiceId field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetSpecialServiceId() string {
	if o == nil || IsNil(o.SpecialServiceId) {
		var ret string
		return ret
	}
	return *o.SpecialServiceId
}

// GetSpecialServiceIdOk returns a tuple with the SpecialServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) GetSpecialServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialServiceId) {
		return nil, false
	}
	return o.SpecialServiceId, true
}

// HasSpecialServiceId returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInner) HasSpecialServiceId() bool {
	if o != nil && !IsNil(o.SpecialServiceId) {
		return true
	}

	return false
}

// SetSpecialServiceId gets a reference to the given string and assigns it to the SpecialServiceId field.
func (o *DomesticShipmentResponseRateSpecialServicesInner) SetSpecialServiceId(v string) {
	o.SpecialServiceId = &v
}

func (o DomesticShipmentResponseRateSpecialServicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomesticShipmentResponseRateSpecialServicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.InputParameters) {
		toSerialize["inputParameters"] = o.InputParameters
	}
	if !IsNil(o.SpecialServiceId) {
		toSerialize["specialServiceId"] = o.SpecialServiceId
	}
	return toSerialize, nil
}

type NullableDomesticShipmentResponseRateSpecialServicesInner struct {
	value *DomesticShipmentResponseRateSpecialServicesInner
	isSet bool
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInner) Get() *DomesticShipmentResponseRateSpecialServicesInner {
	return v.value
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInner) Set(val *DomesticShipmentResponseRateSpecialServicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomesticShipmentResponseRateSpecialServicesInner(val *DomesticShipmentResponseRateSpecialServicesInner) *NullableDomesticShipmentResponseRateSpecialServicesInner {
	return &NullableDomesticShipmentResponseRateSpecialServicesInner{value: val, isSet: true}
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


