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
	"bytes"
	"fmt"
)

// checks if the ReplyKeyboardMarkup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplyKeyboardMarkup{}

// ReplyKeyboardMarkup This object represents a [custom keyboard](https://core.telegram.org/bots/features#keyboards) with reply options (see [Introduction to bots](https://core.telegram.org/bots/features#keyboards) for details and examples). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of [KeyboardButton](https://core.telegram.org/bots/api/#keyboardbutton) objects
	Keyboard [][]KeyboardButton `json:"keyboard"`
	// *Optional*. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to *false*, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent *bool `json:"is_persistent,omitempty"`
	// *Optional*. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to *false*, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard *bool `json:"resize_keyboard,omitempty"`
	// *Optional*. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to *false*.
	OneTimeKeyboard *bool `json:"one_time_keyboard,omitempty"`
	// *Optional*. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	// *Optional*. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.    *Example:* A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
	Selective *bool `json:"selective,omitempty"`
}

type _ReplyKeyboardMarkup ReplyKeyboardMarkup

// NewReplyKeyboardMarkup instantiates a new ReplyKeyboardMarkup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplyKeyboardMarkup(keyboard [][]KeyboardButton) *ReplyKeyboardMarkup {
	this := ReplyKeyboardMarkup{}
	this.Keyboard = keyboard
	var isPersistent bool = false
	this.IsPersistent = &isPersistent
	var resizeKeyboard bool = false
	this.ResizeKeyboard = &resizeKeyboard
	var oneTimeKeyboard bool = false
	this.OneTimeKeyboard = &oneTimeKeyboard
	return &this
}

// NewReplyKeyboardMarkupWithDefaults instantiates a new ReplyKeyboardMarkup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplyKeyboardMarkupWithDefaults() *ReplyKeyboardMarkup {
	this := ReplyKeyboardMarkup{}
	var isPersistent bool = false
	this.IsPersistent = &isPersistent
	var resizeKeyboard bool = false
	this.ResizeKeyboard = &resizeKeyboard
	var oneTimeKeyboard bool = false
	this.OneTimeKeyboard = &oneTimeKeyboard
	return &this
}

// GetKeyboard returns the Keyboard field value
func (o *ReplyKeyboardMarkup) GetKeyboard() [][]KeyboardButton {
	if o == nil {
		var ret [][]KeyboardButton
		return ret
	}

	return o.Keyboard
}

// GetKeyboardOk returns a tuple with the Keyboard field value
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetKeyboardOk() ([][]KeyboardButton, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keyboard, true
}

// SetKeyboard sets field value
func (o *ReplyKeyboardMarkup) SetKeyboard(v [][]KeyboardButton) {
	o.Keyboard = v
}

// GetIsPersistent returns the IsPersistent field value if set, zero value otherwise.
func (o *ReplyKeyboardMarkup) GetIsPersistent() bool {
	if o == nil || IsNil(o.IsPersistent) {
		var ret bool
		return ret
	}
	return *o.IsPersistent
}

// GetIsPersistentOk returns a tuple with the IsPersistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetIsPersistentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPersistent) {
		return nil, false
	}
	return o.IsPersistent, true
}

// HasIsPersistent returns a boolean if a field has been set.
func (o *ReplyKeyboardMarkup) HasIsPersistent() bool {
	if o != nil && !IsNil(o.IsPersistent) {
		return true
	}

	return false
}

// SetIsPersistent gets a reference to the given bool and assigns it to the IsPersistent field.
func (o *ReplyKeyboardMarkup) SetIsPersistent(v bool) {
	o.IsPersistent = &v
}


// GetResizeKeyboard returns the ResizeKeyboard field value if set, zero value otherwise.
func (o *ReplyKeyboardMarkup) GetResizeKeyboard() bool {
	if o == nil || IsNil(o.ResizeKeyboard) {
		var ret bool
		return ret
	}
	return *o.ResizeKeyboard
}

// GetResizeKeyboardOk returns a tuple with the ResizeKeyboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetResizeKeyboardOk() (*bool, bool) {
	if o == nil || IsNil(o.ResizeKeyboard) {
		return nil, false
	}
	return o.ResizeKeyboard, true
}

// HasResizeKeyboard returns a boolean if a field has been set.
func (o *ReplyKeyboardMarkup) HasResizeKeyboard() bool {
	if o != nil && !IsNil(o.ResizeKeyboard) {
		return true
	}

	return false
}

// SetResizeKeyboard gets a reference to the given bool and assigns it to the ResizeKeyboard field.
func (o *ReplyKeyboardMarkup) SetResizeKeyboard(v bool) {
	o.ResizeKeyboard = &v
}


// GetOneTimeKeyboard returns the OneTimeKeyboard field value if set, zero value otherwise.
func (o *ReplyKeyboardMarkup) GetOneTimeKeyboard() bool {
	if o == nil || IsNil(o.OneTimeKeyboard) {
		var ret bool
		return ret
	}
	return *o.OneTimeKeyboard
}

// GetOneTimeKeyboardOk returns a tuple with the OneTimeKeyboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetOneTimeKeyboardOk() (*bool, bool) {
	if o == nil || IsNil(o.OneTimeKeyboard) {
		return nil, false
	}
	return o.OneTimeKeyboard, true
}

// HasOneTimeKeyboard returns a boolean if a field has been set.
func (o *ReplyKeyboardMarkup) HasOneTimeKeyboard() bool {
	if o != nil && !IsNil(o.OneTimeKeyboard) {
		return true
	}

	return false
}

// SetOneTimeKeyboard gets a reference to the given bool and assigns it to the OneTimeKeyboard field.
func (o *ReplyKeyboardMarkup) SetOneTimeKeyboard(v bool) {
	o.OneTimeKeyboard = &v
}


// GetInputFieldPlaceholder returns the InputFieldPlaceholder field value if set, zero value otherwise.
func (o *ReplyKeyboardMarkup) GetInputFieldPlaceholder() string {
	if o == nil || IsNil(o.InputFieldPlaceholder) {
		var ret string
		return ret
	}
	return *o.InputFieldPlaceholder
}

// GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetInputFieldPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.InputFieldPlaceholder) {
		return nil, false
	}
	return o.InputFieldPlaceholder, true
}

// HasInputFieldPlaceholder returns a boolean if a field has been set.
func (o *ReplyKeyboardMarkup) HasInputFieldPlaceholder() bool {
	if o != nil && !IsNil(o.InputFieldPlaceholder) {
		return true
	}

	return false
}

// SetInputFieldPlaceholder gets a reference to the given string and assigns it to the InputFieldPlaceholder field.
func (o *ReplyKeyboardMarkup) SetInputFieldPlaceholder(v string) {
	o.InputFieldPlaceholder = &v
}


// GetSelective returns the Selective field value if set, zero value otherwise.
func (o *ReplyKeyboardMarkup) GetSelective() bool {
	if o == nil || IsNil(o.Selective) {
		var ret bool
		return ret
	}
	return *o.Selective
}

// GetSelectiveOk returns a tuple with the Selective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardMarkup) GetSelectiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Selective) {
		return nil, false
	}
	return o.Selective, true
}

// HasSelective returns a boolean if a field has been set.
func (o *ReplyKeyboardMarkup) HasSelective() bool {
	if o != nil && !IsNil(o.Selective) {
		return true
	}

	return false
}

// SetSelective gets a reference to the given bool and assigns it to the Selective field.
func (o *ReplyKeyboardMarkup) SetSelective(v bool) {
	o.Selective = &v
}


func (o ReplyKeyboardMarkup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplyKeyboardMarkup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyboard"] = o.Keyboard
	if !IsNil(o.IsPersistent) {
		toSerialize["is_persistent"] = o.IsPersistent
	}
	if !IsNil(o.ResizeKeyboard) {
		toSerialize["resize_keyboard"] = o.ResizeKeyboard
	}
	if !IsNil(o.OneTimeKeyboard) {
		toSerialize["one_time_keyboard"] = o.OneTimeKeyboard
	}
	if !IsNil(o.InputFieldPlaceholder) {
		toSerialize["input_field_placeholder"] = o.InputFieldPlaceholder
	}
	if !IsNil(o.Selective) {
		toSerialize["selective"] = o.Selective
	}
	return toSerialize, nil
}

func (o *ReplyKeyboardMarkup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keyboard",
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

	varReplyKeyboardMarkup := _ReplyKeyboardMarkup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplyKeyboardMarkup)

	if err != nil {
		return err
	}

	*o = ReplyKeyboardMarkup(varReplyKeyboardMarkup)

	return err
}

type NullableReplyKeyboardMarkup struct {
	value *ReplyKeyboardMarkup
	isSet bool
}

func (v NullableReplyKeyboardMarkup) Get() *ReplyKeyboardMarkup {
	return v.value
}

func (v *NullableReplyKeyboardMarkup) Set(val *ReplyKeyboardMarkup) {
	v.value = val
	v.isSet = true
}

func (v NullableReplyKeyboardMarkup) IsSet() bool {
	return v.isSet
}

func (v *NullableReplyKeyboardMarkup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplyKeyboardMarkup(val *ReplyKeyboardMarkup) *NullableReplyKeyboardMarkup {
	return &NullableReplyKeyboardMarkup{value: val, isSet: true}
}

func (v NullableReplyKeyboardMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplyKeyboardMarkup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


