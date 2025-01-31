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

// checks if the ParcelTypesInnerDimensionRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelTypesInnerDimensionRulesInner{}

// ParcelTypesInnerDimensionRulesInner struct for ParcelTypesInnerDimensionRulesInner
type ParcelTypesInnerDimensionRulesInner struct {
	// This measures the parcel's maximum length and determine parcel’s girth.
	MaxLengthPlusGirth *float32 `json:"maxLengthPlusGirth,omitempty"`
	MaxParcelDimensions *ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions `json:"maxParcelDimensions,omitempty"`
	// This measures the parcel's minimum length and determine parcel’s girth.
	MinLengthPlusGirth *float32 `json:"minLengthPlusGirth,omitempty"`
	MinParcelDimensions *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions `json:"minParcelDimensions,omitempty"`
	Required *bool `json:"required,omitempty"`
	// UnitofMesurement is a standard for measuring the physical quantities of specified dimension parameters.
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
}

// NewParcelTypesInnerDimensionRulesInner instantiates a new ParcelTypesInnerDimensionRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelTypesInnerDimensionRulesInner() *ParcelTypesInnerDimensionRulesInner {
	this := ParcelTypesInnerDimensionRulesInner{}
	return &this
}

// NewParcelTypesInnerDimensionRulesInnerWithDefaults instantiates a new ParcelTypesInnerDimensionRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelTypesInnerDimensionRulesInnerWithDefaults() *ParcelTypesInnerDimensionRulesInner {
	this := ParcelTypesInnerDimensionRulesInner{}
	return &this
}

// GetMaxLengthPlusGirth returns the MaxLengthPlusGirth field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetMaxLengthPlusGirth() float32 {
	if o == nil || IsNil(o.MaxLengthPlusGirth) {
		var ret float32
		return ret
	}
	return *o.MaxLengthPlusGirth
}

// GetMaxLengthPlusGirthOk returns a tuple with the MaxLengthPlusGirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetMaxLengthPlusGirthOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxLengthPlusGirth) {
		return nil, false
	}
	return o.MaxLengthPlusGirth, true
}

// HasMaxLengthPlusGirth returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasMaxLengthPlusGirth() bool {
	if o != nil && !IsNil(o.MaxLengthPlusGirth) {
		return true
	}

	return false
}

// SetMaxLengthPlusGirth gets a reference to the given float32 and assigns it to the MaxLengthPlusGirth field.
func (o *ParcelTypesInnerDimensionRulesInner) SetMaxLengthPlusGirth(v float32) {
	o.MaxLengthPlusGirth = &v
}

// GetMaxParcelDimensions returns the MaxParcelDimensions field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetMaxParcelDimensions() ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions {
	if o == nil || IsNil(o.MaxParcelDimensions) {
		var ret ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions
		return ret
	}
	return *o.MaxParcelDimensions
}

// GetMaxParcelDimensionsOk returns a tuple with the MaxParcelDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetMaxParcelDimensionsOk() (*ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions, bool) {
	if o == nil || IsNil(o.MaxParcelDimensions) {
		return nil, false
	}
	return o.MaxParcelDimensions, true
}

// HasMaxParcelDimensions returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasMaxParcelDimensions() bool {
	if o != nil && !IsNil(o.MaxParcelDimensions) {
		return true
	}

	return false
}

// SetMaxParcelDimensions gets a reference to the given ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions and assigns it to the MaxParcelDimensions field.
func (o *ParcelTypesInnerDimensionRulesInner) SetMaxParcelDimensions(v ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions) {
	o.MaxParcelDimensions = &v
}

// GetMinLengthPlusGirth returns the MinLengthPlusGirth field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetMinLengthPlusGirth() float32 {
	if o == nil || IsNil(o.MinLengthPlusGirth) {
		var ret float32
		return ret
	}
	return *o.MinLengthPlusGirth
}

// GetMinLengthPlusGirthOk returns a tuple with the MinLengthPlusGirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetMinLengthPlusGirthOk() (*float32, bool) {
	if o == nil || IsNil(o.MinLengthPlusGirth) {
		return nil, false
	}
	return o.MinLengthPlusGirth, true
}

// HasMinLengthPlusGirth returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasMinLengthPlusGirth() bool {
	if o != nil && !IsNil(o.MinLengthPlusGirth) {
		return true
	}

	return false
}

// SetMinLengthPlusGirth gets a reference to the given float32 and assigns it to the MinLengthPlusGirth field.
func (o *ParcelTypesInnerDimensionRulesInner) SetMinLengthPlusGirth(v float32) {
	o.MinLengthPlusGirth = &v
}

// GetMinParcelDimensions returns the MinParcelDimensions field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetMinParcelDimensions() ParcelTypesInnerDimensionRulesInnerMinParcelDimensions {
	if o == nil || IsNil(o.MinParcelDimensions) {
		var ret ParcelTypesInnerDimensionRulesInnerMinParcelDimensions
		return ret
	}
	return *o.MinParcelDimensions
}

// GetMinParcelDimensionsOk returns a tuple with the MinParcelDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetMinParcelDimensionsOk() (*ParcelTypesInnerDimensionRulesInnerMinParcelDimensions, bool) {
	if o == nil || IsNil(o.MinParcelDimensions) {
		return nil, false
	}
	return o.MinParcelDimensions, true
}

// HasMinParcelDimensions returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasMinParcelDimensions() bool {
	if o != nil && !IsNil(o.MinParcelDimensions) {
		return true
	}

	return false
}

// SetMinParcelDimensions gets a reference to the given ParcelTypesInnerDimensionRulesInnerMinParcelDimensions and assigns it to the MinParcelDimensions field.
func (o *ParcelTypesInnerDimensionRulesInner) SetMinParcelDimensions(v ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) {
	o.MinParcelDimensions = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ParcelTypesInnerDimensionRulesInner) SetRequired(v bool) {
	o.Required = &v
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInner) GetUnitOfMeasurement() string {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInner) GetUnitOfMeasurementOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		return nil, false
	}
	return o.UnitOfMeasurement, true
}

// HasUnitOfMeasurement returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInner) HasUnitOfMeasurement() bool {
	if o != nil && !IsNil(o.UnitOfMeasurement) {
		return true
	}

	return false
}

// SetUnitOfMeasurement gets a reference to the given string and assigns it to the UnitOfMeasurement field.
func (o *ParcelTypesInnerDimensionRulesInner) SetUnitOfMeasurement(v string) {
	o.UnitOfMeasurement = &v
}

func (o ParcelTypesInnerDimensionRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelTypesInnerDimensionRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxLengthPlusGirth) {
		toSerialize["maxLengthPlusGirth"] = o.MaxLengthPlusGirth
	}
	if !IsNil(o.MaxParcelDimensions) {
		toSerialize["maxParcelDimensions"] = o.MaxParcelDimensions
	}
	if !IsNil(o.MinLengthPlusGirth) {
		toSerialize["minLengthPlusGirth"] = o.MinLengthPlusGirth
	}
	if !IsNil(o.MinParcelDimensions) {
		toSerialize["minParcelDimensions"] = o.MinParcelDimensions
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.UnitOfMeasurement) {
		toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	}
	return toSerialize, nil
}

type NullableParcelTypesInnerDimensionRulesInner struct {
	value *ParcelTypesInnerDimensionRulesInner
	isSet bool
}

func (v NullableParcelTypesInnerDimensionRulesInner) Get() *ParcelTypesInnerDimensionRulesInner {
	return v.value
}

func (v *NullableParcelTypesInnerDimensionRulesInner) Set(val *ParcelTypesInnerDimensionRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelTypesInnerDimensionRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelTypesInnerDimensionRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelTypesInnerDimensionRulesInner(val *ParcelTypesInnerDimensionRulesInner) *NullableParcelTypesInnerDimensionRulesInner {
	return &NullableParcelTypesInnerDimensionRulesInner{value: val, isSet: true}
}

func (v NullableParcelTypesInnerDimensionRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelTypesInnerDimensionRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


