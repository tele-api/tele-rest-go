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

// checks if the InlineKeyboardMarkup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineKeyboardMarkup{}

// InlineKeyboardMarkup This object represents an [inline keyboard](https://core.telegram.org/bots/features#inline-keyboards) that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of [InlineKeyboardButton](https://core.telegram.org/bots/api/#inlinekeyboardbutton) objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type _InlineKeyboardMarkup InlineKeyboardMarkup

// NewInlineKeyboardMarkup instantiates a new InlineKeyboardMarkup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineKeyboardMarkup(inlineKeyboard [][]InlineKeyboardButton) *InlineKeyboardMarkup {
	this := InlineKeyboardMarkup{}
	this.InlineKeyboard = inlineKeyboard
	return &this
}

// NewInlineKeyboardMarkupWithDefaults instantiates a new InlineKeyboardMarkup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineKeyboardMarkupWithDefaults() *InlineKeyboardMarkup {
	this := InlineKeyboardMarkup{}
	return &this
}

// GetInlineKeyboard returns the InlineKeyboard field value
func (o *InlineKeyboardMarkup) GetInlineKeyboard() [][]InlineKeyboardButton {
	if o == nil {
		var ret [][]InlineKeyboardButton
		return ret
	}

	return o.InlineKeyboard
}

// GetInlineKeyboardOk returns a tuple with the InlineKeyboard field value
// and a boolean to check if the value has been set.
func (o *InlineKeyboardMarkup) GetInlineKeyboardOk() ([][]InlineKeyboardButton, bool) {
	if o == nil {
		return nil, false
	}
	return o.InlineKeyboard, true
}

// SetInlineKeyboard sets field value
func (o *InlineKeyboardMarkup) SetInlineKeyboard(v [][]InlineKeyboardButton) {
	o.InlineKeyboard = v
}

func (o InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineKeyboardMarkup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inline_keyboard"] = o.InlineKeyboard
	return toSerialize, nil
}

func (o *InlineKeyboardMarkup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inline_keyboard",
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

	varInlineKeyboardMarkup := _InlineKeyboardMarkup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineKeyboardMarkup)

	if err != nil {
		return err
	}

	*o = InlineKeyboardMarkup(varInlineKeyboardMarkup)

	return err
}

type NullableInlineKeyboardMarkup struct {
	value *InlineKeyboardMarkup
	isSet bool
}

func (v NullableInlineKeyboardMarkup) Get() *InlineKeyboardMarkup {
	return v.value
}

func (v *NullableInlineKeyboardMarkup) Set(val *InlineKeyboardMarkup) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineKeyboardMarkup) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineKeyboardMarkup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineKeyboardMarkup(val *InlineKeyboardMarkup) *NullableInlineKeyboardMarkup {
	return &NullableInlineKeyboardMarkup{value: val, isSet: true}
}

func (v NullableInlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineKeyboardMarkup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


