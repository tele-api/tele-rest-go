/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the GiftInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftInfo{}

// GiftInfo Describes a service message about a regular gift that was sent or received.
type GiftInfo struct {
	Gift Gift `json:"gift"`
	// *Optional*. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
	OwnedGiftId *string `json:"owned_gift_id,omitempty"`
	// *Optional*. Number of Telegram Stars that can be claimed by the receiver by converting the gift; omitted if conversion to Telegram Stars is impossible
	ConvertStarCount *int32 `json:"convert_star_count,omitempty"`
	// *Optional*. Number of Telegram Stars that were prepaid by the sender for the ability to upgrade the gift
	PrepaidUpgradeStarCount *int32 `json:"prepaid_upgrade_star_count,omitempty"`
	// *Optional*. True, if the gift can be upgraded to a unique gift
	CanBeUpgraded *bool `json:"can_be_upgraded,omitempty"`
	// *Optional*. Text of the message that was added to the gift
	Text *string `json:"text,omitempty"`
	// *Optional*. Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`
	// *Optional*. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
	IsPrivate *bool `json:"is_private,omitempty"`
}

type _GiftInfo GiftInfo

// NewGiftInfo instantiates a new GiftInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftInfo(gift Gift) *GiftInfo {
	this := GiftInfo{}
	this.Gift = gift
	var canBeUpgraded bool = true
	this.CanBeUpgraded = &canBeUpgraded
	var isPrivate bool = true
	this.IsPrivate = &isPrivate
	return &this
}

// NewGiftInfoWithDefaults instantiates a new GiftInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftInfoWithDefaults() *GiftInfo {
	this := GiftInfo{}
	var canBeUpgraded bool = true
	this.CanBeUpgraded = &canBeUpgraded
	var isPrivate bool = true
	this.IsPrivate = &isPrivate
	return &this
}

// GetGift returns the Gift field value
func (o *GiftInfo) GetGift() Gift {
	if o == nil {
		var ret Gift
		return ret
	}

	return o.Gift
}

// GetGiftOk returns a tuple with the Gift field value
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetGiftOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gift, true
}

// SetGift sets field value
func (o *GiftInfo) SetGift(v Gift) {
	o.Gift = v
}

// GetOwnedGiftId returns the OwnedGiftId field value if set, zero value otherwise.
func (o *GiftInfo) GetOwnedGiftId() string {
	if o == nil || IsNil(o.OwnedGiftId) {
		var ret string
		return ret
	}
	return *o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedGiftId) {
		return nil, false
	}
	return o.OwnedGiftId, true
}

// HasOwnedGiftId returns a boolean if a field has been set.
func (o *GiftInfo) HasOwnedGiftId() bool {
	if o != nil && !IsNil(o.OwnedGiftId) {
		return true
	}

	return false
}

// SetOwnedGiftId gets a reference to the given string and assigns it to the OwnedGiftId field.
func (o *GiftInfo) SetOwnedGiftId(v string) {
	o.OwnedGiftId = &v
}


// GetConvertStarCount returns the ConvertStarCount field value if set, zero value otherwise.
func (o *GiftInfo) GetConvertStarCount() int32 {
	if o == nil || IsNil(o.ConvertStarCount) {
		var ret int32
		return ret
	}
	return *o.ConvertStarCount
}

// GetConvertStarCountOk returns a tuple with the ConvertStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetConvertStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConvertStarCount) {
		return nil, false
	}
	return o.ConvertStarCount, true
}

// HasConvertStarCount returns a boolean if a field has been set.
func (o *GiftInfo) HasConvertStarCount() bool {
	if o != nil && !IsNil(o.ConvertStarCount) {
		return true
	}

	return false
}

// SetConvertStarCount gets a reference to the given int32 and assigns it to the ConvertStarCount field.
func (o *GiftInfo) SetConvertStarCount(v int32) {
	o.ConvertStarCount = &v
}


// GetPrepaidUpgradeStarCount returns the PrepaidUpgradeStarCount field value if set, zero value otherwise.
func (o *GiftInfo) GetPrepaidUpgradeStarCount() int32 {
	if o == nil || IsNil(o.PrepaidUpgradeStarCount) {
		var ret int32
		return ret
	}
	return *o.PrepaidUpgradeStarCount
}

// GetPrepaidUpgradeStarCountOk returns a tuple with the PrepaidUpgradeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetPrepaidUpgradeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrepaidUpgradeStarCount) {
		return nil, false
	}
	return o.PrepaidUpgradeStarCount, true
}

// HasPrepaidUpgradeStarCount returns a boolean if a field has been set.
func (o *GiftInfo) HasPrepaidUpgradeStarCount() bool {
	if o != nil && !IsNil(o.PrepaidUpgradeStarCount) {
		return true
	}

	return false
}

// SetPrepaidUpgradeStarCount gets a reference to the given int32 and assigns it to the PrepaidUpgradeStarCount field.
func (o *GiftInfo) SetPrepaidUpgradeStarCount(v int32) {
	o.PrepaidUpgradeStarCount = &v
}


// GetCanBeUpgraded returns the CanBeUpgraded field value if set, zero value otherwise.
func (o *GiftInfo) GetCanBeUpgraded() bool {
	if o == nil || IsNil(o.CanBeUpgraded) {
		var ret bool
		return ret
	}
	return *o.CanBeUpgraded
}

// GetCanBeUpgradedOk returns a tuple with the CanBeUpgraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetCanBeUpgradedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeUpgraded) {
		return nil, false
	}
	return o.CanBeUpgraded, true
}

// HasCanBeUpgraded returns a boolean if a field has been set.
func (o *GiftInfo) HasCanBeUpgraded() bool {
	if o != nil && !IsNil(o.CanBeUpgraded) {
		return true
	}

	return false
}

// SetCanBeUpgraded gets a reference to the given bool and assigns it to the CanBeUpgraded field.
func (o *GiftInfo) SetCanBeUpgraded(v bool) {
	o.CanBeUpgraded = &v
}


// GetText returns the Text field value if set, zero value otherwise.
func (o *GiftInfo) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GiftInfo) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *GiftInfo) SetText(v string) {
	o.Text = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *GiftInfo) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *GiftInfo) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *GiftInfo) SetEntities(v []MessageEntity) {
	o.Entities = v
}


// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *GiftInfo) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftInfo) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *GiftInfo) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *GiftInfo) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}


func (o GiftInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gift"] = o.Gift
	if !IsNil(o.OwnedGiftId) {
		toSerialize["owned_gift_id"] = o.OwnedGiftId
	}
	if !IsNil(o.ConvertStarCount) {
		toSerialize["convert_star_count"] = o.ConvertStarCount
	}
	if !IsNil(o.PrepaidUpgradeStarCount) {
		toSerialize["prepaid_upgrade_star_count"] = o.PrepaidUpgradeStarCount
	}
	if !IsNil(o.CanBeUpgraded) {
		toSerialize["can_be_upgraded"] = o.CanBeUpgraded
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}
	return toSerialize, nil
}

func (o *GiftInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gift",
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

	varGiftInfo := _GiftInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiftInfo)

	if err != nil {
		return err
	}

	*o = GiftInfo(varGiftInfo)

	return err
}

type NullableGiftInfo struct {
	value *GiftInfo
	isSet bool
}

func (v NullableGiftInfo) Get() *GiftInfo {
	return v.value
}

func (v *NullableGiftInfo) Set(val *GiftInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftInfo(val *GiftInfo) *NullableGiftInfo {
	return &NullableGiftInfo{value: val, isSet: true}
}

func (v NullableGiftInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


