# EditMessageMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**EditMessageTextResponseResult**](EditMessageTextResponseResult.md) |  | 

## Methods

### NewEditMessageMediaResponse

`func NewEditMessageMediaResponse(ok bool, result EditMessageTextResponseResult, ) *EditMessageMediaResponse`

NewEditMessageMediaResponse instantiates a new EditMessageMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMessageMediaResponseWithDefaults

`func NewEditMessageMediaResponseWithDefaults() *EditMessageMediaResponse`

NewEditMessageMediaResponseWithDefaults instantiates a new EditMessageMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *EditMessageMediaResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *EditMessageMediaResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *EditMessageMediaResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *EditMessageMediaResponse) GetResult() EditMessageTextResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EditMessageMediaResponse) GetResultOk() (*EditMessageTextResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EditMessageMediaResponse) SetResult(v EditMessageTextResponseResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


