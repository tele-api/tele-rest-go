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

// checks if the ChatMemberRestricted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMemberRestricted{}

// ChatMemberRestricted Represents a [chat member](https://core.telegram.org/bots/api/#chatmember) that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	// The member's status in the chat, always “restricted”
	Status string `json:"status"`
	User User `json:"user"`
	// *True*, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`
	// *True*, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`
	// *True*, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`
	// *True*, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`
	// *True*, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`
	// *True*, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`
	// *True*, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`
	// *True*, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`
	// *True*, if the user is allowed to send polls and checklists
	CanSendPolls bool `json:"can_send_polls"`
	// *True*, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// *True*, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// *True*, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	// *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	// *True*, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages"`
	// *True*, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics"`
	// Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
	UntilDate int32 `json:"until_date"`
}

type _ChatMemberRestricted ChatMemberRestricted

// NewChatMemberRestricted instantiates a new ChatMemberRestricted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMemberRestricted(status string, user User, isMember bool, canSendMessages bool, canSendAudios bool, canSendDocuments bool, canSendPhotos bool, canSendVideos bool, canSendVideoNotes bool, canSendVoiceNotes bool, canSendPolls bool, canSendOtherMessages bool, canAddWebPagePreviews bool, canChangeInfo bool, canInviteUsers bool, canPinMessages bool, canManageTopics bool, untilDate int32) *ChatMemberRestricted {
	this := ChatMemberRestricted{}
	this.Status = status
	this.User = user
	this.IsMember = isMember
	this.CanSendMessages = canSendMessages
	this.CanSendAudios = canSendAudios
	this.CanSendDocuments = canSendDocuments
	this.CanSendPhotos = canSendPhotos
	this.CanSendVideos = canSendVideos
	this.CanSendVideoNotes = canSendVideoNotes
	this.CanSendVoiceNotes = canSendVoiceNotes
	this.CanSendPolls = canSendPolls
	this.CanSendOtherMessages = canSendOtherMessages
	this.CanAddWebPagePreviews = canAddWebPagePreviews
	this.CanChangeInfo = canChangeInfo
	this.CanInviteUsers = canInviteUsers
	this.CanPinMessages = canPinMessages
	this.CanManageTopics = canManageTopics
	this.UntilDate = untilDate
	return &this
}

// NewChatMemberRestrictedWithDefaults instantiates a new ChatMemberRestricted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMemberRestrictedWithDefaults() *ChatMemberRestricted {
	this := ChatMemberRestricted{}
	var status string = "restricted"
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *ChatMemberRestricted) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChatMemberRestricted) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value
func (o *ChatMemberRestricted) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ChatMemberRestricted) SetUser(v User) {
	o.User = v
}

// GetIsMember returns the IsMember field value
func (o *ChatMemberRestricted) GetIsMember() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMember
}

// GetIsMemberOk returns a tuple with the IsMember field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetIsMemberOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMember, true
}

// SetIsMember sets field value
func (o *ChatMemberRestricted) SetIsMember(v bool) {
	o.IsMember = v
}

// GetCanSendMessages returns the CanSendMessages field value
func (o *ChatMemberRestricted) GetCanSendMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendMessages
}

// GetCanSendMessagesOk returns a tuple with the CanSendMessages field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendMessages, true
}

// SetCanSendMessages sets field value
func (o *ChatMemberRestricted) SetCanSendMessages(v bool) {
	o.CanSendMessages = v
}

// GetCanSendAudios returns the CanSendAudios field value
func (o *ChatMemberRestricted) GetCanSendAudios() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendAudios
}

// GetCanSendAudiosOk returns a tuple with the CanSendAudios field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendAudiosOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendAudios, true
}

// SetCanSendAudios sets field value
func (o *ChatMemberRestricted) SetCanSendAudios(v bool) {
	o.CanSendAudios = v
}

// GetCanSendDocuments returns the CanSendDocuments field value
func (o *ChatMemberRestricted) GetCanSendDocuments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendDocuments
}

// GetCanSendDocumentsOk returns a tuple with the CanSendDocuments field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendDocumentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendDocuments, true
}

// SetCanSendDocuments sets field value
func (o *ChatMemberRestricted) SetCanSendDocuments(v bool) {
	o.CanSendDocuments = v
}

// GetCanSendPhotos returns the CanSendPhotos field value
func (o *ChatMemberRestricted) GetCanSendPhotos() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendPhotos
}

// GetCanSendPhotosOk returns a tuple with the CanSendPhotos field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendPhotosOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendPhotos, true
}

// SetCanSendPhotos sets field value
func (o *ChatMemberRestricted) SetCanSendPhotos(v bool) {
	o.CanSendPhotos = v
}

// GetCanSendVideos returns the CanSendVideos field value
func (o *ChatMemberRestricted) GetCanSendVideos() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendVideos
}

// GetCanSendVideosOk returns a tuple with the CanSendVideos field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendVideosOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendVideos, true
}

// SetCanSendVideos sets field value
func (o *ChatMemberRestricted) SetCanSendVideos(v bool) {
	o.CanSendVideos = v
}

// GetCanSendVideoNotes returns the CanSendVideoNotes field value
func (o *ChatMemberRestricted) GetCanSendVideoNotes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendVideoNotes
}

// GetCanSendVideoNotesOk returns a tuple with the CanSendVideoNotes field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendVideoNotesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendVideoNotes, true
}

// SetCanSendVideoNotes sets field value
func (o *ChatMemberRestricted) SetCanSendVideoNotes(v bool) {
	o.CanSendVideoNotes = v
}

// GetCanSendVoiceNotes returns the CanSendVoiceNotes field value
func (o *ChatMemberRestricted) GetCanSendVoiceNotes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendVoiceNotes
}

// GetCanSendVoiceNotesOk returns a tuple with the CanSendVoiceNotes field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendVoiceNotesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendVoiceNotes, true
}

// SetCanSendVoiceNotes sets field value
func (o *ChatMemberRestricted) SetCanSendVoiceNotes(v bool) {
	o.CanSendVoiceNotes = v
}

// GetCanSendPolls returns the CanSendPolls field value
func (o *ChatMemberRestricted) GetCanSendPolls() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendPolls
}

// GetCanSendPollsOk returns a tuple with the CanSendPolls field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendPollsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendPolls, true
}

// SetCanSendPolls sets field value
func (o *ChatMemberRestricted) SetCanSendPolls(v bool) {
	o.CanSendPolls = v
}

// GetCanSendOtherMessages returns the CanSendOtherMessages field value
func (o *ChatMemberRestricted) GetCanSendOtherMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSendOtherMessages
}

// GetCanSendOtherMessagesOk returns a tuple with the CanSendOtherMessages field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanSendOtherMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSendOtherMessages, true
}

// SetCanSendOtherMessages sets field value
func (o *ChatMemberRestricted) SetCanSendOtherMessages(v bool) {
	o.CanSendOtherMessages = v
}

// GetCanAddWebPagePreviews returns the CanAddWebPagePreviews field value
func (o *ChatMemberRestricted) GetCanAddWebPagePreviews() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanAddWebPagePreviews
}

// GetCanAddWebPagePreviewsOk returns a tuple with the CanAddWebPagePreviews field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanAddWebPagePreviewsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanAddWebPagePreviews, true
}

// SetCanAddWebPagePreviews sets field value
func (o *ChatMemberRestricted) SetCanAddWebPagePreviews(v bool) {
	o.CanAddWebPagePreviews = v
}

// GetCanChangeInfo returns the CanChangeInfo field value
func (o *ChatMemberRestricted) GetCanChangeInfo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanChangeInfo
}

// GetCanChangeInfoOk returns a tuple with the CanChangeInfo field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanChangeInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanChangeInfo, true
}

// SetCanChangeInfo sets field value
func (o *ChatMemberRestricted) SetCanChangeInfo(v bool) {
	o.CanChangeInfo = v
}

// GetCanInviteUsers returns the CanInviteUsers field value
func (o *ChatMemberRestricted) GetCanInviteUsers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanInviteUsers
}

// GetCanInviteUsersOk returns a tuple with the CanInviteUsers field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanInviteUsersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanInviteUsers, true
}

// SetCanInviteUsers sets field value
func (o *ChatMemberRestricted) SetCanInviteUsers(v bool) {
	o.CanInviteUsers = v
}

// GetCanPinMessages returns the CanPinMessages field value
func (o *ChatMemberRestricted) GetCanPinMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanPinMessages
}

// GetCanPinMessagesOk returns a tuple with the CanPinMessages field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanPinMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanPinMessages, true
}

// SetCanPinMessages sets field value
func (o *ChatMemberRestricted) SetCanPinMessages(v bool) {
	o.CanPinMessages = v
}

// GetCanManageTopics returns the CanManageTopics field value
func (o *ChatMemberRestricted) GetCanManageTopics() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanManageTopics
}

// GetCanManageTopicsOk returns a tuple with the CanManageTopics field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetCanManageTopicsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanManageTopics, true
}

// SetCanManageTopics sets field value
func (o *ChatMemberRestricted) SetCanManageTopics(v bool) {
	o.CanManageTopics = v
}

// GetUntilDate returns the UntilDate field value
func (o *ChatMemberRestricted) GetUntilDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UntilDate
}

// GetUntilDateOk returns a tuple with the UntilDate field value
// and a boolean to check if the value has been set.
func (o *ChatMemberRestricted) GetUntilDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UntilDate, true
}

// SetUntilDate sets field value
func (o *ChatMemberRestricted) SetUntilDate(v int32) {
	o.UntilDate = v
}

func (o ChatMemberRestricted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMemberRestricted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["user"] = o.User
	toSerialize["is_member"] = o.IsMember
	toSerialize["can_send_messages"] = o.CanSendMessages
	toSerialize["can_send_audios"] = o.CanSendAudios
	toSerialize["can_send_documents"] = o.CanSendDocuments
	toSerialize["can_send_photos"] = o.CanSendPhotos
	toSerialize["can_send_videos"] = o.CanSendVideos
	toSerialize["can_send_video_notes"] = o.CanSendVideoNotes
	toSerialize["can_send_voice_notes"] = o.CanSendVoiceNotes
	toSerialize["can_send_polls"] = o.CanSendPolls
	toSerialize["can_send_other_messages"] = o.CanSendOtherMessages
	toSerialize["can_add_web_page_previews"] = o.CanAddWebPagePreviews
	toSerialize["can_change_info"] = o.CanChangeInfo
	toSerialize["can_invite_users"] = o.CanInviteUsers
	toSerialize["can_pin_messages"] = o.CanPinMessages
	toSerialize["can_manage_topics"] = o.CanManageTopics
	toSerialize["until_date"] = o.UntilDate
	return toSerialize, nil
}

func (o *ChatMemberRestricted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"user",
		"is_member",
		"can_send_messages",
		"can_send_audios",
		"can_send_documents",
		"can_send_photos",
		"can_send_videos",
		"can_send_video_notes",
		"can_send_voice_notes",
		"can_send_polls",
		"can_send_other_messages",
		"can_add_web_page_previews",
		"can_change_info",
		"can_invite_users",
		"can_pin_messages",
		"can_manage_topics",
		"until_date",
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

	varChatMemberRestricted := _ChatMemberRestricted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatMemberRestricted)

	if err != nil {
		return err
	}

	*o = ChatMemberRestricted(varChatMemberRestricted)

	return err
}

type NullableChatMemberRestricted struct {
	value *ChatMemberRestricted
	isSet bool
}

func (v NullableChatMemberRestricted) Get() *ChatMemberRestricted {
	return v.value
}

func (v *NullableChatMemberRestricted) Set(val *ChatMemberRestricted) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMemberRestricted) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMemberRestricted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMemberRestricted(val *ChatMemberRestricted) *NullableChatMemberRestricted {
	return &NullableChatMemberRestricted{value: val, isSet: true}
}

func (v NullableChatMemberRestricted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMemberRestricted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


