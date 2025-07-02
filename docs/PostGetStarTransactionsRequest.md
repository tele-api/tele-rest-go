# PostGetStarTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Number of transactions to skip in the response | [optional] 
**Limit** | Pointer to **int32** | The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100. | [optional] [default to 100]

## Methods

### NewPostGetStarTransactionsRequest

`func NewPostGetStarTransactionsRequest() *PostGetStarTransactionsRequest`

NewPostGetStarTransactionsRequest instantiates a new PostGetStarTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetStarTransactionsRequestWithDefaults

`func NewPostGetStarTransactionsRequestWithDefaults() *PostGetStarTransactionsRequest`

NewPostGetStarTransactionsRequestWithDefaults instantiates a new PostGetStarTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *PostGetStarTransactionsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PostGetStarTransactionsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PostGetStarTransactionsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PostGetStarTransactionsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PostGetStarTransactionsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostGetStarTransactionsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostGetStarTransactionsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostGetStarTransactionsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


