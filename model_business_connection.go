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

// checks if the BusinessConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessConnection{}

// BusinessConnection Describes the connection of the bot with a business account.
type BusinessConnection struct {
	// Unique identifier of the business connection
	Id string `json:"id"`
	User User `json:"user"`
	// Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserChatId int32 `json:"user_chat_id"`
	// Date the connection was established in Unix time
	Date int32 `json:"date"`
	Rights *BusinessBotRights `json:"rights,omitempty"`
	// True, if the connection is active
	IsEnabled bool `json:"is_enabled"`
}

type _BusinessConnection BusinessConnection

// NewBusinessConnection instantiates a new BusinessConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessConnection(id string, user User, userChatId int32, date int32, isEnabled bool) *BusinessConnection {
	this := BusinessConnection{}
	this.Id = id
	this.User = user
	this.UserChatId = userChatId
	this.Date = date
	this.IsEnabled = isEnabled
	return &this
}

// NewBusinessConnectionWithDefaults instantiates a new BusinessConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessConnectionWithDefaults() *BusinessConnection {
	this := BusinessConnection{}
	return &this
}

// GetId returns the Id field value
func (o *BusinessConnection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BusinessConnection) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *BusinessConnection) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BusinessConnection) SetUser(v User) {
	o.User = v
}

// GetUserChatId returns the UserChatId field value
func (o *BusinessConnection) GetUserChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserChatId
}

// GetUserChatIdOk returns a tuple with the UserChatId field value
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetUserChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserChatId, true
}

// SetUserChatId sets field value
func (o *BusinessConnection) SetUserChatId(v int32) {
	o.UserChatId = v
}

// GetDate returns the Date field value
func (o *BusinessConnection) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *BusinessConnection) SetDate(v int32) {
	o.Date = v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *BusinessConnection) GetRights() BusinessBotRights {
	if o == nil || IsNil(o.Rights) {
		var ret BusinessBotRights
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetRightsOk() (*BusinessBotRights, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *BusinessConnection) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given BusinessBotRights and assigns it to the Rights field.
func (o *BusinessConnection) SetRights(v BusinessBotRights) {
	o.Rights = &v
}


// GetIsEnabled returns the IsEnabled field value
func (o *BusinessConnection) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *BusinessConnection) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *BusinessConnection) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o BusinessConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user"] = o.User
	toSerialize["user_chat_id"] = o.UserChatId
	toSerialize["date"] = o.Date
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	toSerialize["is_enabled"] = o.IsEnabled
	return toSerialize, nil
}

func (o *BusinessConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user",
		"user_chat_id",
		"date",
		"is_enabled",
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

	varBusinessConnection := _BusinessConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBusinessConnection)

	if err != nil {
		return err
	}

	*o = BusinessConnection(varBusinessConnection)

	return err
}

type NullableBusinessConnection struct {
	value *BusinessConnection
	isSet bool
}

func (v NullableBusinessConnection) Get() *BusinessConnection {
	return v.value
}

func (v *NullableBusinessConnection) Set(val *BusinessConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessConnection(val *BusinessConnection) *NullableBusinessConnection {
	return &NullableBusinessConnection{value: val, isSet: true}
}

func (v NullableBusinessConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


