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

// checks if the MultipieceShipmentRequestToAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceShipmentRequestToAddress{}

// MultipieceShipmentRequestToAddress struct for MultipieceShipmentRequestToAddress
type MultipieceShipmentRequestToAddress struct {
	// The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential).
	AddressLine1 string `json:"addressLine1"`
	// The addressLine2 contains Street address or Landmark (if any).
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The addressLine3 contains P.O. Box (if any) near the address.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The name of the city or town to where the address belongs.
	CityTown *string `json:"cityTown,omitempty"`
	//  The country in which the address is located. Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines.
	CountryCode string `json:"countryCode"`
	// Name of the recipient to which this address points.
	Name *string `json:"name,omitempty"`
	// This is recipient's phone number. Enter the digits with or without spaces or hyphens. The maximum limit of characters for Phone number is 10 digits. 
	Phone *string `json:"phone,omitempty"`
	// The Postal Code or ZIP Code of the address. For US addresses, use either the 5-digit or 9-digit ZIP Code in one of the following formats: '12345' or '12345-6789'. If you use a different format, such as 12345- or 123451234, will receive an error.
	PostalCode string `json:"postalCode"`
	// The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code. 
	StateProvince string `json:"stateProvince"`
}

type _MultipieceShipmentRequestToAddress MultipieceShipmentRequestToAddress

// NewMultipieceShipmentRequestToAddress instantiates a new MultipieceShipmentRequestToAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceShipmentRequestToAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string) *MultipieceShipmentRequestToAddress {
	this := MultipieceShipmentRequestToAddress{}
	this.AddressLine1 = addressLine1
	this.CountryCode = countryCode
	this.PostalCode = postalCode
	this.StateProvince = stateProvince
	return &this
}

// NewMultipieceShipmentRequestToAddressWithDefaults instantiates a new MultipieceShipmentRequestToAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceShipmentRequestToAddressWithDefaults() *MultipieceShipmentRequestToAddress {
	this := MultipieceShipmentRequestToAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *MultipieceShipmentRequestToAddress) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *MultipieceShipmentRequestToAddress) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *MultipieceShipmentRequestToAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *MultipieceShipmentRequestToAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *MultipieceShipmentRequestToAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *MultipieceShipmentRequestToAddress) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *MultipieceShipmentRequestToAddress) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *MultipieceShipmentRequestToAddress) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *MultipieceShipmentRequestToAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *MultipieceShipmentRequestToAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *MultipieceShipmentRequestToAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCountryCode returns the CountryCode field value
func (o *MultipieceShipmentRequestToAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *MultipieceShipmentRequestToAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MultipieceShipmentRequestToAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MultipieceShipmentRequestToAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MultipieceShipmentRequestToAddress) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *MultipieceShipmentRequestToAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *MultipieceShipmentRequestToAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *MultipieceShipmentRequestToAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value
func (o *MultipieceShipmentRequestToAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *MultipieceShipmentRequestToAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetStateProvince returns the StateProvince field value
func (o *MultipieceShipmentRequestToAddress) GetStateProvince() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value
// and a boolean to check if the value has been set.
func (o *MultipieceShipmentRequestToAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateProvince, true
}

// SetStateProvince sets field value
func (o *MultipieceShipmentRequestToAddress) SetStateProvince(v string) {
	o.StateProvince = v
}

func (o MultipieceShipmentRequestToAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceShipmentRequestToAddress) ToMap() (map[string]interface{}, error) {
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

func (o *MultipieceShipmentRequestToAddress) UnmarshalJSON(data []byte) (err error) {
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

	varMultipieceShipmentRequestToAddress := _MultipieceShipmentRequestToAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultipieceShipmentRequestToAddress)

	if err != nil {
		return err
	}

	*o = MultipieceShipmentRequestToAddress(varMultipieceShipmentRequestToAddress)

	return err
}

type NullableMultipieceShipmentRequestToAddress struct {
	value *MultipieceShipmentRequestToAddress
	isSet bool
}

func (v NullableMultipieceShipmentRequestToAddress) Get() *MultipieceShipmentRequestToAddress {
	return v.value
}

func (v *NullableMultipieceShipmentRequestToAddress) Set(val *MultipieceShipmentRequestToAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceShipmentRequestToAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceShipmentRequestToAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceShipmentRequestToAddress(val *MultipieceShipmentRequestToAddress) *NullableMultipieceShipmentRequestToAddress {
	return &NullableMultipieceShipmentRequestToAddress{value: val, isSet: true}
}

func (v NullableMultipieceShipmentRequestToAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceShipmentRequestToAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


