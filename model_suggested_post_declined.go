/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-09-09T23:46:30.762261440Z[Etc/UTC]
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
)

// checks if the SuggestedPostDeclined type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostDeclined{}

// SuggestedPostDeclined Describes a service message about the rejection of a suggested post.
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	// *Optional*. Comment with which the post was declined
	Comment *string `json:"comment,omitempty"`
}

// NewSuggestedPostDeclined instantiates a new SuggestedPostDeclined object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostDeclined() *SuggestedPostDeclined {
	this := SuggestedPostDeclined{}
	return &this
}

// NewSuggestedPostDeclinedWithDefaults instantiates a new SuggestedPostDeclined object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostDeclinedWithDefaults() *SuggestedPostDeclined {
	this := SuggestedPostDeclined{}
	return &this
}

// GetSuggestedPostMessage returns the SuggestedPostMessage field value if set, zero value otherwise.
func (o *SuggestedPostDeclined) GetSuggestedPostMessage() Message {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		var ret Message
		return ret
	}
	return *o.SuggestedPostMessage
}

// GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostDeclined) GetSuggestedPostMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		return nil, false
	}
	return o.SuggestedPostMessage, true
}

// HasSuggestedPostMessage returns a boolean if a field has been set.
func (o *SuggestedPostDeclined) HasSuggestedPostMessage() bool {
	if o != nil && !IsNil(o.SuggestedPostMessage) {
		return true
	}

	return false
}

// SetSuggestedPostMessage gets a reference to the given Message and assigns it to the SuggestedPostMessage field.
func (o *SuggestedPostDeclined) SetSuggestedPostMessage(v Message) {
	o.SuggestedPostMessage = &v
}


// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SuggestedPostDeclined) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostDeclined) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SuggestedPostDeclined) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SuggestedPostDeclined) SetComment(v string) {
	o.Comment = &v
}


func (o SuggestedPostDeclined) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostDeclined) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedPostMessage) {
		toSerialize["suggested_post_message"] = o.SuggestedPostMessage
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableSuggestedPostDeclined struct {
	value *SuggestedPostDeclined
	isSet bool
}

func (v NullableSuggestedPostDeclined) Get() *SuggestedPostDeclined {
	return v.value
}

func (v *NullableSuggestedPostDeclined) Set(val *SuggestedPostDeclined) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostDeclined) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostDeclined) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostDeclined(val *SuggestedPostDeclined) *NullableSuggestedPostDeclined {
	return &NullableSuggestedPostDeclined{value: val, isSet: true}
}

func (v NullableSuggestedPostDeclined) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostDeclined) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


