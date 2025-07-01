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


// SendMessagePostRequestReplyMarkup Additional interface options. A JSON-serialized object for an [inline keyboard](https://core.telegram.org/bots/features#inline-keyboards), [custom reply keyboard](https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a reply from the user
type SendMessagePostRequestReplyMarkup struct {
	ForceReply *ForceReply
	InlineKeyboardMarkup *InlineKeyboardMarkup
	ReplyKeyboardMarkup *ReplyKeyboardMarkup
	ReplyKeyboardRemove *ReplyKeyboardRemove
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SendMessagePostRequestReplyMarkup) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ForceReply
	err = json.Unmarshal(data, &dst.ForceReply);
	if err == nil {
		jsonForceReply, _ := json.Marshal(dst.ForceReply)
		if string(jsonForceReply) == "{}" { // empty struct
			dst.ForceReply = nil
		} else {
			return nil // data stored in dst.ForceReply, return on the first match
		}
	} else {
		dst.ForceReply = nil
	}

	// try to unmarshal JSON data into InlineKeyboardMarkup
	err = json.Unmarshal(data, &dst.InlineKeyboardMarkup);
	if err == nil {
		jsonInlineKeyboardMarkup, _ := json.Marshal(dst.InlineKeyboardMarkup)
		if string(jsonInlineKeyboardMarkup) == "{}" { // empty struct
			dst.InlineKeyboardMarkup = nil
		} else {
			return nil // data stored in dst.InlineKeyboardMarkup, return on the first match
		}
	} else {
		dst.InlineKeyboardMarkup = nil
	}

	// try to unmarshal JSON data into ReplyKeyboardMarkup
	err = json.Unmarshal(data, &dst.ReplyKeyboardMarkup);
	if err == nil {
		jsonReplyKeyboardMarkup, _ := json.Marshal(dst.ReplyKeyboardMarkup)
		if string(jsonReplyKeyboardMarkup) == "{}" { // empty struct
			dst.ReplyKeyboardMarkup = nil
		} else {
			return nil // data stored in dst.ReplyKeyboardMarkup, return on the first match
		}
	} else {
		dst.ReplyKeyboardMarkup = nil
	}

	// try to unmarshal JSON data into ReplyKeyboardRemove
	err = json.Unmarshal(data, &dst.ReplyKeyboardRemove);
	if err == nil {
		jsonReplyKeyboardRemove, _ := json.Marshal(dst.ReplyKeyboardRemove)
		if string(jsonReplyKeyboardRemove) == "{}" { // empty struct
			dst.ReplyKeyboardRemove = nil
		} else {
			return nil // data stored in dst.ReplyKeyboardRemove, return on the first match
		}
	} else {
		dst.ReplyKeyboardRemove = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SendMessagePostRequestReplyMarkup)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SendMessagePostRequestReplyMarkup) MarshalJSON() ([]byte, error) {
	if src.ForceReply != nil {
		return json.Marshal(&src.ForceReply)
	}

	if src.InlineKeyboardMarkup != nil {
		return json.Marshal(&src.InlineKeyboardMarkup)
	}

	if src.ReplyKeyboardMarkup != nil {
		return json.Marshal(&src.ReplyKeyboardMarkup)
	}

	if src.ReplyKeyboardRemove != nil {
		return json.Marshal(&src.ReplyKeyboardRemove)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableSendMessagePostRequestReplyMarkup struct {
	value *SendMessagePostRequestReplyMarkup
	isSet bool
}

func (v NullableSendMessagePostRequestReplyMarkup) Get() *SendMessagePostRequestReplyMarkup {
	return v.value
}

func (v *NullableSendMessagePostRequestReplyMarkup) Set(val *SendMessagePostRequestReplyMarkup) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessagePostRequestReplyMarkup) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessagePostRequestReplyMarkup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessagePostRequestReplyMarkup(val *SendMessagePostRequestReplyMarkup) *NullableSendMessagePostRequestReplyMarkup {
	return &NullableSendMessagePostRequestReplyMarkup{value: val, isSet: true}
}

func (v NullableSendMessagePostRequestReplyMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessagePostRequestReplyMarkup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


