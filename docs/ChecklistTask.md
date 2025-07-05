# ChecklistTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the task | 
**Text** | **string** | Text of the task | 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the task text | [optional] 
**CompletedByUser** | Pointer to [**User**](User.md) |  | [optional] 
**CompletionDate** | Pointer to **int32** | *Optional*. Point in time (Unix timestamp) when the task was completed; 0 if the task wasn&#39;t completed | [optional] 

## Methods

### NewChecklistTask

`func NewChecklistTask(id int32, text string, ) *ChecklistTask`

NewChecklistTask instantiates a new ChecklistTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecklistTaskWithDefaults

`func NewChecklistTaskWithDefaults() *ChecklistTask`

NewChecklistTaskWithDefaults instantiates a new ChecklistTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChecklistTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChecklistTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChecklistTask) SetId(v int32)`

SetId sets Id field to given value.


### GetText

`func (o *ChecklistTask) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChecklistTask) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChecklistTask) SetText(v string)`

SetText sets Text field to given value.


### GetTextEntities

`func (o *ChecklistTask) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *ChecklistTask) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *ChecklistTask) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *ChecklistTask) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.

### GetCompletedByUser

`func (o *ChecklistTask) GetCompletedByUser() User`

GetCompletedByUser returns the CompletedByUser field if non-nil, zero value otherwise.

### GetCompletedByUserOk

`func (o *ChecklistTask) GetCompletedByUserOk() (*User, bool)`

GetCompletedByUserOk returns a tuple with the CompletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUser

`func (o *ChecklistTask) SetCompletedByUser(v User)`

SetCompletedByUser sets CompletedByUser field to given value.

### HasCompletedByUser

`func (o *ChecklistTask) HasCompletedByUser() bool`

HasCompletedByUser returns a boolean if a field has been set.

### GetCompletionDate

`func (o *ChecklistTask) GetCompletionDate() int32`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *ChecklistTask) GetCompletionDateOk() (*int32, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *ChecklistTask) SetCompletionDate(v int32)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *ChecklistTask) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


