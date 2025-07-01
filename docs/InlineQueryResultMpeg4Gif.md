# InlineQueryResultMpeg4Gif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *mpeg4\\_gif* | [default to "mpeg4_gif"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**Mpeg4Url** | **string** | A valid URL for the MPEG4 file | 
**Mpeg4Width** | Pointer to **int32** | *Optional*. Video width | [optional] 
**Mpeg4Height** | Pointer to **int32** | *Optional*. Video height | [optional] 
**Mpeg4Duration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**ThumbnailUrl** | **string** | URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result | 
**ThumbnailMimeType** | Pointer to **string** | *Optional*. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg” | [optional] [default to "image/jpeg"]
**Title** | Pointer to **string** | *Optional*. Title for the result | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultMpeg4Gif

`func NewInlineQueryResultMpeg4Gif(type_ string, id string, mpeg4Url string, thumbnailUrl string, ) *InlineQueryResultMpeg4Gif`

NewInlineQueryResultMpeg4Gif instantiates a new InlineQueryResultMpeg4Gif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultMpeg4GifWithDefaults

`func NewInlineQueryResultMpeg4GifWithDefaults() *InlineQueryResultMpeg4Gif`

NewInlineQueryResultMpeg4GifWithDefaults instantiates a new InlineQueryResultMpeg4Gif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultMpeg4Gif) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultMpeg4Gif) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultMpeg4Gif) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultMpeg4Gif) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultMpeg4Gif) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultMpeg4Gif) SetId(v string)`

SetId sets Id field to given value.


### GetMpeg4Url

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4Url() string`

GetMpeg4Url returns the Mpeg4Url field if non-nil, zero value otherwise.

### GetMpeg4UrlOk

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4UrlOk() (*string, bool)`

GetMpeg4UrlOk returns a tuple with the Mpeg4Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Url

`func (o *InlineQueryResultMpeg4Gif) SetMpeg4Url(v string)`

SetMpeg4Url sets Mpeg4Url field to given value.


### GetMpeg4Width

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4Width() int32`

GetMpeg4Width returns the Mpeg4Width field if non-nil, zero value otherwise.

### GetMpeg4WidthOk

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4WidthOk() (*int32, bool)`

GetMpeg4WidthOk returns a tuple with the Mpeg4Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Width

`func (o *InlineQueryResultMpeg4Gif) SetMpeg4Width(v int32)`

SetMpeg4Width sets Mpeg4Width field to given value.

### HasMpeg4Width

`func (o *InlineQueryResultMpeg4Gif) HasMpeg4Width() bool`

HasMpeg4Width returns a boolean if a field has been set.

### GetMpeg4Height

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4Height() int32`

GetMpeg4Height returns the Mpeg4Height field if non-nil, zero value otherwise.

### GetMpeg4HeightOk

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4HeightOk() (*int32, bool)`

GetMpeg4HeightOk returns a tuple with the Mpeg4Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Height

`func (o *InlineQueryResultMpeg4Gif) SetMpeg4Height(v int32)`

SetMpeg4Height sets Mpeg4Height field to given value.

### HasMpeg4Height

`func (o *InlineQueryResultMpeg4Gif) HasMpeg4Height() bool`

HasMpeg4Height returns a boolean if a field has been set.

### GetMpeg4Duration

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4Duration() int32`

GetMpeg4Duration returns the Mpeg4Duration field if non-nil, zero value otherwise.

### GetMpeg4DurationOk

`func (o *InlineQueryResultMpeg4Gif) GetMpeg4DurationOk() (*int32, bool)`

GetMpeg4DurationOk returns a tuple with the Mpeg4Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4Duration

`func (o *InlineQueryResultMpeg4Gif) SetMpeg4Duration(v int32)`

SetMpeg4Duration sets Mpeg4Duration field to given value.

### HasMpeg4Duration

`func (o *InlineQueryResultMpeg4Gif) HasMpeg4Duration() bool`

HasMpeg4Duration returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *InlineQueryResultMpeg4Gif) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultMpeg4Gif) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultMpeg4Gif) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetThumbnailMimeType

`func (o *InlineQueryResultMpeg4Gif) GetThumbnailMimeType() string`

GetThumbnailMimeType returns the ThumbnailMimeType field if non-nil, zero value otherwise.

### GetThumbnailMimeTypeOk

`func (o *InlineQueryResultMpeg4Gif) GetThumbnailMimeTypeOk() (*string, bool)`

GetThumbnailMimeTypeOk returns a tuple with the ThumbnailMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailMimeType

`func (o *InlineQueryResultMpeg4Gif) SetThumbnailMimeType(v string)`

SetThumbnailMimeType sets ThumbnailMimeType field to given value.

### HasThumbnailMimeType

`func (o *InlineQueryResultMpeg4Gif) HasThumbnailMimeType() bool`

HasThumbnailMimeType returns a boolean if a field has been set.

### GetTitle

`func (o *InlineQueryResultMpeg4Gif) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultMpeg4Gif) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultMpeg4Gif) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineQueryResultMpeg4Gif) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCaption

`func (o *InlineQueryResultMpeg4Gif) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultMpeg4Gif) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultMpeg4Gif) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultMpeg4Gif) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultMpeg4Gif) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultMpeg4Gif) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultMpeg4Gif) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultMpeg4Gif) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultMpeg4Gif) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultMpeg4Gif) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultMpeg4Gif) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultMpeg4Gif) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultMpeg4Gif) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultMpeg4Gif) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultMpeg4Gif) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultMpeg4Gif) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultMpeg4Gif) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultMpeg4Gif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultMpeg4Gif) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultMpeg4Gif) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultMpeg4Gif) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultMpeg4Gif) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultMpeg4Gif) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultMpeg4Gif) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


