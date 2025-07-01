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

// InputPaidMedia - This object describes the paid media to be sent. Currently, it can be one of  * [InputPaidMediaPhoto](https://core.telegram.org/bots/api/#inputpaidmediaphoto) * [InputPaidMediaVideo](https://core.telegram.org/bots/api/#inputpaidmediavideo)
type InputPaidMedia struct {
	InputPaidMediaPhoto *InputPaidMediaPhoto
	InputPaidMediaVideo *InputPaidMediaVideo
}

// InputPaidMediaPhotoAsInputPaidMedia is a convenience function that returns InputPaidMediaPhoto wrapped in InputPaidMedia
func InputPaidMediaPhotoAsInputPaidMedia(v *InputPaidMediaPhoto) InputPaidMedia {
	return InputPaidMedia{
		InputPaidMediaPhoto: v,
	}
}

// InputPaidMediaVideoAsInputPaidMedia is a convenience function that returns InputPaidMediaVideo wrapped in InputPaidMedia
func InputPaidMediaVideoAsInputPaidMedia(v *InputPaidMediaVideo) InputPaidMedia {
	return InputPaidMedia{
		InputPaidMediaVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputPaidMedia) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputPaidMediaPhoto
	err = newStrictDecoder(data).Decode(&dst.InputPaidMediaPhoto)
	if err == nil {
		jsonInputPaidMediaPhoto, _ := json.Marshal(dst.InputPaidMediaPhoto)
		if string(jsonInputPaidMediaPhoto) == "{}" { // empty struct
			dst.InputPaidMediaPhoto = nil
		} else {
			if err = validator.Validate(dst.InputPaidMediaPhoto); err != nil {
				dst.InputPaidMediaPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputPaidMediaPhoto = nil
	}

	// try to unmarshal data into InputPaidMediaVideo
	err = newStrictDecoder(data).Decode(&dst.InputPaidMediaVideo)
	if err == nil {
		jsonInputPaidMediaVideo, _ := json.Marshal(dst.InputPaidMediaVideo)
		if string(jsonInputPaidMediaVideo) == "{}" { // empty struct
			dst.InputPaidMediaVideo = nil
		} else {
			if err = validator.Validate(dst.InputPaidMediaVideo); err != nil {
				dst.InputPaidMediaVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputPaidMediaVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputPaidMediaPhoto = nil
		dst.InputPaidMediaVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputPaidMedia)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputPaidMedia)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputPaidMedia) MarshalJSON() ([]byte, error) {
	if src.InputPaidMediaPhoto != nil {
		return json.Marshal(&src.InputPaidMediaPhoto)
	}

	if src.InputPaidMediaVideo != nil {
		return json.Marshal(&src.InputPaidMediaVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputPaidMedia) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InputPaidMediaPhoto != nil {
		return obj.InputPaidMediaPhoto
	}

	if obj.InputPaidMediaVideo != nil {
		return obj.InputPaidMediaVideo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputPaidMedia) GetActualInstanceValue() (interface{}) {
	if obj.InputPaidMediaPhoto != nil {
		return *obj.InputPaidMediaPhoto
	}

	if obj.InputPaidMediaVideo != nil {
		return *obj.InputPaidMediaVideo
	}

	// all schemas are nil
	return nil
}

type NullableInputPaidMedia struct {
	value *InputPaidMedia
	isSet bool
}

func (v NullableInputPaidMedia) Get() *InputPaidMedia {
	return v.value
}

func (v *NullableInputPaidMedia) Set(val *InputPaidMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPaidMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPaidMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPaidMedia(val *InputPaidMedia) *NullableInputPaidMedia {
	return &NullableInputPaidMedia{value: val, isSet: true}
}

func (v NullableInputPaidMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPaidMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


