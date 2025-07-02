# SendMediaGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Media** | [**[]SendMediaGroupRequestMediaInner**](SendMediaGroupRequestMediaInner.md) | A JSON-serialized array describing messages to be sent, must include 2-10 items | 
**DisableNotification** | Pointer to **bool** | Sends messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent messages from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 

## Methods

### NewSendMediaGroupRequest

`func NewSendMediaGroupRequest(chatId SendMessageRequestChatId, media []SendMediaGroupRequestMediaInner, ) *SendMediaGroupRequest`

NewSendMediaGroupRequest instantiates a new SendMediaGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMediaGroupRequestWithDefaults

`func NewSendMediaGroupRequestWithDefaults() *SendMediaGroupRequest`

NewSendMediaGroupRequestWithDefaults instantiates a new SendMediaGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendMediaGroupRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendMediaGroupRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendMediaGroupRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendMediaGroupRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendMediaGroupRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendMediaGroupRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendMediaGroupRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendMediaGroupRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendMediaGroupRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendMediaGroupRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendMediaGroupRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetMedia

`func (o *SendMediaGroupRequest) GetMedia() []SendMediaGroupRequestMediaInner`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SendMediaGroupRequest) GetMediaOk() (*[]SendMediaGroupRequestMediaInner, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SendMediaGroupRequest) SetMedia(v []SendMediaGroupRequestMediaInner)`

SetMedia sets Media field to given value.


### GetDisableNotification

`func (o *SendMediaGroupRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendMediaGroupRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendMediaGroupRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendMediaGroupRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendMediaGroupRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendMediaGroupRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendMediaGroupRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendMediaGroupRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendMediaGroupRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendMediaGroupRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendMediaGroupRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendMediaGroupRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendMediaGroupRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendMediaGroupRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendMediaGroupRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendMediaGroupRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendMediaGroupRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendMediaGroupRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendMediaGroupRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendMediaGroupRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


