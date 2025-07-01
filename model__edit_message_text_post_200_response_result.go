/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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
)


// EditMessageTextPost200ResponseResult struct for EditMessageTextPost200ResponseResult
type EditMessageTextPost200ResponseResult struct {
	Message *Message
	Bool *bool
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EditMessageTextPost200ResponseResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Message
	err = json.Unmarshal(data, &dst.Message);
	if err == nil {
		jsonMessage, _ := json.Marshal(dst.Message)
		if string(jsonMessage) == "{}" { // empty struct
			dst.Message = nil
		} else {
			return nil // data stored in dst.Message, return on the first match
		}
	} else {
		dst.Message = nil
	}

	// try to unmarshal JSON data into Bool
	err = json.Unmarshal(data, &dst.Bool);
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			return nil // data stored in dst.Bool, return on the first match
		}
	} else {
		dst.Bool = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EditMessageTextPost200ResponseResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EditMessageTextPost200ResponseResult) MarshalJSON() ([]byte, error) {
	if src.Message != nil {
		return json.Marshal(&src.Message)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableEditMessageTextPost200ResponseResult struct {
	value *EditMessageTextPost200ResponseResult
	isSet bool
}

func (v NullableEditMessageTextPost200ResponseResult) Get() *EditMessageTextPost200ResponseResult {
	return v.value
}

func (v *NullableEditMessageTextPost200ResponseResult) Set(val *EditMessageTextPost200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageTextPost200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageTextPost200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageTextPost200ResponseResult(val *EditMessageTextPost200ResponseResult) *NullableEditMessageTextPost200ResponseResult {
	return &NullableEditMessageTextPost200ResponseResult{value: val, isSet: true}
}

func (v NullableEditMessageTextPost200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageTextPost200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


