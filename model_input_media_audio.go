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

// checks if the InputMediaAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputMediaAudio{}

// InputMediaAudio Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type of the result, must be *audio*
	Type string `json:"type"`
	// File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\<file\\_attach\\_name\\>” to upload a new one using multipart/form-data under \\<file\\_attach\\_name\\> name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the thumbnail was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Thumbnail *string `json:"thumbnail,omitempty"`
	// *Optional*. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Duration of the audio in seconds
	Duration *int32 `json:"duration,omitempty"`
	// *Optional*. Performer of the audio
	Performer *string `json:"performer,omitempty"`
	// *Optional*. Title of the audio
	Title *string `json:"title,omitempty"`
}

type _InputMediaAudio InputMediaAudio

// NewInputMediaAudio instantiates a new InputMediaAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputMediaAudio(type_ string, media string) *InputMediaAudio {
	this := InputMediaAudio{}
	this.Type = type_
	this.Media = media
	return &this
}

// NewInputMediaAudioWithDefaults instantiates a new InputMediaAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputMediaAudioWithDefaults() *InputMediaAudio {
	this := InputMediaAudio{}
	var type_ string = "audio"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputMediaAudio) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputMediaAudio) SetType(v string) {
	o.Type = v
}

// GetMedia returns the Media field value
func (o *InputMediaAudio) GetMedia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Media
}

// GetMediaOk returns a tuple with the Media field value
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetMediaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Media, true
}

// SetMedia sets field value
func (o *InputMediaAudio) SetMedia(v string) {
	o.Media = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *InputMediaAudio) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *InputMediaAudio) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *InputMediaAudio) SetThumbnail(v string) {
	o.Thumbnail = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InputMediaAudio) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InputMediaAudio) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InputMediaAudio) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InputMediaAudio) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InputMediaAudio) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InputMediaAudio) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InputMediaAudio) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InputMediaAudio) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InputMediaAudio) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InputMediaAudio) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InputMediaAudio) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InputMediaAudio) SetDuration(v int32) {
	o.Duration = &v
}


// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *InputMediaAudio) GetPerformer() string {
	if o == nil || IsNil(o.Performer) {
		var ret string
		return ret
	}
	return *o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetPerformerOk() (*string, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *InputMediaAudio) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given string and assigns it to the Performer field.
func (o *InputMediaAudio) SetPerformer(v string) {
	o.Performer = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InputMediaAudio) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaAudio) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InputMediaAudio) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InputMediaAudio) SetTitle(v string) {
	o.Title = &v
}


func (o InputMediaAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputMediaAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["media"] = o.Media
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.CaptionEntities) {
		toSerialize["caption_entities"] = o.CaptionEntities
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *InputMediaAudio) UnmarshalJSON(data []byte) (err error) {
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

	varInputMediaAudio := _InputMediaAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputMediaAudio)

	if err != nil {
		return err
	}

	*o = InputMediaAudio(varInputMediaAudio)

	return err
}

type NullableInputMediaAudio struct {
	value *InputMediaAudio
	isSet bool
}

func (v NullableInputMediaAudio) Get() *InputMediaAudio {
	return v.value
}

func (v *NullableInputMediaAudio) Set(val *InputMediaAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMediaAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMediaAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMediaAudio(val *InputMediaAudio) *NullableInputMediaAudio {
	return &NullableInputMediaAudio{value: val, isSet: true}
}

func (v NullableInputMediaAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMediaAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


