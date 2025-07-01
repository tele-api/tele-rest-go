# AnswerWebAppQueryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebAppQueryId** | **string** | Unique identifier for the query to be answered | 
**Result** | [**InlineQueryResult**](InlineQueryResult.md) |  | 

## Methods

### NewAnswerWebAppQueryPostRequest

`func NewAnswerWebAppQueryPostRequest(webAppQueryId string, result InlineQueryResult, ) *AnswerWebAppQueryPostRequest`

NewAnswerWebAppQueryPostRequest instantiates a new AnswerWebAppQueryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWebAppQueryPostRequestWithDefaults

`func NewAnswerWebAppQueryPostRequestWithDefaults() *AnswerWebAppQueryPostRequest`

NewAnswerWebAppQueryPostRequestWithDefaults instantiates a new AnswerWebAppQueryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebAppQueryId

`func (o *AnswerWebAppQueryPostRequest) GetWebAppQueryId() string`

GetWebAppQueryId returns the WebAppQueryId field if non-nil, zero value otherwise.

### GetWebAppQueryIdOk

`func (o *AnswerWebAppQueryPostRequest) GetWebAppQueryIdOk() (*string, bool)`

GetWebAppQueryIdOk returns a tuple with the WebAppQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppQueryId

`func (o *AnswerWebAppQueryPostRequest) SetWebAppQueryId(v string)`

SetWebAppQueryId sets WebAppQueryId field to given value.


### GetResult

`func (o *AnswerWebAppQueryPostRequest) GetResult() InlineQueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AnswerWebAppQueryPostRequest) GetResultOk() (*InlineQueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AnswerWebAppQueryPostRequest) SetResult(v InlineQueryResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


