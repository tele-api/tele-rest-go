/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the StoryAreaTypeLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryAreaTypeLocation{}

// StoryAreaTypeLocation Describes a story area pointing to a location. Currently, a story can have up to 10 location areas.
type StoryAreaTypeLocation struct {
	// Type of the area, always “location”
	Type string `json:"type"`
	// Location latitude in degrees
	Latitude float32 `json:"latitude"`
	// Location longitude in degrees
	Longitude float32 `json:"longitude"`
	Address *LocationAddress `json:"address,omitempty"`
}

type _StoryAreaTypeLocation StoryAreaTypeLocation

// NewStoryAreaTypeLocation instantiates a new StoryAreaTypeLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryAreaTypeLocation(type_ string, latitude float32, longitude float32) *StoryAreaTypeLocation {
	this := StoryAreaTypeLocation{}
	this.Type = type_
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewStoryAreaTypeLocationWithDefaults instantiates a new StoryAreaTypeLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryAreaTypeLocationWithDefaults() *StoryAreaTypeLocation {
	this := StoryAreaTypeLocation{}
	var type_ string = "location"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StoryAreaTypeLocation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLocation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StoryAreaTypeLocation) SetType(v string) {
	o.Type = v
}

// GetLatitude returns the Latitude field value
func (o *StoryAreaTypeLocation) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLocation) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *StoryAreaTypeLocation) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *StoryAreaTypeLocation) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLocation) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *StoryAreaTypeLocation) SetLongitude(v float32) {
	o.Longitude = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StoryAreaTypeLocation) GetAddress() LocationAddress {
	if o == nil || IsNil(o.Address) {
		var ret LocationAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLocation) GetAddressOk() (*LocationAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StoryAreaTypeLocation) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given LocationAddress and assigns it to the Address field.
func (o *StoryAreaTypeLocation) SetAddress(v LocationAddress) {
	o.Address = &v
}


func (o StoryAreaTypeLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryAreaTypeLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

func (o *StoryAreaTypeLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"latitude",
		"longitude",
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

	varStoryAreaTypeLocation := _StoryAreaTypeLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStoryAreaTypeLocation)

	if err != nil {
		return err
	}

	*o = StoryAreaTypeLocation(varStoryAreaTypeLocation)

	return err
}

type NullableStoryAreaTypeLocation struct {
	value *StoryAreaTypeLocation
	isSet bool
}

func (v NullableStoryAreaTypeLocation) Get() *StoryAreaTypeLocation {
	return v.value
}

func (v *NullableStoryAreaTypeLocation) Set(val *StoryAreaTypeLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaTypeLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaTypeLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaTypeLocation(val *StoryAreaTypeLocation) *NullableStoryAreaTypeLocation {
	return &NullableStoryAreaTypeLocation{value: val, isSet: true}
}

func (v NullableStoryAreaTypeLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaTypeLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


