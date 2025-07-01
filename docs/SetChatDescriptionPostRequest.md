# SetChatDescriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**Description** | Pointer to **string** | New chat description, 0-255 characters | [optional] 

## Methods

### NewSetChatDescriptionPostRequest

`func NewSetChatDescriptionPostRequest(chatId SendMessagePostRequestChatId, ) *SetChatDescriptionPostRequest`

NewSetChatDescriptionPostRequest instantiates a new SetChatDescriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatDescriptionPostRequestWithDefaults

`func NewSetChatDescriptionPostRequestWithDefaults() *SetChatDescriptionPostRequest`

NewSetChatDescriptionPostRequestWithDefaults instantiates a new SetChatDescriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatDescriptionPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatDescriptionPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatDescriptionPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetDescription

`func (o *SetChatDescriptionPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetChatDescriptionPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetChatDescriptionPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetChatDescriptionPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


