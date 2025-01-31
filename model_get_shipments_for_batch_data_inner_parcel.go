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

// checks if the GetShipmentsForBatchDataInnerParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentsForBatchDataInnerParcel{}

// GetShipmentsForBatchDataInnerParcel Parcel dimension, weight and unit of measurement details
type GetShipmentsForBatchDataInnerParcel struct {
	Dimension *GetShipmentsForBatchDataInnerParcelDimension `json:"dimension,omitempty"`
	Weight *GetShipmentsForBatchDataInnerParcelWeight `json:"weight,omitempty"`
}

// NewGetShipmentsForBatchDataInnerParcel instantiates a new GetShipmentsForBatchDataInnerParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentsForBatchDataInnerParcel() *GetShipmentsForBatchDataInnerParcel {
	this := GetShipmentsForBatchDataInnerParcel{}
	return &this
}

// NewGetShipmentsForBatchDataInnerParcelWithDefaults instantiates a new GetShipmentsForBatchDataInnerParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentsForBatchDataInnerParcelWithDefaults() *GetShipmentsForBatchDataInnerParcel {
	this := GetShipmentsForBatchDataInnerParcel{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerParcel) GetDimension() GetShipmentsForBatchDataInnerParcelDimension {
	if o == nil || IsNil(o.Dimension) {
		var ret GetShipmentsForBatchDataInnerParcelDimension
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerParcel) GetDimensionOk() (*GetShipmentsForBatchDataInnerParcelDimension, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerParcel) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given GetShipmentsForBatchDataInnerParcelDimension and assigns it to the Dimension field.
func (o *GetShipmentsForBatchDataInnerParcel) SetDimension(v GetShipmentsForBatchDataInnerParcelDimension) {
	o.Dimension = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerParcel) GetWeight() GetShipmentsForBatchDataInnerParcelWeight {
	if o == nil || IsNil(o.Weight) {
		var ret GetShipmentsForBatchDataInnerParcelWeight
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerParcel) GetWeightOk() (*GetShipmentsForBatchDataInnerParcelWeight, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerParcel) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given GetShipmentsForBatchDataInnerParcelWeight and assigns it to the Weight field.
func (o *GetShipmentsForBatchDataInnerParcel) SetWeight(v GetShipmentsForBatchDataInnerParcelWeight) {
	o.Weight = &v
}

func (o GetShipmentsForBatchDataInnerParcel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShipmentsForBatchDataInnerParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetShipmentsForBatchDataInnerParcel struct {
	value *GetShipmentsForBatchDataInnerParcel
	isSet bool
}

func (v NullableGetShipmentsForBatchDataInnerParcel) Get() *GetShipmentsForBatchDataInnerParcel {
	return v.value
}

func (v *NullableGetShipmentsForBatchDataInnerParcel) Set(val *GetShipmentsForBatchDataInnerParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentsForBatchDataInnerParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentsForBatchDataInnerParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentsForBatchDataInnerParcel(val *GetShipmentsForBatchDataInnerParcel) *NullableGetShipmentsForBatchDataInnerParcel {
	return &NullableGetShipmentsForBatchDataInnerParcel{value: val, isSet: true}
}

func (v NullableGetShipmentsForBatchDataInnerParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShipmentsForBatchDataInnerParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


