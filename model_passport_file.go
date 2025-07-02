/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the PassportFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PassportFile{}

// PassportFile This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// File size in bytes
	FileSize int32 `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int32 `json:"file_date"`
}

type _PassportFile PassportFile

// NewPassportFile instantiates a new PassportFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassportFile(fileId string, fileUniqueId string, fileSize int32, fileDate int32) *PassportFile {
	this := PassportFile{}
	this.FileId = fileId
	this.FileUniqueId = fileUniqueId
	this.FileSize = fileSize
	this.FileDate = fileDate
	return &this
}

// NewPassportFileWithDefaults instantiates a new PassportFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassportFileWithDefaults() *PassportFile {
	this := PassportFile{}
	return &this
}

// GetFileId returns the FileId field value
func (o *PassportFile) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *PassportFile) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *PassportFile) SetFileId(v string) {
	o.FileId = v
}

// GetFileUniqueId returns the FileUniqueId field value
func (o *PassportFile) GetFileUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUniqueId
}

// GetFileUniqueIdOk returns a tuple with the FileUniqueId field value
// and a boolean to check if the value has been set.
func (o *PassportFile) GetFileUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUniqueId, true
}

// SetFileUniqueId sets field value
func (o *PassportFile) SetFileUniqueId(v string) {
	o.FileUniqueId = v
}

// GetFileSize returns the FileSize field value
func (o *PassportFile) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *PassportFile) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *PassportFile) SetFileSize(v int32) {
	o.FileSize = v
}

// GetFileDate returns the FileDate field value
func (o *PassportFile) GetFileDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileDate
}

// GetFileDateOk returns a tuple with the FileDate field value
// and a boolean to check if the value has been set.
func (o *PassportFile) GetFileDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileDate, true
}

// SetFileDate sets field value
func (o *PassportFile) SetFileDate(v int32) {
	o.FileDate = v
}

func (o PassportFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassportFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_id"] = o.FileId
	toSerialize["file_unique_id"] = o.FileUniqueId
	toSerialize["file_size"] = o.FileSize
	toSerialize["file_date"] = o.FileDate
	return toSerialize, nil
}

func (o *PassportFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_id",
		"file_unique_id",
		"file_size",
		"file_date",
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

	varPassportFile := _PassportFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPassportFile)

	if err != nil {
		return err
	}

	*o = PassportFile(varPassportFile)

	return err
}

type NullablePassportFile struct {
	value *PassportFile
	isSet bool
}

func (v NullablePassportFile) Get() *PassportFile {
	return v.value
}

func (v *NullablePassportFile) Set(val *PassportFile) {
	v.value = val
	v.isSet = true
}

func (v NullablePassportFile) IsSet() bool {
	return v.isSet
}

func (v *NullablePassportFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassportFile(val *PassportFile) *NullablePassportFile {
	return &NullablePassportFile{value: val, isSet: true}
}

func (v NullablePassportFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassportFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


