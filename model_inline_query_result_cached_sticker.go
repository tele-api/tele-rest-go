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

// checks if the InlineQueryResultCachedSticker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultCachedSticker{}

// InlineQueryResultCachedSticker Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be *sticker*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier of the sticker
	StickerFileId string `json:"sticker_file_id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

type _InlineQueryResultCachedSticker InlineQueryResultCachedSticker

// NewInlineQueryResultCachedSticker instantiates a new InlineQueryResultCachedSticker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultCachedSticker(type_ string, id string, stickerFileId string) *InlineQueryResultCachedSticker {
	this := InlineQueryResultCachedSticker{}
	this.Type = type_
	this.Id = id
	this.StickerFileId = stickerFileId
	return &this
}

// NewInlineQueryResultCachedStickerWithDefaults instantiates a new InlineQueryResultCachedSticker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultCachedStickerWithDefaults() *InlineQueryResultCachedSticker {
	this := InlineQueryResultCachedSticker{}
	var type_ string = "sticker"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultCachedSticker) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedSticker) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultCachedSticker) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultCachedSticker) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedSticker) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultCachedSticker) SetId(v string) {
	o.Id = v
}

// GetStickerFileId returns the StickerFileId field value
func (o *InlineQueryResultCachedSticker) GetStickerFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StickerFileId
}

// GetStickerFileIdOk returns a tuple with the StickerFileId field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedSticker) GetStickerFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StickerFileId, true
}

// SetStickerFileId sets field value
func (o *InlineQueryResultCachedSticker) SetStickerFileId(v string) {
	o.StickerFileId = v
}

// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultCachedSticker) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedSticker) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultCachedSticker) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultCachedSticker) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultCachedSticker) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultCachedSticker) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultCachedSticker) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultCachedSticker) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


func (o InlineQueryResultCachedSticker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultCachedSticker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["sticker_file_id"] = o.StickerFileId
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	return toSerialize, nil
}

func (o *InlineQueryResultCachedSticker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"sticker_file_id",
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

	varInlineQueryResultCachedSticker := _InlineQueryResultCachedSticker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultCachedSticker)

	if err != nil {
		return err
	}

	*o = InlineQueryResultCachedSticker(varInlineQueryResultCachedSticker)

	return err
}

type NullableInlineQueryResultCachedSticker struct {
	value *InlineQueryResultCachedSticker
	isSet bool
}

func (v NullableInlineQueryResultCachedSticker) Get() *InlineQueryResultCachedSticker {
	return v.value
}

func (v *NullableInlineQueryResultCachedSticker) Set(val *InlineQueryResultCachedSticker) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultCachedSticker) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultCachedSticker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultCachedSticker(val *InlineQueryResultCachedSticker) *NullableInlineQueryResultCachedSticker {
	return &NullableInlineQueryResultCachedSticker{value: val, isSet: true}
}

func (v NullableInlineQueryResultCachedSticker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultCachedSticker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


