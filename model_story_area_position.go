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

// checks if the StoryAreaPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryAreaPosition{}

// StoryAreaPosition Describes the position of a clickable area within a story.
type StoryAreaPosition struct {
	// The abscissa of the area's center, as a percentage of the media width
	XPercentage float32 `json:"x_percentage"`
	// The ordinate of the area's center, as a percentage of the media height
	YPercentage float32 `json:"y_percentage"`
	// The width of the area's rectangle, as a percentage of the media width
	WidthPercentage float32 `json:"width_percentage"`
	// The height of the area's rectangle, as a percentage of the media height
	HeightPercentage float32 `json:"height_percentage"`
	// The clockwise rotation angle of the rectangle, in degrees; 0-360
	RotationAngle float32 `json:"rotation_angle"`
	// The radius of the rectangle corner rounding, as a percentage of the media width
	CornerRadiusPercentage float32 `json:"corner_radius_percentage"`
}

type _StoryAreaPosition StoryAreaPosition

// NewStoryAreaPosition instantiates a new StoryAreaPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryAreaPosition(xPercentage float32, yPercentage float32, widthPercentage float32, heightPercentage float32, rotationAngle float32, cornerRadiusPercentage float32) *StoryAreaPosition {
	this := StoryAreaPosition{}
	this.XPercentage = xPercentage
	this.YPercentage = yPercentage
	this.WidthPercentage = widthPercentage
	this.HeightPercentage = heightPercentage
	this.RotationAngle = rotationAngle
	this.CornerRadiusPercentage = cornerRadiusPercentage
	return &this
}

// NewStoryAreaPositionWithDefaults instantiates a new StoryAreaPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryAreaPositionWithDefaults() *StoryAreaPosition {
	this := StoryAreaPosition{}
	return &this
}

// GetXPercentage returns the XPercentage field value
func (o *StoryAreaPosition) GetXPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.XPercentage
}

// GetXPercentageOk returns a tuple with the XPercentage field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetXPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XPercentage, true
}

// SetXPercentage sets field value
func (o *StoryAreaPosition) SetXPercentage(v float32) {
	o.XPercentage = v
}

// GetYPercentage returns the YPercentage field value
func (o *StoryAreaPosition) GetYPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.YPercentage
}

// GetYPercentageOk returns a tuple with the YPercentage field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetYPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YPercentage, true
}

// SetYPercentage sets field value
func (o *StoryAreaPosition) SetYPercentage(v float32) {
	o.YPercentage = v
}

// GetWidthPercentage returns the WidthPercentage field value
func (o *StoryAreaPosition) GetWidthPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WidthPercentage
}

// GetWidthPercentageOk returns a tuple with the WidthPercentage field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetWidthPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidthPercentage, true
}

// SetWidthPercentage sets field value
func (o *StoryAreaPosition) SetWidthPercentage(v float32) {
	o.WidthPercentage = v
}

// GetHeightPercentage returns the HeightPercentage field value
func (o *StoryAreaPosition) GetHeightPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HeightPercentage
}

// GetHeightPercentageOk returns a tuple with the HeightPercentage field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetHeightPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeightPercentage, true
}

// SetHeightPercentage sets field value
func (o *StoryAreaPosition) SetHeightPercentage(v float32) {
	o.HeightPercentage = v
}

// GetRotationAngle returns the RotationAngle field value
func (o *StoryAreaPosition) GetRotationAngle() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RotationAngle
}

// GetRotationAngleOk returns a tuple with the RotationAngle field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetRotationAngleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationAngle, true
}

// SetRotationAngle sets field value
func (o *StoryAreaPosition) SetRotationAngle(v float32) {
	o.RotationAngle = v
}

// GetCornerRadiusPercentage returns the CornerRadiusPercentage field value
func (o *StoryAreaPosition) GetCornerRadiusPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CornerRadiusPercentage
}

// GetCornerRadiusPercentageOk returns a tuple with the CornerRadiusPercentage field value
// and a boolean to check if the value has been set.
func (o *StoryAreaPosition) GetCornerRadiusPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CornerRadiusPercentage, true
}

// SetCornerRadiusPercentage sets field value
func (o *StoryAreaPosition) SetCornerRadiusPercentage(v float32) {
	o.CornerRadiusPercentage = v
}

func (o StoryAreaPosition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryAreaPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x_percentage"] = o.XPercentage
	toSerialize["y_percentage"] = o.YPercentage
	toSerialize["width_percentage"] = o.WidthPercentage
	toSerialize["height_percentage"] = o.HeightPercentage
	toSerialize["rotation_angle"] = o.RotationAngle
	toSerialize["corner_radius_percentage"] = o.CornerRadiusPercentage
	return toSerialize, nil
}

func (o *StoryAreaPosition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x_percentage",
		"y_percentage",
		"width_percentage",
		"height_percentage",
		"rotation_angle",
		"corner_radius_percentage",
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

	varStoryAreaPosition := _StoryAreaPosition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStoryAreaPosition)

	if err != nil {
		return err
	}

	*o = StoryAreaPosition(varStoryAreaPosition)

	return err
}

type NullableStoryAreaPosition struct {
	value *StoryAreaPosition
	isSet bool
}

func (v NullableStoryAreaPosition) Get() *StoryAreaPosition {
	return v.value
}

func (v *NullableStoryAreaPosition) Set(val *StoryAreaPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaPosition(val *StoryAreaPosition) *NullableStoryAreaPosition {
	return &NullableStoryAreaPosition{value: val, isSet: true}
}

func (v NullableStoryAreaPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


