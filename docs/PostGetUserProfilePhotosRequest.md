# PostGetUserProfilePhotosRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user | 
**Offset** | Pointer to **int32** | Sequential number of the first photo to be returned. By default, all photos are returned. | [optional] 
**Limit** | Pointer to **int32** | Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100. | [optional] [default to 100]

## Methods

### NewPostGetUserProfilePhotosRequest

`func NewPostGetUserProfilePhotosRequest(userId int32, ) *PostGetUserProfilePhotosRequest`

NewPostGetUserProfilePhotosRequest instantiates a new PostGetUserProfilePhotosRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetUserProfilePhotosRequestWithDefaults

`func NewPostGetUserProfilePhotosRequestWithDefaults() *PostGetUserProfilePhotosRequest`

NewPostGetUserProfilePhotosRequestWithDefaults instantiates a new PostGetUserProfilePhotosRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostGetUserProfilePhotosRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostGetUserProfilePhotosRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostGetUserProfilePhotosRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetOffset

`func (o *PostGetUserProfilePhotosRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PostGetUserProfilePhotosRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PostGetUserProfilePhotosRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PostGetUserProfilePhotosRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PostGetUserProfilePhotosRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostGetUserProfilePhotosRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostGetUserProfilePhotosRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PostGetUserProfilePhotosRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


