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

// checks if the SendVenueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendVenueRequest{}

// SendVenueRequest Request parameters for sendVenue
type SendVenueRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	// Latitude of the venue
	Latitude float32 `json:"latitude"`
	// Longitude of the venue
	Longitude float32 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Foursquare identifier of the venue
	FoursquareId *string `json:"foursquare_id,omitempty"`
	// Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	// Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	// Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
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

type _SendVenueRequest SendVenueRequest

// NewSendVenueRequest instantiates a new SendVenueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendVenueRequest(chatId SendMessageRequestChatId, latitude float32, longitude float32, title string, address string) *SendVenueRequest {
	this := SendVenueRequest{}
	this.ChatId = chatId
	this.Latitude = latitude
	this.Longitude = longitude
	this.Title = title
	this.Address = address
	return &this
}

// NewSendVenueRequestWithDefaults instantiates a new SendVenueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendVenueRequestWithDefaults() *SendVenueRequest {
	this := SendVenueRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *SendVenueRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *SendVenueRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *SendVenueRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetChatId returns the ChatId field value
func (o *SendVenueRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendVenueRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendVenueRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendVenueRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendVenueRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetLatitude returns the Latitude field value
func (o *SendVenueRequest) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *SendVenueRequest) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *SendVenueRequest) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *SendVenueRequest) SetLongitude(v float32) {
	o.Longitude = v
}

// GetTitle returns the Title field value
func (o *SendVenueRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SendVenueRequest) SetTitle(v string) {
	o.Title = v
}

// GetAddress returns the Address field value
func (o *SendVenueRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *SendVenueRequest) SetAddress(v string) {
	o.Address = v
}

// GetFoursquareId returns the FoursquareId field value if set, zero value otherwise.
func (o *SendVenueRequest) GetFoursquareId() string {
	if o == nil || IsNil(o.FoursquareId) {
		var ret string
		return ret
	}
	return *o.FoursquareId
}

// GetFoursquareIdOk returns a tuple with the FoursquareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetFoursquareIdOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareId) {
		return nil, false
	}
	return o.FoursquareId, true
}

// HasFoursquareId returns a boolean if a field has been set.
func (o *SendVenueRequest) HasFoursquareId() bool {
	if o != nil && !IsNil(o.FoursquareId) {
		return true
	}

	return false
}

// SetFoursquareId gets a reference to the given string and assigns it to the FoursquareId field.
func (o *SendVenueRequest) SetFoursquareId(v string) {
	o.FoursquareId = &v
}


// GetFoursquareType returns the FoursquareType field value if set, zero value otherwise.
func (o *SendVenueRequest) GetFoursquareType() string {
	if o == nil || IsNil(o.FoursquareType) {
		var ret string
		return ret
	}
	return *o.FoursquareType
}

// GetFoursquareTypeOk returns a tuple with the FoursquareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetFoursquareTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FoursquareType) {
		return nil, false
	}
	return o.FoursquareType, true
}

// HasFoursquareType returns a boolean if a field has been set.
func (o *SendVenueRequest) HasFoursquareType() bool {
	if o != nil && !IsNil(o.FoursquareType) {
		return true
	}

	return false
}

// SetFoursquareType gets a reference to the given string and assigns it to the FoursquareType field.
func (o *SendVenueRequest) SetFoursquareType(v string) {
	o.FoursquareType = &v
}


// GetGooglePlaceId returns the GooglePlaceId field value if set, zero value otherwise.
func (o *SendVenueRequest) GetGooglePlaceId() string {
	if o == nil || IsNil(o.GooglePlaceId) {
		var ret string
		return ret
	}
	return *o.GooglePlaceId
}

// GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetGooglePlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceId) {
		return nil, false
	}
	return o.GooglePlaceId, true
}

// HasGooglePlaceId returns a boolean if a field has been set.
func (o *SendVenueRequest) HasGooglePlaceId() bool {
	if o != nil && !IsNil(o.GooglePlaceId) {
		return true
	}

	return false
}

// SetGooglePlaceId gets a reference to the given string and assigns it to the GooglePlaceId field.
func (o *SendVenueRequest) SetGooglePlaceId(v string) {
	o.GooglePlaceId = &v
}


// GetGooglePlaceType returns the GooglePlaceType field value if set, zero value otherwise.
func (o *SendVenueRequest) GetGooglePlaceType() string {
	if o == nil || IsNil(o.GooglePlaceType) {
		var ret string
		return ret
	}
	return *o.GooglePlaceType
}

// GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetGooglePlaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GooglePlaceType) {
		return nil, false
	}
	return o.GooglePlaceType, true
}

// HasGooglePlaceType returns a boolean if a field has been set.
func (o *SendVenueRequest) HasGooglePlaceType() bool {
	if o != nil && !IsNil(o.GooglePlaceType) {
		return true
	}

	return false
}

// SetGooglePlaceType gets a reference to the given string and assigns it to the GooglePlaceType field.
func (o *SendVenueRequest) SetGooglePlaceType(v string) {
	o.GooglePlaceType = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *SendVenueRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *SendVenueRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *SendVenueRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *SendVenueRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *SendVenueRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *SendVenueRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *SendVenueRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *SendVenueRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *SendVenueRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetMessageEffectId returns the MessageEffectId field value if set, zero value otherwise.
func (o *SendVenueRequest) GetMessageEffectId() string {
	if o == nil || IsNil(o.MessageEffectId) {
		var ret string
		return ret
	}
	return *o.MessageEffectId
}

// GetMessageEffectIdOk returns a tuple with the MessageEffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetMessageEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageEffectId) {
		return nil, false
	}
	return o.MessageEffectId, true
}

// HasMessageEffectId returns a boolean if a field has been set.
func (o *SendVenueRequest) HasMessageEffectId() bool {
	if o != nil && !IsNil(o.MessageEffectId) {
		return true
	}

	return false
}

// SetMessageEffectId gets a reference to the given string and assigns it to the MessageEffectId field.
func (o *SendVenueRequest) SetMessageEffectId(v string) {
	o.MessageEffectId = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *SendVenueRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *SendVenueRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *SendVenueRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *SendVenueRequest) GetReplyMarkup() SendMessageRequestReplyMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret SendMessageRequestReplyMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVenueRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *SendVenueRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given SendMessageRequestReplyMarkup and assigns it to the ReplyMarkup field.
func (o *SendVenueRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup) {
	o.ReplyMarkup = &v
}


func (o SendVenueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendVenueRequest) ToMap() (map[string]interface{}, error) {
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

func (o *SendVenueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
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

	varSendVenueRequest := _SendVenueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendVenueRequest)

	if err != nil {
		return err
	}

	*o = SendVenueRequest(varSendVenueRequest)

	return err
}

type NullableSendVenueRequest struct {
	value *SendVenueRequest
	isSet bool
}

func (v NullableSendVenueRequest) Get() *SendVenueRequest {
	return v.value
}

func (v *NullableSendVenueRequest) Set(val *SendVenueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendVenueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendVenueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendVenueRequest(val *SendVenueRequest) *NullableSendVenueRequest {
	return &NullableSendVenueRequest{value: val, isSet: true}
}

func (v NullableSendVenueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendVenueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


