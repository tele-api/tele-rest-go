# SetChatDescriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**Description** | Pointer to **string** | New chat description, 0-255 characters | [optional] 

## Methods

### NewSetChatDescriptionRequest

`func NewSetChatDescriptionRequest(chatId SendMessageRequestChatId, ) *SetChatDescriptionRequest`

NewSetChatDescriptionRequest instantiates a new SetChatDescriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatDescriptionRequestWithDefaults

`func NewSetChatDescriptionRequestWithDefaults() *SetChatDescriptionRequest`

NewSetChatDescriptionRequestWithDefaults instantiates a new SetChatDescriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatDescriptionRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatDescriptionRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatDescriptionRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetDescription

`func (o *SetChatDescriptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetChatDescriptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetChatDescriptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetChatDescriptionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


