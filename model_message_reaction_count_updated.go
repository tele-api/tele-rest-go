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

// checks if the MessageReactionCountUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionCountUpdated{}

// MessageReactionCountUpdated This object represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	Chat Chat `json:"chat"`
	// Unique message identifier inside the chat
	MessageId int32 `json:"message_id"`
	// Date of the change in Unix time
	Date int32 `json:"date"`
	// List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
}

type _MessageReactionCountUpdated MessageReactionCountUpdated

// NewMessageReactionCountUpdated instantiates a new MessageReactionCountUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionCountUpdated(chat Chat, messageId int32, date int32, reactions []ReactionCount) *MessageReactionCountUpdated {
	this := MessageReactionCountUpdated{}
	this.Chat = chat
	this.MessageId = messageId
	this.Date = date
	this.Reactions = reactions
	return &this
}

// NewMessageReactionCountUpdatedWithDefaults instantiates a new MessageReactionCountUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionCountUpdatedWithDefaults() *MessageReactionCountUpdated {
	this := MessageReactionCountUpdated{}
	return &this
}

// GetChat returns the Chat field value
func (o *MessageReactionCountUpdated) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountUpdated) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *MessageReactionCountUpdated) SetChat(v Chat) {
	o.Chat = v
}

// GetMessageId returns the MessageId field value
func (o *MessageReactionCountUpdated) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountUpdated) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *MessageReactionCountUpdated) SetMessageId(v int32) {
	o.MessageId = v
}

// GetDate returns the Date field value
func (o *MessageReactionCountUpdated) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountUpdated) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MessageReactionCountUpdated) SetDate(v int32) {
	o.Date = v
}

// GetReactions returns the Reactions field value
func (o *MessageReactionCountUpdated) GetReactions() []ReactionCount {
	if o == nil {
		var ret []ReactionCount
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *MessageReactionCountUpdated) GetReactionsOk() ([]ReactionCount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reactions, true
}

// SetReactions sets field value
func (o *MessageReactionCountUpdated) SetReactions(v []ReactionCount) {
	o.Reactions = v
}

func (o MessageReactionCountUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionCountUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["message_id"] = o.MessageId
	toSerialize["date"] = o.Date
	toSerialize["reactions"] = o.Reactions
	return toSerialize, nil
}

func (o *MessageReactionCountUpdated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"message_id",
		"date",
		"reactions",
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

	varMessageReactionCountUpdated := _MessageReactionCountUpdated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReactionCountUpdated)

	if err != nil {
		return err
	}

	*o = MessageReactionCountUpdated(varMessageReactionCountUpdated)

	return err
}

type NullableMessageReactionCountUpdated struct {
	value *MessageReactionCountUpdated
	isSet bool
}

func (v NullableMessageReactionCountUpdated) Get() *MessageReactionCountUpdated {
	return v.value
}

func (v *NullableMessageReactionCountUpdated) Set(val *MessageReactionCountUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionCountUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionCountUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionCountUpdated(val *MessageReactionCountUpdated) *NullableMessageReactionCountUpdated {
	return &NullableMessageReactionCountUpdated{value: val, isSet: true}
}

func (v NullableMessageReactionCountUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionCountUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


