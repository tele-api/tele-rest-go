# PhotoSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**Width** | **int32** | Photo width | 
**Height** | **int32** | Photo height | 
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes | [optional] 

## Methods

### NewPhotoSize

`func NewPhotoSize(fileId string, fileUniqueId string, width int32, height int32, ) *PhotoSize`

NewPhotoSize instantiates a new PhotoSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoSizeWithDefaults

`func NewPhotoSizeWithDefaults() *PhotoSize`

NewPhotoSizeWithDefaults instantiates a new PhotoSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *PhotoSize) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PhotoSize) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PhotoSize) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *PhotoSize) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *PhotoSize) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *PhotoSize) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetWidth

`func (o *PhotoSize) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PhotoSize) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PhotoSize) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *PhotoSize) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PhotoSize) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PhotoSize) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetFileSize

`func (o *PhotoSize) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *PhotoSize) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *PhotoSize) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *PhotoSize) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


