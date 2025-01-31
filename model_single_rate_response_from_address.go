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

// checks if the SingleRateResponseFromAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleRateResponseFromAddress{}

// SingleRateResponseFromAddress Sender address details
type SingleRateResponseFromAddress struct {
	// The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential).
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The addressLine2 contains Street address or Landmark (if any).
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The addressLine3 contains P.O. Box (if any) near the address.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The name of the city or town to where the sender's address belongs.
	CityTown *string `json:"cityTown,omitempty"`
	// The name of the company, in case if the sender address is not residential.
	Company *string `json:"company,omitempty"`
	// The two-character ISO Code of the source country from this ISO country list. The country in which the address is located. Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines.
	CountryCode *string `json:"countryCode,omitempty"`
	// Name of the sender to which this address points.
	Name *string `json:"name,omitempty"`
	// The Postal Code or ZIP Code of the address. For US addresses, use either the 5-digit or 9-digit ZIP Code in one of the following formats: '12345' or '12345-6789'. If you use a different format, such as 12345- or 123451234, will receive an error.
	PostalCode *string `json:"postalCode,omitempty"`
	// The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code. 
	StateProvince *string `json:"stateProvince,omitempty"`
}

// NewSingleRateResponseFromAddress instantiates a new SingleRateResponseFromAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleRateResponseFromAddress() *SingleRateResponseFromAddress {
	this := SingleRateResponseFromAddress{}
	return &this
}

// NewSingleRateResponseFromAddressWithDefaults instantiates a new SingleRateResponseFromAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleRateResponseFromAddressWithDefaults() *SingleRateResponseFromAddress {
	this := SingleRateResponseFromAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *SingleRateResponseFromAddress) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *SingleRateResponseFromAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *SingleRateResponseFromAddress) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *SingleRateResponseFromAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *SingleRateResponseFromAddress) SetCompany(v string) {
	o.Company = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SingleRateResponseFromAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SingleRateResponseFromAddress) SetName(v string) {
	o.Name = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *SingleRateResponseFromAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *SingleRateResponseFromAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseFromAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *SingleRateResponseFromAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *SingleRateResponseFromAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

func (o SingleRateResponseFromAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleRateResponseFromAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	return toSerialize, nil
}

type NullableSingleRateResponseFromAddress struct {
	value *SingleRateResponseFromAddress
	isSet bool
}

func (v NullableSingleRateResponseFromAddress) Get() *SingleRateResponseFromAddress {
	return v.value
}

func (v *NullableSingleRateResponseFromAddress) Set(val *SingleRateResponseFromAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleRateResponseFromAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleRateResponseFromAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleRateResponseFromAddress(val *SingleRateResponseFromAddress) *NullableSingleRateResponseFromAddress {
	return &NullableSingleRateResponseFromAddress{value: val, isSet: true}
}

func (v NullableSingleRateResponseFromAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleRateResponseFromAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


