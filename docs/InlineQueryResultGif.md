# InlineQueryResultGif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *gif* | [default to "gif"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**GifUrl** | **string** | A valid URL for the GIF file | 
**GifWidth** | Pointer to **int32** | *Optional*. Width of the GIF | [optional] 
**GifHeight** | Pointer to **int32** | *Optional*. Height of the GIF | [optional] 
**GifDuration** | Pointer to **int32** | *Optional*. Duration of the GIF in seconds | [optional] 
**ThumbnailUrl** | **string** | URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result | 
**ThumbnailMimeType** | Pointer to **string** | *Optional*. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg” | [optional] [default to "image/jpeg"]
**Title** | Pointer to **string** | *Optional*. Title for the result | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the GIF file to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultGif

`func NewInlineQueryResultGif(type_ string, id string, gifUrl string, thumbnailUrl string, ) *InlineQueryResultGif`

NewInlineQueryResultGif instantiates a new InlineQueryResultGif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultGifWithDefaults

`func NewInlineQueryResultGifWithDefaults() *InlineQueryResultGif`

NewInlineQueryResultGifWithDefaults instantiates a new InlineQueryResultGif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultGif) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultGif) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultGif) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultGif) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultGif) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultGif) SetId(v string)`

SetId sets Id field to given value.


### GetGifUrl

`func (o *InlineQueryResultGif) GetGifUrl() string`

GetGifUrl returns the GifUrl field if non-nil, zero value otherwise.

### GetGifUrlOk

`func (o *InlineQueryResultGif) GetGifUrlOk() (*string, bool)`

GetGifUrlOk returns a tuple with the GifUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifUrl

`func (o *InlineQueryResultGif) SetGifUrl(v string)`

SetGifUrl sets GifUrl field to given value.


### GetGifWidth

`func (o *InlineQueryResultGif) GetGifWidth() int32`

GetGifWidth returns the GifWidth field if non-nil, zero value otherwise.

### GetGifWidthOk

`func (o *InlineQueryResultGif) GetGifWidthOk() (*int32, bool)`

GetGifWidthOk returns a tuple with the GifWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifWidth

`func (o *InlineQueryResultGif) SetGifWidth(v int32)`

SetGifWidth sets GifWidth field to given value.

### HasGifWidth

`func (o *InlineQueryResultGif) HasGifWidth() bool`

HasGifWidth returns a boolean if a field has been set.

### GetGifHeight

`func (o *InlineQueryResultGif) GetGifHeight() int32`

GetGifHeight returns the GifHeight field if non-nil, zero value otherwise.

### GetGifHeightOk

`func (o *InlineQueryResultGif) GetGifHeightOk() (*int32, bool)`

GetGifHeightOk returns a tuple with the GifHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifHeight

`func (o *InlineQueryResultGif) SetGifHeight(v int32)`

SetGifHeight sets GifHeight field to given value.

### HasGifHeight

`func (o *InlineQueryResultGif) HasGifHeight() bool`

HasGifHeight returns a boolean if a field has been set.

### GetGifDuration

`func (o *InlineQueryResultGif) GetGifDuration() int32`

GetGifDuration returns the GifDuration field if non-nil, zero value otherwise.

### GetGifDurationOk

`func (o *InlineQueryResultGif) GetGifDurationOk() (*int32, bool)`

GetGifDurationOk returns a tuple with the GifDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifDuration

`func (o *InlineQueryResultGif) SetGifDuration(v int32)`

SetGifDuration sets GifDuration field to given value.

### HasGifDuration

`func (o *InlineQueryResultGif) HasGifDuration() bool`

HasGifDuration returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *InlineQueryResultGif) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultGif) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultGif) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetThumbnailMimeType

`func (o *InlineQueryResultGif) GetThumbnailMimeType() string`

GetThumbnailMimeType returns the ThumbnailMimeType field if non-nil, zero value otherwise.

### GetThumbnailMimeTypeOk

`func (o *InlineQueryResultGif) GetThumbnailMimeTypeOk() (*string, bool)`

GetThumbnailMimeTypeOk returns a tuple with the ThumbnailMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailMimeType

`func (o *InlineQueryResultGif) SetThumbnailMimeType(v string)`

SetThumbnailMimeType sets ThumbnailMimeType field to given value.

### HasThumbnailMimeType

`func (o *InlineQueryResultGif) HasThumbnailMimeType() bool`

HasThumbnailMimeType returns a boolean if a field has been set.

### GetTitle

`func (o *InlineQueryResultGif) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultGif) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultGif) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineQueryResultGif) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCaption

`func (o *InlineQueryResultGif) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultGif) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultGif) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultGif) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultGif) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultGif) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultGif) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultGif) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultGif) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultGif) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultGif) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultGif) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultGif) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultGif) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultGif) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultGif) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultGif) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultGif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultGif) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultGif) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultGif) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultGif) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultGif) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultGif) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


