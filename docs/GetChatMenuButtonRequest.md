# GetChatMenuButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **int32** | Unique identifier for the target private chat. If not specified, default bot&#39;s menu button will be returned | [optional] 

## Methods

### NewGetChatMenuButtonRequest

`func NewGetChatMenuButtonRequest() *GetChatMenuButtonRequest`

NewGetChatMenuButtonRequest instantiates a new GetChatMenuButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChatMenuButtonRequestWithDefaults

`func NewGetChatMenuButtonRequestWithDefaults() *GetChatMenuButtonRequest`

NewGetChatMenuButtonRequestWithDefaults instantiates a new GetChatMenuButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *GetChatMenuButtonRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *GetChatMenuButtonRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *GetChatMenuButtonRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *GetChatMenuButtonRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


