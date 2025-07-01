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

// checks if the OwnedGiftUnique type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnedGiftUnique{}

// OwnedGiftUnique Describes a unique gift received and owned by a user or a chat.
type OwnedGiftUnique struct {
	// Type of the gift, always “unique”
	Type string `json:"type"`
	Gift UniqueGift `json:"gift"`
	// *Optional*. Unique identifier of the received gift for the bot; for gifts received on behalf of business accounts only
	OwnedGiftId *string `json:"owned_gift_id,omitempty"`
	SenderUser *User `json:"sender_user,omitempty"`
	// Date the gift was sent in Unix time
	SendDate int32 `json:"send_date"`
	// *Optional*. True, if the gift is displayed on the account's profile page; for gifts received on behalf of business accounts only
	IsSaved *bool `json:"is_saved,omitempty"`
	// *Optional*. True, if the gift can be transferred to another owner; for gifts received on behalf of business accounts only
	CanBeTransferred *bool `json:"can_be_transferred,omitempty"`
	// *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift
	TransferStarCount *int32 `json:"transfer_star_count,omitempty"`
}

type _OwnedGiftUnique OwnedGiftUnique

// NewOwnedGiftUnique instantiates a new OwnedGiftUnique object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnedGiftUnique(type_ string, gift UniqueGift, sendDate int32) *OwnedGiftUnique {
	this := OwnedGiftUnique{}
	this.Type = type_
	this.Gift = gift
	this.SendDate = sendDate
	var isSaved bool = true
	this.IsSaved = &isSaved
	var canBeTransferred bool = true
	this.CanBeTransferred = &canBeTransferred
	return &this
}

// NewOwnedGiftUniqueWithDefaults instantiates a new OwnedGiftUnique object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnedGiftUniqueWithDefaults() *OwnedGiftUnique {
	this := OwnedGiftUnique{}
	var type_ string = "unique"
	this.Type = type_
	var isSaved bool = true
	this.IsSaved = &isSaved
	var canBeTransferred bool = true
	this.CanBeTransferred = &canBeTransferred
	return &this
}

// GetType returns the Type field value
func (o *OwnedGiftUnique) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OwnedGiftUnique) SetType(v string) {
	o.Type = v
}

// GetGift returns the Gift field value
func (o *OwnedGiftUnique) GetGift() UniqueGift {
	if o == nil {
		var ret UniqueGift
		return ret
	}

	return o.Gift
}

// GetGiftOk returns a tuple with the Gift field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetGiftOk() (*UniqueGift, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gift, true
}

// SetGift sets field value
func (o *OwnedGiftUnique) SetGift(v UniqueGift) {
	o.Gift = v
}

// GetOwnedGiftId returns the OwnedGiftId field value if set, zero value otherwise.
func (o *OwnedGiftUnique) GetOwnedGiftId() string {
	if o == nil || IsNil(o.OwnedGiftId) {
		var ret string
		return ret
	}
	return *o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedGiftId) {
		return nil, false
	}
	return o.OwnedGiftId, true
}

// HasOwnedGiftId returns a boolean if a field has been set.
func (o *OwnedGiftUnique) HasOwnedGiftId() bool {
	if o != nil && !IsNil(o.OwnedGiftId) {
		return true
	}

	return false
}

// SetOwnedGiftId gets a reference to the given string and assigns it to the OwnedGiftId field.
func (o *OwnedGiftUnique) SetOwnedGiftId(v string) {
	o.OwnedGiftId = &v
}


// GetSenderUser returns the SenderUser field value if set, zero value otherwise.
func (o *OwnedGiftUnique) GetSenderUser() User {
	if o == nil || IsNil(o.SenderUser) {
		var ret User
		return ret
	}
	return *o.SenderUser
}

// GetSenderUserOk returns a tuple with the SenderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetSenderUserOk() (*User, bool) {
	if o == nil || IsNil(o.SenderUser) {
		return nil, false
	}
	return o.SenderUser, true
}

// HasSenderUser returns a boolean if a field has been set.
func (o *OwnedGiftUnique) HasSenderUser() bool {
	if o != nil && !IsNil(o.SenderUser) {
		return true
	}

	return false
}

// SetSenderUser gets a reference to the given User and assigns it to the SenderUser field.
func (o *OwnedGiftUnique) SetSenderUser(v User) {
	o.SenderUser = &v
}


// GetSendDate returns the SendDate field value
func (o *OwnedGiftUnique) GetSendDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetSendDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendDate, true
}

// SetSendDate sets field value
func (o *OwnedGiftUnique) SetSendDate(v int32) {
	o.SendDate = v
}

// GetIsSaved returns the IsSaved field value if set, zero value otherwise.
func (o *OwnedGiftUnique) GetIsSaved() bool {
	if o == nil || IsNil(o.IsSaved) {
		var ret bool
		return ret
	}
	return *o.IsSaved
}

// GetIsSavedOk returns a tuple with the IsSaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetIsSavedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSaved) {
		return nil, false
	}
	return o.IsSaved, true
}

// HasIsSaved returns a boolean if a field has been set.
func (o *OwnedGiftUnique) HasIsSaved() bool {
	if o != nil && !IsNil(o.IsSaved) {
		return true
	}

	return false
}

// SetIsSaved gets a reference to the given bool and assigns it to the IsSaved field.
func (o *OwnedGiftUnique) SetIsSaved(v bool) {
	o.IsSaved = &v
}


// GetCanBeTransferred returns the CanBeTransferred field value if set, zero value otherwise.
func (o *OwnedGiftUnique) GetCanBeTransferred() bool {
	if o == nil || IsNil(o.CanBeTransferred) {
		var ret bool
		return ret
	}
	return *o.CanBeTransferred
}

// GetCanBeTransferredOk returns a tuple with the CanBeTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetCanBeTransferredOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeTransferred) {
		return nil, false
	}
	return o.CanBeTransferred, true
}

// HasCanBeTransferred returns a boolean if a field has been set.
func (o *OwnedGiftUnique) HasCanBeTransferred() bool {
	if o != nil && !IsNil(o.CanBeTransferred) {
		return true
	}

	return false
}

// SetCanBeTransferred gets a reference to the given bool and assigns it to the CanBeTransferred field.
func (o *OwnedGiftUnique) SetCanBeTransferred(v bool) {
	o.CanBeTransferred = &v
}


// GetTransferStarCount returns the TransferStarCount field value if set, zero value otherwise.
func (o *OwnedGiftUnique) GetTransferStarCount() int32 {
	if o == nil || IsNil(o.TransferStarCount) {
		var ret int32
		return ret
	}
	return *o.TransferStarCount
}

// GetTransferStarCountOk returns a tuple with the TransferStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftUnique) GetTransferStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransferStarCount) {
		return nil, false
	}
	return o.TransferStarCount, true
}

// HasTransferStarCount returns a boolean if a field has been set.
func (o *OwnedGiftUnique) HasTransferStarCount() bool {
	if o != nil && !IsNil(o.TransferStarCount) {
		return true
	}

	return false
}

// SetTransferStarCount gets a reference to the given int32 and assigns it to the TransferStarCount field.
func (o *OwnedGiftUnique) SetTransferStarCount(v int32) {
	o.TransferStarCount = &v
}


func (o OwnedGiftUnique) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnedGiftUnique) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["gift"] = o.Gift
	if !IsNil(o.OwnedGiftId) {
		toSerialize["owned_gift_id"] = o.OwnedGiftId
	}
	if !IsNil(o.SenderUser) {
		toSerialize["sender_user"] = o.SenderUser
	}
	toSerialize["send_date"] = o.SendDate
	if !IsNil(o.IsSaved) {
		toSerialize["is_saved"] = o.IsSaved
	}
	if !IsNil(o.CanBeTransferred) {
		toSerialize["can_be_transferred"] = o.CanBeTransferred
	}
	if !IsNil(o.TransferStarCount) {
		toSerialize["transfer_star_count"] = o.TransferStarCount
	}
	return toSerialize, nil
}

func (o *OwnedGiftUnique) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"gift",
		"send_date",
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

	varOwnedGiftUnique := _OwnedGiftUnique{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOwnedGiftUnique)

	if err != nil {
		return err
	}

	*o = OwnedGiftUnique(varOwnedGiftUnique)

	return err
}

type NullableOwnedGiftUnique struct {
	value *OwnedGiftUnique
	isSet bool
}

func (v NullableOwnedGiftUnique) Get() *OwnedGiftUnique {
	return v.value
}

func (v *NullableOwnedGiftUnique) Set(val *OwnedGiftUnique) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnedGiftUnique) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnedGiftUnique) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnedGiftUnique(val *OwnedGiftUnique) *NullableOwnedGiftUnique {
	return &NullableOwnedGiftUnique{value: val, isSet: true}
}

func (v NullableOwnedGiftUnique) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnedGiftUnique) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


