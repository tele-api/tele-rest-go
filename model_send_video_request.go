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

// checks if the SendVideoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendVideoRequest{}

// SendVideoRequest Request parameters for sendVideo
type SendVideoRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	Video NullableString `json:"video"`
	// Duration of sent video in seconds
	Duration *int32 `json:"duration,omitempty"`
	// Video width
	Width *int32 `json:"width,omitempty"`
	// Video height
	Height *int32 `json:"height,omitempty"`
	Thumbnail NullableString `json:"thumbnail,omitempty"`
	Cover NullableString `json:"cover,omitempty"`
	// Start timestamp for the video in the message
	StartTimestamp *int32 `json:"start_timestamp,omitempty"`
	// Video caption (may also be used when resending videos by *file\\_id*), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Pass *True*, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Pass *True* if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	// Pass *True* if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
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

type _SendVideoRequest SendVideoRequest

// NewSendVideoRequest instantiates a new SendVideoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendVideoRequest(chatId SendMessageRequestChatId, video NullableString) *SendVideoRequest {
	this := SendVideoRequest{}
	this.ChatId = chatId
	this.Video = video
	return &this
}

// NewSendVideoRequestWithDefaults instantiates a new SendVideoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendVideoRequestWithDefaults() *SendVideoRequest {
	this := SendVideoRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *SendVideoRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *SendVideoRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *SendVideoRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *SendVideoRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendVideoRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendVideoRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendVideoRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendVideoRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetVideo returns the Video field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SendVideoRequest) GetVideo() string {
	if o == nil || o.Video.Get() == nil {
		var ret string
		return ret
	}

	return *o.Video.Get()
}

// GetVideoOk returns a tuple with the Video field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVideoRequest) GetVideoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Video.Get(), o.Video.IsSet()
}

// SetVideo sets field value
func (o *SendVideoRequest) SetVideo(v string) {
	o.Video.Set(&v)
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SendVideoRequest) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SendVideoRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *SendVideoRequest) SetDuration(v int32) {
	o.Duration = &v
}


// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SendVideoRequest) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SendVideoRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *SendVideoRequest) SetWidth(v int32) {
	o.Width = &v
}


// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SendVideoRequest) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SendVideoRequest) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *SendVideoRequest) SetHeight(v int32) {
	o.Height = &v
}


// GetThumbnail returns the Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendVideoRequest) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbnail.Get()
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVideoRequest) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		return nil, false
	}
	return o.Thumbnail.Get(), o.Thumbnail.IsSet()
}

// HasThumbnail returns a boolean if a field has been set.
func (o *SendVideoRequest) HasThumbnail() bool {
	if o != nil && o.Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given NullableString and assigns it to the Thumbnail field.
func (o *SendVideoRequest) SetThumbnail(v string) {
	o.Thumbnail.Set(&v)
}

// SetThumbnailNil sets the value for Thumbnail to be an explicit nil
func (o *SendVideoRequest) SetThumbnailNil() {
	o.Thumbnail.Set(nil)
}

// UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
func (o *SendVideoRequest) UnsetThumbnail() {
	o.Thumbnail.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendVideoRequest) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVideoRequest) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover.Get()) {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *SendVideoRequest) HasCover() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *SendVideoRequest) SetCover(v string) {
	o.Cover.Set(&v)
}

// SetCoverNil sets the value for Cover to be an explicit nil
func (o *SendVideoRequest) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *SendVideoRequest) UnsetCover() {
	o.Cover.Unset()
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *SendVideoRequest) GetStartTimestamp() int32 {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret int32
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *SendVideoRequest) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given int32 and assigns it to the StartTimestamp field.
func (o *SendVideoRequest) SetStartTimestamp(v int32) {
	o.StartTimestamp = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *SendVideoRequest) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *SendVideoRequest) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *SendVideoRequest) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *SendVideoRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *SendVideoRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *SendVideoRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *SendVideoRequest) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *SendVideoRequest) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *SendVideoRequest) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *SendVideoRequest) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *SendVideoRequest) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *SendVideoRequest) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetHasSpoiler returns the HasSpoiler field value if set, zero value otherwise.
func (o *SendVideoRequest) GetHasSpoiler() bool {
	if o == nil || IsNil(o.HasSpoiler) {
		var ret bool
		return ret
	}
	return *o.HasSpoiler
}

// GetHasSpoilerOk returns a tuple with the HasSpoiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetHasSpoilerOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSpoiler) {
		return nil, false
	}
	return o.HasSpoiler, true
}

// HasHasSpoiler returns a boolean if a field has been set.
func (o *SendVideoRequest) HasHasSpoiler() bool {
	if o != nil && !IsNil(o.HasSpoiler) {
		return true
	}

	return false
}

// SetHasSpoiler gets a reference to the given bool and assigns it to the HasSpoiler field.
func (o *SendVideoRequest) SetHasSpoiler(v bool) {
	o.HasSpoiler = &v
}


// GetSupportsStreaming returns the SupportsStreaming field value if set, zero value otherwise.
func (o *SendVideoRequest) GetSupportsStreaming() bool {
	if o == nil || IsNil(o.SupportsStreaming) {
		var ret bool
		return ret
	}
	return *o.SupportsStreaming
}

// GetSupportsStreamingOk returns a tuple with the SupportsStreaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetSupportsStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsStreaming) {
		return nil, false
	}
	return o.SupportsStreaming, true
}

// HasSupportsStreaming returns a boolean if a field has been set.
func (o *SendVideoRequest) HasSupportsStreaming() bool {
	if o != nil && !IsNil(o.SupportsStreaming) {
		return true
	}

	return false
}

// SetSupportsStreaming gets a reference to the given bool and assigns it to the SupportsStreaming field.
func (o *SendVideoRequest) SetSupportsStreaming(v bool) {
	o.SupportsStreaming = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *SendVideoRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *SendVideoRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *SendVideoRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *SendVideoRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *SendVideoRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *SendVideoRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *SendVideoRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *SendVideoRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *SendVideoRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetMessageEffectId returns the MessageEffectId field value if set, zero value otherwise.
func (o *SendVideoRequest) GetMessageEffectId() string {
	if o == nil || IsNil(o.MessageEffectId) {
		var ret string
		return ret
	}
	return *o.MessageEffectId
}

// GetMessageEffectIdOk returns a tuple with the MessageEffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetMessageEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageEffectId) {
		return nil, false
	}
	return o.MessageEffectId, true
}

// HasMessageEffectId returns a boolean if a field has been set.
func (o *SendVideoRequest) HasMessageEffectId() bool {
	if o != nil && !IsNil(o.MessageEffectId) {
		return true
	}

	return false
}

// SetMessageEffectId gets a reference to the given string and assigns it to the MessageEffectId field.
func (o *SendVideoRequest) SetMessageEffectId(v string) {
	o.MessageEffectId = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *SendVideoRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *SendVideoRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *SendVideoRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *SendVideoRequest) GetReplyMarkup() SendMessageRequestReplyMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret SendMessageRequestReplyMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVideoRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *SendVideoRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given SendMessageRequestReplyMarkup and assigns it to the ReplyMarkup field.
func (o *SendVideoRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup) {
	o.ReplyMarkup = &v
}


func (o SendVideoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendVideoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["video"] = o.Video.Get()
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if o.Thumbnail.IsSet() {
		toSerialize["thumbnail"] = o.Thumbnail.Get()
	}
	if o.Cover.IsSet() {
		toSerialize["cover"] = o.Cover.Get()
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["start_timestamp"] = o.StartTimestamp
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.CaptionEntities) {
		toSerialize["caption_entities"] = o.CaptionEntities
	}
	if !IsNil(o.ShowCaptionAboveMedia) {
		toSerialize["show_caption_above_media"] = o.ShowCaptionAboveMedia
	}
	if !IsNil(o.HasSpoiler) {
		toSerialize["has_spoiler"] = o.HasSpoiler
	}
	if !IsNil(o.SupportsStreaming) {
		toSerialize["supports_streaming"] = o.SupportsStreaming
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

func (o *SendVideoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"video",
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

	varSendVideoRequest := _SendVideoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendVideoRequest)

	if err != nil {
		return err
	}

	*o = SendVideoRequest(varSendVideoRequest)

	return err
}

type NullableSendVideoRequest struct {
	value *SendVideoRequest
	isSet bool
}

func (v NullableSendVideoRequest) Get() *SendVideoRequest {
	return v.value
}

func (v *NullableSendVideoRequest) Set(val *SendVideoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendVideoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendVideoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendVideoRequest(val *SendVideoRequest) *NullableSendVideoRequest {
	return &NullableSendVideoRequest{value: val, isSet: true}
}

func (v NullableSendVideoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendVideoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


