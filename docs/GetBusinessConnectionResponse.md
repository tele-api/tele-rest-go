# GetBusinessConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**BusinessConnection**](BusinessConnection.md) |  | 

## Methods

### NewGetBusinessConnectionResponse

`func NewGetBusinessConnectionResponse(ok bool, result BusinessConnection, ) *GetBusinessConnectionResponse`

NewGetBusinessConnectionResponse instantiates a new GetBusinessConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessConnectionResponseWithDefaults

`func NewGetBusinessConnectionResponseWithDefaults() *GetBusinessConnectionResponse`

NewGetBusinessConnectionResponseWithDefaults instantiates a new GetBusinessConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetBusinessConnectionResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetBusinessConnectionResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetBusinessConnectionResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetBusinessConnectionResponse) GetResult() BusinessConnection`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBusinessConnectionResponse) GetResultOk() (*BusinessConnection, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBusinessConnectionResponse) SetResult(v BusinessConnection)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


