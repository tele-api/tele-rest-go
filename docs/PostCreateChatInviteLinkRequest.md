# PostCreateChatInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 
**ExpireDate** | Pointer to **int32** | Point in time (Unix timestamp) when the link will expire | [optional] 
**MemberLimit** | Pointer to **int32** | The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999 | [optional] 
**CreatesJoinRequest** | Pointer to **bool** | *True*, if users joining the chat via the link need to be approved by chat administrators. If *True*, *member\\_limit* can&#39;t be specified | [optional] 

## Methods

### NewPostCreateChatInviteLinkRequest

`func NewPostCreateChatInviteLinkRequest(chatId PostSendMessageRequestChatId, ) *PostCreateChatInviteLinkRequest`

NewPostCreateChatInviteLinkRequest instantiates a new PostCreateChatInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateChatInviteLinkRequestWithDefaults

`func NewPostCreateChatInviteLinkRequestWithDefaults() *PostCreateChatInviteLinkRequest`

NewPostCreateChatInviteLinkRequestWithDefaults instantiates a new PostCreateChatInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostCreateChatInviteLinkRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostCreateChatInviteLinkRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostCreateChatInviteLinkRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *PostCreateChatInviteLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCreateChatInviteLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCreateChatInviteLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCreateChatInviteLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpireDate

`func (o *PostCreateChatInviteLinkRequest) GetExpireDate() int32`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *PostCreateChatInviteLinkRequest) GetExpireDateOk() (*int32, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *PostCreateChatInviteLinkRequest) SetExpireDate(v int32)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *PostCreateChatInviteLinkRequest) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetMemberLimit

`func (o *PostCreateChatInviteLinkRequest) GetMemberLimit() int32`

GetMemberLimit returns the MemberLimit field if non-nil, zero value otherwise.

### GetMemberLimitOk

`func (o *PostCreateChatInviteLinkRequest) GetMemberLimitOk() (*int32, bool)`

GetMemberLimitOk returns a tuple with the MemberLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLimit

`func (o *PostCreateChatInviteLinkRequest) SetMemberLimit(v int32)`

SetMemberLimit sets MemberLimit field to given value.

### HasMemberLimit

`func (o *PostCreateChatInviteLinkRequest) HasMemberLimit() bool`

HasMemberLimit returns a boolean if a field has been set.

### GetCreatesJoinRequest

`func (o *PostCreateChatInviteLinkRequest) GetCreatesJoinRequest() bool`

GetCreatesJoinRequest returns the CreatesJoinRequest field if non-nil, zero value otherwise.

### GetCreatesJoinRequestOk

`func (o *PostCreateChatInviteLinkRequest) GetCreatesJoinRequestOk() (*bool, bool)`

GetCreatesJoinRequestOk returns a tuple with the CreatesJoinRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatesJoinRequest

`func (o *PostCreateChatInviteLinkRequest) SetCreatesJoinRequest(v bool)`

SetCreatesJoinRequest sets CreatesJoinRequest field to given value.

### HasCreatesJoinRequest

`func (o *PostCreateChatInviteLinkRequest) HasCreatesJoinRequest() bool`

HasCreatesJoinRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


