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

// checks if the RateShopResponseRatesInnerDeliveryCommitment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateShopResponseRatesInnerDeliveryCommitment{}

// RateShopResponseRatesInnerDeliveryCommitment It specifies the delivery commitment provided by carrier for this shipment
type RateShopResponseRatesInnerDeliveryCommitment struct {
	// It specifies the additional details for the delivery commitment
	AdditionalDetails *string `json:"additionalDetails,omitempty"`
	// It specifies estimated delivery date time
	EstimatedDeliveryDateTime *string `json:"estimatedDeliveryDateTime,omitempty"`
	// It specifies if guarantee is provided.
	Guarantee *string `json:"guarantee,omitempty"`
	// It specifies the maximum estimated number of days
	MaxEstimatedNumberOfDays *string `json:"maxEstimatedNumberOfDays,omitempty"`
	// It specifies the minimum estimated number of days
	MinEstimatedNumberOfDays *string `json:"minEstimatedNumberOfDays,omitempty"`
}

// NewRateShopResponseRatesInnerDeliveryCommitment instantiates a new RateShopResponseRatesInnerDeliveryCommitment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateShopResponseRatesInnerDeliveryCommitment() *RateShopResponseRatesInnerDeliveryCommitment {
	this := RateShopResponseRatesInnerDeliveryCommitment{}
	return &this
}

// NewRateShopResponseRatesInnerDeliveryCommitmentWithDefaults instantiates a new RateShopResponseRatesInnerDeliveryCommitment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateShopResponseRatesInnerDeliveryCommitmentWithDefaults() *RateShopResponseRatesInnerDeliveryCommitment {
	this := RateShopResponseRatesInnerDeliveryCommitment{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetAdditionalDetails() string {
	if o == nil || IsNil(o.AdditionalDetails) {
		var ret string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetAdditionalDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalDetails) {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) HasAdditionalDetails() bool {
	if o != nil && !IsNil(o.AdditionalDetails) {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given string and assigns it to the AdditionalDetails field.
func (o *RateShopResponseRatesInnerDeliveryCommitment) SetAdditionalDetails(v string) {
	o.AdditionalDetails = &v
}

// GetEstimatedDeliveryDateTime returns the EstimatedDeliveryDateTime field value if set, zero value otherwise.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetEstimatedDeliveryDateTime() string {
	if o == nil || IsNil(o.EstimatedDeliveryDateTime) {
		var ret string
		return ret
	}
	return *o.EstimatedDeliveryDateTime
}

// GetEstimatedDeliveryDateTimeOk returns a tuple with the EstimatedDeliveryDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetEstimatedDeliveryDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDateTime) {
		return nil, false
	}
	return o.EstimatedDeliveryDateTime, true
}

// HasEstimatedDeliveryDateTime returns a boolean if a field has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) HasEstimatedDeliveryDateTime() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDateTime) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDateTime gets a reference to the given string and assigns it to the EstimatedDeliveryDateTime field.
func (o *RateShopResponseRatesInnerDeliveryCommitment) SetEstimatedDeliveryDateTime(v string) {
	o.EstimatedDeliveryDateTime = &v
}

// GetGuarantee returns the Guarantee field value if set, zero value otherwise.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetGuarantee() string {
	if o == nil || IsNil(o.Guarantee) {
		var ret string
		return ret
	}
	return *o.Guarantee
}

// GetGuaranteeOk returns a tuple with the Guarantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetGuaranteeOk() (*string, bool) {
	if o == nil || IsNil(o.Guarantee) {
		return nil, false
	}
	return o.Guarantee, true
}

// HasGuarantee returns a boolean if a field has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) HasGuarantee() bool {
	if o != nil && !IsNil(o.Guarantee) {
		return true
	}

	return false
}

// SetGuarantee gets a reference to the given string and assigns it to the Guarantee field.
func (o *RateShopResponseRatesInnerDeliveryCommitment) SetGuarantee(v string) {
	o.Guarantee = &v
}

// GetMaxEstimatedNumberOfDays returns the MaxEstimatedNumberOfDays field value if set, zero value otherwise.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetMaxEstimatedNumberOfDays() string {
	if o == nil || IsNil(o.MaxEstimatedNumberOfDays) {
		var ret string
		return ret
	}
	return *o.MaxEstimatedNumberOfDays
}

// GetMaxEstimatedNumberOfDaysOk returns a tuple with the MaxEstimatedNumberOfDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetMaxEstimatedNumberOfDaysOk() (*string, bool) {
	if o == nil || IsNil(o.MaxEstimatedNumberOfDays) {
		return nil, false
	}
	return o.MaxEstimatedNumberOfDays, true
}

// HasMaxEstimatedNumberOfDays returns a boolean if a field has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) HasMaxEstimatedNumberOfDays() bool {
	if o != nil && !IsNil(o.MaxEstimatedNumberOfDays) {
		return true
	}

	return false
}

// SetMaxEstimatedNumberOfDays gets a reference to the given string and assigns it to the MaxEstimatedNumberOfDays field.
func (o *RateShopResponseRatesInnerDeliveryCommitment) SetMaxEstimatedNumberOfDays(v string) {
	o.MaxEstimatedNumberOfDays = &v
}

// GetMinEstimatedNumberOfDays returns the MinEstimatedNumberOfDays field value if set, zero value otherwise.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetMinEstimatedNumberOfDays() string {
	if o == nil || IsNil(o.MinEstimatedNumberOfDays) {
		var ret string
		return ret
	}
	return *o.MinEstimatedNumberOfDays
}

// GetMinEstimatedNumberOfDaysOk returns a tuple with the MinEstimatedNumberOfDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) GetMinEstimatedNumberOfDaysOk() (*string, bool) {
	if o == nil || IsNil(o.MinEstimatedNumberOfDays) {
		return nil, false
	}
	return o.MinEstimatedNumberOfDays, true
}

// HasMinEstimatedNumberOfDays returns a boolean if a field has been set.
func (o *RateShopResponseRatesInnerDeliveryCommitment) HasMinEstimatedNumberOfDays() bool {
	if o != nil && !IsNil(o.MinEstimatedNumberOfDays) {
		return true
	}

	return false
}

// SetMinEstimatedNumberOfDays gets a reference to the given string and assigns it to the MinEstimatedNumberOfDays field.
func (o *RateShopResponseRatesInnerDeliveryCommitment) SetMinEstimatedNumberOfDays(v string) {
	o.MinEstimatedNumberOfDays = &v
}

func (o RateShopResponseRatesInnerDeliveryCommitment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateShopResponseRatesInnerDeliveryCommitment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalDetails) {
		toSerialize["additionalDetails"] = o.AdditionalDetails
	}
	if !IsNil(o.EstimatedDeliveryDateTime) {
		toSerialize["estimatedDeliveryDateTime"] = o.EstimatedDeliveryDateTime
	}
	if !IsNil(o.Guarantee) {
		toSerialize["guarantee"] = o.Guarantee
	}
	if !IsNil(o.MaxEstimatedNumberOfDays) {
		toSerialize["maxEstimatedNumberOfDays"] = o.MaxEstimatedNumberOfDays
	}
	if !IsNil(o.MinEstimatedNumberOfDays) {
		toSerialize["minEstimatedNumberOfDays"] = o.MinEstimatedNumberOfDays
	}
	return toSerialize, nil
}

type NullableRateShopResponseRatesInnerDeliveryCommitment struct {
	value *RateShopResponseRatesInnerDeliveryCommitment
	isSet bool
}

func (v NullableRateShopResponseRatesInnerDeliveryCommitment) Get() *RateShopResponseRatesInnerDeliveryCommitment {
	return v.value
}

func (v *NullableRateShopResponseRatesInnerDeliveryCommitment) Set(val *RateShopResponseRatesInnerDeliveryCommitment) {
	v.value = val
	v.isSet = true
}

func (v NullableRateShopResponseRatesInnerDeliveryCommitment) IsSet() bool {
	return v.isSet
}

func (v *NullableRateShopResponseRatesInnerDeliveryCommitment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateShopResponseRatesInnerDeliveryCommitment(val *RateShopResponseRatesInnerDeliveryCommitment) *NullableRateShopResponseRatesInnerDeliveryCommitment {
	return &NullableRateShopResponseRatesInnerDeliveryCommitment{value: val, isSet: true}
}

func (v NullableRateShopResponseRatesInnerDeliveryCommitment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateShopResponseRatesInnerDeliveryCommitment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


