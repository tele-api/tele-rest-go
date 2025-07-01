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


// BotCommandScope This object represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:  * [BotCommandScopeDefault](https://core.telegram.org/bots/api/#botcommandscopedefault) * [BotCommandScopeAllPrivateChats](https://core.telegram.org/bots/api/#botcommandscopeallprivatechats) * [BotCommandScopeAllGroupChats](https://core.telegram.org/bots/api/#botcommandscopeallgroupchats) * [BotCommandScopeAllChatAdministrators](https://core.telegram.org/bots/api/#botcommandscopeallchatadministrators) * [BotCommandScopeChat](https://core.telegram.org/bots/api/#botcommandscopechat) * [BotCommandScopeChatAdministrators](https://core.telegram.org/bots/api/#botcommandscopechatadministrators) * [BotCommandScopeChatMember](https://core.telegram.org/bots/api/#botcommandscopechatmember)
type BotCommandScope struct {
	BotCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators
	BotCommandScopeAllGroupChats *BotCommandScopeAllGroupChats
	BotCommandScopeAllPrivateChats *BotCommandScopeAllPrivateChats
	BotCommandScopeChat *BotCommandScopeChat
	BotCommandScopeChatAdministrators *BotCommandScopeChatAdministrators
	BotCommandScopeChatMember *BotCommandScopeChatMember
	BotCommandScopeDefault *BotCommandScopeDefault
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BotCommandScope) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BotCommandScopeAllChatAdministrators
	err = json.Unmarshal(data, &dst.BotCommandScopeAllChatAdministrators);
	if err == nil {
		jsonBotCommandScopeAllChatAdministrators, _ := json.Marshal(dst.BotCommandScopeAllChatAdministrators)
		if string(jsonBotCommandScopeAllChatAdministrators) == "{}" { // empty struct
			dst.BotCommandScopeAllChatAdministrators = nil
		} else {
			return nil // data stored in dst.BotCommandScopeAllChatAdministrators, return on the first match
		}
	} else {
		dst.BotCommandScopeAllChatAdministrators = nil
	}

	// try to unmarshal JSON data into BotCommandScopeAllGroupChats
	err = json.Unmarshal(data, &dst.BotCommandScopeAllGroupChats);
	if err == nil {
		jsonBotCommandScopeAllGroupChats, _ := json.Marshal(dst.BotCommandScopeAllGroupChats)
		if string(jsonBotCommandScopeAllGroupChats) == "{}" { // empty struct
			dst.BotCommandScopeAllGroupChats = nil
		} else {
			return nil // data stored in dst.BotCommandScopeAllGroupChats, return on the first match
		}
	} else {
		dst.BotCommandScopeAllGroupChats = nil
	}

	// try to unmarshal JSON data into BotCommandScopeAllPrivateChats
	err = json.Unmarshal(data, &dst.BotCommandScopeAllPrivateChats);
	if err == nil {
		jsonBotCommandScopeAllPrivateChats, _ := json.Marshal(dst.BotCommandScopeAllPrivateChats)
		if string(jsonBotCommandScopeAllPrivateChats) == "{}" { // empty struct
			dst.BotCommandScopeAllPrivateChats = nil
		} else {
			return nil // data stored in dst.BotCommandScopeAllPrivateChats, return on the first match
		}
	} else {
		dst.BotCommandScopeAllPrivateChats = nil
	}

	// try to unmarshal JSON data into BotCommandScopeChat
	err = json.Unmarshal(data, &dst.BotCommandScopeChat);
	if err == nil {
		jsonBotCommandScopeChat, _ := json.Marshal(dst.BotCommandScopeChat)
		if string(jsonBotCommandScopeChat) == "{}" { // empty struct
			dst.BotCommandScopeChat = nil
		} else {
			return nil // data stored in dst.BotCommandScopeChat, return on the first match
		}
	} else {
		dst.BotCommandScopeChat = nil
	}

	// try to unmarshal JSON data into BotCommandScopeChatAdministrators
	err = json.Unmarshal(data, &dst.BotCommandScopeChatAdministrators);
	if err == nil {
		jsonBotCommandScopeChatAdministrators, _ := json.Marshal(dst.BotCommandScopeChatAdministrators)
		if string(jsonBotCommandScopeChatAdministrators) == "{}" { // empty struct
			dst.BotCommandScopeChatAdministrators = nil
		} else {
			return nil // data stored in dst.BotCommandScopeChatAdministrators, return on the first match
		}
	} else {
		dst.BotCommandScopeChatAdministrators = nil
	}

	// try to unmarshal JSON data into BotCommandScopeChatMember
	err = json.Unmarshal(data, &dst.BotCommandScopeChatMember);
	if err == nil {
		jsonBotCommandScopeChatMember, _ := json.Marshal(dst.BotCommandScopeChatMember)
		if string(jsonBotCommandScopeChatMember) == "{}" { // empty struct
			dst.BotCommandScopeChatMember = nil
		} else {
			return nil // data stored in dst.BotCommandScopeChatMember, return on the first match
		}
	} else {
		dst.BotCommandScopeChatMember = nil
	}

	// try to unmarshal JSON data into BotCommandScopeDefault
	err = json.Unmarshal(data, &dst.BotCommandScopeDefault);
	if err == nil {
		jsonBotCommandScopeDefault, _ := json.Marshal(dst.BotCommandScopeDefault)
		if string(jsonBotCommandScopeDefault) == "{}" { // empty struct
			dst.BotCommandScopeDefault = nil
		} else {
			return nil // data stored in dst.BotCommandScopeDefault, return on the first match
		}
	} else {
		dst.BotCommandScopeDefault = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BotCommandScope)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BotCommandScope) MarshalJSON() ([]byte, error) {
	if src.BotCommandScopeAllChatAdministrators != nil {
		return json.Marshal(&src.BotCommandScopeAllChatAdministrators)
	}

	if src.BotCommandScopeAllGroupChats != nil {
		return json.Marshal(&src.BotCommandScopeAllGroupChats)
	}

	if src.BotCommandScopeAllPrivateChats != nil {
		return json.Marshal(&src.BotCommandScopeAllPrivateChats)
	}

	if src.BotCommandScopeChat != nil {
		return json.Marshal(&src.BotCommandScopeChat)
	}

	if src.BotCommandScopeChatAdministrators != nil {
		return json.Marshal(&src.BotCommandScopeChatAdministrators)
	}

	if src.BotCommandScopeChatMember != nil {
		return json.Marshal(&src.BotCommandScopeChatMember)
	}

	if src.BotCommandScopeDefault != nil {
		return json.Marshal(&src.BotCommandScopeDefault)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableBotCommandScope struct {
	value *BotCommandScope
	isSet bool
}

func (v NullableBotCommandScope) Get() *BotCommandScope {
	return v.value
}

func (v *NullableBotCommandScope) Set(val *BotCommandScope) {
	v.value = val
	v.isSet = true
}

func (v NullableBotCommandScope) IsSet() bool {
	return v.isSet
}

func (v *NullableBotCommandScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotCommandScope(val *BotCommandScope) *NullableBotCommandScope {
	return &NullableBotCommandScope{value: val, isSet: true}
}

func (v NullableBotCommandScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotCommandScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


