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

// checks if the UniqueGiftSymbol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGiftSymbol{}

// UniqueGiftSymbol This object describes the symbol shown on the pattern of a unique gift.
type UniqueGiftSymbol struct {
	// Name of the symbol
	Name string `json:"name"`
	Sticker Sticker `json:"sticker"`
	// The number of unique gifts that receive this model for every 1000 gifts upgraded
	RarityPerMille int32 `json:"rarity_per_mille"`
}

type _UniqueGiftSymbol UniqueGiftSymbol

// NewUniqueGiftSymbol instantiates a new UniqueGiftSymbol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGiftSymbol(name string, sticker Sticker, rarityPerMille int32) *UniqueGiftSymbol {
	this := UniqueGiftSymbol{}
	this.Name = name
	this.Sticker = sticker
	this.RarityPerMille = rarityPerMille
	return &this
}

// NewUniqueGiftSymbolWithDefaults instantiates a new UniqueGiftSymbol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftSymbolWithDefaults() *UniqueGiftSymbol {
	this := UniqueGiftSymbol{}
	return &this
}

// GetName returns the Name field value
func (o *UniqueGiftSymbol) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftSymbol) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UniqueGiftSymbol) SetName(v string) {
	o.Name = v
}

// GetSticker returns the Sticker field value
func (o *UniqueGiftSymbol) GetSticker() Sticker {
	if o == nil {
		var ret Sticker
		return ret
	}

	return o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftSymbol) GetStickerOk() (*Sticker, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticker, true
}

// SetSticker sets field value
func (o *UniqueGiftSymbol) SetSticker(v Sticker) {
	o.Sticker = v
}

// GetRarityPerMille returns the RarityPerMille field value
func (o *UniqueGiftSymbol) GetRarityPerMille() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RarityPerMille
}

// GetRarityPerMilleOk returns a tuple with the RarityPerMille field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftSymbol) GetRarityPerMilleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RarityPerMille, true
}

// SetRarityPerMille sets field value
func (o *UniqueGiftSymbol) SetRarityPerMille(v int32) {
	o.RarityPerMille = v
}

func (o UniqueGiftSymbol) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGiftSymbol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["sticker"] = o.Sticker
	toSerialize["rarity_per_mille"] = o.RarityPerMille
	return toSerialize, nil
}

func (o *UniqueGiftSymbol) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sticker",
		"rarity_per_mille",
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

	varUniqueGiftSymbol := _UniqueGiftSymbol{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGiftSymbol)

	if err != nil {
		return err
	}

	*o = UniqueGiftSymbol(varUniqueGiftSymbol)

	return err
}

type NullableUniqueGiftSymbol struct {
	value *UniqueGiftSymbol
	isSet bool
}

func (v NullableUniqueGiftSymbol) Get() *UniqueGiftSymbol {
	return v.value
}

func (v *NullableUniqueGiftSymbol) Set(val *UniqueGiftSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGiftSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGiftSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGiftSymbol(val *UniqueGiftSymbol) *NullableUniqueGiftSymbol {
	return &NullableUniqueGiftSymbol{value: val, isSet: true}
}

func (v NullableUniqueGiftSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGiftSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


