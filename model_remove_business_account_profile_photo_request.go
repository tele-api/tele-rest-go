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

// checks if the RemoveBusinessAccountProfilePhotoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveBusinessAccountProfilePhotoRequest{}

// RemoveBusinessAccountProfilePhotoRequest Request parameters for removeBusinessAccountProfilePhoto
type RemoveBusinessAccountProfilePhotoRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Pass *True* to remove the public photo, which is visible even if the main photo is hidden by the business account's privacy settings. After the main photo is removed, the previous profile photo (if present) becomes the main photo.
	IsPublic *bool `json:"is_public,omitempty"`
}

type _RemoveBusinessAccountProfilePhotoRequest RemoveBusinessAccountProfilePhotoRequest

// NewRemoveBusinessAccountProfilePhotoRequest instantiates a new RemoveBusinessAccountProfilePhotoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveBusinessAccountProfilePhotoRequest(businessConnectionId string) *RemoveBusinessAccountProfilePhotoRequest {
	this := RemoveBusinessAccountProfilePhotoRequest{}
	this.BusinessConnectionId = businessConnectionId
	return &this
}

// NewRemoveBusinessAccountProfilePhotoRequestWithDefaults instantiates a new RemoveBusinessAccountProfilePhotoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveBusinessAccountProfilePhotoRequestWithDefaults() *RemoveBusinessAccountProfilePhotoRequest {
	this := RemoveBusinessAccountProfilePhotoRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *RemoveBusinessAccountProfilePhotoRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *RemoveBusinessAccountProfilePhotoRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *RemoveBusinessAccountProfilePhotoRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *RemoveBusinessAccountProfilePhotoRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveBusinessAccountProfilePhotoRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *RemoveBusinessAccountProfilePhotoRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *RemoveBusinessAccountProfilePhotoRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}


func (o RemoveBusinessAccountProfilePhotoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveBusinessAccountProfilePhotoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	return toSerialize, nil
}

func (o *RemoveBusinessAccountProfilePhotoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
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

	varRemoveBusinessAccountProfilePhotoRequest := _RemoveBusinessAccountProfilePhotoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveBusinessAccountProfilePhotoRequest)

	if err != nil {
		return err
	}

	*o = RemoveBusinessAccountProfilePhotoRequest(varRemoveBusinessAccountProfilePhotoRequest)

	return err
}

type NullableRemoveBusinessAccountProfilePhotoRequest struct {
	value *RemoveBusinessAccountProfilePhotoRequest
	isSet bool
}

func (v NullableRemoveBusinessAccountProfilePhotoRequest) Get() *RemoveBusinessAccountProfilePhotoRequest {
	return v.value
}

func (v *NullableRemoveBusinessAccountProfilePhotoRequest) Set(val *RemoveBusinessAccountProfilePhotoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveBusinessAccountProfilePhotoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveBusinessAccountProfilePhotoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveBusinessAccountProfilePhotoRequest(val *RemoveBusinessAccountProfilePhotoRequest) *NullableRemoveBusinessAccountProfilePhotoRequest {
	return &NullableRemoveBusinessAccountProfilePhotoRequest{value: val, isSet: true}
}

func (v NullableRemoveBusinessAccountProfilePhotoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveBusinessAccountProfilePhotoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


