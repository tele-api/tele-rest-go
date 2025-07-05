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

// checks if the BanChatMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BanChatMemberRequest{}

// BanChatMemberRequest Request parameters for banChatMember
type BanChatMemberRequest struct {
	ChatId BanChatMemberRequestChatId `json:"chat_id"`
	// Unique identifier of the target user
	UserId int32 `json:"user_id"`
	// Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	UntilDate *int32 `json:"until_date,omitempty"`
	// Pass *True* to delete all messages from the chat for the user that is being removed. If *False*, the user will be able to see messages in the group that were sent before the user was removed. Always *True* for supergroups and channels.
	RevokeMessages *bool `json:"revoke_messages,omitempty"`
}

type _BanChatMemberRequest BanChatMemberRequest

// NewBanChatMemberRequest instantiates a new BanChatMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanChatMemberRequest(chatId BanChatMemberRequestChatId, userId int32) *BanChatMemberRequest {
	this := BanChatMemberRequest{}
	this.ChatId = chatId
	this.UserId = userId
	return &this
}

// NewBanChatMemberRequestWithDefaults instantiates a new BanChatMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanChatMemberRequestWithDefaults() *BanChatMemberRequest {
	this := BanChatMemberRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *BanChatMemberRequest) GetChatId() BanChatMemberRequestChatId {
	if o == nil {
		var ret BanChatMemberRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *BanChatMemberRequest) GetChatIdOk() (*BanChatMemberRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *BanChatMemberRequest) SetChatId(v BanChatMemberRequestChatId) {
	o.ChatId = v
}

// GetUserId returns the UserId field value
func (o *BanChatMemberRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *BanChatMemberRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *BanChatMemberRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetUntilDate returns the UntilDate field value if set, zero value otherwise.
func (o *BanChatMemberRequest) GetUntilDate() int32 {
	if o == nil || IsNil(o.UntilDate) {
		var ret int32
		return ret
	}
	return *o.UntilDate
}

// GetUntilDateOk returns a tuple with the UntilDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanChatMemberRequest) GetUntilDateOk() (*int32, bool) {
	if o == nil || IsNil(o.UntilDate) {
		return nil, false
	}
	return o.UntilDate, true
}

// HasUntilDate returns a boolean if a field has been set.
func (o *BanChatMemberRequest) HasUntilDate() bool {
	if o != nil && !IsNil(o.UntilDate) {
		return true
	}

	return false
}

// SetUntilDate gets a reference to the given int32 and assigns it to the UntilDate field.
func (o *BanChatMemberRequest) SetUntilDate(v int32) {
	o.UntilDate = &v
}


// GetRevokeMessages returns the RevokeMessages field value if set, zero value otherwise.
func (o *BanChatMemberRequest) GetRevokeMessages() bool {
	if o == nil || IsNil(o.RevokeMessages) {
		var ret bool
		return ret
	}
	return *o.RevokeMessages
}

// GetRevokeMessagesOk returns a tuple with the RevokeMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanChatMemberRequest) GetRevokeMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.RevokeMessages) {
		return nil, false
	}
	return o.RevokeMessages, true
}

// HasRevokeMessages returns a boolean if a field has been set.
func (o *BanChatMemberRequest) HasRevokeMessages() bool {
	if o != nil && !IsNil(o.RevokeMessages) {
		return true
	}

	return false
}

// SetRevokeMessages gets a reference to the given bool and assigns it to the RevokeMessages field.
func (o *BanChatMemberRequest) SetRevokeMessages(v bool) {
	o.RevokeMessages = &v
}


func (o BanChatMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BanChatMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.UntilDate) {
		toSerialize["until_date"] = o.UntilDate
	}
	if !IsNil(o.RevokeMessages) {
		toSerialize["revoke_messages"] = o.RevokeMessages
	}
	return toSerialize, nil
}

func (o *BanChatMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"user_id",
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

	varBanChatMemberRequest := _BanChatMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBanChatMemberRequest)

	if err != nil {
		return err
	}

	*o = BanChatMemberRequest(varBanChatMemberRequest)

	return err
}

type NullableBanChatMemberRequest struct {
	value *BanChatMemberRequest
	isSet bool
}

func (v NullableBanChatMemberRequest) Get() *BanChatMemberRequest {
	return v.value
}

func (v *NullableBanChatMemberRequest) Set(val *BanChatMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBanChatMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBanChatMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanChatMemberRequest(val *BanChatMemberRequest) *NullableBanChatMemberRequest {
	return &NullableBanChatMemberRequest{value: val, isSet: true}
}

func (v NullableBanChatMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanChatMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


