# InlineQueryResultVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *video* | [default to "video"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**VideoUrl** | **string** | A valid URL for the embedded video player or video file | 
**MimeType** | **string** | MIME type of the content of the video URL, “text/html” or “video/mp4” | 
**ThumbnailUrl** | **string** | URL of the thumbnail (JPEG only) for the video | 
**Title** | **string** | Title for the result | 
**Caption** | Pointer to **string** | *Optional*. Caption of the video to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**VideoWidth** | Pointer to **int32** | *Optional*. Video width | [optional] 
**VideoHeight** | Pointer to **int32** | *Optional*. Video height | [optional] 
**VideoDuration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**Description** | Pointer to **string** | *Optional*. Short description of the result | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultVideo

`func NewInlineQueryResultVideo(type_ string, id string, videoUrl string, mimeType string, thumbnailUrl string, title string, ) *InlineQueryResultVideo`

NewInlineQueryResultVideo instantiates a new InlineQueryResultVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultVideoWithDefaults

`func NewInlineQueryResultVideoWithDefaults() *InlineQueryResultVideo`

NewInlineQueryResultVideoWithDefaults instantiates a new InlineQueryResultVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultVideo) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultVideo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultVideo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultVideo) SetId(v string)`

SetId sets Id field to given value.


### GetVideoUrl

`func (o *InlineQueryResultVideo) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *InlineQueryResultVideo) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *InlineQueryResultVideo) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.


### GetMimeType

`func (o *InlineQueryResultVideo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *InlineQueryResultVideo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *InlineQueryResultVideo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetThumbnailUrl

`func (o *InlineQueryResultVideo) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultVideo) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultVideo) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetTitle

`func (o *InlineQueryResultVideo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultVideo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultVideo) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCaption

`func (o *InlineQueryResultVideo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultVideo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultVideo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultVideo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultVideo) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultVideo) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultVideo) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultVideo) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultVideo) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultVideo) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultVideo) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultVideo) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultVideo) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultVideo) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultVideo) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultVideo) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetVideoWidth

`func (o *InlineQueryResultVideo) GetVideoWidth() int32`

GetVideoWidth returns the VideoWidth field if non-nil, zero value otherwise.

### GetVideoWidthOk

`func (o *InlineQueryResultVideo) GetVideoWidthOk() (*int32, bool)`

GetVideoWidthOk returns a tuple with the VideoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoWidth

`func (o *InlineQueryResultVideo) SetVideoWidth(v int32)`

SetVideoWidth sets VideoWidth field to given value.

### HasVideoWidth

`func (o *InlineQueryResultVideo) HasVideoWidth() bool`

HasVideoWidth returns a boolean if a field has been set.

### GetVideoHeight

`func (o *InlineQueryResultVideo) GetVideoHeight() int32`

GetVideoHeight returns the VideoHeight field if non-nil, zero value otherwise.

### GetVideoHeightOk

`func (o *InlineQueryResultVideo) GetVideoHeightOk() (*int32, bool)`

GetVideoHeightOk returns a tuple with the VideoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoHeight

`func (o *InlineQueryResultVideo) SetVideoHeight(v int32)`

SetVideoHeight sets VideoHeight field to given value.

### HasVideoHeight

`func (o *InlineQueryResultVideo) HasVideoHeight() bool`

HasVideoHeight returns a boolean if a field has been set.

### GetVideoDuration

`func (o *InlineQueryResultVideo) GetVideoDuration() int32`

GetVideoDuration returns the VideoDuration field if non-nil, zero value otherwise.

### GetVideoDurationOk

`func (o *InlineQueryResultVideo) GetVideoDurationOk() (*int32, bool)`

GetVideoDurationOk returns a tuple with the VideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDuration

`func (o *InlineQueryResultVideo) SetVideoDuration(v int32)`

SetVideoDuration sets VideoDuration field to given value.

### HasVideoDuration

`func (o *InlineQueryResultVideo) HasVideoDuration() bool`

HasVideoDuration returns a boolean if a field has been set.

### GetDescription

`func (o *InlineQueryResultVideo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineQueryResultVideo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineQueryResultVideo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineQueryResultVideo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultVideo) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultVideo) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultVideo) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultVideo) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultVideo) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultVideo) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultVideo) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultVideo) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


