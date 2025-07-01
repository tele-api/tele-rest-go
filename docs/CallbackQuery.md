# CallbackQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this query | 
**From** | [**User**](User.md) |  | 
**Message** | Pointer to [**MaybeInaccessibleMessage**](MaybeInaccessibleMessage.md) |  | [optional] 
**InlineMessageId** | Pointer to **string** | *Optional*. Identifier of the message sent via the bot in inline mode, that originated the query. | [optional] 
**ChatInstance** | **string** | Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in [games](https://core.telegram.org/bots/api/#games). | 
**Data** | Pointer to **string** | *Optional*. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data. | [optional] 
**GameShortName** | Pointer to **string** | *Optional*. Short name of a [Game](https://core.telegram.org/bots/api/#games) to be returned, serves as the unique identifier for the game | [optional] 

## Methods

### NewCallbackQuery

`func NewCallbackQuery(id string, from User, chatInstance string, ) *CallbackQuery`

NewCallbackQuery instantiates a new CallbackQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackQueryWithDefaults

`func NewCallbackQueryWithDefaults() *CallbackQuery`

NewCallbackQueryWithDefaults instantiates a new CallbackQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallbackQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallbackQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallbackQuery) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *CallbackQuery) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallbackQuery) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallbackQuery) SetFrom(v User)`

SetFrom sets From field to given value.


### GetMessage

`func (o *CallbackQuery) GetMessage() MaybeInaccessibleMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CallbackQuery) GetMessageOk() (*MaybeInaccessibleMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CallbackQuery) SetMessage(v MaybeInaccessibleMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CallbackQuery) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *CallbackQuery) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *CallbackQuery) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *CallbackQuery) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *CallbackQuery) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.

### GetChatInstance

`func (o *CallbackQuery) GetChatInstance() string`

GetChatInstance returns the ChatInstance field if non-nil, zero value otherwise.

### GetChatInstanceOk

`func (o *CallbackQuery) GetChatInstanceOk() (*string, bool)`

GetChatInstanceOk returns a tuple with the ChatInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInstance

`func (o *CallbackQuery) SetChatInstance(v string)`

SetChatInstance sets ChatInstance field to given value.


### GetData

`func (o *CallbackQuery) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallbackQuery) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallbackQuery) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CallbackQuery) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGameShortName

`func (o *CallbackQuery) GetGameShortName() string`

GetGameShortName returns the GameShortName field if non-nil, zero value otherwise.

### GetGameShortNameOk

`func (o *CallbackQuery) GetGameShortNameOk() (*string, bool)`

GetGameShortNameOk returns a tuple with the GameShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameShortName

`func (o *CallbackQuery) SetGameShortName(v string)`

SetGameShortName sets GameShortName field to given value.

### HasGameShortName

`func (o *CallbackQuery) HasGameShortName() bool`

HasGameShortName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


