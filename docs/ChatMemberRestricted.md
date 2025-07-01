# ChatMemberRestricted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The member&#39;s status in the chat, always “restricted” | [default to "restricted"]
**User** | [**User**](User.md) |  | 
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
**CanChangeInfo** | **bool** | *True*, if the user is allowed to change the chat title, photo and other settings | 
**CanInviteUsers** | **bool** | *True*, if the user is allowed to invite new users to the chat | 
**CanPinMessages** | **bool** | *True*, if the user is allowed to pin messages | 
**CanManageTopics** | **bool** | *True*, if the user is allowed to create forum topics | 
**UntilDate** | **int32** | Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever | 

## Methods

### NewChatMemberRestricted

`func NewChatMemberRestricted(status string, user User, isMember bool, canSendMessages bool, canSendAudios bool, canSendDocuments bool, canSendPhotos bool, canSendVideos bool, canSendVideoNotes bool, canSendVoiceNotes bool, canSendPolls bool, canSendOtherMessages bool, canAddWebPagePreviews bool, canChangeInfo bool, canInviteUsers bool, canPinMessages bool, canManageTopics bool, untilDate int32, ) *ChatMemberRestricted`

NewChatMemberRestricted instantiates a new ChatMemberRestricted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMemberRestrictedWithDefaults

`func NewChatMemberRestrictedWithDefaults() *ChatMemberRestricted`

NewChatMemberRestrictedWithDefaults instantiates a new ChatMemberRestricted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatMemberRestricted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatMemberRestricted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatMemberRestricted) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *ChatMemberRestricted) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatMemberRestricted) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatMemberRestricted) SetUser(v User)`

SetUser sets User field to given value.


### GetIsMember

`func (o *ChatMemberRestricted) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ChatMemberRestricted) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ChatMemberRestricted) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.


### GetCanSendMessages

`func (o *ChatMemberRestricted) GetCanSendMessages() bool`

GetCanSendMessages returns the CanSendMessages field if non-nil, zero value otherwise.

### GetCanSendMessagesOk

`func (o *ChatMemberRestricted) GetCanSendMessagesOk() (*bool, bool)`

GetCanSendMessagesOk returns a tuple with the CanSendMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessages

`func (o *ChatMemberRestricted) SetCanSendMessages(v bool)`

SetCanSendMessages sets CanSendMessages field to given value.


### GetCanSendAudios

`func (o *ChatMemberRestricted) GetCanSendAudios() bool`

GetCanSendAudios returns the CanSendAudios field if non-nil, zero value otherwise.

### GetCanSendAudiosOk

`func (o *ChatMemberRestricted) GetCanSendAudiosOk() (*bool, bool)`

GetCanSendAudiosOk returns a tuple with the CanSendAudios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendAudios

`func (o *ChatMemberRestricted) SetCanSendAudios(v bool)`

SetCanSendAudios sets CanSendAudios field to given value.


### GetCanSendDocuments

`func (o *ChatMemberRestricted) GetCanSendDocuments() bool`

GetCanSendDocuments returns the CanSendDocuments field if non-nil, zero value otherwise.

### GetCanSendDocumentsOk

`func (o *ChatMemberRestricted) GetCanSendDocumentsOk() (*bool, bool)`

GetCanSendDocumentsOk returns a tuple with the CanSendDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendDocuments

`func (o *ChatMemberRestricted) SetCanSendDocuments(v bool)`

SetCanSendDocuments sets CanSendDocuments field to given value.


### GetCanSendPhotos

`func (o *ChatMemberRestricted) GetCanSendPhotos() bool`

GetCanSendPhotos returns the CanSendPhotos field if non-nil, zero value otherwise.

### GetCanSendPhotosOk

`func (o *ChatMemberRestricted) GetCanSendPhotosOk() (*bool, bool)`

GetCanSendPhotosOk returns a tuple with the CanSendPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPhotos

`func (o *ChatMemberRestricted) SetCanSendPhotos(v bool)`

SetCanSendPhotos sets CanSendPhotos field to given value.


### GetCanSendVideos

`func (o *ChatMemberRestricted) GetCanSendVideos() bool`

GetCanSendVideos returns the CanSendVideos field if non-nil, zero value otherwise.

### GetCanSendVideosOk

`func (o *ChatMemberRestricted) GetCanSendVideosOk() (*bool, bool)`

GetCanSendVideosOk returns a tuple with the CanSendVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideos

`func (o *ChatMemberRestricted) SetCanSendVideos(v bool)`

SetCanSendVideos sets CanSendVideos field to given value.


### GetCanSendVideoNotes

`func (o *ChatMemberRestricted) GetCanSendVideoNotes() bool`

GetCanSendVideoNotes returns the CanSendVideoNotes field if non-nil, zero value otherwise.

### GetCanSendVideoNotesOk

`func (o *ChatMemberRestricted) GetCanSendVideoNotesOk() (*bool, bool)`

GetCanSendVideoNotesOk returns a tuple with the CanSendVideoNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideoNotes

`func (o *ChatMemberRestricted) SetCanSendVideoNotes(v bool)`

SetCanSendVideoNotes sets CanSendVideoNotes field to given value.


### GetCanSendVoiceNotes

`func (o *ChatMemberRestricted) GetCanSendVoiceNotes() bool`

GetCanSendVoiceNotes returns the CanSendVoiceNotes field if non-nil, zero value otherwise.

### GetCanSendVoiceNotesOk

`func (o *ChatMemberRestricted) GetCanSendVoiceNotesOk() (*bool, bool)`

GetCanSendVoiceNotesOk returns a tuple with the CanSendVoiceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVoiceNotes

`func (o *ChatMemberRestricted) SetCanSendVoiceNotes(v bool)`

SetCanSendVoiceNotes sets CanSendVoiceNotes field to given value.


### GetCanSendPolls

`func (o *ChatMemberRestricted) GetCanSendPolls() bool`

GetCanSendPolls returns the CanSendPolls field if non-nil, zero value otherwise.

### GetCanSendPollsOk

`func (o *ChatMemberRestricted) GetCanSendPollsOk() (*bool, bool)`

GetCanSendPollsOk returns a tuple with the CanSendPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPolls

`func (o *ChatMemberRestricted) SetCanSendPolls(v bool)`

SetCanSendPolls sets CanSendPolls field to given value.


### GetCanSendOtherMessages

`func (o *ChatMemberRestricted) GetCanSendOtherMessages() bool`

GetCanSendOtherMessages returns the CanSendOtherMessages field if non-nil, zero value otherwise.

### GetCanSendOtherMessagesOk

`func (o *ChatMemberRestricted) GetCanSendOtherMessagesOk() (*bool, bool)`

GetCanSendOtherMessagesOk returns a tuple with the CanSendOtherMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendOtherMessages

`func (o *ChatMemberRestricted) SetCanSendOtherMessages(v bool)`

SetCanSendOtherMessages sets CanSendOtherMessages field to given value.


### GetCanAddWebPagePreviews

`func (o *ChatMemberRestricted) GetCanAddWebPagePreviews() bool`

GetCanAddWebPagePreviews returns the CanAddWebPagePreviews field if non-nil, zero value otherwise.

### GetCanAddWebPagePreviewsOk

`func (o *ChatMemberRestricted) GetCanAddWebPagePreviewsOk() (*bool, bool)`

GetCanAddWebPagePreviewsOk returns a tuple with the CanAddWebPagePreviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddWebPagePreviews

`func (o *ChatMemberRestricted) SetCanAddWebPagePreviews(v bool)`

SetCanAddWebPagePreviews sets CanAddWebPagePreviews field to given value.


### GetCanChangeInfo

`func (o *ChatMemberRestricted) GetCanChangeInfo() bool`

GetCanChangeInfo returns the CanChangeInfo field if non-nil, zero value otherwise.

### GetCanChangeInfoOk

`func (o *ChatMemberRestricted) GetCanChangeInfoOk() (*bool, bool)`

GetCanChangeInfoOk returns a tuple with the CanChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeInfo

`func (o *ChatMemberRestricted) SetCanChangeInfo(v bool)`

SetCanChangeInfo sets CanChangeInfo field to given value.


### GetCanInviteUsers

`func (o *ChatMemberRestricted) GetCanInviteUsers() bool`

GetCanInviteUsers returns the CanInviteUsers field if non-nil, zero value otherwise.

### GetCanInviteUsersOk

`func (o *ChatMemberRestricted) GetCanInviteUsersOk() (*bool, bool)`

GetCanInviteUsersOk returns a tuple with the CanInviteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsers

`func (o *ChatMemberRestricted) SetCanInviteUsers(v bool)`

SetCanInviteUsers sets CanInviteUsers field to given value.


### GetCanPinMessages

`func (o *ChatMemberRestricted) GetCanPinMessages() bool`

GetCanPinMessages returns the CanPinMessages field if non-nil, zero value otherwise.

### GetCanPinMessagesOk

`func (o *ChatMemberRestricted) GetCanPinMessagesOk() (*bool, bool)`

GetCanPinMessagesOk returns a tuple with the CanPinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPinMessages

`func (o *ChatMemberRestricted) SetCanPinMessages(v bool)`

SetCanPinMessages sets CanPinMessages field to given value.


### GetCanManageTopics

`func (o *ChatMemberRestricted) GetCanManageTopics() bool`

GetCanManageTopics returns the CanManageTopics field if non-nil, zero value otherwise.

### GetCanManageTopicsOk

`func (o *ChatMemberRestricted) GetCanManageTopicsOk() (*bool, bool)`

GetCanManageTopicsOk returns a tuple with the CanManageTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTopics

`func (o *ChatMemberRestricted) SetCanManageTopics(v bool)`

SetCanManageTopics sets CanManageTopics field to given value.


### GetUntilDate

`func (o *ChatMemberRestricted) GetUntilDate() int32`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *ChatMemberRestricted) GetUntilDateOk() (*int32, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *ChatMemberRestricted) SetUntilDate(v int32)`

SetUntilDate sets UntilDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


