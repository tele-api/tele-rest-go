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

// checks if the Giveaway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Giveaway{}

// Giveaway This object represents a message about a scheduled giveaway.
type Giveaway struct {
	// The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`
	// Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int32 `json:"winners_selection_date"`
	// The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int32 `json:"winner_count"`
	// *Optional*. *True*, if only users who join the chats after the giveaway started should be eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`
	// *Optional*. *True*, if the list of giveaway winners will be visible to everyone
	HasPublicWinners *bool `json:"has_public_winners,omitempty"`
	// *Optional*. Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`
	// *Optional*. A list of two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	CountryCodes []string `json:"country_codes,omitempty"`
	// *Optional*. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int32 `json:"prize_star_count,omitempty"`
	// *Optional*. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int32 `json:"premium_subscription_month_count,omitempty"`
}

type _Giveaway Giveaway

// NewGiveaway instantiates a new Giveaway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiveaway(chats []Chat, winnersSelectionDate int32, winnerCount int32) *Giveaway {
	this := Giveaway{}
	this.Chats = chats
	this.WinnersSelectionDate = winnersSelectionDate
	this.WinnerCount = winnerCount
	var onlyNewMembers bool = true
	this.OnlyNewMembers = &onlyNewMembers
	var hasPublicWinners bool = true
	this.HasPublicWinners = &hasPublicWinners
	return &this
}

// NewGiveawayWithDefaults instantiates a new Giveaway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiveawayWithDefaults() *Giveaway {
	this := Giveaway{}
	var onlyNewMembers bool = true
	this.OnlyNewMembers = &onlyNewMembers
	var hasPublicWinners bool = true
	this.HasPublicWinners = &hasPublicWinners
	return &this
}

// GetChats returns the Chats field value
func (o *Giveaway) GetChats() []Chat {
	if o == nil {
		var ret []Chat
		return ret
	}

	return o.Chats
}

// GetChatsOk returns a tuple with the Chats field value
// and a boolean to check if the value has been set.
func (o *Giveaway) GetChatsOk() ([]Chat, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chats, true
}

// SetChats sets field value
func (o *Giveaway) SetChats(v []Chat) {
	o.Chats = v
}

// GetWinnersSelectionDate returns the WinnersSelectionDate field value
func (o *Giveaway) GetWinnersSelectionDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WinnersSelectionDate
}

// GetWinnersSelectionDateOk returns a tuple with the WinnersSelectionDate field value
// and a boolean to check if the value has been set.
func (o *Giveaway) GetWinnersSelectionDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnersSelectionDate, true
}

// SetWinnersSelectionDate sets field value
func (o *Giveaway) SetWinnersSelectionDate(v int32) {
	o.WinnersSelectionDate = v
}

// GetWinnerCount returns the WinnerCount field value
func (o *Giveaway) GetWinnerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WinnerCount
}

// GetWinnerCountOk returns a tuple with the WinnerCount field value
// and a boolean to check if the value has been set.
func (o *Giveaway) GetWinnerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnerCount, true
}

// SetWinnerCount sets field value
func (o *Giveaway) SetWinnerCount(v int32) {
	o.WinnerCount = v
}

// GetOnlyNewMembers returns the OnlyNewMembers field value if set, zero value otherwise.
func (o *Giveaway) GetOnlyNewMembers() bool {
	if o == nil || IsNil(o.OnlyNewMembers) {
		var ret bool
		return ret
	}
	return *o.OnlyNewMembers
}

// GetOnlyNewMembersOk returns a tuple with the OnlyNewMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetOnlyNewMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyNewMembers) {
		return nil, false
	}
	return o.OnlyNewMembers, true
}

// HasOnlyNewMembers returns a boolean if a field has been set.
func (o *Giveaway) HasOnlyNewMembers() bool {
	if o != nil && !IsNil(o.OnlyNewMembers) {
		return true
	}

	return false
}

// SetOnlyNewMembers gets a reference to the given bool and assigns it to the OnlyNewMembers field.
func (o *Giveaway) SetOnlyNewMembers(v bool) {
	o.OnlyNewMembers = &v
}


// GetHasPublicWinners returns the HasPublicWinners field value if set, zero value otherwise.
func (o *Giveaway) GetHasPublicWinners() bool {
	if o == nil || IsNil(o.HasPublicWinners) {
		var ret bool
		return ret
	}
	return *o.HasPublicWinners
}

// GetHasPublicWinnersOk returns a tuple with the HasPublicWinners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetHasPublicWinnersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPublicWinners) {
		return nil, false
	}
	return o.HasPublicWinners, true
}

// HasHasPublicWinners returns a boolean if a field has been set.
func (o *Giveaway) HasHasPublicWinners() bool {
	if o != nil && !IsNil(o.HasPublicWinners) {
		return true
	}

	return false
}

// SetHasPublicWinners gets a reference to the given bool and assigns it to the HasPublicWinners field.
func (o *Giveaway) SetHasPublicWinners(v bool) {
	o.HasPublicWinners = &v
}


// GetPrizeDescription returns the PrizeDescription field value if set, zero value otherwise.
func (o *Giveaway) GetPrizeDescription() string {
	if o == nil || IsNil(o.PrizeDescription) {
		var ret string
		return ret
	}
	return *o.PrizeDescription
}

// GetPrizeDescriptionOk returns a tuple with the PrizeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetPrizeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PrizeDescription) {
		return nil, false
	}
	return o.PrizeDescription, true
}

// HasPrizeDescription returns a boolean if a field has been set.
func (o *Giveaway) HasPrizeDescription() bool {
	if o != nil && !IsNil(o.PrizeDescription) {
		return true
	}

	return false
}

// SetPrizeDescription gets a reference to the given string and assigns it to the PrizeDescription field.
func (o *Giveaway) SetPrizeDescription(v string) {
	o.PrizeDescription = &v
}


// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *Giveaway) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *Giveaway) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *Giveaway) SetCountryCodes(v []string) {
	o.CountryCodes = v
}


// GetPrizeStarCount returns the PrizeStarCount field value if set, zero value otherwise.
func (o *Giveaway) GetPrizeStarCount() int32 {
	if o == nil || IsNil(o.PrizeStarCount) {
		var ret int32
		return ret
	}
	return *o.PrizeStarCount
}

// GetPrizeStarCountOk returns a tuple with the PrizeStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetPrizeStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrizeStarCount) {
		return nil, false
	}
	return o.PrizeStarCount, true
}

// HasPrizeStarCount returns a boolean if a field has been set.
func (o *Giveaway) HasPrizeStarCount() bool {
	if o != nil && !IsNil(o.PrizeStarCount) {
		return true
	}

	return false
}

// SetPrizeStarCount gets a reference to the given int32 and assigns it to the PrizeStarCount field.
func (o *Giveaway) SetPrizeStarCount(v int32) {
	o.PrizeStarCount = &v
}


// GetPremiumSubscriptionMonthCount returns the PremiumSubscriptionMonthCount field value if set, zero value otherwise.
func (o *Giveaway) GetPremiumSubscriptionMonthCount() int32 {
	if o == nil || IsNil(o.PremiumSubscriptionMonthCount) {
		var ret int32
		return ret
	}
	return *o.PremiumSubscriptionMonthCount
}

// GetPremiumSubscriptionMonthCountOk returns a tuple with the PremiumSubscriptionMonthCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetPremiumSubscriptionMonthCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PremiumSubscriptionMonthCount) {
		return nil, false
	}
	return o.PremiumSubscriptionMonthCount, true
}

// HasPremiumSubscriptionMonthCount returns a boolean if a field has been set.
func (o *Giveaway) HasPremiumSubscriptionMonthCount() bool {
	if o != nil && !IsNil(o.PremiumSubscriptionMonthCount) {
		return true
	}

	return false
}

// SetPremiumSubscriptionMonthCount gets a reference to the given int32 and assigns it to the PremiumSubscriptionMonthCount field.
func (o *Giveaway) SetPremiumSubscriptionMonthCount(v int32) {
	o.PremiumSubscriptionMonthCount = &v
}


func (o Giveaway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Giveaway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chats"] = o.Chats
	toSerialize["winners_selection_date"] = o.WinnersSelectionDate
	toSerialize["winner_count"] = o.WinnerCount
	if !IsNil(o.OnlyNewMembers) {
		toSerialize["only_new_members"] = o.OnlyNewMembers
	}
	if !IsNil(o.HasPublicWinners) {
		toSerialize["has_public_winners"] = o.HasPublicWinners
	}
	if !IsNil(o.PrizeDescription) {
		toSerialize["prize_description"] = o.PrizeDescription
	}
	if !IsNil(o.CountryCodes) {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if !IsNil(o.PrizeStarCount) {
		toSerialize["prize_star_count"] = o.PrizeStarCount
	}
	if !IsNil(o.PremiumSubscriptionMonthCount) {
		toSerialize["premium_subscription_month_count"] = o.PremiumSubscriptionMonthCount
	}
	return toSerialize, nil
}

func (o *Giveaway) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chats",
		"winners_selection_date",
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

	varGiveaway := _Giveaway{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiveaway)

	if err != nil {
		return err
	}

	*o = Giveaway(varGiveaway)

	return err
}

type NullableGiveaway struct {
	value *Giveaway
	isSet bool
}

func (v NullableGiveaway) Get() *Giveaway {
	return v.value
}

func (v *NullableGiveaway) Set(val *Giveaway) {
	v.value = val
	v.isSet = true
}

func (v NullableGiveaway) IsSet() bool {
	return v.isSet
}

func (v *NullableGiveaway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiveaway(val *Giveaway) *NullableGiveaway {
	return &NullableGiveaway{value: val, isSet: true}
}

func (v NullableGiveaway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiveaway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


