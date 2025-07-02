# AnswerWebAppQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebAppQueryId** | **string** | Unique identifier for the query to be answered | 
**Result** | [**InlineQueryResult**](InlineQueryResult.md) |  | 

## Methods

### NewAnswerWebAppQueryRequest

`func NewAnswerWebAppQueryRequest(webAppQueryId string, result InlineQueryResult, ) *AnswerWebAppQueryRequest`

NewAnswerWebAppQueryRequest instantiates a new AnswerWebAppQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWebAppQueryRequestWithDefaults

`func NewAnswerWebAppQueryRequestWithDefaults() *AnswerWebAppQueryRequest`

NewAnswerWebAppQueryRequestWithDefaults instantiates a new AnswerWebAppQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebAppQueryId

`func (o *AnswerWebAppQueryRequest) GetWebAppQueryId() string`

GetWebAppQueryId returns the WebAppQueryId field if non-nil, zero value otherwise.

### GetWebAppQueryIdOk

`func (o *AnswerWebAppQueryRequest) GetWebAppQueryIdOk() (*string, bool)`

GetWebAppQueryIdOk returns a tuple with the WebAppQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppQueryId

`func (o *AnswerWebAppQueryRequest) SetWebAppQueryId(v string)`

SetWebAppQueryId sets WebAppQueryId field to given value.


### GetResult

`func (o *AnswerWebAppQueryRequest) GetResult() InlineQueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AnswerWebAppQueryRequest) GetResultOk() (*InlineQueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AnswerWebAppQueryRequest) SetResult(v InlineQueryResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


