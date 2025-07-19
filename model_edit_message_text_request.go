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

// checks if the EditMessageTextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditMessageTextRequest{}

// EditMessageTextRequest Request parameters for editMessageText
type EditMessageTextRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId *EditMessageTextRequestChatId `json:"chat_id,omitempty"`
	// Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit
	MessageId *int32 `json:"message_id,omitempty"`
	// Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`
	// Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of *parse\\_mode*
	Entities []MessageEntity `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type _EditMessageTextRequest EditMessageTextRequest

// NewEditMessageTextRequest instantiates a new EditMessageTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMessageTextRequest(text string) *EditMessageTextRequest {
	this := EditMessageTextRequest{}
	this.Text = text
	return &this
}

// NewEditMessageTextRequestWithDefaults instantiates a new EditMessageTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMessageTextRequestWithDefaults() *EditMessageTextRequest {
	this := EditMessageTextRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *EditMessageTextRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetChatId() EditMessageTextRequestChatId {
	if o == nil || IsNil(o.ChatId) {
		var ret EditMessageTextRequestChatId
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetChatIdOk() (*EditMessageTextRequestChatId, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given EditMessageTextRequestChatId and assigns it to the ChatId field.
func (o *EditMessageTextRequest) SetChatId(v EditMessageTextRequestChatId) {
	o.ChatId = &v
}


// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *EditMessageTextRequest) SetMessageId(v int32) {
	o.MessageId = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *EditMessageTextRequest) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


// GetText returns the Text field value
func (o *EditMessageTextRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *EditMessageTextRequest) SetText(v string) {
	o.Text = v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *EditMessageTextRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *EditMessageTextRequest) SetEntities(v []MessageEntity) {
	o.Entities = v
}


// GetLinkPreviewOptions returns the LinkPreviewOptions field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetLinkPreviewOptions() LinkPreviewOptions {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		var ret LinkPreviewOptions
		return ret
	}
	return *o.LinkPreviewOptions
}

// GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool) {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		return nil, false
	}
	return o.LinkPreviewOptions, true
}

// HasLinkPreviewOptions returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasLinkPreviewOptions() bool {
	if o != nil && !IsNil(o.LinkPreviewOptions) {
		return true
	}

	return false
}

// SetLinkPreviewOptions gets a reference to the given LinkPreviewOptions and assigns it to the LinkPreviewOptions field.
func (o *EditMessageTextRequest) SetLinkPreviewOptions(v LinkPreviewOptions) {
	o.LinkPreviewOptions = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *EditMessageTextRequest) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageTextRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *EditMessageTextRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *EditMessageTextRequest) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o EditMessageTextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditMessageTextRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["text"] = o.Text
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.LinkPreviewOptions) {
		toSerialize["link_preview_options"] = o.LinkPreviewOptions
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *EditMessageTextRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varEditMessageTextRequest := _EditMessageTextRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditMessageTextRequest)

	if err != nil {
		return err
	}

	*o = EditMessageTextRequest(varEditMessageTextRequest)

	return err
}

type NullableEditMessageTextRequest struct {
	value *EditMessageTextRequest
	isSet bool
}

func (v NullableEditMessageTextRequest) Get() *EditMessageTextRequest {
	return v.value
}

func (v *NullableEditMessageTextRequest) Set(val *EditMessageTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageTextRequest(val *EditMessageTextRequest) *NullableEditMessageTextRequest {
	return &NullableEditMessageTextRequest{value: val, isSet: true}
}

func (v NullableEditMessageTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


