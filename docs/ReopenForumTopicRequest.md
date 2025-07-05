# ReopenForumTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**MessageThreadId** | **int32** | Unique identifier for the target message thread of the forum topic | 

## Methods

### NewReopenForumTopicRequest

`func NewReopenForumTopicRequest(chatId BotCommandScopeChatChatId, messageThreadId int32, ) *ReopenForumTopicRequest`

NewReopenForumTopicRequest instantiates a new ReopenForumTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReopenForumTopicRequestWithDefaults

`func NewReopenForumTopicRequestWithDefaults() *ReopenForumTopicRequest`

NewReopenForumTopicRequestWithDefaults instantiates a new ReopenForumTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ReopenForumTopicRequest) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ReopenForumTopicRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ReopenForumTopicRequest) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *ReopenForumTopicRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *ReopenForumTopicRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *ReopenForumTopicRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


