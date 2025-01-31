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

// checks if the GetStatusDetailedResponseLabelDetailsResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatusDetailedResponseLabelDetailsResultsInner{}

// GetStatusDetailedResponseLabelDetailsResultsInner struct for GetStatusDetailedResponseLabelDetailsResultsInner
type GetStatusDetailedResponseLabelDetailsResultsInner struct {
	Url *string `json:"url,omitempty"`
	//  This indicates the shipment identifiers.
	ShipmentIdentifiers []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner `json:"shipmentIdentifiers,omitempty"`
}

// NewGetStatusDetailedResponseLabelDetailsResultsInner instantiates a new GetStatusDetailedResponseLabelDetailsResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatusDetailedResponseLabelDetailsResultsInner() *GetStatusDetailedResponseLabelDetailsResultsInner {
	this := GetStatusDetailedResponseLabelDetailsResultsInner{}
	return &this
}

// NewGetStatusDetailedResponseLabelDetailsResultsInnerWithDefaults instantiates a new GetStatusDetailedResponseLabelDetailsResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatusDetailedResponseLabelDetailsResultsInnerWithDefaults() *GetStatusDetailedResponseLabelDetailsResultsInner {
	this := GetStatusDetailedResponseLabelDetailsResultsInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) SetUrl(v string) {
	o.Url = &v
}

// GetShipmentIdentifiers returns the ShipmentIdentifiers field value if set, zero value otherwise.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetShipmentIdentifiers() []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner {
	if o == nil || IsNil(o.ShipmentIdentifiers) {
		var ret []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner
		return ret
	}
	return o.ShipmentIdentifiers
}

// GetShipmentIdentifiersOk returns a tuple with the ShipmentIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetShipmentIdentifiersOk() ([]GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner, bool) {
	if o == nil || IsNil(o.ShipmentIdentifiers) {
		return nil, false
	}
	return o.ShipmentIdentifiers, true
}

// HasShipmentIdentifiers returns a boolean if a field has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) HasShipmentIdentifiers() bool {
	if o != nil && !IsNil(o.ShipmentIdentifiers) {
		return true
	}

	return false
}

// SetShipmentIdentifiers gets a reference to the given []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner and assigns it to the ShipmentIdentifiers field.
func (o *GetStatusDetailedResponseLabelDetailsResultsInner) SetShipmentIdentifiers(v []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) {
	o.ShipmentIdentifiers = v
}

func (o GetStatusDetailedResponseLabelDetailsResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatusDetailedResponseLabelDetailsResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ShipmentIdentifiers) {
		toSerialize["shipmentIdentifiers"] = o.ShipmentIdentifiers
	}
	return toSerialize, nil
}

type NullableGetStatusDetailedResponseLabelDetailsResultsInner struct {
	value *GetStatusDetailedResponseLabelDetailsResultsInner
	isSet bool
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInner) Get() *GetStatusDetailedResponseLabelDetailsResultsInner {
	return v.value
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInner) Set(val *GetStatusDetailedResponseLabelDetailsResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatusDetailedResponseLabelDetailsResultsInner(val *GetStatusDetailedResponseLabelDetailsResultsInner) *NullableGetStatusDetailedResponseLabelDetailsResultsInner {
	return &NullableGetStatusDetailedResponseLabelDetailsResultsInner{value: val, isSet: true}
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


