/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ShipmentERR type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentERR{}

// ShipmentERR struct for ShipmentERR
type ShipmentERR struct {
	// The external ID of the shipment. User can provide any custom value to it for its own reference
	ExternalId *string `json:"externalId,omitempty"`
	FromAddress ShipmentFromAddress `json:"fromAddress"`
	Parcel ShipmentParcel `json:"parcel"`
	// A unique identifier associated with the Carrier account used by client users during shipment process.
	CarrierAccountId string `json:"carrierAccountId"`
	// >-Packaging type specific to the carrier, e.g., PKG, LGENV.
	ParcelType string `json:"parcelType"`
	// >-The abbreviated name of the carrier-specific service. Required for creating a shipment. Optional for rating a parcel.
	ServiceId string `json:"serviceId"`
	// Current Shipment date
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	SpecialServices []SpecialServiceERRInner `json:"specialServices,omitempty"`
	ShipmentOptions *ShipmentOptions `json:"shipmentOptions,omitempty"`
	// Additional metadata that needs to be stored for this shipment can be added here. For now, `costAccountName` is supported.
	Metadata []ShipmentMetadataInner `json:"metadata,omitempty"`
	ToAddress ShipmentToAddress `json:"toAddress"`
}

type _ShipmentERR ShipmentERR

// NewShipmentERR instantiates a new ShipmentERR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentERR(fromAddress ShipmentFromAddress, parcel ShipmentParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentToAddress) *ShipmentERR {
	this := ShipmentERR{}
	this.FromAddress = fromAddress
	this.Parcel = parcel
	this.CarrierAccountId = carrierAccountId
	this.ParcelType = parcelType
	this.ServiceId = serviceId
	this.ToAddress = toAddress
	return &this
}

// NewShipmentERRWithDefaults instantiates a new ShipmentERR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentERRWithDefaults() *ShipmentERR {
	this := ShipmentERR{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ShipmentERR) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ShipmentERR) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ShipmentERR) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFromAddress returns the FromAddress field value
func (o *ShipmentERR) GetFromAddress() ShipmentFromAddress {
	if o == nil {
		var ret ShipmentFromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetFromAddressOk() (*ShipmentFromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *ShipmentERR) SetFromAddress(v ShipmentFromAddress) {
	o.FromAddress = v
}

// GetParcel returns the Parcel field value
func (o *ShipmentERR) GetParcel() ShipmentParcel {
	if o == nil {
		var ret ShipmentParcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetParcelOk() (*ShipmentParcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *ShipmentERR) SetParcel(v ShipmentParcel) {
	o.Parcel = v
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *ShipmentERR) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *ShipmentERR) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetParcelType returns the ParcelType field value
func (o *ShipmentERR) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *ShipmentERR) SetParcelType(v string) {
	o.ParcelType = v
}

// GetServiceId returns the ServiceId field value
func (o *ShipmentERR) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ShipmentERR) SetServiceId(v string) {
	o.ServiceId = v
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *ShipmentERR) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *ShipmentERR) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *ShipmentERR) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *ShipmentERR) GetSpecialServices() []SpecialServiceERRInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialServiceERRInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetSpecialServicesOk() ([]SpecialServiceERRInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *ShipmentERR) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialServiceERRInner and assigns it to the SpecialServices field.
func (o *ShipmentERR) SetSpecialServices(v []SpecialServiceERRInner) {
	o.SpecialServices = v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *ShipmentERR) GetShipmentOptions() ShipmentOptions {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptions
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetShipmentOptionsOk() (*ShipmentOptions, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *ShipmentERR) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptions and assigns it to the ShipmentOptions field.
func (o *ShipmentERR) SetShipmentOptions(v ShipmentOptions) {
	o.ShipmentOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShipmentERR) GetMetadata() []ShipmentMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []ShipmentMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetMetadataOk() ([]ShipmentMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShipmentERR) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ShipmentMetadataInner and assigns it to the Metadata field.
func (o *ShipmentERR) SetMetadata(v []ShipmentMetadataInner) {
	o.Metadata = v
}

// GetToAddress returns the ToAddress field value
func (o *ShipmentERR) GetToAddress() ShipmentToAddress {
	if o == nil {
		var ret ShipmentToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentERR) GetToAddressOk() (*ShipmentToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *ShipmentERR) SetToAddress(v ShipmentToAddress) {
	o.ToAddress = v
}

func (o ShipmentERR) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentERR) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["parcel"] = o.Parcel
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["parcelType"] = o.ParcelType
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.DateOfShipment) {
		toSerialize["dateOfShipment"] = o.DateOfShipment
	}
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	if !IsNil(o.ShipmentOptions) {
		toSerialize["shipmentOptions"] = o.ShipmentOptions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["toAddress"] = o.ToAddress
	return toSerialize, nil
}

func (o *ShipmentERR) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromAddress",
		"parcel",
		"carrierAccountId",
		"parcelType",
		"serviceId",
		"toAddress",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varShipmentERR := _ShipmentERR{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentERR)

	if err != nil {
		return err
	}

	*o = ShipmentERR(varShipmentERR)

	return err
}

type NullableShipmentERR struct {
	value *ShipmentERR
	isSet bool
}

func (v NullableShipmentERR) Get() *ShipmentERR {
	return v.value
}

func (v *NullableShipmentERR) Set(val *ShipmentERR) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentERR) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentERR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentERR(val *ShipmentERR) *NullableShipmentERR {
	return &NullableShipmentERR{value: val, isSet: true}
}

func (v NullableShipmentERR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentERR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


