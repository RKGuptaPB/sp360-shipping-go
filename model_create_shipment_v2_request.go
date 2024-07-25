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

// CreateShipmentV2Request - struct for CreateShipmentV2Request
type CreateShipmentV2Request struct {
	ShipmentDomesticByRateGroup *ShipmentDomesticByRateGroup
	ShipmentDomesticByRulesetCarrier *ShipmentDomesticByRulesetCarrier
	ShipmentDomesticByRulesetRateGroup *ShipmentDomesticByRulesetRateGroup
	ShipmentDomesticWithCarrier *ShipmentDomesticWithCarrier
	ShipmentDomesticWithCarrierAccount *ShipmentDomesticWithCarrierAccount
}

// ShipmentDomesticByRateGroupAsCreateShipmentV2Request is a convenience function that returns ShipmentDomesticByRateGroup wrapped in CreateShipmentV2Request
func ShipmentDomesticByRateGroupAsCreateShipmentV2Request(v *ShipmentDomesticByRateGroup) CreateShipmentV2Request {
	return CreateShipmentV2Request{
		ShipmentDomesticByRateGroup: v,
	}
}

// ShipmentDomesticByRulesetCarrierAsCreateShipmentV2Request is a convenience function that returns ShipmentDomesticByRulesetCarrier wrapped in CreateShipmentV2Request
func ShipmentDomesticByRulesetCarrierAsCreateShipmentV2Request(v *ShipmentDomesticByRulesetCarrier) CreateShipmentV2Request {
	return CreateShipmentV2Request{
		ShipmentDomesticByRulesetCarrier: v,
	}
}

// ShipmentDomesticByRulesetRateGroupAsCreateShipmentV2Request is a convenience function that returns ShipmentDomesticByRulesetRateGroup wrapped in CreateShipmentV2Request
func ShipmentDomesticByRulesetRateGroupAsCreateShipmentV2Request(v *ShipmentDomesticByRulesetRateGroup) CreateShipmentV2Request {
	return CreateShipmentV2Request{
		ShipmentDomesticByRulesetRateGroup: v,
	}
}

// ShipmentDomesticWithCarrierAsCreateShipmentV2Request is a convenience function that returns ShipmentDomesticWithCarrier wrapped in CreateShipmentV2Request
func ShipmentDomesticWithCarrierAsCreateShipmentV2Request(v *ShipmentDomesticWithCarrier) CreateShipmentV2Request {
	return CreateShipmentV2Request{
		ShipmentDomesticWithCarrier: v,
	}
}

// ShipmentDomesticWithCarrierAccountAsCreateShipmentV2Request is a convenience function that returns ShipmentDomesticWithCarrierAccount wrapped in CreateShipmentV2Request
func ShipmentDomesticWithCarrierAccountAsCreateShipmentV2Request(v *ShipmentDomesticWithCarrierAccount) CreateShipmentV2Request {
	return CreateShipmentV2Request{
		ShipmentDomesticWithCarrierAccount: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateShipmentV2Request) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ShipmentDomesticByRateGroup
	err = newStrictDecoder(data).Decode(&dst.ShipmentDomesticByRateGroup)
	if err == nil {
		jsonShipmentDomesticByRateGroup, _ := json.Marshal(dst.ShipmentDomesticByRateGroup)
		if string(jsonShipmentDomesticByRateGroup) == "{}" { // empty struct
			dst.ShipmentDomesticByRateGroup = nil
		} else {
			if err = validator.Validate(dst.ShipmentDomesticByRateGroup); err != nil {
				dst.ShipmentDomesticByRateGroup = nil
			} else {
				match++
			}
		}
	} else {
		dst.ShipmentDomesticByRateGroup = nil
	}

	// try to unmarshal data into ShipmentDomesticByRulesetCarrier
	err = newStrictDecoder(data).Decode(&dst.ShipmentDomesticByRulesetCarrier)
	if err == nil {
		jsonShipmentDomesticByRulesetCarrier, _ := json.Marshal(dst.ShipmentDomesticByRulesetCarrier)
		if string(jsonShipmentDomesticByRulesetCarrier) == "{}" { // empty struct
			dst.ShipmentDomesticByRulesetCarrier = nil
		} else {
			if err = validator.Validate(dst.ShipmentDomesticByRulesetCarrier); err != nil {
				dst.ShipmentDomesticByRulesetCarrier = nil
			} else {
				match++
			}
		}
	} else {
		dst.ShipmentDomesticByRulesetCarrier = nil
	}

	// try to unmarshal data into ShipmentDomesticByRulesetRateGroup
	err = newStrictDecoder(data).Decode(&dst.ShipmentDomesticByRulesetRateGroup)
	if err == nil {
		jsonShipmentDomesticByRulesetRateGroup, _ := json.Marshal(dst.ShipmentDomesticByRulesetRateGroup)
		if string(jsonShipmentDomesticByRulesetRateGroup) == "{}" { // empty struct
			dst.ShipmentDomesticByRulesetRateGroup = nil
		} else {
			if err = validator.Validate(dst.ShipmentDomesticByRulesetRateGroup); err != nil {
				dst.ShipmentDomesticByRulesetRateGroup = nil
			} else {
				match++
			}
		}
	} else {
		dst.ShipmentDomesticByRulesetRateGroup = nil
	}

	// try to unmarshal data into ShipmentDomesticWithCarrier
	err = newStrictDecoder(data).Decode(&dst.ShipmentDomesticWithCarrier)
	if err == nil {
		jsonShipmentDomesticWithCarrier, _ := json.Marshal(dst.ShipmentDomesticWithCarrier)
		if string(jsonShipmentDomesticWithCarrier) == "{}" { // empty struct
			dst.ShipmentDomesticWithCarrier = nil
		} else {
			if err = validator.Validate(dst.ShipmentDomesticWithCarrier); err != nil {
				dst.ShipmentDomesticWithCarrier = nil
			} else {
				match++
			}
		}
	} else {
		dst.ShipmentDomesticWithCarrier = nil
	}

	// try to unmarshal data into ShipmentDomesticWithCarrierAccount
	err = newStrictDecoder(data).Decode(&dst.ShipmentDomesticWithCarrierAccount)
	if err == nil {
		jsonShipmentDomesticWithCarrierAccount, _ := json.Marshal(dst.ShipmentDomesticWithCarrierAccount)
		if string(jsonShipmentDomesticWithCarrierAccount) == "{}" { // empty struct
			dst.ShipmentDomesticWithCarrierAccount = nil
		} else {
			if err = validator.Validate(dst.ShipmentDomesticWithCarrierAccount); err != nil {
				dst.ShipmentDomesticWithCarrierAccount = nil
			} else {
				match++
			}
		}
	} else {
		dst.ShipmentDomesticWithCarrierAccount = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ShipmentDomesticByRateGroup = nil
		dst.ShipmentDomesticByRulesetCarrier = nil
		dst.ShipmentDomesticByRulesetRateGroup = nil
		dst.ShipmentDomesticWithCarrier = nil
		dst.ShipmentDomesticWithCarrierAccount = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateShipmentV2Request)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateShipmentV2Request)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateShipmentV2Request) MarshalJSON() ([]byte, error) {
	if src.ShipmentDomesticByRateGroup != nil {
		return json.Marshal(&src.ShipmentDomesticByRateGroup)
	}

	if src.ShipmentDomesticByRulesetCarrier != nil {
		return json.Marshal(&src.ShipmentDomesticByRulesetCarrier)
	}

	if src.ShipmentDomesticByRulesetRateGroup != nil {
		return json.Marshal(&src.ShipmentDomesticByRulesetRateGroup)
	}

	if src.ShipmentDomesticWithCarrier != nil {
		return json.Marshal(&src.ShipmentDomesticWithCarrier)
	}

	if src.ShipmentDomesticWithCarrierAccount != nil {
		return json.Marshal(&src.ShipmentDomesticWithCarrierAccount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateShipmentV2Request) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ShipmentDomesticByRateGroup != nil {
		return obj.ShipmentDomesticByRateGroup
	}

	if obj.ShipmentDomesticByRulesetCarrier != nil {
		return obj.ShipmentDomesticByRulesetCarrier
	}

	if obj.ShipmentDomesticByRulesetRateGroup != nil {
		return obj.ShipmentDomesticByRulesetRateGroup
	}

	if obj.ShipmentDomesticWithCarrier != nil {
		return obj.ShipmentDomesticWithCarrier
	}

	if obj.ShipmentDomesticWithCarrierAccount != nil {
		return obj.ShipmentDomesticWithCarrierAccount
	}

	// all schemas are nil
	return nil
}

type NullableCreateShipmentV2Request struct {
	value *CreateShipmentV2Request
	isSet bool
}

func (v NullableCreateShipmentV2Request) Get() *CreateShipmentV2Request {
	return v.value
}

func (v *NullableCreateShipmentV2Request) Set(val *CreateShipmentV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipmentV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipmentV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipmentV2Request(val *CreateShipmentV2Request) *NullableCreateShipmentV2Request {
	return &NullableCreateShipmentV2Request{value: val, isSet: true}
}

func (v NullableCreateShipmentV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipmentV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


