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

// checks if the SchedulePickupUPSResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupUPSResponse{}

// SchedulePickupUPSResponse struct for SchedulePickupUPSResponse
type SchedulePickupUPSResponse struct {
	// description
	PackageLocation *string `json:"packageLocation,omitempty"`
	// description
	PickupConfirmationNumber *string `json:"pickupConfirmationNumber,omitempty"`
	// description
	PickupId *string `json:"pickupId,omitempty"`
	// description
	Carrier *string `json:"carrier,omitempty"`
	// description
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	PickupAddress *SchedulePickupDHLEXPRequestPickupAddress `json:"pickupAddress,omitempty"`
	// description
	PickupSummary []SchedulePickupUPSResponsePickupSummaryInner `json:"pickupSummary,omitempty"`
	// description
	SpecialInstructions *string `json:"specialInstructions,omitempty"`
	// description
	PickupTotalWeight *float32 `json:"pickupTotalWeight,omitempty"`
	// description
	PickupTotalWeightUnit *string `json:"pickupTotalWeightUnit,omitempty"`
	// description
	PickupOptions []SchedulePickupUPSRequestPickupOptionsInner `json:"pickupOptions,omitempty"`
	PickupRate *SchedulePickupUPSResponsePickupRate `json:"pickupRate,omitempty"`
	// description
	TotalCarrierCharge *float32 `json:"totalCarrierCharge,omitempty"`
}

// NewSchedulePickupUPSResponse instantiates a new SchedulePickupUPSResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupUPSResponse() *SchedulePickupUPSResponse {
	this := SchedulePickupUPSResponse{}
	return &this
}

// NewSchedulePickupUPSResponseWithDefaults instantiates a new SchedulePickupUPSResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupUPSResponseWithDefaults() *SchedulePickupUPSResponse {
	this := SchedulePickupUPSResponse{}
	return &this
}

// GetPackageLocation returns the PackageLocation field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPackageLocation() string {
	if o == nil || IsNil(o.PackageLocation) {
		var ret string
		return ret
	}
	return *o.PackageLocation
}

// GetPackageLocationOk returns a tuple with the PackageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPackageLocationOk() (*string, bool) {
	if o == nil || IsNil(o.PackageLocation) {
		return nil, false
	}
	return o.PackageLocation, true
}

// HasPackageLocation returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPackageLocation() bool {
	if o != nil && !IsNil(o.PackageLocation) {
		return true
	}

	return false
}

// SetPackageLocation gets a reference to the given string and assigns it to the PackageLocation field.
func (o *SchedulePickupUPSResponse) SetPackageLocation(v string) {
	o.PackageLocation = &v
}

// GetPickupConfirmationNumber returns the PickupConfirmationNumber field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupConfirmationNumber() string {
	if o == nil || IsNil(o.PickupConfirmationNumber) {
		var ret string
		return ret
	}
	return *o.PickupConfirmationNumber
}

// GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupConfirmationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PickupConfirmationNumber) {
		return nil, false
	}
	return o.PickupConfirmationNumber, true
}

// HasPickupConfirmationNumber returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupConfirmationNumber() bool {
	if o != nil && !IsNil(o.PickupConfirmationNumber) {
		return true
	}

	return false
}

// SetPickupConfirmationNumber gets a reference to the given string and assigns it to the PickupConfirmationNumber field.
func (o *SchedulePickupUPSResponse) SetPickupConfirmationNumber(v string) {
	o.PickupConfirmationNumber = &v
}

// GetPickupId returns the PickupId field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupId() string {
	if o == nil || IsNil(o.PickupId) {
		var ret string
		return ret
	}
	return *o.PickupId
}

// GetPickupIdOk returns a tuple with the PickupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PickupId) {
		return nil, false
	}
	return o.PickupId, true
}

// HasPickupId returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupId() bool {
	if o != nil && !IsNil(o.PickupId) {
		return true
	}

	return false
}

// SetPickupId gets a reference to the given string and assigns it to the PickupId field.
func (o *SchedulePickupUPSResponse) SetPickupId(v string) {
	o.PickupId = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *SchedulePickupUPSResponse) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *SchedulePickupUPSResponse) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetPickupAddress returns the PickupAddress field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress {
	if o == nil || IsNil(o.PickupAddress) {
		var ret SchedulePickupDHLEXPRequestPickupAddress
		return ret
	}
	return *o.PickupAddress
}

// GetPickupAddressOk returns a tuple with the PickupAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool) {
	if o == nil || IsNil(o.PickupAddress) {
		return nil, false
	}
	return o.PickupAddress, true
}

// HasPickupAddress returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupAddress() bool {
	if o != nil && !IsNil(o.PickupAddress) {
		return true
	}

	return false
}

// SetPickupAddress gets a reference to the given SchedulePickupDHLEXPRequestPickupAddress and assigns it to the PickupAddress field.
func (o *SchedulePickupUPSResponse) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress) {
	o.PickupAddress = &v
}

// GetPickupSummary returns the PickupSummary field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupSummary() []SchedulePickupUPSResponsePickupSummaryInner {
	if o == nil || IsNil(o.PickupSummary) {
		var ret []SchedulePickupUPSResponsePickupSummaryInner
		return ret
	}
	return o.PickupSummary
}

// GetPickupSummaryOk returns a tuple with the PickupSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupSummaryOk() ([]SchedulePickupUPSResponsePickupSummaryInner, bool) {
	if o == nil || IsNil(o.PickupSummary) {
		return nil, false
	}
	return o.PickupSummary, true
}

// HasPickupSummary returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupSummary() bool {
	if o != nil && !IsNil(o.PickupSummary) {
		return true
	}

	return false
}

// SetPickupSummary gets a reference to the given []SchedulePickupUPSResponsePickupSummaryInner and assigns it to the PickupSummary field.
func (o *SchedulePickupUPSResponse) SetPickupSummary(v []SchedulePickupUPSResponsePickupSummaryInner) {
	o.PickupSummary = v
}

// GetSpecialInstructions returns the SpecialInstructions field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetSpecialInstructions() string {
	if o == nil || IsNil(o.SpecialInstructions) {
		var ret string
		return ret
	}
	return *o.SpecialInstructions
}

// GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetSpecialInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialInstructions) {
		return nil, false
	}
	return o.SpecialInstructions, true
}

// HasSpecialInstructions returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasSpecialInstructions() bool {
	if o != nil && !IsNil(o.SpecialInstructions) {
		return true
	}

	return false
}

// SetSpecialInstructions gets a reference to the given string and assigns it to the SpecialInstructions field.
func (o *SchedulePickupUPSResponse) SetSpecialInstructions(v string) {
	o.SpecialInstructions = &v
}

// GetPickupTotalWeight returns the PickupTotalWeight field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupTotalWeight() float32 {
	if o == nil || IsNil(o.PickupTotalWeight) {
		var ret float32
		return ret
	}
	return *o.PickupTotalWeight
}

// GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupTotalWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.PickupTotalWeight) {
		return nil, false
	}
	return o.PickupTotalWeight, true
}

// HasPickupTotalWeight returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupTotalWeight() bool {
	if o != nil && !IsNil(o.PickupTotalWeight) {
		return true
	}

	return false
}

// SetPickupTotalWeight gets a reference to the given float32 and assigns it to the PickupTotalWeight field.
func (o *SchedulePickupUPSResponse) SetPickupTotalWeight(v float32) {
	o.PickupTotalWeight = &v
}

// GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupTotalWeightUnit() string {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		var ret string
		return ret
	}
	return *o.PickupTotalWeightUnit
}

// GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupTotalWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		return nil, false
	}
	return o.PickupTotalWeightUnit, true
}

// HasPickupTotalWeightUnit returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupTotalWeightUnit() bool {
	if o != nil && !IsNil(o.PickupTotalWeightUnit) {
		return true
	}

	return false
}

// SetPickupTotalWeightUnit gets a reference to the given string and assigns it to the PickupTotalWeightUnit field.
func (o *SchedulePickupUPSResponse) SetPickupTotalWeightUnit(v string) {
	o.PickupTotalWeightUnit = &v
}

// GetPickupOptions returns the PickupOptions field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupOptions() []SchedulePickupUPSRequestPickupOptionsInner {
	if o == nil || IsNil(o.PickupOptions) {
		var ret []SchedulePickupUPSRequestPickupOptionsInner
		return ret
	}
	return o.PickupOptions
}

// GetPickupOptionsOk returns a tuple with the PickupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupOptionsOk() ([]SchedulePickupUPSRequestPickupOptionsInner, bool) {
	if o == nil || IsNil(o.PickupOptions) {
		return nil, false
	}
	return o.PickupOptions, true
}

// HasPickupOptions returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupOptions() bool {
	if o != nil && !IsNil(o.PickupOptions) {
		return true
	}

	return false
}

// SetPickupOptions gets a reference to the given []SchedulePickupUPSRequestPickupOptionsInner and assigns it to the PickupOptions field.
func (o *SchedulePickupUPSResponse) SetPickupOptions(v []SchedulePickupUPSRequestPickupOptionsInner) {
	o.PickupOptions = v
}

// GetPickupRate returns the PickupRate field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetPickupRate() SchedulePickupUPSResponsePickupRate {
	if o == nil || IsNil(o.PickupRate) {
		var ret SchedulePickupUPSResponsePickupRate
		return ret
	}
	return *o.PickupRate
}

// GetPickupRateOk returns a tuple with the PickupRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool) {
	if o == nil || IsNil(o.PickupRate) {
		return nil, false
	}
	return o.PickupRate, true
}

// HasPickupRate returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasPickupRate() bool {
	if o != nil && !IsNil(o.PickupRate) {
		return true
	}

	return false
}

// SetPickupRate gets a reference to the given SchedulePickupUPSResponsePickupRate and assigns it to the PickupRate field.
func (o *SchedulePickupUPSResponse) SetPickupRate(v SchedulePickupUPSResponsePickupRate) {
	o.PickupRate = &v
}

// GetTotalCarrierCharge returns the TotalCarrierCharge field value if set, zero value otherwise.
func (o *SchedulePickupUPSResponse) GetTotalCarrierCharge() float32 {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		var ret float32
		return ret
	}
	return *o.TotalCarrierCharge
}

// GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupUPSResponse) GetTotalCarrierChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		return nil, false
	}
	return o.TotalCarrierCharge, true
}

// HasTotalCarrierCharge returns a boolean if a field has been set.
func (o *SchedulePickupUPSResponse) HasTotalCarrierCharge() bool {
	if o != nil && !IsNil(o.TotalCarrierCharge) {
		return true
	}

	return false
}

// SetTotalCarrierCharge gets a reference to the given float32 and assigns it to the TotalCarrierCharge field.
func (o *SchedulePickupUPSResponse) SetTotalCarrierCharge(v float32) {
	o.TotalCarrierCharge = &v
}

func (o SchedulePickupUPSResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupUPSResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageLocation) {
		toSerialize["packageLocation"] = o.PackageLocation
	}
	if !IsNil(o.PickupConfirmationNumber) {
		toSerialize["pickupConfirmationNumber"] = o.PickupConfirmationNumber
	}
	if !IsNil(o.PickupId) {
		toSerialize["pickupId"] = o.PickupId
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
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
	if !IsNil(o.SpecialInstructions) {
		toSerialize["specialInstructions"] = o.SpecialInstructions
	}
	if !IsNil(o.PickupTotalWeight) {
		toSerialize["pickupTotalWeight"] = o.PickupTotalWeight
	}
	if !IsNil(o.PickupTotalWeightUnit) {
		toSerialize["pickupTotalWeightUnit"] = o.PickupTotalWeightUnit
	}
	if !IsNil(o.PickupOptions) {
		toSerialize["pickupOptions"] = o.PickupOptions
	}
	if !IsNil(o.PickupRate) {
		toSerialize["pickupRate"] = o.PickupRate
	}
	if !IsNil(o.TotalCarrierCharge) {
		toSerialize["totalCarrierCharge"] = o.TotalCarrierCharge
	}
	return toSerialize, nil
}

type NullableSchedulePickupUPSResponse struct {
	value *SchedulePickupUPSResponse
	isSet bool
}

func (v NullableSchedulePickupUPSResponse) Get() *SchedulePickupUPSResponse {
	return v.value
}

func (v *NullableSchedulePickupUPSResponse) Set(val *SchedulePickupUPSResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupUPSResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupUPSResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupUPSResponse(val *SchedulePickupUPSResponse) *NullableSchedulePickupUPSResponse {
	return &NullableSchedulePickupUPSResponse{value: val, isSet: true}
}

func (v NullableSchedulePickupUPSResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupUPSResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


