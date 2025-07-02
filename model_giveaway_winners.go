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

// checks if the GiveawayWinners type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiveawayWinners{}

// GiveawayWinners This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	Chat Chat `json:"chat"`
	// Identifier of the message with the giveaway in the chat
	GiveawayMessageId int32 `json:"giveaway_message_id"`
	// Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int32 `json:"winners_selection_date"`
	// Total number of winners in the giveaway
	WinnerCount int32 `json:"winner_count"`
	// List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`
	// *Optional*. The number of other chats the user had to join in order to be eligible for the giveaway
	AdditionalChatCount *int32 `json:"additional_chat_count,omitempty"`
	// *Optional*. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int32 `json:"prize_star_count,omitempty"`
	// *Optional*. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int32 `json:"premium_subscription_month_count,omitempty"`
	// *Optional*. Number of undistributed prizes
	UnclaimedPrizeCount *int32 `json:"unclaimed_prize_count,omitempty"`
	// *Optional*. *True*, if only users who had joined the chats after the giveaway started were eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`
	// *Optional*. *True*, if the giveaway was canceled because the payment for it was refunded
	WasRefunded *bool `json:"was_refunded,omitempty"`
	// *Optional*. Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`
}

type _GiveawayWinners GiveawayWinners

// NewGiveawayWinners instantiates a new GiveawayWinners object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiveawayWinners(chat Chat, giveawayMessageId int32, winnersSelectionDate int32, winnerCount int32, winners []User) *GiveawayWinners {
	this := GiveawayWinners{}
	this.Chat = chat
	this.GiveawayMessageId = giveawayMessageId
	this.WinnersSelectionDate = winnersSelectionDate
	this.WinnerCount = winnerCount
	this.Winners = winners
	var onlyNewMembers bool = true
	this.OnlyNewMembers = &onlyNewMembers
	var wasRefunded bool = true
	this.WasRefunded = &wasRefunded
	return &this
}

// NewGiveawayWinnersWithDefaults instantiates a new GiveawayWinners object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiveawayWinnersWithDefaults() *GiveawayWinners {
	this := GiveawayWinners{}
	var onlyNewMembers bool = true
	this.OnlyNewMembers = &onlyNewMembers
	var wasRefunded bool = true
	this.WasRefunded = &wasRefunded
	return &this
}

// GetChat returns the Chat field value
func (o *GiveawayWinners) GetChat() Chat {
	if o == nil {
		var ret Chat
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetChatOk() (*Chat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *GiveawayWinners) SetChat(v Chat) {
	o.Chat = v
}

// GetGiveawayMessageId returns the GiveawayMessageId field value
func (o *GiveawayWinners) GetGiveawayMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GiveawayMessageId
}

// GetGiveawayMessageIdOk returns a tuple with the GiveawayMessageId field value
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetGiveawayMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GiveawayMessageId, true
}

// SetGiveawayMessageId sets field value
func (o *GiveawayWinners) SetGiveawayMessageId(v int32) {
	o.GiveawayMessageId = v
}

// GetWinnersSelectionDate returns the WinnersSelectionDate field value
func (o *GiveawayWinners) GetWinnersSelectionDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WinnersSelectionDate
}

// GetWinnersSelectionDateOk returns a tuple with the WinnersSelectionDate field value
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetWinnersSelectionDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnersSelectionDate, true
}

// SetWinnersSelectionDate sets field value
func (o *GiveawayWinners) SetWinnersSelectionDate(v int32) {
	o.WinnersSelectionDate = v
}

// GetWinnerCount returns the WinnerCount field value
func (o *GiveawayWinners) GetWinnerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WinnerCount
}

// GetWinnerCountOk returns a tuple with the WinnerCount field value
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetWinnerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnerCount, true
}

// SetWinnerCount sets field value
func (o *GiveawayWinners) SetWinnerCount(v int32) {
	o.WinnerCount = v
}

// GetWinners returns the Winners field value
func (o *GiveawayWinners) GetWinners() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Winners
}

// GetWinnersOk returns a tuple with the Winners field value
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetWinnersOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Winners, true
}

// SetWinners sets field value
func (o *GiveawayWinners) SetWinners(v []User) {
	o.Winners = v
}

// GetAdditionalChatCount returns the AdditionalChatCount field value if set, zero value otherwise.
func (o *GiveawayWinners) GetAdditionalChatCount() int32 {
	if o == nil || IsNil(o.AdditionalChatCount) {
		var ret int32
		return ret
	}
	return *o.AdditionalChatCount
}

// GetAdditionalChatCountOk returns a tuple with the AdditionalChatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetAdditionalChatCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AdditionalChatCount) {
		return nil, false
	}
	return o.AdditionalChatCount, true
}

// HasAdditionalChatCount returns a boolean if a field has been set.
func (o *GiveawayWinners) HasAdditionalChatCount() bool {
	if o != nil && !IsNil(o.AdditionalChatCount) {
		return true
	}

	return false
}

// SetAdditionalChatCount gets a reference to the given int32 and assigns it to the AdditionalChatCount field.
func (o *GiveawayWinners) SetAdditionalChatCount(v int32) {
	o.AdditionalChatCount = &v
}


// GetPrizeStarCount returns the PrizeStarCount field value if set, zero value otherwise.
func (o *GiveawayWinners) GetPrizeStarCount() int32 {
	if o == nil || IsNil(o.PrizeStarCount) {
		var ret int32
		return ret
	}
	return *o.PrizeStarCount
}

// GetPrizeStarCountOk returns a tuple with the PrizeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetPrizeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrizeStarCount) {
		return nil, false
	}
	return o.PrizeStarCount, true
}

// HasPrizeStarCount returns a boolean if a field has been set.
func (o *GiveawayWinners) HasPrizeStarCount() bool {
	if o != nil && !IsNil(o.PrizeStarCount) {
		return true
	}

	return false
}

// SetPrizeStarCount gets a reference to the given int32 and assigns it to the PrizeStarCount field.
func (o *GiveawayWinners) SetPrizeStarCount(v int32) {
	o.PrizeStarCount = &v
}


// GetPremiumSubscriptionMonthCount returns the PremiumSubscriptionMonthCount field value if set, zero value otherwise.
func (o *GiveawayWinners) GetPremiumSubscriptionMonthCount() int32 {
	if o == nil || IsNil(o.PremiumSubscriptionMonthCount) {
		var ret int32
		return ret
	}
	return *o.PremiumSubscriptionMonthCount
}

// GetPremiumSubscriptionMonthCountOk returns a tuple with the PremiumSubscriptionMonthCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetPremiumSubscriptionMonthCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PremiumSubscriptionMonthCount) {
		return nil, false
	}
	return o.PremiumSubscriptionMonthCount, true
}

// HasPremiumSubscriptionMonthCount returns a boolean if a field has been set.
func (o *GiveawayWinners) HasPremiumSubscriptionMonthCount() bool {
	if o != nil && !IsNil(o.PremiumSubscriptionMonthCount) {
		return true
	}

	return false
}

// SetPremiumSubscriptionMonthCount gets a reference to the given int32 and assigns it to the PremiumSubscriptionMonthCount field.
func (o *GiveawayWinners) SetPremiumSubscriptionMonthCount(v int32) {
	o.PremiumSubscriptionMonthCount = &v
}


// GetUnclaimedPrizeCount returns the UnclaimedPrizeCount field value if set, zero value otherwise.
func (o *GiveawayWinners) GetUnclaimedPrizeCount() int32 {
	if o == nil || IsNil(o.UnclaimedPrizeCount) {
		var ret int32
		return ret
	}
	return *o.UnclaimedPrizeCount
}

// GetUnclaimedPrizeCountOk returns a tuple with the UnclaimedPrizeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetUnclaimedPrizeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnclaimedPrizeCount) {
		return nil, false
	}
	return o.UnclaimedPrizeCount, true
}

// HasUnclaimedPrizeCount returns a boolean if a field has been set.
func (o *GiveawayWinners) HasUnclaimedPrizeCount() bool {
	if o != nil && !IsNil(o.UnclaimedPrizeCount) {
		return true
	}

	return false
}

// SetUnclaimedPrizeCount gets a reference to the given int32 and assigns it to the UnclaimedPrizeCount field.
func (o *GiveawayWinners) SetUnclaimedPrizeCount(v int32) {
	o.UnclaimedPrizeCount = &v
}


// GetOnlyNewMembers returns the OnlyNewMembers field value if set, zero value otherwise.
func (o *GiveawayWinners) GetOnlyNewMembers() bool {
	if o == nil || IsNil(o.OnlyNewMembers) {
		var ret bool
		return ret
	}
	return *o.OnlyNewMembers
}

// GetOnlyNewMembersOk returns a tuple with the OnlyNewMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetOnlyNewMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyNewMembers) {
		return nil, false
	}
	return o.OnlyNewMembers, true
}

// HasOnlyNewMembers returns a boolean if a field has been set.
func (o *GiveawayWinners) HasOnlyNewMembers() bool {
	if o != nil && !IsNil(o.OnlyNewMembers) {
		return true
	}

	return false
}

// SetOnlyNewMembers gets a reference to the given bool and assigns it to the OnlyNewMembers field.
func (o *GiveawayWinners) SetOnlyNewMembers(v bool) {
	o.OnlyNewMembers = &v
}


// GetWasRefunded returns the WasRefunded field value if set, zero value otherwise.
func (o *GiveawayWinners) GetWasRefunded() bool {
	if o == nil || IsNil(o.WasRefunded) {
		var ret bool
		return ret
	}
	return *o.WasRefunded
}

// GetWasRefundedOk returns a tuple with the WasRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetWasRefundedOk() (*bool, bool) {
	if o == nil || IsNil(o.WasRefunded) {
		return nil, false
	}
	return o.WasRefunded, true
}

// HasWasRefunded returns a boolean if a field has been set.
func (o *GiveawayWinners) HasWasRefunded() bool {
	if o != nil && !IsNil(o.WasRefunded) {
		return true
	}

	return false
}

// SetWasRefunded gets a reference to the given bool and assigns it to the WasRefunded field.
func (o *GiveawayWinners) SetWasRefunded(v bool) {
	o.WasRefunded = &v
}


// GetPrizeDescription returns the PrizeDescription field value if set, zero value otherwise.
func (o *GiveawayWinners) GetPrizeDescription() string {
	if o == nil || IsNil(o.PrizeDescription) {
		var ret string
		return ret
	}
	return *o.PrizeDescription
}

// GetPrizeDescriptionOk returns a tuple with the PrizeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiveawayWinners) GetPrizeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PrizeDescription) {
		return nil, false
	}
	return o.PrizeDescription, true
}

// HasPrizeDescription returns a boolean if a field has been set.
func (o *GiveawayWinners) HasPrizeDescription() bool {
	if o != nil && !IsNil(o.PrizeDescription) {
		return true
	}

	return false
}

// SetPrizeDescription gets a reference to the given string and assigns it to the PrizeDescription field.
func (o *GiveawayWinners) SetPrizeDescription(v string) {
	o.PrizeDescription = &v
}


func (o GiveawayWinners) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiveawayWinners) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["giveaway_message_id"] = o.GiveawayMessageId
	toSerialize["winners_selection_date"] = o.WinnersSelectionDate
	toSerialize["winner_count"] = o.WinnerCount
	toSerialize["winners"] = o.Winners
	if !IsNil(o.AdditionalChatCount) {
		toSerialize["additional_chat_count"] = o.AdditionalChatCount
	}
	if !IsNil(o.PrizeStarCount) {
		toSerialize["prize_star_count"] = o.PrizeStarCount
	}
	if !IsNil(o.PremiumSubscriptionMonthCount) {
		toSerialize["premium_subscription_month_count"] = o.PremiumSubscriptionMonthCount
	}
	if !IsNil(o.UnclaimedPrizeCount) {
		toSerialize["unclaimed_prize_count"] = o.UnclaimedPrizeCount
	}
	if !IsNil(o.OnlyNewMembers) {
		toSerialize["only_new_members"] = o.OnlyNewMembers
	}
	if !IsNil(o.WasRefunded) {
		toSerialize["was_refunded"] = o.WasRefunded
	}
	if !IsNil(o.PrizeDescription) {
		toSerialize["prize_description"] = o.PrizeDescription
	}
	return toSerialize, nil
}

func (o *GiveawayWinners) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"giveaway_message_id",
		"winners_selection_date",
		"winner_count",
		"winners",
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

	varGiveawayWinners := _GiveawayWinners{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiveawayWinners)

	if err != nil {
		return err
	}

	*o = GiveawayWinners(varGiveawayWinners)

	return err
}

type NullableGiveawayWinners struct {
	value *GiveawayWinners
	isSet bool
}

func (v NullableGiveawayWinners) Get() *GiveawayWinners {
	return v.value
}

func (v *NullableGiveawayWinners) Set(val *GiveawayWinners) {
	v.value = val
	v.isSet = true
}

func (v NullableGiveawayWinners) IsSet() bool {
	return v.isSet
}

func (v *NullableGiveawayWinners) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiveawayWinners(val *GiveawayWinners) *NullableGiveawayWinners {
	return &NullableGiveawayWinners{value: val, isSet: true}
}

func (v NullableGiveawayWinners) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiveawayWinners) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


