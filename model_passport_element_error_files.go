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

// checks if the PassportElementErrorFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PassportElementErrorFiles{}

// PassportElementErrorFiles Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Error source, must be *files*
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration”
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

type _PassportElementErrorFiles PassportElementErrorFiles

// NewPassportElementErrorFiles instantiates a new PassportElementErrorFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassportElementErrorFiles(source string, type_ string, fileHashes []string, message string) *PassportElementErrorFiles {
	this := PassportElementErrorFiles{}
	this.Source = source
	this.Type = type_
	this.FileHashes = fileHashes
	this.Message = message
	return &this
}

// NewPassportElementErrorFilesWithDefaults instantiates a new PassportElementErrorFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassportElementErrorFilesWithDefaults() *PassportElementErrorFiles {
	this := PassportElementErrorFiles{}
	var source string = "files"
	this.Source = source
	return &this
}

// GetSource returns the Source field value
func (o *PassportElementErrorFiles) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorFiles) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PassportElementErrorFiles) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *PassportElementErrorFiles) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorFiles) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PassportElementErrorFiles) SetType(v string) {
	o.Type = v
}

// GetFileHashes returns the FileHashes field value
func (o *PassportElementErrorFiles) GetFileHashes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FileHashes
}

// GetFileHashesOk returns a tuple with the FileHashes field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorFiles) GetFileHashesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileHashes, true
}

// SetFileHashes sets field value
func (o *PassportElementErrorFiles) SetFileHashes(v []string) {
	o.FileHashes = v
}

// GetMessage returns the Message field value
func (o *PassportElementErrorFiles) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PassportElementErrorFiles) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PassportElementErrorFiles) SetMessage(v string) {
	o.Message = v
}

func (o PassportElementErrorFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassportElementErrorFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["type"] = o.Type
	toSerialize["file_hashes"] = o.FileHashes
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *PassportElementErrorFiles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"type",
		"file_hashes",
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

	varPassportElementErrorFiles := _PassportElementErrorFiles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPassportElementErrorFiles)

	if err != nil {
		return err
	}

	*o = PassportElementErrorFiles(varPassportElementErrorFiles)

	return err
}

type NullablePassportElementErrorFiles struct {
	value *PassportElementErrorFiles
	isSet bool
}

func (v NullablePassportElementErrorFiles) Get() *PassportElementErrorFiles {
	return v.value
}

func (v *NullablePassportElementErrorFiles) Set(val *PassportElementErrorFiles) {
	v.value = val
	v.isSet = true
}

func (v NullablePassportElementErrorFiles) IsSet() bool {
	return v.isSet
}

func (v *NullablePassportElementErrorFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassportElementErrorFiles(val *PassportElementErrorFiles) *NullablePassportElementErrorFiles {
	return &NullablePassportElementErrorFiles{value: val, isSet: true}
}

func (v NullablePassportElementErrorFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassportElementErrorFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


