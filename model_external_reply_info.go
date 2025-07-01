/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Telegram Bot API service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package tele_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExternalReplyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalReplyInfo{}

// ExternalReplyInfo This object contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Origin MessageOrigin `json:"origin"`
	Chat *Chat `json:"chat,omitempty"`
	// *Optional*. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	MessageId *int32 `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation *Animation `json:"animation,omitempty"`
	Audio *Audio `json:"audio,omitempty"`
	Document *Document `json:"document,omitempty"`
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	// *Optional*. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
	Story *Story `json:"story,omitempty"`
	Video *Video `json:"video,omitempty"`
	VideoNote *VideoNote `json:"video_note,omitempty"`
	Voice *Voice `json:"voice,omitempty"`
	// *Optional*. *True*, if the message media is covered by a spoiler animation
	HasMediaSpoiler *bool `json:"has_media_spoiler,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Dice *Dice `json:"dice,omitempty"`
	Game *Game `json:"game,omitempty"`
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	Invoice *Invoice `json:"invoice,omitempty"`
	Location *Location `json:"location,omitempty"`
	Poll *Poll `json:"poll,omitempty"`
	Venue *Venue `json:"venue,omitempty"`
}

type _ExternalReplyInfo ExternalReplyInfo

// NewExternalReplyInfo instantiates a new ExternalReplyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalReplyInfo(origin MessageOrigin) *ExternalReplyInfo {
	this := ExternalReplyInfo{}
	this.Origin = origin
	var hasMediaSpoiler bool = true
	this.HasMediaSpoiler = &hasMediaSpoiler
	return &this
}

// NewExternalReplyInfoWithDefaults instantiates a new ExternalReplyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalReplyInfoWithDefaults() *ExternalReplyInfo {
	this := ExternalReplyInfo{}
	var hasMediaSpoiler bool = true
	this.HasMediaSpoiler = &hasMediaSpoiler
	return &this
}

// GetOrigin returns the Origin field value
func (o *ExternalReplyInfo) GetOrigin() MessageOrigin {
	if o == nil {
		var ret MessageOrigin
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetOriginOk() (*MessageOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *ExternalReplyInfo) SetOrigin(v MessageOrigin) {
	o.Origin = v
}

// GetChat returns the Chat field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetChat() Chat {
	if o == nil || IsNil(o.Chat) {
		var ret Chat
		return ret
	}
	return *o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasChat() bool {
	if o != nil && !IsNil(o.Chat) {
		return true
	}

	return false
}

// SetChat gets a reference to the given Chat and assigns it to the Chat field.
func (o *ExternalReplyInfo) SetChat(v Chat) {
	o.Chat = &v
}


// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *ExternalReplyInfo) SetMessageId(v int32) {
	o.MessageId = &v
}


// GetLinkPreviewOptions returns the LinkPreviewOptions field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetLinkPreviewOptions() LinkPreviewOptions {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		var ret LinkPreviewOptions
		return ret
	}
	return *o.LinkPreviewOptions
}

// GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool) {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		return nil, false
	}
	return o.LinkPreviewOptions, true
}

// HasLinkPreviewOptions returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasLinkPreviewOptions() bool {
	if o != nil && !IsNil(o.LinkPreviewOptions) {
		return true
	}

	return false
}

// SetLinkPreviewOptions gets a reference to the given LinkPreviewOptions and assigns it to the LinkPreviewOptions field.
func (o *ExternalReplyInfo) SetLinkPreviewOptions(v LinkPreviewOptions) {
	o.LinkPreviewOptions = &v
}


// GetAnimation returns the Animation field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetAnimation() Animation {
	if o == nil || IsNil(o.Animation) {
		var ret Animation
		return ret
	}
	return *o.Animation
}

// GetAnimationOk returns a tuple with the Animation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetAnimationOk() (*Animation, bool) {
	if o == nil || IsNil(o.Animation) {
		return nil, false
	}
	return o.Animation, true
}

// HasAnimation returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasAnimation() bool {
	if o != nil && !IsNil(o.Animation) {
		return true
	}

	return false
}

// SetAnimation gets a reference to the given Animation and assigns it to the Animation field.
func (o *ExternalReplyInfo) SetAnimation(v Animation) {
	o.Animation = &v
}


// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetAudio() Audio {
	if o == nil || IsNil(o.Audio) {
		var ret Audio
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetAudioOk() (*Audio, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasAudio() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given Audio and assigns it to the Audio field.
func (o *ExternalReplyInfo) SetAudio(v Audio) {
	o.Audio = &v
}


// GetDocument returns the Document field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetDocument() Document {
	if o == nil || IsNil(o.Document) {
		var ret Document
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Document and assigns it to the Document field.
func (o *ExternalReplyInfo) SetDocument(v Document) {
	o.Document = &v
}


// GetPaidMedia returns the PaidMedia field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetPaidMedia() PaidMediaInfo {
	if o == nil || IsNil(o.PaidMedia) {
		var ret PaidMediaInfo
		return ret
	}
	return *o.PaidMedia
}

// GetPaidMediaOk returns a tuple with the PaidMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetPaidMediaOk() (*PaidMediaInfo, bool) {
	if o == nil || IsNil(o.PaidMedia) {
		return nil, false
	}
	return o.PaidMedia, true
}

// HasPaidMedia returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasPaidMedia() bool {
	if o != nil && !IsNil(o.PaidMedia) {
		return true
	}

	return false
}

// SetPaidMedia gets a reference to the given PaidMediaInfo and assigns it to the PaidMedia field.
func (o *ExternalReplyInfo) SetPaidMedia(v PaidMediaInfo) {
	o.PaidMedia = &v
}


// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetPhoto() []PhotoSize {
	if o == nil || IsNil(o.Photo) {
		var ret []PhotoSize
		return ret
	}
	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetPhotoOk() ([]PhotoSize, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given []PhotoSize and assigns it to the Photo field.
func (o *ExternalReplyInfo) SetPhoto(v []PhotoSize) {
	o.Photo = v
}


// GetSticker returns the Sticker field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetSticker() Sticker {
	if o == nil || IsNil(o.Sticker) {
		var ret Sticker
		return ret
	}
	return *o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetStickerOk() (*Sticker, bool) {
	if o == nil || IsNil(o.Sticker) {
		return nil, false
	}
	return o.Sticker, true
}

// HasSticker returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasSticker() bool {
	if o != nil && !IsNil(o.Sticker) {
		return true
	}

	return false
}

// SetSticker gets a reference to the given Sticker and assigns it to the Sticker field.
func (o *ExternalReplyInfo) SetSticker(v Sticker) {
	o.Sticker = &v
}


// GetStory returns the Story field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetStory() Story {
	if o == nil || IsNil(o.Story) {
		var ret Story
		return ret
	}
	return *o.Story
}

// GetStoryOk returns a tuple with the Story field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetStoryOk() (*Story, bool) {
	if o == nil || IsNil(o.Story) {
		return nil, false
	}
	return o.Story, true
}

// HasStory returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasStory() bool {
	if o != nil && !IsNil(o.Story) {
		return true
	}

	return false
}

// SetStory gets a reference to the given Story and assigns it to the Story field.
func (o *ExternalReplyInfo) SetStory(v Story) {
	o.Story = &v
}


// GetVideo returns the Video field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetVideo() Video {
	if o == nil || IsNil(o.Video) {
		var ret Video
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetVideoOk() (*Video, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given Video and assigns it to the Video field.
func (o *ExternalReplyInfo) SetVideo(v Video) {
	o.Video = &v
}


// GetVideoNote returns the VideoNote field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetVideoNote() VideoNote {
	if o == nil || IsNil(o.VideoNote) {
		var ret VideoNote
		return ret
	}
	return *o.VideoNote
}

// GetVideoNoteOk returns a tuple with the VideoNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetVideoNoteOk() (*VideoNote, bool) {
	if o == nil || IsNil(o.VideoNote) {
		return nil, false
	}
	return o.VideoNote, true
}

// HasVideoNote returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasVideoNote() bool {
	if o != nil && !IsNil(o.VideoNote) {
		return true
	}

	return false
}

// SetVideoNote gets a reference to the given VideoNote and assigns it to the VideoNote field.
func (o *ExternalReplyInfo) SetVideoNote(v VideoNote) {
	o.VideoNote = &v
}


// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetVoice() Voice {
	if o == nil || IsNil(o.Voice) {
		var ret Voice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetVoiceOk() (*Voice, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given Voice and assigns it to the Voice field.
func (o *ExternalReplyInfo) SetVoice(v Voice) {
	o.Voice = &v
}


// GetHasMediaSpoiler returns the HasMediaSpoiler field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetHasMediaSpoiler() bool {
	if o == nil || IsNil(o.HasMediaSpoiler) {
		var ret bool
		return ret
	}
	return *o.HasMediaSpoiler
}

// GetHasMediaSpoilerOk returns a tuple with the HasMediaSpoiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetHasMediaSpoilerOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMediaSpoiler) {
		return nil, false
	}
	return o.HasMediaSpoiler, true
}

// HasHasMediaSpoiler returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasHasMediaSpoiler() bool {
	if o != nil && !IsNil(o.HasMediaSpoiler) {
		return true
	}

	return false
}

// SetHasMediaSpoiler gets a reference to the given bool and assigns it to the HasMediaSpoiler field.
func (o *ExternalReplyInfo) SetHasMediaSpoiler(v bool) {
	o.HasMediaSpoiler = &v
}


// GetContact returns the Contact field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *ExternalReplyInfo) SetContact(v Contact) {
	o.Contact = &v
}


// GetDice returns the Dice field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetDice() Dice {
	if o == nil || IsNil(o.Dice) {
		var ret Dice
		return ret
	}
	return *o.Dice
}

// GetDiceOk returns a tuple with the Dice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetDiceOk() (*Dice, bool) {
	if o == nil || IsNil(o.Dice) {
		return nil, false
	}
	return o.Dice, true
}

// HasDice returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasDice() bool {
	if o != nil && !IsNil(o.Dice) {
		return true
	}

	return false
}

// SetDice gets a reference to the given Dice and assigns it to the Dice field.
func (o *ExternalReplyInfo) SetDice(v Dice) {
	o.Dice = &v
}


// GetGame returns the Game field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetGame() Game {
	if o == nil || IsNil(o.Game) {
		var ret Game
		return ret
	}
	return *o.Game
}

// GetGameOk returns a tuple with the Game field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetGameOk() (*Game, bool) {
	if o == nil || IsNil(o.Game) {
		return nil, false
	}
	return o.Game, true
}

// HasGame returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasGame() bool {
	if o != nil && !IsNil(o.Game) {
		return true
	}

	return false
}

// SetGame gets a reference to the given Game and assigns it to the Game field.
func (o *ExternalReplyInfo) SetGame(v Game) {
	o.Game = &v
}


// GetGiveaway returns the Giveaway field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetGiveaway() Giveaway {
	if o == nil || IsNil(o.Giveaway) {
		var ret Giveaway
		return ret
	}
	return *o.Giveaway
}

// GetGiveawayOk returns a tuple with the Giveaway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetGiveawayOk() (*Giveaway, bool) {
	if o == nil || IsNil(o.Giveaway) {
		return nil, false
	}
	return o.Giveaway, true
}

// HasGiveaway returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasGiveaway() bool {
	if o != nil && !IsNil(o.Giveaway) {
		return true
	}

	return false
}

// SetGiveaway gets a reference to the given Giveaway and assigns it to the Giveaway field.
func (o *ExternalReplyInfo) SetGiveaway(v Giveaway) {
	o.Giveaway = &v
}


// GetGiveawayWinners returns the GiveawayWinners field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetGiveawayWinners() GiveawayWinners {
	if o == nil || IsNil(o.GiveawayWinners) {
		var ret GiveawayWinners
		return ret
	}
	return *o.GiveawayWinners
}

// GetGiveawayWinnersOk returns a tuple with the GiveawayWinners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetGiveawayWinnersOk() (*GiveawayWinners, bool) {
	if o == nil || IsNil(o.GiveawayWinners) {
		return nil, false
	}
	return o.GiveawayWinners, true
}

// HasGiveawayWinners returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasGiveawayWinners() bool {
	if o != nil && !IsNil(o.GiveawayWinners) {
		return true
	}

	return false
}

// SetGiveawayWinners gets a reference to the given GiveawayWinners and assigns it to the GiveawayWinners field.
func (o *ExternalReplyInfo) SetGiveawayWinners(v GiveawayWinners) {
	o.GiveawayWinners = &v
}


// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetInvoice() Invoice {
	if o == nil || IsNil(o.Invoice) {
		var ret Invoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetInvoiceOk() (*Invoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given Invoice and assigns it to the Invoice field.
func (o *ExternalReplyInfo) SetInvoice(v Invoice) {
	o.Invoice = &v
}


// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *ExternalReplyInfo) SetLocation(v Location) {
	o.Location = &v
}


// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetPoll() Poll {
	if o == nil || IsNil(o.Poll) {
		var ret Poll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetPollOk() (*Poll, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given Poll and assigns it to the Poll field.
func (o *ExternalReplyInfo) SetPoll(v Poll) {
	o.Poll = &v
}


// GetVenue returns the Venue field value if set, zero value otherwise.
func (o *ExternalReplyInfo) GetVenue() Venue {
	if o == nil || IsNil(o.Venue) {
		var ret Venue
		return ret
	}
	return *o.Venue
}

// GetVenueOk returns a tuple with the Venue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReplyInfo) GetVenueOk() (*Venue, bool) {
	if o == nil || IsNil(o.Venue) {
		return nil, false
	}
	return o.Venue, true
}

// HasVenue returns a boolean if a field has been set.
func (o *ExternalReplyInfo) HasVenue() bool {
	if o != nil && !IsNil(o.Venue) {
		return true
	}

	return false
}

// SetVenue gets a reference to the given Venue and assigns it to the Venue field.
func (o *ExternalReplyInfo) SetVenue(v Venue) {
	o.Venue = &v
}


func (o ExternalReplyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalReplyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin"] = o.Origin
	if !IsNil(o.Chat) {
		toSerialize["chat"] = o.Chat
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.LinkPreviewOptions) {
		toSerialize["link_preview_options"] = o.LinkPreviewOptions
	}
	if !IsNil(o.Animation) {
		toSerialize["animation"] = o.Animation
	}
	if !IsNil(o.Audio) {
		toSerialize["audio"] = o.Audio
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.PaidMedia) {
		toSerialize["paid_media"] = o.PaidMedia
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.Sticker) {
		toSerialize["sticker"] = o.Sticker
	}
	if !IsNil(o.Story) {
		toSerialize["story"] = o.Story
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.VideoNote) {
		toSerialize["video_note"] = o.VideoNote
	}
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	if !IsNil(o.HasMediaSpoiler) {
		toSerialize["has_media_spoiler"] = o.HasMediaSpoiler
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Dice) {
		toSerialize["dice"] = o.Dice
	}
	if !IsNil(o.Game) {
		toSerialize["game"] = o.Game
	}
	if !IsNil(o.Giveaway) {
		toSerialize["giveaway"] = o.Giveaway
	}
	if !IsNil(o.GiveawayWinners) {
		toSerialize["giveaway_winners"] = o.GiveawayWinners
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}
	if !IsNil(o.Venue) {
		toSerialize["venue"] = o.Venue
	}
	return toSerialize, nil
}

func (o *ExternalReplyInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"origin",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExternalReplyInfo := _ExternalReplyInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalReplyInfo)

	if err != nil {
		return err
	}

	*o = ExternalReplyInfo(varExternalReplyInfo)

	return err
}

type NullableExternalReplyInfo struct {
	value *ExternalReplyInfo
	isSet bool
}

func (v NullableExternalReplyInfo) Get() *ExternalReplyInfo {
	return v.value
}

func (v *NullableExternalReplyInfo) Set(val *ExternalReplyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalReplyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalReplyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalReplyInfo(val *ExternalReplyInfo) *NullableExternalReplyInfo {
	return &NullableExternalReplyInfo{value: val, isSet: true}
}

func (v NullableExternalReplyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalReplyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


