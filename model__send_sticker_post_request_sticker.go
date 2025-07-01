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

// SendStickerPostRequestSticker - Sticker to send. Pass a file\\_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files). Video and animated stickers can't be sent via an HTTP URL.
type SendStickerPostRequestSticker struct {
	Any *interface{}
	String *string
}

// interface{}AsSendStickerPostRequestSticker is a convenience function that returns interface{} wrapped in SendStickerPostRequestSticker
func AnyAsSendStickerPostRequestSticker(v *interface{}) SendStickerPostRequestSticker {
	return SendStickerPostRequestSticker{
		Any: v,
	}
}

// stringAsSendStickerPostRequestSticker is a convenience function that returns string wrapped in SendStickerPostRequestSticker
func StringAsSendStickerPostRequestSticker(v *string) SendStickerPostRequestSticker {
	return SendStickerPostRequestSticker{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SendStickerPostRequestSticker) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Any
	err = newStrictDecoder(data).Decode(&dst.Any)
	if err == nil {
		jsonAny, _ := json.Marshal(dst.Any)
		if string(jsonAny) == "{}" { // empty struct
			dst.Any = nil
		} else {
			if err = validator.Validate(dst.Any); err != nil {
				dst.Any = nil
			} else {
				match++
			}
		}
	} else {
		dst.Any = nil
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
		dst.Any = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SendStickerPostRequestSticker)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SendStickerPostRequestSticker)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SendStickerPostRequestSticker) MarshalJSON() ([]byte, error) {
	if src.Any != nil {
		return json.Marshal(&src.Any)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SendStickerPostRequestSticker) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Any != nil {
		return obj.Any
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SendStickerPostRequestSticker) GetActualInstanceValue() (interface{}) {
	if obj.Any != nil {
		return *obj.Any
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSendStickerPostRequestSticker struct {
	value *SendStickerPostRequestSticker
	isSet bool
}

func (v NullableSendStickerPostRequestSticker) Get() *SendStickerPostRequestSticker {
	return v.value
}

func (v *NullableSendStickerPostRequestSticker) Set(val *SendStickerPostRequestSticker) {
	v.value = val
	v.isSet = true
}

func (v NullableSendStickerPostRequestSticker) IsSet() bool {
	return v.isSet
}

func (v *NullableSendStickerPostRequestSticker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendStickerPostRequestSticker(val *SendStickerPostRequestSticker) *NullableSendStickerPostRequestSticker {
	return &NullableSendStickerPostRequestSticker{value: val, isSet: true}
}

func (v NullableSendStickerPostRequestSticker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendStickerPostRequestSticker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


