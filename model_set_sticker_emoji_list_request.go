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

// checks if the SetStickerEmojiListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetStickerEmojiListRequest{}

// SetStickerEmojiListRequest Request parameters for setStickerEmojiList
type SetStickerEmojiListRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
}

type _SetStickerEmojiListRequest SetStickerEmojiListRequest

// NewSetStickerEmojiListRequest instantiates a new SetStickerEmojiListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetStickerEmojiListRequest(sticker string, emojiList []string) *SetStickerEmojiListRequest {
	this := SetStickerEmojiListRequest{}
	this.Sticker = sticker
	this.EmojiList = emojiList
	return &this
}

// NewSetStickerEmojiListRequestWithDefaults instantiates a new SetStickerEmojiListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetStickerEmojiListRequestWithDefaults() *SetStickerEmojiListRequest {
	this := SetStickerEmojiListRequest{}
	return &this
}

// GetSticker returns the Sticker field value
func (o *SetStickerEmojiListRequest) GetSticker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value
// and a boolean to check if the value has been set.
func (o *SetStickerEmojiListRequest) GetStickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticker, true
}

// SetSticker sets field value
func (o *SetStickerEmojiListRequest) SetSticker(v string) {
	o.Sticker = v
}

// GetEmojiList returns the EmojiList field value
func (o *SetStickerEmojiListRequest) GetEmojiList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmojiList
}

// GetEmojiListOk returns a tuple with the EmojiList field value
// and a boolean to check if the value has been set.
func (o *SetStickerEmojiListRequest) GetEmojiListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmojiList, true
}

// SetEmojiList sets field value
func (o *SetStickerEmojiListRequest) SetEmojiList(v []string) {
	o.EmojiList = v
}

func (o SetStickerEmojiListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetStickerEmojiListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sticker"] = o.Sticker
	toSerialize["emoji_list"] = o.EmojiList
	return toSerialize, nil
}

func (o *SetStickerEmojiListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sticker",
		"emoji_list",
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

	varSetStickerEmojiListRequest := _SetStickerEmojiListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetStickerEmojiListRequest)

	if err != nil {
		return err
	}

	*o = SetStickerEmojiListRequest(varSetStickerEmojiListRequest)

	return err
}

type NullableSetStickerEmojiListRequest struct {
	value *SetStickerEmojiListRequest
	isSet bool
}

func (v NullableSetStickerEmojiListRequest) Get() *SetStickerEmojiListRequest {
	return v.value
}

func (v *NullableSetStickerEmojiListRequest) Set(val *SetStickerEmojiListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetStickerEmojiListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetStickerEmojiListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetStickerEmojiListRequest(val *SetStickerEmojiListRequest) *NullableSetStickerEmojiListRequest {
	return &NullableSetStickerEmojiListRequest{value: val, isSet: true}
}

func (v NullableSetStickerEmojiListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetStickerEmojiListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


