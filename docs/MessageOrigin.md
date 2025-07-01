# MessageOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the message origin, always “channel” | [default to "channel"]
**Date** | **int32** | Date the message was sent originally in Unix time | 
**SenderUser** | [**User**](User.md) |  | 
**SenderUserName** | **string** | Name of the user that sent the message originally | 
**SenderChat** | [**Chat**](Chat.md) |  | 
**AuthorSignature** | Pointer to **string** | *Optional*. Signature of the original post author | [optional] 
**Chat** | [**Chat**](Chat.md) |  | 
**MessageId** | **int32** | Unique message identifier inside the chat | 

## Methods

### NewMessageOrigin

`func NewMessageOrigin(type_ string, date int32, senderUser User, senderUserName string, senderChat Chat, chat Chat, messageId int32, ) *MessageOrigin`

NewMessageOrigin instantiates a new MessageOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOriginWithDefaults

`func NewMessageOriginWithDefaults() *MessageOrigin`

NewMessageOriginWithDefaults instantiates a new MessageOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOrigin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOrigin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOrigin) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *MessageOrigin) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageOrigin) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageOrigin) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSenderUser

`func (o *MessageOrigin) GetSenderUser() User`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *MessageOrigin) GetSenderUserOk() (*User, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *MessageOrigin) SetSenderUser(v User)`

SetSenderUser sets SenderUser field to given value.


### GetSenderUserName

`func (o *MessageOrigin) GetSenderUserName() string`

GetSenderUserName returns the SenderUserName field if non-nil, zero value otherwise.

### GetSenderUserNameOk

`func (o *MessageOrigin) GetSenderUserNameOk() (*string, bool)`

GetSenderUserNameOk returns a tuple with the SenderUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserName

`func (o *MessageOrigin) SetSenderUserName(v string)`

SetSenderUserName sets SenderUserName field to given value.


### GetSenderChat

`func (o *MessageOrigin) GetSenderChat() Chat`

GetSenderChat returns the SenderChat field if non-nil, zero value otherwise.

### GetSenderChatOk

`func (o *MessageOrigin) GetSenderChatOk() (*Chat, bool)`

GetSenderChatOk returns a tuple with the SenderChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderChat

`func (o *MessageOrigin) SetSenderChat(v Chat)`

SetSenderChat sets SenderChat field to given value.


### GetAuthorSignature

`func (o *MessageOrigin) GetAuthorSignature() string`

GetAuthorSignature returns the AuthorSignature field if non-nil, zero value otherwise.

### GetAuthorSignatureOk

`func (o *MessageOrigin) GetAuthorSignatureOk() (*string, bool)`

GetAuthorSignatureOk returns a tuple with the AuthorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorSignature

`func (o *MessageOrigin) SetAuthorSignature(v string)`

SetAuthorSignature sets AuthorSignature field to given value.

### HasAuthorSignature

`func (o *MessageOrigin) HasAuthorSignature() bool`

HasAuthorSignature returns a boolean if a field has been set.

### GetChat

`func (o *MessageOrigin) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MessageOrigin) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MessageOrigin) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageId

`func (o *MessageOrigin) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageOrigin) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageOrigin) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


