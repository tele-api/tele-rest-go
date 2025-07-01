# EditMessageLiveLocationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message to be edited was sent | [optional] 
**ChatId** | Pointer to [**EditMessageTextPostRequestChatId**](EditMessageTextPostRequestChatId.md) |  | [optional] 
**MessageId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit | [optional] 
**InlineMessageId** | Pointer to **string** | Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message | [optional] 
**Latitude** | **float32** | Latitude of new location | 
**Longitude** | **float32** | Longitude of new location | 
**LivePeriod** | Pointer to **int32** | New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current *live\\_period* by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then *live\\_period* remains unchanged | [optional] 
**HorizontalAccuracy** | Pointer to **float32** | The radius of uncertainty for the location, measured in meters; 0-1500 | [optional] 
**Heading** | Pointer to **int32** | Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified. | [optional] 
**ProximityAlertRadius** | Pointer to **int32** | The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified. | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewEditMessageLiveLocationPostRequest

`func NewEditMessageLiveLocationPostRequest(latitude float32, longitude float32, ) *EditMessageLiveLocationPostRequest`

NewEditMessageLiveLocationPostRequest instantiates a new EditMessageLiveLocationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMessageLiveLocationPostRequestWithDefaults

`func NewEditMessageLiveLocationPostRequestWithDefaults() *EditMessageLiveLocationPostRequest`

NewEditMessageLiveLocationPostRequestWithDefaults instantiates a new EditMessageLiveLocationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *EditMessageLiveLocationPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *EditMessageLiveLocationPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *EditMessageLiveLocationPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *EditMessageLiveLocationPostRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *EditMessageLiveLocationPostRequest) GetChatId() EditMessageTextPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditMessageLiveLocationPostRequest) GetChatIdOk() (*EditMessageTextPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditMessageLiveLocationPostRequest) SetChatId(v EditMessageTextPostRequestChatId)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *EditMessageLiveLocationPostRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *EditMessageLiveLocationPostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EditMessageLiveLocationPostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EditMessageLiveLocationPostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EditMessageLiveLocationPostRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *EditMessageLiveLocationPostRequest) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *EditMessageLiveLocationPostRequest) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *EditMessageLiveLocationPostRequest) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *EditMessageLiveLocationPostRequest) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.

### GetLatitude

`func (o *EditMessageLiveLocationPostRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EditMessageLiveLocationPostRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EditMessageLiveLocationPostRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *EditMessageLiveLocationPostRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EditMessageLiveLocationPostRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EditMessageLiveLocationPostRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetLivePeriod

`func (o *EditMessageLiveLocationPostRequest) GetLivePeriod() int32`

GetLivePeriod returns the LivePeriod field if non-nil, zero value otherwise.

### GetLivePeriodOk

`func (o *EditMessageLiveLocationPostRequest) GetLivePeriodOk() (*int32, bool)`

GetLivePeriodOk returns a tuple with the LivePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePeriod

`func (o *EditMessageLiveLocationPostRequest) SetLivePeriod(v int32)`

SetLivePeriod sets LivePeriod field to given value.

### HasLivePeriod

`func (o *EditMessageLiveLocationPostRequest) HasLivePeriod() bool`

HasLivePeriod returns a boolean if a field has been set.

### GetHorizontalAccuracy

`func (o *EditMessageLiveLocationPostRequest) GetHorizontalAccuracy() float32`

GetHorizontalAccuracy returns the HorizontalAccuracy field if non-nil, zero value otherwise.

### GetHorizontalAccuracyOk

`func (o *EditMessageLiveLocationPostRequest) GetHorizontalAccuracyOk() (*float32, bool)`

GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAccuracy

`func (o *EditMessageLiveLocationPostRequest) SetHorizontalAccuracy(v float32)`

SetHorizontalAccuracy sets HorizontalAccuracy field to given value.

### HasHorizontalAccuracy

`func (o *EditMessageLiveLocationPostRequest) HasHorizontalAccuracy() bool`

HasHorizontalAccuracy returns a boolean if a field has been set.

### GetHeading

`func (o *EditMessageLiveLocationPostRequest) GetHeading() int32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *EditMessageLiveLocationPostRequest) GetHeadingOk() (*int32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *EditMessageLiveLocationPostRequest) SetHeading(v int32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *EditMessageLiveLocationPostRequest) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetProximityAlertRadius

`func (o *EditMessageLiveLocationPostRequest) GetProximityAlertRadius() int32`

GetProximityAlertRadius returns the ProximityAlertRadius field if non-nil, zero value otherwise.

### GetProximityAlertRadiusOk

`func (o *EditMessageLiveLocationPostRequest) GetProximityAlertRadiusOk() (*int32, bool)`

GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertRadius

`func (o *EditMessageLiveLocationPostRequest) SetProximityAlertRadius(v int32)`

SetProximityAlertRadius sets ProximityAlertRadius field to given value.

### HasProximityAlertRadius

`func (o *EditMessageLiveLocationPostRequest) HasProximityAlertRadius() bool`

HasProximityAlertRadius returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *EditMessageLiveLocationPostRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *EditMessageLiveLocationPostRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *EditMessageLiveLocationPostRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *EditMessageLiveLocationPostRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


