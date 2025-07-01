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

// checks if the SetMyCommandsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetMyCommandsPostRequest{}

// SetMyCommandsPostRequest struct for SetMyCommandsPostRequest
type SetMyCommandsPostRequest struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []BotCommand `json:"commands"`
	Scope *BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
}

type _SetMyCommandsPostRequest SetMyCommandsPostRequest

// NewSetMyCommandsPostRequest instantiates a new SetMyCommandsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMyCommandsPostRequest(commands []BotCommand) *SetMyCommandsPostRequest {
	this := SetMyCommandsPostRequest{}
	this.Commands = commands
	return &this
}

// NewSetMyCommandsPostRequestWithDefaults instantiates a new SetMyCommandsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMyCommandsPostRequestWithDefaults() *SetMyCommandsPostRequest {
	this := SetMyCommandsPostRequest{}
	return &this
}

// GetCommands returns the Commands field value
func (o *SetMyCommandsPostRequest) GetCommands() []BotCommand {
	if o == nil {
		var ret []BotCommand
		return ret
	}

	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
func (o *SetMyCommandsPostRequest) GetCommandsOk() ([]BotCommand, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commands, true
}

// SetCommands sets field value
func (o *SetMyCommandsPostRequest) SetCommands(v []BotCommand) {
	o.Commands = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SetMyCommandsPostRequest) GetScope() BotCommandScope {
	if o == nil || IsNil(o.Scope) {
		var ret BotCommandScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMyCommandsPostRequest) GetScopeOk() (*BotCommandScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SetMyCommandsPostRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given BotCommandScope and assigns it to the Scope field.
func (o *SetMyCommandsPostRequest) SetScope(v BotCommandScope) {
	o.Scope = &v
}


// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *SetMyCommandsPostRequest) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMyCommandsPostRequest) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *SetMyCommandsPostRequest) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *SetMyCommandsPostRequest) SetLanguageCode(v string) {
	o.LanguageCode = &v
}


func (o SetMyCommandsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMyCommandsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commands"] = o.Commands
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["language_code"] = o.LanguageCode
	}
	return toSerialize, nil
}

func (o *SetMyCommandsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commands",
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

	varSetMyCommandsPostRequest := _SetMyCommandsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetMyCommandsPostRequest)

	if err != nil {
		return err
	}

	*o = SetMyCommandsPostRequest(varSetMyCommandsPostRequest)

	return err
}

type NullableSetMyCommandsPostRequest struct {
	value *SetMyCommandsPostRequest
	isSet bool
}

func (v NullableSetMyCommandsPostRequest) Get() *SetMyCommandsPostRequest {
	return v.value
}

func (v *NullableSetMyCommandsPostRequest) Set(val *SetMyCommandsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMyCommandsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMyCommandsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMyCommandsPostRequest(val *SetMyCommandsPostRequest) *NullableSetMyCommandsPostRequest {
	return &NullableSetMyCommandsPostRequest{value: val, isSet: true}
}

func (v NullableSetMyCommandsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMyCommandsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


