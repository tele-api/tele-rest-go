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
)

// checks if the EditMessageCaptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditMessageCaptionRequest{}

// EditMessageCaptionRequest Request parameters for editMessageCaption
type EditMessageCaptionRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId *EditMessageTextRequestChatId `json:"chat_id,omitempty"`
	// Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit
	MessageId *int32 `json:"message_id,omitempty"`
	// Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// New caption of the message, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the message caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Pass *True*, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// NewEditMessageCaptionRequest instantiates a new EditMessageCaptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMessageCaptionRequest() *EditMessageCaptionRequest {
	this := EditMessageCaptionRequest{}
	return &this
}

// NewEditMessageCaptionRequestWithDefaults instantiates a new EditMessageCaptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMessageCaptionRequestWithDefaults() *EditMessageCaptionRequest {
	this := EditMessageCaptionRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *EditMessageCaptionRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetChatId() EditMessageTextRequestChatId {
	if o == nil || IsNil(o.ChatId) {
		var ret EditMessageTextRequestChatId
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetChatIdOk() (*EditMessageTextRequestChatId, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given EditMessageTextRequestChatId and assigns it to the ChatId field.
func (o *EditMessageCaptionRequest) SetChatId(v EditMessageTextRequestChatId) {
	o.ChatId = &v
}


// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *EditMessageCaptionRequest) SetMessageId(v int32) {
	o.MessageId = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *EditMessageCaptionRequest) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *EditMessageCaptionRequest) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *EditMessageCaptionRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *EditMessageCaptionRequest) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *EditMessageCaptionRequest) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *EditMessageCaptionRequest) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageCaptionRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *EditMessageCaptionRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *EditMessageCaptionRequest) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o EditMessageCaptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditMessageCaptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.InlineMessageId) {
		toSerialize["inline_message_id"] = o.InlineMessageId
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
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

type NullableEditMessageCaptionRequest struct {
	value *EditMessageCaptionRequest
	isSet bool
}

func (v NullableEditMessageCaptionRequest) Get() *EditMessageCaptionRequest {
	return v.value
}

func (v *NullableEditMessageCaptionRequest) Set(val *EditMessageCaptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageCaptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageCaptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageCaptionRequest(val *EditMessageCaptionRequest) *NullableEditMessageCaptionRequest {
	return &NullableEditMessageCaptionRequest{value: val, isSet: true}
}

func (v NullableEditMessageCaptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageCaptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


