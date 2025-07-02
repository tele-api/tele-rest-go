# PostRestrictChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostRestrictChatMemberRequestChatId**](PostRestrictChatMemberRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**Permissions** | [**ChatPermissions**](ChatPermissions.md) |  | 
**UseIndependentChatPermissions** | Pointer to **bool** | Pass *True* if chat permissions are set independently. Otherwise, the *can\\_send\\_other\\_messages* and *can\\_add\\_web\\_page\\_previews* permissions will imply the *can\\_send\\_messages*, *can\\_send\\_audios*, *can\\_send\\_documents*, *can\\_send\\_photos*, *can\\_send\\_videos*, *can\\_send\\_video\\_notes*, and *can\\_send\\_voice\\_notes* permissions; the *can\\_send\\_polls* permission will imply the *can\\_send\\_messages* permission. | [optional] 
**UntilDate** | Pointer to **int32** | Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever | [optional] 

## Methods

### NewPostRestrictChatMemberRequest

`func NewPostRestrictChatMemberRequest(chatId PostRestrictChatMemberRequestChatId, userId int32, permissions ChatPermissions, ) *PostRestrictChatMemberRequest`

NewPostRestrictChatMemberRequest instantiates a new PostRestrictChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRestrictChatMemberRequestWithDefaults

`func NewPostRestrictChatMemberRequestWithDefaults() *PostRestrictChatMemberRequest`

NewPostRestrictChatMemberRequestWithDefaults instantiates a new PostRestrictChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostRestrictChatMemberRequest) GetChatId() PostRestrictChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostRestrictChatMemberRequest) GetChatIdOk() (*PostRestrictChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostRestrictChatMemberRequest) SetChatId(v PostRestrictChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *PostRestrictChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostRestrictChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostRestrictChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetPermissions

`func (o *PostRestrictChatMemberRequest) GetPermissions() ChatPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PostRestrictChatMemberRequest) GetPermissionsOk() (*ChatPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PostRestrictChatMemberRequest) SetPermissions(v ChatPermissions)`

SetPermissions sets Permissions field to given value.


### GetUseIndependentChatPermissions

`func (o *PostRestrictChatMemberRequest) GetUseIndependentChatPermissions() bool`

GetUseIndependentChatPermissions returns the UseIndependentChatPermissions field if non-nil, zero value otherwise.

### GetUseIndependentChatPermissionsOk

`func (o *PostRestrictChatMemberRequest) GetUseIndependentChatPermissionsOk() (*bool, bool)`

GetUseIndependentChatPermissionsOk returns a tuple with the UseIndependentChatPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndependentChatPermissions

`func (o *PostRestrictChatMemberRequest) SetUseIndependentChatPermissions(v bool)`

SetUseIndependentChatPermissions sets UseIndependentChatPermissions field to given value.

### HasUseIndependentChatPermissions

`func (o *PostRestrictChatMemberRequest) HasUseIndependentChatPermissions() bool`

HasUseIndependentChatPermissions returns a boolean if a field has been set.

### GetUntilDate

`func (o *PostRestrictChatMemberRequest) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *PostRestrictChatMemberRequest) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *PostRestrictChatMemberRequest) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *PostRestrictChatMemberRequest) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


