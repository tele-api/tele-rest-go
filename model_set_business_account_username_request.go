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

// checks if the SetBusinessAccountUsernameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetBusinessAccountUsernameRequest{}

// SetBusinessAccountUsernameRequest Request parameters for setBusinessAccountUsername
type SetBusinessAccountUsernameRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// The new value of the username for the business account; 0-32 characters
	Username *string `json:"username,omitempty"`
}

type _SetBusinessAccountUsernameRequest SetBusinessAccountUsernameRequest

// NewSetBusinessAccountUsernameRequest instantiates a new SetBusinessAccountUsernameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBusinessAccountUsernameRequest(businessConnectionId string) *SetBusinessAccountUsernameRequest {
	this := SetBusinessAccountUsernameRequest{}
	this.BusinessConnectionId = businessConnectionId
	return &this
}

// NewSetBusinessAccountUsernameRequestWithDefaults instantiates a new SetBusinessAccountUsernameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBusinessAccountUsernameRequestWithDefaults() *SetBusinessAccountUsernameRequest {
	this := SetBusinessAccountUsernameRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *SetBusinessAccountUsernameRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountUsernameRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *SetBusinessAccountUsernameRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SetBusinessAccountUsernameRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountUsernameRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SetBusinessAccountUsernameRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SetBusinessAccountUsernameRequest) SetUsername(v string) {
	o.Username = &v
}


func (o SetBusinessAccountUsernameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetBusinessAccountUsernameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *SetBusinessAccountUsernameRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSetBusinessAccountUsernameRequest := _SetBusinessAccountUsernameRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetBusinessAccountUsernameRequest)

	if err != nil {
		return err
	}

	*o = SetBusinessAccountUsernameRequest(varSetBusinessAccountUsernameRequest)

	return err
}

type NullableSetBusinessAccountUsernameRequest struct {
	value *SetBusinessAccountUsernameRequest
	isSet bool
}

func (v NullableSetBusinessAccountUsernameRequest) Get() *SetBusinessAccountUsernameRequest {
	return v.value
}

func (v *NullableSetBusinessAccountUsernameRequest) Set(val *SetBusinessAccountUsernameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBusinessAccountUsernameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBusinessAccountUsernameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBusinessAccountUsernameRequest(val *SetBusinessAccountUsernameRequest) *NullableSetBusinessAccountUsernameRequest {
	return &NullableSetBusinessAccountUsernameRequest{value: val, isSet: true}
}

func (v NullableSetBusinessAccountUsernameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBusinessAccountUsernameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


