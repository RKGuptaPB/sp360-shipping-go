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

// checks if the MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}

// MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner description
type MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner struct {
	// description
	Fee *float32 `json:"fee,omitempty"`
	// description
	Name *string `json:"name,omitempty"`
}

// NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner instantiates a new MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	this := MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}
	return &this
}

// NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInnerWithDefaults instantiates a new MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInnerWithDefaults() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	this := MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetFee() float32 {
	if o == nil || IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) SetFee(v float32) {
	o.Fee = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) SetName(v string) {
	o.Name = &v
}

func (o MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner struct {
	value *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner
	isSet bool
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Get() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	return v.value
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Set(val *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner(val *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	return &NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{value: val, isSet: true}
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


