# SendPaidMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**Message**](Message.md) |  | 

## Methods

### NewSendPaidMediaResponse

`func NewSendPaidMediaResponse(ok bool, result Message, ) *SendPaidMediaResponse`

NewSendPaidMediaResponse instantiates a new SendPaidMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPaidMediaResponseWithDefaults

`func NewSendPaidMediaResponseWithDefaults() *SendPaidMediaResponse`

NewSendPaidMediaResponseWithDefaults instantiates a new SendPaidMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *SendPaidMediaResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *SendPaidMediaResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *SendPaidMediaResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *SendPaidMediaResponse) GetResult() Message`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SendPaidMediaResponse) GetResultOk() (*Message, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SendPaidMediaResponse) SetResult(v Message)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


