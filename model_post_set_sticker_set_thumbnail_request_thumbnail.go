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

// PostSetStickerSetThumbnailRequestThumbnail - A **.WEBP** or **.PNG** image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a **.TGS** animation with a thumbnail up to 32 kilobytes in size (see [https://core.telegram.org/stickers#animation-requirements](https://core.telegram.org/stickers#animation-requirements) for animated sticker technical requirements), or a **.WEBM** video with the thumbnail up to 32 kilobytes in size; see [https://core.telegram.org/stickers#video-requirements](https://core.telegram.org/stickers#video-requirements) for video sticker technical requirements. Pass a *file\\_id* as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files). Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
type PostSetStickerSetThumbnailRequestThumbnail struct {
	Any *interface{}
	String *string
}

// interface{}AsPostSetStickerSetThumbnailRequestThumbnail is a convenience function that returns interface{} wrapped in PostSetStickerSetThumbnailRequestThumbnail
func AnyAsPostSetStickerSetThumbnailRequestThumbnail(v *interface{}) PostSetStickerSetThumbnailRequestThumbnail {
	return PostSetStickerSetThumbnailRequestThumbnail{
		Any: v,
	}
}

// stringAsPostSetStickerSetThumbnailRequestThumbnail is a convenience function that returns string wrapped in PostSetStickerSetThumbnailRequestThumbnail
func StringAsPostSetStickerSetThumbnailRequestThumbnail(v *string) PostSetStickerSetThumbnailRequestThumbnail {
	return PostSetStickerSetThumbnailRequestThumbnail{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostSetStickerSetThumbnailRequestThumbnail) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PostSetStickerSetThumbnailRequestThumbnail)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostSetStickerSetThumbnailRequestThumbnail)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostSetStickerSetThumbnailRequestThumbnail) MarshalJSON() ([]byte, error) {
	if src.Any != nil {
		return json.Marshal(&src.Any)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostSetStickerSetThumbnailRequestThumbnail) GetActualInstance() (interface{}) {
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
func (obj PostSetStickerSetThumbnailRequestThumbnail) GetActualInstanceValue() (interface{}) {
	if obj.Any != nil {
		return *obj.Any
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePostSetStickerSetThumbnailRequestThumbnail struct {
	value *PostSetStickerSetThumbnailRequestThumbnail
	isSet bool
}

func (v NullablePostSetStickerSetThumbnailRequestThumbnail) Get() *PostSetStickerSetThumbnailRequestThumbnail {
	return v.value
}

func (v *NullablePostSetStickerSetThumbnailRequestThumbnail) Set(val *PostSetStickerSetThumbnailRequestThumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSetStickerSetThumbnailRequestThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSetStickerSetThumbnailRequestThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSetStickerSetThumbnailRequestThumbnail(val *PostSetStickerSetThumbnailRequestThumbnail) *NullablePostSetStickerSetThumbnailRequestThumbnail {
	return &NullablePostSetStickerSetThumbnailRequestThumbnail{value: val, isSet: true}
}

func (v NullablePostSetStickerSetThumbnailRequestThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSetStickerSetThumbnailRequestThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


