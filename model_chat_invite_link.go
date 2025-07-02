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

// checks if the ChatInviteLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatInviteLink{}

// ChatInviteLink Represents an invite link for a chat.
type ChatInviteLink struct {
	// The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`
	Creator User `json:"creator"`
	// *True*, if users joining the chat via the link need to be approved by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request"`
	// *True*, if the link is primary
	IsPrimary bool `json:"is_primary"`
	// *True*, if the link is revoked
	IsRevoked bool `json:"is_revoked"`
	// *Optional*. Invite link name
	Name *string `json:"name,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate *int32 `json:"expire_date,omitempty"`
	// *Optional*. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int32 `json:"member_limit,omitempty"`
	// *Optional*. Number of pending join requests created using this link
	PendingJoinRequestCount *int32 `json:"pending_join_request_count,omitempty"`
	// *Optional*. The number of seconds the subscription will be active for before the next payment
	SubscriptionPeriod *int32 `json:"subscription_period,omitempty"`
	// *Optional*. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
	SubscriptionPrice *int32 `json:"subscription_price,omitempty"`
}

type _ChatInviteLink ChatInviteLink

// NewChatInviteLink instantiates a new ChatInviteLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatInviteLink(inviteLink string, creator User, createsJoinRequest bool, isPrimary bool, isRevoked bool) *ChatInviteLink {
	this := ChatInviteLink{}
	this.InviteLink = inviteLink
	this.Creator = creator
	this.CreatesJoinRequest = createsJoinRequest
	this.IsPrimary = isPrimary
	this.IsRevoked = isRevoked
	return &this
}

// NewChatInviteLinkWithDefaults instantiates a new ChatInviteLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatInviteLinkWithDefaults() *ChatInviteLink {
	this := ChatInviteLink{}
	return &this
}

// GetInviteLink returns the InviteLink field value
func (o *ChatInviteLink) GetInviteLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetInviteLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteLink, true
}

// SetInviteLink sets field value
func (o *ChatInviteLink) SetInviteLink(v string) {
	o.InviteLink = v
}

// GetCreator returns the Creator field value
func (o *ChatInviteLink) GetCreator() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetCreatorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *ChatInviteLink) SetCreator(v User) {
	o.Creator = v
}

// GetCreatesJoinRequest returns the CreatesJoinRequest field value
func (o *ChatInviteLink) GetCreatesJoinRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreatesJoinRequest
}

// GetCreatesJoinRequestOk returns a tuple with the CreatesJoinRequest field value
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetCreatesJoinRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatesJoinRequest, true
}

// SetCreatesJoinRequest sets field value
func (o *ChatInviteLink) SetCreatesJoinRequest(v bool) {
	o.CreatesJoinRequest = v
}

// GetIsPrimary returns the IsPrimary field value
func (o *ChatInviteLink) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *ChatInviteLink) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetIsRevoked returns the IsRevoked field value
func (o *ChatInviteLink) GetIsRevoked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRevoked
}

// GetIsRevokedOk returns a tuple with the IsRevoked field value
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetIsRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRevoked, true
}

// SetIsRevoked sets field value
func (o *ChatInviteLink) SetIsRevoked(v bool) {
	o.IsRevoked = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChatInviteLink) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChatInviteLink) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChatInviteLink) SetName(v string) {
	o.Name = &v
}


// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *ChatInviteLink) GetExpireDate() int32 {
	if o == nil || IsNil(o.ExpireDate) {
		var ret int32
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetExpireDateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireDate) {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *ChatInviteLink) HasExpireDate() bool {
	if o != nil && !IsNil(o.ExpireDate) {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given int32 and assigns it to the ExpireDate field.
func (o *ChatInviteLink) SetExpireDate(v int32) {
	o.ExpireDate = &v
}


// GetMemberLimit returns the MemberLimit field value if set, zero value otherwise.
func (o *ChatInviteLink) GetMemberLimit() int32 {
	if o == nil || IsNil(o.MemberLimit) {
		var ret int32
		return ret
	}
	return *o.MemberLimit
}

// GetMemberLimitOk returns a tuple with the MemberLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetMemberLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberLimit) {
		return nil, false
	}
	return o.MemberLimit, true
}

// HasMemberLimit returns a boolean if a field has been set.
func (o *ChatInviteLink) HasMemberLimit() bool {
	if o != nil && !IsNil(o.MemberLimit) {
		return true
	}

	return false
}

// SetMemberLimit gets a reference to the given int32 and assigns it to the MemberLimit field.
func (o *ChatInviteLink) SetMemberLimit(v int32) {
	o.MemberLimit = &v
}


// GetPendingJoinRequestCount returns the PendingJoinRequestCount field value if set, zero value otherwise.
func (o *ChatInviteLink) GetPendingJoinRequestCount() int32 {
	if o == nil || IsNil(o.PendingJoinRequestCount) {
		var ret int32
		return ret
	}
	return *o.PendingJoinRequestCount
}

// GetPendingJoinRequestCountOk returns a tuple with the PendingJoinRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetPendingJoinRequestCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PendingJoinRequestCount) {
		return nil, false
	}
	return o.PendingJoinRequestCount, true
}

// HasPendingJoinRequestCount returns a boolean if a field has been set.
func (o *ChatInviteLink) HasPendingJoinRequestCount() bool {
	if o != nil && !IsNil(o.PendingJoinRequestCount) {
		return true
	}

	return false
}

// SetPendingJoinRequestCount gets a reference to the given int32 and assigns it to the PendingJoinRequestCount field.
func (o *ChatInviteLink) SetPendingJoinRequestCount(v int32) {
	o.PendingJoinRequestCount = &v
}


// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *ChatInviteLink) GetSubscriptionPeriod() int32 {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret int32
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetSubscriptionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *ChatInviteLink) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given int32 and assigns it to the SubscriptionPeriod field.
func (o *ChatInviteLink) SetSubscriptionPeriod(v int32) {
	o.SubscriptionPeriod = &v
}


// GetSubscriptionPrice returns the SubscriptionPrice field value if set, zero value otherwise.
func (o *ChatInviteLink) GetSubscriptionPrice() int32 {
	if o == nil || IsNil(o.SubscriptionPrice) {
		var ret int32
		return ret
	}
	return *o.SubscriptionPrice
}

// GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatInviteLink) GetSubscriptionPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionPrice) {
		return nil, false
	}
	return o.SubscriptionPrice, true
}

// HasSubscriptionPrice returns a boolean if a field has been set.
func (o *ChatInviteLink) HasSubscriptionPrice() bool {
	if o != nil && !IsNil(o.SubscriptionPrice) {
		return true
	}

	return false
}

// SetSubscriptionPrice gets a reference to the given int32 and assigns it to the SubscriptionPrice field.
func (o *ChatInviteLink) SetSubscriptionPrice(v int32) {
	o.SubscriptionPrice = &v
}


func (o ChatInviteLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatInviteLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invite_link"] = o.InviteLink
	toSerialize["creator"] = o.Creator
	toSerialize["creates_join_request"] = o.CreatesJoinRequest
	toSerialize["is_primary"] = o.IsPrimary
	toSerialize["is_revoked"] = o.IsRevoked
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpireDate) {
		toSerialize["expire_date"] = o.ExpireDate
	}
	if !IsNil(o.MemberLimit) {
		toSerialize["member_limit"] = o.MemberLimit
	}
	if !IsNil(o.PendingJoinRequestCount) {
		toSerialize["pending_join_request_count"] = o.PendingJoinRequestCount
	}
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscription_period"] = o.SubscriptionPeriod
	}
	if !IsNil(o.SubscriptionPrice) {
		toSerialize["subscription_price"] = o.SubscriptionPrice
	}
	return toSerialize, nil
}

func (o *ChatInviteLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invite_link",
		"creator",
		"creates_join_request",
		"is_primary",
		"is_revoked",
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

	varChatInviteLink := _ChatInviteLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatInviteLink)

	if err != nil {
		return err
	}

	*o = ChatInviteLink(varChatInviteLink)

	return err
}

type NullableChatInviteLink struct {
	value *ChatInviteLink
	isSet bool
}

func (v NullableChatInviteLink) Get() *ChatInviteLink {
	return v.value
}

func (v *NullableChatInviteLink) Set(val *ChatInviteLink) {
	v.value = val
	v.isSet = true
}

func (v NullableChatInviteLink) IsSet() bool {
	return v.isSet
}

func (v *NullableChatInviteLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatInviteLink(val *ChatInviteLink) *NullableChatInviteLink {
	return &NullableChatInviteLink{value: val, isSet: true}
}

func (v NullableChatInviteLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatInviteLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


