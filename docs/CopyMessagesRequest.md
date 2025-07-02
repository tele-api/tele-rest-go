# CopyMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**ForwardMessagesRequestFromChatId**](ForwardMessagesRequestFromChatId.md) |  | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to copy. The identifiers must be specified in a strictly increasing order. | 
**DisableNotification** | Pointer to **bool** | Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent messages from forwarding and saving | [optional] 
**RemoveCaption** | Pointer to **bool** | Pass *True* to copy the messages without their captions | [optional] 

## Methods

### NewCopyMessagesRequest

`func NewCopyMessagesRequest(chatId SendMessageRequestChatId, fromChatId ForwardMessagesRequestFromChatId, messageIds []int32, ) *CopyMessagesRequest`

NewCopyMessagesRequest instantiates a new CopyMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyMessagesRequestWithDefaults

`func NewCopyMessagesRequestWithDefaults() *CopyMessagesRequest`

NewCopyMessagesRequestWithDefaults instantiates a new CopyMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CopyMessagesRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CopyMessagesRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CopyMessagesRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *CopyMessagesRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *CopyMessagesRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *CopyMessagesRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *CopyMessagesRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *CopyMessagesRequest) GetFromChatId() ForwardMessagesRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *CopyMessagesRequest) GetFromChatIdOk() (*ForwardMessagesRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *CopyMessagesRequest) SetFromChatId(v ForwardMessagesRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetMessageIds

`func (o *CopyMessagesRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *CopyMessagesRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *CopyMessagesRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetDisableNotification

`func (o *CopyMessagesRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *CopyMessagesRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *CopyMessagesRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *CopyMessagesRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *CopyMessagesRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *CopyMessagesRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *CopyMessagesRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *CopyMessagesRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetRemoveCaption

`func (o *CopyMessagesRequest) GetRemoveCaption() bool`

GetRemoveCaption returns the RemoveCaption field if non-nil, zero value otherwise.

### GetRemoveCaptionOk

`func (o *CopyMessagesRequest) GetRemoveCaptionOk() (*bool, bool)`

GetRemoveCaptionOk returns a tuple with the RemoveCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveCaption

`func (o *CopyMessagesRequest) SetRemoveCaption(v bool)`

SetRemoveCaption sets RemoveCaption field to given value.

### HasRemoveCaption

`func (o *CopyMessagesRequest) HasRemoveCaption() bool`

HasRemoveCaption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


