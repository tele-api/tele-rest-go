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

// checks if the PostPromoteChatMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPromoteChatMemberRequest{}

// PostPromoteChatMemberRequest struct for PostPromoteChatMemberRequest
type PostPromoteChatMemberRequest struct {
	ChatId PostSendMessageRequestChatId `json:"chat_id"`
	// Unique identifier of the target user
	UserId int32 `json:"user_id"`
	// Pass *True* if the administrator's presence in the chat is hidden
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// Pass *True* if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty"`
	// Pass *True* if the administrator can delete messages of other users
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`
	// Pass *True* if the administrator can manage video chats
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`
	// Pass *True* if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`
	// Pass *True* if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`
	// Pass *True* if the administrator can change chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty"`
	// Pass *True* if the administrator can invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`
	// Pass *True* if the administrator can post stories to the chat
	CanPostStories *bool `json:"can_post_stories,omitempty"`
	// Pass *True* if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories *bool `json:"can_edit_stories,omitempty"`
	// Pass *True* if the administrator can delete stories posted by other users
	CanDeleteStories *bool `json:"can_delete_stories,omitempty"`
	// Pass *True* if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	// Pass *True* if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	// Pass *True* if the administrator can pin messages; for supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	// Pass *True* if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

type _PostPromoteChatMemberRequest PostPromoteChatMemberRequest

// NewPostPromoteChatMemberRequest instantiates a new PostPromoteChatMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPromoteChatMemberRequest(chatId PostSendMessageRequestChatId, userId int32) *PostPromoteChatMemberRequest {
	this := PostPromoteChatMemberRequest{}
	this.ChatId = chatId
	this.UserId = userId
	return &this
}

// NewPostPromoteChatMemberRequestWithDefaults instantiates a new PostPromoteChatMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPromoteChatMemberRequestWithDefaults() *PostPromoteChatMemberRequest {
	this := PostPromoteChatMemberRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *PostPromoteChatMemberRequest) GetChatId() PostSendMessageRequestChatId {
	if o == nil {
		var ret PostSendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *PostPromoteChatMemberRequest) SetChatId(v PostSendMessageRequestChatId) {
	o.ChatId = v
}

// GetUserId returns the UserId field value
func (o *PostPromoteChatMemberRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostPromoteChatMemberRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetIsAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymous) {
		return nil, false
	}
	return o.IsAnonymous, true
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasIsAnonymous() bool {
	if o != nil && !IsNil(o.IsAnonymous) {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given bool and assigns it to the IsAnonymous field.
func (o *PostPromoteChatMemberRequest) SetIsAnonymous(v bool) {
	o.IsAnonymous = &v
}


// GetCanManageChat returns the CanManageChat field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanManageChat() bool {
	if o == nil || IsNil(o.CanManageChat) {
		var ret bool
		return ret
	}
	return *o.CanManageChat
}

// GetCanManageChatOk returns a tuple with the CanManageChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanManageChatOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageChat) {
		return nil, false
	}
	return o.CanManageChat, true
}

// HasCanManageChat returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanManageChat() bool {
	if o != nil && !IsNil(o.CanManageChat) {
		return true
	}

	return false
}

// SetCanManageChat gets a reference to the given bool and assigns it to the CanManageChat field.
func (o *PostPromoteChatMemberRequest) SetCanManageChat(v bool) {
	o.CanManageChat = &v
}


// GetCanDeleteMessages returns the CanDeleteMessages field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanDeleteMessages() bool {
	if o == nil || IsNil(o.CanDeleteMessages) {
		var ret bool
		return ret
	}
	return *o.CanDeleteMessages
}

// GetCanDeleteMessagesOk returns a tuple with the CanDeleteMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanDeleteMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteMessages) {
		return nil, false
	}
	return o.CanDeleteMessages, true
}

// HasCanDeleteMessages returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanDeleteMessages() bool {
	if o != nil && !IsNil(o.CanDeleteMessages) {
		return true
	}

	return false
}

// SetCanDeleteMessages gets a reference to the given bool and assigns it to the CanDeleteMessages field.
func (o *PostPromoteChatMemberRequest) SetCanDeleteMessages(v bool) {
	o.CanDeleteMessages = &v
}


// GetCanManageVideoChats returns the CanManageVideoChats field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanManageVideoChats() bool {
	if o == nil || IsNil(o.CanManageVideoChats) {
		var ret bool
		return ret
	}
	return *o.CanManageVideoChats
}

// GetCanManageVideoChatsOk returns a tuple with the CanManageVideoChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanManageVideoChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageVideoChats) {
		return nil, false
	}
	return o.CanManageVideoChats, true
}

// HasCanManageVideoChats returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanManageVideoChats() bool {
	if o != nil && !IsNil(o.CanManageVideoChats) {
		return true
	}

	return false
}

// SetCanManageVideoChats gets a reference to the given bool and assigns it to the CanManageVideoChats field.
func (o *PostPromoteChatMemberRequest) SetCanManageVideoChats(v bool) {
	o.CanManageVideoChats = &v
}


// GetCanRestrictMembers returns the CanRestrictMembers field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanRestrictMembers() bool {
	if o == nil || IsNil(o.CanRestrictMembers) {
		var ret bool
		return ret
	}
	return *o.CanRestrictMembers
}

// GetCanRestrictMembersOk returns a tuple with the CanRestrictMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanRestrictMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRestrictMembers) {
		return nil, false
	}
	return o.CanRestrictMembers, true
}

// HasCanRestrictMembers returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanRestrictMembers() bool {
	if o != nil && !IsNil(o.CanRestrictMembers) {
		return true
	}

	return false
}

// SetCanRestrictMembers gets a reference to the given bool and assigns it to the CanRestrictMembers field.
func (o *PostPromoteChatMemberRequest) SetCanRestrictMembers(v bool) {
	o.CanRestrictMembers = &v
}


// GetCanPromoteMembers returns the CanPromoteMembers field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanPromoteMembers() bool {
	if o == nil || IsNil(o.CanPromoteMembers) {
		var ret bool
		return ret
	}
	return *o.CanPromoteMembers
}

// GetCanPromoteMembersOk returns a tuple with the CanPromoteMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanPromoteMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPromoteMembers) {
		return nil, false
	}
	return o.CanPromoteMembers, true
}

// HasCanPromoteMembers returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanPromoteMembers() bool {
	if o != nil && !IsNil(o.CanPromoteMembers) {
		return true
	}

	return false
}

// SetCanPromoteMembers gets a reference to the given bool and assigns it to the CanPromoteMembers field.
func (o *PostPromoteChatMemberRequest) SetCanPromoteMembers(v bool) {
	o.CanPromoteMembers = &v
}


// GetCanChangeInfo returns the CanChangeInfo field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanChangeInfo() bool {
	if o == nil || IsNil(o.CanChangeInfo) {
		var ret bool
		return ret
	}
	return *o.CanChangeInfo
}

// GetCanChangeInfoOk returns a tuple with the CanChangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanChangeInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.CanChangeInfo) {
		return nil, false
	}
	return o.CanChangeInfo, true
}

// HasCanChangeInfo returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanChangeInfo() bool {
	if o != nil && !IsNil(o.CanChangeInfo) {
		return true
	}

	return false
}

// SetCanChangeInfo gets a reference to the given bool and assigns it to the CanChangeInfo field.
func (o *PostPromoteChatMemberRequest) SetCanChangeInfo(v bool) {
	o.CanChangeInfo = &v
}


// GetCanInviteUsers returns the CanInviteUsers field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanInviteUsers() bool {
	if o == nil || IsNil(o.CanInviteUsers) {
		var ret bool
		return ret
	}
	return *o.CanInviteUsers
}

// GetCanInviteUsersOk returns a tuple with the CanInviteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanInviteUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanInviteUsers) {
		return nil, false
	}
	return o.CanInviteUsers, true
}

// HasCanInviteUsers returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanInviteUsers() bool {
	if o != nil && !IsNil(o.CanInviteUsers) {
		return true
	}

	return false
}

// SetCanInviteUsers gets a reference to the given bool and assigns it to the CanInviteUsers field.
func (o *PostPromoteChatMemberRequest) SetCanInviteUsers(v bool) {
	o.CanInviteUsers = &v
}


// GetCanPostStories returns the CanPostStories field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanPostStories() bool {
	if o == nil || IsNil(o.CanPostStories) {
		var ret bool
		return ret
	}
	return *o.CanPostStories
}

// GetCanPostStoriesOk returns a tuple with the CanPostStories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanPostStoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPostStories) {
		return nil, false
	}
	return o.CanPostStories, true
}

// HasCanPostStories returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanPostStories() bool {
	if o != nil && !IsNil(o.CanPostStories) {
		return true
	}

	return false
}

// SetCanPostStories gets a reference to the given bool and assigns it to the CanPostStories field.
func (o *PostPromoteChatMemberRequest) SetCanPostStories(v bool) {
	o.CanPostStories = &v
}


// GetCanEditStories returns the CanEditStories field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanEditStories() bool {
	if o == nil || IsNil(o.CanEditStories) {
		var ret bool
		return ret
	}
	return *o.CanEditStories
}

// GetCanEditStoriesOk returns a tuple with the CanEditStories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanEditStoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditStories) {
		return nil, false
	}
	return o.CanEditStories, true
}

// HasCanEditStories returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanEditStories() bool {
	if o != nil && !IsNil(o.CanEditStories) {
		return true
	}

	return false
}

// SetCanEditStories gets a reference to the given bool and assigns it to the CanEditStories field.
func (o *PostPromoteChatMemberRequest) SetCanEditStories(v bool) {
	o.CanEditStories = &v
}


// GetCanDeleteStories returns the CanDeleteStories field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanDeleteStories() bool {
	if o == nil || IsNil(o.CanDeleteStories) {
		var ret bool
		return ret
	}
	return *o.CanDeleteStories
}

// GetCanDeleteStoriesOk returns a tuple with the CanDeleteStories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanDeleteStoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteStories) {
		return nil, false
	}
	return o.CanDeleteStories, true
}

// HasCanDeleteStories returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanDeleteStories() bool {
	if o != nil && !IsNil(o.CanDeleteStories) {
		return true
	}

	return false
}

// SetCanDeleteStories gets a reference to the given bool and assigns it to the CanDeleteStories field.
func (o *PostPromoteChatMemberRequest) SetCanDeleteStories(v bool) {
	o.CanDeleteStories = &v
}


// GetCanPostMessages returns the CanPostMessages field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanPostMessages() bool {
	if o == nil || IsNil(o.CanPostMessages) {
		var ret bool
		return ret
	}
	return *o.CanPostMessages
}

// GetCanPostMessagesOk returns a tuple with the CanPostMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanPostMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPostMessages) {
		return nil, false
	}
	return o.CanPostMessages, true
}

// HasCanPostMessages returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanPostMessages() bool {
	if o != nil && !IsNil(o.CanPostMessages) {
		return true
	}

	return false
}

// SetCanPostMessages gets a reference to the given bool and assigns it to the CanPostMessages field.
func (o *PostPromoteChatMemberRequest) SetCanPostMessages(v bool) {
	o.CanPostMessages = &v
}


// GetCanEditMessages returns the CanEditMessages field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanEditMessages() bool {
	if o == nil || IsNil(o.CanEditMessages) {
		var ret bool
		return ret
	}
	return *o.CanEditMessages
}

// GetCanEditMessagesOk returns a tuple with the CanEditMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanEditMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditMessages) {
		return nil, false
	}
	return o.CanEditMessages, true
}

// HasCanEditMessages returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanEditMessages() bool {
	if o != nil && !IsNil(o.CanEditMessages) {
		return true
	}

	return false
}

// SetCanEditMessages gets a reference to the given bool and assigns it to the CanEditMessages field.
func (o *PostPromoteChatMemberRequest) SetCanEditMessages(v bool) {
	o.CanEditMessages = &v
}


// GetCanPinMessages returns the CanPinMessages field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanPinMessages() bool {
	if o == nil || IsNil(o.CanPinMessages) {
		var ret bool
		return ret
	}
	return *o.CanPinMessages
}

// GetCanPinMessagesOk returns a tuple with the CanPinMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanPinMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPinMessages) {
		return nil, false
	}
	return o.CanPinMessages, true
}

// HasCanPinMessages returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanPinMessages() bool {
	if o != nil && !IsNil(o.CanPinMessages) {
		return true
	}

	return false
}

// SetCanPinMessages gets a reference to the given bool and assigns it to the CanPinMessages field.
func (o *PostPromoteChatMemberRequest) SetCanPinMessages(v bool) {
	o.CanPinMessages = &v
}


// GetCanManageTopics returns the CanManageTopics field value if set, zero value otherwise.
func (o *PostPromoteChatMemberRequest) GetCanManageTopics() bool {
	if o == nil || IsNil(o.CanManageTopics) {
		var ret bool
		return ret
	}
	return *o.CanManageTopics
}

// GetCanManageTopicsOk returns a tuple with the CanManageTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPromoteChatMemberRequest) GetCanManageTopicsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageTopics) {
		return nil, false
	}
	return o.CanManageTopics, true
}

// HasCanManageTopics returns a boolean if a field has been set.
func (o *PostPromoteChatMemberRequest) HasCanManageTopics() bool {
	if o != nil && !IsNil(o.CanManageTopics) {
		return true
	}

	return false
}

// SetCanManageTopics gets a reference to the given bool and assigns it to the CanManageTopics field.
func (o *PostPromoteChatMemberRequest) SetCanManageTopics(v bool) {
	o.CanManageTopics = &v
}


func (o PostPromoteChatMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPromoteChatMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.IsAnonymous) {
		toSerialize["is_anonymous"] = o.IsAnonymous
	}
	if !IsNil(o.CanManageChat) {
		toSerialize["can_manage_chat"] = o.CanManageChat
	}
	if !IsNil(o.CanDeleteMessages) {
		toSerialize["can_delete_messages"] = o.CanDeleteMessages
	}
	if !IsNil(o.CanManageVideoChats) {
		toSerialize["can_manage_video_chats"] = o.CanManageVideoChats
	}
	if !IsNil(o.CanRestrictMembers) {
		toSerialize["can_restrict_members"] = o.CanRestrictMembers
	}
	if !IsNil(o.CanPromoteMembers) {
		toSerialize["can_promote_members"] = o.CanPromoteMembers
	}
	if !IsNil(o.CanChangeInfo) {
		toSerialize["can_change_info"] = o.CanChangeInfo
	}
	if !IsNil(o.CanInviteUsers) {
		toSerialize["can_invite_users"] = o.CanInviteUsers
	}
	if !IsNil(o.CanPostStories) {
		toSerialize["can_post_stories"] = o.CanPostStories
	}
	if !IsNil(o.CanEditStories) {
		toSerialize["can_edit_stories"] = o.CanEditStories
	}
	if !IsNil(o.CanDeleteStories) {
		toSerialize["can_delete_stories"] = o.CanDeleteStories
	}
	if !IsNil(o.CanPostMessages) {
		toSerialize["can_post_messages"] = o.CanPostMessages
	}
	if !IsNil(o.CanEditMessages) {
		toSerialize["can_edit_messages"] = o.CanEditMessages
	}
	if !IsNil(o.CanPinMessages) {
		toSerialize["can_pin_messages"] = o.CanPinMessages
	}
	if !IsNil(o.CanManageTopics) {
		toSerialize["can_manage_topics"] = o.CanManageTopics
	}
	return toSerialize, nil
}

func (o *PostPromoteChatMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"user_id",
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

	varPostPromoteChatMemberRequest := _PostPromoteChatMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostPromoteChatMemberRequest)

	if err != nil {
		return err
	}

	*o = PostPromoteChatMemberRequest(varPostPromoteChatMemberRequest)

	return err
}

type NullablePostPromoteChatMemberRequest struct {
	value *PostPromoteChatMemberRequest
	isSet bool
}

func (v NullablePostPromoteChatMemberRequest) Get() *PostPromoteChatMemberRequest {
	return v.value
}

func (v *NullablePostPromoteChatMemberRequest) Set(val *PostPromoteChatMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPromoteChatMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPromoteChatMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPromoteChatMemberRequest(val *PostPromoteChatMemberRequest) *NullablePostPromoteChatMemberRequest {
	return &NullablePostPromoteChatMemberRequest{value: val, isSet: true}
}

func (v NullablePostPromoteChatMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPromoteChatMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


