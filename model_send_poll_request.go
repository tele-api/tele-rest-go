/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the SendPollRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendPollRequest{}

// SendPollRequest Request parameters for sendPoll
type SendPollRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// Mode for parsing entities in the question. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Currently, only custom emoji entities are allowed
	QuestionParseMode *string `json:"question_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of *question\\_parse\\_mode*
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`
	// A JSON-serialized list of 2-12 answer options
	Options []InputPollOption `json:"options"`
	// *True*, if the poll needs to be anonymous, defaults to *True*
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type *string `json:"type,omitempty"`
	// *True*, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to *False*
	AllowsMultipleAnswers *bool `json:"allows_multiple_answers,omitempty"`
	// 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionId *int32 `json:"correct_option_id,omitempty"`
	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation *string `json:"explanation,omitempty"`
	// Mode for parsing entities in the explanation. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ExplanationParseMode *string `json:"explanation_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of *explanation\\_parse\\_mode*
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with *close\\_date*.
	OpenPeriod *int32 `json:"open_period,omitempty"`
	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with *open\\_period*.
	CloseDate *int32 `json:"close_date,omitempty"`
	// Pass *True* if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed *bool `json:"is_closed,omitempty"`
	// Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup *SendMessageRequestReplyMarkup `json:"reply_markup,omitempty"`
}

type _SendPollRequest SendPollRequest

// NewSendPollRequest instantiates a new SendPollRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendPollRequest(chatId SendMessageRequestChatId, question string, options []InputPollOption) *SendPollRequest {
	this := SendPollRequest{}
	this.ChatId = chatId
	this.Question = question
	this.Options = options
	return &this
}

// NewSendPollRequestWithDefaults instantiates a new SendPollRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendPollRequestWithDefaults() *SendPollRequest {
	this := SendPollRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *SendPollRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *SendPollRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *SendPollRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *SendPollRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendPollRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendPollRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendPollRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendPollRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetQuestion returns the Question field value
func (o *SendPollRequest) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *SendPollRequest) SetQuestion(v string) {
	o.Question = v
}

// GetQuestionParseMode returns the QuestionParseMode field value if set, zero value otherwise.
func (o *SendPollRequest) GetQuestionParseMode() string {
	if o == nil || IsNil(o.QuestionParseMode) {
		var ret string
		return ret
	}
	return *o.QuestionParseMode
}

// GetQuestionParseModeOk returns a tuple with the QuestionParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetQuestionParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionParseMode) {
		return nil, false
	}
	return o.QuestionParseMode, true
}

// HasQuestionParseMode returns a boolean if a field has been set.
func (o *SendPollRequest) HasQuestionParseMode() bool {
	if o != nil && !IsNil(o.QuestionParseMode) {
		return true
	}

	return false
}

// SetQuestionParseMode gets a reference to the given string and assigns it to the QuestionParseMode field.
func (o *SendPollRequest) SetQuestionParseMode(v string) {
	o.QuestionParseMode = &v
}


// GetQuestionEntities returns the QuestionEntities field value if set, zero value otherwise.
func (o *SendPollRequest) GetQuestionEntities() []MessageEntity {
	if o == nil || IsNil(o.QuestionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.QuestionEntities
}

// GetQuestionEntitiesOk returns a tuple with the QuestionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetQuestionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.QuestionEntities) {
		return nil, false
	}
	return o.QuestionEntities, true
}

// HasQuestionEntities returns a boolean if a field has been set.
func (o *SendPollRequest) HasQuestionEntities() bool {
	if o != nil && !IsNil(o.QuestionEntities) {
		return true
	}

	return false
}

// SetQuestionEntities gets a reference to the given []MessageEntity and assigns it to the QuestionEntities field.
func (o *SendPollRequest) SetQuestionEntities(v []MessageEntity) {
	o.QuestionEntities = v
}


// GetOptions returns the Options field value
func (o *SendPollRequest) GetOptions() []InputPollOption {
	if o == nil {
		var ret []InputPollOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetOptionsOk() ([]InputPollOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *SendPollRequest) SetOptions(v []InputPollOption) {
	o.Options = v
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise.
func (o *SendPollRequest) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetIsAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymous) {
		return nil, false
	}
	return o.IsAnonymous, true
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *SendPollRequest) HasIsAnonymous() bool {
	if o != nil && !IsNil(o.IsAnonymous) {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given bool and assigns it to the IsAnonymous field.
func (o *SendPollRequest) SetIsAnonymous(v bool) {
	o.IsAnonymous = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *SendPollRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SendPollRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SendPollRequest) SetType(v string) {
	o.Type = &v
}


// GetAllowsMultipleAnswers returns the AllowsMultipleAnswers field value if set, zero value otherwise.
func (o *SendPollRequest) GetAllowsMultipleAnswers() bool {
	if o == nil || IsNil(o.AllowsMultipleAnswers) {
		var ret bool
		return ret
	}
	return *o.AllowsMultipleAnswers
}

// GetAllowsMultipleAnswersOk returns a tuple with the AllowsMultipleAnswers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetAllowsMultipleAnswersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowsMultipleAnswers) {
		return nil, false
	}
	return o.AllowsMultipleAnswers, true
}

// HasAllowsMultipleAnswers returns a boolean if a field has been set.
func (o *SendPollRequest) HasAllowsMultipleAnswers() bool {
	if o != nil && !IsNil(o.AllowsMultipleAnswers) {
		return true
	}

	return false
}

// SetAllowsMultipleAnswers gets a reference to the given bool and assigns it to the AllowsMultipleAnswers field.
func (o *SendPollRequest) SetAllowsMultipleAnswers(v bool) {
	o.AllowsMultipleAnswers = &v
}


// GetCorrectOptionId returns the CorrectOptionId field value if set, zero value otherwise.
func (o *SendPollRequest) GetCorrectOptionId() int32 {
	if o == nil || IsNil(o.CorrectOptionId) {
		var ret int32
		return ret
	}
	return *o.CorrectOptionId
}

// GetCorrectOptionIdOk returns a tuple with the CorrectOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetCorrectOptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CorrectOptionId) {
		return nil, false
	}
	return o.CorrectOptionId, true
}

// HasCorrectOptionId returns a boolean if a field has been set.
func (o *SendPollRequest) HasCorrectOptionId() bool {
	if o != nil && !IsNil(o.CorrectOptionId) {
		return true
	}

	return false
}

// SetCorrectOptionId gets a reference to the given int32 and assigns it to the CorrectOptionId field.
func (o *SendPollRequest) SetCorrectOptionId(v int32) {
	o.CorrectOptionId = &v
}


// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *SendPollRequest) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *SendPollRequest) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *SendPollRequest) SetExplanation(v string) {
	o.Explanation = &v
}


// GetExplanationParseMode returns the ExplanationParseMode field value if set, zero value otherwise.
func (o *SendPollRequest) GetExplanationParseMode() string {
	if o == nil || IsNil(o.ExplanationParseMode) {
		var ret string
		return ret
	}
	return *o.ExplanationParseMode
}

// GetExplanationParseModeOk returns a tuple with the ExplanationParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetExplanationParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ExplanationParseMode) {
		return nil, false
	}
	return o.ExplanationParseMode, true
}

// HasExplanationParseMode returns a boolean if a field has been set.
func (o *SendPollRequest) HasExplanationParseMode() bool {
	if o != nil && !IsNil(o.ExplanationParseMode) {
		return true
	}

	return false
}

// SetExplanationParseMode gets a reference to the given string and assigns it to the ExplanationParseMode field.
func (o *SendPollRequest) SetExplanationParseMode(v string) {
	o.ExplanationParseMode = &v
}


// GetExplanationEntities returns the ExplanationEntities field value if set, zero value otherwise.
func (o *SendPollRequest) GetExplanationEntities() []MessageEntity {
	if o == nil || IsNil(o.ExplanationEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.ExplanationEntities
}

// GetExplanationEntitiesOk returns a tuple with the ExplanationEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetExplanationEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.ExplanationEntities) {
		return nil, false
	}
	return o.ExplanationEntities, true
}

// HasExplanationEntities returns a boolean if a field has been set.
func (o *SendPollRequest) HasExplanationEntities() bool {
	if o != nil && !IsNil(o.ExplanationEntities) {
		return true
	}

	return false
}

// SetExplanationEntities gets a reference to the given []MessageEntity and assigns it to the ExplanationEntities field.
func (o *SendPollRequest) SetExplanationEntities(v []MessageEntity) {
	o.ExplanationEntities = v
}


// GetOpenPeriod returns the OpenPeriod field value if set, zero value otherwise.
func (o *SendPollRequest) GetOpenPeriod() int32 {
	if o == nil || IsNil(o.OpenPeriod) {
		var ret int32
		return ret
	}
	return *o.OpenPeriod
}

// GetOpenPeriodOk returns a tuple with the OpenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetOpenPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenPeriod) {
		return nil, false
	}
	return o.OpenPeriod, true
}

// HasOpenPeriod returns a boolean if a field has been set.
func (o *SendPollRequest) HasOpenPeriod() bool {
	if o != nil && !IsNil(o.OpenPeriod) {
		return true
	}

	return false
}

// SetOpenPeriod gets a reference to the given int32 and assigns it to the OpenPeriod field.
func (o *SendPollRequest) SetOpenPeriod(v int32) {
	o.OpenPeriod = &v
}


// GetCloseDate returns the CloseDate field value if set, zero value otherwise.
func (o *SendPollRequest) GetCloseDate() int32 {
	if o == nil || IsNil(o.CloseDate) {
		var ret int32
		return ret
	}
	return *o.CloseDate
}

// GetCloseDateOk returns a tuple with the CloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetCloseDateOk() (*int32, bool) {
	if o == nil || IsNil(o.CloseDate) {
		return nil, false
	}
	return o.CloseDate, true
}

// HasCloseDate returns a boolean if a field has been set.
func (o *SendPollRequest) HasCloseDate() bool {
	if o != nil && !IsNil(o.CloseDate) {
		return true
	}

	return false
}

// SetCloseDate gets a reference to the given int32 and assigns it to the CloseDate field.
func (o *SendPollRequest) SetCloseDate(v int32) {
	o.CloseDate = &v
}


// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *SendPollRequest) GetIsClosed() bool {
	if o == nil || IsNil(o.IsClosed) {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetIsClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsClosed) {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *SendPollRequest) HasIsClosed() bool {
	if o != nil && !IsNil(o.IsClosed) {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *SendPollRequest) SetIsClosed(v bool) {
	o.IsClosed = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *SendPollRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *SendPollRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *SendPollRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *SendPollRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *SendPollRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *SendPollRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *SendPollRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *SendPollRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *SendPollRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetMessageEffectId returns the MessageEffectId field value if set, zero value otherwise.
func (o *SendPollRequest) GetMessageEffectId() string {
	if o == nil || IsNil(o.MessageEffectId) {
		var ret string
		return ret
	}
	return *o.MessageEffectId
}

// GetMessageEffectIdOk returns a tuple with the MessageEffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetMessageEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageEffectId) {
		return nil, false
	}
	return o.MessageEffectId, true
}

// HasMessageEffectId returns a boolean if a field has been set.
func (o *SendPollRequest) HasMessageEffectId() bool {
	if o != nil && !IsNil(o.MessageEffectId) {
		return true
	}

	return false
}

// SetMessageEffectId gets a reference to the given string and assigns it to the MessageEffectId field.
func (o *SendPollRequest) SetMessageEffectId(v string) {
	o.MessageEffectId = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *SendPollRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *SendPollRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *SendPollRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *SendPollRequest) GetReplyMarkup() SendMessageRequestReplyMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret SendMessageRequestReplyMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPollRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *SendPollRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given SendMessageRequestReplyMarkup and assigns it to the ReplyMarkup field.
func (o *SendPollRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup) {
	o.ReplyMarkup = &v
}


func (o SendPollRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendPollRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["question"] = o.Question
	if !IsNil(o.QuestionParseMode) {
		toSerialize["question_parse_mode"] = o.QuestionParseMode
	}
	if !IsNil(o.QuestionEntities) {
		toSerialize["question_entities"] = o.QuestionEntities
	}
	toSerialize["options"] = o.Options
	if !IsNil(o.IsAnonymous) {
		toSerialize["is_anonymous"] = o.IsAnonymous
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AllowsMultipleAnswers) {
		toSerialize["allows_multiple_answers"] = o.AllowsMultipleAnswers
	}
	if !IsNil(o.CorrectOptionId) {
		toSerialize["correct_option_id"] = o.CorrectOptionId
	}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	if !IsNil(o.ExplanationParseMode) {
		toSerialize["explanation_parse_mode"] = o.ExplanationParseMode
	}
	if !IsNil(o.ExplanationEntities) {
		toSerialize["explanation_entities"] = o.ExplanationEntities
	}
	if !IsNil(o.OpenPeriod) {
		toSerialize["open_period"] = o.OpenPeriod
	}
	if !IsNil(o.CloseDate) {
		toSerialize["close_date"] = o.CloseDate
	}
	if !IsNil(o.IsClosed) {
		toSerialize["is_closed"] = o.IsClosed
	}
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	if !IsNil(o.AllowPaidBroadcast) {
		toSerialize["allow_paid_broadcast"] = o.AllowPaidBroadcast
	}
	if !IsNil(o.MessageEffectId) {
		toSerialize["message_effect_id"] = o.MessageEffectId
	}
	if !IsNil(o.ReplyParameters) {
		toSerialize["reply_parameters"] = o.ReplyParameters
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *SendPollRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"question",
		"options",
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

	varSendPollRequest := _SendPollRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendPollRequest)

	if err != nil {
		return err
	}

	*o = SendPollRequest(varSendPollRequest)

	return err
}

type NullableSendPollRequest struct {
	value *SendPollRequest
	isSet bool
}

func (v NullableSendPollRequest) Get() *SendPollRequest {
	return v.value
}

func (v *NullableSendPollRequest) Set(val *SendPollRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendPollRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendPollRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendPollRequest(val *SendPollRequest) *NullableSendPollRequest {
	return &NullableSendPollRequest{value: val, isSet: true}
}

func (v NullableSendPollRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendPollRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


