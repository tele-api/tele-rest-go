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

// checks if the ChatBoostSourceGiveaway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatBoostSourceGiveaway{}

// ChatBoostSourceGiveaway The boost was obtained by the creation of a Telegram Premium or a Telegram Star giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription for Telegram Premium giveaways and *prize\\_star\\_count* / 500 times for one year for Telegram Star giveaways.
type ChatBoostSourceGiveaway struct {
	// Source of the boost, always “giveaway”
	Source string `json:"source"`
	// Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	GiveawayMessageId int32 `json:"giveaway_message_id"`
	User *User `json:"user,omitempty"`
	// *Optional*. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int32 `json:"prize_star_count,omitempty"`
	// *Optional*. *True*, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed *bool `json:"is_unclaimed,omitempty"`
}

type _ChatBoostSourceGiveaway ChatBoostSourceGiveaway

// NewChatBoostSourceGiveaway instantiates a new ChatBoostSourceGiveaway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatBoostSourceGiveaway(source string, giveawayMessageId int32) *ChatBoostSourceGiveaway {
	this := ChatBoostSourceGiveaway{}
	this.Source = source
	this.GiveawayMessageId = giveawayMessageId
	var isUnclaimed bool = true
	this.IsUnclaimed = &isUnclaimed
	return &this
}

// NewChatBoostSourceGiveawayWithDefaults instantiates a new ChatBoostSourceGiveaway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatBoostSourceGiveawayWithDefaults() *ChatBoostSourceGiveaway {
	this := ChatBoostSourceGiveaway{}
	var source string = "giveaway"
	this.Source = source
	var isUnclaimed bool = true
	this.IsUnclaimed = &isUnclaimed
	return &this
}

// GetSource returns the Source field value
func (o *ChatBoostSourceGiveaway) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ChatBoostSourceGiveaway) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ChatBoostSourceGiveaway) SetSource(v string) {
	o.Source = v
}

// GetGiveawayMessageId returns the GiveawayMessageId field value
func (o *ChatBoostSourceGiveaway) GetGiveawayMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GiveawayMessageId
}

// GetGiveawayMessageIdOk returns a tuple with the GiveawayMessageId field value
// and a boolean to check if the value has been set.
func (o *ChatBoostSourceGiveaway) GetGiveawayMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GiveawayMessageId, true
}

// SetGiveawayMessageId sets field value
func (o *ChatBoostSourceGiveaway) SetGiveawayMessageId(v int32) {
	o.GiveawayMessageId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ChatBoostSourceGiveaway) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBoostSourceGiveaway) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ChatBoostSourceGiveaway) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *ChatBoostSourceGiveaway) SetUser(v User) {
	o.User = &v
}


// GetPrizeStarCount returns the PrizeStarCount field value if set, zero value otherwise.
func (o *ChatBoostSourceGiveaway) GetPrizeStarCount() int32 {
	if o == nil || IsNil(o.PrizeStarCount) {
		var ret int32
		return ret
	}
	return *o.PrizeStarCount
}

// GetPrizeStarCountOk returns a tuple with the PrizeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBoostSourceGiveaway) GetPrizeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrizeStarCount) {
		return nil, false
	}
	return o.PrizeStarCount, true
}

// HasPrizeStarCount returns a boolean if a field has been set.
func (o *ChatBoostSourceGiveaway) HasPrizeStarCount() bool {
	if o != nil && !IsNil(o.PrizeStarCount) {
		return true
	}

	return false
}

// SetPrizeStarCount gets a reference to the given int32 and assigns it to the PrizeStarCount field.
func (o *ChatBoostSourceGiveaway) SetPrizeStarCount(v int32) {
	o.PrizeStarCount = &v
}


// GetIsUnclaimed returns the IsUnclaimed field value if set, zero value otherwise.
func (o *ChatBoostSourceGiveaway) GetIsUnclaimed() bool {
	if o == nil || IsNil(o.IsUnclaimed) {
		var ret bool
		return ret
	}
	return *o.IsUnclaimed
}

// GetIsUnclaimedOk returns a tuple with the IsUnclaimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBoostSourceGiveaway) GetIsUnclaimedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnclaimed) {
		return nil, false
	}
	return o.IsUnclaimed, true
}

// HasIsUnclaimed returns a boolean if a field has been set.
func (o *ChatBoostSourceGiveaway) HasIsUnclaimed() bool {
	if o != nil && !IsNil(o.IsUnclaimed) {
		return true
	}

	return false
}

// SetIsUnclaimed gets a reference to the given bool and assigns it to the IsUnclaimed field.
func (o *ChatBoostSourceGiveaway) SetIsUnclaimed(v bool) {
	o.IsUnclaimed = &v
}


func (o ChatBoostSourceGiveaway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatBoostSourceGiveaway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["giveaway_message_id"] = o.GiveawayMessageId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.PrizeStarCount) {
		toSerialize["prize_star_count"] = o.PrizeStarCount
	}
	if !IsNil(o.IsUnclaimed) {
		toSerialize["is_unclaimed"] = o.IsUnclaimed
	}
	return toSerialize, nil
}

func (o *ChatBoostSourceGiveaway) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"giveaway_message_id",
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

	varChatBoostSourceGiveaway := _ChatBoostSourceGiveaway{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatBoostSourceGiveaway)

	if err != nil {
		return err
	}

	*o = ChatBoostSourceGiveaway(varChatBoostSourceGiveaway)

	return err
}

type NullableChatBoostSourceGiveaway struct {
	value *ChatBoostSourceGiveaway
	isSet bool
}

func (v NullableChatBoostSourceGiveaway) Get() *ChatBoostSourceGiveaway {
	return v.value
}

func (v *NullableChatBoostSourceGiveaway) Set(val *ChatBoostSourceGiveaway) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBoostSourceGiveaway) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBoostSourceGiveaway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBoostSourceGiveaway(val *ChatBoostSourceGiveaway) *NullableChatBoostSourceGiveaway {
	return &NullableChatBoostSourceGiveaway{value: val, isSet: true}
}

func (v NullableChatBoostSourceGiveaway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBoostSourceGiveaway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


