# PostBanChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostBanChatMemberRequestChatId**](PostBanChatMemberRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**UntilDate** | Pointer to **int32** | Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only. | [optional] 
**RevokeMessages** | Pointer to **bool** | Pass *True* to delete all messages from the chat for the user that is being removed. If *False*, the user will be able to see messages in the group that were sent before the user was removed. Always *True* for supergroups and channels. | [optional] 

## Methods

### NewPostBanChatMemberRequest

`func NewPostBanChatMemberRequest(chatId PostBanChatMemberRequestChatId, userId int32, ) *PostBanChatMemberRequest`

NewPostBanChatMemberRequest instantiates a new PostBanChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBanChatMemberRequestWithDefaults

`func NewPostBanChatMemberRequestWithDefaults() *PostBanChatMemberRequest`

NewPostBanChatMemberRequestWithDefaults instantiates a new PostBanChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostBanChatMemberRequest) GetChatId() PostBanChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostBanChatMemberRequest) GetChatIdOk() (*PostBanChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostBanChatMemberRequest) SetChatId(v PostBanChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *PostBanChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostBanChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostBanChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUntilDate

`func (o *PostBanChatMemberRequest) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *PostBanChatMemberRequest) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *PostBanChatMemberRequest) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *PostBanChatMemberRequest) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### GetRevokeMessages

`func (o *PostBanChatMemberRequest) GetRevokeMessages() bool`

GetRevokeMessages returns the RevokeMessages field if non-nil, zero value otherwise.

### GetRevokeMessagesOk

`func (o *PostBanChatMemberRequest) GetRevokeMessagesOk() (*bool, bool)`

GetRevokeMessagesOk returns a tuple with the RevokeMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeMessages

`func (o *PostBanChatMemberRequest) SetRevokeMessages(v bool)`

SetRevokeMessages sets RevokeMessages field to given value.

### HasRevokeMessages

`func (o *PostBanChatMemberRequest) HasRevokeMessages() bool`

HasRevokeMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


