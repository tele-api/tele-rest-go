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

// checks if the SuggestedPostApproved type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostApproved{}

// SuggestedPostApproved Describes a service message about the approval of a suggested post.
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Price *SuggestedPostPrice `json:"price,omitempty"`
	// Date when the post will be published
	SendDate int32 `json:"send_date"`
}

type _SuggestedPostApproved SuggestedPostApproved

// NewSuggestedPostApproved instantiates a new SuggestedPostApproved object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostApproved(sendDate int32) *SuggestedPostApproved {
	this := SuggestedPostApproved{}
	this.SendDate = sendDate
	return &this
}

// NewSuggestedPostApprovedWithDefaults instantiates a new SuggestedPostApproved object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostApprovedWithDefaults() *SuggestedPostApproved {
	this := SuggestedPostApproved{}
	return &this
}

// GetSuggestedPostMessage returns the SuggestedPostMessage field value if set, zero value otherwise.
func (o *SuggestedPostApproved) GetSuggestedPostMessage() Message {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		var ret Message
		return ret
	}
	return *o.SuggestedPostMessage
}

// GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostApproved) GetSuggestedPostMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		return nil, false
	}
	return o.SuggestedPostMessage, true
}

// HasSuggestedPostMessage returns a boolean if a field has been set.
func (o *SuggestedPostApproved) HasSuggestedPostMessage() bool {
	if o != nil && !IsNil(o.SuggestedPostMessage) {
		return true
	}

	return false
}

// SetSuggestedPostMessage gets a reference to the given Message and assigns it to the SuggestedPostMessage field.
func (o *SuggestedPostApproved) SetSuggestedPostMessage(v Message) {
	o.SuggestedPostMessage = &v
}


// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SuggestedPostApproved) GetPrice() SuggestedPostPrice {
	if o == nil || IsNil(o.Price) {
		var ret SuggestedPostPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostApproved) GetPriceOk() (*SuggestedPostPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SuggestedPostApproved) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given SuggestedPostPrice and assigns it to the Price field.
func (o *SuggestedPostApproved) SetPrice(v SuggestedPostPrice) {
	o.Price = &v
}


// GetSendDate returns the SendDate field value
func (o *SuggestedPostApproved) GetSendDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value
// and a boolean to check if the value has been set.
func (o *SuggestedPostApproved) GetSendDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendDate, true
}

// SetSendDate sets field value
func (o *SuggestedPostApproved) SetSendDate(v int32) {
	o.SendDate = v
}

func (o SuggestedPostApproved) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostApproved) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedPostMessage) {
		toSerialize["suggested_post_message"] = o.SuggestedPostMessage
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	toSerialize["send_date"] = o.SendDate
	return toSerialize, nil
}

func (o *SuggestedPostApproved) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"send_date",
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

	varSuggestedPostApproved := _SuggestedPostApproved{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestedPostApproved)

	if err != nil {
		return err
	}

	*o = SuggestedPostApproved(varSuggestedPostApproved)

	return err
}

type NullableSuggestedPostApproved struct {
	value *SuggestedPostApproved
	isSet bool
}

func (v NullableSuggestedPostApproved) Get() *SuggestedPostApproved {
	return v.value
}

func (v *NullableSuggestedPostApproved) Set(val *SuggestedPostApproved) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostApproved) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostApproved) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostApproved(val *SuggestedPostApproved) *NullableSuggestedPostApproved {
	return &NullableSuggestedPostApproved{value: val, isSet: true}
}

func (v NullableSuggestedPostApproved) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostApproved) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


