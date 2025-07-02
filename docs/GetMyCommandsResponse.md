# GetMyCommandsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**[]BotCommand**](BotCommand.md) |  | 

## Methods

### NewGetMyCommandsResponse

`func NewGetMyCommandsResponse(ok bool, result []BotCommand, ) *GetMyCommandsResponse`

NewGetMyCommandsResponse instantiates a new GetMyCommandsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyCommandsResponseWithDefaults

`func NewGetMyCommandsResponseWithDefaults() *GetMyCommandsResponse`

NewGetMyCommandsResponseWithDefaults instantiates a new GetMyCommandsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetMyCommandsResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetMyCommandsResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetMyCommandsResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetMyCommandsResponse) GetResult() []BotCommand`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMyCommandsResponse) GetResultOk() (*[]BotCommand, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMyCommandsResponse) SetResult(v []BotCommand)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


