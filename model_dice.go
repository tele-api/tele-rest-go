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
	"bytes"
	"fmt"
)

// checks if the Dice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dice{}

// Dice This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`
	// Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int32 `json:"value"`
}

type _Dice Dice

// NewDice instantiates a new Dice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDice(emoji string, value int32) *Dice {
	this := Dice{}
	this.Emoji = emoji
	this.Value = value
	return &this
}

// NewDiceWithDefaults instantiates a new Dice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiceWithDefaults() *Dice {
	this := Dice{}
	return &this
}

// GetEmoji returns the Emoji field value
func (o *Dice) GetEmoji() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value
// and a boolean to check if the value has been set.
func (o *Dice) GetEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emoji, true
}

// SetEmoji sets field value
func (o *Dice) SetEmoji(v string) {
	o.Emoji = v
}

// GetValue returns the Value field value
func (o *Dice) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Dice) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Dice) SetValue(v int32) {
	o.Value = v
}

func (o Dice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emoji"] = o.Emoji
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *Dice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emoji",
		"value",
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

	varDice := _Dice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDice)

	if err != nil {
		return err
	}

	*o = Dice(varDice)

	return err
}

type NullableDice struct {
	value *Dice
	isSet bool
}

func (v NullableDice) Get() *Dice {
	return v.value
}

func (v *NullableDice) Set(val *Dice) {
	v.value = val
	v.isSet = true
}

func (v NullableDice) IsSet() bool {
	return v.isSet
}

func (v *NullableDice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDice(val *Dice) *NullableDice {
	return &NullableDice{value: val, isSet: true}
}

func (v NullableDice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


