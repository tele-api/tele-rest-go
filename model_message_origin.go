/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// MessageOrigin - This object describes the origin of a message. It can be one of  * [MessageOriginUser](https://core.telegram.org/bots/api/#messageoriginuser) * [MessageOriginHiddenUser](https://core.telegram.org/bots/api/#messageoriginhiddenuser) * [MessageOriginChat](https://core.telegram.org/bots/api/#messageoriginchat) * [MessageOriginChannel](https://core.telegram.org/bots/api/#messageoriginchannel)
type MessageOrigin struct {
	MessageOriginChannel *MessageOriginChannel
	MessageOriginChat *MessageOriginChat
	MessageOriginHiddenUser *MessageOriginHiddenUser
	MessageOriginUser *MessageOriginUser
}

// MessageOriginChannelAsMessageOrigin is a convenience function that returns MessageOriginChannel wrapped in MessageOrigin
func MessageOriginChannelAsMessageOrigin(v *MessageOriginChannel) MessageOrigin {
	return MessageOrigin{
		MessageOriginChannel: v,
	}
}

// MessageOriginChatAsMessageOrigin is a convenience function that returns MessageOriginChat wrapped in MessageOrigin
func MessageOriginChatAsMessageOrigin(v *MessageOriginChat) MessageOrigin {
	return MessageOrigin{
		MessageOriginChat: v,
	}
}

// MessageOriginHiddenUserAsMessageOrigin is a convenience function that returns MessageOriginHiddenUser wrapped in MessageOrigin
func MessageOriginHiddenUserAsMessageOrigin(v *MessageOriginHiddenUser) MessageOrigin {
	return MessageOrigin{
		MessageOriginHiddenUser: v,
	}
}

// MessageOriginUserAsMessageOrigin is a convenience function that returns MessageOriginUser wrapped in MessageOrigin
func MessageOriginUserAsMessageOrigin(v *MessageOriginUser) MessageOrigin {
	return MessageOrigin{
		MessageOriginUser: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageOrigin) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MessageOriginChannel
	err = newStrictDecoder(data).Decode(&dst.MessageOriginChannel)
	if err == nil {
		jsonMessageOriginChannel, _ := json.Marshal(dst.MessageOriginChannel)
		if string(jsonMessageOriginChannel) == "{}" { // empty struct
			dst.MessageOriginChannel = nil
		} else {
			if err = validator.Validate(dst.MessageOriginChannel); err != nil {
				dst.MessageOriginChannel = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageOriginChannel = nil
	}

	// try to unmarshal data into MessageOriginChat
	err = newStrictDecoder(data).Decode(&dst.MessageOriginChat)
	if err == nil {
		jsonMessageOriginChat, _ := json.Marshal(dst.MessageOriginChat)
		if string(jsonMessageOriginChat) == "{}" { // empty struct
			dst.MessageOriginChat = nil
		} else {
			if err = validator.Validate(dst.MessageOriginChat); err != nil {
				dst.MessageOriginChat = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageOriginChat = nil
	}

	// try to unmarshal data into MessageOriginHiddenUser
	err = newStrictDecoder(data).Decode(&dst.MessageOriginHiddenUser)
	if err == nil {
		jsonMessageOriginHiddenUser, _ := json.Marshal(dst.MessageOriginHiddenUser)
		if string(jsonMessageOriginHiddenUser) == "{}" { // empty struct
			dst.MessageOriginHiddenUser = nil
		} else {
			if err = validator.Validate(dst.MessageOriginHiddenUser); err != nil {
				dst.MessageOriginHiddenUser = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageOriginHiddenUser = nil
	}

	// try to unmarshal data into MessageOriginUser
	err = newStrictDecoder(data).Decode(&dst.MessageOriginUser)
	if err == nil {
		jsonMessageOriginUser, _ := json.Marshal(dst.MessageOriginUser)
		if string(jsonMessageOriginUser) == "{}" { // empty struct
			dst.MessageOriginUser = nil
		} else {
			if err = validator.Validate(dst.MessageOriginUser); err != nil {
				dst.MessageOriginUser = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageOriginUser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MessageOriginChannel = nil
		dst.MessageOriginChat = nil
		dst.MessageOriginHiddenUser = nil
		dst.MessageOriginUser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MessageOrigin)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MessageOrigin)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageOrigin) MarshalJSON() ([]byte, error) {
	if src.MessageOriginChannel != nil {
		return json.Marshal(&src.MessageOriginChannel)
	}

	if src.MessageOriginChat != nil {
		return json.Marshal(&src.MessageOriginChat)
	}

	if src.MessageOriginHiddenUser != nil {
		return json.Marshal(&src.MessageOriginHiddenUser)
	}

	if src.MessageOriginUser != nil {
		return json.Marshal(&src.MessageOriginUser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageOrigin) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MessageOriginChannel != nil {
		return obj.MessageOriginChannel
	}

	if obj.MessageOriginChat != nil {
		return obj.MessageOriginChat
	}

	if obj.MessageOriginHiddenUser != nil {
		return obj.MessageOriginHiddenUser
	}

	if obj.MessageOriginUser != nil {
		return obj.MessageOriginUser
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MessageOrigin) GetActualInstanceValue() (interface{}) {
	if obj.MessageOriginChannel != nil {
		return *obj.MessageOriginChannel
	}

	if obj.MessageOriginChat != nil {
		return *obj.MessageOriginChat
	}

	if obj.MessageOriginHiddenUser != nil {
		return *obj.MessageOriginHiddenUser
	}

	if obj.MessageOriginUser != nil {
		return *obj.MessageOriginUser
	}

	// all schemas are nil
	return nil
}

type NullableMessageOrigin struct {
	value *MessageOrigin
	isSet bool
}

func (v NullableMessageOrigin) Get() *MessageOrigin {
	return v.value
}

func (v *NullableMessageOrigin) Set(val *MessageOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageOrigin(val *MessageOrigin) *NullableMessageOrigin {
	return &NullableMessageOrigin{value: val, isSet: true}
}

func (v NullableMessageOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


