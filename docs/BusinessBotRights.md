# BusinessBotRights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanReply** | Pointer to **bool** | *Optional*. True, if the bot can send and edit messages in the private chats that had incoming messages in the last 24 hours | [optional] [default to true]
**CanReadMessages** | Pointer to **bool** | *Optional*. True, if the bot can mark incoming private messages as read | [optional] [default to true]
**CanDeleteOutgoingMessages** | Pointer to **bool** | *Optional*. True, if the bot can delete messages sent by the bot | [optional] [default to true]
**CanDeleteAllMessages** | Pointer to **bool** | *Optional*. True, if the bot can delete all private messages in managed chats | [optional] [default to true]
**CanEditName** | Pointer to **bool** | *Optional*. True, if the bot can edit the first and last name of the business account | [optional] [default to true]
**CanEditBio** | Pointer to **bool** | *Optional*. True, if the bot can edit the bio of the business account | [optional] [default to true]
**CanEditProfilePhoto** | Pointer to **bool** | *Optional*. True, if the bot can edit the profile photo of the business account | [optional] [default to true]
**CanEditUsername** | Pointer to **bool** | *Optional*. True, if the bot can edit the username of the business account | [optional] [default to true]
**CanChangeGiftSettings** | Pointer to **bool** | *Optional*. True, if the bot can change the privacy settings pertaining to gifts for the business account | [optional] [default to true]
**CanViewGiftsAndStars** | Pointer to **bool** | *Optional*. True, if the bot can view gifts and the amount of Telegram Stars owned by the business account | [optional] [default to true]
**CanConvertGiftsToStars** | Pointer to **bool** | *Optional*. True, if the bot can convert regular gifts owned by the business account to Telegram Stars | [optional] [default to true]
**CanTransferAndUpgradeGifts** | Pointer to **bool** | *Optional*. True, if the bot can transfer and upgrade gifts owned by the business account | [optional] [default to true]
**CanTransferStars** | Pointer to **bool** | *Optional*. True, if the bot can transfer Telegram Stars received by the business account to its own account, or use them to upgrade and transfer gifts | [optional] [default to true]
**CanManageStories** | Pointer to **bool** | *Optional*. True, if the bot can post, edit and delete stories on behalf of the business account | [optional] [default to true]

## Methods

### NewBusinessBotRights

`func NewBusinessBotRights() *BusinessBotRights`

NewBusinessBotRights instantiates a new BusinessBotRights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessBotRightsWithDefaults

`func NewBusinessBotRightsWithDefaults() *BusinessBotRights`

NewBusinessBotRightsWithDefaults instantiates a new BusinessBotRights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanReply

`func (o *BusinessBotRights) GetCanReply() bool`

GetCanReply returns the CanReply field if non-nil, zero value otherwise.

### GetCanReplyOk

`func (o *BusinessBotRights) GetCanReplyOk() (*bool, bool)`

GetCanReplyOk returns a tuple with the CanReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReply

`func (o *BusinessBotRights) SetCanReply(v bool)`

SetCanReply sets CanReply field to given value.

### HasCanReply

`func (o *BusinessBotRights) HasCanReply() bool`

HasCanReply returns a boolean if a field has been set.

### GetCanReadMessages

`func (o *BusinessBotRights) GetCanReadMessages() bool`

GetCanReadMessages returns the CanReadMessages field if non-nil, zero value otherwise.

### GetCanReadMessagesOk

`func (o *BusinessBotRights) GetCanReadMessagesOk() (*bool, bool)`

GetCanReadMessagesOk returns a tuple with the CanReadMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReadMessages

`func (o *BusinessBotRights) SetCanReadMessages(v bool)`

SetCanReadMessages sets CanReadMessages field to given value.

### HasCanReadMessages

`func (o *BusinessBotRights) HasCanReadMessages() bool`

HasCanReadMessages returns a boolean if a field has been set.

### GetCanDeleteOutgoingMessages

`func (o *BusinessBotRights) GetCanDeleteOutgoingMessages() bool`

GetCanDeleteOutgoingMessages returns the CanDeleteOutgoingMessages field if non-nil, zero value otherwise.

### GetCanDeleteOutgoingMessagesOk

`func (o *BusinessBotRights) GetCanDeleteOutgoingMessagesOk() (*bool, bool)`

GetCanDeleteOutgoingMessagesOk returns a tuple with the CanDeleteOutgoingMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOutgoingMessages

`func (o *BusinessBotRights) SetCanDeleteOutgoingMessages(v bool)`

SetCanDeleteOutgoingMessages sets CanDeleteOutgoingMessages field to given value.

### HasCanDeleteOutgoingMessages

`func (o *BusinessBotRights) HasCanDeleteOutgoingMessages() bool`

HasCanDeleteOutgoingMessages returns a boolean if a field has been set.

### GetCanDeleteAllMessages

`func (o *BusinessBotRights) GetCanDeleteAllMessages() bool`

GetCanDeleteAllMessages returns the CanDeleteAllMessages field if non-nil, zero value otherwise.

### GetCanDeleteAllMessagesOk

`func (o *BusinessBotRights) GetCanDeleteAllMessagesOk() (*bool, bool)`

GetCanDeleteAllMessagesOk returns a tuple with the CanDeleteAllMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAllMessages

`func (o *BusinessBotRights) SetCanDeleteAllMessages(v bool)`

SetCanDeleteAllMessages sets CanDeleteAllMessages field to given value.

### HasCanDeleteAllMessages

`func (o *BusinessBotRights) HasCanDeleteAllMessages() bool`

HasCanDeleteAllMessages returns a boolean if a field has been set.

### GetCanEditName

`func (o *BusinessBotRights) GetCanEditName() bool`

GetCanEditName returns the CanEditName field if non-nil, zero value otherwise.

### GetCanEditNameOk

`func (o *BusinessBotRights) GetCanEditNameOk() (*bool, bool)`

GetCanEditNameOk returns a tuple with the CanEditName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditName

`func (o *BusinessBotRights) SetCanEditName(v bool)`

SetCanEditName sets CanEditName field to given value.

### HasCanEditName

`func (o *BusinessBotRights) HasCanEditName() bool`

HasCanEditName returns a boolean if a field has been set.

### GetCanEditBio

`func (o *BusinessBotRights) GetCanEditBio() bool`

GetCanEditBio returns the CanEditBio field if non-nil, zero value otherwise.

### GetCanEditBioOk

`func (o *BusinessBotRights) GetCanEditBioOk() (*bool, bool)`

GetCanEditBioOk returns a tuple with the CanEditBio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditBio

`func (o *BusinessBotRights) SetCanEditBio(v bool)`

SetCanEditBio sets CanEditBio field to given value.

### HasCanEditBio

`func (o *BusinessBotRights) HasCanEditBio() bool`

HasCanEditBio returns a boolean if a field has been set.

### GetCanEditProfilePhoto

`func (o *BusinessBotRights) GetCanEditProfilePhoto() bool`

GetCanEditProfilePhoto returns the CanEditProfilePhoto field if non-nil, zero value otherwise.

### GetCanEditProfilePhotoOk

`func (o *BusinessBotRights) GetCanEditProfilePhotoOk() (*bool, bool)`

GetCanEditProfilePhotoOk returns a tuple with the CanEditProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditProfilePhoto

`func (o *BusinessBotRights) SetCanEditProfilePhoto(v bool)`

SetCanEditProfilePhoto sets CanEditProfilePhoto field to given value.

### HasCanEditProfilePhoto

`func (o *BusinessBotRights) HasCanEditProfilePhoto() bool`

HasCanEditProfilePhoto returns a boolean if a field has been set.

### GetCanEditUsername

`func (o *BusinessBotRights) GetCanEditUsername() bool`

GetCanEditUsername returns the CanEditUsername field if non-nil, zero value otherwise.

### GetCanEditUsernameOk

`func (o *BusinessBotRights) GetCanEditUsernameOk() (*bool, bool)`

GetCanEditUsernameOk returns a tuple with the CanEditUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditUsername

`func (o *BusinessBotRights) SetCanEditUsername(v bool)`

SetCanEditUsername sets CanEditUsername field to given value.

### HasCanEditUsername

`func (o *BusinessBotRights) HasCanEditUsername() bool`

HasCanEditUsername returns a boolean if a field has been set.

### GetCanChangeGiftSettings

`func (o *BusinessBotRights) GetCanChangeGiftSettings() bool`

GetCanChangeGiftSettings returns the CanChangeGiftSettings field if non-nil, zero value otherwise.

### GetCanChangeGiftSettingsOk

`func (o *BusinessBotRights) GetCanChangeGiftSettingsOk() (*bool, bool)`

GetCanChangeGiftSettingsOk returns a tuple with the CanChangeGiftSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeGiftSettings

`func (o *BusinessBotRights) SetCanChangeGiftSettings(v bool)`

SetCanChangeGiftSettings sets CanChangeGiftSettings field to given value.

### HasCanChangeGiftSettings

`func (o *BusinessBotRights) HasCanChangeGiftSettings() bool`

HasCanChangeGiftSettings returns a boolean if a field has been set.

### GetCanViewGiftsAndStars

`func (o *BusinessBotRights) GetCanViewGiftsAndStars() bool`

GetCanViewGiftsAndStars returns the CanViewGiftsAndStars field if non-nil, zero value otherwise.

### GetCanViewGiftsAndStarsOk

`func (o *BusinessBotRights) GetCanViewGiftsAndStarsOk() (*bool, bool)`

GetCanViewGiftsAndStarsOk returns a tuple with the CanViewGiftsAndStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewGiftsAndStars

`func (o *BusinessBotRights) SetCanViewGiftsAndStars(v bool)`

SetCanViewGiftsAndStars sets CanViewGiftsAndStars field to given value.

### HasCanViewGiftsAndStars

`func (o *BusinessBotRights) HasCanViewGiftsAndStars() bool`

HasCanViewGiftsAndStars returns a boolean if a field has been set.

### GetCanConvertGiftsToStars

`func (o *BusinessBotRights) GetCanConvertGiftsToStars() bool`

GetCanConvertGiftsToStars returns the CanConvertGiftsToStars field if non-nil, zero value otherwise.

### GetCanConvertGiftsToStarsOk

`func (o *BusinessBotRights) GetCanConvertGiftsToStarsOk() (*bool, bool)`

GetCanConvertGiftsToStarsOk returns a tuple with the CanConvertGiftsToStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConvertGiftsToStars

`func (o *BusinessBotRights) SetCanConvertGiftsToStars(v bool)`

SetCanConvertGiftsToStars sets CanConvertGiftsToStars field to given value.

### HasCanConvertGiftsToStars

`func (o *BusinessBotRights) HasCanConvertGiftsToStars() bool`

HasCanConvertGiftsToStars returns a boolean if a field has been set.

### GetCanTransferAndUpgradeGifts

`func (o *BusinessBotRights) GetCanTransferAndUpgradeGifts() bool`

GetCanTransferAndUpgradeGifts returns the CanTransferAndUpgradeGifts field if non-nil, zero value otherwise.

### GetCanTransferAndUpgradeGiftsOk

`func (o *BusinessBotRights) GetCanTransferAndUpgradeGiftsOk() (*bool, bool)`

GetCanTransferAndUpgradeGiftsOk returns a tuple with the CanTransferAndUpgradeGifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransferAndUpgradeGifts

`func (o *BusinessBotRights) SetCanTransferAndUpgradeGifts(v bool)`

SetCanTransferAndUpgradeGifts sets CanTransferAndUpgradeGifts field to given value.

### HasCanTransferAndUpgradeGifts

`func (o *BusinessBotRights) HasCanTransferAndUpgradeGifts() bool`

HasCanTransferAndUpgradeGifts returns a boolean if a field has been set.

### GetCanTransferStars

`func (o *BusinessBotRights) GetCanTransferStars() bool`

GetCanTransferStars returns the CanTransferStars field if non-nil, zero value otherwise.

### GetCanTransferStarsOk

`func (o *BusinessBotRights) GetCanTransferStarsOk() (*bool, bool)`

GetCanTransferStarsOk returns a tuple with the CanTransferStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransferStars

`func (o *BusinessBotRights) SetCanTransferStars(v bool)`

SetCanTransferStars sets CanTransferStars field to given value.

### HasCanTransferStars

`func (o *BusinessBotRights) HasCanTransferStars() bool`

HasCanTransferStars returns a boolean if a field has been set.

### GetCanManageStories

`func (o *BusinessBotRights) GetCanManageStories() bool`

GetCanManageStories returns the CanManageStories field if non-nil, zero value otherwise.

### GetCanManageStoriesOk

`func (o *BusinessBotRights) GetCanManageStoriesOk() (*bool, bool)`

GetCanManageStoriesOk returns a tuple with the CanManageStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageStories

`func (o *BusinessBotRights) SetCanManageStories(v bool)`

SetCanManageStories sets CanManageStories field to given value.

### HasCanManageStories

`func (o *BusinessBotRights) HasCanManageStories() bool`

HasCanManageStories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


