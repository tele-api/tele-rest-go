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

// checks if the SetBusinessAccountBioPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetBusinessAccountBioPostRequest{}

// SetBusinessAccountBioPostRequest struct for SetBusinessAccountBioPostRequest
type SetBusinessAccountBioPostRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// The new value of the bio for the business account; 0-140 characters
	Bio *string `json:"bio,omitempty"`
}

type _SetBusinessAccountBioPostRequest SetBusinessAccountBioPostRequest

// NewSetBusinessAccountBioPostRequest instantiates a new SetBusinessAccountBioPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBusinessAccountBioPostRequest(businessConnectionId string) *SetBusinessAccountBioPostRequest {
	this := SetBusinessAccountBioPostRequest{}
	this.BusinessConnectionId = businessConnectionId
	return &this
}

// NewSetBusinessAccountBioPostRequestWithDefaults instantiates a new SetBusinessAccountBioPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBusinessAccountBioPostRequestWithDefaults() *SetBusinessAccountBioPostRequest {
	this := SetBusinessAccountBioPostRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *SetBusinessAccountBioPostRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountBioPostRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *SetBusinessAccountBioPostRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *SetBusinessAccountBioPostRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountBioPostRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *SetBusinessAccountBioPostRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *SetBusinessAccountBioPostRequest) SetBio(v string) {
	o.Bio = &v
}


func (o SetBusinessAccountBioPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetBusinessAccountBioPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	return toSerialize, nil
}

func (o *SetBusinessAccountBioPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSetBusinessAccountBioPostRequest := _SetBusinessAccountBioPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetBusinessAccountBioPostRequest)

	if err != nil {
		return err
	}

	*o = SetBusinessAccountBioPostRequest(varSetBusinessAccountBioPostRequest)

	return err
}

type NullableSetBusinessAccountBioPostRequest struct {
	value *SetBusinessAccountBioPostRequest
	isSet bool
}

func (v NullableSetBusinessAccountBioPostRequest) Get() *SetBusinessAccountBioPostRequest {
	return v.value
}

func (v *NullableSetBusinessAccountBioPostRequest) Set(val *SetBusinessAccountBioPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBusinessAccountBioPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBusinessAccountBioPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBusinessAccountBioPostRequest(val *SetBusinessAccountBioPostRequest) *NullableSetBusinessAccountBioPostRequest {
	return &NullableSetBusinessAccountBioPostRequest{value: val, isSet: true}
}

func (v NullableSetBusinessAccountBioPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBusinessAccountBioPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


