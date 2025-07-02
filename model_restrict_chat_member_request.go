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

// checks if the RestrictChatMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictChatMemberRequest{}

// RestrictChatMemberRequest Request parameters for restrictChatMember
type RestrictChatMemberRequest struct {
	ChatId BotCommandScopeChatChatId `json:"chat_id"`
	// Unique identifier of the target user
	UserId int32 `json:"user_id"`
	Permissions ChatPermissions `json:"permissions"`
	// Pass *True* if chat permissions are set independently. Otherwise, the *can\\_send\\_other\\_messages* and *can\\_add\\_web\\_page\\_previews* permissions will imply the *can\\_send\\_messages*, *can\\_send\\_audios*, *can\\_send\\_documents*, *can\\_send\\_photos*, *can\\_send\\_videos*, *can\\_send\\_video\\_notes*, and *can\\_send\\_voice\\_notes* permissions; the *can\\_send\\_polls* permission will imply the *can\\_send\\_messages* permission.
	UseIndependentChatPermissions *bool `json:"use_independent_chat_permissions,omitempty"`
	// Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilDate *int32 `json:"until_date,omitempty"`
}

type _RestrictChatMemberRequest RestrictChatMemberRequest

// NewRestrictChatMemberRequest instantiates a new RestrictChatMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictChatMemberRequest(chatId BotCommandScopeChatChatId, userId int32, permissions ChatPermissions) *RestrictChatMemberRequest {
	this := RestrictChatMemberRequest{}
	this.ChatId = chatId
	this.UserId = userId
	this.Permissions = permissions
	return &this
}

// NewRestrictChatMemberRequestWithDefaults instantiates a new RestrictChatMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictChatMemberRequestWithDefaults() *RestrictChatMemberRequest {
	this := RestrictChatMemberRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *RestrictChatMemberRequest) GetChatId() BotCommandScopeChatChatId {
	if o == nil {
		var ret BotCommandScopeChatChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *RestrictChatMemberRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *RestrictChatMemberRequest) SetChatId(v BotCommandScopeChatChatId) {
	o.ChatId = v
}

// GetUserId returns the UserId field value
func (o *RestrictChatMemberRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *RestrictChatMemberRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *RestrictChatMemberRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetPermissions returns the Permissions field value
func (o *RestrictChatMemberRequest) GetPermissions() ChatPermissions {
	if o == nil {
		var ret ChatPermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *RestrictChatMemberRequest) GetPermissionsOk() (*ChatPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *RestrictChatMemberRequest) SetPermissions(v ChatPermissions) {
	o.Permissions = v
}

// GetUseIndependentChatPermissions returns the UseIndependentChatPermissions field value if set, zero value otherwise.
func (o *RestrictChatMemberRequest) GetUseIndependentChatPermissions() bool {
	if o == nil || IsNil(o.UseIndependentChatPermissions) {
		var ret bool
		return ret
	}
	return *o.UseIndependentChatPermissions
}

// GetUseIndependentChatPermissionsOk returns a tuple with the UseIndependentChatPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictChatMemberRequest) GetUseIndependentChatPermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseIndependentChatPermissions) {
		return nil, false
	}
	return o.UseIndependentChatPermissions, true
}

// HasUseIndependentChatPermissions returns a boolean if a field has been set.
func (o *RestrictChatMemberRequest) HasUseIndependentChatPermissions() bool {
	if o != nil && !IsNil(o.UseIndependentChatPermissions) {
		return true
	}

	return false
}

// SetUseIndependentChatPermissions gets a reference to the given bool and assigns it to the UseIndependentChatPermissions field.
func (o *RestrictChatMemberRequest) SetUseIndependentChatPermissions(v bool) {
	o.UseIndependentChatPermissions = &v
}


// GetUntilDate returns the UntilDate field value if set, zero value otherwise.
func (o *RestrictChatMemberRequest) GetUntilDate() int32 {
	if o == nil || IsNil(o.UntilDate) {
		var ret int32
		return ret
	}
	return *o.UntilDate
}

// GetUntilDateOk returns a tuple with the UntilDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictChatMemberRequest) GetUntilDateOk() (*int32, bool) {
	if o == nil || IsNil(o.UntilDate) {
		return nil, false
	}
	return o.UntilDate, true
}

// HasUntilDate returns a boolean if a field has been set.
func (o *RestrictChatMemberRequest) HasUntilDate() bool {
	if o != nil && !IsNil(o.UntilDate) {
		return true
	}

	return false
}

// SetUntilDate gets a reference to the given int32 and assigns it to the UntilDate field.
func (o *RestrictChatMemberRequest) SetUntilDate(v int32) {
	o.UntilDate = &v
}


func (o RestrictChatMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestrictChatMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["user_id"] = o.UserId
	toSerialize["permissions"] = o.Permissions
	if !IsNil(o.UseIndependentChatPermissions) {
		toSerialize["use_independent_chat_permissions"] = o.UseIndependentChatPermissions
	}
	if !IsNil(o.UntilDate) {
		toSerialize["until_date"] = o.UntilDate
	}
	return toSerialize, nil
}

func (o *RestrictChatMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"user_id",
		"permissions",
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

	varRestrictChatMemberRequest := _RestrictChatMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestrictChatMemberRequest)

	if err != nil {
		return err
	}

	*o = RestrictChatMemberRequest(varRestrictChatMemberRequest)

	return err
}

type NullableRestrictChatMemberRequest struct {
	value *RestrictChatMemberRequest
	isSet bool
}

func (v NullableRestrictChatMemberRequest) Get() *RestrictChatMemberRequest {
	return v.value
}

func (v *NullableRestrictChatMemberRequest) Set(val *RestrictChatMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictChatMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictChatMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictChatMemberRequest(val *RestrictChatMemberRequest) *NullableRestrictChatMemberRequest {
	return &NullableRestrictChatMemberRequest{value: val, isSet: true}
}

func (v NullableRestrictChatMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictChatMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


