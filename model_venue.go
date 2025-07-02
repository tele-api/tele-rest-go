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

// checks if the Venue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Venue{}

// Venue This object represents a venue.
type Venue struct {
	Location Location `json:"location"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue
	FoursquareId *string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

type _Venue Venue

// NewVenue instantiates a new Venue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVenue(location Location, title string, address string) *Venue {
	this := Venue{}
	this.Location = location
	this.Title = title
	this.Address = address
	return &this
}

// NewVenueWithDefaults instantiates a new Venue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVenueWithDefaults() *Venue {
	this := Venue{}
	return &this
}

// GetLocation returns the Location field value
func (o *Venue) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Venue) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Venue) SetLocation(v Location) {
	o.Location = v
}

// GetTitle returns the Title field value
func (o *Venue) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Venue) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Venue) SetTitle(v string) {
	o.Title = v
}

// GetAddress returns the Address field value
func (o *Venue) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Venue) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Venue) SetAddress(v string) {
	o.Address = v
}

// GetFoursquareId returns the FoursquareId field value if set, zero value otherwise.
func (o *Venue) GetFoursquareId() string {
	if o == nil || IsNil(o.FoursquareId) {
		var ret string
		return ret
	}
	return *o.FoursquareId
}

// GetFoursquareIdOk returns a tuple with the FoursquareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venue) GetFoursquareIdOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareId) {
		return nil, false
	}
	return o.FoursquareId, true
}

// HasFoursquareId returns a boolean if a field has been set.
func (o *Venue) HasFoursquareId() bool {
	if o != nil && !IsNil(o.FoursquareId) {
		return true
	}

	return false
}

// SetFoursquareId gets a reference to the given string and assigns it to the FoursquareId field.
func (o *Venue) SetFoursquareId(v string) {
	o.FoursquareId = &v
}


// GetFoursquareType returns the FoursquareType field value if set, zero value otherwise.
func (o *Venue) GetFoursquareType() string {
	if o == nil || IsNil(o.FoursquareType) {
		var ret string
		return ret
	}
	return *o.FoursquareType
}

// GetFoursquareTypeOk returns a tuple with the FoursquareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venue) GetFoursquareTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareType) {
		return nil, false
	}
	return o.FoursquareType, true
}

// HasFoursquareType returns a boolean if a field has been set.
func (o *Venue) HasFoursquareType() bool {
	if o != nil && !IsNil(o.FoursquareType) {
		return true
	}

	return false
}

// SetFoursquareType gets a reference to the given string and assigns it to the FoursquareType field.
func (o *Venue) SetFoursquareType(v string) {
	o.FoursquareType = &v
}


// GetGooglePlaceId returns the GooglePlaceId field value if set, zero value otherwise.
func (o *Venue) GetGooglePlaceId() string {
	if o == nil || IsNil(o.GooglePlaceId) {
		var ret string
		return ret
	}
	return *o.GooglePlaceId
}

// GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venue) GetGooglePlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceId) {
		return nil, false
	}
	return o.GooglePlaceId, true
}

// HasGooglePlaceId returns a boolean if a field has been set.
func (o *Venue) HasGooglePlaceId() bool {
	if o != nil && !IsNil(o.GooglePlaceId) {
		return true
	}

	return false
}

// SetGooglePlaceId gets a reference to the given string and assigns it to the GooglePlaceId field.
func (o *Venue) SetGooglePlaceId(v string) {
	o.GooglePlaceId = &v
}


// GetGooglePlaceType returns the GooglePlaceType field value if set, zero value otherwise.
func (o *Venue) GetGooglePlaceType() string {
	if o == nil || IsNil(o.GooglePlaceType) {
		var ret string
		return ret
	}
	return *o.GooglePlaceType
}

// GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venue) GetGooglePlaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceType) {
		return nil, false
	}
	return o.GooglePlaceType, true
}

// HasGooglePlaceType returns a boolean if a field has been set.
func (o *Venue) HasGooglePlaceType() bool {
	if o != nil && !IsNil(o.GooglePlaceType) {
		return true
	}

	return false
}

// SetGooglePlaceType gets a reference to the given string and assigns it to the GooglePlaceType field.
func (o *Venue) SetGooglePlaceType(v string) {
	o.GooglePlaceType = &v
}


func (o Venue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Venue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
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

func (o *Venue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
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

	varVenue := _Venue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVenue)

	if err != nil {
		return err
	}

	*o = Venue(varVenue)

	return err
}

type NullableVenue struct {
	value *Venue
	isSet bool
}

func (v NullableVenue) Get() *Venue {
	return v.value
}

func (v *NullableVenue) Set(val *Venue) {
	v.value = val
	v.isSet = true
}

func (v NullableVenue) IsSet() bool {
	return v.isSet
}

func (v *NullableVenue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVenue(val *Venue) *NullableVenue {
	return &NullableVenue{value: val, isSet: true}
}

func (v NullableVenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVenue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


