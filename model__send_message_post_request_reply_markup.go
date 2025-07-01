/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// SendMessagePostRequestReplyMarkup - Additional interface options. A JSON-serialized object for an [inline keyboard](https://core.telegram.org/bots/features#inline-keyboards), [custom reply keyboard](https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a reply from the user
type SendMessagePostRequestReplyMarkup struct {
	ForceReply *ForceReply
	InlineKeyboardMarkup *InlineKeyboardMarkup
	ReplyKeyboardMarkup *ReplyKeyboardMarkup
	ReplyKeyboardRemove *ReplyKeyboardRemove
}

// ForceReplyAsSendMessagePostRequestReplyMarkup is a convenience function that returns ForceReply wrapped in SendMessagePostRequestReplyMarkup
func ForceReplyAsSendMessagePostRequestReplyMarkup(v *ForceReply) SendMessagePostRequestReplyMarkup {
	return SendMessagePostRequestReplyMarkup{
		ForceReply: v,
	}
}

// InlineKeyboardMarkupAsSendMessagePostRequestReplyMarkup is a convenience function that returns InlineKeyboardMarkup wrapped in SendMessagePostRequestReplyMarkup
func InlineKeyboardMarkupAsSendMessagePostRequestReplyMarkup(v *InlineKeyboardMarkup) SendMessagePostRequestReplyMarkup {
	return SendMessagePostRequestReplyMarkup{
		InlineKeyboardMarkup: v,
	}
}

// ReplyKeyboardMarkupAsSendMessagePostRequestReplyMarkup is a convenience function that returns ReplyKeyboardMarkup wrapped in SendMessagePostRequestReplyMarkup
func ReplyKeyboardMarkupAsSendMessagePostRequestReplyMarkup(v *ReplyKeyboardMarkup) SendMessagePostRequestReplyMarkup {
	return SendMessagePostRequestReplyMarkup{
		ReplyKeyboardMarkup: v,
	}
}

// ReplyKeyboardRemoveAsSendMessagePostRequestReplyMarkup is a convenience function that returns ReplyKeyboardRemove wrapped in SendMessagePostRequestReplyMarkup
func ReplyKeyboardRemoveAsSendMessagePostRequestReplyMarkup(v *ReplyKeyboardRemove) SendMessagePostRequestReplyMarkup {
	return SendMessagePostRequestReplyMarkup{
		ReplyKeyboardRemove: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SendMessagePostRequestReplyMarkup) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ForceReply
	err = newStrictDecoder(data).Decode(&dst.ForceReply)
	if err == nil {
		jsonForceReply, _ := json.Marshal(dst.ForceReply)
		if string(jsonForceReply) == "{}" { // empty struct
			dst.ForceReply = nil
		} else {
			if err = validator.Validate(dst.ForceReply); err != nil {
				dst.ForceReply = nil
			} else {
				match++
			}
		}
	} else {
		dst.ForceReply = nil
	}

	// try to unmarshal data into InlineKeyboardMarkup
	err = newStrictDecoder(data).Decode(&dst.InlineKeyboardMarkup)
	if err == nil {
		jsonInlineKeyboardMarkup, _ := json.Marshal(dst.InlineKeyboardMarkup)
		if string(jsonInlineKeyboardMarkup) == "{}" { // empty struct
			dst.InlineKeyboardMarkup = nil
		} else {
			if err = validator.Validate(dst.InlineKeyboardMarkup); err != nil {
				dst.InlineKeyboardMarkup = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineKeyboardMarkup = nil
	}

	// try to unmarshal data into ReplyKeyboardMarkup
	err = newStrictDecoder(data).Decode(&dst.ReplyKeyboardMarkup)
	if err == nil {
		jsonReplyKeyboardMarkup, _ := json.Marshal(dst.ReplyKeyboardMarkup)
		if string(jsonReplyKeyboardMarkup) == "{}" { // empty struct
			dst.ReplyKeyboardMarkup = nil
		} else {
			if err = validator.Validate(dst.ReplyKeyboardMarkup); err != nil {
				dst.ReplyKeyboardMarkup = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReplyKeyboardMarkup = nil
	}

	// try to unmarshal data into ReplyKeyboardRemove
	err = newStrictDecoder(data).Decode(&dst.ReplyKeyboardRemove)
	if err == nil {
		jsonReplyKeyboardRemove, _ := json.Marshal(dst.ReplyKeyboardRemove)
		if string(jsonReplyKeyboardRemove) == "{}" { // empty struct
			dst.ReplyKeyboardRemove = nil
		} else {
			if err = validator.Validate(dst.ReplyKeyboardRemove); err != nil {
				dst.ReplyKeyboardRemove = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReplyKeyboardRemove = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ForceReply = nil
		dst.InlineKeyboardMarkup = nil
		dst.ReplyKeyboardMarkup = nil
		dst.ReplyKeyboardRemove = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SendMessagePostRequestReplyMarkup)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SendMessagePostRequestReplyMarkup)")
	}
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

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SendMessagePostRequestReplyMarkup) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ForceReply != nil {
		return obj.ForceReply
	}

	if obj.InlineKeyboardMarkup != nil {
		return obj.InlineKeyboardMarkup
	}

	if obj.ReplyKeyboardMarkup != nil {
		return obj.ReplyKeyboardMarkup
	}

	if obj.ReplyKeyboardRemove != nil {
		return obj.ReplyKeyboardRemove
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SendMessagePostRequestReplyMarkup) GetActualInstanceValue() (interface{}) {
	if obj.ForceReply != nil {
		return *obj.ForceReply
	}

	if obj.InlineKeyboardMarkup != nil {
		return *obj.InlineKeyboardMarkup
	}

	if obj.ReplyKeyboardMarkup != nil {
		return *obj.ReplyKeyboardMarkup
	}

	if obj.ReplyKeyboardRemove != nil {
		return *obj.ReplyKeyboardRemove
	}

	// all schemas are nil
	return nil
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


