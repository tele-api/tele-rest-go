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

// checks if the InlineQueryResultAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultAudio{}

// InlineQueryResultAudio Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	// Type of the result, must be *audio*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the audio file
	AudioUrl string `json:"audio_url"`
	// Title
	Title string `json:"title"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Performer
	Performer *string `json:"performer,omitempty"`
	// *Optional*. Audio duration in seconds
	AudioDuration *int32 `json:"audio_duration,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultAudio InlineQueryResultAudio

// NewInlineQueryResultAudio instantiates a new InlineQueryResultAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultAudio(type_ string, id string, audioUrl string, title string) *InlineQueryResultAudio {
	this := InlineQueryResultAudio{}
	this.Type = type_
	this.Id = id
	this.AudioUrl = audioUrl
	this.Title = title
	return &this
}

// NewInlineQueryResultAudioWithDefaults instantiates a new InlineQueryResultAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultAudioWithDefaults() *InlineQueryResultAudio {
	this := InlineQueryResultAudio{}
	var type_ string = "audio"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultAudio) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultAudio) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultAudio) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultAudio) SetId(v string) {
	o.Id = v
}

// GetAudioUrl returns the AudioUrl field value
func (o *InlineQueryResultAudio) GetAudioUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudioUrl
}

// GetAudioUrlOk returns a tuple with the AudioUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetAudioUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioUrl, true
}

// SetAudioUrl sets field value
func (o *InlineQueryResultAudio) SetAudioUrl(v string) {
	o.AudioUrl = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultAudio) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultAudio) SetTitle(v string) {
	o.Title = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultAudio) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultAudio) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultAudio) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetPerformer() string {
	if o == nil || IsNil(o.Performer) {
		var ret string
		return ret
	}
	return *o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetPerformerOk() (*string, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given string and assigns it to the Performer field.
func (o *InlineQueryResultAudio) SetPerformer(v string) {
	o.Performer = &v
}


// GetAudioDuration returns the AudioDuration field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetAudioDuration() int32 {
	if o == nil || IsNil(o.AudioDuration) {
		var ret int32
		return ret
	}
	return *o.AudioDuration
}

// GetAudioDurationOk returns a tuple with the AudioDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetAudioDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.AudioDuration) {
		return nil, false
	}
	return o.AudioDuration, true
}

// HasAudioDuration returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasAudioDuration() bool {
	if o != nil && !IsNil(o.AudioDuration) {
		return true
	}

	return false
}

// SetAudioDuration gets a reference to the given int32 and assigns it to the AudioDuration field.
func (o *InlineQueryResultAudio) SetAudioDuration(v int32) {
	o.AudioDuration = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultAudio) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultAudio) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultAudio) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultAudio) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultAudio) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["audio_url"] = o.AudioUrl
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
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
	}
	if !IsNil(o.AudioDuration) {
		toSerialize["audio_duration"] = o.AudioDuration
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	return toSerialize, nil
}

func (o *InlineQueryResultAudio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"audio_url",
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

	varInlineQueryResultAudio := _InlineQueryResultAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultAudio)

	if err != nil {
		return err
	}

	*o = InlineQueryResultAudio(varInlineQueryResultAudio)

	return err
}

type NullableInlineQueryResultAudio struct {
	value *InlineQueryResultAudio
	isSet bool
}

func (v NullableInlineQueryResultAudio) Get() *InlineQueryResultAudio {
	return v.value
}

func (v *NullableInlineQueryResultAudio) Set(val *InlineQueryResultAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultAudio(val *InlineQueryResultAudio) *NullableInlineQueryResultAudio {
	return &NullableInlineQueryResultAudio{value: val, isSet: true}
}

func (v NullableInlineQueryResultAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


