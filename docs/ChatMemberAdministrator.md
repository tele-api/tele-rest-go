# ChatMemberAdministrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “administrator” | [default to "administrator"]
**User** | [**User**](User.md) |  | 
**CanBeEdited** | **bool** | *True*, if the bot is allowed to edit administrator privileges of that user | 
**IsAnonymous** | **bool** | *True*, if the user&#39;s presence in the chat is hidden | 
**CanManageChat** | **bool** | *True*, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege. | 
**CanDeleteMessages** | **bool** | *True*, if the administrator can delete messages of other users | 
**CanManageVideoChats** | **bool** | *True*, if the administrator can manage video chats | 
**CanRestrictMembers** | **bool** | *True*, if the administrator can restrict, ban or unban chat members, or access supergroup statistics | 
**CanPromoteMembers** | **bool** | *True*, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user) | 
**CanChangeInfo** | **bool** | *True*, if the user is allowed to change the chat title, photo and other settings | 
**CanInviteUsers** | **bool** | *True*, if the user is allowed to invite new users to the chat | 
**CanPostStories** | **bool** | *True*, if the administrator can post stories to the chat | 
**CanEditStories** | **bool** | *True*, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat&#39;s story archive | 
**CanDeleteStories** | **bool** | *True*, if the administrator can delete stories posted by other users | 
**CanPostMessages** | Pointer to **bool** | *Optional*. *True*, if the administrator can post messages in the channel, or access channel statistics; for channels only | [optional] 
**CanEditMessages** | Pointer to **bool** | *Optional*. *True*, if the administrator can edit messages of other users and can pin messages; for channels only | [optional] 
**CanPinMessages** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to pin messages; for groups and supergroups only | [optional] 
**CanManageTopics** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only | [optional] 
**CustomTitle** | Pointer to **string** | *Optional*. Custom title for this user | [optional] 

## Methods

### NewChatMemberAdministrator

`func NewChatMemberAdministrator(status string, user User, canBeEdited bool, isAnonymous bool, canManageChat bool, canDeleteMessages bool, canManageVideoChats bool, canRestrictMembers bool, canPromoteMembers bool, canChangeInfo bool, canInviteUsers bool, canPostStories bool, canEditStories bool, canDeleteStories bool, ) *ChatMemberAdministrator`

NewChatMemberAdministrator instantiates a new ChatMemberAdministrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberAdministratorWithDefaults

`func NewChatMemberAdministratorWithDefaults() *ChatMemberAdministrator`

NewChatMemberAdministratorWithDefaults instantiates a new ChatMemberAdministrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberAdministrator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberAdministrator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberAdministrator) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberAdministrator) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberAdministrator) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberAdministrator) SetUser(v User)`

SetUser sets User field to given value.


### GetCanBeEdited

`func (o *ChatMemberAdministrator) GetCanBeEdited() bool`

GetCanBeEdited returns the CanBeEdited field if non-nil, zero value otherwise.

### GetCanBeEditedOk

`func (o *ChatMemberAdministrator) GetCanBeEditedOk() (*bool, bool)`

GetCanBeEditedOk returns a tuple with the CanBeEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeEdited

`func (o *ChatMemberAdministrator) SetCanBeEdited(v bool)`

SetCanBeEdited sets CanBeEdited field to given value.


### GetIsAnonymous

`func (o *ChatMemberAdministrator) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *ChatMemberAdministrator) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *ChatMemberAdministrator) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.


### GetCanManageChat

`func (o *ChatMemberAdministrator) GetCanManageChat() bool`

GetCanManageChat returns the CanManageChat field if non-nil, zero value otherwise.

### GetCanManageChatOk

`func (o *ChatMemberAdministrator) GetCanManageChatOk() (*bool, bool)`

GetCanManageChatOk returns a tuple with the CanManageChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageChat

`func (o *ChatMemberAdministrator) SetCanManageChat(v bool)`

SetCanManageChat sets CanManageChat field to given value.


### GetCanDeleteMessages

`func (o *ChatMemberAdministrator) GetCanDeleteMessages() bool`

GetCanDeleteMessages returns the CanDeleteMessages field if non-nil, zero value otherwise.

### GetCanDeleteMessagesOk

`func (o *ChatMemberAdministrator) GetCanDeleteMessagesOk() (*bool, bool)`

GetCanDeleteMessagesOk returns a tuple with the CanDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteMessages

`func (o *ChatMemberAdministrator) SetCanDeleteMessages(v bool)`

SetCanDeleteMessages sets CanDeleteMessages field to given value.


### GetCanManageVideoChats

`func (o *ChatMemberAdministrator) GetCanManageVideoChats() bool`

GetCanManageVideoChats returns the CanManageVideoChats field if non-nil, zero value otherwise.

### GetCanManageVideoChatsOk

`func (o *ChatMemberAdministrator) GetCanManageVideoChatsOk() (*bool, bool)`

GetCanManageVideoChatsOk returns a tuple with the CanManageVideoChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageVideoChats

`func (o *ChatMemberAdministrator) SetCanManageVideoChats(v bool)`

SetCanManageVideoChats sets CanManageVideoChats field to given value.


### GetCanRestrictMembers

`func (o *ChatMemberAdministrator) GetCanRestrictMembers() bool`

GetCanRestrictMembers returns the CanRestrictMembers field if non-nil, zero value otherwise.

### GetCanRestrictMembersOk

`func (o *ChatMemberAdministrator) GetCanRestrictMembersOk() (*bool, bool)`

GetCanRestrictMembersOk returns a tuple with the CanRestrictMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRestrictMembers

`func (o *ChatMemberAdministrator) SetCanRestrictMembers(v bool)`

SetCanRestrictMembers sets CanRestrictMembers field to given value.


### GetCanPromoteMembers

`func (o *ChatMemberAdministrator) GetCanPromoteMembers() bool`

GetCanPromoteMembers returns the CanPromoteMembers field if non-nil, zero value otherwise.

### GetCanPromoteMembersOk

`func (o *ChatMemberAdministrator) GetCanPromoteMembersOk() (*bool, bool)`

GetCanPromoteMembersOk returns a tuple with the CanPromoteMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPromoteMembers

`func (o *ChatMemberAdministrator) SetCanPromoteMembers(v bool)`

SetCanPromoteMembers sets CanPromoteMembers field to given value.


### GetCanChangeInfo

`func (o *ChatMemberAdministrator) GetCanChangeInfo() bool`

GetCanChangeInfo returns the CanChangeInfo field if non-nil, zero value otherwise.

### GetCanChangeInfoOk

`func (o *ChatMemberAdministrator) GetCanChangeInfoOk() (*bool, bool)`

GetCanChangeInfoOk returns a tuple with the CanChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeInfo

`func (o *ChatMemberAdministrator) SetCanChangeInfo(v bool)`

SetCanChangeInfo sets CanChangeInfo field to given value.


### GetCanInviteUsers

`func (o *ChatMemberAdministrator) GetCanInviteUsers() bool`

GetCanInviteUsers returns the CanInviteUsers field if non-nil, zero value otherwise.

### GetCanInviteUsersOk

`func (o *ChatMemberAdministrator) GetCanInviteUsersOk() (*bool, bool)`

GetCanInviteUsersOk returns a tuple with the CanInviteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsers

`func (o *ChatMemberAdministrator) SetCanInviteUsers(v bool)`

SetCanInviteUsers sets CanInviteUsers field to given value.


### GetCanPostStories

`func (o *ChatMemberAdministrator) GetCanPostStories() bool`

GetCanPostStories returns the CanPostStories field if non-nil, zero value otherwise.

### GetCanPostStoriesOk

`func (o *ChatMemberAdministrator) GetCanPostStoriesOk() (*bool, bool)`

GetCanPostStoriesOk returns a tuple with the CanPostStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostStories

`func (o *ChatMemberAdministrator) SetCanPostStories(v bool)`

SetCanPostStories sets CanPostStories field to given value.


### GetCanEditStories

`func (o *ChatMemberAdministrator) GetCanEditStories() bool`

GetCanEditStories returns the CanEditStories field if non-nil, zero value otherwise.

### GetCanEditStoriesOk

`func (o *ChatMemberAdministrator) GetCanEditStoriesOk() (*bool, bool)`

GetCanEditStoriesOk returns a tuple with the CanEditStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditStories

`func (o *ChatMemberAdministrator) SetCanEditStories(v bool)`

SetCanEditStories sets CanEditStories field to given value.


### GetCanDeleteStories

`func (o *ChatMemberAdministrator) GetCanDeleteStories() bool`

GetCanDeleteStories returns the CanDeleteStories field if non-nil, zero value otherwise.

### GetCanDeleteStoriesOk

`func (o *ChatMemberAdministrator) GetCanDeleteStoriesOk() (*bool, bool)`

GetCanDeleteStoriesOk returns a tuple with the CanDeleteStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteStories

`func (o *ChatMemberAdministrator) SetCanDeleteStories(v bool)`

SetCanDeleteStories sets CanDeleteStories field to given value.


### GetCanPostMessages

`func (o *ChatMemberAdministrator) GetCanPostMessages() bool`

GetCanPostMessages returns the CanPostMessages field if non-nil, zero value otherwise.

### GetCanPostMessagesOk

`func (o *ChatMemberAdministrator) GetCanPostMessagesOk() (*bool, bool)`

GetCanPostMessagesOk returns a tuple with the CanPostMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostMessages

`func (o *ChatMemberAdministrator) SetCanPostMessages(v bool)`

SetCanPostMessages sets CanPostMessages field to given value.

### HasCanPostMessages

`func (o *ChatMemberAdministrator) HasCanPostMessages() bool`

HasCanPostMessages returns a boolean if a field has been set.

### GetCanEditMessages

`func (o *ChatMemberAdministrator) GetCanEditMessages() bool`

GetCanEditMessages returns the CanEditMessages field if non-nil, zero value otherwise.

### GetCanEditMessagesOk

`func (o *ChatMemberAdministrator) GetCanEditMessagesOk() (*bool, bool)`

GetCanEditMessagesOk returns a tuple with the CanEditMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditMessages

`func (o *ChatMemberAdministrator) SetCanEditMessages(v bool)`

SetCanEditMessages sets CanEditMessages field to given value.

### HasCanEditMessages

`func (o *ChatMemberAdministrator) HasCanEditMessages() bool`

HasCanEditMessages returns a boolean if a field has been set.

### GetCanPinMessages

`func (o *ChatMemberAdministrator) GetCanPinMessages() bool`

GetCanPinMessages returns the CanPinMessages field if non-nil, zero value otherwise.

### GetCanPinMessagesOk

`func (o *ChatMemberAdministrator) GetCanPinMessagesOk() (*bool, bool)`

GetCanPinMessagesOk returns a tuple with the CanPinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPinMessages

`func (o *ChatMemberAdministrator) SetCanPinMessages(v bool)`

SetCanPinMessages sets CanPinMessages field to given value.

### HasCanPinMessages

`func (o *ChatMemberAdministrator) HasCanPinMessages() bool`

HasCanPinMessages returns a boolean if a field has been set.

### GetCanManageTopics

`func (o *ChatMemberAdministrator) GetCanManageTopics() bool`

GetCanManageTopics returns the CanManageTopics field if non-nil, zero value otherwise.

### GetCanManageTopicsOk

`func (o *ChatMemberAdministrator) GetCanManageTopicsOk() (*bool, bool)`

GetCanManageTopicsOk returns a tuple with the CanManageTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTopics

`func (o *ChatMemberAdministrator) SetCanManageTopics(v bool)`

SetCanManageTopics sets CanManageTopics field to given value.

### HasCanManageTopics

`func (o *ChatMemberAdministrator) HasCanManageTopics() bool`

HasCanManageTopics returns a boolean if a field has been set.

### GetCustomTitle

`func (o *ChatMemberAdministrator) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *ChatMemberAdministrator) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *ChatMemberAdministrator) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *ChatMemberAdministrator) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


