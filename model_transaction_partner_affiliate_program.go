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

// checks if the TransactionPartnerAffiliateProgram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionPartnerAffiliateProgram{}

// TransactionPartnerAffiliateProgram Describes the affiliate program that issued the affiliate commission received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	// Type of the transaction partner, always “affiliate\\_program”
	Type string `json:"type"`
	SponsorUser *User `json:"sponsor_user,omitempty"`
	// The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users
	CommissionPerMille int32 `json:"commission_per_mille"`
}

type _TransactionPartnerAffiliateProgram TransactionPartnerAffiliateProgram

// NewTransactionPartnerAffiliateProgram instantiates a new TransactionPartnerAffiliateProgram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionPartnerAffiliateProgram(type_ string, commissionPerMille int32) *TransactionPartnerAffiliateProgram {
	this := TransactionPartnerAffiliateProgram{}
	this.Type = type_
	this.CommissionPerMille = commissionPerMille
	return &this
}

// NewTransactionPartnerAffiliateProgramWithDefaults instantiates a new TransactionPartnerAffiliateProgram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionPartnerAffiliateProgramWithDefaults() *TransactionPartnerAffiliateProgram {
	this := TransactionPartnerAffiliateProgram{}
	var type_ string = "affiliate_program"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TransactionPartnerAffiliateProgram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionPartnerAffiliateProgram) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionPartnerAffiliateProgram) SetType(v string) {
	o.Type = v
}

// GetSponsorUser returns the SponsorUser field value if set, zero value otherwise.
func (o *TransactionPartnerAffiliateProgram) GetSponsorUser() User {
	if o == nil || IsNil(o.SponsorUser) {
		var ret User
		return ret
	}
	return *o.SponsorUser
}

// GetSponsorUserOk returns a tuple with the SponsorUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerAffiliateProgram) GetSponsorUserOk() (*User, bool) {
	if o == nil || IsNil(o.SponsorUser) {
		return nil, false
	}
	return o.SponsorUser, true
}

// HasSponsorUser returns a boolean if a field has been set.
func (o *TransactionPartnerAffiliateProgram) HasSponsorUser() bool {
	if o != nil && !IsNil(o.SponsorUser) {
		return true
	}

	return false
}

// SetSponsorUser gets a reference to the given User and assigns it to the SponsorUser field.
func (o *TransactionPartnerAffiliateProgram) SetSponsorUser(v User) {
	o.SponsorUser = &v
}


// GetCommissionPerMille returns the CommissionPerMille field value
func (o *TransactionPartnerAffiliateProgram) GetCommissionPerMille() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommissionPerMille
}

// GetCommissionPerMilleOk returns a tuple with the CommissionPerMille field value
// and a boolean to check if the value has been set.
func (o *TransactionPartnerAffiliateProgram) GetCommissionPerMilleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommissionPerMille, true
}

// SetCommissionPerMille sets field value
func (o *TransactionPartnerAffiliateProgram) SetCommissionPerMille(v int32) {
	o.CommissionPerMille = v
}

func (o TransactionPartnerAffiliateProgram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionPartnerAffiliateProgram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.SponsorUser) {
		toSerialize["sponsor_user"] = o.SponsorUser
	}
	toSerialize["commission_per_mille"] = o.CommissionPerMille
	return toSerialize, nil
}

func (o *TransactionPartnerAffiliateProgram) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"commission_per_mille",
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

	varTransactionPartnerAffiliateProgram := _TransactionPartnerAffiliateProgram{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionPartnerAffiliateProgram)

	if err != nil {
		return err
	}

	*o = TransactionPartnerAffiliateProgram(varTransactionPartnerAffiliateProgram)

	return err
}

type NullableTransactionPartnerAffiliateProgram struct {
	value *TransactionPartnerAffiliateProgram
	isSet bool
}

func (v NullableTransactionPartnerAffiliateProgram) Get() *TransactionPartnerAffiliateProgram {
	return v.value
}

func (v *NullableTransactionPartnerAffiliateProgram) Set(val *TransactionPartnerAffiliateProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPartnerAffiliateProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPartnerAffiliateProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPartnerAffiliateProgram(val *TransactionPartnerAffiliateProgram) *NullableTransactionPartnerAffiliateProgram {
	return &NullableTransactionPartnerAffiliateProgram{value: val, isSet: true}
}

func (v NullableTransactionPartnerAffiliateProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPartnerAffiliateProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


