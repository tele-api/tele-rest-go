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

// BackgroundType - This object describes the type of a background. Currently, it can be one of  * [BackgroundTypeFill](https://core.telegram.org/bots/api/#backgroundtypefill) * [BackgroundTypeWallpaper](https://core.telegram.org/bots/api/#backgroundtypewallpaper) * [BackgroundTypePattern](https://core.telegram.org/bots/api/#backgroundtypepattern) * [BackgroundTypeChatTheme](https://core.telegram.org/bots/api/#backgroundtypechattheme)
type BackgroundType struct {
	BackgroundTypeChatTheme *BackgroundTypeChatTheme
	BackgroundTypeFill *BackgroundTypeFill
	BackgroundTypePattern *BackgroundTypePattern
	BackgroundTypeWallpaper *BackgroundTypeWallpaper
}

// BackgroundTypeChatThemeAsBackgroundType is a convenience function that returns BackgroundTypeChatTheme wrapped in BackgroundType
func BackgroundTypeChatThemeAsBackgroundType(v *BackgroundTypeChatTheme) BackgroundType {
	return BackgroundType{
		BackgroundTypeChatTheme: v,
	}
}

// BackgroundTypeFillAsBackgroundType is a convenience function that returns BackgroundTypeFill wrapped in BackgroundType
func BackgroundTypeFillAsBackgroundType(v *BackgroundTypeFill) BackgroundType {
	return BackgroundType{
		BackgroundTypeFill: v,
	}
}

// BackgroundTypePatternAsBackgroundType is a convenience function that returns BackgroundTypePattern wrapped in BackgroundType
func BackgroundTypePatternAsBackgroundType(v *BackgroundTypePattern) BackgroundType {
	return BackgroundType{
		BackgroundTypePattern: v,
	}
}

// BackgroundTypeWallpaperAsBackgroundType is a convenience function that returns BackgroundTypeWallpaper wrapped in BackgroundType
func BackgroundTypeWallpaperAsBackgroundType(v *BackgroundTypeWallpaper) BackgroundType {
	return BackgroundType{
		BackgroundTypeWallpaper: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BackgroundType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackgroundTypeChatTheme
	err = newStrictDecoder(data).Decode(&dst.BackgroundTypeChatTheme)
	if err == nil {
		jsonBackgroundTypeChatTheme, _ := json.Marshal(dst.BackgroundTypeChatTheme)
		if string(jsonBackgroundTypeChatTheme) == "{}" { // empty struct
			dst.BackgroundTypeChatTheme = nil
		} else {
			if err = validator.Validate(dst.BackgroundTypeChatTheme); err != nil {
				dst.BackgroundTypeChatTheme = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundTypeChatTheme = nil
	}

	// try to unmarshal data into BackgroundTypeFill
	err = newStrictDecoder(data).Decode(&dst.BackgroundTypeFill)
	if err == nil {
		jsonBackgroundTypeFill, _ := json.Marshal(dst.BackgroundTypeFill)
		if string(jsonBackgroundTypeFill) == "{}" { // empty struct
			dst.BackgroundTypeFill = nil
		} else {
			if err = validator.Validate(dst.BackgroundTypeFill); err != nil {
				dst.BackgroundTypeFill = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundTypeFill = nil
	}

	// try to unmarshal data into BackgroundTypePattern
	err = newStrictDecoder(data).Decode(&dst.BackgroundTypePattern)
	if err == nil {
		jsonBackgroundTypePattern, _ := json.Marshal(dst.BackgroundTypePattern)
		if string(jsonBackgroundTypePattern) == "{}" { // empty struct
			dst.BackgroundTypePattern = nil
		} else {
			if err = validator.Validate(dst.BackgroundTypePattern); err != nil {
				dst.BackgroundTypePattern = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundTypePattern = nil
	}

	// try to unmarshal data into BackgroundTypeWallpaper
	err = newStrictDecoder(data).Decode(&dst.BackgroundTypeWallpaper)
	if err == nil {
		jsonBackgroundTypeWallpaper, _ := json.Marshal(dst.BackgroundTypeWallpaper)
		if string(jsonBackgroundTypeWallpaper) == "{}" { // empty struct
			dst.BackgroundTypeWallpaper = nil
		} else {
			if err = validator.Validate(dst.BackgroundTypeWallpaper); err != nil {
				dst.BackgroundTypeWallpaper = nil
			} else {
				match++
			}
		}
	} else {
		dst.BackgroundTypeWallpaper = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackgroundTypeChatTheme = nil
		dst.BackgroundTypeFill = nil
		dst.BackgroundTypePattern = nil
		dst.BackgroundTypeWallpaper = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BackgroundType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BackgroundType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BackgroundType) MarshalJSON() ([]byte, error) {
	if src.BackgroundTypeChatTheme != nil {
		return json.Marshal(&src.BackgroundTypeChatTheme)
	}

	if src.BackgroundTypeFill != nil {
		return json.Marshal(&src.BackgroundTypeFill)
	}

	if src.BackgroundTypePattern != nil {
		return json.Marshal(&src.BackgroundTypePattern)
	}

	if src.BackgroundTypeWallpaper != nil {
		return json.Marshal(&src.BackgroundTypeWallpaper)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BackgroundType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BackgroundTypeChatTheme != nil {
		return obj.BackgroundTypeChatTheme
	}

	if obj.BackgroundTypeFill != nil {
		return obj.BackgroundTypeFill
	}

	if obj.BackgroundTypePattern != nil {
		return obj.BackgroundTypePattern
	}

	if obj.BackgroundTypeWallpaper != nil {
		return obj.BackgroundTypeWallpaper
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BackgroundType) GetActualInstanceValue() (interface{}) {
	if obj.BackgroundTypeChatTheme != nil {
		return *obj.BackgroundTypeChatTheme
	}

	if obj.BackgroundTypeFill != nil {
		return *obj.BackgroundTypeFill
	}

	if obj.BackgroundTypePattern != nil {
		return *obj.BackgroundTypePattern
	}

	if obj.BackgroundTypeWallpaper != nil {
		return *obj.BackgroundTypeWallpaper
	}

	// all schemas are nil
	return nil
}

type NullableBackgroundType struct {
	value *BackgroundType
	isSet bool
}

func (v NullableBackgroundType) Get() *BackgroundType {
	return v.value
}

func (v *NullableBackgroundType) Set(val *BackgroundType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundType(val *BackgroundType) *NullableBackgroundType {
	return &NullableBackgroundType{value: val, isSet: true}
}

func (v NullableBackgroundType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


