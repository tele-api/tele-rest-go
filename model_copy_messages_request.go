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

// checks if the CopyMessagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyMessagesRequest{}

// CopyMessagesRequest Request parameters for copyMessages
type CopyMessagesRequest struct {
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	FromChatId ForwardMessagesRequestFromChatId `json:"from_chat_id"`
	// A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to copy. The identifiers must be specified in a strictly increasing order.
	MessageIds []int32 `json:"message_ids"`
	// Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass *True* to copy the messages without their captions
	RemoveCaption *bool `json:"remove_caption,omitempty"`
}

type _CopyMessagesRequest CopyMessagesRequest

// NewCopyMessagesRequest instantiates a new CopyMessagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyMessagesRequest(chatId SendMessageRequestChatId, fromChatId ForwardMessagesRequestFromChatId, messageIds []int32) *CopyMessagesRequest {
	this := CopyMessagesRequest{}
	this.ChatId = chatId
	this.FromChatId = fromChatId
	this.MessageIds = messageIds
	return &this
}

// NewCopyMessagesRequestWithDefaults instantiates a new CopyMessagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyMessagesRequestWithDefaults() *CopyMessagesRequest {
	this := CopyMessagesRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *CopyMessagesRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *CopyMessagesRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *CopyMessagesRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *CopyMessagesRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *CopyMessagesRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetFromChatId returns the FromChatId field value
func (o *CopyMessagesRequest) GetFromChatId() ForwardMessagesRequestFromChatId {
	if o == nil {
		var ret ForwardMessagesRequestFromChatId
		return ret
	}

	return o.FromChatId
}

// GetFromChatIdOk returns a tuple with the FromChatId field value
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetFromChatIdOk() (*ForwardMessagesRequestFromChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromChatId, true
}

// SetFromChatId sets field value
func (o *CopyMessagesRequest) SetFromChatId(v ForwardMessagesRequestFromChatId) {
	o.FromChatId = v
}

// GetMessageIds returns the MessageIds field value
func (o *CopyMessagesRequest) GetMessageIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIds
}

// GetMessageIdsOk returns a tuple with the MessageIds field value
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetMessageIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIds, true
}

// SetMessageIds sets field value
func (o *CopyMessagesRequest) SetMessageIds(v []int32) {
	o.MessageIds = v
}

// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *CopyMessagesRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *CopyMessagesRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *CopyMessagesRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *CopyMessagesRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *CopyMessagesRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *CopyMessagesRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetRemoveCaption returns the RemoveCaption field value if set, zero value otherwise.
func (o *CopyMessagesRequest) GetRemoveCaption() bool {
	if o == nil || IsNil(o.RemoveCaption) {
		var ret bool
		return ret
	}
	return *o.RemoveCaption
}

// GetRemoveCaptionOk returns a tuple with the RemoveCaption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessagesRequest) GetRemoveCaptionOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveCaption) {
		return nil, false
	}
	return o.RemoveCaption, true
}

// HasRemoveCaption returns a boolean if a field has been set.
func (o *CopyMessagesRequest) HasRemoveCaption() bool {
	if o != nil && !IsNil(o.RemoveCaption) {
		return true
	}

	return false
}

// SetRemoveCaption gets a reference to the given bool and assigns it to the RemoveCaption field.
func (o *CopyMessagesRequest) SetRemoveCaption(v bool) {
	o.RemoveCaption = &v
}


func (o CopyMessagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyMessagesRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RemoveCaption) {
		toSerialize["remove_caption"] = o.RemoveCaption
	}
	return toSerialize, nil
}

func (o *CopyMessagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCopyMessagesRequest := _CopyMessagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCopyMessagesRequest)

	if err != nil {
		return err
	}

	*o = CopyMessagesRequest(varCopyMessagesRequest)

	return err
}

type NullableCopyMessagesRequest struct {
	value *CopyMessagesRequest
	isSet bool
}

func (v NullableCopyMessagesRequest) Get() *CopyMessagesRequest {
	return v.value
}

func (v *NullableCopyMessagesRequest) Set(val *CopyMessagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyMessagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyMessagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyMessagesRequest(val *CopyMessagesRequest) *NullableCopyMessagesRequest {
	return &NullableCopyMessagesRequest{value: val, isSet: true}
}

func (v NullableCopyMessagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyMessagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


