# ChecklistTasksDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChecklistMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**MarkedAsDoneTaskIds** | Pointer to **[]int32** | *Optional*. Identifiers of the tasks that were marked as done | [optional] 
**MarkedAsNotDoneTaskIds** | Pointer to **[]int32** | *Optional*. Identifiers of the tasks that were marked as not done | [optional] 

## Methods

### NewChecklistTasksDone

`func NewChecklistTasksDone() *ChecklistTasksDone`

NewChecklistTasksDone instantiates a new ChecklistTasksDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecklistTasksDoneWithDefaults

`func NewChecklistTasksDoneWithDefaults() *ChecklistTasksDone`

NewChecklistTasksDoneWithDefaults instantiates a new ChecklistTasksDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecklistMessage

`func (o *ChecklistTasksDone) GetChecklistMessage() Message`

GetChecklistMessage returns the ChecklistMessage field if non-nil, zero value otherwise.

### GetChecklistMessageOk

`func (o *ChecklistTasksDone) GetChecklistMessageOk() (*Message, bool)`

GetChecklistMessageOk returns a tuple with the ChecklistMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistMessage

`func (o *ChecklistTasksDone) SetChecklistMessage(v Message)`

SetChecklistMessage sets ChecklistMessage field to given value.

### HasChecklistMessage

`func (o *ChecklistTasksDone) HasChecklistMessage() bool`

HasChecklistMessage returns a boolean if a field has been set.

### GetMarkedAsDoneTaskIds

`func (o *ChecklistTasksDone) GetMarkedAsDoneTaskIds() []int32`

GetMarkedAsDoneTaskIds returns the MarkedAsDoneTaskIds field if non-nil, zero value otherwise.

### GetMarkedAsDoneTaskIdsOk

`func (o *ChecklistTasksDone) GetMarkedAsDoneTaskIdsOk() (*[]int32, bool)`

GetMarkedAsDoneTaskIdsOk returns a tuple with the MarkedAsDoneTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedAsDoneTaskIds

`func (o *ChecklistTasksDone) SetMarkedAsDoneTaskIds(v []int32)`

SetMarkedAsDoneTaskIds sets MarkedAsDoneTaskIds field to given value.

### HasMarkedAsDoneTaskIds

`func (o *ChecklistTasksDone) HasMarkedAsDoneTaskIds() bool`

HasMarkedAsDoneTaskIds returns a boolean if a field has been set.

### GetMarkedAsNotDoneTaskIds

`func (o *ChecklistTasksDone) GetMarkedAsNotDoneTaskIds() []int32`

GetMarkedAsNotDoneTaskIds returns the MarkedAsNotDoneTaskIds field if non-nil, zero value otherwise.

### GetMarkedAsNotDoneTaskIdsOk

`func (o *ChecklistTasksDone) GetMarkedAsNotDoneTaskIdsOk() (*[]int32, bool)`

GetMarkedAsNotDoneTaskIdsOk returns a tuple with the MarkedAsNotDoneTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedAsNotDoneTaskIds

`func (o *ChecklistTasksDone) SetMarkedAsNotDoneTaskIds(v []int32)`

SetMarkedAsNotDoneTaskIds sets MarkedAsNotDoneTaskIds field to given value.

### HasMarkedAsNotDoneTaskIds

`func (o *ChecklistTasksDone) HasMarkedAsNotDoneTaskIds() bool`

HasMarkedAsNotDoneTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


