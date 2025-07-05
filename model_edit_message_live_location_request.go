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

// checks if the EditMessageLiveLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditMessageLiveLocationRequest{}

// EditMessageLiveLocationRequest Request parameters for editMessageLiveLocation
type EditMessageLiveLocationRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId *EditMessageTextRequestChatId `json:"chat_id,omitempty"`
	// Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit
	MessageId *int32 `json:"message_id,omitempty"`
	// Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// Latitude of new location
	Latitude float32 `json:"latitude"`
	// Longitude of new location
	Longitude float32 `json:"longitude"`
	// New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current *live\\_period* by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then *live\\_period* remains unchanged
	LivePeriod *int32 `json:"live_period,omitempty"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int32 `json:"heading,omitempty"`
	// The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int32 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type _EditMessageLiveLocationRequest EditMessageLiveLocationRequest

// NewEditMessageLiveLocationRequest instantiates a new EditMessageLiveLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMessageLiveLocationRequest(latitude float32, longitude float32) *EditMessageLiveLocationRequest {
	this := EditMessageLiveLocationRequest{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewEditMessageLiveLocationRequestWithDefaults instantiates a new EditMessageLiveLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMessageLiveLocationRequestWithDefaults() *EditMessageLiveLocationRequest {
	this := EditMessageLiveLocationRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *EditMessageLiveLocationRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetChatId() EditMessageTextRequestChatId {
	if o == nil || IsNil(o.ChatId) {
		var ret EditMessageTextRequestChatId
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetChatIdOk() (*EditMessageTextRequestChatId, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given EditMessageTextRequestChatId and assigns it to the ChatId field.
func (o *EditMessageLiveLocationRequest) SetChatId(v EditMessageTextRequestChatId) {
	o.ChatId = &v
}


// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *EditMessageLiveLocationRequest) SetMessageId(v int32) {
	o.MessageId = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *EditMessageLiveLocationRequest) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


// GetLatitude returns the Latitude field value
func (o *EditMessageLiveLocationRequest) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *EditMessageLiveLocationRequest) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *EditMessageLiveLocationRequest) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *EditMessageLiveLocationRequest) SetLongitude(v float32) {
	o.Longitude = v
}

// GetLivePeriod returns the LivePeriod field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetLivePeriod() int32 {
	if o == nil || IsNil(o.LivePeriod) {
		var ret int32
		return ret
	}
	return *o.LivePeriod
}

// GetLivePeriodOk returns a tuple with the LivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetLivePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.LivePeriod) {
		return nil, false
	}
	return o.LivePeriod, true
}

// HasLivePeriod returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasLivePeriod() bool {
	if o != nil && !IsNil(o.LivePeriod) {
		return true
	}

	return false
}

// SetLivePeriod gets a reference to the given int32 and assigns it to the LivePeriod field.
func (o *EditMessageLiveLocationRequest) SetLivePeriod(v int32) {
	o.LivePeriod = &v
}


// GetHorizontalAccuracy returns the HorizontalAccuracy field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetHorizontalAccuracy() float32 {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		var ret float32
		return ret
	}
	return *o.HorizontalAccuracy
}

// GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetHorizontalAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		return nil, false
	}
	return o.HorizontalAccuracy, true
}

// HasHorizontalAccuracy returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasHorizontalAccuracy() bool {
	if o != nil && !IsNil(o.HorizontalAccuracy) {
		return true
	}

	return false
}

// SetHorizontalAccuracy gets a reference to the given float32 and assigns it to the HorizontalAccuracy field.
func (o *EditMessageLiveLocationRequest) SetHorizontalAccuracy(v float32) {
	o.HorizontalAccuracy = &v
}


// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetHeading() int32 {
	if o == nil || IsNil(o.Heading) {
		var ret int32
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetHeadingOk() (*int32, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given int32 and assigns it to the Heading field.
func (o *EditMessageLiveLocationRequest) SetHeading(v int32) {
	o.Heading = &v
}


// GetProximityAlertRadius returns the ProximityAlertRadius field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetProximityAlertRadius() int32 {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		var ret int32
		return ret
	}
	return *o.ProximityAlertRadius
}

// GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetProximityAlertRadiusOk() (*int32, bool) {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		return nil, false
	}
	return o.ProximityAlertRadius, true
}

// HasProximityAlertRadius returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasProximityAlertRadius() bool {
	if o != nil && !IsNil(o.ProximityAlertRadius) {
		return true
	}

	return false
}

// SetProximityAlertRadius gets a reference to the given int32 and assigns it to the ProximityAlertRadius field.
func (o *EditMessageLiveLocationRequest) SetProximityAlertRadius(v int32) {
	o.ProximityAlertRadius = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *EditMessageLiveLocationRequest) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMessageLiveLocationRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *EditMessageLiveLocationRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *EditMessageLiveLocationRequest) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o EditMessageLiveLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditMessageLiveLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.InlineMessageId) {
		toSerialize["inline_message_id"] = o.InlineMessageId
	}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	if !IsNil(o.LivePeriod) {
		toSerialize["live_period"] = o.LivePeriod
	}
	if !IsNil(o.HorizontalAccuracy) {
		toSerialize["horizontal_accuracy"] = o.HorizontalAccuracy
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
	return toSerialize, nil
}

func (o *EditMessageLiveLocationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEditMessageLiveLocationRequest := _EditMessageLiveLocationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditMessageLiveLocationRequest)

	if err != nil {
		return err
	}

	*o = EditMessageLiveLocationRequest(varEditMessageLiveLocationRequest)

	return err
}

type NullableEditMessageLiveLocationRequest struct {
	value *EditMessageLiveLocationRequest
	isSet bool
}

func (v NullableEditMessageLiveLocationRequest) Get() *EditMessageLiveLocationRequest {
	return v.value
}

func (v *NullableEditMessageLiveLocationRequest) Set(val *EditMessageLiveLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMessageLiveLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMessageLiveLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMessageLiveLocationRequest(val *EditMessageLiveLocationRequest) *NullableEditMessageLiveLocationRequest {
	return &NullableEditMessageLiveLocationRequest{value: val, isSet: true}
}

func (v NullableEditMessageLiveLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMessageLiveLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


