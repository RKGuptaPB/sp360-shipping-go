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

// checks if the DomesticShipmentResponseRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomesticShipmentResponseRate{}

// DomesticShipmentResponseRate struct for DomesticShipmentResponseRate
type DomesticShipmentResponseRate struct {
	// The base service charge is payable to the carrier, excluding special service charges.
	BaseCharge *float32 `json:"baseCharge,omitempty"`
	// Carrier is a service used to transport the parcels or couriers from one place to another.
	Carrier *string `json:"carrier,omitempty"`
	// A three-character (all uppercase letter) symbol of a currency according to the international ISO standard. As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof. For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR. 
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc. 
	ParcelType *string `json:"parcelType,omitempty"`
	// Its value can be CONTRACT_RATES, COMMERCIAL or COMMERCIAL_BASE for USPS and COMMERCIAL for other carriers depending on the Pitney Bowes contract/subscription
	RateTypeId *string `json:"rateTypeId,omitempty"`
	// The unique identifier given to the carrier specific service.
	ServiceId *string `json:"serviceId,omitempty"`
	// It provides a carrier-service based special or extra service.
	SpecialServices []DomesticShipmentResponseRateSpecialServicesInner `json:"specialServices,omitempty"`
	// The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties, except as noted below:
	TotalCarrierCharge *float32 `json:"totalCarrierCharge,omitempty"`
}

// NewDomesticShipmentResponseRate instantiates a new DomesticShipmentResponseRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomesticShipmentResponseRate() *DomesticShipmentResponseRate {
	this := DomesticShipmentResponseRate{}
	return &this
}

// NewDomesticShipmentResponseRateWithDefaults instantiates a new DomesticShipmentResponseRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomesticShipmentResponseRateWithDefaults() *DomesticShipmentResponseRate {
	this := DomesticShipmentResponseRate{}
	return &this
}

// GetBaseCharge returns the BaseCharge field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetBaseCharge() float32 {
	if o == nil || IsNil(o.BaseCharge) {
		var ret float32
		return ret
	}
	return *o.BaseCharge
}

// GetBaseChargeOk returns a tuple with the BaseCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetBaseChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCharge) {
		return nil, false
	}
	return o.BaseCharge, true
}

// HasBaseCharge returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasBaseCharge() bool {
	if o != nil && !IsNil(o.BaseCharge) {
		return true
	}

	return false
}

// SetBaseCharge gets a reference to the given float32 and assigns it to the BaseCharge field.
func (o *DomesticShipmentResponseRate) SetBaseCharge(v float32) {
	o.BaseCharge = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *DomesticShipmentResponseRate) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *DomesticShipmentResponseRate) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *DomesticShipmentResponseRate) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetRateTypeId returns the RateTypeId field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetRateTypeId() string {
	if o == nil || IsNil(o.RateTypeId) {
		var ret string
		return ret
	}
	return *o.RateTypeId
}

// GetRateTypeIdOk returns a tuple with the RateTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetRateTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RateTypeId) {
		return nil, false
	}
	return o.RateTypeId, true
}

// HasRateTypeId returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasRateTypeId() bool {
	if o != nil && !IsNil(o.RateTypeId) {
		return true
	}

	return false
}

// SetRateTypeId gets a reference to the given string and assigns it to the RateTypeId field.
func (o *DomesticShipmentResponseRate) SetRateTypeId(v string) {
	o.RateTypeId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DomesticShipmentResponseRate) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetSpecialServices() []DomesticShipmentResponseRateSpecialServicesInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []DomesticShipmentResponseRateSpecialServicesInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetSpecialServicesOk() ([]DomesticShipmentResponseRateSpecialServicesInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []DomesticShipmentResponseRateSpecialServicesInner and assigns it to the SpecialServices field.
func (o *DomesticShipmentResponseRate) SetSpecialServices(v []DomesticShipmentResponseRateSpecialServicesInner) {
	o.SpecialServices = v
}

// GetTotalCarrierCharge returns the TotalCarrierCharge field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRate) GetTotalCarrierCharge() float32 {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		var ret float32
		return ret
	}
	return *o.TotalCarrierCharge
}

// GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRate) GetTotalCarrierChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		return nil, false
	}
	return o.TotalCarrierCharge, true
}

// HasTotalCarrierCharge returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRate) HasTotalCarrierCharge() bool {
	if o != nil && !IsNil(o.TotalCarrierCharge) {
		return true
	}

	return false
}

// SetTotalCarrierCharge gets a reference to the given float32 and assigns it to the TotalCarrierCharge field.
func (o *DomesticShipmentResponseRate) SetTotalCarrierCharge(v float32) {
	o.TotalCarrierCharge = &v
}

func (o DomesticShipmentResponseRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomesticShipmentResponseRate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TotalCarrierCharge) {
		toSerialize["totalCarrierCharge"] = o.TotalCarrierCharge
	}
	return toSerialize, nil
}

type NullableDomesticShipmentResponseRate struct {
	value *DomesticShipmentResponseRate
	isSet bool
}

func (v NullableDomesticShipmentResponseRate) Get() *DomesticShipmentResponseRate {
	return v.value
}

func (v *NullableDomesticShipmentResponseRate) Set(val *DomesticShipmentResponseRate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomesticShipmentResponseRate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomesticShipmentResponseRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomesticShipmentResponseRate(val *DomesticShipmentResponseRate) *NullableDomesticShipmentResponseRate {
	return &NullableDomesticShipmentResponseRate{value: val, isSet: true}
}

func (v NullableDomesticShipmentResponseRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomesticShipmentResponseRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


