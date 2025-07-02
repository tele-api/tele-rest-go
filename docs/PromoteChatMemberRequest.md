# PromoteChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**IsAnonymous** | Pointer to **bool** | Pass *True* if the administrator&#39;s presence in the chat is hidden | [optional] 
**CanManageChat** | Pointer to **bool** | Pass *True* if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege. | [optional] 
**CanDeleteMessages** | Pointer to **bool** | Pass *True* if the administrator can delete messages of other users | [optional] 
**CanManageVideoChats** | Pointer to **bool** | Pass *True* if the administrator can manage video chats | [optional] 
**CanRestrictMembers** | Pointer to **bool** | Pass *True* if the administrator can restrict, ban or unban chat members, or access supergroup statistics | [optional] 
**CanPromoteMembers** | Pointer to **bool** | Pass *True* if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him) | [optional] 
**CanChangeInfo** | Pointer to **bool** | Pass *True* if the administrator can change chat title, photo and other settings | [optional] 
**CanInviteUsers** | Pointer to **bool** | Pass *True* if the administrator can invite new users to the chat | [optional] 
**CanPostStories** | Pointer to **bool** | Pass *True* if the administrator can post stories to the chat | [optional] 
**CanEditStories** | Pointer to **bool** | Pass *True* if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat&#39;s story archive | [optional] 
**CanDeleteStories** | Pointer to **bool** | Pass *True* if the administrator can delete stories posted by other users | [optional] 
**CanPostMessages** | Pointer to **bool** | Pass *True* if the administrator can post messages in the channel, or access channel statistics; for channels only | [optional] 
**CanEditMessages** | Pointer to **bool** | Pass *True* if the administrator can edit messages of other users and can pin messages; for channels only | [optional] 
**CanPinMessages** | Pointer to **bool** | Pass *True* if the administrator can pin messages; for supergroups only | [optional] 
**CanManageTopics** | Pointer to **bool** | Pass *True* if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only | [optional] 

## Methods

### NewPromoteChatMemberRequest

`func NewPromoteChatMemberRequest(chatId SendMessageRequestChatId, userId int32, ) *PromoteChatMemberRequest`

NewPromoteChatMemberRequest instantiates a new PromoteChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteChatMemberRequestWithDefaults

`func NewPromoteChatMemberRequestWithDefaults() *PromoteChatMemberRequest`

NewPromoteChatMemberRequestWithDefaults instantiates a new PromoteChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PromoteChatMemberRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PromoteChatMemberRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PromoteChatMemberRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *PromoteChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PromoteChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PromoteChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetIsAnonymous

`func (o *PromoteChatMemberRequest) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *PromoteChatMemberRequest) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *PromoteChatMemberRequest) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *PromoteChatMemberRequest) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetCanManageChat

`func (o *PromoteChatMemberRequest) GetCanManageChat() bool`

GetCanManageChat returns the CanManageChat field if non-nil, zero value otherwise.

### GetCanManageChatOk

`func (o *PromoteChatMemberRequest) GetCanManageChatOk() (*bool, bool)`

GetCanManageChatOk returns a tuple with the CanManageChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageChat

`func (o *PromoteChatMemberRequest) SetCanManageChat(v bool)`

SetCanManageChat sets CanManageChat field to given value.

### HasCanManageChat

`func (o *PromoteChatMemberRequest) HasCanManageChat() bool`

HasCanManageChat returns a boolean if a field has been set.

### GetCanDeleteMessages

`func (o *PromoteChatMemberRequest) GetCanDeleteMessages() bool`

GetCanDeleteMessages returns the CanDeleteMessages field if non-nil, zero value otherwise.

### GetCanDeleteMessagesOk

`func (o *PromoteChatMemberRequest) GetCanDeleteMessagesOk() (*bool, bool)`

GetCanDeleteMessagesOk returns a tuple with the CanDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteMessages

`func (o *PromoteChatMemberRequest) SetCanDeleteMessages(v bool)`

SetCanDeleteMessages sets CanDeleteMessages field to given value.

### HasCanDeleteMessages

`func (o *PromoteChatMemberRequest) HasCanDeleteMessages() bool`

HasCanDeleteMessages returns a boolean if a field has been set.

### GetCanManageVideoChats

`func (o *PromoteChatMemberRequest) GetCanManageVideoChats() bool`

GetCanManageVideoChats returns the CanManageVideoChats field if non-nil, zero value otherwise.

### GetCanManageVideoChatsOk

`func (o *PromoteChatMemberRequest) GetCanManageVideoChatsOk() (*bool, bool)`

GetCanManageVideoChatsOk returns a tuple with the CanManageVideoChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageVideoChats

`func (o *PromoteChatMemberRequest) SetCanManageVideoChats(v bool)`

SetCanManageVideoChats sets CanManageVideoChats field to given value.

### HasCanManageVideoChats

`func (o *PromoteChatMemberRequest) HasCanManageVideoChats() bool`

HasCanManageVideoChats returns a boolean if a field has been set.

### GetCanRestrictMembers

`func (o *PromoteChatMemberRequest) GetCanRestrictMembers() bool`

GetCanRestrictMembers returns the CanRestrictMembers field if non-nil, zero value otherwise.

### GetCanRestrictMembersOk

`func (o *PromoteChatMemberRequest) GetCanRestrictMembersOk() (*bool, bool)`

GetCanRestrictMembersOk returns a tuple with the CanRestrictMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRestrictMembers

`func (o *PromoteChatMemberRequest) SetCanRestrictMembers(v bool)`

SetCanRestrictMembers sets CanRestrictMembers field to given value.

### HasCanRestrictMembers

`func (o *PromoteChatMemberRequest) HasCanRestrictMembers() bool`

HasCanRestrictMembers returns a boolean if a field has been set.

### GetCanPromoteMembers

`func (o *PromoteChatMemberRequest) GetCanPromoteMembers() bool`

GetCanPromoteMembers returns the CanPromoteMembers field if non-nil, zero value otherwise.

### GetCanPromoteMembersOk

`func (o *PromoteChatMemberRequest) GetCanPromoteMembersOk() (*bool, bool)`

GetCanPromoteMembersOk returns a tuple with the CanPromoteMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPromoteMembers

`func (o *PromoteChatMemberRequest) SetCanPromoteMembers(v bool)`

SetCanPromoteMembers sets CanPromoteMembers field to given value.

### HasCanPromoteMembers

`func (o *PromoteChatMemberRequest) HasCanPromoteMembers() bool`

HasCanPromoteMembers returns a boolean if a field has been set.

### GetCanChangeInfo

`func (o *PromoteChatMemberRequest) GetCanChangeInfo() bool`

GetCanChangeInfo returns the CanChangeInfo field if non-nil, zero value otherwise.

### GetCanChangeInfoOk

`func (o *PromoteChatMemberRequest) GetCanChangeInfoOk() (*bool, bool)`

GetCanChangeInfoOk returns a tuple with the CanChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeInfo

`func (o *PromoteChatMemberRequest) SetCanChangeInfo(v bool)`

SetCanChangeInfo sets CanChangeInfo field to given value.

### HasCanChangeInfo

`func (o *PromoteChatMemberRequest) HasCanChangeInfo() bool`

HasCanChangeInfo returns a boolean if a field has been set.

### GetCanInviteUsers

`func (o *PromoteChatMemberRequest) GetCanInviteUsers() bool`

GetCanInviteUsers returns the CanInviteUsers field if non-nil, zero value otherwise.

### GetCanInviteUsersOk

`func (o *PromoteChatMemberRequest) GetCanInviteUsersOk() (*bool, bool)`

GetCanInviteUsersOk returns a tuple with the CanInviteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsers

`func (o *PromoteChatMemberRequest) SetCanInviteUsers(v bool)`

SetCanInviteUsers sets CanInviteUsers field to given value.

### HasCanInviteUsers

`func (o *PromoteChatMemberRequest) HasCanInviteUsers() bool`

HasCanInviteUsers returns a boolean if a field has been set.

### GetCanPostStories

`func (o *PromoteChatMemberRequest) GetCanPostStories() bool`

GetCanPostStories returns the CanPostStories field if non-nil, zero value otherwise.

### GetCanPostStoriesOk

`func (o *PromoteChatMemberRequest) GetCanPostStoriesOk() (*bool, bool)`

GetCanPostStoriesOk returns a tuple with the CanPostStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostStories

`func (o *PromoteChatMemberRequest) SetCanPostStories(v bool)`

SetCanPostStories sets CanPostStories field to given value.

### HasCanPostStories

`func (o *PromoteChatMemberRequest) HasCanPostStories() bool`

HasCanPostStories returns a boolean if a field has been set.

### GetCanEditStories

`func (o *PromoteChatMemberRequest) GetCanEditStories() bool`

GetCanEditStories returns the CanEditStories field if non-nil, zero value otherwise.

### GetCanEditStoriesOk

`func (o *PromoteChatMemberRequest) GetCanEditStoriesOk() (*bool, bool)`

GetCanEditStoriesOk returns a tuple with the CanEditStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditStories

`func (o *PromoteChatMemberRequest) SetCanEditStories(v bool)`

SetCanEditStories sets CanEditStories field to given value.

### HasCanEditStories

`func (o *PromoteChatMemberRequest) HasCanEditStories() bool`

HasCanEditStories returns a boolean if a field has been set.

### GetCanDeleteStories

`func (o *PromoteChatMemberRequest) GetCanDeleteStories() bool`

GetCanDeleteStories returns the CanDeleteStories field if non-nil, zero value otherwise.

### GetCanDeleteStoriesOk

`func (o *PromoteChatMemberRequest) GetCanDeleteStoriesOk() (*bool, bool)`

GetCanDeleteStoriesOk returns a tuple with the CanDeleteStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteStories

`func (o *PromoteChatMemberRequest) SetCanDeleteStories(v bool)`

SetCanDeleteStories sets CanDeleteStories field to given value.

### HasCanDeleteStories

`func (o *PromoteChatMemberRequest) HasCanDeleteStories() bool`

HasCanDeleteStories returns a boolean if a field has been set.

### GetCanPostMessages

`func (o *PromoteChatMemberRequest) GetCanPostMessages() bool`

GetCanPostMessages returns the CanPostMessages field if non-nil, zero value otherwise.

### GetCanPostMessagesOk

`func (o *PromoteChatMemberRequest) GetCanPostMessagesOk() (*bool, bool)`

GetCanPostMessagesOk returns a tuple with the CanPostMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostMessages

`func (o *PromoteChatMemberRequest) SetCanPostMessages(v bool)`

SetCanPostMessages sets CanPostMessages field to given value.

### HasCanPostMessages

`func (o *PromoteChatMemberRequest) HasCanPostMessages() bool`

HasCanPostMessages returns a boolean if a field has been set.

### GetCanEditMessages

`func (o *PromoteChatMemberRequest) GetCanEditMessages() bool`

GetCanEditMessages returns the CanEditMessages field if non-nil, zero value otherwise.

### GetCanEditMessagesOk

`func (o *PromoteChatMemberRequest) GetCanEditMessagesOk() (*bool, bool)`

GetCanEditMessagesOk returns a tuple with the CanEditMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditMessages

`func (o *PromoteChatMemberRequest) SetCanEditMessages(v bool)`

SetCanEditMessages sets CanEditMessages field to given value.

### HasCanEditMessages

`func (o *PromoteChatMemberRequest) HasCanEditMessages() bool`

HasCanEditMessages returns a boolean if a field has been set.

### GetCanPinMessages

`func (o *PromoteChatMemberRequest) GetCanPinMessages() bool`

GetCanPinMessages returns the CanPinMessages field if non-nil, zero value otherwise.

### GetCanPinMessagesOk

`func (o *PromoteChatMemberRequest) GetCanPinMessagesOk() (*bool, bool)`

GetCanPinMessagesOk returns a tuple with the CanPinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPinMessages

`func (o *PromoteChatMemberRequest) SetCanPinMessages(v bool)`

SetCanPinMessages sets CanPinMessages field to given value.

### HasCanPinMessages

`func (o *PromoteChatMemberRequest) HasCanPinMessages() bool`

HasCanPinMessages returns a boolean if a field has been set.

### GetCanManageTopics

`func (o *PromoteChatMemberRequest) GetCanManageTopics() bool`

GetCanManageTopics returns the CanManageTopics field if non-nil, zero value otherwise.

### GetCanManageTopicsOk

`func (o *PromoteChatMemberRequest) GetCanManageTopicsOk() (*bool, bool)`

GetCanManageTopicsOk returns a tuple with the CanManageTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTopics

`func (o *PromoteChatMemberRequest) SetCanManageTopics(v bool)`

SetCanManageTopics sets CanManageTopics field to given value.

### HasCanManageTopics

`func (o *PromoteChatMemberRequest) HasCanManageTopics() bool`

HasCanManageTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


