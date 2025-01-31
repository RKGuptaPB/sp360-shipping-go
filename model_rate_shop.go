/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RateShop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateShop{}

// RateShop struct for RateShop
type RateShop struct {
	// Defines the date of the Shipment in the format YYYY:MM: DD., required for future shipment rating
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	FromAddress RateShopFromAddress `json:"fromAddress"`
	Parcel RateShopParcel `json:"parcel"`
	// It provides one or more carrier accounts Ids for rate shop. It can be referred from `Get Carrier Accounts` API
	CarrierAccounts []string `json:"carrierAccounts,omitempty"`
	// Parcel Type required for rating, it's value can be referred from response of `Get Parcel Types` API
	ParcelType *string `json:"parcelType,omitempty"`
	ToAddress SingleRateToAddress `json:"toAddress"`
	// isHazmat if set to true, only services which support Hazardous material shipment would be visible in the response
	IsHazmat *bool `json:"isHazmat,omitempty"`
}

type _RateShop RateShop

// NewRateShop instantiates a new RateShop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateShop(fromAddress RateShopFromAddress, parcel RateShopParcel, toAddress SingleRateToAddress) *RateShop {
	this := RateShop{}
	this.FromAddress = fromAddress
	this.Parcel = parcel
	this.ToAddress = toAddress
	return &this
}

// NewRateShopWithDefaults instantiates a new RateShop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateShopWithDefaults() *RateShop {
	this := RateShop{}
	return &this
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *RateShop) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShop) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *RateShop) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *RateShop) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetFromAddress returns the FromAddress field value
func (o *RateShop) GetFromAddress() RateShopFromAddress {
	if o == nil {
		var ret RateShopFromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *RateShop) GetFromAddressOk() (*RateShopFromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *RateShop) SetFromAddress(v RateShopFromAddress) {
	o.FromAddress = v
}

// GetParcel returns the Parcel field value
func (o *RateShop) GetParcel() RateShopParcel {
	if o == nil {
		var ret RateShopParcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *RateShop) GetParcelOk() (*RateShopParcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *RateShop) SetParcel(v RateShopParcel) {
	o.Parcel = v
}

// GetCarrierAccounts returns the CarrierAccounts field value if set, zero value otherwise.
func (o *RateShop) GetCarrierAccounts() []string {
	if o == nil || IsNil(o.CarrierAccounts) {
		var ret []string
		return ret
	}
	return o.CarrierAccounts
}

// GetCarrierAccountsOk returns a tuple with the CarrierAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShop) GetCarrierAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.CarrierAccounts) {
		return nil, false
	}
	return o.CarrierAccounts, true
}

// HasCarrierAccounts returns a boolean if a field has been set.
func (o *RateShop) HasCarrierAccounts() bool {
	if o != nil && !IsNil(o.CarrierAccounts) {
		return true
	}

	return false
}

// SetCarrierAccounts gets a reference to the given []string and assigns it to the CarrierAccounts field.
func (o *RateShop) SetCarrierAccounts(v []string) {
	o.CarrierAccounts = v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *RateShop) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShop) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *RateShop) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *RateShop) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetToAddress returns the ToAddress field value
func (o *RateShop) GetToAddress() SingleRateToAddress {
	if o == nil {
		var ret SingleRateToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *RateShop) GetToAddressOk() (*SingleRateToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *RateShop) SetToAddress(v SingleRateToAddress) {
	o.ToAddress = v
}

// GetIsHazmat returns the IsHazmat field value if set, zero value otherwise.
func (o *RateShop) GetIsHazmat() bool {
	if o == nil || IsNil(o.IsHazmat) {
		var ret bool
		return ret
	}
	return *o.IsHazmat
}

// GetIsHazmatOk returns a tuple with the IsHazmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShop) GetIsHazmatOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHazmat) {
		return nil, false
	}
	return o.IsHazmat, true
}

// HasIsHazmat returns a boolean if a field has been set.
func (o *RateShop) HasIsHazmat() bool {
	if o != nil && !IsNil(o.IsHazmat) {
		return true
	}

	return false
}

// SetIsHazmat gets a reference to the given bool and assigns it to the IsHazmat field.
func (o *RateShop) SetIsHazmat(v bool) {
	o.IsHazmat = &v
}

func (o RateShop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateShop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateOfShipment) {
		toSerialize["dateOfShipment"] = o.DateOfShipment
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["parcel"] = o.Parcel
	if !IsNil(o.CarrierAccounts) {
		toSerialize["carrierAccounts"] = o.CarrierAccounts
	}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	toSerialize["toAddress"] = o.ToAddress
	if !IsNil(o.IsHazmat) {
		toSerialize["isHazmat"] = o.IsHazmat
	}
	return toSerialize, nil
}

func (o *RateShop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromAddress",
		"parcel",
		"toAddress",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRateShop := _RateShop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRateShop)

	if err != nil {
		return err
	}

	*o = RateShop(varRateShop)

	return err
}

type NullableRateShop struct {
	value *RateShop
	isSet bool
}

func (v NullableRateShop) Get() *RateShop {
	return v.value
}

func (v *NullableRateShop) Set(val *RateShop) {
	v.value = val
	v.isSet = true
}

func (v NullableRateShop) IsSet() bool {
	return v.isSet
}

func (v *NullableRateShop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateShop(val *RateShop) *NullableRateShop {
	return &NullableRateShop{value: val, isSet: true}
}

func (v NullableRateShop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateShop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


