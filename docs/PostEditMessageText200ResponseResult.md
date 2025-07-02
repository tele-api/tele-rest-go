# PostEditMessageText200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent | 
**MessageThreadId** | Pointer to **int32** | *Optional*. Unique identifier of a message thread to which the message belongs; for supergroups only | [optional] 
**From** | Pointer to [**User**](User.md) |  | [optional] 
**SenderChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**SenderBoostCount** | Pointer to **int32** | *Optional*. If the sender of the message boosted the chat, the number of boosts added by the user | [optional] 
**SenderBusinessBot** | Pointer to [**User**](User.md) |  | [optional] 
**Date** | **int32** | Date the message was sent in Unix time. It is always a positive number, representing a valid date. | 
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
**IsFromOffline** | Pointer to **bool** | *Optional*. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message | [optional] [default to true]
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
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. True, if the caption must be shown above the message media | [optional] [default to true]
**HasMediaSpoiler** | Pointer to **bool** | *Optional*. *True*, if the message media is covered by a spoiler animation | [optional] [default to true]
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

### NewPostEditMessageText200ResponseResult

`func NewPostEditMessageText200ResponseResult(messageId int32, date int32, chat Chat, ) *PostEditMessageText200ResponseResult`

NewPostEditMessageText200ResponseResult instantiates a new PostEditMessageText200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditMessageText200ResponseResultWithDefaults

`func NewPostEditMessageText200ResponseResultWithDefaults() *PostEditMessageText200ResponseResult`

NewPostEditMessageText200ResponseResultWithDefaults instantiates a new PostEditMessageText200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *PostEditMessageText200ResponseResult) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PostEditMessageText200ResponseResult) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PostEditMessageText200ResponseResult) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetMessageThreadId

`func (o *PostEditMessageText200ResponseResult) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostEditMessageText200ResponseResult) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostEditMessageText200ResponseResult) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostEditMessageText200ResponseResult) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetFrom

`func (o *PostEditMessageText200ResponseResult) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PostEditMessageText200ResponseResult) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PostEditMessageText200ResponseResult) SetFrom(v User)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PostEditMessageText200ResponseResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSenderChat

`func (o *PostEditMessageText200ResponseResult) GetSenderChat() Chat`

GetSenderChat returns the SenderChat field if non-nil, zero value otherwise.

### GetSenderChatOk

`func (o *PostEditMessageText200ResponseResult) GetSenderChatOk() (*Chat, bool)`

GetSenderChatOk returns a tuple with the SenderChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderChat

`func (o *PostEditMessageText200ResponseResult) SetSenderChat(v Chat)`

SetSenderChat sets SenderChat field to given value.

### HasSenderChat

`func (o *PostEditMessageText200ResponseResult) HasSenderChat() bool`

HasSenderChat returns a boolean if a field has been set.

### GetSenderBoostCount

`func (o *PostEditMessageText200ResponseResult) GetSenderBoostCount() int32`

GetSenderBoostCount returns the SenderBoostCount field if non-nil, zero value otherwise.

### GetSenderBoostCountOk

`func (o *PostEditMessageText200ResponseResult) GetSenderBoostCountOk() (*int32, bool)`

GetSenderBoostCountOk returns a tuple with the SenderBoostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBoostCount

`func (o *PostEditMessageText200ResponseResult) SetSenderBoostCount(v int32)`

SetSenderBoostCount sets SenderBoostCount field to given value.

### HasSenderBoostCount

`func (o *PostEditMessageText200ResponseResult) HasSenderBoostCount() bool`

HasSenderBoostCount returns a boolean if a field has been set.

### GetSenderBusinessBot

`func (o *PostEditMessageText200ResponseResult) GetSenderBusinessBot() User`

GetSenderBusinessBot returns the SenderBusinessBot field if non-nil, zero value otherwise.

### GetSenderBusinessBotOk

`func (o *PostEditMessageText200ResponseResult) GetSenderBusinessBotOk() (*User, bool)`

GetSenderBusinessBotOk returns a tuple with the SenderBusinessBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBusinessBot

`func (o *PostEditMessageText200ResponseResult) SetSenderBusinessBot(v User)`

SetSenderBusinessBot sets SenderBusinessBot field to given value.

### HasSenderBusinessBot

`func (o *PostEditMessageText200ResponseResult) HasSenderBusinessBot() bool`

HasSenderBusinessBot returns a boolean if a field has been set.

### GetDate

`func (o *PostEditMessageText200ResponseResult) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PostEditMessageText200ResponseResult) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PostEditMessageText200ResponseResult) SetDate(v int32)`

SetDate sets Date field to given value.


### GetBusinessConnectionId

`func (o *PostEditMessageText200ResponseResult) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostEditMessageText200ResponseResult) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostEditMessageText200ResponseResult) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *PostEditMessageText200ResponseResult) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChat

`func (o *PostEditMessageText200ResponseResult) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *PostEditMessageText200ResponseResult) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *PostEditMessageText200ResponseResult) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetForwardOrigin

`func (o *PostEditMessageText200ResponseResult) GetForwardOrigin() MessageOrigin`

GetForwardOrigin returns the ForwardOrigin field if non-nil, zero value otherwise.

### GetForwardOriginOk

`func (o *PostEditMessageText200ResponseResult) GetForwardOriginOk() (*MessageOrigin, bool)`

GetForwardOriginOk returns a tuple with the ForwardOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOrigin

`func (o *PostEditMessageText200ResponseResult) SetForwardOrigin(v MessageOrigin)`

SetForwardOrigin sets ForwardOrigin field to given value.

### HasForwardOrigin

`func (o *PostEditMessageText200ResponseResult) HasForwardOrigin() bool`

HasForwardOrigin returns a boolean if a field has been set.

### GetIsTopicMessage

`func (o *PostEditMessageText200ResponseResult) GetIsTopicMessage() bool`

GetIsTopicMessage returns the IsTopicMessage field if non-nil, zero value otherwise.

### GetIsTopicMessageOk

`func (o *PostEditMessageText200ResponseResult) GetIsTopicMessageOk() (*bool, bool)`

GetIsTopicMessageOk returns a tuple with the IsTopicMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTopicMessage

`func (o *PostEditMessageText200ResponseResult) SetIsTopicMessage(v bool)`

SetIsTopicMessage sets IsTopicMessage field to given value.

### HasIsTopicMessage

`func (o *PostEditMessageText200ResponseResult) HasIsTopicMessage() bool`

HasIsTopicMessage returns a boolean if a field has been set.

### GetIsAutomaticForward

`func (o *PostEditMessageText200ResponseResult) GetIsAutomaticForward() bool`

GetIsAutomaticForward returns the IsAutomaticForward field if non-nil, zero value otherwise.

### GetIsAutomaticForwardOk

`func (o *PostEditMessageText200ResponseResult) GetIsAutomaticForwardOk() (*bool, bool)`

GetIsAutomaticForwardOk returns a tuple with the IsAutomaticForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticForward

`func (o *PostEditMessageText200ResponseResult) SetIsAutomaticForward(v bool)`

SetIsAutomaticForward sets IsAutomaticForward field to given value.

### HasIsAutomaticForward

`func (o *PostEditMessageText200ResponseResult) HasIsAutomaticForward() bool`

HasIsAutomaticForward returns a boolean if a field has been set.

### GetReplyToMessage

`func (o *PostEditMessageText200ResponseResult) GetReplyToMessage() Message`

GetReplyToMessage returns the ReplyToMessage field if non-nil, zero value otherwise.

### GetReplyToMessageOk

`func (o *PostEditMessageText200ResponseResult) GetReplyToMessageOk() (*Message, bool)`

GetReplyToMessageOk returns a tuple with the ReplyToMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToMessage

`func (o *PostEditMessageText200ResponseResult) SetReplyToMessage(v Message)`

SetReplyToMessage sets ReplyToMessage field to given value.

### HasReplyToMessage

`func (o *PostEditMessageText200ResponseResult) HasReplyToMessage() bool`

HasReplyToMessage returns a boolean if a field has been set.

### GetExternalReply

`func (o *PostEditMessageText200ResponseResult) GetExternalReply() ExternalReplyInfo`

GetExternalReply returns the ExternalReply field if non-nil, zero value otherwise.

### GetExternalReplyOk

`func (o *PostEditMessageText200ResponseResult) GetExternalReplyOk() (*ExternalReplyInfo, bool)`

GetExternalReplyOk returns a tuple with the ExternalReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReply

`func (o *PostEditMessageText200ResponseResult) SetExternalReply(v ExternalReplyInfo)`

SetExternalReply sets ExternalReply field to given value.

### HasExternalReply

`func (o *PostEditMessageText200ResponseResult) HasExternalReply() bool`

HasExternalReply returns a boolean if a field has been set.

### GetQuote

`func (o *PostEditMessageText200ResponseResult) GetQuote() TextQuote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *PostEditMessageText200ResponseResult) GetQuoteOk() (*TextQuote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *PostEditMessageText200ResponseResult) SetQuote(v TextQuote)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *PostEditMessageText200ResponseResult) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReplyToStory

`func (o *PostEditMessageText200ResponseResult) GetReplyToStory() Story`

GetReplyToStory returns the ReplyToStory field if non-nil, zero value otherwise.

### GetReplyToStoryOk

`func (o *PostEditMessageText200ResponseResult) GetReplyToStoryOk() (*Story, bool)`

GetReplyToStoryOk returns a tuple with the ReplyToStory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToStory

`func (o *PostEditMessageText200ResponseResult) SetReplyToStory(v Story)`

SetReplyToStory sets ReplyToStory field to given value.

### HasReplyToStory

`func (o *PostEditMessageText200ResponseResult) HasReplyToStory() bool`

HasReplyToStory returns a boolean if a field has been set.

### GetViaBot

`func (o *PostEditMessageText200ResponseResult) GetViaBot() User`

GetViaBot returns the ViaBot field if non-nil, zero value otherwise.

### GetViaBotOk

`func (o *PostEditMessageText200ResponseResult) GetViaBotOk() (*User, bool)`

GetViaBotOk returns a tuple with the ViaBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaBot

`func (o *PostEditMessageText200ResponseResult) SetViaBot(v User)`

SetViaBot sets ViaBot field to given value.

### HasViaBot

`func (o *PostEditMessageText200ResponseResult) HasViaBot() bool`

HasViaBot returns a boolean if a field has been set.

### GetEditDate

`func (o *PostEditMessageText200ResponseResult) GetEditDate() int32`

GetEditDate returns the EditDate field if non-nil, zero value otherwise.

### GetEditDateOk

`func (o *PostEditMessageText200ResponseResult) GetEditDateOk() (*int32, bool)`

GetEditDateOk returns a tuple with the EditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditDate

`func (o *PostEditMessageText200ResponseResult) SetEditDate(v int32)`

SetEditDate sets EditDate field to given value.

### HasEditDate

`func (o *PostEditMessageText200ResponseResult) HasEditDate() bool`

HasEditDate returns a boolean if a field has been set.

### GetHasProtectedContent

`func (o *PostEditMessageText200ResponseResult) GetHasProtectedContent() bool`

GetHasProtectedContent returns the HasProtectedContent field if non-nil, zero value otherwise.

### GetHasProtectedContentOk

`func (o *PostEditMessageText200ResponseResult) GetHasProtectedContentOk() (*bool, bool)`

GetHasProtectedContentOk returns a tuple with the HasProtectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProtectedContent

`func (o *PostEditMessageText200ResponseResult) SetHasProtectedContent(v bool)`

SetHasProtectedContent sets HasProtectedContent field to given value.

### HasHasProtectedContent

`func (o *PostEditMessageText200ResponseResult) HasHasProtectedContent() bool`

HasHasProtectedContent returns a boolean if a field has been set.

### GetIsFromOffline

`func (o *PostEditMessageText200ResponseResult) GetIsFromOffline() bool`

GetIsFromOffline returns the IsFromOffline field if non-nil, zero value otherwise.

### GetIsFromOfflineOk

`func (o *PostEditMessageText200ResponseResult) GetIsFromOfflineOk() (*bool, bool)`

GetIsFromOfflineOk returns a tuple with the IsFromOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromOffline

`func (o *PostEditMessageText200ResponseResult) SetIsFromOffline(v bool)`

SetIsFromOffline sets IsFromOffline field to given value.

### HasIsFromOffline

`func (o *PostEditMessageText200ResponseResult) HasIsFromOffline() bool`

HasIsFromOffline returns a boolean if a field has been set.

### GetMediaGroupId

`func (o *PostEditMessageText200ResponseResult) GetMediaGroupId() string`

GetMediaGroupId returns the MediaGroupId field if non-nil, zero value otherwise.

### GetMediaGroupIdOk

`func (o *PostEditMessageText200ResponseResult) GetMediaGroupIdOk() (*string, bool)`

GetMediaGroupIdOk returns a tuple with the MediaGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaGroupId

`func (o *PostEditMessageText200ResponseResult) SetMediaGroupId(v string)`

SetMediaGroupId sets MediaGroupId field to given value.

### HasMediaGroupId

`func (o *PostEditMessageText200ResponseResult) HasMediaGroupId() bool`

HasMediaGroupId returns a boolean if a field has been set.

### GetAuthorSignature

`func (o *PostEditMessageText200ResponseResult) GetAuthorSignature() string`

GetAuthorSignature returns the AuthorSignature field if non-nil, zero value otherwise.

### GetAuthorSignatureOk

`func (o *PostEditMessageText200ResponseResult) GetAuthorSignatureOk() (*string, bool)`

GetAuthorSignatureOk returns a tuple with the AuthorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorSignature

`func (o *PostEditMessageText200ResponseResult) SetAuthorSignature(v string)`

SetAuthorSignature sets AuthorSignature field to given value.

### HasAuthorSignature

`func (o *PostEditMessageText200ResponseResult) HasAuthorSignature() bool`

HasAuthorSignature returns a boolean if a field has been set.

### GetPaidStarCount

`func (o *PostEditMessageText200ResponseResult) GetPaidStarCount() int32`

GetPaidStarCount returns the PaidStarCount field if non-nil, zero value otherwise.

### GetPaidStarCountOk

`func (o *PostEditMessageText200ResponseResult) GetPaidStarCountOk() (*int32, bool)`

GetPaidStarCountOk returns a tuple with the PaidStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidStarCount

`func (o *PostEditMessageText200ResponseResult) SetPaidStarCount(v int32)`

SetPaidStarCount sets PaidStarCount field to given value.

### HasPaidStarCount

`func (o *PostEditMessageText200ResponseResult) HasPaidStarCount() bool`

HasPaidStarCount returns a boolean if a field has been set.

### GetText

`func (o *PostEditMessageText200ResponseResult) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostEditMessageText200ResponseResult) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostEditMessageText200ResponseResult) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostEditMessageText200ResponseResult) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEntities

`func (o *PostEditMessageText200ResponseResult) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PostEditMessageText200ResponseResult) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PostEditMessageText200ResponseResult) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PostEditMessageText200ResponseResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetLinkPreviewOptions

`func (o *PostEditMessageText200ResponseResult) GetLinkPreviewOptions() LinkPreviewOptions`

GetLinkPreviewOptions returns the LinkPreviewOptions field if non-nil, zero value otherwise.

### GetLinkPreviewOptionsOk

`func (o *PostEditMessageText200ResponseResult) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool)`

GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPreviewOptions

`func (o *PostEditMessageText200ResponseResult) SetLinkPreviewOptions(v LinkPreviewOptions)`

SetLinkPreviewOptions sets LinkPreviewOptions field to given value.

### HasLinkPreviewOptions

`func (o *PostEditMessageText200ResponseResult) HasLinkPreviewOptions() bool`

HasLinkPreviewOptions returns a boolean if a field has been set.

### GetEffectId

`func (o *PostEditMessageText200ResponseResult) GetEffectId() string`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *PostEditMessageText200ResponseResult) GetEffectIdOk() (*string, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *PostEditMessageText200ResponseResult) SetEffectId(v string)`

SetEffectId sets EffectId field to given value.

### HasEffectId

`func (o *PostEditMessageText200ResponseResult) HasEffectId() bool`

HasEffectId returns a boolean if a field has been set.

### GetAnimation

`func (o *PostEditMessageText200ResponseResult) GetAnimation() Animation`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *PostEditMessageText200ResponseResult) GetAnimationOk() (*Animation, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *PostEditMessageText200ResponseResult) SetAnimation(v Animation)`

SetAnimation sets Animation field to given value.

### HasAnimation

`func (o *PostEditMessageText200ResponseResult) HasAnimation() bool`

HasAnimation returns a boolean if a field has been set.

### GetAudio

`func (o *PostEditMessageText200ResponseResult) GetAudio() Audio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *PostEditMessageText200ResponseResult) GetAudioOk() (*Audio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *PostEditMessageText200ResponseResult) SetAudio(v Audio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *PostEditMessageText200ResponseResult) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *PostEditMessageText200ResponseResult) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *PostEditMessageText200ResponseResult) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *PostEditMessageText200ResponseResult) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *PostEditMessageText200ResponseResult) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetPaidMedia

`func (o *PostEditMessageText200ResponseResult) GetPaidMedia() PaidMediaInfo`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *PostEditMessageText200ResponseResult) GetPaidMediaOk() (*PaidMediaInfo, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *PostEditMessageText200ResponseResult) SetPaidMedia(v PaidMediaInfo)`

SetPaidMedia sets PaidMedia field to given value.

### HasPaidMedia

`func (o *PostEditMessageText200ResponseResult) HasPaidMedia() bool`

HasPaidMedia returns a boolean if a field has been set.

### GetPhoto

`func (o *PostEditMessageText200ResponseResult) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PostEditMessageText200ResponseResult) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PostEditMessageText200ResponseResult) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *PostEditMessageText200ResponseResult) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetSticker

`func (o *PostEditMessageText200ResponseResult) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *PostEditMessageText200ResponseResult) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *PostEditMessageText200ResponseResult) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *PostEditMessageText200ResponseResult) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetStory

`func (o *PostEditMessageText200ResponseResult) GetStory() Story`

GetStory returns the Story field if non-nil, zero value otherwise.

### GetStoryOk

`func (o *PostEditMessageText200ResponseResult) GetStoryOk() (*Story, bool)`

GetStoryOk returns a tuple with the Story field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStory

`func (o *PostEditMessageText200ResponseResult) SetStory(v Story)`

SetStory sets Story field to given value.

### HasStory

`func (o *PostEditMessageText200ResponseResult) HasStory() bool`

HasStory returns a boolean if a field has been set.

### GetVideo

`func (o *PostEditMessageText200ResponseResult) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PostEditMessageText200ResponseResult) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PostEditMessageText200ResponseResult) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PostEditMessageText200ResponseResult) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoNote

`func (o *PostEditMessageText200ResponseResult) GetVideoNote() VideoNote`

GetVideoNote returns the VideoNote field if non-nil, zero value otherwise.

### GetVideoNoteOk

`func (o *PostEditMessageText200ResponseResult) GetVideoNoteOk() (*VideoNote, bool)`

GetVideoNoteOk returns a tuple with the VideoNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoNote

`func (o *PostEditMessageText200ResponseResult) SetVideoNote(v VideoNote)`

SetVideoNote sets VideoNote field to given value.

### HasVideoNote

`func (o *PostEditMessageText200ResponseResult) HasVideoNote() bool`

HasVideoNote returns a boolean if a field has been set.

### GetVoice

`func (o *PostEditMessageText200ResponseResult) GetVoice() Voice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *PostEditMessageText200ResponseResult) GetVoiceOk() (*Voice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *PostEditMessageText200ResponseResult) SetVoice(v Voice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *PostEditMessageText200ResponseResult) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetCaption

`func (o *PostEditMessageText200ResponseResult) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *PostEditMessageText200ResponseResult) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *PostEditMessageText200ResponseResult) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *PostEditMessageText200ResponseResult) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *PostEditMessageText200ResponseResult) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *PostEditMessageText200ResponseResult) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *PostEditMessageText200ResponseResult) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *PostEditMessageText200ResponseResult) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *PostEditMessageText200ResponseResult) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *PostEditMessageText200ResponseResult) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *PostEditMessageText200ResponseResult) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *PostEditMessageText200ResponseResult) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetHasMediaSpoiler

`func (o *PostEditMessageText200ResponseResult) GetHasMediaSpoiler() bool`

GetHasMediaSpoiler returns the HasMediaSpoiler field if non-nil, zero value otherwise.

### GetHasMediaSpoilerOk

`func (o *PostEditMessageText200ResponseResult) GetHasMediaSpoilerOk() (*bool, bool)`

GetHasMediaSpoilerOk returns a tuple with the HasMediaSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMediaSpoiler

`func (o *PostEditMessageText200ResponseResult) SetHasMediaSpoiler(v bool)`

SetHasMediaSpoiler sets HasMediaSpoiler field to given value.

### HasHasMediaSpoiler

`func (o *PostEditMessageText200ResponseResult) HasHasMediaSpoiler() bool`

HasHasMediaSpoiler returns a boolean if a field has been set.

### GetContact

`func (o *PostEditMessageText200ResponseResult) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PostEditMessageText200ResponseResult) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PostEditMessageText200ResponseResult) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PostEditMessageText200ResponseResult) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDice

`func (o *PostEditMessageText200ResponseResult) GetDice() Dice`

GetDice returns the Dice field if non-nil, zero value otherwise.

### GetDiceOk

`func (o *PostEditMessageText200ResponseResult) GetDiceOk() (*Dice, bool)`

GetDiceOk returns a tuple with the Dice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDice

`func (o *PostEditMessageText200ResponseResult) SetDice(v Dice)`

SetDice sets Dice field to given value.

### HasDice

`func (o *PostEditMessageText200ResponseResult) HasDice() bool`

HasDice returns a boolean if a field has been set.

### GetGame

`func (o *PostEditMessageText200ResponseResult) GetGame() Game`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *PostEditMessageText200ResponseResult) GetGameOk() (*Game, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *PostEditMessageText200ResponseResult) SetGame(v Game)`

SetGame sets Game field to given value.

### HasGame

`func (o *PostEditMessageText200ResponseResult) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetPoll

`func (o *PostEditMessageText200ResponseResult) GetPoll() Poll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *PostEditMessageText200ResponseResult) GetPollOk() (*Poll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *PostEditMessageText200ResponseResult) SetPoll(v Poll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *PostEditMessageText200ResponseResult) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetVenue

`func (o *PostEditMessageText200ResponseResult) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *PostEditMessageText200ResponseResult) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *PostEditMessageText200ResponseResult) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *PostEditMessageText200ResponseResult) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetLocation

`func (o *PostEditMessageText200ResponseResult) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostEditMessageText200ResponseResult) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostEditMessageText200ResponseResult) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PostEditMessageText200ResponseResult) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNewChatMembers

`func (o *PostEditMessageText200ResponseResult) GetNewChatMembers() []User`

GetNewChatMembers returns the NewChatMembers field if non-nil, zero value otherwise.

### GetNewChatMembersOk

`func (o *PostEditMessageText200ResponseResult) GetNewChatMembersOk() (*[]User, bool)`

GetNewChatMembersOk returns a tuple with the NewChatMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatMembers

`func (o *PostEditMessageText200ResponseResult) SetNewChatMembers(v []User)`

SetNewChatMembers sets NewChatMembers field to given value.

### HasNewChatMembers

`func (o *PostEditMessageText200ResponseResult) HasNewChatMembers() bool`

HasNewChatMembers returns a boolean if a field has been set.

### GetLeftChatMember

`func (o *PostEditMessageText200ResponseResult) GetLeftChatMember() User`

GetLeftChatMember returns the LeftChatMember field if non-nil, zero value otherwise.

### GetLeftChatMemberOk

`func (o *PostEditMessageText200ResponseResult) GetLeftChatMemberOk() (*User, bool)`

GetLeftChatMemberOk returns a tuple with the LeftChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftChatMember

`func (o *PostEditMessageText200ResponseResult) SetLeftChatMember(v User)`

SetLeftChatMember sets LeftChatMember field to given value.

### HasLeftChatMember

`func (o *PostEditMessageText200ResponseResult) HasLeftChatMember() bool`

HasLeftChatMember returns a boolean if a field has been set.

### GetNewChatTitle

`func (o *PostEditMessageText200ResponseResult) GetNewChatTitle() string`

GetNewChatTitle returns the NewChatTitle field if non-nil, zero value otherwise.

### GetNewChatTitleOk

`func (o *PostEditMessageText200ResponseResult) GetNewChatTitleOk() (*string, bool)`

GetNewChatTitleOk returns a tuple with the NewChatTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatTitle

`func (o *PostEditMessageText200ResponseResult) SetNewChatTitle(v string)`

SetNewChatTitle sets NewChatTitle field to given value.

### HasNewChatTitle

`func (o *PostEditMessageText200ResponseResult) HasNewChatTitle() bool`

HasNewChatTitle returns a boolean if a field has been set.

### GetNewChatPhoto

`func (o *PostEditMessageText200ResponseResult) GetNewChatPhoto() []PhotoSize`

GetNewChatPhoto returns the NewChatPhoto field if non-nil, zero value otherwise.

### GetNewChatPhotoOk

`func (o *PostEditMessageText200ResponseResult) GetNewChatPhotoOk() (*[]PhotoSize, bool)`

GetNewChatPhotoOk returns a tuple with the NewChatPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewChatPhoto

`func (o *PostEditMessageText200ResponseResult) SetNewChatPhoto(v []PhotoSize)`

SetNewChatPhoto sets NewChatPhoto field to given value.

### HasNewChatPhoto

`func (o *PostEditMessageText200ResponseResult) HasNewChatPhoto() bool`

HasNewChatPhoto returns a boolean if a field has been set.

### GetDeleteChatPhoto

`func (o *PostEditMessageText200ResponseResult) GetDeleteChatPhoto() bool`

GetDeleteChatPhoto returns the DeleteChatPhoto field if non-nil, zero value otherwise.

### GetDeleteChatPhotoOk

`func (o *PostEditMessageText200ResponseResult) GetDeleteChatPhotoOk() (*bool, bool)`

GetDeleteChatPhotoOk returns a tuple with the DeleteChatPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteChatPhoto

`func (o *PostEditMessageText200ResponseResult) SetDeleteChatPhoto(v bool)`

SetDeleteChatPhoto sets DeleteChatPhoto field to given value.

### HasDeleteChatPhoto

`func (o *PostEditMessageText200ResponseResult) HasDeleteChatPhoto() bool`

HasDeleteChatPhoto returns a boolean if a field has been set.

### GetGroupChatCreated

`func (o *PostEditMessageText200ResponseResult) GetGroupChatCreated() bool`

GetGroupChatCreated returns the GroupChatCreated field if non-nil, zero value otherwise.

### GetGroupChatCreatedOk

`func (o *PostEditMessageText200ResponseResult) GetGroupChatCreatedOk() (*bool, bool)`

GetGroupChatCreatedOk returns a tuple with the GroupChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupChatCreated

`func (o *PostEditMessageText200ResponseResult) SetGroupChatCreated(v bool)`

SetGroupChatCreated sets GroupChatCreated field to given value.

### HasGroupChatCreated

`func (o *PostEditMessageText200ResponseResult) HasGroupChatCreated() bool`

HasGroupChatCreated returns a boolean if a field has been set.

### GetSupergroupChatCreated

`func (o *PostEditMessageText200ResponseResult) GetSupergroupChatCreated() bool`

GetSupergroupChatCreated returns the SupergroupChatCreated field if non-nil, zero value otherwise.

### GetSupergroupChatCreatedOk

`func (o *PostEditMessageText200ResponseResult) GetSupergroupChatCreatedOk() (*bool, bool)`

GetSupergroupChatCreatedOk returns a tuple with the SupergroupChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupergroupChatCreated

`func (o *PostEditMessageText200ResponseResult) SetSupergroupChatCreated(v bool)`

SetSupergroupChatCreated sets SupergroupChatCreated field to given value.

### HasSupergroupChatCreated

`func (o *PostEditMessageText200ResponseResult) HasSupergroupChatCreated() bool`

HasSupergroupChatCreated returns a boolean if a field has been set.

### GetChannelChatCreated

`func (o *PostEditMessageText200ResponseResult) GetChannelChatCreated() bool`

GetChannelChatCreated returns the ChannelChatCreated field if non-nil, zero value otherwise.

### GetChannelChatCreatedOk

`func (o *PostEditMessageText200ResponseResult) GetChannelChatCreatedOk() (*bool, bool)`

GetChannelChatCreatedOk returns a tuple with the ChannelChatCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelChatCreated

`func (o *PostEditMessageText200ResponseResult) SetChannelChatCreated(v bool)`

SetChannelChatCreated sets ChannelChatCreated field to given value.

### HasChannelChatCreated

`func (o *PostEditMessageText200ResponseResult) HasChannelChatCreated() bool`

HasChannelChatCreated returns a boolean if a field has been set.

### GetMessageAutoDeleteTimerChanged

`func (o *PostEditMessageText200ResponseResult) GetMessageAutoDeleteTimerChanged() MessageAutoDeleteTimerChanged`

GetMessageAutoDeleteTimerChanged returns the MessageAutoDeleteTimerChanged field if non-nil, zero value otherwise.

### GetMessageAutoDeleteTimerChangedOk

`func (o *PostEditMessageText200ResponseResult) GetMessageAutoDeleteTimerChangedOk() (*MessageAutoDeleteTimerChanged, bool)`

GetMessageAutoDeleteTimerChangedOk returns a tuple with the MessageAutoDeleteTimerChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageAutoDeleteTimerChanged

`func (o *PostEditMessageText200ResponseResult) SetMessageAutoDeleteTimerChanged(v MessageAutoDeleteTimerChanged)`

SetMessageAutoDeleteTimerChanged sets MessageAutoDeleteTimerChanged field to given value.

### HasMessageAutoDeleteTimerChanged

`func (o *PostEditMessageText200ResponseResult) HasMessageAutoDeleteTimerChanged() bool`

HasMessageAutoDeleteTimerChanged returns a boolean if a field has been set.

### GetMigrateToChatId

`func (o *PostEditMessageText200ResponseResult) GetMigrateToChatId() int32`

GetMigrateToChatId returns the MigrateToChatId field if non-nil, zero value otherwise.

### GetMigrateToChatIdOk

`func (o *PostEditMessageText200ResponseResult) GetMigrateToChatIdOk() (*int32, bool)`

GetMigrateToChatIdOk returns a tuple with the MigrateToChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateToChatId

`func (o *PostEditMessageText200ResponseResult) SetMigrateToChatId(v int32)`

SetMigrateToChatId sets MigrateToChatId field to given value.

### HasMigrateToChatId

`func (o *PostEditMessageText200ResponseResult) HasMigrateToChatId() bool`

HasMigrateToChatId returns a boolean if a field has been set.

### GetMigrateFromChatId

`func (o *PostEditMessageText200ResponseResult) GetMigrateFromChatId() int32`

GetMigrateFromChatId returns the MigrateFromChatId field if non-nil, zero value otherwise.

### GetMigrateFromChatIdOk

`func (o *PostEditMessageText200ResponseResult) GetMigrateFromChatIdOk() (*int32, bool)`

GetMigrateFromChatIdOk returns a tuple with the MigrateFromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateFromChatId

`func (o *PostEditMessageText200ResponseResult) SetMigrateFromChatId(v int32)`

SetMigrateFromChatId sets MigrateFromChatId field to given value.

### HasMigrateFromChatId

`func (o *PostEditMessageText200ResponseResult) HasMigrateFromChatId() bool`

HasMigrateFromChatId returns a boolean if a field has been set.

### GetPinnedMessage

`func (o *PostEditMessageText200ResponseResult) GetPinnedMessage() MaybeInaccessibleMessage`

GetPinnedMessage returns the PinnedMessage field if non-nil, zero value otherwise.

### GetPinnedMessageOk

`func (o *PostEditMessageText200ResponseResult) GetPinnedMessageOk() (*MaybeInaccessibleMessage, bool)`

GetPinnedMessageOk returns a tuple with the PinnedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedMessage

`func (o *PostEditMessageText200ResponseResult) SetPinnedMessage(v MaybeInaccessibleMessage)`

SetPinnedMessage sets PinnedMessage field to given value.

### HasPinnedMessage

`func (o *PostEditMessageText200ResponseResult) HasPinnedMessage() bool`

HasPinnedMessage returns a boolean if a field has been set.

### GetInvoice

`func (o *PostEditMessageText200ResponseResult) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *PostEditMessageText200ResponseResult) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *PostEditMessageText200ResponseResult) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *PostEditMessageText200ResponseResult) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetSuccessfulPayment

`func (o *PostEditMessageText200ResponseResult) GetSuccessfulPayment() SuccessfulPayment`

GetSuccessfulPayment returns the SuccessfulPayment field if non-nil, zero value otherwise.

### GetSuccessfulPaymentOk

`func (o *PostEditMessageText200ResponseResult) GetSuccessfulPaymentOk() (*SuccessfulPayment, bool)`

GetSuccessfulPaymentOk returns a tuple with the SuccessfulPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulPayment

`func (o *PostEditMessageText200ResponseResult) SetSuccessfulPayment(v SuccessfulPayment)`

SetSuccessfulPayment sets SuccessfulPayment field to given value.

### HasSuccessfulPayment

`func (o *PostEditMessageText200ResponseResult) HasSuccessfulPayment() bool`

HasSuccessfulPayment returns a boolean if a field has been set.

### GetRefundedPayment

`func (o *PostEditMessageText200ResponseResult) GetRefundedPayment() RefundedPayment`

GetRefundedPayment returns the RefundedPayment field if non-nil, zero value otherwise.

### GetRefundedPaymentOk

`func (o *PostEditMessageText200ResponseResult) GetRefundedPaymentOk() (*RefundedPayment, bool)`

GetRefundedPaymentOk returns a tuple with the RefundedPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedPayment

`func (o *PostEditMessageText200ResponseResult) SetRefundedPayment(v RefundedPayment)`

SetRefundedPayment sets RefundedPayment field to given value.

### HasRefundedPayment

`func (o *PostEditMessageText200ResponseResult) HasRefundedPayment() bool`

HasRefundedPayment returns a boolean if a field has been set.

### GetUsersShared

`func (o *PostEditMessageText200ResponseResult) GetUsersShared() UsersShared`

GetUsersShared returns the UsersShared field if non-nil, zero value otherwise.

### GetUsersSharedOk

`func (o *PostEditMessageText200ResponseResult) GetUsersSharedOk() (*UsersShared, bool)`

GetUsersSharedOk returns a tuple with the UsersShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersShared

`func (o *PostEditMessageText200ResponseResult) SetUsersShared(v UsersShared)`

SetUsersShared sets UsersShared field to given value.

### HasUsersShared

`func (o *PostEditMessageText200ResponseResult) HasUsersShared() bool`

HasUsersShared returns a boolean if a field has been set.

### GetChatShared

`func (o *PostEditMessageText200ResponseResult) GetChatShared() ChatShared`

GetChatShared returns the ChatShared field if non-nil, zero value otherwise.

### GetChatSharedOk

`func (o *PostEditMessageText200ResponseResult) GetChatSharedOk() (*ChatShared, bool)`

GetChatSharedOk returns a tuple with the ChatShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatShared

`func (o *PostEditMessageText200ResponseResult) SetChatShared(v ChatShared)`

SetChatShared sets ChatShared field to given value.

### HasChatShared

`func (o *PostEditMessageText200ResponseResult) HasChatShared() bool`

HasChatShared returns a boolean if a field has been set.

### GetGift

`func (o *PostEditMessageText200ResponseResult) GetGift() GiftInfo`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *PostEditMessageText200ResponseResult) GetGiftOk() (*GiftInfo, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *PostEditMessageText200ResponseResult) SetGift(v GiftInfo)`

SetGift sets Gift field to given value.

### HasGift

`func (o *PostEditMessageText200ResponseResult) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetUniqueGift

`func (o *PostEditMessageText200ResponseResult) GetUniqueGift() UniqueGiftInfo`

GetUniqueGift returns the UniqueGift field if non-nil, zero value otherwise.

### GetUniqueGiftOk

`func (o *PostEditMessageText200ResponseResult) GetUniqueGiftOk() (*UniqueGiftInfo, bool)`

GetUniqueGiftOk returns a tuple with the UniqueGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGift

`func (o *PostEditMessageText200ResponseResult) SetUniqueGift(v UniqueGiftInfo)`

SetUniqueGift sets UniqueGift field to given value.

### HasUniqueGift

`func (o *PostEditMessageText200ResponseResult) HasUniqueGift() bool`

HasUniqueGift returns a boolean if a field has been set.

### GetConnectedWebsite

`func (o *PostEditMessageText200ResponseResult) GetConnectedWebsite() string`

GetConnectedWebsite returns the ConnectedWebsite field if non-nil, zero value otherwise.

### GetConnectedWebsiteOk

`func (o *PostEditMessageText200ResponseResult) GetConnectedWebsiteOk() (*string, bool)`

GetConnectedWebsiteOk returns a tuple with the ConnectedWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWebsite

`func (o *PostEditMessageText200ResponseResult) SetConnectedWebsite(v string)`

SetConnectedWebsite sets ConnectedWebsite field to given value.

### HasConnectedWebsite

`func (o *PostEditMessageText200ResponseResult) HasConnectedWebsite() bool`

HasConnectedWebsite returns a boolean if a field has been set.

### GetWriteAccessAllowed

`func (o *PostEditMessageText200ResponseResult) GetWriteAccessAllowed() WriteAccessAllowed`

GetWriteAccessAllowed returns the WriteAccessAllowed field if non-nil, zero value otherwise.

### GetWriteAccessAllowedOk

`func (o *PostEditMessageText200ResponseResult) GetWriteAccessAllowedOk() (*WriteAccessAllowed, bool)`

GetWriteAccessAllowedOk returns a tuple with the WriteAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAccessAllowed

`func (o *PostEditMessageText200ResponseResult) SetWriteAccessAllowed(v WriteAccessAllowed)`

SetWriteAccessAllowed sets WriteAccessAllowed field to given value.

### HasWriteAccessAllowed

`func (o *PostEditMessageText200ResponseResult) HasWriteAccessAllowed() bool`

HasWriteAccessAllowed returns a boolean if a field has been set.

### GetPassportData

`func (o *PostEditMessageText200ResponseResult) GetPassportData() PassportData`

GetPassportData returns the PassportData field if non-nil, zero value otherwise.

### GetPassportDataOk

`func (o *PostEditMessageText200ResponseResult) GetPassportDataOk() (*PassportData, bool)`

GetPassportDataOk returns a tuple with the PassportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportData

`func (o *PostEditMessageText200ResponseResult) SetPassportData(v PassportData)`

SetPassportData sets PassportData field to given value.

### HasPassportData

`func (o *PostEditMessageText200ResponseResult) HasPassportData() bool`

HasPassportData returns a boolean if a field has been set.

### GetProximityAlertTriggered

`func (o *PostEditMessageText200ResponseResult) GetProximityAlertTriggered() ProximityAlertTriggered`

GetProximityAlertTriggered returns the ProximityAlertTriggered field if non-nil, zero value otherwise.

### GetProximityAlertTriggeredOk

`func (o *PostEditMessageText200ResponseResult) GetProximityAlertTriggeredOk() (*ProximityAlertTriggered, bool)`

GetProximityAlertTriggeredOk returns a tuple with the ProximityAlertTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertTriggered

`func (o *PostEditMessageText200ResponseResult) SetProximityAlertTriggered(v ProximityAlertTriggered)`

SetProximityAlertTriggered sets ProximityAlertTriggered field to given value.

### HasProximityAlertTriggered

`func (o *PostEditMessageText200ResponseResult) HasProximityAlertTriggered() bool`

HasProximityAlertTriggered returns a boolean if a field has been set.

### GetBoostAdded

`func (o *PostEditMessageText200ResponseResult) GetBoostAdded() ChatBoostAdded`

GetBoostAdded returns the BoostAdded field if non-nil, zero value otherwise.

### GetBoostAddedOk

`func (o *PostEditMessageText200ResponseResult) GetBoostAddedOk() (*ChatBoostAdded, bool)`

GetBoostAddedOk returns a tuple with the BoostAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostAdded

`func (o *PostEditMessageText200ResponseResult) SetBoostAdded(v ChatBoostAdded)`

SetBoostAdded sets BoostAdded field to given value.

### HasBoostAdded

`func (o *PostEditMessageText200ResponseResult) HasBoostAdded() bool`

HasBoostAdded returns a boolean if a field has been set.

### GetChatBackgroundSet

`func (o *PostEditMessageText200ResponseResult) GetChatBackgroundSet() ChatBackground`

GetChatBackgroundSet returns the ChatBackgroundSet field if non-nil, zero value otherwise.

### GetChatBackgroundSetOk

`func (o *PostEditMessageText200ResponseResult) GetChatBackgroundSetOk() (*ChatBackground, bool)`

GetChatBackgroundSetOk returns a tuple with the ChatBackgroundSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatBackgroundSet

`func (o *PostEditMessageText200ResponseResult) SetChatBackgroundSet(v ChatBackground)`

SetChatBackgroundSet sets ChatBackgroundSet field to given value.

### HasChatBackgroundSet

`func (o *PostEditMessageText200ResponseResult) HasChatBackgroundSet() bool`

HasChatBackgroundSet returns a boolean if a field has been set.

### GetForumTopicCreated

`func (o *PostEditMessageText200ResponseResult) GetForumTopicCreated() ForumTopicCreated`

GetForumTopicCreated returns the ForumTopicCreated field if non-nil, zero value otherwise.

### GetForumTopicCreatedOk

`func (o *PostEditMessageText200ResponseResult) GetForumTopicCreatedOk() (*ForumTopicCreated, bool)`

GetForumTopicCreatedOk returns a tuple with the ForumTopicCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicCreated

`func (o *PostEditMessageText200ResponseResult) SetForumTopicCreated(v ForumTopicCreated)`

SetForumTopicCreated sets ForumTopicCreated field to given value.

### HasForumTopicCreated

`func (o *PostEditMessageText200ResponseResult) HasForumTopicCreated() bool`

HasForumTopicCreated returns a boolean if a field has been set.

### GetForumTopicEdited

`func (o *PostEditMessageText200ResponseResult) GetForumTopicEdited() ForumTopicEdited`

GetForumTopicEdited returns the ForumTopicEdited field if non-nil, zero value otherwise.

### GetForumTopicEditedOk

`func (o *PostEditMessageText200ResponseResult) GetForumTopicEditedOk() (*ForumTopicEdited, bool)`

GetForumTopicEditedOk returns a tuple with the ForumTopicEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicEdited

`func (o *PostEditMessageText200ResponseResult) SetForumTopicEdited(v ForumTopicEdited)`

SetForumTopicEdited sets ForumTopicEdited field to given value.

### HasForumTopicEdited

`func (o *PostEditMessageText200ResponseResult) HasForumTopicEdited() bool`

HasForumTopicEdited returns a boolean if a field has been set.

### GetForumTopicClosed

`func (o *PostEditMessageText200ResponseResult) GetForumTopicClosed() interface{}`

GetForumTopicClosed returns the ForumTopicClosed field if non-nil, zero value otherwise.

### GetForumTopicClosedOk

`func (o *PostEditMessageText200ResponseResult) GetForumTopicClosedOk() (*interface{}, bool)`

GetForumTopicClosedOk returns a tuple with the ForumTopicClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicClosed

`func (o *PostEditMessageText200ResponseResult) SetForumTopicClosed(v interface{})`

SetForumTopicClosed sets ForumTopicClosed field to given value.

### HasForumTopicClosed

`func (o *PostEditMessageText200ResponseResult) HasForumTopicClosed() bool`

HasForumTopicClosed returns a boolean if a field has been set.

### SetForumTopicClosedNil

`func (o *PostEditMessageText200ResponseResult) SetForumTopicClosedNil(b bool)`

 SetForumTopicClosedNil sets the value for ForumTopicClosed to be an explicit nil

### UnsetForumTopicClosed
`func (o *PostEditMessageText200ResponseResult) UnsetForumTopicClosed()`

UnsetForumTopicClosed ensures that no value is present for ForumTopicClosed, not even an explicit nil
### GetForumTopicReopened

`func (o *PostEditMessageText200ResponseResult) GetForumTopicReopened() interface{}`

GetForumTopicReopened returns the ForumTopicReopened field if non-nil, zero value otherwise.

### GetForumTopicReopenedOk

`func (o *PostEditMessageText200ResponseResult) GetForumTopicReopenedOk() (*interface{}, bool)`

GetForumTopicReopenedOk returns a tuple with the ForumTopicReopened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicReopened

`func (o *PostEditMessageText200ResponseResult) SetForumTopicReopened(v interface{})`

SetForumTopicReopened sets ForumTopicReopened field to given value.

### HasForumTopicReopened

`func (o *PostEditMessageText200ResponseResult) HasForumTopicReopened() bool`

HasForumTopicReopened returns a boolean if a field has been set.

### SetForumTopicReopenedNil

`func (o *PostEditMessageText200ResponseResult) SetForumTopicReopenedNil(b bool)`

 SetForumTopicReopenedNil sets the value for ForumTopicReopened to be an explicit nil

### UnsetForumTopicReopened
`func (o *PostEditMessageText200ResponseResult) UnsetForumTopicReopened()`

UnsetForumTopicReopened ensures that no value is present for ForumTopicReopened, not even an explicit nil
### GetGeneralForumTopicHidden

`func (o *PostEditMessageText200ResponseResult) GetGeneralForumTopicHidden() interface{}`

GetGeneralForumTopicHidden returns the GeneralForumTopicHidden field if non-nil, zero value otherwise.

### GetGeneralForumTopicHiddenOk

`func (o *PostEditMessageText200ResponseResult) GetGeneralForumTopicHiddenOk() (*interface{}, bool)`

GetGeneralForumTopicHiddenOk returns a tuple with the GeneralForumTopicHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralForumTopicHidden

`func (o *PostEditMessageText200ResponseResult) SetGeneralForumTopicHidden(v interface{})`

SetGeneralForumTopicHidden sets GeneralForumTopicHidden field to given value.

### HasGeneralForumTopicHidden

`func (o *PostEditMessageText200ResponseResult) HasGeneralForumTopicHidden() bool`

HasGeneralForumTopicHidden returns a boolean if a field has been set.

### SetGeneralForumTopicHiddenNil

`func (o *PostEditMessageText200ResponseResult) SetGeneralForumTopicHiddenNil(b bool)`

 SetGeneralForumTopicHiddenNil sets the value for GeneralForumTopicHidden to be an explicit nil

### UnsetGeneralForumTopicHidden
`func (o *PostEditMessageText200ResponseResult) UnsetGeneralForumTopicHidden()`

UnsetGeneralForumTopicHidden ensures that no value is present for GeneralForumTopicHidden, not even an explicit nil
### GetGeneralForumTopicUnhidden

`func (o *PostEditMessageText200ResponseResult) GetGeneralForumTopicUnhidden() interface{}`

GetGeneralForumTopicUnhidden returns the GeneralForumTopicUnhidden field if non-nil, zero value otherwise.

### GetGeneralForumTopicUnhiddenOk

`func (o *PostEditMessageText200ResponseResult) GetGeneralForumTopicUnhiddenOk() (*interface{}, bool)`

GetGeneralForumTopicUnhiddenOk returns a tuple with the GeneralForumTopicUnhidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralForumTopicUnhidden

`func (o *PostEditMessageText200ResponseResult) SetGeneralForumTopicUnhidden(v interface{})`

SetGeneralForumTopicUnhidden sets GeneralForumTopicUnhidden field to given value.

### HasGeneralForumTopicUnhidden

`func (o *PostEditMessageText200ResponseResult) HasGeneralForumTopicUnhidden() bool`

HasGeneralForumTopicUnhidden returns a boolean if a field has been set.

### SetGeneralForumTopicUnhiddenNil

`func (o *PostEditMessageText200ResponseResult) SetGeneralForumTopicUnhiddenNil(b bool)`

 SetGeneralForumTopicUnhiddenNil sets the value for GeneralForumTopicUnhidden to be an explicit nil

### UnsetGeneralForumTopicUnhidden
`func (o *PostEditMessageText200ResponseResult) UnsetGeneralForumTopicUnhidden()`

UnsetGeneralForumTopicUnhidden ensures that no value is present for GeneralForumTopicUnhidden, not even an explicit nil
### GetGiveawayCreated

`func (o *PostEditMessageText200ResponseResult) GetGiveawayCreated() GiveawayCreated`

GetGiveawayCreated returns the GiveawayCreated field if non-nil, zero value otherwise.

### GetGiveawayCreatedOk

`func (o *PostEditMessageText200ResponseResult) GetGiveawayCreatedOk() (*GiveawayCreated, bool)`

GetGiveawayCreatedOk returns a tuple with the GiveawayCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayCreated

`func (o *PostEditMessageText200ResponseResult) SetGiveawayCreated(v GiveawayCreated)`

SetGiveawayCreated sets GiveawayCreated field to given value.

### HasGiveawayCreated

`func (o *PostEditMessageText200ResponseResult) HasGiveawayCreated() bool`

HasGiveawayCreated returns a boolean if a field has been set.

### GetGiveaway

`func (o *PostEditMessageText200ResponseResult) GetGiveaway() Giveaway`

GetGiveaway returns the Giveaway field if non-nil, zero value otherwise.

### GetGiveawayOk

`func (o *PostEditMessageText200ResponseResult) GetGiveawayOk() (*Giveaway, bool)`

GetGiveawayOk returns a tuple with the Giveaway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveaway

`func (o *PostEditMessageText200ResponseResult) SetGiveaway(v Giveaway)`

SetGiveaway sets Giveaway field to given value.

### HasGiveaway

`func (o *PostEditMessageText200ResponseResult) HasGiveaway() bool`

HasGiveaway returns a boolean if a field has been set.

### GetGiveawayWinners

`func (o *PostEditMessageText200ResponseResult) GetGiveawayWinners() GiveawayWinners`

GetGiveawayWinners returns the GiveawayWinners field if non-nil, zero value otherwise.

### GetGiveawayWinnersOk

`func (o *PostEditMessageText200ResponseResult) GetGiveawayWinnersOk() (*GiveawayWinners, bool)`

GetGiveawayWinnersOk returns a tuple with the GiveawayWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayWinners

`func (o *PostEditMessageText200ResponseResult) SetGiveawayWinners(v GiveawayWinners)`

SetGiveawayWinners sets GiveawayWinners field to given value.

### HasGiveawayWinners

`func (o *PostEditMessageText200ResponseResult) HasGiveawayWinners() bool`

HasGiveawayWinners returns a boolean if a field has been set.

### GetGiveawayCompleted

`func (o *PostEditMessageText200ResponseResult) GetGiveawayCompleted() GiveawayCompleted`

GetGiveawayCompleted returns the GiveawayCompleted field if non-nil, zero value otherwise.

### GetGiveawayCompletedOk

`func (o *PostEditMessageText200ResponseResult) GetGiveawayCompletedOk() (*GiveawayCompleted, bool)`

GetGiveawayCompletedOk returns a tuple with the GiveawayCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayCompleted

`func (o *PostEditMessageText200ResponseResult) SetGiveawayCompleted(v GiveawayCompleted)`

SetGiveawayCompleted sets GiveawayCompleted field to given value.

### HasGiveawayCompleted

`func (o *PostEditMessageText200ResponseResult) HasGiveawayCompleted() bool`

HasGiveawayCompleted returns a boolean if a field has been set.

### GetPaidMessagePriceChanged

`func (o *PostEditMessageText200ResponseResult) GetPaidMessagePriceChanged() PaidMessagePriceChanged`

GetPaidMessagePriceChanged returns the PaidMessagePriceChanged field if non-nil, zero value otherwise.

### GetPaidMessagePriceChangedOk

`func (o *PostEditMessageText200ResponseResult) GetPaidMessagePriceChangedOk() (*PaidMessagePriceChanged, bool)`

GetPaidMessagePriceChangedOk returns a tuple with the PaidMessagePriceChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMessagePriceChanged

`func (o *PostEditMessageText200ResponseResult) SetPaidMessagePriceChanged(v PaidMessagePriceChanged)`

SetPaidMessagePriceChanged sets PaidMessagePriceChanged field to given value.

### HasPaidMessagePriceChanged

`func (o *PostEditMessageText200ResponseResult) HasPaidMessagePriceChanged() bool`

HasPaidMessagePriceChanged returns a boolean if a field has been set.

### GetVideoChatScheduled

`func (o *PostEditMessageText200ResponseResult) GetVideoChatScheduled() VideoChatScheduled`

GetVideoChatScheduled returns the VideoChatScheduled field if non-nil, zero value otherwise.

### GetVideoChatScheduledOk

`func (o *PostEditMessageText200ResponseResult) GetVideoChatScheduledOk() (*VideoChatScheduled, bool)`

GetVideoChatScheduledOk returns a tuple with the VideoChatScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatScheduled

`func (o *PostEditMessageText200ResponseResult) SetVideoChatScheduled(v VideoChatScheduled)`

SetVideoChatScheduled sets VideoChatScheduled field to given value.

### HasVideoChatScheduled

`func (o *PostEditMessageText200ResponseResult) HasVideoChatScheduled() bool`

HasVideoChatScheduled returns a boolean if a field has been set.

### GetVideoChatStarted

`func (o *PostEditMessageText200ResponseResult) GetVideoChatStarted() interface{}`

GetVideoChatStarted returns the VideoChatStarted field if non-nil, zero value otherwise.

### GetVideoChatStartedOk

`func (o *PostEditMessageText200ResponseResult) GetVideoChatStartedOk() (*interface{}, bool)`

GetVideoChatStartedOk returns a tuple with the VideoChatStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatStarted

`func (o *PostEditMessageText200ResponseResult) SetVideoChatStarted(v interface{})`

SetVideoChatStarted sets VideoChatStarted field to given value.

### HasVideoChatStarted

`func (o *PostEditMessageText200ResponseResult) HasVideoChatStarted() bool`

HasVideoChatStarted returns a boolean if a field has been set.

### SetVideoChatStartedNil

`func (o *PostEditMessageText200ResponseResult) SetVideoChatStartedNil(b bool)`

 SetVideoChatStartedNil sets the value for VideoChatStarted to be an explicit nil

### UnsetVideoChatStarted
`func (o *PostEditMessageText200ResponseResult) UnsetVideoChatStarted()`

UnsetVideoChatStarted ensures that no value is present for VideoChatStarted, not even an explicit nil
### GetVideoChatEnded

`func (o *PostEditMessageText200ResponseResult) GetVideoChatEnded() VideoChatEnded`

GetVideoChatEnded returns the VideoChatEnded field if non-nil, zero value otherwise.

### GetVideoChatEndedOk

`func (o *PostEditMessageText200ResponseResult) GetVideoChatEndedOk() (*VideoChatEnded, bool)`

GetVideoChatEndedOk returns a tuple with the VideoChatEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatEnded

`func (o *PostEditMessageText200ResponseResult) SetVideoChatEnded(v VideoChatEnded)`

SetVideoChatEnded sets VideoChatEnded field to given value.

### HasVideoChatEnded

`func (o *PostEditMessageText200ResponseResult) HasVideoChatEnded() bool`

HasVideoChatEnded returns a boolean if a field has been set.

### GetVideoChatParticipantsInvited

`func (o *PostEditMessageText200ResponseResult) GetVideoChatParticipantsInvited() VideoChatParticipantsInvited`

GetVideoChatParticipantsInvited returns the VideoChatParticipantsInvited field if non-nil, zero value otherwise.

### GetVideoChatParticipantsInvitedOk

`func (o *PostEditMessageText200ResponseResult) GetVideoChatParticipantsInvitedOk() (*VideoChatParticipantsInvited, bool)`

GetVideoChatParticipantsInvitedOk returns a tuple with the VideoChatParticipantsInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatParticipantsInvited

`func (o *PostEditMessageText200ResponseResult) SetVideoChatParticipantsInvited(v VideoChatParticipantsInvited)`

SetVideoChatParticipantsInvited sets VideoChatParticipantsInvited field to given value.

### HasVideoChatParticipantsInvited

`func (o *PostEditMessageText200ResponseResult) HasVideoChatParticipantsInvited() bool`

HasVideoChatParticipantsInvited returns a boolean if a field has been set.

### GetWebAppData

`func (o *PostEditMessageText200ResponseResult) GetWebAppData() WebAppData`

GetWebAppData returns the WebAppData field if non-nil, zero value otherwise.

### GetWebAppDataOk

`func (o *PostEditMessageText200ResponseResult) GetWebAppDataOk() (*WebAppData, bool)`

GetWebAppDataOk returns a tuple with the WebAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppData

`func (o *PostEditMessageText200ResponseResult) SetWebAppData(v WebAppData)`

SetWebAppData sets WebAppData field to given value.

### HasWebAppData

`func (o *PostEditMessageText200ResponseResult) HasWebAppData() bool`

HasWebAppData returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *PostEditMessageText200ResponseResult) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *PostEditMessageText200ResponseResult) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *PostEditMessageText200ResponseResult) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *PostEditMessageText200ResponseResult) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


