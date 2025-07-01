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


// ChatBoostSource This object describes the source of a chat boost. It can be one of  * [ChatBoostSourcePremium](https://core.telegram.org/bots/api/#chatboostsourcepremium) * [ChatBoostSourceGiftCode](https://core.telegram.org/bots/api/#chatboostsourcegiftcode) * [ChatBoostSourceGiveaway](https://core.telegram.org/bots/api/#chatboostsourcegiveaway)
type ChatBoostSource struct {
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway
	ChatBoostSourcePremium *ChatBoostSourcePremium
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ChatBoostSource) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ChatBoostSourceGiftCode
	err = json.Unmarshal(data, &dst.ChatBoostSourceGiftCode);
	if err == nil {
		jsonChatBoostSourceGiftCode, _ := json.Marshal(dst.ChatBoostSourceGiftCode)
		if string(jsonChatBoostSourceGiftCode) == "{}" { // empty struct
			dst.ChatBoostSourceGiftCode = nil
		} else {
			return nil // data stored in dst.ChatBoostSourceGiftCode, return on the first match
		}
	} else {
		dst.ChatBoostSourceGiftCode = nil
	}

	// try to unmarshal JSON data into ChatBoostSourceGiveaway
	err = json.Unmarshal(data, &dst.ChatBoostSourceGiveaway);
	if err == nil {
		jsonChatBoostSourceGiveaway, _ := json.Marshal(dst.ChatBoostSourceGiveaway)
		if string(jsonChatBoostSourceGiveaway) == "{}" { // empty struct
			dst.ChatBoostSourceGiveaway = nil
		} else {
			return nil // data stored in dst.ChatBoostSourceGiveaway, return on the first match
		}
	} else {
		dst.ChatBoostSourceGiveaway = nil
	}

	// try to unmarshal JSON data into ChatBoostSourcePremium
	err = json.Unmarshal(data, &dst.ChatBoostSourcePremium);
	if err == nil {
		jsonChatBoostSourcePremium, _ := json.Marshal(dst.ChatBoostSourcePremium)
		if string(jsonChatBoostSourcePremium) == "{}" { // empty struct
			dst.ChatBoostSourcePremium = nil
		} else {
			return nil // data stored in dst.ChatBoostSourcePremium, return on the first match
		}
	} else {
		dst.ChatBoostSourcePremium = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ChatBoostSource)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatBoostSource) MarshalJSON() ([]byte, error) {
	if src.ChatBoostSourceGiftCode != nil {
		return json.Marshal(&src.ChatBoostSourceGiftCode)
	}

	if src.ChatBoostSourceGiveaway != nil {
		return json.Marshal(&src.ChatBoostSourceGiveaway)
	}

	if src.ChatBoostSourcePremium != nil {
		return json.Marshal(&src.ChatBoostSourcePremium)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableChatBoostSource struct {
	value *ChatBoostSource
	isSet bool
}

func (v NullableChatBoostSource) Get() *ChatBoostSource {
	return v.value
}

func (v *NullableChatBoostSource) Set(val *ChatBoostSource) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBoostSource) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBoostSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBoostSource(val *ChatBoostSource) *NullableChatBoostSource {
	return &NullableChatBoostSource{value: val, isSet: true}
}

func (v NullableChatBoostSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBoostSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


