/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the CreateChatInviteLinkPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatInviteLinkPost200Response{}

// CreateChatInviteLinkPost200Response struct for CreateChatInviteLinkPost200Response
type CreateChatInviteLinkPost200Response struct {
	Ok bool `json:"ok"`
	Result ChatInviteLink `json:"result"`
}

type _CreateChatInviteLinkPost200Response CreateChatInviteLinkPost200Response

// NewCreateChatInviteLinkPost200Response instantiates a new CreateChatInviteLinkPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatInviteLinkPost200Response(ok bool, result ChatInviteLink) *CreateChatInviteLinkPost200Response {
	this := CreateChatInviteLinkPost200Response{}
	this.Ok = ok
	this.Result = result
	return &this
}

// NewCreateChatInviteLinkPost200ResponseWithDefaults instantiates a new CreateChatInviteLinkPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatInviteLinkPost200ResponseWithDefaults() *CreateChatInviteLinkPost200Response {
	this := CreateChatInviteLinkPost200Response{}
	var ok bool = true
	this.Ok = ok
	return &this
}

// GetOk returns the Ok field value
func (o *CreateChatInviteLinkPost200Response) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *CreateChatInviteLinkPost200Response) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *CreateChatInviteLinkPost200Response) SetOk(v bool) {
	o.Ok = v
}

// GetResult returns the Result field value
func (o *CreateChatInviteLinkPost200Response) GetResult() ChatInviteLink {
	if o == nil {
		var ret ChatInviteLink
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CreateChatInviteLinkPost200Response) GetResultOk() (*ChatInviteLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CreateChatInviteLinkPost200Response) SetResult(v ChatInviteLink) {
	o.Result = v
}

func (o CreateChatInviteLinkPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatInviteLinkPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ok"] = o.Ok
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CreateChatInviteLinkPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ok",
		"result",
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

	varCreateChatInviteLinkPost200Response := _CreateChatInviteLinkPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatInviteLinkPost200Response)

	if err != nil {
		return err
	}

	*o = CreateChatInviteLinkPost200Response(varCreateChatInviteLinkPost200Response)

	return err
}

type NullableCreateChatInviteLinkPost200Response struct {
	value *CreateChatInviteLinkPost200Response
	isSet bool
}

func (v NullableCreateChatInviteLinkPost200Response) Get() *CreateChatInviteLinkPost200Response {
	return v.value
}

func (v *NullableCreateChatInviteLinkPost200Response) Set(val *CreateChatInviteLinkPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatInviteLinkPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatInviteLinkPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatInviteLinkPost200Response(val *CreateChatInviteLinkPost200Response) *NullableCreateChatInviteLinkPost200Response {
	return &NullableCreateChatInviteLinkPost200Response{value: val, isSet: true}
}

func (v NullableCreateChatInviteLinkPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatInviteLinkPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


