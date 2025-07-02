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

// checks if the InlineQueryResultLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultLocation{}

// InlineQueryResultLocation Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use *input\\_message\\_content* to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	// Type of the result, must be *location*
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Location latitude in degrees
	Latitude float32 `json:"latitude"`
	// Location longitude in degrees
	Longitude float32 `json:"longitude"`
	// Location title
	Title string `json:"title"`
	// *Optional*. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	// *Optional*. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int32 `json:"live_period,omitempty"`
	// *Optional*. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int32 `json:"heading,omitempty"`
	// *Optional*. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int32 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`
}

type _InlineQueryResultLocation InlineQueryResultLocation

// NewInlineQueryResultLocation instantiates a new InlineQueryResultLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultLocation(type_ string, id string, latitude float32, longitude float32, title string) *InlineQueryResultLocation {
	this := InlineQueryResultLocation{}
	this.Type = type_
	this.Id = id
	this.Latitude = latitude
	this.Longitude = longitude
	this.Title = title
	return &this
}

// NewInlineQueryResultLocationWithDefaults instantiates a new InlineQueryResultLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultLocationWithDefaults() *InlineQueryResultLocation {
	this := InlineQueryResultLocation{}
	var type_ string = "location"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InlineQueryResultLocation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineQueryResultLocation) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InlineQueryResultLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQueryResultLocation) SetId(v string) {
	o.Id = v
}

// GetLatitude returns the Latitude field value
func (o *InlineQueryResultLocation) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *InlineQueryResultLocation) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *InlineQueryResultLocation) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *InlineQueryResultLocation) SetLongitude(v float32) {
	o.Longitude = v
}

// GetTitle returns the Title field value
func (o *InlineQueryResultLocation) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineQueryResultLocation) SetTitle(v string) {
	o.Title = v
}

// GetHorizontalAccuracy returns the HorizontalAccuracy field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetHorizontalAccuracy() float32 {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		var ret float32
		return ret
	}
	return *o.HorizontalAccuracy
}

// GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetHorizontalAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		return nil, false
	}
	return o.HorizontalAccuracy, true
}

// HasHorizontalAccuracy returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasHorizontalAccuracy() bool {
	if o != nil && !IsNil(o.HorizontalAccuracy) {
		return true
	}

	return false
}

// SetHorizontalAccuracy gets a reference to the given float32 and assigns it to the HorizontalAccuracy field.
func (o *InlineQueryResultLocation) SetHorizontalAccuracy(v float32) {
	o.HorizontalAccuracy = &v
}


// GetLivePeriod returns the LivePeriod field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetLivePeriod() int32 {
	if o == nil || IsNil(o.LivePeriod) {
		var ret int32
		return ret
	}
	return *o.LivePeriod
}

// GetLivePeriodOk returns a tuple with the LivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetLivePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.LivePeriod) {
		return nil, false
	}
	return o.LivePeriod, true
}

// HasLivePeriod returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasLivePeriod() bool {
	if o != nil && !IsNil(o.LivePeriod) {
		return true
	}

	return false
}

// SetLivePeriod gets a reference to the given int32 and assigns it to the LivePeriod field.
func (o *InlineQueryResultLocation) SetLivePeriod(v int32) {
	o.LivePeriod = &v
}


// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetHeading() int32 {
	if o == nil || IsNil(o.Heading) {
		var ret int32
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetHeadingOk() (*int32, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given int32 and assigns it to the Heading field.
func (o *InlineQueryResultLocation) SetHeading(v int32) {
	o.Heading = &v
}


// GetProximityAlertRadius returns the ProximityAlertRadius field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetProximityAlertRadius() int32 {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		var ret int32
		return ret
	}
	return *o.ProximityAlertRadius
}

// GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetProximityAlertRadiusOk() (*int32, bool) {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		return nil, false
	}
	return o.ProximityAlertRadius, true
}

// HasProximityAlertRadius returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasProximityAlertRadius() bool {
	if o != nil && !IsNil(o.ProximityAlertRadius) {
		return true
	}

	return false
}

// SetProximityAlertRadius gets a reference to the given int32 and assigns it to the ProximityAlertRadius field.
func (o *InlineQueryResultLocation) SetProximityAlertRadius(v int32) {
	o.ProximityAlertRadius = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *InlineQueryResultLocation) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


// GetInputMessageContent returns the InputMessageContent field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetInputMessageContent() InputMessageContent {
	if o == nil || IsNil(o.InputMessageContent) {
		var ret InputMessageContent
		return ret
	}
	return *o.InputMessageContent
}

// GetInputMessageContentOk returns a tuple with the InputMessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetInputMessageContentOk() (*InputMessageContent, bool) {
	if o == nil || IsNil(o.InputMessageContent) {
		return nil, false
	}
	return o.InputMessageContent, true
}

// HasInputMessageContent returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasInputMessageContent() bool {
	if o != nil && !IsNil(o.InputMessageContent) {
		return true
	}

	return false
}

// SetInputMessageContent gets a reference to the given InputMessageContent and assigns it to the InputMessageContent field.
func (o *InlineQueryResultLocation) SetInputMessageContent(v InputMessageContent) {
	o.InputMessageContent = &v
}


// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *InlineQueryResultLocation) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}


// GetThumbnailWidth returns the ThumbnailWidth field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetThumbnailWidth() int32 {
	if o == nil || IsNil(o.ThumbnailWidth) {
		var ret int32
		return ret
	}
	return *o.ThumbnailWidth
}

// GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetThumbnailWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailWidth) {
		return nil, false
	}
	return o.ThumbnailWidth, true
}

// HasThumbnailWidth returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasThumbnailWidth() bool {
	if o != nil && !IsNil(o.ThumbnailWidth) {
		return true
	}

	return false
}

// SetThumbnailWidth gets a reference to the given int32 and assigns it to the ThumbnailWidth field.
func (o *InlineQueryResultLocation) SetThumbnailWidth(v int32) {
	o.ThumbnailWidth = &v
}


// GetThumbnailHeight returns the ThumbnailHeight field value if set, zero value otherwise.
func (o *InlineQueryResultLocation) GetThumbnailHeight() int32 {
	if o == nil || IsNil(o.ThumbnailHeight) {
		var ret int32
		return ret
	}
	return *o.ThumbnailHeight
}

// GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultLocation) GetThumbnailHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailHeight) {
		return nil, false
	}
	return o.ThumbnailHeight, true
}

// HasThumbnailHeight returns a boolean if a field has been set.
func (o *InlineQueryResultLocation) HasThumbnailHeight() bool {
	if o != nil && !IsNil(o.ThumbnailHeight) {
		return true
	}

	return false
}

// SetThumbnailHeight gets a reference to the given int32 and assigns it to the ThumbnailHeight field.
func (o *InlineQueryResultLocation) SetThumbnailHeight(v int32) {
	o.ThumbnailHeight = &v
}


func (o InlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["title"] = o.Title
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

func (o *InlineQueryResultLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"latitude",
		"longitude",
		"title",
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

	varInlineQueryResultLocation := _InlineQueryResultLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultLocation)

	if err != nil {
		return err
	}

	*o = InlineQueryResultLocation(varInlineQueryResultLocation)

	return err
}

type NullableInlineQueryResultLocation struct {
	value *InlineQueryResultLocation
	isSet bool
}

func (v NullableInlineQueryResultLocation) Get() *InlineQueryResultLocation {
	return v.value
}

func (v *NullableInlineQueryResultLocation) Set(val *InlineQueryResultLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultLocation(val *InlineQueryResultLocation) *NullableInlineQueryResultLocation {
	return &NullableInlineQueryResultLocation{value: val, isSet: true}
}

func (v NullableInlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


