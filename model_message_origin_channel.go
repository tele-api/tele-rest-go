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

// checks if the MessageOriginChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageOriginChannel{}

// MessageOriginChannel The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	// Type of the message origin, always “channel”
	Type string `json:"type"`
	// Date the message was sent originally in Unix time
	Date int32 `json:"date"`
	Chat Chat `json:"chat"`
	// Unique message identifier inside the chat
	MessageId int32 `json:"message_id"`
	// *Optional*. Signature of the original post author
	AuthorSignature *string `json:"author_signature,omitempty"`
}

type _MessageOriginChannel MessageOriginChannel

// NewMessageOriginChannel instantiates a new MessageOriginChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageOriginChannel(type_ string, date int32, chat Chat, messageId int32) *MessageOriginChannel {
	this := MessageOriginChannel{}
	this.Type = type_
	this.Date = date
	this.Chat = chat
	this.MessageId = messageId
	return &this
}

// NewMessageOriginChannelWithDefaults instantiates a new MessageOriginChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageOriginChannelWithDefaults() *MessageOriginChannel {
	this := MessageOriginChannel{}
	var type_ string = "channel"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *MessageOriginChannel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageOriginChannel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageOriginChannel) SetType(v string) {
	o.Type = v
}

// GetDate returns the Date field value
func (o *MessageOriginChannel) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MessageOriginChannel) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MessageOriginChannel) SetDate(v int32) {
	o.Date = v
}

// GetChat returns the Chat field value
func (o *MessageOriginChannel) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *MessageOriginChannel) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *MessageOriginChannel) SetChat(v Chat) {
	o.Chat = v
}

// GetMessageId returns the MessageId field value
func (o *MessageOriginChannel) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *MessageOriginChannel) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *MessageOriginChannel) SetMessageId(v int32) {
	o.MessageId = v
}

// GetAuthorSignature returns the AuthorSignature field value if set, zero value otherwise.
func (o *MessageOriginChannel) GetAuthorSignature() string {
	if o == nil || IsNil(o.AuthorSignature) {
		var ret string
		return ret
	}
	return *o.AuthorSignature
}

// GetAuthorSignatureOk returns a tuple with the AuthorSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageOriginChannel) GetAuthorSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorSignature) {
		return nil, false
	}
	return o.AuthorSignature, true
}

// HasAuthorSignature returns a boolean if a field has been set.
func (o *MessageOriginChannel) HasAuthorSignature() bool {
	if o != nil && !IsNil(o.AuthorSignature) {
		return true
	}

	return false
}

// SetAuthorSignature gets a reference to the given string and assigns it to the AuthorSignature field.
func (o *MessageOriginChannel) SetAuthorSignature(v string) {
	o.AuthorSignature = &v
}


func (o MessageOriginChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageOriginChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["date"] = o.Date
	toSerialize["chat"] = o.Chat
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.AuthorSignature) {
		toSerialize["author_signature"] = o.AuthorSignature
	}
	return toSerialize, nil
}

func (o *MessageOriginChannel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"date",
		"chat",
		"message_id",
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

	varMessageOriginChannel := _MessageOriginChannel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageOriginChannel)

	if err != nil {
		return err
	}

	*o = MessageOriginChannel(varMessageOriginChannel)

	return err
}

type NullableMessageOriginChannel struct {
	value *MessageOriginChannel
	isSet bool
}

func (v NullableMessageOriginChannel) Get() *MessageOriginChannel {
	return v.value
}

func (v *NullableMessageOriginChannel) Set(val *MessageOriginChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageOriginChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageOriginChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageOriginChannel(val *MessageOriginChannel) *NullableMessageOriginChannel {
	return &NullableMessageOriginChannel{value: val, isSet: true}
}

func (v NullableMessageOriginChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageOriginChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


