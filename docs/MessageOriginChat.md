# MessageOriginChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the message origin, always “chat” | [default to "chat"]
**Date** | **int32** | Date the message was sent originally in Unix time | 
**SenderChat** | [**Chat**](Chat.md) |  | 
**AuthorSignature** | Pointer to **string** | *Optional*. For messages originally sent by an anonymous chat administrator, original message author signature | [optional] 

## Methods

### NewMessageOriginChat

`func NewMessageOriginChat(type_ string, date int32, senderChat Chat, ) *MessageOriginChat`

NewMessageOriginChat instantiates a new MessageOriginChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOriginChatWithDefaults

`func NewMessageOriginChatWithDefaults() *MessageOriginChat`

NewMessageOriginChatWithDefaults instantiates a new MessageOriginChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOriginChat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOriginChat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOriginChat) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *MessageOriginChat) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageOriginChat) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageOriginChat) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSenderChat

`func (o *MessageOriginChat) GetSenderChat() Chat`

GetSenderChat returns the SenderChat field if non-nil, zero value otherwise.

### GetSenderChatOk

`func (o *MessageOriginChat) GetSenderChatOk() (*Chat, bool)`

GetSenderChatOk returns a tuple with the SenderChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderChat

`func (o *MessageOriginChat) SetSenderChat(v Chat)`

SetSenderChat sets SenderChat field to given value.


### GetAuthorSignature

`func (o *MessageOriginChat) GetAuthorSignature() string`

GetAuthorSignature returns the AuthorSignature field if non-nil, zero value otherwise.

### GetAuthorSignatureOk

`func (o *MessageOriginChat) GetAuthorSignatureOk() (*string, bool)`

GetAuthorSignatureOk returns a tuple with the AuthorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorSignature

`func (o *MessageOriginChat) SetAuthorSignature(v string)`

SetAuthorSignature sets AuthorSignature field to given value.

### HasAuthorSignature

`func (o *MessageOriginChat) HasAuthorSignature() bool`

HasAuthorSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


