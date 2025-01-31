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

// checks if the SchedulePickupFedexRequestPickupSummaryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupFedexRequestPickupSummaryInner{}

// SchedulePickupFedexRequestPickupSummaryInner struct for SchedulePickupFedexRequestPickupSummaryInner
type SchedulePickupFedexRequestPickupSummaryInner struct {
	// description
	ServiceId *string `json:"serviceId,omitempty"`
	// description
	PackageCount *float32 `json:"packageCount,omitempty"`
	// description
	TotalWeight *float32 `json:"totalWeight,omitempty"`
	// description
	WeightUnit *string `json:"weightUnit,omitempty"`
	// description
	ParcelType *string `json:"parcelType,omitempty"`
	// description
	ToAddressCountryCode *string `json:"toAddressCountryCode,omitempty"`
}

// NewSchedulePickupFedexRequestPickupSummaryInner instantiates a new SchedulePickupFedexRequestPickupSummaryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupFedexRequestPickupSummaryInner() *SchedulePickupFedexRequestPickupSummaryInner {
	this := SchedulePickupFedexRequestPickupSummaryInner{}
	return &this
}

// NewSchedulePickupFedexRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupFedexRequestPickupSummaryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupFedexRequestPickupSummaryInnerWithDefaults() *SchedulePickupFedexRequestPickupSummaryInner {
	this := SchedulePickupFedexRequestPickupSummaryInner{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetPackageCount returns the PackageCount field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetPackageCount() float32 {
	if o == nil || IsNil(o.PackageCount) {
		var ret float32
		return ret
	}
	return *o.PackageCount
}

// GetPackageCountOk returns a tuple with the PackageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool) {
	if o == nil || IsNil(o.PackageCount) {
		return nil, false
	}
	return o.PackageCount, true
}

// HasPackageCount returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasPackageCount() bool {
	if o != nil && !IsNil(o.PackageCount) {
		return true
	}

	return false
}

// SetPackageCount gets a reference to the given float32 and assigns it to the PackageCount field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetPackageCount(v float32) {
	o.PackageCount = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetTotalWeight() float32 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float32
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float32 and assigns it to the TotalWeight field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetTotalWeight(v float32) {
	o.TotalWeight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetToAddressCountryCode returns the ToAddressCountryCode field value if set, zero value otherwise.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetToAddressCountryCode() string {
	if o == nil || IsNil(o.ToAddressCountryCode) {
		var ret string
		return ret
	}
	return *o.ToAddressCountryCode
}

// GetToAddressCountryCodeOk returns a tuple with the ToAddressCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) GetToAddressCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ToAddressCountryCode) {
		return nil, false
	}
	return o.ToAddressCountryCode, true
}

// HasToAddressCountryCode returns a boolean if a field has been set.
func (o *SchedulePickupFedexRequestPickupSummaryInner) HasToAddressCountryCode() bool {
	if o != nil && !IsNil(o.ToAddressCountryCode) {
		return true
	}

	return false
}

// SetToAddressCountryCode gets a reference to the given string and assigns it to the ToAddressCountryCode field.
func (o *SchedulePickupFedexRequestPickupSummaryInner) SetToAddressCountryCode(v string) {
	o.ToAddressCountryCode = &v
}

func (o SchedulePickupFedexRequestPickupSummaryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupFedexRequestPickupSummaryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.PackageCount) {
		toSerialize["packageCount"] = o.PackageCount
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["totalWeight"] = o.TotalWeight
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.ToAddressCountryCode) {
		toSerialize["toAddressCountryCode"] = o.ToAddressCountryCode
	}
	return toSerialize, nil
}

type NullableSchedulePickupFedexRequestPickupSummaryInner struct {
	value *SchedulePickupFedexRequestPickupSummaryInner
	isSet bool
}

func (v NullableSchedulePickupFedexRequestPickupSummaryInner) Get() *SchedulePickupFedexRequestPickupSummaryInner {
	return v.value
}

func (v *NullableSchedulePickupFedexRequestPickupSummaryInner) Set(val *SchedulePickupFedexRequestPickupSummaryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupFedexRequestPickupSummaryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupFedexRequestPickupSummaryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupFedexRequestPickupSummaryInner(val *SchedulePickupFedexRequestPickupSummaryInner) *NullableSchedulePickupFedexRequestPickupSummaryInner {
	return &NullableSchedulePickupFedexRequestPickupSummaryInner{value: val, isSet: true}
}

func (v NullableSchedulePickupFedexRequestPickupSummaryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupFedexRequestPickupSummaryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


