# ChatFullInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier. | 
**Type** | **string** | Type of the chat, can be either “private”, “group”, “supergroup” or “channel” | 
**Title** | Pointer to **string** | *Optional*. Title, for supergroups, channels and group chats | [optional] 
**Username** | Pointer to **string** | *Optional*. Username, for private chats, supergroups and channels if available | [optional] 
**FirstName** | Pointer to **string** | *Optional*. First name of the other party in a private chat | [optional] 
**LastName** | Pointer to **string** | *Optional*. Last name of the other party in a private chat | [optional] 
**IsForum** | Pointer to **bool** | *Optional*. *True*, if the supergroup chat is a forum (has [topics](https://telegram.org/blog/topics-in-groups-collectible-usernames#topics-in-groups) enabled) | [optional] [default to true]
**AccentColorId** | **int32** | Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See [accent colors](https://core.telegram.org/bots/api/#accent-colors) for more details. | 
**MaxReactionCount** | **int32** | The maximum number of reactions that can be set on a message in the chat | 
**Photo** | Pointer to [**ChatPhoto**](ChatPhoto.md) |  | [optional] 
**ActiveUsernames** | Pointer to **[]string** | *Optional*. If non-empty, the list of all [active chat usernames](https://telegram.org/blog/topics-in-groups-collectible-usernames#collectible-usernames); for private chats, supergroups and channels | [optional] 
**Birthdate** | Pointer to [**Birthdate**](Birthdate.md) |  | [optional] 
**BusinessIntro** | Pointer to [**BusinessIntro**](BusinessIntro.md) |  | [optional] 
**BusinessLocation** | Pointer to [**BusinessLocation**](BusinessLocation.md) |  | [optional] 
**BusinessOpeningHours** | Pointer to [**BusinessOpeningHours**](BusinessOpeningHours.md) |  | [optional] 
**PersonalChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**AvailableReactions** | Pointer to [**[]ReactionType**](ReactionType.md) | *Optional*. List of available reactions allowed in the chat. If omitted, then all [emoji reactions](https://core.telegram.org/bots/api/#reactiontypeemoji) are allowed. | [optional] 
**BackgroundCustomEmojiId** | Pointer to **string** | *Optional*. Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background | [optional] 
**ProfileAccentColorId** | Pointer to **int32** | *Optional*. Identifier of the accent color for the chat&#39;s profile background. See [profile accent colors](https://core.telegram.org/bots/api/#profile-accent-colors) for more details. | [optional] 
**ProfileBackgroundCustomEmojiId** | Pointer to **string** | *Optional*. Custom emoji identifier of the emoji chosen by the chat for its profile background | [optional] 
**EmojiStatusCustomEmojiId** | Pointer to **string** | *Optional*. Custom emoji identifier of the emoji status of the chat or the other party in a private chat | [optional] 
**EmojiStatusExpirationDate** | Pointer to **int32** | *Optional*. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any | [optional] 
**Bio** | Pointer to **string** | *Optional*. Bio of the other party in a private chat | [optional] 
**HasPrivateForwards** | Pointer to **bool** | *Optional*. *True*, if privacy settings of the other party in the private chat allows to use &#x60;tg://user?id&#x3D;&lt;user_id&gt;&#x60; links only in chats with the user | [optional] [default to true]
**HasRestrictedVoiceAndVideoMessages** | Pointer to **bool** | *Optional*. *True*, if the privacy settings of the other party restrict sending voice and video note messages in the private chat | [optional] [default to true]
**JoinToSendMessages** | Pointer to **bool** | *Optional*. *True*, if users need to join the supergroup before they can send messages | [optional] [default to true]
**JoinByRequest** | Pointer to **bool** | *Optional*. *True*, if all users directly joining the supergroup without using an invite link need to be approved by supergroup administrators | [optional] [default to true]
**Description** | Pointer to **string** | *Optional*. Description, for groups, supergroups and channel chats | [optional] 
**InviteLink** | Pointer to **string** | *Optional*. Primary invite link, for groups, supergroups and channel chats | [optional] 
**PinnedMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Permissions** | Pointer to [**ChatPermissions**](ChatPermissions.md) |  | [optional] 
**AcceptedGiftTypes** | [**AcceptedGiftTypes**](AcceptedGiftTypes.md) |  | 
**CanSendPaidMedia** | Pointer to **bool** | *Optional*. *True*, if paid media messages can be sent or forwarded to the channel chat. The field is available only for channel chats. | [optional] [default to true]
**SlowModeDelay** | Pointer to **int32** | *Optional*. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds | [optional] 
**UnrestrictBoostCount** | Pointer to **int32** | *Optional*. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions | [optional] 
**MessageAutoDeleteTime** | Pointer to **int32** | *Optional*. The time after which all messages sent to the chat will be automatically deleted; in seconds | [optional] 
**HasAggressiveAntiSpamEnabled** | Pointer to **bool** | *Optional*. *True*, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators. | [optional] [default to true]
**HasHiddenMembers** | Pointer to **bool** | *Optional*. *True*, if non-administrators can only get the list of bots and administrators in the chat | [optional] [default to true]
**HasProtectedContent** | Pointer to **bool** | *Optional*. *True*, if messages from the chat can&#39;t be forwarded to other chats | [optional] [default to true]
**HasVisibleHistory** | Pointer to **bool** | *Optional*. *True*, if new chat members will have access to old messages; available only to chat administrators | [optional] [default to true]
**StickerSetName** | Pointer to **string** | *Optional*. For supergroups, name of the group sticker set | [optional] 
**CanSetStickerSet** | Pointer to **bool** | *Optional*. *True*, if the bot can change the group sticker set | [optional] [default to true]
**CustomEmojiStickerSetName** | Pointer to **string** | *Optional*. For supergroups, the name of the group&#39;s custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group. | [optional] 
**LinkedChatId** | Pointer to **int32** | *Optional*. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. | [optional] 
**Location** | Pointer to [**ChatLocation**](ChatLocation.md) |  | [optional] 

## Methods

### NewChatFullInfo

`func NewChatFullInfo(id int32, type_ string, accentColorId int32, maxReactionCount int32, acceptedGiftTypes AcceptedGiftTypes, ) *ChatFullInfo`

NewChatFullInfo instantiates a new ChatFullInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatFullInfoWithDefaults

`func NewChatFullInfoWithDefaults() *ChatFullInfo`

NewChatFullInfoWithDefaults instantiates a new ChatFullInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatFullInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatFullInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatFullInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *ChatFullInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatFullInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatFullInfo) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ChatFullInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChatFullInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChatFullInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChatFullInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsername

`func (o *ChatFullInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChatFullInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChatFullInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ChatFullInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *ChatFullInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ChatFullInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ChatFullInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ChatFullInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ChatFullInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ChatFullInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ChatFullInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ChatFullInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsForum

`func (o *ChatFullInfo) GetIsForum() bool`

GetIsForum returns the IsForum field if non-nil, zero value otherwise.

### GetIsForumOk

`func (o *ChatFullInfo) GetIsForumOk() (*bool, bool)`

GetIsForumOk returns a tuple with the IsForum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForum

`func (o *ChatFullInfo) SetIsForum(v bool)`

SetIsForum sets IsForum field to given value.

### HasIsForum

`func (o *ChatFullInfo) HasIsForum() bool`

HasIsForum returns a boolean if a field has been set.

### GetAccentColorId

`func (o *ChatFullInfo) GetAccentColorId() int32`

GetAccentColorId returns the AccentColorId field if non-nil, zero value otherwise.

### GetAccentColorIdOk

`func (o *ChatFullInfo) GetAccentColorIdOk() (*int32, bool)`

GetAccentColorIdOk returns a tuple with the AccentColorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColorId

`func (o *ChatFullInfo) SetAccentColorId(v int32)`

SetAccentColorId sets AccentColorId field to given value.


### GetMaxReactionCount

`func (o *ChatFullInfo) GetMaxReactionCount() int32`

GetMaxReactionCount returns the MaxReactionCount field if non-nil, zero value otherwise.

### GetMaxReactionCountOk

`func (o *ChatFullInfo) GetMaxReactionCountOk() (*int32, bool)`

GetMaxReactionCountOk returns a tuple with the MaxReactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReactionCount

`func (o *ChatFullInfo) SetMaxReactionCount(v int32)`

SetMaxReactionCount sets MaxReactionCount field to given value.


### GetPhoto

`func (o *ChatFullInfo) GetPhoto() ChatPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ChatFullInfo) GetPhotoOk() (*ChatPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ChatFullInfo) SetPhoto(v ChatPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ChatFullInfo) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetActiveUsernames

`func (o *ChatFullInfo) GetActiveUsernames() []string`

GetActiveUsernames returns the ActiveUsernames field if non-nil, zero value otherwise.

### GetActiveUsernamesOk

`func (o *ChatFullInfo) GetActiveUsernamesOk() (*[]string, bool)`

GetActiveUsernamesOk returns a tuple with the ActiveUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsernames

`func (o *ChatFullInfo) SetActiveUsernames(v []string)`

SetActiveUsernames sets ActiveUsernames field to given value.

### HasActiveUsernames

`func (o *ChatFullInfo) HasActiveUsernames() bool`

HasActiveUsernames returns a boolean if a field has been set.

### GetBirthdate

`func (o *ChatFullInfo) GetBirthdate() Birthdate`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *ChatFullInfo) GetBirthdateOk() (*Birthdate, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *ChatFullInfo) SetBirthdate(v Birthdate)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *ChatFullInfo) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetBusinessIntro

`func (o *ChatFullInfo) GetBusinessIntro() BusinessIntro`

GetBusinessIntro returns the BusinessIntro field if non-nil, zero value otherwise.

### GetBusinessIntroOk

`func (o *ChatFullInfo) GetBusinessIntroOk() (*BusinessIntro, bool)`

GetBusinessIntroOk returns a tuple with the BusinessIntro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIntro

`func (o *ChatFullInfo) SetBusinessIntro(v BusinessIntro)`

SetBusinessIntro sets BusinessIntro field to given value.

### HasBusinessIntro

`func (o *ChatFullInfo) HasBusinessIntro() bool`

HasBusinessIntro returns a boolean if a field has been set.

### GetBusinessLocation

`func (o *ChatFullInfo) GetBusinessLocation() BusinessLocation`

GetBusinessLocation returns the BusinessLocation field if non-nil, zero value otherwise.

### GetBusinessLocationOk

`func (o *ChatFullInfo) GetBusinessLocationOk() (*BusinessLocation, bool)`

GetBusinessLocationOk returns a tuple with the BusinessLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLocation

`func (o *ChatFullInfo) SetBusinessLocation(v BusinessLocation)`

SetBusinessLocation sets BusinessLocation field to given value.

### HasBusinessLocation

`func (o *ChatFullInfo) HasBusinessLocation() bool`

HasBusinessLocation returns a boolean if a field has been set.

### GetBusinessOpeningHours

`func (o *ChatFullInfo) GetBusinessOpeningHours() BusinessOpeningHours`

GetBusinessOpeningHours returns the BusinessOpeningHours field if non-nil, zero value otherwise.

### GetBusinessOpeningHoursOk

`func (o *ChatFullInfo) GetBusinessOpeningHoursOk() (*BusinessOpeningHours, bool)`

GetBusinessOpeningHoursOk returns a tuple with the BusinessOpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessOpeningHours

`func (o *ChatFullInfo) SetBusinessOpeningHours(v BusinessOpeningHours)`

SetBusinessOpeningHours sets BusinessOpeningHours field to given value.

### HasBusinessOpeningHours

`func (o *ChatFullInfo) HasBusinessOpeningHours() bool`

HasBusinessOpeningHours returns a boolean if a field has been set.

### GetPersonalChat

`func (o *ChatFullInfo) GetPersonalChat() Chat`

GetPersonalChat returns the PersonalChat field if non-nil, zero value otherwise.

### GetPersonalChatOk

`func (o *ChatFullInfo) GetPersonalChatOk() (*Chat, bool)`

GetPersonalChatOk returns a tuple with the PersonalChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalChat

`func (o *ChatFullInfo) SetPersonalChat(v Chat)`

SetPersonalChat sets PersonalChat field to given value.

### HasPersonalChat

`func (o *ChatFullInfo) HasPersonalChat() bool`

HasPersonalChat returns a boolean if a field has been set.

### GetAvailableReactions

`func (o *ChatFullInfo) GetAvailableReactions() []ReactionType`

GetAvailableReactions returns the AvailableReactions field if non-nil, zero value otherwise.

### GetAvailableReactionsOk

`func (o *ChatFullInfo) GetAvailableReactionsOk() (*[]ReactionType, bool)`

GetAvailableReactionsOk returns a tuple with the AvailableReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReactions

`func (o *ChatFullInfo) SetAvailableReactions(v []ReactionType)`

SetAvailableReactions sets AvailableReactions field to given value.

### HasAvailableReactions

`func (o *ChatFullInfo) HasAvailableReactions() bool`

HasAvailableReactions returns a boolean if a field has been set.

### GetBackgroundCustomEmojiId

`func (o *ChatFullInfo) GetBackgroundCustomEmojiId() string`

GetBackgroundCustomEmojiId returns the BackgroundCustomEmojiId field if non-nil, zero value otherwise.

### GetBackgroundCustomEmojiIdOk

`func (o *ChatFullInfo) GetBackgroundCustomEmojiIdOk() (*string, bool)`

GetBackgroundCustomEmojiIdOk returns a tuple with the BackgroundCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundCustomEmojiId

`func (o *ChatFullInfo) SetBackgroundCustomEmojiId(v string)`

SetBackgroundCustomEmojiId sets BackgroundCustomEmojiId field to given value.

### HasBackgroundCustomEmojiId

`func (o *ChatFullInfo) HasBackgroundCustomEmojiId() bool`

HasBackgroundCustomEmojiId returns a boolean if a field has been set.

### GetProfileAccentColorId

`func (o *ChatFullInfo) GetProfileAccentColorId() int32`

GetProfileAccentColorId returns the ProfileAccentColorId field if non-nil, zero value otherwise.

### GetProfileAccentColorIdOk

`func (o *ChatFullInfo) GetProfileAccentColorIdOk() (*int32, bool)`

GetProfileAccentColorIdOk returns a tuple with the ProfileAccentColorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAccentColorId

`func (o *ChatFullInfo) SetProfileAccentColorId(v int32)`

SetProfileAccentColorId sets ProfileAccentColorId field to given value.

### HasProfileAccentColorId

`func (o *ChatFullInfo) HasProfileAccentColorId() bool`

HasProfileAccentColorId returns a boolean if a field has been set.

### GetProfileBackgroundCustomEmojiId

`func (o *ChatFullInfo) GetProfileBackgroundCustomEmojiId() string`

GetProfileBackgroundCustomEmojiId returns the ProfileBackgroundCustomEmojiId field if non-nil, zero value otherwise.

### GetProfileBackgroundCustomEmojiIdOk

`func (o *ChatFullInfo) GetProfileBackgroundCustomEmojiIdOk() (*string, bool)`

GetProfileBackgroundCustomEmojiIdOk returns a tuple with the ProfileBackgroundCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileBackgroundCustomEmojiId

`func (o *ChatFullInfo) SetProfileBackgroundCustomEmojiId(v string)`

SetProfileBackgroundCustomEmojiId sets ProfileBackgroundCustomEmojiId field to given value.

### HasProfileBackgroundCustomEmojiId

`func (o *ChatFullInfo) HasProfileBackgroundCustomEmojiId() bool`

HasProfileBackgroundCustomEmojiId returns a boolean if a field has been set.

### GetEmojiStatusCustomEmojiId

`func (o *ChatFullInfo) GetEmojiStatusCustomEmojiId() string`

GetEmojiStatusCustomEmojiId returns the EmojiStatusCustomEmojiId field if non-nil, zero value otherwise.

### GetEmojiStatusCustomEmojiIdOk

`func (o *ChatFullInfo) GetEmojiStatusCustomEmojiIdOk() (*string, bool)`

GetEmojiStatusCustomEmojiIdOk returns a tuple with the EmojiStatusCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusCustomEmojiId

`func (o *ChatFullInfo) SetEmojiStatusCustomEmojiId(v string)`

SetEmojiStatusCustomEmojiId sets EmojiStatusCustomEmojiId field to given value.

### HasEmojiStatusCustomEmojiId

`func (o *ChatFullInfo) HasEmojiStatusCustomEmojiId() bool`

HasEmojiStatusCustomEmojiId returns a boolean if a field has been set.

### GetEmojiStatusExpirationDate

`func (o *ChatFullInfo) GetEmojiStatusExpirationDate() int32`

GetEmojiStatusExpirationDate returns the EmojiStatusExpirationDate field if non-nil, zero value otherwise.

### GetEmojiStatusExpirationDateOk

`func (o *ChatFullInfo) GetEmojiStatusExpirationDateOk() (*int32, bool)`

GetEmojiStatusExpirationDateOk returns a tuple with the EmojiStatusExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusExpirationDate

`func (o *ChatFullInfo) SetEmojiStatusExpirationDate(v int32)`

SetEmojiStatusExpirationDate sets EmojiStatusExpirationDate field to given value.

### HasEmojiStatusExpirationDate

`func (o *ChatFullInfo) HasEmojiStatusExpirationDate() bool`

HasEmojiStatusExpirationDate returns a boolean if a field has been set.

### GetBio

`func (o *ChatFullInfo) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *ChatFullInfo) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *ChatFullInfo) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *ChatFullInfo) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetHasPrivateForwards

`func (o *ChatFullInfo) GetHasPrivateForwards() bool`

GetHasPrivateForwards returns the HasPrivateForwards field if non-nil, zero value otherwise.

### GetHasPrivateForwardsOk

`func (o *ChatFullInfo) GetHasPrivateForwardsOk() (*bool, bool)`

GetHasPrivateForwardsOk returns a tuple with the HasPrivateForwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateForwards

`func (o *ChatFullInfo) SetHasPrivateForwards(v bool)`

SetHasPrivateForwards sets HasPrivateForwards field to given value.

### HasHasPrivateForwards

`func (o *ChatFullInfo) HasHasPrivateForwards() bool`

HasHasPrivateForwards returns a boolean if a field has been set.

### GetHasRestrictedVoiceAndVideoMessages

`func (o *ChatFullInfo) GetHasRestrictedVoiceAndVideoMessages() bool`

GetHasRestrictedVoiceAndVideoMessages returns the HasRestrictedVoiceAndVideoMessages field if non-nil, zero value otherwise.

### GetHasRestrictedVoiceAndVideoMessagesOk

`func (o *ChatFullInfo) GetHasRestrictedVoiceAndVideoMessagesOk() (*bool, bool)`

GetHasRestrictedVoiceAndVideoMessagesOk returns a tuple with the HasRestrictedVoiceAndVideoMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRestrictedVoiceAndVideoMessages

`func (o *ChatFullInfo) SetHasRestrictedVoiceAndVideoMessages(v bool)`

SetHasRestrictedVoiceAndVideoMessages sets HasRestrictedVoiceAndVideoMessages field to given value.

### HasHasRestrictedVoiceAndVideoMessages

`func (o *ChatFullInfo) HasHasRestrictedVoiceAndVideoMessages() bool`

HasHasRestrictedVoiceAndVideoMessages returns a boolean if a field has been set.

### GetJoinToSendMessages

`func (o *ChatFullInfo) GetJoinToSendMessages() bool`

GetJoinToSendMessages returns the JoinToSendMessages field if non-nil, zero value otherwise.

### GetJoinToSendMessagesOk

`func (o *ChatFullInfo) GetJoinToSendMessagesOk() (*bool, bool)`

GetJoinToSendMessagesOk returns a tuple with the JoinToSendMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinToSendMessages

`func (o *ChatFullInfo) SetJoinToSendMessages(v bool)`

SetJoinToSendMessages sets JoinToSendMessages field to given value.

### HasJoinToSendMessages

`func (o *ChatFullInfo) HasJoinToSendMessages() bool`

HasJoinToSendMessages returns a boolean if a field has been set.

### GetJoinByRequest

`func (o *ChatFullInfo) GetJoinByRequest() bool`

GetJoinByRequest returns the JoinByRequest field if non-nil, zero value otherwise.

### GetJoinByRequestOk

`func (o *ChatFullInfo) GetJoinByRequestOk() (*bool, bool)`

GetJoinByRequestOk returns a tuple with the JoinByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinByRequest

`func (o *ChatFullInfo) SetJoinByRequest(v bool)`

SetJoinByRequest sets JoinByRequest field to given value.

### HasJoinByRequest

`func (o *ChatFullInfo) HasJoinByRequest() bool`

HasJoinByRequest returns a boolean if a field has been set.

### GetDescription

`func (o *ChatFullInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChatFullInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChatFullInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChatFullInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInviteLink

`func (o *ChatFullInfo) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ChatFullInfo) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ChatFullInfo) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *ChatFullInfo) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetPinnedMessage

`func (o *ChatFullInfo) GetPinnedMessage() Message`

GetPinnedMessage returns the PinnedMessage field if non-nil, zero value otherwise.

### GetPinnedMessageOk

`func (o *ChatFullInfo) GetPinnedMessageOk() (*Message, bool)`

GetPinnedMessageOk returns a tuple with the PinnedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedMessage

`func (o *ChatFullInfo) SetPinnedMessage(v Message)`

SetPinnedMessage sets PinnedMessage field to given value.

### HasPinnedMessage

`func (o *ChatFullInfo) HasPinnedMessage() bool`

HasPinnedMessage returns a boolean if a field has been set.

### GetPermissions

`func (o *ChatFullInfo) GetPermissions() ChatPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ChatFullInfo) GetPermissionsOk() (*ChatPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ChatFullInfo) SetPermissions(v ChatPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ChatFullInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAcceptedGiftTypes

`func (o *ChatFullInfo) GetAcceptedGiftTypes() AcceptedGiftTypes`

GetAcceptedGiftTypes returns the AcceptedGiftTypes field if non-nil, zero value otherwise.

### GetAcceptedGiftTypesOk

`func (o *ChatFullInfo) GetAcceptedGiftTypesOk() (*AcceptedGiftTypes, bool)`

GetAcceptedGiftTypesOk returns a tuple with the AcceptedGiftTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedGiftTypes

`func (o *ChatFullInfo) SetAcceptedGiftTypes(v AcceptedGiftTypes)`

SetAcceptedGiftTypes sets AcceptedGiftTypes field to given value.


### GetCanSendPaidMedia

`func (o *ChatFullInfo) GetCanSendPaidMedia() bool`

GetCanSendPaidMedia returns the CanSendPaidMedia field if non-nil, zero value otherwise.

### GetCanSendPaidMediaOk

`func (o *ChatFullInfo) GetCanSendPaidMediaOk() (*bool, bool)`

GetCanSendPaidMediaOk returns a tuple with the CanSendPaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendPaidMedia

`func (o *ChatFullInfo) SetCanSendPaidMedia(v bool)`

SetCanSendPaidMedia sets CanSendPaidMedia field to given value.

### HasCanSendPaidMedia

`func (o *ChatFullInfo) HasCanSendPaidMedia() bool`

HasCanSendPaidMedia returns a boolean if a field has been set.

### GetSlowModeDelay

`func (o *ChatFullInfo) GetSlowModeDelay() int32`

GetSlowModeDelay returns the SlowModeDelay field if non-nil, zero value otherwise.

### GetSlowModeDelayOk

`func (o *ChatFullInfo) GetSlowModeDelayOk() (*int32, bool)`

GetSlowModeDelayOk returns a tuple with the SlowModeDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowModeDelay

`func (o *ChatFullInfo) SetSlowModeDelay(v int32)`

SetSlowModeDelay sets SlowModeDelay field to given value.

### HasSlowModeDelay

`func (o *ChatFullInfo) HasSlowModeDelay() bool`

HasSlowModeDelay returns a boolean if a field has been set.

### GetUnrestrictBoostCount

`func (o *ChatFullInfo) GetUnrestrictBoostCount() int32`

GetUnrestrictBoostCount returns the UnrestrictBoostCount field if non-nil, zero value otherwise.

### GetUnrestrictBoostCountOk

`func (o *ChatFullInfo) GetUnrestrictBoostCountOk() (*int32, bool)`

GetUnrestrictBoostCountOk returns a tuple with the UnrestrictBoostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrestrictBoostCount

`func (o *ChatFullInfo) SetUnrestrictBoostCount(v int32)`

SetUnrestrictBoostCount sets UnrestrictBoostCount field to given value.

### HasUnrestrictBoostCount

`func (o *ChatFullInfo) HasUnrestrictBoostCount() bool`

HasUnrestrictBoostCount returns a boolean if a field has been set.

### GetMessageAutoDeleteTime

`func (o *ChatFullInfo) GetMessageAutoDeleteTime() int32`

GetMessageAutoDeleteTime returns the MessageAutoDeleteTime field if non-nil, zero value otherwise.

### GetMessageAutoDeleteTimeOk

`func (o *ChatFullInfo) GetMessageAutoDeleteTimeOk() (*int32, bool)`

GetMessageAutoDeleteTimeOk returns a tuple with the MessageAutoDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageAutoDeleteTime

`func (o *ChatFullInfo) SetMessageAutoDeleteTime(v int32)`

SetMessageAutoDeleteTime sets MessageAutoDeleteTime field to given value.

### HasMessageAutoDeleteTime

`func (o *ChatFullInfo) HasMessageAutoDeleteTime() bool`

HasMessageAutoDeleteTime returns a boolean if a field has been set.

### GetHasAggressiveAntiSpamEnabled

`func (o *ChatFullInfo) GetHasAggressiveAntiSpamEnabled() bool`

GetHasAggressiveAntiSpamEnabled returns the HasAggressiveAntiSpamEnabled field if non-nil, zero value otherwise.

### GetHasAggressiveAntiSpamEnabledOk

`func (o *ChatFullInfo) GetHasAggressiveAntiSpamEnabledOk() (*bool, bool)`

GetHasAggressiveAntiSpamEnabledOk returns a tuple with the HasAggressiveAntiSpamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAggressiveAntiSpamEnabled

`func (o *ChatFullInfo) SetHasAggressiveAntiSpamEnabled(v bool)`

SetHasAggressiveAntiSpamEnabled sets HasAggressiveAntiSpamEnabled field to given value.

### HasHasAggressiveAntiSpamEnabled

`func (o *ChatFullInfo) HasHasAggressiveAntiSpamEnabled() bool`

HasHasAggressiveAntiSpamEnabled returns a boolean if a field has been set.

### GetHasHiddenMembers

`func (o *ChatFullInfo) GetHasHiddenMembers() bool`

GetHasHiddenMembers returns the HasHiddenMembers field if non-nil, zero value otherwise.

### GetHasHiddenMembersOk

`func (o *ChatFullInfo) GetHasHiddenMembersOk() (*bool, bool)`

GetHasHiddenMembersOk returns a tuple with the HasHiddenMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHiddenMembers

`func (o *ChatFullInfo) SetHasHiddenMembers(v bool)`

SetHasHiddenMembers sets HasHiddenMembers field to given value.

### HasHasHiddenMembers

`func (o *ChatFullInfo) HasHasHiddenMembers() bool`

HasHasHiddenMembers returns a boolean if a field has been set.

### GetHasProtectedContent

`func (o *ChatFullInfo) GetHasProtectedContent() bool`

GetHasProtectedContent returns the HasProtectedContent field if non-nil, zero value otherwise.

### GetHasProtectedContentOk

`func (o *ChatFullInfo) GetHasProtectedContentOk() (*bool, bool)`

GetHasProtectedContentOk returns a tuple with the HasProtectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProtectedContent

`func (o *ChatFullInfo) SetHasProtectedContent(v bool)`

SetHasProtectedContent sets HasProtectedContent field to given value.

### HasHasProtectedContent

`func (o *ChatFullInfo) HasHasProtectedContent() bool`

HasHasProtectedContent returns a boolean if a field has been set.

### GetHasVisibleHistory

`func (o *ChatFullInfo) GetHasVisibleHistory() bool`

GetHasVisibleHistory returns the HasVisibleHistory field if non-nil, zero value otherwise.

### GetHasVisibleHistoryOk

`func (o *ChatFullInfo) GetHasVisibleHistoryOk() (*bool, bool)`

GetHasVisibleHistoryOk returns a tuple with the HasVisibleHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVisibleHistory

`func (o *ChatFullInfo) SetHasVisibleHistory(v bool)`

SetHasVisibleHistory sets HasVisibleHistory field to given value.

### HasHasVisibleHistory

`func (o *ChatFullInfo) HasHasVisibleHistory() bool`

HasHasVisibleHistory returns a boolean if a field has been set.

### GetStickerSetName

`func (o *ChatFullInfo) GetStickerSetName() string`

GetStickerSetName returns the StickerSetName field if non-nil, zero value otherwise.

### GetStickerSetNameOk

`func (o *ChatFullInfo) GetStickerSetNameOk() (*string, bool)`

GetStickerSetNameOk returns a tuple with the StickerSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerSetName

`func (o *ChatFullInfo) SetStickerSetName(v string)`

SetStickerSetName sets StickerSetName field to given value.

### HasStickerSetName

`func (o *ChatFullInfo) HasStickerSetName() bool`

HasStickerSetName returns a boolean if a field has been set.

### GetCanSetStickerSet

`func (o *ChatFullInfo) GetCanSetStickerSet() bool`

GetCanSetStickerSet returns the CanSetStickerSet field if non-nil, zero value otherwise.

### GetCanSetStickerSetOk

`func (o *ChatFullInfo) GetCanSetStickerSetOk() (*bool, bool)`

GetCanSetStickerSetOk returns a tuple with the CanSetStickerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetStickerSet

`func (o *ChatFullInfo) SetCanSetStickerSet(v bool)`

SetCanSetStickerSet sets CanSetStickerSet field to given value.

### HasCanSetStickerSet

`func (o *ChatFullInfo) HasCanSetStickerSet() bool`

HasCanSetStickerSet returns a boolean if a field has been set.

### GetCustomEmojiStickerSetName

`func (o *ChatFullInfo) GetCustomEmojiStickerSetName() string`

GetCustomEmojiStickerSetName returns the CustomEmojiStickerSetName field if non-nil, zero value otherwise.

### GetCustomEmojiStickerSetNameOk

`func (o *ChatFullInfo) GetCustomEmojiStickerSetNameOk() (*string, bool)`

GetCustomEmojiStickerSetNameOk returns a tuple with the CustomEmojiStickerSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiStickerSetName

`func (o *ChatFullInfo) SetCustomEmojiStickerSetName(v string)`

SetCustomEmojiStickerSetName sets CustomEmojiStickerSetName field to given value.

### HasCustomEmojiStickerSetName

`func (o *ChatFullInfo) HasCustomEmojiStickerSetName() bool`

HasCustomEmojiStickerSetName returns a boolean if a field has been set.

### GetLinkedChatId

`func (o *ChatFullInfo) GetLinkedChatId() int32`

GetLinkedChatId returns the LinkedChatId field if non-nil, zero value otherwise.

### GetLinkedChatIdOk

`func (o *ChatFullInfo) GetLinkedChatIdOk() (*int32, bool)`

GetLinkedChatIdOk returns a tuple with the LinkedChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedChatId

`func (o *ChatFullInfo) SetLinkedChatId(v int32)`

SetLinkedChatId sets LinkedChatId field to given value.

### HasLinkedChatId

`func (o *ChatFullInfo) HasLinkedChatId() bool`

HasLinkedChatId returns a boolean if a field has been set.

### GetLocation

`func (o *ChatFullInfo) GetLocation() ChatLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChatFullInfo) GetLocationOk() (*ChatLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChatFullInfo) SetLocation(v ChatLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChatFullInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


