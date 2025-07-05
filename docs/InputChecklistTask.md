# InputChecklistTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the task; must be positive and unique among all task identifiers currently present in the checklist | 
**Text** | **string** | Text of the task; 1-100 characters after entities parsing | 
**ParseMode** | Pointer to **string** | Optional. Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details. | [optional] 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the text, which can be specified instead of parse\\_mode. Currently, only *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities are allowed. | [optional] 

## Methods

### NewInputChecklistTask

`func NewInputChecklistTask(id int32, text string, ) *InputChecklistTask`

NewInputChecklistTask instantiates a new InputChecklistTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputChecklistTaskWithDefaults

`func NewInputChecklistTaskWithDefaults() *InputChecklistTask`

NewInputChecklistTaskWithDefaults instantiates a new InputChecklistTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputChecklistTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputChecklistTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputChecklistTask) SetId(v int32)`

SetId sets Id field to given value.


### GetText

`func (o *InputChecklistTask) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InputChecklistTask) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InputChecklistTask) SetText(v string)`

SetText sets Text field to given value.


### GetParseMode

`func (o *InputChecklistTask) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InputChecklistTask) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InputChecklistTask) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InputChecklistTask) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetTextEntities

`func (o *InputChecklistTask) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *InputChecklistTask) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *InputChecklistTask) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *InputChecklistTask) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


