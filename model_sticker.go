/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the Sticker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sticker{}

// Sticker This object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Type of the sticker, currently one of “regular”, “mask”, “custom\\_emoji”. The type of the sticker is independent from its format, which is determined by the fields *is\\_animated* and *is\\_video*.
	Type string `json:"type"`
	// Sticker width
	Width int32 `json:"width"`
	// Sticker height
	Height int32 `json:"height"`
	// *True*, if the sticker is [animated](https://telegram.org/blog/animated-stickers)
	IsAnimated bool `json:"is_animated"`
	// *True*, if the sticker is a [video sticker](https://telegram.org/blog/video-stickers-better-reactions)
	IsVideo bool `json:"is_video"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// *Optional*. Emoji associated with the sticker
	Emoji *string `json:"emoji,omitempty"`
	// *Optional*. Name of the sticker set to which the sticker belongs
	SetName *string `json:"set_name,omitempty"`
	PremiumAnimation *File `json:"premium_animation,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// *Optional*. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
	// *Optional*. *True*, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	NeedsRepainting *bool `json:"needs_repainting,omitempty"`
	// *Optional*. File size in bytes
	FileSize *int32 `json:"file_size,omitempty"`
}

type _Sticker Sticker

// NewSticker instantiates a new Sticker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSticker(fileId string, fileUniqueId string, type_ string, width int32, height int32, isAnimated bool, isVideo bool) *Sticker {
	this := Sticker{}
	this.FileId = fileId
	this.FileUniqueId = fileUniqueId
	this.Type = type_
	this.Width = width
	this.Height = height
	this.IsAnimated = isAnimated
	this.IsVideo = isVideo
	var needsRepainting bool = true
	this.NeedsRepainting = &needsRepainting
	return &this
}

// NewStickerWithDefaults instantiates a new Sticker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerWithDefaults() *Sticker {
	this := Sticker{}
	var needsRepainting bool = true
	this.NeedsRepainting = &needsRepainting
	return &this
}

// GetFileId returns the FileId field value
func (o *Sticker) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *Sticker) SetFileId(v string) {
	o.FileId = v
}

// GetFileUniqueId returns the FileUniqueId field value
func (o *Sticker) GetFileUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUniqueId
}

// GetFileUniqueIdOk returns a tuple with the FileUniqueId field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetFileUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUniqueId, true
}

// SetFileUniqueId sets field value
func (o *Sticker) SetFileUniqueId(v string) {
	o.FileUniqueId = v
}

// GetType returns the Type field value
func (o *Sticker) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Sticker) SetType(v string) {
	o.Type = v
}

// GetWidth returns the Width field value
func (o *Sticker) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Sticker) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *Sticker) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Sticker) SetHeight(v int32) {
	o.Height = v
}

// GetIsAnimated returns the IsAnimated field value
func (o *Sticker) GetIsAnimated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnimated
}

// GetIsAnimatedOk returns a tuple with the IsAnimated field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetIsAnimatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnimated, true
}

// SetIsAnimated sets field value
func (o *Sticker) SetIsAnimated(v bool) {
	o.IsAnimated = v
}

// GetIsVideo returns the IsVideo field value
func (o *Sticker) GetIsVideo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVideo
}

// GetIsVideoOk returns a tuple with the IsVideo field value
// and a boolean to check if the value has been set.
func (o *Sticker) GetIsVideoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVideo, true
}

// SetIsVideo sets field value
func (o *Sticker) SetIsVideo(v bool) {
	o.IsVideo = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *Sticker) GetThumbnail() PhotoSize {
	if o == nil || IsNil(o.Thumbnail) {
		var ret PhotoSize
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetThumbnailOk() (*PhotoSize, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *Sticker) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given PhotoSize and assigns it to the Thumbnail field.
func (o *Sticker) SetThumbnail(v PhotoSize) {
	o.Thumbnail = &v
}


// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *Sticker) GetEmoji() string {
	if o == nil || IsNil(o.Emoji) {
		var ret string
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.Emoji) {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *Sticker) HasEmoji() bool {
	if o != nil && !IsNil(o.Emoji) {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given string and assigns it to the Emoji field.
func (o *Sticker) SetEmoji(v string) {
	o.Emoji = &v
}


// GetSetName returns the SetName field value if set, zero value otherwise.
func (o *Sticker) GetSetName() string {
	if o == nil || IsNil(o.SetName) {
		var ret string
		return ret
	}
	return *o.SetName
}

// GetSetNameOk returns a tuple with the SetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.SetName) {
		return nil, false
	}
	return o.SetName, true
}

// HasSetName returns a boolean if a field has been set.
func (o *Sticker) HasSetName() bool {
	if o != nil && !IsNil(o.SetName) {
		return true
	}

	return false
}

// SetSetName gets a reference to the given string and assigns it to the SetName field.
func (o *Sticker) SetSetName(v string) {
	o.SetName = &v
}


// GetPremiumAnimation returns the PremiumAnimation field value if set, zero value otherwise.
func (o *Sticker) GetPremiumAnimation() File {
	if o == nil || IsNil(o.PremiumAnimation) {
		var ret File
		return ret
	}
	return *o.PremiumAnimation
}

// GetPremiumAnimationOk returns a tuple with the PremiumAnimation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetPremiumAnimationOk() (*File, bool) {
	if o == nil || IsNil(o.PremiumAnimation) {
		return nil, false
	}
	return o.PremiumAnimation, true
}

// HasPremiumAnimation returns a boolean if a field has been set.
func (o *Sticker) HasPremiumAnimation() bool {
	if o != nil && !IsNil(o.PremiumAnimation) {
		return true
	}

	return false
}

// SetPremiumAnimation gets a reference to the given File and assigns it to the PremiumAnimation field.
func (o *Sticker) SetPremiumAnimation(v File) {
	o.PremiumAnimation = &v
}


// GetMaskPosition returns the MaskPosition field value if set, zero value otherwise.
func (o *Sticker) GetMaskPosition() MaskPosition {
	if o == nil || IsNil(o.MaskPosition) {
		var ret MaskPosition
		return ret
	}
	return *o.MaskPosition
}

// GetMaskPositionOk returns a tuple with the MaskPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetMaskPositionOk() (*MaskPosition, bool) {
	if o == nil || IsNil(o.MaskPosition) {
		return nil, false
	}
	return o.MaskPosition, true
}

// HasMaskPosition returns a boolean if a field has been set.
func (o *Sticker) HasMaskPosition() bool {
	if o != nil && !IsNil(o.MaskPosition) {
		return true
	}

	return false
}

// SetMaskPosition gets a reference to the given MaskPosition and assigns it to the MaskPosition field.
func (o *Sticker) SetMaskPosition(v MaskPosition) {
	o.MaskPosition = &v
}


// GetCustomEmojiId returns the CustomEmojiId field value if set, zero value otherwise.
func (o *Sticker) GetCustomEmojiId() string {
	if o == nil || IsNil(o.CustomEmojiId) {
		var ret string
		return ret
	}
	return *o.CustomEmojiId
}

// GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetCustomEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEmojiId) {
		return nil, false
	}
	return o.CustomEmojiId, true
}

// HasCustomEmojiId returns a boolean if a field has been set.
func (o *Sticker) HasCustomEmojiId() bool {
	if o != nil && !IsNil(o.CustomEmojiId) {
		return true
	}

	return false
}

// SetCustomEmojiId gets a reference to the given string and assigns it to the CustomEmojiId field.
func (o *Sticker) SetCustomEmojiId(v string) {
	o.CustomEmojiId = &v
}


// GetNeedsRepainting returns the NeedsRepainting field value if set, zero value otherwise.
func (o *Sticker) GetNeedsRepainting() bool {
	if o == nil || IsNil(o.NeedsRepainting) {
		var ret bool
		return ret
	}
	return *o.NeedsRepainting
}

// GetNeedsRepaintingOk returns a tuple with the NeedsRepainting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetNeedsRepaintingOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsRepainting) {
		return nil, false
	}
	return o.NeedsRepainting, true
}

// HasNeedsRepainting returns a boolean if a field has been set.
func (o *Sticker) HasNeedsRepainting() bool {
	if o != nil && !IsNil(o.NeedsRepainting) {
		return true
	}

	return false
}

// SetNeedsRepainting gets a reference to the given bool and assigns it to the NeedsRepainting field.
func (o *Sticker) SetNeedsRepainting(v bool) {
	o.NeedsRepainting = &v
}


// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *Sticker) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sticker) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *Sticker) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *Sticker) SetFileSize(v int32) {
	o.FileSize = &v
}


func (o Sticker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sticker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_id"] = o.FileId
	toSerialize["file_unique_id"] = o.FileUniqueId
	toSerialize["type"] = o.Type
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["is_animated"] = o.IsAnimated
	toSerialize["is_video"] = o.IsVideo
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Emoji) {
		toSerialize["emoji"] = o.Emoji
	}
	if !IsNil(o.SetName) {
		toSerialize["set_name"] = o.SetName
	}
	if !IsNil(o.PremiumAnimation) {
		toSerialize["premium_animation"] = o.PremiumAnimation
	}
	if !IsNil(o.MaskPosition) {
		toSerialize["mask_position"] = o.MaskPosition
	}
	if !IsNil(o.CustomEmojiId) {
		toSerialize["custom_emoji_id"] = o.CustomEmojiId
	}
	if !IsNil(o.NeedsRepainting) {
		toSerialize["needs_repainting"] = o.NeedsRepainting
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	return toSerialize, nil
}

func (o *Sticker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_id",
		"file_unique_id",
		"type",
		"width",
		"height",
		"is_animated",
		"is_video",
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

	varSticker := _Sticker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSticker)

	if err != nil {
		return err
	}

	*o = Sticker(varSticker)

	return err
}

type NullableSticker struct {
	value *Sticker
	isSet bool
}

func (v NullableSticker) Get() *Sticker {
	return v.value
}

func (v *NullableSticker) Set(val *Sticker) {
	v.value = val
	v.isSet = true
}

func (v NullableSticker) IsSet() bool {
	return v.isSet
}

func (v *NullableSticker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSticker(val *Sticker) *NullableSticker {
	return &NullableSticker{value: val, isSet: true}
}

func (v NullableSticker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSticker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


