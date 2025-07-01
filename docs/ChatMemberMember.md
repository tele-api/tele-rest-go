# ChatMemberMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “member” | [default to "member"]
**User** | [**User**](User.md) |  | 
**UntilDate** | Pointer to **int32** | *Optional*. Date when the user&#39;s subscription will expire; Unix time | [optional] 

## Methods

### NewChatMemberMember

`func NewChatMemberMember(status string, user User, ) *ChatMemberMember`

NewChatMemberMember instantiates a new ChatMemberMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberMemberWithDefaults

`func NewChatMemberMemberWithDefaults() *ChatMemberMember`

NewChatMemberMemberWithDefaults instantiates a new ChatMemberMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberMember) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberMember) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberMember) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberMember) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberMember) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberMember) SetUser(v User)`

SetUser sets User field to given value.


### GetUntilDate

`func (o *ChatMemberMember) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *ChatMemberMember) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *ChatMemberMember) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *ChatMemberMember) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


