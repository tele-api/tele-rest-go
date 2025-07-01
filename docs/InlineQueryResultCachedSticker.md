# InlineQueryResultCachedSticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *sticker* | [default to "sticker"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**StickerFileId** | **string** | A valid file identifier of the sticker | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultCachedSticker

`func NewInlineQueryResultCachedSticker(type_ string, id string, stickerFileId string, ) *InlineQueryResultCachedSticker`

NewInlineQueryResultCachedSticker instantiates a new InlineQueryResultCachedSticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultCachedStickerWithDefaults

`func NewInlineQueryResultCachedStickerWithDefaults() *InlineQueryResultCachedSticker`

NewInlineQueryResultCachedStickerWithDefaults instantiates a new InlineQueryResultCachedSticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultCachedSticker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultCachedSticker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultCachedSticker) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultCachedSticker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultCachedSticker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultCachedSticker) SetId(v string)`

SetId sets Id field to given value.


### GetStickerFileId

`func (o *InlineQueryResultCachedSticker) GetStickerFileId() string`

GetStickerFileId returns the StickerFileId field if non-nil, zero value otherwise.

### GetStickerFileIdOk

`func (o *InlineQueryResultCachedSticker) GetStickerFileIdOk() (*string, bool)`

GetStickerFileIdOk returns a tuple with the StickerFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerFileId

`func (o *InlineQueryResultCachedSticker) SetStickerFileId(v string)`

SetStickerFileId sets StickerFileId field to given value.


### GetReplyMarkup

`func (o *InlineQueryResultCachedSticker) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultCachedSticker) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultCachedSticker) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultCachedSticker) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultCachedSticker) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultCachedSticker) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultCachedSticker) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultCachedSticker) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


