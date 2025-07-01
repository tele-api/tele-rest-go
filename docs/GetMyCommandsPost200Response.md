# GetMyCommandsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**[]BotCommand**](BotCommand.md) |  | 

## Methods

### NewGetMyCommandsPost200Response

`func NewGetMyCommandsPost200Response(ok bool, result []BotCommand, ) *GetMyCommandsPost200Response`

NewGetMyCommandsPost200Response instantiates a new GetMyCommandsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyCommandsPost200ResponseWithDefaults

`func NewGetMyCommandsPost200ResponseWithDefaults() *GetMyCommandsPost200Response`

NewGetMyCommandsPost200ResponseWithDefaults instantiates a new GetMyCommandsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetMyCommandsPost200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetMyCommandsPost200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetMyCommandsPost200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetMyCommandsPost200Response) GetResult() []BotCommand`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMyCommandsPost200Response) GetResultOk() (*[]BotCommand, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMyCommandsPost200Response) SetResult(v []BotCommand)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


