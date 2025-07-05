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

// checks if the PinChatMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinChatMessageRequest{}

// PinChatMessageRequest Request parameters for pinChatMessage
type PinChatMessageRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be pinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Identifier of a message to pin
	MessageId int32 `json:"message_id"`
	// Pass *True* if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification *bool `json:"disable_notification,omitempty"`
}

type _PinChatMessageRequest PinChatMessageRequest

// NewPinChatMessageRequest instantiates a new PinChatMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinChatMessageRequest(chatId SendMessageRequestChatId, messageId int32) *PinChatMessageRequest {
	this := PinChatMessageRequest{}
	this.ChatId = chatId
	this.MessageId = messageId
	return &this
}

// NewPinChatMessageRequestWithDefaults instantiates a new PinChatMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinChatMessageRequestWithDefaults() *PinChatMessageRequest {
	this := PinChatMessageRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *PinChatMessageRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinChatMessageRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *PinChatMessageRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *PinChatMessageRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *PinChatMessageRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PinChatMessageRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PinChatMessageRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageId returns the MessageId field value
func (o *PinChatMessageRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *PinChatMessageRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *PinChatMessageRequest) SetMessageId(v int32) {
	o.MessageId = v
}

// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *PinChatMessageRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinChatMessageRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *PinChatMessageRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *PinChatMessageRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


func (o PinChatMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinChatMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	return toSerialize, nil
}

func (o *PinChatMessageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPinChatMessageRequest := _PinChatMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPinChatMessageRequest)

	if err != nil {
		return err
	}

	*o = PinChatMessageRequest(varPinChatMessageRequest)

	return err
}

type NullablePinChatMessageRequest struct {
	value *PinChatMessageRequest
	isSet bool
}

func (v NullablePinChatMessageRequest) Get() *PinChatMessageRequest {
	return v.value
}

func (v *NullablePinChatMessageRequest) Set(val *PinChatMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePinChatMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePinChatMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinChatMessageRequest(val *PinChatMessageRequest) *NullablePinChatMessageRequest {
	return &NullablePinChatMessageRequest{value: val, isSet: true}
}

func (v NullablePinChatMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinChatMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


