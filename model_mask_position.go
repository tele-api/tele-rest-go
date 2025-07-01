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

// checks if the MaskPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskPosition{}

// MaskPosition This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float32 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}

type _MaskPosition MaskPosition

// NewMaskPosition instantiates a new MaskPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskPosition(point string, xShift float32, yShift float32, scale float32) *MaskPosition {
	this := MaskPosition{}
	this.Point = point
	this.XShift = xShift
	this.YShift = yShift
	this.Scale = scale
	return &this
}

// NewMaskPositionWithDefaults instantiates a new MaskPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskPositionWithDefaults() *MaskPosition {
	this := MaskPosition{}
	return &this
}

// GetPoint returns the Point field value
func (o *MaskPosition) GetPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *MaskPosition) GetPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *MaskPosition) SetPoint(v string) {
	o.Point = v
}

// GetXShift returns the XShift field value
func (o *MaskPosition) GetXShift() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.XShift
}

// GetXShiftOk returns a tuple with the XShift field value
// and a boolean to check if the value has been set.
func (o *MaskPosition) GetXShiftOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XShift, true
}

// SetXShift sets field value
func (o *MaskPosition) SetXShift(v float32) {
	o.XShift = v
}

// GetYShift returns the YShift field value
func (o *MaskPosition) GetYShift() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.YShift
}

// GetYShiftOk returns a tuple with the YShift field value
// and a boolean to check if the value has been set.
func (o *MaskPosition) GetYShiftOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YShift, true
}

// SetYShift sets field value
func (o *MaskPosition) SetYShift(v float32) {
	o.YShift = v
}

// GetScale returns the Scale field value
func (o *MaskPosition) GetScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Scale
}

// GetScaleOk returns a tuple with the Scale field value
// and a boolean to check if the value has been set.
func (o *MaskPosition) GetScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scale, true
}

// SetScale sets field value
func (o *MaskPosition) SetScale(v float32) {
	o.Scale = v
}

func (o MaskPosition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["point"] = o.Point
	toSerialize["x_shift"] = o.XShift
	toSerialize["y_shift"] = o.YShift
	toSerialize["scale"] = o.Scale
	return toSerialize, nil
}

func (o *MaskPosition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"point",
		"x_shift",
		"y_shift",
		"scale",
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

	varMaskPosition := _MaskPosition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMaskPosition)

	if err != nil {
		return err
	}

	*o = MaskPosition(varMaskPosition)

	return err
}

type NullableMaskPosition struct {
	value *MaskPosition
	isSet bool
}

func (v NullableMaskPosition) Get() *MaskPosition {
	return v.value
}

func (v *NullableMaskPosition) Set(val *MaskPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskPosition(val *MaskPosition) *NullableMaskPosition {
	return &NullableMaskPosition{value: val, isSet: true}
}

func (v NullableMaskPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


