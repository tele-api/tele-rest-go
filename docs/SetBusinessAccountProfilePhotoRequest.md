# SetBusinessAccountProfilePhotoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**Photo** | [**InputProfilePhoto**](InputProfilePhoto.md) |  | 
**IsPublic** | Pointer to **bool** | Pass True to set the public photo, which will be visible even if the main photo is hidden by the business account&#39;s privacy settings. An account can have only one public photo. | [optional] 

## Methods

### NewSetBusinessAccountProfilePhotoRequest

`func NewSetBusinessAccountProfilePhotoRequest(businessConnectionId string, photo InputProfilePhoto, ) *SetBusinessAccountProfilePhotoRequest`

NewSetBusinessAccountProfilePhotoRequest instantiates a new SetBusinessAccountProfilePhotoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBusinessAccountProfilePhotoRequestWithDefaults

`func NewSetBusinessAccountProfilePhotoRequestWithDefaults() *SetBusinessAccountProfilePhotoRequest`

NewSetBusinessAccountProfilePhotoRequestWithDefaults instantiates a new SetBusinessAccountProfilePhotoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SetBusinessAccountProfilePhotoRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SetBusinessAccountProfilePhotoRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SetBusinessAccountProfilePhotoRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetPhoto

`func (o *SetBusinessAccountProfilePhotoRequest) GetPhoto() InputProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *SetBusinessAccountProfilePhotoRequest) GetPhotoOk() (*InputProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *SetBusinessAccountProfilePhotoRequest) SetPhoto(v InputProfilePhoto)`

SetPhoto sets Photo field to given value.


### GetIsPublic

`func (o *SetBusinessAccountProfilePhotoRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SetBusinessAccountProfilePhotoRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SetBusinessAccountProfilePhotoRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *SetBusinessAccountProfilePhotoRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


