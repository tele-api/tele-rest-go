# InputTextMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageText** | **string** | Text of the message to be sent, 1-4096 characters | 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in message text, which can be specified instead of *parse\\_mode* | [optional] 
**LinkPreviewOptions** | Pointer to [**LinkPreviewOptions**](LinkPreviewOptions.md) |  | [optional] 

## Methods

### NewInputTextMessageContent

`func NewInputTextMessageContent(messageText string, ) *InputTextMessageContent`

NewInputTextMessageContent instantiates a new InputTextMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTextMessageContentWithDefaults

`func NewInputTextMessageContentWithDefaults() *InputTextMessageContent`

NewInputTextMessageContentWithDefaults instantiates a new InputTextMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageText

`func (o *InputTextMessageContent) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *InputTextMessageContent) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *InputTextMessageContent) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.


### GetParseMode

`func (o *InputTextMessageContent) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InputTextMessageContent) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InputTextMessageContent) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InputTextMessageContent) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetEntities

`func (o *InputTextMessageContent) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *InputTextMessageContent) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *InputTextMessageContent) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *InputTextMessageContent) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetLinkPreviewOptions

`func (o *InputTextMessageContent) GetLinkPreviewOptions() LinkPreviewOptions`

GetLinkPreviewOptions returns the LinkPreviewOptions field if non-nil, zero value otherwise.

### GetLinkPreviewOptionsOk

`func (o *InputTextMessageContent) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool)`

GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPreviewOptions

`func (o *InputTextMessageContent) SetLinkPreviewOptions(v LinkPreviewOptions)`

SetLinkPreviewOptions sets LinkPreviewOptions field to given value.

### HasLinkPreviewOptions

`func (o *InputTextMessageContent) HasLinkPreviewOptions() bool`

HasLinkPreviewOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


