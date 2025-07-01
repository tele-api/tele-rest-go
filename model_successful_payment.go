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

// checks if the SuccessfulPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessfulPayment{}

// SuccessfulPayment This object contains basic information about a successful payment. Note that if the buyer initiates a chargeback with the relevant payment provider following this transaction, the funds may be debited from your balance. This is outside of Telegram's control.
type SuccessfulPayment struct {
	// Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90)
	Currency string `json:"currency"`
	// Total price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int32 `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// *Optional*. Expiration date of the subscription, in Unix time; for recurring payments only
	SubscriptionExpirationDate *int32 `json:"subscription_expiration_date,omitempty"`
	// *Optional*. True, if the payment is a recurring payment for a subscription
	IsRecurring *bool `json:"is_recurring,omitempty"`
	// *Optional*. True, if the payment is the first payment for a subscription
	IsFirstRecurring *bool `json:"is_first_recurring,omitempty"`
	// *Optional*. Identifier of the shipping option chosen by the user
	ShippingOptionId *string `json:"shipping_option_id,omitempty"`
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

type _SuccessfulPayment SuccessfulPayment

// NewSuccessfulPayment instantiates a new SuccessfulPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulPayment(currency string, totalAmount int32, invoicePayload string, telegramPaymentChargeId string, providerPaymentChargeId string) *SuccessfulPayment {
	this := SuccessfulPayment{}
	this.Currency = currency
	this.TotalAmount = totalAmount
	this.InvoicePayload = invoicePayload
	var isRecurring bool = true
	this.IsRecurring = &isRecurring
	var isFirstRecurring bool = true
	this.IsFirstRecurring = &isFirstRecurring
	this.TelegramPaymentChargeId = telegramPaymentChargeId
	this.ProviderPaymentChargeId = providerPaymentChargeId
	return &this
}

// NewSuccessfulPaymentWithDefaults instantiates a new SuccessfulPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulPaymentWithDefaults() *SuccessfulPayment {
	this := SuccessfulPayment{}
	var isRecurring bool = true
	this.IsRecurring = &isRecurring
	var isFirstRecurring bool = true
	this.IsFirstRecurring = &isFirstRecurring
	return &this
}

// GetCurrency returns the Currency field value
func (o *SuccessfulPayment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SuccessfulPayment) SetCurrency(v string) {
	o.Currency = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *SuccessfulPayment) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetTotalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *SuccessfulPayment) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetInvoicePayload returns the InvoicePayload field value
func (o *SuccessfulPayment) GetInvoicePayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoicePayload
}

// GetInvoicePayloadOk returns a tuple with the InvoicePayload field value
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetInvoicePayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicePayload, true
}

// SetInvoicePayload sets field value
func (o *SuccessfulPayment) SetInvoicePayload(v string) {
	o.InvoicePayload = v
}

// GetSubscriptionExpirationDate returns the SubscriptionExpirationDate field value if set, zero value otherwise.
func (o *SuccessfulPayment) GetSubscriptionExpirationDate() int32 {
	if o == nil || IsNil(o.SubscriptionExpirationDate) {
		var ret int32
		return ret
	}
	return *o.SubscriptionExpirationDate
}

// GetSubscriptionExpirationDateOk returns a tuple with the SubscriptionExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetSubscriptionExpirationDateOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionExpirationDate) {
		return nil, false
	}
	return o.SubscriptionExpirationDate, true
}

// HasSubscriptionExpirationDate returns a boolean if a field has been set.
func (o *SuccessfulPayment) HasSubscriptionExpirationDate() bool {
	if o != nil && !IsNil(o.SubscriptionExpirationDate) {
		return true
	}

	return false
}

// SetSubscriptionExpirationDate gets a reference to the given int32 and assigns it to the SubscriptionExpirationDate field.
func (o *SuccessfulPayment) SetSubscriptionExpirationDate(v int32) {
	o.SubscriptionExpirationDate = &v
}


// GetIsRecurring returns the IsRecurring field value if set, zero value otherwise.
func (o *SuccessfulPayment) GetIsRecurring() bool {
	if o == nil || IsNil(o.IsRecurring) {
		var ret bool
		return ret
	}
	return *o.IsRecurring
}

// GetIsRecurringOk returns a tuple with the IsRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetIsRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRecurring) {
		return nil, false
	}
	return o.IsRecurring, true
}

// HasIsRecurring returns a boolean if a field has been set.
func (o *SuccessfulPayment) HasIsRecurring() bool {
	if o != nil && !IsNil(o.IsRecurring) {
		return true
	}

	return false
}

// SetIsRecurring gets a reference to the given bool and assigns it to the IsRecurring field.
func (o *SuccessfulPayment) SetIsRecurring(v bool) {
	o.IsRecurring = &v
}


// GetIsFirstRecurring returns the IsFirstRecurring field value if set, zero value otherwise.
func (o *SuccessfulPayment) GetIsFirstRecurring() bool {
	if o == nil || IsNil(o.IsFirstRecurring) {
		var ret bool
		return ret
	}
	return *o.IsFirstRecurring
}

// GetIsFirstRecurringOk returns a tuple with the IsFirstRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetIsFirstRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFirstRecurring) {
		return nil, false
	}
	return o.IsFirstRecurring, true
}

// HasIsFirstRecurring returns a boolean if a field has been set.
func (o *SuccessfulPayment) HasIsFirstRecurring() bool {
	if o != nil && !IsNil(o.IsFirstRecurring) {
		return true
	}

	return false
}

// SetIsFirstRecurring gets a reference to the given bool and assigns it to the IsFirstRecurring field.
func (o *SuccessfulPayment) SetIsFirstRecurring(v bool) {
	o.IsFirstRecurring = &v
}


// GetShippingOptionId returns the ShippingOptionId field value if set, zero value otherwise.
func (o *SuccessfulPayment) GetShippingOptionId() string {
	if o == nil || IsNil(o.ShippingOptionId) {
		var ret string
		return ret
	}
	return *o.ShippingOptionId
}

// GetShippingOptionIdOk returns a tuple with the ShippingOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetShippingOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingOptionId) {
		return nil, false
	}
	return o.ShippingOptionId, true
}

// HasShippingOptionId returns a boolean if a field has been set.
func (o *SuccessfulPayment) HasShippingOptionId() bool {
	if o != nil && !IsNil(o.ShippingOptionId) {
		return true
	}

	return false
}

// SetShippingOptionId gets a reference to the given string and assigns it to the ShippingOptionId field.
func (o *SuccessfulPayment) SetShippingOptionId(v string) {
	o.ShippingOptionId = &v
}


// GetOrderInfo returns the OrderInfo field value if set, zero value otherwise.
func (o *SuccessfulPayment) GetOrderInfo() OrderInfo {
	if o == nil || IsNil(o.OrderInfo) {
		var ret OrderInfo
		return ret
	}
	return *o.OrderInfo
}

// GetOrderInfoOk returns a tuple with the OrderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetOrderInfoOk() (*OrderInfo, bool) {
	if o == nil || IsNil(o.OrderInfo) {
		return nil, false
	}
	return o.OrderInfo, true
}

// HasOrderInfo returns a boolean if a field has been set.
func (o *SuccessfulPayment) HasOrderInfo() bool {
	if o != nil && !IsNil(o.OrderInfo) {
		return true
	}

	return false
}

// SetOrderInfo gets a reference to the given OrderInfo and assigns it to the OrderInfo field.
func (o *SuccessfulPayment) SetOrderInfo(v OrderInfo) {
	o.OrderInfo = &v
}


// GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field value
func (o *SuccessfulPayment) GetTelegramPaymentChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelegramPaymentChargeId
}

// GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field value
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetTelegramPaymentChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelegramPaymentChargeId, true
}

// SetTelegramPaymentChargeId sets field value
func (o *SuccessfulPayment) SetTelegramPaymentChargeId(v string) {
	o.TelegramPaymentChargeId = v
}

// GetProviderPaymentChargeId returns the ProviderPaymentChargeId field value
func (o *SuccessfulPayment) GetProviderPaymentChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderPaymentChargeId
}

// GetProviderPaymentChargeIdOk returns a tuple with the ProviderPaymentChargeId field value
// and a boolean to check if the value has been set.
func (o *SuccessfulPayment) GetProviderPaymentChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderPaymentChargeId, true
}

// SetProviderPaymentChargeId sets field value
func (o *SuccessfulPayment) SetProviderPaymentChargeId(v string) {
	o.ProviderPaymentChargeId = v
}

func (o SuccessfulPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessfulPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["total_amount"] = o.TotalAmount
	toSerialize["invoice_payload"] = o.InvoicePayload
	if !IsNil(o.SubscriptionExpirationDate) {
		toSerialize["subscription_expiration_date"] = o.SubscriptionExpirationDate
	}
	if !IsNil(o.IsRecurring) {
		toSerialize["is_recurring"] = o.IsRecurring
	}
	if !IsNil(o.IsFirstRecurring) {
		toSerialize["is_first_recurring"] = o.IsFirstRecurring
	}
	if !IsNil(o.ShippingOptionId) {
		toSerialize["shipping_option_id"] = o.ShippingOptionId
	}
	if !IsNil(o.OrderInfo) {
		toSerialize["order_info"] = o.OrderInfo
	}
	toSerialize["telegram_payment_charge_id"] = o.TelegramPaymentChargeId
	toSerialize["provider_payment_charge_id"] = o.ProviderPaymentChargeId
	return toSerialize, nil
}

func (o *SuccessfulPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"total_amount",
		"invoice_payload",
		"telegram_payment_charge_id",
		"provider_payment_charge_id",
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

	varSuccessfulPayment := _SuccessfulPayment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuccessfulPayment)

	if err != nil {
		return err
	}

	*o = SuccessfulPayment(varSuccessfulPayment)

	return err
}

type NullableSuccessfulPayment struct {
	value *SuccessfulPayment
	isSet bool
}

func (v NullableSuccessfulPayment) Get() *SuccessfulPayment {
	return v.value
}

func (v *NullableSuccessfulPayment) Set(val *SuccessfulPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulPayment(val *SuccessfulPayment) *NullableSuccessfulPayment {
	return &NullableSuccessfulPayment{value: val, isSet: true}
}

func (v NullableSuccessfulPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


