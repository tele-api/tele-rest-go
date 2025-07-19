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

// checks if the InlineQueryResultArticle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultArticle{}

// InlineQueryResultArticle Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Type of the result, must be *article*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Title of the result
	Title string `json:"title"`
	InputMessageContent InputMessageContent `json:"input_message_content"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. URL of the result
	Url *string `json:"url,omitempty"`
	// *Optional*. Short description of the result
	Description *string `json:"description,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`
}

type _InlineQueryResultArticle InlineQueryResultArticle

// NewInlineQueryResultArticle instantiates a new InlineQueryResultArticle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultArticle(type_ string, id string, title string, inputMessageContent InputMessageContent) *InlineQueryResultArticle {
	this := InlineQueryResultArticle{}
	this.Type = type_
	this.Id = id
	this.Title = title
	this.InputMessageContent = inputMessageContent
	return &this
}

// NewInlineQueryResultArticleWithDefaults instantiates a new InlineQueryResultArticle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultArticleWithDefaults() *InlineQueryResultArticle {
	this := InlineQueryResultArticle{}
	var type_ string = "article"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultArticle) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultArticle) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultArticle) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultArticle) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultArticle) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultArticle) SetTitle(v string) {
	o.Title = v
}

// GetInputMessageContent returns the InputMessageContent field value
func (o *InlineQueryResultArticle) GetInputMessageContent() InputMessageContent {
	if o == nil {
		var ret InputMessageContent
		return ret
	}

	return o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputMessageContent, true
}

// SetInputMessageContent sets field value
func (o *InlineQueryResultArticle) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = v
}

// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultArticle) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineQueryResultArticle) SetUrl(v string) {
	o.Url = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineQueryResultArticle) SetDescription(v string) {
	o.Description = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *InlineQueryResultArticle) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}


// GetThumbnailWidth returns the ThumbnailWidth field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetThumbnailWidth() int32 {
	if o == nil || IsNil(o.ThumbnailWidth) {
		var ret int32
		return ret
	}
	return *o.ThumbnailWidth
}

// GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetThumbnailWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailWidth) {
		return nil, false
	}
	return o.ThumbnailWidth, true
}

// HasThumbnailWidth returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasThumbnailWidth() bool {
	if o != nil && !IsNil(o.ThumbnailWidth) {
		return true
	}

	return false
}

// SetThumbnailWidth gets a reference to the given int32 and assigns it to the ThumbnailWidth field.
func (o *InlineQueryResultArticle) SetThumbnailWidth(v int32) {
	o.ThumbnailWidth = &v
}


// GetThumbnailHeight returns the ThumbnailHeight field value if set, zero value otherwise.
func (o *InlineQueryResultArticle) GetThumbnailHeight() int32 {
	if o == nil || IsNil(o.ThumbnailHeight) {
		var ret int32
		return ret
	}
	return *o.ThumbnailHeight
}

// GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultArticle) GetThumbnailHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailHeight) {
		return nil, false
	}
	return o.ThumbnailHeight, true
}

// HasThumbnailHeight returns a boolean if a field has been set.
func (o *InlineQueryResultArticle) HasThumbnailHeight() bool {
	if o != nil && !IsNil(o.ThumbnailHeight) {
		return true
	}

	return false
}

// SetThumbnailHeight gets a reference to the given int32 and assigns it to the ThumbnailHeight field.
func (o *InlineQueryResultArticle) SetThumbnailHeight(v int32) {
	o.ThumbnailHeight = &v
}


func (o InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultArticle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["input_message_content"] = o.InputMessageContent
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *InlineQueryResultArticle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"title",
		"input_message_content",
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

	varInlineQueryResultArticle := _InlineQueryResultArticle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultArticle)

	if err != nil {
		return err
	}

	*o = InlineQueryResultArticle(varInlineQueryResultArticle)

	return err
}

type NullableInlineQueryResultArticle struct {
	value *InlineQueryResultArticle
	isSet bool
}

func (v NullableInlineQueryResultArticle) Get() *InlineQueryResultArticle {
	return v.value
}

func (v *NullableInlineQueryResultArticle) Set(val *InlineQueryResultArticle) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultArticle) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultArticle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultArticle(val *InlineQueryResultArticle) *NullableInlineQueryResultArticle {
	return &NullableInlineQueryResultArticle{value: val, isSet: true}
}

func (v NullableInlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultArticle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


