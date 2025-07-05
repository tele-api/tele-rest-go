# PostStoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**Story**](Story.md) |  | 

## Methods

### NewPostStoryResponse

`func NewPostStoryResponse(ok bool, result Story, ) *PostStoryResponse`

NewPostStoryResponse instantiates a new PostStoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStoryResponseWithDefaults

`func NewPostStoryResponseWithDefaults() *PostStoryResponse`

NewPostStoryResponseWithDefaults instantiates a new PostStoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *PostStoryResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PostStoryResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PostStoryResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *PostStoryResponse) GetResult() Story`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostStoryResponse) GetResultOk() (*Story, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostStoryResponse) SetResult(v Story)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


