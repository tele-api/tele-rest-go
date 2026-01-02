/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.3.0
 *    * - **Modified**: 2026-01-01T02:06:16.405655294Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
	// *Optional*. *True*, if the gift can only be purchased by Telegram Premium subscribers
	IsPremium *bool `json:"is_premium,omitempty"`
	// *Optional*. *True*, if the gift can be used (after being upgraded) to customize a user's appearance
	HasColors *bool `json:"has_colors,omitempty"`
	// *Optional*. The total number of gifts of this type that can be sent by all users; for limited gifts only
	TotalCount *int32 `json:"total_count,omitempty"`
	// *Optional*. The number of remaining gifts of this type that can be sent by all users; for limited gifts only
	RemainingCount *int32 `json:"remaining_count,omitempty"`
	// *Optional*. The total number of gifts of this type that can be sent by the bot; for limited gifts only
	PersonalTotalCount *int32 `json:"personal_total_count,omitempty"`
	// *Optional*. The number of remaining gifts of this type that can be sent by the bot; for limited gifts only
	PersonalRemainingCount *int32 `json:"personal_remaining_count,omitempty"`
	Background *GiftBackground `json:"background,omitempty"`
	// *Optional*. The total number of different unique gifts that can be obtained by upgrading the gift
	UniqueGiftVariantCount *int32 `json:"unique_gift_variant_count,omitempty"`
	PublisherChat *Chat `json:"publisher_chat,omitempty"`
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
	var isPremium bool = true
	this.IsPremium = &isPremium
	var hasColors bool = true
	this.HasColors = &hasColors
	return &this
}

// NewGiftWithDefaults instantiates a new Gift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftWithDefaults() *Gift {
	this := Gift{}
	var isPremium bool = true
	this.IsPremium = &isPremium
	var hasColors bool = true
	this.HasColors = &hasColors
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


// GetIsPremium returns the IsPremium field value if set, zero value otherwise.
func (o *Gift) GetIsPremium() bool {
	if o == nil || IsNil(o.IsPremium) {
		var ret bool
		return ret
	}
	return *o.IsPremium
}

// GetIsPremiumOk returns a tuple with the IsPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetIsPremiumOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPremium) {
		return nil, false
	}
	return o.IsPremium, true
}

// HasIsPremium returns a boolean if a field has been set.
func (o *Gift) HasIsPremium() bool {
	if o != nil && !IsNil(o.IsPremium) {
		return true
	}

	return false
}

// SetIsPremium gets a reference to the given bool and assigns it to the IsPremium field.
func (o *Gift) SetIsPremium(v bool) {
	o.IsPremium = &v
}


// GetHasColors returns the HasColors field value if set, zero value otherwise.
func (o *Gift) GetHasColors() bool {
	if o == nil || IsNil(o.HasColors) {
		var ret bool
		return ret
	}
	return *o.HasColors
}

// GetHasColorsOk returns a tuple with the HasColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetHasColorsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasColors) {
		return nil, false
	}
	return o.HasColors, true
}

// HasHasColors returns a boolean if a field has been set.
func (o *Gift) HasHasColors() bool {
	if o != nil && !IsNil(o.HasColors) {
		return true
	}

	return false
}

// SetHasColors gets a reference to the given bool and assigns it to the HasColors field.
func (o *Gift) SetHasColors(v bool) {
	o.HasColors = &v
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


// GetPersonalTotalCount returns the PersonalTotalCount field value if set, zero value otherwise.
func (o *Gift) GetPersonalTotalCount() int32 {
	if o == nil || IsNil(o.PersonalTotalCount) {
		var ret int32
		return ret
	}
	return *o.PersonalTotalCount
}

// GetPersonalTotalCountOk returns a tuple with the PersonalTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetPersonalTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonalTotalCount) {
		return nil, false
	}
	return o.PersonalTotalCount, true
}

// HasPersonalTotalCount returns a boolean if a field has been set.
func (o *Gift) HasPersonalTotalCount() bool {
	if o != nil && !IsNil(o.PersonalTotalCount) {
		return true
	}

	return false
}

// SetPersonalTotalCount gets a reference to the given int32 and assigns it to the PersonalTotalCount field.
func (o *Gift) SetPersonalTotalCount(v int32) {
	o.PersonalTotalCount = &v
}


// GetPersonalRemainingCount returns the PersonalRemainingCount field value if set, zero value otherwise.
func (o *Gift) GetPersonalRemainingCount() int32 {
	if o == nil || IsNil(o.PersonalRemainingCount) {
		var ret int32
		return ret
	}
	return *o.PersonalRemainingCount
}

// GetPersonalRemainingCountOk returns a tuple with the PersonalRemainingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetPersonalRemainingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonalRemainingCount) {
		return nil, false
	}
	return o.PersonalRemainingCount, true
}

// HasPersonalRemainingCount returns a boolean if a field has been set.
func (o *Gift) HasPersonalRemainingCount() bool {
	if o != nil && !IsNil(o.PersonalRemainingCount) {
		return true
	}

	return false
}

// SetPersonalRemainingCount gets a reference to the given int32 and assigns it to the PersonalRemainingCount field.
func (o *Gift) SetPersonalRemainingCount(v int32) {
	o.PersonalRemainingCount = &v
}


// GetBackground returns the Background field value if set, zero value otherwise.
func (o *Gift) GetBackground() GiftBackground {
	if o == nil || IsNil(o.Background) {
		var ret GiftBackground
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetBackgroundOk() (*GiftBackground, bool) {
	if o == nil || IsNil(o.Background) {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *Gift) HasBackground() bool {
	if o != nil && !IsNil(o.Background) {
		return true
	}

	return false
}

// SetBackground gets a reference to the given GiftBackground and assigns it to the Background field.
func (o *Gift) SetBackground(v GiftBackground) {
	o.Background = &v
}


// GetUniqueGiftVariantCount returns the UniqueGiftVariantCount field value if set, zero value otherwise.
func (o *Gift) GetUniqueGiftVariantCount() int32 {
	if o == nil || IsNil(o.UniqueGiftVariantCount) {
		var ret int32
		return ret
	}
	return *o.UniqueGiftVariantCount
}

// GetUniqueGiftVariantCountOk returns a tuple with the UniqueGiftVariantCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetUniqueGiftVariantCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueGiftVariantCount) {
		return nil, false
	}
	return o.UniqueGiftVariantCount, true
}

// HasUniqueGiftVariantCount returns a boolean if a field has been set.
func (o *Gift) HasUniqueGiftVariantCount() bool {
	if o != nil && !IsNil(o.UniqueGiftVariantCount) {
		return true
	}

	return false
}

// SetUniqueGiftVariantCount gets a reference to the given int32 and assigns it to the UniqueGiftVariantCount field.
func (o *Gift) SetUniqueGiftVariantCount(v int32) {
	o.UniqueGiftVariantCount = &v
}


// GetPublisherChat returns the PublisherChat field value if set, zero value otherwise.
func (o *Gift) GetPublisherChat() Chat {
	if o == nil || IsNil(o.PublisherChat) {
		var ret Chat
		return ret
	}
	return *o.PublisherChat
}

// GetPublisherChatOk returns a tuple with the PublisherChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gift) GetPublisherChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.PublisherChat) {
		return nil, false
	}
	return o.PublisherChat, true
}

// HasPublisherChat returns a boolean if a field has been set.
func (o *Gift) HasPublisherChat() bool {
	if o != nil && !IsNil(o.PublisherChat) {
		return true
	}

	return false
}

// SetPublisherChat gets a reference to the given Chat and assigns it to the PublisherChat field.
func (o *Gift) SetPublisherChat(v Chat) {
	o.PublisherChat = &v
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
	if !IsNil(o.IsPremium) {
		toSerialize["is_premium"] = o.IsPremium
	}
	if !IsNil(o.HasColors) {
		toSerialize["has_colors"] = o.HasColors
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !IsNil(o.RemainingCount) {
		toSerialize["remaining_count"] = o.RemainingCount
	}
	if !IsNil(o.PersonalTotalCount) {
		toSerialize["personal_total_count"] = o.PersonalTotalCount
	}
	if !IsNil(o.PersonalRemainingCount) {
		toSerialize["personal_remaining_count"] = o.PersonalRemainingCount
	}
	if !IsNil(o.Background) {
		toSerialize["background"] = o.Background
	}
	if !IsNil(o.UniqueGiftVariantCount) {
		toSerialize["unique_gift_variant_count"] = o.UniqueGiftVariantCount
	}
	if !IsNil(o.PublisherChat) {
		toSerialize["publisher_chat"] = o.PublisherChat
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

