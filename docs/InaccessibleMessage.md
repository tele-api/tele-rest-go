# InaccessibleMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**MessageId** | **int32** | Unique message identifier inside the chat | 
**Date** | **int32** | Always 0. The field can be used to differentiate regular and inaccessible messages. | 

## Methods

### NewInaccessibleMessage

`func NewInaccessibleMessage(chat Chat, messageId int32, date int32, ) *InaccessibleMessage`

NewInaccessibleMessage instantiates a new InaccessibleMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInaccessibleMessageWithDefaults

`func NewInaccessibleMessageWithDefaults() *InaccessibleMessage`

NewInaccessibleMessageWithDefaults instantiates a new InaccessibleMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *InaccessibleMessage) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *InaccessibleMessage) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *InaccessibleMessage) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageId

`func (o *InaccessibleMessage) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *InaccessibleMessage) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *InaccessibleMessage) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetDate

`func (o *InaccessibleMessage) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InaccessibleMessage) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InaccessibleMessage) SetDate(v int32)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


