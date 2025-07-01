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

// checks if the StarTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StarTransaction{}

// StarTransaction Describes a Telegram Star transaction. Note that if the buyer initiates a chargeback with the payment provider from whom they acquired Stars (e.g., Apple, Google) following this transaction, the refunded Stars will be deducted from the bot's balance. This is outside of Telegram's control.
type StarTransaction struct {
	// Unique identifier of the transaction. Coincides with the identifier of the original transaction for refund transactions. Coincides with *SuccessfulPayment.telegram\\_payment\\_charge\\_id* for successful incoming payments from users.
	Id string `json:"id"`
	// Integer amount of Telegram Stars transferred by the transaction
	Amount int32 `json:"amount"`
	// *Optional*. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999
	NanostarAmount *int32 `json:"nanostar_amount,omitempty"`
	// Date the transaction was created in Unix time
	Date int32 `json:"date"`
	Source *TransactionPartner `json:"source,omitempty"`
	Receiver *TransactionPartner `json:"receiver,omitempty"`
}

type _StarTransaction StarTransaction

// NewStarTransaction instantiates a new StarTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStarTransaction(id string, amount int32, date int32) *StarTransaction {
	this := StarTransaction{}
	this.Id = id
	this.Amount = amount
	this.Date = date
	return &this
}

// NewStarTransactionWithDefaults instantiates a new StarTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStarTransactionWithDefaults() *StarTransaction {
	this := StarTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *StarTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StarTransaction) SetId(v string) {
	o.Id = v
}

// GetAmount returns the Amount field value
func (o *StarTransaction) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StarTransaction) SetAmount(v int32) {
	o.Amount = v
}

// GetNanostarAmount returns the NanostarAmount field value if set, zero value otherwise.
func (o *StarTransaction) GetNanostarAmount() int32 {
	if o == nil || IsNil(o.NanostarAmount) {
		var ret int32
		return ret
	}
	return *o.NanostarAmount
}

// GetNanostarAmountOk returns a tuple with the NanostarAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetNanostarAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.NanostarAmount) {
		return nil, false
	}
	return o.NanostarAmount, true
}

// HasNanostarAmount returns a boolean if a field has been set.
func (o *StarTransaction) HasNanostarAmount() bool {
	if o != nil && !IsNil(o.NanostarAmount) {
		return true
	}

	return false
}

// SetNanostarAmount gets a reference to the given int32 and assigns it to the NanostarAmount field.
func (o *StarTransaction) SetNanostarAmount(v int32) {
	o.NanostarAmount = &v
}


// GetDate returns the Date field value
func (o *StarTransaction) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *StarTransaction) SetDate(v int32) {
	o.Date = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StarTransaction) GetSource() TransactionPartner {
	if o == nil || IsNil(o.Source) {
		var ret TransactionPartner
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetSourceOk() (*TransactionPartner, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StarTransaction) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given TransactionPartner and assigns it to the Source field.
func (o *StarTransaction) SetSource(v TransactionPartner) {
	o.Source = &v
}


// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *StarTransaction) GetReceiver() TransactionPartner {
	if o == nil || IsNil(o.Receiver) {
		var ret TransactionPartner
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarTransaction) GetReceiverOk() (*TransactionPartner, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *StarTransaction) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given TransactionPartner and assigns it to the Receiver field.
func (o *StarTransaction) SetReceiver(v TransactionPartner) {
	o.Receiver = &v
}


func (o StarTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StarTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["amount"] = o.Amount
	if !IsNil(o.NanostarAmount) {
		toSerialize["nanostar_amount"] = o.NanostarAmount
	}
	toSerialize["date"] = o.Date
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	return toSerialize, nil
}

func (o *StarTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"amount",
		"date",
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

	varStarTransaction := _StarTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStarTransaction)

	if err != nil {
		return err
	}

	*o = StarTransaction(varStarTransaction)

	return err
}

type NullableStarTransaction struct {
	value *StarTransaction
	isSet bool
}

func (v NullableStarTransaction) Get() *StarTransaction {
	return v.value
}

func (v *NullableStarTransaction) Set(val *StarTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableStarTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableStarTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStarTransaction(val *StarTransaction) *NullableStarTransaction {
	return &NullableStarTransaction{value: val, isSet: true}
}

func (v NullableStarTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStarTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


