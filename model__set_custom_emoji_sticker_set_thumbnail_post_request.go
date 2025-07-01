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

// checks if the SetCustomEmojiStickerSetThumbnailPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetCustomEmojiStickerSetThumbnailPostRequest{}

// SetCustomEmojiStickerSetThumbnailPostRequest struct for SetCustomEmojiStickerSetThumbnailPostRequest
type SetCustomEmojiStickerSetThumbnailPostRequest struct {
	// Sticker set name
	Name string `json:"name"`
	// Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
}

type _SetCustomEmojiStickerSetThumbnailPostRequest SetCustomEmojiStickerSetThumbnailPostRequest

// NewSetCustomEmojiStickerSetThumbnailPostRequest instantiates a new SetCustomEmojiStickerSetThumbnailPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCustomEmojiStickerSetThumbnailPostRequest(name string) *SetCustomEmojiStickerSetThumbnailPostRequest {
	this := SetCustomEmojiStickerSetThumbnailPostRequest{}
	this.Name = name
	return &this
}

// NewSetCustomEmojiStickerSetThumbnailPostRequestWithDefaults instantiates a new SetCustomEmojiStickerSetThumbnailPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCustomEmojiStickerSetThumbnailPostRequestWithDefaults() *SetCustomEmojiStickerSetThumbnailPostRequest {
	this := SetCustomEmojiStickerSetThumbnailPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) SetName(v string) {
	o.Name = v
}

// GetCustomEmojiId returns the CustomEmojiId field value if set, zero value otherwise.
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) GetCustomEmojiId() string {
	if o == nil || IsNil(o.CustomEmojiId) {
		var ret string
		return ret
	}
	return *o.CustomEmojiId
}

// GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) GetCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEmojiId) {
		return nil, false
	}
	return o.CustomEmojiId, true
}

// HasCustomEmojiId returns a boolean if a field has been set.
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) HasCustomEmojiId() bool {
	if o != nil && !IsNil(o.CustomEmojiId) {
		return true
	}

	return false
}

// SetCustomEmojiId gets a reference to the given string and assigns it to the CustomEmojiId field.
func (o *SetCustomEmojiStickerSetThumbnailPostRequest) SetCustomEmojiId(v string) {
	o.CustomEmojiId = &v
}


func (o SetCustomEmojiStickerSetThumbnailPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetCustomEmojiStickerSetThumbnailPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.CustomEmojiId) {
		toSerialize["custom_emoji_id"] = o.CustomEmojiId
	}
	return toSerialize, nil
}

func (o *SetCustomEmojiStickerSetThumbnailPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSetCustomEmojiStickerSetThumbnailPostRequest := _SetCustomEmojiStickerSetThumbnailPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetCustomEmojiStickerSetThumbnailPostRequest)

	if err != nil {
		return err
	}

	*o = SetCustomEmojiStickerSetThumbnailPostRequest(varSetCustomEmojiStickerSetThumbnailPostRequest)

	return err
}

type NullableSetCustomEmojiStickerSetThumbnailPostRequest struct {
	value *SetCustomEmojiStickerSetThumbnailPostRequest
	isSet bool
}

func (v NullableSetCustomEmojiStickerSetThumbnailPostRequest) Get() *SetCustomEmojiStickerSetThumbnailPostRequest {
	return v.value
}

func (v *NullableSetCustomEmojiStickerSetThumbnailPostRequest) Set(val *SetCustomEmojiStickerSetThumbnailPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCustomEmojiStickerSetThumbnailPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCustomEmojiStickerSetThumbnailPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCustomEmojiStickerSetThumbnailPostRequest(val *SetCustomEmojiStickerSetThumbnailPostRequest) *NullableSetCustomEmojiStickerSetThumbnailPostRequest {
	return &NullableSetCustomEmojiStickerSetThumbnailPostRequest{value: val, isSet: true}
}

func (v NullableSetCustomEmojiStickerSetThumbnailPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCustomEmojiStickerSetThumbnailPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


