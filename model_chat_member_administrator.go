/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the ChatMemberAdministrator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMemberAdministrator{}

// ChatMemberAdministrator Represents a [chat member](https://core.telegram.org/bots/api/#chatmember) that has some additional privileges.
type ChatMemberAdministrator struct {
	// The member's status in the chat, always “administrator”
	Status string `json:"status"`
	User User `json:"user"`
	// *True*, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`
	// *True*, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	// *True*, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	// *True*, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`
	// *True*, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	// *True*, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`
	// *True*, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`
	// *True*, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	// *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	// *True*, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`
	// *True*, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`
	// *True*, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`
	// *Optional*. *True*, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	// *Optional*. *True*, if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
	// *Optional*. Custom title for this user
	CustomTitle *string `json:"custom_title,omitempty"`
}

type _ChatMemberAdministrator ChatMemberAdministrator

// NewChatMemberAdministrator instantiates a new ChatMemberAdministrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMemberAdministrator(status string, user User, canBeEdited bool, isAnonymous bool, canManageChat bool, canDeleteMessages bool, canManageVideoChats bool, canRestrictMembers bool, canPromoteMembers bool, canChangeInfo bool, canInviteUsers bool, canPostStories bool, canEditStories bool, canDeleteStories bool) *ChatMemberAdministrator {
	this := ChatMemberAdministrator{}
	this.Status = status
	this.User = user
	this.CanBeEdited = canBeEdited
	this.IsAnonymous = isAnonymous
	this.CanManageChat = canManageChat
	this.CanDeleteMessages = canDeleteMessages
	this.CanManageVideoChats = canManageVideoChats
	this.CanRestrictMembers = canRestrictMembers
	this.CanPromoteMembers = canPromoteMembers
	this.CanChangeInfo = canChangeInfo
	this.CanInviteUsers = canInviteUsers
	this.CanPostStories = canPostStories
	this.CanEditStories = canEditStories
	this.CanDeleteStories = canDeleteStories
	return &this
}

// NewChatMemberAdministratorWithDefaults instantiates a new ChatMemberAdministrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMemberAdministratorWithDefaults() *ChatMemberAdministrator {
	this := ChatMemberAdministrator{}
	var status string = "administrator"
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *ChatMemberAdministrator) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChatMemberAdministrator) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value
func (o *ChatMemberAdministrator) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ChatMemberAdministrator) SetUser(v User) {
	o.User = v
}

// GetCanBeEdited returns the CanBeEdited field value
func (o *ChatMemberAdministrator) GetCanBeEdited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanBeEdited
}

// GetCanBeEditedOk returns a tuple with the CanBeEdited field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanBeEditedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanBeEdited, true
}

// SetCanBeEdited sets field value
func (o *ChatMemberAdministrator) SetCanBeEdited(v bool) {
	o.CanBeEdited = v
}

// GetIsAnonymous returns the IsAnonymous field value
func (o *ChatMemberAdministrator) GetIsAnonymous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetIsAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnonymous, true
}

// SetIsAnonymous sets field value
func (o *ChatMemberAdministrator) SetIsAnonymous(v bool) {
	o.IsAnonymous = v
}

// GetCanManageChat returns the CanManageChat field value
func (o *ChatMemberAdministrator) GetCanManageChat() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanManageChat
}

// GetCanManageChatOk returns a tuple with the CanManageChat field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanManageChatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanManageChat, true
}

// SetCanManageChat sets field value
func (o *ChatMemberAdministrator) SetCanManageChat(v bool) {
	o.CanManageChat = v
}

// GetCanDeleteMessages returns the CanDeleteMessages field value
func (o *ChatMemberAdministrator) GetCanDeleteMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDeleteMessages
}

// GetCanDeleteMessagesOk returns a tuple with the CanDeleteMessages field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanDeleteMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanDeleteMessages, true
}

// SetCanDeleteMessages sets field value
func (o *ChatMemberAdministrator) SetCanDeleteMessages(v bool) {
	o.CanDeleteMessages = v
}

// GetCanManageVideoChats returns the CanManageVideoChats field value
func (o *ChatMemberAdministrator) GetCanManageVideoChats() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanManageVideoChats
}

// GetCanManageVideoChatsOk returns a tuple with the CanManageVideoChats field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanManageVideoChatsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanManageVideoChats, true
}

// SetCanManageVideoChats sets field value
func (o *ChatMemberAdministrator) SetCanManageVideoChats(v bool) {
	o.CanManageVideoChats = v
}

// GetCanRestrictMembers returns the CanRestrictMembers field value
func (o *ChatMemberAdministrator) GetCanRestrictMembers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRestrictMembers
}

// GetCanRestrictMembersOk returns a tuple with the CanRestrictMembers field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanRestrictMembersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanRestrictMembers, true
}

// SetCanRestrictMembers sets field value
func (o *ChatMemberAdministrator) SetCanRestrictMembers(v bool) {
	o.CanRestrictMembers = v
}

// GetCanPromoteMembers returns the CanPromoteMembers field value
func (o *ChatMemberAdministrator) GetCanPromoteMembers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanPromoteMembers
}

// GetCanPromoteMembersOk returns a tuple with the CanPromoteMembers field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanPromoteMembersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanPromoteMembers, true
}

// SetCanPromoteMembers sets field value
func (o *ChatMemberAdministrator) SetCanPromoteMembers(v bool) {
	o.CanPromoteMembers = v
}

// GetCanChangeInfo returns the CanChangeInfo field value
func (o *ChatMemberAdministrator) GetCanChangeInfo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanChangeInfo
}

// GetCanChangeInfoOk returns a tuple with the CanChangeInfo field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanChangeInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanChangeInfo, true
}

// SetCanChangeInfo sets field value
func (o *ChatMemberAdministrator) SetCanChangeInfo(v bool) {
	o.CanChangeInfo = v
}

// GetCanInviteUsers returns the CanInviteUsers field value
func (o *ChatMemberAdministrator) GetCanInviteUsers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanInviteUsers
}

// GetCanInviteUsersOk returns a tuple with the CanInviteUsers field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanInviteUsersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanInviteUsers, true
}

// SetCanInviteUsers sets field value
func (o *ChatMemberAdministrator) SetCanInviteUsers(v bool) {
	o.CanInviteUsers = v
}

// GetCanPostStories returns the CanPostStories field value
func (o *ChatMemberAdministrator) GetCanPostStories() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanPostStories
}

// GetCanPostStoriesOk returns a tuple with the CanPostStories field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanPostStoriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanPostStories, true
}

// SetCanPostStories sets field value
func (o *ChatMemberAdministrator) SetCanPostStories(v bool) {
	o.CanPostStories = v
}

// GetCanEditStories returns the CanEditStories field value
func (o *ChatMemberAdministrator) GetCanEditStories() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditStories
}

// GetCanEditStoriesOk returns a tuple with the CanEditStories field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanEditStoriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditStories, true
}

// SetCanEditStories sets field value
func (o *ChatMemberAdministrator) SetCanEditStories(v bool) {
	o.CanEditStories = v
}

// GetCanDeleteStories returns the CanDeleteStories field value
func (o *ChatMemberAdministrator) GetCanDeleteStories() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDeleteStories
}

// GetCanDeleteStoriesOk returns a tuple with the CanDeleteStories field value
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanDeleteStoriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanDeleteStories, true
}

// SetCanDeleteStories sets field value
func (o *ChatMemberAdministrator) SetCanDeleteStories(v bool) {
	o.CanDeleteStories = v
}

// GetCanPostMessages returns the CanPostMessages field value if set, zero value otherwise.
func (o *ChatMemberAdministrator) GetCanPostMessages() bool {
	if o == nil || IsNil(o.CanPostMessages) {
		var ret bool
		return ret
	}
	return *o.CanPostMessages
}

// GetCanPostMessagesOk returns a tuple with the CanPostMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanPostMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPostMessages) {
		return nil, false
	}
	return o.CanPostMessages, true
}

// HasCanPostMessages returns a boolean if a field has been set.
func (o *ChatMemberAdministrator) HasCanPostMessages() bool {
	if o != nil && !IsNil(o.CanPostMessages) {
		return true
	}

	return false
}

// SetCanPostMessages gets a reference to the given bool and assigns it to the CanPostMessages field.
func (o *ChatMemberAdministrator) SetCanPostMessages(v bool) {
	o.CanPostMessages = &v
}


// GetCanEditMessages returns the CanEditMessages field value if set, zero value otherwise.
func (o *ChatMemberAdministrator) GetCanEditMessages() bool {
	if o == nil || IsNil(o.CanEditMessages) {
		var ret bool
		return ret
	}
	return *o.CanEditMessages
}

// GetCanEditMessagesOk returns a tuple with the CanEditMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanEditMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditMessages) {
		return nil, false
	}
	return o.CanEditMessages, true
}

// HasCanEditMessages returns a boolean if a field has been set.
func (o *ChatMemberAdministrator) HasCanEditMessages() bool {
	if o != nil && !IsNil(o.CanEditMessages) {
		return true
	}

	return false
}

// SetCanEditMessages gets a reference to the given bool and assigns it to the CanEditMessages field.
func (o *ChatMemberAdministrator) SetCanEditMessages(v bool) {
	o.CanEditMessages = &v
}


// GetCanPinMessages returns the CanPinMessages field value if set, zero value otherwise.
func (o *ChatMemberAdministrator) GetCanPinMessages() bool {
	if o == nil || IsNil(o.CanPinMessages) {
		var ret bool
		return ret
	}
	return *o.CanPinMessages
}

// GetCanPinMessagesOk returns a tuple with the CanPinMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanPinMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPinMessages) {
		return nil, false
	}
	return o.CanPinMessages, true
}

// HasCanPinMessages returns a boolean if a field has been set.
func (o *ChatMemberAdministrator) HasCanPinMessages() bool {
	if o != nil && !IsNil(o.CanPinMessages) {
		return true
	}

	return false
}

// SetCanPinMessages gets a reference to the given bool and assigns it to the CanPinMessages field.
func (o *ChatMemberAdministrator) SetCanPinMessages(v bool) {
	o.CanPinMessages = &v
}


// GetCanManageTopics returns the CanManageTopics field value if set, zero value otherwise.
func (o *ChatMemberAdministrator) GetCanManageTopics() bool {
	if o == nil || IsNil(o.CanManageTopics) {
		var ret bool
		return ret
	}
	return *o.CanManageTopics
}

// GetCanManageTopicsOk returns a tuple with the CanManageTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCanManageTopicsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageTopics) {
		return nil, false
	}
	return o.CanManageTopics, true
}

// HasCanManageTopics returns a boolean if a field has been set.
func (o *ChatMemberAdministrator) HasCanManageTopics() bool {
	if o != nil && !IsNil(o.CanManageTopics) {
		return true
	}

	return false
}

// SetCanManageTopics gets a reference to the given bool and assigns it to the CanManageTopics field.
func (o *ChatMemberAdministrator) SetCanManageTopics(v bool) {
	o.CanManageTopics = &v
}


// GetCustomTitle returns the CustomTitle field value if set, zero value otherwise.
func (o *ChatMemberAdministrator) GetCustomTitle() string {
	if o == nil || IsNil(o.CustomTitle) {
		var ret string
		return ret
	}
	return *o.CustomTitle
}

// GetCustomTitleOk returns a tuple with the CustomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberAdministrator) GetCustomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTitle) {
		return nil, false
	}
	return o.CustomTitle, true
}

// HasCustomTitle returns a boolean if a field has been set.
func (o *ChatMemberAdministrator) HasCustomTitle() bool {
	if o != nil && !IsNil(o.CustomTitle) {
		return true
	}

	return false
}

// SetCustomTitle gets a reference to the given string and assigns it to the CustomTitle field.
func (o *ChatMemberAdministrator) SetCustomTitle(v string) {
	o.CustomTitle = &v
}


func (o ChatMemberAdministrator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMemberAdministrator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["user"] = o.User
	toSerialize["can_be_edited"] = o.CanBeEdited
	toSerialize["is_anonymous"] = o.IsAnonymous
	toSerialize["can_manage_chat"] = o.CanManageChat
	toSerialize["can_delete_messages"] = o.CanDeleteMessages
	toSerialize["can_manage_video_chats"] = o.CanManageVideoChats
	toSerialize["can_restrict_members"] = o.CanRestrictMembers
	toSerialize["can_promote_members"] = o.CanPromoteMembers
	toSerialize["can_change_info"] = o.CanChangeInfo
	toSerialize["can_invite_users"] = o.CanInviteUsers
	toSerialize["can_post_stories"] = o.CanPostStories
	toSerialize["can_edit_stories"] = o.CanEditStories
	toSerialize["can_delete_stories"] = o.CanDeleteStories
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
	if !IsNil(o.CustomTitle) {
		toSerialize["custom_title"] = o.CustomTitle
	}
	return toSerialize, nil
}

func (o *ChatMemberAdministrator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"user",
		"can_be_edited",
		"is_anonymous",
		"can_manage_chat",
		"can_delete_messages",
		"can_manage_video_chats",
		"can_restrict_members",
		"can_promote_members",
		"can_change_info",
		"can_invite_users",
		"can_post_stories",
		"can_edit_stories",
		"can_delete_stories",
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

	varChatMemberAdministrator := _ChatMemberAdministrator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatMemberAdministrator)

	if err != nil {
		return err
	}

	*o = ChatMemberAdministrator(varChatMemberAdministrator)

	return err
}

type NullableChatMemberAdministrator struct {
	value *ChatMemberAdministrator
	isSet bool
}

func (v NullableChatMemberAdministrator) Get() *ChatMemberAdministrator {
	return v.value
}

func (v *NullableChatMemberAdministrator) Set(val *ChatMemberAdministrator) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMemberAdministrator) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMemberAdministrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMemberAdministrator(val *ChatMemberAdministrator) *NullableChatMemberAdministrator {
	return &NullableChatMemberAdministrator{value: val, isSet: true}
}

func (v NullableChatMemberAdministrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMemberAdministrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


