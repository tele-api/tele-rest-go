# InlineQueryResultCachedVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *video* | [default to "video"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**VideoFileId** | **string** | A valid file identifier for the video file | 
**Title** | **string** | Title for the result | 
**Description** | Pointer to **string** | *Optional*. Short description of the result | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the video to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultCachedVideo

`func NewInlineQueryResultCachedVideo(type_ string, id string, videoFileId string, title string, ) *InlineQueryResultCachedVideo`

NewInlineQueryResultCachedVideo instantiates a new InlineQueryResultCachedVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultCachedVideoWithDefaults

`func NewInlineQueryResultCachedVideoWithDefaults() *InlineQueryResultCachedVideo`

NewInlineQueryResultCachedVideoWithDefaults instantiates a new InlineQueryResultCachedVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultCachedVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultCachedVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultCachedVideo) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultCachedVideo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultCachedVideo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultCachedVideo) SetId(v string)`

SetId sets Id field to given value.


### GetVideoFileId

`func (o *InlineQueryResultCachedVideo) GetVideoFileId() string`

GetVideoFileId returns the VideoFileId field if non-nil, zero value otherwise.

### GetVideoFileIdOk

`func (o *InlineQueryResultCachedVideo) GetVideoFileIdOk() (*string, bool)`

GetVideoFileIdOk returns a tuple with the VideoFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFileId

`func (o *InlineQueryResultCachedVideo) SetVideoFileId(v string)`

SetVideoFileId sets VideoFileId field to given value.


### GetTitle

`func (o *InlineQueryResultCachedVideo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultCachedVideo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultCachedVideo) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InlineQueryResultCachedVideo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineQueryResultCachedVideo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineQueryResultCachedVideo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineQueryResultCachedVideo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCaption

`func (o *InlineQueryResultCachedVideo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultCachedVideo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultCachedVideo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultCachedVideo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultCachedVideo) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultCachedVideo) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultCachedVideo) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultCachedVideo) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultCachedVideo) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultCachedVideo) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultCachedVideo) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultCachedVideo) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultCachedVideo) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultCachedVideo) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultCachedVideo) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultCachedVideo) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultCachedVideo) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultCachedVideo) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultCachedVideo) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultCachedVideo) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultCachedVideo) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultCachedVideo) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultCachedVideo) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultCachedVideo) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


