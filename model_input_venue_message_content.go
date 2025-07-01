/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the InputVenueMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputVenueMessageContent{}

// InputVenueMessageContent Represents the [content](https://core.telegram.org/bots/api/#inputmessagecontent) of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the venue in degrees
	Longitude float32 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue, if known
	FoursquareId *string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

type _InputVenueMessageContent InputVenueMessageContent

// NewInputVenueMessageContent instantiates a new InputVenueMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVenueMessageContent(latitude float32, longitude float32, title string, address string) *InputVenueMessageContent {
	this := InputVenueMessageContent{}
	this.Latitude = latitude
	this.Longitude = longitude
	this.Title = title
	this.Address = address
	return &this
}

// NewInputVenueMessageContentWithDefaults instantiates a new InputVenueMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVenueMessageContentWithDefaults() *InputVenueMessageContent {
	this := InputVenueMessageContent{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *InputVenueMessageContent) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *InputVenueMessageContent) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *InputVenueMessageContent) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *InputVenueMessageContent) SetLongitude(v float32) {
	o.Longitude = v
}

// GetTitle returns the Title field value
func (o *InputVenueMessageContent) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InputVenueMessageContent) SetTitle(v string) {
	o.Title = v
}

// GetAddress returns the Address field value
func (o *InputVenueMessageContent) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InputVenueMessageContent) SetAddress(v string) {
	o.Address = v
}

// GetFoursquareId returns the FoursquareId field value if set, zero value otherwise.
func (o *InputVenueMessageContent) GetFoursquareId() string {
	if o == nil || IsNil(o.FoursquareId) {
		var ret string
		return ret
	}
	return *o.FoursquareId
}

// GetFoursquareIdOk returns a tuple with the FoursquareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetFoursquareIdOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareId) {
		return nil, false
	}
	return o.FoursquareId, true
}

// HasFoursquareId returns a boolean if a field has been set.
func (o *InputVenueMessageContent) HasFoursquareId() bool {
	if o != nil && !IsNil(o.FoursquareId) {
		return true
	}

	return false
}

// SetFoursquareId gets a reference to the given string and assigns it to the FoursquareId field.
func (o *InputVenueMessageContent) SetFoursquareId(v string) {
	o.FoursquareId = &v
}


// GetFoursquareType returns the FoursquareType field value if set, zero value otherwise.
func (o *InputVenueMessageContent) GetFoursquareType() string {
	if o == nil || IsNil(o.FoursquareType) {
		var ret string
		return ret
	}
	return *o.FoursquareType
}

// GetFoursquareTypeOk returns a tuple with the FoursquareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetFoursquareTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareType) {
		return nil, false
	}
	return o.FoursquareType, true
}

// HasFoursquareType returns a boolean if a field has been set.
func (o *InputVenueMessageContent) HasFoursquareType() bool {
	if o != nil && !IsNil(o.FoursquareType) {
		return true
	}

	return false
}

// SetFoursquareType gets a reference to the given string and assigns it to the FoursquareType field.
func (o *InputVenueMessageContent) SetFoursquareType(v string) {
	o.FoursquareType = &v
}


// GetGooglePlaceId returns the GooglePlaceId field value if set, zero value otherwise.
func (o *InputVenueMessageContent) GetGooglePlaceId() string {
	if o == nil || IsNil(o.GooglePlaceId) {
		var ret string
		return ret
	}
	return *o.GooglePlaceId
}

// GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetGooglePlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceId) {
		return nil, false
	}
	return o.GooglePlaceId, true
}

// HasGooglePlaceId returns a boolean if a field has been set.
func (o *InputVenueMessageContent) HasGooglePlaceId() bool {
	if o != nil && !IsNil(o.GooglePlaceId) {
		return true
	}

	return false
}

// SetGooglePlaceId gets a reference to the given string and assigns it to the GooglePlaceId field.
func (o *InputVenueMessageContent) SetGooglePlaceId(v string) {
	o.GooglePlaceId = &v
}


// GetGooglePlaceType returns the GooglePlaceType field value if set, zero value otherwise.
func (o *InputVenueMessageContent) GetGooglePlaceType() string {
	if o == nil || IsNil(o.GooglePlaceType) {
		var ret string
		return ret
	}
	return *o.GooglePlaceType
}

// GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVenueMessageContent) GetGooglePlaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceType) {
		return nil, false
	}
	return o.GooglePlaceType, true
}

// HasGooglePlaceType returns a boolean if a field has been set.
func (o *InputVenueMessageContent) HasGooglePlaceType() bool {
	if o != nil && !IsNil(o.GooglePlaceType) {
		return true
	}

	return false
}

// SetGooglePlaceType gets a reference to the given string and assigns it to the GooglePlaceType field.
func (o *InputVenueMessageContent) SetGooglePlaceType(v string) {
	o.GooglePlaceType = &v
}


func (o InputVenueMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputVenueMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["title"] = o.Title
	toSerialize["address"] = o.Address
	if !IsNil(o.FoursquareId) {
		toSerialize["foursquare_id"] = o.FoursquareId
	}
	if !IsNil(o.FoursquareType) {
		toSerialize["foursquare_type"] = o.FoursquareType
	}
	if !IsNil(o.GooglePlaceId) {
		toSerialize["google_place_id"] = o.GooglePlaceId
	}
	if !IsNil(o.GooglePlaceType) {
		toSerialize["google_place_type"] = o.GooglePlaceType
	}
	return toSerialize, nil
}

func (o *InputVenueMessageContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"latitude",
		"longitude",
		"title",
		"address",
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

	varInputVenueMessageContent := _InputVenueMessageContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputVenueMessageContent)

	if err != nil {
		return err
	}

	*o = InputVenueMessageContent(varInputVenueMessageContent)

	return err
}

type NullableInputVenueMessageContent struct {
	value *InputVenueMessageContent
	isSet bool
}

func (v NullableInputVenueMessageContent) Get() *InputVenueMessageContent {
	return v.value
}

func (v *NullableInputVenueMessageContent) Set(val *InputVenueMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVenueMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVenueMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVenueMessageContent(val *InputVenueMessageContent) *NullableInputVenueMessageContent {
	return &NullableInputVenueMessageContent{value: val, isSet: true}
}

func (v NullableInputVenueMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVenueMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


