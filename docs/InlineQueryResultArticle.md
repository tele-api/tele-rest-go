# InlineQueryResultArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *article* | [default to "article"]
**Id** | **string** | Unique identifier for this result, 1-64 Bytes | 
**Title** | **string** | Title of the result | 
**InputMessageContent** | [**InputMessageContent**](InputMessageContent.md) |  | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**Url** | Pointer to **string** | *Optional*. URL of the result | [optional] 
**Description** | Pointer to **string** | *Optional*. Short description of the result | [optional] 
**ThumbnailUrl** | Pointer to **string** | *Optional*. Url of the thumbnail for the result | [optional] 
**ThumbnailWidth** | Pointer to **int32** | *Optional*. Thumbnail width | [optional] 
**ThumbnailHeight** | Pointer to **int32** | *Optional*. Thumbnail height | [optional] 

## Methods

### NewInlineQueryResultArticle

`func NewInlineQueryResultArticle(type_ string, id string, title string, inputMessageContent InputMessageContent, ) *InlineQueryResultArticle`

NewInlineQueryResultArticle instantiates a new InlineQueryResultArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultArticleWithDefaults

`func NewInlineQueryResultArticleWithDefaults() *InlineQueryResultArticle`

NewInlineQueryResultArticleWithDefaults instantiates a new InlineQueryResultArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultArticle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultArticle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultArticle) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultArticle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultArticle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultArticle) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *InlineQueryResultArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetInputMessageContent

`func (o *InlineQueryResultArticle) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultArticle) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultArticle) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.


### GetReplyMarkup

`func (o *InlineQueryResultArticle) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultArticle) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultArticle) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultArticle) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetUrl

`func (o *InlineQueryResultArticle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineQueryResultArticle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineQueryResultArticle) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineQueryResultArticle) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *InlineQueryResultArticle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineQueryResultArticle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineQueryResultArticle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineQueryResultArticle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *InlineQueryResultArticle) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultArticle) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultArticle) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *InlineQueryResultArticle) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetThumbnailWidth

`func (o *InlineQueryResultArticle) GetThumbnailWidth() int32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *InlineQueryResultArticle) GetThumbnailWidthOk() (*int32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *InlineQueryResultArticle) SetThumbnailWidth(v int32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *InlineQueryResultArticle) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### GetThumbnailHeight

`func (o *InlineQueryResultArticle) GetThumbnailHeight() int32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *InlineQueryResultArticle) GetThumbnailHeightOk() (*int32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *InlineQueryResultArticle) SetThumbnailHeight(v int32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *InlineQueryResultArticle) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


