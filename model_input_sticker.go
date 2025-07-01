/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the InputSticker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputSticker{}

// InputSticker This object describes a sticker to be added to a sticker set.
type InputSticker struct {
	// The added sticker. Pass a *file\\_id* as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass “attach://\\<file\\_attach\\_name\\>” to upload a new file using multipart/form-data under \\<file\\_attach\\_name\\> name. Animated and video stickers can't be uploaded via HTTP URL. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Sticker string `json:"sticker"`
	// Format of the added sticker, must be one of “static” for a **.WEBP** or **.PNG** image, “animated” for a **.TGS** animation, “video” for a **.WEBM** video
	Format string `json:"format"`
	// List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// *Optional*. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom\\_emoji” stickers only.
	Keywords []string `json:"keywords,omitempty"`
}

type _InputSticker InputSticker

// NewInputSticker instantiates a new InputSticker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSticker(sticker string, format string, emojiList []string) *InputSticker {
	this := InputSticker{}
	this.Sticker = sticker
	this.Format = format
	this.EmojiList = emojiList
	return &this
}

// NewInputStickerWithDefaults instantiates a new InputSticker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputStickerWithDefaults() *InputSticker {
	this := InputSticker{}
	return &this
}

// GetSticker returns the Sticker field value
func (o *InputSticker) GetSticker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value
// and a boolean to check if the value has been set.
func (o *InputSticker) GetStickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticker, true
}

// SetSticker sets field value
func (o *InputSticker) SetSticker(v string) {
	o.Sticker = v
}

// GetFormat returns the Format field value
func (o *InputSticker) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *InputSticker) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *InputSticker) SetFormat(v string) {
	o.Format = v
}

// GetEmojiList returns the EmojiList field value
func (o *InputSticker) GetEmojiList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmojiList
}

// GetEmojiListOk returns a tuple with the EmojiList field value
// and a boolean to check if the value has been set.
func (o *InputSticker) GetEmojiListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmojiList, true
}

// SetEmojiList sets field value
func (o *InputSticker) SetEmojiList(v []string) {
	o.EmojiList = v
}

// GetMaskPosition returns the MaskPosition field value if set, zero value otherwise.
func (o *InputSticker) GetMaskPosition() MaskPosition {
	if o == nil || IsNil(o.MaskPosition) {
		var ret MaskPosition
		return ret
	}
	return *o.MaskPosition
}

// GetMaskPositionOk returns a tuple with the MaskPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSticker) GetMaskPositionOk() (*MaskPosition, bool) {
	if o == nil || IsNil(o.MaskPosition) {
		return nil, false
	}
	return o.MaskPosition, true
}

// HasMaskPosition returns a boolean if a field has been set.
func (o *InputSticker) HasMaskPosition() bool {
	if o != nil && !IsNil(o.MaskPosition) {
		return true
	}

	return false
}

// SetMaskPosition gets a reference to the given MaskPosition and assigns it to the MaskPosition field.
func (o *InputSticker) SetMaskPosition(v MaskPosition) {
	o.MaskPosition = &v
}


// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *InputSticker) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSticker) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *InputSticker) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *InputSticker) SetKeywords(v []string) {
	o.Keywords = v
}


func (o InputSticker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputSticker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sticker"] = o.Sticker
	toSerialize["format"] = o.Format
	toSerialize["emoji_list"] = o.EmojiList
	if !IsNil(o.MaskPosition) {
		toSerialize["mask_position"] = o.MaskPosition
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	return toSerialize, nil
}

func (o *InputSticker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sticker",
		"format",
		"emoji_list",
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

	varInputSticker := _InputSticker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputSticker)

	if err != nil {
		return err
	}

	*o = InputSticker(varInputSticker)

	return err
}

type NullableInputSticker struct {
	value *InputSticker
	isSet bool
}

func (v NullableInputSticker) Get() *InputSticker {
	return v.value
}

func (v *NullableInputSticker) Set(val *InputSticker) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSticker) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSticker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSticker(val *InputSticker) *NullableInputSticker {
	return &NullableInputSticker{value: val, isSet: true}
}

func (v NullableInputSticker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSticker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


