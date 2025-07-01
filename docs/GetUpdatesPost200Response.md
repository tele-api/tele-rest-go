# GetUpdatesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**[]Update**](Update.md) |  | 

## Methods

### NewGetUpdatesPost200Response

`func NewGetUpdatesPost200Response(ok bool, result []Update, ) *GetUpdatesPost200Response`

NewGetUpdatesPost200Response instantiates a new GetUpdatesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUpdatesPost200ResponseWithDefaults

`func NewGetUpdatesPost200ResponseWithDefaults() *GetUpdatesPost200Response`

NewGetUpdatesPost200ResponseWithDefaults instantiates a new GetUpdatesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GetUpdatesPost200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetUpdatesPost200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetUpdatesPost200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *GetUpdatesPost200Response) GetResult() []Update`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUpdatesPost200Response) GetResultOk() (*[]Update, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUpdatesPost200Response) SetResult(v []Update)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


