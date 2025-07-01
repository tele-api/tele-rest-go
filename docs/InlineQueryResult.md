# InlineQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *voice* | [default to "voice"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**AudioFileId** | **string** | A valid file identifier for the audio file | 
**Caption** | Pointer to **string** | *Optional*. Caption, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the voice message caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | [**InputMessageContent**](InputMessageContent.md) |  | 
**Title** | **string** | Recording title | 
**DocumentFileId** | **string** | A valid file identifier for the file | 
**Description** | Pointer to **string** | *Optional*. Short description of the result | [optional] 
**GifFileId** | **string** | A valid file identifier for the GIF file | 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**Mpeg4FileId** | **string** | A valid file identifier for the MPEG4 file | 
**PhotoFileId** | **string** | A valid file identifier of the photo | 
**StickerFileId** | **string** | A valid file identifier of the sticker | 
**VideoFileId** | **string** | A valid file identifier for the video file | 
**VoiceFileId** | **string** | A valid file identifier for the voice message | 
**Url** | Pointer to **string** | *Optional*. URL of the result | [optional] 
**ThumbnailUrl** | **string** | URL of the thumbnail (JPEG only) for the video | 
**ThumbnailWidth** | Pointer to **int32** | *Optional*. Thumbnail width | [optional] 
**ThumbnailHeight** | Pointer to **int32** | *Optional*. Thumbnail height | [optional] 
**AudioUrl** | **string** | A valid URL for the audio file | 
**Performer** | Pointer to **string** | *Optional*. Performer | [optional] 
**AudioDuration** | Pointer to **int32** | *Optional*. Audio duration in seconds | [optional] 
**PhoneNumber** | **string** | Contact&#39;s phone number | 
**FirstName** | **string** | Contact&#39;s first name | 
**LastName** | Pointer to **string** | *Optional*. Contact&#39;s last name | [optional] 
**Vcard** | Pointer to **string** | *Optional*. Additional data about the contact in the form of a [vCard](https://en.wikipedia.org/wiki/VCard), 0-2048 bytes | [optional] 
**GameShortName** | **string** | Short name of the game | 
**DocumentUrl** | **string** | A valid URL for the file | 
**MimeType** | **string** | MIME type of the content of the video URL, “text/html” or “video/mp4” | 
**GifUrl** | **string** | A valid URL for the GIF file | 
**GifWidth** | Pointer to **int32** | *Optional*. Width of the GIF | [optional] 
**GifHeight** | Pointer to **int32** | *Optional*. Height of the GIF | [optional] 
**GifDuration** | Pointer to **int32** | *Optional*. Duration of the GIF in seconds | [optional] 
**ThumbnailMimeType** | Pointer to **string** | *Optional*. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg” | [optional] [default to "image/jpeg"]
**Latitude** | **float32** | Latitude of the venue location in degrees | 
**Longitude** | **float32** | Longitude of the venue location in degrees | 
**HorizontalAccuracy** | Pointer to **float32** | *Optional*. The radius of uncertainty for the location, measured in meters; 0-1500 | [optional] 
**LivePeriod** | Pointer to **int32** | *Optional*. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely. | [optional] 
**Heading** | Pointer to **int32** | *Optional*. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified. | [optional] 
**ProximityAlertRadius** | Pointer to **int32** | *Optional*. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified. | [optional] 
**Mpeg4Url** | **string** | A valid URL for the MPEG4 file | 
**Mpeg4Width** | Pointer to **int32** | *Optional*. Video width | [optional] 
**Mpeg4Height** | Pointer to **int32** | *Optional*. Video height | [optional] 
**Mpeg4Duration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**PhotoUrl** | **string** | A valid URL of the photo. Photo must be in **JPEG** format. Photo size must not exceed 5MB | 
**PhotoWidth** | Pointer to **int32** | *Optional*. Width of the photo | [optional] 
**PhotoHeight** | Pointer to **int32** | *Optional*. Height of the photo | [optional] 
**Address** | **string** | Address of the venue | 
**FoursquareId** | Pointer to **string** | *Optional*. Foursquare identifier of the venue if known | [optional] 
**FoursquareType** | Pointer to **string** | *Optional*. Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.) | [optional] 
**GooglePlaceId** | Pointer to **string** | *Optional*. Google Places identifier of the venue | [optional] 
**GooglePlaceType** | Pointer to **string** | *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).) | [optional] 
**VideoUrl** | **string** | A valid URL for the embedded video player or video file | 
**VideoWidth** | Pointer to **int32** | *Optional*. Video width | [optional] 
**VideoHeight** | Pointer to **int32** | *Optional*. Video height | [optional] 
**VideoDuration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**VoiceUrl** | **string** | A valid URL for the voice recording | 
**VoiceDuration** | Pointer to **int32** | *Optional*. Recording duration in seconds | [optional] 

## Methods

### NewInlineQueryResult

`func NewInlineQueryResult(type_ string, id string, audioFileId string, inputMessageContent InputMessageContent, title string, documentFileId string, gifFileId string, mpeg4FileId string, photoFileId string, stickerFileId string, videoFileId string, voiceFileId string, thumbnailUrl string, audioUrl string, phoneNumber string, firstName string, gameShortName string, documentUrl string, mimeType string, gifUrl string, latitude float32, longitude float32, mpeg4Url string, photoUrl string, address string, videoUrl string, voiceUrl string, ) *InlineQueryResult`

NewInlineQueryResult instantiates a new InlineQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultWithDefaults

`func NewInlineQueryResultWithDefaults() *InlineQueryResult`

NewInlineQueryResultWithDefaults instantiates a new InlineQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResult) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResult) SetId(v string)`

SetId sets Id field to given value.


### GetAudioFileId

`func (o *InlineQueryResult) GetAudioFileId() string`

GetAudioFileId returns the AudioFileId field if non-nil, zero value otherwise.

### GetAudioFileIdOk

`func (o *InlineQueryResult) GetAudioFileIdOk() (*string, bool)`

GetAudioFileIdOk returns a tuple with the AudioFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFileId

`func (o *InlineQueryResult) SetAudioFileId(v string)`

SetAudioFileId sets AudioFileId field to given value.


### GetCaption

`func (o *InlineQueryResult) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResult) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResult) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResult) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResult) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResult) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResult) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResult) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResult) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResult) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResult) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResult) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResult) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResult) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResult) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResult) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResult) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResult) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResult) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.


### GetTitle

`func (o *InlineQueryResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDocumentFileId

`func (o *InlineQueryResult) GetDocumentFileId() string`

GetDocumentFileId returns the DocumentFileId field if non-nil, zero value otherwise.

### GetDocumentFileIdOk

`func (o *InlineQueryResult) GetDocumentFileIdOk() (*string, bool)`

GetDocumentFileIdOk returns a tuple with the DocumentFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFileId

`func (o *InlineQueryResult) SetDocumentFileId(v string)`

SetDocumentFileId sets DocumentFileId field to given value.


### GetDescription

`func (o *InlineQueryResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineQueryResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineQueryResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineQueryResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGifFileId

`func (o *InlineQueryResult) GetGifFileId() string`

GetGifFileId returns the GifFileId field if non-nil, zero value otherwise.

### GetGifFileIdOk

`func (o *InlineQueryResult) GetGifFileIdOk() (*string, bool)`

GetGifFileIdOk returns a tuple with the GifFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifFileId

`func (o *InlineQueryResult) SetGifFileId(v string)`

SetGifFileId sets GifFileId field to given value.


### GetShowCaptionAboveMedia

`func (o *InlineQueryResult) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResult) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResult) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResult) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetMpeg4FileId

`func (o *InlineQueryResult) GetMpeg4FileId() string`

GetMpeg4FileId returns the Mpeg4FileId field if non-nil, zero value otherwise.

### GetMpeg4FileIdOk

`func (o *InlineQueryResult) GetMpeg4FileIdOk() (*string, bool)`

GetMpeg4FileIdOk returns a tuple with the Mpeg4FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4FileId

`func (o *InlineQueryResult) SetMpeg4FileId(v string)`

SetMpeg4FileId sets Mpeg4FileId field to given value.


### GetPhotoFileId

`func (o *InlineQueryResult) GetPhotoFileId() string`

GetPhotoFileId returns the PhotoFileId field if non-nil, zero value otherwise.

### GetPhotoFileIdOk

`func (o *InlineQueryResult) GetPhotoFileIdOk() (*string, bool)`

GetPhotoFileIdOk returns a tuple with the PhotoFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoFileId

`func (o *InlineQueryResult) SetPhotoFileId(v string)`

SetPhotoFileId sets PhotoFileId field to given value.


### GetStickerFileId

`func (o *InlineQueryResult) GetStickerFileId() string`

GetStickerFileId returns the StickerFileId field if non-nil, zero value otherwise.

### GetStickerFileIdOk

`func (o *InlineQueryResult) GetStickerFileIdOk() (*string, bool)`

GetStickerFileIdOk returns a tuple with the StickerFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerFileId

`func (o *InlineQueryResult) SetStickerFileId(v string)`

SetStickerFileId sets StickerFileId field to given value.


### GetVideoFileId

`func (o *InlineQueryResult) GetVideoFileId() string`

GetVideoFileId returns the VideoFileId field if non-nil, zero value otherwise.

### GetVideoFileIdOk

`func (o *InlineQueryResult) GetVideoFileIdOk() (*string, bool)`

GetVideoFileIdOk returns a tuple with the VideoFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFileId

`func (o *InlineQueryResult) SetVideoFileId(v string)`

SetVideoFileId sets VideoFileId field to given value.


### GetVoiceFileId

`func (o *InlineQueryResult) GetVoiceFileId() string`

GetVoiceFileId returns the VoiceFileId field if non-nil, zero value otherwise.

### GetVoiceFileIdOk

`func (o *InlineQueryResult) GetVoiceFileIdOk() (*string, bool)`

GetVoiceFileIdOk returns a tuple with the VoiceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFileId

`func (o *InlineQueryResult) SetVoiceFileId(v string)`

SetVoiceFileId sets VoiceFileId field to given value.


### GetUrl

`func (o *InlineQueryResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineQueryResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineQueryResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineQueryResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *InlineQueryResult) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResult) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResult) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetThumbnailWidth

`func (o *InlineQueryResult) GetThumbnailWidth() int32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *InlineQueryResult) GetThumbnailWidthOk() (*int32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *InlineQueryResult) SetThumbnailWidth(v int32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *InlineQueryResult) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### GetThumbnailHeight

`func (o *InlineQueryResult) GetThumbnailHeight() int32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *InlineQueryResult) GetThumbnailHeightOk() (*int32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *InlineQueryResult) SetThumbnailHeight(v int32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *InlineQueryResult) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### GetAudioUrl

`func (o *InlineQueryResult) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *InlineQueryResult) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *InlineQueryResult) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.


### GetPerformer

`func (o *InlineQueryResult) GetPerformer() string`

GetPerformer returns the Performer field if non-nil, zero value otherwise.

### GetPerformerOk

`func (o *InlineQueryResult) GetPerformerOk() (*string, bool)`

GetPerformerOk returns a tuple with the Performer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformer

`func (o *InlineQueryResult) SetPerformer(v string)`

SetPerformer sets Performer field to given value.

### HasPerformer

`func (o *InlineQueryResult) HasPerformer() bool`

HasPerformer returns a boolean if a field has been set.

### GetAudioDuration

`func (o *InlineQueryResult) GetAudioDuration() int32`

GetAudioDuration returns the AudioDuration field if non-nil, zero value otherwise.

### GetAudioDurationOk

`func (o *InlineQueryResult) GetAudioDurationOk() (*int32, bool)`

GetAudioDurationOk returns a tuple with the AudioDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDuration

`func (o *InlineQueryResult) SetAudioDuration(v int32)`

SetAudioDuration sets AudioDuration field to given value.

### HasAudioDuration

`func (o *InlineQueryResult) HasAudioDuration() bool`

HasAudioDuration returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *InlineQueryResult) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InlineQueryResult) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InlineQueryResult) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetFirstName

`func (o *InlineQueryResult) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineQueryResult) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineQueryResult) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *InlineQueryResult) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineQueryResult) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineQueryResult) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineQueryResult) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetVcard

`func (o *InlineQueryResult) GetVcard() string`

GetVcard returns the Vcard field if non-nil, zero value otherwise.

### GetVcardOk

`func (o *InlineQueryResult) GetVcardOk() (*string, bool)`

GetVcardOk returns a tuple with the Vcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcard

`func (o *InlineQueryResult) SetVcard(v string)`

SetVcard sets Vcard field to given value.

### HasVcard

`func (o *InlineQueryResult) HasVcard() bool`

HasVcard returns a boolean if a field has been set.

### GetGameShortName

`func (o *InlineQueryResult) GetGameShortName() string`

GetGameShortName returns the GameShortName field if non-nil, zero value otherwise.

### GetGameShortNameOk

`func (o *InlineQueryResult) GetGameShortNameOk() (*string, bool)`

GetGameShortNameOk returns a tuple with the GameShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameShortName

`func (o *InlineQueryResult) SetGameShortName(v string)`

SetGameShortName sets GameShortName field to given value.


### GetDocumentUrl

`func (o *InlineQueryResult) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *InlineQueryResult) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *InlineQueryResult) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetMimeType

`func (o *InlineQueryResult) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *InlineQueryResult) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *InlineQueryResult) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetGifUrl

`func (o *InlineQueryResult) GetGifUrl() string`

GetGifUrl returns the GifUrl field if non-nil, zero value otherwise.

### GetGifUrlOk

`func (o *InlineQueryResult) GetGifUrlOk() (*string, bool)`

GetGifUrlOk returns a tuple with the GifUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifUrl

`func (o *InlineQueryResult) SetGifUrl(v string)`

SetGifUrl sets GifUrl field to given value.


### GetGifWidth

`func (o *InlineQueryResult) GetGifWidth() int32`

GetGifWidth returns the GifWidth field if non-nil, zero value otherwise.

### GetGifWidthOk

`func (o *InlineQueryResult) GetGifWidthOk() (*int32, bool)`

GetGifWidthOk returns a tuple with the GifWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifWidth

`func (o *InlineQueryResult) SetGifWidth(v int32)`

SetGifWidth sets GifWidth field to given value.

### HasGifWidth

`func (o *InlineQueryResult) HasGifWidth() bool`

HasGifWidth returns a boolean if a field has been set.

### GetGifHeight

`func (o *InlineQueryResult) GetGifHeight() int32`

GetGifHeight returns the GifHeight field if non-nil, zero value otherwise.

### GetGifHeightOk

`func (o *InlineQueryResult) GetGifHeightOk() (*int32, bool)`

GetGifHeightOk returns a tuple with the GifHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifHeight

`func (o *InlineQueryResult) SetGifHeight(v int32)`

SetGifHeight sets GifHeight field to given value.

### HasGifHeight

`func (o *InlineQueryResult) HasGifHeight() bool`

HasGifHeight returns a boolean if a field has been set.

### GetGifDuration

`func (o *InlineQueryResult) GetGifDuration() int32`

GetGifDuration returns the GifDuration field if non-nil, zero value otherwise.

### GetGifDurationOk

`func (o *InlineQueryResult) GetGifDurationOk() (*int32, bool)`

GetGifDurationOk returns a tuple with the GifDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifDuration

`func (o *InlineQueryResult) SetGifDuration(v int32)`

SetGifDuration sets GifDuration field to given value.

### HasGifDuration

`func (o *InlineQueryResult) HasGifDuration() bool`

HasGifDuration returns a boolean if a field has been set.

### GetThumbnailMimeType

`func (o *InlineQueryResult) GetThumbnailMimeType() string`

GetThumbnailMimeType returns the ThumbnailMimeType field if non-nil, zero value otherwise.

### GetThumbnailMimeTypeOk

`func (o *InlineQueryResult) GetThumbnailMimeTypeOk() (*string, bool)`

GetThumbnailMimeTypeOk returns a tuple with the ThumbnailMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailMimeType

`func (o *InlineQueryResult) SetThumbnailMimeType(v string)`

SetThumbnailMimeType sets ThumbnailMimeType field to given value.

### HasThumbnailMimeType

`func (o *InlineQueryResult) HasThumbnailMimeType() bool`

HasThumbnailMimeType returns a boolean if a field has been set.

### GetLatitude

`func (o *InlineQueryResult) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InlineQueryResult) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InlineQueryResult) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *InlineQueryResult) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InlineQueryResult) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InlineQueryResult) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetHorizontalAccuracy

`func (o *InlineQueryResult) GetHorizontalAccuracy() float32`

GetHorizontalAccuracy returns the HorizontalAccuracy field if non-nil, zero value otherwise.

### GetHorizontalAccuracyOk

`func (o *InlineQueryResult) GetHorizontalAccuracyOk() (*float32, bool)`

GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAccuracy

`func (o *InlineQueryResult) SetHorizontalAccuracy(v float32)`

SetHorizontalAccuracy sets HorizontalAccuracy field to given value.

### HasHorizontalAccuracy

`func (o *InlineQueryResult) HasHorizontalAccuracy() bool`

HasHorizontalAccuracy returns a boolean if a field has been set.

### GetLivePeriod

`func (o *InlineQueryResult) GetLivePeriod() int32`

GetLivePeriod returns the LivePeriod field if non-nil, zero value otherwise.

### GetLivePeriodOk

`func (o *InlineQueryResult) GetLivePeriodOk() (*int32, bool)`

GetLivePeriodOk returns a tuple with the LivePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePeriod

`func (o *InlineQueryResult) SetLivePeriod(v int32)`

SetLivePeriod sets LivePeriod field to given value.

### HasLivePeriod

`func (o *InlineQueryResult) HasLivePeriod() bool`

HasLivePeriod returns a boolean if a field has been set.

### GetHeading

`func (o *InlineQueryResult) GetHeading() int32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *InlineQueryResult) GetHeadingOk() (*int32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *InlineQueryResult) SetHeading(v int32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *InlineQueryResult) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetProximityAlertRadius

`func (o *InlineQueryResult) GetProximityAlertRadius() int32`

GetProximityAlertRadius returns the ProximityAlertRadius field if non-nil, zero value otherwise.

### GetProximityAlertRadiusOk

`func (o *InlineQueryResult) GetProximityAlertRadiusOk() (*int32, bool)`

GetProximityAlertRadiusOk returns a tuple with the ProximityAlertRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertRadius

`func (o *InlineQueryResult) SetProximityAlertRadius(v int32)`

SetProximityAlertRadius sets ProximityAlertRadius field to given value.

### HasProximityAlertRadius

`func (o *InlineQueryResult) HasProximityAlertRadius() bool`

HasProximityAlertRadius returns a boolean if a field has been set.

### GetMpeg4Url

`func (o *InlineQueryResult) GetMpeg4Url() string`

GetMpeg4Url returns the Mpeg4Url field if non-nil, zero value otherwise.

### GetMpeg4UrlOk

`func (o *InlineQueryResult) GetMpeg4UrlOk() (*string, bool)`

GetMpeg4UrlOk returns a tuple with the Mpeg4Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Url

`func (o *InlineQueryResult) SetMpeg4Url(v string)`

SetMpeg4Url sets Mpeg4Url field to given value.


### GetMpeg4Width

`func (o *InlineQueryResult) GetMpeg4Width() int32`

GetMpeg4Width returns the Mpeg4Width field if non-nil, zero value otherwise.

### GetMpeg4WidthOk

`func (o *InlineQueryResult) GetMpeg4WidthOk() (*int32, bool)`

GetMpeg4WidthOk returns a tuple with the Mpeg4Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Width

`func (o *InlineQueryResult) SetMpeg4Width(v int32)`

SetMpeg4Width sets Mpeg4Width field to given value.

### HasMpeg4Width

`func (o *InlineQueryResult) HasMpeg4Width() bool`

HasMpeg4Width returns a boolean if a field has been set.

### GetMpeg4Height

`func (o *InlineQueryResult) GetMpeg4Height() int32`

GetMpeg4Height returns the Mpeg4Height field if non-nil, zero value otherwise.

### GetMpeg4HeightOk

`func (o *InlineQueryResult) GetMpeg4HeightOk() (*int32, bool)`

GetMpeg4HeightOk returns a tuple with the Mpeg4Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Height

`func (o *InlineQueryResult) SetMpeg4Height(v int32)`

SetMpeg4Height sets Mpeg4Height field to given value.

### HasMpeg4Height

`func (o *InlineQueryResult) HasMpeg4Height() bool`

HasMpeg4Height returns a boolean if a field has been set.

### GetMpeg4Duration

`func (o *InlineQueryResult) GetMpeg4Duration() int32`

GetMpeg4Duration returns the Mpeg4Duration field if non-nil, zero value otherwise.

### GetMpeg4DurationOk

`func (o *InlineQueryResult) GetMpeg4DurationOk() (*int32, bool)`

GetMpeg4DurationOk returns a tuple with the Mpeg4Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Duration

`func (o *InlineQueryResult) SetMpeg4Duration(v int32)`

SetMpeg4Duration sets Mpeg4Duration field to given value.

### HasMpeg4Duration

`func (o *InlineQueryResult) HasMpeg4Duration() bool`

HasMpeg4Duration returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *InlineQueryResult) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *InlineQueryResult) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *InlineQueryResult) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.


### GetPhotoWidth

`func (o *InlineQueryResult) GetPhotoWidth() int32`

GetPhotoWidth returns the PhotoWidth field if non-nil, zero value otherwise.

### GetPhotoWidthOk

`func (o *InlineQueryResult) GetPhotoWidthOk() (*int32, bool)`

GetPhotoWidthOk returns a tuple with the PhotoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoWidth

`func (o *InlineQueryResult) SetPhotoWidth(v int32)`

SetPhotoWidth sets PhotoWidth field to given value.

### HasPhotoWidth

`func (o *InlineQueryResult) HasPhotoWidth() bool`

HasPhotoWidth returns a boolean if a field has been set.

### GetPhotoHeight

`func (o *InlineQueryResult) GetPhotoHeight() int32`

GetPhotoHeight returns the PhotoHeight field if non-nil, zero value otherwise.

### GetPhotoHeightOk

`func (o *InlineQueryResult) GetPhotoHeightOk() (*int32, bool)`

GetPhotoHeightOk returns a tuple with the PhotoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoHeight

`func (o *InlineQueryResult) SetPhotoHeight(v int32)`

SetPhotoHeight sets PhotoHeight field to given value.

### HasPhotoHeight

`func (o *InlineQueryResult) HasPhotoHeight() bool`

HasPhotoHeight returns a boolean if a field has been set.

### GetAddress

`func (o *InlineQueryResult) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineQueryResult) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineQueryResult) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFoursquareId

`func (o *InlineQueryResult) GetFoursquareId() string`

GetFoursquareId returns the FoursquareId field if non-nil, zero value otherwise.

### GetFoursquareIdOk

`func (o *InlineQueryResult) GetFoursquareIdOk() (*string, bool)`

GetFoursquareIdOk returns a tuple with the FoursquareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareId

`func (o *InlineQueryResult) SetFoursquareId(v string)`

SetFoursquareId sets FoursquareId field to given value.

### HasFoursquareId

`func (o *InlineQueryResult) HasFoursquareId() bool`

HasFoursquareId returns a boolean if a field has been set.

### GetFoursquareType

`func (o *InlineQueryResult) GetFoursquareType() string`

GetFoursquareType returns the FoursquareType field if non-nil, zero value otherwise.

### GetFoursquareTypeOk

`func (o *InlineQueryResult) GetFoursquareTypeOk() (*string, bool)`

GetFoursquareTypeOk returns a tuple with the FoursquareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareType

`func (o *InlineQueryResult) SetFoursquareType(v string)`

SetFoursquareType sets FoursquareType field to given value.

### HasFoursquareType

`func (o *InlineQueryResult) HasFoursquareType() bool`

HasFoursquareType returns a boolean if a field has been set.

### GetGooglePlaceId

`func (o *InlineQueryResult) GetGooglePlaceId() string`

GetGooglePlaceId returns the GooglePlaceId field if non-nil, zero value otherwise.

### GetGooglePlaceIdOk

`func (o *InlineQueryResult) GetGooglePlaceIdOk() (*string, bool)`

GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceId

`func (o *InlineQueryResult) SetGooglePlaceId(v string)`

SetGooglePlaceId sets GooglePlaceId field to given value.

### HasGooglePlaceId

`func (o *InlineQueryResult) HasGooglePlaceId() bool`

HasGooglePlaceId returns a boolean if a field has been set.

### GetGooglePlaceType

`func (o *InlineQueryResult) GetGooglePlaceType() string`

GetGooglePlaceType returns the GooglePlaceType field if non-nil, zero value otherwise.

### GetGooglePlaceTypeOk

`func (o *InlineQueryResult) GetGooglePlaceTypeOk() (*string, bool)`

GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceType

`func (o *InlineQueryResult) SetGooglePlaceType(v string)`

SetGooglePlaceType sets GooglePlaceType field to given value.

### HasGooglePlaceType

`func (o *InlineQueryResult) HasGooglePlaceType() bool`

HasGooglePlaceType returns a boolean if a field has been set.

### GetVideoUrl

`func (o *InlineQueryResult) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *InlineQueryResult) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *InlineQueryResult) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.


### GetVideoWidth

`func (o *InlineQueryResult) GetVideoWidth() int32`

GetVideoWidth returns the VideoWidth field if non-nil, zero value otherwise.

### GetVideoWidthOk

`func (o *InlineQueryResult) GetVideoWidthOk() (*int32, bool)`

GetVideoWidthOk returns a tuple with the VideoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoWidth

`func (o *InlineQueryResult) SetVideoWidth(v int32)`

SetVideoWidth sets VideoWidth field to given value.

### HasVideoWidth

`func (o *InlineQueryResult) HasVideoWidth() bool`

HasVideoWidth returns a boolean if a field has been set.

### GetVideoHeight

`func (o *InlineQueryResult) GetVideoHeight() int32`

GetVideoHeight returns the VideoHeight field if non-nil, zero value otherwise.

### GetVideoHeightOk

`func (o *InlineQueryResult) GetVideoHeightOk() (*int32, bool)`

GetVideoHeightOk returns a tuple with the VideoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoHeight

`func (o *InlineQueryResult) SetVideoHeight(v int32)`

SetVideoHeight sets VideoHeight field to given value.

### HasVideoHeight

`func (o *InlineQueryResult) HasVideoHeight() bool`

HasVideoHeight returns a boolean if a field has been set.

### GetVideoDuration

`func (o *InlineQueryResult) GetVideoDuration() int32`

GetVideoDuration returns the VideoDuration field if non-nil, zero value otherwise.

### GetVideoDurationOk

`func (o *InlineQueryResult) GetVideoDurationOk() (*int32, bool)`

GetVideoDurationOk returns a tuple with the VideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDuration

`func (o *InlineQueryResult) SetVideoDuration(v int32)`

SetVideoDuration sets VideoDuration field to given value.

### HasVideoDuration

`func (o *InlineQueryResult) HasVideoDuration() bool`

HasVideoDuration returns a boolean if a field has been set.

### GetVoiceUrl

`func (o *InlineQueryResult) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *InlineQueryResult) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *InlineQueryResult) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.


### GetVoiceDuration

`func (o *InlineQueryResult) GetVoiceDuration() int32`

GetVoiceDuration returns the VoiceDuration field if non-nil, zero value otherwise.

### GetVoiceDurationOk

`func (o *InlineQueryResult) GetVoiceDurationOk() (*int32, bool)`

GetVoiceDurationOk returns a tuple with the VoiceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDuration

`func (o *InlineQueryResult) SetVoiceDuration(v int32)`

SetVoiceDuration sets VoiceDuration field to given value.

### HasVoiceDuration

`func (o *InlineQueryResult) HasVoiceDuration() bool`

HasVoiceDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


