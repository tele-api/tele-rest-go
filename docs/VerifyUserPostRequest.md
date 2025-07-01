# VerifyUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user | 
**CustomDescription** | Pointer to **string** | Custom description for the verification; 0-70 characters. Must be empty if the organization isn&#39;t allowed to provide a custom verification description. | [optional] 

## Methods

### NewVerifyUserPostRequest

`func NewVerifyUserPostRequest(userId int32, ) *VerifyUserPostRequest`

NewVerifyUserPostRequest instantiates a new VerifyUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyUserPostRequestWithDefaults

`func NewVerifyUserPostRequestWithDefaults() *VerifyUserPostRequest`

NewVerifyUserPostRequestWithDefaults instantiates a new VerifyUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *VerifyUserPostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VerifyUserPostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VerifyUserPostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCustomDescription

`func (o *VerifyUserPostRequest) GetCustomDescription() string`

GetCustomDescription returns the CustomDescription field if non-nil, zero value otherwise.

### GetCustomDescriptionOk

`func (o *VerifyUserPostRequest) GetCustomDescriptionOk() (*string, bool)`

GetCustomDescriptionOk returns a tuple with the CustomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDescription

`func (o *VerifyUserPostRequest) SetCustomDescription(v string)`

SetCustomDescription sets CustomDescription field to given value.

### HasCustomDescription

`func (o *VerifyUserPostRequest) HasCustomDescription() bool`

HasCustomDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


