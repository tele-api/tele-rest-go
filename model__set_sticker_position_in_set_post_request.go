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

// checks if the SetStickerPositionInSetPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetStickerPositionInSetPostRequest{}

// SetStickerPositionInSetPostRequest struct for SetStickerPositionInSetPostRequest
type SetStickerPositionInSetPostRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// New sticker position in the set, zero-based
	Position int32 `json:"position"`
}

type _SetStickerPositionInSetPostRequest SetStickerPositionInSetPostRequest

// NewSetStickerPositionInSetPostRequest instantiates a new SetStickerPositionInSetPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetStickerPositionInSetPostRequest(sticker string, position int32) *SetStickerPositionInSetPostRequest {
	this := SetStickerPositionInSetPostRequest{}
	this.Sticker = sticker
	this.Position = position
	return &this
}

// NewSetStickerPositionInSetPostRequestWithDefaults instantiates a new SetStickerPositionInSetPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetStickerPositionInSetPostRequestWithDefaults() *SetStickerPositionInSetPostRequest {
	this := SetStickerPositionInSetPostRequest{}
	return &this
}

// GetSticker returns the Sticker field value
func (o *SetStickerPositionInSetPostRequest) GetSticker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value
// and a boolean to check if the value has been set.
func (o *SetStickerPositionInSetPostRequest) GetStickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticker, true
}

// SetSticker sets field value
func (o *SetStickerPositionInSetPostRequest) SetSticker(v string) {
	o.Sticker = v
}

// GetPosition returns the Position field value
func (o *SetStickerPositionInSetPostRequest) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *SetStickerPositionInSetPostRequest) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *SetStickerPositionInSetPostRequest) SetPosition(v int32) {
	o.Position = v
}

func (o SetStickerPositionInSetPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetStickerPositionInSetPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sticker"] = o.Sticker
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

func (o *SetStickerPositionInSetPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sticker",
		"position",
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

	varSetStickerPositionInSetPostRequest := _SetStickerPositionInSetPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetStickerPositionInSetPostRequest)

	if err != nil {
		return err
	}

	*o = SetStickerPositionInSetPostRequest(varSetStickerPositionInSetPostRequest)

	return err
}

type NullableSetStickerPositionInSetPostRequest struct {
	value *SetStickerPositionInSetPostRequest
	isSet bool
}

func (v NullableSetStickerPositionInSetPostRequest) Get() *SetStickerPositionInSetPostRequest {
	return v.value
}

func (v *NullableSetStickerPositionInSetPostRequest) Set(val *SetStickerPositionInSetPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetStickerPositionInSetPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetStickerPositionInSetPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetStickerPositionInSetPostRequest(val *SetStickerPositionInSetPostRequest) *NullableSetStickerPositionInSetPostRequest {
	return &NullableSetStickerPositionInSetPostRequest{value: val, isSet: true}
}

func (v NullableSetStickerPositionInSetPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetStickerPositionInSetPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


