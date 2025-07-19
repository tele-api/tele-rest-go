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

// checks if the InlineQueryResultCachedMpeg4Gif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultCachedMpeg4Gif{}

// InlineQueryResultCachedMpeg4Gif Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be *mpeg4\\_gif*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the MPEG4 file
	Mpeg4FileId string `json:"mpeg4_file_id"`
	// *Optional*. Title for the result
	Title *string `json:"title,omitempty"`
	// *Optional*. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Pass *True*, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultCachedMpeg4Gif InlineQueryResultCachedMpeg4Gif

// NewInlineQueryResultCachedMpeg4Gif instantiates a new InlineQueryResultCachedMpeg4Gif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultCachedMpeg4Gif(type_ string, id string, mpeg4FileId string) *InlineQueryResultCachedMpeg4Gif {
	this := InlineQueryResultCachedMpeg4Gif{}
	this.Type = type_
	this.Id = id
	this.Mpeg4FileId = mpeg4FileId
	return &this
}

// NewInlineQueryResultCachedMpeg4GifWithDefaults instantiates a new InlineQueryResultCachedMpeg4Gif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultCachedMpeg4GifWithDefaults() *InlineQueryResultCachedMpeg4Gif {
	this := InlineQueryResultCachedMpeg4Gif{}
	var type_ string = "mpeg4_gif"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultCachedMpeg4Gif) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultCachedMpeg4Gif) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultCachedMpeg4Gif) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultCachedMpeg4Gif) SetId(v string) {
	o.Id = v
}

// GetMpeg4FileId returns the Mpeg4FileId field value
func (o *InlineQueryResultCachedMpeg4Gif) GetMpeg4FileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mpeg4FileId
}

// GetMpeg4FileIdOk returns a tuple with the Mpeg4FileId field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetMpeg4FileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mpeg4FileId, true
}

// SetMpeg4FileId sets field value
func (o *InlineQueryResultCachedMpeg4Gif) SetMpeg4FileId(v string) {
	o.Mpeg4FileId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineQueryResultCachedMpeg4Gif) SetTitle(v string) {
	o.Title = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultCachedMpeg4Gif) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultCachedMpeg4Gif) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultCachedMpeg4Gif) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *InlineQueryResultCachedMpeg4Gif) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultCachedMpeg4Gif) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultCachedMpeg4Gif) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedMpeg4Gif) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultCachedMpeg4Gif) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultCachedMpeg4Gif) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultCachedMpeg4Gif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultCachedMpeg4Gif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["mpeg4_file_id"] = o.Mpeg4FileId
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
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
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	return toSerialize, nil
}

func (o *InlineQueryResultCachedMpeg4Gif) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"mpeg4_file_id",
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

	varInlineQueryResultCachedMpeg4Gif := _InlineQueryResultCachedMpeg4Gif{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultCachedMpeg4Gif)

	if err != nil {
		return err
	}

	*o = InlineQueryResultCachedMpeg4Gif(varInlineQueryResultCachedMpeg4Gif)

	return err
}

type NullableInlineQueryResultCachedMpeg4Gif struct {
	value *InlineQueryResultCachedMpeg4Gif
	isSet bool
}

func (v NullableInlineQueryResultCachedMpeg4Gif) Get() *InlineQueryResultCachedMpeg4Gif {
	return v.value
}

func (v *NullableInlineQueryResultCachedMpeg4Gif) Set(val *InlineQueryResultCachedMpeg4Gif) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultCachedMpeg4Gif) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultCachedMpeg4Gif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultCachedMpeg4Gif(val *InlineQueryResultCachedMpeg4Gif) *NullableInlineQueryResultCachedMpeg4Gif {
	return &NullableInlineQueryResultCachedMpeg4Gif{value: val, isSet: true}
}

func (v NullableInlineQueryResultCachedMpeg4Gif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultCachedMpeg4Gif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


