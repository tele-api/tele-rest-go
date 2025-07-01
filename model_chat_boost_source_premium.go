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

// checks if the ChatBoostSourcePremium type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatBoostSourcePremium{}

// ChatBoostSourcePremium The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	// Source of the boost, always “premium”
	Source string `json:"source"`
	User User `json:"user"`
}

type _ChatBoostSourcePremium ChatBoostSourcePremium

// NewChatBoostSourcePremium instantiates a new ChatBoostSourcePremium object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatBoostSourcePremium(source string, user User) *ChatBoostSourcePremium {
	this := ChatBoostSourcePremium{}
	this.Source = source
	this.User = user
	return &this
}

// NewChatBoostSourcePremiumWithDefaults instantiates a new ChatBoostSourcePremium object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatBoostSourcePremiumWithDefaults() *ChatBoostSourcePremium {
	this := ChatBoostSourcePremium{}
	var source string = "premium"
	this.Source = source
	return &this
}

// GetSource returns the Source field value
func (o *ChatBoostSourcePremium) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ChatBoostSourcePremium) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ChatBoostSourcePremium) SetSource(v string) {
	o.Source = v
}

// GetUser returns the User field value
func (o *ChatBoostSourcePremium) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ChatBoostSourcePremium) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ChatBoostSourcePremium) SetUser(v User) {
	o.User = v
}

func (o ChatBoostSourcePremium) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatBoostSourcePremium) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *ChatBoostSourcePremium) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"user",
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

	varChatBoostSourcePremium := _ChatBoostSourcePremium{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatBoostSourcePremium)

	if err != nil {
		return err
	}

	*o = ChatBoostSourcePremium(varChatBoostSourcePremium)

	return err
}

type NullableChatBoostSourcePremium struct {
	value *ChatBoostSourcePremium
	isSet bool
}

func (v NullableChatBoostSourcePremium) Get() *ChatBoostSourcePremium {
	return v.value
}

func (v *NullableChatBoostSourcePremium) Set(val *ChatBoostSourcePremium) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBoostSourcePremium) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBoostSourcePremium) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBoostSourcePremium(val *ChatBoostSourcePremium) *NullableChatBoostSourcePremium {
	return &NullableChatBoostSourcePremium{value: val, isSet: true}
}

func (v NullableChatBoostSourcePremium) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBoostSourcePremium) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


