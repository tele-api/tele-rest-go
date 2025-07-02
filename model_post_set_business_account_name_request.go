/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the PostSetBusinessAccountNameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSetBusinessAccountNameRequest{}

// PostSetBusinessAccountNameRequest struct for PostSetBusinessAccountNameRequest
type PostSetBusinessAccountNameRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// The new value of the first name for the business account; 1-64 characters
	FirstName string `json:"first_name"`
	// The new value of the last name for the business account; 0-64 characters
	LastName *string `json:"last_name,omitempty"`
}

type _PostSetBusinessAccountNameRequest PostSetBusinessAccountNameRequest

// NewPostSetBusinessAccountNameRequest instantiates a new PostSetBusinessAccountNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSetBusinessAccountNameRequest(businessConnectionId string, firstName string) *PostSetBusinessAccountNameRequest {
	this := PostSetBusinessAccountNameRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.FirstName = firstName
	return &this
}

// NewPostSetBusinessAccountNameRequestWithDefaults instantiates a new PostSetBusinessAccountNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSetBusinessAccountNameRequestWithDefaults() *PostSetBusinessAccountNameRequest {
	this := PostSetBusinessAccountNameRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *PostSetBusinessAccountNameRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *PostSetBusinessAccountNameRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *PostSetBusinessAccountNameRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetFirstName returns the FirstName field value
func (o *PostSetBusinessAccountNameRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *PostSetBusinessAccountNameRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *PostSetBusinessAccountNameRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PostSetBusinessAccountNameRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetBusinessAccountNameRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PostSetBusinessAccountNameRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PostSetBusinessAccountNameRequest) SetLastName(v string) {
	o.LastName = &v
}


func (o PostSetBusinessAccountNameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSetBusinessAccountNameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	return toSerialize, nil
}

func (o *PostSetBusinessAccountNameRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"first_name",
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

	varPostSetBusinessAccountNameRequest := _PostSetBusinessAccountNameRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSetBusinessAccountNameRequest)

	if err != nil {
		return err
	}

	*o = PostSetBusinessAccountNameRequest(varPostSetBusinessAccountNameRequest)

	return err
}

type NullablePostSetBusinessAccountNameRequest struct {
	value *PostSetBusinessAccountNameRequest
	isSet bool
}

func (v NullablePostSetBusinessAccountNameRequest) Get() *PostSetBusinessAccountNameRequest {
	return v.value
}

func (v *NullablePostSetBusinessAccountNameRequest) Set(val *PostSetBusinessAccountNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSetBusinessAccountNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSetBusinessAccountNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSetBusinessAccountNameRequest(val *PostSetBusinessAccountNameRequest) *NullablePostSetBusinessAccountNameRequest {
	return &NullablePostSetBusinessAccountNameRequest{value: val, isSet: true}
}

func (v NullablePostSetBusinessAccountNameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSetBusinessAccountNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


