# VerifyChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**CustomDescription** | Pointer to **string** | Custom description for the verification; 0-70 characters. Must be empty if the organization isn&#39;t allowed to provide a custom verification description. | [optional] 

## Methods

### NewVerifyChatRequest

`func NewVerifyChatRequest(chatId SendMessageRequestChatId, ) *VerifyChatRequest`

NewVerifyChatRequest instantiates a new VerifyChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyChatRequestWithDefaults

`func NewVerifyChatRequestWithDefaults() *VerifyChatRequest`

NewVerifyChatRequestWithDefaults instantiates a new VerifyChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *VerifyChatRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *VerifyChatRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *VerifyChatRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetCustomDescription

`func (o *VerifyChatRequest) GetCustomDescription() string`

GetCustomDescription returns the CustomDescription field if non-nil, zero value otherwise.

### GetCustomDescriptionOk

`func (o *VerifyChatRequest) GetCustomDescriptionOk() (*string, bool)`

GetCustomDescriptionOk returns a tuple with the CustomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDescription

`func (o *VerifyChatRequest) SetCustomDescription(v string)`

SetCustomDescription sets CustomDescription field to given value.

### HasCustomDescription

`func (o *VerifyChatRequest) HasCustomDescription() bool`

HasCustomDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


