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

// checks if the InputMediaDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputMediaDocument{}

// InputMediaDocument Represents a general file to be sent.
type InputMediaDocument struct {
	// Type of the result, must be *document*
	Type string `json:"type"`
	// File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\<file\\_attach\\_name\\>” to upload a new one using multipart/form-data under \\<file\\_attach\\_name\\> name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the thumbnail was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Thumbnail *string `json:"thumbnail,omitempty"`
	// *Optional*. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always *True*, if the document is sent as part of an album.
	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
}

type _InputMediaDocument InputMediaDocument

// NewInputMediaDocument instantiates a new InputMediaDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputMediaDocument(type_ string, media string) *InputMediaDocument {
	this := InputMediaDocument{}
	this.Type = type_
	this.Media = media
	return &this
}

// NewInputMediaDocumentWithDefaults instantiates a new InputMediaDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputMediaDocumentWithDefaults() *InputMediaDocument {
	this := InputMediaDocument{}
	var type_ string = "document"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputMediaDocument) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputMediaDocument) SetType(v string) {
	o.Type = v
}

// GetMedia returns the Media field value
func (o *InputMediaDocument) GetMedia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Media
}

// GetMediaOk returns a tuple with the Media field value
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetMediaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Media, true
}

// SetMedia sets field value
func (o *InputMediaDocument) SetMedia(v string) {
	o.Media = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *InputMediaDocument) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *InputMediaDocument) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *InputMediaDocument) SetThumbnail(v string) {
	o.Thumbnail = &v
}


// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *InputMediaDocument) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *InputMediaDocument) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *InputMediaDocument) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InputMediaDocument) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InputMediaDocument) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InputMediaDocument) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *InputMediaDocument) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *InputMediaDocument) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *InputMediaDocument) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetDisableContentTypeDetection returns the DisableContentTypeDetection field value if set, zero value otherwise.
func (o *InputMediaDocument) GetDisableContentTypeDetection() bool {
	if o == nil || IsNil(o.DisableContentTypeDetection) {
		var ret bool
		return ret
	}
	return *o.DisableContentTypeDetection
}

// GetDisableContentTypeDetectionOk returns a tuple with the DisableContentTypeDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputMediaDocument) GetDisableContentTypeDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableContentTypeDetection) {
		return nil, false
	}
	return o.DisableContentTypeDetection, true
}

// HasDisableContentTypeDetection returns a boolean if a field has been set.
func (o *InputMediaDocument) HasDisableContentTypeDetection() bool {
	if o != nil && !IsNil(o.DisableContentTypeDetection) {
		return true
	}

	return false
}

// SetDisableContentTypeDetection gets a reference to the given bool and assigns it to the DisableContentTypeDetection field.
func (o *InputMediaDocument) SetDisableContentTypeDetection(v bool) {
	o.DisableContentTypeDetection = &v
}


func (o InputMediaDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputMediaDocument) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DisableContentTypeDetection) {
		toSerialize["disable_content_type_detection"] = o.DisableContentTypeDetection
	}
	return toSerialize, nil
}

func (o *InputMediaDocument) UnmarshalJSON(data []byte) (err error) {
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

	varInputMediaDocument := _InputMediaDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputMediaDocument)

	if err != nil {
		return err
	}

	*o = InputMediaDocument(varInputMediaDocument)

	return err
}

type NullableInputMediaDocument struct {
	value *InputMediaDocument
	isSet bool
}

func (v NullableInputMediaDocument) Get() *InputMediaDocument {
	return v.value
}

func (v *NullableInputMediaDocument) Set(val *InputMediaDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMediaDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMediaDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMediaDocument(val *InputMediaDocument) *NullableInputMediaDocument {
	return &NullableInputMediaDocument{value: val, isSet: true}
}

func (v NullableInputMediaDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMediaDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


