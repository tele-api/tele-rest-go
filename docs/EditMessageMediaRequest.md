# EditMessageMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message to be edited was sent | [optional] 
**ChatId** | Pointer to [**EditMessageTextRequestChatId**](EditMessageTextRequestChatId.md) |  | [optional] 
**MessageId** | Pointer to **int32** | Required if *inline\\_message\\_id* is not specified. Identifier of the message to edit | [optional] 
**InlineMessageId** | Pointer to **string** | Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message | [optional] 
**Media** | [**InputMedia**](InputMedia.md) |  | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewEditMessageMediaRequest

`func NewEditMessageMediaRequest(media InputMedia, ) *EditMessageMediaRequest`

NewEditMessageMediaRequest instantiates a new EditMessageMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMessageMediaRequestWithDefaults

`func NewEditMessageMediaRequestWithDefaults() *EditMessageMediaRequest`

NewEditMessageMediaRequestWithDefaults instantiates a new EditMessageMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *EditMessageMediaRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *EditMessageMediaRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *EditMessageMediaRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *EditMessageMediaRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *EditMessageMediaRequest) GetChatId() EditMessageTextRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditMessageMediaRequest) GetChatIdOk() (*EditMessageTextRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditMessageMediaRequest) SetChatId(v EditMessageTextRequestChatId)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *EditMessageMediaRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *EditMessageMediaRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EditMessageMediaRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EditMessageMediaRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EditMessageMediaRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *EditMessageMediaRequest) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *EditMessageMediaRequest) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *EditMessageMediaRequest) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *EditMessageMediaRequest) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.

### GetMedia

`func (o *EditMessageMediaRequest) GetMedia() InputMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *EditMessageMediaRequest) GetMediaOk() (*InputMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *EditMessageMediaRequest) SetMedia(v InputMedia)`

SetMedia sets Media field to given value.


### GetReplyMarkup

`func (o *EditMessageMediaRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *EditMessageMediaRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *EditMessageMediaRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *EditMessageMediaRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


