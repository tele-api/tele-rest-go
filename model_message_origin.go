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


// MessageOrigin This object describes the origin of a message. It can be one of  * [MessageOriginUser](https://core.telegram.org/bots/api/#messageoriginuser) * [MessageOriginHiddenUser](https://core.telegram.org/bots/api/#messageoriginhiddenuser) * [MessageOriginChat](https://core.telegram.org/bots/api/#messageoriginchat) * [MessageOriginChannel](https://core.telegram.org/bots/api/#messageoriginchannel)
type MessageOrigin struct {
	MessageOriginChannel *MessageOriginChannel
	MessageOriginChat *MessageOriginChat
	MessageOriginHiddenUser *MessageOriginHiddenUser
	MessageOriginUser *MessageOriginUser
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MessageOrigin) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MessageOriginChannel
	err = json.Unmarshal(data, &dst.MessageOriginChannel);
	if err == nil {
		jsonMessageOriginChannel, _ := json.Marshal(dst.MessageOriginChannel)
		if string(jsonMessageOriginChannel) == "{}" { // empty struct
			dst.MessageOriginChannel = nil
		} else {
			return nil // data stored in dst.MessageOriginChannel, return on the first match
		}
	} else {
		dst.MessageOriginChannel = nil
	}

	// try to unmarshal JSON data into MessageOriginChat
	err = json.Unmarshal(data, &dst.MessageOriginChat);
	if err == nil {
		jsonMessageOriginChat, _ := json.Marshal(dst.MessageOriginChat)
		if string(jsonMessageOriginChat) == "{}" { // empty struct
			dst.MessageOriginChat = nil
		} else {
			return nil // data stored in dst.MessageOriginChat, return on the first match
		}
	} else {
		dst.MessageOriginChat = nil
	}

	// try to unmarshal JSON data into MessageOriginHiddenUser
	err = json.Unmarshal(data, &dst.MessageOriginHiddenUser);
	if err == nil {
		jsonMessageOriginHiddenUser, _ := json.Marshal(dst.MessageOriginHiddenUser)
		if string(jsonMessageOriginHiddenUser) == "{}" { // empty struct
			dst.MessageOriginHiddenUser = nil
		} else {
			return nil // data stored in dst.MessageOriginHiddenUser, return on the first match
		}
	} else {
		dst.MessageOriginHiddenUser = nil
	}

	// try to unmarshal JSON data into MessageOriginUser
	err = json.Unmarshal(data, &dst.MessageOriginUser);
	if err == nil {
		jsonMessageOriginUser, _ := json.Marshal(dst.MessageOriginUser)
		if string(jsonMessageOriginUser) == "{}" { // empty struct
			dst.MessageOriginUser = nil
		} else {
			return nil // data stored in dst.MessageOriginUser, return on the first match
		}
	} else {
		dst.MessageOriginUser = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MessageOrigin)")
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

	return nil, nil // no data in anyOf schemas
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


