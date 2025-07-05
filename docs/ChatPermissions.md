# ChatPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSendMessages** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues | [optional] 
**CanSendAudios** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send audios | [optional] 
**CanSendDocuments** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send documents | [optional] 
**CanSendPhotos** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send photos | [optional] 
**CanSendVideos** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send videos | [optional] 
**CanSendVideoNotes** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send video notes | [optional] 
**CanSendVoiceNotes** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send voice notes | [optional] 
**CanSendPolls** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send polls and checklists | [optional] 
**CanSendOtherMessages** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to send animations, games, stickers and use inline bots | [optional] 
**CanAddWebPagePreviews** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to add web page previews to their messages | [optional] 
**CanChangeInfo** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups | [optional] 
**CanInviteUsers** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to invite new users to the chat | [optional] 
**CanPinMessages** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to pin messages. Ignored in public supergroups | [optional] 
**CanManageTopics** | Pointer to **bool** | *Optional*. *True*, if the user is allowed to create forum topics. If omitted defaults to the value of can\\_pin\\_messages | [optional] 

## Methods

### NewChatPermissions

`func NewChatPermissions() *ChatPermissions`

NewChatPermissions instantiates a new ChatPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPermissionsWithDefaults

`func NewChatPermissionsWithDefaults() *ChatPermissions`

NewChatPermissionsWithDefaults instantiates a new ChatPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSendMessages

`func (o *ChatPermissions) GetCanSendMessages() bool`

GetCanSendMessages returns the CanSendMessages field if non-nil, zero value otherwise.

### GetCanSendMessagesOk

`func (o *ChatPermissions) GetCanSendMessagesOk() (*bool, bool)`

GetCanSendMessagesOk returns a tuple with the CanSendMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessages

`func (o *ChatPermissions) SetCanSendMessages(v bool)`

SetCanSendMessages sets CanSendMessages field to given value.

### HasCanSendMessages

`func (o *ChatPermissions) HasCanSendMessages() bool`

HasCanSendMessages returns a boolean if a field has been set.

### GetCanSendAudios

`func (o *ChatPermissions) GetCanSendAudios() bool`

GetCanSendAudios returns the CanSendAudios field if non-nil, zero value otherwise.

### GetCanSendAudiosOk

`func (o *ChatPermissions) GetCanSendAudiosOk() (*bool, bool)`

GetCanSendAudiosOk returns a tuple with the CanSendAudios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendAudios

`func (o *ChatPermissions) SetCanSendAudios(v bool)`

SetCanSendAudios sets CanSendAudios field to given value.

### HasCanSendAudios

`func (o *ChatPermissions) HasCanSendAudios() bool`

HasCanSendAudios returns a boolean if a field has been set.

### GetCanSendDocuments

`func (o *ChatPermissions) GetCanSendDocuments() bool`

GetCanSendDocuments returns the CanSendDocuments field if non-nil, zero value otherwise.

### GetCanSendDocumentsOk

`func (o *ChatPermissions) GetCanSendDocumentsOk() (*bool, bool)`

GetCanSendDocumentsOk returns a tuple with the CanSendDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendDocuments

`func (o *ChatPermissions) SetCanSendDocuments(v bool)`

SetCanSendDocuments sets CanSendDocuments field to given value.

### HasCanSendDocuments

`func (o *ChatPermissions) HasCanSendDocuments() bool`

HasCanSendDocuments returns a boolean if a field has been set.

### GetCanSendPhotos

`func (o *ChatPermissions) GetCanSendPhotos() bool`

GetCanSendPhotos returns the CanSendPhotos field if non-nil, zero value otherwise.

### GetCanSendPhotosOk

`func (o *ChatPermissions) GetCanSendPhotosOk() (*bool, bool)`

GetCanSendPhotosOk returns a tuple with the CanSendPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPhotos

`func (o *ChatPermissions) SetCanSendPhotos(v bool)`

SetCanSendPhotos sets CanSendPhotos field to given value.

### HasCanSendPhotos

`func (o *ChatPermissions) HasCanSendPhotos() bool`

HasCanSendPhotos returns a boolean if a field has been set.

### GetCanSendVideos

`func (o *ChatPermissions) GetCanSendVideos() bool`

GetCanSendVideos returns the CanSendVideos field if non-nil, zero value otherwise.

### GetCanSendVideosOk

`func (o *ChatPermissions) GetCanSendVideosOk() (*bool, bool)`

GetCanSendVideosOk returns a tuple with the CanSendVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideos

`func (o *ChatPermissions) SetCanSendVideos(v bool)`

SetCanSendVideos sets CanSendVideos field to given value.

### HasCanSendVideos

`func (o *ChatPermissions) HasCanSendVideos() bool`

HasCanSendVideos returns a boolean if a field has been set.

### GetCanSendVideoNotes

`func (o *ChatPermissions) GetCanSendVideoNotes() bool`

GetCanSendVideoNotes returns the CanSendVideoNotes field if non-nil, zero value otherwise.

### GetCanSendVideoNotesOk

`func (o *ChatPermissions) GetCanSendVideoNotesOk() (*bool, bool)`

GetCanSendVideoNotesOk returns a tuple with the CanSendVideoNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVideoNotes

`func (o *ChatPermissions) SetCanSendVideoNotes(v bool)`

SetCanSendVideoNotes sets CanSendVideoNotes field to given value.

### HasCanSendVideoNotes

`func (o *ChatPermissions) HasCanSendVideoNotes() bool`

HasCanSendVideoNotes returns a boolean if a field has been set.

### GetCanSendVoiceNotes

`func (o *ChatPermissions) GetCanSendVoiceNotes() bool`

GetCanSendVoiceNotes returns the CanSendVoiceNotes field if non-nil, zero value otherwise.

### GetCanSendVoiceNotesOk

`func (o *ChatPermissions) GetCanSendVoiceNotesOk() (*bool, bool)`

GetCanSendVoiceNotesOk returns a tuple with the CanSendVoiceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendVoiceNotes

`func (o *ChatPermissions) SetCanSendVoiceNotes(v bool)`

SetCanSendVoiceNotes sets CanSendVoiceNotes field to given value.

### HasCanSendVoiceNotes

`func (o *ChatPermissions) HasCanSendVoiceNotes() bool`

HasCanSendVoiceNotes returns a boolean if a field has been set.

### GetCanSendPolls

`func (o *ChatPermissions) GetCanSendPolls() bool`

GetCanSendPolls returns the CanSendPolls field if non-nil, zero value otherwise.

### GetCanSendPollsOk

`func (o *ChatPermissions) GetCanSendPollsOk() (*bool, bool)`

GetCanSendPollsOk returns a tuple with the CanSendPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPolls

`func (o *ChatPermissions) SetCanSendPolls(v bool)`

SetCanSendPolls sets CanSendPolls field to given value.

### HasCanSendPolls

`func (o *ChatPermissions) HasCanSendPolls() bool`

HasCanSendPolls returns a boolean if a field has been set.

### GetCanSendOtherMessages

`func (o *ChatPermissions) GetCanSendOtherMessages() bool`

GetCanSendOtherMessages returns the CanSendOtherMessages field if non-nil, zero value otherwise.

### GetCanSendOtherMessagesOk

`func (o *ChatPermissions) GetCanSendOtherMessagesOk() (*bool, bool)`

GetCanSendOtherMessagesOk returns a tuple with the CanSendOtherMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendOtherMessages

`func (o *ChatPermissions) SetCanSendOtherMessages(v bool)`

SetCanSendOtherMessages sets CanSendOtherMessages field to given value.

### HasCanSendOtherMessages

`func (o *ChatPermissions) HasCanSendOtherMessages() bool`

HasCanSendOtherMessages returns a boolean if a field has been set.

### GetCanAddWebPagePreviews

`func (o *ChatPermissions) GetCanAddWebPagePreviews() bool`

GetCanAddWebPagePreviews returns the CanAddWebPagePreviews field if non-nil, zero value otherwise.

### GetCanAddWebPagePreviewsOk

`func (o *ChatPermissions) GetCanAddWebPagePreviewsOk() (*bool, bool)`

GetCanAddWebPagePreviewsOk returns a tuple with the CanAddWebPagePreviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddWebPagePreviews

`func (o *ChatPermissions) SetCanAddWebPagePreviews(v bool)`

SetCanAddWebPagePreviews sets CanAddWebPagePreviews field to given value.

### HasCanAddWebPagePreviews

`func (o *ChatPermissions) HasCanAddWebPagePreviews() bool`

HasCanAddWebPagePreviews returns a boolean if a field has been set.

### GetCanChangeInfo

`func (o *ChatPermissions) GetCanChangeInfo() bool`

GetCanChangeInfo returns the CanChangeInfo field if non-nil, zero value otherwise.

### GetCanChangeInfoOk

`func (o *ChatPermissions) GetCanChangeInfoOk() (*bool, bool)`

GetCanChangeInfoOk returns a tuple with the CanChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeInfo

`func (o *ChatPermissions) SetCanChangeInfo(v bool)`

SetCanChangeInfo sets CanChangeInfo field to given value.

### HasCanChangeInfo

`func (o *ChatPermissions) HasCanChangeInfo() bool`

HasCanChangeInfo returns a boolean if a field has been set.

### GetCanInviteUsers

`func (o *ChatPermissions) GetCanInviteUsers() bool`

GetCanInviteUsers returns the CanInviteUsers field if non-nil, zero value otherwise.

### GetCanInviteUsersOk

`func (o *ChatPermissions) GetCanInviteUsersOk() (*bool, bool)`

GetCanInviteUsersOk returns a tuple with the CanInviteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsers

`func (o *ChatPermissions) SetCanInviteUsers(v bool)`

SetCanInviteUsers sets CanInviteUsers field to given value.

### HasCanInviteUsers

`func (o *ChatPermissions) HasCanInviteUsers() bool`

HasCanInviteUsers returns a boolean if a field has been set.

### GetCanPinMessages

`func (o *ChatPermissions) GetCanPinMessages() bool`

GetCanPinMessages returns the CanPinMessages field if non-nil, zero value otherwise.

### GetCanPinMessagesOk

`func (o *ChatPermissions) GetCanPinMessagesOk() (*bool, bool)`

GetCanPinMessagesOk returns a tuple with the CanPinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPinMessages

`func (o *ChatPermissions) SetCanPinMessages(v bool)`

SetCanPinMessages sets CanPinMessages field to given value.

### HasCanPinMessages

`func (o *ChatPermissions) HasCanPinMessages() bool`

HasCanPinMessages returns a boolean if a field has been set.

### GetCanManageTopics

`func (o *ChatPermissions) GetCanManageTopics() bool`

GetCanManageTopics returns the CanManageTopics field if non-nil, zero value otherwise.

### GetCanManageTopicsOk

`func (o *ChatPermissions) GetCanManageTopicsOk() (*bool, bool)`

GetCanManageTopicsOk returns a tuple with the CanManageTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTopics

`func (o *ChatPermissions) SetCanManageTopics(v bool)`

SetCanManageTopics sets CanManageTopics field to given value.

### HasCanManageTopics

`func (o *ChatPermissions) HasCanManageTopics() bool`

HasCanManageTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


