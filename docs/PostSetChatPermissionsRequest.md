# PostSetChatPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostRestrictChatMemberRequestChatId**](PostRestrictChatMemberRequestChatId.md) |  | 
**Permissions** | [**ChatPermissions**](ChatPermissions.md) |  | 
**UseIndependentChatPermissions** | Pointer to **bool** | Pass *True* if chat permissions are set independently. Otherwise, the *can\\_send\\_other\\_messages* and *can\\_add\\_web\\_page\\_previews* permissions will imply the *can\\_send\\_messages*, *can\\_send\\_audios*, *can\\_send\\_documents*, *can\\_send\\_photos*, *can\\_send\\_videos*, *can\\_send\\_video\\_notes*, and *can\\_send\\_voice\\_notes* permissions; the *can\\_send\\_polls* permission will imply the *can\\_send\\_messages* permission. | [optional] 

## Methods

### NewPostSetChatPermissionsRequest

`func NewPostSetChatPermissionsRequest(chatId PostRestrictChatMemberRequestChatId, permissions ChatPermissions, ) *PostSetChatPermissionsRequest`

NewPostSetChatPermissionsRequest instantiates a new PostSetChatPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetChatPermissionsRequestWithDefaults

`func NewPostSetChatPermissionsRequestWithDefaults() *PostSetChatPermissionsRequest`

NewPostSetChatPermissionsRequestWithDefaults instantiates a new PostSetChatPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostSetChatPermissionsRequest) GetChatId() PostRestrictChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostSetChatPermissionsRequest) GetChatIdOk() (*PostRestrictChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostSetChatPermissionsRequest) SetChatId(v PostRestrictChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetPermissions

`func (o *PostSetChatPermissionsRequest) GetPermissions() ChatPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PostSetChatPermissionsRequest) GetPermissionsOk() (*ChatPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PostSetChatPermissionsRequest) SetPermissions(v ChatPermissions)`

SetPermissions sets Permissions field to given value.


### GetUseIndependentChatPermissions

`func (o *PostSetChatPermissionsRequest) GetUseIndependentChatPermissions() bool`

GetUseIndependentChatPermissions returns the UseIndependentChatPermissions field if non-nil, zero value otherwise.

### GetUseIndependentChatPermissionsOk

`func (o *PostSetChatPermissionsRequest) GetUseIndependentChatPermissionsOk() (*bool, bool)`

GetUseIndependentChatPermissionsOk returns a tuple with the UseIndependentChatPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndependentChatPermissions

`func (o *PostSetChatPermissionsRequest) SetUseIndependentChatPermissions(v bool)`

SetUseIndependentChatPermissions sets UseIndependentChatPermissions field to given value.

### HasUseIndependentChatPermissions

`func (o *PostSetChatPermissionsRequest) HasUseIndependentChatPermissions() bool`

HasUseIndependentChatPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


