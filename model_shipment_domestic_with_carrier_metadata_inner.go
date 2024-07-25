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

// checks if the ShipmentDomesticWithCarrierMetadataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDomesticWithCarrierMetadataInner{}

// ShipmentDomesticWithCarrierMetadataInner The data that provides information about other interrelated/ required data.<br /> Here, metadata details consist of CostAccount Name and Value.
type ShipmentDomesticWithCarrierMetadataInner struct {
	// Name of the Cost Account which are linked to Shipment.
	Name *string `json:"name,omitempty"`
	// Indicates the value for the CostAccount.
	Value *string `json:"value,omitempty"`
}

// NewShipmentDomesticWithCarrierMetadataInner instantiates a new ShipmentDomesticWithCarrierMetadataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDomesticWithCarrierMetadataInner() *ShipmentDomesticWithCarrierMetadataInner {
	this := ShipmentDomesticWithCarrierMetadataInner{}
	return &this
}

// NewShipmentDomesticWithCarrierMetadataInnerWithDefaults instantiates a new ShipmentDomesticWithCarrierMetadataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDomesticWithCarrierMetadataInnerWithDefaults() *ShipmentDomesticWithCarrierMetadataInner {
	this := ShipmentDomesticWithCarrierMetadataInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierMetadataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierMetadataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierMetadataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipmentDomesticWithCarrierMetadataInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierMetadataInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierMetadataInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierMetadataInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ShipmentDomesticWithCarrierMetadataInner) SetValue(v string) {
	o.Value = &v
}

func (o ShipmentDomesticWithCarrierMetadataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDomesticWithCarrierMetadataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableShipmentDomesticWithCarrierMetadataInner struct {
	value *ShipmentDomesticWithCarrierMetadataInner
	isSet bool
}

func (v NullableShipmentDomesticWithCarrierMetadataInner) Get() *ShipmentDomesticWithCarrierMetadataInner {
	return v.value
}

func (v *NullableShipmentDomesticWithCarrierMetadataInner) Set(val *ShipmentDomesticWithCarrierMetadataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDomesticWithCarrierMetadataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDomesticWithCarrierMetadataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDomesticWithCarrierMetadataInner(val *ShipmentDomesticWithCarrierMetadataInner) *NullableShipmentDomesticWithCarrierMetadataInner {
	return &NullableShipmentDomesticWithCarrierMetadataInner{value: val, isSet: true}
}

func (v NullableShipmentDomesticWithCarrierMetadataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDomesticWithCarrierMetadataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


