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

// checks if the GiveawayCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiveawayCompleted{}

// GiveawayCompleted This object represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	// Number of winners in the giveaway
	WinnerCount int32 `json:"winner_count"`
	// *Optional*. Number of undistributed prizes
	UnclaimedPrizeCount *int32 `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage *Message `json:"giveaway_message,omitempty"`
	// *Optional*. *True*, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
	IsStarGiveaway *bool `json:"is_star_giveaway,omitempty"`
}

type _GiveawayCompleted GiveawayCompleted

// NewGiveawayCompleted instantiates a new GiveawayCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiveawayCompleted(winnerCount int32) *GiveawayCompleted {
	this := GiveawayCompleted{}
	this.WinnerCount = winnerCount
	var isStarGiveaway bool = true
	this.IsStarGiveaway = &isStarGiveaway
	return &this
}

// NewGiveawayCompletedWithDefaults instantiates a new GiveawayCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiveawayCompletedWithDefaults() *GiveawayCompleted {
	this := GiveawayCompleted{}
	var isStarGiveaway bool = true
	this.IsStarGiveaway = &isStarGiveaway
	return &this
}

// GetWinnerCount returns the WinnerCount field value
func (o *GiveawayCompleted) GetWinnerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WinnerCount
}

// GetWinnerCountOk returns a tuple with the WinnerCount field value
// and a boolean to check if the value has been set.
func (o *GiveawayCompleted) GetWinnerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnerCount, true
}

// SetWinnerCount sets field value
func (o *GiveawayCompleted) SetWinnerCount(v int32) {
	o.WinnerCount = v
}

// GetUnclaimedPrizeCount returns the UnclaimedPrizeCount field value if set, zero value otherwise.
func (o *GiveawayCompleted) GetUnclaimedPrizeCount() int32 {
	if o == nil || IsNil(o.UnclaimedPrizeCount) {
		var ret int32
		return ret
	}
	return *o.UnclaimedPrizeCount
}

// GetUnclaimedPrizeCountOk returns a tuple with the UnclaimedPrizeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayCompleted) GetUnclaimedPrizeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnclaimedPrizeCount) {
		return nil, false
	}
	return o.UnclaimedPrizeCount, true
}

// HasUnclaimedPrizeCount returns a boolean if a field has been set.
func (o *GiveawayCompleted) HasUnclaimedPrizeCount() bool {
	if o != nil && !IsNil(o.UnclaimedPrizeCount) {
		return true
	}

	return false
}

// SetUnclaimedPrizeCount gets a reference to the given int32 and assigns it to the UnclaimedPrizeCount field.
func (o *GiveawayCompleted) SetUnclaimedPrizeCount(v int32) {
	o.UnclaimedPrizeCount = &v
}


// GetGiveawayMessage returns the GiveawayMessage field value if set, zero value otherwise.
func (o *GiveawayCompleted) GetGiveawayMessage() Message {
	if o == nil || IsNil(o.GiveawayMessage) {
		var ret Message
		return ret
	}
	return *o.GiveawayMessage
}

// GetGiveawayMessageOk returns a tuple with the GiveawayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayCompleted) GetGiveawayMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.GiveawayMessage) {
		return nil, false
	}
	return o.GiveawayMessage, true
}

// HasGiveawayMessage returns a boolean if a field has been set.
func (o *GiveawayCompleted) HasGiveawayMessage() bool {
	if o != nil && !IsNil(o.GiveawayMessage) {
		return true
	}

	return false
}

// SetGiveawayMessage gets a reference to the given Message and assigns it to the GiveawayMessage field.
func (o *GiveawayCompleted) SetGiveawayMessage(v Message) {
	o.GiveawayMessage = &v
}


// GetIsStarGiveaway returns the IsStarGiveaway field value if set, zero value otherwise.
func (o *GiveawayCompleted) GetIsStarGiveaway() bool {
	if o == nil || IsNil(o.IsStarGiveaway) {
		var ret bool
		return ret
	}
	return *o.IsStarGiveaway
}

// GetIsStarGiveawayOk returns a tuple with the IsStarGiveaway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayCompleted) GetIsStarGiveawayOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStarGiveaway) {
		return nil, false
	}
	return o.IsStarGiveaway, true
}

// HasIsStarGiveaway returns a boolean if a field has been set.
func (o *GiveawayCompleted) HasIsStarGiveaway() bool {
	if o != nil && !IsNil(o.IsStarGiveaway) {
		return true
	}

	return false
}

// SetIsStarGiveaway gets a reference to the given bool and assigns it to the IsStarGiveaway field.
func (o *GiveawayCompleted) SetIsStarGiveaway(v bool) {
	o.IsStarGiveaway = &v
}


func (o GiveawayCompleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiveawayCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["winner_count"] = o.WinnerCount
	if !IsNil(o.UnclaimedPrizeCount) {
		toSerialize["unclaimed_prize_count"] = o.UnclaimedPrizeCount
	}
	if !IsNil(o.GiveawayMessage) {
		toSerialize["giveaway_message"] = o.GiveawayMessage
	}
	if !IsNil(o.IsStarGiveaway) {
		toSerialize["is_star_giveaway"] = o.IsStarGiveaway
	}
	return toSerialize, nil
}

func (o *GiveawayCompleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"winner_count",
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

	varGiveawayCompleted := _GiveawayCompleted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiveawayCompleted)

	if err != nil {
		return err
	}

	*o = GiveawayCompleted(varGiveawayCompleted)

	return err
}

type NullableGiveawayCompleted struct {
	value *GiveawayCompleted
	isSet bool
}

func (v NullableGiveawayCompleted) Get() *GiveawayCompleted {
	return v.value
}

func (v *NullableGiveawayCompleted) Set(val *GiveawayCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableGiveawayCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableGiveawayCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiveawayCompleted(val *GiveawayCompleted) *NullableGiveawayCompleted {
	return &NullableGiveawayCompleted{value: val, isSet: true}
}

func (v NullableGiveawayCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiveawayCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


