# EditMessageTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message to be edited was sent | [optional] 
**ChatId** | Pointer to [**EditMessageTextRequestChatId**](EditMessageTextRequestChatId.md) |  | [optional] 
**MessageId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit | [optional] 
**InlineMessageId** | Pointer to **string** | Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message | [optional] 
**Text** | **string** | New text of the message, 1-4096 characters after entities parsing | 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in message text, which can be specified instead of *parse\\_mode* | [optional] 
**LinkPreviewOptions** | Pointer to [**LinkPreviewOptions**](LinkPreviewOptions.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewEditMessageTextRequest

`func NewEditMessageTextRequest(text string, ) *EditMessageTextRequest`

NewEditMessageTextRequest instantiates a new EditMessageTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMessageTextRequestWithDefaults

`func NewEditMessageTextRequestWithDefaults() *EditMessageTextRequest`

NewEditMessageTextRequestWithDefaults instantiates a new EditMessageTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *EditMessageTextRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *EditMessageTextRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *EditMessageTextRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *EditMessageTextRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *EditMessageTextRequest) GetChatId() EditMessageTextRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditMessageTextRequest) GetChatIdOk() (*EditMessageTextRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditMessageTextRequest) SetChatId(v EditMessageTextRequestChatId)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *EditMessageTextRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *EditMessageTextRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EditMessageTextRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EditMessageTextRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EditMessageTextRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *EditMessageTextRequest) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *EditMessageTextRequest) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *EditMessageTextRequest) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *EditMessageTextRequest) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.

### GetText

`func (o *EditMessageTextRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EditMessageTextRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EditMessageTextRequest) SetText(v string)`

SetText sets Text field to given value.


### GetParseMode

`func (o *EditMessageTextRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *EditMessageTextRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *EditMessageTextRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *EditMessageTextRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetEntities

`func (o *EditMessageTextRequest) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EditMessageTextRequest) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EditMessageTextRequest) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *EditMessageTextRequest) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetLinkPreviewOptions

`func (o *EditMessageTextRequest) GetLinkPreviewOptions() LinkPreviewOptions`

GetLinkPreviewOptions returns the LinkPreviewOptions field if non-nil, zero value otherwise.

### GetLinkPreviewOptionsOk

`func (o *EditMessageTextRequest) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool)`

GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPreviewOptions

`func (o *EditMessageTextRequest) SetLinkPreviewOptions(v LinkPreviewOptions)`

SetLinkPreviewOptions sets LinkPreviewOptions field to given value.

### HasLinkPreviewOptions

`func (o *EditMessageTextRequest) HasLinkPreviewOptions() bool`

HasLinkPreviewOptions returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *EditMessageTextRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *EditMessageTextRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *EditMessageTextRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *EditMessageTextRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


