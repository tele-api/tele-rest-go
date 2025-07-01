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

// checks if the InputProfilePhotoStatic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputProfilePhotoStatic{}

// InputProfilePhotoStatic A static profile photo in the .JPG format.
type InputProfilePhotoStatic struct {
	// Type of the profile photo, must be *static*
	Type string `json:"type"`
	// The static profile photo. Profile photos can't be reused and can only be uploaded as a new file, so you can pass “attach://\\<file\\_attach\\_name\\>” if the photo was uploaded using multipart/form-data under \\<file\\_attach\\_name\\>. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files)
	Photo string `json:"photo"`
}

type _InputProfilePhotoStatic InputProfilePhotoStatic

// NewInputProfilePhotoStatic instantiates a new InputProfilePhotoStatic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputProfilePhotoStatic(type_ string, photo string) *InputProfilePhotoStatic {
	this := InputProfilePhotoStatic{}
	this.Type = type_
	this.Photo = photo
	return &this
}

// NewInputProfilePhotoStaticWithDefaults instantiates a new InputProfilePhotoStatic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputProfilePhotoStaticWithDefaults() *InputProfilePhotoStatic {
	this := InputProfilePhotoStatic{}
	var type_ string = "static"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InputProfilePhotoStatic) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputProfilePhotoStatic) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputProfilePhotoStatic) SetType(v string) {
	o.Type = v
}

// GetPhoto returns the Photo field value
func (o *InputProfilePhotoStatic) GetPhoto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value
// and a boolean to check if the value has been set.
func (o *InputProfilePhotoStatic) GetPhotoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Photo, true
}

// SetPhoto sets field value
func (o *InputProfilePhotoStatic) SetPhoto(v string) {
	o.Photo = v
}

func (o InputProfilePhotoStatic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputProfilePhotoStatic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["photo"] = o.Photo
	return toSerialize, nil
}

func (o *InputProfilePhotoStatic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"photo",
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

	varInputProfilePhotoStatic := _InputProfilePhotoStatic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputProfilePhotoStatic)

	if err != nil {
		return err
	}

	*o = InputProfilePhotoStatic(varInputProfilePhotoStatic)

	return err
}

type NullableInputProfilePhotoStatic struct {
	value *InputProfilePhotoStatic
	isSet bool
}

func (v NullableInputProfilePhotoStatic) Get() *InputProfilePhotoStatic {
	return v.value
}

func (v *NullableInputProfilePhotoStatic) Set(val *InputProfilePhotoStatic) {
	v.value = val
	v.isSet = true
}

func (v NullableInputProfilePhotoStatic) IsSet() bool {
	return v.isSet
}

func (v *NullableInputProfilePhotoStatic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputProfilePhotoStatic(val *InputProfilePhotoStatic) *NullableInputProfilePhotoStatic {
	return &NullableInputProfilePhotoStatic{value: val, isSet: true}
}

func (v NullableInputProfilePhotoStatic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputProfilePhotoStatic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


