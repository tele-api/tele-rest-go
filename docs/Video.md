# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**Width** | **int32** | Video width as defined by the sender | 
**Height** | **int32** | Video height as defined by the sender | 
**Duration** | **int32** | Duration of the video in seconds as defined by the sender | 
**Thumbnail** | Pointer to [**PhotoSize**](PhotoSize.md) |  | [optional] 
**Cover** | Pointer to [**[]PhotoSize**](PhotoSize.md) | *Optional*. Available sizes of the cover of the video in the message | [optional] 
**StartTimestamp** | Pointer to **int32** | *Optional*. Timestamp in seconds from which the video will play in the message | [optional] 
**FileName** | Pointer to **string** | *Optional*. Original filename as defined by the sender | [optional] 
**MimeType** | Pointer to **string** | *Optional*. MIME type of the file as defined by the sender | [optional] 
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value. | [optional] 

## Methods

### NewVideo

`func NewVideo(fileId string, fileUniqueId string, width int32, height int32, duration int32, ) *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *Video) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Video) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Video) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *Video) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *Video) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *Video) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetWidth

`func (o *Video) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Video) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Video) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Video) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Video) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Video) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDuration

`func (o *Video) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Video) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Video) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetThumbnail

`func (o *Video) GetThumbnail() PhotoSize`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Video) GetThumbnailOk() (*PhotoSize, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Video) SetThumbnail(v PhotoSize)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Video) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCover

`func (o *Video) GetCover() []PhotoSize`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *Video) GetCoverOk() (*[]PhotoSize, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *Video) SetCover(v []PhotoSize)`

SetCover sets Cover field to given value.

### HasCover

`func (o *Video) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *Video) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Video) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Video) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *Video) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetFileName

`func (o *Video) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Video) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Video) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Video) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMimeType

`func (o *Video) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Video) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Video) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Video) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileSize

`func (o *Video) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Video) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Video) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Video) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


