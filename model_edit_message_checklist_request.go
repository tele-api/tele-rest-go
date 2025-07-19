/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the EditMessageChecklistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditMessageChecklistRequest{}

// EditMessageChecklistRequest Request parameters for editMessageChecklist
type EditMessageChecklistRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier for the target chat
	ChatId int32 `json:"chat_id"`
	// Unique identifier for the target message
	MessageId int32 `json:"message_id"`
	Checklist InputChecklist `json:"checklist"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type _EditMessageChecklistRequest EditMessageChecklistRequest

// NewEditMessageChecklistRequest instantiates a new EditMessageChecklistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMessageChecklistRequest(businessConnectionId string, chatId int32, messageId int32, checklist InputChecklist) *EditMessageChecklistRequest {
	this := EditMessageChecklistRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.ChatId = chatId
	this.MessageId = messageId
	this.Checklist = checklist
	return &this
}

// NewEditMessageChecklistRequestWithDefaults instantiates a new EditMessageChecklistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMessageChecklistRequestWithDefaults() *EditMessageChecklistRequest {
	this := EditMessageChecklistRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *EditMessageChecklistRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *EditMessageChecklistRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *EditMessageChecklistRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetChatId returns the ChatId field value
func (o *EditMessageChecklistRequest) GetChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *EditMessageChecklistRequest) GetChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *EditMessageChecklistRequest) SetChatId(v int32) {
	o.ChatId = v
}

// GetMessageId returns the MessageId field value
func (o *EditMessageChecklistRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *EditMessageChecklistRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *EditMessageChecklistRequest) SetMessageId(v int32) {
	o.MessageId = v
}

// GetChecklist returns the Checklist field value
func (o *EditMessageChecklistRequest) GetChecklist() InputChecklist {
	if o == nil {
		var ret InputChecklist
		return ret
	}

	return o.Checklist
}

// GetChecklistOk returns a tuple with the Checklist field value
// and a boolean to check if the value has been set.
func (o *EditMessageChecklistRequest) GetChecklistOk() (*InputChecklist, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checklist, true
}

// SetChecklist sets field value
func (o *EditMessageChecklistRequest) SetChecklist(v InputChecklist) {
	o.Checklist = v
}

// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *EditMessageChecklistRequest) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageChecklistRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *EditMessageChecklistRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *EditMessageChecklistRequest) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o EditMessageChecklistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditMessageChecklistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["chat_id"] = o.ChatId
	toSerialize["message_id"] = o.MessageId
	toSerialize["checklist"] = o.Checklist
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *EditMessageChecklistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"chat_id",
		"message_id",
		"checklist",
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

	varEditMessageChecklistRequest := _EditMessageChecklistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditMessageChecklistRequest)

	if err != nil {
		return err
	}

	*o = EditMessageChecklistRequest(varEditMessageChecklistRequest)

	return err
}

type NullableEditMessageChecklistRequest struct {
	value *EditMessageChecklistRequest
	isSet bool
}

func (v NullableEditMessageChecklistRequest) Get() *EditMessageChecklistRequest {
	return v.value
}

func (v *NullableEditMessageChecklistRequest) Set(val *EditMessageChecklistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageChecklistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageChecklistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageChecklistRequest(val *EditMessageChecklistRequest) *NullableEditMessageChecklistRequest {
	return &NullableEditMessageChecklistRequest{value: val, isSet: true}
}

func (v NullableEditMessageChecklistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageChecklistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


