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

// checks if the ChatFullInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatFullInfo{}

// ChatFullInfo This object contains full information about a chat.
type ChatFullInfo struct {
	// Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int32 `json:"id"`
	// Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// *Optional*. Title, for supergroups, channels and group chats
	Title *string `json:"title,omitempty"`
	// *Optional*. Username, for private chats, supergroups and channels if available
	Username *string `json:"username,omitempty"`
	// *Optional*. First name of the other party in a private chat
	FirstName *string `json:"first_name,omitempty"`
	// *Optional*. Last name of the other party in a private chat
	LastName *string `json:"last_name,omitempty"`
	// *Optional*. *True*, if the supergroup chat is a forum (has [topics](https://telegram.org/blog/topics-in-groups-collectible-usernames#topics-in-groups) enabled)
	IsForum *bool `json:"is_forum,omitempty"`
	// Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See [accent colors](https://core.telegram.org/bots/api/#accent-colors) for more details.
	AccentColorId int32 `json:"accent_color_id"`
	// The maximum number of reactions that can be set on a message in the chat
	MaxReactionCount int32 `json:"max_reaction_count"`
	Photo *ChatPhoto `json:"photo,omitempty"`
	// *Optional*. If non-empty, the list of all [active chat usernames](https://telegram.org/blog/topics-in-groups-collectible-usernames#collectible-usernames); for private chats, supergroups and channels
	ActiveUsernames []string `json:"active_usernames,omitempty"`
	Birthdate *Birthdate `json:"birthdate,omitempty"`
	BusinessIntro *BusinessIntro `json:"business_intro,omitempty"`
	BusinessLocation *BusinessLocation `json:"business_location,omitempty"`
	BusinessOpeningHours *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat *Chat `json:"personal_chat,omitempty"`
	// *Optional*. List of available reactions allowed in the chat. If omitted, then all [emoji reactions](https://core.telegram.org/bots/api/#reactiontypeemoji) are allowed.
	AvailableReactions []ReactionType `json:"available_reactions,omitempty"`
	// *Optional*. Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background
	BackgroundCustomEmojiId *string `json:"background_custom_emoji_id,omitempty"`
	// *Optional*. Identifier of the accent color for the chat's profile background. See [profile accent colors](https://core.telegram.org/bots/api/#profile-accent-colors) for more details.
	ProfileAccentColorId *int32 `json:"profile_accent_color_id,omitempty"`
	// *Optional*. Custom emoji identifier of the emoji chosen by the chat for its profile background
	ProfileBackgroundCustomEmojiId *string `json:"profile_background_custom_emoji_id,omitempty"`
	// *Optional*. Custom emoji identifier of the emoji status of the chat or the other party in a private chat
	EmojiStatusCustomEmojiId *string `json:"emoji_status_custom_emoji_id,omitempty"`
	// *Optional*. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any
	EmojiStatusExpirationDate *int32 `json:"emoji_status_expiration_date,omitempty"`
	// *Optional*. Bio of the other party in a private chat
	Bio *string `json:"bio,omitempty"`
	// *Optional*. *True*, if privacy settings of the other party in the private chat allows to use `tg://user?id=<user_id>` links only in chats with the user
	HasPrivateForwards *bool `json:"has_private_forwards,omitempty"`
	// *Optional*. *True*, if the privacy settings of the other party restrict sending voice and video note messages in the private chat
	HasRestrictedVoiceAndVideoMessages *bool `json:"has_restricted_voice_and_video_messages,omitempty"`
	// *Optional*. *True*, if users need to join the supergroup before they can send messages
	JoinToSendMessages *bool `json:"join_to_send_messages,omitempty"`
	// *Optional*. *True*, if all users directly joining the supergroup without using an invite link need to be approved by supergroup administrators
	JoinByRequest *bool `json:"join_by_request,omitempty"`
	// *Optional*. Description, for groups, supergroups and channel chats
	Description *string `json:"description,omitempty"`
	// *Optional*. Primary invite link, for groups, supergroups and channel chats
	InviteLink *string `json:"invite_link,omitempty"`
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	AcceptedGiftTypes AcceptedGiftTypes `json:"accepted_gift_types"`
	// *Optional*. *True*, if paid media messages can be sent or forwarded to the channel chat. The field is available only for channel chats.
	CanSendPaidMedia *bool `json:"can_send_paid_media,omitempty"`
	// *Optional*. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds
	SlowModeDelay *int32 `json:"slow_mode_delay,omitempty"`
	// *Optional*. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions
	UnrestrictBoostCount *int32 `json:"unrestrict_boost_count,omitempty"`
	// *Optional*. The time after which all messages sent to the chat will be automatically deleted; in seconds
	MessageAutoDeleteTime *int32 `json:"message_auto_delete_time,omitempty"`
	// *Optional*. *True*, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators.
	HasAggressiveAntiSpamEnabled *bool `json:"has_aggressive_anti_spam_enabled,omitempty"`
	// *Optional*. *True*, if non-administrators can only get the list of bots and administrators in the chat
	HasHiddenMembers *bool `json:"has_hidden_members,omitempty"`
	// *Optional*. *True*, if messages from the chat can't be forwarded to other chats
	HasProtectedContent *bool `json:"has_protected_content,omitempty"`
	// *Optional*. *True*, if new chat members will have access to old messages; available only to chat administrators
	HasVisibleHistory *bool `json:"has_visible_history,omitempty"`
	// *Optional*. For supergroups, name of the group sticker set
	StickerSetName *string `json:"sticker_set_name,omitempty"`
	// *Optional*. *True*, if the bot can change the group sticker set
	CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`
	// *Optional*. For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group.
	CustomEmojiStickerSetName *string `json:"custom_emoji_sticker_set_name,omitempty"`
	// *Optional*. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	LinkedChatId *int32 `json:"linked_chat_id,omitempty"`
	Location *ChatLocation `json:"location,omitempty"`
}

type _ChatFullInfo ChatFullInfo

// NewChatFullInfo instantiates a new ChatFullInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatFullInfo(id int32, type_ string, accentColorId int32, maxReactionCount int32, acceptedGiftTypes AcceptedGiftTypes) *ChatFullInfo {
	this := ChatFullInfo{}
	this.Id = id
	this.Type = type_
	var isForum bool = true
	this.IsForum = &isForum
	this.AccentColorId = accentColorId
	this.MaxReactionCount = maxReactionCount
	var hasPrivateForwards bool = true
	this.HasPrivateForwards = &hasPrivateForwards
	var hasRestrictedVoiceAndVideoMessages bool = true
	this.HasRestrictedVoiceAndVideoMessages = &hasRestrictedVoiceAndVideoMessages
	var joinToSendMessages bool = true
	this.JoinToSendMessages = &joinToSendMessages
	var joinByRequest bool = true
	this.JoinByRequest = &joinByRequest
	this.AcceptedGiftTypes = acceptedGiftTypes
	var canSendPaidMedia bool = true
	this.CanSendPaidMedia = &canSendPaidMedia
	var hasAggressiveAntiSpamEnabled bool = true
	this.HasAggressiveAntiSpamEnabled = &hasAggressiveAntiSpamEnabled
	var hasHiddenMembers bool = true
	this.HasHiddenMembers = &hasHiddenMembers
	var hasProtectedContent bool = true
	this.HasProtectedContent = &hasProtectedContent
	var hasVisibleHistory bool = true
	this.HasVisibleHistory = &hasVisibleHistory
	var canSetStickerSet bool = true
	this.CanSetStickerSet = &canSetStickerSet
	return &this
}

// NewChatFullInfoWithDefaults instantiates a new ChatFullInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatFullInfoWithDefaults() *ChatFullInfo {
	this := ChatFullInfo{}
	var isForum bool = true
	this.IsForum = &isForum
	var hasPrivateForwards bool = true
	this.HasPrivateForwards = &hasPrivateForwards
	var hasRestrictedVoiceAndVideoMessages bool = true
	this.HasRestrictedVoiceAndVideoMessages = &hasRestrictedVoiceAndVideoMessages
	var joinToSendMessages bool = true
	this.JoinToSendMessages = &joinToSendMessages
	var joinByRequest bool = true
	this.JoinByRequest = &joinByRequest
	var canSendPaidMedia bool = true
	this.CanSendPaidMedia = &canSendPaidMedia
	var hasAggressiveAntiSpamEnabled bool = true
	this.HasAggressiveAntiSpamEnabled = &hasAggressiveAntiSpamEnabled
	var hasHiddenMembers bool = true
	this.HasHiddenMembers = &hasHiddenMembers
	var hasProtectedContent bool = true
	this.HasProtectedContent = &hasProtectedContent
	var hasVisibleHistory bool = true
	this.HasVisibleHistory = &hasVisibleHistory
	var canSetStickerSet bool = true
	this.CanSetStickerSet = &canSetStickerSet
	return &this
}

// GetId returns the Id field value
func (o *ChatFullInfo) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatFullInfo) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ChatFullInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChatFullInfo) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ChatFullInfo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ChatFullInfo) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ChatFullInfo) SetTitle(v string) {
	o.Title = &v
}


// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ChatFullInfo) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ChatFullInfo) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ChatFullInfo) SetUsername(v string) {
	o.Username = &v
}


// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ChatFullInfo) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ChatFullInfo) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ChatFullInfo) SetFirstName(v string) {
	o.FirstName = &v
}


// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ChatFullInfo) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ChatFullInfo) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ChatFullInfo) SetLastName(v string) {
	o.LastName = &v
}


// GetIsForum returns the IsForum field value if set, zero value otherwise.
func (o *ChatFullInfo) GetIsForum() bool {
	if o == nil || IsNil(o.IsForum) {
		var ret bool
		return ret
	}
	return *o.IsForum
}

// GetIsForumOk returns a tuple with the IsForum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetIsForumOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForum) {
		return nil, false
	}
	return o.IsForum, true
}

// HasIsForum returns a boolean if a field has been set.
func (o *ChatFullInfo) HasIsForum() bool {
	if o != nil && !IsNil(o.IsForum) {
		return true
	}

	return false
}

// SetIsForum gets a reference to the given bool and assigns it to the IsForum field.
func (o *ChatFullInfo) SetIsForum(v bool) {
	o.IsForum = &v
}


// GetAccentColorId returns the AccentColorId field value
func (o *ChatFullInfo) GetAccentColorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccentColorId
}

// GetAccentColorIdOk returns a tuple with the AccentColorId field value
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetAccentColorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccentColorId, true
}

// SetAccentColorId sets field value
func (o *ChatFullInfo) SetAccentColorId(v int32) {
	o.AccentColorId = v
}

// GetMaxReactionCount returns the MaxReactionCount field value
func (o *ChatFullInfo) GetMaxReactionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxReactionCount
}

// GetMaxReactionCountOk returns a tuple with the MaxReactionCount field value
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetMaxReactionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxReactionCount, true
}

// SetMaxReactionCount sets field value
func (o *ChatFullInfo) SetMaxReactionCount(v int32) {
	o.MaxReactionCount = v
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *ChatFullInfo) GetPhoto() ChatPhoto {
	if o == nil || IsNil(o.Photo) {
		var ret ChatPhoto
		return ret
	}
	return *o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetPhotoOk() (*ChatPhoto, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *ChatFullInfo) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given ChatPhoto and assigns it to the Photo field.
func (o *ChatFullInfo) SetPhoto(v ChatPhoto) {
	o.Photo = &v
}


// GetActiveUsernames returns the ActiveUsernames field value if set, zero value otherwise.
func (o *ChatFullInfo) GetActiveUsernames() []string {
	if o == nil || IsNil(o.ActiveUsernames) {
		var ret []string
		return ret
	}
	return o.ActiveUsernames
}

// GetActiveUsernamesOk returns a tuple with the ActiveUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetActiveUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ActiveUsernames) {
		return nil, false
	}
	return o.ActiveUsernames, true
}

// HasActiveUsernames returns a boolean if a field has been set.
func (o *ChatFullInfo) HasActiveUsernames() bool {
	if o != nil && !IsNil(o.ActiveUsernames) {
		return true
	}

	return false
}

// SetActiveUsernames gets a reference to the given []string and assigns it to the ActiveUsernames field.
func (o *ChatFullInfo) SetActiveUsernames(v []string) {
	o.ActiveUsernames = v
}


// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBirthdate() Birthdate {
	if o == nil || IsNil(o.Birthdate) {
		var ret Birthdate
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBirthdateOk() (*Birthdate, bool) {
	if o == nil || IsNil(o.Birthdate) {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBirthdate() bool {
	if o != nil && !IsNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given Birthdate and assigns it to the Birthdate field.
func (o *ChatFullInfo) SetBirthdate(v Birthdate) {
	o.Birthdate = &v
}


// GetBusinessIntro returns the BusinessIntro field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBusinessIntro() BusinessIntro {
	if o == nil || IsNil(o.BusinessIntro) {
		var ret BusinessIntro
		return ret
	}
	return *o.BusinessIntro
}

// GetBusinessIntroOk returns a tuple with the BusinessIntro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBusinessIntroOk() (*BusinessIntro, bool) {
	if o == nil || IsNil(o.BusinessIntro) {
		return nil, false
	}
	return o.BusinessIntro, true
}

// HasBusinessIntro returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBusinessIntro() bool {
	if o != nil && !IsNil(o.BusinessIntro) {
		return true
	}

	return false
}

// SetBusinessIntro gets a reference to the given BusinessIntro and assigns it to the BusinessIntro field.
func (o *ChatFullInfo) SetBusinessIntro(v BusinessIntro) {
	o.BusinessIntro = &v
}


// GetBusinessLocation returns the BusinessLocation field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBusinessLocation() BusinessLocation {
	if o == nil || IsNil(o.BusinessLocation) {
		var ret BusinessLocation
		return ret
	}
	return *o.BusinessLocation
}

// GetBusinessLocationOk returns a tuple with the BusinessLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBusinessLocationOk() (*BusinessLocation, bool) {
	if o == nil || IsNil(o.BusinessLocation) {
		return nil, false
	}
	return o.BusinessLocation, true
}

// HasBusinessLocation returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBusinessLocation() bool {
	if o != nil && !IsNil(o.BusinessLocation) {
		return true
	}

	return false
}

// SetBusinessLocation gets a reference to the given BusinessLocation and assigns it to the BusinessLocation field.
func (o *ChatFullInfo) SetBusinessLocation(v BusinessLocation) {
	o.BusinessLocation = &v
}


// GetBusinessOpeningHours returns the BusinessOpeningHours field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBusinessOpeningHours() BusinessOpeningHours {
	if o == nil || IsNil(o.BusinessOpeningHours) {
		var ret BusinessOpeningHours
		return ret
	}
	return *o.BusinessOpeningHours
}

// GetBusinessOpeningHoursOk returns a tuple with the BusinessOpeningHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBusinessOpeningHoursOk() (*BusinessOpeningHours, bool) {
	if o == nil || IsNil(o.BusinessOpeningHours) {
		return nil, false
	}
	return o.BusinessOpeningHours, true
}

// HasBusinessOpeningHours returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBusinessOpeningHours() bool {
	if o != nil && !IsNil(o.BusinessOpeningHours) {
		return true
	}

	return false
}

// SetBusinessOpeningHours gets a reference to the given BusinessOpeningHours and assigns it to the BusinessOpeningHours field.
func (o *ChatFullInfo) SetBusinessOpeningHours(v BusinessOpeningHours) {
	o.BusinessOpeningHours = &v
}


// GetPersonalChat returns the PersonalChat field value if set, zero value otherwise.
func (o *ChatFullInfo) GetPersonalChat() Chat {
	if o == nil || IsNil(o.PersonalChat) {
		var ret Chat
		return ret
	}
	return *o.PersonalChat
}

// GetPersonalChatOk returns a tuple with the PersonalChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetPersonalChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.PersonalChat) {
		return nil, false
	}
	return o.PersonalChat, true
}

// HasPersonalChat returns a boolean if a field has been set.
func (o *ChatFullInfo) HasPersonalChat() bool {
	if o != nil && !IsNil(o.PersonalChat) {
		return true
	}

	return false
}

// SetPersonalChat gets a reference to the given Chat and assigns it to the PersonalChat field.
func (o *ChatFullInfo) SetPersonalChat(v Chat) {
	o.PersonalChat = &v
}


// GetAvailableReactions returns the AvailableReactions field value if set, zero value otherwise.
func (o *ChatFullInfo) GetAvailableReactions() []ReactionType {
	if o == nil || IsNil(o.AvailableReactions) {
		var ret []ReactionType
		return ret
	}
	return o.AvailableReactions
}

// GetAvailableReactionsOk returns a tuple with the AvailableReactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetAvailableReactionsOk() ([]ReactionType, bool) {
	if o == nil || IsNil(o.AvailableReactions) {
		return nil, false
	}
	return o.AvailableReactions, true
}

// HasAvailableReactions returns a boolean if a field has been set.
func (o *ChatFullInfo) HasAvailableReactions() bool {
	if o != nil && !IsNil(o.AvailableReactions) {
		return true
	}

	return false
}

// SetAvailableReactions gets a reference to the given []ReactionType and assigns it to the AvailableReactions field.
func (o *ChatFullInfo) SetAvailableReactions(v []ReactionType) {
	o.AvailableReactions = v
}


// GetBackgroundCustomEmojiId returns the BackgroundCustomEmojiId field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBackgroundCustomEmojiId() string {
	if o == nil || IsNil(o.BackgroundCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.BackgroundCustomEmojiId
}

// GetBackgroundCustomEmojiIdOk returns a tuple with the BackgroundCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBackgroundCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundCustomEmojiId) {
		return nil, false
	}
	return o.BackgroundCustomEmojiId, true
}

// HasBackgroundCustomEmojiId returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBackgroundCustomEmojiId() bool {
	if o != nil && !IsNil(o.BackgroundCustomEmojiId) {
		return true
	}

	return false
}

// SetBackgroundCustomEmojiId gets a reference to the given string and assigns it to the BackgroundCustomEmojiId field.
func (o *ChatFullInfo) SetBackgroundCustomEmojiId(v string) {
	o.BackgroundCustomEmojiId = &v
}


// GetProfileAccentColorId returns the ProfileAccentColorId field value if set, zero value otherwise.
func (o *ChatFullInfo) GetProfileAccentColorId() int32 {
	if o == nil || IsNil(o.ProfileAccentColorId) {
		var ret int32
		return ret
	}
	return *o.ProfileAccentColorId
}

// GetProfileAccentColorIdOk returns a tuple with the ProfileAccentColorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetProfileAccentColorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProfileAccentColorId) {
		return nil, false
	}
	return o.ProfileAccentColorId, true
}

// HasProfileAccentColorId returns a boolean if a field has been set.
func (o *ChatFullInfo) HasProfileAccentColorId() bool {
	if o != nil && !IsNil(o.ProfileAccentColorId) {
		return true
	}

	return false
}

// SetProfileAccentColorId gets a reference to the given int32 and assigns it to the ProfileAccentColorId field.
func (o *ChatFullInfo) SetProfileAccentColorId(v int32) {
	o.ProfileAccentColorId = &v
}


// GetProfileBackgroundCustomEmojiId returns the ProfileBackgroundCustomEmojiId field value if set, zero value otherwise.
func (o *ChatFullInfo) GetProfileBackgroundCustomEmojiId() string {
	if o == nil || IsNil(o.ProfileBackgroundCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.ProfileBackgroundCustomEmojiId
}

// GetProfileBackgroundCustomEmojiIdOk returns a tuple with the ProfileBackgroundCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetProfileBackgroundCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileBackgroundCustomEmojiId) {
		return nil, false
	}
	return o.ProfileBackgroundCustomEmojiId, true
}

// HasProfileBackgroundCustomEmojiId returns a boolean if a field has been set.
func (o *ChatFullInfo) HasProfileBackgroundCustomEmojiId() bool {
	if o != nil && !IsNil(o.ProfileBackgroundCustomEmojiId) {
		return true
	}

	return false
}

// SetProfileBackgroundCustomEmojiId gets a reference to the given string and assigns it to the ProfileBackgroundCustomEmojiId field.
func (o *ChatFullInfo) SetProfileBackgroundCustomEmojiId(v string) {
	o.ProfileBackgroundCustomEmojiId = &v
}


// GetEmojiStatusCustomEmojiId returns the EmojiStatusCustomEmojiId field value if set, zero value otherwise.
func (o *ChatFullInfo) GetEmojiStatusCustomEmojiId() string {
	if o == nil || IsNil(o.EmojiStatusCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiStatusCustomEmojiId
}

// GetEmojiStatusCustomEmojiIdOk returns a tuple with the EmojiStatusCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetEmojiStatusCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiStatusCustomEmojiId) {
		return nil, false
	}
	return o.EmojiStatusCustomEmojiId, true
}

// HasEmojiStatusCustomEmojiId returns a boolean if a field has been set.
func (o *ChatFullInfo) HasEmojiStatusCustomEmojiId() bool {
	if o != nil && !IsNil(o.EmojiStatusCustomEmojiId) {
		return true
	}

	return false
}

// SetEmojiStatusCustomEmojiId gets a reference to the given string and assigns it to the EmojiStatusCustomEmojiId field.
func (o *ChatFullInfo) SetEmojiStatusCustomEmojiId(v string) {
	o.EmojiStatusCustomEmojiId = &v
}


// GetEmojiStatusExpirationDate returns the EmojiStatusExpirationDate field value if set, zero value otherwise.
func (o *ChatFullInfo) GetEmojiStatusExpirationDate() int32 {
	if o == nil || IsNil(o.EmojiStatusExpirationDate) {
		var ret int32
		return ret
	}
	return *o.EmojiStatusExpirationDate
}

// GetEmojiStatusExpirationDateOk returns a tuple with the EmojiStatusExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetEmojiStatusExpirationDateOk() (*int32, bool) {
	if o == nil || IsNil(o.EmojiStatusExpirationDate) {
		return nil, false
	}
	return o.EmojiStatusExpirationDate, true
}

// HasEmojiStatusExpirationDate returns a boolean if a field has been set.
func (o *ChatFullInfo) HasEmojiStatusExpirationDate() bool {
	if o != nil && !IsNil(o.EmojiStatusExpirationDate) {
		return true
	}

	return false
}

// SetEmojiStatusExpirationDate gets a reference to the given int32 and assigns it to the EmojiStatusExpirationDate field.
func (o *ChatFullInfo) SetEmojiStatusExpirationDate(v int32) {
	o.EmojiStatusExpirationDate = &v
}


// GetBio returns the Bio field value if set, zero value otherwise.
func (o *ChatFullInfo) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *ChatFullInfo) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *ChatFullInfo) SetBio(v string) {
	o.Bio = &v
}


// GetHasPrivateForwards returns the HasPrivateForwards field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasPrivateForwards() bool {
	if o == nil || IsNil(o.HasPrivateForwards) {
		var ret bool
		return ret
	}
	return *o.HasPrivateForwards
}

// GetHasPrivateForwardsOk returns a tuple with the HasPrivateForwards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasPrivateForwardsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPrivateForwards) {
		return nil, false
	}
	return o.HasPrivateForwards, true
}

// HasHasPrivateForwards returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasPrivateForwards() bool {
	if o != nil && !IsNil(o.HasPrivateForwards) {
		return true
	}

	return false
}

// SetHasPrivateForwards gets a reference to the given bool and assigns it to the HasPrivateForwards field.
func (o *ChatFullInfo) SetHasPrivateForwards(v bool) {
	o.HasPrivateForwards = &v
}


// GetHasRestrictedVoiceAndVideoMessages returns the HasRestrictedVoiceAndVideoMessages field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasRestrictedVoiceAndVideoMessages() bool {
	if o == nil || IsNil(o.HasRestrictedVoiceAndVideoMessages) {
		var ret bool
		return ret
	}
	return *o.HasRestrictedVoiceAndVideoMessages
}

// GetHasRestrictedVoiceAndVideoMessagesOk returns a tuple with the HasRestrictedVoiceAndVideoMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasRestrictedVoiceAndVideoMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRestrictedVoiceAndVideoMessages) {
		return nil, false
	}
	return o.HasRestrictedVoiceAndVideoMessages, true
}

// HasHasRestrictedVoiceAndVideoMessages returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasRestrictedVoiceAndVideoMessages() bool {
	if o != nil && !IsNil(o.HasRestrictedVoiceAndVideoMessages) {
		return true
	}

	return false
}

// SetHasRestrictedVoiceAndVideoMessages gets a reference to the given bool and assigns it to the HasRestrictedVoiceAndVideoMessages field.
func (o *ChatFullInfo) SetHasRestrictedVoiceAndVideoMessages(v bool) {
	o.HasRestrictedVoiceAndVideoMessages = &v
}


// GetJoinToSendMessages returns the JoinToSendMessages field value if set, zero value otherwise.
func (o *ChatFullInfo) GetJoinToSendMessages() bool {
	if o == nil || IsNil(o.JoinToSendMessages) {
		var ret bool
		return ret
	}
	return *o.JoinToSendMessages
}

// GetJoinToSendMessagesOk returns a tuple with the JoinToSendMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetJoinToSendMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.JoinToSendMessages) {
		return nil, false
	}
	return o.JoinToSendMessages, true
}

// HasJoinToSendMessages returns a boolean if a field has been set.
func (o *ChatFullInfo) HasJoinToSendMessages() bool {
	if o != nil && !IsNil(o.JoinToSendMessages) {
		return true
	}

	return false
}

// SetJoinToSendMessages gets a reference to the given bool and assigns it to the JoinToSendMessages field.
func (o *ChatFullInfo) SetJoinToSendMessages(v bool) {
	o.JoinToSendMessages = &v
}


// GetJoinByRequest returns the JoinByRequest field value if set, zero value otherwise.
func (o *ChatFullInfo) GetJoinByRequest() bool {
	if o == nil || IsNil(o.JoinByRequest) {
		var ret bool
		return ret
	}
	return *o.JoinByRequest
}

// GetJoinByRequestOk returns a tuple with the JoinByRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetJoinByRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.JoinByRequest) {
		return nil, false
	}
	return o.JoinByRequest, true
}

// HasJoinByRequest returns a boolean if a field has been set.
func (o *ChatFullInfo) HasJoinByRequest() bool {
	if o != nil && !IsNil(o.JoinByRequest) {
		return true
	}

	return false
}

// SetJoinByRequest gets a reference to the given bool and assigns it to the JoinByRequest field.
func (o *ChatFullInfo) SetJoinByRequest(v bool) {
	o.JoinByRequest = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChatFullInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChatFullInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChatFullInfo) SetDescription(v string) {
	o.Description = &v
}


// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
func (o *ChatFullInfo) GetInviteLink() string {
	if o == nil || IsNil(o.InviteLink) {
		var ret string
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetInviteLinkOk() (*string, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *ChatFullInfo) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given string and assigns it to the InviteLink field.
func (o *ChatFullInfo) SetInviteLink(v string) {
	o.InviteLink = &v
}


// GetPinnedMessage returns the PinnedMessage field value if set, zero value otherwise.
func (o *ChatFullInfo) GetPinnedMessage() Message {
	if o == nil || IsNil(o.PinnedMessage) {
		var ret Message
		return ret
	}
	return *o.PinnedMessage
}

// GetPinnedMessageOk returns a tuple with the PinnedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetPinnedMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.PinnedMessage) {
		return nil, false
	}
	return o.PinnedMessage, true
}

// HasPinnedMessage returns a boolean if a field has been set.
func (o *ChatFullInfo) HasPinnedMessage() bool {
	if o != nil && !IsNil(o.PinnedMessage) {
		return true
	}

	return false
}

// SetPinnedMessage gets a reference to the given Message and assigns it to the PinnedMessage field.
func (o *ChatFullInfo) SetPinnedMessage(v Message) {
	o.PinnedMessage = &v
}


// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ChatFullInfo) GetPermissions() ChatPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret ChatPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetPermissionsOk() (*ChatPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ChatFullInfo) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ChatPermissions and assigns it to the Permissions field.
func (o *ChatFullInfo) SetPermissions(v ChatPermissions) {
	o.Permissions = &v
}


// GetAcceptedGiftTypes returns the AcceptedGiftTypes field value
func (o *ChatFullInfo) GetAcceptedGiftTypes() AcceptedGiftTypes {
	if o == nil {
		var ret AcceptedGiftTypes
		return ret
	}

	return o.AcceptedGiftTypes
}

// GetAcceptedGiftTypesOk returns a tuple with the AcceptedGiftTypes field value
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetAcceptedGiftTypesOk() (*AcceptedGiftTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedGiftTypes, true
}

// SetAcceptedGiftTypes sets field value
func (o *ChatFullInfo) SetAcceptedGiftTypes(v AcceptedGiftTypes) {
	o.AcceptedGiftTypes = v
}

// GetCanSendPaidMedia returns the CanSendPaidMedia field value if set, zero value otherwise.
func (o *ChatFullInfo) GetCanSendPaidMedia() bool {
	if o == nil || IsNil(o.CanSendPaidMedia) {
		var ret bool
		return ret
	}
	return *o.CanSendPaidMedia
}

// GetCanSendPaidMediaOk returns a tuple with the CanSendPaidMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetCanSendPaidMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSendPaidMedia) {
		return nil, false
	}
	return o.CanSendPaidMedia, true
}

// HasCanSendPaidMedia returns a boolean if a field has been set.
func (o *ChatFullInfo) HasCanSendPaidMedia() bool {
	if o != nil && !IsNil(o.CanSendPaidMedia) {
		return true
	}

	return false
}

// SetCanSendPaidMedia gets a reference to the given bool and assigns it to the CanSendPaidMedia field.
func (o *ChatFullInfo) SetCanSendPaidMedia(v bool) {
	o.CanSendPaidMedia = &v
}


// GetSlowModeDelay returns the SlowModeDelay field value if set, zero value otherwise.
func (o *ChatFullInfo) GetSlowModeDelay() int32 {
	if o == nil || IsNil(o.SlowModeDelay) {
		var ret int32
		return ret
	}
	return *o.SlowModeDelay
}

// GetSlowModeDelayOk returns a tuple with the SlowModeDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetSlowModeDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.SlowModeDelay) {
		return nil, false
	}
	return o.SlowModeDelay, true
}

// HasSlowModeDelay returns a boolean if a field has been set.
func (o *ChatFullInfo) HasSlowModeDelay() bool {
	if o != nil && !IsNil(o.SlowModeDelay) {
		return true
	}

	return false
}

// SetSlowModeDelay gets a reference to the given int32 and assigns it to the SlowModeDelay field.
func (o *ChatFullInfo) SetSlowModeDelay(v int32) {
	o.SlowModeDelay = &v
}


// GetUnrestrictBoostCount returns the UnrestrictBoostCount field value if set, zero value otherwise.
func (o *ChatFullInfo) GetUnrestrictBoostCount() int32 {
	if o == nil || IsNil(o.UnrestrictBoostCount) {
		var ret int32
		return ret
	}
	return *o.UnrestrictBoostCount
}

// GetUnrestrictBoostCountOk returns a tuple with the UnrestrictBoostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetUnrestrictBoostCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnrestrictBoostCount) {
		return nil, false
	}
	return o.UnrestrictBoostCount, true
}

// HasUnrestrictBoostCount returns a boolean if a field has been set.
func (o *ChatFullInfo) HasUnrestrictBoostCount() bool {
	if o != nil && !IsNil(o.UnrestrictBoostCount) {
		return true
	}

	return false
}

// SetUnrestrictBoostCount gets a reference to the given int32 and assigns it to the UnrestrictBoostCount field.
func (o *ChatFullInfo) SetUnrestrictBoostCount(v int32) {
	o.UnrestrictBoostCount = &v
}


// GetMessageAutoDeleteTime returns the MessageAutoDeleteTime field value if set, zero value otherwise.
func (o *ChatFullInfo) GetMessageAutoDeleteTime() int32 {
	if o == nil || IsNil(o.MessageAutoDeleteTime) {
		var ret int32
		return ret
	}
	return *o.MessageAutoDeleteTime
}

// GetMessageAutoDeleteTimeOk returns a tuple with the MessageAutoDeleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetMessageAutoDeleteTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageAutoDeleteTime) {
		return nil, false
	}
	return o.MessageAutoDeleteTime, true
}

// HasMessageAutoDeleteTime returns a boolean if a field has been set.
func (o *ChatFullInfo) HasMessageAutoDeleteTime() bool {
	if o != nil && !IsNil(o.MessageAutoDeleteTime) {
		return true
	}

	return false
}

// SetMessageAutoDeleteTime gets a reference to the given int32 and assigns it to the MessageAutoDeleteTime field.
func (o *ChatFullInfo) SetMessageAutoDeleteTime(v int32) {
	o.MessageAutoDeleteTime = &v
}


// GetHasAggressiveAntiSpamEnabled returns the HasAggressiveAntiSpamEnabled field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasAggressiveAntiSpamEnabled() bool {
	if o == nil || IsNil(o.HasAggressiveAntiSpamEnabled) {
		var ret bool
		return ret
	}
	return *o.HasAggressiveAntiSpamEnabled
}

// GetHasAggressiveAntiSpamEnabledOk returns a tuple with the HasAggressiveAntiSpamEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasAggressiveAntiSpamEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAggressiveAntiSpamEnabled) {
		return nil, false
	}
	return o.HasAggressiveAntiSpamEnabled, true
}

// HasHasAggressiveAntiSpamEnabled returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasAggressiveAntiSpamEnabled() bool {
	if o != nil && !IsNil(o.HasAggressiveAntiSpamEnabled) {
		return true
	}

	return false
}

// SetHasAggressiveAntiSpamEnabled gets a reference to the given bool and assigns it to the HasAggressiveAntiSpamEnabled field.
func (o *ChatFullInfo) SetHasAggressiveAntiSpamEnabled(v bool) {
	o.HasAggressiveAntiSpamEnabled = &v
}


// GetHasHiddenMembers returns the HasHiddenMembers field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasHiddenMembers() bool {
	if o == nil || IsNil(o.HasHiddenMembers) {
		var ret bool
		return ret
	}
	return *o.HasHiddenMembers
}

// GetHasHiddenMembersOk returns a tuple with the HasHiddenMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasHiddenMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHiddenMembers) {
		return nil, false
	}
	return o.HasHiddenMembers, true
}

// HasHasHiddenMembers returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasHiddenMembers() bool {
	if o != nil && !IsNil(o.HasHiddenMembers) {
		return true
	}

	return false
}

// SetHasHiddenMembers gets a reference to the given bool and assigns it to the HasHiddenMembers field.
func (o *ChatFullInfo) SetHasHiddenMembers(v bool) {
	o.HasHiddenMembers = &v
}


// GetHasProtectedContent returns the HasProtectedContent field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasProtectedContent() bool {
	if o == nil || IsNil(o.HasProtectedContent) {
		var ret bool
		return ret
	}
	return *o.HasProtectedContent
}

// GetHasProtectedContentOk returns a tuple with the HasProtectedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasProtectedContentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasProtectedContent) {
		return nil, false
	}
	return o.HasProtectedContent, true
}

// HasHasProtectedContent returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasProtectedContent() bool {
	if o != nil && !IsNil(o.HasProtectedContent) {
		return true
	}

	return false
}

// SetHasProtectedContent gets a reference to the given bool and assigns it to the HasProtectedContent field.
func (o *ChatFullInfo) SetHasProtectedContent(v bool) {
	o.HasProtectedContent = &v
}


// GetHasVisibleHistory returns the HasVisibleHistory field value if set, zero value otherwise.
func (o *ChatFullInfo) GetHasVisibleHistory() bool {
	if o == nil || IsNil(o.HasVisibleHistory) {
		var ret bool
		return ret
	}
	return *o.HasVisibleHistory
}

// GetHasVisibleHistoryOk returns a tuple with the HasVisibleHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetHasVisibleHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.HasVisibleHistory) {
		return nil, false
	}
	return o.HasVisibleHistory, true
}

// HasHasVisibleHistory returns a boolean if a field has been set.
func (o *ChatFullInfo) HasHasVisibleHistory() bool {
	if o != nil && !IsNil(o.HasVisibleHistory) {
		return true
	}

	return false
}

// SetHasVisibleHistory gets a reference to the given bool and assigns it to the HasVisibleHistory field.
func (o *ChatFullInfo) SetHasVisibleHistory(v bool) {
	o.HasVisibleHistory = &v
}


// GetStickerSetName returns the StickerSetName field value if set, zero value otherwise.
func (o *ChatFullInfo) GetStickerSetName() string {
	if o == nil || IsNil(o.StickerSetName) {
		var ret string
		return ret
	}
	return *o.StickerSetName
}

// GetStickerSetNameOk returns a tuple with the StickerSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetStickerSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.StickerSetName) {
		return nil, false
	}
	return o.StickerSetName, true
}

// HasStickerSetName returns a boolean if a field has been set.
func (o *ChatFullInfo) HasStickerSetName() bool {
	if o != nil && !IsNil(o.StickerSetName) {
		return true
	}

	return false
}

// SetStickerSetName gets a reference to the given string and assigns it to the StickerSetName field.
func (o *ChatFullInfo) SetStickerSetName(v string) {
	o.StickerSetName = &v
}


// GetCanSetStickerSet returns the CanSetStickerSet field value if set, zero value otherwise.
func (o *ChatFullInfo) GetCanSetStickerSet() bool {
	if o == nil || IsNil(o.CanSetStickerSet) {
		var ret bool
		return ret
	}
	return *o.CanSetStickerSet
}

// GetCanSetStickerSetOk returns a tuple with the CanSetStickerSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetCanSetStickerSetOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSetStickerSet) {
		return nil, false
	}
	return o.CanSetStickerSet, true
}

// HasCanSetStickerSet returns a boolean if a field has been set.
func (o *ChatFullInfo) HasCanSetStickerSet() bool {
	if o != nil && !IsNil(o.CanSetStickerSet) {
		return true
	}

	return false
}

// SetCanSetStickerSet gets a reference to the given bool and assigns it to the CanSetStickerSet field.
func (o *ChatFullInfo) SetCanSetStickerSet(v bool) {
	o.CanSetStickerSet = &v
}


// GetCustomEmojiStickerSetName returns the CustomEmojiStickerSetName field value if set, zero value otherwise.
func (o *ChatFullInfo) GetCustomEmojiStickerSetName() string {
	if o == nil || IsNil(o.CustomEmojiStickerSetName) {
		var ret string
		return ret
	}
	return *o.CustomEmojiStickerSetName
}

// GetCustomEmojiStickerSetNameOk returns a tuple with the CustomEmojiStickerSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetCustomEmojiStickerSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEmojiStickerSetName) {
		return nil, false
	}
	return o.CustomEmojiStickerSetName, true
}

// HasCustomEmojiStickerSetName returns a boolean if a field has been set.
func (o *ChatFullInfo) HasCustomEmojiStickerSetName() bool {
	if o != nil && !IsNil(o.CustomEmojiStickerSetName) {
		return true
	}

	return false
}

// SetCustomEmojiStickerSetName gets a reference to the given string and assigns it to the CustomEmojiStickerSetName field.
func (o *ChatFullInfo) SetCustomEmojiStickerSetName(v string) {
	o.CustomEmojiStickerSetName = &v
}


// GetLinkedChatId returns the LinkedChatId field value if set, zero value otherwise.
func (o *ChatFullInfo) GetLinkedChatId() int32 {
	if o == nil || IsNil(o.LinkedChatId) {
		var ret int32
		return ret
	}
	return *o.LinkedChatId
}

// GetLinkedChatIdOk returns a tuple with the LinkedChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetLinkedChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkedChatId) {
		return nil, false
	}
	return o.LinkedChatId, true
}

// HasLinkedChatId returns a boolean if a field has been set.
func (o *ChatFullInfo) HasLinkedChatId() bool {
	if o != nil && !IsNil(o.LinkedChatId) {
		return true
	}

	return false
}

// SetLinkedChatId gets a reference to the given int32 and assigns it to the LinkedChatId field.
func (o *ChatFullInfo) SetLinkedChatId(v int32) {
	o.LinkedChatId = &v
}


// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ChatFullInfo) GetLocation() ChatLocation {
	if o == nil || IsNil(o.Location) {
		var ret ChatLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatFullInfo) GetLocationOk() (*ChatLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ChatFullInfo) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ChatLocation and assigns it to the Location field.
func (o *ChatFullInfo) SetLocation(v ChatLocation) {
	o.Location = &v
}


func (o ChatFullInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatFullInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.IsForum) {
		toSerialize["is_forum"] = o.IsForum
	}
	toSerialize["accent_color_id"] = o.AccentColorId
	toSerialize["max_reaction_count"] = o.MaxReactionCount
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.ActiveUsernames) {
		toSerialize["active_usernames"] = o.ActiveUsernames
	}
	if !IsNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !IsNil(o.BusinessIntro) {
		toSerialize["business_intro"] = o.BusinessIntro
	}
	if !IsNil(o.BusinessLocation) {
		toSerialize["business_location"] = o.BusinessLocation
	}
	if !IsNil(o.BusinessOpeningHours) {
		toSerialize["business_opening_hours"] = o.BusinessOpeningHours
	}
	if !IsNil(o.PersonalChat) {
		toSerialize["personal_chat"] = o.PersonalChat
	}
	if !IsNil(o.AvailableReactions) {
		toSerialize["available_reactions"] = o.AvailableReactions
	}
	if !IsNil(o.BackgroundCustomEmojiId) {
		toSerialize["background_custom_emoji_id"] = o.BackgroundCustomEmojiId
	}
	if !IsNil(o.ProfileAccentColorId) {
		toSerialize["profile_accent_color_id"] = o.ProfileAccentColorId
	}
	if !IsNil(o.ProfileBackgroundCustomEmojiId) {
		toSerialize["profile_background_custom_emoji_id"] = o.ProfileBackgroundCustomEmojiId
	}
	if !IsNil(o.EmojiStatusCustomEmojiId) {
		toSerialize["emoji_status_custom_emoji_id"] = o.EmojiStatusCustomEmojiId
	}
	if !IsNil(o.EmojiStatusExpirationDate) {
		toSerialize["emoji_status_expiration_date"] = o.EmojiStatusExpirationDate
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.HasPrivateForwards) {
		toSerialize["has_private_forwards"] = o.HasPrivateForwards
	}
	if !IsNil(o.HasRestrictedVoiceAndVideoMessages) {
		toSerialize["has_restricted_voice_and_video_messages"] = o.HasRestrictedVoiceAndVideoMessages
	}
	if !IsNil(o.JoinToSendMessages) {
		toSerialize["join_to_send_messages"] = o.JoinToSendMessages
	}
	if !IsNil(o.JoinByRequest) {
		toSerialize["join_by_request"] = o.JoinByRequest
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InviteLink) {
		toSerialize["invite_link"] = o.InviteLink
	}
	if !IsNil(o.PinnedMessage) {
		toSerialize["pinned_message"] = o.PinnedMessage
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	toSerialize["accepted_gift_types"] = o.AcceptedGiftTypes
	if !IsNil(o.CanSendPaidMedia) {
		toSerialize["can_send_paid_media"] = o.CanSendPaidMedia
	}
	if !IsNil(o.SlowModeDelay) {
		toSerialize["slow_mode_delay"] = o.SlowModeDelay
	}
	if !IsNil(o.UnrestrictBoostCount) {
		toSerialize["unrestrict_boost_count"] = o.UnrestrictBoostCount
	}
	if !IsNil(o.MessageAutoDeleteTime) {
		toSerialize["message_auto_delete_time"] = o.MessageAutoDeleteTime
	}
	if !IsNil(o.HasAggressiveAntiSpamEnabled) {
		toSerialize["has_aggressive_anti_spam_enabled"] = o.HasAggressiveAntiSpamEnabled
	}
	if !IsNil(o.HasHiddenMembers) {
		toSerialize["has_hidden_members"] = o.HasHiddenMembers
	}
	if !IsNil(o.HasProtectedContent) {
		toSerialize["has_protected_content"] = o.HasProtectedContent
	}
	if !IsNil(o.HasVisibleHistory) {
		toSerialize["has_visible_history"] = o.HasVisibleHistory
	}
	if !IsNil(o.StickerSetName) {
		toSerialize["sticker_set_name"] = o.StickerSetName
	}
	if !IsNil(o.CanSetStickerSet) {
		toSerialize["can_set_sticker_set"] = o.CanSetStickerSet
	}
	if !IsNil(o.CustomEmojiStickerSetName) {
		toSerialize["custom_emoji_sticker_set_name"] = o.CustomEmojiStickerSetName
	}
	if !IsNil(o.LinkedChatId) {
		toSerialize["linked_chat_id"] = o.LinkedChatId
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

func (o *ChatFullInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"accent_color_id",
		"max_reaction_count",
		"accepted_gift_types",
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

	varChatFullInfo := _ChatFullInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatFullInfo)

	if err != nil {
		return err
	}

	*o = ChatFullInfo(varChatFullInfo)

	return err
}

type NullableChatFullInfo struct {
	value *ChatFullInfo
	isSet bool
}

func (v NullableChatFullInfo) Get() *ChatFullInfo {
	return v.value
}

func (v *NullableChatFullInfo) Set(val *ChatFullInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChatFullInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChatFullInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatFullInfo(val *ChatFullInfo) *NullableChatFullInfo {
	return &NullableChatFullInfo{value: val, isSet: true}
}

func (v NullableChatFullInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatFullInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


