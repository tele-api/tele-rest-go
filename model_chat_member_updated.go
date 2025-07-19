/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the ChatMemberUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMemberUpdated{}

// ChatMemberUpdated This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat Chat `json:"chat"`
	From User `json:"from"`
	// Date the change was done in Unix time
	Date int32 `json:"date"`
	OldChatMember ChatMember `json:"old_chat_member"`
	NewChatMember ChatMember `json:"new_chat_member"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	// *Optional*. *True*, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaJoinRequest *bool `json:"via_join_request,omitempty"`
	// *Optional*. *True*, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink *bool `json:"via_chat_folder_invite_link,omitempty"`
}

type _ChatMemberUpdated ChatMemberUpdated

// NewChatMemberUpdated instantiates a new ChatMemberUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMemberUpdated(chat Chat, from User, date int32, oldChatMember ChatMember, newChatMember ChatMember) *ChatMemberUpdated {
	this := ChatMemberUpdated{}
	this.Chat = chat
	this.From = from
	this.Date = date
	this.OldChatMember = oldChatMember
	this.NewChatMember = newChatMember
	return &this
}

// NewChatMemberUpdatedWithDefaults instantiates a new ChatMemberUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMemberUpdatedWithDefaults() *ChatMemberUpdated {
	this := ChatMemberUpdated{}
	return &this
}

// GetChat returns the Chat field value
func (o *ChatMemberUpdated) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *ChatMemberUpdated) SetChat(v Chat) {
	o.Chat = v
}

// GetFrom returns the From field value
func (o *ChatMemberUpdated) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *ChatMemberUpdated) SetFrom(v User) {
	o.From = v
}

// GetDate returns the Date field value
func (o *ChatMemberUpdated) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ChatMemberUpdated) SetDate(v int32) {
	o.Date = v
}

// GetOldChatMember returns the OldChatMember field value
func (o *ChatMemberUpdated) GetOldChatMember() ChatMember {
	if o == nil {
		var ret ChatMember
		return ret
	}

	return o.OldChatMember
}

// GetOldChatMemberOk returns a tuple with the OldChatMember field value
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetOldChatMemberOk() (*ChatMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldChatMember, true
}

// SetOldChatMember sets field value
func (o *ChatMemberUpdated) SetOldChatMember(v ChatMember) {
	o.OldChatMember = v
}

// GetNewChatMember returns the NewChatMember field value
func (o *ChatMemberUpdated) GetNewChatMember() ChatMember {
	if o == nil {
		var ret ChatMember
		return ret
	}

	return o.NewChatMember
}

// GetNewChatMemberOk returns a tuple with the NewChatMember field value
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetNewChatMemberOk() (*ChatMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewChatMember, true
}

// SetNewChatMember sets field value
func (o *ChatMemberUpdated) SetNewChatMember(v ChatMember) {
	o.NewChatMember = v
}

// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
func (o *ChatMemberUpdated) GetInviteLink() ChatInviteLink {
	if o == nil || IsNil(o.InviteLink) {
		var ret ChatInviteLink
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetInviteLinkOk() (*ChatInviteLink, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *ChatMemberUpdated) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given ChatInviteLink and assigns it to the InviteLink field.
func (o *ChatMemberUpdated) SetInviteLink(v ChatInviteLink) {
	o.InviteLink = &v
}


// GetViaJoinRequest returns the ViaJoinRequest field value if set, zero value otherwise.
func (o *ChatMemberUpdated) GetViaJoinRequest() bool {
	if o == nil || IsNil(o.ViaJoinRequest) {
		var ret bool
		return ret
	}
	return *o.ViaJoinRequest
}

// GetViaJoinRequestOk returns a tuple with the ViaJoinRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetViaJoinRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.ViaJoinRequest) {
		return nil, false
	}
	return o.ViaJoinRequest, true
}

// HasViaJoinRequest returns a boolean if a field has been set.
func (o *ChatMemberUpdated) HasViaJoinRequest() bool {
	if o != nil && !IsNil(o.ViaJoinRequest) {
		return true
	}

	return false
}

// SetViaJoinRequest gets a reference to the given bool and assigns it to the ViaJoinRequest field.
func (o *ChatMemberUpdated) SetViaJoinRequest(v bool) {
	o.ViaJoinRequest = &v
}


// GetViaChatFolderInviteLink returns the ViaChatFolderInviteLink field value if set, zero value otherwise.
func (o *ChatMemberUpdated) GetViaChatFolderInviteLink() bool {
	if o == nil || IsNil(o.ViaChatFolderInviteLink) {
		var ret bool
		return ret
	}
	return *o.ViaChatFolderInviteLink
}

// GetViaChatFolderInviteLinkOk returns a tuple with the ViaChatFolderInviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberUpdated) GetViaChatFolderInviteLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.ViaChatFolderInviteLink) {
		return nil, false
	}
	return o.ViaChatFolderInviteLink, true
}

// HasViaChatFolderInviteLink returns a boolean if a field has been set.
func (o *ChatMemberUpdated) HasViaChatFolderInviteLink() bool {
	if o != nil && !IsNil(o.ViaChatFolderInviteLink) {
		return true
	}

	return false
}

// SetViaChatFolderInviteLink gets a reference to the given bool and assigns it to the ViaChatFolderInviteLink field.
func (o *ChatMemberUpdated) SetViaChatFolderInviteLink(v bool) {
	o.ViaChatFolderInviteLink = &v
}


func (o ChatMemberUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMemberUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["from"] = o.From
	toSerialize["date"] = o.Date
	toSerialize["old_chat_member"] = o.OldChatMember
	toSerialize["new_chat_member"] = o.NewChatMember
	if !IsNil(o.InviteLink) {
		toSerialize["invite_link"] = o.InviteLink
	}
	if !IsNil(o.ViaJoinRequest) {
		toSerialize["via_join_request"] = o.ViaJoinRequest
	}
	if !IsNil(o.ViaChatFolderInviteLink) {
		toSerialize["via_chat_folder_invite_link"] = o.ViaChatFolderInviteLink
	}
	return toSerialize, nil
}

func (o *ChatMemberUpdated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"from",
		"date",
		"old_chat_member",
		"new_chat_member",
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

	varChatMemberUpdated := _ChatMemberUpdated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatMemberUpdated)

	if err != nil {
		return err
	}

	*o = ChatMemberUpdated(varChatMemberUpdated)

	return err
}

type NullableChatMemberUpdated struct {
	value *ChatMemberUpdated
	isSet bool
}

func (v NullableChatMemberUpdated) Get() *ChatMemberUpdated {
	return v.value
}

func (v *NullableChatMemberUpdated) Set(val *ChatMemberUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMemberUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMemberUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMemberUpdated(val *ChatMemberUpdated) *NullableChatMemberUpdated {
	return &NullableChatMemberUpdated{value: val, isSet: true}
}

func (v NullableChatMemberUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMemberUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


