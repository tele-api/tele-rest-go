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

// checks if the InlineQueryResultVoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultVoice{}

// InlineQueryResultVoice Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	// Type of the result, must be *voice*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the voice recording
	VoiceUrl string `json:"voice_url"`
	// Recording title
	Title string `json:"title"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the voice message caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Recording duration in seconds
	VoiceDuration *int32 `json:"voice_duration,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultVoice InlineQueryResultVoice

// NewInlineQueryResultVoice instantiates a new InlineQueryResultVoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultVoice(type_ string, id string, voiceUrl string, title string) *InlineQueryResultVoice {
	this := InlineQueryResultVoice{}
	this.Type = type_
	this.Id = id
	this.VoiceUrl = voiceUrl
	this.Title = title
	return &this
}

// NewInlineQueryResultVoiceWithDefaults instantiates a new InlineQueryResultVoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultVoiceWithDefaults() *InlineQueryResultVoice {
	this := InlineQueryResultVoice{}
	var type_ string = "voice"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultVoice) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultVoice) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultVoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultVoice) SetId(v string) {
	o.Id = v
}

// GetVoiceUrl returns the VoiceUrl field value
func (o *InlineQueryResultVoice) GetVoiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VoiceUrl
}

// GetVoiceUrlOk returns a tuple with the VoiceUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetVoiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoiceUrl, true
}

// SetVoiceUrl sets field value
func (o *InlineQueryResultVoice) SetVoiceUrl(v string) {
	o.VoiceUrl = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultVoice) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultVoice) SetTitle(v string) {
	o.Title = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultVoice) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultVoice) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultVoice) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetVoiceDuration returns the VoiceDuration field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetVoiceDuration() int32 {
	if o == nil || IsNil(o.VoiceDuration) {
		var ret int32
		return ret
	}
	return *o.VoiceDuration
}

// GetVoiceDurationOk returns a tuple with the VoiceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetVoiceDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceDuration) {
		return nil, false
	}
	return o.VoiceDuration, true
}

// HasVoiceDuration returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasVoiceDuration() bool {
	if o != nil && !IsNil(o.VoiceDuration) {
		return true
	}

	return false
}

// SetVoiceDuration gets a reference to the given int32 and assigns it to the VoiceDuration field.
func (o *InlineQueryResultVoice) SetVoiceDuration(v int32) {
	o.VoiceDuration = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultVoice) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultVoice) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVoice) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultVoice) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultVoice) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultVoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultVoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["voice_url"] = o.VoiceUrl
	toSerialize["title"] = o.Title
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.CaptionEntities) {
		toSerialize["caption_entities"] = o.CaptionEntities
	}
	if !IsNil(o.VoiceDuration) {
		toSerialize["voice_duration"] = o.VoiceDuration
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	return toSerialize, nil
}

func (o *InlineQueryResultVoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"voice_url",
		"title",
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

	varInlineQueryResultVoice := _InlineQueryResultVoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultVoice)

	if err != nil {
		return err
	}

	*o = InlineQueryResultVoice(varInlineQueryResultVoice)

	return err
}

type NullableInlineQueryResultVoice struct {
	value *InlineQueryResultVoice
	isSet bool
}

func (v NullableInlineQueryResultVoice) Get() *InlineQueryResultVoice {
	return v.value
}

func (v *NullableInlineQueryResultVoice) Set(val *InlineQueryResultVoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultVoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultVoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultVoice(val *InlineQueryResultVoice) *NullableInlineQueryResultVoice {
	return &NullableInlineQueryResultVoice{value: val, isSet: true}
}

func (v NullableInlineQueryResultVoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultVoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


