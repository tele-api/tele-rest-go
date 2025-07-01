# UserProfilePhotos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total number of profile pictures the target user has | 
**Photos** | [**[][]PhotoSize**]([]PhotoSize.md) | Requested profile pictures (in up to 4 sizes each) | 

## Methods

### NewUserProfilePhotos

`func NewUserProfilePhotos(totalCount int32, photos [][]PhotoSize, ) *UserProfilePhotos`

NewUserProfilePhotos instantiates a new UserProfilePhotos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfilePhotosWithDefaults

`func NewUserProfilePhotosWithDefaults() *UserProfilePhotos`

NewUserProfilePhotosWithDefaults instantiates a new UserProfilePhotos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *UserProfilePhotos) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserProfilePhotos) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserProfilePhotos) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPhotos

`func (o *UserProfilePhotos) GetPhotos() [][]PhotoSize`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *UserProfilePhotos) GetPhotosOk() (*[][]PhotoSize, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *UserProfilePhotos) SetPhotos(v [][]PhotoSize)`

SetPhotos sets Photos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


