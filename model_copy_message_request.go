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

// checks if the CopyMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyMessageRequest{}

// CopyMessageRequest Request parameters for copyMessage
type CopyMessageRequest struct {
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	FromChatId ForwardMessageRequestFromChatId `json:"from_chat_id"`
	// Message identifier in the chat specified in *from\\_chat\\_id*
	MessageId int32 `json:"message_id"`
	// New start timestamp for the copied video in the message
	VideoStartTimestamp *int32 `json:"video_start_timestamp,omitempty"`
	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the new caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Pass *True*, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup *SendMessageRequestReplyMarkup `json:"reply_markup,omitempty"`
}

type _CopyMessageRequest CopyMessageRequest

// NewCopyMessageRequest instantiates a new CopyMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyMessageRequest(chatId SendMessageRequestChatId, fromChatId ForwardMessageRequestFromChatId, messageId int32) *CopyMessageRequest {
	this := CopyMessageRequest{}
	this.ChatId = chatId
	this.FromChatId = fromChatId
	this.MessageId = messageId
	return &this
}

// NewCopyMessageRequestWithDefaults instantiates a new CopyMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyMessageRequestWithDefaults() *CopyMessageRequest {
	this := CopyMessageRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *CopyMessageRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *CopyMessageRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *CopyMessageRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetFromChatId returns the FromChatId field value
func (o *CopyMessageRequest) GetFromChatId() ForwardMessageRequestFromChatId {
	if o == nil {
		var ret ForwardMessageRequestFromChatId
		return ret
	}

	return o.FromChatId
}

// GetFromChatIdOk returns a tuple with the FromChatId field value
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetFromChatIdOk() (*ForwardMessageRequestFromChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromChatId, true
}

// SetFromChatId sets field value
func (o *CopyMessageRequest) SetFromChatId(v ForwardMessageRequestFromChatId) {
	o.FromChatId = v
}

// GetMessageId returns the MessageId field value
func (o *CopyMessageRequest) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *CopyMessageRequest) SetMessageId(v int32) {
	o.MessageId = v
}

// GetVideoStartTimestamp returns the VideoStartTimestamp field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetVideoStartTimestamp() int32 {
	if o == nil || IsNil(o.VideoStartTimestamp) {
		var ret int32
		return ret
	}
	return *o.VideoStartTimestamp
}

// GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetVideoStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoStartTimestamp) {
		return nil, false
	}
	return o.VideoStartTimestamp, true
}

// HasVideoStartTimestamp returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasVideoStartTimestamp() bool {
	if o != nil && !IsNil(o.VideoStartTimestamp) {
		return true
	}

	return false
}

// SetVideoStartTimestamp gets a reference to the given int32 and assigns it to the VideoStartTimestamp field.
func (o *CopyMessageRequest) SetVideoStartTimestamp(v int32) {
	o.VideoStartTimestamp = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *CopyMessageRequest) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *CopyMessageRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *CopyMessageRequest) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *CopyMessageRequest) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *CopyMessageRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *CopyMessageRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *CopyMessageRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *CopyMessageRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *CopyMessageRequest) GetReplyMarkup() SendMessageRequestReplyMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret SendMessageRequestReplyMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMessageRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *CopyMessageRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given SendMessageRequestReplyMarkup and assigns it to the ReplyMarkup field.
func (o *CopyMessageRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup) {
	o.ReplyMarkup = &v
}


func (o CopyMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["from_chat_id"] = o.FromChatId
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.VideoStartTimestamp) {
		toSerialize["video_start_timestamp"] = o.VideoStartTimestamp
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
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	if !IsNil(o.AllowPaidBroadcast) {
		toSerialize["allow_paid_broadcast"] = o.AllowPaidBroadcast
	}
	if !IsNil(o.ReplyParameters) {
		toSerialize["reply_parameters"] = o.ReplyParameters
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *CopyMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"from_chat_id",
		"message_id",
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

	varCopyMessageRequest := _CopyMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCopyMessageRequest)

	if err != nil {
		return err
	}

	*o = CopyMessageRequest(varCopyMessageRequest)

	return err
}

type NullableCopyMessageRequest struct {
	value *CopyMessageRequest
	isSet bool
}

func (v NullableCopyMessageRequest) Get() *CopyMessageRequest {
	return v.value
}

func (v *NullableCopyMessageRequest) Set(val *CopyMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyMessageRequest(val *CopyMessageRequest) *NullableCopyMessageRequest {
	return &NullableCopyMessageRequest{value: val, isSet: true}
}

func (v NullableCopyMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


