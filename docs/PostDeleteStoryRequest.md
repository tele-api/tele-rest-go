# PostDeleteStoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**StoryId** | **int32** | Unique identifier of the story to delete | 

## Methods

### NewPostDeleteStoryRequest

`func NewPostDeleteStoryRequest(businessConnectionId string, storyId int32, ) *PostDeleteStoryRequest`

NewPostDeleteStoryRequest instantiates a new PostDeleteStoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteStoryRequestWithDefaults

`func NewPostDeleteStoryRequestWithDefaults() *PostDeleteStoryRequest`

NewPostDeleteStoryRequestWithDefaults instantiates a new PostDeleteStoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostDeleteStoryRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostDeleteStoryRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostDeleteStoryRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetStoryId

`func (o *PostDeleteStoryRequest) GetStoryId() int32`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *PostDeleteStoryRequest) GetStoryIdOk() (*int32, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *PostDeleteStoryRequest) SetStoryId(v int32)`

SetStoryId sets StoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


