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

// StoryAreaType - Describes the type of a clickable area on a story. Currently, it can be one of  * [StoryAreaTypeLocation](https://core.telegram.org/bots/api/#storyareatypelocation) * [StoryAreaTypeSuggestedReaction](https://core.telegram.org/bots/api/#storyareatypesuggestedreaction) * [StoryAreaTypeLink](https://core.telegram.org/bots/api/#storyareatypelink) * [StoryAreaTypeWeather](https://core.telegram.org/bots/api/#storyareatypeweather) * [StoryAreaTypeUniqueGift](https://core.telegram.org/bots/api/#storyareatypeuniquegift)
type StoryAreaType struct {
	StoryAreaTypeLink *StoryAreaTypeLink
	StoryAreaTypeLocation *StoryAreaTypeLocation
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction
	StoryAreaTypeUniqueGift *StoryAreaTypeUniqueGift
	StoryAreaTypeWeather *StoryAreaTypeWeather
}

// StoryAreaTypeLinkAsStoryAreaType is a convenience function that returns StoryAreaTypeLink wrapped in StoryAreaType
func StoryAreaTypeLinkAsStoryAreaType(v *StoryAreaTypeLink) StoryAreaType {
	return StoryAreaType{
		StoryAreaTypeLink: v,
	}
}

// StoryAreaTypeLocationAsStoryAreaType is a convenience function that returns StoryAreaTypeLocation wrapped in StoryAreaType
func StoryAreaTypeLocationAsStoryAreaType(v *StoryAreaTypeLocation) StoryAreaType {
	return StoryAreaType{
		StoryAreaTypeLocation: v,
	}
}

// StoryAreaTypeSuggestedReactionAsStoryAreaType is a convenience function that returns StoryAreaTypeSuggestedReaction wrapped in StoryAreaType
func StoryAreaTypeSuggestedReactionAsStoryAreaType(v *StoryAreaTypeSuggestedReaction) StoryAreaType {
	return StoryAreaType{
		StoryAreaTypeSuggestedReaction: v,
	}
}

// StoryAreaTypeUniqueGiftAsStoryAreaType is a convenience function that returns StoryAreaTypeUniqueGift wrapped in StoryAreaType
func StoryAreaTypeUniqueGiftAsStoryAreaType(v *StoryAreaTypeUniqueGift) StoryAreaType {
	return StoryAreaType{
		StoryAreaTypeUniqueGift: v,
	}
}

// StoryAreaTypeWeatherAsStoryAreaType is a convenience function that returns StoryAreaTypeWeather wrapped in StoryAreaType
func StoryAreaTypeWeatherAsStoryAreaType(v *StoryAreaTypeWeather) StoryAreaType {
	return StoryAreaType{
		StoryAreaTypeWeather: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StoryAreaType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StoryAreaTypeLink
	err = newStrictDecoder(data).Decode(&dst.StoryAreaTypeLink)
	if err == nil {
		jsonStoryAreaTypeLink, _ := json.Marshal(dst.StoryAreaTypeLink)
		if string(jsonStoryAreaTypeLink) == "{}" { // empty struct
			dst.StoryAreaTypeLink = nil
		} else {
			if err = validator.Validate(dst.StoryAreaTypeLink); err != nil {
				dst.StoryAreaTypeLink = nil
			} else {
				match++
			}
		}
	} else {
		dst.StoryAreaTypeLink = nil
	}

	// try to unmarshal data into StoryAreaTypeLocation
	err = newStrictDecoder(data).Decode(&dst.StoryAreaTypeLocation)
	if err == nil {
		jsonStoryAreaTypeLocation, _ := json.Marshal(dst.StoryAreaTypeLocation)
		if string(jsonStoryAreaTypeLocation) == "{}" { // empty struct
			dst.StoryAreaTypeLocation = nil
		} else {
			if err = validator.Validate(dst.StoryAreaTypeLocation); err != nil {
				dst.StoryAreaTypeLocation = nil
			} else {
				match++
			}
		}
	} else {
		dst.StoryAreaTypeLocation = nil
	}

	// try to unmarshal data into StoryAreaTypeSuggestedReaction
	err = newStrictDecoder(data).Decode(&dst.StoryAreaTypeSuggestedReaction)
	if err == nil {
		jsonStoryAreaTypeSuggestedReaction, _ := json.Marshal(dst.StoryAreaTypeSuggestedReaction)
		if string(jsonStoryAreaTypeSuggestedReaction) == "{}" { // empty struct
			dst.StoryAreaTypeSuggestedReaction = nil
		} else {
			if err = validator.Validate(dst.StoryAreaTypeSuggestedReaction); err != nil {
				dst.StoryAreaTypeSuggestedReaction = nil
			} else {
				match++
			}
		}
	} else {
		dst.StoryAreaTypeSuggestedReaction = nil
	}

	// try to unmarshal data into StoryAreaTypeUniqueGift
	err = newStrictDecoder(data).Decode(&dst.StoryAreaTypeUniqueGift)
	if err == nil {
		jsonStoryAreaTypeUniqueGift, _ := json.Marshal(dst.StoryAreaTypeUniqueGift)
		if string(jsonStoryAreaTypeUniqueGift) == "{}" { // empty struct
			dst.StoryAreaTypeUniqueGift = nil
		} else {
			if err = validator.Validate(dst.StoryAreaTypeUniqueGift); err != nil {
				dst.StoryAreaTypeUniqueGift = nil
			} else {
				match++
			}
		}
	} else {
		dst.StoryAreaTypeUniqueGift = nil
	}

	// try to unmarshal data into StoryAreaTypeWeather
	err = newStrictDecoder(data).Decode(&dst.StoryAreaTypeWeather)
	if err == nil {
		jsonStoryAreaTypeWeather, _ := json.Marshal(dst.StoryAreaTypeWeather)
		if string(jsonStoryAreaTypeWeather) == "{}" { // empty struct
			dst.StoryAreaTypeWeather = nil
		} else {
			if err = validator.Validate(dst.StoryAreaTypeWeather); err != nil {
				dst.StoryAreaTypeWeather = nil
			} else {
				match++
			}
		}
	} else {
		dst.StoryAreaTypeWeather = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.StoryAreaTypeLink = nil
		dst.StoryAreaTypeLocation = nil
		dst.StoryAreaTypeSuggestedReaction = nil
		dst.StoryAreaTypeUniqueGift = nil
		dst.StoryAreaTypeWeather = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StoryAreaType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StoryAreaType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StoryAreaType) MarshalJSON() ([]byte, error) {
	if src.StoryAreaTypeLink != nil {
		return json.Marshal(&src.StoryAreaTypeLink)
	}

	if src.StoryAreaTypeLocation != nil {
		return json.Marshal(&src.StoryAreaTypeLocation)
	}

	if src.StoryAreaTypeSuggestedReaction != nil {
		return json.Marshal(&src.StoryAreaTypeSuggestedReaction)
	}

	if src.StoryAreaTypeUniqueGift != nil {
		return json.Marshal(&src.StoryAreaTypeUniqueGift)
	}

	if src.StoryAreaTypeWeather != nil {
		return json.Marshal(&src.StoryAreaTypeWeather)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StoryAreaType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StoryAreaTypeLink != nil {
		return obj.StoryAreaTypeLink
	}

	if obj.StoryAreaTypeLocation != nil {
		return obj.StoryAreaTypeLocation
	}

	if obj.StoryAreaTypeSuggestedReaction != nil {
		return obj.StoryAreaTypeSuggestedReaction
	}

	if obj.StoryAreaTypeUniqueGift != nil {
		return obj.StoryAreaTypeUniqueGift
	}

	if obj.StoryAreaTypeWeather != nil {
		return obj.StoryAreaTypeWeather
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StoryAreaType) GetActualInstanceValue() (interface{}) {
	if obj.StoryAreaTypeLink != nil {
		return *obj.StoryAreaTypeLink
	}

	if obj.StoryAreaTypeLocation != nil {
		return *obj.StoryAreaTypeLocation
	}

	if obj.StoryAreaTypeSuggestedReaction != nil {
		return *obj.StoryAreaTypeSuggestedReaction
	}

	if obj.StoryAreaTypeUniqueGift != nil {
		return *obj.StoryAreaTypeUniqueGift
	}

	if obj.StoryAreaTypeWeather != nil {
		return *obj.StoryAreaTypeWeather
	}

	// all schemas are nil
	return nil
}

type NullableStoryAreaType struct {
	value *StoryAreaType
	isSet bool
}

func (v NullableStoryAreaType) Get() *StoryAreaType {
	return v.value
}

func (v *NullableStoryAreaType) Set(val *StoryAreaType) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaType) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaType(val *StoryAreaType) *NullableStoryAreaType {
	return &NullableStoryAreaType{value: val, isSet: true}
}

func (v NullableStoryAreaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


