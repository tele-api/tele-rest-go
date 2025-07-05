# TextQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Text of the quoted part of a message that is replied to by the given message | 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the quote. Currently, only *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities are kept in quotes. | [optional] 
**Position** | **int32** | Approximate quote position in the original message in UTF-16 code units as specified by the sender | 
**IsManual** | Pointer to **bool** | *Optional*. *True*, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server. | [optional] [default to true]

## Methods

### NewTextQuote

`func NewTextQuote(text string, position int32, ) *TextQuote`

NewTextQuote instantiates a new TextQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextQuoteWithDefaults

`func NewTextQuoteWithDefaults() *TextQuote`

NewTextQuoteWithDefaults instantiates a new TextQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TextQuote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextQuote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextQuote) SetText(v string)`

SetText sets Text field to given value.


### GetEntities

`func (o *TextQuote) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TextQuote) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TextQuote) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *TextQuote) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetPosition

`func (o *TextQuote) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TextQuote) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TextQuote) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetIsManual

`func (o *TextQuote) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TextQuote) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TextQuote) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TextQuote) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


