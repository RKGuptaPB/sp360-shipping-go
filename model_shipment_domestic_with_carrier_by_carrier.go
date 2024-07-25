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

// checks if the ShipmentDomesticWithCarrierByCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDomesticWithCarrierByCarrier{}

// ShipmentDomesticWithCarrierByCarrier The shipment is grouped by Carrier and their Service.
type ShipmentDomesticWithCarrierByCarrier struct {
	// A unique identifier associated with the specific carrier, i.e. CarrierID, which must be valid.
	Carrier *string `json:"carrier,omitempty"`
	// Indicates a unique identifier associated with the carrier specific service, that is ServiceID, which must be valid.
	Service *string `json:"service,omitempty"`
}

// NewShipmentDomesticWithCarrierByCarrier instantiates a new ShipmentDomesticWithCarrierByCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDomesticWithCarrierByCarrier() *ShipmentDomesticWithCarrierByCarrier {
	this := ShipmentDomesticWithCarrierByCarrier{}
	return &this
}

// NewShipmentDomesticWithCarrierByCarrierWithDefaults instantiates a new ShipmentDomesticWithCarrierByCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDomesticWithCarrierByCarrierWithDefaults() *ShipmentDomesticWithCarrierByCarrier {
	this := ShipmentDomesticWithCarrierByCarrier{}
	return &this
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierByCarrier) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierByCarrier) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierByCarrier) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *ShipmentDomesticWithCarrierByCarrier) SetCarrier(v string) {
	o.Carrier = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierByCarrier) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierByCarrier) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierByCarrier) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ShipmentDomesticWithCarrierByCarrier) SetService(v string) {
	o.Service = &v
}

func (o ShipmentDomesticWithCarrierByCarrier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDomesticWithCarrierByCarrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	return toSerialize, nil
}

type NullableShipmentDomesticWithCarrierByCarrier struct {
	value *ShipmentDomesticWithCarrierByCarrier
	isSet bool
}

func (v NullableShipmentDomesticWithCarrierByCarrier) Get() *ShipmentDomesticWithCarrierByCarrier {
	return v.value
}

func (v *NullableShipmentDomesticWithCarrierByCarrier) Set(val *ShipmentDomesticWithCarrierByCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDomesticWithCarrierByCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDomesticWithCarrierByCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDomesticWithCarrierByCarrier(val *ShipmentDomesticWithCarrierByCarrier) *NullableShipmentDomesticWithCarrierByCarrier {
	return &NullableShipmentDomesticWithCarrierByCarrier{value: val, isSet: true}
}

func (v NullableShipmentDomesticWithCarrierByCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDomesticWithCarrierByCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


