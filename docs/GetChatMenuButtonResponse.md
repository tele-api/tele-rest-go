# GetChatMenuButtonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**MenuButton**](MenuButton.md) |  | 

## Methods

### NewGetChatMenuButtonResponse

`func NewGetChatMenuButtonResponse(ok bool, result MenuButton, ) *GetChatMenuButtonResponse`

NewGetChatMenuButtonResponse instantiates a new GetChatMenuButtonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChatMenuButtonResponseWithDefaults

`func NewGetChatMenuButtonResponseWithDefaults() *GetChatMenuButtonResponse`

NewGetChatMenuButtonResponseWithDefaults instantiates a new GetChatMenuButtonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetChatMenuButtonResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetChatMenuButtonResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetChatMenuButtonResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetChatMenuButtonResponse) GetResult() MenuButton`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetChatMenuButtonResponse) GetResultOk() (*MenuButton, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetChatMenuButtonResponse) SetResult(v MenuButton)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


