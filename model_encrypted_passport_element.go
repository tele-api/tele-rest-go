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

// checks if the EncryptedPassportElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptedPassportElement{}

// EncryptedPassportElement Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of “personal\\_details”, “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “address”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration”, “phone\\_number”, “email”.
	Type string `json:"type"`
	// *Optional*. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal\\_details”, “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport” and “address” types. Can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials).
	Data *string `json:"data,omitempty"`
	// *Optional*. User's verified phone number; available only for “phone\\_number” type
	PhoneNumber *string `json:"phone_number,omitempty"`
	// *Optional*. User's verified email address; available only for “email” type
	Email *string `json:"email,omitempty"`
	// *Optional*. Array of encrypted files with documents provided by the user; available only for “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration” and “temporary\\_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials).
	Files []PassportFile `json:"files,omitempty"`
	FrontSide *PassportFile `json:"front_side,omitempty"`
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	Selfie *PassportFile `json:"selfie,omitempty"`
	// *Optional*. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration” and “temporary\\_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials).
	Translation []PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in [PassportElementErrorUnspecified](https://core.telegram.org/bots/api/#passportelementerrorunspecified)
	Hash string `json:"hash"`
}

type _EncryptedPassportElement EncryptedPassportElement

// NewEncryptedPassportElement instantiates a new EncryptedPassportElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptedPassportElement(type_ string, hash string) *EncryptedPassportElement {
	this := EncryptedPassportElement{}
	this.Type = type_
	this.Hash = hash
	return &this
}

// NewEncryptedPassportElementWithDefaults instantiates a new EncryptedPassportElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptedPassportElementWithDefaults() *EncryptedPassportElement {
	this := EncryptedPassportElement{}
	return &this
}

// GetType returns the Type field value
func (o *EncryptedPassportElement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EncryptedPassportElement) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *EncryptedPassportElement) SetData(v string) {
	o.Data = &v
}


// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *EncryptedPassportElement) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}


// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EncryptedPassportElement) SetEmail(v string) {
	o.Email = &v
}


// GetFiles returns the Files field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetFiles() []PassportFile {
	if o == nil || IsNil(o.Files) {
		var ret []PassportFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetFilesOk() ([]PassportFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []PassportFile and assigns it to the Files field.
func (o *EncryptedPassportElement) SetFiles(v []PassportFile) {
	o.Files = v
}


// GetFrontSide returns the FrontSide field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetFrontSide() PassportFile {
	if o == nil || IsNil(o.FrontSide) {
		var ret PassportFile
		return ret
	}
	return *o.FrontSide
}

// GetFrontSideOk returns a tuple with the FrontSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetFrontSideOk() (*PassportFile, bool) {
	if o == nil || IsNil(o.FrontSide) {
		return nil, false
	}
	return o.FrontSide, true
}

// HasFrontSide returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasFrontSide() bool {
	if o != nil && !IsNil(o.FrontSide) {
		return true
	}

	return false
}

// SetFrontSide gets a reference to the given PassportFile and assigns it to the FrontSide field.
func (o *EncryptedPassportElement) SetFrontSide(v PassportFile) {
	o.FrontSide = &v
}


// GetReverseSide returns the ReverseSide field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetReverseSide() PassportFile {
	if o == nil || IsNil(o.ReverseSide) {
		var ret PassportFile
		return ret
	}
	return *o.ReverseSide
}

// GetReverseSideOk returns a tuple with the ReverseSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetReverseSideOk() (*PassportFile, bool) {
	if o == nil || IsNil(o.ReverseSide) {
		return nil, false
	}
	return o.ReverseSide, true
}

// HasReverseSide returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasReverseSide() bool {
	if o != nil && !IsNil(o.ReverseSide) {
		return true
	}

	return false
}

// SetReverseSide gets a reference to the given PassportFile and assigns it to the ReverseSide field.
func (o *EncryptedPassportElement) SetReverseSide(v PassportFile) {
	o.ReverseSide = &v
}


// GetSelfie returns the Selfie field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetSelfie() PassportFile {
	if o == nil || IsNil(o.Selfie) {
		var ret PassportFile
		return ret
	}
	return *o.Selfie
}

// GetSelfieOk returns a tuple with the Selfie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetSelfieOk() (*PassportFile, bool) {
	if o == nil || IsNil(o.Selfie) {
		return nil, false
	}
	return o.Selfie, true
}

// HasSelfie returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasSelfie() bool {
	if o != nil && !IsNil(o.Selfie) {
		return true
	}

	return false
}

// SetSelfie gets a reference to the given PassportFile and assigns it to the Selfie field.
func (o *EncryptedPassportElement) SetSelfie(v PassportFile) {
	o.Selfie = &v
}


// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *EncryptedPassportElement) GetTranslation() []PassportFile {
	if o == nil || IsNil(o.Translation) {
		var ret []PassportFile
		return ret
	}
	return o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetTranslationOk() ([]PassportFile, bool) {
	if o == nil || IsNil(o.Translation) {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *EncryptedPassportElement) HasTranslation() bool {
	if o != nil && !IsNil(o.Translation) {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given []PassportFile and assigns it to the Translation field.
func (o *EncryptedPassportElement) SetTranslation(v []PassportFile) {
	o.Translation = v
}


// GetHash returns the Hash field value
func (o *EncryptedPassportElement) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *EncryptedPassportElement) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *EncryptedPassportElement) SetHash(v string) {
	o.Hash = v
}

func (o EncryptedPassportElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptedPassportElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.FrontSide) {
		toSerialize["front_side"] = o.FrontSide
	}
	if !IsNil(o.ReverseSide) {
		toSerialize["reverse_side"] = o.ReverseSide
	}
	if !IsNil(o.Selfie) {
		toSerialize["selfie"] = o.Selfie
	}
	if !IsNil(o.Translation) {
		toSerialize["translation"] = o.Translation
	}
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *EncryptedPassportElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"hash",
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

	varEncryptedPassportElement := _EncryptedPassportElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEncryptedPassportElement)

	if err != nil {
		return err
	}

	*o = EncryptedPassportElement(varEncryptedPassportElement)

	return err
}

type NullableEncryptedPassportElement struct {
	value *EncryptedPassportElement
	isSet bool
}

func (v NullableEncryptedPassportElement) Get() *EncryptedPassportElement {
	return v.value
}

func (v *NullableEncryptedPassportElement) Set(val *EncryptedPassportElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptedPassportElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptedPassportElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptedPassportElement(val *EncryptedPassportElement) *NullableEncryptedPassportElement {
	return &NullableEncryptedPassportElement{value: val, isSet: true}
}

func (v NullableEncryptedPassportElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptedPassportElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


