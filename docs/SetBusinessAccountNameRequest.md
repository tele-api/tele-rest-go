# SetBusinessAccountNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**FirstName** | **string** | The new value of the first name for the business account; 1-64 characters | 
**LastName** | Pointer to **string** | The new value of the last name for the business account; 0-64 characters | [optional] 

## Methods

### NewSetBusinessAccountNameRequest

`func NewSetBusinessAccountNameRequest(businessConnectionId string, firstName string, ) *SetBusinessAccountNameRequest`

NewSetBusinessAccountNameRequest instantiates a new SetBusinessAccountNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBusinessAccountNameRequestWithDefaults

`func NewSetBusinessAccountNameRequestWithDefaults() *SetBusinessAccountNameRequest`

NewSetBusinessAccountNameRequestWithDefaults instantiates a new SetBusinessAccountNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SetBusinessAccountNameRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SetBusinessAccountNameRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SetBusinessAccountNameRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetFirstName

`func (o *SetBusinessAccountNameRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SetBusinessAccountNameRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SetBusinessAccountNameRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SetBusinessAccountNameRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SetBusinessAccountNameRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SetBusinessAccountNameRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SetBusinessAccountNameRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


