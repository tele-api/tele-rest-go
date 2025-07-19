/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the InputProfilePhotoAnimated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputProfilePhotoAnimated{}

// InputProfilePhotoAnimated An animated profile photo in the MPEG4 format.
type InputProfilePhotoAnimated struct {
	// Type of the profile photo, must be *animated*
	Type string `json:"type"`
	// The animated profile photo. Profile photos can't be reused and can only be uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the photo was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Animation string `json:"animation"`
	// *Optional*. Timestamp in seconds of the frame that will be used as the static profile photo. Defaults to 0.0.
	MainFrameTimestamp *float32 `json:"main_frame_timestamp,omitempty"`
}

type _InputProfilePhotoAnimated InputProfilePhotoAnimated

// NewInputProfilePhotoAnimated instantiates a new InputProfilePhotoAnimated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputProfilePhotoAnimated(type_ string, animation string) *InputProfilePhotoAnimated {
	this := InputProfilePhotoAnimated{}
	this.Type = type_
	this.Animation = animation
	return &this
}

// NewInputProfilePhotoAnimatedWithDefaults instantiates a new InputProfilePhotoAnimated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputProfilePhotoAnimatedWithDefaults() *InputProfilePhotoAnimated {
	this := InputProfilePhotoAnimated{}
	var type_ string = "animated"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputProfilePhotoAnimated) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputProfilePhotoAnimated) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputProfilePhotoAnimated) SetType(v string) {
	o.Type = v
}

// GetAnimation returns the Animation field value
func (o *InputProfilePhotoAnimated) GetAnimation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Animation
}

// GetAnimationOk returns a tuple with the Animation field value
// and a boolean to check if the value has been set.
func (o *InputProfilePhotoAnimated) GetAnimationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Animation, true
}

// SetAnimation sets field value
func (o *InputProfilePhotoAnimated) SetAnimation(v string) {
	o.Animation = v
}

// GetMainFrameTimestamp returns the MainFrameTimestamp field value if set, zero value otherwise.
func (o *InputProfilePhotoAnimated) GetMainFrameTimestamp() float32 {
	if o == nil || IsNil(o.MainFrameTimestamp) {
		var ret float32
		return ret
	}
	return *o.MainFrameTimestamp
}

// GetMainFrameTimestampOk returns a tuple with the MainFrameTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputProfilePhotoAnimated) GetMainFrameTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.MainFrameTimestamp) {
		return nil, false
	}
	return o.MainFrameTimestamp, true
}

// HasMainFrameTimestamp returns a boolean if a field has been set.
func (o *InputProfilePhotoAnimated) HasMainFrameTimestamp() bool {
	if o != nil && !IsNil(o.MainFrameTimestamp) {
		return true
	}

	return false
}

// SetMainFrameTimestamp gets a reference to the given float32 and assigns it to the MainFrameTimestamp field.
func (o *InputProfilePhotoAnimated) SetMainFrameTimestamp(v float32) {
	o.MainFrameTimestamp = &v
}


func (o InputProfilePhotoAnimated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputProfilePhotoAnimated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["animation"] = o.Animation
	if !IsNil(o.MainFrameTimestamp) {
		toSerialize["main_frame_timestamp"] = o.MainFrameTimestamp
	}
	return toSerialize, nil
}

func (o *InputProfilePhotoAnimated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"animation",
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

	varInputProfilePhotoAnimated := _InputProfilePhotoAnimated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputProfilePhotoAnimated)

	if err != nil {
		return err
	}

	*o = InputProfilePhotoAnimated(varInputProfilePhotoAnimated)

	return err
}

type NullableInputProfilePhotoAnimated struct {
	value *InputProfilePhotoAnimated
	isSet bool
}

func (v NullableInputProfilePhotoAnimated) Get() *InputProfilePhotoAnimated {
	return v.value
}

func (v *NullableInputProfilePhotoAnimated) Set(val *InputProfilePhotoAnimated) {
	v.value = val
	v.isSet = true
}

func (v NullableInputProfilePhotoAnimated) IsSet() bool {
	return v.isSet
}

func (v *NullableInputProfilePhotoAnimated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputProfilePhotoAnimated(val *InputProfilePhotoAnimated) *NullableInputProfilePhotoAnimated {
	return &NullableInputProfilePhotoAnimated{value: val, isSet: true}
}

func (v NullableInputProfilePhotoAnimated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputProfilePhotoAnimated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


