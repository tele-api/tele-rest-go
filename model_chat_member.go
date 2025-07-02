/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// ChatMember - This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:  * [ChatMemberOwner](https://core.telegram.org/bots/api/#chatmemberowner) * [ChatMemberAdministrator](https://core.telegram.org/bots/api/#chatmemberadministrator) * [ChatMemberMember](https://core.telegram.org/bots/api/#chatmembermember) * [ChatMemberRestricted](https://core.telegram.org/bots/api/#chatmemberrestricted) * [ChatMemberLeft](https://core.telegram.org/bots/api/#chatmemberleft) * [ChatMemberBanned](https://core.telegram.org/bots/api/#chatmemberbanned)
type ChatMember struct {
	ChatMemberAdministrator *ChatMemberAdministrator
	ChatMemberBanned *ChatMemberBanned
	ChatMemberLeft *ChatMemberLeft
	ChatMemberMember *ChatMemberMember
	ChatMemberOwner *ChatMemberOwner
	ChatMemberRestricted *ChatMemberRestricted
}

// ChatMemberAdministratorAsChatMember is a convenience function that returns ChatMemberAdministrator wrapped in ChatMember
func ChatMemberAdministratorAsChatMember(v *ChatMemberAdministrator) ChatMember {
	return ChatMember{
		ChatMemberAdministrator: v,
	}
}

// ChatMemberBannedAsChatMember is a convenience function that returns ChatMemberBanned wrapped in ChatMember
func ChatMemberBannedAsChatMember(v *ChatMemberBanned) ChatMember {
	return ChatMember{
		ChatMemberBanned: v,
	}
}

// ChatMemberLeftAsChatMember is a convenience function that returns ChatMemberLeft wrapped in ChatMember
func ChatMemberLeftAsChatMember(v *ChatMemberLeft) ChatMember {
	return ChatMember{
		ChatMemberLeft: v,
	}
}

// ChatMemberMemberAsChatMember is a convenience function that returns ChatMemberMember wrapped in ChatMember
func ChatMemberMemberAsChatMember(v *ChatMemberMember) ChatMember {
	return ChatMember{
		ChatMemberMember: v,
	}
}

// ChatMemberOwnerAsChatMember is a convenience function that returns ChatMemberOwner wrapped in ChatMember
func ChatMemberOwnerAsChatMember(v *ChatMemberOwner) ChatMember {
	return ChatMember{
		ChatMemberOwner: v,
	}
}

// ChatMemberRestrictedAsChatMember is a convenience function that returns ChatMemberRestricted wrapped in ChatMember
func ChatMemberRestrictedAsChatMember(v *ChatMemberRestricted) ChatMember {
	return ChatMember{
		ChatMemberRestricted: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChatMember) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatMemberAdministrator
	err = newStrictDecoder(data).Decode(&dst.ChatMemberAdministrator)
	if err == nil {
		jsonChatMemberAdministrator, _ := json.Marshal(dst.ChatMemberAdministrator)
		if string(jsonChatMemberAdministrator) == "{}" { // empty struct
			dst.ChatMemberAdministrator = nil
		} else {
			if err = validator.Validate(dst.ChatMemberAdministrator); err != nil {
				dst.ChatMemberAdministrator = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberAdministrator = nil
	}

	// try to unmarshal data into ChatMemberBanned
	err = newStrictDecoder(data).Decode(&dst.ChatMemberBanned)
	if err == nil {
		jsonChatMemberBanned, _ := json.Marshal(dst.ChatMemberBanned)
		if string(jsonChatMemberBanned) == "{}" { // empty struct
			dst.ChatMemberBanned = nil
		} else {
			if err = validator.Validate(dst.ChatMemberBanned); err != nil {
				dst.ChatMemberBanned = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberBanned = nil
	}

	// try to unmarshal data into ChatMemberLeft
	err = newStrictDecoder(data).Decode(&dst.ChatMemberLeft)
	if err == nil {
		jsonChatMemberLeft, _ := json.Marshal(dst.ChatMemberLeft)
		if string(jsonChatMemberLeft) == "{}" { // empty struct
			dst.ChatMemberLeft = nil
		} else {
			if err = validator.Validate(dst.ChatMemberLeft); err != nil {
				dst.ChatMemberLeft = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberLeft = nil
	}

	// try to unmarshal data into ChatMemberMember
	err = newStrictDecoder(data).Decode(&dst.ChatMemberMember)
	if err == nil {
		jsonChatMemberMember, _ := json.Marshal(dst.ChatMemberMember)
		if string(jsonChatMemberMember) == "{}" { // empty struct
			dst.ChatMemberMember = nil
		} else {
			if err = validator.Validate(dst.ChatMemberMember); err != nil {
				dst.ChatMemberMember = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberMember = nil
	}

	// try to unmarshal data into ChatMemberOwner
	err = newStrictDecoder(data).Decode(&dst.ChatMemberOwner)
	if err == nil {
		jsonChatMemberOwner, _ := json.Marshal(dst.ChatMemberOwner)
		if string(jsonChatMemberOwner) == "{}" { // empty struct
			dst.ChatMemberOwner = nil
		} else {
			if err = validator.Validate(dst.ChatMemberOwner); err != nil {
				dst.ChatMemberOwner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberOwner = nil
	}

	// try to unmarshal data into ChatMemberRestricted
	err = newStrictDecoder(data).Decode(&dst.ChatMemberRestricted)
	if err == nil {
		jsonChatMemberRestricted, _ := json.Marshal(dst.ChatMemberRestricted)
		if string(jsonChatMemberRestricted) == "{}" { // empty struct
			dst.ChatMemberRestricted = nil
		} else {
			if err = validator.Validate(dst.ChatMemberRestricted); err != nil {
				dst.ChatMemberRestricted = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatMemberRestricted = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChatMemberAdministrator = nil
		dst.ChatMemberBanned = nil
		dst.ChatMemberLeft = nil
		dst.ChatMemberMember = nil
		dst.ChatMemberOwner = nil
		dst.ChatMemberRestricted = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChatMember)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChatMember)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatMember) MarshalJSON() ([]byte, error) {
	if src.ChatMemberAdministrator != nil {
		return json.Marshal(&src.ChatMemberAdministrator)
	}

	if src.ChatMemberBanned != nil {
		return json.Marshal(&src.ChatMemberBanned)
	}

	if src.ChatMemberLeft != nil {
		return json.Marshal(&src.ChatMemberLeft)
	}

	if src.ChatMemberMember != nil {
		return json.Marshal(&src.ChatMemberMember)
	}

	if src.ChatMemberOwner != nil {
		return json.Marshal(&src.ChatMemberOwner)
	}

	if src.ChatMemberRestricted != nil {
		return json.Marshal(&src.ChatMemberRestricted)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChatMember) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChatMemberAdministrator != nil {
		return obj.ChatMemberAdministrator
	}

	if obj.ChatMemberBanned != nil {
		return obj.ChatMemberBanned
	}

	if obj.ChatMemberLeft != nil {
		return obj.ChatMemberLeft
	}

	if obj.ChatMemberMember != nil {
		return obj.ChatMemberMember
	}

	if obj.ChatMemberOwner != nil {
		return obj.ChatMemberOwner
	}

	if obj.ChatMemberRestricted != nil {
		return obj.ChatMemberRestricted
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ChatMember) GetActualInstanceValue() (interface{}) {
	if obj.ChatMemberAdministrator != nil {
		return *obj.ChatMemberAdministrator
	}

	if obj.ChatMemberBanned != nil {
		return *obj.ChatMemberBanned
	}

	if obj.ChatMemberLeft != nil {
		return *obj.ChatMemberLeft
	}

	if obj.ChatMemberMember != nil {
		return *obj.ChatMemberMember
	}

	if obj.ChatMemberOwner != nil {
		return *obj.ChatMemberOwner
	}

	if obj.ChatMemberRestricted != nil {
		return *obj.ChatMemberRestricted
	}

	// all schemas are nil
	return nil
}

type NullableChatMember struct {
	value *ChatMember
	isSet bool
}

func (v NullableChatMember) Get() *ChatMember {
	return v.value
}

func (v *NullableChatMember) Set(val *ChatMember) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMember) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMember(val *ChatMember) *NullableChatMember {
	return &NullableChatMember{value: val, isSet: true}
}

func (v NullableChatMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


