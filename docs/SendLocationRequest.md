# SendLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Latitude** | **float32** | Latitude of the location | 
**Longitude** | **float32** | Longitude of the location | 
**HorizontalAccuracy** | Pointer to **float32** | The radius of uncertainty for the location, measured in meters; 0-1500 | [optional] 
**LivePeriod** | Pointer to **int32** | Period in seconds during which the location will be updated (see [Live Locations](https://telegram.org/blog/live-locations), should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely. | [optional] 
**Heading** | Pointer to **int32** | For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified. | [optional] 
**ProximityAlertRadius** | Pointer to **int32** | For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified. | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendLocationRequest

`func NewSendLocationRequest(chatId SendMessageRequestChatId, latitude float32, longitude float32, ) *SendLocationRequest`

NewSendLocationRequest instantiates a new SendLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendLocationRequestWithDefaults

`func NewSendLocationRequestWithDefaults() *SendLocationRequest`

NewSendLocationRequestWithDefaults instantiates a new SendLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendLocationRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendLocationRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendLocationRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendLocationRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendLocationRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendLocationRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendLocationRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendLocationRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendLocationRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendLocationRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendLocationRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetLatitude

`func (o *SendLocationRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SendLocationRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SendLocationRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SendLocationRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SendLocationRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SendLocationRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetHorizontalAccuracy

`func (o *SendLocationRequest) GetHorizontalAccuracy() float32`

GetHorizontalAccuracy returns the HorizontalAccuracy field if non-nil, zero value otherwise.

### GetHorizontalAccuracyOk

`func (o *SendLocationRequest) GetHorizontalAccuracyOk() (*float32, bool)`

GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAccuracy

`func (o *SendLocationRequest) SetHorizontalAccuracy(v float32)`

SetHorizontalAccuracy sets HorizontalAccuracy field to given value.

### HasHorizontalAccuracy

`func (o *SendLocationRequest) HasHorizontalAccuracy() bool`

HasHorizontalAccuracy returns a boolean if a field has been set.

### GetLivePeriod

`func (o *SendLocationRequest) GetLivePeriod() int32`

GetLivePeriod returns the LivePeriod field if non-nil, zero value otherwise.

### GetLivePeriodOk

`func (o *SendLocationRequest) GetLivePeriodOk() (*int32, bool)`

GetLivePeriodOk returns a tuple with the LivePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePeriod

`func (o *SendLocationRequest) SetLivePeriod(v int32)`

SetLivePeriod sets LivePeriod field to given value.

### HasLivePeriod

`func (o *SendLocationRequest) HasLivePeriod() bool`

HasLivePeriod returns a boolean if a field has been set.

### GetHeading

`func (o *SendLocationRequest) GetHeading() int32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *SendLocationRequest) GetHeadingOk() (*int32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *SendLocationRequest) SetHeading(v int32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *SendLocationRequest) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetProximityAlertRadius

`func (o *SendLocationRequest) GetProximityAlertRadius() int32`

GetProximityAlertRadius returns the ProximityAlertRadius field if non-nil, zero value otherwise.

### GetProximityAlertRadiusOk

`func (o *SendLocationRequest) GetProximityAlertRadiusOk() (*int32, bool)`

GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertRadius

`func (o *SendLocationRequest) SetProximityAlertRadius(v int32)`

SetProximityAlertRadius sets ProximityAlertRadius field to given value.

### HasProximityAlertRadius

`func (o *SendLocationRequest) HasProximityAlertRadius() bool`

HasProximityAlertRadius returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendLocationRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendLocationRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendLocationRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendLocationRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendLocationRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendLocationRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendLocationRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendLocationRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendLocationRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendLocationRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendLocationRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendLocationRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendLocationRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendLocationRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendLocationRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendLocationRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendLocationRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendLocationRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendLocationRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendLocationRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendLocationRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendLocationRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendLocationRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendLocationRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


