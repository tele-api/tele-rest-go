/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the InlineQueryResultGif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultGif{}

// InlineQueryResultGif Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	// Type of the result, must be *gif*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the GIF file
	GifUrl string `json:"gif_url"`
	// *Optional*. Width of the GIF
	GifWidth *int32 `json:"gif_width,omitempty"`
	// *Optional*. Height of the GIF
	GifHeight *int32 `json:"gif_height,omitempty"`
	// *Optional*. Duration of the GIF in seconds
	GifDuration *int32 `json:"gif_duration,omitempty"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`
	// *Optional*. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType *string `json:"thumbnail_mime_type,omitempty"`
	// *Optional*. Title for the result
	Title *string `json:"title,omitempty"`
	// *Optional*. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
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

type _InlineQueryResultGif InlineQueryResultGif

// NewInlineQueryResultGif instantiates a new InlineQueryResultGif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultGif(type_ string, id string, gifUrl string, thumbnailUrl string) *InlineQueryResultGif {
	this := InlineQueryResultGif{}
	this.Type = type_
	this.Id = id
	this.GifUrl = gifUrl
	this.ThumbnailUrl = thumbnailUrl
	var thumbnailMimeType string = "image/jpeg"
	this.ThumbnailMimeType = &thumbnailMimeType
	return &this
}

// NewInlineQueryResultGifWithDefaults instantiates a new InlineQueryResultGif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultGifWithDefaults() *InlineQueryResultGif {
	this := InlineQueryResultGif{}
	var type_ string = "gif"
	this.Type = type_
	var thumbnailMimeType string = "image/jpeg"
	this.ThumbnailMimeType = &thumbnailMimeType
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultGif) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultGif) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultGif) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultGif) SetId(v string) {
	o.Id = v
}

// GetGifUrl returns the GifUrl field value
func (o *InlineQueryResultGif) GetGifUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GifUrl
}

// GetGifUrlOk returns a tuple with the GifUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetGifUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GifUrl, true
}

// SetGifUrl sets field value
func (o *InlineQueryResultGif) SetGifUrl(v string) {
	o.GifUrl = v
}

// GetGifWidth returns the GifWidth field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetGifWidth() int32 {
	if o == nil || IsNil(o.GifWidth) {
		var ret int32
		return ret
	}
	return *o.GifWidth
}

// GetGifWidthOk returns a tuple with the GifWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetGifWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.GifWidth) {
		return nil, false
	}
	return o.GifWidth, true
}

// HasGifWidth returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasGifWidth() bool {
	if o != nil && !IsNil(o.GifWidth) {
		return true
	}

	return false
}

// SetGifWidth gets a reference to the given int32 and assigns it to the GifWidth field.
func (o *InlineQueryResultGif) SetGifWidth(v int32) {
	o.GifWidth = &v
}


// GetGifHeight returns the GifHeight field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetGifHeight() int32 {
	if o == nil || IsNil(o.GifHeight) {
		var ret int32
		return ret
	}
	return *o.GifHeight
}

// GetGifHeightOk returns a tuple with the GifHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetGifHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.GifHeight) {
		return nil, false
	}
	return o.GifHeight, true
}

// HasGifHeight returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasGifHeight() bool {
	if o != nil && !IsNil(o.GifHeight) {
		return true
	}

	return false
}

// SetGifHeight gets a reference to the given int32 and assigns it to the GifHeight field.
func (o *InlineQueryResultGif) SetGifHeight(v int32) {
	o.GifHeight = &v
}


// GetGifDuration returns the GifDuration field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetGifDuration() int32 {
	if o == nil || IsNil(o.GifDuration) {
		var ret int32
		return ret
	}
	return *o.GifDuration
}

// GetGifDurationOk returns a tuple with the GifDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetGifDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.GifDuration) {
		return nil, false
	}
	return o.GifDuration, true
}

// HasGifDuration returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasGifDuration() bool {
	if o != nil && !IsNil(o.GifDuration) {
		return true
	}

	return false
}

// SetGifDuration gets a reference to the given int32 and assigns it to the GifDuration field.
func (o *InlineQueryResultGif) SetGifDuration(v int32) {
	o.GifDuration = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *InlineQueryResultGif) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *InlineQueryResultGif) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetThumbnailMimeType returns the ThumbnailMimeType field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetThumbnailMimeType() string {
	if o == nil || IsNil(o.ThumbnailMimeType) {
		var ret string
		return ret
	}
	return *o.ThumbnailMimeType
}

// GetThumbnailMimeTypeOk returns a tuple with the ThumbnailMimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetThumbnailMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailMimeType) {
		return nil, false
	}
	return o.ThumbnailMimeType, true
}

// HasThumbnailMimeType returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasThumbnailMimeType() bool {
	if o != nil && !IsNil(o.ThumbnailMimeType) {
		return true
	}

	return false
}

// SetThumbnailMimeType gets a reference to the given string and assigns it to the ThumbnailMimeType field.
func (o *InlineQueryResultGif) SetThumbnailMimeType(v string) {
	o.ThumbnailMimeType = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineQueryResultGif) SetTitle(v string) {
	o.Title = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultGif) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultGif) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultGif) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *InlineQueryResultGif) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultGif) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultGif) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultGif) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultGif) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultGif) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultGif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultGif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["gif_url"] = o.GifUrl
	if !IsNil(o.GifWidth) {
		toSerialize["gif_width"] = o.GifWidth
	}
	if !IsNil(o.GifHeight) {
		toSerialize["gif_height"] = o.GifHeight
	}
	if !IsNil(o.GifDuration) {
		toSerialize["gif_duration"] = o.GifDuration
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

func (o *InlineQueryResultGif) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"gif_url",
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

	varInlineQueryResultGif := _InlineQueryResultGif{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultGif)

	if err != nil {
		return err
	}

	*o = InlineQueryResultGif(varInlineQueryResultGif)

	return err
}

type NullableInlineQueryResultGif struct {
	value *InlineQueryResultGif
	isSet bool
}

func (v NullableInlineQueryResultGif) Get() *InlineQueryResultGif {
	return v.value
}

func (v *NullableInlineQueryResultGif) Set(val *InlineQueryResultGif) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultGif) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultGif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultGif(val *InlineQueryResultGif) *NullableInlineQueryResultGif {
	return &NullableInlineQueryResultGif{value: val, isSet: true}
}

func (v NullableInlineQueryResultGif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultGif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


