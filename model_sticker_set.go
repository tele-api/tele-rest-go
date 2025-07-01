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

// checks if the StickerSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StickerSet{}

// StickerSet This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`
	// Sticker set title
	Title string `json:"title"`
	// Type of stickers in the set, currently one of “regular”, “mask”, “custom\\_emoji”
	StickerType string `json:"sticker_type"`
	// List of all set stickers
	Stickers []Sticker `json:"stickers"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

type _StickerSet StickerSet

// NewStickerSet instantiates a new StickerSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStickerSet(name string, title string, stickerType string, stickers []Sticker) *StickerSet {
	this := StickerSet{}
	this.Name = name
	this.Title = title
	this.StickerType = stickerType
	this.Stickers = stickers
	return &this
}

// NewStickerSetWithDefaults instantiates a new StickerSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerSetWithDefaults() *StickerSet {
	this := StickerSet{}
	return &this
}

// GetName returns the Name field value
func (o *StickerSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StickerSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StickerSet) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *StickerSet) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StickerSet) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *StickerSet) SetTitle(v string) {
	o.Title = v
}

// GetStickerType returns the StickerType field value
func (o *StickerSet) GetStickerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StickerType
}

// GetStickerTypeOk returns a tuple with the StickerType field value
// and a boolean to check if the value has been set.
func (o *StickerSet) GetStickerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StickerType, true
}

// SetStickerType sets field value
func (o *StickerSet) SetStickerType(v string) {
	o.StickerType = v
}

// GetStickers returns the Stickers field value
func (o *StickerSet) GetStickers() []Sticker {
	if o == nil {
		var ret []Sticker
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *StickerSet) GetStickersOk() ([]Sticker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *StickerSet) SetStickers(v []Sticker) {
	o.Stickers = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *StickerSet) GetThumbnail() PhotoSize {
	if o == nil || IsNil(o.Thumbnail) {
		var ret PhotoSize
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StickerSet) GetThumbnailOk() (*PhotoSize, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *StickerSet) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given PhotoSize and assigns it to the Thumbnail field.
func (o *StickerSet) SetThumbnail(v PhotoSize) {
	o.Thumbnail = &v
}


func (o StickerSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StickerSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["sticker_type"] = o.StickerType
	toSerialize["stickers"] = o.Stickers
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	return toSerialize, nil
}

func (o *StickerSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"title",
		"sticker_type",
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

	varStickerSet := _StickerSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStickerSet)

	if err != nil {
		return err
	}

	*o = StickerSet(varStickerSet)

	return err
}

type NullableStickerSet struct {
	value *StickerSet
	isSet bool
}

func (v NullableStickerSet) Get() *StickerSet {
	return v.value
}

func (v *NullableStickerSet) Set(val *StickerSet) {
	v.value = val
	v.isSet = true
}

func (v NullableStickerSet) IsSet() bool {
	return v.isSet
}

func (v *NullableStickerSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStickerSet(val *StickerSet) *NullableStickerSet {
	return &NullableStickerSet{value: val, isSet: true}
}

func (v NullableStickerSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStickerSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


