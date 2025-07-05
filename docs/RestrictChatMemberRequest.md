# RestrictChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**Permissions** | [**ChatPermissions**](ChatPermissions.md) |  | 
**UseIndependentChatPermissions** | Pointer to **bool** | Pass *True* if chat permissions are set independently. Otherwise, the *can\\_send\\_other\\_messages* and *can\\_add\\_web\\_page\\_previews* permissions will imply the *can\\_send\\_messages*, *can\\_send\\_audios*, *can\\_send\\_documents*, *can\\_send\\_photos*, *can\\_send\\_videos*, *can\\_send\\_video\\_notes*, and *can\\_send\\_voice\\_notes* permissions; the *can\\_send\\_polls* permission will imply the *can\\_send\\_messages* permission. | [optional] 
**UntilDate** | Pointer to **int32** | Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever | [optional] 

## Methods

### NewRestrictChatMemberRequest

`func NewRestrictChatMemberRequest(chatId BotCommandScopeChatChatId, userId int32, permissions ChatPermissions, ) *RestrictChatMemberRequest`

NewRestrictChatMemberRequest instantiates a new RestrictChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictChatMemberRequestWithDefaults

`func NewRestrictChatMemberRequestWithDefaults() *RestrictChatMemberRequest`

NewRestrictChatMemberRequestWithDefaults instantiates a new RestrictChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *RestrictChatMemberRequest) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *RestrictChatMemberRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *RestrictChatMemberRequest) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *RestrictChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RestrictChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RestrictChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetPermissions

`func (o *RestrictChatMemberRequest) GetPermissions() ChatPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RestrictChatMemberRequest) GetPermissionsOk() (*ChatPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RestrictChatMemberRequest) SetPermissions(v ChatPermissions)`

SetPermissions sets Permissions field to given value.


### GetUseIndependentChatPermissions

`func (o *RestrictChatMemberRequest) GetUseIndependentChatPermissions() bool`

GetUseIndependentChatPermissions returns the UseIndependentChatPermissions field if non-nil, zero value otherwise.

### GetUseIndependentChatPermissionsOk

`func (o *RestrictChatMemberRequest) GetUseIndependentChatPermissionsOk() (*bool, bool)`

GetUseIndependentChatPermissionsOk returns a tuple with the UseIndependentChatPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndependentChatPermissions

`func (o *RestrictChatMemberRequest) SetUseIndependentChatPermissions(v bool)`

SetUseIndependentChatPermissions sets UseIndependentChatPermissions field to given value.

### HasUseIndependentChatPermissions

`func (o *RestrictChatMemberRequest) HasUseIndependentChatPermissions() bool`

HasUseIndependentChatPermissions returns a boolean if a field has been set.

### GetUntilDate

`func (o *RestrictChatMemberRequest) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *RestrictChatMemberRequest) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *RestrictChatMemberRequest) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *RestrictChatMemberRequest) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


