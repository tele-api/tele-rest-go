# PostCopyMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**MessageId**](MessageId.md) |  | 

## Methods

### NewPostCopyMessage200Response

`func NewPostCopyMessage200Response(ok bool, result MessageId, ) *PostCopyMessage200Response`

NewPostCopyMessage200Response instantiates a new PostCopyMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCopyMessage200ResponseWithDefaults

`func NewPostCopyMessage200ResponseWithDefaults() *PostCopyMessage200Response`

NewPostCopyMessage200ResponseWithDefaults instantiates a new PostCopyMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *PostCopyMessage200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PostCopyMessage200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PostCopyMessage200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *PostCopyMessage200Response) GetResult() MessageId`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostCopyMessage200Response) GetResultOk() (*MessageId, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostCopyMessage200Response) SetResult(v MessageId)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


