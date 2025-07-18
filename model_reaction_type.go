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

// ReactionType - This object describes the type of a reaction. Currently, it can be one of  * [ReactionTypeEmoji](https://core.telegram.org/bots/api/#reactiontypeemoji) * [ReactionTypeCustomEmoji](https://core.telegram.org/bots/api/#reactiontypecustomemoji) * [ReactionTypePaid](https://core.telegram.org/bots/api/#reactiontypepaid)
type ReactionType struct {
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji
	ReactionTypeEmoji *ReactionTypeEmoji
	ReactionTypePaid *ReactionTypePaid
}

// ReactionTypeCustomEmojiAsReactionType is a convenience function that returns ReactionTypeCustomEmoji wrapped in ReactionType
func ReactionTypeCustomEmojiAsReactionType(v *ReactionTypeCustomEmoji) ReactionType {
	return ReactionType{
		ReactionTypeCustomEmoji: v,
	}
}

// ReactionTypeEmojiAsReactionType is a convenience function that returns ReactionTypeEmoji wrapped in ReactionType
func ReactionTypeEmojiAsReactionType(v *ReactionTypeEmoji) ReactionType {
	return ReactionType{
		ReactionTypeEmoji: v,
	}
}

// ReactionTypePaidAsReactionType is a convenience function that returns ReactionTypePaid wrapped in ReactionType
func ReactionTypePaidAsReactionType(v *ReactionTypePaid) ReactionType {
	return ReactionType{
		ReactionTypePaid: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReactionType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ReactionTypeCustomEmoji
	err = newStrictDecoder(data).Decode(&dst.ReactionTypeCustomEmoji)
	if err == nil {
		jsonReactionTypeCustomEmoji, _ := json.Marshal(dst.ReactionTypeCustomEmoji)
		if string(jsonReactionTypeCustomEmoji) == "{}" { // empty struct
			dst.ReactionTypeCustomEmoji = nil
		} else {
			if err = validator.Validate(dst.ReactionTypeCustomEmoji); err != nil {
				dst.ReactionTypeCustomEmoji = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReactionTypeCustomEmoji = nil
	}

	// try to unmarshal data into ReactionTypeEmoji
	err = newStrictDecoder(data).Decode(&dst.ReactionTypeEmoji)
	if err == nil {
		jsonReactionTypeEmoji, _ := json.Marshal(dst.ReactionTypeEmoji)
		if string(jsonReactionTypeEmoji) == "{}" { // empty struct
			dst.ReactionTypeEmoji = nil
		} else {
			if err = validator.Validate(dst.ReactionTypeEmoji); err != nil {
				dst.ReactionTypeEmoji = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReactionTypeEmoji = nil
	}

	// try to unmarshal data into ReactionTypePaid
	err = newStrictDecoder(data).Decode(&dst.ReactionTypePaid)
	if err == nil {
		jsonReactionTypePaid, _ := json.Marshal(dst.ReactionTypePaid)
		if string(jsonReactionTypePaid) == "{}" { // empty struct
			dst.ReactionTypePaid = nil
		} else {
			if err = validator.Validate(dst.ReactionTypePaid); err != nil {
				dst.ReactionTypePaid = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReactionTypePaid = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ReactionTypeCustomEmoji = nil
		dst.ReactionTypeEmoji = nil
		dst.ReactionTypePaid = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ReactionType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReactionType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReactionType) MarshalJSON() ([]byte, error) {
	if src.ReactionTypeCustomEmoji != nil {
		return json.Marshal(&src.ReactionTypeCustomEmoji)
	}

	if src.ReactionTypeEmoji != nil {
		return json.Marshal(&src.ReactionTypeEmoji)
	}

	if src.ReactionTypePaid != nil {
		return json.Marshal(&src.ReactionTypePaid)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReactionType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ReactionTypeCustomEmoji != nil {
		return obj.ReactionTypeCustomEmoji
	}

	if obj.ReactionTypeEmoji != nil {
		return obj.ReactionTypeEmoji
	}

	if obj.ReactionTypePaid != nil {
		return obj.ReactionTypePaid
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ReactionType) GetActualInstanceValue() (interface{}) {
	if obj.ReactionTypeCustomEmoji != nil {
		return *obj.ReactionTypeCustomEmoji
	}

	if obj.ReactionTypeEmoji != nil {
		return *obj.ReactionTypeEmoji
	}

	if obj.ReactionTypePaid != nil {
		return *obj.ReactionTypePaid
	}

	// all schemas are nil
	return nil
}

type NullableReactionType struct {
	value *ReactionType
	isSet bool
}

func (v NullableReactionType) Get() *ReactionType {
	return v.value
}

func (v *NullableReactionType) Set(val *ReactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionType(val *ReactionType) *NullableReactionType {
	return &NullableReactionType{value: val, isSet: true}
}

func (v NullableReactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


