# ChatMemberUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**From** | [**User**](User.md) |  | 
**Date** | **int32** | Date the change was done in Unix time | 
**OldChatMember** | [**ChatMember**](ChatMember.md) |  | 
**NewChatMember** | [**ChatMember**](ChatMember.md) |  | 
**InviteLink** | Pointer to [**ChatInviteLink**](ChatInviteLink.md) |  | [optional] 
**ViaJoinRequest** | Pointer to **bool** | *Optional*. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator | [optional] 
**ViaChatFolderInviteLink** | Pointer to **bool** | *Optional*. True, if the user joined the chat via a chat folder invite link | [optional] 

## Methods

### NewChatMemberUpdated

`func NewChatMemberUpdated(chat Chat, from User, date int32, oldChatMember ChatMember, newChatMember ChatMember, ) *ChatMemberUpdated`

NewChatMemberUpdated instantiates a new ChatMemberUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberUpdatedWithDefaults

`func NewChatMemberUpdatedWithDefaults() *ChatMemberUpdated`

NewChatMemberUpdatedWithDefaults instantiates a new ChatMemberUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatMemberUpdated) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatMemberUpdated) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatMemberUpdated) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetFrom

`func (o *ChatMemberUpdated) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChatMemberUpdated) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChatMemberUpdated) SetFrom(v User)`

SetFrom sets From field to given value.


### GetDate

`func (o *ChatMemberUpdated) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ChatMemberUpdated) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ChatMemberUpdated) SetDate(v int32)`

SetDate sets Date field to given value.


### GetOldChatMember

`func (o *ChatMemberUpdated) GetOldChatMember() ChatMember`

GetOldChatMember returns the OldChatMember field if non-nil, zero value otherwise.

### GetOldChatMemberOk

`func (o *ChatMemberUpdated) GetOldChatMemberOk() (*ChatMember, bool)`

GetOldChatMemberOk returns a tuple with the OldChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldChatMember

`func (o *ChatMemberUpdated) SetOldChatMember(v ChatMember)`

SetOldChatMember sets OldChatMember field to given value.


### GetNewChatMember

`func (o *ChatMemberUpdated) GetNewChatMember() ChatMember`

GetNewChatMember returns the NewChatMember field if non-nil, zero value otherwise.

### GetNewChatMemberOk

`func (o *ChatMemberUpdated) GetNewChatMemberOk() (*ChatMember, bool)`

GetNewChatMemberOk returns a tuple with the NewChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatMember

`func (o *ChatMemberUpdated) SetNewChatMember(v ChatMember)`

SetNewChatMember sets NewChatMember field to given value.


### GetInviteLink

`func (o *ChatMemberUpdated) GetInviteLink() ChatInviteLink`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ChatMemberUpdated) GetInviteLinkOk() (*ChatInviteLink, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ChatMemberUpdated) SetInviteLink(v ChatInviteLink)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *ChatMemberUpdated) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetViaJoinRequest

`func (o *ChatMemberUpdated) GetViaJoinRequest() bool`

GetViaJoinRequest returns the ViaJoinRequest field if non-nil, zero value otherwise.

### GetViaJoinRequestOk

`func (o *ChatMemberUpdated) GetViaJoinRequestOk() (*bool, bool)`

GetViaJoinRequestOk returns a tuple with the ViaJoinRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaJoinRequest

`func (o *ChatMemberUpdated) SetViaJoinRequest(v bool)`

SetViaJoinRequest sets ViaJoinRequest field to given value.

### HasViaJoinRequest

`func (o *ChatMemberUpdated) HasViaJoinRequest() bool`

HasViaJoinRequest returns a boolean if a field has been set.

### GetViaChatFolderInviteLink

`func (o *ChatMemberUpdated) GetViaChatFolderInviteLink() bool`

GetViaChatFolderInviteLink returns the ViaChatFolderInviteLink field if non-nil, zero value otherwise.

### GetViaChatFolderInviteLinkOk

`func (o *ChatMemberUpdated) GetViaChatFolderInviteLinkOk() (*bool, bool)`

GetViaChatFolderInviteLinkOk returns a tuple with the ViaChatFolderInviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaChatFolderInviteLink

`func (o *ChatMemberUpdated) SetViaChatFolderInviteLink(v bool)`

SetViaChatFolderInviteLink sets ViaChatFolderInviteLink field to given value.

### HasViaChatFolderInviteLink

`func (o *ChatMemberUpdated) HasViaChatFolderInviteLink() bool`

HasViaChatFolderInviteLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


