# Checklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the checklist | 
**TitleEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the checklist title | [optional] 
**Tasks** | [**[]ChecklistTask**](ChecklistTask.md) | List of tasks in the checklist | 
**OthersCanAddTasks** | Pointer to **bool** | *Optional*. *True*, if users other than the creator of the list can add tasks to the list | [optional] [default to true]
**OthersCanMarkTasksAsDone** | Pointer to **bool** | *Optional*. *True*, if users other than the creator of the list can mark tasks as done or not done | [optional] [default to true]

## Methods

### NewChecklist

`func NewChecklist(title string, tasks []ChecklistTask, ) *Checklist`

NewChecklist instantiates a new Checklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecklistWithDefaults

`func NewChecklistWithDefaults() *Checklist`

NewChecklistWithDefaults instantiates a new Checklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Checklist) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Checklist) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Checklist) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTitleEntities

`func (o *Checklist) GetTitleEntities() []MessageEntity`

GetTitleEntities returns the TitleEntities field if non-nil, zero value otherwise.

### GetTitleEntitiesOk

`func (o *Checklist) GetTitleEntitiesOk() (*[]MessageEntity, bool)`

GetTitleEntitiesOk returns a tuple with the TitleEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEntities

`func (o *Checklist) SetTitleEntities(v []MessageEntity)`

SetTitleEntities sets TitleEntities field to given value.

### HasTitleEntities

`func (o *Checklist) HasTitleEntities() bool`

HasTitleEntities returns a boolean if a field has been set.

### GetTasks

`func (o *Checklist) GetTasks() []ChecklistTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Checklist) GetTasksOk() (*[]ChecklistTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Checklist) SetTasks(v []ChecklistTask)`

SetTasks sets Tasks field to given value.


### GetOthersCanAddTasks

`func (o *Checklist) GetOthersCanAddTasks() bool`

GetOthersCanAddTasks returns the OthersCanAddTasks field if non-nil, zero value otherwise.

### GetOthersCanAddTasksOk

`func (o *Checklist) GetOthersCanAddTasksOk() (*bool, bool)`

GetOthersCanAddTasksOk returns a tuple with the OthersCanAddTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanAddTasks

`func (o *Checklist) SetOthersCanAddTasks(v bool)`

SetOthersCanAddTasks sets OthersCanAddTasks field to given value.

### HasOthersCanAddTasks

`func (o *Checklist) HasOthersCanAddTasks() bool`

HasOthersCanAddTasks returns a boolean if a field has been set.

### GetOthersCanMarkTasksAsDone

`func (o *Checklist) GetOthersCanMarkTasksAsDone() bool`

GetOthersCanMarkTasksAsDone returns the OthersCanMarkTasksAsDone field if non-nil, zero value otherwise.

### GetOthersCanMarkTasksAsDoneOk

`func (o *Checklist) GetOthersCanMarkTasksAsDoneOk() (*bool, bool)`

GetOthersCanMarkTasksAsDoneOk returns a tuple with the OthersCanMarkTasksAsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanMarkTasksAsDone

`func (o *Checklist) SetOthersCanMarkTasksAsDone(v bool)`

SetOthersCanMarkTasksAsDone sets OthersCanMarkTasksAsDone field to given value.

### HasOthersCanMarkTasksAsDone

`func (o *Checklist) HasOthersCanMarkTasksAsDone() bool`

HasOthersCanMarkTasksAsDone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


