/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the EditForumTopicRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditForumTopicRequest{}

// EditForumTopicRequest Request parameters for editForumTopic
type EditForumTopicRequest struct {
	ChatId BotCommandScopeChatChatId `json:"chat_id"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int32 `json:"message_thread_id"`
	// New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	Name *string `json:"name,omitempty"`
	// New unique identifier of the custom emoji shown as the topic icon. Use [getForumTopicIconStickers](https://core.telegram.org/bots/api/#getforumtopiciconstickers) to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type _EditForumTopicRequest EditForumTopicRequest

// NewEditForumTopicRequest instantiates a new EditForumTopicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditForumTopicRequest(chatId BotCommandScopeChatChatId, messageThreadId int32) *EditForumTopicRequest {
	this := EditForumTopicRequest{}
	this.ChatId = chatId
	this.MessageThreadId = messageThreadId
	return &this
}

// NewEditForumTopicRequestWithDefaults instantiates a new EditForumTopicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditForumTopicRequestWithDefaults() *EditForumTopicRequest {
	this := EditForumTopicRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *EditForumTopicRequest) GetChatId() BotCommandScopeChatChatId {
	if o == nil {
		var ret BotCommandScopeChatChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *EditForumTopicRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *EditForumTopicRequest) SetChatId(v BotCommandScopeChatChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value
func (o *EditForumTopicRequest) GetMessageThreadId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value
// and a boolean to check if the value has been set.
func (o *EditForumTopicRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageThreadId, true
}

// SetMessageThreadId sets field value
func (o *EditForumTopicRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditForumTopicRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditForumTopicRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditForumTopicRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditForumTopicRequest) SetName(v string) {
	o.Name = &v
}


// GetIconCustomEmojiId returns the IconCustomEmojiId field value if set, zero value otherwise.
func (o *EditForumTopicRequest) GetIconCustomEmojiId() string {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.IconCustomEmojiId
}

// GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditForumTopicRequest) GetIconCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		return nil, false
	}
	return o.IconCustomEmojiId, true
}

// HasIconCustomEmojiId returns a boolean if a field has been set.
func (o *EditForumTopicRequest) HasIconCustomEmojiId() bool {
	if o != nil && !IsNil(o.IconCustomEmojiId) {
		return true
	}

	return false
}

// SetIconCustomEmojiId gets a reference to the given string and assigns it to the IconCustomEmojiId field.
func (o *EditForumTopicRequest) SetIconCustomEmojiId(v string) {
	o.IconCustomEmojiId = &v
}


func (o EditForumTopicRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditForumTopicRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["message_thread_id"] = o.MessageThreadId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IconCustomEmojiId) {
		toSerialize["icon_custom_emoji_id"] = o.IconCustomEmojiId
	}
	return toSerialize, nil
}

func (o *EditForumTopicRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"message_thread_id",
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

	varEditForumTopicRequest := _EditForumTopicRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditForumTopicRequest)

	if err != nil {
		return err
	}

	*o = EditForumTopicRequest(varEditForumTopicRequest)

	return err
}

type NullableEditForumTopicRequest struct {
	value *EditForumTopicRequest
	isSet bool
}

func (v NullableEditForumTopicRequest) Get() *EditForumTopicRequest {
	return v.value
}

func (v *NullableEditForumTopicRequest) Set(val *EditForumTopicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditForumTopicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditForumTopicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditForumTopicRequest(val *EditForumTopicRequest) *NullableEditForumTopicRequest {
	return &NullableEditForumTopicRequest{value: val, isSet: true}
}

func (v NullableEditForumTopicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditForumTopicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


