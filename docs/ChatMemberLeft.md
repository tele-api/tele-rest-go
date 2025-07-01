# ChatMemberLeft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “left” | [default to "left"]
**User** | [**User**](User.md) |  | 

## Methods

### NewChatMemberLeft

`func NewChatMemberLeft(status string, user User, ) *ChatMemberLeft`

NewChatMemberLeft instantiates a new ChatMemberLeft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberLeftWithDefaults

`func NewChatMemberLeftWithDefaults() *ChatMemberLeft`

NewChatMemberLeftWithDefaults instantiates a new ChatMemberLeft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberLeft) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberLeft) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberLeft) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberLeft) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberLeft) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberLeft) SetUser(v User)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


