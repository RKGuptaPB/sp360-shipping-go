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

// checks if the JobStatusPrintStatusTransactionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobStatusPrintStatusTransactionInner{}

// JobStatusPrintStatusTransactionInner Array containing print statuses.
type JobStatusPrintStatusTransactionInner struct {
	// The status transaction ID of the Print document.
	Name *string `json:"name,omitempty"`
	// Status of the Print document.
	Status *string `json:"status,omitempty"`
}

// NewJobStatusPrintStatusTransactionInner instantiates a new JobStatusPrintStatusTransactionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStatusPrintStatusTransactionInner() *JobStatusPrintStatusTransactionInner {
	this := JobStatusPrintStatusTransactionInner{}
	return &this
}

// NewJobStatusPrintStatusTransactionInnerWithDefaults instantiates a new JobStatusPrintStatusTransactionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStatusPrintStatusTransactionInnerWithDefaults() *JobStatusPrintStatusTransactionInner {
	this := JobStatusPrintStatusTransactionInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JobStatusPrintStatusTransactionInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatusPrintStatusTransactionInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JobStatusPrintStatusTransactionInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JobStatusPrintStatusTransactionInner) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JobStatusPrintStatusTransactionInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatusPrintStatusTransactionInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JobStatusPrintStatusTransactionInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JobStatusPrintStatusTransactionInner) SetStatus(v string) {
	o.Status = &v
}

func (o JobStatusPrintStatusTransactionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobStatusPrintStatusTransactionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableJobStatusPrintStatusTransactionInner struct {
	value *JobStatusPrintStatusTransactionInner
	isSet bool
}

func (v NullableJobStatusPrintStatusTransactionInner) Get() *JobStatusPrintStatusTransactionInner {
	return v.value
}

func (v *NullableJobStatusPrintStatusTransactionInner) Set(val *JobStatusPrintStatusTransactionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStatusPrintStatusTransactionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStatusPrintStatusTransactionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStatusPrintStatusTransactionInner(val *JobStatusPrintStatusTransactionInner) *NullableJobStatusPrintStatusTransactionInner {
	return &NullableJobStatusPrintStatusTransactionInner{value: val, isSet: true}
}

func (v NullableJobStatusPrintStatusTransactionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStatusPrintStatusTransactionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


