/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the InlineQueryResultVenue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultVenue{}

// InlineQueryResultVenue Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	// Type of the result, must be *venue*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Latitude of the venue location in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the venue location in degrees
	Longitude float32 `json:"longitude"`
	// Title of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue if known
	FoursquareId *string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`
}

type _InlineQueryResultVenue InlineQueryResultVenue

// NewInlineQueryResultVenue instantiates a new InlineQueryResultVenue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultVenue(type_ string, id string, latitude float32, longitude float32, title string, address string) *InlineQueryResultVenue {
	this := InlineQueryResultVenue{}
	this.Type = type_
	this.Id = id
	this.Latitude = latitude
	this.Longitude = longitude
	this.Title = title
	this.Address = address
	return &this
}

// NewInlineQueryResultVenueWithDefaults instantiates a new InlineQueryResultVenue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultVenueWithDefaults() *InlineQueryResultVenue {
	this := InlineQueryResultVenue{}
	var type_ string = "venue"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultVenue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultVenue) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultVenue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultVenue) SetId(v string) {
	o.Id = v
}

// GetLatitude returns the Latitude field value
func (o *InlineQueryResultVenue) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *InlineQueryResultVenue) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *InlineQueryResultVenue) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *InlineQueryResultVenue) SetLongitude(v float32) {
	o.Longitude = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultVenue) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultVenue) SetTitle(v string) {
	o.Title = v
}

// GetAddress returns the Address field value
func (o *InlineQueryResultVenue) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineQueryResultVenue) SetAddress(v string) {
	o.Address = v
}

// GetFoursquareId returns the FoursquareId field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetFoursquareId() string {
	if o == nil || IsNil(o.FoursquareId) {
		var ret string
		return ret
	}
	return *o.FoursquareId
}

// GetFoursquareIdOk returns a tuple with the FoursquareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetFoursquareIdOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareId) {
		return nil, false
	}
	return o.FoursquareId, true
}

// HasFoursquareId returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasFoursquareId() bool {
	if o != nil && !IsNil(o.FoursquareId) {
		return true
	}

	return false
}

// SetFoursquareId gets a reference to the given string and assigns it to the FoursquareId field.
func (o *InlineQueryResultVenue) SetFoursquareId(v string) {
	o.FoursquareId = &v
}


// GetFoursquareType returns the FoursquareType field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetFoursquareType() string {
	if o == nil || IsNil(o.FoursquareType) {
		var ret string
		return ret
	}
	return *o.FoursquareType
}

// GetFoursquareTypeOk returns a tuple with the FoursquareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetFoursquareTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareType) {
		return nil, false
	}
	return o.FoursquareType, true
}

// HasFoursquareType returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasFoursquareType() bool {
	if o != nil && !IsNil(o.FoursquareType) {
		return true
	}

	return false
}

// SetFoursquareType gets a reference to the given string and assigns it to the FoursquareType field.
func (o *InlineQueryResultVenue) SetFoursquareType(v string) {
	o.FoursquareType = &v
}


// GetGooglePlaceId returns the GooglePlaceId field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetGooglePlaceId() string {
	if o == nil || IsNil(o.GooglePlaceId) {
		var ret string
		return ret
	}
	return *o.GooglePlaceId
}

// GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetGooglePlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceId) {
		return nil, false
	}
	return o.GooglePlaceId, true
}

// HasGooglePlaceId returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasGooglePlaceId() bool {
	if o != nil && !IsNil(o.GooglePlaceId) {
		return true
	}

	return false
}

// SetGooglePlaceId gets a reference to the given string and assigns it to the GooglePlaceId field.
func (o *InlineQueryResultVenue) SetGooglePlaceId(v string) {
	o.GooglePlaceId = &v
}


// GetGooglePlaceType returns the GooglePlaceType field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetGooglePlaceType() string {
	if o == nil || IsNil(o.GooglePlaceType) {
		var ret string
		return ret
	}
	return *o.GooglePlaceType
}

// GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetGooglePlaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceType) {
		return nil, false
	}
	return o.GooglePlaceType, true
}

// HasGooglePlaceType returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasGooglePlaceType() bool {
	if o != nil && !IsNil(o.GooglePlaceType) {
		return true
	}

	return false
}

// SetGooglePlaceType gets a reference to the given string and assigns it to the GooglePlaceType field.
func (o *InlineQueryResultVenue) SetGooglePlaceType(v string) {
	o.GooglePlaceType = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultVenue) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultVenue) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *InlineQueryResultVenue) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}


// GetThumbnailWidth returns the ThumbnailWidth field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetThumbnailWidth() int32 {
	if o == nil || IsNil(o.ThumbnailWidth) {
		var ret int32
		return ret
	}
	return *o.ThumbnailWidth
}

// GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetThumbnailWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailWidth) {
		return nil, false
	}
	return o.ThumbnailWidth, true
}

// HasThumbnailWidth returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasThumbnailWidth() bool {
	if o != nil && !IsNil(o.ThumbnailWidth) {
		return true
	}

	return false
}

// SetThumbnailWidth gets a reference to the given int32 and assigns it to the ThumbnailWidth field.
func (o *InlineQueryResultVenue) SetThumbnailWidth(v int32) {
	o.ThumbnailWidth = &v
}


// GetThumbnailHeight returns the ThumbnailHeight field value if set, zero value otherwise.
func (o *InlineQueryResultVenue) GetThumbnailHeight() int32 {
	if o == nil || IsNil(o.ThumbnailHeight) {
		var ret int32
		return ret
	}
	return *o.ThumbnailHeight
}

// GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultVenue) GetThumbnailHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailHeight) {
		return nil, false
	}
	return o.ThumbnailHeight, true
}

// HasThumbnailHeight returns a boolean if a field has been set.
func (o *InlineQueryResultVenue) HasThumbnailHeight() bool {
	if o != nil && !IsNil(o.ThumbnailHeight) {
		return true
	}

	return false
}

// SetThumbnailHeight gets a reference to the given int32 and assigns it to the ThumbnailHeight field.
func (o *InlineQueryResultVenue) SetThumbnailHeight(v int32) {
	o.ThumbnailHeight = &v
}


func (o InlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultVenue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
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
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	if !IsNil(o.InputMessageContent) {
		toSerialize["input_message_content"] = o.InputMessageContent
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if !IsNil(o.ThumbnailWidth) {
		toSerialize["thumbnail_width"] = o.ThumbnailWidth
	}
	if !IsNil(o.ThumbnailHeight) {
		toSerialize["thumbnail_height"] = o.ThumbnailHeight
	}
	return toSerialize, nil
}

func (o *InlineQueryResultVenue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varInlineQueryResultVenue := _InlineQueryResultVenue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultVenue)

	if err != nil {
		return err
	}

	*o = InlineQueryResultVenue(varInlineQueryResultVenue)

	return err
}

type NullableInlineQueryResultVenue struct {
	value *InlineQueryResultVenue
	isSet bool
}

func (v NullableInlineQueryResultVenue) Get() *InlineQueryResultVenue {
	return v.value
}

func (v *NullableInlineQueryResultVenue) Set(val *InlineQueryResultVenue) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultVenue) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultVenue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultVenue(val *InlineQueryResultVenue) *NullableInlineQueryResultVenue {
	return &NullableInlineQueryResultVenue{value: val, isSet: true}
}

func (v NullableInlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultVenue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


