# VideoNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**Length** | **int32** | Video width and height (diameter of the video message) as defined by the sender | 
**Duration** | **int32** | Duration of the video in seconds as defined by the sender | 
**Thumbnail** | Pointer to [**PhotoSize**](PhotoSize.md) |  | [optional] 
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes | [optional] 

## Methods

### NewVideoNote

`func NewVideoNote(fileId string, fileUniqueId string, length int32, duration int32, ) *VideoNote`

NewVideoNote instantiates a new VideoNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoNoteWithDefaults

`func NewVideoNoteWithDefaults() *VideoNote`

NewVideoNoteWithDefaults instantiates a new VideoNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *VideoNote) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *VideoNote) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *VideoNote) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *VideoNote) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *VideoNote) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *VideoNote) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetLength

`func (o *VideoNote) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VideoNote) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VideoNote) SetLength(v int32)`

SetLength sets Length field to given value.


### GetDuration

`func (o *VideoNote) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoNote) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoNote) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetThumbnail

`func (o *VideoNote) GetThumbnail() PhotoSize`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *VideoNote) GetThumbnailOk() (*PhotoSize, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *VideoNote) SetThumbnail(v PhotoSize)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *VideoNote) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetFileSize

`func (o *VideoNote) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *VideoNote) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *VideoNote) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *VideoNote) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


