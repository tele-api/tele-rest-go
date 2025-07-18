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

// checks if the UniqueGiftInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueGiftInfo{}

// UniqueGiftInfo Describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	Gift UniqueGift `json:"gift"`
	// Origin of the gift. Currently, either “upgrade” for gifts upgraded from regular gifts, “transfer” for gifts transferred from other users or channels, or “resale” for gifts bought from other users
	Origin string `json:"origin"`
	// *Optional*. For gifts bought from other users, the price paid for the gift
	LastResaleStarCount *int32 `json:"last_resale_star_count,omitempty"`
	// *Optional*. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
	OwnedGiftId *string `json:"owned_gift_id,omitempty"`
	// *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift
	TransferStarCount *int32 `json:"transfer_star_count,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the gift can be transferred. If it is in the past, then the gift can be transferred now
	NextTransferDate *int32 `json:"next_transfer_date,omitempty"`
}

type _UniqueGiftInfo UniqueGiftInfo

// NewUniqueGiftInfo instantiates a new UniqueGiftInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueGiftInfo(gift UniqueGift, origin string) *UniqueGiftInfo {
	this := UniqueGiftInfo{}
	this.Gift = gift
	this.Origin = origin
	return &this
}

// NewUniqueGiftInfoWithDefaults instantiates a new UniqueGiftInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueGiftInfoWithDefaults() *UniqueGiftInfo {
	this := UniqueGiftInfo{}
	return &this
}

// GetGift returns the Gift field value
func (o *UniqueGiftInfo) GetGift() UniqueGift {
	if o == nil {
		var ret UniqueGift
		return ret
	}

	return o.Gift
}

// GetGiftOk returns a tuple with the Gift field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetGiftOk() (*UniqueGift, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gift, true
}

// SetGift sets field value
func (o *UniqueGiftInfo) SetGift(v UniqueGift) {
	o.Gift = v
}

// GetOrigin returns the Origin field value
func (o *UniqueGiftInfo) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *UniqueGiftInfo) SetOrigin(v string) {
	o.Origin = v
}

// GetLastResaleStarCount returns the LastResaleStarCount field value if set, zero value otherwise.
func (o *UniqueGiftInfo) GetLastResaleStarCount() int32 {
	if o == nil || IsNil(o.LastResaleStarCount) {
		var ret int32
		return ret
	}
	return *o.LastResaleStarCount
}

// GetLastResaleStarCountOk returns a tuple with the LastResaleStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetLastResaleStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LastResaleStarCount) {
		return nil, false
	}
	return o.LastResaleStarCount, true
}

// HasLastResaleStarCount returns a boolean if a field has been set.
func (o *UniqueGiftInfo) HasLastResaleStarCount() bool {
	if o != nil && !IsNil(o.LastResaleStarCount) {
		return true
	}

	return false
}

// SetLastResaleStarCount gets a reference to the given int32 and assigns it to the LastResaleStarCount field.
func (o *UniqueGiftInfo) SetLastResaleStarCount(v int32) {
	o.LastResaleStarCount = &v
}


// GetOwnedGiftId returns the OwnedGiftId field value if set, zero value otherwise.
func (o *UniqueGiftInfo) GetOwnedGiftId() string {
	if o == nil || IsNil(o.OwnedGiftId) {
		var ret string
		return ret
	}
	return *o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedGiftId) {
		return nil, false
	}
	return o.OwnedGiftId, true
}

// HasOwnedGiftId returns a boolean if a field has been set.
func (o *UniqueGiftInfo) HasOwnedGiftId() bool {
	if o != nil && !IsNil(o.OwnedGiftId) {
		return true
	}

	return false
}

// SetOwnedGiftId gets a reference to the given string and assigns it to the OwnedGiftId field.
func (o *UniqueGiftInfo) SetOwnedGiftId(v string) {
	o.OwnedGiftId = &v
}


// GetTransferStarCount returns the TransferStarCount field value if set, zero value otherwise.
func (o *UniqueGiftInfo) GetTransferStarCount() int32 {
	if o == nil || IsNil(o.TransferStarCount) {
		var ret int32
		return ret
	}
	return *o.TransferStarCount
}

// GetTransferStarCountOk returns a tuple with the TransferStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetTransferStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransferStarCount) {
		return nil, false
	}
	return o.TransferStarCount, true
}

// HasTransferStarCount returns a boolean if a field has been set.
func (o *UniqueGiftInfo) HasTransferStarCount() bool {
	if o != nil && !IsNil(o.TransferStarCount) {
		return true
	}

	return false
}

// SetTransferStarCount gets a reference to the given int32 and assigns it to the TransferStarCount field.
func (o *UniqueGiftInfo) SetTransferStarCount(v int32) {
	o.TransferStarCount = &v
}


// GetNextTransferDate returns the NextTransferDate field value if set, zero value otherwise.
func (o *UniqueGiftInfo) GetNextTransferDate() int32 {
	if o == nil || IsNil(o.NextTransferDate) {
		var ret int32
		return ret
	}
	return *o.NextTransferDate
}

// GetNextTransferDateOk returns a tuple with the NextTransferDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueGiftInfo) GetNextTransferDateOk() (*int32, bool) {
	if o == nil || IsNil(o.NextTransferDate) {
		return nil, false
	}
	return o.NextTransferDate, true
}

// HasNextTransferDate returns a boolean if a field has been set.
func (o *UniqueGiftInfo) HasNextTransferDate() bool {
	if o != nil && !IsNil(o.NextTransferDate) {
		return true
	}

	return false
}

// SetNextTransferDate gets a reference to the given int32 and assigns it to the NextTransferDate field.
func (o *UniqueGiftInfo) SetNextTransferDate(v int32) {
	o.NextTransferDate = &v
}


func (o UniqueGiftInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueGiftInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gift"] = o.Gift
	toSerialize["origin"] = o.Origin
	if !IsNil(o.LastResaleStarCount) {
		toSerialize["last_resale_star_count"] = o.LastResaleStarCount
	}
	if !IsNil(o.OwnedGiftId) {
		toSerialize["owned_gift_id"] = o.OwnedGiftId
	}
	if !IsNil(o.TransferStarCount) {
		toSerialize["transfer_star_count"] = o.TransferStarCount
	}
	if !IsNil(o.NextTransferDate) {
		toSerialize["next_transfer_date"] = o.NextTransferDate
	}
	return toSerialize, nil
}

func (o *UniqueGiftInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gift",
		"origin",
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

	varUniqueGiftInfo := _UniqueGiftInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniqueGiftInfo)

	if err != nil {
		return err
	}

	*o = UniqueGiftInfo(varUniqueGiftInfo)

	return err
}

type NullableUniqueGiftInfo struct {
	value *UniqueGiftInfo
	isSet bool
}

func (v NullableUniqueGiftInfo) Get() *UniqueGiftInfo {
	return v.value
}

func (v *NullableUniqueGiftInfo) Set(val *UniqueGiftInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueGiftInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueGiftInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueGiftInfo(val *UniqueGiftInfo) *NullableUniqueGiftInfo {
	return &NullableUniqueGiftInfo{value: val, isSet: true}
}

func (v NullableUniqueGiftInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueGiftInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


