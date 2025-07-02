/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the InlineQueryResultMpeg4Gif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultMpeg4Gif{}

// InlineQueryResultMpeg4Gif Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be *mpeg4\\_gif*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the MPEG4 file
	Mpeg4Url string `json:"mpeg4_url"`
	// *Optional*. Video width
	Mpeg4Width *int32 `json:"mpeg4_width,omitempty"`
	// *Optional*. Video height
	Mpeg4Height *int32 `json:"mpeg4_height,omitempty"`
	// *Optional*. Video duration in seconds
	Mpeg4Duration *int32 `json:"mpeg4_duration,omitempty"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`
	// *Optional*. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType *string `json:"thumbnail_mime_type,omitempty"`
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

type _InlineQueryResultMpeg4Gif InlineQueryResultMpeg4Gif

// NewInlineQueryResultMpeg4Gif instantiates a new InlineQueryResultMpeg4Gif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultMpeg4Gif(type_ string, id string, mpeg4Url string, thumbnailUrl string) *InlineQueryResultMpeg4Gif {
	this := InlineQueryResultMpeg4Gif{}
	this.Type = type_
	this.Id = id
	this.Mpeg4Url = mpeg4Url
	this.ThumbnailUrl = thumbnailUrl
	var thumbnailMimeType string = "image/jpeg"
	this.ThumbnailMimeType = &thumbnailMimeType
	return &this
}

// NewInlineQueryResultMpeg4GifWithDefaults instantiates a new InlineQueryResultMpeg4Gif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultMpeg4GifWithDefaults() *InlineQueryResultMpeg4Gif {
	this := InlineQueryResultMpeg4Gif{}
	var type_ string = "mpeg4_gif"
	this.Type = type_
	var thumbnailMimeType string = "image/jpeg"
	this.ThumbnailMimeType = &thumbnailMimeType
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultMpeg4Gif) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultMpeg4Gif) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultMpeg4Gif) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultMpeg4Gif) SetId(v string) {
	o.Id = v
}

// GetMpeg4Url returns the Mpeg4Url field value
func (o *InlineQueryResultMpeg4Gif) GetMpeg4Url() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mpeg4Url
}

// GetMpeg4UrlOk returns a tuple with the Mpeg4Url field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4UrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mpeg4Url, true
}

// SetMpeg4Url sets field value
func (o *InlineQueryResultMpeg4Gif) SetMpeg4Url(v string) {
	o.Mpeg4Url = v
}

// GetMpeg4Width returns the Mpeg4Width field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4Width() int32 {
	if o == nil || IsNil(o.Mpeg4Width) {
		var ret int32
		return ret
	}
	return *o.Mpeg4Width
}

// GetMpeg4WidthOk returns a tuple with the Mpeg4Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4WidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Mpeg4Width) {
		return nil, false
	}
	return o.Mpeg4Width, true
}

// HasMpeg4Width returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasMpeg4Width() bool {
	if o != nil && !IsNil(o.Mpeg4Width) {
		return true
	}

	return false
}

// SetMpeg4Width gets a reference to the given int32 and assigns it to the Mpeg4Width field.
func (o *InlineQueryResultMpeg4Gif) SetMpeg4Width(v int32) {
	o.Mpeg4Width = &v
}


// GetMpeg4Height returns the Mpeg4Height field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4Height() int32 {
	if o == nil || IsNil(o.Mpeg4Height) {
		var ret int32
		return ret
	}
	return *o.Mpeg4Height
}

// GetMpeg4HeightOk returns a tuple with the Mpeg4Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4HeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Mpeg4Height) {
		return nil, false
	}
	return o.Mpeg4Height, true
}

// HasMpeg4Height returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasMpeg4Height() bool {
	if o != nil && !IsNil(o.Mpeg4Height) {
		return true
	}

	return false
}

// SetMpeg4Height gets a reference to the given int32 and assigns it to the Mpeg4Height field.
func (o *InlineQueryResultMpeg4Gif) SetMpeg4Height(v int32) {
	o.Mpeg4Height = &v
}


// GetMpeg4Duration returns the Mpeg4Duration field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4Duration() int32 {
	if o == nil || IsNil(o.Mpeg4Duration) {
		var ret int32
		return ret
	}
	return *o.Mpeg4Duration
}

// GetMpeg4DurationOk returns a tuple with the Mpeg4Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetMpeg4DurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Mpeg4Duration) {
		return nil, false
	}
	return o.Mpeg4Duration, true
}

// HasMpeg4Duration returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasMpeg4Duration() bool {
	if o != nil && !IsNil(o.Mpeg4Duration) {
		return true
	}

	return false
}

// SetMpeg4Duration gets a reference to the given int32 and assigns it to the Mpeg4Duration field.
func (o *InlineQueryResultMpeg4Gif) SetMpeg4Duration(v int32) {
	o.Mpeg4Duration = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *InlineQueryResultMpeg4Gif) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *InlineQueryResultMpeg4Gif) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetThumbnailMimeType returns the ThumbnailMimeType field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetThumbnailMimeType() string {
	if o == nil || IsNil(o.ThumbnailMimeType) {
		var ret string
		return ret
	}
	return *o.ThumbnailMimeType
}

// GetThumbnailMimeTypeOk returns a tuple with the ThumbnailMimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetThumbnailMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailMimeType) {
		return nil, false
	}
	return o.ThumbnailMimeType, true
}

// HasThumbnailMimeType returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasThumbnailMimeType() bool {
	if o != nil && !IsNil(o.ThumbnailMimeType) {
		return true
	}

	return false
}

// SetThumbnailMimeType gets a reference to the given string and assigns it to the ThumbnailMimeType field.
func (o *InlineQueryResultMpeg4Gif) SetThumbnailMimeType(v string) {
	o.ThumbnailMimeType = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineQueryResultMpeg4Gif) SetTitle(v string) {
	o.Title = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultMpeg4Gif) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultMpeg4Gif) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultMpeg4Gif) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *InlineQueryResultMpeg4Gif) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultMpeg4Gif) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultMpeg4Gif) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultMpeg4Gif) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultMpeg4Gif) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultMpeg4Gif) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultMpeg4Gif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultMpeg4Gif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["mpeg4_url"] = o.Mpeg4Url
	if !IsNil(o.Mpeg4Width) {
		toSerialize["mpeg4_width"] = o.Mpeg4Width
	}
	if !IsNil(o.Mpeg4Height) {
		toSerialize["mpeg4_height"] = o.Mpeg4Height
	}
	if !IsNil(o.Mpeg4Duration) {
		toSerialize["mpeg4_duration"] = o.Mpeg4Duration
	}
	toSerialize["thumbnail_url"] = o.ThumbnailUrl
	if !IsNil(o.ThumbnailMimeType) {
		toSerialize["thumbnail_mime_type"] = o.ThumbnailMimeType
	}
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

func (o *InlineQueryResultMpeg4Gif) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"mpeg4_url",
		"thumbnail_url",
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

	varInlineQueryResultMpeg4Gif := _InlineQueryResultMpeg4Gif{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultMpeg4Gif)

	if err != nil {
		return err
	}

	*o = InlineQueryResultMpeg4Gif(varInlineQueryResultMpeg4Gif)

	return err
}

type NullableInlineQueryResultMpeg4Gif struct {
	value *InlineQueryResultMpeg4Gif
	isSet bool
}

func (v NullableInlineQueryResultMpeg4Gif) Get() *InlineQueryResultMpeg4Gif {
	return v.value
}

func (v *NullableInlineQueryResultMpeg4Gif) Set(val *InlineQueryResultMpeg4Gif) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultMpeg4Gif) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultMpeg4Gif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultMpeg4Gif(val *InlineQueryResultMpeg4Gif) *NullableInlineQueryResultMpeg4Gif {
	return &NullableInlineQueryResultMpeg4Gif{value: val, isSet: true}
}

func (v NullableInlineQueryResultMpeg4Gif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultMpeg4Gif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


