# PostCopyMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**PostForwardMessagesRequestFromChatId**](PostForwardMessagesRequestFromChatId.md) |  | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to copy. The identifiers must be specified in a strictly increasing order. | 
**DisableNotification** | Pointer to **bool** | Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent messages from forwarding and saving | [optional] 
**RemoveCaption** | Pointer to **bool** | Pass *True* to copy the messages without their captions | [optional] 

## Methods

### NewPostCopyMessagesRequest

`func NewPostCopyMessagesRequest(chatId PostSendMessageRequestChatId, fromChatId PostForwardMessagesRequestFromChatId, messageIds []int32, ) *PostCopyMessagesRequest`

NewPostCopyMessagesRequest instantiates a new PostCopyMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCopyMessagesRequestWithDefaults

`func NewPostCopyMessagesRequestWithDefaults() *PostCopyMessagesRequest`

NewPostCopyMessagesRequestWithDefaults instantiates a new PostCopyMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostCopyMessagesRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostCopyMessagesRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostCopyMessagesRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *PostCopyMessagesRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostCopyMessagesRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostCopyMessagesRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostCopyMessagesRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *PostCopyMessagesRequest) GetFromChatId() PostForwardMessagesRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *PostCopyMessagesRequest) GetFromChatIdOk() (*PostForwardMessagesRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *PostCopyMessagesRequest) SetFromChatId(v PostForwardMessagesRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetMessageIds

`func (o *PostCopyMessagesRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *PostCopyMessagesRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *PostCopyMessagesRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetDisableNotification

`func (o *PostCopyMessagesRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *PostCopyMessagesRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *PostCopyMessagesRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *PostCopyMessagesRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *PostCopyMessagesRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PostCopyMessagesRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PostCopyMessagesRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PostCopyMessagesRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetRemoveCaption

`func (o *PostCopyMessagesRequest) GetRemoveCaption() bool`

GetRemoveCaption returns the RemoveCaption field if non-nil, zero value otherwise.

### GetRemoveCaptionOk

`func (o *PostCopyMessagesRequest) GetRemoveCaptionOk() (*bool, bool)`

GetRemoveCaptionOk returns a tuple with the RemoveCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveCaption

`func (o *PostCopyMessagesRequest) SetRemoveCaption(v bool)`

SetRemoveCaption sets RemoveCaption field to given value.

### HasRemoveCaption

`func (o *PostCopyMessagesRequest) HasRemoveCaption() bool`

HasRemoveCaption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


