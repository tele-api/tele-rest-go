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

// checks if the Birthdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Birthdate{}

// Birthdate Describes the birthdate of a user.
type Birthdate struct {
	// Day of the user's birth; 1-31
	Day int32 `json:"day"`
	// Month of the user's birth; 1-12
	Month int32 `json:"month"`
	// *Optional*. Year of the user's birth
	Year *int32 `json:"year,omitempty"`
}

type _Birthdate Birthdate

// NewBirthdate instantiates a new Birthdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBirthdate(day int32, month int32) *Birthdate {
	this := Birthdate{}
	this.Day = day
	this.Month = month
	return &this
}

// NewBirthdateWithDefaults instantiates a new Birthdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBirthdateWithDefaults() *Birthdate {
	this := Birthdate{}
	return &this
}

// GetDay returns the Day field value
func (o *Birthdate) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *Birthdate) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *Birthdate) SetDay(v int32) {
	o.Day = v
}

// GetMonth returns the Month field value
func (o *Birthdate) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *Birthdate) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *Birthdate) SetMonth(v int32) {
	o.Month = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Birthdate) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Birthdate) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Birthdate) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *Birthdate) SetYear(v int32) {
	o.Year = &v
}


func (o Birthdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Birthdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["day"] = o.Day
	toSerialize["month"] = o.Month
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

func (o *Birthdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"day",
		"month",
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

	varBirthdate := _Birthdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBirthdate)

	if err != nil {
		return err
	}

	*o = Birthdate(varBirthdate)

	return err
}

type NullableBirthdate struct {
	value *Birthdate
	isSet bool
}

func (v NullableBirthdate) Get() *Birthdate {
	return v.value
}

func (v *NullableBirthdate) Set(val *Birthdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBirthdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBirthdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBirthdate(val *Birthdate) *NullableBirthdate {
	return &NullableBirthdate{value: val, isSet: true}
}

func (v NullableBirthdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBirthdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


