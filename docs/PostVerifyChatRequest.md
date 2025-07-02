# PostVerifyChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**CustomDescription** | Pointer to **string** | Custom description for the verification; 0-70 characters. Must be empty if the organization isn&#39;t allowed to provide a custom verification description. | [optional] 

## Methods

### NewPostVerifyChatRequest

`func NewPostVerifyChatRequest(chatId PostSendMessageRequestChatId, ) *PostVerifyChatRequest`

NewPostVerifyChatRequest instantiates a new PostVerifyChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVerifyChatRequestWithDefaults

`func NewPostVerifyChatRequestWithDefaults() *PostVerifyChatRequest`

NewPostVerifyChatRequestWithDefaults instantiates a new PostVerifyChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostVerifyChatRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostVerifyChatRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostVerifyChatRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetCustomDescription

`func (o *PostVerifyChatRequest) GetCustomDescription() string`

GetCustomDescription returns the CustomDescription field if non-nil, zero value otherwise.

### GetCustomDescriptionOk

`func (o *PostVerifyChatRequest) GetCustomDescriptionOk() (*string, bool)`

GetCustomDescriptionOk returns a tuple with the CustomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDescription

`func (o *PostVerifyChatRequest) SetCustomDescription(v string)`

SetCustomDescription sets CustomDescription field to given value.

### HasCustomDescription

`func (o *PostVerifyChatRequest) HasCustomDescription() bool`

HasCustomDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


