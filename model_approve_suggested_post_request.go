/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-08-17T02:10:51.183543732Z[Etc/UTC]
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

// checks if the ApproveSuggestedPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveSuggestedPostRequest{}

// ApproveSuggestedPostRequest Request parameters for approveSuggestedPost
type ApproveSuggestedPostRequest struct {
	// Unique identifier for the target direct messages chat
	ChatId int32 `json:"chat_id"`
	// Identifier of a suggested post message to approve
	MessageId int32 `json:"message_id"`
	// Point in time (Unix timestamp) when the post is expected to be published; omit if the date has already been specified when the suggested post was created. If specified, then the date must be not more than 2678400 seconds (30 days) in the future
	SendDate *int32 `json:"send_date,omitempty"`
}

type _ApproveSuggestedPostRequest ApproveSuggestedPostRequest

// NewApproveSuggestedPostRequest instantiates a new ApproveSuggestedPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveSuggestedPostRequest(chatId int32, messageId int32) *ApproveSuggestedPostRequest {
	this := ApproveSuggestedPostRequest{}
	this.ChatId = chatId
	this.MessageId = messageId
	return &this
}

// NewApproveSuggestedPostRequestWithDefaults instantiates a new ApproveSuggestedPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveSuggestedPostRequestWithDefaults() *ApproveSuggestedPostRequest {
	this := ApproveSuggestedPostRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *ApproveSuggestedPostRequest) GetChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *ApproveSuggestedPostRequest) GetChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *ApproveSuggestedPostRequest) SetChatId(v int32) {
	o.ChatId = v
}

// GetMessageId returns the MessageId field value
func (o *ApproveSuggestedPostRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *ApproveSuggestedPostRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *ApproveSuggestedPostRequest) SetMessageId(v int32) {
	o.MessageId = v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *ApproveSuggestedPostRequest) GetSendDate() int32 {
	if o == nil || IsNil(o.SendDate) {
		var ret int32
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveSuggestedPostRequest) GetSendDateOk() (*int32, bool) {
	if o == nil || IsNil(o.SendDate) {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *ApproveSuggestedPostRequest) HasSendDate() bool {
	if o != nil && !IsNil(o.SendDate) {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given int32 and assigns it to the SendDate field.
func (o *ApproveSuggestedPostRequest) SetSendDate(v int32) {
	o.SendDate = &v
}


func (o ApproveSuggestedPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproveSuggestedPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.SendDate) {
		toSerialize["send_date"] = o.SendDate
	}
	return toSerialize, nil
}

func (o *ApproveSuggestedPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varApproveSuggestedPostRequest := _ApproveSuggestedPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApproveSuggestedPostRequest)

	if err != nil {
		return err
	}

	*o = ApproveSuggestedPostRequest(varApproveSuggestedPostRequest)

	return err
}

type NullableApproveSuggestedPostRequest struct {
	value *ApproveSuggestedPostRequest
	isSet bool
}

func (v NullableApproveSuggestedPostRequest) Get() *ApproveSuggestedPostRequest {
	return v.value
}

func (v *NullableApproveSuggestedPostRequest) Set(val *ApproveSuggestedPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveSuggestedPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveSuggestedPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveSuggestedPostRequest(val *ApproveSuggestedPostRequest) *NullableApproveSuggestedPostRequest {
	return &NullableApproveSuggestedPostRequest{value: val, isSet: true}
}

func (v NullableApproveSuggestedPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveSuggestedPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


