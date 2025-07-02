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
)

// checks if the BusinessIntro type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessIntro{}

// BusinessIntro Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	// *Optional*. Title text of the business intro
	Title *string `json:"title,omitempty"`
	// *Optional*. Message text of the business intro
	Message *string `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

// NewBusinessIntro instantiates a new BusinessIntro object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessIntro() *BusinessIntro {
	this := BusinessIntro{}
	return &this
}

// NewBusinessIntroWithDefaults instantiates a new BusinessIntro object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessIntroWithDefaults() *BusinessIntro {
	this := BusinessIntro{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BusinessIntro) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIntro) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BusinessIntro) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BusinessIntro) SetTitle(v string) {
	o.Title = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BusinessIntro) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIntro) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BusinessIntro) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BusinessIntro) SetMessage(v string) {
	o.Message = &v
}


// GetSticker returns the Sticker field value if set, zero value otherwise.
func (o *BusinessIntro) GetSticker() Sticker {
	if o == nil || IsNil(o.Sticker) {
		var ret Sticker
		return ret
	}
	return *o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIntro) GetStickerOk() (*Sticker, bool) {
	if o == nil || IsNil(o.Sticker) {
		return nil, false
	}
	return o.Sticker, true
}

// HasSticker returns a boolean if a field has been set.
func (o *BusinessIntro) HasSticker() bool {
	if o != nil && !IsNil(o.Sticker) {
		return true
	}

	return false
}

// SetSticker gets a reference to the given Sticker and assigns it to the Sticker field.
func (o *BusinessIntro) SetSticker(v Sticker) {
	o.Sticker = &v
}


func (o BusinessIntro) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessIntro) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Sticker) {
		toSerialize["sticker"] = o.Sticker
	}
	return toSerialize, nil
}

type NullableBusinessIntro struct {
	value *BusinessIntro
	isSet bool
}

func (v NullableBusinessIntro) Get() *BusinessIntro {
	return v.value
}

func (v *NullableBusinessIntro) Set(val *BusinessIntro) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessIntro) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessIntro) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessIntro(val *BusinessIntro) *NullableBusinessIntro {
	return &NullableBusinessIntro{value: val, isSet: true}
}

func (v NullableBusinessIntro) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessIntro) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


