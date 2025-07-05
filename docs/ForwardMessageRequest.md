# ForwardMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**ForwardMessageRequestFromChatId**](ForwardMessageRequestFromChatId.md) |  | 
**VideoStartTimestamp** | Pointer to **int32** | New start timestamp for the forwarded video in the message | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the forwarded message from forwarding and saving | [optional] 
**MessageId** | **int32** | Message identifier in the chat specified in *from\\_chat\\_id* | 

## Methods

### NewForwardMessageRequest

`func NewForwardMessageRequest(chatId SendMessageRequestChatId, fromChatId ForwardMessageRequestFromChatId, messageId int32, ) *ForwardMessageRequest`

NewForwardMessageRequest instantiates a new ForwardMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardMessageRequestWithDefaults

`func NewForwardMessageRequestWithDefaults() *ForwardMessageRequest`

NewForwardMessageRequestWithDefaults instantiates a new ForwardMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ForwardMessageRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ForwardMessageRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ForwardMessageRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *ForwardMessageRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *ForwardMessageRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *ForwardMessageRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *ForwardMessageRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *ForwardMessageRequest) GetFromChatId() ForwardMessageRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *ForwardMessageRequest) GetFromChatIdOk() (*ForwardMessageRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *ForwardMessageRequest) SetFromChatId(v ForwardMessageRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetVideoStartTimestamp

`func (o *ForwardMessageRequest) GetVideoStartTimestamp() int32`

GetVideoStartTimestamp returns the VideoStartTimestamp field if non-nil, zero value otherwise.

### GetVideoStartTimestampOk

`func (o *ForwardMessageRequest) GetVideoStartTimestampOk() (*int32, bool)`

GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoStartTimestamp

`func (o *ForwardMessageRequest) SetVideoStartTimestamp(v int32)`

SetVideoStartTimestamp sets VideoStartTimestamp field to given value.

### HasVideoStartTimestamp

`func (o *ForwardMessageRequest) HasVideoStartTimestamp() bool`

HasVideoStartTimestamp returns a boolean if a field has been set.

### GetDisableNotification

`func (o *ForwardMessageRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *ForwardMessageRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *ForwardMessageRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *ForwardMessageRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *ForwardMessageRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *ForwardMessageRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *ForwardMessageRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *ForwardMessageRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetMessageId

`func (o *ForwardMessageRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ForwardMessageRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ForwardMessageRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


