# InputChecklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the checklist; 1-255 characters after entities parsing | 
**ParseMode** | Pointer to **string** | Optional. Mode for parsing entities in the title. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**TitleEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the title, which can be specified instead of parse\\_mode. Currently, only *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities are allowed. | [optional] 
**Tasks** | [**[]InputChecklistTask**](InputChecklistTask.md) | List of 1-30 tasks in the checklist | 
**OthersCanAddTasks** | Pointer to **bool** | *Optional*. Pass *True* if other users can add tasks to the checklist | [optional] 
**OthersCanMarkTasksAsDone** | Pointer to **bool** | *Optional*. Pass *True* if other users can mark tasks as done or not done in the checklist | [optional] 

## Methods

### NewInputChecklist

`func NewInputChecklist(title string, tasks []InputChecklistTask, ) *InputChecklist`

NewInputChecklist instantiates a new InputChecklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputChecklistWithDefaults

`func NewInputChecklistWithDefaults() *InputChecklist`

NewInputChecklistWithDefaults instantiates a new InputChecklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputChecklist) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputChecklist) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputChecklist) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetParseMode

`func (o *InputChecklist) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InputChecklist) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InputChecklist) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InputChecklist) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetTitleEntities

`func (o *InputChecklist) GetTitleEntities() []MessageEntity`

GetTitleEntities returns the TitleEntities field if non-nil, zero value otherwise.

### GetTitleEntitiesOk

`func (o *InputChecklist) GetTitleEntitiesOk() (*[]MessageEntity, bool)`

GetTitleEntitiesOk returns a tuple with the TitleEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEntities

`func (o *InputChecklist) SetTitleEntities(v []MessageEntity)`

SetTitleEntities sets TitleEntities field to given value.

### HasTitleEntities

`func (o *InputChecklist) HasTitleEntities() bool`

HasTitleEntities returns a boolean if a field has been set.

### GetTasks

`func (o *InputChecklist) GetTasks() []InputChecklistTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputChecklist) GetTasksOk() (*[]InputChecklistTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputChecklist) SetTasks(v []InputChecklistTask)`

SetTasks sets Tasks field to given value.


### GetOthersCanAddTasks

`func (o *InputChecklist) GetOthersCanAddTasks() bool`

GetOthersCanAddTasks returns the OthersCanAddTasks field if non-nil, zero value otherwise.

### GetOthersCanAddTasksOk

`func (o *InputChecklist) GetOthersCanAddTasksOk() (*bool, bool)`

GetOthersCanAddTasksOk returns a tuple with the OthersCanAddTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanAddTasks

`func (o *InputChecklist) SetOthersCanAddTasks(v bool)`

SetOthersCanAddTasks sets OthersCanAddTasks field to given value.

### HasOthersCanAddTasks

`func (o *InputChecklist) HasOthersCanAddTasks() bool`

HasOthersCanAddTasks returns a boolean if a field has been set.

### GetOthersCanMarkTasksAsDone

`func (o *InputChecklist) GetOthersCanMarkTasksAsDone() bool`

GetOthersCanMarkTasksAsDone returns the OthersCanMarkTasksAsDone field if non-nil, zero value otherwise.

### GetOthersCanMarkTasksAsDoneOk

`func (o *InputChecklist) GetOthersCanMarkTasksAsDoneOk() (*bool, bool)`

GetOthersCanMarkTasksAsDoneOk returns a tuple with the OthersCanMarkTasksAsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanMarkTasksAsDone

`func (o *InputChecklist) SetOthersCanMarkTasksAsDone(v bool)`

SetOthersCanMarkTasksAsDone sets OthersCanMarkTasksAsDone field to given value.

### HasOthersCanMarkTasksAsDone

`func (o *InputChecklist) HasOthersCanMarkTasksAsDone() bool`

HasOthersCanMarkTasksAsDone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


