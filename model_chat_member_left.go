/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the ChatMemberLeft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMemberLeft{}

// ChatMemberLeft Represents a [chat member](https://core.telegram.org/bots/api/#chatmember) that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	// The member's status in the chat, always “left”
	Status string `json:"status"`
	User User `json:"user"`
}

type _ChatMemberLeft ChatMemberLeft

// NewChatMemberLeft instantiates a new ChatMemberLeft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMemberLeft(status string, user User) *ChatMemberLeft {
	this := ChatMemberLeft{}
	this.Status = status
	this.User = user
	return &this
}

// NewChatMemberLeftWithDefaults instantiates a new ChatMemberLeft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMemberLeftWithDefaults() *ChatMemberLeft {
	this := ChatMemberLeft{}
	var status string = "left"
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *ChatMemberLeft) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChatMemberLeft) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChatMemberLeft) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value
func (o *ChatMemberLeft) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ChatMemberLeft) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ChatMemberLeft) SetUser(v User) {
	o.User = v
}

func (o ChatMemberLeft) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMemberLeft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *ChatMemberLeft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"user",
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

	varChatMemberLeft := _ChatMemberLeft{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatMemberLeft)

	if err != nil {
		return err
	}

	*o = ChatMemberLeft(varChatMemberLeft)

	return err
}

type NullableChatMemberLeft struct {
	value *ChatMemberLeft
	isSet bool
}

func (v NullableChatMemberLeft) Get() *ChatMemberLeft {
	return v.value
}

func (v *NullableChatMemberLeft) Set(val *ChatMemberLeft) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMemberLeft) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMemberLeft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMemberLeft(val *ChatMemberLeft) *NullableChatMemberLeft {
	return &NullableChatMemberLeft{value: val, isSet: true}
}

func (v NullableChatMemberLeft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMemberLeft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


