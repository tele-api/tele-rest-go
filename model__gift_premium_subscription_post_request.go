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

// checks if the GiftPremiumSubscriptionPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftPremiumSubscriptionPostRequest{}

// GiftPremiumSubscriptionPostRequest struct for GiftPremiumSubscriptionPostRequest
type GiftPremiumSubscriptionPostRequest struct {
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

type _GiftPremiumSubscriptionPostRequest GiftPremiumSubscriptionPostRequest

// NewGiftPremiumSubscriptionPostRequest instantiates a new GiftPremiumSubscriptionPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftPremiumSubscriptionPostRequest(userId int32, monthCount int32, starCount int32) *GiftPremiumSubscriptionPostRequest {
	this := GiftPremiumSubscriptionPostRequest{}
	this.UserId = userId
	this.MonthCount = monthCount
	this.StarCount = starCount
	return &this
}

// NewGiftPremiumSubscriptionPostRequestWithDefaults instantiates a new GiftPremiumSubscriptionPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftPremiumSubscriptionPostRequestWithDefaults() *GiftPremiumSubscriptionPostRequest {
	this := GiftPremiumSubscriptionPostRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *GiftPremiumSubscriptionPostRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GiftPremiumSubscriptionPostRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetMonthCount returns the MonthCount field value
func (o *GiftPremiumSubscriptionPostRequest) GetMonthCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MonthCount
}

// GetMonthCountOk returns a tuple with the MonthCount field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetMonthCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthCount, true
}

// SetMonthCount sets field value
func (o *GiftPremiumSubscriptionPostRequest) SetMonthCount(v int32) {
	o.MonthCount = v
}

// GetStarCount returns the StarCount field value
func (o *GiftPremiumSubscriptionPostRequest) GetStarCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StarCount
}

// GetStarCountOk returns a tuple with the StarCount field value
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetStarCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarCount, true
}

// SetStarCount sets field value
func (o *GiftPremiumSubscriptionPostRequest) SetStarCount(v int32) {
	o.StarCount = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionPostRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionPostRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *GiftPremiumSubscriptionPostRequest) SetText(v string) {
	o.Text = &v
}


// GetTextParseMode returns the TextParseMode field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionPostRequest) GetTextParseMode() string {
	if o == nil || IsNil(o.TextParseMode) {
		var ret string
		return ret
	}
	return *o.TextParseMode
}

// GetTextParseModeOk returns a tuple with the TextParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetTextParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.TextParseMode) {
		return nil, false
	}
	return o.TextParseMode, true
}

// HasTextParseMode returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionPostRequest) HasTextParseMode() bool {
	if o != nil && !IsNil(o.TextParseMode) {
		return true
	}

	return false
}

// SetTextParseMode gets a reference to the given string and assigns it to the TextParseMode field.
func (o *GiftPremiumSubscriptionPostRequest) SetTextParseMode(v string) {
	o.TextParseMode = &v
}


// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *GiftPremiumSubscriptionPostRequest) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftPremiumSubscriptionPostRequest) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *GiftPremiumSubscriptionPostRequest) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *GiftPremiumSubscriptionPostRequest) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


func (o GiftPremiumSubscriptionPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftPremiumSubscriptionPostRequest) ToMap() (map[string]interface{}, error) {
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

func (o *GiftPremiumSubscriptionPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGiftPremiumSubscriptionPostRequest := _GiftPremiumSubscriptionPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGiftPremiumSubscriptionPostRequest)

	if err != nil {
		return err
	}

	*o = GiftPremiumSubscriptionPostRequest(varGiftPremiumSubscriptionPostRequest)

	return err
}

type NullableGiftPremiumSubscriptionPostRequest struct {
	value *GiftPremiumSubscriptionPostRequest
	isSet bool
}

func (v NullableGiftPremiumSubscriptionPostRequest) Get() *GiftPremiumSubscriptionPostRequest {
	return v.value
}

func (v *NullableGiftPremiumSubscriptionPostRequest) Set(val *GiftPremiumSubscriptionPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftPremiumSubscriptionPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftPremiumSubscriptionPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftPremiumSubscriptionPostRequest(val *GiftPremiumSubscriptionPostRequest) *NullableGiftPremiumSubscriptionPostRequest {
	return &NullableGiftPremiumSubscriptionPostRequest{value: val, isSet: true}
}

func (v NullableGiftPremiumSubscriptionPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftPremiumSubscriptionPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


