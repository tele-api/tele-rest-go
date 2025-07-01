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

// checks if the ConvertGiftToStarsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertGiftToStarsPostRequest{}

// ConvertGiftToStarsPostRequest struct for ConvertGiftToStarsPostRequest
type ConvertGiftToStarsPostRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier of the regular gift that should be converted to Telegram Stars
	OwnedGiftId string `json:"owned_gift_id"`
}

type _ConvertGiftToStarsPostRequest ConvertGiftToStarsPostRequest

// NewConvertGiftToStarsPostRequest instantiates a new ConvertGiftToStarsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertGiftToStarsPostRequest(businessConnectionId string, ownedGiftId string) *ConvertGiftToStarsPostRequest {
	this := ConvertGiftToStarsPostRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.OwnedGiftId = ownedGiftId
	return &this
}

// NewConvertGiftToStarsPostRequestWithDefaults instantiates a new ConvertGiftToStarsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertGiftToStarsPostRequestWithDefaults() *ConvertGiftToStarsPostRequest {
	this := ConvertGiftToStarsPostRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *ConvertGiftToStarsPostRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *ConvertGiftToStarsPostRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *ConvertGiftToStarsPostRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetOwnedGiftId returns the OwnedGiftId field value
func (o *ConvertGiftToStarsPostRequest) GetOwnedGiftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value
// and a boolean to check if the value has been set.
func (o *ConvertGiftToStarsPostRequest) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedGiftId, true
}

// SetOwnedGiftId sets field value
func (o *ConvertGiftToStarsPostRequest) SetOwnedGiftId(v string) {
	o.OwnedGiftId = v
}

func (o ConvertGiftToStarsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertGiftToStarsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["owned_gift_id"] = o.OwnedGiftId
	return toSerialize, nil
}

func (o *ConvertGiftToStarsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"owned_gift_id",
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

	varConvertGiftToStarsPostRequest := _ConvertGiftToStarsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConvertGiftToStarsPostRequest)

	if err != nil {
		return err
	}

	*o = ConvertGiftToStarsPostRequest(varConvertGiftToStarsPostRequest)

	return err
}

type NullableConvertGiftToStarsPostRequest struct {
	value *ConvertGiftToStarsPostRequest
	isSet bool
}

func (v NullableConvertGiftToStarsPostRequest) Get() *ConvertGiftToStarsPostRequest {
	return v.value
}

func (v *NullableConvertGiftToStarsPostRequest) Set(val *ConvertGiftToStarsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertGiftToStarsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertGiftToStarsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertGiftToStarsPostRequest(val *ConvertGiftToStarsPostRequest) *NullableConvertGiftToStarsPostRequest {
	return &NullableConvertGiftToStarsPostRequest{value: val, isSet: true}
}

func (v NullableConvertGiftToStarsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertGiftToStarsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


