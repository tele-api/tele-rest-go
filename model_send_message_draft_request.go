/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.3.0
 *    * - **Modified**: 2026-01-01T02:06:16.405655294Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// checks if the SendMessageDraftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendMessageDraftRequest{}

// SendMessageDraftRequest Request parameters for sendMessageDraft
type SendMessageDraftRequest struct {
	// Unique identifier for the target private chat
	ChatId int32 `json:"chat_id"`
	// Unique identifier for the target message thread
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	// Unique identifier of the message draft; must be non-zero. Changes of drafts with the same identifier are animated
	DraftId int32 `json:"draft_id"`
	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`
	// Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of *parse\\_mode*
	Entities []MessageEntity `json:"entities,omitempty"`
}

type _SendMessageDraftRequest SendMessageDraftRequest

// NewSendMessageDraftRequest instantiates a new SendMessageDraftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessageDraftRequest(chatId int32, draftId int32, text string) *SendMessageDraftRequest {
	this := SendMessageDraftRequest{}
	this.ChatId = chatId
	this.DraftId = draftId
	this.Text = text
	return &this
}

// NewSendMessageDraftRequestWithDefaults instantiates a new SendMessageDraftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessageDraftRequestWithDefaults() *SendMessageDraftRequest {
	this := SendMessageDraftRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *SendMessageDraftRequest) GetChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendMessageDraftRequest) SetChatId(v int32) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendMessageDraftRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendMessageDraftRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendMessageDraftRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetDraftId returns the DraftId field value
func (o *SendMessageDraftRequest) GetDraftId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DraftId
}

// GetDraftIdOk returns a tuple with the DraftId field value
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetDraftIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftId, true
}

// SetDraftId sets field value
func (o *SendMessageDraftRequest) SetDraftId(v int32) {
	o.DraftId = v
}

// GetText returns the Text field value
func (o *SendMessageDraftRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *SendMessageDraftRequest) SetText(v string) {
	o.Text = v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *SendMessageDraftRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *SendMessageDraftRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *SendMessageDraftRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *SendMessageDraftRequest) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageDraftRequest) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *SendMessageDraftRequest) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *SendMessageDraftRequest) SetEntities(v []MessageEntity) {
	o.Entities = v
}


func (o SendMessageDraftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendMessageDraftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["draft_id"] = o.DraftId
	toSerialize["text"] = o.Text
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	return toSerialize, nil
}

func (o *SendMessageDraftRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"draft_id",
		"text",
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

	varSendMessageDraftRequest := _SendMessageDraftRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendMessageDraftRequest)

	if err != nil {
		return err
	}

	*o = SendMessageDraftRequest(varSendMessageDraftRequest)

	return err
}

type NullableSendMessageDraftRequest struct {
	value *SendMessageDraftRequest
	isSet bool
}

func (v NullableSendMessageDraftRequest) Get() *SendMessageDraftRequest {
	return v.value
}

func (v *NullableSendMessageDraftRequest) Set(val *SendMessageDraftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessageDraftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessageDraftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessageDraftRequest(val *SendMessageDraftRequest) *NullableSendMessageDraftRequest {
	return &NullableSendMessageDraftRequest{value: val, isSet: true}
}

func (v NullableSendMessageDraftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessageDraftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

