# SetChatAdministratorCustomTitlePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**RestrictChatMemberPostRequestChatId**](RestrictChatMemberPostRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**CustomTitle** | **string** | New custom title for the administrator; 0-16 characters, emoji are not allowed | 

## Methods

### NewSetChatAdministratorCustomTitlePostRequest

`func NewSetChatAdministratorCustomTitlePostRequest(chatId RestrictChatMemberPostRequestChatId, userId int32, customTitle string, ) *SetChatAdministratorCustomTitlePostRequest`

NewSetChatAdministratorCustomTitlePostRequest instantiates a new SetChatAdministratorCustomTitlePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatAdministratorCustomTitlePostRequestWithDefaults

`func NewSetChatAdministratorCustomTitlePostRequestWithDefaults() *SetChatAdministratorCustomTitlePostRequest`

NewSetChatAdministratorCustomTitlePostRequestWithDefaults instantiates a new SetChatAdministratorCustomTitlePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatAdministratorCustomTitlePostRequest) GetChatId() RestrictChatMemberPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatAdministratorCustomTitlePostRequest) GetChatIdOk() (*RestrictChatMemberPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatAdministratorCustomTitlePostRequest) SetChatId(v RestrictChatMemberPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *SetChatAdministratorCustomTitlePostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SetChatAdministratorCustomTitlePostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SetChatAdministratorCustomTitlePostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCustomTitle

`func (o *SetChatAdministratorCustomTitlePostRequest) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *SetChatAdministratorCustomTitlePostRequest) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *SetChatAdministratorCustomTitlePostRequest) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


