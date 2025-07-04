# RemoveBusinessAccountProfilePhotoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**IsPublic** | Pointer to **bool** | Pass *True* to remove the public photo, which is visible even if the main photo is hidden by the business account&#39;s privacy settings. After the main photo is removed, the previous profile photo (if present) becomes the main photo. | [optional] 

## Methods

### NewRemoveBusinessAccountProfilePhotoRequest

`func NewRemoveBusinessAccountProfilePhotoRequest(businessConnectionId string, ) *RemoveBusinessAccountProfilePhotoRequest`

NewRemoveBusinessAccountProfilePhotoRequest instantiates a new RemoveBusinessAccountProfilePhotoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveBusinessAccountProfilePhotoRequestWithDefaults

`func NewRemoveBusinessAccountProfilePhotoRequestWithDefaults() *RemoveBusinessAccountProfilePhotoRequest`

NewRemoveBusinessAccountProfilePhotoRequestWithDefaults instantiates a new RemoveBusinessAccountProfilePhotoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *RemoveBusinessAccountProfilePhotoRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *RemoveBusinessAccountProfilePhotoRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *RemoveBusinessAccountProfilePhotoRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetIsPublic

`func (o *RemoveBusinessAccountProfilePhotoRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *RemoveBusinessAccountProfilePhotoRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *RemoveBusinessAccountProfilePhotoRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *RemoveBusinessAccountProfilePhotoRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


