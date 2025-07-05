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

// checks if the InlineQueryResultDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultDocument{}

// InlineQueryResultDocument Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the file. Currently, only **.PDF** and **.ZIP** files can be sent using this method.
type InlineQueryResultDocument struct {
	// Type of the result, must be *document*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// *Optional*. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// A valid URL for the file
	DocumentUrl string `json:"document_url"`
	// MIME type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`
	// *Optional*. Short description of the result
	Description *string `json:"description,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// *Optional*. URL of the thumbnail (JPEG only) for the file
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`
}

type _InlineQueryResultDocument InlineQueryResultDocument

// NewInlineQueryResultDocument instantiates a new InlineQueryResultDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultDocument(type_ string, id string, title string, documentUrl string, mimeType string) *InlineQueryResultDocument {
	this := InlineQueryResultDocument{}
	this.Type = type_
	this.Id = id
	this.Title = title
	this.DocumentUrl = documentUrl
	this.MimeType = mimeType
	return &this
}

// NewInlineQueryResultDocumentWithDefaults instantiates a new InlineQueryResultDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultDocumentWithDefaults() *InlineQueryResultDocument {
	this := InlineQueryResultDocument{}
	var type_ string = "document"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultDocument) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultDocument) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultDocument) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultDocument) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultDocument) SetTitle(v string) {
	o.Title = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InlineQueryResultDocument) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InlineQueryResultDocument) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InlineQueryResultDocument) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetDocumentUrl returns the DocumentUrl field value
func (o *InlineQueryResultDocument) GetDocumentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentUrl, true
}

// SetDocumentUrl sets field value
func (o *InlineQueryResultDocument) SetDocumentUrl(v string) {
	o.DocumentUrl = v
}

// GetMimeType returns the MimeType field value
func (o *InlineQueryResultDocument) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *InlineQueryResultDocument) SetMimeType(v string) {
	o.MimeType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineQueryResultDocument) SetDescription(v string) {
	o.Description = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultDocument) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultDocument) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *InlineQueryResultDocument) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}


// GetThumbnailWidth returns the ThumbnailWidth field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetThumbnailWidth() int32 {
	if o == nil || IsNil(o.ThumbnailWidth) {
		var ret int32
		return ret
	}
	return *o.ThumbnailWidth
}

// GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetThumbnailWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailWidth) {
		return nil, false
	}
	return o.ThumbnailWidth, true
}

// HasThumbnailWidth returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasThumbnailWidth() bool {
	if o != nil && !IsNil(o.ThumbnailWidth) {
		return true
	}

	return false
}

// SetThumbnailWidth gets a reference to the given int32 and assigns it to the ThumbnailWidth field.
func (o *InlineQueryResultDocument) SetThumbnailWidth(v int32) {
	o.ThumbnailWidth = &v
}


// GetThumbnailHeight returns the ThumbnailHeight field value if set, zero value otherwise.
func (o *InlineQueryResultDocument) GetThumbnailHeight() int32 {
	if o == nil || IsNil(o.ThumbnailHeight) {
		var ret int32
		return ret
	}
	return *o.ThumbnailHeight
}

// GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultDocument) GetThumbnailHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailHeight) {
		return nil, false
	}
	return o.ThumbnailHeight, true
}

// HasThumbnailHeight returns a boolean if a field has been set.
func (o *InlineQueryResultDocument) HasThumbnailHeight() bool {
	if o != nil && !IsNil(o.ThumbnailHeight) {
		return true
	}

	return false
}

// SetThumbnailHeight gets a reference to the given int32 and assigns it to the ThumbnailHeight field.
func (o *InlineQueryResultDocument) SetThumbnailHeight(v int32) {
	o.ThumbnailHeight = &v
}


func (o InlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
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
	toSerialize["document_url"] = o.DocumentUrl
	toSerialize["mime_type"] = o.MimeType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if !IsNil(o.ThumbnailWidth) {
		toSerialize["thumbnail_width"] = o.ThumbnailWidth
	}
	if !IsNil(o.ThumbnailHeight) {
		toSerialize["thumbnail_height"] = o.ThumbnailHeight
	}
	return toSerialize, nil
}

func (o *InlineQueryResultDocument) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"title",
		"document_url",
		"mime_type",
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

	varInlineQueryResultDocument := _InlineQueryResultDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultDocument)

	if err != nil {
		return err
	}

	*o = InlineQueryResultDocument(varInlineQueryResultDocument)

	return err
}

type NullableInlineQueryResultDocument struct {
	value *InlineQueryResultDocument
	isSet bool
}

func (v NullableInlineQueryResultDocument) Get() *InlineQueryResultDocument {
	return v.value
}

func (v *NullableInlineQueryResultDocument) Set(val *InlineQueryResultDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultDocument(val *InlineQueryResultDocument) *NullableInlineQueryResultDocument {
	return &NullableInlineQueryResultDocument{value: val, isSet: true}
}

func (v NullableInlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


