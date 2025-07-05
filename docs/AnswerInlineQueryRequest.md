# AnswerInlineQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineQueryId** | **string** | Unique identifier for the answered query | 
**Results** | [**[]InlineQueryResult**](InlineQueryResult.md) | A JSON-serialized array of results for the inline query | 
**CacheTime** | Pointer to **int32** | The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300. | [optional] [default to 300]
**IsPersonal** | Pointer to **bool** | Pass *True* if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query. | [optional] 
**NextOffset** | Pointer to **string** | Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don&#39;t support pagination. Offset length can&#39;t exceed 64 bytes. | [optional] 
**Button** | Pointer to [**InlineQueryResultsButton**](InlineQueryResultsButton.md) |  | [optional] 

## Methods

### NewAnswerInlineQueryRequest

`func NewAnswerInlineQueryRequest(inlineQueryId string, results []InlineQueryResult, ) *AnswerInlineQueryRequest`

NewAnswerInlineQueryRequest instantiates a new AnswerInlineQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerInlineQueryRequestWithDefaults

`func NewAnswerInlineQueryRequestWithDefaults() *AnswerInlineQueryRequest`

NewAnswerInlineQueryRequestWithDefaults instantiates a new AnswerInlineQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineQueryId

`func (o *AnswerInlineQueryRequest) GetInlineQueryId() string`

GetInlineQueryId returns the InlineQueryId field if non-nil, zero value otherwise.

### GetInlineQueryIdOk

`func (o *AnswerInlineQueryRequest) GetInlineQueryIdOk() (*string, bool)`

GetInlineQueryIdOk returns a tuple with the InlineQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineQueryId

`func (o *AnswerInlineQueryRequest) SetInlineQueryId(v string)`

SetInlineQueryId sets InlineQueryId field to given value.


### GetResults

`func (o *AnswerInlineQueryRequest) GetResults() []InlineQueryResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AnswerInlineQueryRequest) GetResultsOk() (*[]InlineQueryResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AnswerInlineQueryRequest) SetResults(v []InlineQueryResult)`

SetResults sets Results field to given value.


### GetCacheTime

`func (o *AnswerInlineQueryRequest) GetCacheTime() int32`

GetCacheTime returns the CacheTime field if non-nil, zero value otherwise.

### GetCacheTimeOk

`func (o *AnswerInlineQueryRequest) GetCacheTimeOk() (*int32, bool)`

GetCacheTimeOk returns a tuple with the CacheTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTime

`func (o *AnswerInlineQueryRequest) SetCacheTime(v int32)`

SetCacheTime sets CacheTime field to given value.

### HasCacheTime

`func (o *AnswerInlineQueryRequest) HasCacheTime() bool`

HasCacheTime returns a boolean if a field has been set.

### GetIsPersonal

`func (o *AnswerInlineQueryRequest) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *AnswerInlineQueryRequest) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *AnswerInlineQueryRequest) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.

### HasIsPersonal

`func (o *AnswerInlineQueryRequest) HasIsPersonal() bool`

HasIsPersonal returns a boolean if a field has been set.

### GetNextOffset

`func (o *AnswerInlineQueryRequest) GetNextOffset() string`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *AnswerInlineQueryRequest) GetNextOffsetOk() (*string, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *AnswerInlineQueryRequest) SetNextOffset(v string)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *AnswerInlineQueryRequest) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetButton

`func (o *AnswerInlineQueryRequest) GetButton() InlineQueryResultsButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *AnswerInlineQueryRequest) GetButtonOk() (*InlineQueryResultsButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *AnswerInlineQueryRequest) SetButton(v InlineQueryResultsButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *AnswerInlineQueryRequest) HasButton() bool`

HasButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


