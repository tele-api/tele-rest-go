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

// checks if the CreateChatSubscriptionInviteLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatSubscriptionInviteLinkRequest{}

// CreateChatSubscriptionInviteLinkRequest Request parameters for createChatSubscriptionInviteLink
type CreateChatSubscriptionInviteLinkRequest struct {
	ChatId CreateChatSubscriptionInviteLinkRequestChatId `json:"chat_id"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	// The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPeriod int32 `json:"subscription_period"`
	// The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-10000
	SubscriptionPrice int32 `json:"subscription_price"`
}

type _CreateChatSubscriptionInviteLinkRequest CreateChatSubscriptionInviteLinkRequest

// NewCreateChatSubscriptionInviteLinkRequest instantiates a new CreateChatSubscriptionInviteLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatSubscriptionInviteLinkRequest(chatId CreateChatSubscriptionInviteLinkRequestChatId, subscriptionPeriod int32, subscriptionPrice int32) *CreateChatSubscriptionInviteLinkRequest {
	this := CreateChatSubscriptionInviteLinkRequest{}
	this.ChatId = chatId
	this.SubscriptionPeriod = subscriptionPeriod
	this.SubscriptionPrice = subscriptionPrice
	return &this
}

// NewCreateChatSubscriptionInviteLinkRequestWithDefaults instantiates a new CreateChatSubscriptionInviteLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatSubscriptionInviteLinkRequestWithDefaults() *CreateChatSubscriptionInviteLinkRequest {
	this := CreateChatSubscriptionInviteLinkRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *CreateChatSubscriptionInviteLinkRequest) GetChatId() CreateChatSubscriptionInviteLinkRequestChatId {
	if o == nil {
		var ret CreateChatSubscriptionInviteLinkRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *CreateChatSubscriptionInviteLinkRequest) GetChatIdOk() (*CreateChatSubscriptionInviteLinkRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *CreateChatSubscriptionInviteLinkRequest) SetChatId(v CreateChatSubscriptionInviteLinkRequestChatId) {
	o.ChatId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateChatSubscriptionInviteLinkRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatSubscriptionInviteLinkRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateChatSubscriptionInviteLinkRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateChatSubscriptionInviteLinkRequest) SetName(v string) {
	o.Name = &v
}


// GetSubscriptionPeriod returns the SubscriptionPeriod field value
func (o *CreateChatSubscriptionInviteLinkRequest) GetSubscriptionPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value
// and a boolean to check if the value has been set.
func (o *CreateChatSubscriptionInviteLinkRequest) GetSubscriptionPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionPeriod, true
}

// SetSubscriptionPeriod sets field value
func (o *CreateChatSubscriptionInviteLinkRequest) SetSubscriptionPeriod(v int32) {
	o.SubscriptionPeriod = v
}

// GetSubscriptionPrice returns the SubscriptionPrice field value
func (o *CreateChatSubscriptionInviteLinkRequest) GetSubscriptionPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubscriptionPrice
}

// GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field value
// and a boolean to check if the value has been set.
func (o *CreateChatSubscriptionInviteLinkRequest) GetSubscriptionPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionPrice, true
}

// SetSubscriptionPrice sets field value
func (o *CreateChatSubscriptionInviteLinkRequest) SetSubscriptionPrice(v int32) {
	o.SubscriptionPrice = v
}

func (o CreateChatSubscriptionInviteLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatSubscriptionInviteLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["subscription_period"] = o.SubscriptionPeriod
	toSerialize["subscription_price"] = o.SubscriptionPrice
	return toSerialize, nil
}

func (o *CreateChatSubscriptionInviteLinkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"subscription_period",
		"subscription_price",
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

	varCreateChatSubscriptionInviteLinkRequest := _CreateChatSubscriptionInviteLinkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatSubscriptionInviteLinkRequest)

	if err != nil {
		return err
	}

	*o = CreateChatSubscriptionInviteLinkRequest(varCreateChatSubscriptionInviteLinkRequest)

	return err
}

type NullableCreateChatSubscriptionInviteLinkRequest struct {
	value *CreateChatSubscriptionInviteLinkRequest
	isSet bool
}

func (v NullableCreateChatSubscriptionInviteLinkRequest) Get() *CreateChatSubscriptionInviteLinkRequest {
	return v.value
}

func (v *NullableCreateChatSubscriptionInviteLinkRequest) Set(val *CreateChatSubscriptionInviteLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatSubscriptionInviteLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatSubscriptionInviteLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatSubscriptionInviteLinkRequest(val *CreateChatSubscriptionInviteLinkRequest) *NullableCreateChatSubscriptionInviteLinkRequest {
	return &NullableCreateChatSubscriptionInviteLinkRequest{value: val, isSet: true}
}

func (v NullableCreateChatSubscriptionInviteLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatSubscriptionInviteLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


