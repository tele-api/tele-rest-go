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

// checks if the PostSendGiftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSendGiftRequest{}

// PostSendGiftRequest struct for PostSendGiftRequest
type PostSendGiftRequest struct {
	// Required if *chat\\_id* is not specified. Unique identifier of the target user who will receive the gift.
	UserId *int32 `json:"user_id,omitempty"`
	ChatId *PostSendGiftRequestChatId `json:"chat_id,omitempty"`
	// Identifier of the gift
	GiftId string `json:"gift_id"`
	// Pass *True* to pay for the gift upgrade from the bot's balance, thereby making the upgrade free for the receiver
	PayForUpgrade *bool `json:"pay_for_upgrade,omitempty"`
	// Text that will be shown along with the gift; 0-128 characters
	Text *string `json:"text,omitempty"`
	// Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of *text\\_parse\\_mode*. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type _PostSendGiftRequest PostSendGiftRequest

// NewPostSendGiftRequest instantiates a new PostSendGiftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSendGiftRequest(giftId string) *PostSendGiftRequest {
	this := PostSendGiftRequest{}
	this.GiftId = giftId
	return &this
}

// NewPostSendGiftRequestWithDefaults instantiates a new PostSendGiftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSendGiftRequestWithDefaults() *PostSendGiftRequest {
	this := PostSendGiftRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *PostSendGiftRequest) SetUserId(v int32) {
	o.UserId = &v
}


// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetChatId() PostSendGiftRequestChatId {
	if o == nil || IsNil(o.ChatId) {
		var ret PostSendGiftRequestChatId
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetChatIdOk() (*PostSendGiftRequestChatId, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given PostSendGiftRequestChatId and assigns it to the ChatId field.
func (o *PostSendGiftRequest) SetChatId(v PostSendGiftRequestChatId) {
	o.ChatId = &v
}


// GetGiftId returns the GiftId field value
func (o *PostSendGiftRequest) GetGiftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GiftId
}

// GetGiftIdOk returns a tuple with the GiftId field value
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetGiftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GiftId, true
}

// SetGiftId sets field value
func (o *PostSendGiftRequest) SetGiftId(v string) {
	o.GiftId = v
}

// GetPayForUpgrade returns the PayForUpgrade field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetPayForUpgrade() bool {
	if o == nil || IsNil(o.PayForUpgrade) {
		var ret bool
		return ret
	}
	return *o.PayForUpgrade
}

// GetPayForUpgradeOk returns a tuple with the PayForUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetPayForUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.PayForUpgrade) {
		return nil, false
	}
	return o.PayForUpgrade, true
}

// HasPayForUpgrade returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasPayForUpgrade() bool {
	if o != nil && !IsNil(o.PayForUpgrade) {
		return true
	}

	return false
}

// SetPayForUpgrade gets a reference to the given bool and assigns it to the PayForUpgrade field.
func (o *PostSendGiftRequest) SetPayForUpgrade(v bool) {
	o.PayForUpgrade = &v
}


// GetText returns the Text field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PostSendGiftRequest) SetText(v string) {
	o.Text = &v
}


// GetTextParseMode returns the TextParseMode field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetTextParseMode() string {
	if o == nil || IsNil(o.TextParseMode) {
		var ret string
		return ret
	}
	return *o.TextParseMode
}

// GetTextParseModeOk returns a tuple with the TextParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetTextParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.TextParseMode) {
		return nil, false
	}
	return o.TextParseMode, true
}

// HasTextParseMode returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasTextParseMode() bool {
	if o != nil && !IsNil(o.TextParseMode) {
		return true
	}

	return false
}

// SetTextParseMode gets a reference to the given string and assigns it to the TextParseMode field.
func (o *PostSendGiftRequest) SetTextParseMode(v string) {
	o.TextParseMode = &v
}


// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *PostSendGiftRequest) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSendGiftRequest) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *PostSendGiftRequest) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *PostSendGiftRequest) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


func (o PostSendGiftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSendGiftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	toSerialize["gift_id"] = o.GiftId
	if !IsNil(o.PayForUpgrade) {
		toSerialize["pay_for_upgrade"] = o.PayForUpgrade
	}
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

func (o *PostSendGiftRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gift_id",
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

	varPostSendGiftRequest := _PostSendGiftRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSendGiftRequest)

	if err != nil {
		return err
	}

	*o = PostSendGiftRequest(varPostSendGiftRequest)

	return err
}

type NullablePostSendGiftRequest struct {
	value *PostSendGiftRequest
	isSet bool
}

func (v NullablePostSendGiftRequest) Get() *PostSendGiftRequest {
	return v.value
}

func (v *NullablePostSendGiftRequest) Set(val *PostSendGiftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSendGiftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSendGiftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSendGiftRequest(val *PostSendGiftRequest) *NullablePostSendGiftRequest {
	return &NullablePostSendGiftRequest{value: val, isSet: true}
}

func (v NullablePostSendGiftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSendGiftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


