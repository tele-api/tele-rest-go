# ExternalReplyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | [**MessageOrigin**](MessageOrigin.md) |  | 
**Chat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**MessageId** | Pointer to **int32** | *Optional*. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel. | [optional] 
**LinkPreviewOptions** | Pointer to [**LinkPreviewOptions**](LinkPreviewOptions.md) |  | [optional] 
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
**HasMediaSpoiler** | Pointer to **bool** | *Optional*. *True*, if the message media is covered by a spoiler animation | [optional] [default to true]
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Dice** | Pointer to [**Dice**](Dice.md) |  | [optional] 
**Game** | Pointer to [**Game**](Game.md) |  | [optional] 
**Giveaway** | Pointer to [**Giveaway**](Giveaway.md) |  | [optional] 
**GiveawayWinners** | Pointer to [**GiveawayWinners**](GiveawayWinners.md) |  | [optional] 
**Invoice** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Poll** | Pointer to [**Poll**](Poll.md) |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 

## Methods

### NewExternalReplyInfo

`func NewExternalReplyInfo(origin MessageOrigin, ) *ExternalReplyInfo`

NewExternalReplyInfo instantiates a new ExternalReplyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalReplyInfoWithDefaults

`func NewExternalReplyInfoWithDefaults() *ExternalReplyInfo`

NewExternalReplyInfoWithDefaults instantiates a new ExternalReplyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *ExternalReplyInfo) GetOrigin() MessageOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ExternalReplyInfo) GetOriginOk() (*MessageOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ExternalReplyInfo) SetOrigin(v MessageOrigin)`

SetOrigin sets Origin field to given value.


### GetChat

`func (o *ExternalReplyInfo) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ExternalReplyInfo) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ExternalReplyInfo) SetChat(v Chat)`

SetChat sets Chat field to given value.

### HasChat

`func (o *ExternalReplyInfo) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetMessageId

`func (o *ExternalReplyInfo) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ExternalReplyInfo) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ExternalReplyInfo) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ExternalReplyInfo) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetLinkPreviewOptions

`func (o *ExternalReplyInfo) GetLinkPreviewOptions() LinkPreviewOptions`

GetLinkPreviewOptions returns the LinkPreviewOptions field if non-nil, zero value otherwise.

### GetLinkPreviewOptionsOk

`func (o *ExternalReplyInfo) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool)`

GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPreviewOptions

`func (o *ExternalReplyInfo) SetLinkPreviewOptions(v LinkPreviewOptions)`

SetLinkPreviewOptions sets LinkPreviewOptions field to given value.

### HasLinkPreviewOptions

`func (o *ExternalReplyInfo) HasLinkPreviewOptions() bool`

HasLinkPreviewOptions returns a boolean if a field has been set.

### GetAnimation

`func (o *ExternalReplyInfo) GetAnimation() Animation`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *ExternalReplyInfo) GetAnimationOk() (*Animation, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *ExternalReplyInfo) SetAnimation(v Animation)`

SetAnimation sets Animation field to given value.

### HasAnimation

`func (o *ExternalReplyInfo) HasAnimation() bool`

HasAnimation returns a boolean if a field has been set.

### GetAudio

`func (o *ExternalReplyInfo) GetAudio() Audio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *ExternalReplyInfo) GetAudioOk() (*Audio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *ExternalReplyInfo) SetAudio(v Audio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *ExternalReplyInfo) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *ExternalReplyInfo) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ExternalReplyInfo) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ExternalReplyInfo) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ExternalReplyInfo) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetPaidMedia

`func (o *ExternalReplyInfo) GetPaidMedia() PaidMediaInfo`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *ExternalReplyInfo) GetPaidMediaOk() (*PaidMediaInfo, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *ExternalReplyInfo) SetPaidMedia(v PaidMediaInfo)`

SetPaidMedia sets PaidMedia field to given value.

### HasPaidMedia

`func (o *ExternalReplyInfo) HasPaidMedia() bool`

HasPaidMedia returns a boolean if a field has been set.

### GetPhoto

`func (o *ExternalReplyInfo) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ExternalReplyInfo) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ExternalReplyInfo) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ExternalReplyInfo) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetSticker

`func (o *ExternalReplyInfo) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *ExternalReplyInfo) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *ExternalReplyInfo) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *ExternalReplyInfo) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetStory

`func (o *ExternalReplyInfo) GetStory() Story`

GetStory returns the Story field if non-nil, zero value otherwise.

### GetStoryOk

`func (o *ExternalReplyInfo) GetStoryOk() (*Story, bool)`

GetStoryOk returns a tuple with the Story field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStory

`func (o *ExternalReplyInfo) SetStory(v Story)`

SetStory sets Story field to given value.

### HasStory

`func (o *ExternalReplyInfo) HasStory() bool`

HasStory returns a boolean if a field has been set.

### GetVideo

`func (o *ExternalReplyInfo) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *ExternalReplyInfo) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *ExternalReplyInfo) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *ExternalReplyInfo) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoNote

`func (o *ExternalReplyInfo) GetVideoNote() VideoNote`

GetVideoNote returns the VideoNote field if non-nil, zero value otherwise.

### GetVideoNoteOk

`func (o *ExternalReplyInfo) GetVideoNoteOk() (*VideoNote, bool)`

GetVideoNoteOk returns a tuple with the VideoNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoNote

`func (o *ExternalReplyInfo) SetVideoNote(v VideoNote)`

SetVideoNote sets VideoNote field to given value.

### HasVideoNote

`func (o *ExternalReplyInfo) HasVideoNote() bool`

HasVideoNote returns a boolean if a field has been set.

### GetVoice

`func (o *ExternalReplyInfo) GetVoice() Voice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *ExternalReplyInfo) GetVoiceOk() (*Voice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *ExternalReplyInfo) SetVoice(v Voice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *ExternalReplyInfo) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetHasMediaSpoiler

`func (o *ExternalReplyInfo) GetHasMediaSpoiler() bool`

GetHasMediaSpoiler returns the HasMediaSpoiler field if non-nil, zero value otherwise.

### GetHasMediaSpoilerOk

`func (o *ExternalReplyInfo) GetHasMediaSpoilerOk() (*bool, bool)`

GetHasMediaSpoilerOk returns a tuple with the HasMediaSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMediaSpoiler

`func (o *ExternalReplyInfo) SetHasMediaSpoiler(v bool)`

SetHasMediaSpoiler sets HasMediaSpoiler field to given value.

### HasHasMediaSpoiler

`func (o *ExternalReplyInfo) HasHasMediaSpoiler() bool`

HasHasMediaSpoiler returns a boolean if a field has been set.

### GetContact

`func (o *ExternalReplyInfo) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ExternalReplyInfo) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ExternalReplyInfo) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ExternalReplyInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDice

`func (o *ExternalReplyInfo) GetDice() Dice`

GetDice returns the Dice field if non-nil, zero value otherwise.

### GetDiceOk

`func (o *ExternalReplyInfo) GetDiceOk() (*Dice, bool)`

GetDiceOk returns a tuple with the Dice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDice

`func (o *ExternalReplyInfo) SetDice(v Dice)`

SetDice sets Dice field to given value.

### HasDice

`func (o *ExternalReplyInfo) HasDice() bool`

HasDice returns a boolean if a field has been set.

### GetGame

`func (o *ExternalReplyInfo) GetGame() Game`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *ExternalReplyInfo) GetGameOk() (*Game, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *ExternalReplyInfo) SetGame(v Game)`

SetGame sets Game field to given value.

### HasGame

`func (o *ExternalReplyInfo) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetGiveaway

`func (o *ExternalReplyInfo) GetGiveaway() Giveaway`

GetGiveaway returns the Giveaway field if non-nil, zero value otherwise.

### GetGiveawayOk

`func (o *ExternalReplyInfo) GetGiveawayOk() (*Giveaway, bool)`

GetGiveawayOk returns a tuple with the Giveaway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveaway

`func (o *ExternalReplyInfo) SetGiveaway(v Giveaway)`

SetGiveaway sets Giveaway field to given value.

### HasGiveaway

`func (o *ExternalReplyInfo) HasGiveaway() bool`

HasGiveaway returns a boolean if a field has been set.

### GetGiveawayWinners

`func (o *ExternalReplyInfo) GetGiveawayWinners() GiveawayWinners`

GetGiveawayWinners returns the GiveawayWinners field if non-nil, zero value otherwise.

### GetGiveawayWinnersOk

`func (o *ExternalReplyInfo) GetGiveawayWinnersOk() (*GiveawayWinners, bool)`

GetGiveawayWinnersOk returns a tuple with the GiveawayWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayWinners

`func (o *ExternalReplyInfo) SetGiveawayWinners(v GiveawayWinners)`

SetGiveawayWinners sets GiveawayWinners field to given value.

### HasGiveawayWinners

`func (o *ExternalReplyInfo) HasGiveawayWinners() bool`

HasGiveawayWinners returns a boolean if a field has been set.

### GetInvoice

`func (o *ExternalReplyInfo) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ExternalReplyInfo) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ExternalReplyInfo) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ExternalReplyInfo) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetLocation

`func (o *ExternalReplyInfo) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalReplyInfo) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalReplyInfo) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalReplyInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPoll

`func (o *ExternalReplyInfo) GetPoll() Poll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *ExternalReplyInfo) GetPollOk() (*Poll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *ExternalReplyInfo) SetPoll(v Poll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *ExternalReplyInfo) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetVenue

`func (o *ExternalReplyInfo) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *ExternalReplyInfo) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *ExternalReplyInfo) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *ExternalReplyInfo) HasVenue() bool`

HasVenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


