# PostSavePreparedInlineMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user that can use the prepared message | 
**Result** | [**InlineQueryResult**](InlineQueryResult.md) |  | 
**AllowUserChats** | Pointer to **bool** | Pass *True* if the message can be sent to private chats with users | [optional] 
**AllowBotChats** | Pointer to **bool** | Pass *True* if the message can be sent to private chats with bots | [optional] 
**AllowGroupChats** | Pointer to **bool** | Pass *True* if the message can be sent to group and supergroup chats | [optional] 
**AllowChannelChats** | Pointer to **bool** | Pass *True* if the message can be sent to channel chats | [optional] 

## Methods

### NewPostSavePreparedInlineMessageRequest

`func NewPostSavePreparedInlineMessageRequest(userId int32, result InlineQueryResult, ) *PostSavePreparedInlineMessageRequest`

NewPostSavePreparedInlineMessageRequest instantiates a new PostSavePreparedInlineMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSavePreparedInlineMessageRequestWithDefaults

`func NewPostSavePreparedInlineMessageRequestWithDefaults() *PostSavePreparedInlineMessageRequest`

NewPostSavePreparedInlineMessageRequestWithDefaults instantiates a new PostSavePreparedInlineMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostSavePreparedInlineMessageRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostSavePreparedInlineMessageRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostSavePreparedInlineMessageRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetResult

`func (o *PostSavePreparedInlineMessageRequest) GetResult() InlineQueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostSavePreparedInlineMessageRequest) GetResultOk() (*InlineQueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostSavePreparedInlineMessageRequest) SetResult(v InlineQueryResult)`

SetResult sets Result field to given value.


### GetAllowUserChats

`func (o *PostSavePreparedInlineMessageRequest) GetAllowUserChats() bool`

GetAllowUserChats returns the AllowUserChats field if non-nil, zero value otherwise.

### GetAllowUserChatsOk

`func (o *PostSavePreparedInlineMessageRequest) GetAllowUserChatsOk() (*bool, bool)`

GetAllowUserChatsOk returns a tuple with the AllowUserChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserChats

`func (o *PostSavePreparedInlineMessageRequest) SetAllowUserChats(v bool)`

SetAllowUserChats sets AllowUserChats field to given value.

### HasAllowUserChats

`func (o *PostSavePreparedInlineMessageRequest) HasAllowUserChats() bool`

HasAllowUserChats returns a boolean if a field has been set.

### GetAllowBotChats

`func (o *PostSavePreparedInlineMessageRequest) GetAllowBotChats() bool`

GetAllowBotChats returns the AllowBotChats field if non-nil, zero value otherwise.

### GetAllowBotChatsOk

`func (o *PostSavePreparedInlineMessageRequest) GetAllowBotChatsOk() (*bool, bool)`

GetAllowBotChatsOk returns a tuple with the AllowBotChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBotChats

`func (o *PostSavePreparedInlineMessageRequest) SetAllowBotChats(v bool)`

SetAllowBotChats sets AllowBotChats field to given value.

### HasAllowBotChats

`func (o *PostSavePreparedInlineMessageRequest) HasAllowBotChats() bool`

HasAllowBotChats returns a boolean if a field has been set.

### GetAllowGroupChats

`func (o *PostSavePreparedInlineMessageRequest) GetAllowGroupChats() bool`

GetAllowGroupChats returns the AllowGroupChats field if non-nil, zero value otherwise.

### GetAllowGroupChatsOk

`func (o *PostSavePreparedInlineMessageRequest) GetAllowGroupChatsOk() (*bool, bool)`

GetAllowGroupChatsOk returns a tuple with the AllowGroupChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGroupChats

`func (o *PostSavePreparedInlineMessageRequest) SetAllowGroupChats(v bool)`

SetAllowGroupChats sets AllowGroupChats field to given value.

### HasAllowGroupChats

`func (o *PostSavePreparedInlineMessageRequest) HasAllowGroupChats() bool`

HasAllowGroupChats returns a boolean if a field has been set.

### GetAllowChannelChats

`func (o *PostSavePreparedInlineMessageRequest) GetAllowChannelChats() bool`

GetAllowChannelChats returns the AllowChannelChats field if non-nil, zero value otherwise.

### GetAllowChannelChatsOk

`func (o *PostSavePreparedInlineMessageRequest) GetAllowChannelChatsOk() (*bool, bool)`

GetAllowChannelChatsOk returns a tuple with the AllowChannelChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChannelChats

`func (o *PostSavePreparedInlineMessageRequest) SetAllowChannelChats(v bool)`

SetAllowChannelChats sets AllowChannelChats field to given value.

### HasAllowChannelChats

`func (o *PostSavePreparedInlineMessageRequest) HasAllowChannelChats() bool`

HasAllowChannelChats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


