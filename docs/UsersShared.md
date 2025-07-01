# UsersShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int32** | Identifier of the request | 
**Users** | [**[]SharedUser**](SharedUser.md) | Information about users shared with the bot. | 

## Methods

### NewUsersShared

`func NewUsersShared(requestId int32, users []SharedUser, ) *UsersShared`

NewUsersShared instantiates a new UsersShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersSharedWithDefaults

`func NewUsersSharedWithDefaults() *UsersShared`

NewUsersSharedWithDefaults instantiates a new UsersShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UsersShared) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UsersShared) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UsersShared) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.


### GetUsers

`func (o *UsersShared) GetUsers() []SharedUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersShared) GetUsersOk() (*[]SharedUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersShared) SetUsers(v []SharedUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


