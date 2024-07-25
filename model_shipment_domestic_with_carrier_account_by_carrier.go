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

// checks if the ShipmentDomesticWithCarrierAccountByCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDomesticWithCarrierAccountByCarrier{}

// ShipmentDomesticWithCarrierAccountByCarrier The shipment is grouped by Carrier and their Service.
type ShipmentDomesticWithCarrierAccountByCarrier struct {
	// This is a unique identifier associated with the specific sub-carrier account, which must be valid.<br /> This is used in the shipment creation (if this value is defined, Carrier properties will be skipped).
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	// Indicates a unique identifier associated with the carrier specific service, that is ServiceID, which must be valid.
	Service *string `json:"service,omitempty"`
}

// NewShipmentDomesticWithCarrierAccountByCarrier instantiates a new ShipmentDomesticWithCarrierAccountByCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDomesticWithCarrierAccountByCarrier() *ShipmentDomesticWithCarrierAccountByCarrier {
	this := ShipmentDomesticWithCarrierAccountByCarrier{}
	return &this
}

// NewShipmentDomesticWithCarrierAccountByCarrierWithDefaults instantiates a new ShipmentDomesticWithCarrierAccountByCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDomesticWithCarrierAccountByCarrierWithDefaults() *ShipmentDomesticWithCarrierAccountByCarrier {
	this := ShipmentDomesticWithCarrierAccountByCarrier{}
	return &this
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ShipmentDomesticWithCarrierAccountByCarrier) SetService(v string) {
	o.Service = &v
}

func (o ShipmentDomesticWithCarrierAccountByCarrier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDomesticWithCarrierAccountByCarrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	return toSerialize, nil
}

type NullableShipmentDomesticWithCarrierAccountByCarrier struct {
	value *ShipmentDomesticWithCarrierAccountByCarrier
	isSet bool
}

func (v NullableShipmentDomesticWithCarrierAccountByCarrier) Get() *ShipmentDomesticWithCarrierAccountByCarrier {
	return v.value
}

func (v *NullableShipmentDomesticWithCarrierAccountByCarrier) Set(val *ShipmentDomesticWithCarrierAccountByCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDomesticWithCarrierAccountByCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDomesticWithCarrierAccountByCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDomesticWithCarrierAccountByCarrier(val *ShipmentDomesticWithCarrierAccountByCarrier) *NullableShipmentDomesticWithCarrierAccountByCarrier {
	return &NullableShipmentDomesticWithCarrierAccountByCarrier{value: val, isSet: true}
}

func (v NullableShipmentDomesticWithCarrierAccountByCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDomesticWithCarrierAccountByCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


