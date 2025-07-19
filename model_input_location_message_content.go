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

// checks if the InputLocationMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputLocationMessageContent{}

// InputLocationMessageContent Represents the [content](https://core.telegram.org/bots/api/#inputmessagecontent) of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float32 `json:"longitude"`
	// *Optional*. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	// *Optional*. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int32 `json:"live_period,omitempty"`
	// *Optional*. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int32 `json:"heading,omitempty"`
	// *Optional*. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int32 `json:"proximity_alert_radius,omitempty"`
}

type _InputLocationMessageContent InputLocationMessageContent

// NewInputLocationMessageContent instantiates a new InputLocationMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputLocationMessageContent(latitude float32, longitude float32) *InputLocationMessageContent {
	this := InputLocationMessageContent{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewInputLocationMessageContentWithDefaults instantiates a new InputLocationMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputLocationMessageContentWithDefaults() *InputLocationMessageContent {
	this := InputLocationMessageContent{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *InputLocationMessageContent) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *InputLocationMessageContent) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *InputLocationMessageContent) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *InputLocationMessageContent) SetLongitude(v float32) {
	o.Longitude = v
}

// GetHorizontalAccuracy returns the HorizontalAccuracy field value if set, zero value otherwise.
func (o *InputLocationMessageContent) GetHorizontalAccuracy() float32 {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		var ret float32
		return ret
	}
	return *o.HorizontalAccuracy
}

// GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetHorizontalAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		return nil, false
	}
	return o.HorizontalAccuracy, true
}

// HasHorizontalAccuracy returns a boolean if a field has been set.
func (o *InputLocationMessageContent) HasHorizontalAccuracy() bool {
	if o != nil && !IsNil(o.HorizontalAccuracy) {
		return true
	}

	return false
}

// SetHorizontalAccuracy gets a reference to the given float32 and assigns it to the HorizontalAccuracy field.
func (o *InputLocationMessageContent) SetHorizontalAccuracy(v float32) {
	o.HorizontalAccuracy = &v
}


// GetLivePeriod returns the LivePeriod field value if set, zero value otherwise.
func (o *InputLocationMessageContent) GetLivePeriod() int32 {
	if o == nil || IsNil(o.LivePeriod) {
		var ret int32
		return ret
	}
	return *o.LivePeriod
}

// GetLivePeriodOk returns a tuple with the LivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetLivePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.LivePeriod) {
		return nil, false
	}
	return o.LivePeriod, true
}

// HasLivePeriod returns a boolean if a field has been set.
func (o *InputLocationMessageContent) HasLivePeriod() bool {
	if o != nil && !IsNil(o.LivePeriod) {
		return true
	}

	return false
}

// SetLivePeriod gets a reference to the given int32 and assigns it to the LivePeriod field.
func (o *InputLocationMessageContent) SetLivePeriod(v int32) {
	o.LivePeriod = &v
}


// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *InputLocationMessageContent) GetHeading() int32 {
	if o == nil || IsNil(o.Heading) {
		var ret int32
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetHeadingOk() (*int32, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *InputLocationMessageContent) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given int32 and assigns it to the Heading field.
func (o *InputLocationMessageContent) SetHeading(v int32) {
	o.Heading = &v
}


// GetProximityAlertRadius returns the ProximityAlertRadius field value if set, zero value otherwise.
func (o *InputLocationMessageContent) GetProximityAlertRadius() int32 {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		var ret int32
		return ret
	}
	return *o.ProximityAlertRadius
}

// GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputLocationMessageContent) GetProximityAlertRadiusOk() (*int32, bool) {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		return nil, false
	}
	return o.ProximityAlertRadius, true
}

// HasProximityAlertRadius returns a boolean if a field has been set.
func (o *InputLocationMessageContent) HasProximityAlertRadius() bool {
	if o != nil && !IsNil(o.ProximityAlertRadius) {
		return true
	}

	return false
}

// SetProximityAlertRadius gets a reference to the given int32 and assigns it to the ProximityAlertRadius field.
func (o *InputLocationMessageContent) SetProximityAlertRadius(v int32) {
	o.ProximityAlertRadius = &v
}


func (o InputLocationMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputLocationMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	if !IsNil(o.HorizontalAccuracy) {
		toSerialize["horizontal_accuracy"] = o.HorizontalAccuracy
	}
	if !IsNil(o.LivePeriod) {
		toSerialize["live_period"] = o.LivePeriod
	}
	if !IsNil(o.Heading) {
		toSerialize["heading"] = o.Heading
	}
	if !IsNil(o.ProximityAlertRadius) {
		toSerialize["proximity_alert_radius"] = o.ProximityAlertRadius
	}
	return toSerialize, nil
}

func (o *InputLocationMessageContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varInputLocationMessageContent := _InputLocationMessageContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputLocationMessageContent)

	if err != nil {
		return err
	}

	*o = InputLocationMessageContent(varInputLocationMessageContent)

	return err
}

type NullableInputLocationMessageContent struct {
	value *InputLocationMessageContent
	isSet bool
}

func (v NullableInputLocationMessageContent) Get() *InputLocationMessageContent {
	return v.value
}

func (v *NullableInputLocationMessageContent) Set(val *InputLocationMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputLocationMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputLocationMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputLocationMessageContent(val *InputLocationMessageContent) *NullableInputLocationMessageContent {
	return &NullableInputLocationMessageContent{value: val, isSet: true}
}

func (v NullableInputLocationMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputLocationMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


