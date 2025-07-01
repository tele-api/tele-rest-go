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

// checks if the MessageReactionUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionUpdated{}

// MessageReactionUpdated This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	Chat Chat `json:"chat"`
	// Unique identifier of the message inside the chat
	MessageId int32 `json:"message_id"`
	User *User `json:"user,omitempty"`
	ActorChat *Chat `json:"actor_chat,omitempty"`
	// Date of the change in Unix time
	Date int32 `json:"date"`
	// Previous list of reaction types that were set by the user
	OldReaction []ReactionType `json:"old_reaction"`
	// New list of reaction types that have been set by the user
	NewReaction []ReactionType `json:"new_reaction"`
}

type _MessageReactionUpdated MessageReactionUpdated

// NewMessageReactionUpdated instantiates a new MessageReactionUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionUpdated(chat Chat, messageId int32, date int32, oldReaction []ReactionType, newReaction []ReactionType) *MessageReactionUpdated {
	this := MessageReactionUpdated{}
	this.Chat = chat
	this.MessageId = messageId
	this.Date = date
	this.OldReaction = oldReaction
	this.NewReaction = newReaction
	return &this
}

// NewMessageReactionUpdatedWithDefaults instantiates a new MessageReactionUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionUpdatedWithDefaults() *MessageReactionUpdated {
	this := MessageReactionUpdated{}
	return &this
}

// GetChat returns the Chat field value
func (o *MessageReactionUpdated) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *MessageReactionUpdated) SetChat(v Chat) {
	o.Chat = v
}

// GetMessageId returns the MessageId field value
func (o *MessageReactionUpdated) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *MessageReactionUpdated) SetMessageId(v int32) {
	o.MessageId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MessageReactionUpdated) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MessageReactionUpdated) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *MessageReactionUpdated) SetUser(v User) {
	o.User = &v
}


// GetActorChat returns the ActorChat field value if set, zero value otherwise.
func (o *MessageReactionUpdated) GetActorChat() Chat {
	if o == nil || IsNil(o.ActorChat) {
		var ret Chat
		return ret
	}
	return *o.ActorChat
}

// GetActorChatOk returns a tuple with the ActorChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetActorChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.ActorChat) {
		return nil, false
	}
	return o.ActorChat, true
}

// HasActorChat returns a boolean if a field has been set.
func (o *MessageReactionUpdated) HasActorChat() bool {
	if o != nil && !IsNil(o.ActorChat) {
		return true
	}

	return false
}

// SetActorChat gets a reference to the given Chat and assigns it to the ActorChat field.
func (o *MessageReactionUpdated) SetActorChat(v Chat) {
	o.ActorChat = &v
}


// GetDate returns the Date field value
func (o *MessageReactionUpdated) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MessageReactionUpdated) SetDate(v int32) {
	o.Date = v
}

// GetOldReaction returns the OldReaction field value
func (o *MessageReactionUpdated) GetOldReaction() []ReactionType {
	if o == nil {
		var ret []ReactionType
		return ret
	}

	return o.OldReaction
}

// GetOldReactionOk returns a tuple with the OldReaction field value
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetOldReactionOk() ([]ReactionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldReaction, true
}

// SetOldReaction sets field value
func (o *MessageReactionUpdated) SetOldReaction(v []ReactionType) {
	o.OldReaction = v
}

// GetNewReaction returns the NewReaction field value
func (o *MessageReactionUpdated) GetNewReaction() []ReactionType {
	if o == nil {
		var ret []ReactionType
		return ret
	}

	return o.NewReaction
}

// GetNewReactionOk returns a tuple with the NewReaction field value
// and a boolean to check if the value has been set.
func (o *MessageReactionUpdated) GetNewReactionOk() ([]ReactionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewReaction, true
}

// SetNewReaction sets field value
func (o *MessageReactionUpdated) SetNewReaction(v []ReactionType) {
	o.NewReaction = v
}

func (o MessageReactionUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.ActorChat) {
		toSerialize["actor_chat"] = o.ActorChat
	}
	toSerialize["date"] = o.Date
	toSerialize["old_reaction"] = o.OldReaction
	toSerialize["new_reaction"] = o.NewReaction
	return toSerialize, nil
}

func (o *MessageReactionUpdated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"message_id",
		"date",
		"old_reaction",
		"new_reaction",
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

	varMessageReactionUpdated := _MessageReactionUpdated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReactionUpdated)

	if err != nil {
		return err
	}

	*o = MessageReactionUpdated(varMessageReactionUpdated)

	return err
}

type NullableMessageReactionUpdated struct {
	value *MessageReactionUpdated
	isSet bool
}

func (v NullableMessageReactionUpdated) Get() *MessageReactionUpdated {
	return v.value
}

func (v *NullableMessageReactionUpdated) Set(val *MessageReactionUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionUpdated(val *MessageReactionUpdated) *NullableMessageReactionUpdated {
	return &NullableMessageReactionUpdated{value: val, isSet: true}
}

func (v NullableMessageReactionUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


