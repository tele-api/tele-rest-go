/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// InputStoryContent - This object describes the content of a story to post. Currently, it can be one of  * [InputStoryContentPhoto](https://core.telegram.org/bots/api/#inputstorycontentphoto) * [InputStoryContentVideo](https://core.telegram.org/bots/api/#inputstorycontentvideo)
type InputStoryContent struct {
	InputStoryContentPhoto *InputStoryContentPhoto
	InputStoryContentVideo *InputStoryContentVideo
}

// InputStoryContentPhotoAsInputStoryContent is a convenience function that returns InputStoryContentPhoto wrapped in InputStoryContent
func InputStoryContentPhotoAsInputStoryContent(v *InputStoryContentPhoto) InputStoryContent {
	return InputStoryContent{
		InputStoryContentPhoto: v,
	}
}

// InputStoryContentVideoAsInputStoryContent is a convenience function that returns InputStoryContentVideo wrapped in InputStoryContent
func InputStoryContentVideoAsInputStoryContent(v *InputStoryContentVideo) InputStoryContent {
	return InputStoryContent{
		InputStoryContentVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputStoryContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputStoryContentPhoto
	err = newStrictDecoder(data).Decode(&dst.InputStoryContentPhoto)
	if err == nil {
		jsonInputStoryContentPhoto, _ := json.Marshal(dst.InputStoryContentPhoto)
		if string(jsonInputStoryContentPhoto) == "{}" { // empty struct
			dst.InputStoryContentPhoto = nil
		} else {
			if err = validator.Validate(dst.InputStoryContentPhoto); err != nil {
				dst.InputStoryContentPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputStoryContentPhoto = nil
	}

	// try to unmarshal data into InputStoryContentVideo
	err = newStrictDecoder(data).Decode(&dst.InputStoryContentVideo)
	if err == nil {
		jsonInputStoryContentVideo, _ := json.Marshal(dst.InputStoryContentVideo)
		if string(jsonInputStoryContentVideo) == "{}" { // empty struct
			dst.InputStoryContentVideo = nil
		} else {
			if err = validator.Validate(dst.InputStoryContentVideo); err != nil {
				dst.InputStoryContentVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputStoryContentVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputStoryContentPhoto = nil
		dst.InputStoryContentVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputStoryContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputStoryContent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputStoryContent) MarshalJSON() ([]byte, error) {
	if src.InputStoryContentPhoto != nil {
		return json.Marshal(&src.InputStoryContentPhoto)
	}

	if src.InputStoryContentVideo != nil {
		return json.Marshal(&src.InputStoryContentVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputStoryContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InputStoryContentPhoto != nil {
		return obj.InputStoryContentPhoto
	}

	if obj.InputStoryContentVideo != nil {
		return obj.InputStoryContentVideo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputStoryContent) GetActualInstanceValue() (interface{}) {
	if obj.InputStoryContentPhoto != nil {
		return *obj.InputStoryContentPhoto
	}

	if obj.InputStoryContentVideo != nil {
		return *obj.InputStoryContentVideo
	}

	// all schemas are nil
	return nil
}

type NullableInputStoryContent struct {
	value *InputStoryContent
	isSet bool
}

func (v NullableInputStoryContent) Get() *InputStoryContent {
	return v.value
}

func (v *NullableInputStoryContent) Set(val *InputStoryContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputStoryContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputStoryContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputStoryContent(val *InputStoryContent) *NullableInputStoryContent {
	return &NullableInputStoryContent{value: val, isSet: true}
}

func (v NullableInputStoryContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputStoryContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


