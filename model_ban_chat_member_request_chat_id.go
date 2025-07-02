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

// BanChatMemberRequestChatId - Unique identifier for the target group or username of the target supergroup or channel (in the format `@channelusername`)
type BanChatMemberRequestChatId struct {
	Int32 *int32
	String *string
}

// int32AsBanChatMemberRequestChatId is a convenience function that returns int32 wrapped in BanChatMemberRequestChatId
func Int32AsBanChatMemberRequestChatId(v *int32) BanChatMemberRequestChatId {
	return BanChatMemberRequestChatId{
		Int32: v,
	}
}

// stringAsBanChatMemberRequestChatId is a convenience function that returns string wrapped in BanChatMemberRequestChatId
func StringAsBanChatMemberRequestChatId(v *string) BanChatMemberRequestChatId {
	return BanChatMemberRequestChatId{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BanChatMemberRequestChatId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BanChatMemberRequestChatId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BanChatMemberRequestChatId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BanChatMemberRequestChatId) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BanChatMemberRequestChatId) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BanChatMemberRequestChatId) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableBanChatMemberRequestChatId struct {
	value *BanChatMemberRequestChatId
	isSet bool
}

func (v NullableBanChatMemberRequestChatId) Get() *BanChatMemberRequestChatId {
	return v.value
}

func (v *NullableBanChatMemberRequestChatId) Set(val *BanChatMemberRequestChatId) {
	v.value = val
	v.isSet = true
}

func (v NullableBanChatMemberRequestChatId) IsSet() bool {
	return v.isSet
}

func (v *NullableBanChatMemberRequestChatId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanChatMemberRequestChatId(val *BanChatMemberRequestChatId) *NullableBanChatMemberRequestChatId {
	return &NullableBanChatMemberRequestChatId{value: val, isSet: true}
}

func (v NullableBanChatMemberRequestChatId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanChatMemberRequestChatId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


