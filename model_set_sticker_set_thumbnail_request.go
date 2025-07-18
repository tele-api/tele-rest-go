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

// checks if the SetStickerSetThumbnailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetStickerSetThumbnailRequest{}

// SetStickerSetThumbnailRequest Request parameters for setStickerSetThumbnail
type SetStickerSetThumbnailRequest struct {
	// Sticker set name
	Name string `json:"name"`
	// User identifier of the sticker set owner
	UserId int32 `json:"user_id"`
	Thumbnail NullableString `json:"thumbnail,omitempty"`
	// Format of the thumbnail, must be one of “static” for a **.WEBP** or **.PNG** image, “animated” for a **.TGS** animation, or “video” for a **.WEBM** video
	Format string `json:"format"`
}

type _SetStickerSetThumbnailRequest SetStickerSetThumbnailRequest

// NewSetStickerSetThumbnailRequest instantiates a new SetStickerSetThumbnailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetStickerSetThumbnailRequest(name string, userId int32, format string) *SetStickerSetThumbnailRequest {
	this := SetStickerSetThumbnailRequest{}
	this.Name = name
	this.UserId = userId
	this.Format = format
	return &this
}

// NewSetStickerSetThumbnailRequestWithDefaults instantiates a new SetStickerSetThumbnailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetStickerSetThumbnailRequestWithDefaults() *SetStickerSetThumbnailRequest {
	this := SetStickerSetThumbnailRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SetStickerSetThumbnailRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SetStickerSetThumbnailRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SetStickerSetThumbnailRequest) SetName(v string) {
	o.Name = v
}

// GetUserId returns the UserId field value
func (o *SetStickerSetThumbnailRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SetStickerSetThumbnailRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SetStickerSetThumbnailRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetStickerSetThumbnailRequest) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbnail.Get()
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetStickerSetThumbnailRequest) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail.Get()) {
		return nil, false
	}
	return o.Thumbnail.Get(), o.Thumbnail.IsSet()
}

// HasThumbnail returns a boolean if a field has been set.
func (o *SetStickerSetThumbnailRequest) HasThumbnail() bool {
	if o != nil && o.Thumbnail.IsSet() {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given NullableString and assigns it to the Thumbnail field.
func (o *SetStickerSetThumbnailRequest) SetThumbnail(v string) {
	o.Thumbnail.Set(&v)
}

// SetThumbnailNil sets the value for Thumbnail to be an explicit nil
func (o *SetStickerSetThumbnailRequest) SetThumbnailNil() {
	o.Thumbnail.Set(nil)
}

// UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
func (o *SetStickerSetThumbnailRequest) UnsetThumbnail() {
	o.Thumbnail.Unset()
}

// GetFormat returns the Format field value
func (o *SetStickerSetThumbnailRequest) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *SetStickerSetThumbnailRequest) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *SetStickerSetThumbnailRequest) SetFormat(v string) {
	o.Format = v
}

func (o SetStickerSetThumbnailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetStickerSetThumbnailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["user_id"] = o.UserId
	if o.Thumbnail.IsSet() {
		toSerialize["thumbnail"] = o.Thumbnail.Get()
	}
	toSerialize["format"] = o.Format
	return toSerialize, nil
}

func (o *SetStickerSetThumbnailRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"user_id",
		"format",
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

	varSetStickerSetThumbnailRequest := _SetStickerSetThumbnailRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetStickerSetThumbnailRequest)

	if err != nil {
		return err
	}

	*o = SetStickerSetThumbnailRequest(varSetStickerSetThumbnailRequest)

	return err
}

type NullableSetStickerSetThumbnailRequest struct {
	value *SetStickerSetThumbnailRequest
	isSet bool
}

func (v NullableSetStickerSetThumbnailRequest) Get() *SetStickerSetThumbnailRequest {
	return v.value
}

func (v *NullableSetStickerSetThumbnailRequest) Set(val *SetStickerSetThumbnailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetStickerSetThumbnailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetStickerSetThumbnailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetStickerSetThumbnailRequest(val *SetStickerSetThumbnailRequest) *NullableSetStickerSetThumbnailRequest {
	return &NullableSetStickerSetThumbnailRequest{value: val, isSet: true}
}

func (v NullableSetStickerSetThumbnailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetStickerSetThumbnailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


