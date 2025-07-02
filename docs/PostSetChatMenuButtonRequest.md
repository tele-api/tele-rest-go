# PostSetChatMenuButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **int32** | Unique identifier for the target private chat. If not specified, default bot&#39;s menu button will be changed | [optional] 
**MenuButton** | Pointer to [**MenuButton**](MenuButton.md) |  | [optional] 

## Methods

### NewPostSetChatMenuButtonRequest

`func NewPostSetChatMenuButtonRequest() *PostSetChatMenuButtonRequest`

NewPostSetChatMenuButtonRequest instantiates a new PostSetChatMenuButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetChatMenuButtonRequestWithDefaults

`func NewPostSetChatMenuButtonRequestWithDefaults() *PostSetChatMenuButtonRequest`

NewPostSetChatMenuButtonRequestWithDefaults instantiates a new PostSetChatMenuButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostSetChatMenuButtonRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostSetChatMenuButtonRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostSetChatMenuButtonRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *PostSetChatMenuButtonRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMenuButton

`func (o *PostSetChatMenuButtonRequest) GetMenuButton() MenuButton`

GetMenuButton returns the MenuButton field if non-nil, zero value otherwise.

### GetMenuButtonOk

`func (o *PostSetChatMenuButtonRequest) GetMenuButtonOk() (*MenuButton, bool)`

GetMenuButtonOk returns a tuple with the MenuButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuButton

`func (o *PostSetChatMenuButtonRequest) SetMenuButton(v MenuButton)`

SetMenuButton sets MenuButton field to given value.

### HasMenuButton

`func (o *PostSetChatMenuButtonRequest) HasMenuButton() bool`

HasMenuButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


