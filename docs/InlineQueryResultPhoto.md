# InlineQueryResultPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *photo* | [default to "photo"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**PhotoUrl** | **string** | A valid URL of the photo. Photo must be in **JPEG** format. Photo size must not exceed 5MB | 
**ThumbnailUrl** | **string** | URL of the thumbnail for the photo | 
**PhotoWidth** | Pointer to **int32** | *Optional*. Width of the photo | [optional] 
**PhotoHeight** | Pointer to **int32** | *Optional*. Height of the photo | [optional] 
**Title** | Pointer to **string** | *Optional*. Title for the result | [optional] 
**Description** | Pointer to **string** | *Optional*. Short description of the result | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the photo to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the photo caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultPhoto

`func NewInlineQueryResultPhoto(type_ string, id string, photoUrl string, thumbnailUrl string, ) *InlineQueryResultPhoto`

NewInlineQueryResultPhoto instantiates a new InlineQueryResultPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultPhotoWithDefaults

`func NewInlineQueryResultPhotoWithDefaults() *InlineQueryResultPhoto`

NewInlineQueryResultPhotoWithDefaults instantiates a new InlineQueryResultPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultPhoto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultPhoto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultPhoto) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultPhoto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultPhoto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultPhoto) SetId(v string)`

SetId sets Id field to given value.


### GetPhotoUrl

`func (o *InlineQueryResultPhoto) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *InlineQueryResultPhoto) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *InlineQueryResultPhoto) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.


### GetThumbnailUrl

`func (o *InlineQueryResultPhoto) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultPhoto) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultPhoto) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetPhotoWidth

`func (o *InlineQueryResultPhoto) GetPhotoWidth() int32`

GetPhotoWidth returns the PhotoWidth field if non-nil, zero value otherwise.

### GetPhotoWidthOk

`func (o *InlineQueryResultPhoto) GetPhotoWidthOk() (*int32, bool)`

GetPhotoWidthOk returns a tuple with the PhotoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoWidth

`func (o *InlineQueryResultPhoto) SetPhotoWidth(v int32)`

SetPhotoWidth sets PhotoWidth field to given value.

### HasPhotoWidth

`func (o *InlineQueryResultPhoto) HasPhotoWidth() bool`

HasPhotoWidth returns a boolean if a field has been set.

### GetPhotoHeight

`func (o *InlineQueryResultPhoto) GetPhotoHeight() int32`

GetPhotoHeight returns the PhotoHeight field if non-nil, zero value otherwise.

### GetPhotoHeightOk

`func (o *InlineQueryResultPhoto) GetPhotoHeightOk() (*int32, bool)`

GetPhotoHeightOk returns a tuple with the PhotoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoHeight

`func (o *InlineQueryResultPhoto) SetPhotoHeight(v int32)`

SetPhotoHeight sets PhotoHeight field to given value.

### HasPhotoHeight

`func (o *InlineQueryResultPhoto) HasPhotoHeight() bool`

HasPhotoHeight returns a boolean if a field has been set.

### GetTitle

`func (o *InlineQueryResultPhoto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultPhoto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultPhoto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineQueryResultPhoto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InlineQueryResultPhoto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineQueryResultPhoto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineQueryResultPhoto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineQueryResultPhoto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCaption

`func (o *InlineQueryResultPhoto) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultPhoto) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultPhoto) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultPhoto) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultPhoto) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultPhoto) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultPhoto) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultPhoto) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultPhoto) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultPhoto) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultPhoto) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultPhoto) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InlineQueryResultPhoto) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InlineQueryResultPhoto) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InlineQueryResultPhoto) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InlineQueryResultPhoto) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultPhoto) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultPhoto) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultPhoto) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultPhoto) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultPhoto) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultPhoto) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultPhoto) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultPhoto) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


