# CopyMessagesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**FromChatId** | [**ForwardMessagesPostRequestFromChatId**](ForwardMessagesPostRequestFromChatId.md) |  | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages in the chat *from\\_chat\\_id* to copy. The identifiers must be specified in a strictly increasing order. | 
**DisableNotification** | Pointer to **bool** | Sends the messages [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent messages from forwarding and saving | [optional] 
**RemoveCaption** | Pointer to **bool** | Pass *True* to copy the messages without their captions | [optional] 

## Methods

### NewCopyMessagesPostRequest

`func NewCopyMessagesPostRequest(chatId SendMessagePostRequestChatId, fromChatId ForwardMessagesPostRequestFromChatId, messageIds []int32, ) *CopyMessagesPostRequest`

NewCopyMessagesPostRequest instantiates a new CopyMessagesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyMessagesPostRequestWithDefaults

`func NewCopyMessagesPostRequestWithDefaults() *CopyMessagesPostRequest`

NewCopyMessagesPostRequestWithDefaults instantiates a new CopyMessagesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CopyMessagesPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CopyMessagesPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CopyMessagesPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *CopyMessagesPostRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *CopyMessagesPostRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *CopyMessagesPostRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *CopyMessagesPostRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFromChatId

`func (o *CopyMessagesPostRequest) GetFromChatId() ForwardMessagesPostRequestFromChatId`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *CopyMessagesPostRequest) GetFromChatIdOk() (*ForwardMessagesPostRequestFromChatId, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *CopyMessagesPostRequest) SetFromChatId(v ForwardMessagesPostRequestFromChatId)`

SetFromChatId sets FromChatId field to given value.


### GetMessageIds

`func (o *CopyMessagesPostRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *CopyMessagesPostRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *CopyMessagesPostRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetDisableNotification

`func (o *CopyMessagesPostRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *CopyMessagesPostRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *CopyMessagesPostRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *CopyMessagesPostRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *CopyMessagesPostRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *CopyMessagesPostRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *CopyMessagesPostRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *CopyMessagesPostRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetRemoveCaption

`func (o *CopyMessagesPostRequest) GetRemoveCaption() bool`

GetRemoveCaption returns the RemoveCaption field if non-nil, zero value otherwise.

### GetRemoveCaptionOk

`func (o *CopyMessagesPostRequest) GetRemoveCaptionOk() (*bool, bool)`

GetRemoveCaptionOk returns a tuple with the RemoveCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveCaption

`func (o *CopyMessagesPostRequest) SetRemoveCaption(v bool)`

SetRemoveCaption sets RemoveCaption field to given value.

### HasRemoveCaption

`func (o *CopyMessagesPostRequest) HasRemoveCaption() bool`

HasRemoveCaption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


