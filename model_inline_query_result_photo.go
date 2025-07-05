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

// checks if the InlineQueryResultPhoto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultPhoto{}

// InlineQueryResultPhoto Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type of the result, must be *photo*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL of the photo. Photo must be in **JPEG** format. Photo size must not exceed 5MB
	PhotoUrl string `json:"photo_url"`
	// URL of the thumbnail for the photo
	ThumbnailUrl string `json:"thumbnail_url"`
	// *Optional*. Width of the photo
	PhotoWidth *int32 `json:"photo_width,omitempty"`
	// *Optional*. Height of the photo
	PhotoHeight *int32 `json:"photo_height,omitempty"`
	// *Optional*. Title for the result
	Title *string `json:"title,omitempty"`
	// *Optional*. Short description of the result
	Description *string `json:"description,omitempty"`
	// *Optional*. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the photo caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Pass *True*, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultPhoto InlineQueryResultPhoto

// NewInlineQueryResultPhoto instantiates a new InlineQueryResultPhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultPhoto(type_ string, id string, photoUrl string, thumbnailUrl string) *InlineQueryResultPhoto {
	this := InlineQueryResultPhoto{}
	this.Type = type_
	this.Id = id
	this.PhotoUrl = photoUrl
	this.ThumbnailUrl = thumbnailUrl
	return &this
}

// NewInlineQueryResultPhotoWithDefaults instantiates a new InlineQueryResultPhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultPhotoWithDefaults() *InlineQueryResultPhoto {
	this := InlineQueryResultPhoto{}
	var type_ string = "photo"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultPhoto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultPhoto) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultPhoto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultPhoto) SetId(v string) {
	o.Id = v
}

// GetPhotoUrl returns the PhotoUrl field value
func (o *InlineQueryResultPhoto) GetPhotoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetPhotoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhotoUrl, true
}

// SetPhotoUrl sets field value
func (o *InlineQueryResultPhoto) SetPhotoUrl(v string) {
	o.PhotoUrl = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *InlineQueryResultPhoto) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *InlineQueryResultPhoto) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetPhotoWidth returns the PhotoWidth field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetPhotoWidth() int32 {
	if o == nil || IsNil(o.PhotoWidth) {
		var ret int32
		return ret
	}
	return *o.PhotoWidth
}

// GetPhotoWidthOk returns a tuple with the PhotoWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetPhotoWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoWidth) {
		return nil, false
	}
	return o.PhotoWidth, true
}

// HasPhotoWidth returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasPhotoWidth() bool {
	if o != nil && !IsNil(o.PhotoWidth) {
		return true
	}

	return false
}

// SetPhotoWidth gets a reference to the given int32 and assigns it to the PhotoWidth field.
func (o *InlineQueryResultPhoto) SetPhotoWidth(v int32) {
	o.PhotoWidth = &v
}


// GetPhotoHeight returns the PhotoHeight field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetPhotoHeight() int32 {
	if o == nil || IsNil(o.PhotoHeight) {
		var ret int32
		return ret
	}
	return *o.PhotoHeight
}

// GetPhotoHeightOk returns a tuple with the PhotoHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetPhotoHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoHeight) {
		return nil, false
	}
	return o.PhotoHeight, true
}

// HasPhotoHeight returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasPhotoHeight() bool {
	if o != nil && !IsNil(o.PhotoHeight) {
		return true
	}

	return false
}

// SetPhotoHeight gets a reference to the given int32 and assigns it to the PhotoHeight field.
func (o *InlineQueryResultPhoto) SetPhotoHeight(v int32) {
	o.PhotoHeight = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineQueryResultPhoto) SetTitle(v string) {
	o.Title = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineQueryResultPhoto) SetDescription(v string) {
	o.Description = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultPhoto) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultPhoto) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultPhoto) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetShowCaptionAboveMedia() bool {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		var ret bool
		return ret
	}
	return *o.ShowCaptionAboveMedia
}

// GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetShowCaptionAboveMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCaptionAboveMedia) {
		return nil, false
	}
	return o.ShowCaptionAboveMedia, true
}

// HasShowCaptionAboveMedia returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasShowCaptionAboveMedia() bool {
	if o != nil && !IsNil(o.ShowCaptionAboveMedia) {
		return true
	}

	return false
}

// SetShowCaptionAboveMedia gets a reference to the given bool and assigns it to the ShowCaptionAboveMedia field.
func (o *InlineQueryResultPhoto) SetShowCaptionAboveMedia(v bool) {
	o.ShowCaptionAboveMedia = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultPhoto) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultPhoto) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultPhoto) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultPhoto) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultPhoto) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultPhoto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["photo_url"] = o.PhotoUrl
	toSerialize["thumbnail_url"] = o.ThumbnailUrl
	if !IsNil(o.PhotoWidth) {
		toSerialize["photo_width"] = o.PhotoWidth
	}
	if !IsNil(o.PhotoHeight) {
		toSerialize["photo_height"] = o.PhotoHeight
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *InlineQueryResultPhoto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"photo_url",
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

	varInlineQueryResultPhoto := _InlineQueryResultPhoto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultPhoto)

	if err != nil {
		return err
	}

	*o = InlineQueryResultPhoto(varInlineQueryResultPhoto)

	return err
}

type NullableInlineQueryResultPhoto struct {
	value *InlineQueryResultPhoto
	isSet bool
}

func (v NullableInlineQueryResultPhoto) Get() *InlineQueryResultPhoto {
	return v.value
}

func (v *NullableInlineQueryResultPhoto) Set(val *InlineQueryResultPhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultPhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultPhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultPhoto(val *InlineQueryResultPhoto) *NullableInlineQueryResultPhoto {
	return &NullableInlineQueryResultPhoto{value: val, isSet: true}
}

func (v NullableInlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultPhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


