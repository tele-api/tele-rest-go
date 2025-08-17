# ApproveSuggestedPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | **int32** | Unique identifier for the target direct messages chat | 
**MessageId** | **int32** | Identifier of a suggested post message to approve | 
**SendDate** | Pointer to **int32** | Point in time (Unix timestamp) when the post is expected to be published; omit if the date has already been specified when the suggested post was created. If specified, then the date must be not more than 2678400 seconds (30 days) in the future | [optional] 

## Methods

### NewApproveSuggestedPostRequest

`func NewApproveSuggestedPostRequest(chatId int32, messageId int32, ) *ApproveSuggestedPostRequest`

NewApproveSuggestedPostRequest instantiates a new ApproveSuggestedPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveSuggestedPostRequestWithDefaults

`func NewApproveSuggestedPostRequestWithDefaults() *ApproveSuggestedPostRequest`

NewApproveSuggestedPostRequestWithDefaults instantiates a new ApproveSuggestedPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ApproveSuggestedPostRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ApproveSuggestedPostRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ApproveSuggestedPostRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *ApproveSuggestedPostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ApproveSuggestedPostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ApproveSuggestedPostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetSendDate

`func (o *ApproveSuggestedPostRequest) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ApproveSuggestedPostRequest) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ApproveSuggestedPostRequest) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ApproveSuggestedPostRequest) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


