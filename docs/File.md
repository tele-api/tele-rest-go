# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value. | [optional] 
**FilePath** | Pointer to **string** | *Optional*. File path. Use &#x60;https://api.telegram.org/file/bot&lt;token&gt;/&lt;file_path&gt;&#x60; to get the file. | [optional] 

## Methods

### NewFile

`func NewFile(fileId string, fileUniqueId string, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *File) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *File) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *File) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *File) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *File) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *File) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetFileSize

`func (o *File) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *File) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *File) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *File) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFilePath

`func (o *File) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *File) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *File) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *File) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


