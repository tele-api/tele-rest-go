# AnswerWebAppQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**SentWebAppMessage**](SentWebAppMessage.md) |  | 

## Methods

### NewAnswerWebAppQueryResponse

`func NewAnswerWebAppQueryResponse(ok bool, result SentWebAppMessage, ) *AnswerWebAppQueryResponse`

NewAnswerWebAppQueryResponse instantiates a new AnswerWebAppQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWebAppQueryResponseWithDefaults

`func NewAnswerWebAppQueryResponseWithDefaults() *AnswerWebAppQueryResponse`

NewAnswerWebAppQueryResponseWithDefaults instantiates a new AnswerWebAppQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *AnswerWebAppQueryResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AnswerWebAppQueryResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AnswerWebAppQueryResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *AnswerWebAppQueryResponse) GetResult() SentWebAppMessage`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AnswerWebAppQueryResponse) GetResultOk() (*SentWebAppMessage, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AnswerWebAppQueryResponse) SetResult(v SentWebAppMessage)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


