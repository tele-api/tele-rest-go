# ChecklistTasksAdded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChecklistMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Tasks** | [**[]ChecklistTask**](ChecklistTask.md) | List of tasks added to the checklist | 

## Methods

### NewChecklistTasksAdded

`func NewChecklistTasksAdded(tasks []ChecklistTask, ) *ChecklistTasksAdded`

NewChecklistTasksAdded instantiates a new ChecklistTasksAdded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecklistTasksAddedWithDefaults

`func NewChecklistTasksAddedWithDefaults() *ChecklistTasksAdded`

NewChecklistTasksAddedWithDefaults instantiates a new ChecklistTasksAdded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecklistMessage

`func (o *ChecklistTasksAdded) GetChecklistMessage() Message`

GetChecklistMessage returns the ChecklistMessage field if non-nil, zero value otherwise.

### GetChecklistMessageOk

`func (o *ChecklistTasksAdded) GetChecklistMessageOk() (*Message, bool)`

GetChecklistMessageOk returns a tuple with the ChecklistMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistMessage

`func (o *ChecklistTasksAdded) SetChecklistMessage(v Message)`

SetChecklistMessage sets ChecklistMessage field to given value.

### HasChecklistMessage

`func (o *ChecklistTasksAdded) HasChecklistMessage() bool`

HasChecklistMessage returns a boolean if a field has been set.

### GetTasks

`func (o *ChecklistTasksAdded) GetTasks() []ChecklistTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ChecklistTasksAdded) GetTasksOk() (*[]ChecklistTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ChecklistTasksAdded) SetTasks(v []ChecklistTask)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


