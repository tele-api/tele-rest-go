# SetChatTitleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**Title** | **string** | New chat title, 1-128 characters | 

## Methods

### NewSetChatTitleRequest

`func NewSetChatTitleRequest(chatId SendMessageRequestChatId, title string, ) *SetChatTitleRequest`

NewSetChatTitleRequest instantiates a new SetChatTitleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatTitleRequestWithDefaults

`func NewSetChatTitleRequestWithDefaults() *SetChatTitleRequest`

NewSetChatTitleRequestWithDefaults instantiates a new SetChatTitleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatTitleRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatTitleRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatTitleRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetTitle

`func (o *SetChatTitleRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SetChatTitleRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SetChatTitleRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


