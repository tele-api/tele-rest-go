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

// checks if the BackgroundFillGradient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundFillGradient{}

// BackgroundFillGradient The background is a gradient fill.
type BackgroundFillGradient struct {
	// Type of the background fill, always “gradient”
	Type string `json:"type"`
	// Top color of the gradient in the RGB24 format
	TopColor int32 `json:"top_color"`
	// Bottom color of the gradient in the RGB24 format
	BottomColor int32 `json:"bottom_color"`
	// Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int32 `json:"rotation_angle"`
}

type _BackgroundFillGradient BackgroundFillGradient

// NewBackgroundFillGradient instantiates a new BackgroundFillGradient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundFillGradient(type_ string, topColor int32, bottomColor int32, rotationAngle int32) *BackgroundFillGradient {
	this := BackgroundFillGradient{}
	this.Type = type_
	this.TopColor = topColor
	this.BottomColor = bottomColor
	this.RotationAngle = rotationAngle
	return &this
}

// NewBackgroundFillGradientWithDefaults instantiates a new BackgroundFillGradient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundFillGradientWithDefaults() *BackgroundFillGradient {
	this := BackgroundFillGradient{}
	var type_ string = "gradient"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BackgroundFillGradient) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackgroundFillGradient) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BackgroundFillGradient) SetType(v string) {
	o.Type = v
}

// GetTopColor returns the TopColor field value
func (o *BackgroundFillGradient) GetTopColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TopColor
}

// GetTopColorOk returns a tuple with the TopColor field value
// and a boolean to check if the value has been set.
func (o *BackgroundFillGradient) GetTopColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopColor, true
}

// SetTopColor sets field value
func (o *BackgroundFillGradient) SetTopColor(v int32) {
	o.TopColor = v
}

// GetBottomColor returns the BottomColor field value
func (o *BackgroundFillGradient) GetBottomColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BottomColor
}

// GetBottomColorOk returns a tuple with the BottomColor field value
// and a boolean to check if the value has been set.
func (o *BackgroundFillGradient) GetBottomColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BottomColor, true
}

// SetBottomColor sets field value
func (o *BackgroundFillGradient) SetBottomColor(v int32) {
	o.BottomColor = v
}

// GetRotationAngle returns the RotationAngle field value
func (o *BackgroundFillGradient) GetRotationAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RotationAngle
}

// GetRotationAngleOk returns a tuple with the RotationAngle field value
// and a boolean to check if the value has been set.
func (o *BackgroundFillGradient) GetRotationAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationAngle, true
}

// SetRotationAngle sets field value
func (o *BackgroundFillGradient) SetRotationAngle(v int32) {
	o.RotationAngle = v
}

func (o BackgroundFillGradient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundFillGradient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["top_color"] = o.TopColor
	toSerialize["bottom_color"] = o.BottomColor
	toSerialize["rotation_angle"] = o.RotationAngle
	return toSerialize, nil
}

func (o *BackgroundFillGradient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"top_color",
		"bottom_color",
		"rotation_angle",
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

	varBackgroundFillGradient := _BackgroundFillGradient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackgroundFillGradient)

	if err != nil {
		return err
	}

	*o = BackgroundFillGradient(varBackgroundFillGradient)

	return err
}

type NullableBackgroundFillGradient struct {
	value *BackgroundFillGradient
	isSet bool
}

func (v NullableBackgroundFillGradient) Get() *BackgroundFillGradient {
	return v.value
}

func (v *NullableBackgroundFillGradient) Set(val *BackgroundFillGradient) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundFillGradient) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundFillGradient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundFillGradient(val *BackgroundFillGradient) *NullableBackgroundFillGradient {
	return &NullableBackgroundFillGradient{value: val, isSet: true}
}

func (v NullableBackgroundFillGradient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundFillGradient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


