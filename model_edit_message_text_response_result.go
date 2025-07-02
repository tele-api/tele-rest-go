/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Telegram Bot API service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package tele_rest

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// EditMessageTextResponseResult - struct for EditMessageTextResponseResult
type EditMessageTextResponseResult struct {
	Message *Message
	Bool *bool
}

// MessageAsEditMessageTextResponseResult is a convenience function that returns Message wrapped in EditMessageTextResponseResult
func MessageAsEditMessageTextResponseResult(v *Message) EditMessageTextResponseResult {
	return EditMessageTextResponseResult{
		Message: v,
	}
}

// boolAsEditMessageTextResponseResult is a convenience function that returns bool wrapped in EditMessageTextResponseResult
func BoolAsEditMessageTextResponseResult(v *bool) EditMessageTextResponseResult {
	return EditMessageTextResponseResult{
		Bool: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EditMessageTextResponseResult) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Message
	err = newStrictDecoder(data).Decode(&dst.Message)
	if err == nil {
		jsonMessage, _ := json.Marshal(dst.Message)
		if string(jsonMessage) == "{}" { // empty struct
			dst.Message = nil
		} else {
			if err = validator.Validate(dst.Message); err != nil {
				dst.Message = nil
			} else {
				match++
			}
		}
	} else {
		dst.Message = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Message = nil
		dst.Bool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EditMessageTextResponseResult)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EditMessageTextResponseResult)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EditMessageTextResponseResult) MarshalJSON() ([]byte, error) {
	if src.Message != nil {
		return json.Marshal(&src.Message)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EditMessageTextResponseResult) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Message != nil {
		return obj.Message
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EditMessageTextResponseResult) GetActualInstanceValue() (interface{}) {
	if obj.Message != nil {
		return *obj.Message
	}

	if obj.Bool != nil {
		return *obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableEditMessageTextResponseResult struct {
	value *EditMessageTextResponseResult
	isSet bool
}

func (v NullableEditMessageTextResponseResult) Get() *EditMessageTextResponseResult {
	return v.value
}

func (v *NullableEditMessageTextResponseResult) Set(val *EditMessageTextResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageTextResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageTextResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageTextResponseResult(val *EditMessageTextResponseResult) *NullableEditMessageTextResponseResult {
	return &NullableEditMessageTextResponseResult{value: val, isSet: true}
}

func (v NullableEditMessageTextResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageTextResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


