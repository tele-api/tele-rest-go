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

// checks if the PostBanChatSenderChatRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBanChatSenderChatRequest{}

// PostBanChatSenderChatRequest struct for PostBanChatSenderChatRequest
type PostBanChatSenderChatRequest struct {
	ChatId PostSendMessageRequestChatId `json:"chat_id"`
	// Unique identifier of the target sender chat
	SenderChatId int32 `json:"sender_chat_id"`
}

type _PostBanChatSenderChatRequest PostBanChatSenderChatRequest

// NewPostBanChatSenderChatRequest instantiates a new PostBanChatSenderChatRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBanChatSenderChatRequest(chatId PostSendMessageRequestChatId, senderChatId int32) *PostBanChatSenderChatRequest {
	this := PostBanChatSenderChatRequest{}
	this.ChatId = chatId
	this.SenderChatId = senderChatId
	return &this
}

// NewPostBanChatSenderChatRequestWithDefaults instantiates a new PostBanChatSenderChatRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBanChatSenderChatRequestWithDefaults() *PostBanChatSenderChatRequest {
	this := PostBanChatSenderChatRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *PostBanChatSenderChatRequest) GetChatId() PostSendMessageRequestChatId {
	if o == nil {
		var ret PostSendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PostBanChatSenderChatRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PostBanChatSenderChatRequest) SetChatId(v PostSendMessageRequestChatId) {
	o.ChatId = v
}

// GetSenderChatId returns the SenderChatId field value
func (o *PostBanChatSenderChatRequest) GetSenderChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SenderChatId
}

// GetSenderChatIdOk returns a tuple with the SenderChatId field value
// and a boolean to check if the value has been set.
func (o *PostBanChatSenderChatRequest) GetSenderChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderChatId, true
}

// SetSenderChatId sets field value
func (o *PostBanChatSenderChatRequest) SetSenderChatId(v int32) {
	o.SenderChatId = v
}

func (o PostBanChatSenderChatRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBanChatSenderChatRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["sender_chat_id"] = o.SenderChatId
	return toSerialize, nil
}

func (o *PostBanChatSenderChatRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"sender_chat_id",
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

	varPostBanChatSenderChatRequest := _PostBanChatSenderChatRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostBanChatSenderChatRequest)

	if err != nil {
		return err
	}

	*o = PostBanChatSenderChatRequest(varPostBanChatSenderChatRequest)

	return err
}

type NullablePostBanChatSenderChatRequest struct {
	value *PostBanChatSenderChatRequest
	isSet bool
}

func (v NullablePostBanChatSenderChatRequest) Get() *PostBanChatSenderChatRequest {
	return v.value
}

func (v *NullablePostBanChatSenderChatRequest) Set(val *PostBanChatSenderChatRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBanChatSenderChatRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBanChatSenderChatRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBanChatSenderChatRequest(val *PostBanChatSenderChatRequest) *NullablePostBanChatSenderChatRequest {
	return &NullablePostBanChatSenderChatRequest{value: val, isSet: true}
}

func (v NullablePostBanChatSenderChatRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBanChatSenderChatRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


