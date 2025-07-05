# MaybeInaccessibleMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | Unique message identifier inside the chat | 
**MessageThreadId** | Pointer to **int32** | *Optional*. Unique identifier of a message thread to which the message belongs; for supergroups only | [optional] 
**From** | Pointer to [**User**](User.md) |  | [optional] 
**SenderChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**SenderBoostCount** | Pointer to **int32** | *Optional*. If the sender of the message boosted the chat, the number of boosts added by the user | [optional] 
**SenderBusinessBot** | Pointer to [**User**](User.md) |  | [optional] 
**Date** | **int32** | Always 0. The field can be used to differentiate regular and inaccessible messages. | 
**BusinessConnectionId** | Pointer to **string** | *Optional*. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier. | [optional] 
**Chat** | [**Chat**](Chat.md) |  | 
**ForwardOrigin** | Pointer to [**MessageOrigin**](MessageOrigin.md) |  | [optional] 
**IsTopicMessage** | Pointer to **bool** | *Optional*. *True*, if the message is sent to a forum topic | [optional] [default to true]
**IsAutomaticForward** | Pointer to **bool** | *Optional*. *True*, if the message is a channel post that was automatically forwarded to the connected discussion group | [optional] [default to true]
**ReplyToMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**ExternalReply** | Pointer to [**ExternalReplyInfo**](ExternalReplyInfo.md) |  | [optional] 
**Quote** | Pointer to [**TextQuote**](TextQuote.md) |  | [optional] 
**ReplyToStory** | Pointer to [**Story**](Story.md) |  | [optional] 
**ViaBot** | Pointer to [**User**](User.md) |  | [optional] 
**EditDate** | Pointer to **int32** | *Optional*. Date the message was last edited in Unix time | [optional] 
**HasProtectedContent** | Pointer to **bool** | *Optional*. *True*, if the message can&#39;t be forwarded | [optional] [default to true]
**IsFromOffline** | Pointer to **bool** | *Optional*. *True*, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message | [optional] [default to true]
**MediaGroupId** | Pointer to **string** | *Optional*. The unique identifier of a media message group this message belongs to | [optional] 
**AuthorSignature** | Pointer to **string** | *Optional*. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator | [optional] 
**PaidStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars that were paid by the sender of the message to send it | [optional] 
**Text** | Pointer to **string** | *Optional*. For text messages, the actual UTF-8 text of the message | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text | [optional] 
**LinkPreviewOptions** | Pointer to [**LinkPreviewOptions**](LinkPreviewOptions.md) |  | [optional] 
**EffectId** | Pointer to **string** | *Optional*. Unique identifier of the message effect added to the message | [optional] 
**Animation** | Pointer to [**Animation**](Animation.md) |  | [optional] 
**Audio** | Pointer to [**Audio**](Audio.md) |  | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**PaidMedia** | Pointer to [**PaidMediaInfo**](PaidMediaInfo.md) |  | [optional] 
**Photo** | Pointer to [**[]PhotoSize**](PhotoSize.md) | *Optional*. Message is a photo, available sizes of the photo | [optional] 
**Sticker** | Pointer to [**Sticker**](Sticker.md) |  | [optional] 
**Story** | Pointer to [**Story**](Story.md) |  | [optional] 
**Video** | Pointer to [**Video**](Video.md) |  | [optional] 
**VideoNote** | Pointer to [**VideoNote**](VideoNote.md) |  | [optional] 
**Voice** | Pointer to [**Voice**](Voice.md) |  | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption for the animation, audio, document, paid media, photo, video or voice | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. *True*, if the caption must be shown above the message media | [optional] [default to true]
**HasMediaSpoiler** | Pointer to **bool** | *Optional*. *True*, if the message media is covered by a spoiler animation | [optional] [default to true]
**Checklist** | Pointer to [**Checklist**](Checklist.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Dice** | Pointer to [**Dice**](Dice.md) |  | [optional] 
**Game** | Pointer to [**Game**](Game.md) |  | [optional] 
**Poll** | Pointer to [**Poll**](Poll.md) |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**NewChatMembers** | Pointer to [**[]User**](User.md) | *Optional*. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members) | [optional] 
**LeftChatMember** | Pointer to [**User**](User.md) |  | [optional] 
**NewChatTitle** | Pointer to **string** | *Optional*. A chat title was changed to this value | [optional] 
**NewChatPhoto** | Pointer to [**[]PhotoSize**](PhotoSize.md) | *Optional*. A chat photo was change to this value | [optional] 
**DeleteChatPhoto** | Pointer to **bool** | *Optional*. Service message: the chat photo was deleted | [optional] [default to true]
**GroupChatCreated** | Pointer to **bool** | *Optional*. Service message: the group has been created | [optional] [default to true]
**SupergroupChatCreated** | Pointer to **bool** | *Optional*. Service message: the supergroup has been created. This field can&#39;t be received in a message coming through updates, because bot can&#39;t be a member of a supergroup when it is created. It can only be found in reply\\_to\\_message if someone replies to a very first message in a directly created supergroup. | [optional] [default to true]
**ChannelChatCreated** | Pointer to **bool** | *Optional*. Service message: the channel has been created. This field can&#39;t be received in a message coming through updates, because bot can&#39;t be a member of a channel when it is created. It can only be found in reply\\_to\\_message if someone replies to a very first message in a channel. | [optional] [default to true]
**MessageAutoDeleteTimerChanged** | Pointer to [**MessageAutoDeleteTimerChanged**](MessageAutoDeleteTimerChanged.md) |  | [optional] 
**MigrateToChatId** | Pointer to **int32** | *Optional*. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier. | [optional] 
**MigrateFromChatId** | Pointer to **int32** | *Optional*. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier. | [optional] 
**PinnedMessage** | Pointer to [**MaybeInaccessibleMessage**](MaybeInaccessibleMessage.md) |  | [optional] 
**Invoice** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**SuccessfulPayment** | Pointer to [**SuccessfulPayment**](SuccessfulPayment.md) |  | [optional] 
**RefundedPayment** | Pointer to [**RefundedPayment**](RefundedPayment.md) |  | [optional] 
**UsersShared** | Pointer to [**UsersShared**](UsersShared.md) |  | [optional] 
**ChatShared** | Pointer to [**ChatShared**](ChatShared.md) |  | [optional] 
**Gift** | Pointer to [**GiftInfo**](GiftInfo.md) |  | [optional] 
**UniqueGift** | Pointer to [**UniqueGiftInfo**](UniqueGiftInfo.md) |  | [optional] 
**ConnectedWebsite** | Pointer to **string** | *Optional*. The domain name of the website on which the user has logged in. [More about Telegram Login »](https://core.telegram.org/widgets/login) | [optional] 
**WriteAccessAllowed** | Pointer to [**WriteAccessAllowed**](WriteAccessAllowed.md) |  | [optional] 
**PassportData** | Pointer to [**PassportData**](PassportData.md) |  | [optional] 
**ProximityAlertTriggered** | Pointer to [**ProximityAlertTriggered**](ProximityAlertTriggered.md) |  | [optional] 
**BoostAdded** | Pointer to [**ChatBoostAdded**](ChatBoostAdded.md) |  | [optional] 
**ChatBackgroundSet** | Pointer to [**ChatBackground**](ChatBackground.md) |  | [optional] 
**ChecklistTasksDone** | Pointer to [**ChecklistTasksDone**](ChecklistTasksDone.md) |  | [optional] 
**ChecklistTasksAdded** | Pointer to [**ChecklistTasksAdded**](ChecklistTasksAdded.md) |  | [optional] 
**DirectMessagePriceChanged** | Pointer to [**DirectMessagePriceChanged**](DirectMessagePriceChanged.md) |  | [optional] 
**ForumTopicCreated** | Pointer to [**ForumTopicCreated**](ForumTopicCreated.md) |  | [optional] 
**ForumTopicEdited** | Pointer to [**ForumTopicEdited**](ForumTopicEdited.md) |  | [optional] 
**ForumTopicClosed** | Pointer to **interface{}** |  | [optional] 
**ForumTopicReopened** | Pointer to **interface{}** |  | [optional] 
**GeneralForumTopicHidden** | Pointer to **interface{}** |  | [optional] 
**GeneralForumTopicUnhidden** | Pointer to **interface{}** |  | [optional] 
**GiveawayCreated** | Pointer to [**GiveawayCreated**](GiveawayCreated.md) |  | [optional] 
**Giveaway** | Pointer to [**Giveaway**](Giveaway.md) |  | [optional] 
**GiveawayWinners** | Pointer to [**GiveawayWinners**](GiveawayWinners.md) |  | [optional] 
**GiveawayCompleted** | Pointer to [**GiveawayCompleted**](GiveawayCompleted.md) |  | [optional] 
**PaidMessagePriceChanged** | Pointer to [**PaidMessagePriceChanged**](PaidMessagePriceChanged.md) |  | [optional] 
**VideoChatScheduled** | Pointer to [**VideoChatScheduled**](VideoChatScheduled.md) |  | [optional] 
**VideoChatStarted** | Pointer to **interface{}** |  | [optional] 
**VideoChatEnded** | Pointer to [**VideoChatEnded**](VideoChatEnded.md) |  | [optional] 
**VideoChatParticipantsInvited** | Pointer to [**VideoChatParticipantsInvited**](VideoChatParticipantsInvited.md) |  | [optional] 
**WebAppData** | Pointer to [**WebAppData**](WebAppData.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewMaybeInaccessibleMessage

`func NewMaybeInaccessibleMessage(messageId int32, date int32, chat Chat, ) *MaybeInaccessibleMessage`

NewMaybeInaccessibleMessage instantiates a new MaybeInaccessibleMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaybeInaccessibleMessageWithDefaults

`func NewMaybeInaccessibleMessageWithDefaults() *MaybeInaccessibleMessage`

NewMaybeInaccessibleMessageWithDefaults instantiates a new MaybeInaccessibleMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MaybeInaccessibleMessage) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MaybeInaccessibleMessage) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MaybeInaccessibleMessage) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetMessageThreadId

`func (o *MaybeInaccessibleMessage) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *MaybeInaccessibleMessage) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *MaybeInaccessibleMessage) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *MaybeInaccessibleMessage) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFrom

`func (o *MaybeInaccessibleMessage) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MaybeInaccessibleMessage) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MaybeInaccessibleMessage) SetFrom(v User)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MaybeInaccessibleMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSenderChat

`func (o *MaybeInaccessibleMessage) GetSenderChat() Chat`

GetSenderChat returns the SenderChat field if non-nil, zero value otherwise.

### GetSenderChatOk

`func (o *MaybeInaccessibleMessage) GetSenderChatOk() (*Chat, bool)`

GetSenderChatOk returns a tuple with the SenderChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderChat

`func (o *MaybeInaccessibleMessage) SetSenderChat(v Chat)`

SetSenderChat sets SenderChat field to given value.

### HasSenderChat

`func (o *MaybeInaccessibleMessage) HasSenderChat() bool`

HasSenderChat returns a boolean if a field has been set.

### GetSenderBoostCount

`func (o *MaybeInaccessibleMessage) GetSenderBoostCount() int32`

GetSenderBoostCount returns the SenderBoostCount field if non-nil, zero value otherwise.

### GetSenderBoostCountOk

`func (o *MaybeInaccessibleMessage) GetSenderBoostCountOk() (*int32, bool)`

GetSenderBoostCountOk returns a tuple with the SenderBoostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBoostCount

`func (o *MaybeInaccessibleMessage) SetSenderBoostCount(v int32)`

SetSenderBoostCount sets SenderBoostCount field to given value.

### HasSenderBoostCount

`func (o *MaybeInaccessibleMessage) HasSenderBoostCount() bool`

HasSenderBoostCount returns a boolean if a field has been set.

### GetSenderBusinessBot

`func (o *MaybeInaccessibleMessage) GetSenderBusinessBot() User`

GetSenderBusinessBot returns the SenderBusinessBot field if non-nil, zero value otherwise.

### GetSenderBusinessBotOk

`func (o *MaybeInaccessibleMessage) GetSenderBusinessBotOk() (*User, bool)`

GetSenderBusinessBotOk returns a tuple with the SenderBusinessBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBusinessBot

`func (o *MaybeInaccessibleMessage) SetSenderBusinessBot(v User)`

SetSenderBusinessBot sets SenderBusinessBot field to given value.

### HasSenderBusinessBot

`func (o *MaybeInaccessibleMessage) HasSenderBusinessBot() bool`

HasSenderBusinessBot returns a boolean if a field has been set.

### GetDate

`func (o *MaybeInaccessibleMessage) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MaybeInaccessibleMessage) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MaybeInaccessibleMessage) SetDate(v int32)`

SetDate sets Date field to given value.


### GetBusinessConnectionId

`func (o *MaybeInaccessibleMessage) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *MaybeInaccessibleMessage) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *MaybeInaccessibleMessage) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *MaybeInaccessibleMessage) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChat

`func (o *MaybeInaccessibleMessage) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MaybeInaccessibleMessage) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MaybeInaccessibleMessage) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetForwardOrigin

`func (o *MaybeInaccessibleMessage) GetForwardOrigin() MessageOrigin`

GetForwardOrigin returns the ForwardOrigin field if non-nil, zero value otherwise.

### GetForwardOriginOk

`func (o *MaybeInaccessibleMessage) GetForwardOriginOk() (*MessageOrigin, bool)`

GetForwardOriginOk returns a tuple with the ForwardOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOrigin

`func (o *MaybeInaccessibleMessage) SetForwardOrigin(v MessageOrigin)`

SetForwardOrigin sets ForwardOrigin field to given value.

### HasForwardOrigin

`func (o *MaybeInaccessibleMessage) HasForwardOrigin() bool`

HasForwardOrigin returns a boolean if a field has been set.

### GetIsTopicMessage

`func (o *MaybeInaccessibleMessage) GetIsTopicMessage() bool`

GetIsTopicMessage returns the IsTopicMessage field if non-nil, zero value otherwise.

### GetIsTopicMessageOk

`func (o *MaybeInaccessibleMessage) GetIsTopicMessageOk() (*bool, bool)`

GetIsTopicMessageOk returns a tuple with the IsTopicMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTopicMessage

`func (o *MaybeInaccessibleMessage) SetIsTopicMessage(v bool)`

SetIsTopicMessage sets IsTopicMessage field to given value.

### HasIsTopicMessage

`func (o *MaybeInaccessibleMessage) HasIsTopicMessage() bool`

HasIsTopicMessage returns a boolean if a field has been set.

### GetIsAutomaticForward

`func (o *MaybeInaccessibleMessage) GetIsAutomaticForward() bool`

GetIsAutomaticForward returns the IsAutomaticForward field if non-nil, zero value otherwise.

### GetIsAutomaticForwardOk

`func (o *MaybeInaccessibleMessage) GetIsAutomaticForwardOk() (*bool, bool)`

GetIsAutomaticForwardOk returns a tuple with the IsAutomaticForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticForward

`func (o *MaybeInaccessibleMessage) SetIsAutomaticForward(v bool)`

SetIsAutomaticForward sets IsAutomaticForward field to given value.

### HasIsAutomaticForward

`func (o *MaybeInaccessibleMessage) HasIsAutomaticForward() bool`

HasIsAutomaticForward returns a boolean if a field has been set.

### GetReplyToMessage

`func (o *MaybeInaccessibleMessage) GetReplyToMessage() Message`

GetReplyToMessage returns the ReplyToMessage field if non-nil, zero value otherwise.

### GetReplyToMessageOk

`func (o *MaybeInaccessibleMessage) GetReplyToMessageOk() (*Message, bool)`

GetReplyToMessageOk returns a tuple with the ReplyToMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToMessage

`func (o *MaybeInaccessibleMessage) SetReplyToMessage(v Message)`

SetReplyToMessage sets ReplyToMessage field to given value.

### HasReplyToMessage

`func (o *MaybeInaccessibleMessage) HasReplyToMessage() bool`

HasReplyToMessage returns a boolean if a field has been set.

### GetExternalReply

`func (o *MaybeInaccessibleMessage) GetExternalReply() ExternalReplyInfo`

GetExternalReply returns the ExternalReply field if non-nil, zero value otherwise.

### GetExternalReplyOk

`func (o *MaybeInaccessibleMessage) GetExternalReplyOk() (*ExternalReplyInfo, bool)`

GetExternalReplyOk returns a tuple with the ExternalReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReply

`func (o *MaybeInaccessibleMessage) SetExternalReply(v ExternalReplyInfo)`

SetExternalReply sets ExternalReply field to given value.

### HasExternalReply

`func (o *MaybeInaccessibleMessage) HasExternalReply() bool`

HasExternalReply returns a boolean if a field has been set.

### GetQuote

`func (o *MaybeInaccessibleMessage) GetQuote() TextQuote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *MaybeInaccessibleMessage) GetQuoteOk() (*TextQuote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *MaybeInaccessibleMessage) SetQuote(v TextQuote)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *MaybeInaccessibleMessage) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReplyToStory

`func (o *MaybeInaccessibleMessage) GetReplyToStory() Story`

GetReplyToStory returns the ReplyToStory field if non-nil, zero value otherwise.

### GetReplyToStoryOk

`func (o *MaybeInaccessibleMessage) GetReplyToStoryOk() (*Story, bool)`

GetReplyToStoryOk returns a tuple with the ReplyToStory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToStory

`func (o *MaybeInaccessibleMessage) SetReplyToStory(v Story)`

SetReplyToStory sets ReplyToStory field to given value.

### HasReplyToStory

`func (o *MaybeInaccessibleMessage) HasReplyToStory() bool`

HasReplyToStory returns a boolean if a field has been set.

### GetViaBot

`func (o *MaybeInaccessibleMessage) GetViaBot() User`

GetViaBot returns the ViaBot field if non-nil, zero value otherwise.

### GetViaBotOk

`func (o *MaybeInaccessibleMessage) GetViaBotOk() (*User, bool)`

GetViaBotOk returns a tuple with the ViaBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaBot

`func (o *MaybeInaccessibleMessage) SetViaBot(v User)`

SetViaBot sets ViaBot field to given value.

### HasViaBot

`func (o *MaybeInaccessibleMessage) HasViaBot() bool`

HasViaBot returns a boolean if a field has been set.

### GetEditDate

`func (o *MaybeInaccessibleMessage) GetEditDate() int32`

GetEditDate returns the EditDate field if non-nil, zero value otherwise.

### GetEditDateOk

`func (o *MaybeInaccessibleMessage) GetEditDateOk() (*int32, bool)`

GetEditDateOk returns a tuple with the EditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditDate

`func (o *MaybeInaccessibleMessage) SetEditDate(v int32)`

SetEditDate sets EditDate field to given value.

### HasEditDate

`func (o *MaybeInaccessibleMessage) HasEditDate() bool`

HasEditDate returns a boolean if a field has been set.

### GetHasProtectedContent

`func (o *MaybeInaccessibleMessage) GetHasProtectedContent() bool`

GetHasProtectedContent returns the HasProtectedContent field if non-nil, zero value otherwise.

### GetHasProtectedContentOk

`func (o *MaybeInaccessibleMessage) GetHasProtectedContentOk() (*bool, bool)`

GetHasProtectedContentOk returns a tuple with the HasProtectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProtectedContent

`func (o *MaybeInaccessibleMessage) SetHasProtectedContent(v bool)`

SetHasProtectedContent sets HasProtectedContent field to given value.

### HasHasProtectedContent

`func (o *MaybeInaccessibleMessage) HasHasProtectedContent() bool`

HasHasProtectedContent returns a boolean if a field has been set.

### GetIsFromOffline

`func (o *MaybeInaccessibleMessage) GetIsFromOffline() bool`

GetIsFromOffline returns the IsFromOffline field if non-nil, zero value otherwise.

### GetIsFromOfflineOk

`func (o *MaybeInaccessibleMessage) GetIsFromOfflineOk() (*bool, bool)`

GetIsFromOfflineOk returns a tuple with the IsFromOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromOffline

`func (o *MaybeInaccessibleMessage) SetIsFromOffline(v bool)`

SetIsFromOffline sets IsFromOffline field to given value.

### HasIsFromOffline

`func (o *MaybeInaccessibleMessage) HasIsFromOffline() bool`

HasIsFromOffline returns a boolean if a field has been set.

### GetMediaGroupId

`func (o *MaybeInaccessibleMessage) GetMediaGroupId() string`

GetMediaGroupId returns the MediaGroupId field if non-nil, zero value otherwise.

### GetMediaGroupIdOk

`func (o *MaybeInaccessibleMessage) GetMediaGroupIdOk() (*string, bool)`

GetMediaGroupIdOk returns a tuple with the MediaGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaGroupId

`func (o *MaybeInaccessibleMessage) SetMediaGroupId(v string)`

SetMediaGroupId sets MediaGroupId field to given value.

### HasMediaGroupId

`func (o *MaybeInaccessibleMessage) HasMediaGroupId() bool`

HasMediaGroupId returns a boolean if a field has been set.

### GetAuthorSignature

`func (o *MaybeInaccessibleMessage) GetAuthorSignature() string`

GetAuthorSignature returns the AuthorSignature field if non-nil, zero value otherwise.

### GetAuthorSignatureOk

`func (o *MaybeInaccessibleMessage) GetAuthorSignatureOk() (*string, bool)`

GetAuthorSignatureOk returns a tuple with the AuthorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorSignature

`func (o *MaybeInaccessibleMessage) SetAuthorSignature(v string)`

SetAuthorSignature sets AuthorSignature field to given value.

### HasAuthorSignature

`func (o *MaybeInaccessibleMessage) HasAuthorSignature() bool`

HasAuthorSignature returns a boolean if a field has been set.

### GetPaidStarCount

`func (o *MaybeInaccessibleMessage) GetPaidStarCount() int32`

GetPaidStarCount returns the PaidStarCount field if non-nil, zero value otherwise.

### GetPaidStarCountOk

`func (o *MaybeInaccessibleMessage) GetPaidStarCountOk() (*int32, bool)`

GetPaidStarCountOk returns a tuple with the PaidStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidStarCount

`func (o *MaybeInaccessibleMessage) SetPaidStarCount(v int32)`

SetPaidStarCount sets PaidStarCount field to given value.

### HasPaidStarCount

`func (o *MaybeInaccessibleMessage) HasPaidStarCount() bool`

HasPaidStarCount returns a boolean if a field has been set.

### GetText

`func (o *MaybeInaccessibleMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MaybeInaccessibleMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MaybeInaccessibleMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MaybeInaccessibleMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEntities

`func (o *MaybeInaccessibleMessage) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *MaybeInaccessibleMessage) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *MaybeInaccessibleMessage) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *MaybeInaccessibleMessage) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetLinkPreviewOptions

`func (o *MaybeInaccessibleMessage) GetLinkPreviewOptions() LinkPreviewOptions`

GetLinkPreviewOptions returns the LinkPreviewOptions field if non-nil, zero value otherwise.

### GetLinkPreviewOptionsOk

`func (o *MaybeInaccessibleMessage) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool)`

GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPreviewOptions

`func (o *MaybeInaccessibleMessage) SetLinkPreviewOptions(v LinkPreviewOptions)`

SetLinkPreviewOptions sets LinkPreviewOptions field to given value.

### HasLinkPreviewOptions

`func (o *MaybeInaccessibleMessage) HasLinkPreviewOptions() bool`

HasLinkPreviewOptions returns a boolean if a field has been set.

### GetEffectId

`func (o *MaybeInaccessibleMessage) GetEffectId() string`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *MaybeInaccessibleMessage) GetEffectIdOk() (*string, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *MaybeInaccessibleMessage) SetEffectId(v string)`

SetEffectId sets EffectId field to given value.

### HasEffectId

`func (o *MaybeInaccessibleMessage) HasEffectId() bool`

HasEffectId returns a boolean if a field has been set.

### GetAnimation

`func (o *MaybeInaccessibleMessage) GetAnimation() Animation`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *MaybeInaccessibleMessage) GetAnimationOk() (*Animation, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *MaybeInaccessibleMessage) SetAnimation(v Animation)`

SetAnimation sets Animation field to given value.

### HasAnimation

`func (o *MaybeInaccessibleMessage) HasAnimation() bool`

HasAnimation returns a boolean if a field has been set.

### GetAudio

`func (o *MaybeInaccessibleMessage) GetAudio() Audio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *MaybeInaccessibleMessage) GetAudioOk() (*Audio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *MaybeInaccessibleMessage) SetAudio(v Audio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *MaybeInaccessibleMessage) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *MaybeInaccessibleMessage) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *MaybeInaccessibleMessage) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *MaybeInaccessibleMessage) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *MaybeInaccessibleMessage) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetPaidMedia

`func (o *MaybeInaccessibleMessage) GetPaidMedia() PaidMediaInfo`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *MaybeInaccessibleMessage) GetPaidMediaOk() (*PaidMediaInfo, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *MaybeInaccessibleMessage) SetPaidMedia(v PaidMediaInfo)`

SetPaidMedia sets PaidMedia field to given value.

### HasPaidMedia

`func (o *MaybeInaccessibleMessage) HasPaidMedia() bool`

HasPaidMedia returns a boolean if a field has been set.

### GetPhoto

`func (o *MaybeInaccessibleMessage) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MaybeInaccessibleMessage) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MaybeInaccessibleMessage) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MaybeInaccessibleMessage) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetSticker

`func (o *MaybeInaccessibleMessage) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *MaybeInaccessibleMessage) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *MaybeInaccessibleMessage) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *MaybeInaccessibleMessage) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetStory

`func (o *MaybeInaccessibleMessage) GetStory() Story`

GetStory returns the Story field if non-nil, zero value otherwise.

### GetStoryOk

`func (o *MaybeInaccessibleMessage) GetStoryOk() (*Story, bool)`

GetStoryOk returns a tuple with the Story field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStory

`func (o *MaybeInaccessibleMessage) SetStory(v Story)`

SetStory sets Story field to given value.

### HasStory

`func (o *MaybeInaccessibleMessage) HasStory() bool`

HasStory returns a boolean if a field has been set.

### GetVideo

`func (o *MaybeInaccessibleMessage) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MaybeInaccessibleMessage) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *MaybeInaccessibleMessage) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *MaybeInaccessibleMessage) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoNote

`func (o *MaybeInaccessibleMessage) GetVideoNote() VideoNote`

GetVideoNote returns the VideoNote field if non-nil, zero value otherwise.

### GetVideoNoteOk

`func (o *MaybeInaccessibleMessage) GetVideoNoteOk() (*VideoNote, bool)`

GetVideoNoteOk returns a tuple with the VideoNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoNote

`func (o *MaybeInaccessibleMessage) SetVideoNote(v VideoNote)`

SetVideoNote sets VideoNote field to given value.

### HasVideoNote

`func (o *MaybeInaccessibleMessage) HasVideoNote() bool`

HasVideoNote returns a boolean if a field has been set.

### GetVoice

`func (o *MaybeInaccessibleMessage) GetVoice() Voice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *MaybeInaccessibleMessage) GetVoiceOk() (*Voice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *MaybeInaccessibleMessage) SetVoice(v Voice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *MaybeInaccessibleMessage) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetCaption

`func (o *MaybeInaccessibleMessage) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *MaybeInaccessibleMessage) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *MaybeInaccessibleMessage) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *MaybeInaccessibleMessage) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *MaybeInaccessibleMessage) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *MaybeInaccessibleMessage) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *MaybeInaccessibleMessage) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *MaybeInaccessibleMessage) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *MaybeInaccessibleMessage) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *MaybeInaccessibleMessage) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *MaybeInaccessibleMessage) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *MaybeInaccessibleMessage) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetHasMediaSpoiler

`func (o *MaybeInaccessibleMessage) GetHasMediaSpoiler() bool`

GetHasMediaSpoiler returns the HasMediaSpoiler field if non-nil, zero value otherwise.

### GetHasMediaSpoilerOk

`func (o *MaybeInaccessibleMessage) GetHasMediaSpoilerOk() (*bool, bool)`

GetHasMediaSpoilerOk returns a tuple with the HasMediaSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMediaSpoiler

`func (o *MaybeInaccessibleMessage) SetHasMediaSpoiler(v bool)`

SetHasMediaSpoiler sets HasMediaSpoiler field to given value.

### HasHasMediaSpoiler

`func (o *MaybeInaccessibleMessage) HasHasMediaSpoiler() bool`

HasHasMediaSpoiler returns a boolean if a field has been set.

### GetChecklist

`func (o *MaybeInaccessibleMessage) GetChecklist() Checklist`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *MaybeInaccessibleMessage) GetChecklistOk() (*Checklist, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *MaybeInaccessibleMessage) SetChecklist(v Checklist)`

SetChecklist sets Checklist field to given value.

### HasChecklist

`func (o *MaybeInaccessibleMessage) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### GetContact

`func (o *MaybeInaccessibleMessage) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *MaybeInaccessibleMessage) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *MaybeInaccessibleMessage) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *MaybeInaccessibleMessage) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDice

`func (o *MaybeInaccessibleMessage) GetDice() Dice`

GetDice returns the Dice field if non-nil, zero value otherwise.

### GetDiceOk

`func (o *MaybeInaccessibleMessage) GetDiceOk() (*Dice, bool)`

GetDiceOk returns a tuple with the Dice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDice

`func (o *MaybeInaccessibleMessage) SetDice(v Dice)`

SetDice sets Dice field to given value.

### HasDice

`func (o *MaybeInaccessibleMessage) HasDice() bool`

HasDice returns a boolean if a field has been set.

### GetGame

`func (o *MaybeInaccessibleMessage) GetGame() Game`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *MaybeInaccessibleMessage) GetGameOk() (*Game, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *MaybeInaccessibleMessage) SetGame(v Game)`

SetGame sets Game field to given value.

### HasGame

`func (o *MaybeInaccessibleMessage) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetPoll

`func (o *MaybeInaccessibleMessage) GetPoll() Poll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *MaybeInaccessibleMessage) GetPollOk() (*Poll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *MaybeInaccessibleMessage) SetPoll(v Poll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *MaybeInaccessibleMessage) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetVenue

`func (o *MaybeInaccessibleMessage) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *MaybeInaccessibleMessage) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *MaybeInaccessibleMessage) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *MaybeInaccessibleMessage) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetLocation

`func (o *MaybeInaccessibleMessage) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MaybeInaccessibleMessage) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MaybeInaccessibleMessage) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MaybeInaccessibleMessage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNewChatMembers

`func (o *MaybeInaccessibleMessage) GetNewChatMembers() []User`

GetNewChatMembers returns the NewChatMembers field if non-nil, zero value otherwise.

### GetNewChatMembersOk

`func (o *MaybeInaccessibleMessage) GetNewChatMembersOk() (*[]User, bool)`

GetNewChatMembersOk returns a tuple with the NewChatMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatMembers

`func (o *MaybeInaccessibleMessage) SetNewChatMembers(v []User)`

SetNewChatMembers sets NewChatMembers field to given value.

### HasNewChatMembers

`func (o *MaybeInaccessibleMessage) HasNewChatMembers() bool`

HasNewChatMembers returns a boolean if a field has been set.

### GetLeftChatMember

`func (o *MaybeInaccessibleMessage) GetLeftChatMember() User`

GetLeftChatMember returns the LeftChatMember field if non-nil, zero value otherwise.

### GetLeftChatMemberOk

`func (o *MaybeInaccessibleMessage) GetLeftChatMemberOk() (*User, bool)`

GetLeftChatMemberOk returns a tuple with the LeftChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftChatMember

`func (o *MaybeInaccessibleMessage) SetLeftChatMember(v User)`

SetLeftChatMember sets LeftChatMember field to given value.

### HasLeftChatMember

`func (o *MaybeInaccessibleMessage) HasLeftChatMember() bool`

HasLeftChatMember returns a boolean if a field has been set.

### GetNewChatTitle

`func (o *MaybeInaccessibleMessage) GetNewChatTitle() string`

GetNewChatTitle returns the NewChatTitle field if non-nil, zero value otherwise.

### GetNewChatTitleOk

`func (o *MaybeInaccessibleMessage) GetNewChatTitleOk() (*string, bool)`

GetNewChatTitleOk returns a tuple with the NewChatTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatTitle

`func (o *MaybeInaccessibleMessage) SetNewChatTitle(v string)`

SetNewChatTitle sets NewChatTitle field to given value.

### HasNewChatTitle

`func (o *MaybeInaccessibleMessage) HasNewChatTitle() bool`

HasNewChatTitle returns a boolean if a field has been set.

### GetNewChatPhoto

`func (o *MaybeInaccessibleMessage) GetNewChatPhoto() []PhotoSize`

GetNewChatPhoto returns the NewChatPhoto field if non-nil, zero value otherwise.

### GetNewChatPhotoOk

`func (o *MaybeInaccessibleMessage) GetNewChatPhotoOk() (*[]PhotoSize, bool)`

GetNewChatPhotoOk returns a tuple with the NewChatPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatPhoto

`func (o *MaybeInaccessibleMessage) SetNewChatPhoto(v []PhotoSize)`

SetNewChatPhoto sets NewChatPhoto field to given value.

### HasNewChatPhoto

`func (o *MaybeInaccessibleMessage) HasNewChatPhoto() bool`

HasNewChatPhoto returns a boolean if a field has been set.

### GetDeleteChatPhoto

`func (o *MaybeInaccessibleMessage) GetDeleteChatPhoto() bool`

GetDeleteChatPhoto returns the DeleteChatPhoto field if non-nil, zero value otherwise.

### GetDeleteChatPhotoOk

`func (o *MaybeInaccessibleMessage) GetDeleteChatPhotoOk() (*bool, bool)`

GetDeleteChatPhotoOk returns a tuple with the DeleteChatPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteChatPhoto

`func (o *MaybeInaccessibleMessage) SetDeleteChatPhoto(v bool)`

SetDeleteChatPhoto sets DeleteChatPhoto field to given value.

### HasDeleteChatPhoto

`func (o *MaybeInaccessibleMessage) HasDeleteChatPhoto() bool`

HasDeleteChatPhoto returns a boolean if a field has been set.

### GetGroupChatCreated

`func (o *MaybeInaccessibleMessage) GetGroupChatCreated() bool`

GetGroupChatCreated returns the GroupChatCreated field if non-nil, zero value otherwise.

### GetGroupChatCreatedOk

`func (o *MaybeInaccessibleMessage) GetGroupChatCreatedOk() (*bool, bool)`

GetGroupChatCreatedOk returns a tuple with the GroupChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupChatCreated

`func (o *MaybeInaccessibleMessage) SetGroupChatCreated(v bool)`

SetGroupChatCreated sets GroupChatCreated field to given value.

### HasGroupChatCreated

`func (o *MaybeInaccessibleMessage) HasGroupChatCreated() bool`

HasGroupChatCreated returns a boolean if a field has been set.

### GetSupergroupChatCreated

`func (o *MaybeInaccessibleMessage) GetSupergroupChatCreated() bool`

GetSupergroupChatCreated returns the SupergroupChatCreated field if non-nil, zero value otherwise.

### GetSupergroupChatCreatedOk

`func (o *MaybeInaccessibleMessage) GetSupergroupChatCreatedOk() (*bool, bool)`

GetSupergroupChatCreatedOk returns a tuple with the SupergroupChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupergroupChatCreated

`func (o *MaybeInaccessibleMessage) SetSupergroupChatCreated(v bool)`

SetSupergroupChatCreated sets SupergroupChatCreated field to given value.

### HasSupergroupChatCreated

`func (o *MaybeInaccessibleMessage) HasSupergroupChatCreated() bool`

HasSupergroupChatCreated returns a boolean if a field has been set.

### GetChannelChatCreated

`func (o *MaybeInaccessibleMessage) GetChannelChatCreated() bool`

GetChannelChatCreated returns the ChannelChatCreated field if non-nil, zero value otherwise.

### GetChannelChatCreatedOk

`func (o *MaybeInaccessibleMessage) GetChannelChatCreatedOk() (*bool, bool)`

GetChannelChatCreatedOk returns a tuple with the ChannelChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelChatCreated

`func (o *MaybeInaccessibleMessage) SetChannelChatCreated(v bool)`

SetChannelChatCreated sets ChannelChatCreated field to given value.

### HasChannelChatCreated

`func (o *MaybeInaccessibleMessage) HasChannelChatCreated() bool`

HasChannelChatCreated returns a boolean if a field has been set.

### GetMessageAutoDeleteTimerChanged

`func (o *MaybeInaccessibleMessage) GetMessageAutoDeleteTimerChanged() MessageAutoDeleteTimerChanged`

GetMessageAutoDeleteTimerChanged returns the MessageAutoDeleteTimerChanged field if non-nil, zero value otherwise.

### GetMessageAutoDeleteTimerChangedOk

`func (o *MaybeInaccessibleMessage) GetMessageAutoDeleteTimerChangedOk() (*MessageAutoDeleteTimerChanged, bool)`

GetMessageAutoDeleteTimerChangedOk returns a tuple with the MessageAutoDeleteTimerChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageAutoDeleteTimerChanged

`func (o *MaybeInaccessibleMessage) SetMessageAutoDeleteTimerChanged(v MessageAutoDeleteTimerChanged)`

SetMessageAutoDeleteTimerChanged sets MessageAutoDeleteTimerChanged field to given value.

### HasMessageAutoDeleteTimerChanged

`func (o *MaybeInaccessibleMessage) HasMessageAutoDeleteTimerChanged() bool`

HasMessageAutoDeleteTimerChanged returns a boolean if a field has been set.

### GetMigrateToChatId

`func (o *MaybeInaccessibleMessage) GetMigrateToChatId() int32`

GetMigrateToChatId returns the MigrateToChatId field if non-nil, zero value otherwise.

### GetMigrateToChatIdOk

`func (o *MaybeInaccessibleMessage) GetMigrateToChatIdOk() (*int32, bool)`

GetMigrateToChatIdOk returns a tuple with the MigrateToChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateToChatId

`func (o *MaybeInaccessibleMessage) SetMigrateToChatId(v int32)`

SetMigrateToChatId sets MigrateToChatId field to given value.

### HasMigrateToChatId

`func (o *MaybeInaccessibleMessage) HasMigrateToChatId() bool`

HasMigrateToChatId returns a boolean if a field has been set.

### GetMigrateFromChatId

`func (o *MaybeInaccessibleMessage) GetMigrateFromChatId() int32`

GetMigrateFromChatId returns the MigrateFromChatId field if non-nil, zero value otherwise.

### GetMigrateFromChatIdOk

`func (o *MaybeInaccessibleMessage) GetMigrateFromChatIdOk() (*int32, bool)`

GetMigrateFromChatIdOk returns a tuple with the MigrateFromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateFromChatId

`func (o *MaybeInaccessibleMessage) SetMigrateFromChatId(v int32)`

SetMigrateFromChatId sets MigrateFromChatId field to given value.

### HasMigrateFromChatId

`func (o *MaybeInaccessibleMessage) HasMigrateFromChatId() bool`

HasMigrateFromChatId returns a boolean if a field has been set.

### GetPinnedMessage

`func (o *MaybeInaccessibleMessage) GetPinnedMessage() MaybeInaccessibleMessage`

GetPinnedMessage returns the PinnedMessage field if non-nil, zero value otherwise.

### GetPinnedMessageOk

`func (o *MaybeInaccessibleMessage) GetPinnedMessageOk() (*MaybeInaccessibleMessage, bool)`

GetPinnedMessageOk returns a tuple with the PinnedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedMessage

`func (o *MaybeInaccessibleMessage) SetPinnedMessage(v MaybeInaccessibleMessage)`

SetPinnedMessage sets PinnedMessage field to given value.

### HasPinnedMessage

`func (o *MaybeInaccessibleMessage) HasPinnedMessage() bool`

HasPinnedMessage returns a boolean if a field has been set.

### GetInvoice

`func (o *MaybeInaccessibleMessage) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MaybeInaccessibleMessage) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MaybeInaccessibleMessage) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MaybeInaccessibleMessage) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetSuccessfulPayment

`func (o *MaybeInaccessibleMessage) GetSuccessfulPayment() SuccessfulPayment`

GetSuccessfulPayment returns the SuccessfulPayment field if non-nil, zero value otherwise.

### GetSuccessfulPaymentOk

`func (o *MaybeInaccessibleMessage) GetSuccessfulPaymentOk() (*SuccessfulPayment, bool)`

GetSuccessfulPaymentOk returns a tuple with the SuccessfulPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulPayment

`func (o *MaybeInaccessibleMessage) SetSuccessfulPayment(v SuccessfulPayment)`

SetSuccessfulPayment sets SuccessfulPayment field to given value.

### HasSuccessfulPayment

`func (o *MaybeInaccessibleMessage) HasSuccessfulPayment() bool`

HasSuccessfulPayment returns a boolean if a field has been set.

### GetRefundedPayment

`func (o *MaybeInaccessibleMessage) GetRefundedPayment() RefundedPayment`

GetRefundedPayment returns the RefundedPayment field if non-nil, zero value otherwise.

### GetRefundedPaymentOk

`func (o *MaybeInaccessibleMessage) GetRefundedPaymentOk() (*RefundedPayment, bool)`

GetRefundedPaymentOk returns a tuple with the RefundedPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedPayment

`func (o *MaybeInaccessibleMessage) SetRefundedPayment(v RefundedPayment)`

SetRefundedPayment sets RefundedPayment field to given value.

### HasRefundedPayment

`func (o *MaybeInaccessibleMessage) HasRefundedPayment() bool`

HasRefundedPayment returns a boolean if a field has been set.

### GetUsersShared

`func (o *MaybeInaccessibleMessage) GetUsersShared() UsersShared`

GetUsersShared returns the UsersShared field if non-nil, zero value otherwise.

### GetUsersSharedOk

`func (o *MaybeInaccessibleMessage) GetUsersSharedOk() (*UsersShared, bool)`

GetUsersSharedOk returns a tuple with the UsersShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersShared

`func (o *MaybeInaccessibleMessage) SetUsersShared(v UsersShared)`

SetUsersShared sets UsersShared field to given value.

### HasUsersShared

`func (o *MaybeInaccessibleMessage) HasUsersShared() bool`

HasUsersShared returns a boolean if a field has been set.

### GetChatShared

`func (o *MaybeInaccessibleMessage) GetChatShared() ChatShared`

GetChatShared returns the ChatShared field if non-nil, zero value otherwise.

### GetChatSharedOk

`func (o *MaybeInaccessibleMessage) GetChatSharedOk() (*ChatShared, bool)`

GetChatSharedOk returns a tuple with the ChatShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatShared

`func (o *MaybeInaccessibleMessage) SetChatShared(v ChatShared)`

SetChatShared sets ChatShared field to given value.

### HasChatShared

`func (o *MaybeInaccessibleMessage) HasChatShared() bool`

HasChatShared returns a boolean if a field has been set.

### GetGift

`func (o *MaybeInaccessibleMessage) GetGift() GiftInfo`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *MaybeInaccessibleMessage) GetGiftOk() (*GiftInfo, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *MaybeInaccessibleMessage) SetGift(v GiftInfo)`

SetGift sets Gift field to given value.

### HasGift

`func (o *MaybeInaccessibleMessage) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetUniqueGift

`func (o *MaybeInaccessibleMessage) GetUniqueGift() UniqueGiftInfo`

GetUniqueGift returns the UniqueGift field if non-nil, zero value otherwise.

### GetUniqueGiftOk

`func (o *MaybeInaccessibleMessage) GetUniqueGiftOk() (*UniqueGiftInfo, bool)`

GetUniqueGiftOk returns a tuple with the UniqueGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGift

`func (o *MaybeInaccessibleMessage) SetUniqueGift(v UniqueGiftInfo)`

SetUniqueGift sets UniqueGift field to given value.

### HasUniqueGift

`func (o *MaybeInaccessibleMessage) HasUniqueGift() bool`

HasUniqueGift returns a boolean if a field has been set.

### GetConnectedWebsite

`func (o *MaybeInaccessibleMessage) GetConnectedWebsite() string`

GetConnectedWebsite returns the ConnectedWebsite field if non-nil, zero value otherwise.

### GetConnectedWebsiteOk

`func (o *MaybeInaccessibleMessage) GetConnectedWebsiteOk() (*string, bool)`

GetConnectedWebsiteOk returns a tuple with the ConnectedWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWebsite

`func (o *MaybeInaccessibleMessage) SetConnectedWebsite(v string)`

SetConnectedWebsite sets ConnectedWebsite field to given value.

### HasConnectedWebsite

`func (o *MaybeInaccessibleMessage) HasConnectedWebsite() bool`

HasConnectedWebsite returns a boolean if a field has been set.

### GetWriteAccessAllowed

`func (o *MaybeInaccessibleMessage) GetWriteAccessAllowed() WriteAccessAllowed`

GetWriteAccessAllowed returns the WriteAccessAllowed field if non-nil, zero value otherwise.

### GetWriteAccessAllowedOk

`func (o *MaybeInaccessibleMessage) GetWriteAccessAllowedOk() (*WriteAccessAllowed, bool)`

GetWriteAccessAllowedOk returns a tuple with the WriteAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAccessAllowed

`func (o *MaybeInaccessibleMessage) SetWriteAccessAllowed(v WriteAccessAllowed)`

SetWriteAccessAllowed sets WriteAccessAllowed field to given value.

### HasWriteAccessAllowed

`func (o *MaybeInaccessibleMessage) HasWriteAccessAllowed() bool`

HasWriteAccessAllowed returns a boolean if a field has been set.

### GetPassportData

`func (o *MaybeInaccessibleMessage) GetPassportData() PassportData`

GetPassportData returns the PassportData field if non-nil, zero value otherwise.

### GetPassportDataOk

`func (o *MaybeInaccessibleMessage) GetPassportDataOk() (*PassportData, bool)`

GetPassportDataOk returns a tuple with the PassportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportData

`func (o *MaybeInaccessibleMessage) SetPassportData(v PassportData)`

SetPassportData sets PassportData field to given value.

### HasPassportData

`func (o *MaybeInaccessibleMessage) HasPassportData() bool`

HasPassportData returns a boolean if a field has been set.

### GetProximityAlertTriggered

`func (o *MaybeInaccessibleMessage) GetProximityAlertTriggered() ProximityAlertTriggered`

GetProximityAlertTriggered returns the ProximityAlertTriggered field if non-nil, zero value otherwise.

### GetProximityAlertTriggeredOk

`func (o *MaybeInaccessibleMessage) GetProximityAlertTriggeredOk() (*ProximityAlertTriggered, bool)`

GetProximityAlertTriggeredOk returns a tuple with the ProximityAlertTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertTriggered

`func (o *MaybeInaccessibleMessage) SetProximityAlertTriggered(v ProximityAlertTriggered)`

SetProximityAlertTriggered sets ProximityAlertTriggered field to given value.

### HasProximityAlertTriggered

`func (o *MaybeInaccessibleMessage) HasProximityAlertTriggered() bool`

HasProximityAlertTriggered returns a boolean if a field has been set.

### GetBoostAdded

`func (o *MaybeInaccessibleMessage) GetBoostAdded() ChatBoostAdded`

GetBoostAdded returns the BoostAdded field if non-nil, zero value otherwise.

### GetBoostAddedOk

`func (o *MaybeInaccessibleMessage) GetBoostAddedOk() (*ChatBoostAdded, bool)`

GetBoostAddedOk returns a tuple with the BoostAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostAdded

`func (o *MaybeInaccessibleMessage) SetBoostAdded(v ChatBoostAdded)`

SetBoostAdded sets BoostAdded field to given value.

### HasBoostAdded

`func (o *MaybeInaccessibleMessage) HasBoostAdded() bool`

HasBoostAdded returns a boolean if a field has been set.

### GetChatBackgroundSet

`func (o *MaybeInaccessibleMessage) GetChatBackgroundSet() ChatBackground`

GetChatBackgroundSet returns the ChatBackgroundSet field if non-nil, zero value otherwise.

### GetChatBackgroundSetOk

`func (o *MaybeInaccessibleMessage) GetChatBackgroundSetOk() (*ChatBackground, bool)`

GetChatBackgroundSetOk returns a tuple with the ChatBackgroundSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatBackgroundSet

`func (o *MaybeInaccessibleMessage) SetChatBackgroundSet(v ChatBackground)`

SetChatBackgroundSet sets ChatBackgroundSet field to given value.

### HasChatBackgroundSet

`func (o *MaybeInaccessibleMessage) HasChatBackgroundSet() bool`

HasChatBackgroundSet returns a boolean if a field has been set.

### GetChecklistTasksDone

`func (o *MaybeInaccessibleMessage) GetChecklistTasksDone() ChecklistTasksDone`

GetChecklistTasksDone returns the ChecklistTasksDone field if non-nil, zero value otherwise.

### GetChecklistTasksDoneOk

`func (o *MaybeInaccessibleMessage) GetChecklistTasksDoneOk() (*ChecklistTasksDone, bool)`

GetChecklistTasksDoneOk returns a tuple with the ChecklistTasksDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistTasksDone

`func (o *MaybeInaccessibleMessage) SetChecklistTasksDone(v ChecklistTasksDone)`

SetChecklistTasksDone sets ChecklistTasksDone field to given value.

### HasChecklistTasksDone

`func (o *MaybeInaccessibleMessage) HasChecklistTasksDone() bool`

HasChecklistTasksDone returns a boolean if a field has been set.

### GetChecklistTasksAdded

`func (o *MaybeInaccessibleMessage) GetChecklistTasksAdded() ChecklistTasksAdded`

GetChecklistTasksAdded returns the ChecklistTasksAdded field if non-nil, zero value otherwise.

### GetChecklistTasksAddedOk

`func (o *MaybeInaccessibleMessage) GetChecklistTasksAddedOk() (*ChecklistTasksAdded, bool)`

GetChecklistTasksAddedOk returns a tuple with the ChecklistTasksAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistTasksAdded

`func (o *MaybeInaccessibleMessage) SetChecklistTasksAdded(v ChecklistTasksAdded)`

SetChecklistTasksAdded sets ChecklistTasksAdded field to given value.

### HasChecklistTasksAdded

`func (o *MaybeInaccessibleMessage) HasChecklistTasksAdded() bool`

HasChecklistTasksAdded returns a boolean if a field has been set.

### GetDirectMessagePriceChanged

`func (o *MaybeInaccessibleMessage) GetDirectMessagePriceChanged() DirectMessagePriceChanged`

GetDirectMessagePriceChanged returns the DirectMessagePriceChanged field if non-nil, zero value otherwise.

### GetDirectMessagePriceChangedOk

`func (o *MaybeInaccessibleMessage) GetDirectMessagePriceChangedOk() (*DirectMessagePriceChanged, bool)`

GetDirectMessagePriceChangedOk returns a tuple with the DirectMessagePriceChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMessagePriceChanged

`func (o *MaybeInaccessibleMessage) SetDirectMessagePriceChanged(v DirectMessagePriceChanged)`

SetDirectMessagePriceChanged sets DirectMessagePriceChanged field to given value.

### HasDirectMessagePriceChanged

`func (o *MaybeInaccessibleMessage) HasDirectMessagePriceChanged() bool`

HasDirectMessagePriceChanged returns a boolean if a field has been set.

### GetForumTopicCreated

`func (o *MaybeInaccessibleMessage) GetForumTopicCreated() ForumTopicCreated`

GetForumTopicCreated returns the ForumTopicCreated field if non-nil, zero value otherwise.

### GetForumTopicCreatedOk

`func (o *MaybeInaccessibleMessage) GetForumTopicCreatedOk() (*ForumTopicCreated, bool)`

GetForumTopicCreatedOk returns a tuple with the ForumTopicCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicCreated

`func (o *MaybeInaccessibleMessage) SetForumTopicCreated(v ForumTopicCreated)`

SetForumTopicCreated sets ForumTopicCreated field to given value.

### HasForumTopicCreated

`func (o *MaybeInaccessibleMessage) HasForumTopicCreated() bool`

HasForumTopicCreated returns a boolean if a field has been set.

### GetForumTopicEdited

`func (o *MaybeInaccessibleMessage) GetForumTopicEdited() ForumTopicEdited`

GetForumTopicEdited returns the ForumTopicEdited field if non-nil, zero value otherwise.

### GetForumTopicEditedOk

`func (o *MaybeInaccessibleMessage) GetForumTopicEditedOk() (*ForumTopicEdited, bool)`

GetForumTopicEditedOk returns a tuple with the ForumTopicEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicEdited

`func (o *MaybeInaccessibleMessage) SetForumTopicEdited(v ForumTopicEdited)`

SetForumTopicEdited sets ForumTopicEdited field to given value.

### HasForumTopicEdited

`func (o *MaybeInaccessibleMessage) HasForumTopicEdited() bool`

HasForumTopicEdited returns a boolean if a field has been set.

### GetForumTopicClosed

`func (o *MaybeInaccessibleMessage) GetForumTopicClosed() interface{}`

GetForumTopicClosed returns the ForumTopicClosed field if non-nil, zero value otherwise.

### GetForumTopicClosedOk

`func (o *MaybeInaccessibleMessage) GetForumTopicClosedOk() (*interface{}, bool)`

GetForumTopicClosedOk returns a tuple with the ForumTopicClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicClosed

`func (o *MaybeInaccessibleMessage) SetForumTopicClosed(v interface{})`

SetForumTopicClosed sets ForumTopicClosed field to given value.

### HasForumTopicClosed

`func (o *MaybeInaccessibleMessage) HasForumTopicClosed() bool`

HasForumTopicClosed returns a boolean if a field has been set.

### SetForumTopicClosedNil

`func (o *MaybeInaccessibleMessage) SetForumTopicClosedNil(b bool)`

 SetForumTopicClosedNil sets the value for ForumTopicClosed to be an explicit nil

### UnsetForumTopicClosed
`func (o *MaybeInaccessibleMessage) UnsetForumTopicClosed()`

UnsetForumTopicClosed ensures that no value is present for ForumTopicClosed, not even an explicit nil
### GetForumTopicReopened

`func (o *MaybeInaccessibleMessage) GetForumTopicReopened() interface{}`

GetForumTopicReopened returns the ForumTopicReopened field if non-nil, zero value otherwise.

### GetForumTopicReopenedOk

`func (o *MaybeInaccessibleMessage) GetForumTopicReopenedOk() (*interface{}, bool)`

GetForumTopicReopenedOk returns a tuple with the ForumTopicReopened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicReopened

`func (o *MaybeInaccessibleMessage) SetForumTopicReopened(v interface{})`

SetForumTopicReopened sets ForumTopicReopened field to given value.

### HasForumTopicReopened

`func (o *MaybeInaccessibleMessage) HasForumTopicReopened() bool`

HasForumTopicReopened returns a boolean if a field has been set.

### SetForumTopicReopenedNil

`func (o *MaybeInaccessibleMessage) SetForumTopicReopenedNil(b bool)`

 SetForumTopicReopenedNil sets the value for ForumTopicReopened to be an explicit nil

### UnsetForumTopicReopened
`func (o *MaybeInaccessibleMessage) UnsetForumTopicReopened()`

UnsetForumTopicReopened ensures that no value is present for ForumTopicReopened, not even an explicit nil
### GetGeneralForumTopicHidden

`func (o *MaybeInaccessibleMessage) GetGeneralForumTopicHidden() interface{}`

GetGeneralForumTopicHidden returns the GeneralForumTopicHidden field if non-nil, zero value otherwise.

### GetGeneralForumTopicHiddenOk

`func (o *MaybeInaccessibleMessage) GetGeneralForumTopicHiddenOk() (*interface{}, bool)`

GetGeneralForumTopicHiddenOk returns a tuple with the GeneralForumTopicHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralForumTopicHidden

`func (o *MaybeInaccessibleMessage) SetGeneralForumTopicHidden(v interface{})`

SetGeneralForumTopicHidden sets GeneralForumTopicHidden field to given value.

### HasGeneralForumTopicHidden

`func (o *MaybeInaccessibleMessage) HasGeneralForumTopicHidden() bool`

HasGeneralForumTopicHidden returns a boolean if a field has been set.

### SetGeneralForumTopicHiddenNil

`func (o *MaybeInaccessibleMessage) SetGeneralForumTopicHiddenNil(b bool)`

 SetGeneralForumTopicHiddenNil sets the value for GeneralForumTopicHidden to be an explicit nil

### UnsetGeneralForumTopicHidden
`func (o *MaybeInaccessibleMessage) UnsetGeneralForumTopicHidden()`

UnsetGeneralForumTopicHidden ensures that no value is present for GeneralForumTopicHidden, not even an explicit nil
### GetGeneralForumTopicUnhidden

`func (o *MaybeInaccessibleMessage) GetGeneralForumTopicUnhidden() interface{}`

GetGeneralForumTopicUnhidden returns the GeneralForumTopicUnhidden field if non-nil, zero value otherwise.

### GetGeneralForumTopicUnhiddenOk

`func (o *MaybeInaccessibleMessage) GetGeneralForumTopicUnhiddenOk() (*interface{}, bool)`

GetGeneralForumTopicUnhiddenOk returns a tuple with the GeneralForumTopicUnhidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralForumTopicUnhidden

`func (o *MaybeInaccessibleMessage) SetGeneralForumTopicUnhidden(v interface{})`

SetGeneralForumTopicUnhidden sets GeneralForumTopicUnhidden field to given value.

### HasGeneralForumTopicUnhidden

`func (o *MaybeInaccessibleMessage) HasGeneralForumTopicUnhidden() bool`

HasGeneralForumTopicUnhidden returns a boolean if a field has been set.

### SetGeneralForumTopicUnhiddenNil

`func (o *MaybeInaccessibleMessage) SetGeneralForumTopicUnhiddenNil(b bool)`

 SetGeneralForumTopicUnhiddenNil sets the value for GeneralForumTopicUnhidden to be an explicit nil

### UnsetGeneralForumTopicUnhidden
`func (o *MaybeInaccessibleMessage) UnsetGeneralForumTopicUnhidden()`

UnsetGeneralForumTopicUnhidden ensures that no value is present for GeneralForumTopicUnhidden, not even an explicit nil
### GetGiveawayCreated

`func (o *MaybeInaccessibleMessage) GetGiveawayCreated() GiveawayCreated`

GetGiveawayCreated returns the GiveawayCreated field if non-nil, zero value otherwise.

### GetGiveawayCreatedOk

`func (o *MaybeInaccessibleMessage) GetGiveawayCreatedOk() (*GiveawayCreated, bool)`

GetGiveawayCreatedOk returns a tuple with the GiveawayCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayCreated

`func (o *MaybeInaccessibleMessage) SetGiveawayCreated(v GiveawayCreated)`

SetGiveawayCreated sets GiveawayCreated field to given value.

### HasGiveawayCreated

`func (o *MaybeInaccessibleMessage) HasGiveawayCreated() bool`

HasGiveawayCreated returns a boolean if a field has been set.

### GetGiveaway

`func (o *MaybeInaccessibleMessage) GetGiveaway() Giveaway`

GetGiveaway returns the Giveaway field if non-nil, zero value otherwise.

### GetGiveawayOk

`func (o *MaybeInaccessibleMessage) GetGiveawayOk() (*Giveaway, bool)`

GetGiveawayOk returns a tuple with the Giveaway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveaway

`func (o *MaybeInaccessibleMessage) SetGiveaway(v Giveaway)`

SetGiveaway sets Giveaway field to given value.

### HasGiveaway

`func (o *MaybeInaccessibleMessage) HasGiveaway() bool`

HasGiveaway returns a boolean if a field has been set.

### GetGiveawayWinners

`func (o *MaybeInaccessibleMessage) GetGiveawayWinners() GiveawayWinners`

GetGiveawayWinners returns the GiveawayWinners field if non-nil, zero value otherwise.

### GetGiveawayWinnersOk

`func (o *MaybeInaccessibleMessage) GetGiveawayWinnersOk() (*GiveawayWinners, bool)`

GetGiveawayWinnersOk returns a tuple with the GiveawayWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayWinners

`func (o *MaybeInaccessibleMessage) SetGiveawayWinners(v GiveawayWinners)`

SetGiveawayWinners sets GiveawayWinners field to given value.

### HasGiveawayWinners

`func (o *MaybeInaccessibleMessage) HasGiveawayWinners() bool`

HasGiveawayWinners returns a boolean if a field has been set.

### GetGiveawayCompleted

`func (o *MaybeInaccessibleMessage) GetGiveawayCompleted() GiveawayCompleted`

GetGiveawayCompleted returns the GiveawayCompleted field if non-nil, zero value otherwise.

### GetGiveawayCompletedOk

`func (o *MaybeInaccessibleMessage) GetGiveawayCompletedOk() (*GiveawayCompleted, bool)`

GetGiveawayCompletedOk returns a tuple with the GiveawayCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayCompleted

`func (o *MaybeInaccessibleMessage) SetGiveawayCompleted(v GiveawayCompleted)`

SetGiveawayCompleted sets GiveawayCompleted field to given value.

### HasGiveawayCompleted

`func (o *MaybeInaccessibleMessage) HasGiveawayCompleted() bool`

HasGiveawayCompleted returns a boolean if a field has been set.

### GetPaidMessagePriceChanged

`func (o *MaybeInaccessibleMessage) GetPaidMessagePriceChanged() PaidMessagePriceChanged`

GetPaidMessagePriceChanged returns the PaidMessagePriceChanged field if non-nil, zero value otherwise.

### GetPaidMessagePriceChangedOk

`func (o *MaybeInaccessibleMessage) GetPaidMessagePriceChangedOk() (*PaidMessagePriceChanged, bool)`

GetPaidMessagePriceChangedOk returns a tuple with the PaidMessagePriceChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMessagePriceChanged

`func (o *MaybeInaccessibleMessage) SetPaidMessagePriceChanged(v PaidMessagePriceChanged)`

SetPaidMessagePriceChanged sets PaidMessagePriceChanged field to given value.

### HasPaidMessagePriceChanged

`func (o *MaybeInaccessibleMessage) HasPaidMessagePriceChanged() bool`

HasPaidMessagePriceChanged returns a boolean if a field has been set.

### GetVideoChatScheduled

`func (o *MaybeInaccessibleMessage) GetVideoChatScheduled() VideoChatScheduled`

GetVideoChatScheduled returns the VideoChatScheduled field if non-nil, zero value otherwise.

### GetVideoChatScheduledOk

`func (o *MaybeInaccessibleMessage) GetVideoChatScheduledOk() (*VideoChatScheduled, bool)`

GetVideoChatScheduledOk returns a tuple with the VideoChatScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatScheduled

`func (o *MaybeInaccessibleMessage) SetVideoChatScheduled(v VideoChatScheduled)`

SetVideoChatScheduled sets VideoChatScheduled field to given value.

### HasVideoChatScheduled

`func (o *MaybeInaccessibleMessage) HasVideoChatScheduled() bool`

HasVideoChatScheduled returns a boolean if a field has been set.

### GetVideoChatStarted

`func (o *MaybeInaccessibleMessage) GetVideoChatStarted() interface{}`

GetVideoChatStarted returns the VideoChatStarted field if non-nil, zero value otherwise.

### GetVideoChatStartedOk

`func (o *MaybeInaccessibleMessage) GetVideoChatStartedOk() (*interface{}, bool)`

GetVideoChatStartedOk returns a tuple with the VideoChatStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatStarted

`func (o *MaybeInaccessibleMessage) SetVideoChatStarted(v interface{})`

SetVideoChatStarted sets VideoChatStarted field to given value.

### HasVideoChatStarted

`func (o *MaybeInaccessibleMessage) HasVideoChatStarted() bool`

HasVideoChatStarted returns a boolean if a field has been set.

### SetVideoChatStartedNil

`func (o *MaybeInaccessibleMessage) SetVideoChatStartedNil(b bool)`

 SetVideoChatStartedNil sets the value for VideoChatStarted to be an explicit nil

### UnsetVideoChatStarted
`func (o *MaybeInaccessibleMessage) UnsetVideoChatStarted()`

UnsetVideoChatStarted ensures that no value is present for VideoChatStarted, not even an explicit nil
### GetVideoChatEnded

`func (o *MaybeInaccessibleMessage) GetVideoChatEnded() VideoChatEnded`

GetVideoChatEnded returns the VideoChatEnded field if non-nil, zero value otherwise.

### GetVideoChatEndedOk

`func (o *MaybeInaccessibleMessage) GetVideoChatEndedOk() (*VideoChatEnded, bool)`

GetVideoChatEndedOk returns a tuple with the VideoChatEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatEnded

`func (o *MaybeInaccessibleMessage) SetVideoChatEnded(v VideoChatEnded)`

SetVideoChatEnded sets VideoChatEnded field to given value.

### HasVideoChatEnded

`func (o *MaybeInaccessibleMessage) HasVideoChatEnded() bool`

HasVideoChatEnded returns a boolean if a field has been set.

### GetVideoChatParticipantsInvited

`func (o *MaybeInaccessibleMessage) GetVideoChatParticipantsInvited() VideoChatParticipantsInvited`

GetVideoChatParticipantsInvited returns the VideoChatParticipantsInvited field if non-nil, zero value otherwise.

### GetVideoChatParticipantsInvitedOk

`func (o *MaybeInaccessibleMessage) GetVideoChatParticipantsInvitedOk() (*VideoChatParticipantsInvited, bool)`

GetVideoChatParticipantsInvitedOk returns a tuple with the VideoChatParticipantsInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatParticipantsInvited

`func (o *MaybeInaccessibleMessage) SetVideoChatParticipantsInvited(v VideoChatParticipantsInvited)`

SetVideoChatParticipantsInvited sets VideoChatParticipantsInvited field to given value.

### HasVideoChatParticipantsInvited

`func (o *MaybeInaccessibleMessage) HasVideoChatParticipantsInvited() bool`

HasVideoChatParticipantsInvited returns a boolean if a field has been set.

### GetWebAppData

`func (o *MaybeInaccessibleMessage) GetWebAppData() WebAppData`

GetWebAppData returns the WebAppData field if non-nil, zero value otherwise.

### GetWebAppDataOk

`func (o *MaybeInaccessibleMessage) GetWebAppDataOk() (*WebAppData, bool)`

GetWebAppDataOk returns a tuple with the WebAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppData

`func (o *MaybeInaccessibleMessage) SetWebAppData(v WebAppData)`

SetWebAppData sets WebAppData field to given value.

### HasWebAppData

`func (o *MaybeInaccessibleMessage) HasWebAppData() bool`

HasWebAppData returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *MaybeInaccessibleMessage) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *MaybeInaccessibleMessage) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *MaybeInaccessibleMessage) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *MaybeInaccessibleMessage) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


