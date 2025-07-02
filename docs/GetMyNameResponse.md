# GetMyNameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**BotName**](BotName.md) |  | 

## Methods

### NewGetMyNameResponse

`func NewGetMyNameResponse(ok bool, result BotName, ) *GetMyNameResponse`

NewGetMyNameResponse instantiates a new GetMyNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyNameResponseWithDefaults

`func NewGetMyNameResponseWithDefaults() *GetMyNameResponse`

NewGetMyNameResponseWithDefaults instantiates a new GetMyNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetMyNameResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetMyNameResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetMyNameResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetMyNameResponse) GetResult() BotName`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMyNameResponse) GetResultOk() (*BotName, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMyNameResponse) SetResult(v BotName)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


