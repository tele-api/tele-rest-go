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


// StoryAreaType Describes the type of a clickable area on a story. Currently, it can be one of  * [StoryAreaTypeLocation](https://core.telegram.org/bots/api/#storyareatypelocation) * [StoryAreaTypeSuggestedReaction](https://core.telegram.org/bots/api/#storyareatypesuggestedreaction) * [StoryAreaTypeLink](https://core.telegram.org/bots/api/#storyareatypelink) * [StoryAreaTypeWeather](https://core.telegram.org/bots/api/#storyareatypeweather) * [StoryAreaTypeUniqueGift](https://core.telegram.org/bots/api/#storyareatypeuniquegift)
type StoryAreaType struct {
	StoryAreaTypeLink *StoryAreaTypeLink
	StoryAreaTypeLocation *StoryAreaTypeLocation
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction
	StoryAreaTypeUniqueGift *StoryAreaTypeUniqueGift
	StoryAreaTypeWeather *StoryAreaTypeWeather
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StoryAreaType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into StoryAreaTypeLink
	err = json.Unmarshal(data, &dst.StoryAreaTypeLink);
	if err == nil {
		jsonStoryAreaTypeLink, _ := json.Marshal(dst.StoryAreaTypeLink)
		if string(jsonStoryAreaTypeLink) == "{}" { // empty struct
			dst.StoryAreaTypeLink = nil
		} else {
			return nil // data stored in dst.StoryAreaTypeLink, return on the first match
		}
	} else {
		dst.StoryAreaTypeLink = nil
	}

	// try to unmarshal JSON data into StoryAreaTypeLocation
	err = json.Unmarshal(data, &dst.StoryAreaTypeLocation);
	if err == nil {
		jsonStoryAreaTypeLocation, _ := json.Marshal(dst.StoryAreaTypeLocation)
		if string(jsonStoryAreaTypeLocation) == "{}" { // empty struct
			dst.StoryAreaTypeLocation = nil
		} else {
			return nil // data stored in dst.StoryAreaTypeLocation, return on the first match
		}
	} else {
		dst.StoryAreaTypeLocation = nil
	}

	// try to unmarshal JSON data into StoryAreaTypeSuggestedReaction
	err = json.Unmarshal(data, &dst.StoryAreaTypeSuggestedReaction);
	if err == nil {
		jsonStoryAreaTypeSuggestedReaction, _ := json.Marshal(dst.StoryAreaTypeSuggestedReaction)
		if string(jsonStoryAreaTypeSuggestedReaction) == "{}" { // empty struct
			dst.StoryAreaTypeSuggestedReaction = nil
		} else {
			return nil // data stored in dst.StoryAreaTypeSuggestedReaction, return on the first match
		}
	} else {
		dst.StoryAreaTypeSuggestedReaction = nil
	}

	// try to unmarshal JSON data into StoryAreaTypeUniqueGift
	err = json.Unmarshal(data, &dst.StoryAreaTypeUniqueGift);
	if err == nil {
		jsonStoryAreaTypeUniqueGift, _ := json.Marshal(dst.StoryAreaTypeUniqueGift)
		if string(jsonStoryAreaTypeUniqueGift) == "{}" { // empty struct
			dst.StoryAreaTypeUniqueGift = nil
		} else {
			return nil // data stored in dst.StoryAreaTypeUniqueGift, return on the first match
		}
	} else {
		dst.StoryAreaTypeUniqueGift = nil
	}

	// try to unmarshal JSON data into StoryAreaTypeWeather
	err = json.Unmarshal(data, &dst.StoryAreaTypeWeather);
	if err == nil {
		jsonStoryAreaTypeWeather, _ := json.Marshal(dst.StoryAreaTypeWeather)
		if string(jsonStoryAreaTypeWeather) == "{}" { // empty struct
			dst.StoryAreaTypeWeather = nil
		} else {
			return nil // data stored in dst.StoryAreaTypeWeather, return on the first match
		}
	} else {
		dst.StoryAreaTypeWeather = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(StoryAreaType)")
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

	return nil, nil // no data in anyOf schemas
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


