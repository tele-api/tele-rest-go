/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.3.0
 *    * - **Modified**: 2026-01-01T02:06:16.405655294Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
	"bytes"
	"fmt"
)

// checks if the UniqueGiftColors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGiftColors{}

// UniqueGiftColors This object contains information about the color scheme for a user's name, message replies and link previews based on a unique gift.
type UniqueGiftColors struct {
	// Custom emoji identifier of the unique gift's model
	ModelCustomEmojiId string `json:"model_custom_emoji_id"`
	// Custom emoji identifier of the unique gift's symbol
	SymbolCustomEmojiId string `json:"symbol_custom_emoji_id"`
	// Main color used in light themes; RGB format
	LightThemeMainColor int32 `json:"light_theme_main_color"`
	// List of 1-3 additional colors used in light themes; RGB format
	LightThemeOtherColors []int32 `json:"light_theme_other_colors"`
	// Main color used in dark themes; RGB format
	DarkThemeMainColor int32 `json:"dark_theme_main_color"`
	// List of 1-3 additional colors used in dark themes; RGB format
	DarkThemeOtherColors []int32 `json:"dark_theme_other_colors"`
}

type _UniqueGiftColors UniqueGiftColors

// NewUniqueGiftColors instantiates a new UniqueGiftColors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGiftColors(modelCustomEmojiId string, symbolCustomEmojiId string, lightThemeMainColor int32, lightThemeOtherColors []int32, darkThemeMainColor int32, darkThemeOtherColors []int32) *UniqueGiftColors {
	this := UniqueGiftColors{}
	this.ModelCustomEmojiId = modelCustomEmojiId
	this.SymbolCustomEmojiId = symbolCustomEmojiId
	this.LightThemeMainColor = lightThemeMainColor
	this.LightThemeOtherColors = lightThemeOtherColors
	this.DarkThemeMainColor = darkThemeMainColor
	this.DarkThemeOtherColors = darkThemeOtherColors
	return &this
}

// NewUniqueGiftColorsWithDefaults instantiates a new UniqueGiftColors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftColorsWithDefaults() *UniqueGiftColors {
	this := UniqueGiftColors{}
	return &this
}

// GetModelCustomEmojiId returns the ModelCustomEmojiId field value
func (o *UniqueGiftColors) GetModelCustomEmojiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelCustomEmojiId
}

// GetModelCustomEmojiIdOk returns a tuple with the ModelCustomEmojiId field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetModelCustomEmojiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelCustomEmojiId, true
}

// SetModelCustomEmojiId sets field value
func (o *UniqueGiftColors) SetModelCustomEmojiId(v string) {
	o.ModelCustomEmojiId = v
}

// GetSymbolCustomEmojiId returns the SymbolCustomEmojiId field value
func (o *UniqueGiftColors) GetSymbolCustomEmojiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SymbolCustomEmojiId
}

// GetSymbolCustomEmojiIdOk returns a tuple with the SymbolCustomEmojiId field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetSymbolCustomEmojiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SymbolCustomEmojiId, true
}

// SetSymbolCustomEmojiId sets field value
func (o *UniqueGiftColors) SetSymbolCustomEmojiId(v string) {
	o.SymbolCustomEmojiId = v
}

// GetLightThemeMainColor returns the LightThemeMainColor field value
func (o *UniqueGiftColors) GetLightThemeMainColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LightThemeMainColor
}

// GetLightThemeMainColorOk returns a tuple with the LightThemeMainColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetLightThemeMainColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LightThemeMainColor, true
}

// SetLightThemeMainColor sets field value
func (o *UniqueGiftColors) SetLightThemeMainColor(v int32) {
	o.LightThemeMainColor = v
}

// GetLightThemeOtherColors returns the LightThemeOtherColors field value
func (o *UniqueGiftColors) GetLightThemeOtherColors() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.LightThemeOtherColors
}

// GetLightThemeOtherColorsOk returns a tuple with the LightThemeOtherColors field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetLightThemeOtherColorsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LightThemeOtherColors, true
}

// SetLightThemeOtherColors sets field value
func (o *UniqueGiftColors) SetLightThemeOtherColors(v []int32) {
	o.LightThemeOtherColors = v
}

// GetDarkThemeMainColor returns the DarkThemeMainColor field value
func (o *UniqueGiftColors) GetDarkThemeMainColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DarkThemeMainColor
}

// GetDarkThemeMainColorOk returns a tuple with the DarkThemeMainColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetDarkThemeMainColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkThemeMainColor, true
}

// SetDarkThemeMainColor sets field value
func (o *UniqueGiftColors) SetDarkThemeMainColor(v int32) {
	o.DarkThemeMainColor = v
}

// GetDarkThemeOtherColors returns the DarkThemeOtherColors field value
func (o *UniqueGiftColors) GetDarkThemeOtherColors() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DarkThemeOtherColors
}

// GetDarkThemeOtherColorsOk returns a tuple with the DarkThemeOtherColors field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftColors) GetDarkThemeOtherColorsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DarkThemeOtherColors, true
}

// SetDarkThemeOtherColors sets field value
func (o *UniqueGiftColors) SetDarkThemeOtherColors(v []int32) {
	o.DarkThemeOtherColors = v
}

func (o UniqueGiftColors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGiftColors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model_custom_emoji_id"] = o.ModelCustomEmojiId
	toSerialize["symbol_custom_emoji_id"] = o.SymbolCustomEmojiId
	toSerialize["light_theme_main_color"] = o.LightThemeMainColor
	toSerialize["light_theme_other_colors"] = o.LightThemeOtherColors
	toSerialize["dark_theme_main_color"] = o.DarkThemeMainColor
	toSerialize["dark_theme_other_colors"] = o.DarkThemeOtherColors
	return toSerialize, nil
}

func (o *UniqueGiftColors) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model_custom_emoji_id",
		"symbol_custom_emoji_id",
		"light_theme_main_color",
		"light_theme_other_colors",
		"dark_theme_main_color",
		"dark_theme_other_colors",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUniqueGiftColors := _UniqueGiftColors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGiftColors)

	if err != nil {
		return err
	}

	*o = UniqueGiftColors(varUniqueGiftColors)

	return err
}

type NullableUniqueGiftColors struct {
	value *UniqueGiftColors
	isSet bool
}

func (v NullableUniqueGiftColors) Get() *UniqueGiftColors {
	return v.value
}

func (v *NullableUniqueGiftColors) Set(val *UniqueGiftColors) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGiftColors) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGiftColors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGiftColors(val *UniqueGiftColors) *NullableUniqueGiftColors {
	return &NullableUniqueGiftColors{value: val, isSet: true}
}

func (v NullableUniqueGiftColors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGiftColors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

