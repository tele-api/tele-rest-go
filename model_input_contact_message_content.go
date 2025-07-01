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

// checks if the InputContactMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputContactMessageContent{}

// InputContactMessageContent Represents the [content](https://core.telegram.org/bots/api/#inputmessagecontent) of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// *Optional*. Contact's last name
	LastName *string `json:"last_name,omitempty"`
	// *Optional*. Additional data about the contact in the form of a [vCard](https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard *string `json:"vcard,omitempty"`
}

type _InputContactMessageContent InputContactMessageContent

// NewInputContactMessageContent instantiates a new InputContactMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputContactMessageContent(phoneNumber string, firstName string) *InputContactMessageContent {
	this := InputContactMessageContent{}
	this.PhoneNumber = phoneNumber
	this.FirstName = firstName
	return &this
}

// NewInputContactMessageContentWithDefaults instantiates a new InputContactMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputContactMessageContentWithDefaults() *InputContactMessageContent {
	this := InputContactMessageContent{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *InputContactMessageContent) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *InputContactMessageContent) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *InputContactMessageContent) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetFirstName returns the FirstName field value
func (o *InputContactMessageContent) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *InputContactMessageContent) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *InputContactMessageContent) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InputContactMessageContent) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputContactMessageContent) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InputContactMessageContent) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InputContactMessageContent) SetLastName(v string) {
	o.LastName = &v
}


// GetVcard returns the Vcard field value if set, zero value otherwise.
func (o *InputContactMessageContent) GetVcard() string {
	if o == nil || IsNil(o.Vcard) {
		var ret string
		return ret
	}
	return *o.Vcard
}

// GetVcardOk returns a tuple with the Vcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputContactMessageContent) GetVcardOk() (*string, bool) {
	if o == nil || IsNil(o.Vcard) {
		return nil, false
	}
	return o.Vcard, true
}

// HasVcard returns a boolean if a field has been set.
func (o *InputContactMessageContent) HasVcard() bool {
	if o != nil && !IsNil(o.Vcard) {
		return true
	}

	return false
}

// SetVcard gets a reference to the given string and assigns it to the Vcard field.
func (o *InputContactMessageContent) SetVcard(v string) {
	o.Vcard = &v
}


func (o InputContactMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputContactMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Vcard) {
		toSerialize["vcard"] = o.Vcard
	}
	return toSerialize, nil
}

func (o *InputContactMessageContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
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

	varInputContactMessageContent := _InputContactMessageContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputContactMessageContent)

	if err != nil {
		return err
	}

	*o = InputContactMessageContent(varInputContactMessageContent)

	return err
}

type NullableInputContactMessageContent struct {
	value *InputContactMessageContent
	isSet bool
}

func (v NullableInputContactMessageContent) Get() *InputContactMessageContent {
	return v.value
}

func (v *NullableInputContactMessageContent) Set(val *InputContactMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputContactMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputContactMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputContactMessageContent(val *InputContactMessageContent) *NullableInputContactMessageContent {
	return &NullableInputContactMessageContent{value: val, isSet: true}
}

func (v NullableInputContactMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputContactMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


