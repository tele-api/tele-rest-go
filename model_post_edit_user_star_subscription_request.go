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

// checks if the PostEditUserStarSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostEditUserStarSubscriptionRequest{}

// PostEditUserStarSubscriptionRequest struct for PostEditUserStarSubscriptionRequest
type PostEditUserStarSubscriptionRequest struct {
	// Identifier of the user whose subscription will be edited
	UserId int32 `json:"user_id"`
	// Telegram payment identifier for the subscription
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Pass *True* to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass *False* to allow the user to re-enable a subscription that was previously canceled by the bot.
	IsCanceled bool `json:"is_canceled"`
}

type _PostEditUserStarSubscriptionRequest PostEditUserStarSubscriptionRequest

// NewPostEditUserStarSubscriptionRequest instantiates a new PostEditUserStarSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostEditUserStarSubscriptionRequest(userId int32, telegramPaymentChargeId string, isCanceled bool) *PostEditUserStarSubscriptionRequest {
	this := PostEditUserStarSubscriptionRequest{}
	this.UserId = userId
	this.TelegramPaymentChargeId = telegramPaymentChargeId
	this.IsCanceled = isCanceled
	return &this
}

// NewPostEditUserStarSubscriptionRequestWithDefaults instantiates a new PostEditUserStarSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostEditUserStarSubscriptionRequestWithDefaults() *PostEditUserStarSubscriptionRequest {
	this := PostEditUserStarSubscriptionRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PostEditUserStarSubscriptionRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostEditUserStarSubscriptionRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostEditUserStarSubscriptionRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field value
func (o *PostEditUserStarSubscriptionRequest) GetTelegramPaymentChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelegramPaymentChargeId
}

// GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field value
// and a boolean to check if the value has been set.
func (o *PostEditUserStarSubscriptionRequest) GetTelegramPaymentChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelegramPaymentChargeId, true
}

// SetTelegramPaymentChargeId sets field value
func (o *PostEditUserStarSubscriptionRequest) SetTelegramPaymentChargeId(v string) {
	o.TelegramPaymentChargeId = v
}

// GetIsCanceled returns the IsCanceled field value
func (o *PostEditUserStarSubscriptionRequest) GetIsCanceled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCanceled
}

// GetIsCanceledOk returns a tuple with the IsCanceled field value
// and a boolean to check if the value has been set.
func (o *PostEditUserStarSubscriptionRequest) GetIsCanceledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCanceled, true
}

// SetIsCanceled sets field value
func (o *PostEditUserStarSubscriptionRequest) SetIsCanceled(v bool) {
	o.IsCanceled = v
}

func (o PostEditUserStarSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEditUserStarSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["telegram_payment_charge_id"] = o.TelegramPaymentChargeId
	toSerialize["is_canceled"] = o.IsCanceled
	return toSerialize, nil
}

func (o *PostEditUserStarSubscriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"telegram_payment_charge_id",
		"is_canceled",
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

	varPostEditUserStarSubscriptionRequest := _PostEditUserStarSubscriptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEditUserStarSubscriptionRequest)

	if err != nil {
		return err
	}

	*o = PostEditUserStarSubscriptionRequest(varPostEditUserStarSubscriptionRequest)

	return err
}

type NullablePostEditUserStarSubscriptionRequest struct {
	value *PostEditUserStarSubscriptionRequest
	isSet bool
}

func (v NullablePostEditUserStarSubscriptionRequest) Get() *PostEditUserStarSubscriptionRequest {
	return v.value
}

func (v *NullablePostEditUserStarSubscriptionRequest) Set(val *PostEditUserStarSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEditUserStarSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEditUserStarSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEditUserStarSubscriptionRequest(val *PostEditUserStarSubscriptionRequest) *NullablePostEditUserStarSubscriptionRequest {
	return &NullablePostEditUserStarSubscriptionRequest{value: val, isSet: true}
}

func (v NullablePostEditUserStarSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEditUserStarSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


