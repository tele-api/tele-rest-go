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

// checks if the ForumTopic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumTopic{}

// ForumTopic This object represents a forum topic.
type ForumTopic struct {
	// Unique identifier of the forum topic
	MessageThreadId int32 `json:"message_thread_id"`
	// Name of the topic
	Name string `json:"name"`
	// Color of the topic icon in RGB format
	IconColor int32 `json:"icon_color"`
	// *Optional*. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type _ForumTopic ForumTopic

// NewForumTopic instantiates a new ForumTopic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumTopic(messageThreadId int32, name string, iconColor int32) *ForumTopic {
	this := ForumTopic{}
	this.MessageThreadId = messageThreadId
	this.Name = name
	this.IconColor = iconColor
	return &this
}

// NewForumTopicWithDefaults instantiates a new ForumTopic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumTopicWithDefaults() *ForumTopic {
	this := ForumTopic{}
	return &this
}

// GetMessageThreadId returns the MessageThreadId field value
func (o *ForumTopic) GetMessageThreadId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value
// and a boolean to check if the value has been set.
func (o *ForumTopic) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageThreadId, true
}

// SetMessageThreadId sets field value
func (o *ForumTopic) SetMessageThreadId(v int32) {
	o.MessageThreadId = v
}

// GetName returns the Name field value
func (o *ForumTopic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ForumTopic) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ForumTopic) SetName(v string) {
	o.Name = v
}

// GetIconColor returns the IconColor field value
func (o *ForumTopic) GetIconColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IconColor
}

// GetIconColorOk returns a tuple with the IconColor field value
// and a boolean to check if the value has been set.
func (o *ForumTopic) GetIconColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconColor, true
}

// SetIconColor sets field value
func (o *ForumTopic) SetIconColor(v int32) {
	o.IconColor = v
}

// GetIconCustomEmojiId returns the IconCustomEmojiId field value if set, zero value otherwise.
func (o *ForumTopic) GetIconCustomEmojiId() string {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.IconCustomEmojiId
}

// GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopic) GetIconCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		return nil, false
	}
	return o.IconCustomEmojiId, true
}

// HasIconCustomEmojiId returns a boolean if a field has been set.
func (o *ForumTopic) HasIconCustomEmojiId() bool {
	if o != nil && !IsNil(o.IconCustomEmojiId) {
		return true
	}

	return false
}

// SetIconCustomEmojiId gets a reference to the given string and assigns it to the IconCustomEmojiId field.
func (o *ForumTopic) SetIconCustomEmojiId(v string) {
	o.IconCustomEmojiId = &v
}


func (o ForumTopic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumTopic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_thread_id"] = o.MessageThreadId
	toSerialize["name"] = o.Name
	toSerialize["icon_color"] = o.IconColor
	if !IsNil(o.IconCustomEmojiId) {
		toSerialize["icon_custom_emoji_id"] = o.IconCustomEmojiId
	}
	return toSerialize, nil
}

func (o *ForumTopic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_thread_id",
		"name",
		"icon_color",
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

	varForumTopic := _ForumTopic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForumTopic)

	if err != nil {
		return err
	}

	*o = ForumTopic(varForumTopic)

	return err
}

type NullableForumTopic struct {
	value *ForumTopic
	isSet bool
}

func (v NullableForumTopic) Get() *ForumTopic {
	return v.value
}

func (v *NullableForumTopic) Set(val *ForumTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableForumTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableForumTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumTopic(val *ForumTopic) *NullableForumTopic {
	return &NullableForumTopic{value: val, isSet: true}
}

func (v NullableForumTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


