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

// MultipieceRates200Response - struct for MultipieceRates200Response
type MultipieceRates200Response struct {
	MultipieceRateShopResponse *MultipieceRateShopResponse
	MultipieceRatesResponse *MultipieceRatesResponse
}

// MultipieceRateShopResponseAsMultipieceRates200Response is a convenience function that returns MultipieceRateShopResponse wrapped in MultipieceRates200Response
func MultipieceRateShopResponseAsMultipieceRates200Response(v *MultipieceRateShopResponse) MultipieceRates200Response {
	return MultipieceRates200Response{
		MultipieceRateShopResponse: v,
	}
}

// MultipieceRatesResponseAsMultipieceRates200Response is a convenience function that returns MultipieceRatesResponse wrapped in MultipieceRates200Response
func MultipieceRatesResponseAsMultipieceRates200Response(v *MultipieceRatesResponse) MultipieceRates200Response {
	return MultipieceRates200Response{
		MultipieceRatesResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MultipieceRates200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MultipieceRateShopResponse
	err = newStrictDecoder(data).Decode(&dst.MultipieceRateShopResponse)
	if err == nil {
		jsonMultipieceRateShopResponse, _ := json.Marshal(dst.MultipieceRateShopResponse)
		if string(jsonMultipieceRateShopResponse) == "{}" { // empty struct
			dst.MultipieceRateShopResponse = nil
		} else {
			if err = validator.Validate(dst.MultipieceRateShopResponse); err != nil {
				dst.MultipieceRateShopResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MultipieceRateShopResponse = nil
	}

	// try to unmarshal data into MultipieceRatesResponse
	err = newStrictDecoder(data).Decode(&dst.MultipieceRatesResponse)
	if err == nil {
		jsonMultipieceRatesResponse, _ := json.Marshal(dst.MultipieceRatesResponse)
		if string(jsonMultipieceRatesResponse) == "{}" { // empty struct
			dst.MultipieceRatesResponse = nil
		} else {
			if err = validator.Validate(dst.MultipieceRatesResponse); err != nil {
				dst.MultipieceRatesResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MultipieceRatesResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MultipieceRateShopResponse = nil
		dst.MultipieceRatesResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MultipieceRates200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MultipieceRates200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MultipieceRates200Response) MarshalJSON() ([]byte, error) {
	if src.MultipieceRateShopResponse != nil {
		return json.Marshal(&src.MultipieceRateShopResponse)
	}

	if src.MultipieceRatesResponse != nil {
		return json.Marshal(&src.MultipieceRatesResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MultipieceRates200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MultipieceRateShopResponse != nil {
		return obj.MultipieceRateShopResponse
	}

	if obj.MultipieceRatesResponse != nil {
		return obj.MultipieceRatesResponse
	}

	// all schemas are nil
	return nil
}

type NullableMultipieceRates200Response struct {
	value *MultipieceRates200Response
	isSet bool
}

func (v NullableMultipieceRates200Response) Get() *MultipieceRates200Response {
	return v.value
}

func (v *NullableMultipieceRates200Response) Set(val *MultipieceRates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceRates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceRates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceRates200Response(val *MultipieceRates200Response) *NullableMultipieceRates200Response {
	return &NullableMultipieceRates200Response{value: val, isSet: true}
}

func (v NullableMultipieceRates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceRates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


