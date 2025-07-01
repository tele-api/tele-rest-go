# SetBusinessAccountBioPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**Bio** | Pointer to **string** | The new value of the bio for the business account; 0-140 characters | [optional] 

## Methods

### NewSetBusinessAccountBioPostRequest

`func NewSetBusinessAccountBioPostRequest(businessConnectionId string, ) *SetBusinessAccountBioPostRequest`

NewSetBusinessAccountBioPostRequest instantiates a new SetBusinessAccountBioPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBusinessAccountBioPostRequestWithDefaults

`func NewSetBusinessAccountBioPostRequestWithDefaults() *SetBusinessAccountBioPostRequest`

NewSetBusinessAccountBioPostRequestWithDefaults instantiates a new SetBusinessAccountBioPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SetBusinessAccountBioPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SetBusinessAccountBioPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SetBusinessAccountBioPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetBio

`func (o *SetBusinessAccountBioPostRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *SetBusinessAccountBioPostRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *SetBusinessAccountBioPostRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *SetBusinessAccountBioPostRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


