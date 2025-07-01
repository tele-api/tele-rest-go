# MessageOriginChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the message origin, always “channel” | [default to "channel"]
**Date** | **int32** | Date the message was sent originally in Unix time | 
**Chat** | [**Chat**](Chat.md) |  | 
**MessageId** | **int32** | Unique message identifier inside the chat | 
**AuthorSignature** | Pointer to **string** | *Optional*. Signature of the original post author | [optional] 

## Methods

### NewMessageOriginChannel

`func NewMessageOriginChannel(type_ string, date int32, chat Chat, messageId int32, ) *MessageOriginChannel`

NewMessageOriginChannel instantiates a new MessageOriginChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOriginChannelWithDefaults

`func NewMessageOriginChannelWithDefaults() *MessageOriginChannel`

NewMessageOriginChannelWithDefaults instantiates a new MessageOriginChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOriginChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOriginChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOriginChannel) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *MessageOriginChannel) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageOriginChannel) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageOriginChannel) SetDate(v int32)`

SetDate sets Date field to given value.


### GetChat

`func (o *MessageOriginChannel) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MessageOriginChannel) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MessageOriginChannel) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageId

`func (o *MessageOriginChannel) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageOriginChannel) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageOriginChannel) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetAuthorSignature

`func (o *MessageOriginChannel) GetAuthorSignature() string`

GetAuthorSignature returns the AuthorSignature field if non-nil, zero value otherwise.

### GetAuthorSignatureOk

`func (o *MessageOriginChannel) GetAuthorSignatureOk() (*string, bool)`

GetAuthorSignatureOk returns a tuple with the AuthorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorSignature

`func (o *MessageOriginChannel) SetAuthorSignature(v string)`

SetAuthorSignature sets AuthorSignature field to given value.

### HasAuthorSignature

`func (o *MessageOriginChannel) HasAuthorSignature() bool`

HasAuthorSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


