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
)

// checks if the ForumTopicEdited type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumTopicEdited{}

// ForumTopicEdited This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// *Optional*. New name of the topic, if it was edited
	Name *string `json:"name,omitempty"`
	// *Optional*. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

// NewForumTopicEdited instantiates a new ForumTopicEdited object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumTopicEdited() *ForumTopicEdited {
	this := ForumTopicEdited{}
	return &this
}

// NewForumTopicEditedWithDefaults instantiates a new ForumTopicEdited object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumTopicEditedWithDefaults() *ForumTopicEdited {
	this := ForumTopicEdited{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ForumTopicEdited) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicEdited) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ForumTopicEdited) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ForumTopicEdited) SetName(v string) {
	o.Name = &v
}


// GetIconCustomEmojiId returns the IconCustomEmojiId field value if set, zero value otherwise.
func (o *ForumTopicEdited) GetIconCustomEmojiId() string {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		var ret string
		return ret
	}
	return *o.IconCustomEmojiId
}

// GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicEdited) GetIconCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.IconCustomEmojiId) {
		return nil, false
	}
	return o.IconCustomEmojiId, true
}

// HasIconCustomEmojiId returns a boolean if a field has been set.
func (o *ForumTopicEdited) HasIconCustomEmojiId() bool {
	if o != nil && !IsNil(o.IconCustomEmojiId) {
		return true
	}

	return false
}

// SetIconCustomEmojiId gets a reference to the given string and assigns it to the IconCustomEmojiId field.
func (o *ForumTopicEdited) SetIconCustomEmojiId(v string) {
	o.IconCustomEmojiId = &v
}


func (o ForumTopicEdited) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumTopicEdited) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IconCustomEmojiId) {
		toSerialize["icon_custom_emoji_id"] = o.IconCustomEmojiId
	}
	return toSerialize, nil
}

type NullableForumTopicEdited struct {
	value *ForumTopicEdited
	isSet bool
}

func (v NullableForumTopicEdited) Get() *ForumTopicEdited {
	return v.value
}

func (v *NullableForumTopicEdited) Set(val *ForumTopicEdited) {
	v.value = val
	v.isSet = true
}

func (v NullableForumTopicEdited) IsSet() bool {
	return v.isSet
}

func (v *NullableForumTopicEdited) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumTopicEdited(val *ForumTopicEdited) *NullableForumTopicEdited {
	return &NullableForumTopicEdited{value: val, isSet: true}
}

func (v NullableForumTopicEdited) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumTopicEdited) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


