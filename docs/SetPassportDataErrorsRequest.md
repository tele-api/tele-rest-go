# SetPassportDataErrorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier | 
**Errors** | [**[]PassportElementError**](PassportElementError.md) | A JSON-serialized array describing the errors | 

## Methods

### NewSetPassportDataErrorsRequest

`func NewSetPassportDataErrorsRequest(userId int32, errors []PassportElementError, ) *SetPassportDataErrorsRequest`

NewSetPassportDataErrorsRequest instantiates a new SetPassportDataErrorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPassportDataErrorsRequestWithDefaults

`func NewSetPassportDataErrorsRequestWithDefaults() *SetPassportDataErrorsRequest`

NewSetPassportDataErrorsRequestWithDefaults instantiates a new SetPassportDataErrorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SetPassportDataErrorsRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SetPassportDataErrorsRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SetPassportDataErrorsRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetErrors

`func (o *SetPassportDataErrorsRequest) GetErrors() []PassportElementError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SetPassportDataErrorsRequest) GetErrorsOk() (*[]PassportElementError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SetPassportDataErrorsRequest) SetErrors(v []PassportElementError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


