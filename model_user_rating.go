/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.3.0
 *    * - **Modified**: 2026-01-01T02:06:16.405655294Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// checks if the UserRating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRating{}

// UserRating This object describes the rating of a user based on their Telegram Star spendings.
type UserRating struct {
	// Current level of the user, indicating their reliability when purchasing digital goods and services. A higher level suggests a more trustworthy customer; a negative level is likely reason for concern.
	Level int32 `json:"level"`
	// Numerical value of the user's rating; the higher the rating, the better
	Rating int32 `json:"rating"`
	// The rating value required to get the current level
	CurrentLevelRating int32 `json:"current_level_rating"`
	// *Optional*. The rating value required to get to the next level; omitted if the maximum level was reached
	NextLevelRating *int32 `json:"next_level_rating,omitempty"`
}

type _UserRating UserRating

// NewUserRating instantiates a new UserRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRating(level int32, rating int32, currentLevelRating int32) *UserRating {
	this := UserRating{}
	this.Level = level
	this.Rating = rating
	this.CurrentLevelRating = currentLevelRating
	return &this
}

// NewUserRatingWithDefaults instantiates a new UserRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRatingWithDefaults() *UserRating {
	this := UserRating{}
	return &this
}

// GetLevel returns the Level field value
func (o *UserRating) GetLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *UserRating) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *UserRating) SetLevel(v int32) {
	o.Level = v
}

// GetRating returns the Rating field value
func (o *UserRating) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *UserRating) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *UserRating) SetRating(v int32) {
	o.Rating = v
}

// GetCurrentLevelRating returns the CurrentLevelRating field value
func (o *UserRating) GetCurrentLevelRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentLevelRating
}

// GetCurrentLevelRatingOk returns a tuple with the CurrentLevelRating field value
// and a boolean to check if the value has been set.
func (o *UserRating) GetCurrentLevelRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLevelRating, true
}

// SetCurrentLevelRating sets field value
func (o *UserRating) SetCurrentLevelRating(v int32) {
	o.CurrentLevelRating = v
}

// GetNextLevelRating returns the NextLevelRating field value if set, zero value otherwise.
func (o *UserRating) GetNextLevelRating() int32 {
	if o == nil || IsNil(o.NextLevelRating) {
		var ret int32
		return ret
	}
	return *o.NextLevelRating
}

// GetNextLevelRatingOk returns a tuple with the NextLevelRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRating) GetNextLevelRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.NextLevelRating) {
		return nil, false
	}
	return o.NextLevelRating, true
}

// HasNextLevelRating returns a boolean if a field has been set.
func (o *UserRating) HasNextLevelRating() bool {
	if o != nil && !IsNil(o.NextLevelRating) {
		return true
	}

	return false
}

// SetNextLevelRating gets a reference to the given int32 and assigns it to the NextLevelRating field.
func (o *UserRating) SetNextLevelRating(v int32) {
	o.NextLevelRating = &v
}


func (o UserRating) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["rating"] = o.Rating
	toSerialize["current_level_rating"] = o.CurrentLevelRating
	if !IsNil(o.NextLevelRating) {
		toSerialize["next_level_rating"] = o.NextLevelRating
	}
	return toSerialize, nil
}

func (o *UserRating) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
		"rating",
		"current_level_rating",
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

	varUserRating := _UserRating{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRating)

	if err != nil {
		return err
	}

	*o = UserRating(varUserRating)

	return err
}

type NullableUserRating struct {
	value *UserRating
	isSet bool
}

func (v NullableUserRating) Get() *UserRating {
	return v.value
}

func (v *NullableUserRating) Set(val *UserRating) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRating) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRating(val *UserRating) *NullableUserRating {
	return &NullableUserRating{value: val, isSet: true}
}

func (v NullableUserRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

