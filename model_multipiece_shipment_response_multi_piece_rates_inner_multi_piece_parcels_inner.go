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

// checks if the MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner{}

// MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner description
type MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner struct {
	// description
	ParcelType *string `json:"parcelType,omitempty"`
	Parcel *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel `json:"parcel,omitempty"`
	ParcelRate *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate `json:"parcelRate,omitempty"`
	// description
	ParcelTrackingNumber *string `json:"parcelTrackingNumber,omitempty"`
}

// NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner instantiates a new MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner {
	this := MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner{}
	return &this
}

// NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerWithDefaults instantiates a new MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerWithDefaults() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner {
	this := MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner{}
	return &this
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcel() MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelOk() (*MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel and assigns it to the Parcel field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) SetParcel(v MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcel) {
	o.Parcel = &v
}

// GetParcelRate returns the ParcelRate field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelRate() MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate {
	if o == nil || IsNil(o.ParcelRate) {
		var ret MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate
		return ret
	}
	return *o.ParcelRate
}

// GetParcelRateOk returns a tuple with the ParcelRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelRateOk() (*MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate, bool) {
	if o == nil || IsNil(o.ParcelRate) {
		return nil, false
	}
	return o.ParcelRate, true
}

// HasParcelRate returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) HasParcelRate() bool {
	if o != nil && !IsNil(o.ParcelRate) {
		return true
	}

	return false
}

// SetParcelRate gets a reference to the given MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate and assigns it to the ParcelRate field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) SetParcelRate(v MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRate) {
	o.ParcelRate = &v
}

// GetParcelTrackingNumber returns the ParcelTrackingNumber field value if set, zero value otherwise.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelTrackingNumber() string {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ParcelTrackingNumber
}

// GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) GetParcelTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		return nil, false
	}
	return o.ParcelTrackingNumber, true
}

// HasParcelTrackingNumber returns a boolean if a field has been set.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) HasParcelTrackingNumber() bool {
	if o != nil && !IsNil(o.ParcelTrackingNumber) {
		return true
	}

	return false
}

// SetParcelTrackingNumber gets a reference to the given string and assigns it to the ParcelTrackingNumber field.
func (o *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) SetParcelTrackingNumber(v string) {
	o.ParcelTrackingNumber = &v
}

func (o MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.ParcelRate) {
		toSerialize["parcelRate"] = o.ParcelRate
	}
	if !IsNil(o.ParcelTrackingNumber) {
		toSerialize["parcelTrackingNumber"] = o.ParcelTrackingNumber
	}
	return toSerialize, nil
}

type NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner struct {
	value *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner
	isSet bool
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) Get() *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner {
	return v.value
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) Set(val *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner(val *MultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner {
	return &NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner{value: val, isSet: true}
}

func (v NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


