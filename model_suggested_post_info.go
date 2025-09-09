/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.2.0
 *    * - **Modified**: 2025-09-09T23:46:30.762261440Z[Etc/UTC]
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

// checks if the SuggestedPostInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedPostInfo{}

// SuggestedPostInfo Contains information about a suggested post.
type SuggestedPostInfo struct {
	// State of the suggested post. Currently, it can be one of “pending”, “approved”, “declined”.
	State string `json:"state"`
	Price *SuggestedPostPrice `json:"price,omitempty"`
	// *Optional*. Proposed send date of the post. If the field is omitted, then the post can be published at any time within 30 days at the sole discretion of the user or administrator who approves it.
	SendDate *int32 `json:"send_date,omitempty"`
}

type _SuggestedPostInfo SuggestedPostInfo

// NewSuggestedPostInfo instantiates a new SuggestedPostInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedPostInfo(state string) *SuggestedPostInfo {
	this := SuggestedPostInfo{}
	this.State = state
	return &this
}

// NewSuggestedPostInfoWithDefaults instantiates a new SuggestedPostInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedPostInfoWithDefaults() *SuggestedPostInfo {
	this := SuggestedPostInfo{}
	return &this
}

// GetState returns the State field value
func (o *SuggestedPostInfo) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SuggestedPostInfo) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SuggestedPostInfo) SetState(v string) {
	o.State = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SuggestedPostInfo) GetPrice() SuggestedPostPrice {
	if o == nil || IsNil(o.Price) {
		var ret SuggestedPostPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostInfo) GetPriceOk() (*SuggestedPostPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SuggestedPostInfo) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given SuggestedPostPrice and assigns it to the Price field.
func (o *SuggestedPostInfo) SetPrice(v SuggestedPostPrice) {
	o.Price = &v
}


// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *SuggestedPostInfo) GetSendDate() int32 {
	if o == nil || IsNil(o.SendDate) {
		var ret int32
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedPostInfo) GetSendDateOk() (*int32, bool) {
	if o == nil || IsNil(o.SendDate) {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *SuggestedPostInfo) HasSendDate() bool {
	if o != nil && !IsNil(o.SendDate) {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given int32 and assigns it to the SendDate field.
func (o *SuggestedPostInfo) SetSendDate(v int32) {
	o.SendDate = &v
}


func (o SuggestedPostInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedPostInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.SendDate) {
		toSerialize["send_date"] = o.SendDate
	}
	return toSerialize, nil
}

func (o *SuggestedPostInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varSuggestedPostInfo := _SuggestedPostInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestedPostInfo)

	if err != nil {
		return err
	}

	*o = SuggestedPostInfo(varSuggestedPostInfo)

	return err
}

type NullableSuggestedPostInfo struct {
	value *SuggestedPostInfo
	isSet bool
}

func (v NullableSuggestedPostInfo) Get() *SuggestedPostInfo {
	return v.value
}

func (v *NullableSuggestedPostInfo) Set(val *SuggestedPostInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedPostInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedPostInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedPostInfo(val *SuggestedPostInfo) *NullableSuggestedPostInfo {
	return &NullableSuggestedPostInfo{value: val, isSet: true}
}

func (v NullableSuggestedPostInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedPostInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


