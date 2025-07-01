# ChosenInlineResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultId** | **string** | The unique identifier for the result that was chosen | 
**From** | [**User**](User.md) |  | 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**InlineMessageId** | Pointer to **string** | *Optional*. Identifier of the sent inline message. Available only if there is an [inline keyboard](https://core.telegram.org/bots/api/#inlinekeyboardmarkup) attached to the message. Will be also received in [callback queries](https://core.telegram.org/bots/api/#callbackquery) and can be used to [edit](https://core.telegram.org/bots/api/#updating-messages) the message. | [optional] 
**Query** | **string** | The query that was used to obtain the result | 

## Methods

### NewChosenInlineResult

`func NewChosenInlineResult(resultId string, from User, query string, ) *ChosenInlineResult`

NewChosenInlineResult instantiates a new ChosenInlineResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChosenInlineResultWithDefaults

`func NewChosenInlineResultWithDefaults() *ChosenInlineResult`

NewChosenInlineResultWithDefaults instantiates a new ChosenInlineResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultId

`func (o *ChosenInlineResult) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *ChosenInlineResult) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *ChosenInlineResult) SetResultId(v string)`

SetResultId sets ResultId field to given value.


### GetFrom

`func (o *ChosenInlineResult) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChosenInlineResult) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChosenInlineResult) SetFrom(v User)`

SetFrom sets From field to given value.


### GetLocation

`func (o *ChosenInlineResult) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChosenInlineResult) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChosenInlineResult) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChosenInlineResult) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInlineMessageId

`func (o *ChosenInlineResult) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *ChosenInlineResult) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *ChosenInlineResult) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *ChosenInlineResult) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.

### GetQuery

`func (o *ChosenInlineResult) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ChosenInlineResult) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ChosenInlineResult) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


