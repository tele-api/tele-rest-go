# SavePreparedInlineMessagePostRequest

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

### NewSavePreparedInlineMessagePostRequest

`func NewSavePreparedInlineMessagePostRequest(userId int32, result InlineQueryResult, ) *SavePreparedInlineMessagePostRequest`

NewSavePreparedInlineMessagePostRequest instantiates a new SavePreparedInlineMessagePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavePreparedInlineMessagePostRequestWithDefaults

`func NewSavePreparedInlineMessagePostRequestWithDefaults() *SavePreparedInlineMessagePostRequest`

NewSavePreparedInlineMessagePostRequestWithDefaults instantiates a new SavePreparedInlineMessagePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SavePreparedInlineMessagePostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SavePreparedInlineMessagePostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SavePreparedInlineMessagePostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetResult

`func (o *SavePreparedInlineMessagePostRequest) GetResult() InlineQueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SavePreparedInlineMessagePostRequest) GetResultOk() (*InlineQueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SavePreparedInlineMessagePostRequest) SetResult(v InlineQueryResult)`

SetResult sets Result field to given value.


### GetAllowUserChats

`func (o *SavePreparedInlineMessagePostRequest) GetAllowUserChats() bool`

GetAllowUserChats returns the AllowUserChats field if non-nil, zero value otherwise.

### GetAllowUserChatsOk

`func (o *SavePreparedInlineMessagePostRequest) GetAllowUserChatsOk() (*bool, bool)`

GetAllowUserChatsOk returns a tuple with the AllowUserChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserChats

`func (o *SavePreparedInlineMessagePostRequest) SetAllowUserChats(v bool)`

SetAllowUserChats sets AllowUserChats field to given value.

### HasAllowUserChats

`func (o *SavePreparedInlineMessagePostRequest) HasAllowUserChats() bool`

HasAllowUserChats returns a boolean if a field has been set.

### GetAllowBotChats

`func (o *SavePreparedInlineMessagePostRequest) GetAllowBotChats() bool`

GetAllowBotChats returns the AllowBotChats field if non-nil, zero value otherwise.

### GetAllowBotChatsOk

`func (o *SavePreparedInlineMessagePostRequest) GetAllowBotChatsOk() (*bool, bool)`

GetAllowBotChatsOk returns a tuple with the AllowBotChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBotChats

`func (o *SavePreparedInlineMessagePostRequest) SetAllowBotChats(v bool)`

SetAllowBotChats sets AllowBotChats field to given value.

### HasAllowBotChats

`func (o *SavePreparedInlineMessagePostRequest) HasAllowBotChats() bool`

HasAllowBotChats returns a boolean if a field has been set.

### GetAllowGroupChats

`func (o *SavePreparedInlineMessagePostRequest) GetAllowGroupChats() bool`

GetAllowGroupChats returns the AllowGroupChats field if non-nil, zero value otherwise.

### GetAllowGroupChatsOk

`func (o *SavePreparedInlineMessagePostRequest) GetAllowGroupChatsOk() (*bool, bool)`

GetAllowGroupChatsOk returns a tuple with the AllowGroupChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGroupChats

`func (o *SavePreparedInlineMessagePostRequest) SetAllowGroupChats(v bool)`

SetAllowGroupChats sets AllowGroupChats field to given value.

### HasAllowGroupChats

`func (o *SavePreparedInlineMessagePostRequest) HasAllowGroupChats() bool`

HasAllowGroupChats returns a boolean if a field has been set.

### GetAllowChannelChats

`func (o *SavePreparedInlineMessagePostRequest) GetAllowChannelChats() bool`

GetAllowChannelChats returns the AllowChannelChats field if non-nil, zero value otherwise.

### GetAllowChannelChatsOk

`func (o *SavePreparedInlineMessagePostRequest) GetAllowChannelChatsOk() (*bool, bool)`

GetAllowChannelChatsOk returns a tuple with the AllowChannelChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChannelChats

`func (o *SavePreparedInlineMessagePostRequest) SetAllowChannelChats(v bool)`

SetAllowChannelChats sets AllowChannelChats field to given value.

### HasAllowChannelChats

`func (o *SavePreparedInlineMessagePostRequest) HasAllowChannelChats() bool`

HasAllowChannelChats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


