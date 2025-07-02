# SetGameScoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**EditMessageTextResponseResult**](EditMessageTextResponseResult.md) |  | 

## Methods

### NewSetGameScoreResponse

`func NewSetGameScoreResponse(ok bool, result EditMessageTextResponseResult, ) *SetGameScoreResponse`

NewSetGameScoreResponse instantiates a new SetGameScoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGameScoreResponseWithDefaults

`func NewSetGameScoreResponseWithDefaults() *SetGameScoreResponse`

NewSetGameScoreResponseWithDefaults instantiates a new SetGameScoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *SetGameScoreResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *SetGameScoreResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *SetGameScoreResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *SetGameScoreResponse) GetResult() EditMessageTextResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SetGameScoreResponse) GetResultOk() (*EditMessageTextResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SetGameScoreResponse) SetResult(v EditMessageTextResponseResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


