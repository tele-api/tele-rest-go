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

// InputMedia - This object represents the content of a media message to be sent. It should be one of  * [InputMediaAnimation](https://core.telegram.org/bots/api/#inputmediaanimation) * [InputMediaDocument](https://core.telegram.org/bots/api/#inputmediadocument) * [InputMediaAudio](https://core.telegram.org/bots/api/#inputmediaaudio) * [InputMediaPhoto](https://core.telegram.org/bots/api/#inputmediaphoto) * [InputMediaVideo](https://core.telegram.org/bots/api/#inputmediavideo)
type InputMedia struct {
	InputMediaAnimation *InputMediaAnimation
	InputMediaAudio *InputMediaAudio
	InputMediaDocument *InputMediaDocument
	InputMediaPhoto *InputMediaPhoto
	InputMediaVideo *InputMediaVideo
}

// InputMediaAnimationAsInputMedia is a convenience function that returns InputMediaAnimation wrapped in InputMedia
func InputMediaAnimationAsInputMedia(v *InputMediaAnimation) InputMedia {
	return InputMedia{
		InputMediaAnimation: v,
	}
}

// InputMediaAudioAsInputMedia is a convenience function that returns InputMediaAudio wrapped in InputMedia
func InputMediaAudioAsInputMedia(v *InputMediaAudio) InputMedia {
	return InputMedia{
		InputMediaAudio: v,
	}
}

// InputMediaDocumentAsInputMedia is a convenience function that returns InputMediaDocument wrapped in InputMedia
func InputMediaDocumentAsInputMedia(v *InputMediaDocument) InputMedia {
	return InputMedia{
		InputMediaDocument: v,
	}
}

// InputMediaPhotoAsInputMedia is a convenience function that returns InputMediaPhoto wrapped in InputMedia
func InputMediaPhotoAsInputMedia(v *InputMediaPhoto) InputMedia {
	return InputMedia{
		InputMediaPhoto: v,
	}
}

// InputMediaVideoAsInputMedia is a convenience function that returns InputMediaVideo wrapped in InputMedia
func InputMediaVideoAsInputMedia(v *InputMediaVideo) InputMedia {
	return InputMedia{
		InputMediaVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputMedia) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputMediaAnimation
	err = newStrictDecoder(data).Decode(&dst.InputMediaAnimation)
	if err == nil {
		jsonInputMediaAnimation, _ := json.Marshal(dst.InputMediaAnimation)
		if string(jsonInputMediaAnimation) == "{}" { // empty struct
			dst.InputMediaAnimation = nil
		} else {
			if err = validator.Validate(dst.InputMediaAnimation); err != nil {
				dst.InputMediaAnimation = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputMediaAnimation = nil
	}

	// try to unmarshal data into InputMediaAudio
	err = newStrictDecoder(data).Decode(&dst.InputMediaAudio)
	if err == nil {
		jsonInputMediaAudio, _ := json.Marshal(dst.InputMediaAudio)
		if string(jsonInputMediaAudio) == "{}" { // empty struct
			dst.InputMediaAudio = nil
		} else {
			if err = validator.Validate(dst.InputMediaAudio); err != nil {
				dst.InputMediaAudio = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputMediaAudio = nil
	}

	// try to unmarshal data into InputMediaDocument
	err = newStrictDecoder(data).Decode(&dst.InputMediaDocument)
	if err == nil {
		jsonInputMediaDocument, _ := json.Marshal(dst.InputMediaDocument)
		if string(jsonInputMediaDocument) == "{}" { // empty struct
			dst.InputMediaDocument = nil
		} else {
			if err = validator.Validate(dst.InputMediaDocument); err != nil {
				dst.InputMediaDocument = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputMediaDocument = nil
	}

	// try to unmarshal data into InputMediaPhoto
	err = newStrictDecoder(data).Decode(&dst.InputMediaPhoto)
	if err == nil {
		jsonInputMediaPhoto, _ := json.Marshal(dst.InputMediaPhoto)
		if string(jsonInputMediaPhoto) == "{}" { // empty struct
			dst.InputMediaPhoto = nil
		} else {
			if err = validator.Validate(dst.InputMediaPhoto); err != nil {
				dst.InputMediaPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputMediaPhoto = nil
	}

	// try to unmarshal data into InputMediaVideo
	err = newStrictDecoder(data).Decode(&dst.InputMediaVideo)
	if err == nil {
		jsonInputMediaVideo, _ := json.Marshal(dst.InputMediaVideo)
		if string(jsonInputMediaVideo) == "{}" { // empty struct
			dst.InputMediaVideo = nil
		} else {
			if err = validator.Validate(dst.InputMediaVideo); err != nil {
				dst.InputMediaVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputMediaVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputMediaAnimation = nil
		dst.InputMediaAudio = nil
		dst.InputMediaDocument = nil
		dst.InputMediaPhoto = nil
		dst.InputMediaVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputMedia)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputMedia)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputMedia) MarshalJSON() ([]byte, error) {
	if src.InputMediaAnimation != nil {
		return json.Marshal(&src.InputMediaAnimation)
	}

	if src.InputMediaAudio != nil {
		return json.Marshal(&src.InputMediaAudio)
	}

	if src.InputMediaDocument != nil {
		return json.Marshal(&src.InputMediaDocument)
	}

	if src.InputMediaPhoto != nil {
		return json.Marshal(&src.InputMediaPhoto)
	}

	if src.InputMediaVideo != nil {
		return json.Marshal(&src.InputMediaVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputMedia) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InputMediaAnimation != nil {
		return obj.InputMediaAnimation
	}

	if obj.InputMediaAudio != nil {
		return obj.InputMediaAudio
	}

	if obj.InputMediaDocument != nil {
		return obj.InputMediaDocument
	}

	if obj.InputMediaPhoto != nil {
		return obj.InputMediaPhoto
	}

	if obj.InputMediaVideo != nil {
		return obj.InputMediaVideo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputMedia) GetActualInstanceValue() (interface{}) {
	if obj.InputMediaAnimation != nil {
		return *obj.InputMediaAnimation
	}

	if obj.InputMediaAudio != nil {
		return *obj.InputMediaAudio
	}

	if obj.InputMediaDocument != nil {
		return *obj.InputMediaDocument
	}

	if obj.InputMediaPhoto != nil {
		return *obj.InputMediaPhoto
	}

	if obj.InputMediaVideo != nil {
		return *obj.InputMediaVideo
	}

	// all schemas are nil
	return nil
}

type NullableInputMedia struct {
	value *InputMedia
	isSet bool
}

func (v NullableInputMedia) Get() *InputMedia {
	return v.value
}

func (v *NullableInputMedia) Set(val *InputMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMedia(val *InputMedia) *NullableInputMedia {
	return &NullableInputMedia{value: val, isSet: true}
}

func (v NullableInputMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


