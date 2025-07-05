# GetGameHighScoresRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Target user id | 
**ChatId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Unique identifier for the target chat | [optional] 
**MessageId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Identifier of the sent message | [optional] 
**InlineMessageId** | Pointer to **string** | Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message | [optional] 

## Methods

### NewGetGameHighScoresRequest

`func NewGetGameHighScoresRequest(userId int32, ) *GetGameHighScoresRequest`

NewGetGameHighScoresRequest instantiates a new GetGameHighScoresRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameHighScoresRequestWithDefaults

`func NewGetGameHighScoresRequestWithDefaults() *GetGameHighScoresRequest`

NewGetGameHighScoresRequestWithDefaults instantiates a new GetGameHighScoresRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetGameHighScoresRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetGameHighScoresRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetGameHighScoresRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetChatId

`func (o *GetGameHighScoresRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *GetGameHighScoresRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *GetGameHighScoresRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *GetGameHighScoresRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *GetGameHighScoresRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *GetGameHighScoresRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *GetGameHighScoresRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *GetGameHighScoresRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *GetGameHighScoresRequest) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *GetGameHighScoresRequest) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *GetGameHighScoresRequest) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *GetGameHighScoresRequest) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


