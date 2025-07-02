# PostGetUpdatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as [getUpdates](https://core.telegram.org/bots/api/#getupdates) is called with an *offset* higher than its *update\\_id*. The negative offset can be specified to retrieve updates starting from *-offset* update from the end of the updates queue. All previous updates will be forgotten. | [optional] 
**Limit** | Pointer to **int32** | Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100. | [optional] [default to 100]
**Timeout** | Pointer to **int32** | Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only. | [optional] [default to 0]
**AllowedUpdates** | Pointer to **[]string** | A JSON-serialized list of the update types you want your bot to receive. For example, specify &#x60;[\&quot;message\&quot;, \&quot;edited_channel_post\&quot;, \&quot;callback_query\&quot;]&#x60; to only receive updates of these types. See [Update](https://core.telegram.org/bots/api/#update) for a complete list of available update types. Specify an empty list to receive all update types except *chat\\_member*, *message\\_reaction*, and *message\\_reaction\\_count* (default). If not specified, the previous setting will be used.    Please note that this parameter doesn&#39;t affect updates created before the call to getUpdates, so unwanted updates may be received for a short period of time. | [optional] 

## Methods

### NewPostGetUpdatesRequest

`func NewPostGetUpdatesRequest() *PostGetUpdatesRequest`

NewPostGetUpdatesRequest instantiates a new PostGetUpdatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetUpdatesRequestWithDefaults

`func NewPostGetUpdatesRequestWithDefaults() *PostGetUpdatesRequest`

NewPostGetUpdatesRequestWithDefaults instantiates a new PostGetUpdatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *PostGetUpdatesRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PostGetUpdatesRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PostGetUpdatesRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PostGetUpdatesRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PostGetUpdatesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostGetUpdatesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostGetUpdatesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostGetUpdatesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTimeout

`func (o *PostGetUpdatesRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PostGetUpdatesRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PostGetUpdatesRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PostGetUpdatesRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAllowedUpdates

`func (o *PostGetUpdatesRequest) GetAllowedUpdates() []string`

GetAllowedUpdates returns the AllowedUpdates field if non-nil, zero value otherwise.

### GetAllowedUpdatesOk

`func (o *PostGetUpdatesRequest) GetAllowedUpdatesOk() (*[]string, bool)`

GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUpdates

`func (o *PostGetUpdatesRequest) SetAllowedUpdates(v []string)`

SetAllowedUpdates sets AllowedUpdates field to given value.

### HasAllowedUpdates

`func (o *PostGetUpdatesRequest) HasAllowedUpdates() bool`

HasAllowedUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


