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

// checks if the OwnedGiftRegular type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnedGiftRegular{}

// OwnedGiftRegular Describes a regular gift owned by a user or a chat.
type OwnedGiftRegular struct {
	// Type of the gift, always “regular”
	Type string `json:"type"`
	Gift Gift `json:"gift"`
	// *Optional*. Unique identifier of the gift for the bot; for gifts received on behalf of business accounts only
	OwnedGiftId *string `json:"owned_gift_id,omitempty"`
	SenderUser *User `json:"sender_user,omitempty"`
	// Date the gift was sent in Unix time
	SendDate int32 `json:"send_date"`
	// *Optional*. Text of the message that was added to the gift
	Text *string `json:"text,omitempty"`
	// *Optional*. Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`
	// *Optional*. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
	IsPrivate *bool `json:"is_private,omitempty"`
	// *Optional*. True, if the gift is displayed on the account's profile page; for gifts received on behalf of business accounts only
	IsSaved *bool `json:"is_saved,omitempty"`
	// *Optional*. True, if the gift can be upgraded to a unique gift; for gifts received on behalf of business accounts only
	CanBeUpgraded *bool `json:"can_be_upgraded,omitempty"`
	// *Optional*. True, if the gift was refunded and isn't available anymore
	WasRefunded *bool `json:"was_refunded,omitempty"`
	// *Optional*. Number of Telegram Stars that can be claimed by the receiver instead of the gift; omitted if the gift cannot be converted to Telegram Stars
	ConvertStarCount *int32 `json:"convert_star_count,omitempty"`
	// *Optional*. Number of Telegram Stars that were paid by the sender for the ability to upgrade the gift
	PrepaidUpgradeStarCount *int32 `json:"prepaid_upgrade_star_count,omitempty"`
}

type _OwnedGiftRegular OwnedGiftRegular

// NewOwnedGiftRegular instantiates a new OwnedGiftRegular object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnedGiftRegular(type_ string, gift Gift, sendDate int32) *OwnedGiftRegular {
	this := OwnedGiftRegular{}
	this.Type = type_
	this.Gift = gift
	this.SendDate = sendDate
	var isPrivate bool = true
	this.IsPrivate = &isPrivate
	var isSaved bool = true
	this.IsSaved = &isSaved
	var canBeUpgraded bool = true
	this.CanBeUpgraded = &canBeUpgraded
	var wasRefunded bool = true
	this.WasRefunded = &wasRefunded
	return &this
}

// NewOwnedGiftRegularWithDefaults instantiates a new OwnedGiftRegular object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnedGiftRegularWithDefaults() *OwnedGiftRegular {
	this := OwnedGiftRegular{}
	var type_ string = "regular"
	this.Type = type_
	var isPrivate bool = true
	this.IsPrivate = &isPrivate
	var isSaved bool = true
	this.IsSaved = &isSaved
	var canBeUpgraded bool = true
	this.CanBeUpgraded = &canBeUpgraded
	var wasRefunded bool = true
	this.WasRefunded = &wasRefunded
	return &this
}

// GetType returns the Type field value
func (o *OwnedGiftRegular) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OwnedGiftRegular) SetType(v string) {
	o.Type = v
}

// GetGift returns the Gift field value
func (o *OwnedGiftRegular) GetGift() Gift {
	if o == nil {
		var ret Gift
		return ret
	}

	return o.Gift
}

// GetGiftOk returns a tuple with the Gift field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetGiftOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gift, true
}

// SetGift sets field value
func (o *OwnedGiftRegular) SetGift(v Gift) {
	o.Gift = v
}

// GetOwnedGiftId returns the OwnedGiftId field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetOwnedGiftId() string {
	if o == nil || IsNil(o.OwnedGiftId) {
		var ret string
		return ret
	}
	return *o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedGiftId) {
		return nil, false
	}
	return o.OwnedGiftId, true
}

// HasOwnedGiftId returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasOwnedGiftId() bool {
	if o != nil && !IsNil(o.OwnedGiftId) {
		return true
	}

	return false
}

// SetOwnedGiftId gets a reference to the given string and assigns it to the OwnedGiftId field.
func (o *OwnedGiftRegular) SetOwnedGiftId(v string) {
	o.OwnedGiftId = &v
}


// GetSenderUser returns the SenderUser field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetSenderUser() User {
	if o == nil || IsNil(o.SenderUser) {
		var ret User
		return ret
	}
	return *o.SenderUser
}

// GetSenderUserOk returns a tuple with the SenderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetSenderUserOk() (*User, bool) {
	if o == nil || IsNil(o.SenderUser) {
		return nil, false
	}
	return o.SenderUser, true
}

// HasSenderUser returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasSenderUser() bool {
	if o != nil && !IsNil(o.SenderUser) {
		return true
	}

	return false
}

// SetSenderUser gets a reference to the given User and assigns it to the SenderUser field.
func (o *OwnedGiftRegular) SetSenderUser(v User) {
	o.SenderUser = &v
}


// GetSendDate returns the SendDate field value
func (o *OwnedGiftRegular) GetSendDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetSendDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendDate, true
}

// SetSendDate sets field value
func (o *OwnedGiftRegular) SetSendDate(v int32) {
	o.SendDate = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *OwnedGiftRegular) SetText(v string) {
	o.Text = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *OwnedGiftRegular) SetEntities(v []MessageEntity) {
	o.Entities = v
}


// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *OwnedGiftRegular) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}


// GetIsSaved returns the IsSaved field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetIsSaved() bool {
	if o == nil || IsNil(o.IsSaved) {
		var ret bool
		return ret
	}
	return *o.IsSaved
}

// GetIsSavedOk returns a tuple with the IsSaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetIsSavedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSaved) {
		return nil, false
	}
	return o.IsSaved, true
}

// HasIsSaved returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasIsSaved() bool {
	if o != nil && !IsNil(o.IsSaved) {
		return true
	}

	return false
}

// SetIsSaved gets a reference to the given bool and assigns it to the IsSaved field.
func (o *OwnedGiftRegular) SetIsSaved(v bool) {
	o.IsSaved = &v
}


// GetCanBeUpgraded returns the CanBeUpgraded field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetCanBeUpgraded() bool {
	if o == nil || IsNil(o.CanBeUpgraded) {
		var ret bool
		return ret
	}
	return *o.CanBeUpgraded
}

// GetCanBeUpgradedOk returns a tuple with the CanBeUpgraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetCanBeUpgradedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeUpgraded) {
		return nil, false
	}
	return o.CanBeUpgraded, true
}

// HasCanBeUpgraded returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasCanBeUpgraded() bool {
	if o != nil && !IsNil(o.CanBeUpgraded) {
		return true
	}

	return false
}

// SetCanBeUpgraded gets a reference to the given bool and assigns it to the CanBeUpgraded field.
func (o *OwnedGiftRegular) SetCanBeUpgraded(v bool) {
	o.CanBeUpgraded = &v
}


// GetWasRefunded returns the WasRefunded field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetWasRefunded() bool {
	if o == nil || IsNil(o.WasRefunded) {
		var ret bool
		return ret
	}
	return *o.WasRefunded
}

// GetWasRefundedOk returns a tuple with the WasRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetWasRefundedOk() (*bool, bool) {
	if o == nil || IsNil(o.WasRefunded) {
		return nil, false
	}
	return o.WasRefunded, true
}

// HasWasRefunded returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasWasRefunded() bool {
	if o != nil && !IsNil(o.WasRefunded) {
		return true
	}

	return false
}

// SetWasRefunded gets a reference to the given bool and assigns it to the WasRefunded field.
func (o *OwnedGiftRegular) SetWasRefunded(v bool) {
	o.WasRefunded = &v
}


// GetConvertStarCount returns the ConvertStarCount field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetConvertStarCount() int32 {
	if o == nil || IsNil(o.ConvertStarCount) {
		var ret int32
		return ret
	}
	return *o.ConvertStarCount
}

// GetConvertStarCountOk returns a tuple with the ConvertStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetConvertStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConvertStarCount) {
		return nil, false
	}
	return o.ConvertStarCount, true
}

// HasConvertStarCount returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasConvertStarCount() bool {
	if o != nil && !IsNil(o.ConvertStarCount) {
		return true
	}

	return false
}

// SetConvertStarCount gets a reference to the given int32 and assigns it to the ConvertStarCount field.
func (o *OwnedGiftRegular) SetConvertStarCount(v int32) {
	o.ConvertStarCount = &v
}


// GetPrepaidUpgradeStarCount returns the PrepaidUpgradeStarCount field value if set, zero value otherwise.
func (o *OwnedGiftRegular) GetPrepaidUpgradeStarCount() int32 {
	if o == nil || IsNil(o.PrepaidUpgradeStarCount) {
		var ret int32
		return ret
	}
	return *o.PrepaidUpgradeStarCount
}

// GetPrepaidUpgradeStarCountOk returns a tuple with the PrepaidUpgradeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGiftRegular) GetPrepaidUpgradeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrepaidUpgradeStarCount) {
		return nil, false
	}
	return o.PrepaidUpgradeStarCount, true
}

// HasPrepaidUpgradeStarCount returns a boolean if a field has been set.
func (o *OwnedGiftRegular) HasPrepaidUpgradeStarCount() bool {
	if o != nil && !IsNil(o.PrepaidUpgradeStarCount) {
		return true
	}

	return false
}

// SetPrepaidUpgradeStarCount gets a reference to the given int32 and assigns it to the PrepaidUpgradeStarCount field.
func (o *OwnedGiftRegular) SetPrepaidUpgradeStarCount(v int32) {
	o.PrepaidUpgradeStarCount = &v
}


func (o OwnedGiftRegular) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnedGiftRegular) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}
	if !IsNil(o.IsSaved) {
		toSerialize["is_saved"] = o.IsSaved
	}
	if !IsNil(o.CanBeUpgraded) {
		toSerialize["can_be_upgraded"] = o.CanBeUpgraded
	}
	if !IsNil(o.WasRefunded) {
		toSerialize["was_refunded"] = o.WasRefunded
	}
	if !IsNil(o.ConvertStarCount) {
		toSerialize["convert_star_count"] = o.ConvertStarCount
	}
	if !IsNil(o.PrepaidUpgradeStarCount) {
		toSerialize["prepaid_upgrade_star_count"] = o.PrepaidUpgradeStarCount
	}
	return toSerialize, nil
}

func (o *OwnedGiftRegular) UnmarshalJSON(data []byte) (err error) {
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

	varOwnedGiftRegular := _OwnedGiftRegular{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOwnedGiftRegular)

	if err != nil {
		return err
	}

	*o = OwnedGiftRegular(varOwnedGiftRegular)

	return err
}

type NullableOwnedGiftRegular struct {
	value *OwnedGiftRegular
	isSet bool
}

func (v NullableOwnedGiftRegular) Get() *OwnedGiftRegular {
	return v.value
}

func (v *NullableOwnedGiftRegular) Set(val *OwnedGiftRegular) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnedGiftRegular) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnedGiftRegular) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnedGiftRegular(val *OwnedGiftRegular) *NullableOwnedGiftRegular {
	return &NullableOwnedGiftRegular{value: val, isSet: true}
}

func (v NullableOwnedGiftRegular) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnedGiftRegular) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


