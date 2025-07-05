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

// checks if the MessageEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEntity{}

// MessageEntity This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (`@username`), “hashtag” (`#hashtag` or `#hashtag@chatusername`), “cashtag” (`$USD` or `$USD@chatusername`), “bot\\_command” (`/start@jobs_bot`), “url” (`https://telegram.org`), “email” (`do-not-reply@telegram.org`), “phone\\_number” (`+1-212-555-0123`), “bold” (**bold text**), “italic” (*italic text*), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “expandable\\_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block), “text\\_link” (for clickable text URLs), “text\\_mention” (for users [without usernames](https://telegram.org/blog/edit#new-mentions)), “custom\\_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`
	// Offset in [UTF-16 code units](https://core.telegram.org/api/entities#entity-length) to the start of the entity
	Offset int32 `json:"offset"`
	// Length of the entity in [UTF-16 code units](https://core.telegram.org/api/entities#entity-length)
	Length int32 `json:"length"`
	// *Optional*. For “text\\_link” only, URL that will be opened after user taps on the text
	Url *string `json:"url,omitempty"`
	User *User `json:"user,omitempty"`
	// *Optional*. For “pre” only, the programming language of the entity text
	Language *string `json:"language,omitempty"`
	// *Optional*. For “custom\\_emoji” only, unique identifier of the custom emoji. Use [getCustomEmojiStickers](https://core.telegram.org/bots/api/#getcustomemojistickers) to get full information about the sticker
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
}

type _MessageEntity MessageEntity

// NewMessageEntity instantiates a new MessageEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEntity(type_ string, offset int32, length int32) *MessageEntity {
	this := MessageEntity{}
	this.Type = type_
	this.Offset = offset
	this.Length = length
	return &this
}

// NewMessageEntityWithDefaults instantiates a new MessageEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEntityWithDefaults() *MessageEntity {
	this := MessageEntity{}
	return &this
}

// GetType returns the Type field value
func (o *MessageEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageEntity) SetType(v string) {
	o.Type = v
}

// GetOffset returns the Offset field value
func (o *MessageEntity) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *MessageEntity) SetOffset(v int32) {
	o.Offset = v
}

// GetLength returns the Length field value
func (o *MessageEntity) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *MessageEntity) SetLength(v int32) {
	o.Length = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageEntity) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEntity) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageEntity) SetUrl(v string) {
	o.Url = &v
}


// GetUser returns the User field value if set, zero value otherwise.
func (o *MessageEntity) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MessageEntity) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *MessageEntity) SetUser(v User) {
	o.User = &v
}


// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MessageEntity) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MessageEntity) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MessageEntity) SetLanguage(v string) {
	o.Language = &v
}


// GetCustomEmojiId returns the CustomEmojiId field value if set, zero value otherwise.
func (o *MessageEntity) GetCustomEmojiId() string {
	if o == nil || IsNil(o.CustomEmojiId) {
		var ret string
		return ret
	}
	return *o.CustomEmojiId
}

// GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEntity) GetCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEmojiId) {
		return nil, false
	}
	return o.CustomEmojiId, true
}

// HasCustomEmojiId returns a boolean if a field has been set.
func (o *MessageEntity) HasCustomEmojiId() bool {
	if o != nil && !IsNil(o.CustomEmojiId) {
		return true
	}

	return false
}

// SetCustomEmojiId gets a reference to the given string and assigns it to the CustomEmojiId field.
func (o *MessageEntity) SetCustomEmojiId(v string) {
	o.CustomEmojiId = &v
}


func (o MessageEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["offset"] = o.Offset
	toSerialize["length"] = o.Length
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.CustomEmojiId) {
		toSerialize["custom_emoji_id"] = o.CustomEmojiId
	}
	return toSerialize, nil
}

func (o *MessageEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"offset",
		"length",
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

	varMessageEntity := _MessageEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEntity)

	if err != nil {
		return err
	}

	*o = MessageEntity(varMessageEntity)

	return err
}

type NullableMessageEntity struct {
	value *MessageEntity
	isSet bool
}

func (v NullableMessageEntity) Get() *MessageEntity {
	return v.value
}

func (v *NullableMessageEntity) Set(val *MessageEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEntity(val *MessageEntity) *NullableMessageEntity {
	return &NullableMessageEntity{value: val, isSet: true}
}

func (v NullableMessageEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


