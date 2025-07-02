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

// checks if the PostForwardMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostForwardMessageRequest{}

// PostForwardMessageRequest struct for PostForwardMessageRequest
type PostForwardMessageRequest struct {
	ChatId PostSendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	FromChatId PostForwardMessageRequestFromChatId `json:"from_chat_id"`
	// New start timestamp for the forwarded video in the message
	VideoStartTimestamp *int32 `json:"video_start_timestamp,omitempty"`
	// Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Message identifier in the chat specified in *from\\_chat\\_id*
	MessageId int32 `json:"message_id"`
}

type _PostForwardMessageRequest PostForwardMessageRequest

// NewPostForwardMessageRequest instantiates a new PostForwardMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostForwardMessageRequest(chatId PostSendMessageRequestChatId, fromChatId PostForwardMessageRequestFromChatId, messageId int32) *PostForwardMessageRequest {
	this := PostForwardMessageRequest{}
	this.ChatId = chatId
	this.FromChatId = fromChatId
	this.MessageId = messageId
	return &this
}

// NewPostForwardMessageRequestWithDefaults instantiates a new PostForwardMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostForwardMessageRequestWithDefaults() *PostForwardMessageRequest {
	this := PostForwardMessageRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *PostForwardMessageRequest) GetChatId() PostSendMessageRequestChatId {
	if o == nil {
		var ret PostSendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PostForwardMessageRequest) SetChatId(v PostSendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *PostForwardMessageRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *PostForwardMessageRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *PostForwardMessageRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetFromChatId returns the FromChatId field value
func (o *PostForwardMessageRequest) GetFromChatId() PostForwardMessageRequestFromChatId {
	if o == nil {
		var ret PostForwardMessageRequestFromChatId
		return ret
	}

	return o.FromChatId
}

// GetFromChatIdOk returns a tuple with the FromChatId field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetFromChatIdOk() (*PostForwardMessageRequestFromChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromChatId, true
}

// SetFromChatId sets field value
func (o *PostForwardMessageRequest) SetFromChatId(v PostForwardMessageRequestFromChatId) {
	o.FromChatId = v
}

// GetVideoStartTimestamp returns the VideoStartTimestamp field value if set, zero value otherwise.
func (o *PostForwardMessageRequest) GetVideoStartTimestamp() int32 {
	if o == nil || IsNil(o.VideoStartTimestamp) {
		var ret int32
		return ret
	}
	return *o.VideoStartTimestamp
}

// GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetVideoStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoStartTimestamp) {
		return nil, false
	}
	return o.VideoStartTimestamp, true
}

// HasVideoStartTimestamp returns a boolean if a field has been set.
func (o *PostForwardMessageRequest) HasVideoStartTimestamp() bool {
	if o != nil && !IsNil(o.VideoStartTimestamp) {
		return true
	}

	return false
}

// SetVideoStartTimestamp gets a reference to the given int32 and assigns it to the VideoStartTimestamp field.
func (o *PostForwardMessageRequest) SetVideoStartTimestamp(v int32) {
	o.VideoStartTimestamp = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *PostForwardMessageRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *PostForwardMessageRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *PostForwardMessageRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *PostForwardMessageRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *PostForwardMessageRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *PostForwardMessageRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetMessageId returns the MessageId field value
func (o *PostForwardMessageRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *PostForwardMessageRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *PostForwardMessageRequest) SetMessageId(v int32) {
	o.MessageId = v
}

func (o PostForwardMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostForwardMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["from_chat_id"] = o.FromChatId
	if !IsNil(o.VideoStartTimestamp) {
		toSerialize["video_start_timestamp"] = o.VideoStartTimestamp
	}
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	toSerialize["message_id"] = o.MessageId
	return toSerialize, nil
}

func (o *PostForwardMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"from_chat_id",
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

	varPostForwardMessageRequest := _PostForwardMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostForwardMessageRequest)

	if err != nil {
		return err
	}

	*o = PostForwardMessageRequest(varPostForwardMessageRequest)

	return err
}

type NullablePostForwardMessageRequest struct {
	value *PostForwardMessageRequest
	isSet bool
}

func (v NullablePostForwardMessageRequest) Get() *PostForwardMessageRequest {
	return v.value
}

func (v *NullablePostForwardMessageRequest) Set(val *PostForwardMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostForwardMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostForwardMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostForwardMessageRequest(val *PostForwardMessageRequest) *NullablePostForwardMessageRequest {
	return &NullablePostForwardMessageRequest{value: val, isSet: true}
}

func (v NullablePostForwardMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostForwardMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


