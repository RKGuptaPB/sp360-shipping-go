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

// GetRates200Response - struct for GetRates200Response
type GetRates200Response struct {
	RateShopResponse *RateShopResponse
	SingleRateResponse *SingleRateResponse
}

// RateShopResponseAsGetRates200Response is a convenience function that returns RateShopResponse wrapped in GetRates200Response
func RateShopResponseAsGetRates200Response(v *RateShopResponse) GetRates200Response {
	return GetRates200Response{
		RateShopResponse: v,
	}
}

// SingleRateResponseAsGetRates200Response is a convenience function that returns SingleRateResponse wrapped in GetRates200Response
func SingleRateResponseAsGetRates200Response(v *SingleRateResponse) GetRates200Response {
	return GetRates200Response{
		SingleRateResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRates200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RateShopResponse
	err = newStrictDecoder(data).Decode(&dst.RateShopResponse)
	if err == nil {
		jsonRateShopResponse, _ := json.Marshal(dst.RateShopResponse)
		if string(jsonRateShopResponse) == "{}" { // empty struct
			dst.RateShopResponse = nil
		} else {
			if err = validator.Validate(dst.RateShopResponse); err != nil {
				dst.RateShopResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.RateShopResponse = nil
	}

	// try to unmarshal data into SingleRateResponse
	err = newStrictDecoder(data).Decode(&dst.SingleRateResponse)
	if err == nil {
		jsonSingleRateResponse, _ := json.Marshal(dst.SingleRateResponse)
		if string(jsonSingleRateResponse) == "{}" { // empty struct
			dst.SingleRateResponse = nil
		} else {
			if err = validator.Validate(dst.SingleRateResponse); err != nil {
				dst.SingleRateResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.SingleRateResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RateShopResponse = nil
		dst.SingleRateResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetRates200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetRates200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRates200Response) MarshalJSON() ([]byte, error) {
	if src.RateShopResponse != nil {
		return json.Marshal(&src.RateShopResponse)
	}

	if src.SingleRateResponse != nil {
		return json.Marshal(&src.SingleRateResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRates200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RateShopResponse != nil {
		return obj.RateShopResponse
	}

	if obj.SingleRateResponse != nil {
		return obj.SingleRateResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetRates200Response struct {
	value *GetRates200Response
	isSet bool
}

func (v NullableGetRates200Response) Get() *GetRates200Response {
	return v.value
}

func (v *NullableGetRates200Response) Set(val *GetRates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRates200Response(val *GetRates200Response) *NullableGetRates200Response {
	return &NullableGetRates200Response{value: val, isSet: true}
}

func (v NullableGetRates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


