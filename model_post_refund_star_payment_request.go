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

// checks if the PostRefundStarPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRefundStarPaymentRequest{}

// PostRefundStarPaymentRequest struct for PostRefundStarPaymentRequest
type PostRefundStarPaymentRequest struct {
	// Identifier of the user whose payment will be refunded
	UserId int32 `json:"user_id"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
}

type _PostRefundStarPaymentRequest PostRefundStarPaymentRequest

// NewPostRefundStarPaymentRequest instantiates a new PostRefundStarPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRefundStarPaymentRequest(userId int32, telegramPaymentChargeId string) *PostRefundStarPaymentRequest {
	this := PostRefundStarPaymentRequest{}
	this.UserId = userId
	this.TelegramPaymentChargeId = telegramPaymentChargeId
	return &this
}

// NewPostRefundStarPaymentRequestWithDefaults instantiates a new PostRefundStarPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRefundStarPaymentRequestWithDefaults() *PostRefundStarPaymentRequest {
	this := PostRefundStarPaymentRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PostRefundStarPaymentRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostRefundStarPaymentRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostRefundStarPaymentRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field value
func (o *PostRefundStarPaymentRequest) GetTelegramPaymentChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelegramPaymentChargeId
}

// GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field value
// and a boolean to check if the value has been set.
func (o *PostRefundStarPaymentRequest) GetTelegramPaymentChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelegramPaymentChargeId, true
}

// SetTelegramPaymentChargeId sets field value
func (o *PostRefundStarPaymentRequest) SetTelegramPaymentChargeId(v string) {
	o.TelegramPaymentChargeId = v
}

func (o PostRefundStarPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRefundStarPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["telegram_payment_charge_id"] = o.TelegramPaymentChargeId
	return toSerialize, nil
}

func (o *PostRefundStarPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"telegram_payment_charge_id",
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

	varPostRefundStarPaymentRequest := _PostRefundStarPaymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostRefundStarPaymentRequest)

	if err != nil {
		return err
	}

	*o = PostRefundStarPaymentRequest(varPostRefundStarPaymentRequest)

	return err
}

type NullablePostRefundStarPaymentRequest struct {
	value *PostRefundStarPaymentRequest
	isSet bool
}

func (v NullablePostRefundStarPaymentRequest) Get() *PostRefundStarPaymentRequest {
	return v.value
}

func (v *NullablePostRefundStarPaymentRequest) Set(val *PostRefundStarPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRefundStarPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRefundStarPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRefundStarPaymentRequest(val *PostRefundStarPaymentRequest) *NullablePostRefundStarPaymentRequest {
	return &NullablePostRefundStarPaymentRequest{value: val, isSet: true}
}

func (v NullablePostRefundStarPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRefundStarPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


