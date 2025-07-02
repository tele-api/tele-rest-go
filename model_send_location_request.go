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

// checks if the SendLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendLocationRequest{}

// SendLocationRequest Request parameters for sendLocation
type SendLocationRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	// Latitude of the location
	Latitude float32 `json:"latitude"`
	// Longitude of the location
	Longitude float32 `json:"longitude"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	// Period in seconds during which the location will be updated (see [Live Locations](https://telegram.org/blog/live-locations), should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int32 `json:"live_period,omitempty"`
	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int32 `json:"heading,omitempty"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int32 `json:"proximity_alert_radius,omitempty"`
	// Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup *SendMessageRequestReplyMarkup `json:"reply_markup,omitempty"`
}

type _SendLocationRequest SendLocationRequest

// NewSendLocationRequest instantiates a new SendLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendLocationRequest(chatId SendMessageRequestChatId, latitude float32, longitude float32) *SendLocationRequest {
	this := SendLocationRequest{}
	this.ChatId = chatId
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewSendLocationRequestWithDefaults instantiates a new SendLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendLocationRequestWithDefaults() *SendLocationRequest {
	this := SendLocationRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *SendLocationRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *SendLocationRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *SendLocationRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *SendLocationRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendLocationRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendLocationRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendLocationRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendLocationRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetLatitude returns the Latitude field value
func (o *SendLocationRequest) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *SendLocationRequest) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *SendLocationRequest) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *SendLocationRequest) SetLongitude(v float32) {
	o.Longitude = v
}

// GetHorizontalAccuracy returns the HorizontalAccuracy field value if set, zero value otherwise.
func (o *SendLocationRequest) GetHorizontalAccuracy() float32 {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		var ret float32
		return ret
	}
	return *o.HorizontalAccuracy
}

// GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetHorizontalAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HorizontalAccuracy) {
		return nil, false
	}
	return o.HorizontalAccuracy, true
}

// HasHorizontalAccuracy returns a boolean if a field has been set.
func (o *SendLocationRequest) HasHorizontalAccuracy() bool {
	if o != nil && !IsNil(o.HorizontalAccuracy) {
		return true
	}

	return false
}

// SetHorizontalAccuracy gets a reference to the given float32 and assigns it to the HorizontalAccuracy field.
func (o *SendLocationRequest) SetHorizontalAccuracy(v float32) {
	o.HorizontalAccuracy = &v
}


// GetLivePeriod returns the LivePeriod field value if set, zero value otherwise.
func (o *SendLocationRequest) GetLivePeriod() int32 {
	if o == nil || IsNil(o.LivePeriod) {
		var ret int32
		return ret
	}
	return *o.LivePeriod
}

// GetLivePeriodOk returns a tuple with the LivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetLivePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.LivePeriod) {
		return nil, false
	}
	return o.LivePeriod, true
}

// HasLivePeriod returns a boolean if a field has been set.
func (o *SendLocationRequest) HasLivePeriod() bool {
	if o != nil && !IsNil(o.LivePeriod) {
		return true
	}

	return false
}

// SetLivePeriod gets a reference to the given int32 and assigns it to the LivePeriod field.
func (o *SendLocationRequest) SetLivePeriod(v int32) {
	o.LivePeriod = &v
}


// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *SendLocationRequest) GetHeading() int32 {
	if o == nil || IsNil(o.Heading) {
		var ret int32
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetHeadingOk() (*int32, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *SendLocationRequest) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given int32 and assigns it to the Heading field.
func (o *SendLocationRequest) SetHeading(v int32) {
	o.Heading = &v
}


// GetProximityAlertRadius returns the ProximityAlertRadius field value if set, zero value otherwise.
func (o *SendLocationRequest) GetProximityAlertRadius() int32 {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		var ret int32
		return ret
	}
	return *o.ProximityAlertRadius
}

// GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetProximityAlertRadiusOk() (*int32, bool) {
	if o == nil || IsNil(o.ProximityAlertRadius) {
		return nil, false
	}
	return o.ProximityAlertRadius, true
}

// HasProximityAlertRadius returns a boolean if a field has been set.
func (o *SendLocationRequest) HasProximityAlertRadius() bool {
	if o != nil && !IsNil(o.ProximityAlertRadius) {
		return true
	}

	return false
}

// SetProximityAlertRadius gets a reference to the given int32 and assigns it to the ProximityAlertRadius field.
func (o *SendLocationRequest) SetProximityAlertRadius(v int32) {
	o.ProximityAlertRadius = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *SendLocationRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *SendLocationRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *SendLocationRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *SendLocationRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *SendLocationRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *SendLocationRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *SendLocationRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *SendLocationRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *SendLocationRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetMessageEffectId returns the MessageEffectId field value if set, zero value otherwise.
func (o *SendLocationRequest) GetMessageEffectId() string {
	if o == nil || IsNil(o.MessageEffectId) {
		var ret string
		return ret
	}
	return *o.MessageEffectId
}

// GetMessageEffectIdOk returns a tuple with the MessageEffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetMessageEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageEffectId) {
		return nil, false
	}
	return o.MessageEffectId, true
}

// HasMessageEffectId returns a boolean if a field has been set.
func (o *SendLocationRequest) HasMessageEffectId() bool {
	if o != nil && !IsNil(o.MessageEffectId) {
		return true
	}

	return false
}

// SetMessageEffectId gets a reference to the given string and assigns it to the MessageEffectId field.
func (o *SendLocationRequest) SetMessageEffectId(v string) {
	o.MessageEffectId = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *SendLocationRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *SendLocationRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *SendLocationRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *SendLocationRequest) GetReplyMarkup() SendMessageRequestReplyMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret SendMessageRequestReplyMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLocationRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *SendLocationRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given SendMessageRequestReplyMarkup and assigns it to the ReplyMarkup field.
func (o *SendLocationRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup) {
	o.ReplyMarkup = &v
}


func (o SendLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
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
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	if !IsNil(o.AllowPaidBroadcast) {
		toSerialize["allow_paid_broadcast"] = o.AllowPaidBroadcast
	}
	if !IsNil(o.MessageEffectId) {
		toSerialize["message_effect_id"] = o.MessageEffectId
	}
	if !IsNil(o.ReplyParameters) {
		toSerialize["reply_parameters"] = o.ReplyParameters
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *SendLocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
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

	varSendLocationRequest := _SendLocationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendLocationRequest)

	if err != nil {
		return err
	}

	*o = SendLocationRequest(varSendLocationRequest)

	return err
}

type NullableSendLocationRequest struct {
	value *SendLocationRequest
	isSet bool
}

func (v NullableSendLocationRequest) Get() *SendLocationRequest {
	return v.value
}

func (v *NullableSendLocationRequest) Set(val *SendLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendLocationRequest(val *SendLocationRequest) *NullableSendLocationRequest {
	return &NullableSendLocationRequest{value: val, isSet: true}
}

func (v NullableSendLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


