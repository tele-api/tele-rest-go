/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the GiftPremiumSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftPremiumSubscriptionRequest{}

// GiftPremiumSubscriptionRequest Request parameters for giftPremiumSubscription
type GiftPremiumSubscriptionRequest struct {
	// Unique identifier of the target user who will receive a Telegram Premium subscription
	UserId int32 `json:"user_id"`
	// Number of months the Telegram Premium subscription will be active for the user; must be one of 3, 6, or 12
	MonthCount int32 `json:"month_count"`
	// Number of Telegram Stars to pay for the Telegram Premium subscription; must be 1000 for 3 months, 1500 for 6 months, and 2500 for 12 months
	StarCount int32 `json:"star_count"`
	// Text that will be shown along with the service message about the subscription; 0-128 characters
	Text *string `json:"text,omitempty"`
	// Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of *text\\_parse\\_mode*. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type _GiftPremiumSubscriptionRequest GiftPremiumSubscriptionRequest

// NewGiftPremiumSubscriptionRequest instantiates a new GiftPremiumSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftPremiumSubscriptionRequest(userId int32, monthCount int32, starCount int32) *GiftPremiumSubscriptionRequest {
	this := GiftPremiumSubscriptionRequest{}
	this.UserId = userId
	this.MonthCount = monthCount
	this.StarCount = starCount
	return &this
}

// NewGiftPremiumSubscriptionRequestWithDefaults instantiates a new GiftPremiumSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftPremiumSubscriptionRequestWithDefaults() *GiftPremiumSubscriptionRequest {
	this := GiftPremiumSubscriptionRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *GiftPremiumSubscriptionRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GiftPremiumSubscriptionRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetMonthCount returns the MonthCount field value
func (o *GiftPremiumSubscriptionRequest) GetMonthCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MonthCount
}

// GetMonthCountOk returns a tuple with the MonthCount field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetMonthCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthCount, true
}

// SetMonthCount sets field value
func (o *GiftPremiumSubscriptionRequest) SetMonthCount(v int32) {
	o.MonthCount = v
}

// GetStarCount returns the StarCount field value
func (o *GiftPremiumSubscriptionRequest) GetStarCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StarCount
}

// GetStarCountOk returns a tuple with the StarCount field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetStarCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarCount, true
}

// SetStarCount sets field value
func (o *GiftPremiumSubscriptionRequest) SetStarCount(v int32) {
	o.StarCount = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *GiftPremiumSubscriptionRequest) SetText(v string) {
	o.Text = &v
}


// GetTextParseMode returns the TextParseMode field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionRequest) GetTextParseMode() string {
	if o == nil || IsNil(o.TextParseMode) {
		var ret string
		return ret
	}
	return *o.TextParseMode
}

// GetTextParseModeOk returns a tuple with the TextParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetTextParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.TextParseMode) {
		return nil, false
	}
	return o.TextParseMode, true
}

// HasTextParseMode returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionRequest) HasTextParseMode() bool {
	if o != nil && !IsNil(o.TextParseMode) {
		return true
	}

	return false
}

// SetTextParseMode gets a reference to the given string and assigns it to the TextParseMode field.
func (o *GiftPremiumSubscriptionRequest) SetTextParseMode(v string) {
	o.TextParseMode = &v
}


// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionRequest) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionRequest) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionRequest) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *GiftPremiumSubscriptionRequest) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


func (o GiftPremiumSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftPremiumSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["month_count"] = o.MonthCount
	toSerialize["star_count"] = o.StarCount
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.TextParseMode) {
		toSerialize["text_parse_mode"] = o.TextParseMode
	}
	if !IsNil(o.TextEntities) {
		toSerialize["text_entities"] = o.TextEntities
	}
	return toSerialize, nil
}

func (o *GiftPremiumSubscriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"month_count",
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

	varGiftPremiumSubscriptionRequest := _GiftPremiumSubscriptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiftPremiumSubscriptionRequest)

	if err != nil {
		return err
	}

	*o = GiftPremiumSubscriptionRequest(varGiftPremiumSubscriptionRequest)

	return err
}

type NullableGiftPremiumSubscriptionRequest struct {
	value *GiftPremiumSubscriptionRequest
	isSet bool
}

func (v NullableGiftPremiumSubscriptionRequest) Get() *GiftPremiumSubscriptionRequest {
	return v.value
}

func (v *NullableGiftPremiumSubscriptionRequest) Set(val *GiftPremiumSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftPremiumSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftPremiumSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftPremiumSubscriptionRequest(val *GiftPremiumSubscriptionRequest) *NullableGiftPremiumSubscriptionRequest {
	return &NullableGiftPremiumSubscriptionRequest{value: val, isSet: true}
}

func (v NullableGiftPremiumSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftPremiumSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


