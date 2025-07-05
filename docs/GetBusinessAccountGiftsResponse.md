# GetBusinessAccountGiftsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**OwnedGifts**](OwnedGifts.md) |  | 

## Methods

### NewGetBusinessAccountGiftsResponse

`func NewGetBusinessAccountGiftsResponse(ok bool, result OwnedGifts, ) *GetBusinessAccountGiftsResponse`

NewGetBusinessAccountGiftsResponse instantiates a new GetBusinessAccountGiftsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessAccountGiftsResponseWithDefaults

`func NewGetBusinessAccountGiftsResponseWithDefaults() *GetBusinessAccountGiftsResponse`

NewGetBusinessAccountGiftsResponseWithDefaults instantiates a new GetBusinessAccountGiftsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetBusinessAccountGiftsResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetBusinessAccountGiftsResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetBusinessAccountGiftsResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetBusinessAccountGiftsResponse) GetResult() OwnedGifts`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBusinessAccountGiftsResponse) GetResultOk() (*OwnedGifts, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBusinessAccountGiftsResponse) SetResult(v OwnedGifts)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


