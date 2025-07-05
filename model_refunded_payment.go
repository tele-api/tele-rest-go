/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the RefundedPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundedPayment{}

// RefundedPayment This object contains basic information about a refunded payment.
type RefundedPayment struct {
	// Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90). Currently, always “XTR”
	Currency string `json:"currency"`
	// Total refunded price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of `US$ 1.45`, `total_amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int32 `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// *Optional*. Provider payment identifier
	ProviderPaymentChargeId *string `json:"provider_payment_charge_id,omitempty"`
}

type _RefundedPayment RefundedPayment

// NewRefundedPayment instantiates a new RefundedPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundedPayment(currency string, totalAmount int32, invoicePayload string, telegramPaymentChargeId string) *RefundedPayment {
	this := RefundedPayment{}
	this.Currency = currency
	this.TotalAmount = totalAmount
	this.InvoicePayload = invoicePayload
	this.TelegramPaymentChargeId = telegramPaymentChargeId
	return &this
}

// NewRefundedPaymentWithDefaults instantiates a new RefundedPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundedPaymentWithDefaults() *RefundedPayment {
	this := RefundedPayment{}
	var currency string = "XTR"
	this.Currency = currency
	return &this
}

// GetCurrency returns the Currency field value
func (o *RefundedPayment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RefundedPayment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RefundedPayment) SetCurrency(v string) {
	o.Currency = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *RefundedPayment) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *RefundedPayment) GetTotalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *RefundedPayment) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetInvoicePayload returns the InvoicePayload field value
func (o *RefundedPayment) GetInvoicePayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoicePayload
}

// GetInvoicePayloadOk returns a tuple with the InvoicePayload field value
// and a boolean to check if the value has been set.
func (o *RefundedPayment) GetInvoicePayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicePayload, true
}

// SetInvoicePayload sets field value
func (o *RefundedPayment) SetInvoicePayload(v string) {
	o.InvoicePayload = v
}

// GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field value
func (o *RefundedPayment) GetTelegramPaymentChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelegramPaymentChargeId
}

// GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field value
// and a boolean to check if the value has been set.
func (o *RefundedPayment) GetTelegramPaymentChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelegramPaymentChargeId, true
}

// SetTelegramPaymentChargeId sets field value
func (o *RefundedPayment) SetTelegramPaymentChargeId(v string) {
	o.TelegramPaymentChargeId = v
}

// GetProviderPaymentChargeId returns the ProviderPaymentChargeId field value if set, zero value otherwise.
func (o *RefundedPayment) GetProviderPaymentChargeId() string {
	if o == nil || IsNil(o.ProviderPaymentChargeId) {
		var ret string
		return ret
	}
	return *o.ProviderPaymentChargeId
}

// GetProviderPaymentChargeIdOk returns a tuple with the ProviderPaymentChargeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundedPayment) GetProviderPaymentChargeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderPaymentChargeId) {
		return nil, false
	}
	return o.ProviderPaymentChargeId, true
}

// HasProviderPaymentChargeId returns a boolean if a field has been set.
func (o *RefundedPayment) HasProviderPaymentChargeId() bool {
	if o != nil && !IsNil(o.ProviderPaymentChargeId) {
		return true
	}

	return false
}

// SetProviderPaymentChargeId gets a reference to the given string and assigns it to the ProviderPaymentChargeId field.
func (o *RefundedPayment) SetProviderPaymentChargeId(v string) {
	o.ProviderPaymentChargeId = &v
}


func (o RefundedPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundedPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["total_amount"] = o.TotalAmount
	toSerialize["invoice_payload"] = o.InvoicePayload
	toSerialize["telegram_payment_charge_id"] = o.TelegramPaymentChargeId
	if !IsNil(o.ProviderPaymentChargeId) {
		toSerialize["provider_payment_charge_id"] = o.ProviderPaymentChargeId
	}
	return toSerialize, nil
}

func (o *RefundedPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"total_amount",
		"invoice_payload",
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

	varRefundedPayment := _RefundedPayment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundedPayment)

	if err != nil {
		return err
	}

	*o = RefundedPayment(varRefundedPayment)

	return err
}

type NullableRefundedPayment struct {
	value *RefundedPayment
	isSet bool
}

func (v NullableRefundedPayment) Get() *RefundedPayment {
	return v.value
}

func (v *NullableRefundedPayment) Set(val *RefundedPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundedPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundedPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundedPayment(val *RefundedPayment) *NullableRefundedPayment {
	return &NullableRefundedPayment{value: val, isSet: true}
}

func (v NullableRefundedPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundedPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


