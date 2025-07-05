# SetChatMenuButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **int32** | Unique identifier for the target private chat. If not specified, default bot&#39;s menu button will be changed | [optional] 
**MenuButton** | Pointer to [**MenuButton**](MenuButton.md) |  | [optional] 

## Methods

### NewSetChatMenuButtonRequest

`func NewSetChatMenuButtonRequest() *SetChatMenuButtonRequest`

NewSetChatMenuButtonRequest instantiates a new SetChatMenuButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatMenuButtonRequestWithDefaults

`func NewSetChatMenuButtonRequestWithDefaults() *SetChatMenuButtonRequest`

NewSetChatMenuButtonRequestWithDefaults instantiates a new SetChatMenuButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatMenuButtonRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatMenuButtonRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatMenuButtonRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SetChatMenuButtonRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMenuButton

`func (o *SetChatMenuButtonRequest) GetMenuButton() MenuButton`

GetMenuButton returns the MenuButton field if non-nil, zero value otherwise.

### GetMenuButtonOk

`func (o *SetChatMenuButtonRequest) GetMenuButtonOk() (*MenuButton, bool)`

GetMenuButtonOk returns a tuple with the MenuButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuButton

`func (o *SetChatMenuButtonRequest) SetMenuButton(v MenuButton)`

SetMenuButton sets MenuButton field to given value.

### HasMenuButton

`func (o *SetChatMenuButtonRequest) HasMenuButton() bool`

HasMenuButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


