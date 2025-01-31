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

// checks if the GetShipmentsForBatchDataInnerParcelWeight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentsForBatchDataInnerParcelWeight{}

// GetShipmentsForBatchDataInnerParcelWeight struct for GetShipmentsForBatchDataInnerParcelWeight
type GetShipmentsForBatchDataInnerParcelWeight struct {
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
	Weight *float32 `json:"weight,omitempty"`
}

// NewGetShipmentsForBatchDataInnerParcelWeight instantiates a new GetShipmentsForBatchDataInnerParcelWeight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentsForBatchDataInnerParcelWeight() *GetShipmentsForBatchDataInnerParcelWeight {
	this := GetShipmentsForBatchDataInnerParcelWeight{}
	return &this
}

// NewGetShipmentsForBatchDataInnerParcelWeightWithDefaults instantiates a new GetShipmentsForBatchDataInnerParcelWeight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentsForBatchDataInnerParcelWeightWithDefaults() *GetShipmentsForBatchDataInnerParcelWeight {
	this := GetShipmentsForBatchDataInnerParcelWeight{}
	return &this
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerParcelWeight) GetUnitOfMeasurement() string {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerParcelWeight) GetUnitOfMeasurementOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		return nil, false
	}
	return o.UnitOfMeasurement, true
}

// HasUnitOfMeasurement returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerParcelWeight) HasUnitOfMeasurement() bool {
	if o != nil && !IsNil(o.UnitOfMeasurement) {
		return true
	}

	return false
}

// SetUnitOfMeasurement gets a reference to the given string and assigns it to the UnitOfMeasurement field.
func (o *GetShipmentsForBatchDataInnerParcelWeight) SetUnitOfMeasurement(v string) {
	o.UnitOfMeasurement = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerParcelWeight) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerParcelWeight) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerParcelWeight) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *GetShipmentsForBatchDataInnerParcelWeight) SetWeight(v float32) {
	o.Weight = &v
}

func (o GetShipmentsForBatchDataInnerParcelWeight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShipmentsForBatchDataInnerParcelWeight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnitOfMeasurement) {
		toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGetShipmentsForBatchDataInnerParcelWeight struct {
	value *GetShipmentsForBatchDataInnerParcelWeight
	isSet bool
}

func (v NullableGetShipmentsForBatchDataInnerParcelWeight) Get() *GetShipmentsForBatchDataInnerParcelWeight {
	return v.value
}

func (v *NullableGetShipmentsForBatchDataInnerParcelWeight) Set(val *GetShipmentsForBatchDataInnerParcelWeight) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentsForBatchDataInnerParcelWeight) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentsForBatchDataInnerParcelWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentsForBatchDataInnerParcelWeight(val *GetShipmentsForBatchDataInnerParcelWeight) *NullableGetShipmentsForBatchDataInnerParcelWeight {
	return &NullableGetShipmentsForBatchDataInnerParcelWeight{value: val, isSet: true}
}

func (v NullableGetShipmentsForBatchDataInnerParcelWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShipmentsForBatchDataInnerParcelWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


