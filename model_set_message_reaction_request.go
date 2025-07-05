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

// checks if the SetMessageReactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetMessageReactionRequest{}

// SetMessageReactionRequest Request parameters for setMessageReaction
type SetMessageReactionRequest struct {
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	MessageId int32 `json:"message_id"`
	// A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	Reaction []ReactionType `json:"reaction,omitempty"`
	// Pass *True* to set the reaction with a big animation
	IsBig *bool `json:"is_big,omitempty"`
}

type _SetMessageReactionRequest SetMessageReactionRequest

// NewSetMessageReactionRequest instantiates a new SetMessageReactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMessageReactionRequest(chatId SendMessageRequestChatId, messageId int32) *SetMessageReactionRequest {
	this := SetMessageReactionRequest{}
	this.ChatId = chatId
	this.MessageId = messageId
	return &this
}

// NewSetMessageReactionRequestWithDefaults instantiates a new SetMessageReactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMessageReactionRequestWithDefaults() *SetMessageReactionRequest {
	this := SetMessageReactionRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *SetMessageReactionRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SetMessageReactionRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SetMessageReactionRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageId returns the MessageId field value
func (o *SetMessageReactionRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *SetMessageReactionRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *SetMessageReactionRequest) SetMessageId(v int32) {
	o.MessageId = v
}

// GetReaction returns the Reaction field value if set, zero value otherwise.
func (o *SetMessageReactionRequest) GetReaction() []ReactionType {
	if o == nil || IsNil(o.Reaction) {
		var ret []ReactionType
		return ret
	}
	return o.Reaction
}

// GetReactionOk returns a tuple with the Reaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMessageReactionRequest) GetReactionOk() ([]ReactionType, bool) {
	if o == nil || IsNil(o.Reaction) {
		return nil, false
	}
	return o.Reaction, true
}

// HasReaction returns a boolean if a field has been set.
func (o *SetMessageReactionRequest) HasReaction() bool {
	if o != nil && !IsNil(o.Reaction) {
		return true
	}

	return false
}

// SetReaction gets a reference to the given []ReactionType and assigns it to the Reaction field.
func (o *SetMessageReactionRequest) SetReaction(v []ReactionType) {
	o.Reaction = v
}


// GetIsBig returns the IsBig field value if set, zero value otherwise.
func (o *SetMessageReactionRequest) GetIsBig() bool {
	if o == nil || IsNil(o.IsBig) {
		var ret bool
		return ret
	}
	return *o.IsBig
}

// GetIsBigOk returns a tuple with the IsBig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMessageReactionRequest) GetIsBigOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBig) {
		return nil, false
	}
	return o.IsBig, true
}

// HasIsBig returns a boolean if a field has been set.
func (o *SetMessageReactionRequest) HasIsBig() bool {
	if o != nil && !IsNil(o.IsBig) {
		return true
	}

	return false
}

// SetIsBig gets a reference to the given bool and assigns it to the IsBig field.
func (o *SetMessageReactionRequest) SetIsBig(v bool) {
	o.IsBig = &v
}


func (o SetMessageReactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMessageReactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.Reaction) {
		toSerialize["reaction"] = o.Reaction
	}
	if !IsNil(o.IsBig) {
		toSerialize["is_big"] = o.IsBig
	}
	return toSerialize, nil
}

func (o *SetMessageReactionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
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

	varSetMessageReactionRequest := _SetMessageReactionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetMessageReactionRequest)

	if err != nil {
		return err
	}

	*o = SetMessageReactionRequest(varSetMessageReactionRequest)

	return err
}

type NullableSetMessageReactionRequest struct {
	value *SetMessageReactionRequest
	isSet bool
}

func (v NullableSetMessageReactionRequest) Get() *SetMessageReactionRequest {
	return v.value
}

func (v *NullableSetMessageReactionRequest) Set(val *SetMessageReactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMessageReactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMessageReactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMessageReactionRequest(val *SetMessageReactionRequest) *NullableSetMessageReactionRequest {
	return &NullableSetMessageReactionRequest{value: val, isSet: true}
}

func (v NullableSetMessageReactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMessageReactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


