# SendAnimationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Animation** | **NullableString** |  | 
**Duration** | Pointer to **int32** | Duration of sent animation in seconds | [optional] 
**Width** | Pointer to **int32** | Animation width | [optional] 
**Height** | Pointer to **int32** | Animation height | [optional] 
**Thumbnail** | Pointer to **NullableString** |  | [optional] 
**Caption** | Pointer to **string** | Animation caption (may also be used when resending animation by *file\\_id*), 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the animation caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | Pass *True*, if the caption must be shown above the message media | [optional] 
**HasSpoiler** | Pointer to **bool** | Pass *True* if the animation needs to be covered with a spoiler animation | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendAnimationRequest

`func NewSendAnimationRequest(chatId SendMessageRequestChatId, animation NullableString, ) *SendAnimationRequest`

NewSendAnimationRequest instantiates a new SendAnimationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendAnimationRequestWithDefaults

`func NewSendAnimationRequestWithDefaults() *SendAnimationRequest`

NewSendAnimationRequestWithDefaults instantiates a new SendAnimationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendAnimationRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendAnimationRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendAnimationRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendAnimationRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendAnimationRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendAnimationRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendAnimationRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendAnimationRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendAnimationRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendAnimationRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendAnimationRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetAnimation

`func (o *SendAnimationRequest) GetAnimation() string`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *SendAnimationRequest) GetAnimationOk() (*string, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *SendAnimationRequest) SetAnimation(v string)`

SetAnimation sets Animation field to given value.


### SetAnimationNil

`func (o *SendAnimationRequest) SetAnimationNil(b bool)`

 SetAnimationNil sets the value for Animation to be an explicit nil

### UnsetAnimation
`func (o *SendAnimationRequest) UnsetAnimation()`

UnsetAnimation ensures that no value is present for Animation, not even an explicit nil
### GetDuration

`func (o *SendAnimationRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SendAnimationRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SendAnimationRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SendAnimationRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWidth

`func (o *SendAnimationRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SendAnimationRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SendAnimationRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SendAnimationRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *SendAnimationRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SendAnimationRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SendAnimationRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SendAnimationRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetThumbnail

`func (o *SendAnimationRequest) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *SendAnimationRequest) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *SendAnimationRequest) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *SendAnimationRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *SendAnimationRequest) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *SendAnimationRequest) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetCaption

`func (o *SendAnimationRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SendAnimationRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SendAnimationRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *SendAnimationRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *SendAnimationRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *SendAnimationRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *SendAnimationRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *SendAnimationRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *SendAnimationRequest) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *SendAnimationRequest) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *SendAnimationRequest) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *SendAnimationRequest) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *SendAnimationRequest) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *SendAnimationRequest) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *SendAnimationRequest) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *SendAnimationRequest) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetHasSpoiler

`func (o *SendAnimationRequest) GetHasSpoiler() bool`

GetHasSpoiler returns the HasSpoiler field if non-nil, zero value otherwise.

### GetHasSpoilerOk

`func (o *SendAnimationRequest) GetHasSpoilerOk() (*bool, bool)`

GetHasSpoilerOk returns a tuple with the HasSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpoiler

`func (o *SendAnimationRequest) SetHasSpoiler(v bool)`

SetHasSpoiler sets HasSpoiler field to given value.

### HasHasSpoiler

`func (o *SendAnimationRequest) HasHasSpoiler() bool`

HasHasSpoiler returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendAnimationRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendAnimationRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendAnimationRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendAnimationRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendAnimationRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendAnimationRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendAnimationRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendAnimationRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendAnimationRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendAnimationRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendAnimationRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendAnimationRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendAnimationRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendAnimationRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendAnimationRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendAnimationRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendAnimationRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendAnimationRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendAnimationRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendAnimationRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendAnimationRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendAnimationRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendAnimationRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendAnimationRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


