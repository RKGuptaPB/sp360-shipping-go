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

// checks if the MultipieceRatesRequestMultiPieceParcelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceRatesRequestMultiPieceParcelsInner{}

// MultipieceRatesRequestMultiPieceParcelsInner description
type MultipieceRatesRequestMultiPieceParcelsInner struct {
	// description
	ParcelType *string `json:"parcelType,omitempty"`
	// description
	ParcelId *string `json:"parcelId,omitempty"`
	Parcel *MultipieceRatesRequestMultiPieceParcelsInnerParcel `json:"parcel,omitempty"`
}

// NewMultipieceRatesRequestMultiPieceParcelsInner instantiates a new MultipieceRatesRequestMultiPieceParcelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceRatesRequestMultiPieceParcelsInner() *MultipieceRatesRequestMultiPieceParcelsInner {
	this := MultipieceRatesRequestMultiPieceParcelsInner{}
	return &this
}

// NewMultipieceRatesRequestMultiPieceParcelsInnerWithDefaults instantiates a new MultipieceRatesRequestMultiPieceParcelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceRatesRequestMultiPieceParcelsInnerWithDefaults() *MultipieceRatesRequestMultiPieceParcelsInner {
	this := MultipieceRatesRequestMultiPieceParcelsInner{}
	return &this
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetParcelId returns the ParcelId field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcelId() string {
	if o == nil || IsNil(o.ParcelId) {
		var ret string
		return ret
	}
	return *o.ParcelId
}

// GetParcelIdOk returns a tuple with the ParcelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelId) {
		return nil, false
	}
	return o.ParcelId, true
}

// HasParcelId returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) HasParcelId() bool {
	if o != nil && !IsNil(o.ParcelId) {
		return true
	}

	return false
}

// SetParcelId gets a reference to the given string and assigns it to the ParcelId field.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) SetParcelId(v string) {
	o.ParcelId = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcel() MultipieceRatesRequestMultiPieceParcelsInnerParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret MultipieceRatesRequestMultiPieceParcelsInnerParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) GetParcelOk() (*MultipieceRatesRequestMultiPieceParcelsInnerParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given MultipieceRatesRequestMultiPieceParcelsInnerParcel and assigns it to the Parcel field.
func (o *MultipieceRatesRequestMultiPieceParcelsInner) SetParcel(v MultipieceRatesRequestMultiPieceParcelsInnerParcel) {
	o.Parcel = &v
}

func (o MultipieceRatesRequestMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceRatesRequestMultiPieceParcelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.ParcelId) {
		toSerialize["parcelId"] = o.ParcelId
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	return toSerialize, nil
}

type NullableMultipieceRatesRequestMultiPieceParcelsInner struct {
	value *MultipieceRatesRequestMultiPieceParcelsInner
	isSet bool
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInner) Get() *MultipieceRatesRequestMultiPieceParcelsInner {
	return v.value
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInner) Set(val *MultipieceRatesRequestMultiPieceParcelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceRatesRequestMultiPieceParcelsInner(val *MultipieceRatesRequestMultiPieceParcelsInner) *NullableMultipieceRatesRequestMultiPieceParcelsInner {
	return &NullableMultipieceRatesRequestMultiPieceParcelsInner{value: val, isSet: true}
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


