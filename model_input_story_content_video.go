/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the InputStoryContentVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputStoryContentVideo{}

// InputStoryContentVideo Describes a video to post as a story.
type InputStoryContentVideo struct {
	// Type of the content, must be *video*
	Type string `json:"type"`
	// The video to post as a story. The video must be of the size 720x1280, streamable, encoded with H.265 codec, with key frames added each second in the MPEG4 format, and must not exceed 30 MB. The video can't be reused and can only be uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the video was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Video string `json:"video"`
	// *Optional*. Precise duration of the video in seconds; 0-60
	Duration *float32 `json:"duration,omitempty"`
	// *Optional*. Timestamp in seconds of the frame that will be used as the static cover for the story. Defaults to 0.0.
	CoverFrameTimestamp *float32 `json:"cover_frame_timestamp,omitempty"`
	// *Optional*. Pass *True* if the video has no sound
	IsAnimation *bool `json:"is_animation,omitempty"`
}

type _InputStoryContentVideo InputStoryContentVideo

// NewInputStoryContentVideo instantiates a new InputStoryContentVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputStoryContentVideo(type_ string, video string) *InputStoryContentVideo {
	this := InputStoryContentVideo{}
	this.Type = type_
	this.Video = video
	return &this
}

// NewInputStoryContentVideoWithDefaults instantiates a new InputStoryContentVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputStoryContentVideoWithDefaults() *InputStoryContentVideo {
	this := InputStoryContentVideo{}
	var type_ string = "video"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputStoryContentVideo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputStoryContentVideo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputStoryContentVideo) SetType(v string) {
	o.Type = v
}

// GetVideo returns the Video field value
func (o *InputStoryContentVideo) GetVideo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Video
}

// GetVideoOk returns a tuple with the Video field value
// and a boolean to check if the value has been set.
func (o *InputStoryContentVideo) GetVideoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Video, true
}

// SetVideo sets field value
func (o *InputStoryContentVideo) SetVideo(v string) {
	o.Video = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InputStoryContentVideo) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputStoryContentVideo) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InputStoryContentVideo) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *InputStoryContentVideo) SetDuration(v float32) {
	o.Duration = &v
}


// GetCoverFrameTimestamp returns the CoverFrameTimestamp field value if set, zero value otherwise.
func (o *InputStoryContentVideo) GetCoverFrameTimestamp() float32 {
	if o == nil || IsNil(o.CoverFrameTimestamp) {
		var ret float32
		return ret
	}
	return *o.CoverFrameTimestamp
}

// GetCoverFrameTimestampOk returns a tuple with the CoverFrameTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputStoryContentVideo) GetCoverFrameTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.CoverFrameTimestamp) {
		return nil, false
	}
	return o.CoverFrameTimestamp, true
}

// HasCoverFrameTimestamp returns a boolean if a field has been set.
func (o *InputStoryContentVideo) HasCoverFrameTimestamp() bool {
	if o != nil && !IsNil(o.CoverFrameTimestamp) {
		return true
	}

	return false
}

// SetCoverFrameTimestamp gets a reference to the given float32 and assigns it to the CoverFrameTimestamp field.
func (o *InputStoryContentVideo) SetCoverFrameTimestamp(v float32) {
	o.CoverFrameTimestamp = &v
}


// GetIsAnimation returns the IsAnimation field value if set, zero value otherwise.
func (o *InputStoryContentVideo) GetIsAnimation() bool {
	if o == nil || IsNil(o.IsAnimation) {
		var ret bool
		return ret
	}
	return *o.IsAnimation
}

// GetIsAnimationOk returns a tuple with the IsAnimation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputStoryContentVideo) GetIsAnimationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnimation) {
		return nil, false
	}
	return o.IsAnimation, true
}

// HasIsAnimation returns a boolean if a field has been set.
func (o *InputStoryContentVideo) HasIsAnimation() bool {
	if o != nil && !IsNil(o.IsAnimation) {
		return true
	}

	return false
}

// SetIsAnimation gets a reference to the given bool and assigns it to the IsAnimation field.
func (o *InputStoryContentVideo) SetIsAnimation(v bool) {
	o.IsAnimation = &v
}


func (o InputStoryContentVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputStoryContentVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["video"] = o.Video
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.CoverFrameTimestamp) {
		toSerialize["cover_frame_timestamp"] = o.CoverFrameTimestamp
	}
	if !IsNil(o.IsAnimation) {
		toSerialize["is_animation"] = o.IsAnimation
	}
	return toSerialize, nil
}

func (o *InputStoryContentVideo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"video",
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

	varInputStoryContentVideo := _InputStoryContentVideo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputStoryContentVideo)

	if err != nil {
		return err
	}

	*o = InputStoryContentVideo(varInputStoryContentVideo)

	return err
}

type NullableInputStoryContentVideo struct {
	value *InputStoryContentVideo
	isSet bool
}

func (v NullableInputStoryContentVideo) Get() *InputStoryContentVideo {
	return v.value
}

func (v *NullableInputStoryContentVideo) Set(val *InputStoryContentVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableInputStoryContentVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableInputStoryContentVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputStoryContentVideo(val *InputStoryContentVideo) *NullableInputStoryContentVideo {
	return &NullableInputStoryContentVideo{value: val, isSet: true}
}

func (v NullableInputStoryContentVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputStoryContentVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


