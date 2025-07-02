# PostForwardMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**PostForwardMessagesRequestFromChatId**](PostForwardMessagesRequestFromChatId.md) |  | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to forward. The identifiers must be specified in a strictly increasing order. | 
**DisableNotification** | Pointer to **bool** | Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the forwarded messages from forwarding and saving | [optional] 

## Methods

### NewPostForwardMessagesRequest

`func NewPostForwardMessagesRequest(chatId PostSendMessageRequestChatId, fromChatId PostForwardMessagesRequestFromChatId, messageIds []int32, ) *PostForwardMessagesRequest`

NewPostForwardMessagesRequest instantiates a new PostForwardMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostForwardMessagesRequestWithDefaults

`func NewPostForwardMessagesRequestWithDefaults() *PostForwardMessagesRequest`

NewPostForwardMessagesRequestWithDefaults instantiates a new PostForwardMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostForwardMessagesRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostForwardMessagesRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostForwardMessagesRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *PostForwardMessagesRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostForwardMessagesRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostForwardMessagesRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostForwardMessagesRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *PostForwardMessagesRequest) GetFromChatId() PostForwardMessagesRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *PostForwardMessagesRequest) GetFromChatIdOk() (*PostForwardMessagesRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *PostForwardMessagesRequest) SetFromChatId(v PostForwardMessagesRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetMessageIds

`func (o *PostForwardMessagesRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *PostForwardMessagesRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *PostForwardMessagesRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetDisableNotification

`func (o *PostForwardMessagesRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *PostForwardMessagesRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *PostForwardMessagesRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *PostForwardMessagesRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *PostForwardMessagesRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PostForwardMessagesRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PostForwardMessagesRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PostForwardMessagesRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


