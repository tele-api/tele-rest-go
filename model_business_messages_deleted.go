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

// checks if the BusinessMessagesDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessMessagesDeleted{}

// BusinessMessagesDeleted This object is received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	Chat Chat `json:"chat"`
	// The list of identifiers of deleted messages in the chat of the business account
	MessageIds []int32 `json:"message_ids"`
}

type _BusinessMessagesDeleted BusinessMessagesDeleted

// NewBusinessMessagesDeleted instantiates a new BusinessMessagesDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessMessagesDeleted(businessConnectionId string, chat Chat, messageIds []int32) *BusinessMessagesDeleted {
	this := BusinessMessagesDeleted{}
	this.BusinessConnectionId = businessConnectionId
	this.Chat = chat
	this.MessageIds = messageIds
	return &this
}

// NewBusinessMessagesDeletedWithDefaults instantiates a new BusinessMessagesDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessMessagesDeletedWithDefaults() *BusinessMessagesDeleted {
	this := BusinessMessagesDeleted{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *BusinessMessagesDeleted) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *BusinessMessagesDeleted) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *BusinessMessagesDeleted) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetChat returns the Chat field value
func (o *BusinessMessagesDeleted) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *BusinessMessagesDeleted) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *BusinessMessagesDeleted) SetChat(v Chat) {
	o.Chat = v
}

// GetMessageIds returns the MessageIds field value
func (o *BusinessMessagesDeleted) GetMessageIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIds
}

// GetMessageIdsOk returns a tuple with the MessageIds field value
// and a boolean to check if the value has been set.
func (o *BusinessMessagesDeleted) GetMessageIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIds, true
}

// SetMessageIds sets field value
func (o *BusinessMessagesDeleted) SetMessageIds(v []int32) {
	o.MessageIds = v
}

func (o BusinessMessagesDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessMessagesDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["chat"] = o.Chat
	toSerialize["message_ids"] = o.MessageIds
	return toSerialize, nil
}

func (o *BusinessMessagesDeleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"chat",
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

	varBusinessMessagesDeleted := _BusinessMessagesDeleted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBusinessMessagesDeleted)

	if err != nil {
		return err
	}

	*o = BusinessMessagesDeleted(varBusinessMessagesDeleted)

	return err
}

type NullableBusinessMessagesDeleted struct {
	value *BusinessMessagesDeleted
	isSet bool
}

func (v NullableBusinessMessagesDeleted) Get() *BusinessMessagesDeleted {
	return v.value
}

func (v *NullableBusinessMessagesDeleted) Set(val *BusinessMessagesDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessMessagesDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessMessagesDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessMessagesDeleted(val *BusinessMessagesDeleted) *NullableBusinessMessagesDeleted {
	return &NullableBusinessMessagesDeleted{value: val, isSet: true}
}

func (v NullableBusinessMessagesDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessMessagesDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


