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

// checks if the ReturnLabelParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLabelParcel{}

// ReturnLabelParcel struct for ReturnLabelParcel
type ReturnLabelParcel struct {
	// Length is always the greatest of the three dimensions. The other two dimensions are used in the calculation of the girth.
	Length int32 `json:"length"`
	// There is no strict rule as to which element is the width or the height, but typically, by convention the width is the second greatest dimension of a parcel.
	Width int32 `json:"width"`
	// By convention the height is the smallest dimension of the parcel.
	Height int32 `json:"height"`
	// DimUnit is a standard for measuring the physical quantities of specified dimension parameters.
	DimUnit string `json:"dimUnit"`
	// WeightUnit is a standard for measuring the physical quantities of specified weight.
	WeightUnit string `json:"weightUnit"`
	// Weight is the measure of how heavy an object is.
	Weight float32 `json:"weight"`
}

type _ReturnLabelParcel ReturnLabelParcel

// NewReturnLabelParcel instantiates a new ReturnLabelParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLabelParcel(length int32, width int32, height int32, dimUnit string, weightUnit string, weight float32) *ReturnLabelParcel {
	this := ReturnLabelParcel{}
	this.Length = length
	this.Width = width
	this.Height = height
	this.DimUnit = dimUnit
	this.WeightUnit = weightUnit
	this.Weight = weight
	return &this
}

// NewReturnLabelParcelWithDefaults instantiates a new ReturnLabelParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLabelParcelWithDefaults() *ReturnLabelParcel {
	this := ReturnLabelParcel{}
	return &this
}

// GetLength returns the Length field value
func (o *ReturnLabelParcel) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *ReturnLabelParcel) SetLength(v int32) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *ReturnLabelParcel) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ReturnLabelParcel) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *ReturnLabelParcel) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ReturnLabelParcel) SetHeight(v int32) {
	o.Height = v
}

// GetDimUnit returns the DimUnit field value
func (o *ReturnLabelParcel) GetDimUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimUnit
}

// GetDimUnitOk returns a tuple with the DimUnit field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetDimUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimUnit, true
}

// SetDimUnit sets field value
func (o *ReturnLabelParcel) SetDimUnit(v string) {
	o.DimUnit = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *ReturnLabelParcel) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *ReturnLabelParcel) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetWeight returns the Weight field value
func (o *ReturnLabelParcel) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelParcel) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ReturnLabelParcel) SetWeight(v float32) {
	o.Weight = v
}

func (o ReturnLabelParcel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLabelParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["dimUnit"] = o.DimUnit
	toSerialize["weightUnit"] = o.WeightUnit
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *ReturnLabelParcel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"length",
		"width",
		"height",
		"dimUnit",
		"weightUnit",
		"weight",
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

	varReturnLabelParcel := _ReturnLabelParcel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReturnLabelParcel)

	if err != nil {
		return err
	}

	*o = ReturnLabelParcel(varReturnLabelParcel)

	return err
}

type NullableReturnLabelParcel struct {
	value *ReturnLabelParcel
	isSet bool
}

func (v NullableReturnLabelParcel) Get() *ReturnLabelParcel {
	return v.value
}

func (v *NullableReturnLabelParcel) Set(val *ReturnLabelParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLabelParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLabelParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLabelParcel(val *ReturnLabelParcel) *NullableReturnLabelParcel {
	return &NullableReturnLabelParcel{value: val, isSet: true}
}

func (v NullableReturnLabelParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLabelParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


