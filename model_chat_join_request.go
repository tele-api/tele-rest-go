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

// checks if the ChatJoinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatJoinRequest{}

// ChatJoinRequest Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat Chat `json:"chat"`
	From User `json:"from"`
	// Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	UserChatId int32 `json:"user_chat_id"`
	// Date the request was sent in Unix time
	Date int32 `json:"date"`
	// *Optional*. Bio of the user.
	Bio *string `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type _ChatJoinRequest ChatJoinRequest

// NewChatJoinRequest instantiates a new ChatJoinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatJoinRequest(chat Chat, from User, userChatId int32, date int32) *ChatJoinRequest {
	this := ChatJoinRequest{}
	this.Chat = chat
	this.From = from
	this.UserChatId = userChatId
	this.Date = date
	return &this
}

// NewChatJoinRequestWithDefaults instantiates a new ChatJoinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatJoinRequestWithDefaults() *ChatJoinRequest {
	this := ChatJoinRequest{}
	return &this
}

// GetChat returns the Chat field value
func (o *ChatJoinRequest) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *ChatJoinRequest) SetChat(v Chat) {
	o.Chat = v
}

// GetFrom returns the From field value
func (o *ChatJoinRequest) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *ChatJoinRequest) SetFrom(v User) {
	o.From = v
}

// GetUserChatId returns the UserChatId field value
func (o *ChatJoinRequest) GetUserChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserChatId
}

// GetUserChatIdOk returns a tuple with the UserChatId field value
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetUserChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserChatId, true
}

// SetUserChatId sets field value
func (o *ChatJoinRequest) SetUserChatId(v int32) {
	o.UserChatId = v
}

// GetDate returns the Date field value
func (o *ChatJoinRequest) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ChatJoinRequest) SetDate(v int32) {
	o.Date = v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *ChatJoinRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *ChatJoinRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *ChatJoinRequest) SetBio(v string) {
	o.Bio = &v
}


// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
func (o *ChatJoinRequest) GetInviteLink() ChatInviteLink {
	if o == nil || IsNil(o.InviteLink) {
		var ret ChatInviteLink
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatJoinRequest) GetInviteLinkOk() (*ChatInviteLink, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *ChatJoinRequest) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given ChatInviteLink and assigns it to the InviteLink field.
func (o *ChatJoinRequest) SetInviteLink(v ChatInviteLink) {
	o.InviteLink = &v
}


func (o ChatJoinRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatJoinRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["from"] = o.From
	toSerialize["user_chat_id"] = o.UserChatId
	toSerialize["date"] = o.Date
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.InviteLink) {
		toSerialize["invite_link"] = o.InviteLink
	}
	return toSerialize, nil
}

func (o *ChatJoinRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"from",
		"user_chat_id",
		"date",
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

	varChatJoinRequest := _ChatJoinRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatJoinRequest)

	if err != nil {
		return err
	}

	*o = ChatJoinRequest(varChatJoinRequest)

	return err
}

type NullableChatJoinRequest struct {
	value *ChatJoinRequest
	isSet bool
}

func (v NullableChatJoinRequest) Get() *ChatJoinRequest {
	return v.value
}

func (v *NullableChatJoinRequest) Set(val *ChatJoinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatJoinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatJoinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatJoinRequest(val *ChatJoinRequest) *NullableChatJoinRequest {
	return &NullableChatJoinRequest{value: val, isSet: true}
}

func (v NullableChatJoinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatJoinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


