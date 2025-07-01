/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the UniqueGiftBackdrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGiftBackdrop{}

// UniqueGiftBackdrop This object describes the backdrop of a unique gift.
type UniqueGiftBackdrop struct {
	// Name of the backdrop
	Name string `json:"name"`
	Colors UniqueGiftBackdropColors `json:"colors"`
	// The number of unique gifts that receive this backdrop for every 1000 gifts upgraded
	RarityPerMille int32 `json:"rarity_per_mille"`
}

type _UniqueGiftBackdrop UniqueGiftBackdrop

// NewUniqueGiftBackdrop instantiates a new UniqueGiftBackdrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGiftBackdrop(name string, colors UniqueGiftBackdropColors, rarityPerMille int32) *UniqueGiftBackdrop {
	this := UniqueGiftBackdrop{}
	this.Name = name
	this.Colors = colors
	this.RarityPerMille = rarityPerMille
	return &this
}

// NewUniqueGiftBackdropWithDefaults instantiates a new UniqueGiftBackdrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftBackdropWithDefaults() *UniqueGiftBackdrop {
	this := UniqueGiftBackdrop{}
	return &this
}

// GetName returns the Name field value
func (o *UniqueGiftBackdrop) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdrop) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UniqueGiftBackdrop) SetName(v string) {
	o.Name = v
}

// GetColors returns the Colors field value
func (o *UniqueGiftBackdrop) GetColors() UniqueGiftBackdropColors {
	if o == nil {
		var ret UniqueGiftBackdropColors
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdrop) GetColorsOk() (*UniqueGiftBackdropColors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Colors, true
}

// SetColors sets field value
func (o *UniqueGiftBackdrop) SetColors(v UniqueGiftBackdropColors) {
	o.Colors = v
}

// GetRarityPerMille returns the RarityPerMille field value
func (o *UniqueGiftBackdrop) GetRarityPerMille() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RarityPerMille
}

// GetRarityPerMilleOk returns a tuple with the RarityPerMille field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdrop) GetRarityPerMilleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RarityPerMille, true
}

// SetRarityPerMille sets field value
func (o *UniqueGiftBackdrop) SetRarityPerMille(v int32) {
	o.RarityPerMille = v
}

func (o UniqueGiftBackdrop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGiftBackdrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["colors"] = o.Colors
	toSerialize["rarity_per_mille"] = o.RarityPerMille
	return toSerialize, nil
}

func (o *UniqueGiftBackdrop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"colors",
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

	varUniqueGiftBackdrop := _UniqueGiftBackdrop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGiftBackdrop)

	if err != nil {
		return err
	}

	*o = UniqueGiftBackdrop(varUniqueGiftBackdrop)

	return err
}

type NullableUniqueGiftBackdrop struct {
	value *UniqueGiftBackdrop
	isSet bool
}

func (v NullableUniqueGiftBackdrop) Get() *UniqueGiftBackdrop {
	return v.value
}

func (v *NullableUniqueGiftBackdrop) Set(val *UniqueGiftBackdrop) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGiftBackdrop) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGiftBackdrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGiftBackdrop(val *UniqueGiftBackdrop) *NullableUniqueGiftBackdrop {
	return &NullableUniqueGiftBackdrop{value: val, isSet: true}
}

func (v NullableUniqueGiftBackdrop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGiftBackdrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


