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

// checks if the SchedulePickupUPSRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupUPSRequest{}

// SchedulePickupUPSRequest struct for SchedulePickupUPSRequest
type SchedulePickupUPSRequest struct {
	// description
	PackageLocation *string `json:"packageLocation,omitempty"`
	// description
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	PickupAddress *SchedulePickupDHLEXPRequestPickupAddress `json:"pickupAddress,omitempty"`
	// description
	PickupSummary []SchedulePickupUPSRequestPickupSummaryInner `json:"pickupSummary,omitempty"`
	// description
	Additionalnotes *string `json:"additionalnotes,omitempty"`
	// description
	PickupStartTime *string `json:"pickupStartTime,omitempty"`
	// description
	PickupTotalWeight *float32 `json:"pickupTotalWeight,omitempty"`
	// description
	PickupTotalWeightUnit *string `json:"pickupTotalWeightUnit,omitempty"`
	// description
	Reference *string `json:"reference,omitempty"`
	// description
	PickupOptions []SchedulePickupUPSRequestPickupOptionsInner `json:"pickupOptions,omitempty"`
}

// NewSchedulePickupUPSRequest instantiates a new SchedulePickupUPSRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupUPSRequest() *SchedulePickupUPSRequest {
	this := SchedulePickupUPSRequest{}
	return &this
}

// NewSchedulePickupUPSRequestWithDefaults instantiates a new SchedulePickupUPSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupUPSRequestWithDefaults() *SchedulePickupUPSRequest {
	this := SchedulePickupUPSRequest{}
	return &this
}

// GetPackageLocation returns the PackageLocation field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPackageLocation() string {
	if o == nil || IsNil(o.PackageLocation) {
		var ret string
		return ret
	}
	return *o.PackageLocation
}

// GetPackageLocationOk returns a tuple with the PackageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPackageLocationOk() (*string, bool) {
	if o == nil || IsNil(o.PackageLocation) {
		return nil, false
	}
	return o.PackageLocation, true
}

// HasPackageLocation returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPackageLocation() bool {
	if o != nil && !IsNil(o.PackageLocation) {
		return true
	}

	return false
}

// SetPackageLocation gets a reference to the given string and assigns it to the PackageLocation field.
func (o *SchedulePickupUPSRequest) SetPackageLocation(v string) {
	o.PackageLocation = &v
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *SchedulePickupUPSRequest) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetPickupAddress returns the PickupAddress field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress {
	if o == nil || IsNil(o.PickupAddress) {
		var ret SchedulePickupDHLEXPRequestPickupAddress
		return ret
	}
	return *o.PickupAddress
}

// GetPickupAddressOk returns a tuple with the PickupAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool) {
	if o == nil || IsNil(o.PickupAddress) {
		return nil, false
	}
	return o.PickupAddress, true
}

// HasPickupAddress returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupAddress() bool {
	if o != nil && !IsNil(o.PickupAddress) {
		return true
	}

	return false
}

// SetPickupAddress gets a reference to the given SchedulePickupDHLEXPRequestPickupAddress and assigns it to the PickupAddress field.
func (o *SchedulePickupUPSRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress) {
	o.PickupAddress = &v
}

// GetPickupSummary returns the PickupSummary field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupSummary() []SchedulePickupUPSRequestPickupSummaryInner {
	if o == nil || IsNil(o.PickupSummary) {
		var ret []SchedulePickupUPSRequestPickupSummaryInner
		return ret
	}
	return o.PickupSummary
}

// GetPickupSummaryOk returns a tuple with the PickupSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupSummaryOk() ([]SchedulePickupUPSRequestPickupSummaryInner, bool) {
	if o == nil || IsNil(o.PickupSummary) {
		return nil, false
	}
	return o.PickupSummary, true
}

// HasPickupSummary returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupSummary() bool {
	if o != nil && !IsNil(o.PickupSummary) {
		return true
	}

	return false
}

// SetPickupSummary gets a reference to the given []SchedulePickupUPSRequestPickupSummaryInner and assigns it to the PickupSummary field.
func (o *SchedulePickupUPSRequest) SetPickupSummary(v []SchedulePickupUPSRequestPickupSummaryInner) {
	o.PickupSummary = v
}

// GetAdditionalnotes returns the Additionalnotes field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetAdditionalnotes() string {
	if o == nil || IsNil(o.Additionalnotes) {
		var ret string
		return ret
	}
	return *o.Additionalnotes
}

// GetAdditionalnotesOk returns a tuple with the Additionalnotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetAdditionalnotesOk() (*string, bool) {
	if o == nil || IsNil(o.Additionalnotes) {
		return nil, false
	}
	return o.Additionalnotes, true
}

// HasAdditionalnotes returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasAdditionalnotes() bool {
	if o != nil && !IsNil(o.Additionalnotes) {
		return true
	}

	return false
}

// SetAdditionalnotes gets a reference to the given string and assigns it to the Additionalnotes field.
func (o *SchedulePickupUPSRequest) SetAdditionalnotes(v string) {
	o.Additionalnotes = &v
}

// GetPickupStartTime returns the PickupStartTime field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupStartTime() string {
	if o == nil || IsNil(o.PickupStartTime) {
		var ret string
		return ret
	}
	return *o.PickupStartTime
}

// GetPickupStartTimeOk returns a tuple with the PickupStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PickupStartTime) {
		return nil, false
	}
	return o.PickupStartTime, true
}

// HasPickupStartTime returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupStartTime() bool {
	if o != nil && !IsNil(o.PickupStartTime) {
		return true
	}

	return false
}

// SetPickupStartTime gets a reference to the given string and assigns it to the PickupStartTime field.
func (o *SchedulePickupUPSRequest) SetPickupStartTime(v string) {
	o.PickupStartTime = &v
}

// GetPickupTotalWeight returns the PickupTotalWeight field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupTotalWeight() float32 {
	if o == nil || IsNil(o.PickupTotalWeight) {
		var ret float32
		return ret
	}
	return *o.PickupTotalWeight
}

// GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupTotalWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.PickupTotalWeight) {
		return nil, false
	}
	return o.PickupTotalWeight, true
}

// HasPickupTotalWeight returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupTotalWeight() bool {
	if o != nil && !IsNil(o.PickupTotalWeight) {
		return true
	}

	return false
}

// SetPickupTotalWeight gets a reference to the given float32 and assigns it to the PickupTotalWeight field.
func (o *SchedulePickupUPSRequest) SetPickupTotalWeight(v float32) {
	o.PickupTotalWeight = &v
}

// GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupTotalWeightUnit() string {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		var ret string
		return ret
	}
	return *o.PickupTotalWeightUnit
}

// GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupTotalWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		return nil, false
	}
	return o.PickupTotalWeightUnit, true
}

// HasPickupTotalWeightUnit returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupTotalWeightUnit() bool {
	if o != nil && !IsNil(o.PickupTotalWeightUnit) {
		return true
	}

	return false
}

// SetPickupTotalWeightUnit gets a reference to the given string and assigns it to the PickupTotalWeightUnit field.
func (o *SchedulePickupUPSRequest) SetPickupTotalWeightUnit(v string) {
	o.PickupTotalWeightUnit = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *SchedulePickupUPSRequest) SetReference(v string) {
	o.Reference = &v
}

// GetPickupOptions returns the PickupOptions field value if set, zero value otherwise.
func (o *SchedulePickupUPSRequest) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner {
	if o == nil || IsNil(o.PickupOptions) {
		var ret []SchedulePickupUPSRequestPickupOptionsInner
		return ret
	}
	return o.PickupOptions
}

// GetPickupOptionsOk returns a tuple with the PickupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSRequest) GetPickupOptionsOk() ([]SchedulePickupUPSRequestPickupOptionsInner, bool) {
	if o == nil || IsNil(o.PickupOptions) {
		return nil, false
	}
	return o.PickupOptions, true
}

// HasPickupOptions returns a boolean if a field has been set.
func (o *SchedulePickupUPSRequest) HasPickupOptions() bool {
	if o != nil && !IsNil(o.PickupOptions) {
		return true
	}

	return false
}

// SetPickupOptions gets a reference to the given []SchedulePickupUPSRequestPickupOptionsInner and assigns it to the PickupOptions field.
func (o *SchedulePickupUPSRequest) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner) {
	o.PickupOptions = v
}

func (o SchedulePickupUPSRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupUPSRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageLocation) {
		toSerialize["packageLocation"] = o.PackageLocation
	}
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.PickupAddress) {
		toSerialize["pickupAddress"] = o.PickupAddress
	}
	if !IsNil(o.PickupSummary) {
		toSerialize["pickupSummary"] = o.PickupSummary
	}
	if !IsNil(o.Additionalnotes) {
		toSerialize["additionalnotes"] = o.Additionalnotes
	}
	if !IsNil(o.PickupStartTime) {
		toSerialize["pickupStartTime"] = o.PickupStartTime
	}
	if !IsNil(o.PickupTotalWeight) {
		toSerialize["pickupTotalWeight"] = o.PickupTotalWeight
	}
	if !IsNil(o.PickupTotalWeightUnit) {
		toSerialize["pickupTotalWeightUnit"] = o.PickupTotalWeightUnit
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.PickupOptions) {
		toSerialize["pickupOptions"] = o.PickupOptions
	}
	return toSerialize, nil
}

type NullableSchedulePickupUPSRequest struct {
	value *SchedulePickupUPSRequest
	isSet bool
}

func (v NullableSchedulePickupUPSRequest) Get() *SchedulePickupUPSRequest {
	return v.value
}

func (v *NullableSchedulePickupUPSRequest) Set(val *SchedulePickupUPSRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupUPSRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupUPSRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupUPSRequest(val *SchedulePickupUPSRequest) *NullableSchedulePickupUPSRequest {
	return &NullableSchedulePickupUPSRequest{value: val, isSet: true}
}

func (v NullableSchedulePickupUPSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupUPSRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


