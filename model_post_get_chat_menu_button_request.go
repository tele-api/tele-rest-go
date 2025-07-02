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
)

// checks if the PostGetChatMenuButtonRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGetChatMenuButtonRequest{}

// PostGetChatMenuButtonRequest struct for PostGetChatMenuButtonRequest
type PostGetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
	ChatId *int32 `json:"chat_id,omitempty"`
}

// NewPostGetChatMenuButtonRequest instantiates a new PostGetChatMenuButtonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGetChatMenuButtonRequest() *PostGetChatMenuButtonRequest {
	this := PostGetChatMenuButtonRequest{}
	return &this
}

// NewPostGetChatMenuButtonRequestWithDefaults instantiates a new PostGetChatMenuButtonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGetChatMenuButtonRequestWithDefaults() *PostGetChatMenuButtonRequest {
	this := PostGetChatMenuButtonRequest{}
	return &this
}

// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *PostGetChatMenuButtonRequest) GetChatId() int32 {
	if o == nil || IsNil(o.ChatId) {
		var ret int32
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGetChatMenuButtonRequest) GetChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *PostGetChatMenuButtonRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given int32 and assigns it to the ChatId field.
func (o *PostGetChatMenuButtonRequest) SetChatId(v int32) {
	o.ChatId = &v
}


func (o PostGetChatMenuButtonRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGetChatMenuButtonRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	return toSerialize, nil
}

type NullablePostGetChatMenuButtonRequest struct {
	value *PostGetChatMenuButtonRequest
	isSet bool
}

func (v NullablePostGetChatMenuButtonRequest) Get() *PostGetChatMenuButtonRequest {
	return v.value
}

func (v *NullablePostGetChatMenuButtonRequest) Set(val *PostGetChatMenuButtonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGetChatMenuButtonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGetChatMenuButtonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGetChatMenuButtonRequest(val *PostGetChatMenuButtonRequest) *NullablePostGetChatMenuButtonRequest {
	return &NullablePostGetChatMenuButtonRequest{value: val, isSet: true}
}

func (v NullablePostGetChatMenuButtonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGetChatMenuButtonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


