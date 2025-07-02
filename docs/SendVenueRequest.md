# SendVenueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Latitude** | **float32** | Latitude of the venue | 
**Longitude** | **float32** | Longitude of the venue | 
**Title** | **string** | Name of the venue | 
**Address** | **string** | Address of the venue | 
**FoursquareId** | Pointer to **string** | Foursquare identifier of the venue | [optional] 
**FoursquareType** | Pointer to **string** | Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.) | [optional] 
**GooglePlaceId** | Pointer to **string** | Google Places identifier of the venue | [optional] 
**GooglePlaceType** | Pointer to **string** | Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).) | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendVenueRequest

`func NewSendVenueRequest(chatId SendMessageRequestChatId, latitude float32, longitude float32, title string, address string, ) *SendVenueRequest`

NewSendVenueRequest instantiates a new SendVenueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendVenueRequestWithDefaults

`func NewSendVenueRequestWithDefaults() *SendVenueRequest`

NewSendVenueRequestWithDefaults instantiates a new SendVenueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendVenueRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendVenueRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendVenueRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendVenueRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendVenueRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendVenueRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendVenueRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendVenueRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendVenueRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendVenueRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendVenueRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetLatitude

`func (o *SendVenueRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SendVenueRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SendVenueRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SendVenueRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SendVenueRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SendVenueRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetTitle

`func (o *SendVenueRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendVenueRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendVenueRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAddress

`func (o *SendVenueRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SendVenueRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SendVenueRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFoursquareId

`func (o *SendVenueRequest) GetFoursquareId() string`

GetFoursquareId returns the FoursquareId field if non-nil, zero value otherwise.

### GetFoursquareIdOk

`func (o *SendVenueRequest) GetFoursquareIdOk() (*string, bool)`

GetFoursquareIdOk returns a tuple with the FoursquareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareId

`func (o *SendVenueRequest) SetFoursquareId(v string)`

SetFoursquareId sets FoursquareId field to given value.

### HasFoursquareId

`func (o *SendVenueRequest) HasFoursquareId() bool`

HasFoursquareId returns a boolean if a field has been set.

### GetFoursquareType

`func (o *SendVenueRequest) GetFoursquareType() string`

GetFoursquareType returns the FoursquareType field if non-nil, zero value otherwise.

### GetFoursquareTypeOk

`func (o *SendVenueRequest) GetFoursquareTypeOk() (*string, bool)`

GetFoursquareTypeOk returns a tuple with the FoursquareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareType

`func (o *SendVenueRequest) SetFoursquareType(v string)`

SetFoursquareType sets FoursquareType field to given value.

### HasFoursquareType

`func (o *SendVenueRequest) HasFoursquareType() bool`

HasFoursquareType returns a boolean if a field has been set.

### GetGooglePlaceId

`func (o *SendVenueRequest) GetGooglePlaceId() string`

GetGooglePlaceId returns the GooglePlaceId field if non-nil, zero value otherwise.

### GetGooglePlaceIdOk

`func (o *SendVenueRequest) GetGooglePlaceIdOk() (*string, bool)`

GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceId

`func (o *SendVenueRequest) SetGooglePlaceId(v string)`

SetGooglePlaceId sets GooglePlaceId field to given value.

### HasGooglePlaceId

`func (o *SendVenueRequest) HasGooglePlaceId() bool`

HasGooglePlaceId returns a boolean if a field has been set.

### GetGooglePlaceType

`func (o *SendVenueRequest) GetGooglePlaceType() string`

GetGooglePlaceType returns the GooglePlaceType field if non-nil, zero value otherwise.

### GetGooglePlaceTypeOk

`func (o *SendVenueRequest) GetGooglePlaceTypeOk() (*string, bool)`

GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceType

`func (o *SendVenueRequest) SetGooglePlaceType(v string)`

SetGooglePlaceType sets GooglePlaceType field to given value.

### HasGooglePlaceType

`func (o *SendVenueRequest) HasGooglePlaceType() bool`

HasGooglePlaceType returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendVenueRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendVenueRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendVenueRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendVenueRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendVenueRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendVenueRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendVenueRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendVenueRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendVenueRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendVenueRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendVenueRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendVenueRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendVenueRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendVenueRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendVenueRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendVenueRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendVenueRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendVenueRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendVenueRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendVenueRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendVenueRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendVenueRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendVenueRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendVenueRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


