/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the InlineQueryResultContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultContact{}

// InlineQueryResultContact Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	// Type of the result, must be *contact*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// *Optional*. Contact's last name
	LastName *string `json:"last_name,omitempty"`
	// *Optional*. Additional data about the contact in the form of a [vCard](https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard *string `json:"vcard,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`
}

type _InlineQueryResultContact InlineQueryResultContact

// NewInlineQueryResultContact instantiates a new InlineQueryResultContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultContact(type_ string, id string, phoneNumber string, firstName string) *InlineQueryResultContact {
	this := InlineQueryResultContact{}
	this.Type = type_
	this.Id = id
	this.PhoneNumber = phoneNumber
	this.FirstName = firstName
	return &this
}

// NewInlineQueryResultContactWithDefaults instantiates a new InlineQueryResultContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultContactWithDefaults() *InlineQueryResultContact {
	this := InlineQueryResultContact{}
	var type_ string = "contact"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultContact) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultContact) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultContact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultContact) SetId(v string) {
	o.Id = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *InlineQueryResultContact) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *InlineQueryResultContact) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetFirstName returns the FirstName field value
func (o *InlineQueryResultContact) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *InlineQueryResultContact) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InlineQueryResultContact) SetLastName(v string) {
	o.LastName = &v
}


// GetVcard returns the Vcard field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetVcard() string {
	if o == nil || IsNil(o.Vcard) {
		var ret string
		return ret
	}
	return *o.Vcard
}

// GetVcardOk returns a tuple with the Vcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetVcardOk() (*string, bool) {
	if o == nil || IsNil(o.Vcard) {
		return nil, false
	}
	return o.Vcard, true
}

// HasVcard returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasVcard() bool {
	if o != nil && !IsNil(o.Vcard) {
		return true
	}

	return false
}

// SetVcard gets a reference to the given string and assigns it to the Vcard field.
func (o *InlineQueryResultContact) SetVcard(v string) {
	o.Vcard = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultContact) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultContact) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *InlineQueryResultContact) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}


// GetThumbnailWidth returns the ThumbnailWidth field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetThumbnailWidth() int32 {
	if o == nil || IsNil(o.ThumbnailWidth) {
		var ret int32
		return ret
	}
	return *o.ThumbnailWidth
}

// GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetThumbnailWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailWidth) {
		return nil, false
	}
	return o.ThumbnailWidth, true
}

// HasThumbnailWidth returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasThumbnailWidth() bool {
	if o != nil && !IsNil(o.ThumbnailWidth) {
		return true
	}

	return false
}

// SetThumbnailWidth gets a reference to the given int32 and assigns it to the ThumbnailWidth field.
func (o *InlineQueryResultContact) SetThumbnailWidth(v int32) {
	o.ThumbnailWidth = &v
}


// GetThumbnailHeight returns the ThumbnailHeight field value if set, zero value otherwise.
func (o *InlineQueryResultContact) GetThumbnailHeight() int32 {
	if o == nil || IsNil(o.ThumbnailHeight) {
		var ret int32
		return ret
	}
	return *o.ThumbnailHeight
}

// GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultContact) GetThumbnailHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailHeight) {
		return nil, false
	}
	return o.ThumbnailHeight, true
}

// HasThumbnailHeight returns a boolean if a field has been set.
func (o *InlineQueryResultContact) HasThumbnailHeight() bool {
	if o != nil && !IsNil(o.ThumbnailHeight) {
		return true
	}

	return false
}

// SetThumbnailHeight gets a reference to the given int32 and assigns it to the ThumbnailHeight field.
func (o *InlineQueryResultContact) SetThumbnailHeight(v int32) {
	o.ThumbnailHeight = &v
}


func (o InlineQueryResultContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Vcard) {
		toSerialize["vcard"] = o.Vcard
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

func (o *InlineQueryResultContact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"phone_number",
		"first_name",
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

	varInlineQueryResultContact := _InlineQueryResultContact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultContact)

	if err != nil {
		return err
	}

	*o = InlineQueryResultContact(varInlineQueryResultContact)

	return err
}

type NullableInlineQueryResultContact struct {
	value *InlineQueryResultContact
	isSet bool
}

func (v NullableInlineQueryResultContact) Get() *InlineQueryResultContact {
	return v.value
}

func (v *NullableInlineQueryResultContact) Set(val *InlineQueryResultContact) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultContact) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultContact(val *InlineQueryResultContact) *NullableInlineQueryResultContact {
	return &NullableInlineQueryResultContact{value: val, isSet: true}
}

func (v NullableInlineQueryResultContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


