# ReadBusinessMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection on behalf of which to read the message | 
**ChatId** | **int32** | Unique identifier of the chat in which the message was received. The chat must have been active in the last 24 hours. | 
**MessageId** | **int32** | Unique identifier of the message to mark as read | 

## Methods

### NewReadBusinessMessageRequest

`func NewReadBusinessMessageRequest(businessConnectionId string, chatId int32, messageId int32, ) *ReadBusinessMessageRequest`

NewReadBusinessMessageRequest instantiates a new ReadBusinessMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadBusinessMessageRequestWithDefaults

`func NewReadBusinessMessageRequestWithDefaults() *ReadBusinessMessageRequest`

NewReadBusinessMessageRequestWithDefaults instantiates a new ReadBusinessMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *ReadBusinessMessageRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *ReadBusinessMessageRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *ReadBusinessMessageRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetChatId

`func (o *ReadBusinessMessageRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ReadBusinessMessageRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ReadBusinessMessageRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *ReadBusinessMessageRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ReadBusinessMessageRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ReadBusinessMessageRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


