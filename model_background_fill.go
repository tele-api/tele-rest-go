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


// BackgroundFill This object describes the way a background is filled based on the selected colors. Currently, it can be one of  * [BackgroundFillSolid](https://core.telegram.org/bots/api/#backgroundfillsolid) * [BackgroundFillGradient](https://core.telegram.org/bots/api/#backgroundfillgradient) * [BackgroundFillFreeformGradient](https://core.telegram.org/bots/api/#backgroundfillfreeformgradient)
type BackgroundFill struct {
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient
	BackgroundFillGradient *BackgroundFillGradient
	BackgroundFillSolid *BackgroundFillSolid
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BackgroundFill) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BackgroundFillFreeformGradient
	err = json.Unmarshal(data, &dst.BackgroundFillFreeformGradient);
	if err == nil {
		jsonBackgroundFillFreeformGradient, _ := json.Marshal(dst.BackgroundFillFreeformGradient)
		if string(jsonBackgroundFillFreeformGradient) == "{}" { // empty struct
			dst.BackgroundFillFreeformGradient = nil
		} else {
			return nil // data stored in dst.BackgroundFillFreeformGradient, return on the first match
		}
	} else {
		dst.BackgroundFillFreeformGradient = nil
	}

	// try to unmarshal JSON data into BackgroundFillGradient
	err = json.Unmarshal(data, &dst.BackgroundFillGradient);
	if err == nil {
		jsonBackgroundFillGradient, _ := json.Marshal(dst.BackgroundFillGradient)
		if string(jsonBackgroundFillGradient) == "{}" { // empty struct
			dst.BackgroundFillGradient = nil
		} else {
			return nil // data stored in dst.BackgroundFillGradient, return on the first match
		}
	} else {
		dst.BackgroundFillGradient = nil
	}

	// try to unmarshal JSON data into BackgroundFillSolid
	err = json.Unmarshal(data, &dst.BackgroundFillSolid);
	if err == nil {
		jsonBackgroundFillSolid, _ := json.Marshal(dst.BackgroundFillSolid)
		if string(jsonBackgroundFillSolid) == "{}" { // empty struct
			dst.BackgroundFillSolid = nil
		} else {
			return nil // data stored in dst.BackgroundFillSolid, return on the first match
		}
	} else {
		dst.BackgroundFillSolid = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BackgroundFill)")
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

	return nil, nil // no data in anyOf schemas
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


