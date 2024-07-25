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

// checks if the ShipmentInternationalBatchFromAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInternationalBatchFromAddress{}

// ShipmentInternationalBatchFromAddress struct for ShipmentInternationalBatchFromAddress
type ShipmentInternationalBatchFromAddress struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	AddressLine3 *string `json:"addressLine3,omitempty"`
	CityTown *string `json:"cityTown,omitempty"`
	CountryCode string `json:"countryCode"`
	Name *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	PostalCode string `json:"postalCode"`
	StateProvince string `json:"stateProvince"`
}

type _ShipmentInternationalBatchFromAddress ShipmentInternationalBatchFromAddress

// NewShipmentInternationalBatchFromAddress instantiates a new ShipmentInternationalBatchFromAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInternationalBatchFromAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string) *ShipmentInternationalBatchFromAddress {
	this := ShipmentInternationalBatchFromAddress{}
	this.AddressLine1 = addressLine1
	this.CountryCode = countryCode
	this.PostalCode = postalCode
	this.StateProvince = stateProvince
	return &this
}

// NewShipmentInternationalBatchFromAddressWithDefaults instantiates a new ShipmentInternationalBatchFromAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInternationalBatchFromAddressWithDefaults() *ShipmentInternationalBatchFromAddress {
	this := ShipmentInternationalBatchFromAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *ShipmentInternationalBatchFromAddress) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *ShipmentInternationalBatchFromAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *ShipmentInternationalBatchFromAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *ShipmentInternationalBatchFromAddress) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *ShipmentInternationalBatchFromAddress) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *ShipmentInternationalBatchFromAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *ShipmentInternationalBatchFromAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *ShipmentInternationalBatchFromAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCountryCode returns the CountryCode field value
func (o *ShipmentInternationalBatchFromAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *ShipmentInternationalBatchFromAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipmentInternationalBatchFromAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipmentInternationalBatchFromAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipmentInternationalBatchFromAddress) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ShipmentInternationalBatchFromAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ShipmentInternationalBatchFromAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ShipmentInternationalBatchFromAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value
func (o *ShipmentInternationalBatchFromAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *ShipmentInternationalBatchFromAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetStateProvince returns the StateProvince field value
func (o *ShipmentInternationalBatchFromAddress) GetStateProvince() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalBatchFromAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateProvince, true
}

// SetStateProvince sets field value
func (o *ShipmentInternationalBatchFromAddress) SetStateProvince(v string) {
	o.StateProvince = v
}

func (o ShipmentInternationalBatchFromAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInternationalBatchFromAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	toSerialize["postalCode"] = o.PostalCode
	toSerialize["stateProvince"] = o.StateProvince
	return toSerialize, nil
}

func (o *ShipmentInternationalBatchFromAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressLine1",
		"countryCode",
		"postalCode",
		"stateProvince",
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

	varShipmentInternationalBatchFromAddress := _ShipmentInternationalBatchFromAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentInternationalBatchFromAddress)

	if err != nil {
		return err
	}

	*o = ShipmentInternationalBatchFromAddress(varShipmentInternationalBatchFromAddress)

	return err
}

type NullableShipmentInternationalBatchFromAddress struct {
	value *ShipmentInternationalBatchFromAddress
	isSet bool
}

func (v NullableShipmentInternationalBatchFromAddress) Get() *ShipmentInternationalBatchFromAddress {
	return v.value
}

func (v *NullableShipmentInternationalBatchFromAddress) Set(val *ShipmentInternationalBatchFromAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInternationalBatchFromAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInternationalBatchFromAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInternationalBatchFromAddress(val *ShipmentInternationalBatchFromAddress) *NullableShipmentInternationalBatchFromAddress {
	return &NullableShipmentInternationalBatchFromAddress{value: val, isSet: true}
}

func (v NullableShipmentInternationalBatchFromAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInternationalBatchFromAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


