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

// checks if the ChatPhoto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatPhoto{}

// ChatPhoto This object represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo. This file\\_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo. This file\\_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

type _ChatPhoto ChatPhoto

// NewChatPhoto instantiates a new ChatPhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatPhoto(smallFileId string, smallFileUniqueId string, bigFileId string, bigFileUniqueId string) *ChatPhoto {
	this := ChatPhoto{}
	this.SmallFileId = smallFileId
	this.SmallFileUniqueId = smallFileUniqueId
	this.BigFileId = bigFileId
	this.BigFileUniqueId = bigFileUniqueId
	return &this
}

// NewChatPhotoWithDefaults instantiates a new ChatPhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatPhotoWithDefaults() *ChatPhoto {
	this := ChatPhoto{}
	return &this
}

// GetSmallFileId returns the SmallFileId field value
func (o *ChatPhoto) GetSmallFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmallFileId
}

// GetSmallFileIdOk returns a tuple with the SmallFileId field value
// and a boolean to check if the value has been set.
func (o *ChatPhoto) GetSmallFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmallFileId, true
}

// SetSmallFileId sets field value
func (o *ChatPhoto) SetSmallFileId(v string) {
	o.SmallFileId = v
}

// GetSmallFileUniqueId returns the SmallFileUniqueId field value
func (o *ChatPhoto) GetSmallFileUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmallFileUniqueId
}

// GetSmallFileUniqueIdOk returns a tuple with the SmallFileUniqueId field value
// and a boolean to check if the value has been set.
func (o *ChatPhoto) GetSmallFileUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmallFileUniqueId, true
}

// SetSmallFileUniqueId sets field value
func (o *ChatPhoto) SetSmallFileUniqueId(v string) {
	o.SmallFileUniqueId = v
}

// GetBigFileId returns the BigFileId field value
func (o *ChatPhoto) GetBigFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BigFileId
}

// GetBigFileIdOk returns a tuple with the BigFileId field value
// and a boolean to check if the value has been set.
func (o *ChatPhoto) GetBigFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BigFileId, true
}

// SetBigFileId sets field value
func (o *ChatPhoto) SetBigFileId(v string) {
	o.BigFileId = v
}

// GetBigFileUniqueId returns the BigFileUniqueId field value
func (o *ChatPhoto) GetBigFileUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BigFileUniqueId
}

// GetBigFileUniqueIdOk returns a tuple with the BigFileUniqueId field value
// and a boolean to check if the value has been set.
func (o *ChatPhoto) GetBigFileUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BigFileUniqueId, true
}

// SetBigFileUniqueId sets field value
func (o *ChatPhoto) SetBigFileUniqueId(v string) {
	o.BigFileUniqueId = v
}

func (o ChatPhoto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatPhoto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["small_file_id"] = o.SmallFileId
	toSerialize["small_file_unique_id"] = o.SmallFileUniqueId
	toSerialize["big_file_id"] = o.BigFileId
	toSerialize["big_file_unique_id"] = o.BigFileUniqueId
	return toSerialize, nil
}

func (o *ChatPhoto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"small_file_id",
		"small_file_unique_id",
		"big_file_id",
		"big_file_unique_id",
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

	varChatPhoto := _ChatPhoto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatPhoto)

	if err != nil {
		return err
	}

	*o = ChatPhoto(varChatPhoto)

	return err
}

type NullableChatPhoto struct {
	value *ChatPhoto
	isSet bool
}

func (v NullableChatPhoto) Get() *ChatPhoto {
	return v.value
}

func (v *NullableChatPhoto) Set(val *ChatPhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatPhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatPhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatPhoto(val *ChatPhoto) *NullableChatPhoto {
	return &NullableChatPhoto{value: val, isSet: true}
}

func (v NullableChatPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatPhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


