/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message This object represents a message.
type Message struct {
	// Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageId int32 `json:"message_id"`
	// *Optional*. Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	From *User `json:"from,omitempty"`
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// *Optional*. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBoostCount *int32 `json:"sender_boost_count,omitempty"`
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`
	// Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int32 `json:"date"`
	// *Optional*. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	Chat Chat `json:"chat"`
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`
	// *Optional*. *True*, if the message is sent to a forum topic
	IsTopicMessage *bool `json:"is_topic_message,omitempty"`
	// *Optional*. *True*, if the message is a channel post that was automatically forwarded to the connected discussion group
	IsAutomaticForward *bool `json:"is_automatic_forward,omitempty"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`
	Quote *TextQuote `json:"quote,omitempty"`
	ReplyToStory *Story `json:"reply_to_story,omitempty"`
	ViaBot *User `json:"via_bot,omitempty"`
	// *Optional*. Date the message was last edited in Unix time
	EditDate *int32 `json:"edit_date,omitempty"`
	// *Optional*. *True*, if the message can't be forwarded
	HasProtectedContent *bool `json:"has_protected_content,omitempty"`
	// *Optional*. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	IsFromOffline *bool `json:"is_from_offline,omitempty"`
	// *Optional*. The unique identifier of a media message group this message belongs to
	MediaGroupId *string `json:"media_group_id,omitempty"`
	// *Optional*. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	AuthorSignature *string `json:"author_signature,omitempty"`
	// *Optional*. The number of Telegram Stars that were paid by the sender of the message to send it
	PaidStarCount *int32 `json:"paid_star_count,omitempty"`
	// *Optional*. For text messages, the actual UTF-8 text of the message
	Text *string `json:"text,omitempty"`
	// *Optional*. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// *Optional*. Unique identifier of the message effect added to the message
	EffectId *string `json:"effect_id,omitempty"`
	Animation *Animation `json:"animation,omitempty"`
	Audio *Audio `json:"audio,omitempty"`
	Document *Document `json:"document,omitempty"`
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	// *Optional*. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
	Story *Story `json:"story,omitempty"`
	Video *Video `json:"video,omitempty"`
	VideoNote *VideoNote `json:"video_note,omitempty"`
	Voice *Voice `json:"voice,omitempty"`
	// *Optional*. Caption for the animation, audio, document, paid media, photo, video or voice
	Caption *string `json:"caption,omitempty"`
	// *Optional*. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// *Optional*. *True*, if the message media is covered by a spoiler animation
	HasMediaSpoiler *bool `json:"has_media_spoiler,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Dice *Dice `json:"dice,omitempty"`
	Game *Game `json:"game,omitempty"`
	Poll *Poll `json:"poll,omitempty"`
	Venue *Venue `json:"venue,omitempty"`
	Location *Location `json:"location,omitempty"`
	// *Optional*. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// *Optional*. A chat title was changed to this value
	NewChatTitle *string `json:"new_chat_title,omitempty"`
	// *Optional*. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
	// *Optional*. Service message: the chat photo was deleted
	DeleteChatPhoto *bool `json:"delete_chat_photo,omitempty"`
	// *Optional*. Service message: the group has been created
	GroupChatCreated *bool `json:"group_chat_created,omitempty"`
	// *Optional*. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply\\_to\\_message if someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated *bool `json:"supergroup_chat_created,omitempty"`
	// *Optional*. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply\\_to\\_message if someone replies to a very first message in a channel.
	ChannelChatCreated *bool `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	// *Optional*. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int32 `json:"migrate_to_chat_id,omitempty"`
	// *Optional*. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId *int32 `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage *MaybeInaccessibleMessage `json:"pinned_message,omitempty"`
	Invoice *Invoice `json:"invoice,omitempty"`
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`
	UsersShared *UsersShared `json:"users_shared,omitempty"`
	ChatShared *ChatShared `json:"chat_shared,omitempty"`
	Gift *GiftInfo `json:"gift,omitempty"`
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`
	// *Optional*. The domain name of the website on which the user has logged in. [More about Telegram Login »](https://core.telegram.org/widgets/login)
	ConnectedWebsite *string `json:"connected_website,omitempty"`
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`
	PassportData *PassportData `json:"passport_data,omitempty"`
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed interface{} `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened interface{} `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden interface{} `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden interface{} `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged *PaidMessagePriceChanged `json:"paid_message_price_changed,omitempty"`
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted interface{} `json:"video_chat_started,omitempty"`
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(messageId int32, date int32, chat Chat) *Message {
	this := Message{}
	this.MessageId = messageId
	this.Date = date
	this.Chat = chat
	var isTopicMessage bool = true
	this.IsTopicMessage = &isTopicMessage
	var isAutomaticForward bool = true
	this.IsAutomaticForward = &isAutomaticForward
	var hasProtectedContent bool = true
	this.HasProtectedContent = &hasProtectedContent
	var isFromOffline bool = true
	this.IsFromOffline = &isFromOffline
	var showCaptionAboveMedia bool = true
	this.ShowCaptionAboveMedia = &showCaptionAboveMedia
	var hasMediaSpoiler bool = true
	this.HasMediaSpoiler = &hasMediaSpoiler
	var deleteChatPhoto bool = true
	this.DeleteChatPhoto = &deleteChatPhoto
	var groupChatCreated bool = true
	this.GroupChatCreated = &groupChatCreated
	var supergroupChatCreated bool = true
	this.SupergroupChatCreated = &supergroupChatCreated
	var channelChatCreated bool = true
	this.ChannelChatCreated = &channelChatCreated
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	var isTopicMessage bool = true
	this.IsTopicMessage = &isTopicMessage
	var isAutomaticForward bool = true
	this.IsAutomaticForward = &isAutomaticForward
	var hasProtectedContent bool = true
	this.HasProtectedContent = &hasProtectedContent
	var isFromOffline bool = true
	this.IsFromOffline = &isFromOffline
	var showCaptionAboveMedia bool = true
	this.ShowCaptionAboveMedia = &showCaptionAboveMedia
	var hasMediaSpoiler bool = true
	this.HasMediaSpoiler = &hasMediaSpoiler
	var deleteChatPhoto bool = true
	this.DeleteChatPhoto = &deleteChatPhoto
	var groupChatCreated bool = true
	this.GroupChatCreated = &groupChatCreated
	var supergroupChatCreated bool = true
	this.SupergroupChatCreated = &supergroupChatCreated
	var channelChatCreated bool = true
	this.ChannelChatCreated = &channelChatCreated
	return &this
}

// GetMessageId returns the MessageId field value
func (o *Message) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *Message) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *Message) SetMessageId(v int32) {
	o.MessageId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *Message) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *Message) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *Message) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetFrom returns the From field value if set, zero value otherwise.
func (o *Message) GetFrom() User {
	if o == nil || IsNil(o.From) {
		var ret User
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetFromOk() (*User, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Message) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given User and assigns it to the From field.
func (o *Message) SetFrom(v User) {
	o.From = &v
}


// GetSenderChat returns the SenderChat field value if set, zero value otherwise.
func (o *Message) GetSenderChat() Chat {
	if o == nil || IsNil(o.SenderChat) {
		var ret Chat
		return ret
	}
	return *o.SenderChat
}

// GetSenderChatOk returns a tuple with the SenderChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSenderChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.SenderChat) {
		return nil, false
	}
	return o.SenderChat, true
}

// HasSenderChat returns a boolean if a field has been set.
func (o *Message) HasSenderChat() bool {
	if o != nil && !IsNil(o.SenderChat) {
		return true
	}

	return false
}

// SetSenderChat gets a reference to the given Chat and assigns it to the SenderChat field.
func (o *Message) SetSenderChat(v Chat) {
	o.SenderChat = &v
}


// GetSenderBoostCount returns the SenderBoostCount field value if set, zero value otherwise.
func (o *Message) GetSenderBoostCount() int32 {
	if o == nil || IsNil(o.SenderBoostCount) {
		var ret int32
		return ret
	}
	return *o.SenderBoostCount
}

// GetSenderBoostCountOk returns a tuple with the SenderBoostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSenderBoostCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SenderBoostCount) {
		return nil, false
	}
	return o.SenderBoostCount, true
}

// HasSenderBoostCount returns a boolean if a field has been set.
func (o *Message) HasSenderBoostCount() bool {
	if o != nil && !IsNil(o.SenderBoostCount) {
		return true
	}

	return false
}

// SetSenderBoostCount gets a reference to the given int32 and assigns it to the SenderBoostCount field.
func (o *Message) SetSenderBoostCount(v int32) {
	o.SenderBoostCount = &v
}


// GetSenderBusinessBot returns the SenderBusinessBot field value if set, zero value otherwise.
func (o *Message) GetSenderBusinessBot() User {
	if o == nil || IsNil(o.SenderBusinessBot) {
		var ret User
		return ret
	}
	return *o.SenderBusinessBot
}

// GetSenderBusinessBotOk returns a tuple with the SenderBusinessBot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSenderBusinessBotOk() (*User, bool) {
	if o == nil || IsNil(o.SenderBusinessBot) {
		return nil, false
	}
	return o.SenderBusinessBot, true
}

// HasSenderBusinessBot returns a boolean if a field has been set.
func (o *Message) HasSenderBusinessBot() bool {
	if o != nil && !IsNil(o.SenderBusinessBot) {
		return true
	}

	return false
}

// SetSenderBusinessBot gets a reference to the given User and assigns it to the SenderBusinessBot field.
func (o *Message) SetSenderBusinessBot(v User) {
	o.SenderBusinessBot = &v
}


// GetDate returns the Date field value
func (o *Message) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Message) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Message) SetDate(v int32) {
	o.Date = v
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *Message) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *Message) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *Message) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChat returns the Chat field value
func (o *Message) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *Message) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *Message) SetChat(v Chat) {
	o.Chat = v
}

// GetForwardOrigin returns the ForwardOrigin field value if set, zero value otherwise.
func (o *Message) GetForwardOrigin() MessageOrigin {
	if o == nil || IsNil(o.ForwardOrigin) {
		var ret MessageOrigin
		return ret
	}
	return *o.ForwardOrigin
}

// GetForwardOriginOk returns a tuple with the ForwardOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetForwardOriginOk() (*MessageOrigin, bool) {
	if o == nil || IsNil(o.ForwardOrigin) {
		return nil, false
	}
	return o.ForwardOrigin, true
}

// HasForwardOrigin returns a boolean if a field has been set.
func (o *Message) HasForwardOrigin() bool {
	if o != nil && !IsNil(o.ForwardOrigin) {
		return true
	}

	return false
}

// SetForwardOrigin gets a reference to the given MessageOrigin and assigns it to the ForwardOrigin field.
func (o *Message) SetForwardOrigin(v MessageOrigin) {
	o.ForwardOrigin = &v
}


// GetIsTopicMessage returns the IsTopicMessage field value if set, zero value otherwise.
func (o *Message) GetIsTopicMessage() bool {
	if o == nil || IsNil(o.IsTopicMessage) {
		var ret bool
		return ret
	}
	return *o.IsTopicMessage
}

// GetIsTopicMessageOk returns a tuple with the IsTopicMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetIsTopicMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTopicMessage) {
		return nil, false
	}
	return o.IsTopicMessage, true
}

// HasIsTopicMessage returns a boolean if a field has been set.
func (o *Message) HasIsTopicMessage() bool {
	if o != nil && !IsNil(o.IsTopicMessage) {
		return true
	}

	return false
}

// SetIsTopicMessage gets a reference to the given bool and assigns it to the IsTopicMessage field.
func (o *Message) SetIsTopicMessage(v bool) {
	o.IsTopicMessage = &v
}


// GetIsAutomaticForward returns the IsAutomaticForward field value if set, zero value otherwise.
func (o *Message) GetIsAutomaticForward() bool {
	if o == nil || IsNil(o.IsAutomaticForward) {
		var ret bool
		return ret
	}
	return *o.IsAutomaticForward
}

// GetIsAutomaticForwardOk returns a tuple with the IsAutomaticForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetIsAutomaticForwardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomaticForward) {
		return nil, false
	}
	return o.IsAutomaticForward, true
}

// HasIsAutomaticForward returns a boolean if a field has been set.
func (o *Message) HasIsAutomaticForward() bool {
	if o != nil && !IsNil(o.IsAutomaticForward) {
		return true
	}

	return false
}

// SetIsAutomaticForward gets a reference to the given bool and assigns it to the IsAutomaticForward field.
func (o *Message) SetIsAutomaticForward(v bool) {
	o.IsAutomaticForward = &v
}


// GetReplyToMessage returns the ReplyToMessage field value if set, zero value otherwise.
func (o *Message) GetReplyToMessage() Message {
	if o == nil || IsNil(o.ReplyToMessage) {
		var ret Message
		return ret
	}
	return *o.ReplyToMessage
}

// GetReplyToMessageOk returns a tuple with the ReplyToMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetReplyToMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.ReplyToMessage) {
		return nil, false
	}
	return o.ReplyToMessage, true
}

// HasReplyToMessage returns a boolean if a field has been set.
func (o *Message) HasReplyToMessage() bool {
	if o != nil && !IsNil(o.ReplyToMessage) {
		return true
	}

	return false
}

// SetReplyToMessage gets a reference to the given Message and assigns it to the ReplyToMessage field.
func (o *Message) SetReplyToMessage(v Message) {
	o.ReplyToMessage = &v
}


// GetExternalReply returns the ExternalReply field value if set, zero value otherwise.
func (o *Message) GetExternalReply() ExternalReplyInfo {
	if o == nil || IsNil(o.ExternalReply) {
		var ret ExternalReplyInfo
		return ret
	}
	return *o.ExternalReply
}

// GetExternalReplyOk returns a tuple with the ExternalReply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetExternalReplyOk() (*ExternalReplyInfo, bool) {
	if o == nil || IsNil(o.ExternalReply) {
		return nil, false
	}
	return o.ExternalReply, true
}

// HasExternalReply returns a boolean if a field has been set.
func (o *Message) HasExternalReply() bool {
	if o != nil && !IsNil(o.ExternalReply) {
		return true
	}

	return false
}

// SetExternalReply gets a reference to the given ExternalReplyInfo and assigns it to the ExternalReply field.
func (o *Message) SetExternalReply(v ExternalReplyInfo) {
	o.ExternalReply = &v
}


// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *Message) GetQuote() TextQuote {
	if o == nil || IsNil(o.Quote) {
		var ret TextQuote
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetQuoteOk() (*TextQuote, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *Message) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given TextQuote and assigns it to the Quote field.
func (o *Message) SetQuote(v TextQuote) {
	o.Quote = &v
}


// GetReplyToStory returns the ReplyToStory field value if set, zero value otherwise.
func (o *Message) GetReplyToStory() Story {
	if o == nil || IsNil(o.ReplyToStory) {
		var ret Story
		return ret
	}
	return *o.ReplyToStory
}

// GetReplyToStoryOk returns a tuple with the ReplyToStory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetReplyToStoryOk() (*Story, bool) {
	if o == nil || IsNil(o.ReplyToStory) {
		return nil, false
	}
	return o.ReplyToStory, true
}

// HasReplyToStory returns a boolean if a field has been set.
func (o *Message) HasReplyToStory() bool {
	if o != nil && !IsNil(o.ReplyToStory) {
		return true
	}

	return false
}

// SetReplyToStory gets a reference to the given Story and assigns it to the ReplyToStory field.
func (o *Message) SetReplyToStory(v Story) {
	o.ReplyToStory = &v
}


// GetViaBot returns the ViaBot field value if set, zero value otherwise.
func (o *Message) GetViaBot() User {
	if o == nil || IsNil(o.ViaBot) {
		var ret User
		return ret
	}
	return *o.ViaBot
}

// GetViaBotOk returns a tuple with the ViaBot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetViaBotOk() (*User, bool) {
	if o == nil || IsNil(o.ViaBot) {
		return nil, false
	}
	return o.ViaBot, true
}

// HasViaBot returns a boolean if a field has been set.
func (o *Message) HasViaBot() bool {
	if o != nil && !IsNil(o.ViaBot) {
		return true
	}

	return false
}

// SetViaBot gets a reference to the given User and assigns it to the ViaBot field.
func (o *Message) SetViaBot(v User) {
	o.ViaBot = &v
}


// GetEditDate returns the EditDate field value if set, zero value otherwise.
func (o *Message) GetEditDate() int32 {
	if o == nil || IsNil(o.EditDate) {
		var ret int32
		return ret
	}
	return *o.EditDate
}

// GetEditDateOk returns a tuple with the EditDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetEditDateOk() (*int32, bool) {
	if o == nil || IsNil(o.EditDate) {
		return nil, false
	}
	return o.EditDate, true
}

// HasEditDate returns a boolean if a field has been set.
func (o *Message) HasEditDate() bool {
	if o != nil && !IsNil(o.EditDate) {
		return true
	}

	return false
}

// SetEditDate gets a reference to the given int32 and assigns it to the EditDate field.
func (o *Message) SetEditDate(v int32) {
	o.EditDate = &v
}


// GetHasProtectedContent returns the HasProtectedContent field value if set, zero value otherwise.
func (o *Message) GetHasProtectedContent() bool {
	if o == nil || IsNil(o.HasProtectedContent) {
		var ret bool
		return ret
	}
	return *o.HasProtectedContent
}

// GetHasProtectedContentOk returns a tuple with the HasProtectedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetHasProtectedContentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasProtectedContent) {
		return nil, false
	}
	return o.HasProtectedContent, true
}

// HasHasProtectedContent returns a boolean if a field has been set.
func (o *Message) HasHasProtectedContent() bool {
	if o != nil && !IsNil(o.HasProtectedContent) {
		return true
	}

	return false
}

// SetHasProtectedContent gets a reference to the given bool and assigns it to the HasProtectedContent field.
func (o *Message) SetHasProtectedContent(v bool) {
	o.HasProtectedContent = &v
}


// GetIsFromOffline returns the IsFromOffline field value if set, zero value otherwise.
func (o *Message) GetIsFromOffline() bool {
	if o == nil || IsNil(o.IsFromOffline) {
		var ret bool
		return ret
	}
	return *o.IsFromOffline
}

// GetIsFromOfflineOk returns a tuple with the IsFromOffline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetIsFromOfflineOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFromOffline) {
		return nil, false
	}
	return o.IsFromOffline, true
}

// HasIsFromOffline returns a boolean if a field has been set.
func (o *Message) HasIsFromOffline() bool {
	if o != nil && !IsNil(o.IsFromOffline) {
		return true
	}

	return false
}

// SetIsFromOffline gets a reference to the given bool and assigns it to the IsFromOffline field.
func (o *Message) SetIsFromOffline(v bool) {
	o.IsFromOffline = &v
}


// GetMediaGroupId returns the MediaGroupId field value if set, zero value otherwise.
func (o *Message) GetMediaGroupId() string {
	if o == nil || IsNil(o.MediaGroupId) {
		var ret string
		return ret
	}
	return *o.MediaGroupId
}

// GetMediaGroupIdOk returns a tuple with the MediaGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMediaGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaGroupId) {
		return nil, false
	}
	return o.MediaGroupId, true
}

// HasMediaGroupId returns a boolean if a field has been set.
func (o *Message) HasMediaGroupId() bool {
	if o != nil && !IsNil(o.MediaGroupId) {
		return true
	}

	return false
}

// SetMediaGroupId gets a reference to the given string and assigns it to the MediaGroupId field.
func (o *Message) SetMediaGroupId(v string) {
	o.MediaGroupId = &v
}


// GetAuthorSignature returns the AuthorSignature field value if set, zero value otherwise.
func (o *Message) GetAuthorSignature() string {
	if o == nil || IsNil(o.AuthorSignature) {
		var ret string
		return ret
	}
	return *o.AuthorSignature
}

// GetAuthorSignatureOk returns a tuple with the AuthorSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAuthorSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorSignature) {
		return nil, false
	}
	return o.AuthorSignature, true
}

// HasAuthorSignature returns a boolean if a field has been set.
func (o *Message) HasAuthorSignature() bool {
	if o != nil && !IsNil(o.AuthorSignature) {
		return true
	}

	return false
}

// SetAuthorSignature gets a reference to the given string and assigns it to the AuthorSignature field.
func (o *Message) SetAuthorSignature(v string) {
	o.AuthorSignature = &v
}


// GetPaidStarCount returns the PaidStarCount field value if set, zero value otherwise.
func (o *Message) GetPaidStarCount() int32 {
	if o == nil || IsNil(o.PaidStarCount) {
		var ret int32
		return ret
	}
	return *o.PaidStarCount
}

// GetPaidStarCountOk returns a tuple with the PaidStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPaidStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PaidStarCount) {
		return nil, false
	}
	return o.PaidStarCount, true
}

// HasPaidStarCount returns a boolean if a field has been set.
func (o *Message) HasPaidStarCount() bool {
	if o != nil && !IsNil(o.PaidStarCount) {
		return true
	}

	return false
}

// SetPaidStarCount gets a reference to the given int32 and assigns it to the PaidStarCount field.
func (o *Message) SetPaidStarCount(v int32) {
	o.PaidStarCount = &v
}


// GetText returns the Text field value if set, zero value otherwise.
func (o *Message) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Message) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Message) SetText(v string) {
	o.Text = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *Message) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *Message) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *Message) SetEntities(v []MessageEntity) {
	o.Entities = v
}


// GetLinkPreviewOptions returns the LinkPreviewOptions field value if set, zero value otherwise.
func (o *Message) GetLinkPreviewOptions() LinkPreviewOptions {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		var ret LinkPreviewOptions
		return ret
	}
	return *o.LinkPreviewOptions
}

// GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool) {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		return nil, false
	}
	return o.LinkPreviewOptions, true
}

// HasLinkPreviewOptions returns a boolean if a field has been set.
func (o *Message) HasLinkPreviewOptions() bool {
	if o != nil && !IsNil(o.LinkPreviewOptions) {
		return true
	}

	return false
}

// SetLinkPreviewOptions gets a reference to the given LinkPreviewOptions and assigns it to the LinkPreviewOptions field.
func (o *Message) SetLinkPreviewOptions(v LinkPreviewOptions) {
	o.LinkPreviewOptions = &v
}


// GetEffectId returns the EffectId field value if set, zero value otherwise.
func (o *Message) GetEffectId() string {
	if o == nil || IsNil(o.EffectId) {
		var ret string
		return ret
	}
	return *o.EffectId
}

// GetEffectIdOk returns a tuple with the EffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.EffectId) {
		return nil, false
	}
	return o.EffectId, true
}

// HasEffectId returns a boolean if a field has been set.
func (o *Message) HasEffectId() bool {
	if o != nil && !IsNil(o.EffectId) {
		return true
	}

	return false
}

// SetEffectId gets a reference to the given string and assigns it to the EffectId field.
func (o *Message) SetEffectId(v string) {
	o.EffectId = &v
}


// GetAnimation returns the Animation field value if set, zero value otherwise.
func (o *Message) GetAnimation() Animation {
	if o == nil || IsNil(o.Animation) {
		var ret Animation
		return ret
	}
	return *o.Animation
}

// GetAnimationOk returns a tuple with the Animation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAnimationOk() (*Animation, bool) {
	if o == nil || IsNil(o.Animation) {
		return nil, false
	}
	return o.Animation, true
}

// HasAnimation returns a boolean if a field has been set.
func (o *Message) HasAnimation() bool {
	if o != nil && !IsNil(o.Animation) {
		return true
	}

	return false
}

// SetAnimation gets a reference to the given Animation and assigns it to the Animation field.
func (o *Message) SetAnimation(v Animation) {
	o.Animation = &v
}


// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *Message) GetAudio() Audio {
	if o == nil || IsNil(o.Audio) {
		var ret Audio
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAudioOk() (*Audio, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *Message) HasAudio() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given Audio and assigns it to the Audio field.
func (o *Message) SetAudio(v Audio) {
	o.Audio = &v
}


// GetDocument returns the Document field value if set, zero value otherwise.
func (o *Message) GetDocument() Document {
	if o == nil || IsNil(o.Document) {
		var ret Document
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *Message) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Document and assigns it to the Document field.
func (o *Message) SetDocument(v Document) {
	o.Document = &v
}


// GetPaidMedia returns the PaidMedia field value if set, zero value otherwise.
func (o *Message) GetPaidMedia() PaidMediaInfo {
	if o == nil || IsNil(o.PaidMedia) {
		var ret PaidMediaInfo
		return ret
	}
	return *o.PaidMedia
}

// GetPaidMediaOk returns a tuple with the PaidMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPaidMediaOk() (*PaidMediaInfo, bool) {
	if o == nil || IsNil(o.PaidMedia) {
		return nil, false
	}
	return o.PaidMedia, true
}

// HasPaidMedia returns a boolean if a field has been set.
func (o *Message) HasPaidMedia() bool {
	if o != nil && !IsNil(o.PaidMedia) {
		return true
	}

	return false
}

// SetPaidMedia gets a reference to the given PaidMediaInfo and assigns it to the PaidMedia field.
func (o *Message) SetPaidMedia(v PaidMediaInfo) {
	o.PaidMedia = &v
}


// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *Message) GetPhoto() []PhotoSize {
	if o == nil || IsNil(o.Photo) {
		var ret []PhotoSize
		return ret
	}
	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPhotoOk() ([]PhotoSize, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *Message) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given []PhotoSize and assigns it to the Photo field.
func (o *Message) SetPhoto(v []PhotoSize) {
	o.Photo = v
}


// GetSticker returns the Sticker field value if set, zero value otherwise.
func (o *Message) GetSticker() Sticker {
	if o == nil || IsNil(o.Sticker) {
		var ret Sticker
		return ret
	}
	return *o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetStickerOk() (*Sticker, bool) {
	if o == nil || IsNil(o.Sticker) {
		return nil, false
	}
	return o.Sticker, true
}

// HasSticker returns a boolean if a field has been set.
func (o *Message) HasSticker() bool {
	if o != nil && !IsNil(o.Sticker) {
		return true
	}

	return false
}

// SetSticker gets a reference to the given Sticker and assigns it to the Sticker field.
func (o *Message) SetSticker(v Sticker) {
	o.Sticker = &v
}


// GetStory returns the Story field value if set, zero value otherwise.
func (o *Message) GetStory() Story {
	if o == nil || IsNil(o.Story) {
		var ret Story
		return ret
	}
	return *o.Story
}

// GetStoryOk returns a tuple with the Story field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetStoryOk() (*Story, bool) {
	if o == nil || IsNil(o.Story) {
		return nil, false
	}
	return o.Story, true
}

// HasStory returns a boolean if a field has been set.
func (o *Message) HasStory() bool {
	if o != nil && !IsNil(o.Story) {
		return true
	}

	return false
}

// SetStory gets a reference to the given Story and assigns it to the Story field.
func (o *Message) SetStory(v Story) {
	o.Story = &v
}


// GetVideo returns the Video field value if set, zero value otherwise.
func (o *Message) GetVideo() Video {
	if o == nil || IsNil(o.Video) {
		var ret Video
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVideoOk() (*Video, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *Message) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given Video and assigns it to the Video field.
func (o *Message) SetVideo(v Video) {
	o.Video = &v
}


// GetVideoNote returns the VideoNote field value if set, zero value otherwise.
func (o *Message) GetVideoNote() VideoNote {
	if o == nil || IsNil(o.VideoNote) {
		var ret VideoNote
		return ret
	}
	return *o.VideoNote
}

// GetVideoNoteOk returns a tuple with the VideoNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVideoNoteOk() (*VideoNote, bool) {
	if o == nil || IsNil(o.VideoNote) {
		return nil, false
	}
	return o.VideoNote, true
}

// HasVideoNote returns a boolean if a field has been set.
func (o *Message) HasVideoNote() bool {
	if o != nil && !IsNil(o.VideoNote) {
		return true
	}

	return false
}

// SetVideoNote gets a reference to the given VideoNote and assigns it to the VideoNote field.
func (o *Message) SetVideoNote(v VideoNote) {
	o.VideoNote = &v
}


// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *Message) GetVoice() Voice {
	if o == nil || IsNil(o.Voice) {
		var ret Voice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVoiceOk() (*Voice, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *Message) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given Voice and assigns it to the Voice field.
func (o *Message) SetVoice(v Voice) {
	o.Voice = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *Message) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *Message) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *Message) SetCaption(v string) {
	o.Caption = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *Message) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *Message) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *Message) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *Message) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *Message) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *Message) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetHasMediaSpoiler returns the HasMediaSpoiler field value if set, zero value otherwise.
func (o *Message) GetHasMediaSpoiler() bool {
	if o == nil || IsNil(o.HasMediaSpoiler) {
		var ret bool
		return ret
	}
	return *o.HasMediaSpoiler
}

// GetHasMediaSpoilerOk returns a tuple with the HasMediaSpoiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetHasMediaSpoilerOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMediaSpoiler) {
		return nil, false
	}
	return o.HasMediaSpoiler, true
}

// HasHasMediaSpoiler returns a boolean if a field has been set.
func (o *Message) HasHasMediaSpoiler() bool {
	if o != nil && !IsNil(o.HasMediaSpoiler) {
		return true
	}

	return false
}

// SetHasMediaSpoiler gets a reference to the given bool and assigns it to the HasMediaSpoiler field.
func (o *Message) SetHasMediaSpoiler(v bool) {
	o.HasMediaSpoiler = &v
}


// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Message) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Message) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *Message) SetContact(v Contact) {
	o.Contact = &v
}


// GetDice returns the Dice field value if set, zero value otherwise.
func (o *Message) GetDice() Dice {
	if o == nil || IsNil(o.Dice) {
		var ret Dice
		return ret
	}
	return *o.Dice
}

// GetDiceOk returns a tuple with the Dice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetDiceOk() (*Dice, bool) {
	if o == nil || IsNil(o.Dice) {
		return nil, false
	}
	return o.Dice, true
}

// HasDice returns a boolean if a field has been set.
func (o *Message) HasDice() bool {
	if o != nil && !IsNil(o.Dice) {
		return true
	}

	return false
}

// SetDice gets a reference to the given Dice and assigns it to the Dice field.
func (o *Message) SetDice(v Dice) {
	o.Dice = &v
}


// GetGame returns the Game field value if set, zero value otherwise.
func (o *Message) GetGame() Game {
	if o == nil || IsNil(o.Game) {
		var ret Game
		return ret
	}
	return *o.Game
}

// GetGameOk returns a tuple with the Game field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGameOk() (*Game, bool) {
	if o == nil || IsNil(o.Game) {
		return nil, false
	}
	return o.Game, true
}

// HasGame returns a boolean if a field has been set.
func (o *Message) HasGame() bool {
	if o != nil && !IsNil(o.Game) {
		return true
	}

	return false
}

// SetGame gets a reference to the given Game and assigns it to the Game field.
func (o *Message) SetGame(v Game) {
	o.Game = &v
}


// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *Message) GetPoll() Poll {
	if o == nil || IsNil(o.Poll) {
		var ret Poll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPollOk() (*Poll, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *Message) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given Poll and assigns it to the Poll field.
func (o *Message) SetPoll(v Poll) {
	o.Poll = &v
}


// GetVenue returns the Venue field value if set, zero value otherwise.
func (o *Message) GetVenue() Venue {
	if o == nil || IsNil(o.Venue) {
		var ret Venue
		return ret
	}
	return *o.Venue
}

// GetVenueOk returns a tuple with the Venue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVenueOk() (*Venue, bool) {
	if o == nil || IsNil(o.Venue) {
		return nil, false
	}
	return o.Venue, true
}

// HasVenue returns a boolean if a field has been set.
func (o *Message) HasVenue() bool {
	if o != nil && !IsNil(o.Venue) {
		return true
	}

	return false
}

// SetVenue gets a reference to the given Venue and assigns it to the Venue field.
func (o *Message) SetVenue(v Venue) {
	o.Venue = &v
}


// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Message) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Message) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *Message) SetLocation(v Location) {
	o.Location = &v
}


// GetNewChatMembers returns the NewChatMembers field value if set, zero value otherwise.
func (o *Message) GetNewChatMembers() []User {
	if o == nil || IsNil(o.NewChatMembers) {
		var ret []User
		return ret
	}
	return o.NewChatMembers
}

// GetNewChatMembersOk returns a tuple with the NewChatMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetNewChatMembersOk() ([]User, bool) {
	if o == nil || IsNil(o.NewChatMembers) {
		return nil, false
	}
	return o.NewChatMembers, true
}

// HasNewChatMembers returns a boolean if a field has been set.
func (o *Message) HasNewChatMembers() bool {
	if o != nil && !IsNil(o.NewChatMembers) {
		return true
	}

	return false
}

// SetNewChatMembers gets a reference to the given []User and assigns it to the NewChatMembers field.
func (o *Message) SetNewChatMembers(v []User) {
	o.NewChatMembers = v
}


// GetLeftChatMember returns the LeftChatMember field value if set, zero value otherwise.
func (o *Message) GetLeftChatMember() User {
	if o == nil || IsNil(o.LeftChatMember) {
		var ret User
		return ret
	}
	return *o.LeftChatMember
}

// GetLeftChatMemberOk returns a tuple with the LeftChatMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetLeftChatMemberOk() (*User, bool) {
	if o == nil || IsNil(o.LeftChatMember) {
		return nil, false
	}
	return o.LeftChatMember, true
}

// HasLeftChatMember returns a boolean if a field has been set.
func (o *Message) HasLeftChatMember() bool {
	if o != nil && !IsNil(o.LeftChatMember) {
		return true
	}

	return false
}

// SetLeftChatMember gets a reference to the given User and assigns it to the LeftChatMember field.
func (o *Message) SetLeftChatMember(v User) {
	o.LeftChatMember = &v
}


// GetNewChatTitle returns the NewChatTitle field value if set, zero value otherwise.
func (o *Message) GetNewChatTitle() string {
	if o == nil || IsNil(o.NewChatTitle) {
		var ret string
		return ret
	}
	return *o.NewChatTitle
}

// GetNewChatTitleOk returns a tuple with the NewChatTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetNewChatTitleOk() (*string, bool) {
	if o == nil || IsNil(o.NewChatTitle) {
		return nil, false
	}
	return o.NewChatTitle, true
}

// HasNewChatTitle returns a boolean if a field has been set.
func (o *Message) HasNewChatTitle() bool {
	if o != nil && !IsNil(o.NewChatTitle) {
		return true
	}

	return false
}

// SetNewChatTitle gets a reference to the given string and assigns it to the NewChatTitle field.
func (o *Message) SetNewChatTitle(v string) {
	o.NewChatTitle = &v
}


// GetNewChatPhoto returns the NewChatPhoto field value if set, zero value otherwise.
func (o *Message) GetNewChatPhoto() []PhotoSize {
	if o == nil || IsNil(o.NewChatPhoto) {
		var ret []PhotoSize
		return ret
	}
	return o.NewChatPhoto
}

// GetNewChatPhotoOk returns a tuple with the NewChatPhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetNewChatPhotoOk() ([]PhotoSize, bool) {
	if o == nil || IsNil(o.NewChatPhoto) {
		return nil, false
	}
	return o.NewChatPhoto, true
}

// HasNewChatPhoto returns a boolean if a field has been set.
func (o *Message) HasNewChatPhoto() bool {
	if o != nil && !IsNil(o.NewChatPhoto) {
		return true
	}

	return false
}

// SetNewChatPhoto gets a reference to the given []PhotoSize and assigns it to the NewChatPhoto field.
func (o *Message) SetNewChatPhoto(v []PhotoSize) {
	o.NewChatPhoto = v
}


// GetDeleteChatPhoto returns the DeleteChatPhoto field value if set, zero value otherwise.
func (o *Message) GetDeleteChatPhoto() bool {
	if o == nil || IsNil(o.DeleteChatPhoto) {
		var ret bool
		return ret
	}
	return *o.DeleteChatPhoto
}

// GetDeleteChatPhotoOk returns a tuple with the DeleteChatPhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetDeleteChatPhotoOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteChatPhoto) {
		return nil, false
	}
	return o.DeleteChatPhoto, true
}

// HasDeleteChatPhoto returns a boolean if a field has been set.
func (o *Message) HasDeleteChatPhoto() bool {
	if o != nil && !IsNil(o.DeleteChatPhoto) {
		return true
	}

	return false
}

// SetDeleteChatPhoto gets a reference to the given bool and assigns it to the DeleteChatPhoto field.
func (o *Message) SetDeleteChatPhoto(v bool) {
	o.DeleteChatPhoto = &v
}


// GetGroupChatCreated returns the GroupChatCreated field value if set, zero value otherwise.
func (o *Message) GetGroupChatCreated() bool {
	if o == nil || IsNil(o.GroupChatCreated) {
		var ret bool
		return ret
	}
	return *o.GroupChatCreated
}

// GetGroupChatCreatedOk returns a tuple with the GroupChatCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGroupChatCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupChatCreated) {
		return nil, false
	}
	return o.GroupChatCreated, true
}

// HasGroupChatCreated returns a boolean if a field has been set.
func (o *Message) HasGroupChatCreated() bool {
	if o != nil && !IsNil(o.GroupChatCreated) {
		return true
	}

	return false
}

// SetGroupChatCreated gets a reference to the given bool and assigns it to the GroupChatCreated field.
func (o *Message) SetGroupChatCreated(v bool) {
	o.GroupChatCreated = &v
}


// GetSupergroupChatCreated returns the SupergroupChatCreated field value if set, zero value otherwise.
func (o *Message) GetSupergroupChatCreated() bool {
	if o == nil || IsNil(o.SupergroupChatCreated) {
		var ret bool
		return ret
	}
	return *o.SupergroupChatCreated
}

// GetSupergroupChatCreatedOk returns a tuple with the SupergroupChatCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSupergroupChatCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.SupergroupChatCreated) {
		return nil, false
	}
	return o.SupergroupChatCreated, true
}

// HasSupergroupChatCreated returns a boolean if a field has been set.
func (o *Message) HasSupergroupChatCreated() bool {
	if o != nil && !IsNil(o.SupergroupChatCreated) {
		return true
	}

	return false
}

// SetSupergroupChatCreated gets a reference to the given bool and assigns it to the SupergroupChatCreated field.
func (o *Message) SetSupergroupChatCreated(v bool) {
	o.SupergroupChatCreated = &v
}


// GetChannelChatCreated returns the ChannelChatCreated field value if set, zero value otherwise.
func (o *Message) GetChannelChatCreated() bool {
	if o == nil || IsNil(o.ChannelChatCreated) {
		var ret bool
		return ret
	}
	return *o.ChannelChatCreated
}

// GetChannelChatCreatedOk returns a tuple with the ChannelChatCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetChannelChatCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.ChannelChatCreated) {
		return nil, false
	}
	return o.ChannelChatCreated, true
}

// HasChannelChatCreated returns a boolean if a field has been set.
func (o *Message) HasChannelChatCreated() bool {
	if o != nil && !IsNil(o.ChannelChatCreated) {
		return true
	}

	return false
}

// SetChannelChatCreated gets a reference to the given bool and assigns it to the ChannelChatCreated field.
func (o *Message) SetChannelChatCreated(v bool) {
	o.ChannelChatCreated = &v
}


// GetMessageAutoDeleteTimerChanged returns the MessageAutoDeleteTimerChanged field value if set, zero value otherwise.
func (o *Message) GetMessageAutoDeleteTimerChanged() MessageAutoDeleteTimerChanged {
	if o == nil || IsNil(o.MessageAutoDeleteTimerChanged) {
		var ret MessageAutoDeleteTimerChanged
		return ret
	}
	return *o.MessageAutoDeleteTimerChanged
}

// GetMessageAutoDeleteTimerChangedOk returns a tuple with the MessageAutoDeleteTimerChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMessageAutoDeleteTimerChangedOk() (*MessageAutoDeleteTimerChanged, bool) {
	if o == nil || IsNil(o.MessageAutoDeleteTimerChanged) {
		return nil, false
	}
	return o.MessageAutoDeleteTimerChanged, true
}

// HasMessageAutoDeleteTimerChanged returns a boolean if a field has been set.
func (o *Message) HasMessageAutoDeleteTimerChanged() bool {
	if o != nil && !IsNil(o.MessageAutoDeleteTimerChanged) {
		return true
	}

	return false
}

// SetMessageAutoDeleteTimerChanged gets a reference to the given MessageAutoDeleteTimerChanged and assigns it to the MessageAutoDeleteTimerChanged field.
func (o *Message) SetMessageAutoDeleteTimerChanged(v MessageAutoDeleteTimerChanged) {
	o.MessageAutoDeleteTimerChanged = &v
}


// GetMigrateToChatId returns the MigrateToChatId field value if set, zero value otherwise.
func (o *Message) GetMigrateToChatId() int32 {
	if o == nil || IsNil(o.MigrateToChatId) {
		var ret int32
		return ret
	}
	return *o.MigrateToChatId
}

// GetMigrateToChatIdOk returns a tuple with the MigrateToChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMigrateToChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MigrateToChatId) {
		return nil, false
	}
	return o.MigrateToChatId, true
}

// HasMigrateToChatId returns a boolean if a field has been set.
func (o *Message) HasMigrateToChatId() bool {
	if o != nil && !IsNil(o.MigrateToChatId) {
		return true
	}

	return false
}

// SetMigrateToChatId gets a reference to the given int32 and assigns it to the MigrateToChatId field.
func (o *Message) SetMigrateToChatId(v int32) {
	o.MigrateToChatId = &v
}


// GetMigrateFromChatId returns the MigrateFromChatId field value if set, zero value otherwise.
func (o *Message) GetMigrateFromChatId() int32 {
	if o == nil || IsNil(o.MigrateFromChatId) {
		var ret int32
		return ret
	}
	return *o.MigrateFromChatId
}

// GetMigrateFromChatIdOk returns a tuple with the MigrateFromChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMigrateFromChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MigrateFromChatId) {
		return nil, false
	}
	return o.MigrateFromChatId, true
}

// HasMigrateFromChatId returns a boolean if a field has been set.
func (o *Message) HasMigrateFromChatId() bool {
	if o != nil && !IsNil(o.MigrateFromChatId) {
		return true
	}

	return false
}

// SetMigrateFromChatId gets a reference to the given int32 and assigns it to the MigrateFromChatId field.
func (o *Message) SetMigrateFromChatId(v int32) {
	o.MigrateFromChatId = &v
}


// GetPinnedMessage returns the PinnedMessage field value if set, zero value otherwise.
func (o *Message) GetPinnedMessage() MaybeInaccessibleMessage {
	if o == nil || IsNil(o.PinnedMessage) {
		var ret MaybeInaccessibleMessage
		return ret
	}
	return *o.PinnedMessage
}

// GetPinnedMessageOk returns a tuple with the PinnedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPinnedMessageOk() (*MaybeInaccessibleMessage, bool) {
	if o == nil || IsNil(o.PinnedMessage) {
		return nil, false
	}
	return o.PinnedMessage, true
}

// HasPinnedMessage returns a boolean if a field has been set.
func (o *Message) HasPinnedMessage() bool {
	if o != nil && !IsNil(o.PinnedMessage) {
		return true
	}

	return false
}

// SetPinnedMessage gets a reference to the given MaybeInaccessibleMessage and assigns it to the PinnedMessage field.
func (o *Message) SetPinnedMessage(v MaybeInaccessibleMessage) {
	o.PinnedMessage = &v
}


// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *Message) GetInvoice() Invoice {
	if o == nil || IsNil(o.Invoice) {
		var ret Invoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetInvoiceOk() (*Invoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *Message) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given Invoice and assigns it to the Invoice field.
func (o *Message) SetInvoice(v Invoice) {
	o.Invoice = &v
}


// GetSuccessfulPayment returns the SuccessfulPayment field value if set, zero value otherwise.
func (o *Message) GetSuccessfulPayment() SuccessfulPayment {
	if o == nil || IsNil(o.SuccessfulPayment) {
		var ret SuccessfulPayment
		return ret
	}
	return *o.SuccessfulPayment
}

// GetSuccessfulPaymentOk returns a tuple with the SuccessfulPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSuccessfulPaymentOk() (*SuccessfulPayment, bool) {
	if o == nil || IsNil(o.SuccessfulPayment) {
		return nil, false
	}
	return o.SuccessfulPayment, true
}

// HasSuccessfulPayment returns a boolean if a field has been set.
func (o *Message) HasSuccessfulPayment() bool {
	if o != nil && !IsNil(o.SuccessfulPayment) {
		return true
	}

	return false
}

// SetSuccessfulPayment gets a reference to the given SuccessfulPayment and assigns it to the SuccessfulPayment field.
func (o *Message) SetSuccessfulPayment(v SuccessfulPayment) {
	o.SuccessfulPayment = &v
}


// GetRefundedPayment returns the RefundedPayment field value if set, zero value otherwise.
func (o *Message) GetRefundedPayment() RefundedPayment {
	if o == nil || IsNil(o.RefundedPayment) {
		var ret RefundedPayment
		return ret
	}
	return *o.RefundedPayment
}

// GetRefundedPaymentOk returns a tuple with the RefundedPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetRefundedPaymentOk() (*RefundedPayment, bool) {
	if o == nil || IsNil(o.RefundedPayment) {
		return nil, false
	}
	return o.RefundedPayment, true
}

// HasRefundedPayment returns a boolean if a field has been set.
func (o *Message) HasRefundedPayment() bool {
	if o != nil && !IsNil(o.RefundedPayment) {
		return true
	}

	return false
}

// SetRefundedPayment gets a reference to the given RefundedPayment and assigns it to the RefundedPayment field.
func (o *Message) SetRefundedPayment(v RefundedPayment) {
	o.RefundedPayment = &v
}


// GetUsersShared returns the UsersShared field value if set, zero value otherwise.
func (o *Message) GetUsersShared() UsersShared {
	if o == nil || IsNil(o.UsersShared) {
		var ret UsersShared
		return ret
	}
	return *o.UsersShared
}

// GetUsersSharedOk returns a tuple with the UsersShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetUsersSharedOk() (*UsersShared, bool) {
	if o == nil || IsNil(o.UsersShared) {
		return nil, false
	}
	return o.UsersShared, true
}

// HasUsersShared returns a boolean if a field has been set.
func (o *Message) HasUsersShared() bool {
	if o != nil && !IsNil(o.UsersShared) {
		return true
	}

	return false
}

// SetUsersShared gets a reference to the given UsersShared and assigns it to the UsersShared field.
func (o *Message) SetUsersShared(v UsersShared) {
	o.UsersShared = &v
}


// GetChatShared returns the ChatShared field value if set, zero value otherwise.
func (o *Message) GetChatShared() ChatShared {
	if o == nil || IsNil(o.ChatShared) {
		var ret ChatShared
		return ret
	}
	return *o.ChatShared
}

// GetChatSharedOk returns a tuple with the ChatShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetChatSharedOk() (*ChatShared, bool) {
	if o == nil || IsNil(o.ChatShared) {
		return nil, false
	}
	return o.ChatShared, true
}

// HasChatShared returns a boolean if a field has been set.
func (o *Message) HasChatShared() bool {
	if o != nil && !IsNil(o.ChatShared) {
		return true
	}

	return false
}

// SetChatShared gets a reference to the given ChatShared and assigns it to the ChatShared field.
func (o *Message) SetChatShared(v ChatShared) {
	o.ChatShared = &v
}


// GetGift returns the Gift field value if set, zero value otherwise.
func (o *Message) GetGift() GiftInfo {
	if o == nil || IsNil(o.Gift) {
		var ret GiftInfo
		return ret
	}
	return *o.Gift
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGiftOk() (*GiftInfo, bool) {
	if o == nil || IsNil(o.Gift) {
		return nil, false
	}
	return o.Gift, true
}

// HasGift returns a boolean if a field has been set.
func (o *Message) HasGift() bool {
	if o != nil && !IsNil(o.Gift) {
		return true
	}

	return false
}

// SetGift gets a reference to the given GiftInfo and assigns it to the Gift field.
func (o *Message) SetGift(v GiftInfo) {
	o.Gift = &v
}


// GetUniqueGift returns the UniqueGift field value if set, zero value otherwise.
func (o *Message) GetUniqueGift() UniqueGiftInfo {
	if o == nil || IsNil(o.UniqueGift) {
		var ret UniqueGiftInfo
		return ret
	}
	return *o.UniqueGift
}

// GetUniqueGiftOk returns a tuple with the UniqueGift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetUniqueGiftOk() (*UniqueGiftInfo, bool) {
	if o == nil || IsNil(o.UniqueGift) {
		return nil, false
	}
	return o.UniqueGift, true
}

// HasUniqueGift returns a boolean if a field has been set.
func (o *Message) HasUniqueGift() bool {
	if o != nil && !IsNil(o.UniqueGift) {
		return true
	}

	return false
}

// SetUniqueGift gets a reference to the given UniqueGiftInfo and assigns it to the UniqueGift field.
func (o *Message) SetUniqueGift(v UniqueGiftInfo) {
	o.UniqueGift = &v
}


// GetConnectedWebsite returns the ConnectedWebsite field value if set, zero value otherwise.
func (o *Message) GetConnectedWebsite() string {
	if o == nil || IsNil(o.ConnectedWebsite) {
		var ret string
		return ret
	}
	return *o.ConnectedWebsite
}

// GetConnectedWebsiteOk returns a tuple with the ConnectedWebsite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetConnectedWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectedWebsite) {
		return nil, false
	}
	return o.ConnectedWebsite, true
}

// HasConnectedWebsite returns a boolean if a field has been set.
func (o *Message) HasConnectedWebsite() bool {
	if o != nil && !IsNil(o.ConnectedWebsite) {
		return true
	}

	return false
}

// SetConnectedWebsite gets a reference to the given string and assigns it to the ConnectedWebsite field.
func (o *Message) SetConnectedWebsite(v string) {
	o.ConnectedWebsite = &v
}


// GetWriteAccessAllowed returns the WriteAccessAllowed field value if set, zero value otherwise.
func (o *Message) GetWriteAccessAllowed() WriteAccessAllowed {
	if o == nil || IsNil(o.WriteAccessAllowed) {
		var ret WriteAccessAllowed
		return ret
	}
	return *o.WriteAccessAllowed
}

// GetWriteAccessAllowedOk returns a tuple with the WriteAccessAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetWriteAccessAllowedOk() (*WriteAccessAllowed, bool) {
	if o == nil || IsNil(o.WriteAccessAllowed) {
		return nil, false
	}
	return o.WriteAccessAllowed, true
}

// HasWriteAccessAllowed returns a boolean if a field has been set.
func (o *Message) HasWriteAccessAllowed() bool {
	if o != nil && !IsNil(o.WriteAccessAllowed) {
		return true
	}

	return false
}

// SetWriteAccessAllowed gets a reference to the given WriteAccessAllowed and assigns it to the WriteAccessAllowed field.
func (o *Message) SetWriteAccessAllowed(v WriteAccessAllowed) {
	o.WriteAccessAllowed = &v
}


// GetPassportData returns the PassportData field value if set, zero value otherwise.
func (o *Message) GetPassportData() PassportData {
	if o == nil || IsNil(o.PassportData) {
		var ret PassportData
		return ret
	}
	return *o.PassportData
}

// GetPassportDataOk returns a tuple with the PassportData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPassportDataOk() (*PassportData, bool) {
	if o == nil || IsNil(o.PassportData) {
		return nil, false
	}
	return o.PassportData, true
}

// HasPassportData returns a boolean if a field has been set.
func (o *Message) HasPassportData() bool {
	if o != nil && !IsNil(o.PassportData) {
		return true
	}

	return false
}

// SetPassportData gets a reference to the given PassportData and assigns it to the PassportData field.
func (o *Message) SetPassportData(v PassportData) {
	o.PassportData = &v
}


// GetProximityAlertTriggered returns the ProximityAlertTriggered field value if set, zero value otherwise.
func (o *Message) GetProximityAlertTriggered() ProximityAlertTriggered {
	if o == nil || IsNil(o.ProximityAlertTriggered) {
		var ret ProximityAlertTriggered
		return ret
	}
	return *o.ProximityAlertTriggered
}

// GetProximityAlertTriggeredOk returns a tuple with the ProximityAlertTriggered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetProximityAlertTriggeredOk() (*ProximityAlertTriggered, bool) {
	if o == nil || IsNil(o.ProximityAlertTriggered) {
		return nil, false
	}
	return o.ProximityAlertTriggered, true
}

// HasProximityAlertTriggered returns a boolean if a field has been set.
func (o *Message) HasProximityAlertTriggered() bool {
	if o != nil && !IsNil(o.ProximityAlertTriggered) {
		return true
	}

	return false
}

// SetProximityAlertTriggered gets a reference to the given ProximityAlertTriggered and assigns it to the ProximityAlertTriggered field.
func (o *Message) SetProximityAlertTriggered(v ProximityAlertTriggered) {
	o.ProximityAlertTriggered = &v
}


// GetBoostAdded returns the BoostAdded field value if set, zero value otherwise.
func (o *Message) GetBoostAdded() ChatBoostAdded {
	if o == nil || IsNil(o.BoostAdded) {
		var ret ChatBoostAdded
		return ret
	}
	return *o.BoostAdded
}

// GetBoostAddedOk returns a tuple with the BoostAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetBoostAddedOk() (*ChatBoostAdded, bool) {
	if o == nil || IsNil(o.BoostAdded) {
		return nil, false
	}
	return o.BoostAdded, true
}

// HasBoostAdded returns a boolean if a field has been set.
func (o *Message) HasBoostAdded() bool {
	if o != nil && !IsNil(o.BoostAdded) {
		return true
	}

	return false
}

// SetBoostAdded gets a reference to the given ChatBoostAdded and assigns it to the BoostAdded field.
func (o *Message) SetBoostAdded(v ChatBoostAdded) {
	o.BoostAdded = &v
}


// GetChatBackgroundSet returns the ChatBackgroundSet field value if set, zero value otherwise.
func (o *Message) GetChatBackgroundSet() ChatBackground {
	if o == nil || IsNil(o.ChatBackgroundSet) {
		var ret ChatBackground
		return ret
	}
	return *o.ChatBackgroundSet
}

// GetChatBackgroundSetOk returns a tuple with the ChatBackgroundSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetChatBackgroundSetOk() (*ChatBackground, bool) {
	if o == nil || IsNil(o.ChatBackgroundSet) {
		return nil, false
	}
	return o.ChatBackgroundSet, true
}

// HasChatBackgroundSet returns a boolean if a field has been set.
func (o *Message) HasChatBackgroundSet() bool {
	if o != nil && !IsNil(o.ChatBackgroundSet) {
		return true
	}

	return false
}

// SetChatBackgroundSet gets a reference to the given ChatBackground and assigns it to the ChatBackgroundSet field.
func (o *Message) SetChatBackgroundSet(v ChatBackground) {
	o.ChatBackgroundSet = &v
}


// GetForumTopicCreated returns the ForumTopicCreated field value if set, zero value otherwise.
func (o *Message) GetForumTopicCreated() ForumTopicCreated {
	if o == nil || IsNil(o.ForumTopicCreated) {
		var ret ForumTopicCreated
		return ret
	}
	return *o.ForumTopicCreated
}

// GetForumTopicCreatedOk returns a tuple with the ForumTopicCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetForumTopicCreatedOk() (*ForumTopicCreated, bool) {
	if o == nil || IsNil(o.ForumTopicCreated) {
		return nil, false
	}
	return o.ForumTopicCreated, true
}

// HasForumTopicCreated returns a boolean if a field has been set.
func (o *Message) HasForumTopicCreated() bool {
	if o != nil && !IsNil(o.ForumTopicCreated) {
		return true
	}

	return false
}

// SetForumTopicCreated gets a reference to the given ForumTopicCreated and assigns it to the ForumTopicCreated field.
func (o *Message) SetForumTopicCreated(v ForumTopicCreated) {
	o.ForumTopicCreated = &v
}


// GetForumTopicEdited returns the ForumTopicEdited field value if set, zero value otherwise.
func (o *Message) GetForumTopicEdited() ForumTopicEdited {
	if o == nil || IsNil(o.ForumTopicEdited) {
		var ret ForumTopicEdited
		return ret
	}
	return *o.ForumTopicEdited
}

// GetForumTopicEditedOk returns a tuple with the ForumTopicEdited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetForumTopicEditedOk() (*ForumTopicEdited, bool) {
	if o == nil || IsNil(o.ForumTopicEdited) {
		return nil, false
	}
	return o.ForumTopicEdited, true
}

// HasForumTopicEdited returns a boolean if a field has been set.
func (o *Message) HasForumTopicEdited() bool {
	if o != nil && !IsNil(o.ForumTopicEdited) {
		return true
	}

	return false
}

// SetForumTopicEdited gets a reference to the given ForumTopicEdited and assigns it to the ForumTopicEdited field.
func (o *Message) SetForumTopicEdited(v ForumTopicEdited) {
	o.ForumTopicEdited = &v
}


// GetForumTopicClosed returns the ForumTopicClosed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetForumTopicClosed() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ForumTopicClosed
}

// GetForumTopicClosedOk returns a tuple with the ForumTopicClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetForumTopicClosedOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForumTopicClosed, true
}

// HasForumTopicClosed returns a boolean if a field has been set.
func (o *Message) HasForumTopicClosed() bool {
	if o != nil && !IsNil(o.ForumTopicClosed) {
		return true
	}

	return false
}

// SetForumTopicClosed gets a reference to the given interface{} and assigns it to the ForumTopicClosed field.
func (o *Message) SetForumTopicClosed(v interface{}) {
	o.ForumTopicClosed = v
}


// GetForumTopicReopened returns the ForumTopicReopened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetForumTopicReopened() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ForumTopicReopened
}

// GetForumTopicReopenedOk returns a tuple with the ForumTopicReopened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetForumTopicReopenedOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForumTopicReopened, true
}

// HasForumTopicReopened returns a boolean if a field has been set.
func (o *Message) HasForumTopicReopened() bool {
	if o != nil && !IsNil(o.ForumTopicReopened) {
		return true
	}

	return false
}

// SetForumTopicReopened gets a reference to the given interface{} and assigns it to the ForumTopicReopened field.
func (o *Message) SetForumTopicReopened(v interface{}) {
	o.ForumTopicReopened = v
}


// GetGeneralForumTopicHidden returns the GeneralForumTopicHidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetGeneralForumTopicHidden() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GeneralForumTopicHidden
}

// GetGeneralForumTopicHiddenOk returns a tuple with the GeneralForumTopicHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetGeneralForumTopicHiddenOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneralForumTopicHidden, true
}

// HasGeneralForumTopicHidden returns a boolean if a field has been set.
func (o *Message) HasGeneralForumTopicHidden() bool {
	if o != nil && !IsNil(o.GeneralForumTopicHidden) {
		return true
	}

	return false
}

// SetGeneralForumTopicHidden gets a reference to the given interface{} and assigns it to the GeneralForumTopicHidden field.
func (o *Message) SetGeneralForumTopicHidden(v interface{}) {
	o.GeneralForumTopicHidden = v
}


// GetGeneralForumTopicUnhidden returns the GeneralForumTopicUnhidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetGeneralForumTopicUnhidden() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GeneralForumTopicUnhidden
}

// GetGeneralForumTopicUnhiddenOk returns a tuple with the GeneralForumTopicUnhidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetGeneralForumTopicUnhiddenOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneralForumTopicUnhidden, true
}

// HasGeneralForumTopicUnhidden returns a boolean if a field has been set.
func (o *Message) HasGeneralForumTopicUnhidden() bool {
	if o != nil && !IsNil(o.GeneralForumTopicUnhidden) {
		return true
	}

	return false
}

// SetGeneralForumTopicUnhidden gets a reference to the given interface{} and assigns it to the GeneralForumTopicUnhidden field.
func (o *Message) SetGeneralForumTopicUnhidden(v interface{}) {
	o.GeneralForumTopicUnhidden = v
}


// GetGiveawayCreated returns the GiveawayCreated field value if set, zero value otherwise.
func (o *Message) GetGiveawayCreated() GiveawayCreated {
	if o == nil || IsNil(o.GiveawayCreated) {
		var ret GiveawayCreated
		return ret
	}
	return *o.GiveawayCreated
}

// GetGiveawayCreatedOk returns a tuple with the GiveawayCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGiveawayCreatedOk() (*GiveawayCreated, bool) {
	if o == nil || IsNil(o.GiveawayCreated) {
		return nil, false
	}
	return o.GiveawayCreated, true
}

// HasGiveawayCreated returns a boolean if a field has been set.
func (o *Message) HasGiveawayCreated() bool {
	if o != nil && !IsNil(o.GiveawayCreated) {
		return true
	}

	return false
}

// SetGiveawayCreated gets a reference to the given GiveawayCreated and assigns it to the GiveawayCreated field.
func (o *Message) SetGiveawayCreated(v GiveawayCreated) {
	o.GiveawayCreated = &v
}


// GetGiveaway returns the Giveaway field value if set, zero value otherwise.
func (o *Message) GetGiveaway() Giveaway {
	if o == nil || IsNil(o.Giveaway) {
		var ret Giveaway
		return ret
	}
	return *o.Giveaway
}

// GetGiveawayOk returns a tuple with the Giveaway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGiveawayOk() (*Giveaway, bool) {
	if o == nil || IsNil(o.Giveaway) {
		return nil, false
	}
	return o.Giveaway, true
}

// HasGiveaway returns a boolean if a field has been set.
func (o *Message) HasGiveaway() bool {
	if o != nil && !IsNil(o.Giveaway) {
		return true
	}

	return false
}

// SetGiveaway gets a reference to the given Giveaway and assigns it to the Giveaway field.
func (o *Message) SetGiveaway(v Giveaway) {
	o.Giveaway = &v
}


// GetGiveawayWinners returns the GiveawayWinners field value if set, zero value otherwise.
func (o *Message) GetGiveawayWinners() GiveawayWinners {
	if o == nil || IsNil(o.GiveawayWinners) {
		var ret GiveawayWinners
		return ret
	}
	return *o.GiveawayWinners
}

// GetGiveawayWinnersOk returns a tuple with the GiveawayWinners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGiveawayWinnersOk() (*GiveawayWinners, bool) {
	if o == nil || IsNil(o.GiveawayWinners) {
		return nil, false
	}
	return o.GiveawayWinners, true
}

// HasGiveawayWinners returns a boolean if a field has been set.
func (o *Message) HasGiveawayWinners() bool {
	if o != nil && !IsNil(o.GiveawayWinners) {
		return true
	}

	return false
}

// SetGiveawayWinners gets a reference to the given GiveawayWinners and assigns it to the GiveawayWinners field.
func (o *Message) SetGiveawayWinners(v GiveawayWinners) {
	o.GiveawayWinners = &v
}


// GetGiveawayCompleted returns the GiveawayCompleted field value if set, zero value otherwise.
func (o *Message) GetGiveawayCompleted() GiveawayCompleted {
	if o == nil || IsNil(o.GiveawayCompleted) {
		var ret GiveawayCompleted
		return ret
	}
	return *o.GiveawayCompleted
}

// GetGiveawayCompletedOk returns a tuple with the GiveawayCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetGiveawayCompletedOk() (*GiveawayCompleted, bool) {
	if o == nil || IsNil(o.GiveawayCompleted) {
		return nil, false
	}
	return o.GiveawayCompleted, true
}

// HasGiveawayCompleted returns a boolean if a field has been set.
func (o *Message) HasGiveawayCompleted() bool {
	if o != nil && !IsNil(o.GiveawayCompleted) {
		return true
	}

	return false
}

// SetGiveawayCompleted gets a reference to the given GiveawayCompleted and assigns it to the GiveawayCompleted field.
func (o *Message) SetGiveawayCompleted(v GiveawayCompleted) {
	o.GiveawayCompleted = &v
}


// GetPaidMessagePriceChanged returns the PaidMessagePriceChanged field value if set, zero value otherwise.
func (o *Message) GetPaidMessagePriceChanged() PaidMessagePriceChanged {
	if o == nil || IsNil(o.PaidMessagePriceChanged) {
		var ret PaidMessagePriceChanged
		return ret
	}
	return *o.PaidMessagePriceChanged
}

// GetPaidMessagePriceChangedOk returns a tuple with the PaidMessagePriceChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPaidMessagePriceChangedOk() (*PaidMessagePriceChanged, bool) {
	if o == nil || IsNil(o.PaidMessagePriceChanged) {
		return nil, false
	}
	return o.PaidMessagePriceChanged, true
}

// HasPaidMessagePriceChanged returns a boolean if a field has been set.
func (o *Message) HasPaidMessagePriceChanged() bool {
	if o != nil && !IsNil(o.PaidMessagePriceChanged) {
		return true
	}

	return false
}

// SetPaidMessagePriceChanged gets a reference to the given PaidMessagePriceChanged and assigns it to the PaidMessagePriceChanged field.
func (o *Message) SetPaidMessagePriceChanged(v PaidMessagePriceChanged) {
	o.PaidMessagePriceChanged = &v
}


// GetVideoChatScheduled returns the VideoChatScheduled field value if set, zero value otherwise.
func (o *Message) GetVideoChatScheduled() VideoChatScheduled {
	if o == nil || IsNil(o.VideoChatScheduled) {
		var ret VideoChatScheduled
		return ret
	}
	return *o.VideoChatScheduled
}

// GetVideoChatScheduledOk returns a tuple with the VideoChatScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVideoChatScheduledOk() (*VideoChatScheduled, bool) {
	if o == nil || IsNil(o.VideoChatScheduled) {
		return nil, false
	}
	return o.VideoChatScheduled, true
}

// HasVideoChatScheduled returns a boolean if a field has been set.
func (o *Message) HasVideoChatScheduled() bool {
	if o != nil && !IsNil(o.VideoChatScheduled) {
		return true
	}

	return false
}

// SetVideoChatScheduled gets a reference to the given VideoChatScheduled and assigns it to the VideoChatScheduled field.
func (o *Message) SetVideoChatScheduled(v VideoChatScheduled) {
	o.VideoChatScheduled = &v
}


// GetVideoChatStarted returns the VideoChatStarted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetVideoChatStarted() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VideoChatStarted
}

// GetVideoChatStartedOk returns a tuple with the VideoChatStarted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetVideoChatStartedOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VideoChatStarted, true
}

// HasVideoChatStarted returns a boolean if a field has been set.
func (o *Message) HasVideoChatStarted() bool {
	if o != nil && !IsNil(o.VideoChatStarted) {
		return true
	}

	return false
}

// SetVideoChatStarted gets a reference to the given interface{} and assigns it to the VideoChatStarted field.
func (o *Message) SetVideoChatStarted(v interface{}) {
	o.VideoChatStarted = v
}


// GetVideoChatEnded returns the VideoChatEnded field value if set, zero value otherwise.
func (o *Message) GetVideoChatEnded() VideoChatEnded {
	if o == nil || IsNil(o.VideoChatEnded) {
		var ret VideoChatEnded
		return ret
	}
	return *o.VideoChatEnded
}

// GetVideoChatEndedOk returns a tuple with the VideoChatEnded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVideoChatEndedOk() (*VideoChatEnded, bool) {
	if o == nil || IsNil(o.VideoChatEnded) {
		return nil, false
	}
	return o.VideoChatEnded, true
}

// HasVideoChatEnded returns a boolean if a field has been set.
func (o *Message) HasVideoChatEnded() bool {
	if o != nil && !IsNil(o.VideoChatEnded) {
		return true
	}

	return false
}

// SetVideoChatEnded gets a reference to the given VideoChatEnded and assigns it to the VideoChatEnded field.
func (o *Message) SetVideoChatEnded(v VideoChatEnded) {
	o.VideoChatEnded = &v
}


// GetVideoChatParticipantsInvited returns the VideoChatParticipantsInvited field value if set, zero value otherwise.
func (o *Message) GetVideoChatParticipantsInvited() VideoChatParticipantsInvited {
	if o == nil || IsNil(o.VideoChatParticipantsInvited) {
		var ret VideoChatParticipantsInvited
		return ret
	}
	return *o.VideoChatParticipantsInvited
}

// GetVideoChatParticipantsInvitedOk returns a tuple with the VideoChatParticipantsInvited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetVideoChatParticipantsInvitedOk() (*VideoChatParticipantsInvited, bool) {
	if o == nil || IsNil(o.VideoChatParticipantsInvited) {
		return nil, false
	}
	return o.VideoChatParticipantsInvited, true
}

// HasVideoChatParticipantsInvited returns a boolean if a field has been set.
func (o *Message) HasVideoChatParticipantsInvited() bool {
	if o != nil && !IsNil(o.VideoChatParticipantsInvited) {
		return true
	}

	return false
}

// SetVideoChatParticipantsInvited gets a reference to the given VideoChatParticipantsInvited and assigns it to the VideoChatParticipantsInvited field.
func (o *Message) SetVideoChatParticipantsInvited(v VideoChatParticipantsInvited) {
	o.VideoChatParticipantsInvited = &v
}


// GetWebAppData returns the WebAppData field value if set, zero value otherwise.
func (o *Message) GetWebAppData() WebAppData {
	if o == nil || IsNil(o.WebAppData) {
		var ret WebAppData
		return ret
	}
	return *o.WebAppData
}

// GetWebAppDataOk returns a tuple with the WebAppData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetWebAppDataOk() (*WebAppData, bool) {
	if o == nil || IsNil(o.WebAppData) {
		return nil, false
	}
	return o.WebAppData, true
}

// HasWebAppData returns a boolean if a field has been set.
func (o *Message) HasWebAppData() bool {
	if o != nil && !IsNil(o.WebAppData) {
		return true
	}

	return false
}

// SetWebAppData gets a reference to the given WebAppData and assigns it to the WebAppData field.
func (o *Message) SetWebAppData(v WebAppData) {
	o.WebAppData = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *Message) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *Message) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *Message) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.SenderChat) {
		toSerialize["sender_chat"] = o.SenderChat
	}
	if !IsNil(o.SenderBoostCount) {
		toSerialize["sender_boost_count"] = o.SenderBoostCount
	}
	if !IsNil(o.SenderBusinessBot) {
		toSerialize["sender_business_bot"] = o.SenderBusinessBot
	}
	toSerialize["date"] = o.Date
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat"] = o.Chat
	if !IsNil(o.ForwardOrigin) {
		toSerialize["forward_origin"] = o.ForwardOrigin
	}
	if !IsNil(o.IsTopicMessage) {
		toSerialize["is_topic_message"] = o.IsTopicMessage
	}
	if !IsNil(o.IsAutomaticForward) {
		toSerialize["is_automatic_forward"] = o.IsAutomaticForward
	}
	if !IsNil(o.ReplyToMessage) {
		toSerialize["reply_to_message"] = o.ReplyToMessage
	}
	if !IsNil(o.ExternalReply) {
		toSerialize["external_reply"] = o.ExternalReply
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.ReplyToStory) {
		toSerialize["reply_to_story"] = o.ReplyToStory
	}
	if !IsNil(o.ViaBot) {
		toSerialize["via_bot"] = o.ViaBot
	}
	if !IsNil(o.EditDate) {
		toSerialize["edit_date"] = o.EditDate
	}
	if !IsNil(o.HasProtectedContent) {
		toSerialize["has_protected_content"] = o.HasProtectedContent
	}
	if !IsNil(o.IsFromOffline) {
		toSerialize["is_from_offline"] = o.IsFromOffline
	}
	if !IsNil(o.MediaGroupId) {
		toSerialize["media_group_id"] = o.MediaGroupId
	}
	if !IsNil(o.AuthorSignature) {
		toSerialize["author_signature"] = o.AuthorSignature
	}
	if !IsNil(o.PaidStarCount) {
		toSerialize["paid_star_count"] = o.PaidStarCount
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.LinkPreviewOptions) {
		toSerialize["link_preview_options"] = o.LinkPreviewOptions
	}
	if !IsNil(o.EffectId) {
		toSerialize["effect_id"] = o.EffectId
	}
	if !IsNil(o.Animation) {
		toSerialize["animation"] = o.Animation
	}
	if !IsNil(o.Audio) {
		toSerialize["audio"] = o.Audio
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.PaidMedia) {
		toSerialize["paid_media"] = o.PaidMedia
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.Sticker) {
		toSerialize["sticker"] = o.Sticker
	}
	if !IsNil(o.Story) {
		toSerialize["story"] = o.Story
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.VideoNote) {
		toSerialize["video_note"] = o.VideoNote
	}
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.CaptionEntities) {
		toSerialize["caption_entities"] = o.CaptionEntities
	}
	if !IsNil(o.ShowCaptionAboveMedia) {
		toSerialize["show_caption_above_media"] = o.ShowCaptionAboveMedia
	}
	if !IsNil(o.HasMediaSpoiler) {
		toSerialize["has_media_spoiler"] = o.HasMediaSpoiler
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Dice) {
		toSerialize["dice"] = o.Dice
	}
	if !IsNil(o.Game) {
		toSerialize["game"] = o.Game
	}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}
	if !IsNil(o.Venue) {
		toSerialize["venue"] = o.Venue
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.NewChatMembers) {
		toSerialize["new_chat_members"] = o.NewChatMembers
	}
	if !IsNil(o.LeftChatMember) {
		toSerialize["left_chat_member"] = o.LeftChatMember
	}
	if !IsNil(o.NewChatTitle) {
		toSerialize["new_chat_title"] = o.NewChatTitle
	}
	if !IsNil(o.NewChatPhoto) {
		toSerialize["new_chat_photo"] = o.NewChatPhoto
	}
	if !IsNil(o.DeleteChatPhoto) {
		toSerialize["delete_chat_photo"] = o.DeleteChatPhoto
	}
	if !IsNil(o.GroupChatCreated) {
		toSerialize["group_chat_created"] = o.GroupChatCreated
	}
	if !IsNil(o.SupergroupChatCreated) {
		toSerialize["supergroup_chat_created"] = o.SupergroupChatCreated
	}
	if !IsNil(o.ChannelChatCreated) {
		toSerialize["channel_chat_created"] = o.ChannelChatCreated
	}
	if !IsNil(o.MessageAutoDeleteTimerChanged) {
		toSerialize["message_auto_delete_timer_changed"] = o.MessageAutoDeleteTimerChanged
	}
	if !IsNil(o.MigrateToChatId) {
		toSerialize["migrate_to_chat_id"] = o.MigrateToChatId
	}
	if !IsNil(o.MigrateFromChatId) {
		toSerialize["migrate_from_chat_id"] = o.MigrateFromChatId
	}
	if !IsNil(o.PinnedMessage) {
		toSerialize["pinned_message"] = o.PinnedMessage
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.SuccessfulPayment) {
		toSerialize["successful_payment"] = o.SuccessfulPayment
	}
	if !IsNil(o.RefundedPayment) {
		toSerialize["refunded_payment"] = o.RefundedPayment
	}
	if !IsNil(o.UsersShared) {
		toSerialize["users_shared"] = o.UsersShared
	}
	if !IsNil(o.ChatShared) {
		toSerialize["chat_shared"] = o.ChatShared
	}
	if !IsNil(o.Gift) {
		toSerialize["gift"] = o.Gift
	}
	if !IsNil(o.UniqueGift) {
		toSerialize["unique_gift"] = o.UniqueGift
	}
	if !IsNil(o.ConnectedWebsite) {
		toSerialize["connected_website"] = o.ConnectedWebsite
	}
	if !IsNil(o.WriteAccessAllowed) {
		toSerialize["write_access_allowed"] = o.WriteAccessAllowed
	}
	if !IsNil(o.PassportData) {
		toSerialize["passport_data"] = o.PassportData
	}
	if !IsNil(o.ProximityAlertTriggered) {
		toSerialize["proximity_alert_triggered"] = o.ProximityAlertTriggered
	}
	if !IsNil(o.BoostAdded) {
		toSerialize["boost_added"] = o.BoostAdded
	}
	if !IsNil(o.ChatBackgroundSet) {
		toSerialize["chat_background_set"] = o.ChatBackgroundSet
	}
	if !IsNil(o.ForumTopicCreated) {
		toSerialize["forum_topic_created"] = o.ForumTopicCreated
	}
	if !IsNil(o.ForumTopicEdited) {
		toSerialize["forum_topic_edited"] = o.ForumTopicEdited
	}
	if o.ForumTopicClosed != nil {
		toSerialize["forum_topic_closed"] = o.ForumTopicClosed
	}
	if o.ForumTopicReopened != nil {
		toSerialize["forum_topic_reopened"] = o.ForumTopicReopened
	}
	if o.GeneralForumTopicHidden != nil {
		toSerialize["general_forum_topic_hidden"] = o.GeneralForumTopicHidden
	}
	if o.GeneralForumTopicUnhidden != nil {
		toSerialize["general_forum_topic_unhidden"] = o.GeneralForumTopicUnhidden
	}
	if !IsNil(o.GiveawayCreated) {
		toSerialize["giveaway_created"] = o.GiveawayCreated
	}
	if !IsNil(o.Giveaway) {
		toSerialize["giveaway"] = o.Giveaway
	}
	if !IsNil(o.GiveawayWinners) {
		toSerialize["giveaway_winners"] = o.GiveawayWinners
	}
	if !IsNil(o.GiveawayCompleted) {
		toSerialize["giveaway_completed"] = o.GiveawayCompleted
	}
	if !IsNil(o.PaidMessagePriceChanged) {
		toSerialize["paid_message_price_changed"] = o.PaidMessagePriceChanged
	}
	if !IsNil(o.VideoChatScheduled) {
		toSerialize["video_chat_scheduled"] = o.VideoChatScheduled
	}
	if o.VideoChatStarted != nil {
		toSerialize["video_chat_started"] = o.VideoChatStarted
	}
	if !IsNil(o.VideoChatEnded) {
		toSerialize["video_chat_ended"] = o.VideoChatEnded
	}
	if !IsNil(o.VideoChatParticipantsInvited) {
		toSerialize["video_chat_participants_invited"] = o.VideoChatParticipantsInvited
	}
	if !IsNil(o.WebAppData) {
		toSerialize["web_app_data"] = o.WebAppData
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_id",
		"date",
		"chat",
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

	varMessage := _Message{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


