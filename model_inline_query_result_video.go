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

// checks if the InlineQueryResultVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultVideo{}

// InlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the video.  If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you **must** replace its content using *input\\_message\\_content*.
type InlineQueryResultVideo struct {
	// Type of the result, must be *video*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the embedded video player or video file
	VideoUrl string `json:"video_url"`
	// MIME type of the content of the video URL, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`
	// URL of the thumbnail (JPEG only) for the video
	ThumbnailUrl string `json:"thumbnail_url"`
	// Title for the result
	Title string `json:"title"`
	// *Optional*. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Pass *True*, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// *Optional*. Video width
	VideoWidth *int32 `json:"video_width,omitempty"`
	// *Optional*. Video height
	VideoHeight *int32 `json:"video_height,omitempty"`
	// *Optional*. Video duration in seconds
	VideoDuration *int32 `json:"video_duration,omitempty"`
	// *Optional*. Short description of the result
	Description *string `json:"description,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultVideo InlineQueryResultVideo

// NewInlineQueryResultVideo instantiates a new InlineQueryResultVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultVideo(type_ string, id string, videoUrl string, mimeType string, thumbnailUrl string, title string) *InlineQueryResultVideo {
	this := InlineQueryResultVideo{}
	this.Type = type_
	this.Id = id
	this.VideoUrl = videoUrl
	this.MimeType = mimeType
	this.ThumbnailUrl = thumbnailUrl
	this.Title = title
	return &this
}

// NewInlineQueryResultVideoWithDefaults instantiates a new InlineQueryResultVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultVideoWithDefaults() *InlineQueryResultVideo {
	this := InlineQueryResultVideo{}
	var type_ string = "video"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultVideo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultVideo) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultVideo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultVideo) SetId(v string) {
	o.Id = v
}

// GetVideoUrl returns the VideoUrl field value
func (o *InlineQueryResultVideo) GetVideoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VideoUrl, true
}

// SetVideoUrl sets field value
func (o *InlineQueryResultVideo) SetVideoUrl(v string) {
	o.VideoUrl = v
}

// GetMimeType returns the MimeType field value
func (o *InlineQueryResultVideo) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *InlineQueryResultVideo) SetMimeType(v string) {
	o.MimeType = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *InlineQueryResultVideo) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *InlineQueryResultVideo) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultVideo) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultVideo) SetTitle(v string) {
	o.Title = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultVideo) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultVideo) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultVideo) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *InlineQueryResultVideo) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetVideoWidth returns the VideoWidth field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetVideoWidth() int32 {
	if o == nil || IsNil(o.VideoWidth) {
		var ret int32
		return ret
	}
	return *o.VideoWidth
}

// GetVideoWidthOk returns a tuple with the VideoWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetVideoWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoWidth) {
		return nil, false
	}
	return o.VideoWidth, true
}

// HasVideoWidth returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasVideoWidth() bool {
	if o != nil && !IsNil(o.VideoWidth) {
		return true
	}

	return false
}

// SetVideoWidth gets a reference to the given int32 and assigns it to the VideoWidth field.
func (o *InlineQueryResultVideo) SetVideoWidth(v int32) {
	o.VideoWidth = &v
}


// GetVideoHeight returns the VideoHeight field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetVideoHeight() int32 {
	if o == nil || IsNil(o.VideoHeight) {
		var ret int32
		return ret
	}
	return *o.VideoHeight
}

// GetVideoHeightOk returns a tuple with the VideoHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetVideoHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoHeight) {
		return nil, false
	}
	return o.VideoHeight, true
}

// HasVideoHeight returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasVideoHeight() bool {
	if o != nil && !IsNil(o.VideoHeight) {
		return true
	}

	return false
}

// SetVideoHeight gets a reference to the given int32 and assigns it to the VideoHeight field.
func (o *InlineQueryResultVideo) SetVideoHeight(v int32) {
	o.VideoHeight = &v
}


// GetVideoDuration returns the VideoDuration field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetVideoDuration() int32 {
	if o == nil || IsNil(o.VideoDuration) {
		var ret int32
		return ret
	}
	return *o.VideoDuration
}

// GetVideoDurationOk returns a tuple with the VideoDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetVideoDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoDuration) {
		return nil, false
	}
	return o.VideoDuration, true
}

// HasVideoDuration returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasVideoDuration() bool {
	if o != nil && !IsNil(o.VideoDuration) {
		return true
	}

	return false
}

// SetVideoDuration gets a reference to the given int32 and assigns it to the VideoDuration field.
func (o *InlineQueryResultVideo) SetVideoDuration(v int32) {
	o.VideoDuration = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineQueryResultVideo) SetDescription(v string) {
	o.Description = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultVideo) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultVideo) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVideo) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultVideo) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultVideo) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["video_url"] = o.VideoUrl
	toSerialize["mime_type"] = o.MimeType
	toSerialize["thumbnail_url"] = o.ThumbnailUrl
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
	if !IsNil(o.ShowCaptionAboveMedia) {
		toSerialize["show_caption_above_media"] = o.ShowCaptionAboveMedia
	}
	if !IsNil(o.VideoWidth) {
		toSerialize["video_width"] = o.VideoWidth
	}
	if !IsNil(o.VideoHeight) {
		toSerialize["video_height"] = o.VideoHeight
	}
	if !IsNil(o.VideoDuration) {
		toSerialize["video_duration"] = o.VideoDuration
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	return toSerialize, nil
}

func (o *InlineQueryResultVideo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"video_url",
		"mime_type",
		"thumbnail_url",
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

	varInlineQueryResultVideo := _InlineQueryResultVideo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultVideo)

	if err != nil {
		return err
	}

	*o = InlineQueryResultVideo(varInlineQueryResultVideo)

	return err
}

type NullableInlineQueryResultVideo struct {
	value *InlineQueryResultVideo
	isSet bool
}

func (v NullableInlineQueryResultVideo) Get() *InlineQueryResultVideo {
	return v.value
}

func (v *NullableInlineQueryResultVideo) Set(val *InlineQueryResultVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultVideo(val *InlineQueryResultVideo) *NullableInlineQueryResultVideo {
	return &NullableInlineQueryResultVideo{value: val, isSet: true}
}

func (v NullableInlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


