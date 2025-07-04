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

// checks if the UnpinChatMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnpinChatMessageRequest{}

// UnpinChatMessageRequest Request parameters for unpinChatMessage
type UnpinChatMessageRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be unpinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Identifier of the message to unpin. Required if *business\\_connection\\_id* is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
	MessageId *int32 `json:"message_id,omitempty"`
}

type _UnpinChatMessageRequest UnpinChatMessageRequest

// NewUnpinChatMessageRequest instantiates a new UnpinChatMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnpinChatMessageRequest(chatId SendMessageRequestChatId) *UnpinChatMessageRequest {
	this := UnpinChatMessageRequest{}
	this.ChatId = chatId
	return &this
}

// NewUnpinChatMessageRequestWithDefaults instantiates a new UnpinChatMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnpinChatMessageRequestWithDefaults() *UnpinChatMessageRequest {
	this := UnpinChatMessageRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *UnpinChatMessageRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpinChatMessageRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *UnpinChatMessageRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *UnpinChatMessageRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *UnpinChatMessageRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *UnpinChatMessageRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *UnpinChatMessageRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *UnpinChatMessageRequest) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpinChatMessageRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *UnpinChatMessageRequest) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *UnpinChatMessageRequest) SetMessageId(v int32) {
	o.MessageId = &v
}


func (o UnpinChatMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnpinChatMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	return toSerialize, nil
}

func (o *UnpinChatMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
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

	varUnpinChatMessageRequest := _UnpinChatMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnpinChatMessageRequest)

	if err != nil {
		return err
	}

	*o = UnpinChatMessageRequest(varUnpinChatMessageRequest)

	return err
}

type NullableUnpinChatMessageRequest struct {
	value *UnpinChatMessageRequest
	isSet bool
}

func (v NullableUnpinChatMessageRequest) Get() *UnpinChatMessageRequest {
	return v.value
}

func (v *NullableUnpinChatMessageRequest) Set(val *UnpinChatMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnpinChatMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnpinChatMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnpinChatMessageRequest(val *UnpinChatMessageRequest) *NullableUnpinChatMessageRequest {
	return &NullableUnpinChatMessageRequest{value: val, isSet: true}
}

func (v NullableUnpinChatMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnpinChatMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


