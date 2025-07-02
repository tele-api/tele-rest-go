# PostGetMyCommands200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**[]BotCommand**](BotCommand.md) |  | 

## Methods

### NewPostGetMyCommands200Response

`func NewPostGetMyCommands200Response(ok bool, result []BotCommand, ) *PostGetMyCommands200Response`

NewPostGetMyCommands200Response instantiates a new PostGetMyCommands200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetMyCommands200ResponseWithDefaults

`func NewPostGetMyCommands200ResponseWithDefaults() *PostGetMyCommands200Response`

NewPostGetMyCommands200ResponseWithDefaults instantiates a new PostGetMyCommands200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *PostGetMyCommands200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PostGetMyCommands200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PostGetMyCommands200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *PostGetMyCommands200Response) GetResult() []BotCommand`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostGetMyCommands200Response) GetResultOk() (*[]BotCommand, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostGetMyCommands200Response) SetResult(v []BotCommand)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


