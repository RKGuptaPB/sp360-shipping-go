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

// checks if the ReturnLabelResponseLabelLayoutInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLabelResponseLabelLayoutInner{}

// ReturnLabelResponseLabelLayoutInner struct for ReturnLabelResponseLabelLayoutInner
type ReturnLabelResponseLabelLayoutInner struct {
	ContentType *string `json:"contentType,omitempty"`
	Contents *string `json:"contents,omitempty"`
	FileFormat *string `json:"fileFormat,omitempty"`
	Size *string `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewReturnLabelResponseLabelLayoutInner instantiates a new ReturnLabelResponseLabelLayoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLabelResponseLabelLayoutInner() *ReturnLabelResponseLabelLayoutInner {
	this := ReturnLabelResponseLabelLayoutInner{}
	return &this
}

// NewReturnLabelResponseLabelLayoutInnerWithDefaults instantiates a new ReturnLabelResponseLabelLayoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLabelResponseLabelLayoutInnerWithDefaults() *ReturnLabelResponseLabelLayoutInner {
	this := ReturnLabelResponseLabelLayoutInner{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ReturnLabelResponseLabelLayoutInner) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseLabelLayoutInner) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ReturnLabelResponseLabelLayoutInner) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ReturnLabelResponseLabelLayoutInner) SetContentType(v string) {
	o.ContentType = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ReturnLabelResponseLabelLayoutInner) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseLabelLayoutInner) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *ReturnLabelResponseLabelLayoutInner) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *ReturnLabelResponseLabelLayoutInner) SetContents(v string) {
	o.Contents = &v
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *ReturnLabelResponseLabelLayoutInner) GetFileFormat() string {
	if o == nil || IsNil(o.FileFormat) {
		var ret string
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseLabelLayoutInner) GetFileFormatOk() (*string, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *ReturnLabelResponseLabelLayoutInner) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given string and assigns it to the FileFormat field.
func (o *ReturnLabelResponseLabelLayoutInner) SetFileFormat(v string) {
	o.FileFormat = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReturnLabelResponseLabelLayoutInner) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseLabelLayoutInner) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReturnLabelResponseLabelLayoutInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *ReturnLabelResponseLabelLayoutInner) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReturnLabelResponseLabelLayoutInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelResponseLabelLayoutInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReturnLabelResponseLabelLayoutInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReturnLabelResponseLabelLayoutInner) SetType(v string) {
	o.Type = &v
}

func (o ReturnLabelResponseLabelLayoutInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLabelResponseLabelLayoutInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.FileFormat) {
		toSerialize["fileFormat"] = o.FileFormat
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableReturnLabelResponseLabelLayoutInner struct {
	value *ReturnLabelResponseLabelLayoutInner
	isSet bool
}

func (v NullableReturnLabelResponseLabelLayoutInner) Get() *ReturnLabelResponseLabelLayoutInner {
	return v.value
}

func (v *NullableReturnLabelResponseLabelLayoutInner) Set(val *ReturnLabelResponseLabelLayoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLabelResponseLabelLayoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLabelResponseLabelLayoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLabelResponseLabelLayoutInner(val *ReturnLabelResponseLabelLayoutInner) *NullableReturnLabelResponseLabelLayoutInner {
	return &NullableReturnLabelResponseLabelLayoutInner{value: val, isSet: true}
}

func (v NullableReturnLabelResponseLabelLayoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLabelResponseLabelLayoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


