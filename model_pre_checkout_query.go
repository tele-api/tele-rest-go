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

// checks if the PreCheckoutQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreCheckoutQuery{}

// PreCheckoutQuery This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// Unique query identifier
	Id string `json:"id"`
	From User `json:"from"`
	// Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90)
	Currency string `json:"currency"`
	// Total price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int32 `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// *Optional*. Identifier of the shipping option chosen by the user
	ShippingOptionId *string `json:"shipping_option_id,omitempty"`
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

type _PreCheckoutQuery PreCheckoutQuery

// NewPreCheckoutQuery instantiates a new PreCheckoutQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreCheckoutQuery(id string, from User, currency string, totalAmount int32, invoicePayload string) *PreCheckoutQuery {
	this := PreCheckoutQuery{}
	this.Id = id
	this.From = from
	this.Currency = currency
	this.TotalAmount = totalAmount
	this.InvoicePayload = invoicePayload
	return &this
}

// NewPreCheckoutQueryWithDefaults instantiates a new PreCheckoutQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreCheckoutQueryWithDefaults() *PreCheckoutQuery {
	this := PreCheckoutQuery{}
	return &this
}

// GetId returns the Id field value
func (o *PreCheckoutQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PreCheckoutQuery) SetId(v string) {
	o.Id = v
}

// GetFrom returns the From field value
func (o *PreCheckoutQuery) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PreCheckoutQuery) SetFrom(v User) {
	o.From = v
}

// GetCurrency returns the Currency field value
func (o *PreCheckoutQuery) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PreCheckoutQuery) SetCurrency(v string) {
	o.Currency = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *PreCheckoutQuery) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetTotalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PreCheckoutQuery) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetInvoicePayload returns the InvoicePayload field value
func (o *PreCheckoutQuery) GetInvoicePayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoicePayload
}

// GetInvoicePayloadOk returns a tuple with the InvoicePayload field value
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetInvoicePayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicePayload, true
}

// SetInvoicePayload sets field value
func (o *PreCheckoutQuery) SetInvoicePayload(v string) {
	o.InvoicePayload = v
}

// GetShippingOptionId returns the ShippingOptionId field value if set, zero value otherwise.
func (o *PreCheckoutQuery) GetShippingOptionId() string {
	if o == nil || IsNil(o.ShippingOptionId) {
		var ret string
		return ret
	}
	return *o.ShippingOptionId
}

// GetShippingOptionIdOk returns a tuple with the ShippingOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetShippingOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingOptionId) {
		return nil, false
	}
	return o.ShippingOptionId, true
}

// HasShippingOptionId returns a boolean if a field has been set.
func (o *PreCheckoutQuery) HasShippingOptionId() bool {
	if o != nil && !IsNil(o.ShippingOptionId) {
		return true
	}

	return false
}

// SetShippingOptionId gets a reference to the given string and assigns it to the ShippingOptionId field.
func (o *PreCheckoutQuery) SetShippingOptionId(v string) {
	o.ShippingOptionId = &v
}


// GetOrderInfo returns the OrderInfo field value if set, zero value otherwise.
func (o *PreCheckoutQuery) GetOrderInfo() OrderInfo {
	if o == nil || IsNil(o.OrderInfo) {
		var ret OrderInfo
		return ret
	}
	return *o.OrderInfo
}

// GetOrderInfoOk returns a tuple with the OrderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckoutQuery) GetOrderInfoOk() (*OrderInfo, bool) {
	if o == nil || IsNil(o.OrderInfo) {
		return nil, false
	}
	return o.OrderInfo, true
}

// HasOrderInfo returns a boolean if a field has been set.
func (o *PreCheckoutQuery) HasOrderInfo() bool {
	if o != nil && !IsNil(o.OrderInfo) {
		return true
	}

	return false
}

// SetOrderInfo gets a reference to the given OrderInfo and assigns it to the OrderInfo field.
func (o *PreCheckoutQuery) SetOrderInfo(v OrderInfo) {
	o.OrderInfo = &v
}


func (o PreCheckoutQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreCheckoutQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["from"] = o.From
	toSerialize["currency"] = o.Currency
	toSerialize["total_amount"] = o.TotalAmount
	toSerialize["invoice_payload"] = o.InvoicePayload
	if !IsNil(o.ShippingOptionId) {
		toSerialize["shipping_option_id"] = o.ShippingOptionId
	}
	if !IsNil(o.OrderInfo) {
		toSerialize["order_info"] = o.OrderInfo
	}
	return toSerialize, nil
}

func (o *PreCheckoutQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"from",
		"currency",
		"total_amount",
		"invoice_payload",
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

	varPreCheckoutQuery := _PreCheckoutQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreCheckoutQuery)

	if err != nil {
		return err
	}

	*o = PreCheckoutQuery(varPreCheckoutQuery)

	return err
}

type NullablePreCheckoutQuery struct {
	value *PreCheckoutQuery
	isSet bool
}

func (v NullablePreCheckoutQuery) Get() *PreCheckoutQuery {
	return v.value
}

func (v *NullablePreCheckoutQuery) Set(val *PreCheckoutQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePreCheckoutQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePreCheckoutQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreCheckoutQuery(val *PreCheckoutQuery) *NullablePreCheckoutQuery {
	return &NullablePreCheckoutQuery{value: val, isSet: true}
}

func (v NullablePreCheckoutQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreCheckoutQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


