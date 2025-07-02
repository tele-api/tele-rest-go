# PostEditChatInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**InviteLink** | **string** | The invite link to edit | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 
**ExpireDate** | Pointer to **int32** | Point in time (Unix timestamp) when the link will expire | [optional] 
**MemberLimit** | Pointer to **int32** | The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999 | [optional] 
**CreatesJoinRequest** | Pointer to **bool** | *True*, if users joining the chat via the link need to be approved by chat administrators. If *True*, *member\\_limit* can&#39;t be specified | [optional] 

## Methods

### NewPostEditChatInviteLinkRequest

`func NewPostEditChatInviteLinkRequest(chatId PostSendMessageRequestChatId, inviteLink string, ) *PostEditChatInviteLinkRequest`

NewPostEditChatInviteLinkRequest instantiates a new PostEditChatInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditChatInviteLinkRequestWithDefaults

`func NewPostEditChatInviteLinkRequestWithDefaults() *PostEditChatInviteLinkRequest`

NewPostEditChatInviteLinkRequestWithDefaults instantiates a new PostEditChatInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostEditChatInviteLinkRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostEditChatInviteLinkRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostEditChatInviteLinkRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetInviteLink

`func (o *PostEditChatInviteLinkRequest) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *PostEditChatInviteLinkRequest) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *PostEditChatInviteLinkRequest) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetName

`func (o *PostEditChatInviteLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEditChatInviteLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEditChatInviteLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostEditChatInviteLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpireDate

`func (o *PostEditChatInviteLinkRequest) GetExpireDate() int32`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *PostEditChatInviteLinkRequest) GetExpireDateOk() (*int32, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *PostEditChatInviteLinkRequest) SetExpireDate(v int32)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *PostEditChatInviteLinkRequest) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetMemberLimit

`func (o *PostEditChatInviteLinkRequest) GetMemberLimit() int32`

GetMemberLimit returns the MemberLimit field if non-nil, zero value otherwise.

### GetMemberLimitOk

`func (o *PostEditChatInviteLinkRequest) GetMemberLimitOk() (*int32, bool)`

GetMemberLimitOk returns a tuple with the MemberLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLimit

`func (o *PostEditChatInviteLinkRequest) SetMemberLimit(v int32)`

SetMemberLimit sets MemberLimit field to given value.

### HasMemberLimit

`func (o *PostEditChatInviteLinkRequest) HasMemberLimit() bool`

HasMemberLimit returns a boolean if a field has been set.

### GetCreatesJoinRequest

`func (o *PostEditChatInviteLinkRequest) GetCreatesJoinRequest() bool`

GetCreatesJoinRequest returns the CreatesJoinRequest field if non-nil, zero value otherwise.

### GetCreatesJoinRequestOk

`func (o *PostEditChatInviteLinkRequest) GetCreatesJoinRequestOk() (*bool, bool)`

GetCreatesJoinRequestOk returns a tuple with the CreatesJoinRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatesJoinRequest

`func (o *PostEditChatInviteLinkRequest) SetCreatesJoinRequest(v bool)`

SetCreatesJoinRequest sets CreatesJoinRequest field to given value.

### HasCreatesJoinRequest

`func (o *PostEditChatInviteLinkRequest) HasCreatesJoinRequest() bool`

HasCreatesJoinRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


