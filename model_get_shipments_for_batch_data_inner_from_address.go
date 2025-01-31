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

// checks if the GetShipmentsForBatchDataInnerFromAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentsForBatchDataInnerFromAddress{}

// GetShipmentsForBatchDataInnerFromAddress Sender address details
type GetShipmentsForBatchDataInnerFromAddress struct {
	AddressLines []string `json:"addressLines,omitempty"`
	CityTown *string `json:"cityTown,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Name *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Residential *bool `json:"residential,omitempty"`
	StateProvince *string `json:"stateProvince,omitempty"`
}

// NewGetShipmentsForBatchDataInnerFromAddress instantiates a new GetShipmentsForBatchDataInnerFromAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentsForBatchDataInnerFromAddress() *GetShipmentsForBatchDataInnerFromAddress {
	this := GetShipmentsForBatchDataInnerFromAddress{}
	return &this
}

// NewGetShipmentsForBatchDataInnerFromAddressWithDefaults instantiates a new GetShipmentsForBatchDataInnerFromAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentsForBatchDataInnerFromAddressWithDefaults() *GetShipmentsForBatchDataInnerFromAddress {
	this := GetShipmentsForBatchDataInnerFromAddress{}
	return &this
}

// GetAddressLines returns the AddressLines field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetAddressLines() []string {
	if o == nil || IsNil(o.AddressLines) {
		var ret []string
		return ret
	}
	return o.AddressLines
}

// GetAddressLinesOk returns a tuple with the AddressLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetAddressLinesOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressLines) {
		return nil, false
	}
	return o.AddressLines, true
}

// HasAddressLines returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasAddressLines() bool {
	if o != nil && !IsNil(o.AddressLines) {
		return true
	}

	return false
}

// SetAddressLines gets a reference to the given []string and assigns it to the AddressLines field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetAddressLines(v []string) {
	o.AddressLines = v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *GetShipmentsForBatchDataInnerFromAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *GetShipmentsForBatchDataInnerFromAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

func (o GetShipmentsForBatchDataInnerFromAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShipmentsForBatchDataInnerFromAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLines) {
		toSerialize["addressLines"] = o.AddressLines
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	return toSerialize, nil
}

type NullableGetShipmentsForBatchDataInnerFromAddress struct {
	value *GetShipmentsForBatchDataInnerFromAddress
	isSet bool
}

func (v NullableGetShipmentsForBatchDataInnerFromAddress) Get() *GetShipmentsForBatchDataInnerFromAddress {
	return v.value
}

func (v *NullableGetShipmentsForBatchDataInnerFromAddress) Set(val *GetShipmentsForBatchDataInnerFromAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentsForBatchDataInnerFromAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentsForBatchDataInnerFromAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentsForBatchDataInnerFromAddress(val *GetShipmentsForBatchDataInnerFromAddress) *NullableGetShipmentsForBatchDataInnerFromAddress {
	return &NullableGetShipmentsForBatchDataInnerFromAddress{value: val, isSet: true}
}

func (v NullableGetShipmentsForBatchDataInnerFromAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShipmentsForBatchDataInnerFromAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


