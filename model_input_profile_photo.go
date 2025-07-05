/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// InputProfilePhoto - This object describes a profile photo to set. Currently, it can be one of  * [InputProfilePhotoStatic](https://core.telegram.org/bots/api/#inputprofilephotostatic) * [InputProfilePhotoAnimated](https://core.telegram.org/bots/api/#inputprofilephotoanimated)
type InputProfilePhoto struct {
	InputProfilePhotoAnimated *InputProfilePhotoAnimated
	InputProfilePhotoStatic *InputProfilePhotoStatic
}

// InputProfilePhotoAnimatedAsInputProfilePhoto is a convenience function that returns InputProfilePhotoAnimated wrapped in InputProfilePhoto
func InputProfilePhotoAnimatedAsInputProfilePhoto(v *InputProfilePhotoAnimated) InputProfilePhoto {
	return InputProfilePhoto{
		InputProfilePhotoAnimated: v,
	}
}

// InputProfilePhotoStaticAsInputProfilePhoto is a convenience function that returns InputProfilePhotoStatic wrapped in InputProfilePhoto
func InputProfilePhotoStaticAsInputProfilePhoto(v *InputProfilePhotoStatic) InputProfilePhoto {
	return InputProfilePhoto{
		InputProfilePhotoStatic: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputProfilePhoto) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputProfilePhotoAnimated
	err = newStrictDecoder(data).Decode(&dst.InputProfilePhotoAnimated)
	if err == nil {
		jsonInputProfilePhotoAnimated, _ := json.Marshal(dst.InputProfilePhotoAnimated)
		if string(jsonInputProfilePhotoAnimated) == "{}" { // empty struct
			dst.InputProfilePhotoAnimated = nil
		} else {
			if err = validator.Validate(dst.InputProfilePhotoAnimated); err != nil {
				dst.InputProfilePhotoAnimated = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputProfilePhotoAnimated = nil
	}

	// try to unmarshal data into InputProfilePhotoStatic
	err = newStrictDecoder(data).Decode(&dst.InputProfilePhotoStatic)
	if err == nil {
		jsonInputProfilePhotoStatic, _ := json.Marshal(dst.InputProfilePhotoStatic)
		if string(jsonInputProfilePhotoStatic) == "{}" { // empty struct
			dst.InputProfilePhotoStatic = nil
		} else {
			if err = validator.Validate(dst.InputProfilePhotoStatic); err != nil {
				dst.InputProfilePhotoStatic = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputProfilePhotoStatic = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputProfilePhotoAnimated = nil
		dst.InputProfilePhotoStatic = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputProfilePhoto)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputProfilePhoto)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputProfilePhoto) MarshalJSON() ([]byte, error) {
	if src.InputProfilePhotoAnimated != nil {
		return json.Marshal(&src.InputProfilePhotoAnimated)
	}

	if src.InputProfilePhotoStatic != nil {
		return json.Marshal(&src.InputProfilePhotoStatic)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputProfilePhoto) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InputProfilePhotoAnimated != nil {
		return obj.InputProfilePhotoAnimated
	}

	if obj.InputProfilePhotoStatic != nil {
		return obj.InputProfilePhotoStatic
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputProfilePhoto) GetActualInstanceValue() (interface{}) {
	if obj.InputProfilePhotoAnimated != nil {
		return *obj.InputProfilePhotoAnimated
	}

	if obj.InputProfilePhotoStatic != nil {
		return *obj.InputProfilePhotoStatic
	}

	// all schemas are nil
	return nil
}

type NullableInputProfilePhoto struct {
	value *InputProfilePhoto
	isSet bool
}

func (v NullableInputProfilePhoto) Get() *InputProfilePhoto {
	return v.value
}

func (v *NullableInputProfilePhoto) Set(val *InputProfilePhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableInputProfilePhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableInputProfilePhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputProfilePhoto(val *InputProfilePhoto) *NullableInputProfilePhoto {
	return &NullableInputProfilePhoto{value: val, isSet: true}
}

func (v NullableInputProfilePhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputProfilePhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


