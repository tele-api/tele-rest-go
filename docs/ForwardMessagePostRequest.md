# ForwardMessagePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**ForwardMessagePostRequestFromChatId**](ForwardMessagePostRequestFromChatId.md) |  | 
**VideoStartTimestamp** | Pointer to **int32** | New start timestamp for the forwarded video in the message | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the forwarded message from forwarding and saving | [optional] 
**MessageId** | **int32** | Message identifier in the chat specified in *from\\_chat\\_id* | 

## Methods

### NewForwardMessagePostRequest

`func NewForwardMessagePostRequest(chatId SendMessagePostRequestChatId, fromChatId ForwardMessagePostRequestFromChatId, messageId int32, ) *ForwardMessagePostRequest`

NewForwardMessagePostRequest instantiates a new ForwardMessagePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardMessagePostRequestWithDefaults

`func NewForwardMessagePostRequestWithDefaults() *ForwardMessagePostRequest`

NewForwardMessagePostRequestWithDefaults instantiates a new ForwardMessagePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ForwardMessagePostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ForwardMessagePostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ForwardMessagePostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *ForwardMessagePostRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *ForwardMessagePostRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *ForwardMessagePostRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *ForwardMessagePostRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *ForwardMessagePostRequest) GetFromChatId() ForwardMessagePostRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *ForwardMessagePostRequest) GetFromChatIdOk() (*ForwardMessagePostRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *ForwardMessagePostRequest) SetFromChatId(v ForwardMessagePostRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetVideoStartTimestamp

`func (o *ForwardMessagePostRequest) GetVideoStartTimestamp() int32`

GetVideoStartTimestamp returns the VideoStartTimestamp field if non-nil, zero value otherwise.

### GetVideoStartTimestampOk

`func (o *ForwardMessagePostRequest) GetVideoStartTimestampOk() (*int32, bool)`

GetVideoStartTimestampOk returns a tuple with the VideoStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoStartTimestamp

`func (o *ForwardMessagePostRequest) SetVideoStartTimestamp(v int32)`

SetVideoStartTimestamp sets VideoStartTimestamp field to given value.

### HasVideoStartTimestamp

`func (o *ForwardMessagePostRequest) HasVideoStartTimestamp() bool`

HasVideoStartTimestamp returns a boolean if a field has been set.

### GetDisableNotification

`func (o *ForwardMessagePostRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *ForwardMessagePostRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *ForwardMessagePostRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *ForwardMessagePostRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *ForwardMessagePostRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *ForwardMessagePostRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *ForwardMessagePostRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *ForwardMessagePostRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetMessageId

`func (o *ForwardMessagePostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ForwardMessagePostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ForwardMessagePostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


