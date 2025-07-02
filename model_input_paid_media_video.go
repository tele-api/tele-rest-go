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

// checks if the InputPaidMediaVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputPaidMediaVideo{}

// InputPaidMediaVideo The paid media to send is a video.
type InputPaidMediaVideo struct {
	// Type of the media, must be *video*
	Type string `json:"type"`
	// File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\<file\\_attach\\_name\\>” to upload a new one using multipart/form-data under \\<file\\_attach\\_name\\> name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the thumbnail was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Thumbnail *string `json:"thumbnail,omitempty"`
	// *Optional*. Cover for the video in the message. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\<file\\_attach\\_name\\>” to upload a new one using multipart/form-data under \\<file\\_attach\\_name\\> name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Cover *string `json:"cover,omitempty"`
	// *Optional*. Start timestamp for the video in the message
	StartTimestamp *int32 `json:"start_timestamp,omitempty"`
	// *Optional*. Video width
	Width *int32 `json:"width,omitempty"`
	// *Optional*. Video height
	Height *int32 `json:"height,omitempty"`
	// *Optional*. Video duration in seconds
	Duration *int32 `json:"duration,omitempty"`
	// *Optional*. Pass *True* if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
}

type _InputPaidMediaVideo InputPaidMediaVideo

// NewInputPaidMediaVideo instantiates a new InputPaidMediaVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputPaidMediaVideo(type_ string, media string) *InputPaidMediaVideo {
	this := InputPaidMediaVideo{}
	this.Type = type_
	this.Media = media
	return &this
}

// NewInputPaidMediaVideoWithDefaults instantiates a new InputPaidMediaVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputPaidMediaVideoWithDefaults() *InputPaidMediaVideo {
	this := InputPaidMediaVideo{}
	var type_ string = "video"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputPaidMediaVideo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputPaidMediaVideo) SetType(v string) {
	o.Type = v
}

// GetMedia returns the Media field value
func (o *InputPaidMediaVideo) GetMedia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Media
}

// GetMediaOk returns a tuple with the Media field value
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetMediaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Media, true
}

// SetMedia sets field value
func (o *InputPaidMediaVideo) SetMedia(v string) {
	o.Media = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *InputPaidMediaVideo) SetThumbnail(v string) {
	o.Thumbnail = &v
}


// GetCover returns the Cover field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetCover() string {
	if o == nil || IsNil(o.Cover) {
		var ret string
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasCover() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given string and assigns it to the Cover field.
func (o *InputPaidMediaVideo) SetCover(v string) {
	o.Cover = &v
}


// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetStartTimestamp() int32 {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret int32
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given int32 and assigns it to the StartTimestamp field.
func (o *InputPaidMediaVideo) SetStartTimestamp(v int32) {
	o.StartTimestamp = &v
}


// GetWidth returns the Width field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *InputPaidMediaVideo) SetWidth(v int32) {
	o.Width = &v
}


// GetHeight returns the Height field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *InputPaidMediaVideo) SetHeight(v int32) {
	o.Height = &v
}


// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InputPaidMediaVideo) SetDuration(v int32) {
	o.Duration = &v
}


// GetSupportsStreaming returns the SupportsStreaming field value if set, zero value otherwise.
func (o *InputPaidMediaVideo) GetSupportsStreaming() bool {
	if o == nil || IsNil(o.SupportsStreaming) {
		var ret bool
		return ret
	}
	return *o.SupportsStreaming
}

// GetSupportsStreamingOk returns a tuple with the SupportsStreaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPaidMediaVideo) GetSupportsStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsStreaming) {
		return nil, false
	}
	return o.SupportsStreaming, true
}

// HasSupportsStreaming returns a boolean if a field has been set.
func (o *InputPaidMediaVideo) HasSupportsStreaming() bool {
	if o != nil && !IsNil(o.SupportsStreaming) {
		return true
	}

	return false
}

// SetSupportsStreaming gets a reference to the given bool and assigns it to the SupportsStreaming field.
func (o *InputPaidMediaVideo) SetSupportsStreaming(v bool) {
	o.SupportsStreaming = &v
}


func (o InputPaidMediaVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputPaidMediaVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["media"] = o.Media
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Cover) {
		toSerialize["cover"] = o.Cover
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["start_timestamp"] = o.StartTimestamp
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.SupportsStreaming) {
		toSerialize["supports_streaming"] = o.SupportsStreaming
	}
	return toSerialize, nil
}

func (o *InputPaidMediaVideo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"media",
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

	varInputPaidMediaVideo := _InputPaidMediaVideo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputPaidMediaVideo)

	if err != nil {
		return err
	}

	*o = InputPaidMediaVideo(varInputPaidMediaVideo)

	return err
}

type NullableInputPaidMediaVideo struct {
	value *InputPaidMediaVideo
	isSet bool
}

func (v NullableInputPaidMediaVideo) Get() *InputPaidMediaVideo {
	return v.value
}

func (v *NullableInputPaidMediaVideo) Set(val *InputPaidMediaVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPaidMediaVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPaidMediaVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPaidMediaVideo(val *InputPaidMediaVideo) *NullableInputPaidMediaVideo {
	return &NullableInputPaidMediaVideo{value: val, isSet: true}
}

func (v NullableInputPaidMediaVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPaidMediaVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


