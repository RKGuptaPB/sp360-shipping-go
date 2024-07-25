/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// CreateShipment200Response - struct for CreateShipment200Response
type CreateShipment200Response struct {
	DomesticShipmentResponse *DomesticShipmentResponse
	InternationalShipmentResponse *InternationalShipmentResponse
}

// DomesticShipmentResponseAsCreateShipment200Response is a convenience function that returns DomesticShipmentResponse wrapped in CreateShipment200Response
func DomesticShipmentResponseAsCreateShipment200Response(v *DomesticShipmentResponse) CreateShipment200Response {
	return CreateShipment200Response{
		DomesticShipmentResponse: v,
	}
}

// InternationalShipmentResponseAsCreateShipment200Response is a convenience function that returns InternationalShipmentResponse wrapped in CreateShipment200Response
func InternationalShipmentResponseAsCreateShipment200Response(v *InternationalShipmentResponse) CreateShipment200Response {
	return CreateShipment200Response{
		InternationalShipmentResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateShipment200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DomesticShipmentResponse
	err = newStrictDecoder(data).Decode(&dst.DomesticShipmentResponse)
	if err == nil {
		jsonDomesticShipmentResponse, _ := json.Marshal(dst.DomesticShipmentResponse)
		if string(jsonDomesticShipmentResponse) == "{}" { // empty struct
			dst.DomesticShipmentResponse = nil
		} else {
			if err = validator.Validate(dst.DomesticShipmentResponse); err != nil {
				dst.DomesticShipmentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.DomesticShipmentResponse = nil
	}

	// try to unmarshal data into InternationalShipmentResponse
	err = newStrictDecoder(data).Decode(&dst.InternationalShipmentResponse)
	if err == nil {
		jsonInternationalShipmentResponse, _ := json.Marshal(dst.InternationalShipmentResponse)
		if string(jsonInternationalShipmentResponse) == "{}" { // empty struct
			dst.InternationalShipmentResponse = nil
		} else {
			if err = validator.Validate(dst.InternationalShipmentResponse); err != nil {
				dst.InternationalShipmentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.InternationalShipmentResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DomesticShipmentResponse = nil
		dst.InternationalShipmentResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateShipment200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateShipment200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateShipment200Response) MarshalJSON() ([]byte, error) {
	if src.DomesticShipmentResponse != nil {
		return json.Marshal(&src.DomesticShipmentResponse)
	}

	if src.InternationalShipmentResponse != nil {
		return json.Marshal(&src.InternationalShipmentResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateShipment200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DomesticShipmentResponse != nil {
		return obj.DomesticShipmentResponse
	}

	if obj.InternationalShipmentResponse != nil {
		return obj.InternationalShipmentResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateShipment200Response struct {
	value *CreateShipment200Response
	isSet bool
}

func (v NullableCreateShipment200Response) Get() *CreateShipment200Response {
	return v.value
}

func (v *NullableCreateShipment200Response) Set(val *CreateShipment200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipment200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipment200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipment200Response(val *CreateShipment200Response) *NullableCreateShipment200Response {
	return &NullableCreateShipment200Response{value: val, isSet: true}
}

func (v NullableCreateShipment200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipment200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


