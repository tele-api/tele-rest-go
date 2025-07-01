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

// OwnedGift - This object describes a gift received and owned by a user or a chat. Currently, it can be one of  * [OwnedGiftRegular](https://core.telegram.org/bots/api/#ownedgiftregular) * [OwnedGiftUnique](https://core.telegram.org/bots/api/#ownedgiftunique)
type OwnedGift struct {
	OwnedGiftRegular *OwnedGiftRegular
	OwnedGiftUnique *OwnedGiftUnique
}

// OwnedGiftRegularAsOwnedGift is a convenience function that returns OwnedGiftRegular wrapped in OwnedGift
func OwnedGiftRegularAsOwnedGift(v *OwnedGiftRegular) OwnedGift {
	return OwnedGift{
		OwnedGiftRegular: v,
	}
}

// OwnedGiftUniqueAsOwnedGift is a convenience function that returns OwnedGiftUnique wrapped in OwnedGift
func OwnedGiftUniqueAsOwnedGift(v *OwnedGiftUnique) OwnedGift {
	return OwnedGift{
		OwnedGiftUnique: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OwnedGift) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OwnedGiftRegular
	err = newStrictDecoder(data).Decode(&dst.OwnedGiftRegular)
	if err == nil {
		jsonOwnedGiftRegular, _ := json.Marshal(dst.OwnedGiftRegular)
		if string(jsonOwnedGiftRegular) == "{}" { // empty struct
			dst.OwnedGiftRegular = nil
		} else {
			if err = validator.Validate(dst.OwnedGiftRegular); err != nil {
				dst.OwnedGiftRegular = nil
			} else {
				match++
			}
		}
	} else {
		dst.OwnedGiftRegular = nil
	}

	// try to unmarshal data into OwnedGiftUnique
	err = newStrictDecoder(data).Decode(&dst.OwnedGiftUnique)
	if err == nil {
		jsonOwnedGiftUnique, _ := json.Marshal(dst.OwnedGiftUnique)
		if string(jsonOwnedGiftUnique) == "{}" { // empty struct
			dst.OwnedGiftUnique = nil
		} else {
			if err = validator.Validate(dst.OwnedGiftUnique); err != nil {
				dst.OwnedGiftUnique = nil
			} else {
				match++
			}
		}
	} else {
		dst.OwnedGiftUnique = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OwnedGiftRegular = nil
		dst.OwnedGiftUnique = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OwnedGift)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OwnedGift)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OwnedGift) MarshalJSON() ([]byte, error) {
	if src.OwnedGiftRegular != nil {
		return json.Marshal(&src.OwnedGiftRegular)
	}

	if src.OwnedGiftUnique != nil {
		return json.Marshal(&src.OwnedGiftUnique)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OwnedGift) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.OwnedGiftRegular != nil {
		return obj.OwnedGiftRegular
	}

	if obj.OwnedGiftUnique != nil {
		return obj.OwnedGiftUnique
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OwnedGift) GetActualInstanceValue() (interface{}) {
	if obj.OwnedGiftRegular != nil {
		return *obj.OwnedGiftRegular
	}

	if obj.OwnedGiftUnique != nil {
		return *obj.OwnedGiftUnique
	}

	// all schemas are nil
	return nil
}

type NullableOwnedGift struct {
	value *OwnedGift
	isSet bool
}

func (v NullableOwnedGift) Get() *OwnedGift {
	return v.value
}

func (v *NullableOwnedGift) Set(val *OwnedGift) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnedGift) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnedGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnedGift(val *OwnedGift) *NullableOwnedGift {
	return &NullableOwnedGift{value: val, isSet: true}
}

func (v NullableOwnedGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnedGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


