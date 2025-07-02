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

// checks if the ChatBoostRemoved type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatBoostRemoved{}

// ChatBoostRemoved This object represents a boost removed from a chat.
type ChatBoostRemoved struct {
	Chat Chat `json:"chat"`
	// Unique identifier of the boost
	BoostId string `json:"boost_id"`
	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int32 `json:"remove_date"`
	Source ChatBoostSource `json:"source"`
}

type _ChatBoostRemoved ChatBoostRemoved

// NewChatBoostRemoved instantiates a new ChatBoostRemoved object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatBoostRemoved(chat Chat, boostId string, removeDate int32, source ChatBoostSource) *ChatBoostRemoved {
	this := ChatBoostRemoved{}
	this.Chat = chat
	this.BoostId = boostId
	this.RemoveDate = removeDate
	this.Source = source
	return &this
}

// NewChatBoostRemovedWithDefaults instantiates a new ChatBoostRemoved object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatBoostRemovedWithDefaults() *ChatBoostRemoved {
	this := ChatBoostRemoved{}
	return &this
}

// GetChat returns the Chat field value
func (o *ChatBoostRemoved) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *ChatBoostRemoved) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *ChatBoostRemoved) SetChat(v Chat) {
	o.Chat = v
}

// GetBoostId returns the BoostId field value
func (o *ChatBoostRemoved) GetBoostId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoostId
}

// GetBoostIdOk returns a tuple with the BoostId field value
// and a boolean to check if the value has been set.
func (o *ChatBoostRemoved) GetBoostIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoostId, true
}

// SetBoostId sets field value
func (o *ChatBoostRemoved) SetBoostId(v string) {
	o.BoostId = v
}

// GetRemoveDate returns the RemoveDate field value
func (o *ChatBoostRemoved) GetRemoveDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value
// and a boolean to check if the value has been set.
func (o *ChatBoostRemoved) GetRemoveDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveDate, true
}

// SetRemoveDate sets field value
func (o *ChatBoostRemoved) SetRemoveDate(v int32) {
	o.RemoveDate = v
}

// GetSource returns the Source field value
func (o *ChatBoostRemoved) GetSource() ChatBoostSource {
	if o == nil {
		var ret ChatBoostSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ChatBoostRemoved) GetSourceOk() (*ChatBoostSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ChatBoostRemoved) SetSource(v ChatBoostSource) {
	o.Source = v
}

func (o ChatBoostRemoved) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatBoostRemoved) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["boost_id"] = o.BoostId
	toSerialize["remove_date"] = o.RemoveDate
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *ChatBoostRemoved) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"boost_id",
		"remove_date",
		"source",
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

	varChatBoostRemoved := _ChatBoostRemoved{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatBoostRemoved)

	if err != nil {
		return err
	}

	*o = ChatBoostRemoved(varChatBoostRemoved)

	return err
}

type NullableChatBoostRemoved struct {
	value *ChatBoostRemoved
	isSet bool
}

func (v NullableChatBoostRemoved) Get() *ChatBoostRemoved {
	return v.value
}

func (v *NullableChatBoostRemoved) Set(val *ChatBoostRemoved) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBoostRemoved) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBoostRemoved) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBoostRemoved(val *ChatBoostRemoved) *NullableChatBoostRemoved {
	return &NullableChatBoostRemoved{value: val, isSet: true}
}

func (v NullableChatBoostRemoved) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBoostRemoved) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


