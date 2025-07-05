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

// checks if the MessageOriginUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageOriginUser{}

// MessageOriginUser The message was originally sent by a known user.
type MessageOriginUser struct {
	// Type of the message origin, always “user”
	Type string `json:"type"`
	// Date the message was sent originally in Unix time
	Date int32 `json:"date"`
	SenderUser User `json:"sender_user"`
}

type _MessageOriginUser MessageOriginUser

// NewMessageOriginUser instantiates a new MessageOriginUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageOriginUser(type_ string, date int32, senderUser User) *MessageOriginUser {
	this := MessageOriginUser{}
	this.Type = type_
	this.Date = date
	this.SenderUser = senderUser
	return &this
}

// NewMessageOriginUserWithDefaults instantiates a new MessageOriginUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageOriginUserWithDefaults() *MessageOriginUser {
	this := MessageOriginUser{}
	var type_ string = "user"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *MessageOriginUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageOriginUser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageOriginUser) SetType(v string) {
	o.Type = v
}

// GetDate returns the Date field value
func (o *MessageOriginUser) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MessageOriginUser) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MessageOriginUser) SetDate(v int32) {
	o.Date = v
}

// GetSenderUser returns the SenderUser field value
func (o *MessageOriginUser) GetSenderUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.SenderUser
}

// GetSenderUserOk returns a tuple with the SenderUser field value
// and a boolean to check if the value has been set.
func (o *MessageOriginUser) GetSenderUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderUser, true
}

// SetSenderUser sets field value
func (o *MessageOriginUser) SetSenderUser(v User) {
	o.SenderUser = v
}

func (o MessageOriginUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageOriginUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["date"] = o.Date
	toSerialize["sender_user"] = o.SenderUser
	return toSerialize, nil
}

func (o *MessageOriginUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"date",
		"sender_user",
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

	varMessageOriginUser := _MessageOriginUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageOriginUser)

	if err != nil {
		return err
	}

	*o = MessageOriginUser(varMessageOriginUser)

	return err
}

type NullableMessageOriginUser struct {
	value *MessageOriginUser
	isSet bool
}

func (v NullableMessageOriginUser) Get() *MessageOriginUser {
	return v.value
}

func (v *NullableMessageOriginUser) Set(val *MessageOriginUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageOriginUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageOriginUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageOriginUser(val *MessageOriginUser) *NullableMessageOriginUser {
	return &NullableMessageOriginUser{value: val, isSet: true}
}

func (v NullableMessageOriginUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageOriginUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


