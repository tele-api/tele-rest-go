/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the PostForwardMessagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostForwardMessagesRequest{}

// PostForwardMessagesRequest struct for PostForwardMessagesRequest
type PostForwardMessagesRequest struct {
	ChatId PostSendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	FromChatId PostForwardMessagesRequestFromChatId `json:"from_chat_id"`
	// A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to forward. The identifiers must be specified in a strictly increasing order.
	MessageIds []int32 `json:"message_ids"`
	// Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
}

type _PostForwardMessagesRequest PostForwardMessagesRequest

// NewPostForwardMessagesRequest instantiates a new PostForwardMessagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostForwardMessagesRequest(chatId PostSendMessageRequestChatId, fromChatId PostForwardMessagesRequestFromChatId, messageIds []int32) *PostForwardMessagesRequest {
	this := PostForwardMessagesRequest{}
	this.ChatId = chatId
	this.FromChatId = fromChatId
	this.MessageIds = messageIds
	return &this
}

// NewPostForwardMessagesRequestWithDefaults instantiates a new PostForwardMessagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostForwardMessagesRequestWithDefaults() *PostForwardMessagesRequest {
	this := PostForwardMessagesRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *PostForwardMessagesRequest) GetChatId() PostSendMessageRequestChatId {
	if o == nil {
		var ret PostSendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PostForwardMessagesRequest) SetChatId(v PostSendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *PostForwardMessagesRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *PostForwardMessagesRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *PostForwardMessagesRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetFromChatId returns the FromChatId field value
func (o *PostForwardMessagesRequest) GetFromChatId() PostForwardMessagesRequestFromChatId {
	if o == nil {
		var ret PostForwardMessagesRequestFromChatId
		return ret
	}

	return o.FromChatId
}

// GetFromChatIdOk returns a tuple with the FromChatId field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetFromChatIdOk() (*PostForwardMessagesRequestFromChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromChatId, true
}

// SetFromChatId sets field value
func (o *PostForwardMessagesRequest) SetFromChatId(v PostForwardMessagesRequestFromChatId) {
	o.FromChatId = v
}

// GetMessageIds returns the MessageIds field value
func (o *PostForwardMessagesRequest) GetMessageIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIds
}

// GetMessageIdsOk returns a tuple with the MessageIds field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetMessageIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIds, true
}

// SetMessageIds sets field value
func (o *PostForwardMessagesRequest) SetMessageIds(v []int32) {
	o.MessageIds = v
}

// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *PostForwardMessagesRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *PostForwardMessagesRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *PostForwardMessagesRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *PostForwardMessagesRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessagesRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *PostForwardMessagesRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *PostForwardMessagesRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


func (o PostForwardMessagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostForwardMessagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["from_chat_id"] = o.FromChatId
	toSerialize["message_ids"] = o.MessageIds
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	return toSerialize, nil
}

func (o *PostForwardMessagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"from_chat_id",
		"message_ids",
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

	varPostForwardMessagesRequest := _PostForwardMessagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostForwardMessagesRequest)

	if err != nil {
		return err
	}

	*o = PostForwardMessagesRequest(varPostForwardMessagesRequest)

	return err
}

type NullablePostForwardMessagesRequest struct {
	value *PostForwardMessagesRequest
	isSet bool
}

func (v NullablePostForwardMessagesRequest) Get() *PostForwardMessagesRequest {
	return v.value
}

func (v *NullablePostForwardMessagesRequest) Set(val *PostForwardMessagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostForwardMessagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostForwardMessagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostForwardMessagesRequest(val *PostForwardMessagesRequest) *NullablePostForwardMessagesRequest {
	return &NullablePostForwardMessagesRequest{value: val, isSet: true}
}

func (v NullablePostForwardMessagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostForwardMessagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


