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

// checks if the StarAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StarAmount{}

// StarAmount Describes an amount of Telegram Stars.
type StarAmount struct {
	// Integer amount of Telegram Stars, rounded to 0; can be negative
	Amount int32 `json:"amount"`
	// *Optional*. The number of 1/1000000000 shares of Telegram Stars; from -999999999 to 999999999; can be negative if and only if *amount* is non-positive
	NanostarAmount *int32 `json:"nanostar_amount,omitempty"`
}

type _StarAmount StarAmount

// NewStarAmount instantiates a new StarAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStarAmount(amount int32) *StarAmount {
	this := StarAmount{}
	this.Amount = amount
	return &this
}

// NewStarAmountWithDefaults instantiates a new StarAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStarAmountWithDefaults() *StarAmount {
	this := StarAmount{}
	return &this
}

// GetAmount returns the Amount field value
func (o *StarAmount) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StarAmount) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StarAmount) SetAmount(v int32) {
	o.Amount = v
}

// GetNanostarAmount returns the NanostarAmount field value if set, zero value otherwise.
func (o *StarAmount) GetNanostarAmount() int32 {
	if o == nil || IsNil(o.NanostarAmount) {
		var ret int32
		return ret
	}
	return *o.NanostarAmount
}

// GetNanostarAmountOk returns a tuple with the NanostarAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarAmount) GetNanostarAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.NanostarAmount) {
		return nil, false
	}
	return o.NanostarAmount, true
}

// HasNanostarAmount returns a boolean if a field has been set.
func (o *StarAmount) HasNanostarAmount() bool {
	if o != nil && !IsNil(o.NanostarAmount) {
		return true
	}

	return false
}

// SetNanostarAmount gets a reference to the given int32 and assigns it to the NanostarAmount field.
func (o *StarAmount) SetNanostarAmount(v int32) {
	o.NanostarAmount = &v
}


func (o StarAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StarAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.NanostarAmount) {
		toSerialize["nanostar_amount"] = o.NanostarAmount
	}
	return toSerialize, nil
}

func (o *StarAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varStarAmount := _StarAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStarAmount)

	if err != nil {
		return err
	}

	*o = StarAmount(varStarAmount)

	return err
}

type NullableStarAmount struct {
	value *StarAmount
	isSet bool
}

func (v NullableStarAmount) Get() *StarAmount {
	return v.value
}

func (v *NullableStarAmount) Set(val *StarAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableStarAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableStarAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStarAmount(val *StarAmount) *NullableStarAmount {
	return &NullableStarAmount{value: val, isSet: true}
}

func (v NullableStarAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStarAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


