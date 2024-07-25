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

// checks if the CreateBulkShipmentsERRCoversheet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBulkShipmentsERRCoversheet{}

// CreateBulkShipmentsERRCoversheet ShipmentBatch information contains following schema
type CreateBulkShipmentsERRCoversheet struct {
	GroupName *string `json:"groupName,omitempty"`
	// This indicates the envelope of the Bulk Shipment, e.g., DocSize can be 10', 6'X 9', 6'X 9.5' or 9.5'X 12'
	Size string `json:"size"`
	// This indicates the type of the Batch Shipment, e.g., Coversheet.
	Type string `json:"type"`
	// This defines the type of the shipment which is printed. For example Coversheet prints in PDF form.
	Format *string `json:"format,omitempty"`
	Name string `json:"name"`
	// Default `carrierAccountId` to be used for this batch. You can override this value by defining it at shipment level.
	CarrierAccountId string `json:"carrierAccountId"`
	// Default `parcelType` specific to the carrier, e.g., LTR, LGENV to be used for this batch. You can override this value by defining it at shipment level.
	ParcelType string `json:"parcelType"`
	// Default abbreviated name `serviceId` of the carrier-specific service to be used for this batch. You can override this value by defining it at shipment level.
	ServiceId string `json:"serviceId"`
	SpecialServices []SpecialServiceERRInner `json:"specialServices,omitempty"`
	Shipments []ShipmentERRCoversheet `json:"shipments"`
}

type _CreateBulkShipmentsERRCoversheet CreateBulkShipmentsERRCoversheet

// NewCreateBulkShipmentsERRCoversheet instantiates a new CreateBulkShipmentsERRCoversheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBulkShipmentsERRCoversheet(size string, type_ string, name string, carrierAccountId string, parcelType string, serviceId string, shipments []ShipmentERRCoversheet) *CreateBulkShipmentsERRCoversheet {
	this := CreateBulkShipmentsERRCoversheet{}
	this.Size = size
	this.Type = type_
	this.Name = name
	this.CarrierAccountId = carrierAccountId
	this.ParcelType = parcelType
	this.ServiceId = serviceId
	this.Shipments = shipments
	return &this
}

// NewCreateBulkShipmentsERRCoversheetWithDefaults instantiates a new CreateBulkShipmentsERRCoversheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBulkShipmentsERRCoversheetWithDefaults() *CreateBulkShipmentsERRCoversheet {
	this := CreateBulkShipmentsERRCoversheet{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRCoversheet) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRCoversheet) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *CreateBulkShipmentsERRCoversheet) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSize returns the Size field value
func (o *CreateBulkShipmentsERRCoversheet) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetSize(v string) {
	o.Size = v
}

// GetType returns the Type field value
func (o *CreateBulkShipmentsERRCoversheet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetType(v string) {
	o.Type = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRCoversheet) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRCoversheet) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateBulkShipmentsERRCoversheet) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value
func (o *CreateBulkShipmentsERRCoversheet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetName(v string) {
	o.Name = v
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *CreateBulkShipmentsERRCoversheet) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetParcelType returns the ParcelType field value
func (o *CreateBulkShipmentsERRCoversheet) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetParcelType(v string) {
	o.ParcelType = v
}

// GetServiceId returns the ServiceId field value
func (o *CreateBulkShipmentsERRCoversheet) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRCoversheet) GetSpecialServices() []SpecialServiceERRInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialServiceERRInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetSpecialServicesOk() ([]SpecialServiceERRInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRCoversheet) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialServiceERRInner and assigns it to the SpecialServices field.
func (o *CreateBulkShipmentsERRCoversheet) SetSpecialServices(v []SpecialServiceERRInner) {
	o.SpecialServices = v
}

// GetShipments returns the Shipments field value
func (o *CreateBulkShipmentsERRCoversheet) GetShipments() []ShipmentERRCoversheet {
	if o == nil {
		var ret []ShipmentERRCoversheet
		return ret
	}

	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRCoversheet) GetShipmentsOk() ([]ShipmentERRCoversheet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shipments, true
}

// SetShipments sets field value
func (o *CreateBulkShipmentsERRCoversheet) SetShipments(v []ShipmentERRCoversheet) {
	o.Shipments = v
}

func (o CreateBulkShipmentsERRCoversheet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBulkShipmentsERRCoversheet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["name"] = o.Name
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["parcelType"] = o.ParcelType
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	toSerialize["shipments"] = o.Shipments
	return toSerialize, nil
}

func (o *CreateBulkShipmentsERRCoversheet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
		"type",
		"name",
		"carrierAccountId",
		"parcelType",
		"serviceId",
		"shipments",
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

	varCreateBulkShipmentsERRCoversheet := _CreateBulkShipmentsERRCoversheet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBulkShipmentsERRCoversheet)

	if err != nil {
		return err
	}

	*o = CreateBulkShipmentsERRCoversheet(varCreateBulkShipmentsERRCoversheet)

	return err
}

type NullableCreateBulkShipmentsERRCoversheet struct {
	value *CreateBulkShipmentsERRCoversheet
	isSet bool
}

func (v NullableCreateBulkShipmentsERRCoversheet) Get() *CreateBulkShipmentsERRCoversheet {
	return v.value
}

func (v *NullableCreateBulkShipmentsERRCoversheet) Set(val *CreateBulkShipmentsERRCoversheet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBulkShipmentsERRCoversheet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBulkShipmentsERRCoversheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBulkShipmentsERRCoversheet(val *CreateBulkShipmentsERRCoversheet) *NullableCreateBulkShipmentsERRCoversheet {
	return &NullableCreateBulkShipmentsERRCoversheet{value: val, isSet: true}
}

func (v NullableCreateBulkShipmentsERRCoversheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBulkShipmentsERRCoversheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


