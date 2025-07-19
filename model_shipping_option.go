/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the ShippingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingOption{}

// ShippingOption This object represents one shipping option.
type ShippingOption struct {
	// Shipping option identifier
	Id string `json:"id"`
	// Option title
	Title string `json:"title"`
	// List of price portions
	Prices []LabeledPrice `json:"prices"`
}

type _ShippingOption ShippingOption

// NewShippingOption instantiates a new ShippingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingOption(id string, title string, prices []LabeledPrice) *ShippingOption {
	this := ShippingOption{}
	this.Id = id
	this.Title = title
	this.Prices = prices
	return &this
}

// NewShippingOptionWithDefaults instantiates a new ShippingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingOptionWithDefaults() *ShippingOption {
	this := ShippingOption{}
	return &this
}

// GetId returns the Id field value
func (o *ShippingOption) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShippingOption) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *ShippingOption) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ShippingOption) SetTitle(v string) {
	o.Title = v
}

// GetPrices returns the Prices field value
func (o *ShippingOption) GetPrices() []LabeledPrice {
	if o == nil {
		var ret []LabeledPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetPricesOk() ([]LabeledPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *ShippingOption) SetPrices(v []LabeledPrice) {
	o.Prices = v
}

func (o ShippingOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["prices"] = o.Prices
	return toSerialize, nil
}

func (o *ShippingOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"prices",
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

	varShippingOption := _ShippingOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShippingOption)

	if err != nil {
		return err
	}

	*o = ShippingOption(varShippingOption)

	return err
}

type NullableShippingOption struct {
	value *ShippingOption
	isSet bool
}

func (v NullableShippingOption) Get() *ShippingOption {
	return v.value
}

func (v *NullableShippingOption) Set(val *ShippingOption) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingOption) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingOption(val *ShippingOption) *NullableShippingOption {
	return &NullableShippingOption{value: val, isSet: true}
}

func (v NullableShippingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


