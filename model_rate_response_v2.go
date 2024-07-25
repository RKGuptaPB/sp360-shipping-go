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

// checks if the RateResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateResponseV2{}

// RateResponseV2 struct for RateResponseV2
type RateResponseV2 struct {
	// The base service charge is payable to the carrier, excluding special service charges.
	BaseCharge *float32 `json:"baseCharge,omitempty"`
	// Carrier is a service used to transport the parcels or couriers from one place to another.
	Carrier *string `json:"carrier,omitempty"`
	// A three-character (all uppercase letter) symbol of a currency according to the international ISO standard.<br /> As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof. For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR. 
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection.<br /> ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc.
	ParcelType *string `json:"parcelType,omitempty"`
	// The unique identifier given to the carrier specific service.
	ServiceId *string `json:"serviceId,omitempty"`
	// Additional fees or surcharges for the shipment. Each object in the array defines a surcharge and fee.
	Surcharges []RateResponseV2SurchargesInner `json:"surcharges,omitempty"`
	// The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties, except as noted below:
	TotalCarrierCharge *float32 `json:"totalCarrierCharge,omitempty"`
}

// NewRateResponseV2 instantiates a new RateResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateResponseV2() *RateResponseV2 {
	this := RateResponseV2{}
	return &this
}

// NewRateResponseV2WithDefaults instantiates a new RateResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateResponseV2WithDefaults() *RateResponseV2 {
	this := RateResponseV2{}
	return &this
}

// GetBaseCharge returns the BaseCharge field value if set, zero value otherwise.
func (o *RateResponseV2) GetBaseCharge() float32 {
	if o == nil || IsNil(o.BaseCharge) {
		var ret float32
		return ret
	}
	return *o.BaseCharge
}

// GetBaseChargeOk returns a tuple with the BaseCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetBaseChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCharge) {
		return nil, false
	}
	return o.BaseCharge, true
}

// HasBaseCharge returns a boolean if a field has been set.
func (o *RateResponseV2) HasBaseCharge() bool {
	if o != nil && !IsNil(o.BaseCharge) {
		return true
	}

	return false
}

// SetBaseCharge gets a reference to the given float32 and assigns it to the BaseCharge field.
func (o *RateResponseV2) SetBaseCharge(v float32) {
	o.BaseCharge = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *RateResponseV2) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *RateResponseV2) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *RateResponseV2) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *RateResponseV2) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *RateResponseV2) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *RateResponseV2) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *RateResponseV2) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *RateResponseV2) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *RateResponseV2) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *RateResponseV2) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *RateResponseV2) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *RateResponseV2) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSurcharges returns the Surcharges field value if set, zero value otherwise.
func (o *RateResponseV2) GetSurcharges() []RateResponseV2SurchargesInner {
	if o == nil || IsNil(o.Surcharges) {
		var ret []RateResponseV2SurchargesInner
		return ret
	}
	return o.Surcharges
}

// GetSurchargesOk returns a tuple with the Surcharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetSurchargesOk() ([]RateResponseV2SurchargesInner, bool) {
	if o == nil || IsNil(o.Surcharges) {
		return nil, false
	}
	return o.Surcharges, true
}

// HasSurcharges returns a boolean if a field has been set.
func (o *RateResponseV2) HasSurcharges() bool {
	if o != nil && !IsNil(o.Surcharges) {
		return true
	}

	return false
}

// SetSurcharges gets a reference to the given []RateResponseV2SurchargesInner and assigns it to the Surcharges field.
func (o *RateResponseV2) SetSurcharges(v []RateResponseV2SurchargesInner) {
	o.Surcharges = v
}

// GetTotalCarrierCharge returns the TotalCarrierCharge field value if set, zero value otherwise.
func (o *RateResponseV2) GetTotalCarrierCharge() float32 {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		var ret float32
		return ret
	}
	return *o.TotalCarrierCharge
}

// GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2) GetTotalCarrierChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		return nil, false
	}
	return o.TotalCarrierCharge, true
}

// HasTotalCarrierCharge returns a boolean if a field has been set.
func (o *RateResponseV2) HasTotalCarrierCharge() bool {
	if o != nil && !IsNil(o.TotalCarrierCharge) {
		return true
	}

	return false
}

// SetTotalCarrierCharge gets a reference to the given float32 and assigns it to the TotalCarrierCharge field.
func (o *RateResponseV2) SetTotalCarrierCharge(v float32) {
	o.TotalCarrierCharge = &v
}

func (o RateResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateResponseV2) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.Surcharges) {
		toSerialize["surcharges"] = o.Surcharges
	}
	if !IsNil(o.TotalCarrierCharge) {
		toSerialize["totalCarrierCharge"] = o.TotalCarrierCharge
	}
	return toSerialize, nil
}

type NullableRateResponseV2 struct {
	value *RateResponseV2
	isSet bool
}

func (v NullableRateResponseV2) Get() *RateResponseV2 {
	return v.value
}

func (v *NullableRateResponseV2) Set(val *RateResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableRateResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableRateResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateResponseV2(val *RateResponseV2) *NullableRateResponseV2 {
	return &NullableRateResponseV2{value: val, isSet: true}
}

func (v NullableRateResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


