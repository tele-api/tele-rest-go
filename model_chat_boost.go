/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the ChatBoost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatBoost{}

// ChatBoost This object contains information about a chat boost.
type ChatBoost struct {
	// Unique identifier of the boost
	BoostId string `json:"boost_id"`
	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int32 `json:"add_date"`
	// Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	ExpirationDate int32 `json:"expiration_date"`
	Source ChatBoostSource `json:"source"`
}

type _ChatBoost ChatBoost

// NewChatBoost instantiates a new ChatBoost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatBoost(boostId string, addDate int32, expirationDate int32, source ChatBoostSource) *ChatBoost {
	this := ChatBoost{}
	this.BoostId = boostId
	this.AddDate = addDate
	this.ExpirationDate = expirationDate
	this.Source = source
	return &this
}

// NewChatBoostWithDefaults instantiates a new ChatBoost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatBoostWithDefaults() *ChatBoost {
	this := ChatBoost{}
	return &this
}

// GetBoostId returns the BoostId field value
func (o *ChatBoost) GetBoostId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoostId
}

// GetBoostIdOk returns a tuple with the BoostId field value
// and a boolean to check if the value has been set.
func (o *ChatBoost) GetBoostIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoostId, true
}

// SetBoostId sets field value
func (o *ChatBoost) SetBoostId(v string) {
	o.BoostId = v
}

// GetAddDate returns the AddDate field value
func (o *ChatBoost) GetAddDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AddDate
}

// GetAddDateOk returns a tuple with the AddDate field value
// and a boolean to check if the value has been set.
func (o *ChatBoost) GetAddDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddDate, true
}

// SetAddDate sets field value
func (o *ChatBoost) SetAddDate(v int32) {
	o.AddDate = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *ChatBoost) GetExpirationDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *ChatBoost) GetExpirationDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *ChatBoost) SetExpirationDate(v int32) {
	o.ExpirationDate = v
}

// GetSource returns the Source field value
func (o *ChatBoost) GetSource() ChatBoostSource {
	if o == nil {
		var ret ChatBoostSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ChatBoost) GetSourceOk() (*ChatBoostSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ChatBoost) SetSource(v ChatBoostSource) {
	o.Source = v
}

func (o ChatBoost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatBoost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boost_id"] = o.BoostId
	toSerialize["add_date"] = o.AddDate
	toSerialize["expiration_date"] = o.ExpirationDate
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *ChatBoost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boost_id",
		"add_date",
		"expiration_date",
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

	varChatBoost := _ChatBoost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatBoost)

	if err != nil {
		return err
	}

	*o = ChatBoost(varChatBoost)

	return err
}

type NullableChatBoost struct {
	value *ChatBoost
	isSet bool
}

func (v NullableChatBoost) Get() *ChatBoost {
	return v.value
}

func (v *NullableChatBoost) Set(val *ChatBoost) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBoost) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBoost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBoost(val *ChatBoost) *NullableChatBoost {
	return &NullableChatBoost{value: val, isSet: true}
}

func (v NullableChatBoost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBoost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


