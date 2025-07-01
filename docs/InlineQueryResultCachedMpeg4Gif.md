# InlineQueryResultCachedMpeg4Gif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *mpeg4\\_gif* | [default to "mpeg4_gif"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**Mpeg4FileId** | **string** | A valid file identifier for the MPEG4 file | 
**Title** | Pointer to **string** | *Optional*. Title for the result | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultCachedMpeg4Gif

`func NewInlineQueryResultCachedMpeg4Gif(type_ string, id string, mpeg4FileId string, ) *InlineQueryResultCachedMpeg4Gif`

NewInlineQueryResultCachedMpeg4Gif instantiates a new InlineQueryResultCachedMpeg4Gif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultCachedMpeg4GifWithDefaults

`func NewInlineQueryResultCachedMpeg4GifWithDefaults() *InlineQueryResultCachedMpeg4Gif`

NewInlineQueryResultCachedMpeg4GifWithDefaults instantiates a new InlineQueryResultCachedMpeg4Gif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultCachedMpeg4Gif) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultCachedMpeg4Gif) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultCachedMpeg4Gif) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultCachedMpeg4Gif) SetId(v string)`

SetId sets Id field to given value.


### GetMpeg4FileId

`func (o *InlineQueryResultCachedMpeg4Gif) GetMpeg4FileId() string`

GetMpeg4FileId returns the Mpeg4FileId field if non-nil, zero value otherwise.

### GetMpeg4FileIdOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetMpeg4FileIdOk() (*string, bool)`

GetMpeg4FileIdOk returns a tuple with the Mpeg4FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpeg4FileId

`func (o *InlineQueryResultCachedMpeg4Gif) SetMpeg4FileId(v string)`

SetMpeg4FileId sets Mpeg4FileId field to given value.


### GetTitle

`func (o *InlineQueryResultCachedMpeg4Gif) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultCachedMpeg4Gif) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineQueryResultCachedMpeg4Gif) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCaption

`func (o *InlineQueryResultCachedMpeg4Gif) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultCachedMpeg4Gif) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultCachedMpeg4Gif) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultCachedMpeg4Gif) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultCachedMpeg4Gif) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultCachedMpeg4Gif) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultCachedMpeg4Gif) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultCachedMpeg4Gif) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultCachedMpeg4Gif) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultCachedMpeg4Gif) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultCachedMpeg4Gif) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultCachedMpeg4Gif) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultCachedMpeg4Gif) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultCachedMpeg4Gif) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultCachedMpeg4Gif) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultCachedMpeg4Gif) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultCachedMpeg4Gif) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultCachedMpeg4Gif) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


