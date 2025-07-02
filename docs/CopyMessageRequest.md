# CopyMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**ForwardMessageRequestFromChatId**](ForwardMessageRequestFromChatId.md) |  | 
**MessageId** | **int32** | Message identifier in the chat specified in *from\\_chat\\_id* | 
**VideoStartTimestamp** | Pointer to **int32** | New start timestamp for the copied video in the message | [optional] 
**Caption** | Pointer to **string** | New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept | [optional] 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the new caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | Pass *True*, if the caption must be shown above the message media. Ignored if a new caption isn&#39;t specified. | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewCopyMessageRequest

`func NewCopyMessageRequest(chatId SendMessageRequestChatId, fromChatId ForwardMessageRequestFromChatId, messageId int32, ) *CopyMessageRequest`

NewCopyMessageRequest instantiates a new CopyMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyMessageRequestWithDefaults

`func NewCopyMessageRequestWithDefaults() *CopyMessageRequest`

NewCopyMessageRequestWithDefaults instantiates a new CopyMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CopyMessageRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CopyMessageRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CopyMessageRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *CopyMessageRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *CopyMessageRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *CopyMessageRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *CopyMessageRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *CopyMessageRequest) GetFromChatId() ForwardMessageRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *CopyMessageRequest) GetFromChatIdOk() (*ForwardMessageRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *CopyMessageRequest) SetFromChatId(v ForwardMessageRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetMessageId

`func (o *CopyMessageRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *CopyMessageRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *CopyMessageRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetVideoStartTimestamp

`func (o *CopyMessageRequest) GetVideoStartTimestamp() int32`

GetVideoStartTimestamp returns the VideoStartTimestamp field if non-nil, zero value otherwise.

### GetVideoStartTimestampOk

`func (o *CopyMessageRequest) GetVideoStartTimestampOk() (*int32, bool)`

GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoStartTimestamp

`func (o *CopyMessageRequest) SetVideoStartTimestamp(v int32)`

SetVideoStartTimestamp sets VideoStartTimestamp field to given value.

### HasVideoStartTimestamp

`func (o *CopyMessageRequest) HasVideoStartTimestamp() bool`

HasVideoStartTimestamp returns a boolean if a field has been set.

### GetCaption

`func (o *CopyMessageRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CopyMessageRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CopyMessageRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CopyMessageRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *CopyMessageRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *CopyMessageRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *CopyMessageRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *CopyMessageRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *CopyMessageRequest) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *CopyMessageRequest) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *CopyMessageRequest) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *CopyMessageRequest) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *CopyMessageRequest) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *CopyMessageRequest) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *CopyMessageRequest) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *CopyMessageRequest) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetDisableNotification

`func (o *CopyMessageRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *CopyMessageRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *CopyMessageRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *CopyMessageRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *CopyMessageRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *CopyMessageRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *CopyMessageRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *CopyMessageRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *CopyMessageRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *CopyMessageRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *CopyMessageRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *CopyMessageRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetReplyParameters

`func (o *CopyMessageRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *CopyMessageRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *CopyMessageRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *CopyMessageRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *CopyMessageRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *CopyMessageRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *CopyMessageRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *CopyMessageRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


