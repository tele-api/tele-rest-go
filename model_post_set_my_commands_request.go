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

// checks if the PostSetMyCommandsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSetMyCommandsRequest{}

// PostSetMyCommandsRequest struct for PostSetMyCommandsRequest
type PostSetMyCommandsRequest struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []BotCommand `json:"commands"`
	Scope *BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
}

type _PostSetMyCommandsRequest PostSetMyCommandsRequest

// NewPostSetMyCommandsRequest instantiates a new PostSetMyCommandsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSetMyCommandsRequest(commands []BotCommand) *PostSetMyCommandsRequest {
	this := PostSetMyCommandsRequest{}
	this.Commands = commands
	return &this
}

// NewPostSetMyCommandsRequestWithDefaults instantiates a new PostSetMyCommandsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSetMyCommandsRequestWithDefaults() *PostSetMyCommandsRequest {
	this := PostSetMyCommandsRequest{}
	return &this
}

// GetCommands returns the Commands field value
func (o *PostSetMyCommandsRequest) GetCommands() []BotCommand {
	if o == nil {
		var ret []BotCommand
		return ret
	}

	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
func (o *PostSetMyCommandsRequest) GetCommandsOk() ([]BotCommand, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commands, true
}

// SetCommands sets field value
func (o *PostSetMyCommandsRequest) SetCommands(v []BotCommand) {
	o.Commands = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *PostSetMyCommandsRequest) GetScope() BotCommandScope {
	if o == nil || IsNil(o.Scope) {
		var ret BotCommandScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetMyCommandsRequest) GetScopeOk() (*BotCommandScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *PostSetMyCommandsRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given BotCommandScope and assigns it to the Scope field.
func (o *PostSetMyCommandsRequest) SetScope(v BotCommandScope) {
	o.Scope = &v
}


// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *PostSetMyCommandsRequest) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetMyCommandsRequest) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *PostSetMyCommandsRequest) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *PostSetMyCommandsRequest) SetLanguageCode(v string) {
	o.LanguageCode = &v
}


func (o PostSetMyCommandsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSetMyCommandsRequest) ToMap() (map[string]interface{}, error) {
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

func (o *PostSetMyCommandsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostSetMyCommandsRequest := _PostSetMyCommandsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSetMyCommandsRequest)

	if err != nil {
		return err
	}

	*o = PostSetMyCommandsRequest(varPostSetMyCommandsRequest)

	return err
}

type NullablePostSetMyCommandsRequest struct {
	value *PostSetMyCommandsRequest
	isSet bool
}

func (v NullablePostSetMyCommandsRequest) Get() *PostSetMyCommandsRequest {
	return v.value
}

func (v *NullablePostSetMyCommandsRequest) Set(val *PostSetMyCommandsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSetMyCommandsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSetMyCommandsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSetMyCommandsRequest(val *PostSetMyCommandsRequest) *NullablePostSetMyCommandsRequest {
	return &NullablePostSetMyCommandsRequest{value: val, isSet: true}
}

func (v NullablePostSetMyCommandsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSetMyCommandsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


