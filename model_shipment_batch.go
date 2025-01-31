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

// checks if the ShipmentBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentBatch{}

// ShipmentBatch struct for ShipmentBatch
type ShipmentBatch struct {
	// This is the system generated Batch ID.
	BatchId *string `json:"batchId,omitempty"`
	//  The group name of the Batch.
	GroupName *string `json:"groupName,omitempty"`
	// Name of the Batch.
	Name *string `json:"name,omitempty"`
	Status *JobStatus `json:"status,omitempty"`
	// If this is an import batch, this is the URL to upload the .CSV file.
	UploadURL *string `json:"uploadURL,omitempty"`
}

// NewShipmentBatch instantiates a new ShipmentBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentBatch() *ShipmentBatch {
	this := ShipmentBatch{}
	return &this
}

// NewShipmentBatchWithDefaults instantiates a new ShipmentBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentBatchWithDefaults() *ShipmentBatch {
	this := ShipmentBatch{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *ShipmentBatch) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentBatch) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *ShipmentBatch) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *ShipmentBatch) SetBatchId(v string) {
	o.BatchId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ShipmentBatch) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentBatch) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ShipmentBatch) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ShipmentBatch) SetGroupName(v string) {
	o.GroupName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipmentBatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentBatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipmentBatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipmentBatch) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ShipmentBatch) GetStatus() JobStatus {
	if o == nil || IsNil(o.Status) {
		var ret JobStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentBatch) GetStatusOk() (*JobStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ShipmentBatch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given JobStatus and assigns it to the Status field.
func (o *ShipmentBatch) SetStatus(v JobStatus) {
	o.Status = &v
}

// GetUploadURL returns the UploadURL field value if set, zero value otherwise.
func (o *ShipmentBatch) GetUploadURL() string {
	if o == nil || IsNil(o.UploadURL) {
		var ret string
		return ret
	}
	return *o.UploadURL
}

// GetUploadURLOk returns a tuple with the UploadURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentBatch) GetUploadURLOk() (*string, bool) {
	if o == nil || IsNil(o.UploadURL) {
		return nil, false
	}
	return o.UploadURL, true
}

// HasUploadURL returns a boolean if a field has been set.
func (o *ShipmentBatch) HasUploadURL() bool {
	if o != nil && !IsNil(o.UploadURL) {
		return true
	}

	return false
}

// SetUploadURL gets a reference to the given string and assigns it to the UploadURL field.
func (o *ShipmentBatch) SetUploadURL(v string) {
	o.UploadURL = &v
}

func (o ShipmentBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentBatch) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UploadURL) {
		toSerialize["uploadURL"] = o.UploadURL
	}
	return toSerialize, nil
}

type NullableShipmentBatch struct {
	value *ShipmentBatch
	isSet bool
}

func (v NullableShipmentBatch) Get() *ShipmentBatch {
	return v.value
}

func (v *NullableShipmentBatch) Set(val *ShipmentBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentBatch(val *ShipmentBatch) *NullableShipmentBatch {
	return &NullableShipmentBatch{value: val, isSet: true}
}

func (v NullableShipmentBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


