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

// checks if the Animation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Animation{}

// Animation This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width as defined by the sender
	Width int32 `json:"width"`
	// Video height as defined by the sender
	Height int32 `json:"height"`
	// Duration of the video in seconds as defined by the sender
	Duration int32 `json:"duration"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// *Optional*. Original animation filename as defined by the sender
	FileName *string `json:"file_name,omitempty"`
	// *Optional*. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int32 `json:"file_size,omitempty"`
}

type _Animation Animation

// NewAnimation instantiates a new Animation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimation(fileId string, fileUniqueId string, width int32, height int32, duration int32) *Animation {
	this := Animation{}
	this.FileId = fileId
	this.FileUniqueId = fileUniqueId
	this.Width = width
	this.Height = height
	this.Duration = duration
	return &this
}

// NewAnimationWithDefaults instantiates a new Animation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimationWithDefaults() *Animation {
	this := Animation{}
	return &this
}

// GetFileId returns the FileId field value
func (o *Animation) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *Animation) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *Animation) SetFileId(v string) {
	o.FileId = v
}

// GetFileUniqueId returns the FileUniqueId field value
func (o *Animation) GetFileUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUniqueId
}

// GetFileUniqueIdOk returns a tuple with the FileUniqueId field value
// and a boolean to check if the value has been set.
func (o *Animation) GetFileUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUniqueId, true
}

// SetFileUniqueId sets field value
func (o *Animation) SetFileUniqueId(v string) {
	o.FileUniqueId = v
}

// GetWidth returns the Width field value
func (o *Animation) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Animation) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Animation) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *Animation) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Animation) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Animation) SetHeight(v int32) {
	o.Height = v
}

// GetDuration returns the Duration field value
func (o *Animation) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *Animation) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *Animation) SetDuration(v int32) {
	o.Duration = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *Animation) GetThumbnail() PhotoSize {
	if o == nil || IsNil(o.Thumbnail) {
		var ret PhotoSize
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetThumbnailOk() (*PhotoSize, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *Animation) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given PhotoSize and assigns it to the Thumbnail field.
func (o *Animation) SetThumbnail(v PhotoSize) {
	o.Thumbnail = &v
}


// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *Animation) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *Animation) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *Animation) SetFileName(v string) {
	o.FileName = &v
}


// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *Animation) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *Animation) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *Animation) SetMimeType(v string) {
	o.MimeType = &v
}


// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *Animation) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *Animation) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *Animation) SetFileSize(v int32) {
	o.FileSize = &v
}


func (o Animation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Animation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_id"] = o.FileId
	toSerialize["file_unique_id"] = o.FileUniqueId
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["duration"] = o.Duration
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	return toSerialize, nil
}

func (o *Animation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_id",
		"file_unique_id",
		"width",
		"height",
		"duration",
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

	varAnimation := _Animation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnimation)

	if err != nil {
		return err
	}

	*o = Animation(varAnimation)

	return err
}

type NullableAnimation struct {
	value *Animation
	isSet bool
}

func (v NullableAnimation) Get() *Animation {
	return v.value
}

func (v *NullableAnimation) Set(val *Animation) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimation) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimation(val *Animation) *NullableAnimation {
	return &NullableAnimation{value: val, isSet: true}
}

func (v NullableAnimation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


