# SendVideoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Video** | **NullableString** |  | 
**Duration** | Pointer to **int32** | Duration of sent video in seconds | [optional] 
**Width** | Pointer to **int32** | Video width | [optional] 
**Height** | Pointer to **int32** | Video height | [optional] 
**Thumbnail** | Pointer to **NullableString** |  | [optional] 
**Cover** | Pointer to **NullableString** |  | [optional] 
**StartTimestamp** | Pointer to **int32** | Start timestamp for the video in the message | [optional] 
**Caption** | Pointer to **string** | Video caption (may also be used when resending videos by *file\\_id*), 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | Pass *True*, if the caption must be shown above the message media | [optional] 
**HasSpoiler** | Pointer to **bool** | Pass *True* if the video needs to be covered with a spoiler animation | [optional] 
**SupportsStreaming** | Pointer to **bool** | Pass *True* if the uploaded video is suitable for streaming | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendVideoRequest

`func NewSendVideoRequest(chatId SendMessageRequestChatId, video NullableString, ) *SendVideoRequest`

NewSendVideoRequest instantiates a new SendVideoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendVideoRequestWithDefaults

`func NewSendVideoRequestWithDefaults() *SendVideoRequest`

NewSendVideoRequestWithDefaults instantiates a new SendVideoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendVideoRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendVideoRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendVideoRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendVideoRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendVideoRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendVideoRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendVideoRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendVideoRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendVideoRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendVideoRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendVideoRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetVideo

`func (o *SendVideoRequest) GetVideo() string`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *SendVideoRequest) GetVideoOk() (*string, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *SendVideoRequest) SetVideo(v string)`

SetVideo sets Video field to given value.


### SetVideoNil

`func (o *SendVideoRequest) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *SendVideoRequest) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetDuration

`func (o *SendVideoRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SendVideoRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SendVideoRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SendVideoRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWidth

`func (o *SendVideoRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SendVideoRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SendVideoRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SendVideoRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *SendVideoRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SendVideoRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SendVideoRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SendVideoRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetThumbnail

`func (o *SendVideoRequest) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *SendVideoRequest) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *SendVideoRequest) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *SendVideoRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *SendVideoRequest) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *SendVideoRequest) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetCover

`func (o *SendVideoRequest) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *SendVideoRequest) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *SendVideoRequest) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *SendVideoRequest) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *SendVideoRequest) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *SendVideoRequest) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetStartTimestamp

`func (o *SendVideoRequest) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SendVideoRequest) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SendVideoRequest) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *SendVideoRequest) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetCaption

`func (o *SendVideoRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SendVideoRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SendVideoRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *SendVideoRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *SendVideoRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *SendVideoRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *SendVideoRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *SendVideoRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *SendVideoRequest) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *SendVideoRequest) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *SendVideoRequest) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *SendVideoRequest) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *SendVideoRequest) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *SendVideoRequest) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *SendVideoRequest) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *SendVideoRequest) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetHasSpoiler

`func (o *SendVideoRequest) GetHasSpoiler() bool`

GetHasSpoiler returns the HasSpoiler field if non-nil, zero value otherwise.

### GetHasSpoilerOk

`func (o *SendVideoRequest) GetHasSpoilerOk() (*bool, bool)`

GetHasSpoilerOk returns a tuple with the HasSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpoiler

`func (o *SendVideoRequest) SetHasSpoiler(v bool)`

SetHasSpoiler sets HasSpoiler field to given value.

### HasHasSpoiler

`func (o *SendVideoRequest) HasHasSpoiler() bool`

HasHasSpoiler returns a boolean if a field has been set.

### GetSupportsStreaming

`func (o *SendVideoRequest) GetSupportsStreaming() bool`

GetSupportsStreaming returns the SupportsStreaming field if non-nil, zero value otherwise.

### GetSupportsStreamingOk

`func (o *SendVideoRequest) GetSupportsStreamingOk() (*bool, bool)`

GetSupportsStreamingOk returns a tuple with the SupportsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStreaming

`func (o *SendVideoRequest) SetSupportsStreaming(v bool)`

SetSupportsStreaming sets SupportsStreaming field to given value.

### HasSupportsStreaming

`func (o *SendVideoRequest) HasSupportsStreaming() bool`

HasSupportsStreaming returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendVideoRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendVideoRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendVideoRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendVideoRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendVideoRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendVideoRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendVideoRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendVideoRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendVideoRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendVideoRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendVideoRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendVideoRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendVideoRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendVideoRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendVideoRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendVideoRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendVideoRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendVideoRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendVideoRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendVideoRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendVideoRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendVideoRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendVideoRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendVideoRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


