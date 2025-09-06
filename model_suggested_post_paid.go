/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-09-06T05:32:05.180087835Z[Etc/UTC]
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

// checks if the SuggestedPostPaid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostPaid{}

// SuggestedPostPaid Describes a service message about a successful payment for a suggested post.
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	// Currency in which the payment was made. Currently, one of “XTR” for Telegram Stars or “TON” for toncoins
	Currency string `json:"currency"`
	// *Optional*. The amount of the currency that was received by the channel in nanotoncoins; for payments in toncoins only
	Amount *int32 `json:"amount,omitempty"`
	StarAmount *StarAmount `json:"star_amount,omitempty"`
}

type _SuggestedPostPaid SuggestedPostPaid

// NewSuggestedPostPaid instantiates a new SuggestedPostPaid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostPaid(currency string) *SuggestedPostPaid {
	this := SuggestedPostPaid{}
	this.Currency = currency
	return &this
}

// NewSuggestedPostPaidWithDefaults instantiates a new SuggestedPostPaid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostPaidWithDefaults() *SuggestedPostPaid {
	this := SuggestedPostPaid{}
	return &this
}

// GetSuggestedPostMessage returns the SuggestedPostMessage field value if set, zero value otherwise.
func (o *SuggestedPostPaid) GetSuggestedPostMessage() Message {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		var ret Message
		return ret
	}
	return *o.SuggestedPostMessage
}

// GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostPaid) GetSuggestedPostMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.SuggestedPostMessage) {
		return nil, false
	}
	return o.SuggestedPostMessage, true
}

// HasSuggestedPostMessage returns a boolean if a field has been set.
func (o *SuggestedPostPaid) HasSuggestedPostMessage() bool {
	if o != nil && !IsNil(o.SuggestedPostMessage) {
		return true
	}

	return false
}

// SetSuggestedPostMessage gets a reference to the given Message and assigns it to the SuggestedPostMessage field.
func (o *SuggestedPostPaid) SetSuggestedPostMessage(v Message) {
	o.SuggestedPostMessage = &v
}


// GetCurrency returns the Currency field value
func (o *SuggestedPostPaid) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SuggestedPostPaid) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SuggestedPostPaid) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SuggestedPostPaid) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostPaid) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SuggestedPostPaid) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *SuggestedPostPaid) SetAmount(v int32) {
	o.Amount = &v
}


// GetStarAmount returns the StarAmount field value if set, zero value otherwise.
func (o *SuggestedPostPaid) GetStarAmount() StarAmount {
	if o == nil || IsNil(o.StarAmount) {
		var ret StarAmount
		return ret
	}
	return *o.StarAmount
}

// GetStarAmountOk returns a tuple with the StarAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostPaid) GetStarAmountOk() (*StarAmount, bool) {
	if o == nil || IsNil(o.StarAmount) {
		return nil, false
	}
	return o.StarAmount, true
}

// HasStarAmount returns a boolean if a field has been set.
func (o *SuggestedPostPaid) HasStarAmount() bool {
	if o != nil && !IsNil(o.StarAmount) {
		return true
	}

	return false
}

// SetStarAmount gets a reference to the given StarAmount and assigns it to the StarAmount field.
func (o *SuggestedPostPaid) SetStarAmount(v StarAmount) {
	o.StarAmount = &v
}


func (o SuggestedPostPaid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostPaid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedPostMessage) {
		toSerialize["suggested_post_message"] = o.SuggestedPostMessage
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.StarAmount) {
		toSerialize["star_amount"] = o.StarAmount
	}
	return toSerialize, nil
}

func (o *SuggestedPostPaid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
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

	varSuggestedPostPaid := _SuggestedPostPaid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestedPostPaid)

	if err != nil {
		return err
	}

	*o = SuggestedPostPaid(varSuggestedPostPaid)

	return err
}

type NullableSuggestedPostPaid struct {
	value *SuggestedPostPaid
	isSet bool
}

func (v NullableSuggestedPostPaid) Get() *SuggestedPostPaid {
	return v.value
}

func (v *NullableSuggestedPostPaid) Set(val *SuggestedPostPaid) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostPaid) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostPaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostPaid(val *SuggestedPostPaid) *NullableSuggestedPostPaid {
	return &NullableSuggestedPostPaid{value: val, isSet: true}
}

func (v NullableSuggestedPostPaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostPaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


