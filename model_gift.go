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

// checks if the Gift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gift{}

// Gift This object represents a gift that can be sent by the bot.
type Gift struct {
	// Unique identifier of the gift
	Id string `json:"id"`
	Sticker Sticker `json:"sticker"`
	// The number of Telegram Stars that must be paid to send the sticker
	StarCount int32 `json:"star_count"`
	// *Optional*. The number of Telegram Stars that must be paid to upgrade the gift to a unique one
	UpgradeStarCount *int32 `json:"upgrade_star_count,omitempty"`
	// *Optional*. The total number of the gifts of this type that can be sent; for limited gifts only
	TotalCount *int32 `json:"total_count,omitempty"`
	// *Optional*. The number of remaining gifts of this type that can be sent; for limited gifts only
	RemainingCount *int32 `json:"remaining_count,omitempty"`
}

type _Gift Gift

// NewGift instantiates a new Gift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGift(id string, sticker Sticker, starCount int32) *Gift {
	this := Gift{}
	this.Id = id
	this.Sticker = sticker
	this.StarCount = starCount
	return &this
}

// NewGiftWithDefaults instantiates a new Gift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftWithDefaults() *Gift {
	this := Gift{}
	return &this
}

// GetId returns the Id field value
func (o *Gift) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Gift) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Gift) SetId(v string) {
	o.Id = v
}

// GetSticker returns the Sticker field value
func (o *Gift) GetSticker() Sticker {
	if o == nil {
		var ret Sticker
		return ret
	}

	return o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value
// and a boolean to check if the value has been set.
func (o *Gift) GetStickerOk() (*Sticker, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticker, true
}

// SetSticker sets field value
func (o *Gift) SetSticker(v Sticker) {
	o.Sticker = v
}

// GetStarCount returns the StarCount field value
func (o *Gift) GetStarCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StarCount
}

// GetStarCountOk returns a tuple with the StarCount field value
// and a boolean to check if the value has been set.
func (o *Gift) GetStarCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarCount, true
}

// SetStarCount sets field value
func (o *Gift) SetStarCount(v int32) {
	o.StarCount = v
}

// GetUpgradeStarCount returns the UpgradeStarCount field value if set, zero value otherwise.
func (o *Gift) GetUpgradeStarCount() int32 {
	if o == nil || IsNil(o.UpgradeStarCount) {
		var ret int32
		return ret
	}
	return *o.UpgradeStarCount
}

// GetUpgradeStarCountOk returns a tuple with the UpgradeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetUpgradeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UpgradeStarCount) {
		return nil, false
	}
	return o.UpgradeStarCount, true
}

// HasUpgradeStarCount returns a boolean if a field has been set.
func (o *Gift) HasUpgradeStarCount() bool {
	if o != nil && !IsNil(o.UpgradeStarCount) {
		return true
	}

	return false
}

// SetUpgradeStarCount gets a reference to the given int32 and assigns it to the UpgradeStarCount field.
func (o *Gift) SetUpgradeStarCount(v int32) {
	o.UpgradeStarCount = &v
}


// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Gift) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Gift) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *Gift) SetTotalCount(v int32) {
	o.TotalCount = &v
}


// GetRemainingCount returns the RemainingCount field value if set, zero value otherwise.
func (o *Gift) GetRemainingCount() int32 {
	if o == nil || IsNil(o.RemainingCount) {
		var ret int32
		return ret
	}
	return *o.RemainingCount
}

// GetRemainingCountOk returns a tuple with the RemainingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetRemainingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainingCount) {
		return nil, false
	}
	return o.RemainingCount, true
}

// HasRemainingCount returns a boolean if a field has been set.
func (o *Gift) HasRemainingCount() bool {
	if o != nil && !IsNil(o.RemainingCount) {
		return true
	}

	return false
}

// SetRemainingCount gets a reference to the given int32 and assigns it to the RemainingCount field.
func (o *Gift) SetRemainingCount(v int32) {
	o.RemainingCount = &v
}


func (o Gift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["sticker"] = o.Sticker
	toSerialize["star_count"] = o.StarCount
	if !IsNil(o.UpgradeStarCount) {
		toSerialize["upgrade_star_count"] = o.UpgradeStarCount
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !IsNil(o.RemainingCount) {
		toSerialize["remaining_count"] = o.RemainingCount
	}
	return toSerialize, nil
}

func (o *Gift) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"sticker",
		"star_count",
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

	varGift := _Gift{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGift)

	if err != nil {
		return err
	}

	*o = Gift(varGift)

	return err
}

type NullableGift struct {
	value *Gift
	isSet bool
}

func (v NullableGift) Get() *Gift {
	return v.value
}

func (v *NullableGift) Set(val *Gift) {
	v.value = val
	v.isSet = true
}

func (v NullableGift) IsSet() bool {
	return v.isSet
}

func (v *NullableGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGift(val *Gift) *NullableGift {
	return &NullableGift{value: val, isSet: true}
}

func (v NullableGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


