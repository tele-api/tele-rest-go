/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the StoryAreaTypeWeather type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryAreaTypeWeather{}

// StoryAreaTypeWeather Describes a story area containing weather information. Currently, a story can have up to 3 weather areas.
type StoryAreaTypeWeather struct {
	// Type of the area, always “weather”
	Type string `json:"type"`
	// Temperature, in degree Celsius
	Temperature float32 `json:"temperature"`
	// Emoji representing the weather
	Emoji string `json:"emoji"`
	// A color of the area background in the ARGB format
	BackgroundColor int32 `json:"background_color"`
}

type _StoryAreaTypeWeather StoryAreaTypeWeather

// NewStoryAreaTypeWeather instantiates a new StoryAreaTypeWeather object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryAreaTypeWeather(type_ string, temperature float32, emoji string, backgroundColor int32) *StoryAreaTypeWeather {
	this := StoryAreaTypeWeather{}
	this.Type = type_
	this.Temperature = temperature
	this.Emoji = emoji
	this.BackgroundColor = backgroundColor
	return &this
}

// NewStoryAreaTypeWeatherWithDefaults instantiates a new StoryAreaTypeWeather object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryAreaTypeWeatherWithDefaults() *StoryAreaTypeWeather {
	this := StoryAreaTypeWeather{}
	var type_ string = "weather"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StoryAreaTypeWeather) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeWeather) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StoryAreaTypeWeather) SetType(v string) {
	o.Type = v
}

// GetTemperature returns the Temperature field value
func (o *StoryAreaTypeWeather) GetTemperature() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeWeather) GetTemperatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Temperature, true
}

// SetTemperature sets field value
func (o *StoryAreaTypeWeather) SetTemperature(v float32) {
	o.Temperature = v
}

// GetEmoji returns the Emoji field value
func (o *StoryAreaTypeWeather) GetEmoji() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeWeather) GetEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emoji, true
}

// SetEmoji sets field value
func (o *StoryAreaTypeWeather) SetEmoji(v string) {
	o.Emoji = v
}

// GetBackgroundColor returns the BackgroundColor field value
func (o *StoryAreaTypeWeather) GetBackgroundColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeWeather) GetBackgroundColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundColor, true
}

// SetBackgroundColor sets field value
func (o *StoryAreaTypeWeather) SetBackgroundColor(v int32) {
	o.BackgroundColor = v
}

func (o StoryAreaTypeWeather) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryAreaTypeWeather) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["temperature"] = o.Temperature
	toSerialize["emoji"] = o.Emoji
	toSerialize["background_color"] = o.BackgroundColor
	return toSerialize, nil
}

func (o *StoryAreaTypeWeather) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"temperature",
		"emoji",
		"background_color",
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

	varStoryAreaTypeWeather := _StoryAreaTypeWeather{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStoryAreaTypeWeather)

	if err != nil {
		return err
	}

	*o = StoryAreaTypeWeather(varStoryAreaTypeWeather)

	return err
}

type NullableStoryAreaTypeWeather struct {
	value *StoryAreaTypeWeather
	isSet bool
}

func (v NullableStoryAreaTypeWeather) Get() *StoryAreaTypeWeather {
	return v.value
}

func (v *NullableStoryAreaTypeWeather) Set(val *StoryAreaTypeWeather) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaTypeWeather) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaTypeWeather) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaTypeWeather(val *StoryAreaTypeWeather) *NullableStoryAreaTypeWeather {
	return &NullableStoryAreaTypeWeather{value: val, isSet: true}
}

func (v NullableStoryAreaTypeWeather) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaTypeWeather) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


