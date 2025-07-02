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

// checks if the UniqueGiftBackdropColors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGiftBackdropColors{}

// UniqueGiftBackdropColors This object describes the colors of the backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	// The color in the center of the backdrop in RGB format
	CenterColor int32 `json:"center_color"`
	// The color on the edges of the backdrop in RGB format
	EdgeColor int32 `json:"edge_color"`
	// The color to be applied to the symbol in RGB format
	SymbolColor int32 `json:"symbol_color"`
	// The color for the text on the backdrop in RGB format
	TextColor int32 `json:"text_color"`
}

type _UniqueGiftBackdropColors UniqueGiftBackdropColors

// NewUniqueGiftBackdropColors instantiates a new UniqueGiftBackdropColors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGiftBackdropColors(centerColor int32, edgeColor int32, symbolColor int32, textColor int32) *UniqueGiftBackdropColors {
	this := UniqueGiftBackdropColors{}
	this.CenterColor = centerColor
	this.EdgeColor = edgeColor
	this.SymbolColor = symbolColor
	this.TextColor = textColor
	return &this
}

// NewUniqueGiftBackdropColorsWithDefaults instantiates a new UniqueGiftBackdropColors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftBackdropColorsWithDefaults() *UniqueGiftBackdropColors {
	this := UniqueGiftBackdropColors{}
	return &this
}

// GetCenterColor returns the CenterColor field value
func (o *UniqueGiftBackdropColors) GetCenterColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CenterColor
}

// GetCenterColorOk returns a tuple with the CenterColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdropColors) GetCenterColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CenterColor, true
}

// SetCenterColor sets field value
func (o *UniqueGiftBackdropColors) SetCenterColor(v int32) {
	o.CenterColor = v
}

// GetEdgeColor returns the EdgeColor field value
func (o *UniqueGiftBackdropColors) GetEdgeColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EdgeColor
}

// GetEdgeColorOk returns a tuple with the EdgeColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdropColors) GetEdgeColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeColor, true
}

// SetEdgeColor sets field value
func (o *UniqueGiftBackdropColors) SetEdgeColor(v int32) {
	o.EdgeColor = v
}

// GetSymbolColor returns the SymbolColor field value
func (o *UniqueGiftBackdropColors) GetSymbolColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SymbolColor
}

// GetSymbolColorOk returns a tuple with the SymbolColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdropColors) GetSymbolColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SymbolColor, true
}

// SetSymbolColor sets field value
func (o *UniqueGiftBackdropColors) SetSymbolColor(v int32) {
	o.SymbolColor = v
}

// GetTextColor returns the TextColor field value
func (o *UniqueGiftBackdropColors) GetTextColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftBackdropColors) GetTextColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TextColor, true
}

// SetTextColor sets field value
func (o *UniqueGiftBackdropColors) SetTextColor(v int32) {
	o.TextColor = v
}

func (o UniqueGiftBackdropColors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGiftBackdropColors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["center_color"] = o.CenterColor
	toSerialize["edge_color"] = o.EdgeColor
	toSerialize["symbol_color"] = o.SymbolColor
	toSerialize["text_color"] = o.TextColor
	return toSerialize, nil
}

func (o *UniqueGiftBackdropColors) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"center_color",
		"edge_color",
		"symbol_color",
		"text_color",
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

	varUniqueGiftBackdropColors := _UniqueGiftBackdropColors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGiftBackdropColors)

	if err != nil {
		return err
	}

	*o = UniqueGiftBackdropColors(varUniqueGiftBackdropColors)

	return err
}

type NullableUniqueGiftBackdropColors struct {
	value *UniqueGiftBackdropColors
	isSet bool
}

func (v NullableUniqueGiftBackdropColors) Get() *UniqueGiftBackdropColors {
	return v.value
}

func (v *NullableUniqueGiftBackdropColors) Set(val *UniqueGiftBackdropColors) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGiftBackdropColors) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGiftBackdropColors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGiftBackdropColors(val *UniqueGiftBackdropColors) *NullableUniqueGiftBackdropColors {
	return &NullableUniqueGiftBackdropColors{value: val, isSet: true}
}

func (v NullableUniqueGiftBackdropColors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGiftBackdropColors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


