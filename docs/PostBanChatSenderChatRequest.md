# PostBanChatSenderChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**SenderChatId** | **int32** | Unique identifier of the target sender chat | 

## Methods

### NewPostBanChatSenderChatRequest

`func NewPostBanChatSenderChatRequest(chatId PostSendMessageRequestChatId, senderChatId int32, ) *PostBanChatSenderChatRequest`

NewPostBanChatSenderChatRequest instantiates a new PostBanChatSenderChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBanChatSenderChatRequestWithDefaults

`func NewPostBanChatSenderChatRequestWithDefaults() *PostBanChatSenderChatRequest`

NewPostBanChatSenderChatRequestWithDefaults instantiates a new PostBanChatSenderChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostBanChatSenderChatRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostBanChatSenderChatRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostBanChatSenderChatRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetSenderChatId

`func (o *PostBanChatSenderChatRequest) GetSenderChatId() int32`

GetSenderChatId returns the SenderChatId field if non-nil, zero value otherwise.

### GetSenderChatIdOk

`func (o *PostBanChatSenderChatRequest) GetSenderChatIdOk() (*int32, bool)`

GetSenderChatIdOk returns a tuple with the SenderChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderChatId

`func (o *PostBanChatSenderChatRequest) SetSenderChatId(v int32)`

SetSenderChatId sets SenderChatId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


