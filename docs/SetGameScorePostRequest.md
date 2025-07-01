# SetGameScorePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier | 
**Score** | **int32** | New score, must be non-negative | 
**Force** | Pointer to **bool** | Pass *True* if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters | [optional] 
**DisableEditMessage** | Pointer to **bool** | Pass *True* if the game message should not be automatically edited to include the current scoreboard | [optional] 
**ChatId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Unique identifier for the target chat | [optional] 
**MessageId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Identifier of the sent message | [optional] 
**InlineMessageId** | Pointer to **string** | Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message | [optional] 

## Methods

### NewSetGameScorePostRequest

`func NewSetGameScorePostRequest(userId int32, score int32, ) *SetGameScorePostRequest`

NewSetGameScorePostRequest instantiates a new SetGameScorePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGameScorePostRequestWithDefaults

`func NewSetGameScorePostRequestWithDefaults() *SetGameScorePostRequest`

NewSetGameScorePostRequestWithDefaults instantiates a new SetGameScorePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SetGameScorePostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SetGameScorePostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SetGameScorePostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetScore

`func (o *SetGameScorePostRequest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SetGameScorePostRequest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SetGameScorePostRequest) SetScore(v int32)`

SetScore sets Score field to given value.


### GetForce

`func (o *SetGameScorePostRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *SetGameScorePostRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *SetGameScorePostRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *SetGameScorePostRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetDisableEditMessage

`func (o *SetGameScorePostRequest) GetDisableEditMessage() bool`

GetDisableEditMessage returns the DisableEditMessage field if non-nil, zero value otherwise.

### GetDisableEditMessageOk

`func (o *SetGameScorePostRequest) GetDisableEditMessageOk() (*bool, bool)`

GetDisableEditMessageOk returns a tuple with the DisableEditMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEditMessage

`func (o *SetGameScorePostRequest) SetDisableEditMessage(v bool)`

SetDisableEditMessage sets DisableEditMessage field to given value.

### HasDisableEditMessage

`func (o *SetGameScorePostRequest) HasDisableEditMessage() bool`

HasDisableEditMessage returns a boolean if a field has been set.

### GetChatId

`func (o *SetGameScorePostRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetGameScorePostRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetGameScorePostRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SetGameScorePostRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *SetGameScorePostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SetGameScorePostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SetGameScorePostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SetGameScorePostRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *SetGameScorePostRequest) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *SetGameScorePostRequest) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *SetGameScorePostRequest) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *SetGameScorePostRequest) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


