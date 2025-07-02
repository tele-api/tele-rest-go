# SendDiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Emoji** | Pointer to **string** | Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults to “🎲” | [optional] [default to "🎲"]
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessageRequestReplyMarkup**](SendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendDiceRequest

`func NewSendDiceRequest(chatId SendMessageRequestChatId, ) *SendDiceRequest`

NewSendDiceRequest instantiates a new SendDiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDiceRequestWithDefaults

`func NewSendDiceRequestWithDefaults() *SendDiceRequest`

NewSendDiceRequestWithDefaults instantiates a new SendDiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendDiceRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendDiceRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendDiceRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendDiceRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendDiceRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendDiceRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendDiceRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendDiceRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendDiceRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendDiceRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendDiceRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetEmoji

`func (o *SendDiceRequest) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *SendDiceRequest) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *SendDiceRequest) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *SendDiceRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendDiceRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendDiceRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendDiceRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendDiceRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendDiceRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendDiceRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendDiceRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendDiceRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendDiceRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendDiceRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendDiceRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendDiceRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendDiceRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendDiceRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendDiceRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendDiceRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendDiceRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendDiceRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendDiceRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendDiceRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendDiceRequest) GetReplyMarkup() SendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendDiceRequest) GetReplyMarkupOk() (*SendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendDiceRequest) SetReplyMarkup(v SendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendDiceRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


