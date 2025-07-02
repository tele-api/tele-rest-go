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

// checks if the LocationAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationAddress{}

// LocationAddress Describes the physical address of a location.
type LocationAddress struct {
	// The two-letter ISO 3166-1 alpha-2 country code of the country where the location is located
	CountryCode string `json:"country_code"`
	// *Optional*. State of the location
	State *string `json:"state,omitempty"`
	// *Optional*. City of the location
	City *string `json:"city,omitempty"`
	// *Optional*. Street address of the location
	Street *string `json:"street,omitempty"`
}

type _LocationAddress LocationAddress

// NewLocationAddress instantiates a new LocationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAddress(countryCode string) *LocationAddress {
	this := LocationAddress{}
	this.CountryCode = countryCode
	return &this
}

// NewLocationAddressWithDefaults instantiates a new LocationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAddressWithDefaults() *LocationAddress {
	this := LocationAddress{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *LocationAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *LocationAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *LocationAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LocationAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LocationAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LocationAddress) SetState(v string) {
	o.State = &v
}


// GetCity returns the City field value if set, zero value otherwise.
func (o *LocationAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *LocationAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *LocationAddress) SetCity(v string) {
	o.City = &v
}


// GetStreet returns the Street field value if set, zero value otherwise.
func (o *LocationAddress) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAddress) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *LocationAddress) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *LocationAddress) SetStreet(v string) {
	o.Street = &v
}


func (o LocationAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country_code"] = o.CountryCode
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	return toSerialize, nil
}

func (o *LocationAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country_code",
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

	varLocationAddress := _LocationAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocationAddress)

	if err != nil {
		return err
	}

	*o = LocationAddress(varLocationAddress)

	return err
}

type NullableLocationAddress struct {
	value *LocationAddress
	isSet bool
}

func (v NullableLocationAddress) Get() *LocationAddress {
	return v.value
}

func (v *NullableLocationAddress) Set(val *LocationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAddress(val *LocationAddress) *NullableLocationAddress {
	return &NullableLocationAddress{value: val, isSet: true}
}

func (v NullableLocationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


