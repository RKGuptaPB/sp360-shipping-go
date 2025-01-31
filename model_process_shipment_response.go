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

// checks if the ProcessShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessShipmentResponse{}

// ProcessShipmentResponse struct for ProcessShipmentResponse
type ProcessShipmentResponse struct {
	//  This is the system generated Batch ID.
	BatchId *string `json:"batchId,omitempty"`
	// This indicates group name for the Batch.
	GroupName *string `json:"groupName,omitempty"`
	//  The name of the Batch.
	Name *string `json:"name,omitempty"`
	//  The status of the batch will shows as SUBMITED
	Status *string `json:"status,omitempty"`
}

// NewProcessShipmentResponse instantiates a new ProcessShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessShipmentResponse() *ProcessShipmentResponse {
	this := ProcessShipmentResponse{}
	return &this
}

// NewProcessShipmentResponseWithDefaults instantiates a new ProcessShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessShipmentResponseWithDefaults() *ProcessShipmentResponse {
	this := ProcessShipmentResponse{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *ProcessShipmentResponse) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessShipmentResponse) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *ProcessShipmentResponse) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *ProcessShipmentResponse) SetBatchId(v string) {
	o.BatchId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ProcessShipmentResponse) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessShipmentResponse) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ProcessShipmentResponse) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ProcessShipmentResponse) SetGroupName(v string) {
	o.GroupName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessShipmentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessShipmentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessShipmentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessShipmentResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProcessShipmentResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessShipmentResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProcessShipmentResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProcessShipmentResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ProcessShipmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableProcessShipmentResponse struct {
	value *ProcessShipmentResponse
	isSet bool
}

func (v NullableProcessShipmentResponse) Get() *ProcessShipmentResponse {
	return v.value
}

func (v *NullableProcessShipmentResponse) Set(val *ProcessShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessShipmentResponse(val *ProcessShipmentResponse) *NullableProcessShipmentResponse {
	return &NullableProcessShipmentResponse{value: val, isSet: true}
}

func (v NullableProcessShipmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


