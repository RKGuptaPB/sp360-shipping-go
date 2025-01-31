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

// checks if the GetAllShipmentsDataInnerParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllShipmentsDataInnerParcel{}

// GetAllShipmentsDataInnerParcel struct for GetAllShipmentsDataInnerParcel
type GetAllShipmentsDataInnerParcel struct {
	// Length is a part of Dimension object having highest numeric value out of three required parameters (length, width, and height) of Dimension. It helps determine a parcel’s girth.
	Length float32 `json:"length"`
	// Width is a part of Dimension object having lowest numeric value out of three required parameters of dimension (length, width, and height). This helps determine a parcel’s girth.
	Width float32 `json:"width"`
	// Height is a part of Dimension object where it helps determine a parcel’s girth.
	Height float32 `json:"height"`
	// DimUnit is a standard for measuring the physical quantities of specified dimension parameters.
	DimUnit string `json:"dimUnit"`
	// WeightUnit is a standard for measuring the physical quantities of specified weight.
	WeightUnit string `json:"weightUnit"`
	Weight float32 `json:"weight"`
}

type _GetAllShipmentsDataInnerParcel GetAllShipmentsDataInnerParcel

// NewGetAllShipmentsDataInnerParcel instantiates a new GetAllShipmentsDataInnerParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllShipmentsDataInnerParcel(length float32, width float32, height float32, dimUnit string, weightUnit string, weight float32) *GetAllShipmentsDataInnerParcel {
	this := GetAllShipmentsDataInnerParcel{}
	this.Length = length
	this.Width = width
	this.Height = height
	this.DimUnit = dimUnit
	this.WeightUnit = weightUnit
	this.Weight = weight
	return &this
}

// NewGetAllShipmentsDataInnerParcelWithDefaults instantiates a new GetAllShipmentsDataInnerParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllShipmentsDataInnerParcelWithDefaults() *GetAllShipmentsDataInnerParcel {
	this := GetAllShipmentsDataInnerParcel{}
	return &this
}

// GetLength returns the Length field value
func (o *GetAllShipmentsDataInnerParcel) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *GetAllShipmentsDataInnerParcel) SetLength(v float32) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *GetAllShipmentsDataInnerParcel) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *GetAllShipmentsDataInnerParcel) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *GetAllShipmentsDataInnerParcel) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *GetAllShipmentsDataInnerParcel) SetHeight(v float32) {
	o.Height = v
}

// GetDimUnit returns the DimUnit field value
func (o *GetAllShipmentsDataInnerParcel) GetDimUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimUnit
}

// GetDimUnitOk returns a tuple with the DimUnit field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetDimUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimUnit, true
}

// SetDimUnit sets field value
func (o *GetAllShipmentsDataInnerParcel) SetDimUnit(v string) {
	o.DimUnit = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *GetAllShipmentsDataInnerParcel) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *GetAllShipmentsDataInnerParcel) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetWeight returns the Weight field value
func (o *GetAllShipmentsDataInnerParcel) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *GetAllShipmentsDataInnerParcel) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *GetAllShipmentsDataInnerParcel) SetWeight(v float32) {
	o.Weight = v
}

func (o GetAllShipmentsDataInnerParcel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllShipmentsDataInnerParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["dimUnit"] = o.DimUnit
	toSerialize["weightUnit"] = o.WeightUnit
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *GetAllShipmentsDataInnerParcel) UnmarshalJSON(data []byte) (err error) {
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

	varGetAllShipmentsDataInnerParcel := _GetAllShipmentsDataInnerParcel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAllShipmentsDataInnerParcel)

	if err != nil {
		return err
	}

	*o = GetAllShipmentsDataInnerParcel(varGetAllShipmentsDataInnerParcel)

	return err
}

type NullableGetAllShipmentsDataInnerParcel struct {
	value *GetAllShipmentsDataInnerParcel
	isSet bool
}

func (v NullableGetAllShipmentsDataInnerParcel) Get() *GetAllShipmentsDataInnerParcel {
	return v.value
}

func (v *NullableGetAllShipmentsDataInnerParcel) Set(val *GetAllShipmentsDataInnerParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllShipmentsDataInnerParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllShipmentsDataInnerParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllShipmentsDataInnerParcel(val *GetAllShipmentsDataInnerParcel) *NullableGetAllShipmentsDataInnerParcel {
	return &NullableGetAllShipmentsDataInnerParcel{value: val, isSet: true}
}

func (v NullableGetAllShipmentsDataInnerParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllShipmentsDataInnerParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


