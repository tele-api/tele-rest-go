# ChatJoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**From** | [**User**](User.md) |  | 
**UserChatId** | **int32** | Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user. | 
**Date** | **int32** | Date the request was sent in Unix time | 
**Bio** | Pointer to **string** | *Optional*. Bio of the user. | [optional] 
**InviteLink** | Pointer to [**ChatInviteLink**](ChatInviteLink.md) |  | [optional] 

## Methods

### NewChatJoinRequest

`func NewChatJoinRequest(chat Chat, from User, userChatId int32, date int32, ) *ChatJoinRequest`

NewChatJoinRequest instantiates a new ChatJoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatJoinRequestWithDefaults

`func NewChatJoinRequestWithDefaults() *ChatJoinRequest`

NewChatJoinRequestWithDefaults instantiates a new ChatJoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatJoinRequest) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatJoinRequest) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatJoinRequest) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetFrom

`func (o *ChatJoinRequest) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChatJoinRequest) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChatJoinRequest) SetFrom(v User)`

SetFrom sets From field to given value.


### GetUserChatId

`func (o *ChatJoinRequest) GetUserChatId() int32`

GetUserChatId returns the UserChatId field if non-nil, zero value otherwise.

### GetUserChatIdOk

`func (o *ChatJoinRequest) GetUserChatIdOk() (*int32, bool)`

GetUserChatIdOk returns a tuple with the UserChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserChatId

`func (o *ChatJoinRequest) SetUserChatId(v int32)`

SetUserChatId sets UserChatId field to given value.


### GetDate

`func (o *ChatJoinRequest) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ChatJoinRequest) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ChatJoinRequest) SetDate(v int32)`

SetDate sets Date field to given value.


### GetBio

`func (o *ChatJoinRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *ChatJoinRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *ChatJoinRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *ChatJoinRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetInviteLink

`func (o *ChatJoinRequest) GetInviteLink() ChatInviteLink`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ChatJoinRequest) GetInviteLinkOk() (*ChatInviteLink, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ChatJoinRequest) SetInviteLink(v ChatInviteLink)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *ChatJoinRequest) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


