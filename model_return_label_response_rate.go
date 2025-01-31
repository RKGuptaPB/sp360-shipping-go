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

// checks if the ReturnLabelResponseRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLabelResponseRate{}

// ReturnLabelResponseRate struct for ReturnLabelResponseRate
type ReturnLabelResponseRate struct {
	// Description
	BaseCharge *float32 `json:"baseCharge,omitempty"`
	// Description
	Carrier *string `json:"carrier,omitempty"`
	// Description
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Description
	ParcelType *string `json:"parcelType,omitempty"`
	// Description
	RateTypeId *string `json:"rateTypeId,omitempty"`
	// Description
	ServiceId *string `json:"serviceId,omitempty"`
	SpecialServices []ReturnLabelResponseRateSpecialServicesInner `json:"specialServices,omitempty"`
	Surcharges []ReturnLabelResponseRateSurchargesInner `json:"surcharges,omitempty"`
	// Description
	TotalCarrierCharge *float32 `json:"totalCarrierCharge,omitempty"`
}

// NewReturnLabelResponseRate instantiates a new ReturnLabelResponseRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLabelResponseRate() *ReturnLabelResponseRate {
	this := ReturnLabelResponseRate{}
	return &this
}

// NewReturnLabelResponseRateWithDefaults instantiates a new ReturnLabelResponseRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLabelResponseRateWithDefaults() *ReturnLabelResponseRate {
	this := ReturnLabelResponseRate{}
	return &this
}

// GetBaseCharge returns the BaseCharge field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetBaseCharge() float32 {
	if o == nil || IsNil(o.BaseCharge) {
		var ret float32
		return ret
	}
	return *o.BaseCharge
}

// GetBaseChargeOk returns a tuple with the BaseCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetBaseChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCharge) {
		return nil, false
	}
	return o.BaseCharge, true
}

// HasBaseCharge returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasBaseCharge() bool {
	if o != nil && !IsNil(o.BaseCharge) {
		return true
	}

	return false
}

// SetBaseCharge gets a reference to the given float32 and assigns it to the BaseCharge field.
func (o *ReturnLabelResponseRate) SetBaseCharge(v float32) {
	o.BaseCharge = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *ReturnLabelResponseRate) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ReturnLabelResponseRate) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *ReturnLabelResponseRate) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetRateTypeId returns the RateTypeId field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetRateTypeId() string {
	if o == nil || IsNil(o.RateTypeId) {
		var ret string
		return ret
	}
	return *o.RateTypeId
}

// GetRateTypeIdOk returns a tuple with the RateTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetRateTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RateTypeId) {
		return nil, false
	}
	return o.RateTypeId, true
}

// HasRateTypeId returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasRateTypeId() bool {
	if o != nil && !IsNil(o.RateTypeId) {
		return true
	}

	return false
}

// SetRateTypeId gets a reference to the given string and assigns it to the RateTypeId field.
func (o *ReturnLabelResponseRate) SetRateTypeId(v string) {
	o.RateTypeId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ReturnLabelResponseRate) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetSpecialServices() []ReturnLabelResponseRateSpecialServicesInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []ReturnLabelResponseRateSpecialServicesInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetSpecialServicesOk() ([]ReturnLabelResponseRateSpecialServicesInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []ReturnLabelResponseRateSpecialServicesInner and assigns it to the SpecialServices field.
func (o *ReturnLabelResponseRate) SetSpecialServices(v []ReturnLabelResponseRateSpecialServicesInner) {
	o.SpecialServices = v
}

// GetSurcharges returns the Surcharges field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetSurcharges() []ReturnLabelResponseRateSurchargesInner {
	if o == nil || IsNil(o.Surcharges) {
		var ret []ReturnLabelResponseRateSurchargesInner
		return ret
	}
	return o.Surcharges
}

// GetSurchargesOk returns a tuple with the Surcharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetSurchargesOk() ([]ReturnLabelResponseRateSurchargesInner, bool) {
	if o == nil || IsNil(o.Surcharges) {
		return nil, false
	}
	return o.Surcharges, true
}

// HasSurcharges returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasSurcharges() bool {
	if o != nil && !IsNil(o.Surcharges) {
		return true
	}

	return false
}

// SetSurcharges gets a reference to the given []ReturnLabelResponseRateSurchargesInner and assigns it to the Surcharges field.
func (o *ReturnLabelResponseRate) SetSurcharges(v []ReturnLabelResponseRateSurchargesInner) {
	o.Surcharges = v
}

// GetTotalCarrierCharge returns the TotalCarrierCharge field value if set, zero value otherwise.
func (o *ReturnLabelResponseRate) GetTotalCarrierCharge() float32 {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		var ret float32
		return ret
	}
	return *o.TotalCarrierCharge
}

// GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseRate) GetTotalCarrierChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		return nil, false
	}
	return o.TotalCarrierCharge, true
}

// HasTotalCarrierCharge returns a boolean if a field has been set.
func (o *ReturnLabelResponseRate) HasTotalCarrierCharge() bool {
	if o != nil && !IsNil(o.TotalCarrierCharge) {
		return true
	}

	return false
}

// SetTotalCarrierCharge gets a reference to the given float32 and assigns it to the TotalCarrierCharge field.
func (o *ReturnLabelResponseRate) SetTotalCarrierCharge(v float32) {
	o.TotalCarrierCharge = &v
}

func (o ReturnLabelResponseRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLabelResponseRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseCharge) {
		toSerialize["baseCharge"] = o.BaseCharge
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.RateTypeId) {
		toSerialize["rateTypeId"] = o.RateTypeId
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	if !IsNil(o.Surcharges) {
		toSerialize["surcharges"] = o.Surcharges
	}
	if !IsNil(o.TotalCarrierCharge) {
		toSerialize["totalCarrierCharge"] = o.TotalCarrierCharge
	}
	return toSerialize, nil
}

type NullableReturnLabelResponseRate struct {
	value *ReturnLabelResponseRate
	isSet bool
}

func (v NullableReturnLabelResponseRate) Get() *ReturnLabelResponseRate {
	return v.value
}

func (v *NullableReturnLabelResponseRate) Set(val *ReturnLabelResponseRate) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLabelResponseRate) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLabelResponseRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLabelResponseRate(val *ReturnLabelResponseRate) *NullableReturnLabelResponseRate {
	return &NullableReturnLabelResponseRate{value: val, isSet: true}
}

func (v NullableReturnLabelResponseRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLabelResponseRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


