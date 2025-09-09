# DeclineSuggestedPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | **int32** | Unique identifier for the target direct messages chat | 
**MessageId** | **int32** | Identifier of a suggested post message to decline | 
**Comment** | Pointer to **string** | Comment for the creator of the suggested post; 0-128 characters | [optional] 

## Methods

### NewDeclineSuggestedPostRequest

`func NewDeclineSuggestedPostRequest(chatId int32, messageId int32, ) *DeclineSuggestedPostRequest`

NewDeclineSuggestedPostRequest instantiates a new DeclineSuggestedPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclineSuggestedPostRequestWithDefaults

`func NewDeclineSuggestedPostRequestWithDefaults() *DeclineSuggestedPostRequest`

NewDeclineSuggestedPostRequestWithDefaults instantiates a new DeclineSuggestedPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *DeclineSuggestedPostRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *DeclineSuggestedPostRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *DeclineSuggestedPostRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *DeclineSuggestedPostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *DeclineSuggestedPostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *DeclineSuggestedPostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetComment

`func (o *DeclineSuggestedPostRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeclineSuggestedPostRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeclineSuggestedPostRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DeclineSuggestedPostRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


