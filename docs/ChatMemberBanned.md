# ChatMemberBanned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “kicked” | [default to "kicked"]
**User** | [**User**](User.md) |  | 
**UntilDate** | **int32** | Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever | 

## Methods

### NewChatMemberBanned

`func NewChatMemberBanned(status string, user User, untilDate int32, ) *ChatMemberBanned`

NewChatMemberBanned instantiates a new ChatMemberBanned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberBannedWithDefaults

`func NewChatMemberBannedWithDefaults() *ChatMemberBanned`

NewChatMemberBannedWithDefaults instantiates a new ChatMemberBanned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberBanned) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberBanned) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberBanned) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberBanned) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberBanned) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberBanned) SetUser(v User)`

SetUser sets User field to given value.


### GetUntilDate

`func (o *ChatMemberBanned) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *ChatMemberBanned) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *ChatMemberBanned) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


