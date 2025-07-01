# ReplyParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | Identifier of the message that will be replied to in the current chat, or in the chat *chat\\_id* if it is specified | 
**ChatId** | Pointer to [**ReplyParametersChatId**](ReplyParametersChatId.md) |  | [optional] 
**AllowSendingWithoutReply** | Pointer to **bool** | *Optional*. Pass *True* if the message should be sent even if the specified message to be replied to is not found. Always *False* for replies in another chat or forum topic. Always *True* for messages sent on behalf of a business account. | [optional] 
**Quote** | Pointer to **string** | *Optional*. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities. The message will fail to send if the quote isn&#39;t found in the original message. | [optional] 
**QuoteParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the quote. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**QuoteEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of *quote\\_parse\\_mode*. | [optional] 
**QuotePosition** | Pointer to **int32** | *Optional*. Position of the quote in the original message in UTF-16 code units | [optional] 

## Methods

### NewReplyParameters

`func NewReplyParameters(messageId int32, ) *ReplyParameters`

NewReplyParameters instantiates a new ReplyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyParametersWithDefaults

`func NewReplyParametersWithDefaults() *ReplyParameters`

NewReplyParametersWithDefaults instantiates a new ReplyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *ReplyParameters) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ReplyParameters) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ReplyParameters) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetChatId

`func (o *ReplyParameters) GetChatId() ReplyParametersChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ReplyParameters) GetChatIdOk() (*ReplyParametersChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ReplyParameters) SetChatId(v ReplyParametersChatId)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *ReplyParameters) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetAllowSendingWithoutReply

`func (o *ReplyParameters) GetAllowSendingWithoutReply() bool`

GetAllowSendingWithoutReply returns the AllowSendingWithoutReply field if non-nil, zero value otherwise.

### GetAllowSendingWithoutReplyOk

`func (o *ReplyParameters) GetAllowSendingWithoutReplyOk() (*bool, bool)`

GetAllowSendingWithoutReplyOk returns a tuple with the AllowSendingWithoutReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSendingWithoutReply

`func (o *ReplyParameters) SetAllowSendingWithoutReply(v bool)`

SetAllowSendingWithoutReply sets AllowSendingWithoutReply field to given value.

### HasAllowSendingWithoutReply

`func (o *ReplyParameters) HasAllowSendingWithoutReply() bool`

HasAllowSendingWithoutReply returns a boolean if a field has been set.

### GetQuote

`func (o *ReplyParameters) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *ReplyParameters) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *ReplyParameters) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *ReplyParameters) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetQuoteParseMode

`func (o *ReplyParameters) GetQuoteParseMode() string`

GetQuoteParseMode returns the QuoteParseMode field if non-nil, zero value otherwise.

### GetQuoteParseModeOk

`func (o *ReplyParameters) GetQuoteParseModeOk() (*string, bool)`

GetQuoteParseModeOk returns a tuple with the QuoteParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteParseMode

`func (o *ReplyParameters) SetQuoteParseMode(v string)`

SetQuoteParseMode sets QuoteParseMode field to given value.

### HasQuoteParseMode

`func (o *ReplyParameters) HasQuoteParseMode() bool`

HasQuoteParseMode returns a boolean if a field has been set.

### GetQuoteEntities

`func (o *ReplyParameters) GetQuoteEntities() []MessageEntity`

GetQuoteEntities returns the QuoteEntities field if non-nil, zero value otherwise.

### GetQuoteEntitiesOk

`func (o *ReplyParameters) GetQuoteEntitiesOk() (*[]MessageEntity, bool)`

GetQuoteEntitiesOk returns a tuple with the QuoteEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEntities

`func (o *ReplyParameters) SetQuoteEntities(v []MessageEntity)`

SetQuoteEntities sets QuoteEntities field to given value.

### HasQuoteEntities

`func (o *ReplyParameters) HasQuoteEntities() bool`

HasQuoteEntities returns a boolean if a field has been set.

### GetQuotePosition

`func (o *ReplyParameters) GetQuotePosition() int32`

GetQuotePosition returns the QuotePosition field if non-nil, zero value otherwise.

### GetQuotePositionOk

`func (o *ReplyParameters) GetQuotePositionOk() (*int32, bool)`

GetQuotePositionOk returns a tuple with the QuotePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePosition

`func (o *ReplyParameters) SetQuotePosition(v int32)`

SetQuotePosition sets QuotePosition field to given value.

### HasQuotePosition

`func (o *ReplyParameters) HasQuotePosition() bool`

HasQuotePosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


