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

// checks if the EditChatInviteLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditChatInviteLinkRequest{}

// EditChatInviteLinkRequest Request parameters for editChatInviteLink
type EditChatInviteLinkRequest struct {
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// The invite link to edit
	InviteLink string `json:"invite_link"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	// Point in time (Unix timestamp) when the link will expire
	ExpireDate *int32 `json:"expire_date,omitempty"`
	// The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int32 `json:"member_limit,omitempty"`
	// *True*, if users joining the chat via the link need to be approved by chat administrators. If *True*, *member\\_limit* can't be specified
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty"`
}

type _EditChatInviteLinkRequest EditChatInviteLinkRequest

// NewEditChatInviteLinkRequest instantiates a new EditChatInviteLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditChatInviteLinkRequest(chatId SendMessageRequestChatId, inviteLink string) *EditChatInviteLinkRequest {
	this := EditChatInviteLinkRequest{}
	this.ChatId = chatId
	this.InviteLink = inviteLink
	return &this
}

// NewEditChatInviteLinkRequestWithDefaults instantiates a new EditChatInviteLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditChatInviteLinkRequestWithDefaults() *EditChatInviteLinkRequest {
	this := EditChatInviteLinkRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *EditChatInviteLinkRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *EditChatInviteLinkRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetInviteLink returns the InviteLink field value
func (o *EditChatInviteLinkRequest) GetInviteLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetInviteLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteLink, true
}

// SetInviteLink sets field value
func (o *EditChatInviteLinkRequest) SetInviteLink(v string) {
	o.InviteLink = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditChatInviteLinkRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditChatInviteLinkRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditChatInviteLinkRequest) SetName(v string) {
	o.Name = &v
}


// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *EditChatInviteLinkRequest) GetExpireDate() int32 {
	if o == nil || IsNil(o.ExpireDate) {
		var ret int32
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetExpireDateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireDate) {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *EditChatInviteLinkRequest) HasExpireDate() bool {
	if o != nil && !IsNil(o.ExpireDate) {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given int32 and assigns it to the ExpireDate field.
func (o *EditChatInviteLinkRequest) SetExpireDate(v int32) {
	o.ExpireDate = &v
}


// GetMemberLimit returns the MemberLimit field value if set, zero value otherwise.
func (o *EditChatInviteLinkRequest) GetMemberLimit() int32 {
	if o == nil || IsNil(o.MemberLimit) {
		var ret int32
		return ret
	}
	return *o.MemberLimit
}

// GetMemberLimitOk returns a tuple with the MemberLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetMemberLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberLimit) {
		return nil, false
	}
	return o.MemberLimit, true
}

// HasMemberLimit returns a boolean if a field has been set.
func (o *EditChatInviteLinkRequest) HasMemberLimit() bool {
	if o != nil && !IsNil(o.MemberLimit) {
		return true
	}

	return false
}

// SetMemberLimit gets a reference to the given int32 and assigns it to the MemberLimit field.
func (o *EditChatInviteLinkRequest) SetMemberLimit(v int32) {
	o.MemberLimit = &v
}


// GetCreatesJoinRequest returns the CreatesJoinRequest field value if set, zero value otherwise.
func (o *EditChatInviteLinkRequest) GetCreatesJoinRequest() bool {
	if o == nil || IsNil(o.CreatesJoinRequest) {
		var ret bool
		return ret
	}
	return *o.CreatesJoinRequest
}

// GetCreatesJoinRequestOk returns a tuple with the CreatesJoinRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditChatInviteLinkRequest) GetCreatesJoinRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.CreatesJoinRequest) {
		return nil, false
	}
	return o.CreatesJoinRequest, true
}

// HasCreatesJoinRequest returns a boolean if a field has been set.
func (o *EditChatInviteLinkRequest) HasCreatesJoinRequest() bool {
	if o != nil && !IsNil(o.CreatesJoinRequest) {
		return true
	}

	return false
}

// SetCreatesJoinRequest gets a reference to the given bool and assigns it to the CreatesJoinRequest field.
func (o *EditChatInviteLinkRequest) SetCreatesJoinRequest(v bool) {
	o.CreatesJoinRequest = &v
}


func (o EditChatInviteLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditChatInviteLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["invite_link"] = o.InviteLink
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpireDate) {
		toSerialize["expire_date"] = o.ExpireDate
	}
	if !IsNil(o.MemberLimit) {
		toSerialize["member_limit"] = o.MemberLimit
	}
	if !IsNil(o.CreatesJoinRequest) {
		toSerialize["creates_join_request"] = o.CreatesJoinRequest
	}
	return toSerialize, nil
}

func (o *EditChatInviteLinkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEditChatInviteLinkRequest := _EditChatInviteLinkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditChatInviteLinkRequest)

	if err != nil {
		return err
	}

	*o = EditChatInviteLinkRequest(varEditChatInviteLinkRequest)

	return err
}

type NullableEditChatInviteLinkRequest struct {
	value *EditChatInviteLinkRequest
	isSet bool
}

func (v NullableEditChatInviteLinkRequest) Get() *EditChatInviteLinkRequest {
	return v.value
}

func (v *NullableEditChatInviteLinkRequest) Set(val *EditChatInviteLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditChatInviteLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditChatInviteLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditChatInviteLinkRequest(val *EditChatInviteLinkRequest) *NullableEditChatInviteLinkRequest {
	return &NullableEditChatInviteLinkRequest{value: val, isSet: true}
}

func (v NullableEditChatInviteLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditChatInviteLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


