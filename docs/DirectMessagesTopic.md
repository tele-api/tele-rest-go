# DirectMessagesTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicId** | **int32** | Unique identifier of the topic | 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewDirectMessagesTopic

`func NewDirectMessagesTopic(topicId int32, ) *DirectMessagesTopic`

NewDirectMessagesTopic instantiates a new DirectMessagesTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectMessagesTopicWithDefaults

`func NewDirectMessagesTopicWithDefaults() *DirectMessagesTopic`

NewDirectMessagesTopicWithDefaults instantiates a new DirectMessagesTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicId

`func (o *DirectMessagesTopic) GetTopicId() int32`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *DirectMessagesTopic) GetTopicIdOk() (*int32, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *DirectMessagesTopic) SetTopicId(v int32)`

SetTopicId sets TopicId field to given value.


### GetUser

`func (o *DirectMessagesTopic) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DirectMessagesTopic) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DirectMessagesTopic) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *DirectMessagesTopic) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


