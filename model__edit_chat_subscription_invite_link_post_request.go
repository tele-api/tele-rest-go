/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the EditChatSubscriptionInviteLinkPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditChatSubscriptionInviteLinkPostRequest{}

// EditChatSubscriptionInviteLinkPostRequest struct for EditChatSubscriptionInviteLinkPostRequest
type EditChatSubscriptionInviteLinkPostRequest struct {
	ChatId SendMessagePostRequestChatId `json:"chat_id"`
	// The invite link to edit
	InviteLink string `json:"invite_link"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
}

type _EditChatSubscriptionInviteLinkPostRequest EditChatSubscriptionInviteLinkPostRequest

// NewEditChatSubscriptionInviteLinkPostRequest instantiates a new EditChatSubscriptionInviteLinkPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditChatSubscriptionInviteLinkPostRequest(chatId SendMessagePostRequestChatId, inviteLink string) *EditChatSubscriptionInviteLinkPostRequest {
	this := EditChatSubscriptionInviteLinkPostRequest{}
	this.ChatId = chatId
	this.InviteLink = inviteLink
	return &this
}

// NewEditChatSubscriptionInviteLinkPostRequestWithDefaults instantiates a new EditChatSubscriptionInviteLinkPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditChatSubscriptionInviteLinkPostRequestWithDefaults() *EditChatSubscriptionInviteLinkPostRequest {
	this := EditChatSubscriptionInviteLinkPostRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *EditChatSubscriptionInviteLinkPostRequest) GetChatId() SendMessagePostRequestChatId {
	if o == nil {
		var ret SendMessagePostRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *EditChatSubscriptionInviteLinkPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *EditChatSubscriptionInviteLinkPostRequest) SetChatId(v SendMessagePostRequestChatId) {
	o.ChatId = v
}

// GetInviteLink returns the InviteLink field value
func (o *EditChatSubscriptionInviteLinkPostRequest) GetInviteLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value
// and a boolean to check if the value has been set.
func (o *EditChatSubscriptionInviteLinkPostRequest) GetInviteLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteLink, true
}

// SetInviteLink sets field value
func (o *EditChatSubscriptionInviteLinkPostRequest) SetInviteLink(v string) {
	o.InviteLink = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditChatSubscriptionInviteLinkPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditChatSubscriptionInviteLinkPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditChatSubscriptionInviteLinkPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditChatSubscriptionInviteLinkPostRequest) SetName(v string) {
	o.Name = &v
}


func (o EditChatSubscriptionInviteLinkPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditChatSubscriptionInviteLinkPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["invite_link"] = o.InviteLink
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *EditChatSubscriptionInviteLinkPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"invite_link",
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

	varEditChatSubscriptionInviteLinkPostRequest := _EditChatSubscriptionInviteLinkPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditChatSubscriptionInviteLinkPostRequest)

	if err != nil {
		return err
	}

	*o = EditChatSubscriptionInviteLinkPostRequest(varEditChatSubscriptionInviteLinkPostRequest)

	return err
}

type NullableEditChatSubscriptionInviteLinkPostRequest struct {
	value *EditChatSubscriptionInviteLinkPostRequest
	isSet bool
}

func (v NullableEditChatSubscriptionInviteLinkPostRequest) Get() *EditChatSubscriptionInviteLinkPostRequest {
	return v.value
}

func (v *NullableEditChatSubscriptionInviteLinkPostRequest) Set(val *EditChatSubscriptionInviteLinkPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditChatSubscriptionInviteLinkPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditChatSubscriptionInviteLinkPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditChatSubscriptionInviteLinkPostRequest(val *EditChatSubscriptionInviteLinkPostRequest) *NullableEditChatSubscriptionInviteLinkPostRequest {
	return &NullableEditChatSubscriptionInviteLinkPostRequest{value: val, isSet: true}
}

func (v NullableEditChatSubscriptionInviteLinkPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditChatSubscriptionInviteLinkPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


