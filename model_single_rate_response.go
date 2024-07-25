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

// checks if the SingleRateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleRateResponse{}

// SingleRateResponse struct for SingleRateResponse
type SingleRateResponse struct {
	FromAddress *SingleRateResponseFromAddress `json:"fromAddress,omitempty"`
	Parcel *SingleRateResponseParcel `json:"parcel,omitempty"`
	// It displays the rate for the required carrier-service and/or special service combination
	Rates []SingleRateResponseRatesInner `json:"rates,omitempty"`
	ToAddress *SingleRateResponseToAddress `json:"toAddress,omitempty"`
}

// NewSingleRateResponse instantiates a new SingleRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleRateResponse() *SingleRateResponse {
	this := SingleRateResponse{}
	return &this
}

// NewSingleRateResponseWithDefaults instantiates a new SingleRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleRateResponseWithDefaults() *SingleRateResponse {
	this := SingleRateResponse{}
	return &this
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *SingleRateResponse) GetFromAddress() SingleRateResponseFromAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret SingleRateResponseFromAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponse) GetFromAddressOk() (*SingleRateResponseFromAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *SingleRateResponse) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given SingleRateResponseFromAddress and assigns it to the FromAddress field.
func (o *SingleRateResponse) SetFromAddress(v SingleRateResponseFromAddress) {
	o.FromAddress = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *SingleRateResponse) GetParcel() SingleRateResponseParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret SingleRateResponseParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponse) GetParcelOk() (*SingleRateResponseParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *SingleRateResponse) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given SingleRateResponseParcel and assigns it to the Parcel field.
func (o *SingleRateResponse) SetParcel(v SingleRateResponseParcel) {
	o.Parcel = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *SingleRateResponse) GetRates() []SingleRateResponseRatesInner {
	if o == nil || IsNil(o.Rates) {
		var ret []SingleRateResponseRatesInner
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponse) GetRatesOk() ([]SingleRateResponseRatesInner, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *SingleRateResponse) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given []SingleRateResponseRatesInner and assigns it to the Rates field.
func (o *SingleRateResponse) SetRates(v []SingleRateResponseRatesInner) {
	o.Rates = v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *SingleRateResponse) GetToAddress() SingleRateResponseToAddress {
	if o == nil || IsNil(o.ToAddress) {
		var ret SingleRateResponseToAddress
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponse) GetToAddressOk() (*SingleRateResponseToAddress, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *SingleRateResponse) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given SingleRateResponseToAddress and assigns it to the ToAddress field.
func (o *SingleRateResponse) SetToAddress(v SingleRateResponseToAddress) {
	o.ToAddress = &v
}

func (o SingleRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}
	if !IsNil(o.ToAddress) {
		toSerialize["toAddress"] = o.ToAddress
	}
	return toSerialize, nil
}

type NullableSingleRateResponse struct {
	value *SingleRateResponse
	isSet bool
}

func (v NullableSingleRateResponse) Get() *SingleRateResponse {
	return v.value
}

func (v *NullableSingleRateResponse) Set(val *SingleRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleRateResponse(val *SingleRateResponse) *NullableSingleRateResponse {
	return &NullableSingleRateResponse{value: val, isSet: true}
}

func (v NullableSingleRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


