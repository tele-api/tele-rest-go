# StopPollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**Poll**](Poll.md) |  | 

## Methods

### NewStopPollResponse

`func NewStopPollResponse(ok bool, result Poll, ) *StopPollResponse`

NewStopPollResponse instantiates a new StopPollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPollResponseWithDefaults

`func NewStopPollResponseWithDefaults() *StopPollResponse`

NewStopPollResponseWithDefaults instantiates a new StopPollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *StopPollResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *StopPollResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *StopPollResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *StopPollResponse) GetResult() Poll`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StopPollResponse) GetResultOk() (*Poll, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StopPollResponse) SetResult(v Poll)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


