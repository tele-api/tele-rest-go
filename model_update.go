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

// checks if the Update type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Update{}

// Update This [object](https://core.telegram.org/bots/api/#available-types) represents an incoming update.   At most **one** of the optional parameters can be present in any given update.
type Update struct {
	// The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you're using [webhooks](https://core.telegram.org/bots/api/#setwebhook), since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateId int32 `json:"update_id"`
	Message *Message `json:"message,omitempty"`
	EditedMessage *Message `json:"edited_message,omitempty"`
	ChannelPost *Message `json:"channel_post,omitempty"`
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`
	BusinessMessage *Message `json:"business_message,omitempty"`
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`
	Poll *Poll `json:"poll,omitempty"`
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

type _Update Update

// NewUpdate instantiates a new Update object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdate(updateId int32) *Update {
	this := Update{}
	this.UpdateId = updateId
	return &this
}

// NewUpdateWithDefaults instantiates a new Update object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWithDefaults() *Update {
	this := Update{}
	return &this
}

// GetUpdateId returns the UpdateId field value
func (o *Update) GetUpdateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateId
}

// GetUpdateIdOk returns a tuple with the UpdateId field value
// and a boolean to check if the value has been set.
func (o *Update) GetUpdateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateId, true
}

// SetUpdateId sets field value
func (o *Update) SetUpdateId(v int32) {
	o.UpdateId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Update) GetMessage() Message {
	if o == nil || IsNil(o.Message) {
		var ret Message
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Update) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given Message and assigns it to the Message field.
func (o *Update) SetMessage(v Message) {
	o.Message = &v
}


// GetEditedMessage returns the EditedMessage field value if set, zero value otherwise.
func (o *Update) GetEditedMessage() Message {
	if o == nil || IsNil(o.EditedMessage) {
		var ret Message
		return ret
	}
	return *o.EditedMessage
}

// GetEditedMessageOk returns a tuple with the EditedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetEditedMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.EditedMessage) {
		return nil, false
	}
	return o.EditedMessage, true
}

// HasEditedMessage returns a boolean if a field has been set.
func (o *Update) HasEditedMessage() bool {
	if o != nil && !IsNil(o.EditedMessage) {
		return true
	}

	return false
}

// SetEditedMessage gets a reference to the given Message and assigns it to the EditedMessage field.
func (o *Update) SetEditedMessage(v Message) {
	o.EditedMessage = &v
}


// GetChannelPost returns the ChannelPost field value if set, zero value otherwise.
func (o *Update) GetChannelPost() Message {
	if o == nil || IsNil(o.ChannelPost) {
		var ret Message
		return ret
	}
	return *o.ChannelPost
}

// GetChannelPostOk returns a tuple with the ChannelPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetChannelPostOk() (*Message, bool) {
	if o == nil || IsNil(o.ChannelPost) {
		return nil, false
	}
	return o.ChannelPost, true
}

// HasChannelPost returns a boolean if a field has been set.
func (o *Update) HasChannelPost() bool {
	if o != nil && !IsNil(o.ChannelPost) {
		return true
	}

	return false
}

// SetChannelPost gets a reference to the given Message and assigns it to the ChannelPost field.
func (o *Update) SetChannelPost(v Message) {
	o.ChannelPost = &v
}


// GetEditedChannelPost returns the EditedChannelPost field value if set, zero value otherwise.
func (o *Update) GetEditedChannelPost() Message {
	if o == nil || IsNil(o.EditedChannelPost) {
		var ret Message
		return ret
	}
	return *o.EditedChannelPost
}

// GetEditedChannelPostOk returns a tuple with the EditedChannelPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetEditedChannelPostOk() (*Message, bool) {
	if o == nil || IsNil(o.EditedChannelPost) {
		return nil, false
	}
	return o.EditedChannelPost, true
}

// HasEditedChannelPost returns a boolean if a field has been set.
func (o *Update) HasEditedChannelPost() bool {
	if o != nil && !IsNil(o.EditedChannelPost) {
		return true
	}

	return false
}

// SetEditedChannelPost gets a reference to the given Message and assigns it to the EditedChannelPost field.
func (o *Update) SetEditedChannelPost(v Message) {
	o.EditedChannelPost = &v
}


// GetBusinessConnection returns the BusinessConnection field value if set, zero value otherwise.
func (o *Update) GetBusinessConnection() BusinessConnection {
	if o == nil || IsNil(o.BusinessConnection) {
		var ret BusinessConnection
		return ret
	}
	return *o.BusinessConnection
}

// GetBusinessConnectionOk returns a tuple with the BusinessConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetBusinessConnectionOk() (*BusinessConnection, bool) {
	if o == nil || IsNil(o.BusinessConnection) {
		return nil, false
	}
	return o.BusinessConnection, true
}

// HasBusinessConnection returns a boolean if a field has been set.
func (o *Update) HasBusinessConnection() bool {
	if o != nil && !IsNil(o.BusinessConnection) {
		return true
	}

	return false
}

// SetBusinessConnection gets a reference to the given BusinessConnection and assigns it to the BusinessConnection field.
func (o *Update) SetBusinessConnection(v BusinessConnection) {
	o.BusinessConnection = &v
}


// GetBusinessMessage returns the BusinessMessage field value if set, zero value otherwise.
func (o *Update) GetBusinessMessage() Message {
	if o == nil || IsNil(o.BusinessMessage) {
		var ret Message
		return ret
	}
	return *o.BusinessMessage
}

// GetBusinessMessageOk returns a tuple with the BusinessMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetBusinessMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.BusinessMessage) {
		return nil, false
	}
	return o.BusinessMessage, true
}

// HasBusinessMessage returns a boolean if a field has been set.
func (o *Update) HasBusinessMessage() bool {
	if o != nil && !IsNil(o.BusinessMessage) {
		return true
	}

	return false
}

// SetBusinessMessage gets a reference to the given Message and assigns it to the BusinessMessage field.
func (o *Update) SetBusinessMessage(v Message) {
	o.BusinessMessage = &v
}


// GetEditedBusinessMessage returns the EditedBusinessMessage field value if set, zero value otherwise.
func (o *Update) GetEditedBusinessMessage() Message {
	if o == nil || IsNil(o.EditedBusinessMessage) {
		var ret Message
		return ret
	}
	return *o.EditedBusinessMessage
}

// GetEditedBusinessMessageOk returns a tuple with the EditedBusinessMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetEditedBusinessMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.EditedBusinessMessage) {
		return nil, false
	}
	return o.EditedBusinessMessage, true
}

// HasEditedBusinessMessage returns a boolean if a field has been set.
func (o *Update) HasEditedBusinessMessage() bool {
	if o != nil && !IsNil(o.EditedBusinessMessage) {
		return true
	}

	return false
}

// SetEditedBusinessMessage gets a reference to the given Message and assigns it to the EditedBusinessMessage field.
func (o *Update) SetEditedBusinessMessage(v Message) {
	o.EditedBusinessMessage = &v
}


// GetDeletedBusinessMessages returns the DeletedBusinessMessages field value if set, zero value otherwise.
func (o *Update) GetDeletedBusinessMessages() BusinessMessagesDeleted {
	if o == nil || IsNil(o.DeletedBusinessMessages) {
		var ret BusinessMessagesDeleted
		return ret
	}
	return *o.DeletedBusinessMessages
}

// GetDeletedBusinessMessagesOk returns a tuple with the DeletedBusinessMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetDeletedBusinessMessagesOk() (*BusinessMessagesDeleted, bool) {
	if o == nil || IsNil(o.DeletedBusinessMessages) {
		return nil, false
	}
	return o.DeletedBusinessMessages, true
}

// HasDeletedBusinessMessages returns a boolean if a field has been set.
func (o *Update) HasDeletedBusinessMessages() bool {
	if o != nil && !IsNil(o.DeletedBusinessMessages) {
		return true
	}

	return false
}

// SetDeletedBusinessMessages gets a reference to the given BusinessMessagesDeleted and assigns it to the DeletedBusinessMessages field.
func (o *Update) SetDeletedBusinessMessages(v BusinessMessagesDeleted) {
	o.DeletedBusinessMessages = &v
}


// GetMessageReaction returns the MessageReaction field value if set, zero value otherwise.
func (o *Update) GetMessageReaction() MessageReactionUpdated {
	if o == nil || IsNil(o.MessageReaction) {
		var ret MessageReactionUpdated
		return ret
	}
	return *o.MessageReaction
}

// GetMessageReactionOk returns a tuple with the MessageReaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetMessageReactionOk() (*MessageReactionUpdated, bool) {
	if o == nil || IsNil(o.MessageReaction) {
		return nil, false
	}
	return o.MessageReaction, true
}

// HasMessageReaction returns a boolean if a field has been set.
func (o *Update) HasMessageReaction() bool {
	if o != nil && !IsNil(o.MessageReaction) {
		return true
	}

	return false
}

// SetMessageReaction gets a reference to the given MessageReactionUpdated and assigns it to the MessageReaction field.
func (o *Update) SetMessageReaction(v MessageReactionUpdated) {
	o.MessageReaction = &v
}


// GetMessageReactionCount returns the MessageReactionCount field value if set, zero value otherwise.
func (o *Update) GetMessageReactionCount() MessageReactionCountUpdated {
	if o == nil || IsNil(o.MessageReactionCount) {
		var ret MessageReactionCountUpdated
		return ret
	}
	return *o.MessageReactionCount
}

// GetMessageReactionCountOk returns a tuple with the MessageReactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetMessageReactionCountOk() (*MessageReactionCountUpdated, bool) {
	if o == nil || IsNil(o.MessageReactionCount) {
		return nil, false
	}
	return o.MessageReactionCount, true
}

// HasMessageReactionCount returns a boolean if a field has been set.
func (o *Update) HasMessageReactionCount() bool {
	if o != nil && !IsNil(o.MessageReactionCount) {
		return true
	}

	return false
}

// SetMessageReactionCount gets a reference to the given MessageReactionCountUpdated and assigns it to the MessageReactionCount field.
func (o *Update) SetMessageReactionCount(v MessageReactionCountUpdated) {
	o.MessageReactionCount = &v
}


// GetInlineQuery returns the InlineQuery field value if set, zero value otherwise.
func (o *Update) GetInlineQuery() InlineQuery {
	if o == nil || IsNil(o.InlineQuery) {
		var ret InlineQuery
		return ret
	}
	return *o.InlineQuery
}

// GetInlineQueryOk returns a tuple with the InlineQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetInlineQueryOk() (*InlineQuery, bool) {
	if o == nil || IsNil(o.InlineQuery) {
		return nil, false
	}
	return o.InlineQuery, true
}

// HasInlineQuery returns a boolean if a field has been set.
func (o *Update) HasInlineQuery() bool {
	if o != nil && !IsNil(o.InlineQuery) {
		return true
	}

	return false
}

// SetInlineQuery gets a reference to the given InlineQuery and assigns it to the InlineQuery field.
func (o *Update) SetInlineQuery(v InlineQuery) {
	o.InlineQuery = &v
}


// GetChosenInlineResult returns the ChosenInlineResult field value if set, zero value otherwise.
func (o *Update) GetChosenInlineResult() ChosenInlineResult {
	if o == nil || IsNil(o.ChosenInlineResult) {
		var ret ChosenInlineResult
		return ret
	}
	return *o.ChosenInlineResult
}

// GetChosenInlineResultOk returns a tuple with the ChosenInlineResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetChosenInlineResultOk() (*ChosenInlineResult, bool) {
	if o == nil || IsNil(o.ChosenInlineResult) {
		return nil, false
	}
	return o.ChosenInlineResult, true
}

// HasChosenInlineResult returns a boolean if a field has been set.
func (o *Update) HasChosenInlineResult() bool {
	if o != nil && !IsNil(o.ChosenInlineResult) {
		return true
	}

	return false
}

// SetChosenInlineResult gets a reference to the given ChosenInlineResult and assigns it to the ChosenInlineResult field.
func (o *Update) SetChosenInlineResult(v ChosenInlineResult) {
	o.ChosenInlineResult = &v
}


// GetCallbackQuery returns the CallbackQuery field value if set, zero value otherwise.
func (o *Update) GetCallbackQuery() CallbackQuery {
	if o == nil || IsNil(o.CallbackQuery) {
		var ret CallbackQuery
		return ret
	}
	return *o.CallbackQuery
}

// GetCallbackQueryOk returns a tuple with the CallbackQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetCallbackQueryOk() (*CallbackQuery, bool) {
	if o == nil || IsNil(o.CallbackQuery) {
		return nil, false
	}
	return o.CallbackQuery, true
}

// HasCallbackQuery returns a boolean if a field has been set.
func (o *Update) HasCallbackQuery() bool {
	if o != nil && !IsNil(o.CallbackQuery) {
		return true
	}

	return false
}

// SetCallbackQuery gets a reference to the given CallbackQuery and assigns it to the CallbackQuery field.
func (o *Update) SetCallbackQuery(v CallbackQuery) {
	o.CallbackQuery = &v
}


// GetShippingQuery returns the ShippingQuery field value if set, zero value otherwise.
func (o *Update) GetShippingQuery() ShippingQuery {
	if o == nil || IsNil(o.ShippingQuery) {
		var ret ShippingQuery
		return ret
	}
	return *o.ShippingQuery
}

// GetShippingQueryOk returns a tuple with the ShippingQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetShippingQueryOk() (*ShippingQuery, bool) {
	if o == nil || IsNil(o.ShippingQuery) {
		return nil, false
	}
	return o.ShippingQuery, true
}

// HasShippingQuery returns a boolean if a field has been set.
func (o *Update) HasShippingQuery() bool {
	if o != nil && !IsNil(o.ShippingQuery) {
		return true
	}

	return false
}

// SetShippingQuery gets a reference to the given ShippingQuery and assigns it to the ShippingQuery field.
func (o *Update) SetShippingQuery(v ShippingQuery) {
	o.ShippingQuery = &v
}


// GetPreCheckoutQuery returns the PreCheckoutQuery field value if set, zero value otherwise.
func (o *Update) GetPreCheckoutQuery() PreCheckoutQuery {
	if o == nil || IsNil(o.PreCheckoutQuery) {
		var ret PreCheckoutQuery
		return ret
	}
	return *o.PreCheckoutQuery
}

// GetPreCheckoutQueryOk returns a tuple with the PreCheckoutQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetPreCheckoutQueryOk() (*PreCheckoutQuery, bool) {
	if o == nil || IsNil(o.PreCheckoutQuery) {
		return nil, false
	}
	return o.PreCheckoutQuery, true
}

// HasPreCheckoutQuery returns a boolean if a field has been set.
func (o *Update) HasPreCheckoutQuery() bool {
	if o != nil && !IsNil(o.PreCheckoutQuery) {
		return true
	}

	return false
}

// SetPreCheckoutQuery gets a reference to the given PreCheckoutQuery and assigns it to the PreCheckoutQuery field.
func (o *Update) SetPreCheckoutQuery(v PreCheckoutQuery) {
	o.PreCheckoutQuery = &v
}


// GetPurchasedPaidMedia returns the PurchasedPaidMedia field value if set, zero value otherwise.
func (o *Update) GetPurchasedPaidMedia() PaidMediaPurchased {
	if o == nil || IsNil(o.PurchasedPaidMedia) {
		var ret PaidMediaPurchased
		return ret
	}
	return *o.PurchasedPaidMedia
}

// GetPurchasedPaidMediaOk returns a tuple with the PurchasedPaidMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetPurchasedPaidMediaOk() (*PaidMediaPurchased, bool) {
	if o == nil || IsNil(o.PurchasedPaidMedia) {
		return nil, false
	}
	return o.PurchasedPaidMedia, true
}

// HasPurchasedPaidMedia returns a boolean if a field has been set.
func (o *Update) HasPurchasedPaidMedia() bool {
	if o != nil && !IsNil(o.PurchasedPaidMedia) {
		return true
	}

	return false
}

// SetPurchasedPaidMedia gets a reference to the given PaidMediaPurchased and assigns it to the PurchasedPaidMedia field.
func (o *Update) SetPurchasedPaidMedia(v PaidMediaPurchased) {
	o.PurchasedPaidMedia = &v
}


// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *Update) GetPoll() Poll {
	if o == nil || IsNil(o.Poll) {
		var ret Poll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetPollOk() (*Poll, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *Update) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given Poll and assigns it to the Poll field.
func (o *Update) SetPoll(v Poll) {
	o.Poll = &v
}


// GetPollAnswer returns the PollAnswer field value if set, zero value otherwise.
func (o *Update) GetPollAnswer() PollAnswer {
	if o == nil || IsNil(o.PollAnswer) {
		var ret PollAnswer
		return ret
	}
	return *o.PollAnswer
}

// GetPollAnswerOk returns a tuple with the PollAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetPollAnswerOk() (*PollAnswer, bool) {
	if o == nil || IsNil(o.PollAnswer) {
		return nil, false
	}
	return o.PollAnswer, true
}

// HasPollAnswer returns a boolean if a field has been set.
func (o *Update) HasPollAnswer() bool {
	if o != nil && !IsNil(o.PollAnswer) {
		return true
	}

	return false
}

// SetPollAnswer gets a reference to the given PollAnswer and assigns it to the PollAnswer field.
func (o *Update) SetPollAnswer(v PollAnswer) {
	o.PollAnswer = &v
}


// GetMyChatMember returns the MyChatMember field value if set, zero value otherwise.
func (o *Update) GetMyChatMember() ChatMemberUpdated {
	if o == nil || IsNil(o.MyChatMember) {
		var ret ChatMemberUpdated
		return ret
	}
	return *o.MyChatMember
}

// GetMyChatMemberOk returns a tuple with the MyChatMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetMyChatMemberOk() (*ChatMemberUpdated, bool) {
	if o == nil || IsNil(o.MyChatMember) {
		return nil, false
	}
	return o.MyChatMember, true
}

// HasMyChatMember returns a boolean if a field has been set.
func (o *Update) HasMyChatMember() bool {
	if o != nil && !IsNil(o.MyChatMember) {
		return true
	}

	return false
}

// SetMyChatMember gets a reference to the given ChatMemberUpdated and assigns it to the MyChatMember field.
func (o *Update) SetMyChatMember(v ChatMemberUpdated) {
	o.MyChatMember = &v
}


// GetChatMember returns the ChatMember field value if set, zero value otherwise.
func (o *Update) GetChatMember() ChatMemberUpdated {
	if o == nil || IsNil(o.ChatMember) {
		var ret ChatMemberUpdated
		return ret
	}
	return *o.ChatMember
}

// GetChatMemberOk returns a tuple with the ChatMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetChatMemberOk() (*ChatMemberUpdated, bool) {
	if o == nil || IsNil(o.ChatMember) {
		return nil, false
	}
	return o.ChatMember, true
}

// HasChatMember returns a boolean if a field has been set.
func (o *Update) HasChatMember() bool {
	if o != nil && !IsNil(o.ChatMember) {
		return true
	}

	return false
}

// SetChatMember gets a reference to the given ChatMemberUpdated and assigns it to the ChatMember field.
func (o *Update) SetChatMember(v ChatMemberUpdated) {
	o.ChatMember = &v
}


// GetChatJoinRequest returns the ChatJoinRequest field value if set, zero value otherwise.
func (o *Update) GetChatJoinRequest() ChatJoinRequest {
	if o == nil || IsNil(o.ChatJoinRequest) {
		var ret ChatJoinRequest
		return ret
	}
	return *o.ChatJoinRequest
}

// GetChatJoinRequestOk returns a tuple with the ChatJoinRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetChatJoinRequestOk() (*ChatJoinRequest, bool) {
	if o == nil || IsNil(o.ChatJoinRequest) {
		return nil, false
	}
	return o.ChatJoinRequest, true
}

// HasChatJoinRequest returns a boolean if a field has been set.
func (o *Update) HasChatJoinRequest() bool {
	if o != nil && !IsNil(o.ChatJoinRequest) {
		return true
	}

	return false
}

// SetChatJoinRequest gets a reference to the given ChatJoinRequest and assigns it to the ChatJoinRequest field.
func (o *Update) SetChatJoinRequest(v ChatJoinRequest) {
	o.ChatJoinRequest = &v
}


// GetChatBoost returns the ChatBoost field value if set, zero value otherwise.
func (o *Update) GetChatBoost() ChatBoostUpdated {
	if o == nil || IsNil(o.ChatBoost) {
		var ret ChatBoostUpdated
		return ret
	}
	return *o.ChatBoost
}

// GetChatBoostOk returns a tuple with the ChatBoost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetChatBoostOk() (*ChatBoostUpdated, bool) {
	if o == nil || IsNil(o.ChatBoost) {
		return nil, false
	}
	return o.ChatBoost, true
}

// HasChatBoost returns a boolean if a field has been set.
func (o *Update) HasChatBoost() bool {
	if o != nil && !IsNil(o.ChatBoost) {
		return true
	}

	return false
}

// SetChatBoost gets a reference to the given ChatBoostUpdated and assigns it to the ChatBoost field.
func (o *Update) SetChatBoost(v ChatBoostUpdated) {
	o.ChatBoost = &v
}


// GetRemovedChatBoost returns the RemovedChatBoost field value if set, zero value otherwise.
func (o *Update) GetRemovedChatBoost() ChatBoostRemoved {
	if o == nil || IsNil(o.RemovedChatBoost) {
		var ret ChatBoostRemoved
		return ret
	}
	return *o.RemovedChatBoost
}

// GetRemovedChatBoostOk returns a tuple with the RemovedChatBoost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetRemovedChatBoostOk() (*ChatBoostRemoved, bool) {
	if o == nil || IsNil(o.RemovedChatBoost) {
		return nil, false
	}
	return o.RemovedChatBoost, true
}

// HasRemovedChatBoost returns a boolean if a field has been set.
func (o *Update) HasRemovedChatBoost() bool {
	if o != nil && !IsNil(o.RemovedChatBoost) {
		return true
	}

	return false
}

// SetRemovedChatBoost gets a reference to the given ChatBoostRemoved and assigns it to the RemovedChatBoost field.
func (o *Update) SetRemovedChatBoost(v ChatBoostRemoved) {
	o.RemovedChatBoost = &v
}


func (o Update) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Update) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["update_id"] = o.UpdateId
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.EditedMessage) {
		toSerialize["edited_message"] = o.EditedMessage
	}
	if !IsNil(o.ChannelPost) {
		toSerialize["channel_post"] = o.ChannelPost
	}
	if !IsNil(o.EditedChannelPost) {
		toSerialize["edited_channel_post"] = o.EditedChannelPost
	}
	if !IsNil(o.BusinessConnection) {
		toSerialize["business_connection"] = o.BusinessConnection
	}
	if !IsNil(o.BusinessMessage) {
		toSerialize["business_message"] = o.BusinessMessage
	}
	if !IsNil(o.EditedBusinessMessage) {
		toSerialize["edited_business_message"] = o.EditedBusinessMessage
	}
	if !IsNil(o.DeletedBusinessMessages) {
		toSerialize["deleted_business_messages"] = o.DeletedBusinessMessages
	}
	if !IsNil(o.MessageReaction) {
		toSerialize["message_reaction"] = o.MessageReaction
	}
	if !IsNil(o.MessageReactionCount) {
		toSerialize["message_reaction_count"] = o.MessageReactionCount
	}
	if !IsNil(o.InlineQuery) {
		toSerialize["inline_query"] = o.InlineQuery
	}
	if !IsNil(o.ChosenInlineResult) {
		toSerialize["chosen_inline_result"] = o.ChosenInlineResult
	}
	if !IsNil(o.CallbackQuery) {
		toSerialize["callback_query"] = o.CallbackQuery
	}
	if !IsNil(o.ShippingQuery) {
		toSerialize["shipping_query"] = o.ShippingQuery
	}
	if !IsNil(o.PreCheckoutQuery) {
		toSerialize["pre_checkout_query"] = o.PreCheckoutQuery
	}
	if !IsNil(o.PurchasedPaidMedia) {
		toSerialize["purchased_paid_media"] = o.PurchasedPaidMedia
	}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}
	if !IsNil(o.PollAnswer) {
		toSerialize["poll_answer"] = o.PollAnswer
	}
	if !IsNil(o.MyChatMember) {
		toSerialize["my_chat_member"] = o.MyChatMember
	}
	if !IsNil(o.ChatMember) {
		toSerialize["chat_member"] = o.ChatMember
	}
	if !IsNil(o.ChatJoinRequest) {
		toSerialize["chat_join_request"] = o.ChatJoinRequest
	}
	if !IsNil(o.ChatBoost) {
		toSerialize["chat_boost"] = o.ChatBoost
	}
	if !IsNil(o.RemovedChatBoost) {
		toSerialize["removed_chat_boost"] = o.RemovedChatBoost
	}
	return toSerialize, nil
}

func (o *Update) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"update_id",
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

	varUpdate := _Update{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdate)

	if err != nil {
		return err
	}

	*o = Update(varUpdate)

	return err
}

type NullableUpdate struct {
	value *Update
	isSet bool
}

func (v NullableUpdate) Get() *Update {
	return v.value
}

func (v *NullableUpdate) Set(val *Update) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdate(val *Update) *NullableUpdate {
	return &NullableUpdate{value: val, isSet: true}
}

func (v NullableUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


