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

// checks if the SchedulePickupUPSRequestPickupOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupUPSRequestPickupOptionsInner{}

// SchedulePickupUPSRequestPickupOptionsInner struct for SchedulePickupUPSRequestPickupOptionsInner
type SchedulePickupUPSRequestPickupOptionsInner struct {
	// description
	Name *string `json:"name,omitempty"`
	// description
	Value *string `json:"value,omitempty"`
}

// NewSchedulePickupUPSRequestPickupOptionsInner instantiates a new SchedulePickupUPSRequestPickupOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupUPSRequestPickupOptionsInner() *SchedulePickupUPSRequestPickupOptionsInner {
	this := SchedulePickupUPSRequestPickupOptionsInner{}
	return &this
}

// NewSchedulePickupUPSRequestPickupOptionsInnerWithDefaults instantiates a new SchedulePickupUPSRequestPickupOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupUPSRequestPickupOptionsInnerWithDefaults() *SchedulePickupUPSRequestPickupOptionsInner {
	this := SchedulePickupUPSRequestPickupOptionsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequestPickupOptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequestPickupOptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequestPickupOptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchedulePickupUPSRequestPickupOptionsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequestPickupOptionsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequestPickupOptionsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequestPickupOptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SchedulePickupUPSRequestPickupOptionsInner) SetValue(v string) {
	o.Value = &v
}

func (o SchedulePickupUPSRequestPickupOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupUPSRequestPickupOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSchedulePickupUPSRequestPickupOptionsInner struct {
	value *SchedulePickupUPSRequestPickupOptionsInner
	isSet bool
}

func (v NullableSchedulePickupUPSRequestPickupOptionsInner) Get() *SchedulePickupUPSRequestPickupOptionsInner {
	return v.value
}

func (v *NullableSchedulePickupUPSRequestPickupOptionsInner) Set(val *SchedulePickupUPSRequestPickupOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupUPSRequestPickupOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupUPSRequestPickupOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupUPSRequestPickupOptionsInner(val *SchedulePickupUPSRequestPickupOptionsInner) *NullableSchedulePickupUPSRequestPickupOptionsInner {
	return &NullableSchedulePickupUPSRequestPickupOptionsInner{value: val, isSet: true}
}

func (v NullableSchedulePickupUPSRequestPickupOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupUPSRequestPickupOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


