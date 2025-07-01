# ChatMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “kicked” | [default to "kicked"]
**User** | [**User**](User.md) |  | 
**IsAnonymous** | **bool** | *True*, if the user&#39;s presence in the chat is hidden | 
**CustomTitle** | Pointer to **string** | *Optional*. Custom title for this user | [optional] 
**CanBeEdited** | **bool** | *True*, if the bot is allowed to edit administrator privileges of that user | 
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
**CanPinMessages** | **bool** | *True*, if the user is allowed to pin messages | 
**CanManageTopics** | **bool** | *True*, if the user is allowed to create forum topics | 
**UntilDate** | **int32** | Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever | 
**IsMember** | **bool** | *True*, if the user is a member of the chat at the moment of the request | 
**CanSendMessages** | **bool** | *True*, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues | 
**CanSendAudios** | **bool** | *True*, if the user is allowed to send audios | 
**CanSendDocuments** | **bool** | *True*, if the user is allowed to send documents | 
**CanSendPhotos** | **bool** | *True*, if the user is allowed to send photos | 
**CanSendVideos** | **bool** | *True*, if the user is allowed to send videos | 
**CanSendVideoNotes** | **bool** | *True*, if the user is allowed to send video notes | 
**CanSendVoiceNotes** | **bool** | *True*, if the user is allowed to send voice notes | 
**CanSendPolls** | **bool** | *True*, if the user is allowed to send polls | 
**CanSendOtherMessages** | **bool** | *True*, if the user is allowed to send animations, games, stickers and use inline bots | 
**CanAddWebPagePreviews** | **bool** | *True*, if the user is allowed to add web page previews to their messages | 

## Methods

### NewChatMember

`func NewChatMember(status string, user User, isAnonymous bool, canBeEdited bool, canManageChat bool, canDeleteMessages bool, canManageVideoChats bool, canRestrictMembers bool, canPromoteMembers bool, canChangeInfo bool, canInviteUsers bool, canPostStories bool, canEditStories bool, canDeleteStories bool, canPinMessages bool, canManageTopics bool, untilDate int32, isMember bool, canSendMessages bool, canSendAudios bool, canSendDocuments bool, canSendPhotos bool, canSendVideos bool, canSendVideoNotes bool, canSendVoiceNotes bool, canSendPolls bool, canSendOtherMessages bool, canAddWebPagePreviews bool, ) *ChatMember`

NewChatMember instantiates a new ChatMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberWithDefaults

`func NewChatMemberWithDefaults() *ChatMember`

NewChatMemberWithDefaults instantiates a new ChatMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMember) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMember) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMember) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMember) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMember) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMember) SetUser(v User)`

SetUser sets User field to given value.


### GetIsAnonymous

`func (o *ChatMember) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *ChatMember) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *ChatMember) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.


### GetCustomTitle

`func (o *ChatMember) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *ChatMember) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *ChatMember) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *ChatMember) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.

### GetCanBeEdited

`func (o *ChatMember) GetCanBeEdited() bool`

GetCanBeEdited returns the CanBeEdited field if non-nil, zero value otherwise.

### GetCanBeEditedOk

`func (o *ChatMember) GetCanBeEditedOk() (*bool, bool)`

GetCanBeEditedOk returns a tuple with the CanBeEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeEdited

`func (o *ChatMember) SetCanBeEdited(v bool)`

SetCanBeEdited sets CanBeEdited field to given value.


### GetCanManageChat

`func (o *ChatMember) GetCanManageChat() bool`

GetCanManageChat returns the CanManageChat field if non-nil, zero value otherwise.

### GetCanManageChatOk

`func (o *ChatMember) GetCanManageChatOk() (*bool, bool)`

GetCanManageChatOk returns a tuple with the CanManageChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageChat

`func (o *ChatMember) SetCanManageChat(v bool)`

SetCanManageChat sets CanManageChat field to given value.


### GetCanDeleteMessages

`func (o *ChatMember) GetCanDeleteMessages() bool`

GetCanDeleteMessages returns the CanDeleteMessages field if non-nil, zero value otherwise.

### GetCanDeleteMessagesOk

`func (o *ChatMember) GetCanDeleteMessagesOk() (*bool, bool)`

GetCanDeleteMessagesOk returns a tuple with the CanDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteMessages

`func (o *ChatMember) SetCanDeleteMessages(v bool)`

SetCanDeleteMessages sets CanDeleteMessages field to given value.


### GetCanManageVideoChats

`func (o *ChatMember) GetCanManageVideoChats() bool`

GetCanManageVideoChats returns the CanManageVideoChats field if non-nil, zero value otherwise.

### GetCanManageVideoChatsOk

`func (o *ChatMember) GetCanManageVideoChatsOk() (*bool, bool)`

GetCanManageVideoChatsOk returns a tuple with the CanManageVideoChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageVideoChats

`func (o *ChatMember) SetCanManageVideoChats(v bool)`

SetCanManageVideoChats sets CanManageVideoChats field to given value.


### GetCanRestrictMembers

`func (o *ChatMember) GetCanRestrictMembers() bool`

GetCanRestrictMembers returns the CanRestrictMembers field if non-nil, zero value otherwise.

### GetCanRestrictMembersOk

`func (o *ChatMember) GetCanRestrictMembersOk() (*bool, bool)`

GetCanRestrictMembersOk returns a tuple with the CanRestrictMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRestrictMembers

`func (o *ChatMember) SetCanRestrictMembers(v bool)`

SetCanRestrictMembers sets CanRestrictMembers field to given value.


### GetCanPromoteMembers

`func (o *ChatMember) GetCanPromoteMembers() bool`

GetCanPromoteMembers returns the CanPromoteMembers field if non-nil, zero value otherwise.

### GetCanPromoteMembersOk

`func (o *ChatMember) GetCanPromoteMembersOk() (*bool, bool)`

GetCanPromoteMembersOk returns a tuple with the CanPromoteMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPromoteMembers

`func (o *ChatMember) SetCanPromoteMembers(v bool)`

SetCanPromoteMembers sets CanPromoteMembers field to given value.


### GetCanChangeInfo

`func (o *ChatMember) GetCanChangeInfo() bool`

GetCanChangeInfo returns the CanChangeInfo field if non-nil, zero value otherwise.

### GetCanChangeInfoOk

`func (o *ChatMember) GetCanChangeInfoOk() (*bool, bool)`

GetCanChangeInfoOk returns a tuple with the CanChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeInfo

`func (o *ChatMember) SetCanChangeInfo(v bool)`

SetCanChangeInfo sets CanChangeInfo field to given value.


### GetCanInviteUsers

`func (o *ChatMember) GetCanInviteUsers() bool`

GetCanInviteUsers returns the CanInviteUsers field if non-nil, zero value otherwise.

### GetCanInviteUsersOk

`func (o *ChatMember) GetCanInviteUsersOk() (*bool, bool)`

GetCanInviteUsersOk returns a tuple with the CanInviteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsers

`func (o *ChatMember) SetCanInviteUsers(v bool)`

SetCanInviteUsers sets CanInviteUsers field to given value.


### GetCanPostStories

`func (o *ChatMember) GetCanPostStories() bool`

GetCanPostStories returns the CanPostStories field if non-nil, zero value otherwise.

### GetCanPostStoriesOk

`func (o *ChatMember) GetCanPostStoriesOk() (*bool, bool)`

GetCanPostStoriesOk returns a tuple with the CanPostStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostStories

`func (o *ChatMember) SetCanPostStories(v bool)`

SetCanPostStories sets CanPostStories field to given value.


### GetCanEditStories

`func (o *ChatMember) GetCanEditStories() bool`

GetCanEditStories returns the CanEditStories field if non-nil, zero value otherwise.

### GetCanEditStoriesOk

`func (o *ChatMember) GetCanEditStoriesOk() (*bool, bool)`

GetCanEditStoriesOk returns a tuple with the CanEditStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditStories

`func (o *ChatMember) SetCanEditStories(v bool)`

SetCanEditStories sets CanEditStories field to given value.


### GetCanDeleteStories

`func (o *ChatMember) GetCanDeleteStories() bool`

GetCanDeleteStories returns the CanDeleteStories field if non-nil, zero value otherwise.

### GetCanDeleteStoriesOk

`func (o *ChatMember) GetCanDeleteStoriesOk() (*bool, bool)`

GetCanDeleteStoriesOk returns a tuple with the CanDeleteStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteStories

`func (o *ChatMember) SetCanDeleteStories(v bool)`

SetCanDeleteStories sets CanDeleteStories field to given value.


### GetCanPostMessages

`func (o *ChatMember) GetCanPostMessages() bool`

GetCanPostMessages returns the CanPostMessages field if non-nil, zero value otherwise.

### GetCanPostMessagesOk

`func (o *ChatMember) GetCanPostMessagesOk() (*bool, bool)`

GetCanPostMessagesOk returns a tuple with the CanPostMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPostMessages

`func (o *ChatMember) SetCanPostMessages(v bool)`

SetCanPostMessages sets CanPostMessages field to given value.

### HasCanPostMessages

`func (o *ChatMember) HasCanPostMessages() bool`

HasCanPostMessages returns a boolean if a field has been set.

### GetCanEditMessages

`func (o *ChatMember) GetCanEditMessages() bool`

GetCanEditMessages returns the CanEditMessages field if non-nil, zero value otherwise.

### GetCanEditMessagesOk

`func (o *ChatMember) GetCanEditMessagesOk() (*bool, bool)`

GetCanEditMessagesOk returns a tuple with the CanEditMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditMessages

`func (o *ChatMember) SetCanEditMessages(v bool)`

SetCanEditMessages sets CanEditMessages field to given value.

### HasCanEditMessages

`func (o *ChatMember) HasCanEditMessages() bool`

HasCanEditMessages returns a boolean if a field has been set.

### GetCanPinMessages

`func (o *ChatMember) GetCanPinMessages() bool`

GetCanPinMessages returns the CanPinMessages field if non-nil, zero value otherwise.

### GetCanPinMessagesOk

`func (o *ChatMember) GetCanPinMessagesOk() (*bool, bool)`

GetCanPinMessagesOk returns a tuple with the CanPinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPinMessages

`func (o *ChatMember) SetCanPinMessages(v bool)`

SetCanPinMessages sets CanPinMessages field to given value.


### GetCanManageTopics

`func (o *ChatMember) GetCanManageTopics() bool`

GetCanManageTopics returns the CanManageTopics field if non-nil, zero value otherwise.

### GetCanManageTopicsOk

`func (o *ChatMember) GetCanManageTopicsOk() (*bool, bool)`

GetCanManageTopicsOk returns a tuple with the CanManageTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTopics

`func (o *ChatMember) SetCanManageTopics(v bool)`

SetCanManageTopics sets CanManageTopics field to given value.


### GetUntilDate

`func (o *ChatMember) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *ChatMember) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *ChatMember) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.


### GetIsMember

`func (o *ChatMember) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ChatMember) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ChatMember) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.


### GetCanSendMessages

`func (o *ChatMember) GetCanSendMessages() bool`

GetCanSendMessages returns the CanSendMessages field if non-nil, zero value otherwise.

### GetCanSendMessagesOk

`func (o *ChatMember) GetCanSendMessagesOk() (*bool, bool)`

GetCanSendMessagesOk returns a tuple with the CanSendMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessages

`func (o *ChatMember) SetCanSendMessages(v bool)`

SetCanSendMessages sets CanSendMessages field to given value.


### GetCanSendAudios

`func (o *ChatMember) GetCanSendAudios() bool`

GetCanSendAudios returns the CanSendAudios field if non-nil, zero value otherwise.

### GetCanSendAudiosOk

`func (o *ChatMember) GetCanSendAudiosOk() (*bool, bool)`

GetCanSendAudiosOk returns a tuple with the CanSendAudios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendAudios

`func (o *ChatMember) SetCanSendAudios(v bool)`

SetCanSendAudios sets CanSendAudios field to given value.


### GetCanSendDocuments

`func (o *ChatMember) GetCanSendDocuments() bool`

GetCanSendDocuments returns the CanSendDocuments field if non-nil, zero value otherwise.

### GetCanSendDocumentsOk

`func (o *ChatMember) GetCanSendDocumentsOk() (*bool, bool)`

GetCanSendDocumentsOk returns a tuple with the CanSendDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendDocuments

`func (o *ChatMember) SetCanSendDocuments(v bool)`

SetCanSendDocuments sets CanSendDocuments field to given value.


### GetCanSendPhotos

`func (o *ChatMember) GetCanSendPhotos() bool`

GetCanSendPhotos returns the CanSendPhotos field if non-nil, zero value otherwise.

### GetCanSendPhotosOk

`func (o *ChatMember) GetCanSendPhotosOk() (*bool, bool)`

GetCanSendPhotosOk returns a tuple with the CanSendPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPhotos

`func (o *ChatMember) SetCanSendPhotos(v bool)`

SetCanSendPhotos sets CanSendPhotos field to given value.


### GetCanSendVideos

`func (o *ChatMember) GetCanSendVideos() bool`

GetCanSendVideos returns the CanSendVideos field if non-nil, zero value otherwise.

### GetCanSendVideosOk

`func (o *ChatMember) GetCanSendVideosOk() (*bool, bool)`

GetCanSendVideosOk returns a tuple with the CanSendVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideos

`func (o *ChatMember) SetCanSendVideos(v bool)`

SetCanSendVideos sets CanSendVideos field to given value.


### GetCanSendVideoNotes

`func (o *ChatMember) GetCanSendVideoNotes() bool`

GetCanSendVideoNotes returns the CanSendVideoNotes field if non-nil, zero value otherwise.

### GetCanSendVideoNotesOk

`func (o *ChatMember) GetCanSendVideoNotesOk() (*bool, bool)`

GetCanSendVideoNotesOk returns a tuple with the CanSendVideoNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideoNotes

`func (o *ChatMember) SetCanSendVideoNotes(v bool)`

SetCanSendVideoNotes sets CanSendVideoNotes field to given value.


### GetCanSendVoiceNotes

`func (o *ChatMember) GetCanSendVoiceNotes() bool`

GetCanSendVoiceNotes returns the CanSendVoiceNotes field if non-nil, zero value otherwise.

### GetCanSendVoiceNotesOk

`func (o *ChatMember) GetCanSendVoiceNotesOk() (*bool, bool)`

GetCanSendVoiceNotesOk returns a tuple with the CanSendVoiceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVoiceNotes

`func (o *ChatMember) SetCanSendVoiceNotes(v bool)`

SetCanSendVoiceNotes sets CanSendVoiceNotes field to given value.


### GetCanSendPolls

`func (o *ChatMember) GetCanSendPolls() bool`

GetCanSendPolls returns the CanSendPolls field if non-nil, zero value otherwise.

### GetCanSendPollsOk

`func (o *ChatMember) GetCanSendPollsOk() (*bool, bool)`

GetCanSendPollsOk returns a tuple with the CanSendPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPolls

`func (o *ChatMember) SetCanSendPolls(v bool)`

SetCanSendPolls sets CanSendPolls field to given value.


### GetCanSendOtherMessages

`func (o *ChatMember) GetCanSendOtherMessages() bool`

GetCanSendOtherMessages returns the CanSendOtherMessages field if non-nil, zero value otherwise.

### GetCanSendOtherMessagesOk

`func (o *ChatMember) GetCanSendOtherMessagesOk() (*bool, bool)`

GetCanSendOtherMessagesOk returns a tuple with the CanSendOtherMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendOtherMessages

`func (o *ChatMember) SetCanSendOtherMessages(v bool)`

SetCanSendOtherMessages sets CanSendOtherMessages field to given value.


### GetCanAddWebPagePreviews

`func (o *ChatMember) GetCanAddWebPagePreviews() bool`

GetCanAddWebPagePreviews returns the CanAddWebPagePreviews field if non-nil, zero value otherwise.

### GetCanAddWebPagePreviewsOk

`func (o *ChatMember) GetCanAddWebPagePreviewsOk() (*bool, bool)`

GetCanAddWebPagePreviewsOk returns a tuple with the CanAddWebPagePreviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddWebPagePreviews

`func (o *ChatMember) SetCanAddWebPagePreviews(v bool)`

SetCanAddWebPagePreviews sets CanAddWebPagePreviews field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


