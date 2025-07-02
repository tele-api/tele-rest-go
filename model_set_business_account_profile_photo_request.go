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

// checks if the SetBusinessAccountProfilePhotoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetBusinessAccountProfilePhotoRequest{}

// SetBusinessAccountProfilePhotoRequest Request parameters for setBusinessAccountProfilePhoto
type SetBusinessAccountProfilePhotoRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	Photo InputProfilePhoto `json:"photo"`
	// Pass True to set the public photo, which will be visible even if the main photo is hidden by the business account's privacy settings. An account can have only one public photo.
	IsPublic *bool `json:"is_public,omitempty"`
}

type _SetBusinessAccountProfilePhotoRequest SetBusinessAccountProfilePhotoRequest

// NewSetBusinessAccountProfilePhotoRequest instantiates a new SetBusinessAccountProfilePhotoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBusinessAccountProfilePhotoRequest(businessConnectionId string, photo InputProfilePhoto) *SetBusinessAccountProfilePhotoRequest {
	this := SetBusinessAccountProfilePhotoRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.Photo = photo
	return &this
}

// NewSetBusinessAccountProfilePhotoRequestWithDefaults instantiates a new SetBusinessAccountProfilePhotoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBusinessAccountProfilePhotoRequestWithDefaults() *SetBusinessAccountProfilePhotoRequest {
	this := SetBusinessAccountProfilePhotoRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *SetBusinessAccountProfilePhotoRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountProfilePhotoRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *SetBusinessAccountProfilePhotoRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetPhoto returns the Photo field value
func (o *SetBusinessAccountProfilePhotoRequest) GetPhoto() InputProfilePhoto {
	if o == nil {
		var ret InputProfilePhoto
		return ret
	}

	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountProfilePhotoRequest) GetPhotoOk() (*InputProfilePhoto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Photo, true
}

// SetPhoto sets field value
func (o *SetBusinessAccountProfilePhotoRequest) SetPhoto(v InputProfilePhoto) {
	o.Photo = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *SetBusinessAccountProfilePhotoRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountProfilePhotoRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *SetBusinessAccountProfilePhotoRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *SetBusinessAccountProfilePhotoRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}


func (o SetBusinessAccountProfilePhotoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetBusinessAccountProfilePhotoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["photo"] = o.Photo
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	return toSerialize, nil
}

func (o *SetBusinessAccountProfilePhotoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
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

	varSetBusinessAccountProfilePhotoRequest := _SetBusinessAccountProfilePhotoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetBusinessAccountProfilePhotoRequest)

	if err != nil {
		return err
	}

	*o = SetBusinessAccountProfilePhotoRequest(varSetBusinessAccountProfilePhotoRequest)

	return err
}

type NullableSetBusinessAccountProfilePhotoRequest struct {
	value *SetBusinessAccountProfilePhotoRequest
	isSet bool
}

func (v NullableSetBusinessAccountProfilePhotoRequest) Get() *SetBusinessAccountProfilePhotoRequest {
	return v.value
}

func (v *NullableSetBusinessAccountProfilePhotoRequest) Set(val *SetBusinessAccountProfilePhotoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBusinessAccountProfilePhotoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBusinessAccountProfilePhotoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBusinessAccountProfilePhotoRequest(val *SetBusinessAccountProfilePhotoRequest) *NullableSetBusinessAccountProfilePhotoRequest {
	return &NullableSetBusinessAccountProfilePhotoRequest{value: val, isSet: true}
}

func (v NullableSetBusinessAccountProfilePhotoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBusinessAccountProfilePhotoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


