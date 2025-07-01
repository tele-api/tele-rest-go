# InlineQueryResultGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *game* | [default to "game"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**GameShortName** | **string** | Short name of the game | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewInlineQueryResultGame

`func NewInlineQueryResultGame(type_ string, id string, gameShortName string, ) *InlineQueryResultGame`

NewInlineQueryResultGame instantiates a new InlineQueryResultGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultGameWithDefaults

`func NewInlineQueryResultGameWithDefaults() *InlineQueryResultGame`

NewInlineQueryResultGameWithDefaults instantiates a new InlineQueryResultGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultGame) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultGame) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultGame) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultGame) SetId(v string)`

SetId sets Id field to given value.


### GetGameShortName

`func (o *InlineQueryResultGame) GetGameShortName() string`

GetGameShortName returns the GameShortName field if non-nil, zero value otherwise.

### GetGameShortNameOk

`func (o *InlineQueryResultGame) GetGameShortNameOk() (*string, bool)`

GetGameShortNameOk returns a tuple with the GameShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameShortName

`func (o *InlineQueryResultGame) SetGameShortName(v string)`

SetGameShortName sets GameShortName field to given value.


### GetReplyMarkup

`func (o *InlineQueryResultGame) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultGame) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultGame) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultGame) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


