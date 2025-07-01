# ResponseParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrateToChatId** | Pointer to **int32** | *Optional*. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier. | [optional] 
**RetryAfter** | Pointer to **int32** | *Optional*. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated | [optional] 

## Methods

### NewResponseParameters

`func NewResponseParameters() *ResponseParameters`

NewResponseParameters instantiates a new ResponseParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseParametersWithDefaults

`func NewResponseParametersWithDefaults() *ResponseParameters`

NewResponseParametersWithDefaults instantiates a new ResponseParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrateToChatId

`func (o *ResponseParameters) GetMigrateToChatId() int32`

GetMigrateToChatId returns the MigrateToChatId field if non-nil, zero value otherwise.

### GetMigrateToChatIdOk

`func (o *ResponseParameters) GetMigrateToChatIdOk() (*int32, bool)`

GetMigrateToChatIdOk returns a tuple with the MigrateToChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateToChatId

`func (o *ResponseParameters) SetMigrateToChatId(v int32)`

SetMigrateToChatId sets MigrateToChatId field to given value.

### HasMigrateToChatId

`func (o *ResponseParameters) HasMigrateToChatId() bool`

HasMigrateToChatId returns a boolean if a field has been set.

### GetRetryAfter

`func (o *ResponseParameters) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *ResponseParameters) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *ResponseParameters) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *ResponseParameters) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


