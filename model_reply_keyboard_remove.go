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

// checks if the ReplyKeyboardRemove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplyKeyboardRemove{}

// ReplyKeyboardRemove Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see [ReplyKeyboardMarkup](https://core.telegram.org/bots/api/#replykeyboardmarkup)). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use *one\\_time\\_keyboard* in [ReplyKeyboardMarkup](https://core.telegram.org/bots/api/#replykeyboardmarkup))
	RemoveKeyboard bool `json:"remove_keyboard"`
	// *Optional*. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.    *Example:* A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective *bool `json:"selective,omitempty"`
}

type _ReplyKeyboardRemove ReplyKeyboardRemove

// NewReplyKeyboardRemove instantiates a new ReplyKeyboardRemove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplyKeyboardRemove(removeKeyboard bool) *ReplyKeyboardRemove {
	this := ReplyKeyboardRemove{}
	this.RemoveKeyboard = removeKeyboard
	return &this
}

// NewReplyKeyboardRemoveWithDefaults instantiates a new ReplyKeyboardRemove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplyKeyboardRemoveWithDefaults() *ReplyKeyboardRemove {
	this := ReplyKeyboardRemove{}
	var removeKeyboard bool = true
	this.RemoveKeyboard = removeKeyboard
	return &this
}

// GetRemoveKeyboard returns the RemoveKeyboard field value
func (o *ReplyKeyboardRemove) GetRemoveKeyboard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoveKeyboard
}

// GetRemoveKeyboardOk returns a tuple with the RemoveKeyboard field value
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardRemove) GetRemoveKeyboardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveKeyboard, true
}

// SetRemoveKeyboard sets field value
func (o *ReplyKeyboardRemove) SetRemoveKeyboard(v bool) {
	o.RemoveKeyboard = v
}

// GetSelective returns the Selective field value if set, zero value otherwise.
func (o *ReplyKeyboardRemove) GetSelective() bool {
	if o == nil || IsNil(o.Selective) {
		var ret bool
		return ret
	}
	return *o.Selective
}

// GetSelectiveOk returns a tuple with the Selective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyKeyboardRemove) GetSelectiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Selective) {
		return nil, false
	}
	return o.Selective, true
}

// HasSelective returns a boolean if a field has been set.
func (o *ReplyKeyboardRemove) HasSelective() bool {
	if o != nil && !IsNil(o.Selective) {
		return true
	}

	return false
}

// SetSelective gets a reference to the given bool and assigns it to the Selective field.
func (o *ReplyKeyboardRemove) SetSelective(v bool) {
	o.Selective = &v
}


func (o ReplyKeyboardRemove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplyKeyboardRemove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remove_keyboard"] = o.RemoveKeyboard
	if !IsNil(o.Selective) {
		toSerialize["selective"] = o.Selective
	}
	return toSerialize, nil
}

func (o *ReplyKeyboardRemove) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"remove_keyboard",
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

	varReplyKeyboardRemove := _ReplyKeyboardRemove{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplyKeyboardRemove)

	if err != nil {
		return err
	}

	*o = ReplyKeyboardRemove(varReplyKeyboardRemove)

	return err
}

type NullableReplyKeyboardRemove struct {
	value *ReplyKeyboardRemove
	isSet bool
}

func (v NullableReplyKeyboardRemove) Get() *ReplyKeyboardRemove {
	return v.value
}

func (v *NullableReplyKeyboardRemove) Set(val *ReplyKeyboardRemove) {
	v.value = val
	v.isSet = true
}

func (v NullableReplyKeyboardRemove) IsSet() bool {
	return v.isSet
}

func (v *NullableReplyKeyboardRemove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplyKeyboardRemove(val *ReplyKeyboardRemove) *NullableReplyKeyboardRemove {
	return &NullableReplyKeyboardRemove{value: val, isSet: true}
}

func (v NullableReplyKeyboardRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplyKeyboardRemove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


