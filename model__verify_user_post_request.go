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

// checks if the VerifyUserPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyUserPostRequest{}

// VerifyUserPostRequest struct for VerifyUserPostRequest
type VerifyUserPostRequest struct {
	// Unique identifier of the target user
	UserId int32 `json:"user_id"`
	// Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty"`
}

type _VerifyUserPostRequest VerifyUserPostRequest

// NewVerifyUserPostRequest instantiates a new VerifyUserPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyUserPostRequest(userId int32) *VerifyUserPostRequest {
	this := VerifyUserPostRequest{}
	this.UserId = userId
	return &this
}

// NewVerifyUserPostRequestWithDefaults instantiates a new VerifyUserPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyUserPostRequestWithDefaults() *VerifyUserPostRequest {
	this := VerifyUserPostRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *VerifyUserPostRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *VerifyUserPostRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *VerifyUserPostRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetCustomDescription returns the CustomDescription field value if set, zero value otherwise.
func (o *VerifyUserPostRequest) GetCustomDescription() string {
	if o == nil || IsNil(o.CustomDescription) {
		var ret string
		return ret
	}
	return *o.CustomDescription
}

// GetCustomDescriptionOk returns a tuple with the CustomDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyUserPostRequest) GetCustomDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDescription) {
		return nil, false
	}
	return o.CustomDescription, true
}

// HasCustomDescription returns a boolean if a field has been set.
func (o *VerifyUserPostRequest) HasCustomDescription() bool {
	if o != nil && !IsNil(o.CustomDescription) {
		return true
	}

	return false
}

// SetCustomDescription gets a reference to the given string and assigns it to the CustomDescription field.
func (o *VerifyUserPostRequest) SetCustomDescription(v string) {
	o.CustomDescription = &v
}


func (o VerifyUserPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyUserPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.CustomDescription) {
		toSerialize["custom_description"] = o.CustomDescription
	}
	return toSerialize, nil
}

func (o *VerifyUserPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varVerifyUserPostRequest := _VerifyUserPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyUserPostRequest)

	if err != nil {
		return err
	}

	*o = VerifyUserPostRequest(varVerifyUserPostRequest)

	return err
}

type NullableVerifyUserPostRequest struct {
	value *VerifyUserPostRequest
	isSet bool
}

func (v NullableVerifyUserPostRequest) Get() *VerifyUserPostRequest {
	return v.value
}

func (v *NullableVerifyUserPostRequest) Set(val *VerifyUserPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyUserPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyUserPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyUserPostRequest(val *VerifyUserPostRequest) *NullableVerifyUserPostRequest {
	return &NullableVerifyUserPostRequest{value: val, isSet: true}
}

func (v NullableVerifyUserPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyUserPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


