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

// checks if the DirectMessagePriceChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectMessagePriceChanged{}

// DirectMessagePriceChanged Describes a service message about a change in the price of direct messages sent to a channel chat.
type DirectMessagePriceChanged struct {
	// *True*, if direct messages are enabled for the channel chat; false otherwise
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	// *Optional*. The new number of Telegram Stars that must be paid by users for each direct message sent to the channel. Does not apply to users who have been exempted by administrators. Defaults to 0.
	DirectMessageStarCount *int32 `json:"direct_message_star_count,omitempty"`
}

type _DirectMessagePriceChanged DirectMessagePriceChanged

// NewDirectMessagePriceChanged instantiates a new DirectMessagePriceChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectMessagePriceChanged(areDirectMessagesEnabled bool) *DirectMessagePriceChanged {
	this := DirectMessagePriceChanged{}
	this.AreDirectMessagesEnabled = areDirectMessagesEnabled
	var directMessageStarCount int32 = 0
	this.DirectMessageStarCount = &directMessageStarCount
	return &this
}

// NewDirectMessagePriceChangedWithDefaults instantiates a new DirectMessagePriceChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectMessagePriceChangedWithDefaults() *DirectMessagePriceChanged {
	this := DirectMessagePriceChanged{}
	var directMessageStarCount int32 = 0
	this.DirectMessageStarCount = &directMessageStarCount
	return &this
}

// GetAreDirectMessagesEnabled returns the AreDirectMessagesEnabled field value
func (o *DirectMessagePriceChanged) GetAreDirectMessagesEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AreDirectMessagesEnabled
}

// GetAreDirectMessagesEnabledOk returns a tuple with the AreDirectMessagesEnabled field value
// and a boolean to check if the value has been set.
func (o *DirectMessagePriceChanged) GetAreDirectMessagesEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AreDirectMessagesEnabled, true
}

// SetAreDirectMessagesEnabled sets field value
func (o *DirectMessagePriceChanged) SetAreDirectMessagesEnabled(v bool) {
	o.AreDirectMessagesEnabled = v
}

// GetDirectMessageStarCount returns the DirectMessageStarCount field value if set, zero value otherwise.
func (o *DirectMessagePriceChanged) GetDirectMessageStarCount() int32 {
	if o == nil || IsNil(o.DirectMessageStarCount) {
		var ret int32
		return ret
	}
	return *o.DirectMessageStarCount
}

// GetDirectMessageStarCountOk returns a tuple with the DirectMessageStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectMessagePriceChanged) GetDirectMessageStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DirectMessageStarCount) {
		return nil, false
	}
	return o.DirectMessageStarCount, true
}

// HasDirectMessageStarCount returns a boolean if a field has been set.
func (o *DirectMessagePriceChanged) HasDirectMessageStarCount() bool {
	if o != nil && !IsNil(o.DirectMessageStarCount) {
		return true
	}

	return false
}

// SetDirectMessageStarCount gets a reference to the given int32 and assigns it to the DirectMessageStarCount field.
func (o *DirectMessagePriceChanged) SetDirectMessageStarCount(v int32) {
	o.DirectMessageStarCount = &v
}


func (o DirectMessagePriceChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectMessagePriceChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["are_direct_messages_enabled"] = o.AreDirectMessagesEnabled
	if !IsNil(o.DirectMessageStarCount) {
		toSerialize["direct_message_star_count"] = o.DirectMessageStarCount
	}
	return toSerialize, nil
}

func (o *DirectMessagePriceChanged) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"are_direct_messages_enabled",
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

	varDirectMessagePriceChanged := _DirectMessagePriceChanged{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDirectMessagePriceChanged)

	if err != nil {
		return err
	}

	*o = DirectMessagePriceChanged(varDirectMessagePriceChanged)

	return err
}

type NullableDirectMessagePriceChanged struct {
	value *DirectMessagePriceChanged
	isSet bool
}

func (v NullableDirectMessagePriceChanged) Get() *DirectMessagePriceChanged {
	return v.value
}

func (v *NullableDirectMessagePriceChanged) Set(val *DirectMessagePriceChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectMessagePriceChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectMessagePriceChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectMessagePriceChanged(val *DirectMessagePriceChanged) *NullableDirectMessagePriceChanged {
	return &NullableDirectMessagePriceChanged{value: val, isSet: true}
}

func (v NullableDirectMessagePriceChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectMessagePriceChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


