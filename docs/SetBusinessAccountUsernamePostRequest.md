# SetBusinessAccountUsernamePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**Username** | Pointer to **string** | The new value of the username for the business account; 0-32 characters | [optional] 

## Methods

### NewSetBusinessAccountUsernamePostRequest

`func NewSetBusinessAccountUsernamePostRequest(businessConnectionId string, ) *SetBusinessAccountUsernamePostRequest`

NewSetBusinessAccountUsernamePostRequest instantiates a new SetBusinessAccountUsernamePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBusinessAccountUsernamePostRequestWithDefaults

`func NewSetBusinessAccountUsernamePostRequestWithDefaults() *SetBusinessAccountUsernamePostRequest`

NewSetBusinessAccountUsernamePostRequestWithDefaults instantiates a new SetBusinessAccountUsernamePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SetBusinessAccountUsernamePostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SetBusinessAccountUsernamePostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SetBusinessAccountUsernamePostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetUsername

`func (o *SetBusinessAccountUsernamePostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SetBusinessAccountUsernamePostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SetBusinessAccountUsernamePostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SetBusinessAccountUsernamePostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


