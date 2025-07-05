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

// checks if the CreateNewStickerSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNewStickerSetRequest{}

// CreateNewStickerSetRequest Request parameters for createNewStickerSet
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserId int32 `json:"user_id"`
	// Short name of sticker set, to be used in `t.me/addstickers/` URLs (e.g., *animals*). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in `\"_by_<bot_username>\"`. `<bot_username>` is case insensitive. 1-64 characters.
	Name string `json:"name"`
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
	// A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []InputSticker `json:"stickers"`
	// Type of stickers in the set, pass “regular”, “mask”, or “custom\\_emoji”. By default, a regular sticker set is created.
	StickerType *string `json:"sticker_type,omitempty"`
	// Pass *True* if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting *bool `json:"needs_repainting,omitempty"`
}

type _CreateNewStickerSetRequest CreateNewStickerSetRequest

// NewCreateNewStickerSetRequest instantiates a new CreateNewStickerSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNewStickerSetRequest(userId int32, name string, title string, stickers []InputSticker) *CreateNewStickerSetRequest {
	this := CreateNewStickerSetRequest{}
	this.UserId = userId
	this.Name = name
	this.Title = title
	this.Stickers = stickers
	return &this
}

// NewCreateNewStickerSetRequestWithDefaults instantiates a new CreateNewStickerSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNewStickerSetRequestWithDefaults() *CreateNewStickerSetRequest {
	this := CreateNewStickerSetRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *CreateNewStickerSetRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateNewStickerSetRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *CreateNewStickerSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNewStickerSetRequest) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CreateNewStickerSetRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateNewStickerSetRequest) SetTitle(v string) {
	o.Title = v
}

// GetStickers returns the Stickers field value
func (o *CreateNewStickerSetRequest) GetStickers() []InputSticker {
	if o == nil {
		var ret []InputSticker
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetStickersOk() ([]InputSticker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *CreateNewStickerSetRequest) SetStickers(v []InputSticker) {
	o.Stickers = v
}

// GetStickerType returns the StickerType field value if set, zero value otherwise.
func (o *CreateNewStickerSetRequest) GetStickerType() string {
	if o == nil || IsNil(o.StickerType) {
		var ret string
		return ret
	}
	return *o.StickerType
}

// GetStickerTypeOk returns a tuple with the StickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetStickerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StickerType) {
		return nil, false
	}
	return o.StickerType, true
}

// HasStickerType returns a boolean if a field has been set.
func (o *CreateNewStickerSetRequest) HasStickerType() bool {
	if o != nil && !IsNil(o.StickerType) {
		return true
	}

	return false
}

// SetStickerType gets a reference to the given string and assigns it to the StickerType field.
func (o *CreateNewStickerSetRequest) SetStickerType(v string) {
	o.StickerType = &v
}


// GetNeedsRepainting returns the NeedsRepainting field value if set, zero value otherwise.
func (o *CreateNewStickerSetRequest) GetNeedsRepainting() bool {
	if o == nil || IsNil(o.NeedsRepainting) {
		var ret bool
		return ret
	}
	return *o.NeedsRepainting
}

// GetNeedsRepaintingOk returns a tuple with the NeedsRepainting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNewStickerSetRequest) GetNeedsRepaintingOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsRepainting) {
		return nil, false
	}
	return o.NeedsRepainting, true
}

// HasNeedsRepainting returns a boolean if a field has been set.
func (o *CreateNewStickerSetRequest) HasNeedsRepainting() bool {
	if o != nil && !IsNil(o.NeedsRepainting) {
		return true
	}

	return false
}

// SetNeedsRepainting gets a reference to the given bool and assigns it to the NeedsRepainting field.
func (o *CreateNewStickerSetRequest) SetNeedsRepainting(v bool) {
	o.NeedsRepainting = &v
}


func (o CreateNewStickerSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNewStickerSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["stickers"] = o.Stickers
	if !IsNil(o.StickerType) {
		toSerialize["sticker_type"] = o.StickerType
	}
	if !IsNil(o.NeedsRepainting) {
		toSerialize["needs_repainting"] = o.NeedsRepainting
	}
	return toSerialize, nil
}

func (o *CreateNewStickerSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"name",
		"title",
		"stickers",
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

	varCreateNewStickerSetRequest := _CreateNewStickerSetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNewStickerSetRequest)

	if err != nil {
		return err
	}

	*o = CreateNewStickerSetRequest(varCreateNewStickerSetRequest)

	return err
}

type NullableCreateNewStickerSetRequest struct {
	value *CreateNewStickerSetRequest
	isSet bool
}

func (v NullableCreateNewStickerSetRequest) Get() *CreateNewStickerSetRequest {
	return v.value
}

func (v *NullableCreateNewStickerSetRequest) Set(val *CreateNewStickerSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNewStickerSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNewStickerSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNewStickerSetRequest(val *CreateNewStickerSetRequest) *NullableCreateNewStickerSetRequest {
	return &NullableCreateNewStickerSetRequest{value: val, isSet: true}
}

func (v NullableCreateNewStickerSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNewStickerSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


