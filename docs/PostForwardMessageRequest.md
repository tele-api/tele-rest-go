# PostForwardMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**PostForwardMessageRequestFromChatId**](PostForwardMessageRequestFromChatId.md) |  | 
**VideoStartTimestamp** | Pointer to **int32** | New start timestamp for the forwarded video in the message | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the forwarded message from forwarding and saving | [optional] 
**MessageId** | **int32** | Message identifier in the chat specified in *from\\_chat\\_id* | 

## Methods

### NewPostForwardMessageRequest

`func NewPostForwardMessageRequest(chatId PostSendMessageRequestChatId, fromChatId PostForwardMessageRequestFromChatId, messageId int32, ) *PostForwardMessageRequest`

NewPostForwardMessageRequest instantiates a new PostForwardMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostForwardMessageRequestWithDefaults

`func NewPostForwardMessageRequestWithDefaults() *PostForwardMessageRequest`

NewPostForwardMessageRequestWithDefaults instantiates a new PostForwardMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostForwardMessageRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostForwardMessageRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostForwardMessageRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *PostForwardMessageRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostForwardMessageRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostForwardMessageRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostForwardMessageRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *PostForwardMessageRequest) GetFromChatId() PostForwardMessageRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *PostForwardMessageRequest) GetFromChatIdOk() (*PostForwardMessageRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *PostForwardMessageRequest) SetFromChatId(v PostForwardMessageRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetVideoStartTimestamp

`func (o *PostForwardMessageRequest) GetVideoStartTimestamp() int32`

GetVideoStartTimestamp returns the VideoStartTimestamp field if non-nil, zero value otherwise.

### GetVideoStartTimestampOk

`func (o *PostForwardMessageRequest) GetVideoStartTimestampOk() (*int32, bool)`

GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoStartTimestamp

`func (o *PostForwardMessageRequest) SetVideoStartTimestamp(v int32)`

SetVideoStartTimestamp sets VideoStartTimestamp field to given value.

### HasVideoStartTimestamp

`func (o *PostForwardMessageRequest) HasVideoStartTimestamp() bool`

HasVideoStartTimestamp returns a boolean if a field has been set.

### GetDisableNotification

`func (o *PostForwardMessageRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *PostForwardMessageRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *PostForwardMessageRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *PostForwardMessageRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *PostForwardMessageRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PostForwardMessageRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PostForwardMessageRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PostForwardMessageRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetMessageId

`func (o *PostForwardMessageRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PostForwardMessageRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PostForwardMessageRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


