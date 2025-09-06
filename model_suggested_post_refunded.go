/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-09-06T05:32:05.180087835Z[Etc/UTC]
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

// checks if the SuggestedPostRefunded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostRefunded{}

// SuggestedPostRefunded Describes a service message about a payment refund for a suggested post.
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	// Reason for the refund. Currently, one of “post\\_deleted” if the post was deleted within 24 hours of being posted or removed from scheduled messages without being posted, or “payment\\_refunded” if the payer refunded their payment.
	Reason string `json:"reason"`
}

type _SuggestedPostRefunded SuggestedPostRefunded

// NewSuggestedPostRefunded instantiates a new SuggestedPostRefunded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostRefunded(reason string) *SuggestedPostRefunded {
	this := SuggestedPostRefunded{}
	this.Reason = reason
	return &this
}

// NewSuggestedPostRefundedWithDefaults instantiates a new SuggestedPostRefunded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostRefundedWithDefaults() *SuggestedPostRefunded {
	this := SuggestedPostRefunded{}
	return &this
}

// GetSuggestedPostMessage returns the SuggestedPostMessage field value if set, zero value otherwise.
func (o *SuggestedPostRefunded) GetSuggestedPostMessage() Message {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		var ret Message
		return ret
	}
	return *o.SuggestedPostMessage
}

// GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostRefunded) GetSuggestedPostMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		return nil, false
	}
	return o.SuggestedPostMessage, true
}

// HasSuggestedPostMessage returns a boolean if a field has been set.
func (o *SuggestedPostRefunded) HasSuggestedPostMessage() bool {
	if o != nil && !IsNil(o.SuggestedPostMessage) {
		return true
	}

	return false
}

// SetSuggestedPostMessage gets a reference to the given Message and assigns it to the SuggestedPostMessage field.
func (o *SuggestedPostRefunded) SetSuggestedPostMessage(v Message) {
	o.SuggestedPostMessage = &v
}


// GetReason returns the Reason field value
func (o *SuggestedPostRefunded) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SuggestedPostRefunded) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SuggestedPostRefunded) SetReason(v string) {
	o.Reason = v
}

func (o SuggestedPostRefunded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostRefunded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedPostMessage) {
		toSerialize["suggested_post_message"] = o.SuggestedPostMessage
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *SuggestedPostRefunded) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
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

	varSuggestedPostRefunded := _SuggestedPostRefunded{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestedPostRefunded)

	if err != nil {
		return err
	}

	*o = SuggestedPostRefunded(varSuggestedPostRefunded)

	return err
}

type NullableSuggestedPostRefunded struct {
	value *SuggestedPostRefunded
	isSet bool
}

func (v NullableSuggestedPostRefunded) Get() *SuggestedPostRefunded {
	return v.value
}

func (v *NullableSuggestedPostRefunded) Set(val *SuggestedPostRefunded) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostRefunded) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostRefunded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostRefunded(val *SuggestedPostRefunded) *NullableSuggestedPostRefunded {
	return &NullableSuggestedPostRefunded{value: val, isSet: true}
}

func (v NullableSuggestedPostRefunded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostRefunded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


