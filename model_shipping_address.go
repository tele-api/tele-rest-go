/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the ShippingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingAddress{}

// ShippingAddress This object represents a shipping address.
type ShippingAddress struct {
	// Two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code
	CountryCode string `json:"country_code"`
	// State, if applicable
	State string `json:"state"`
	// City
	City string `json:"city"`
	// First line for the address
	StreetLine1 string `json:"street_line1"`
	// Second line for the address
	StreetLine2 string `json:"street_line2"`
	// Address post code
	PostCode string `json:"post_code"`
}

type _ShippingAddress ShippingAddress

// NewShippingAddress instantiates a new ShippingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingAddress(countryCode string, state string, city string, streetLine1 string, streetLine2 string, postCode string) *ShippingAddress {
	this := ShippingAddress{}
	this.CountryCode = countryCode
	this.State = state
	this.City = city
	this.StreetLine1 = streetLine1
	this.StreetLine2 = streetLine2
	this.PostCode = postCode
	return &this
}

// NewShippingAddressWithDefaults instantiates a new ShippingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingAddressWithDefaults() *ShippingAddress {
	this := ShippingAddress{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *ShippingAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *ShippingAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetState returns the State field value
func (o *ShippingAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ShippingAddress) SetState(v string) {
	o.State = v
}

// GetCity returns the City field value
func (o *ShippingAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ShippingAddress) SetCity(v string) {
	o.City = v
}

// GetStreetLine1 returns the StreetLine1 field value
func (o *ShippingAddress) GetStreetLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreetLine1
}

// GetStreetLine1Ok returns a tuple with the StreetLine1 field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetStreetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreetLine1, true
}

// SetStreetLine1 sets field value
func (o *ShippingAddress) SetStreetLine1(v string) {
	o.StreetLine1 = v
}

// GetStreetLine2 returns the StreetLine2 field value
func (o *ShippingAddress) GetStreetLine2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreetLine2
}

// GetStreetLine2Ok returns a tuple with the StreetLine2 field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetStreetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreetLine2, true
}

// SetStreetLine2 sets field value
func (o *ShippingAddress) SetStreetLine2(v string) {
	o.StreetLine2 = v
}

// GetPostCode returns the PostCode field value
func (o *ShippingAddress) GetPostCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
func (o *ShippingAddress) GetPostCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostCode, true
}

// SetPostCode sets field value
func (o *ShippingAddress) SetPostCode(v string) {
	o.PostCode = v
}

func (o ShippingAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country_code"] = o.CountryCode
	toSerialize["state"] = o.State
	toSerialize["city"] = o.City
	toSerialize["street_line1"] = o.StreetLine1
	toSerialize["street_line2"] = o.StreetLine2
	toSerialize["post_code"] = o.PostCode
	return toSerialize, nil
}

func (o *ShippingAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country_code",
		"state",
		"city",
		"street_line1",
		"street_line2",
		"post_code",
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

	varShippingAddress := _ShippingAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShippingAddress)

	if err != nil {
		return err
	}

	*o = ShippingAddress(varShippingAddress)

	return err
}

type NullableShippingAddress struct {
	value *ShippingAddress
	isSet bool
}

func (v NullableShippingAddress) Get() *ShippingAddress {
	return v.value
}

func (v *NullableShippingAddress) Set(val *ShippingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingAddress(val *ShippingAddress) *NullableShippingAddress {
	return &NullableShippingAddress{value: val, isSet: true}
}

func (v NullableShippingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


