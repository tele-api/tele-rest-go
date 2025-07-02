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

// checks if the PostEditChatSubscriptionInviteLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostEditChatSubscriptionInviteLinkRequest{}

// PostEditChatSubscriptionInviteLinkRequest struct for PostEditChatSubscriptionInviteLinkRequest
type PostEditChatSubscriptionInviteLinkRequest struct {
	ChatId PostSendMessageRequestChatId `json:"chat_id"`
	// The invite link to edit
	InviteLink string `json:"invite_link"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
}

type _PostEditChatSubscriptionInviteLinkRequest PostEditChatSubscriptionInviteLinkRequest

// NewPostEditChatSubscriptionInviteLinkRequest instantiates a new PostEditChatSubscriptionInviteLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostEditChatSubscriptionInviteLinkRequest(chatId PostSendMessageRequestChatId, inviteLink string) *PostEditChatSubscriptionInviteLinkRequest {
	this := PostEditChatSubscriptionInviteLinkRequest{}
	this.ChatId = chatId
	this.InviteLink = inviteLink
	return &this
}

// NewPostEditChatSubscriptionInviteLinkRequestWithDefaults instantiates a new PostEditChatSubscriptionInviteLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostEditChatSubscriptionInviteLinkRequestWithDefaults() *PostEditChatSubscriptionInviteLinkRequest {
	this := PostEditChatSubscriptionInviteLinkRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *PostEditChatSubscriptionInviteLinkRequest) GetChatId() PostSendMessageRequestChatId {
	if o == nil {
		var ret PostSendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PostEditChatSubscriptionInviteLinkRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PostEditChatSubscriptionInviteLinkRequest) SetChatId(v PostSendMessageRequestChatId) {
	o.ChatId = v
}

// GetInviteLink returns the InviteLink field value
func (o *PostEditChatSubscriptionInviteLinkRequest) GetInviteLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value
// and a boolean to check if the value has been set.
func (o *PostEditChatSubscriptionInviteLinkRequest) GetInviteLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteLink, true
}

// SetInviteLink sets field value
func (o *PostEditChatSubscriptionInviteLinkRequest) SetInviteLink(v string) {
	o.InviteLink = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostEditChatSubscriptionInviteLinkRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEditChatSubscriptionInviteLinkRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostEditChatSubscriptionInviteLinkRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostEditChatSubscriptionInviteLinkRequest) SetName(v string) {
	o.Name = &v
}


func (o PostEditChatSubscriptionInviteLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEditChatSubscriptionInviteLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["invite_link"] = o.InviteLink
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *PostEditChatSubscriptionInviteLinkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostEditChatSubscriptionInviteLinkRequest := _PostEditChatSubscriptionInviteLinkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEditChatSubscriptionInviteLinkRequest)

	if err != nil {
		return err
	}

	*o = PostEditChatSubscriptionInviteLinkRequest(varPostEditChatSubscriptionInviteLinkRequest)

	return err
}

type NullablePostEditChatSubscriptionInviteLinkRequest struct {
	value *PostEditChatSubscriptionInviteLinkRequest
	isSet bool
}

func (v NullablePostEditChatSubscriptionInviteLinkRequest) Get() *PostEditChatSubscriptionInviteLinkRequest {
	return v.value
}

func (v *NullablePostEditChatSubscriptionInviteLinkRequest) Set(val *PostEditChatSubscriptionInviteLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEditChatSubscriptionInviteLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEditChatSubscriptionInviteLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEditChatSubscriptionInviteLinkRequest(val *PostEditChatSubscriptionInviteLinkRequest) *NullablePostEditChatSubscriptionInviteLinkRequest {
	return &NullablePostEditChatSubscriptionInviteLinkRequest{value: val, isSet: true}
}

func (v NullablePostEditChatSubscriptionInviteLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEditChatSubscriptionInviteLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


