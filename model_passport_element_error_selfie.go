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

// checks if the PassportElementErrorSelfie type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PassportElementErrorSelfie{}

// PassportElementErrorSelfie Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Error source, must be *selfie*
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

type _PassportElementErrorSelfie PassportElementErrorSelfie

// NewPassportElementErrorSelfie instantiates a new PassportElementErrorSelfie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassportElementErrorSelfie(source string, type_ string, fileHash string, message string) *PassportElementErrorSelfie {
	this := PassportElementErrorSelfie{}
	this.Source = source
	this.Type = type_
	this.FileHash = fileHash
	this.Message = message
	return &this
}

// NewPassportElementErrorSelfieWithDefaults instantiates a new PassportElementErrorSelfie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassportElementErrorSelfieWithDefaults() *PassportElementErrorSelfie {
	this := PassportElementErrorSelfie{}
	var source string = "selfie"
	this.Source = source
	return &this
}

// GetSource returns the Source field value
func (o *PassportElementErrorSelfie) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorSelfie) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PassportElementErrorSelfie) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *PassportElementErrorSelfie) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorSelfie) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PassportElementErrorSelfie) SetType(v string) {
	o.Type = v
}

// GetFileHash returns the FileHash field value
func (o *PassportElementErrorSelfie) GetFileHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileHash
}

// GetFileHashOk returns a tuple with the FileHash field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorSelfie) GetFileHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileHash, true
}

// SetFileHash sets field value
func (o *PassportElementErrorSelfie) SetFileHash(v string) {
	o.FileHash = v
}

// GetMessage returns the Message field value
func (o *PassportElementErrorSelfie) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorSelfie) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PassportElementErrorSelfie) SetMessage(v string) {
	o.Message = v
}

func (o PassportElementErrorSelfie) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassportElementErrorSelfie) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["type"] = o.Type
	toSerialize["file_hash"] = o.FileHash
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *PassportElementErrorSelfie) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"type",
		"file_hash",
		"message",
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

	varPassportElementErrorSelfie := _PassportElementErrorSelfie{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPassportElementErrorSelfie)

	if err != nil {
		return err
	}

	*o = PassportElementErrorSelfie(varPassportElementErrorSelfie)

	return err
}

type NullablePassportElementErrorSelfie struct {
	value *PassportElementErrorSelfie
	isSet bool
}

func (v NullablePassportElementErrorSelfie) Get() *PassportElementErrorSelfie {
	return v.value
}

func (v *NullablePassportElementErrorSelfie) Set(val *PassportElementErrorSelfie) {
	v.value = val
	v.isSet = true
}

func (v NullablePassportElementErrorSelfie) IsSet() bool {
	return v.isSet
}

func (v *NullablePassportElementErrorSelfie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassportElementErrorSelfie(val *PassportElementErrorSelfie) *NullablePassportElementErrorSelfie {
	return &NullablePassportElementErrorSelfie{value: val, isSet: true}
}

func (v NullablePassportElementErrorSelfie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassportElementErrorSelfie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


