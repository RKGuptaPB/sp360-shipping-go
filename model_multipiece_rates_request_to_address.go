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

// checks if the MultipieceRatesRequestToAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceRatesRequestToAddress{}

// MultipieceRatesRequestToAddress description
type MultipieceRatesRequestToAddress struct {
	// description
	Company *string `json:"company,omitempty"`
	// description
	Name *string `json:"name,omitempty"`
	// description
	Phone *string `json:"phone,omitempty"`
	// description
	Email *string `json:"email,omitempty"`
	// description
	Residential *bool `json:"residential,omitempty"`
	// description
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// description
	CityTown *string `json:"cityTown,omitempty"`
	// description
	StateProvince *string `json:"stateProvince,omitempty"`
	// description
	PostalCode *string `json:"postalCode,omitempty"`
	// description
	CountryCode *string `json:"countryCode,omitempty"`
}

// NewMultipieceRatesRequestToAddress instantiates a new MultipieceRatesRequestToAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceRatesRequestToAddress() *MultipieceRatesRequestToAddress {
	this := MultipieceRatesRequestToAddress{}
	return &this
}

// NewMultipieceRatesRequestToAddressWithDefaults instantiates a new MultipieceRatesRequestToAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceRatesRequestToAddressWithDefaults() *MultipieceRatesRequestToAddress {
	this := MultipieceRatesRequestToAddress{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *MultipieceRatesRequestToAddress) SetCompany(v string) {
	o.Company = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MultipieceRatesRequestToAddress) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *MultipieceRatesRequestToAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MultipieceRatesRequestToAddress) SetEmail(v string) {
	o.Email = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *MultipieceRatesRequestToAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *MultipieceRatesRequestToAddress) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *MultipieceRatesRequestToAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *MultipieceRatesRequestToAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *MultipieceRatesRequestToAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *MultipieceRatesRequestToAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestToAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *MultipieceRatesRequestToAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *MultipieceRatesRequestToAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o MultipieceRatesRequestToAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceRatesRequestToAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullableMultipieceRatesRequestToAddress struct {
	value *MultipieceRatesRequestToAddress
	isSet bool
}

func (v NullableMultipieceRatesRequestToAddress) Get() *MultipieceRatesRequestToAddress {
	return v.value
}

func (v *NullableMultipieceRatesRequestToAddress) Set(val *MultipieceRatesRequestToAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceRatesRequestToAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceRatesRequestToAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceRatesRequestToAddress(val *MultipieceRatesRequestToAddress) *NullableMultipieceRatesRequestToAddress {
	return &NullableMultipieceRatesRequestToAddress{value: val, isSet: true}
}

func (v NullableMultipieceRatesRequestToAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceRatesRequestToAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


