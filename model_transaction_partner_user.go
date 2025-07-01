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

// checks if the TransactionPartnerUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionPartnerUser{}

// TransactionPartnerUser Describes a transaction with a user.
type TransactionPartnerUser struct {
	// Type of the transaction partner, always “user”
	Type string `json:"type"`
	// Type of the transaction, currently one of “invoice\\_payment” for payments via invoices, “paid\\_media\\_payment” for payments for paid media, “gift\\_purchase” for gifts sent by the bot, “premium\\_purchase” for Telegram Premium subscriptions gifted by the bot, “business\\_account\\_transfer” for direct transfers from managed business accounts
	TransactionType string `json:"transaction_type"`
	User User `json:"user"`
	Affiliate *AffiliateInfo `json:"affiliate,omitempty"`
	// *Optional*. Bot-specified invoice payload. Can be available only for “invoice\\_payment” transactions.
	InvoicePayload *string `json:"invoice_payload,omitempty"`
	// *Optional*. The duration of the paid subscription. Can be available only for “invoice\\_payment” transactions.
	SubscriptionPeriod *int32 `json:"subscription_period,omitempty"`
	// *Optional*. Information about the paid media bought by the user; for “paid\\_media\\_payment” transactions only
	PaidMedia []PaidMedia `json:"paid_media,omitempty"`
	// *Optional*. Bot-specified paid media payload. Can be available only for “paid\\_media\\_payment” transactions.
	PaidMediaPayload *string `json:"paid_media_payload,omitempty"`
	Gift *Gift `json:"gift,omitempty"`
	// *Optional*. Number of months the gifted Telegram Premium subscription will be active for; for “premium\\_purchase” transactions only
	PremiumSubscriptionDuration *int32 `json:"premium_subscription_duration,omitempty"`
}

type _TransactionPartnerUser TransactionPartnerUser

// NewTransactionPartnerUser instantiates a new TransactionPartnerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionPartnerUser(type_ string, transactionType string, user User) *TransactionPartnerUser {
	this := TransactionPartnerUser{}
	this.Type = type_
	this.TransactionType = transactionType
	this.User = user
	return &this
}

// NewTransactionPartnerUserWithDefaults instantiates a new TransactionPartnerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionPartnerUserWithDefaults() *TransactionPartnerUser {
	this := TransactionPartnerUser{}
	var type_ string = "user"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TransactionPartnerUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionPartnerUser) SetType(v string) {
	o.Type = v
}

// GetTransactionType returns the TransactionType field value
func (o *TransactionPartnerUser) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *TransactionPartnerUser) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetUser returns the User field value
func (o *TransactionPartnerUser) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransactionPartnerUser) SetUser(v User) {
	o.User = v
}

// GetAffiliate returns the Affiliate field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetAffiliate() AffiliateInfo {
	if o == nil || IsNil(o.Affiliate) {
		var ret AffiliateInfo
		return ret
	}
	return *o.Affiliate
}

// GetAffiliateOk returns a tuple with the Affiliate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetAffiliateOk() (*AffiliateInfo, bool) {
	if o == nil || IsNil(o.Affiliate) {
		return nil, false
	}
	return o.Affiliate, true
}

// HasAffiliate returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasAffiliate() bool {
	if o != nil && !IsNil(o.Affiliate) {
		return true
	}

	return false
}

// SetAffiliate gets a reference to the given AffiliateInfo and assigns it to the Affiliate field.
func (o *TransactionPartnerUser) SetAffiliate(v AffiliateInfo) {
	o.Affiliate = &v
}


// GetInvoicePayload returns the InvoicePayload field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetInvoicePayload() string {
	if o == nil || IsNil(o.InvoicePayload) {
		var ret string
		return ret
	}
	return *o.InvoicePayload
}

// GetInvoicePayloadOk returns a tuple with the InvoicePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetInvoicePayloadOk() (*string, bool) {
	if o == nil || IsNil(o.InvoicePayload) {
		return nil, false
	}
	return o.InvoicePayload, true
}

// HasInvoicePayload returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasInvoicePayload() bool {
	if o != nil && !IsNil(o.InvoicePayload) {
		return true
	}

	return false
}

// SetInvoicePayload gets a reference to the given string and assigns it to the InvoicePayload field.
func (o *TransactionPartnerUser) SetInvoicePayload(v string) {
	o.InvoicePayload = &v
}


// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetSubscriptionPeriod() int32 {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret int32
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetSubscriptionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given int32 and assigns it to the SubscriptionPeriod field.
func (o *TransactionPartnerUser) SetSubscriptionPeriod(v int32) {
	o.SubscriptionPeriod = &v
}


// GetPaidMedia returns the PaidMedia field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetPaidMedia() []PaidMedia {
	if o == nil || IsNil(o.PaidMedia) {
		var ret []PaidMedia
		return ret
	}
	return o.PaidMedia
}

// GetPaidMediaOk returns a tuple with the PaidMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetPaidMediaOk() ([]PaidMedia, bool) {
	if o == nil || IsNil(o.PaidMedia) {
		return nil, false
	}
	return o.PaidMedia, true
}

// HasPaidMedia returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasPaidMedia() bool {
	if o != nil && !IsNil(o.PaidMedia) {
		return true
	}

	return false
}

// SetPaidMedia gets a reference to the given []PaidMedia and assigns it to the PaidMedia field.
func (o *TransactionPartnerUser) SetPaidMedia(v []PaidMedia) {
	o.PaidMedia = v
}


// GetPaidMediaPayload returns the PaidMediaPayload field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetPaidMediaPayload() string {
	if o == nil || IsNil(o.PaidMediaPayload) {
		var ret string
		return ret
	}
	return *o.PaidMediaPayload
}

// GetPaidMediaPayloadOk returns a tuple with the PaidMediaPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetPaidMediaPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.PaidMediaPayload) {
		return nil, false
	}
	return o.PaidMediaPayload, true
}

// HasPaidMediaPayload returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasPaidMediaPayload() bool {
	if o != nil && !IsNil(o.PaidMediaPayload) {
		return true
	}

	return false
}

// SetPaidMediaPayload gets a reference to the given string and assigns it to the PaidMediaPayload field.
func (o *TransactionPartnerUser) SetPaidMediaPayload(v string) {
	o.PaidMediaPayload = &v
}


// GetGift returns the Gift field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetGift() Gift {
	if o == nil || IsNil(o.Gift) {
		var ret Gift
		return ret
	}
	return *o.Gift
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetGiftOk() (*Gift, bool) {
	if o == nil || IsNil(o.Gift) {
		return nil, false
	}
	return o.Gift, true
}

// HasGift returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasGift() bool {
	if o != nil && !IsNil(o.Gift) {
		return true
	}

	return false
}

// SetGift gets a reference to the given Gift and assigns it to the Gift field.
func (o *TransactionPartnerUser) SetGift(v Gift) {
	o.Gift = &v
}


// GetPremiumSubscriptionDuration returns the PremiumSubscriptionDuration field value if set, zero value otherwise.
func (o *TransactionPartnerUser) GetPremiumSubscriptionDuration() int32 {
	if o == nil || IsNil(o.PremiumSubscriptionDuration) {
		var ret int32
		return ret
	}
	return *o.PremiumSubscriptionDuration
}

// GetPremiumSubscriptionDurationOk returns a tuple with the PremiumSubscriptionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartnerUser) GetPremiumSubscriptionDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.PremiumSubscriptionDuration) {
		return nil, false
	}
	return o.PremiumSubscriptionDuration, true
}

// HasPremiumSubscriptionDuration returns a boolean if a field has been set.
func (o *TransactionPartnerUser) HasPremiumSubscriptionDuration() bool {
	if o != nil && !IsNil(o.PremiumSubscriptionDuration) {
		return true
	}

	return false
}

// SetPremiumSubscriptionDuration gets a reference to the given int32 and assigns it to the PremiumSubscriptionDuration field.
func (o *TransactionPartnerUser) SetPremiumSubscriptionDuration(v int32) {
	o.PremiumSubscriptionDuration = &v
}


func (o TransactionPartnerUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionPartnerUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["transaction_type"] = o.TransactionType
	toSerialize["user"] = o.User
	if !IsNil(o.Affiliate) {
		toSerialize["affiliate"] = o.Affiliate
	}
	if !IsNil(o.InvoicePayload) {
		toSerialize["invoice_payload"] = o.InvoicePayload
	}
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscription_period"] = o.SubscriptionPeriod
	}
	if !IsNil(o.PaidMedia) {
		toSerialize["paid_media"] = o.PaidMedia
	}
	if !IsNil(o.PaidMediaPayload) {
		toSerialize["paid_media_payload"] = o.PaidMediaPayload
	}
	if !IsNil(o.Gift) {
		toSerialize["gift"] = o.Gift
	}
	if !IsNil(o.PremiumSubscriptionDuration) {
		toSerialize["premium_subscription_duration"] = o.PremiumSubscriptionDuration
	}
	return toSerialize, nil
}

func (o *TransactionPartnerUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"transaction_type",
		"user",
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

	varTransactionPartnerUser := _TransactionPartnerUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionPartnerUser)

	if err != nil {
		return err
	}

	*o = TransactionPartnerUser(varTransactionPartnerUser)

	return err
}

type NullableTransactionPartnerUser struct {
	value *TransactionPartnerUser
	isSet bool
}

func (v NullableTransactionPartnerUser) Get() *TransactionPartnerUser {
	return v.value
}

func (v *NullableTransactionPartnerUser) Set(val *TransactionPartnerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPartnerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPartnerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPartnerUser(val *TransactionPartnerUser) *NullableTransactionPartnerUser {
	return &NullableTransactionPartnerUser{value: val, isSet: true}
}

func (v NullableTransactionPartnerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPartnerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


