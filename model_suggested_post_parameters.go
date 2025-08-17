/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-08-17T02:10:51.183543732Z[Etc/UTC]
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

// checks if the SuggestedPostParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostParameters{}

// SuggestedPostParameters Contains parameters of a post that is being suggested by the bot.
type SuggestedPostParameters struct {
	Price *SuggestedPostPrice `json:"price,omitempty"`
	// *Optional*. Proposed send date of the post. If specified, then the date must be between 300 second and 2678400 seconds (30 days) in the future. If the field is omitted, then the post can be published at any time within 30 days at the sole discretion of the user who approves it.
	SendDate *int32 `json:"send_date,omitempty"`
}

// NewSuggestedPostParameters instantiates a new SuggestedPostParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostParameters() *SuggestedPostParameters {
	this := SuggestedPostParameters{}
	return &this
}

// NewSuggestedPostParametersWithDefaults instantiates a new SuggestedPostParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostParametersWithDefaults() *SuggestedPostParameters {
	this := SuggestedPostParameters{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SuggestedPostParameters) GetPrice() SuggestedPostPrice {
	if o == nil || IsNil(o.Price) {
		var ret SuggestedPostPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostParameters) GetPriceOk() (*SuggestedPostPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SuggestedPostParameters) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given SuggestedPostPrice and assigns it to the Price field.
func (o *SuggestedPostParameters) SetPrice(v SuggestedPostPrice) {
	o.Price = &v
}


// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *SuggestedPostParameters) GetSendDate() int32 {
	if o == nil || IsNil(o.SendDate) {
		var ret int32
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostParameters) GetSendDateOk() (*int32, bool) {
	if o == nil || IsNil(o.SendDate) {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *SuggestedPostParameters) HasSendDate() bool {
	if o != nil && !IsNil(o.SendDate) {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given int32 and assigns it to the SendDate field.
func (o *SuggestedPostParameters) SetSendDate(v int32) {
	o.SendDate = &v
}


func (o SuggestedPostParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.SendDate) {
		toSerialize["send_date"] = o.SendDate
	}
	return toSerialize, nil
}

type NullableSuggestedPostParameters struct {
	value *SuggestedPostParameters
	isSet bool
}

func (v NullableSuggestedPostParameters) Get() *SuggestedPostParameters {
	return v.value
}

func (v *NullableSuggestedPostParameters) Set(val *SuggestedPostParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostParameters(val *SuggestedPostParameters) *NullableSuggestedPostParameters {
	return &NullableSuggestedPostParameters{value: val, isSet: true}
}

func (v NullableSuggestedPostParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


