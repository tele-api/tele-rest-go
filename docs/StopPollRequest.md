# StopPollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message to be edited was sent | [optional] 
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageId** | **int32** | Identifier of the original message with the poll | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewStopPollRequest

`func NewStopPollRequest(chatId SendMessageRequestChatId, messageId int32, ) *StopPollRequest`

NewStopPollRequest instantiates a new StopPollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPollRequestWithDefaults

`func NewStopPollRequestWithDefaults() *StopPollRequest`

NewStopPollRequestWithDefaults instantiates a new StopPollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *StopPollRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *StopPollRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *StopPollRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *StopPollRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *StopPollRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *StopPollRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *StopPollRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *StopPollRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *StopPollRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *StopPollRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetReplyMarkup

`func (o *StopPollRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *StopPollRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *StopPollRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *StopPollRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


