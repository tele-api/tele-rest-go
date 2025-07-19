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

// BackgroundFill - This object describes the way a background is filled based on the selected colors. Currently, it can be one of  * [BackgroundFillSolid](https://core.telegram.org/bots/api/#backgroundfillsolid) * [BackgroundFillGradient](https://core.telegram.org/bots/api/#backgroundfillgradient) * [BackgroundFillFreeformGradient](https://core.telegram.org/bots/api/#backgroundfillfreeformgradient)
type BackgroundFill struct {
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient
	BackgroundFillGradient *BackgroundFillGradient
	BackgroundFillSolid *BackgroundFillSolid
}

// BackgroundFillFreeformGradientAsBackgroundFill is a convenience function that returns BackgroundFillFreeformGradient wrapped in BackgroundFill
func BackgroundFillFreeformGradientAsBackgroundFill(v *BackgroundFillFreeformGradient) BackgroundFill {
	return BackgroundFill{
		BackgroundFillFreeformGradient: v,
	}
}

// BackgroundFillGradientAsBackgroundFill is a convenience function that returns BackgroundFillGradient wrapped in BackgroundFill
func BackgroundFillGradientAsBackgroundFill(v *BackgroundFillGradient) BackgroundFill {
	return BackgroundFill{
		BackgroundFillGradient: v,
	}
}

// BackgroundFillSolidAsBackgroundFill is a convenience function that returns BackgroundFillSolid wrapped in BackgroundFill
func BackgroundFillSolidAsBackgroundFill(v *BackgroundFillSolid) BackgroundFill {
	return BackgroundFill{
		BackgroundFillSolid: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BackgroundFill) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackgroundFillFreeformGradient
	err = newStrictDecoder(data).Decode(&dst.BackgroundFillFreeformGradient)
	if err == nil {
		jsonBackgroundFillFreeformGradient, _ := json.Marshal(dst.BackgroundFillFreeformGradient)
		if string(jsonBackgroundFillFreeformGradient) == "{}" { // empty struct
			dst.BackgroundFillFreeformGradient = nil
		} else {
			if err = validator.Validate(dst.BackgroundFillFreeformGradient); err != nil {
				dst.BackgroundFillFreeformGradient = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundFillFreeformGradient = nil
	}

	// try to unmarshal data into BackgroundFillGradient
	err = newStrictDecoder(data).Decode(&dst.BackgroundFillGradient)
	if err == nil {
		jsonBackgroundFillGradient, _ := json.Marshal(dst.BackgroundFillGradient)
		if string(jsonBackgroundFillGradient) == "{}" { // empty struct
			dst.BackgroundFillGradient = nil
		} else {
			if err = validator.Validate(dst.BackgroundFillGradient); err != nil {
				dst.BackgroundFillGradient = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundFillGradient = nil
	}

	// try to unmarshal data into BackgroundFillSolid
	err = newStrictDecoder(data).Decode(&dst.BackgroundFillSolid)
	if err == nil {
		jsonBackgroundFillSolid, _ := json.Marshal(dst.BackgroundFillSolid)
		if string(jsonBackgroundFillSolid) == "{}" { // empty struct
			dst.BackgroundFillSolid = nil
		} else {
			if err = validator.Validate(dst.BackgroundFillSolid); err != nil {
				dst.BackgroundFillSolid = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundFillSolid = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackgroundFillFreeformGradient = nil
		dst.BackgroundFillGradient = nil
		dst.BackgroundFillSolid = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BackgroundFill)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BackgroundFill)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BackgroundFill) MarshalJSON() ([]byte, error) {
	if src.BackgroundFillFreeformGradient != nil {
		return json.Marshal(&src.BackgroundFillFreeformGradient)
	}

	if src.BackgroundFillGradient != nil {
		return json.Marshal(&src.BackgroundFillGradient)
	}

	if src.BackgroundFillSolid != nil {
		return json.Marshal(&src.BackgroundFillSolid)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BackgroundFill) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BackgroundFillFreeformGradient != nil {
		return obj.BackgroundFillFreeformGradient
	}

	if obj.BackgroundFillGradient != nil {
		return obj.BackgroundFillGradient
	}

	if obj.BackgroundFillSolid != nil {
		return obj.BackgroundFillSolid
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BackgroundFill) GetActualInstanceValue() (interface{}) {
	if obj.BackgroundFillFreeformGradient != nil {
		return *obj.BackgroundFillFreeformGradient
	}

	if obj.BackgroundFillGradient != nil {
		return *obj.BackgroundFillGradient
	}

	if obj.BackgroundFillSolid != nil {
		return *obj.BackgroundFillSolid
	}

	// all schemas are nil
	return nil
}

type NullableBackgroundFill struct {
	value *BackgroundFill
	isSet bool
}

func (v NullableBackgroundFill) Get() *BackgroundFill {
	return v.value
}

func (v *NullableBackgroundFill) Set(val *BackgroundFill) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundFill) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundFill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundFill(val *BackgroundFill) *NullableBackgroundFill {
	return &NullableBackgroundFill{value: val, isSet: true}
}

func (v NullableBackgroundFill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundFill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


