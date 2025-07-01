# ChatMemberOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “creator” | [default to "creator"]
**User** | [**User**](User.md) |  | 
**IsAnonymous** | **bool** | *True*, if the user&#39;s presence in the chat is hidden | 
**CustomTitle** | Pointer to **string** | *Optional*. Custom title for this user | [optional] 

## Methods

### NewChatMemberOwner

`func NewChatMemberOwner(status string, user User, isAnonymous bool, ) *ChatMemberOwner`

NewChatMemberOwner instantiates a new ChatMemberOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberOwnerWithDefaults

`func NewChatMemberOwnerWithDefaults() *ChatMemberOwner`

NewChatMemberOwnerWithDefaults instantiates a new ChatMemberOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberOwner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberOwner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberOwner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberOwner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberOwner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberOwner) SetUser(v User)`

SetUser sets User field to given value.


### GetIsAnonymous

`func (o *ChatMemberOwner) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *ChatMemberOwner) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *ChatMemberOwner) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.


### GetCustomTitle

`func (o *ChatMemberOwner) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *ChatMemberOwner) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *ChatMemberOwner) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *ChatMemberOwner) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


