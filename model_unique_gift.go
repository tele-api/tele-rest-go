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

// checks if the UniqueGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGift{}

// UniqueGift This object describes a unique gift that was upgraded from a regular gift.
type UniqueGift struct {
	// Human-readable name of the regular gift from which this unique gift was upgraded
	BaseName string `json:"base_name"`
	// Unique name of the gift. This name can be used in `https://t.me/nft/...` links and story areas
	Name string `json:"name"`
	// Unique number of the upgraded gift among gifts upgraded from the same regular gift
	Number int32 `json:"number"`
	Model UniqueGiftModel `json:"model"`
	Symbol UniqueGiftSymbol `json:"symbol"`
	Backdrop UniqueGiftBackdrop `json:"backdrop"`
}

type _UniqueGift UniqueGift

// NewUniqueGift instantiates a new UniqueGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGift(baseName string, name string, number int32, model UniqueGiftModel, symbol UniqueGiftSymbol, backdrop UniqueGiftBackdrop) *UniqueGift {
	this := UniqueGift{}
	this.BaseName = baseName
	this.Name = name
	this.Number = number
	this.Model = model
	this.Symbol = symbol
	this.Backdrop = backdrop
	return &this
}

// NewUniqueGiftWithDefaults instantiates a new UniqueGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftWithDefaults() *UniqueGift {
	this := UniqueGift{}
	return &this
}

// GetBaseName returns the BaseName field value
func (o *UniqueGift) GetBaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseName
}

// GetBaseNameOk returns a tuple with the BaseName field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetBaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseName, true
}

// SetBaseName sets field value
func (o *UniqueGift) SetBaseName(v string) {
	o.BaseName = v
}

// GetName returns the Name field value
func (o *UniqueGift) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UniqueGift) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *UniqueGift) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *UniqueGift) SetNumber(v int32) {
	o.Number = v
}

// GetModel returns the Model field value
func (o *UniqueGift) GetModel() UniqueGiftModel {
	if o == nil {
		var ret UniqueGiftModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetModelOk() (*UniqueGiftModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *UniqueGift) SetModel(v UniqueGiftModel) {
	o.Model = v
}

// GetSymbol returns the Symbol field value
func (o *UniqueGift) GetSymbol() UniqueGiftSymbol {
	if o == nil {
		var ret UniqueGiftSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetSymbolOk() (*UniqueGiftSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *UniqueGift) SetSymbol(v UniqueGiftSymbol) {
	o.Symbol = v
}

// GetBackdrop returns the Backdrop field value
func (o *UniqueGift) GetBackdrop() UniqueGiftBackdrop {
	if o == nil {
		var ret UniqueGiftBackdrop
		return ret
	}

	return o.Backdrop
}

// GetBackdropOk returns a tuple with the Backdrop field value
// and a boolean to check if the value has been set.
func (o *UniqueGift) GetBackdropOk() (*UniqueGiftBackdrop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backdrop, true
}

// SetBackdrop sets field value
func (o *UniqueGift) SetBackdrop(v UniqueGiftBackdrop) {
	o.Backdrop = v
}

func (o UniqueGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_name"] = o.BaseName
	toSerialize["name"] = o.Name
	toSerialize["number"] = o.Number
	toSerialize["model"] = o.Model
	toSerialize["symbol"] = o.Symbol
	toSerialize["backdrop"] = o.Backdrop
	return toSerialize, nil
}

func (o *UniqueGift) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_name",
		"name",
		"number",
		"model",
		"symbol",
		"backdrop",
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

	varUniqueGift := _UniqueGift{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGift)

	if err != nil {
		return err
	}

	*o = UniqueGift(varUniqueGift)

	return err
}

type NullableUniqueGift struct {
	value *UniqueGift
	isSet bool
}

func (v NullableUniqueGift) Get() *UniqueGift {
	return v.value
}

func (v *NullableUniqueGift) Set(val *UniqueGift) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGift) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGift(val *UniqueGift) *NullableUniqueGift {
	return &NullableUniqueGift{value: val, isSet: true}
}

func (v NullableUniqueGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


