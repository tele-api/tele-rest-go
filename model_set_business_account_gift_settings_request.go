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

// checks if the SetBusinessAccountGiftSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetBusinessAccountGiftSettingsRequest{}

// SetBusinessAccountGiftSettingsRequest Request parameters for setBusinessAccountGiftSettings
type SetBusinessAccountGiftSettingsRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Pass *True*, if a button for sending a gift to the user or by the business account must always be shown in the input field
	ShowGiftButton bool `json:"show_gift_button"`
	AcceptedGiftTypes AcceptedGiftTypes `json:"accepted_gift_types"`
}

type _SetBusinessAccountGiftSettingsRequest SetBusinessAccountGiftSettingsRequest

// NewSetBusinessAccountGiftSettingsRequest instantiates a new SetBusinessAccountGiftSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBusinessAccountGiftSettingsRequest(businessConnectionId string, showGiftButton bool, acceptedGiftTypes AcceptedGiftTypes) *SetBusinessAccountGiftSettingsRequest {
	this := SetBusinessAccountGiftSettingsRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.ShowGiftButton = showGiftButton
	this.AcceptedGiftTypes = acceptedGiftTypes
	return &this
}

// NewSetBusinessAccountGiftSettingsRequestWithDefaults instantiates a new SetBusinessAccountGiftSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBusinessAccountGiftSettingsRequestWithDefaults() *SetBusinessAccountGiftSettingsRequest {
	this := SetBusinessAccountGiftSettingsRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *SetBusinessAccountGiftSettingsRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountGiftSettingsRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *SetBusinessAccountGiftSettingsRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetShowGiftButton returns the ShowGiftButton field value
func (o *SetBusinessAccountGiftSettingsRequest) GetShowGiftButton() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowGiftButton
}

// GetShowGiftButtonOk returns a tuple with the ShowGiftButton field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountGiftSettingsRequest) GetShowGiftButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowGiftButton, true
}

// SetShowGiftButton sets field value
func (o *SetBusinessAccountGiftSettingsRequest) SetShowGiftButton(v bool) {
	o.ShowGiftButton = v
}

// GetAcceptedGiftTypes returns the AcceptedGiftTypes field value
func (o *SetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypes() AcceptedGiftTypes {
	if o == nil {
		var ret AcceptedGiftTypes
		return ret
	}

	return o.AcceptedGiftTypes
}

// GetAcceptedGiftTypesOk returns a tuple with the AcceptedGiftTypes field value
// and a boolean to check if the value has been set.
func (o *SetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypesOk() (*AcceptedGiftTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedGiftTypes, true
}

// SetAcceptedGiftTypes sets field value
func (o *SetBusinessAccountGiftSettingsRequest) SetAcceptedGiftTypes(v AcceptedGiftTypes) {
	o.AcceptedGiftTypes = v
}

func (o SetBusinessAccountGiftSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetBusinessAccountGiftSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["show_gift_button"] = o.ShowGiftButton
	toSerialize["accepted_gift_types"] = o.AcceptedGiftTypes
	return toSerialize, nil
}

func (o *SetBusinessAccountGiftSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"show_gift_button",
		"accepted_gift_types",
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

	varSetBusinessAccountGiftSettingsRequest := _SetBusinessAccountGiftSettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetBusinessAccountGiftSettingsRequest)

	if err != nil {
		return err
	}

	*o = SetBusinessAccountGiftSettingsRequest(varSetBusinessAccountGiftSettingsRequest)

	return err
}

type NullableSetBusinessAccountGiftSettingsRequest struct {
	value *SetBusinessAccountGiftSettingsRequest
	isSet bool
}

func (v NullableSetBusinessAccountGiftSettingsRequest) Get() *SetBusinessAccountGiftSettingsRequest {
	return v.value
}

func (v *NullableSetBusinessAccountGiftSettingsRequest) Set(val *SetBusinessAccountGiftSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBusinessAccountGiftSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBusinessAccountGiftSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBusinessAccountGiftSettingsRequest(val *SetBusinessAccountGiftSettingsRequest) *NullableSetBusinessAccountGiftSettingsRequest {
	return &NullableSetBusinessAccountGiftSettingsRequest{value: val, isSet: true}
}

func (v NullableSetBusinessAccountGiftSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBusinessAccountGiftSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


