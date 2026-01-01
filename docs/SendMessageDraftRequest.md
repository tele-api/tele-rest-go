# SendMessageDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | **int32** | Unique identifier for the target private chat | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread | [optional] 
**DraftId** | **int32** | Unique identifier of the message draft; must be non-zero. Changes of drafts with the same identifier are animated | 
**Text** | **string** | Text of the message to be sent, 1-4096 characters after entities parsing | 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in message text, which can be specified instead of *parse\\_mode* | [optional] 

## Methods

### NewSendMessageDraftRequest

`func NewSendMessageDraftRequest(chatId int32, draftId int32, text string, ) *SendMessageDraftRequest`

NewSendMessageDraftRequest instantiates a new SendMessageDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageDraftRequestWithDefaults

`func NewSendMessageDraftRequestWithDefaults() *SendMessageDraftRequest`

NewSendMessageDraftRequestWithDefaults instantiates a new SendMessageDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SendMessageDraftRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendMessageDraftRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendMessageDraftRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendMessageDraftRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendMessageDraftRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendMessageDraftRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendMessageDraftRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetDraftId

`func (o *SendMessageDraftRequest) GetDraftId() int32`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *SendMessageDraftRequest) GetDraftIdOk() (*int32, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *SendMessageDraftRequest) SetDraftId(v int32)`

SetDraftId sets DraftId field to given value.


### GetText

`func (o *SendMessageDraftRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendMessageDraftRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendMessageDraftRequest) SetText(v string)`

SetText sets Text field to given value.


### GetParseMode

`func (o *SendMessageDraftRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *SendMessageDraftRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *SendMessageDraftRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *SendMessageDraftRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetEntities

`func (o *SendMessageDraftRequest) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SendMessageDraftRequest) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SendMessageDraftRequest) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SendMessageDraftRequest) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


