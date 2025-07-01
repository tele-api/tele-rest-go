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
)

// checks if the ChatPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatPermissions{}

// ChatPermissions Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// *Optional*. *True*, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendMessages *bool `json:"can_send_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to send audios
	CanSendAudios *bool `json:"can_send_audios,omitempty"`
	// *Optional*. *True*, if the user is allowed to send documents
	CanSendDocuments *bool `json:"can_send_documents,omitempty"`
	// *Optional*. *True*, if the user is allowed to send photos
	CanSendPhotos *bool `json:"can_send_photos,omitempty"`
	// *Optional*. *True*, if the user is allowed to send videos
	CanSendVideos *bool `json:"can_send_videos,omitempty"`
	// *Optional*. *True*, if the user is allowed to send video notes
	CanSendVideoNotes *bool `json:"can_send_video_notes,omitempty"`
	// *Optional*. *True*, if the user is allowed to send voice notes
	CanSendVoiceNotes *bool `json:"can_send_voice_notes,omitempty"`
	// *Optional*. *True*, if the user is allowed to send polls
	CanSendPolls *bool `json:"can_send_polls,omitempty"`
	// *Optional*. *True*, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	// *Optional*. *True*, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanChangeInfo *bool `json:"can_change_info,omitempty"`
	// *Optional*. *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`
	// *Optional*. *True*, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to create forum topics. If omitted defaults to the value of can\\_pin\\_messages
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// NewChatPermissions instantiates a new ChatPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatPermissions() *ChatPermissions {
	this := ChatPermissions{}
	return &this
}

// NewChatPermissionsWithDefaults instantiates a new ChatPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatPermissionsWithDefaults() *ChatPermissions {
	this := ChatPermissions{}
	return &this
}

// GetCanSendMessages returns the CanSendMessages field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendMessages() bool {
	if o == nil || IsNil(o.CanSendMessages) {
		var ret bool
		return ret
	}
	return *o.CanSendMessages
}

// GetCanSendMessagesOk returns a tuple with the CanSendMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendMessages) {
		return nil, false
	}
	return o.CanSendMessages, true
}

// HasCanSendMessages returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendMessages() bool {
	if o != nil && !IsNil(o.CanSendMessages) {
		return true
	}

	return false
}

// SetCanSendMessages gets a reference to the given bool and assigns it to the CanSendMessages field.
func (o *ChatPermissions) SetCanSendMessages(v bool) {
	o.CanSendMessages = &v
}


// GetCanSendAudios returns the CanSendAudios field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendAudios() bool {
	if o == nil || IsNil(o.CanSendAudios) {
		var ret bool
		return ret
	}
	return *o.CanSendAudios
}

// GetCanSendAudiosOk returns a tuple with the CanSendAudios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendAudiosOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendAudios) {
		return nil, false
	}
	return o.CanSendAudios, true
}

// HasCanSendAudios returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendAudios() bool {
	if o != nil && !IsNil(o.CanSendAudios) {
		return true
	}

	return false
}

// SetCanSendAudios gets a reference to the given bool and assigns it to the CanSendAudios field.
func (o *ChatPermissions) SetCanSendAudios(v bool) {
	o.CanSendAudios = &v
}


// GetCanSendDocuments returns the CanSendDocuments field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendDocuments() bool {
	if o == nil || IsNil(o.CanSendDocuments) {
		var ret bool
		return ret
	}
	return *o.CanSendDocuments
}

// GetCanSendDocumentsOk returns a tuple with the CanSendDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendDocuments) {
		return nil, false
	}
	return o.CanSendDocuments, true
}

// HasCanSendDocuments returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendDocuments() bool {
	if o != nil && !IsNil(o.CanSendDocuments) {
		return true
	}

	return false
}

// SetCanSendDocuments gets a reference to the given bool and assigns it to the CanSendDocuments field.
func (o *ChatPermissions) SetCanSendDocuments(v bool) {
	o.CanSendDocuments = &v
}


// GetCanSendPhotos returns the CanSendPhotos field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendPhotos() bool {
	if o == nil || IsNil(o.CanSendPhotos) {
		var ret bool
		return ret
	}
	return *o.CanSendPhotos
}

// GetCanSendPhotosOk returns a tuple with the CanSendPhotos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendPhotosOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendPhotos) {
		return nil, false
	}
	return o.CanSendPhotos, true
}

// HasCanSendPhotos returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendPhotos() bool {
	if o != nil && !IsNil(o.CanSendPhotos) {
		return true
	}

	return false
}

// SetCanSendPhotos gets a reference to the given bool and assigns it to the CanSendPhotos field.
func (o *ChatPermissions) SetCanSendPhotos(v bool) {
	o.CanSendPhotos = &v
}


// GetCanSendVideos returns the CanSendVideos field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendVideos() bool {
	if o == nil || IsNil(o.CanSendVideos) {
		var ret bool
		return ret
	}
	return *o.CanSendVideos
}

// GetCanSendVideosOk returns a tuple with the CanSendVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendVideosOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendVideos) {
		return nil, false
	}
	return o.CanSendVideos, true
}

// HasCanSendVideos returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendVideos() bool {
	if o != nil && !IsNil(o.CanSendVideos) {
		return true
	}

	return false
}

// SetCanSendVideos gets a reference to the given bool and assigns it to the CanSendVideos field.
func (o *ChatPermissions) SetCanSendVideos(v bool) {
	o.CanSendVideos = &v
}


// GetCanSendVideoNotes returns the CanSendVideoNotes field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendVideoNotes() bool {
	if o == nil || IsNil(o.CanSendVideoNotes) {
		var ret bool
		return ret
	}
	return *o.CanSendVideoNotes
}

// GetCanSendVideoNotesOk returns a tuple with the CanSendVideoNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendVideoNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendVideoNotes) {
		return nil, false
	}
	return o.CanSendVideoNotes, true
}

// HasCanSendVideoNotes returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendVideoNotes() bool {
	if o != nil && !IsNil(o.CanSendVideoNotes) {
		return true
	}

	return false
}

// SetCanSendVideoNotes gets a reference to the given bool and assigns it to the CanSendVideoNotes field.
func (o *ChatPermissions) SetCanSendVideoNotes(v bool) {
	o.CanSendVideoNotes = &v
}


// GetCanSendVoiceNotes returns the CanSendVoiceNotes field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendVoiceNotes() bool {
	if o == nil || IsNil(o.CanSendVoiceNotes) {
		var ret bool
		return ret
	}
	return *o.CanSendVoiceNotes
}

// GetCanSendVoiceNotesOk returns a tuple with the CanSendVoiceNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendVoiceNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendVoiceNotes) {
		return nil, false
	}
	return o.CanSendVoiceNotes, true
}

// HasCanSendVoiceNotes returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendVoiceNotes() bool {
	if o != nil && !IsNil(o.CanSendVoiceNotes) {
		return true
	}

	return false
}

// SetCanSendVoiceNotes gets a reference to the given bool and assigns it to the CanSendVoiceNotes field.
func (o *ChatPermissions) SetCanSendVoiceNotes(v bool) {
	o.CanSendVoiceNotes = &v
}


// GetCanSendPolls returns the CanSendPolls field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendPolls() bool {
	if o == nil || IsNil(o.CanSendPolls) {
		var ret bool
		return ret
	}
	return *o.CanSendPolls
}

// GetCanSendPollsOk returns a tuple with the CanSendPolls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendPollsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendPolls) {
		return nil, false
	}
	return o.CanSendPolls, true
}

// HasCanSendPolls returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendPolls() bool {
	if o != nil && !IsNil(o.CanSendPolls) {
		return true
	}

	return false
}

// SetCanSendPolls gets a reference to the given bool and assigns it to the CanSendPolls field.
func (o *ChatPermissions) SetCanSendPolls(v bool) {
	o.CanSendPolls = &v
}


// GetCanSendOtherMessages returns the CanSendOtherMessages field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanSendOtherMessages() bool {
	if o == nil || IsNil(o.CanSendOtherMessages) {
		var ret bool
		return ret
	}
	return *o.CanSendOtherMessages
}

// GetCanSendOtherMessagesOk returns a tuple with the CanSendOtherMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanSendOtherMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendOtherMessages) {
		return nil, false
	}
	return o.CanSendOtherMessages, true
}

// HasCanSendOtherMessages returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanSendOtherMessages() bool {
	if o != nil && !IsNil(o.CanSendOtherMessages) {
		return true
	}

	return false
}

// SetCanSendOtherMessages gets a reference to the given bool and assigns it to the CanSendOtherMessages field.
func (o *ChatPermissions) SetCanSendOtherMessages(v bool) {
	o.CanSendOtherMessages = &v
}


// GetCanAddWebPagePreviews returns the CanAddWebPagePreviews field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanAddWebPagePreviews() bool {
	if o == nil || IsNil(o.CanAddWebPagePreviews) {
		var ret bool
		return ret
	}
	return *o.CanAddWebPagePreviews
}

// GetCanAddWebPagePreviewsOk returns a tuple with the CanAddWebPagePreviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanAddWebPagePreviewsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAddWebPagePreviews) {
		return nil, false
	}
	return o.CanAddWebPagePreviews, true
}

// HasCanAddWebPagePreviews returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanAddWebPagePreviews() bool {
	if o != nil && !IsNil(o.CanAddWebPagePreviews) {
		return true
	}

	return false
}

// SetCanAddWebPagePreviews gets a reference to the given bool and assigns it to the CanAddWebPagePreviews field.
func (o *ChatPermissions) SetCanAddWebPagePreviews(v bool) {
	o.CanAddWebPagePreviews = &v
}


// GetCanChangeInfo returns the CanChangeInfo field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanChangeInfo() bool {
	if o == nil || IsNil(o.CanChangeInfo) {
		var ret bool
		return ret
	}
	return *o.CanChangeInfo
}

// GetCanChangeInfoOk returns a tuple with the CanChangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanChangeInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.CanChangeInfo) {
		return nil, false
	}
	return o.CanChangeInfo, true
}

// HasCanChangeInfo returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanChangeInfo() bool {
	if o != nil && !IsNil(o.CanChangeInfo) {
		return true
	}

	return false
}

// SetCanChangeInfo gets a reference to the given bool and assigns it to the CanChangeInfo field.
func (o *ChatPermissions) SetCanChangeInfo(v bool) {
	o.CanChangeInfo = &v
}


// GetCanInviteUsers returns the CanInviteUsers field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanInviteUsers() bool {
	if o == nil || IsNil(o.CanInviteUsers) {
		var ret bool
		return ret
	}
	return *o.CanInviteUsers
}

// GetCanInviteUsersOk returns a tuple with the CanInviteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanInviteUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanInviteUsers) {
		return nil, false
	}
	return o.CanInviteUsers, true
}

// HasCanInviteUsers returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanInviteUsers() bool {
	if o != nil && !IsNil(o.CanInviteUsers) {
		return true
	}

	return false
}

// SetCanInviteUsers gets a reference to the given bool and assigns it to the CanInviteUsers field.
func (o *ChatPermissions) SetCanInviteUsers(v bool) {
	o.CanInviteUsers = &v
}


// GetCanPinMessages returns the CanPinMessages field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanPinMessages() bool {
	if o == nil || IsNil(o.CanPinMessages) {
		var ret bool
		return ret
	}
	return *o.CanPinMessages
}

// GetCanPinMessagesOk returns a tuple with the CanPinMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanPinMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPinMessages) {
		return nil, false
	}
	return o.CanPinMessages, true
}

// HasCanPinMessages returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanPinMessages() bool {
	if o != nil && !IsNil(o.CanPinMessages) {
		return true
	}

	return false
}

// SetCanPinMessages gets a reference to the given bool and assigns it to the CanPinMessages field.
func (o *ChatPermissions) SetCanPinMessages(v bool) {
	o.CanPinMessages = &v
}


// GetCanManageTopics returns the CanManageTopics field value if set, zero value otherwise.
func (o *ChatPermissions) GetCanManageTopics() bool {
	if o == nil || IsNil(o.CanManageTopics) {
		var ret bool
		return ret
	}
	return *o.CanManageTopics
}

// GetCanManageTopicsOk returns a tuple with the CanManageTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatPermissions) GetCanManageTopicsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageTopics) {
		return nil, false
	}
	return o.CanManageTopics, true
}

// HasCanManageTopics returns a boolean if a field has been set.
func (o *ChatPermissions) HasCanManageTopics() bool {
	if o != nil && !IsNil(o.CanManageTopics) {
		return true
	}

	return false
}

// SetCanManageTopics gets a reference to the given bool and assigns it to the CanManageTopics field.
func (o *ChatPermissions) SetCanManageTopics(v bool) {
	o.CanManageTopics = &v
}


func (o ChatPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanSendMessages) {
		toSerialize["can_send_messages"] = o.CanSendMessages
	}
	if !IsNil(o.CanSendAudios) {
		toSerialize["can_send_audios"] = o.CanSendAudios
	}
	if !IsNil(o.CanSendDocuments) {
		toSerialize["can_send_documents"] = o.CanSendDocuments
	}
	if !IsNil(o.CanSendPhotos) {
		toSerialize["can_send_photos"] = o.CanSendPhotos
	}
	if !IsNil(o.CanSendVideos) {
		toSerialize["can_send_videos"] = o.CanSendVideos
	}
	if !IsNil(o.CanSendVideoNotes) {
		toSerialize["can_send_video_notes"] = o.CanSendVideoNotes
	}
	if !IsNil(o.CanSendVoiceNotes) {
		toSerialize["can_send_voice_notes"] = o.CanSendVoiceNotes
	}
	if !IsNil(o.CanSendPolls) {
		toSerialize["can_send_polls"] = o.CanSendPolls
	}
	if !IsNil(o.CanSendOtherMessages) {
		toSerialize["can_send_other_messages"] = o.CanSendOtherMessages
	}
	if !IsNil(o.CanAddWebPagePreviews) {
		toSerialize["can_add_web_page_previews"] = o.CanAddWebPagePreviews
	}
	if !IsNil(o.CanChangeInfo) {
		toSerialize["can_change_info"] = o.CanChangeInfo
	}
	if !IsNil(o.CanInviteUsers) {
		toSerialize["can_invite_users"] = o.CanInviteUsers
	}
	if !IsNil(o.CanPinMessages) {
		toSerialize["can_pin_messages"] = o.CanPinMessages
	}
	if !IsNil(o.CanManageTopics) {
		toSerialize["can_manage_topics"] = o.CanManageTopics
	}
	return toSerialize, nil
}

type NullableChatPermissions struct {
	value *ChatPermissions
	isSet bool
}

func (v NullableChatPermissions) Get() *ChatPermissions {
	return v.value
}

func (v *NullableChatPermissions) Set(val *ChatPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableChatPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableChatPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatPermissions(val *ChatPermissions) *NullableChatPermissions {
	return &NullableChatPermissions{value: val, isSet: true}
}

func (v NullableChatPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


