/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the ForceReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForceReply{}

// ForceReply Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice [privacy mode](https://core.telegram.org/bots/features#privacy-mode). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`
	// *Optional*. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	// *Optional*. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
	Selective *bool `json:"selective,omitempty"`
}

type _ForceReply ForceReply

// NewForceReply instantiates a new ForceReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForceReply(forceReply bool) *ForceReply {
	this := ForceReply{}
	this.ForceReply = forceReply
	return &this
}

// NewForceReplyWithDefaults instantiates a new ForceReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForceReplyWithDefaults() *ForceReply {
	this := ForceReply{}
	var forceReply bool = true
	this.ForceReply = forceReply
	return &this
}

// GetForceReply returns the ForceReply field value
func (o *ForceReply) GetForceReply() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceReply
}

// GetForceReplyOk returns a tuple with the ForceReply field value
// and a boolean to check if the value has been set.
func (o *ForceReply) GetForceReplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceReply, true
}

// SetForceReply sets field value
func (o *ForceReply) SetForceReply(v bool) {
	o.ForceReply = v
}

// GetInputFieldPlaceholder returns the InputFieldPlaceholder field value if set, zero value otherwise.
func (o *ForceReply) GetInputFieldPlaceholder() string {
	if o == nil || IsNil(o.InputFieldPlaceholder) {
		var ret string
		return ret
	}
	return *o.InputFieldPlaceholder
}

// GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForceReply) GetInputFieldPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.InputFieldPlaceholder) {
		return nil, false
	}
	return o.InputFieldPlaceholder, true
}

// HasInputFieldPlaceholder returns a boolean if a field has been set.
func (o *ForceReply) HasInputFieldPlaceholder() bool {
	if o != nil && !IsNil(o.InputFieldPlaceholder) {
		return true
	}

	return false
}

// SetInputFieldPlaceholder gets a reference to the given string and assigns it to the InputFieldPlaceholder field.
func (o *ForceReply) SetInputFieldPlaceholder(v string) {
	o.InputFieldPlaceholder = &v
}


// GetSelective returns the Selective field value if set, zero value otherwise.
func (o *ForceReply) GetSelective() bool {
	if o == nil || IsNil(o.Selective) {
		var ret bool
		return ret
	}
	return *o.Selective
}

// GetSelectiveOk returns a tuple with the Selective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForceReply) GetSelectiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Selective) {
		return nil, false
	}
	return o.Selective, true
}

// HasSelective returns a boolean if a field has been set.
func (o *ForceReply) HasSelective() bool {
	if o != nil && !IsNil(o.Selective) {
		return true
	}

	return false
}

// SetSelective gets a reference to the given bool and assigns it to the Selective field.
func (o *ForceReply) SetSelective(v bool) {
	o.Selective = &v
}


func (o ForceReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForceReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["force_reply"] = o.ForceReply
	if !IsNil(o.InputFieldPlaceholder) {
		toSerialize["input_field_placeholder"] = o.InputFieldPlaceholder
	}
	if !IsNil(o.Selective) {
		toSerialize["selective"] = o.Selective
	}
	return toSerialize, nil
}

func (o *ForceReply) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"force_reply",
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

	varForceReply := _ForceReply{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForceReply)

	if err != nil {
		return err
	}

	*o = ForceReply(varForceReply)

	return err
}

type NullableForceReply struct {
	value *ForceReply
	isSet bool
}

func (v NullableForceReply) Get() *ForceReply {
	return v.value
}

func (v *NullableForceReply) Set(val *ForceReply) {
	v.value = val
	v.isSet = true
}

func (v NullableForceReply) IsSet() bool {
	return v.isSet
}

func (v *NullableForceReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForceReply(val *ForceReply) *NullableForceReply {
	return &NullableForceReply{value: val, isSet: true}
}

func (v NullableForceReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForceReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


